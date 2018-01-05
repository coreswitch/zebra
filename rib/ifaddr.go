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

	"github.com/coreswitch/netutil"
)

type IfAddrFlag uint32

type IfAddrSlice []*IfAddr

type IfAddr struct {
	Flags  IfAddrFlag      `json:"-"`
	Prefix *netutil.Prefix `json:"address"`
	Label  string          `json:"-"`
}

type IfAddrInfo struct {
	Flags   IfAddrFlag
	MsgType uint16
	Prefix  *netutil.Prefix
	Family  int
	Index   IfIndex
	Label   string
	Scope   int
}

func (ifaddr *IfAddrInfo) String() string {
	return fmt.Sprintf("%s %s index:%d", ifaddr.Prefix, ifaddr.Label, ifaddr.Index)
}

const (
	IFADDR_SOURCE_SYSTEM IfAddrFlag = 1 << iota
	IFADDR_SOURCE_CONFIG
)

func (f IfAddrFlag) CheckFlag(flag IfAddrFlag) bool {
	return (f & flag) == flag
}

func (f IfAddrFlag) SetFlag(flag IfAddrFlag) IfAddrFlag {
	f |= flag
	return f
}

func (f IfAddrFlag) SourceSystem() bool {
	return f.CheckFlag(IFADDR_SOURCE_SYSTEM)
}

func (f IfAddrFlag) SourceConfig() bool {
	return f.CheckFlag(IFADDR_SOURCE_CONFIG)
}

func (addrs IfAddrSlice) Lookup(prefix *netutil.Prefix) *IfAddr {
	for _, addr := range addrs {
		if addr.Prefix.Equal(prefix) {
			return addr
		}
	}
	return nil
}

func (addrs IfAddrSlice) Delete(prefix *netutil.Prefix) (IfAddrSlice, *IfAddr) {
	var found *IfAddr
	var newslice IfAddrSlice

	for _, addr := range addrs {
		if addr.Prefix.Equal(prefix) {
			found = addr
		} else {
			newslice = append(newslice, addr)
		}
	}
	return newslice, found
}

var IfAddrAddHook func(ifp *Interface, ifaddr *netutil.Prefix)
var IfAddrDeleteHook func(ifp *Interface, ifaddr *netutil.Prefix)

func EsiIfMask(addr net.IP) int {
	if len(addr) > 0 && addr[0] == 198 {
		return 15
	} else {
		return 12
	}
}

func EsiIfAddrAddHook(ifp *Interface, ifaddr *netutil.Prefix) {
	fmt.Println("EsiIfAddrAddHook", ifp.Name, ifaddr)

	regex, err := regexp.Compile("sproute\\d+")
	if err != nil {
		fmt.Println("Regex compile error:", err)
		return
	}

	if ifp.Name == "sproute0" {
		fmt.Println("Reflecting sproute0 address to other sprouteX interface")
		for _, ifpp := range IfMap {
			if regex.MatchString(ifpp.Name) && ifpp.Name != "sproute0" {
				fmt.Println("Reflecting to interface", ifpp.Name)
				ifc := InterfaceConfigGet(ifpp.Name)
				if ifc == nil {
					fmt.Println("InterfaceConfigGet failed:", ifpp.Name)
					continue
				}
				addr := ifaddr.Copy()
				addr.Length = EsiIfMask(ifaddr.IP)
				ifc.AddrAdd(ifpp.Name, addr)
			}
		}
	}
}

func EsiIfAddrDeleteHook(ifp *Interface, ifaddr *netutil.Prefix) {
	fmt.Println("EsiIfAddrDeleteHook", ifp.Name, ifaddr)

	regex, err := regexp.Compile("sproute\\d+")
	if err != nil {
		fmt.Println("Regex compile error:", err)
		return
	}

	if ifp.Name == "sproute0" {
		fmt.Println("Reflecting sproute0 address to other sprouteX interface")

		for _, ifcp := range InterfaceConfigMap {
			if regex.MatchString(ifcp.Name) && ifcp.Name != "sproute0" {
				fmt.Println("Reflecting to interface", ifcp.Name)
				ifc := InterfaceConfigGet(ifcp.Name)
				if ifc == nil {
					fmt.Println("InterfaceConfigGet failed:", ifcp.Name)
					continue
				}
				addr := ifaddr.Copy()
				addr.Length = EsiIfMask(ifaddr.IP)
				ifc.AddrDelete(ifcp.Name, addr)
			}
		}
	}
}

func IfAddrAdd(ai *IfAddrInfo) {
	// fmt.Println("IfAddrAdd", ai)

	ifp := IfLookupByIndex(ai.Index)
	if ifp == nil {
		return
	}
	afi := ai.Prefix.AFI()
	if afi == AFI_MAX {
		return
	}

	found := ifp.Addrs[afi].Lookup(ai.Prefix)
	if found != nil {
		return
	}
	addr := &IfAddr{Flags: ai.Flags, Prefix: ai.Prefix}
	ifp.Addrs[afi] = append(ifp.Addrs[afi], addr)

	if addr.Flags.SourceSystem() {
		// fmt.Println("Reflect to config")
	}

	if ifp.IsUp() {
		p := ai.Prefix.Copy()
		ri := &Rib{Type: RIB_CONNECTED, Nexthop: NewNexthopIf(ai.Index), IfAddr: addr}
		ifp.Vrf.RibAdd(p, ri)
		ifp.Vrf.RouterIdAdd(ifp, addr)
	}

	IfAddrAddPropagate(ifp, addr)

	if IfAddrAddHook != nil {
		IfAddrAddHook(ifp, ai.Prefix)
	}

	if IfStatusChangeHook != nil {
		IfStatusChangeHook("", false, false)
	}
}

func IfAddrDelete(ai *IfAddrInfo) {
	// fmt.Println("IfAddrDel", ai)

	ifp := IfLookupByIndex(ai.Index)
	if ifp == nil {
		// XXX fmt.Println("Can't find ifp")
		return
	}
	afi := ai.Prefix.AFI()
	if afi == AFI_MAX {
		// XXX fmt.Println("Unknown AFI")
		return
	}

	addrs, found := ifp.Addrs[afi].Delete(ai.Prefix)
	if found == nil {
		// XXX Can't find the address.
		fmt.Println("Can't find address by", ai.Prefix)
		return
	}

	if ifp.IsUp() {
		p := found.Prefix.Copy()
		ri := &Rib{Type: RIB_CONNECTED, Nexthop: NewNexthopIf(ai.Index), IfAddr: found}
		ifp.Vrf.RibDelete(p, ri)
		ifp.Vrf.RouterIdDelete(ifp, found)
	}
	ifp.Addrs[afi] = addrs

	// Recover process if the address is configured.
	if ifp.IsUp() {
		ifc := InterfaceConfigLookup(ifp.Name)
		if ifc != nil {
			addr := ifc.Addrs[afi].Lookup(ai.Prefix)
			if addr != nil {
				fmt.Println("Recover configured address remove from kernel:", addr.Prefix)
				NetlinkIpAddrAdd(ifp, addr.Prefix)
			}
		}
	}

	IfAddrDeletePropagate(ifp, found)

	if IfAddrDeleteHook != nil {
		IfAddrDeleteHook(ifp, ai.Prefix)
	}

	if IfStatusChangeHook != nil {
		IfStatusChangeHook("", false, false)
	}
}

func IfAddrClean() {
	fmt.Printf("IfAddrClean")
	for _, ifc := range InterfaceConfigMap {
		fmt.Println("IfAddrClean ifc:", ifc.Name)
		ifp := IfLookupByName(ifc.Name)
		if ifp != nil {
			addrSlice := make(IfAddrSlice, len(ifc.Addrs[AFI_IP]), len(ifc.Addrs[AFI_IP]))
			copy(addrSlice, ifc.Addrs[AFI_IP])
			ifc.Addrs[AFI_IP] = ifc.Addrs[AFI_IP][:0]

			for _, addr := range addrSlice {
				fmt.Println("IfAddrClean removing addr:", ifp.Name, addr.Prefix)
				NetlinkIpAddrDelete(ifp, addr.Prefix)
			}
		}
	}
}
