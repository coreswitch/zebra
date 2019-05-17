// Copyright 2016, 2017 Zebra Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rib

import (
	"fmt"
	"net"
	"syscall"

	"github.com/coreswitch/component"
	"github.com/coreswitch/netutil"
	"github.com/coreswitch/zebra/fea"
	"github.com/coreswitch/zebra/policy"
)

type Fn struct {
	fn  func() error
	err chan error
}

type Server struct {
	ifChan     chan IfInfo
	ifaddrChan chan IfAddrInfo
	routeChan  chan RouteInfo
	apiSyncCh  chan *Fn
	apiAsyncCh chan *Fn
	sync       chan *Fn
	async      chan *Fn
	pm         *policy.PrefixListMaster
	pmInternal *policy.PrefixListMaster
}

var (
	server        *Server
	NewServerHook func()
)

func NewServer() *Server {
	inst := &Server{
		ifChan:     make(chan IfInfo, 1024),
		ifaddrChan: make(chan IfAddrInfo, 1024),
		routeChan:  make(chan RouteInfo, 1024),
		apiSyncCh:  make(chan *Fn, 1024),
		apiAsyncCh: make(chan *Fn, 1024),
		sync:       make(chan *Fn, 1024),
		async:      make(chan *Fn, 1024),
		pm:         policy.NewPrefixListMaster(),
		pmInternal: policy.NewPrefixListMaster(),
	}
	server = inst
	if NewServerHook != nil {
		NewServerHook()
	}
	return inst
}

// API channel send.
func (s *Server) apiSync(fn func() error) error {
	err := make(chan error)
	s.apiSyncCh <- &Fn{fn: fn, err: err}
	return <-err
}

// API channel async send.
func (s *Server) apiAsync(fn func() error) error {
	err := make(chan error)
	s.apiAsyncCh <- &Fn{fn: fn, err: err}
	return <-err
}

func (s *Server) syncCall(fn func() error) error {
	err := make(chan error)
	s.sync <- &Fn{fn: fn, err: err}
	return <-err
}

func (s *Server) asyncCall(fn func() error, errCh chan error) {
	s.async <- &Fn{fn: fn, err: errCh}
}

func (s *Server) Serv() {
	go func() {
		for {
			select {
			case sync := <-s.sync:
				sync.err <- sync.fn()
			case async := <-s.async:
				err := async.fn()
				if err != nil {
					async.err <- err
				}
			case ifi := <-s.ifChan:
				if ifi.MsgType == syscall.RTM_NEWLINK {
					fmt.Println("If add:", ifi)
					IfUpdate(&ifi)
				} else {
					fmt.Println("If del:", ifi)
					IfDelete(&ifi)
				}
			case ifaddr := <-s.ifaddrChan:
				if ifaddr.MsgType == syscall.RTM_NEWADDR {
					fmt.Println("Addr add:", ifaddr)
					IfAddrAdd(&ifaddr)
				} else {
					fmt.Println("Addr del:", ifaddr)
					IfAddrDelete(&ifaddr)
				}
			case route := <-s.routeChan:
				if route.MsgType == syscall.RTM_NEWROUTE {
					fmt.Println("Route add ", route)
					RibAdd(uint32(route.Table), route.Prefix, &route.Rib)
				} else {
					fmt.Println("Route del ", route)
					RibDelete(uint32(route.Table), route.Prefix, &route.Rib)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case apiSync := <-s.apiSyncCh:
				apiSync.err <- s.syncCall(apiSync.fn)
			case apiAsync := <-s.apiAsyncCh:
				apiAsync.err <- apiAsync.fn()
			}
		}
	}()
}

func WaitAsync(errCh chan error) error {
	err := <-errCh
	fmt.Println("[API] WaitAsync is called")
	return err
}

func (s *Server) vrfAdd(vrfName string, errCh chan error) error {
	fmt.Println("[API] VrfAdd start:", vrfName)
	VrfMutex.Lock()
	defer VrfMutex.Unlock()

	if v := VrfLookupByName(vrfName); v != nil {
		fmt.Println("VrfAdd: vrf already exists", vrfName)
		return fmt.Errorf("VRF already exists")
	}
	index := VrfExtractIndex(vrfName)
	v, err := VrfAssign(vrfName, index)
	if err != nil {
		fmt.Println("VrfAdd: Can not assign vrf index", index)
		return err
	}

	v.IfWatchAdd(IF_WATCH_REGISTER, vrfName, errCh)

	fea.VrfAdd(v.Name, v.Id)
	NetlinkVrfAdd(v.Name, v.Id)

	return nil
}

func (s *Server) vrfAddAsync(vrfName string) error {
	errCh := make(chan error)
	s.asyncCall(func() error {
		return s.vrfAdd(vrfName, errCh)
	}, errCh)
	err := WaitAsync(errCh)
	fmt.Println("[API] VrfAdd end:", vrfName)
	return err
}

func (s *Server) VrfAdd(vrfName string) error {
	return s.apiAsync(func() error {
		return s.vrfAddAsync(vrfName)
	})
}

func (s *Server) vrfDelete(vrfName string, errCh chan error) error {
	fmt.Println("[API] VrfDelete start:", vrfName)
	VrfMutex.Lock()
	defer VrfMutex.Unlock()

	v := VrfLookupByName(vrfName)
	if v == nil {
		fmt.Println("VrfDelete: vrf does not exists", vrfName)
		return fmt.Errorf("VRF does not exists")
	}

	v.IfWatchAdd(IF_WATCH_UNREGISTER, vrfName, errCh)

	//fea.VrfDelete(v.Name, v.Index)
	NetlinkVrfDelete(v.Name, v.Id)

	delete(VrfMap, v.Name)
	VrfTable[v.Id] = nil
	//GobgpVrfDelete(v.Id)
	ZServerStop(v.ZServer)

	return nil
}

func (s *Server) vrfDeleteAsync(vrfName string) error {
	errCh := make(chan error)
	s.asyncCall(func() error {
		return s.vrfDelete(vrfName, errCh)
	}, errCh)
	err := WaitAsync(errCh)
	fmt.Println("[API] VRFDelete end:", vrfName)

	return err
}

var NoVRFDelete bool

func (s *Server) VrfDelete(vrfName string) error {
	if NoVRFDelete {
		fmt.Println("[API] VRFDelete NoVRFDelete")
		return nil
	}
	return s.apiAsync(func() error {
		return s.vrfDeleteAsync(vrfName)
	})
}

// API.
func (s *Server) VIFAdd(ifName string, vlanId uint64) error {
	return s.apiAsync(func() error {
		return s.vifAddAsync(ifName, vlanId)
	})
}

func (s *Server) vifAddAsync(ifName string, vlanId uint64) error {
	errCh := make(chan error)
	s.asyncCall(func() error {
		return s.vifAdd(ifName, vlanId, errCh)
	}, errCh)
	err := WaitAsync(errCh)
	fmt.Println("[API] VIFAdd end:", ifName, vlanId)
	return err
}

func (s *Server) vifAdd(ifName string, vlanId uint64, errCh chan error) error {
	fmt.Println("[API] VIFAdd start:", ifName, vlanId)

	ifp := IfLookupByName(ifName)
	if ifp == nil {
		return fmt.Errorf("Interface %s can't find", ifName)
	}
	vif := ifp.VIFLookup(vlanId)
	if vif != nil {
		return fmt.Errorf("VIF for vlan ID %d already exists", vlanId)
	}

	vlanIfName := fmt.Sprintf("%s.%d", ifp.Name, vlanId)

	v := VrfLookupByName("")
	v.IfWatchAdd(IF_WATCH_REGISTER, vlanIfName, errCh)

	NetlinkVlanAdd(vlanIfName, int(vlanId), int(ifp.Index))

	vif = NewVIF(vlanId)
	ifp.VIFs = append(ifp.VIFs, vif)

	return nil
}

func (s *Server) VIFDelete(ifName string, vlanId uint64) error {
	return s.apiSync(func() error {
		return s.vifDelete(ifName, vlanId)
	})
}

func (s *Server) vifDelete(ifName string, vlanId uint64) error {
	fmt.Println("[API] VIFDelete start:", ifName, vlanId)

	ifp := IfLookupByName(ifName)
	if ifp == nil {
		return fmt.Errorf("Interace %s can't find:", ifName)
	}
	vif := ifp.UnregisterVIF(vlanId)
	if vif == nil {
		return fmt.Errorf("VIF for vlan ID %d does not exists", vlanId)
	}

	vlanIfName := fmt.Sprintf("%s.%d", ifp.Name, vlanId)

	NetlinkVlanDelete(vlanIfName, int(vlanId))

	fmt.Println("[API] VIFDelete end:", ifName, vlanId)

	return nil
}

func (s *Server) IfVrfBind(ifName string, vrfName string) error {
	retry := 5
	err := s.apiAsync(func() error {
		return s.ifVrfBindAsync(ifName, vrfName)
	})
	for err != nil && retry != 0 {
		retry--
		fmt.Printf("IfVrfBind failed, retry count %d\n", retry)
		err = s.apiAsync(func() error {
			return s.ifVrfBindAsync(ifName, vrfName)
		})
		if err == nil {
			return err
		}
	}
	return err
}

func (s *Server) ifVrfBindAsync(ifName string, vrfName string) error {
	errCh := make(chan error)
	s.asyncCall(func() error {
		return s.ifVrfBind(ifName, vrfName, errCh)
	}, errCh)
	err := WaitAsync(errCh)
	fmt.Println("[API] IfVrfBind end:", ifName, vrfName)
	return err
}

func (s *Server) ifVrfBind(ifName string, vrfName string, errCh chan error) error {
	fmt.Println("[API] IfVrfBind start:", ifName, vrfName)

	ifp := IfLookupByName(ifName)
	if ifp == nil {
		return fmt.Errorf("IfVrfBind: Can't find interface %s", ifName)
	}
	vrfIf := IfLookupByName(vrfName)
	if vrfIf == nil {
		return fmt.Errorf("IfVrfBind: Can't find vrfIf %s", vrfName)
	}
	vrf := VrfLookupByName(vrfName)
	if vrf == nil {
		return fmt.Errorf("IfVrfBind: Can't find vrf %s", vrfName)
	}

	vrf.IfWatchAdd(IF_WATCH_REGISTER, ifName, errCh)

	err := NetlinkVrfBindInterface(ifp.Name, ifp.Index, vrfIf.Index)
	fmt.Println("IfVrfBind: NetlinkVrfBindInterface err", err)
	if err != nil {
		return fmt.Errorf("IfVrfBind: NetlinkVrfBindInterface err %s", err)
	}

	return nil
}

func (s *Server) IfVrfUnbind(ifName string, vrfName string) error {
	return s.apiAsync(func() error {
		return s.ifVrfUnbindAsync(ifName, vrfName)
	})
}

func (s *Server) ifVrfUnbindAsync(ifName string, vrfName string) error {
	errCh := make(chan error)
	s.asyncCall(func() error {
		return s.ifVrfUnbind(ifName, vrfName, errCh)
	}, errCh)
	err := WaitAsync(errCh)
	fmt.Println("[API] IfVrfUnbind end:", ifName, vrfName)
	return err
}

func (s *Server) ifVrfUnbind(ifName string, vrfName string, errCh chan error) error {
	fmt.Println("[API] IfVrfUnbind start:", ifName, vrfName)

	vrf := VrfLookupByName(vrfName)
	if vrf == nil {
		return fmt.Errorf("IfVrfUnbind: Can't find vrf %s", vrfName)
	}

	ifp := vrf.IfLookupByName(ifName)
	if ifp == nil {
		fmt.Println("[API] IfVrfUnbind: Can't find interface", ifName, "in vrf", vrfName)
		return fmt.Errorf("IfVrfUnbind: Can't find interface by name: %s", ifName)
	}

	vrf.IfWatchAdd(IF_WATCH_UNREGISTER, ifName, errCh)

	NetlinkVrfUnbindInterface(ifp.Name, ifp.Index)

	return nil
}

func (s *Server) StaticAdd(prefix *netutil.Prefix, nexthop net.IP) error {
	return s.apiSync(func() error {
		return VrfDefault().StaticAdd(prefix, nexthop)
	})
}

func (s *Server) StaticDelete(prefix *netutil.Prefix, nexthop net.IP) error {
	return s.apiSync(func() error {
		return VrfDefault().StaticDelete(prefix, nexthop)
	})
}

func (s *Server) StaticSeg6SegmentsAdd(prefix *netutil.Prefix, nexthop net.IP, mode string, segs []net.IP) error {
	return s.apiSync(func() error {
		return VrfDefault().StaticSeg6SegmentsAdd(prefix, nexthop, mode, segs)
	})
}

func (s *Server) StaticSeg6SegmentsDelete(prefix *netutil.Prefix, nexthop net.IP, mode string, segs []net.IP) error {
	return s.apiSync(func() error {
		return VrfDefault().StaticSeg6SegmentsDelete(prefix, nexthop, mode, segs)
	})
}

func (s *Server) StaticSeg6LocalAdd(prefix *netutil.Prefix, nexthop net.IP, seg6local EncapSEG6Local) error {
	return s.apiSync(func() error {
		return VrfDefault().StaticSeg6LocalAdd(prefix, nexthop, seg6local)
	})
}

func (s *Server) StaticSeg6LocalDelete(prefix *netutil.Prefix, nexthop net.IP, seg6local EncapSEG6Local) error {
	return s.apiSync(func() error {
		return VrfDefault().StaticSeg6LocalDelete(prefix, nexthop, seg6local)
	})
}

func (s *Server) AddrAdd(ifName string, addr *netutil.Prefix) error {
	return s.apiSync(func() error {
		ifc := InterfaceConfigGet(ifName)
		return ifc.AddrAdd(ifName, addr)
	})
}

func (s *Server) AddrDelete(ifName string, addr *netutil.Prefix) error {
	return s.apiSync(func() error {
		ifc := InterfaceConfigGet(ifName)
		return ifc.AddrDelete(ifName, addr)
	})
}

// func (s *Server) IfUp(ifName string) error {
// 	return s.apiSync(func() error {
// 		fmt.Println("[API] IfUp start:", ifName)
// 		ifp := IfLookupByName(ifName)
// 		if ifp != nil {
// 			err := LinkSetUp(ifp)
// 			fmt.Println("[API] IfUp end:", ifName, err)
// 			return err
// 		}
// 		return nil
// 	})
// }

func (s *Server) InterfaceSubscribe(w Watcher, vrfId uint32) error {
	return s.apiSync(func() error {
		vrf := VrfLookupByIndex(vrfId)
		if vrf == nil {
			return fmt.Errorf("Can't find VRF by VRF ID: %d", vrfId)
		}
		t := WATCH_TYPE_INTERFACE
		for _, v := range vrf.Watchers[t] {
			if w == v {
				return nil
			}
		}
		vrf.Watchers[t] = append(vrf.Watchers[t], w)
		NotifyInterfaces(w, vrf)
		return nil
	})
}

func (s *Server) InterfaceUnsubscribe(w Watcher, vrfId uint32) error {
	return s.apiSync(func() error {
		vrf := VrfLookupByIndex(vrfId)
		if vrf == nil {
			return fmt.Errorf("Can't find VRF by VRF ID: %d", vrfId)
		}
		t := WATCH_TYPE_INTERFACE
		for i, v := range vrf.Watchers[t] {
			if w == v {
				vrf.Watchers[t] = append(vrf.Watchers[t][:i], vrf.Watchers[t][i+1:]...)
			}
		}
		return nil
	})
}

func (s *Server) RouterIdSubscribe(w Watcher, vrfId uint32) error {
	return s.apiSync(func() error {
		vrf := VrfLookupByIndex(vrfId)
		if vrf == nil {
			return fmt.Errorf("Can't find VRF by VRF ID: %d", vrfId)
		}
		t := WATCH_TYPE_ROUTER_ID
		for _, v := range vrf.Watchers[t] {
			if w == v {
				return nil
			}
		}
		vrf.Watchers[t] = append(vrf.Watchers[t], w)
		NotifyRouterId(w, vrf)
		return nil
	})
}

func (s *Server) RouterIdUnsubscribe(w Watcher, vrfId uint32) error {
	return s.apiSync(func() error {
		vrf := VrfLookupByIndex(vrfId)
		if vrf == nil {
			return fmt.Errorf("Can't find VRF by VRF ID: %d", vrfId)
		}
		t := WATCH_TYPE_ROUTER_ID
		for i, v := range vrf.Watchers[t] {
			if w == v {
				vrf.Watchers[t] = append(vrf.Watchers[t][:i], vrf.Watchers[t][i+1:]...)
			}
		}
		return nil
	})
}

func RedistWatcherAdd(watchers Watchers, w Watcher) Watchers {
	for _, v := range watchers {
		if w == v {
			return watchers
		}
	}
	watchers = append(watchers, w)
	return watchers
}

func RedistWatcherRemove(watchers Watchers, w Watcher) Watchers {
	for i, v := range watchers {
		if w == v {
			watchers = append(watchers[:i], watchers[i+1:]...)
		}
	}
	return watchers
}

func EsiNewServerHook() {
	p1, _ := netutil.ParsePrefix("172.0.0.0/8")
	p2, _ := netutil.ParsePrefix("198.0.0.0/8")
	any, _ := netutil.ParsePrefix("0.0.0.0/0")
	server.PrefixListOutAdd(policy.NewPrefixListEntry(5, policy.Deny, p1, policy.WithEq(12)))
	server.PrefixListOutAdd(policy.NewPrefixListEntry(10, policy.Deny, p2, policy.WithEq(15)))
	server.PrefixListOutAdd(policy.NewPrefixListEntry(15, policy.Permit, any, policy.WithLe(32)))
}

func (s *Server) PrefixListOutAdd(entry *policy.PrefixListEntry) {
	s.pmInternal.EntryAdd("*redist-out*", entry)
}

func (s *Server) PrefixListOut() *policy.PrefixList {
	return s.pmInternal.Lookup("*redist-out*")
}

func (s *Server) RedistSubscribe(w Watcher, allVrf bool, vrfId uint32, afi int, typ uint8) error {
	return s.apiSync(func() error {
		if allVrf {
			Redist[afi].typ[typ] = RedistWatcherAdd(Redist[afi].typ[typ], w)
			for _, vrf := range VrfMap {
				vrf.RedistSync(w, afi, typ)
			}
		} else {
			vrf := VrfLookupByIndex(vrfId)
			if vrf == nil {
				return fmt.Errorf("Can't find VRF by VRF ID: %d", vrfId)
			}
			vrf.redist[afi].typ[typ] = RedistWatcherAdd(vrf.redist[afi].typ[typ], w)
			vrf.RedistSync(w, afi, typ)
		}
		return nil
	})
}

func (s *Server) RedistUnsubscribe(w Watcher, allVrf bool, vrfId uint32, afi int, typ uint8) error {
	return s.apiSync(func() error {
		if allVrf {
			Redist[afi].typ[typ] = RedistWatcherRemove(Redist[afi].typ[typ], w)
		} else {
			vrf := VrfLookupByIndex(vrfId)
			if vrf == nil {
				return fmt.Errorf("Can't find VRF by VRF ID: %d", vrfId)
			}
			vrf.redist[afi].typ[typ] = RedistWatcherRemove(vrf.redist[afi].typ[typ], w)
		}
		return nil
	})
}

func (s *Server) RedistDefaultSubscribe(w Watcher, allVrf bool, vrfId uint32, afi int) error {
	return s.apiSync(func() error {
		if allVrf {
			Redist[afi].def = RedistWatcherAdd(Redist[afi].def, w)
			for _, vrf := range VrfMap {
				vrf.RedistDefaultSync(w, afi)
			}
		} else {
			vrf := VrfLookupByIndex(vrfId)
			if vrf == nil {
				return fmt.Errorf("Can't find VRF by VRF ID: %d", vrfId)
			}
			vrf.redist[afi].def = RedistWatcherAdd(vrf.redist[afi].def, w)
			vrf.RedistDefaultSync(w, afi)
		}
		return nil
	})
}

func (s *Server) RedistDefaultUnsubscribe(w Watcher, allVrf bool, vrfId uint32, afi int) error {
	return s.apiSync(func() error {
		if allVrf {
			Redist[afi].def = RedistWatcherRemove(Redist[afi].def, w)
		} else {
			vrf := VrfLookupByIndex(vrfId)
			if vrf == nil {
				return fmt.Errorf("Can't find VRF by VRF ID: %d", vrfId)
			}
			vrf.redist[afi].def = RedistWatcherRemove(vrf.redist[afi].def, w)
		}
		return nil
	})
}

func (s *Server) WatcherUnsubscribe(w Watcher) error {
	return s.apiSync(func() error {
		for afi := AFI_IP; afi < AFI_MAX; afi++ {
			Redist[afi].def = RedistWatcherRemove(Redist[afi].def, w)
			for t := RIB_UNKNOWN; t < RIB_MAX; t++ {
				Redist[afi].typ[t] = RedistWatcherRemove(Redist[afi].typ[t], w)
			}
		}
		for _, vrf := range VrfMap {
			for t, _ := range vrf.Watchers {
				for i, v := range vrf.Watchers[t] {
					if w == v {
						vrf.Watchers[t] = append(vrf.Watchers[t][:i], vrf.Watchers[t][i+1:]...)
					}
				}
			}
			for afi := AFI_IP; afi < AFI_MAX; afi++ {
				vrf.redist[afi].def = RedistWatcherRemove(vrf.redist[afi].def, w)
				for t := RIB_UNKNOWN; t < RIB_MAX; t++ {
					vrf.redist[afi].typ[t] = RedistWatcherRemove(vrf.redist[afi].typ[t], w)
				}
			}
		}
		RibClearSrc(w)
		return nil
	})
}

func (s *Server) PrefixListMasterSet(pm *policy.PrefixListMaster) error {
	return s.apiSync(func() error {
		s.pm = pm
		return nil
	})
}

func (s *Server) VrfDistributeListOspfAdd(vrfName string, dlistName string) error {
	return s.apiSync(func() error {
		fmt.Println("[API] VrfDistributeListOspfAdd", vrfName, dlistName)
		vrf := VrfLookupByName(vrfName)
		if vrf == nil {
			return fmt.Errorf("Can't find VRF by VRF name: %s", vrfName)
		}
		vrf.DListOspf = dlistName
		return nil
	})
}

func (s *Server) VrfDistributeListOspfDelete(vrfName string, dlistName string) error {
	return s.apiSync(func() error {
		fmt.Println("[API] VrfDistributeListOspfDelete", vrfName, dlistName)
		vrf := VrfLookupByName(vrfName)
		if vrf == nil {
			return fmt.Errorf("Can't find VRF by VRF name: %s", vrfName)
		}
		vrf.DListOspf = ""
		return nil
	})
}

func (s *Server) PrefixListLookup(plistName string) *policy.PrefixList {
	return s.pm.Lookup(plistName)
}

func (s *Server) Start() component.Component {
	s.Serv()

	NewVrf("", 0)
	NetlinkDumpAndSubscribe(s)

	return s
}

func (s *Server) Stop() component.Component {
	VrfStop()
	VIFClean()
	IfAddrClean()
	return s
}
