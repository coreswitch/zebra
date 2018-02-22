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

	"github.com/coreswitch/netutil"
)

var (
	ArpAddHook    func(int, net.IP, net.HardwareAddr)
	ArpDeleteHook func(int, net.IP, net.HardwareAddr)
)

func (v *Vrf) ArpTable(neigh *Neigh) *netutil.Ptree {
	if len(neigh.IP) == net.IPv4len {
		return v.arpTable[AFI_IP]
	} else if len(neigh.IP) == net.IPv6len {
		return v.arpTable[AFI_IP6]
	} else {
		return nil
	}
}

func ArpDump(ptree *netutil.Ptree) {
	fmt.Println("--- Arp start ---")
	for n := ptree.Top(); n != nil; n = ptree.Next(n) {
		neigh := n.Item.(*Neigh)
		fmt.Println(neigh)
	}
	fmt.Println("--- Arp end ---")
}

func (v *Vrf) ArpAdd(neigh *Neigh) {
	ptree := v.ArpTable(neigh)
	if ptree == nil {
		return
	}
	n := ptree.LookupByMaxBits(neigh.IP)
	if n != nil {
	} else {
		n = ptree.AcquireByMaxBits(neigh.IP)
		if ArpAddHook != nil {
			ArpAddHook(neigh.LinkIndex, neigh.IP, neigh.HardwareAddr)
		}
	}
	n.Item = neigh
	//ArpDump(ptree)
}

func (v *Vrf) ArpDelete(neigh *Neigh) {
	ptree := v.ArpTable(neigh)
	if ptree == nil {
		return
	}
	n := ptree.LookupByMaxBits(neigh.IP)
	if n == nil {
		return
	}
	if ArpDeleteHook != nil {
		ArpDeleteHook(neigh.LinkIndex, neigh.IP, neigh.HardwareAddr)
	}
	n.Item = nil
	ptree.Release(n)
	//ArpDump(ptree)
}
