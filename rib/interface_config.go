// Copyright 2016 Zebra Project.
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
	"regexp"

	"github.com/coreswitch/netutil"
)

type IfcFlag uint32

const IFC_FLAG_MAX = 32

const (
	IFC_FLAG_SHUTDOWN = 1 << iota
	IFC_FLAG_MTU
	IFC_FLAG_DESCRIPTION
)

func (f IfcFlag) SetFlag(flag IfcFlag) IfcFlag {
	f |= flag
	return f
}

func (f IfcFlag) UnsetFlag(flag IfcFlag) IfcFlag {
	f &^= flag
	return f
}

func (f IfcFlag) CheckFlag(flag IfcFlag) bool {
	return (f & flag) == flag
}

type InterfaceConfig struct {
	Flags       IfcFlag
	Name        string
	Description string
	Shutdown    bool
	HwAddr      []byte
	Index       int
	L3Index     int
	VlanId      int
	Addrs       [AFI_MAX]IfAddrSlice
	Interface   *Interface
	Mtu         uint32
}

var InterfaceConfigMap = map[string]*InterfaceConfig{}

func NewInterfaceConfig(ifname string) *InterfaceConfig {
	return &InterfaceConfig{Name: ifname}
}

func InterfaceConfigGet(ifname string) *InterfaceConfig {
	ifc := InterfaceConfigMap[ifname]
	if ifc == nil {
		ifc = NewInterfaceConfig(ifname)
		InterfaceConfigMap[ifname] = ifc
		return ifc
	}
	return ifc
}

func InterfaceConfigLookup(ifname string) *InterfaceConfig {
	return InterfaceConfigMap[ifname]
}

func InterfaceConfigPush() {
	for ifname, _ := range IfMap {
		ConfigPush([]string{"interfaces", "interface", ifname})
	}
}

var ShutdownSkipHook func(ifp *Interface) bool

func EsiShutdownSkipHook(ifp *Interface) bool {
	r := regexp.MustCompile(`eth\d+`)
	if r.MatchString(ifp.Name) {
		return true
	}
	return false
}

func (ifc *InterfaceConfig) ShutdownSync(ifp *Interface) {
	if ShutdownSkipHook != nil {
		if ShutdownSkipHook(ifp) {
			return
		}
	}

	if ifc.Shutdown {
		if ifp.IsUp() {
			LinkSetDown(ifp)
		}
	} else {
		if !ifp.IsUp() {
			LinkSetUp(ifp)
		}
	}
}

func (ifc *InterfaceConfig) MtuSync(ifp *Interface) {
	if ifc.Flags.CheckFlag(IFC_FLAG_MTU) {
		if ifc.Mtu != ifp.Mtu {
			LinkSetMtu(ifp, ifc.Mtu)
		}
	} else {
		if ifp.Mtu != ifp.DefaultMtu {
			LinkSetMtu(ifp, ifp.DefaultMtu)
		}
	}
}

func (ifc *InterfaceConfig) Sync(flags IfcFlag) {
	ifp := IfLookupByName(ifc.Name)
	if ifp == nil {
		return
	}
	for i := 0; i < IFC_FLAG_MAX; i++ {
		flag := IfcFlag(1 << uint(i))
		if flags.CheckFlag(flag) {
			switch flag {
			case IFC_FLAG_SHUTDOWN:
				ifc.ShutdownSync(ifp)
			case IFC_FLAG_MTU:
				ifc.MtuSync(ifp)
			case IFC_FLAG_DESCRIPTION:
				ifp.Description = ifc.Description
			}
		}
	}
}

func (ifc *InterfaceConfig) ShutdownSet() {
	ifc.Shutdown = true
	ifc.Flags = ifc.Flags.SetFlag(IFC_FLAG_SHUTDOWN)
	ifc.Sync(IFC_FLAG_SHUTDOWN)
}

func (ifc *InterfaceConfig) ShutdownUnset() {
	ifc.Shutdown = false
	ifc.Flags = ifc.Flags.UnsetFlag(IFC_FLAG_SHUTDOWN)
	ifc.Sync(IFC_FLAG_SHUTDOWN)
}

func (ifc *InterfaceConfig) MtuSet(mtu uint32) {
	ifc.Mtu = mtu
	ifc.Flags = ifc.Flags.SetFlag(IFC_FLAG_MTU)
	ifc.Sync(IFC_FLAG_MTU)
}

func (ifc *InterfaceConfig) MtuUnset() {
	ifc.Mtu = 0
	ifc.Flags = ifc.Flags.UnsetFlag(IFC_FLAG_MTU)
	ifc.Sync(IFC_FLAG_MTU)
}

func (ifc *InterfaceConfig) DescriptionSet(desc string) {
	ifc.Description = desc
	ifc.Flags = ifc.Flags.SetFlag(IFC_FLAG_DESCRIPTION)
	ifc.Sync(IFC_FLAG_DESCRIPTION)
}

func (ifc *InterfaceConfig) DescriptionUnset() {
	ifc.Description = ""
	ifc.Flags = ifc.Flags.UnsetFlag(IFC_FLAG_DESCRIPTION)
	ifc.Sync(IFC_FLAG_DESCRIPTION)
}

func (ifc *InterfaceConfig) AddrAdd(ifName string, addr *netutil.Prefix) error {
	fmt.Println("[API] AddrAdd start:", ifName, addr)
	ifaddr := &IfAddr{Prefix: addr}
	afi := addr.AFI()

	if lookup := ifc.Addrs[afi].Lookup(addr); lookup == nil {
		ifc.Addrs[afi] = append(ifc.Addrs[afi], ifaddr)

		ifp := IfLookupByName(ifName)
		if ifp == nil {
			fmt.Println("IfcAddrAdd can't find ifp", ifName)
			return fmt.Errorf("IfcAddrAdd can't find ifp %s", ifName)
		}
		NetlinkIpAddrAdd(ifp, addr)
		if !ifp.IsUp() {
			fmt.Println("IfcAddrAdd !ifp.IsUp()", ifName)
		}
	}

	fmt.Println("[API] AddrAdd end:", ifName, addr)

	return nil
}

func (ifc *InterfaceConfig) AddrDelete(ifName string, addr *netutil.Prefix) error {
	afi := addr.AFI()
	if afi == AFI_MAX {
		return fmt.Errorf("AddrDelete unknown AFI")
	}
	addrs, found := ifc.Addrs[afi].Delete(addr)
	if found == nil {
		// XXX Can't find ifc addr
		fmt.Println("IfcAddrDelete can't find address", addr)
		return fmt.Errorf("IfcAddrDelete can't find address %s", addr)
	}
	ifc.Addrs[afi] = addrs

	ifp := IfLookupByName(ifName)
	if ifp == nil {
		fmt.Println("IfcAddrDelete can't find ifp", ifName)
		return fmt.Errorf("IfcAddrDelete can't find ifp %s", ifName)
	}
	NetlinkIpAddrDelete(ifp, addr)

	fmt.Println("[API] AddrDelete end:", ifName, addr)

	return nil
}
