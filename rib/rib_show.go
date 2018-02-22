// Copyright 2017 zebra project
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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/coreswitch/netutil"
	"github.com/coreswitch/zebra/policy"
	"github.com/vishvananda/netlink/nl"
)

var RibShowHeader = `Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area
`

var RibShowHeaderDB = `       > - selected route, * - FIB route
`

type RibShowParam struct {
	afi          int
	database     bool
	longerPrefix bool
	address      net.IP
	prefix       *netutil.Prefix
	anchor       *netutil.Prefix
	separator    bool
	typ          uint8
}

func ribTypeShortString(rib *Rib) string {
	switch rib.Type {
	case RIB_KERNEL:
		return "K "
	case RIB_CONNECTED:
		return "C "
	case RIB_STATIC:
		return "S "
	case RIB_RIP:
		return "R "
	case RIB_OSPF:
		return "O "
	case RIB_BGP:
		return "B "
	case RIB_ISIS:
		return "i "
	}
	return ""
}

func (v *Vrf) RibShowIPv4Nexthop(rib *Rib, n *Nexthop, offset int, out io.Writer) {
	if offset != 0 {
		fmt.Fprintf(out, "%*s", offset, " ")
	}

	if rib.Type != RIB_KERNEL && rib.Type != RIB_CONNECTED {
		fmt.Fprintf(out, " [%d/%d]", rib.Distance, rib.Metric)
	}

	if n.IsIfOnly() {
		fmt.Fprintf(out, " is directly connected %s", v.IfName(n.Index))
	} else {
		if n.IsAddrOnly() {
			fmt.Fprintf(out, " via %v", n.IP)
		} else {
			fmt.Fprintf(out, " via %v, %s", n.IP, v.IfName(n.Index))
		}
	}
	switch n.EncapType {
	case nl.LWTUNNEL_ENCAP_SEG6:
		fmt.Fprintf(out, " encap seg6 %s", n.EncapSeg6.String())
	}
}

func (v *Vrf) RibShowEntryJson(t *ShowTask, p *RibShowParam, rib *Rib) {
	if !p.database && !rib.IsFib() {
		return
	}
	if p.separator {
		t.Str += ","
	} else {
		p.separator = true
	}
	bytes, err := json.Marshal(rib)
	if err != nil {
		t.Str += "{}"
		return
	}
	t.Str += string(bytes)
}

func (v *Vrf) RibShowIPv4Entry(t *ShowTask, rib *Rib, database bool) {
	buf := new(bytes.Buffer)

	if !database && !rib.IsFib() {
		return
	}
	selected := ' '
	fib := ' '
	if database {
		if rib.IsFib() {
			fib = '*'
		}
		if rib.IsSelected() {
			selected = '>'
		}
	}
	fmt.Fprintf(buf, "%s %s %c%c %v", ribTypeShortString(rib), "  ", fib, selected, rib.Prefix)
	offset := len(rib.Prefix.String()) + 9

	for pos, nexthop := range rib.Nexthops {
		if pos == 0 {
			v.RibShowIPv4Nexthop(rib, nexthop, 0, buf)
		} else {
			v.RibShowIPv4Nexthop(rib, nexthop, offset, buf)
		}
	}
	if rib.Aux != nil {
		aspath := &policy.ASPath{}
		aspath.DecodeFromBytes(rib.Aux)
		fmt.Fprintf(buf, " aspath %s", aspath.String())
	}
	if rib.PathId != 0 {
		fmt.Fprintf(buf, " pathid %d", rib.PathId)
	}
	fmt.Fprintf(buf, "\n")
	t.Str += buf.String()
}

func (v *Vrf) RibShowIPv6Entry(t *ShowTask, rib *Rib, database bool) {
	buf := new(bytes.Buffer)

	if !database && !rib.IsFib() {
		return
	}
	selected := ' '
	fib := ' '
	if database {
		if rib.IsFib() {
			fib = '*'
		}
		if rib.IsSelected() {
			selected = '>'
		}
	}
	fmt.Fprintf(buf, "%s %c%c %v [%d/%d]\n", ribTypeShortString(rib), fib, selected, rib.Prefix, rib.Distance, rib.Metric)

	if len(rib.Nexthops) == 1 {
		nhop := rib.Nexthops[0]
		switch nhop.EncapType {
		case nl.LWTUNNEL_ENCAP_SEG6:
			fmt.Fprintf(buf, "      encap seg6 %s\n", nhop.EncapSeg6.String())
		}
		if nhop.IsIfOnly() {
			fmt.Fprintf(buf, "      via %s, directly connected\n", v.IfName(nhop.Index))
		} else {
			if nhop.IsAddrOnly() {
				fmt.Fprintf(buf, "      via %v\n", nhop.IP)
			} else {
				fmt.Fprintf(buf, "      via %v, %s\n", nhop.IP, v.IfName(nhop.Index))
			}
		}
	}
	t.Str += buf.String()
}

func (v *Vrf) RibShow(t *ShowTask) {
	p := t.Index.(*RibShowParam)
	if t.First {
		if t.Json {
			t.Str = "["
		} else {
			t.Str = RibShowHeader
			if p.database {
				t.Str += RibShowHeaderDB
			}
			t.Str += "\n"
		}
	}

	ptree := v.ribTable[p.afi]

	var top *netutil.PtreeNode
	if p.anchor != nil {
		n := ptree.Acquire(p.anchor.IP, p.anchor.Length)
		top = ptree.Next(n)
	} else {
		top = ptree.Top()
	}

	var n *netutil.PtreeNode
	for n = top; n != nil; n = ptree.Next(n) {
		ribs := n.Item.(RibSlice)
		for _, rib := range ribs {
			if p.typ != 0 && rib.Type != p.typ {
				continue
			}
			if t.Json {
				v.RibShowEntryJson(t, p, rib)
			} else {
				if p.afi == AFI_IP {
					v.RibShowIPv4Entry(t, rib, p.database)
				} else {
					v.RibShowIPv6Entry(t, rib, p.database)
				}
			}
		}
		p.anchor = netutil.PrefixFromNode(n)
		t.Continue = true
		ptree.Release(n)
		break
	}

	if n == nil && t.Json {
		t.Str += "]"
	}
}

func RibShow(vrfName string, t *ShowTask) {
	vrf := VrfLookupByName(vrfName)
	if vrf == nil {
		return
	}
	vrf.RibShow(t)
}

func ShowIpRoute(t *ShowTask, Args []interface{}) {
	if t.First {
		param := &RibShowParam{
			afi: AFI_IP,
		}
		t.Index = param
	}
	RibShow("", t)
}

func ShowIpRouteType(t *ShowTask, Args []interface{}) {
	typ := RibStringType(Args[0].(string))
	if t.First {
		param := &RibShowParam{
			afi: AFI_IP,
			typ: typ,
		}
		t.Index = param
	}
	RibShow("", t)
}

func ShowIpRouteDatabase(t *ShowTask, Args []interface{}) {
	if t.First {
		param := &RibShowParam{
			afi:      AFI_IP,
			database: true,
		}
		t.Index = param
	}
	RibShow("", t)
}

func ShowIpRouteVrf(t *ShowTask, Args []interface{}) {
	vrfName := Args[0].(string)
	if t.First {
		param := &RibShowParam{
			afi: AFI_IP,
		}
		t.Index = param
	}
	RibShow(vrfName, t)
}

func ShowIpRouteVrfType(t *ShowTask, Args []interface{}) {
	vrfName := Args[0].(string)
	typ := RibStringType(Args[1].(string))
	if t.First {
		param := &RibShowParam{
			afi: AFI_IP,
			typ: typ,
		}
		t.Index = param
	}
	RibShow(vrfName, t)
}

func ShowIpRouteVrfDatabase(t *ShowTask, Args []interface{}) {
	vrfName := Args[0].(string)
	if t.First {
		param := &RibShowParam{
			afi:      AFI_IP,
			database: true,
		}
		t.Index = param
	}
	RibShow(vrfName, t)
}

func ShowIpv6Route(t *ShowTask, Args []interface{}) {
	if t.First {
		param := &RibShowParam{
			afi: AFI_IP6,
		}
		t.Index = param
	}
	RibShow("", t)
}

func ShowIpv6RouteDatabase(t *ShowTask, Args []interface{}) {
	if t.First {
		param := &RibShowParam{
			afi:      AFI_IP6,
			database: true,
		}
		t.Index = param
	}
	RibShow("", t)
}

func ShowIpv6RouteVrf(t *ShowTask, Args []interface{}) {
	vrfName := Args[0].(string)
	if t.First {
		param := &RibShowParam{
			afi: AFI_IP6,
		}
		t.Index = param
	}
	RibShow(vrfName, t)
}

func ShowIpv6RouteVrfDatabase(t *ShowTask, Args []interface{}) {
	vrfName := Args[0].(string)
	if t.First {
		param := &RibShowParam{
			afi:      AFI_IP6,
			database: true,
		}
		t.Index = param
	}
	RibShow(vrfName, t)
}
