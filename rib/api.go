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
	"net"

	"github.com/coreswitch/cmd"
	"github.com/coreswitch/netutil"
)

var Parser *cmd.Node

func Config(command int, path []string) {
	if command == cmd.Set {
		fmt.Println("[cmd] add", path)
	} else {
		fmt.Println("[cmd] del", path)
	}
	ret, fn, args, _ := Parser.ParseCmd(path)
	if ret == cmd.ParseSuccess {
		fn.(func(int, cmd.Args) int)(command, args)
	}
}

func RouterIdApi(Cmd int, Args cmd.Args) int {
	id := Args[0].(net.IP)
	if Cmd == cmd.Set {
		VrfDefault().RouterIdSet(id)
	} else {
		VrfDefault().RouterIdUnset()
	}
	return cmd.Success
}

func IPv4RouteApi(Cmd int, Args cmd.Args) int {
	prefix := Args[0].(*netutil.Prefix)
	nexthop := Args[1].(net.IP)
	fmt.Println("Static route:", prefix, nexthop)
	if Cmd == cmd.Set {
		server.StaticAdd(prefix, nexthop)
	} else {
		server.StaticDelete(prefix, nexthop)
	}
	return cmd.Success
}

func IPv4VrfRouteApi(Cmd int, Args cmd.Args) int {
	vrfName := Args[0].(string)
	prefix := Args[1].(*netutil.Prefix)
	nexthop := Args[2].(net.IP)

	fmt.Println("Vrf Static route:", vrfName, prefix, nexthop)

	vrf := VrfLookupByName(vrfName)
	if vrf == nil {
		fmt.Println("IPv4VrfStatic: Can't find VRF")
		return cmd.Success
	}
	if Cmd == cmd.Set {
		vrf.StaticAdd(prefix, nexthop)
	} else {
		vrf.StaticDelete(prefix, nexthop)
	}

	return cmd.Success
}

func IPv4VrfRouteApi2(Cmd int, Args cmd.Args) int {
	return cmd.Success
}

func IPv4RouteSeg6SegmentsApi(Cmd int, Args cmd.Args) int {
	prefix := Args[0].(*netutil.Prefix)
	nexthop := Args[1].(net.IP)
	mode := Args[2].(string)
	Args = Args[3:]
	segs := make([]net.IP, 0, len(Args))
	for _, arg := range Args {
		segs = append(segs, arg.(net.IP))
	}
	fmt.Println("Static IPv4 seg6 segments:", prefix, nexthop, mode, segs)
	if Cmd == cmd.Set {
		server.StaticSeg6SegmentsAdd(prefix, nexthop, mode, segs)
	} else {
		server.StaticSeg6SegmentsDelete(prefix, nexthop, mode, segs)
	}
	return cmd.Success
}

func IPv6RouteApi(Cmd int, Args cmd.Args) int {
	prefix := Args[0].(*netutil.Prefix)
	nexthop := Args[1].(net.IP)
	fmt.Println("Static route:", prefix, nexthop)
	if Cmd == cmd.Set {
		server.StaticAdd(prefix, nexthop)
	} else {
		server.StaticDelete(prefix, nexthop)
	}
	return cmd.Success
}

func IPv6RouteSeg6SegmentsApi(Cmd int, Args cmd.Args) int {
	prefix := Args[0].(*netutil.Prefix)
	nexthop := Args[1].(net.IP)
	mode := Args[2].(string)
	Args = Args[3:]
	fmt.Println("Static IPv6 seg6 segments:", prefix, nexthop, mode, Args)
	if Cmd == cmd.Set {
	} else {
	}
	return cmd.Success
}

func IPv6VrfRouteApi(Cmd int, Args cmd.Args) int {
	vrfName := Args[0].(string)
	prefix := Args[1].(*netutil.Prefix)
	nexthop := Args[2].(net.IP)
	fmt.Println("Vrf Static route:", vrfName, prefix, nexthop)
	vrf := VrfLookupByName(vrfName)
	if vrf == nil {
		fmt.Println("Can't find VRF")
		return cmd.Success
	}
	if Cmd == cmd.Set {
		vrf.StaticAdd(prefix, nexthop)
	} else {
		vrf.StaticDelete(prefix, nexthop)
	}
	return cmd.Success
}

var InterfaceVlanBindHook func(int, int)

var InterfaceVlanFea bool

func InterfaceVlanApi(Cmd int, Args cmd.Args) int {
	ifName := Args[0].(string)
	vlanId := Args[1].(uint64)
	ifc := InterfaceConfigGet(ifName)
	ifc.VlanId = int(vlanId)

	if InterfaceVlanFea {
		vlan := VlanDB.Lookup(uint16(vlanId))
		if vlan != nil {
			if InterfaceVlanBindHook != nil {
				InterfaceVlanBindHook(int(vlanId), ifc.Index)
			}
		}
		vlan.Ports = append(vlan.Ports, ifc)
	} else {
		if Cmd == cmd.Set {
			server.VIFAdd(ifName, vlanId)
		} else {
			server.VIFDelete(ifName, vlanId)
		}
	}

	return cmd.Success
}

func InterfaceVrfApi(Cmd int, Args cmd.Args) int {
	ifName := Args[0].(string)
	vrfName := Args[1].(string)

	if Cmd == cmd.Set {
		server.IfVrfBind(ifName, vrfName)
	} else {
		server.IfVrfUnbind(ifName, vrfName)
	}

	return cmd.Success
}

func InterfaceAddress(Cmd int, Args cmd.Args) int {
	ifName := Args[0].(string)
	addr := Args[1].(*netutil.Prefix)
	if Cmd == cmd.Set {
		server.AddrAdd(ifName, addr)
	} else if Cmd == cmd.Delete {
		server.AddrDelete(ifName, addr)
	}
	return cmd.Success
}

func InterfaceShutdown(Cmd int, Args cmd.Args) int {
	ifname := Args[0].(string)
	ifc := InterfaceConfigGet(ifname)
	if Cmd == cmd.Set {
		ifc.ShutdownSet()
	} else {
		ifc.ShutdownUnset()
	}
	return cmd.Success
}

func InterfaceDescription(Cmd int, Args cmd.Args) int {
	ifname := Args[0].(string)
	desc := Args[1].(string)
	ifc := InterfaceConfigGet(ifname)
	if Cmd == cmd.Set {
		ifc.DescriptionSet(desc)
	} else {
		ifc.DescriptionUnset()
	}
	return cmd.Success
}

func InterfaceMtu(Cmd int, Args cmd.Args) int {
	ifname := Args[0].(string)
	ifc := InterfaceConfigGet(ifname)
	mtu := Args[1].(uint64)
	fmt.Println("MTU", ifc, mtu)
	if Cmd == cmd.Set {
		ifc.MtuSet(uint32(mtu))
	} else {
		ifc.MtuUnset()
	}
	return cmd.Success
}

func VrfApi(Cmd int, Args cmd.Args) int {
	name := Args[0].(string)
	if Cmd == cmd.Set {
		server.VrfAdd(name)
	} else {
		server.VrfDelete(name)
	}
	return cmd.Success
}

func VrfHubNodeApi(Cmd int, Args cmd.Args) int {
	name := Args[0].(string)
	hubNode := Args[1].(string)
	if Cmd == cmd.Set {
		fmt.Println("HubNode add", name, hubNode)
		VrfHubNodeAdd(name, hubNode)
	} else {
		fmt.Println("HubNode delete", name, hubNode)
		VrfHubNodeDelete(name, hubNode)
	}
	return cmd.Success
}

func InitAPI() {
	Parser = cmd.NewParser()
	Parser.InstallCmd([]string{"vrf", "name", "WORD"}, VrfApi)
	Parser.InstallCmd([]string{"vlans", "vlan", "<1-4096>"}, VlanApi)
	Parser.InstallCmd([]string{"interfaces", "interface", "WORD", "vlan", "<1-4096>"}, InterfaceVlanApi)
	Parser.InstallCmd([]string{"interfaces", "interface", "WORD", "vlans", "<1-4096>"}, InterfaceVlanApi)
	Parser.InstallCmd([]string{"interfaces", "interface", "WORD", "vrf", "WORD"}, InterfaceVrfApi)
	Parser.InstallCmd([]string{"interfaces", "interface", "WORD", "ipv4", "address", "A.B.C.D/M"}, InterfaceAddress)
	Parser.InstallCmd([]string{"interfaces", "interface", "WORD", "ipv6", "address", "X:X::X:X/M"}, InterfaceAddress)
	Parser.InstallCmd([]string{"interfaces", "interface", "WORD", "shutdown"}, InterfaceShutdown)
	Parser.InstallCmd([]string{"interfaces", "interface", "WORD", "description", "LINE"}, InterfaceDescription)
	Parser.InstallCmd([]string{"interfaces", "interface", "WORD", "mtu", "<68-65535>"}, InterfaceMtu)
	Parser.InstallCmd([]string{"routing-options", "router-id", "A.B.C.D"}, RouterIdApi)
	Parser.InstallCmd([]string{"routing-options", "ipv4", "route", "A.B.C.D/M", "nexthop", "A.B.C.D"}, IPv4RouteApi)
	Parser.InstallCmd([]string{"routing-options", "ipv4", "route-srv6", "A.B.C.D/M", "nexthop", "A.B.C.D", "seg6", "WORD", "segments", "X:X::X:X", "&"}, IPv4RouteSeg6SegmentsApi)

	Parser.InstallCmd([]string{"routing-options", "ipv6", "route", "X:X::X:X/M", "nexthop", "X:X::X:X"}, IPv6RouteApi)
	Parser.InstallCmd([]string{"routing-options", "ipv6", "route-srv6", "X:X::X:X/M", "nexthop", "X:X::X:X", "seg6", "WORD", "segments", "X:X::X:X", "&"}, IPv6RouteSeg6SegmentsApi)

	Parser.InstallCmd([]string{"vrf", "name", "WORD", "hub-node", "WORD"}, VrfHubNodeApi)
	Parser.InstallCmd([]string{"vrf", "name", "WORD", "static", "route", "A.B.C.D/M", "nexthop", "A.B.C.D"}, IPv4VrfRouteApi)
	Parser.InstallCmd([]string{"vrf", "name", "WORD", "static", "route", "A.B.C.D/M", "interface", "WORD"}, IPv4VrfRouteApi2)
}
