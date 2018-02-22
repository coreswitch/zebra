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
	// "time"

	"github.com/coreswitch/component"
	"github.com/coreswitch/netutil"
	"github.com/hash-set/zebra/fea"
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
}

var (
	server *Server
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
	}
	server = inst
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
					//fmt.Println("If delete:", ifi)
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
					RibAdd(route.Table, route.Prefix, &route.Rib)
				} else {
					fmt.Println("Route del ", route)
					RibDelete(route.Table, route.Prefix, &route.Rib)
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

	fea.VrfAdd(v.Name, v.Index)
	NetlinkVrfAdd(v.Name, v.Index)

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
	NetlinkVrfDelete(v.Name, v.Index)

	delete(VrfMap, v.Name)
	VrfTable[v.Index] = nil

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
	return s.apiAsync(func() error {
		return s.vifDeleteAsync(ifName, vlanId)
	})
}

func (s *Server) vifDeleteAsync(ifName string, vlanId uint64) error {
	errCh := make(chan error)
	s.asyncCall(func() error {
		return s.vifDelete(ifName, vlanId, errCh)
	}, errCh)
	err := WaitAsync(errCh)
	fmt.Println("[API] VIFDelete end:", ifName, vlanId)
	return err
}

func (s *Server) vifDelete(ifName string, vlanId uint64, errCh chan error) error {
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

	v := VrfLookupByName("")
	v.IfWatchAdd(IF_WATCH_UNREGISTER, vlanIfName, errCh)

	NetlinkVlanDelete(vlanIfName, int(vlanId))

	return nil
}

func (s *Server) IfVrfBind(ifName string, vrfName string) error {
	return s.apiAsync(func() error {
		return s.ifVrfBindAsync(ifName, vrfName)
	})
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

	NetlinkVrfBindInterface(ifp.Name, ifp.Index, vrfIf.Index)

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

func (s *Server) IfUp(ifName string) error {
	return s.apiSync(func() error {
		ifp := IfLookupByName(ifName)
		if ifp != nil && !ifp.IsUp() {
			return LinkSetUp(ifp)
		}
		return nil
	})
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
