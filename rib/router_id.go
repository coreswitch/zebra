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
	"fmt"
	"net"

	"github.com/coreswitch/netutil"
)

const (
	RouterIdTypeUnset = iota
	RouterIdTypeConfig
	RouterIdTypeLoopback
	RouterIdTypeAutomatic
)

var RouterIdTypeMap = map[int]string{
	RouterIdTypeUnset:     "router id not set",
	RouterIdTypeConfig:    "configured",
	RouterIdTypeLoopback:  "automatic from looback",
	RouterIdTypeAutomatic: "automatic",
}

type RouterId struct {
	Type          int
	Configured    bool
	ConfigId      net.IP
	Id            net.IP
	LoopbackAddrs *netutil.Ptree
	OtherAddrs    *netutil.Ptree
}

func (rid *RouterId) Init() {
	rid.Id = make(net.IP, net.IPv4len)
	rid.LoopbackAddrs = netutil.NewPtree(32)
	rid.LoopbackAddrs.ReverseOrderSet()
	rid.OtherAddrs = netutil.NewPtree(32)
	rid.OtherAddrs.ReverseOrderSet()
}

func (v *Vrf) routerIdAddrsGet(ifp *Interface) *netutil.Ptree {
	if ifp.IsLoopback() {
		return v.routerId.LoopbackAddrs
	} else {
		return v.routerId.OtherAddrs
	}
}

func (v *Vrf) RouterIdAdd(ifp *Interface, addr *IfAddr) {
	if len(addr.Prefix.IP) != net.IPv4len {
		return
	}
	if addr.Prefix.IP.IsLoopback() {
		return
	}
	q := v.routerIdAddrsGet(ifp)
	q.AcquireWithItem(addr.Prefix.IP, 32, addr)
	v.RouterIdUpdate()
}

func (v *Vrf) RouterIdDelete(ifp *Interface, addr *IfAddr) {
	if len(addr.Prefix.IP) != net.IPv4len {
		return
	}
	if addr.Prefix.IP.IsLoopback() {
		return
	}
	q := v.routerIdAddrsGet(ifp)
	q.ReleaseWithItem(addr.Prefix.IP, 32, addr)
	v.RouterIdUpdate()
}

func (v *Vrf) RouterId() net.IP {
	return v.routerId.Id
}

func (v *Vrf) RouterIdUpdate() {
	current := v.routerId.Id

	var update net.IP
	var typ int

	// Check configured router ID.
	if v.routerId.Configured {
		update = v.routerId.ConfigId
		typ = RouterIdTypeConfig
	} else {
		// Check loopback router ID.
		n := v.routerId.LoopbackAddrs.LookupTop()
		if n != nil {
			update = n.Key()
			typ = RouterIdTypeLoopback
		} else {
			// Check automatic router ID.
			n = v.routerId.OtherAddrs.LookupTop()
			if n != nil {
				update = n.Key()
				typ = RouterIdTypeAutomatic
			} else {
				// Otherwise set to 0.0.0.0.
				update = make(net.IP, 4)
			}
		}
	}
	// Compare curent and update
	if !current.Equal(update) {
		copy(v.routerId.Id, update)
		v.routerId.Type = typ

		// Advertise update.
		v.RouterIdUpdateNotification()
	}
}

func (v *Vrf) RouterIdShow() string {
	return fmt.Sprintf("Router ID: %v (%s)", v.routerId.Id, RouterIdTypeMap[v.routerId.Type])
}

func (v *Vrf) RouterIdSet(id net.IP) {
	v.routerId.Configured = true
	v.routerId.ConfigId = id
	v.RouterIdUpdate()
}

func (v *Vrf) RouterIdUnset() {
	v.routerId.Configured = false
	v.RouterIdUpdate()
}

func (v *Vrf) RouterIdUpdateNotification() {
	RouterIdUpdate(v.Index, v.routerId.Id)
}
