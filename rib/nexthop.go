// Copyright 2016 Zebra Project
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
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
)

type EncapSEG6 struct {
	Mode     int
	Segments []net.IP
}

func (e EncapSEG6) String() string {
	seg := netlink.SEG6Encap{Mode: e.Mode}
	seg.Segments = e.Segments
	return seg.String()
}
func (e EncapSEG6) Equal(x EncapSEG6) bool {
	if e.Mode != x.Mode {
		return false
	}
	if len(e.Segments) != len(x.Segments) {
		return false
	}
	for i := range e.Segments {
		if !e.Segments[i].Equal(x.Segments[i]) {
			return false
		}
	}
	return true
}

type EncapSEG6Local netlink.SEG6LocalEncap

func (e EncapSEG6Local) String() string {
	seg := netlink.SEG6LocalEncap{
		Flags:   e.Flags,
		Action:  e.Action,
		Table:   e.Table,
		InAddr:  e.InAddr,
		In6Addr: e.In6Addr,
		Iif:     e.Iif,
		Oif:     e.Oif,
	}
	seg.Segments = e.Segments
	return seg.String()
}
func (e EncapSEG6Local) Equal(o EncapSEG6Local) bool {
	// compare all arrays first
	for i := range e.Flags {
		if e.Flags[i] != o.Flags[i] {
			return false
		}
	}
	if len(e.Segments) != len(o.Segments) {
		return false
	}
	for i := range e.Segments {
		if !e.Segments[i].Equal(o.Segments[i]) {
			return false
		}
	}
	// compare values
	if !e.InAddr.Equal(o.InAddr) || !e.In6Addr.Equal(o.In6Addr) {
		return false
	}
	if e.Action != o.Action || e.Table != o.Table || e.Iif != o.Iif || e.Oif != o.Oif {
		return false
	}
	return true
}

type Nexthop struct {
	net.IP
	Index          IfIndex
	EncapType      int
	EncapSeg6      EncapSEG6
	EncapSeg6Local EncapSEG6Local
}

func (n *Nexthop) AddressString() string {
	if n.IP == nil {
		return ""
	} else {
		return n.IP.String()
	}
}

func (n *Nexthop) InterfaceString() string {
	if n.Index == 0 {
		return ""
	}
	return IfNameByIndex(n.Index)
}

func (n *Nexthop) MarshalJSON() ([]byte, error) {
	return json.Marshal(
		struct {
			Address   string `json:"address,omitempty"`
			Interface string `json:"interface,omitempty"`
		}{
			Address:   n.AddressString(),
			Interface: n.InterfaceString(),
		},
	)
}

func (n Nexthop) String() string {
	strs := []string{}
	if n.IP != nil {
		strs = append(strs, n.IP.String())
	}
	strs = append(strs, fmt.Sprintf("ifindex %d", n.Index))
	switch n.EncapType {
	case nl.LWTUNNEL_ENCAP_SEG6:
		strs = append(strs, fmt.Sprintf("encap seg6 %s", n.EncapSeg6.String()))
	case nl.LWTUNNEL_ENCAP_SEG6_LOCAL:
		strs = append(strs, fmt.Sprintf("encap seg6local %s", n.EncapSeg6Local.String()))
	}
	return fmt.Sprintf("%s", strings.Join(strs, " "))
}

func (n *Nexthop) IsIfOnly() bool {
	if n.IP == nil && n.Index != 0 {
		return true
	}
	return false
}

func (n *Nexthop) IsAddrOnly() bool {
	if n.IP != nil && n.Index == 0 {
		return true
	}
	return false
}

func (n *Nexthop) IsAddrIf() bool {
	if n.IP != nil && n.Index != 0 {
		return true
	}
	return false
}

func (n *Nexthop) Equal(nn *Nexthop) bool {
	if n.Index != nn.Index {
		return false
	}
	if n.IP == nil && nn.IP == nil {
		return true
	}
	if n.IP == nil || nn.IP == nil {
		return false
	}
	if !n.IP.Equal(nn.IP) {
		return false
	}

	switch n.EncapType {
	case nl.LWTUNNEL_ENCAP_SEG6:
		if !n.EncapSeg6.Equal(nn.EncapSeg6) {
			return false
		}
	case nl.LWTUNNEL_ENCAP_SEG6_LOCAL:
		if !n.EncapSeg6Local.Equal(nn.EncapSeg6Local) {
			return false
		}
	}
	return true
}

func NewNexthopIf(index IfIndex) *Nexthop {
	return &Nexthop{IP: nil, Index: index}
}

func NewNexthopAddr(ip net.IP) *Nexthop {
	return &Nexthop{IP: ip, Index: 0}
}

func NewNexthopAddrIf(ip net.IP, index IfIndex) *Nexthop {
	return &Nexthop{IP: ip, Index: index}
}
