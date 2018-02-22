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
)

type Nexthop struct {
	net.IP
	Index IfIndex
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
	if n.IP == nil {
		return fmt.Sprintf("ifindex %d", n.Index)
	} else {
		return n.IP.String() + " " + fmt.Sprintf("ifindex %d", n.Index)
	}
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
	if n.IP.Equal(nn.IP) {
		return true
	}
	return false
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
