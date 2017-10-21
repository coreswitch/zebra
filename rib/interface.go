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
	"regexp"
	"sync"
	"syscall"
	"time"

	"github.com/coreswitch/netutil"
)

type IfIndex uint32

type Interface struct {
	Name        string
	Index       IfIndex
	VrfIndex    int
	Vrf         *Vrf
	IfType      uint8
	Mtu         uint32
	DefaultMtu  uint32
	Metric      int
	Flags       uint32
	Stats       IfStats
	Addrs       [AFI_MAX]IfAddrSlice
	HwAddr      net.HardwareAddr
	Description string
	Master      int
	VIFs        []*VIF
}

type IfInfo struct {
	MsgType uint16
	Name    string
	Index   uint32
	Mtu     uint32
	Metric  int
	Flags   uint32
	IfType  uint8
	HwAddr  net.HardwareAddr
	Table   int
	Master  int
	Boot    bool
}

const (
	IF_WATCH_REGISTER = iota
	IF_WATCH_UNREGISTER
)

type IfWatcher struct {
	Type   int
	IfName string
	Chan   chan error
	Timer  *time.Timer
}

func (ifi *IfInfo) String() string {
	return fmt.Sprintf("%s index:%d mtu:%d metric:%d table:%d master:%d", ifi.Name, ifi.Index, ifi.Mtu, ifi.Metric, ifi.Table, ifi.Master)
}

var (
	IfTable = netutil.NewPtree(32)
	IfMap   = map[string]*Interface{}
	IfMutex sync.Mutex
)

var (
	HardwareAddrSetHook func(int, []byte)
)

func (ifp *Interface) IsUp() bool {
	return (ifp.Flags & syscall.IFF_UP) != 0
}

func (ifp *Interface) IsRunning() bool {
	return (ifp.Flags & syscall.IFF_RUNNING) != 0
}

func (ifp *Interface) IsLoopback() bool {
	return (ifp.Flags & syscall.IFF_LOOPBACK) != 0
}

func (ifp *Interface) IsBroadcast() bool {
	return (ifp.Flags & syscall.IFF_BROADCAST) != 0
}

func IfNameByIndex(index IfIndex) string {
	ifp := IfLookupByIndex(index)
	if ifp == nil {
		return ""
	}
	return ifp.Name
}

func IfLookupByName(name string) *Interface {
	return IfMap[name]
}

func IfLookupByIndex(index IfIndex) *Interface {
	n := IfTable.LookupByUint32(uint32(index))
	if n == nil {
		return nil
	}
	return n.Item.(*Interface)
}

func IfRegister(ifp *Interface) {
	IfMutex.Lock()
	defer IfMutex.Unlock()
	IfMap[ifp.Name] = ifp
	n := IfTable.AcquireByUint32(uint32(ifp.Index))
	n.Item = ifp
}

func IfUnregister(ifp *Interface) {
	IfMutex.Lock()
	defer IfMutex.Unlock()
	delete(IfMap, ifp.Name)
	IfTable.ReleaseByUint32(uint32(ifp.Index))
}

func (v *Vrf) IfLookupByName(name string) *Interface {
	return v.IfMap[name]
}

func (v *Vrf) IfLookupByIndex(index IfIndex) *Interface {
	n := v.IfTable.LookupByUint32(uint32(index))
	if n == nil {
		return nil
	}
	return n.Item.(*Interface)
}

func (v *Vrf) IfRegister(ifp *Interface) {
	v.IfMutex.Lock()
	defer v.IfMutex.Unlock()
	v.IfMap[ifp.Name] = ifp
	n := v.IfTable.AcquireByUint32(uint32(ifp.Index))
	n.Item = ifp

	for w, _ := range v.Watcher {
		if w.Type == IF_WATCH_REGISTER && w.IfName == ifp.Name {
			WatcherDone(v, w)
		}
	}
}

func (v *Vrf) IfUnregister(ifp *Interface) {
	v.IfMutex.Lock()
	defer v.IfMutex.Unlock()
	delete(v.IfMap, ifp.Name)
	v.IfTable.ReleaseByUint32(uint32(ifp.Index))

	for w, _ := range v.Watcher {
		if w.Type == IF_WATCH_UNREGISTER && w.IfName == ifp.Name {
			fmt.Println("IfUnregister watch", w.IfName)
			WatcherDone(v, w)
		}
	}
}

func WatcherDone(v *Vrf, w *IfWatcher) {
	v.WMutex.Lock()
	defer v.WMutex.Unlock()
	w.Chan <- nil
	if w.Timer != nil {
		w.Timer.Stop()
	}
	delete(v.Watcher, w)
}

func (v *Vrf) IfWatchAdd(typ int, ifName string, errCh chan error) *IfWatcher {
	v.WMutex.Lock()
	defer v.WMutex.Unlock()

	w := &IfWatcher{Type: typ, IfName: ifName, Chan: errCh}
	w.Timer = time.AfterFunc(time.Second*5,
		func() {
			fmt.Println("[API] IfWatch time out!")
			w.Timer = nil
			WatcherDone(v, w)
		})
	v.Watcher[w] = true

	return w
}

var IfAddHook func(ifp *Interface)

func EsiIfAddHook(ifp *Interface) {
	fmt.Println("EsiIfAddHook", ifp.Name)
	regex, err := regexp.Compile("sproute\\d+")
	if err != nil {
		fmt.Println("Regex compile error:", err)
		return
	}

	if regex.MatchString(ifp.Name) && ifp.Name != "sproute0" {
		fmt.Println("EsiIfAddHook going to reflect sproute0 address to", ifp.Name)
		sifp := IfLookupByName("sproute0")
		if sifp == nil {
			fmt.Println("EsiIfAddHook sproute0 does not exist")
			return
		}
		if len(sifp.Addrs[AFI_IP]) == 0 {
			fmt.Println("EsiIfAddHook sproute0 address is empty")
			return
		}
		addr := sifp.Addrs[AFI_IP][0].Prefix
		fmt.Println("EsiIfAddHook ok reflecting address", addr)
		EsiIfAddrAddHook(sifp, addr)
	}
}

func IfAdd(ifi *IfInfo) {
	vrf := VrfLookupByIndex(ifi.Table)
	if vrf == nil {
		fmt.Println("IfAdd: VrfLookupByIndex", ifi.Table, "failed")
		vrf, _ = VrfAssign(fmt.Sprintf("vrf%d", ifi.Table), ifi.Table)
	}

	if ifi.Master != 0 {
		master := IfLookupByIndex(IfIndex(ifi.Master))
		if master == nil {
			fmt.Println("IfAdd: Can't find master interface by index:", ifi.Master)
			ifi.Master = 0
		} else {
			masterVrf := VrfLookupByIndex(master.VrfIndex)
			if masterVrf == nil {
				ifi.Master = 0
			} else {
				vrf = masterVrf
				ifi.Table = master.VrfIndex
			}
		}
	}

	ifp := &Interface{
		IfType:     ifi.IfType,
		Name:       ifi.Name,
		Index:      IfIndex(ifi.Index),
		Flags:      ifi.Flags,
		Mtu:        ifi.Mtu,
		DefaultMtu: ifi.Mtu,
		Metric:     1,
		HwAddr:     ifi.HwAddr,
		VrfIndex:   ifi.Table,
		Vrf:        vrf,
		Master:     ifi.Master,
	}

	IfRegister(ifp)
	vrf.IfRegister(ifp)

	ifc := InterfaceConfigGet(ifp.Name)
	ifc.Interface = ifp

	if !ifi.Boot {
		InterfaceConfigSync(ifp, ifc)
	}

	ConfigPush([]string{"interfaces", "interface", ifp.Name})

	if IfStatusChangeHook != nil {
		IfStatusChangeHook(ifp.Name, ifp.IsUp(), ifp.IsRunning())
	}

	if IfAddHook != nil {
		IfAddHook(ifp)
	}

	IfForceUp(ifp.Name)

	RibWalker()
}

func (v *Vrf) IfName(index IfIndex) string {
	ifp := v.IfLookupByIndex(index)
	if ifp == nil {
		return "unknown"
	} else {
		return ifp.Name
	}
}

func IfDownRibRemove(ifp *Interface) {
	for afi := AFI_IP; afi < AFI_MAX; afi++ {
		for _, addr := range ifp.Addrs[afi] {
			p := addr.Prefix.Copy()
			ri := &Rib{Type: RIB_CONNECTED, Nexthop: NewNexthopIf(ifp.Index), IfAddr: addr}
			ifp.Vrf.RibDelete(p, ri)
			ifp.Vrf.RouterIdDelete(ifp, addr)
		}
	}
}

func IfUpRibAdd(ifp *Interface) {
	for afi := AFI_IP; afi < AFI_MAX; afi++ {
		for _, addr := range ifp.Addrs[afi] {
			p := addr.Prefix.Copy()
			ri := &Rib{Type: RIB_CONNECTED, Nexthop: NewNexthopIf(ifp.Index), IfAddr: addr}
			ifp.Vrf.RibAdd(p, ri)
			ifp.Vrf.RouterIdAdd(ifp, addr)
		}
	}
}

var IfStatusChangeHook func(string, bool, bool)
var IfForceUpFlag bool

func IfForceUp(ifName string) {
	time.AfterFunc(time.Second*3, func() {
		server.IfUp(ifName)
	})
}

func IfSync(ifp *Interface, ifi *IfInfo) {
	if ifp.Master != ifi.Master {
		// Figure out new VRF from master.
		var nvrf *Vrf
		if ifi.Master != 0 {
			master := IfLookupByIndex(IfIndex(ifi.Master))
			if master == nil {
				fmt.Println("Can't find master interface")
				return
			}
			nvrf = VrfLookupByIndex(master.VrfIndex)
			if nvrf == nil {
				fmt.Println("Can't find vrf")
				return
			}
		} else {
			nvrf = VrfLookupByIndex(0)
		}

		IfDownRibRemove(ifp)
		// IfDown(ifp)

		// Unregister from current VRF.
		ovrf := VrfLookupByIndex(ifp.VrfIndex)
		if ovrf != nil {
			ovrf.IfUnregister(ifp)
		}

		// Register to new VRF.
		nvrf.IfRegister(ifp)

		// Update interface master.
		ifp.Master = ifi.Master
		ifp.VrfIndex = nvrf.Index
		ifp.Vrf = nvrf

		if ifp.IsUp() {
			IfUpRibAdd(ifp)
		}
	}

	// Handle interface name change.
	if ifi.Name != ifp.Name {
		vrf := VrfLookupByIndex(ifp.VrfIndex)
		if vrf != nil {
			IfUnregister(ifp)
			vrf.IfUnregister(ifp)

			ifp.Name = ifi.Name

			IfRegister(ifp)
			vrf.IfRegister(ifp)
		}
	}

	if ifp.Flags != ifi.Flags {
		if ifp.IsUp() {
			if (ifi.Flags & syscall.IFF_UP) == 0 {
				fmt.Println("Interface status is Up -> Down", ifp.Name)
				IfDownRibRemove(ifp)
				if IfForceUpFlag {
					IfForceUp(ifp.Name)
				}
			}
		} else {
			if (ifi.Flags & syscall.IFF_UP) != 0 {
				fmt.Println("Interface status is Down -> Up", ifp.Name)
				ifc := InterfaceConfigLookup(ifp.Name)
				if ifc != nil {
					fmt.Println("Re-install interface address", ifp.Name)
					for afi := AFI_IP; afi < AFI_MAX; afi++ {
						for _, addr := range ifc.Addrs[afi] {
							fmt.Println("Re-instal addr:", addr.Prefix)
							NetlinkIpAddrAdd(ifp, addr.Prefix)
						}
					}
				}

				for afi := AFI_IP; afi < AFI_MAX; afi++ {
					for _, addr := range ifp.Addrs[afi] {
						p := addr.Prefix.Copy()
						ri := &Rib{Type: RIB_CONNECTED, Nexthop: NewNexthopIf(ifp.Index), IfAddr: addr}
						ifp.Vrf.RibAdd(p, ri)
						ifp.Vrf.RouterIdAdd(ifp, addr)
					}
				}
			}
		}
		ifp.Flags = ifi.Flags

		// Need to update status
		if IfStatusChangeHook != nil {
			IfStatusChangeHook(ifp.Name, ifp.IsUp(), ifp.IsRunning())
		}

		RibWalker()
	}

	if ifi.Mtu != 0 {
		if ifp.Mtu != ifi.Mtu {
			ifp.Mtu = ifi.Mtu
		}
	}
}

func IfUpdate(ifi *IfInfo) {
	ifp := IfLookupByIndex(IfIndex(ifi.Index))
	if ifp == nil {
		IfAdd(ifi)
	} else {
		IfSync(ifp, ifi)
	}
}

func IfDelete(ifi *IfInfo) {
	//fmt.Println("IfDelete: ", ifi.Name, "Master: ", ifi.Master)
	ifp := IfLookupByIndex(IfIndex(ifi.Index))
	if ifp == nil {
		//fmt.Println("Interface already removed from cache: ", ifi.Name)
		return
	}

	// Withdraw connected route.
	IfDownRibRemove(ifp)

	// Bring down the interface.  This may invoke other routes withdraw such as staic route.

	// Remove from VRF table.
	vrf := ifp.Vrf
	vrf.IfUnregister(ifp)
	IfUnregister(ifp)

	ConfigPull([]string{"interfaces", "interface", ifi.Name})
}

func InterfaceIterate(f func(*Interface)) {
	for n := IfTable.Top(); n != nil; n = IfTable.Next(n) {
		ifp := n.Item.(*Interface)
		f(ifp)
	}
}

func (v *Vrf) InterfaceIterate(f func(*Interface)) {
	for n := v.IfTable.Top(); n != nil; n = v.IfTable.Next(n) {
		ifp := n.Item.(*Interface)
		f(ifp)
	}
}

func InterfaceShowBrief(afi int) (line string) {
	if afi != netutil.AFI_IP && afi != netutil.AFI_IP6 {
		return
	}
	if afi == netutil.AFI_IP {
		line = "Interface             IP-Address      Status                Protocol\n"
	}
	InterfaceIterate(func(ifp *Interface) {
		status := ""

		// XXX "administratively down" treatment is needed.
		if ifp.IsUp() {
			status = "up"
		} else {
			status = "down"
		}
		protocol := ""
		if ifp.IsRunning() {
			protocol = "up"
		} else {
			protocol = "down"
		}

		if afi == netutil.AFI_IP {
			addrStr := ""
			if len(ifp.Addrs[afi]) == 0 {
				addrStr = "unasssigned"
			} else {
				for pos, addr := range ifp.Addrs[afi] {
					addrStr = addr.Prefix.IP.String()
					if pos == 0 {
						line += fmt.Sprintf("%-22s%-16s%-22s%s\n", ifp.Name, addrStr, status, protocol)
					} else {
						line += fmt.Sprintf("                      %-16s\n", addrStr)
					}
				}
			}
		} else {
			line += fmt.Sprintf("%-26s[%s/%s]\n", ifp.Name, status, protocol)
			if len(ifp.Addrs[afi]) == 0 {
				line += fmt.Sprintf("    unassigned\n")
			} else {
				for _, addr := range ifp.Addrs[afi] {
					line += fmt.Sprintf("    %s\n", addr.Prefix.IP.String())
				}
			}
		}
	})

	return
}

func InterfaceConfigSync(ifp *Interface, ifc *InterfaceConfig) {
	ifc.ShutdownSync(ifp)

	if ifc.HwAddr != nil {
		HardwareAddrSet(ifp, ifc.HwAddr)
		if HardwareAddrSetHook != nil {
			HardwareAddrSetHook(ifc.L3Index, ifc.HwAddr)
		}
	}
	for _, ifa := range ifc.Addrs[AFI_IP] {
		NetlinkIpAddrAdd(ifp, ifa.Prefix)
	}
	for _, ifa := range ifc.Addrs[AFI_IP6] {
		NetlinkIpAddrAdd(ifp, ifa.Prefix)
	}
}

func InterfaceSyncWithConfig() {
	InterfaceIterate(func(ifp *Interface) {
		ifc := InterfaceConfigGet(ifp.Name)
		ifc.ShutdownSync(ifp)
	})
}

func IfStatusNotifyEtcd(ifname string, up bool, running bool) {
	EtcdSetIfStatus(ifname, up, running)
}
