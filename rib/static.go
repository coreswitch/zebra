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

	"github.com/coreswitch/netutil"
	"github.com/vishvananda/netlink/nl"
)

type Static struct {
	Prefix   *netutil.Prefix
	Distance uint8
	IfName   string
	Nexthops []*Nexthop
}

func (s *Static) Rib() *Rib {
	rib := &Rib{
		Type: RIB_STATIC,
	}
	for _, nhop := range s.Nexthops {
		rib.Nexthops = append(rib.Nexthops, nhop)
		rib.Nexthop = nhop
	}
	return rib
}

func (v *Vrf) StaticAdd(p *netutil.Prefix, naddr net.IP) error {
	node := v.staticTable[AFI_IP].Acquire(p.IP, p.Length)
	nhop := NewNexthopAddr(naddr)

	var static *Static

	if node.Item != nil {
		static = node.Item.(*Static)
		for _, n := range static.Nexthops {
			if n.Equal(nhop) {
				v.staticTable[AFI_IP].Release(node)
				return fmt.Errorf("Same nexthpo exists")
			}
		}
	} else {
		static = &Static{}
		node.Item = static
	}
	static.Nexthops = append(static.Nexthops, nhop)

	v.RibAdd(p, static.Rib())

	return nil
}

func (v *Vrf) StaticDelete(p *netutil.Prefix, naddr net.IP) error {
	fmt.Println("StaticDelete")
	node := v.staticTable[AFI_IP].Lookup(p.IP, p.Length)
	if node == nil {
		fmt.Println("Can't find route")
		return fmt.Errorf("Can't find the route")
	}
	if node.Item == nil {
		fmt.Println("no static info")
		return fmt.Errorf("No static route information")
	}
	static := node.Item.(*Static)

	nhop := NewNexthopAddr(naddr)
	nhops := []*Nexthop{}
	for _, n := range static.Nexthops {
		if !n.Equal(nhop) {
			nhops = append(nhops, n)
		}
	}
	if len(nhops) == len(static.Nexthops) {
		fmt.Println("Can't find the nexthop")
		return fmt.Errorf("Can't find the nexthop")
	}
	static.Nexthops = nhops

	fmt.Println("StaticDelete", p, static, static.Rib())

	if len(nhops) == 0 {
		fmt.Println("RibDeleteByType")
		v.RibDelete(p, static.Rib())
	} else {
		fmt.Println("RibDelete")
		v.RibAdd(p, static.Rib())
	}
	return nil
}

func (v *Vrf) StaticSeg6SegmentsAdd(p *netutil.Prefix, naddr net.IP, mode string, segs []net.IP) error {
	node := v.staticTable[AFI_IP].Acquire(p.IP, p.Length)
	nhop := NewNexthopAddr(naddr)

	switch mode {
	case "inline":
		nhop.EncapSeg6.Mode = nl.SEG6_IPTUN_MODE_INLINE
	case "encap":
		nhop.EncapSeg6.Mode = nl.SEG6_IPTUN_MODE_ENCAP
	default:
		return fmt.Errorf("Unspported seg6 encap mode:", mode)
	}
	nhop.EncapType = nl.LWTUNNEL_ENCAP_SEG6
	nhop.EncapSeg6.Segments = segs

	var static *Static

	if node.Item != nil {
		static = node.Item.(*Static)
		for _, n := range static.Nexthops {
			if n.Equal(nhop) {
				v.staticTable[AFI_IP].Release(node)
				return fmt.Errorf("Same nexthop exists")
			}
		}
	} else {
		static = &Static{}
		node.Item = static
	}
	static.Nexthops = append(static.Nexthops, nhop)

	v.RibAdd(p, static.Rib())

	return nil
}

func (v *Vrf) StaticSeg6SegmentsDelete(p *netutil.Prefix, naddr net.IP, mode string, segs []net.IP) error {
	fmt.Println("StaticSeg6SegmentsDelete")
	node := v.staticTable[AFI_IP].Lookup(p.IP, p.Length)
	if node == nil {
		fmt.Println("Can't find route")
		return fmt.Errorf("Can't find the route")
	}
	if node.Item == nil {
		fmt.Println("no static info")
		return fmt.Errorf("No static route information")
	}
	static := node.Item.(*Static)

	nhop := NewNexthopAddr(naddr)
	switch mode {
	case "inline":
		nhop.EncapSeg6.Mode = nl.SEG6_IPTUN_MODE_INLINE
	case "encap":
		nhop.EncapSeg6.Mode = nl.SEG6_IPTUN_MODE_ENCAP
	default:
		return fmt.Errorf("Unspported seg6 encap mode:", mode)
	}
	nhop.EncapType = nl.LWTUNNEL_ENCAP_SEG6
	nhop.EncapSeg6.Segments = segs

	nhops := []*Nexthop{}
	for _, n := range static.Nexthops {
		if !n.Equal(nhop) {
			nhops = append(nhops, n)
		}
	}
	if len(nhops) == len(static.Nexthops) {
		fmt.Println("Can't find the nexthop")
		return fmt.Errorf("Can't find the nexthop")
	}
	static.Nexthops = nhops

	if len(nhops) == 0 {
		fmt.Println("RibDelete")
		v.RibDelete(p, static.Rib())
	} else {
		fmt.Println("RibAdd with a nexthop deleted")
		v.RibAdd(p, static.Rib())
	}
	return nil
}
