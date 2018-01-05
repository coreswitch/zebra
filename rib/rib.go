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
	"encoding/json"
	"fmt"
	"net"

	"github.com/coreswitch/netutil"
)

const (
	AFI_IP = iota
	AFI_IP6
	AFI_MAX
)

const (
	RIB_UNKNOWN uint8 = iota
	RIB_KERNEL
	RIB_CONNECTED
	RIB_STATIC
	RIB_RIP
	RIB_OSPF
	RIB_ISIS
	RIB_BGP
)

var ribTypeString = map[uint8]string{
	RIB_UNKNOWN:   "unknown",
	RIB_KERNEL:    "kernel",
	RIB_CONNECTED: "connected",
	RIB_STATIC:    "static",
	RIB_RIP:       "rip",
	RIB_OSPF:      "ospf",
	RIB_ISIS:      "isis",
	RIB_BGP:       "bgp",
}

var ribStringType = map[string]uint8{
	"unknown":   RIB_UNKNOWN,
	"kernel":    RIB_KERNEL,
	"connected": RIB_CONNECTED,
	"static":    RIB_STATIC,
	"rip":       RIB_RIP,
	"ospf":      RIB_OSPF,
	"isis":      RIB_ISIS,
	"bgp":       RIB_BGP,
}

const (
	RIB_SUB_OSPF_IA uint8 = iota
	RIB_SUB_OSPF_NSSA_1
	RIB_SUB_OSPF_NSSA_2
	RIB_SUB_OSPF_EXTERNAL_1
	RIB_SUB_OSPF_EXTERNAL_2
	RIB_SUB_BGP_IBGP
	RIB_SUB_BGP_EBGP
	RIB_SUB_BGP_CONFED
	RIB_SUB_ISIS_L1
	RIB_SUB_ISIS_L2
	RIB_SUB_ISIS_IA
)

const (
	DISTANCE_KERNEL    = 0
	DISTANCE_CONNECTED = 0
	DISTANCE_STATIC    = 1
	DISTANCE_RIP       = 120
	DISTANCE_OSPF      = 110
	DISTANCE_ISIS      = 115
	DISTANCE_EBGP      = 20
	DISTANCE_IBGP      = 200
	DISTNACE_INFINITY  = 255
)

var distanceMap = map[uint8]uint8{
	RIB_KERNEL:    DISTANCE_KERNEL,
	RIB_CONNECTED: DISTANCE_CONNECTED,
	RIB_STATIC:    DISTANCE_STATIC,
	RIB_RIP:       DISTANCE_RIP,
	RIB_OSPF:      DISTANCE_OSPF,
	RIB_ISIS:      DISTANCE_ISIS,
	RIB_BGP:       DISTANCE_IBGP, // EBGP default distance is 20.
}

type RibFlag uint

const (
	flagSelected RibFlag = 1 << iota
	flagFib
	flagDistance
	flagMetric
	RIB_FLAG_RESOLVED  // Static route and iBGP (multipath eBGP) next hop will be resolved.
	RIB_FLAG_BLACKHOLE // This is black hole route.
)

type Rib struct {
	Flags    RibFlag
	Prefix   *netutil.Prefix
	Type     uint8
	SubType  uint8
	Distance uint8
	Metric   uint32
	//IfAddr   *IfAddr
	Nexthop  *Nexthop
	Nexthops []*Nexthop
	Src      interface{}
	Redist   bool
}

type RibSlice []*Rib

func (r *Rib) MarshalJSON() ([]byte, error) {
	ribJSON := struct {
		Prefix   *netutil.Prefix `json:"prefix"`
		Nexthops []*Nexthop      `json:"nexthops"`
		Distance uint8           `json:"distance"`
		Metric   uint32          `json:"metric"`
		Type     string          `json:"type"`
	}{
		Prefix:   r.Prefix,
		Distance: r.Distance,
		Metric:   r.Metric,
	}
	ribJSON.Type = RibTypeString(r.Type)
	ribJSON.Nexthops = r.Nexthops
	if r.Nexthop != nil {
		ribJSON.Nexthops = append(ribJSON.Nexthops, r.Nexthop)
	}
	return json.Marshal(ribJSON)
}

func RibTypeString(typ uint8) string {
	if str, ok := ribTypeString[typ]; !ok {
		return "unknown"
	} else {
		return str
	}
}

func RibStringType(str string) uint8 {
	if typ, ok := ribStringType[str]; !ok {
		return RIB_UNKNOWN
	} else {
		return typ
	}
}

func (r *Rib) SetFlag(flag RibFlag) {
	r.Flags |= flag
}

func (r *Rib) UnsetFlag(flag RibFlag) {
	r.Flags &^= flag
}

func (r *Rib) CheckFlag(flag RibFlag) bool {
	return (r.Flags & flag) == flag
}

func (r *Rib) IsFib() bool {
	return r.CheckFlag(flagFib)
}

func (r *Rib) SetFib() {
	r.SetFlag(flagFib)
}

func (r *Rib) UnsetFib() {
	r.UnsetFlag(flagFib)
}

func (r *Rib) IsSelected() bool {
	return r.CheckFlag(flagSelected)
}

func (r *Rib) SetSelected() {
	r.SetFlag(flagSelected)
}

func (r *Rib) UnsetSelected() {
	r.UnsetFlag(flagSelected)
}

func (r *Rib) IsSelectedFib() bool {
	return r.IsSelected() && r.IsFib()
}

func (r *Rib) IsSystem() bool {
	return r.Type == RIB_KERNEL || r.Type == RIB_CONNECTED
}

func (r *Rib) HasDistance() bool {
	return r.CheckFlag(flagDistance)
}

func (r *Rib) HasMetric() bool {
	return r.CheckFlag(flagMetric)
}

func (r *Rib) SetResolved() {
	r.SetFlag(RIB_FLAG_RESOLVED)
}

func (r *Rib) UnsetResolved() {
	r.UnsetFlag(RIB_FLAG_RESOLVED)
}

func (r *Rib) IsResolved() bool {
	return r.CheckFlag(RIB_FLAG_RESOLVED)
}

func (v *Vrf) AfiPtree(p *netutil.Prefix) *netutil.Ptree {
	afi := p.AFI()
	if afi == AFI_MAX {
		return nil
	}
	return v.ribTable[afi]
}

func DistanceCalc(typ uint8, subType uint8) uint8 {
	distance := distanceMap[typ]
	if typ == RIB_BGP && subType == RIB_SUB_BGP_EBGP {
		distance = DISTANCE_EBGP
	}
	return distance
}

func NewRib(p *netutil.Prefix, ri *Rib) *Rib {
	rib := &Rib{
		Flags:    ri.Flags,
		Type:     ri.Type,
		SubType:  ri.SubType,
		Prefix:   p,
		Nexthop:  ri.Nexthop,
		Nexthops: ri.Nexthops,
		Metric:   ri.Metric,
		Src:      ri.Src,
	}

	if ri.HasDistance() {
		rib.Distance = ri.Distance
	} else {
		rib.Distance = DistanceCalc(ri.Type, ri.SubType)
	}
	return rib
}

func (rib *Rib) Equal(ri *Rib) bool {
	if rib.Type != ri.Type {
		return false
	}
	switch rib.Type {
	case RIB_CONNECTED:
		if rib.Nexthop.Equal(ri.Nexthop) && rib.Src == ri.Src {
			return true
		} else {
			return false
		}
	case RIB_STATIC:
		return true
	// 	if rib.Nexthop.Equal(ri.Nexthop) {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// case RIB_BGP:
	// 	if rib.Src == ri.Src {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	default:
		// Other type of routes are considered as implicit withdraw.
		return true
	}
}

var AddPathDefault bool

func (v *Vrf) RibProcess(p *netutil.Prefix, ribs RibSlice, del *Rib, resolve bool) {
	var fib *Rib
	var selected *Rib
	var preSelected *Rib
	var def bool

	// Default route.
	if p.Length == 0 && AddPathDefault {
		def = true
	}

	// Traverse RIBs to find FIB and selected one.
	for _, rib := range ribs {
		if def && rib.Type == RIB_KERNEL {
			rib.SetFib()
			rib.SetSelected()
			continue
		}

		if rib.IsFib() {
			fib = rib
		}
		if rib.IsSelected() {
			preSelected = rib
		}
		if resolve {
			v.Resolve(rib)
		}
		if !rib.IsResolved() {
			continue
		}
		if selected == nil {
			selected = rib
		} else {
			switch {
			case rib.Distance < selected.Distance:
				// Distance is smaller, take it.
				selected = rib
			case rib.Distance == selected.Distance:
				// Same distance and smaller metric, take it.
				if rib.Metric < selected.Metric {
					selected = rib
				}
			case rib.Distance > selected.Distance:
				// Do nothing
			}
		}
	}

	// When deleted RIB is FIB.
	if del != nil && del.IsFib() {
		fib = del
	}

	// Update selected flag.
	if preSelected != selected {
		if preSelected != nil {
			preSelected.UnsetSelected()
		}
		if selected != nil {
			selected.SetSelected()
		}
	}

	// FIB and selected is same.
	if fib == selected && fib != nil {
		if fib.Type == RIB_STATIC || fib.Type == RIB_BGP {
			NetlinkRouteAdd(p, fib, v.Index)
		}
		return
	}

	// Withdraw old FIB
	if fib != nil {
		if !fib.IsSystem() {
			NetlinkRouteDelete(p, fib, v.Index)
		}
		if !fib.Redist {
			RedistIPv4Delete(v.Index, p, fib)
		}
		fib.UnsetFib()
	}

	if selected != nil {
		if !selected.IsSystem() {
			err := NetlinkRouteAdd(p, selected, v.Index)
			if err == nil {
				selected.SetFib()
			}
		}
		if !selected.Redist {
			RedistIPv4Add(v.Index, p, selected, nil)
		}
		selected.SetFib()
	}
}

func (v *Vrf) RibAdd(p *netutil.Prefix, ri *Rib) {
	v.Mutex.Lock()
	defer v.Mutex.Unlock()
	p.ApplyMask()
	ptree := v.AfiPtree(p)
	n := ptree.Acquire(p.IP, p.Length)

	// System route is already in FIB.
	// if ri.IsSystem() {
	// 	ri.SetFib()
	// }

	// Resolve nexthop.
	v.Resolve(ri)

	// Updated ribs.
	var found *Rib
	var ribs RibSlice

	if n.Item != nil && ri.Type == RIB_BGP {
		for _, rib := range n.Item.(RibSlice) {
			if rib.Type == RIB_BGP {
				src := ri.Src.(net.Conn)
				dst := rib.Src.(net.Conn)
				if ClientVersion(src) == 3 && ClientVersion(dst) == 2 {
					ptree.Release(n)
					return
				}

				if src == dst {
					if rib.Nexthop != nil && ri.Nexthop != nil {
						if rib.Nexthop.Equal(ri.Nexthop) {
							fmt.Println("Same source and same nexthop, do nothing")
							ptree.Release(n)
							return
						}
					}
				}
			}
		}
	}

	// Check this rib replace existing one or not.
	if n.Item != nil {
		for _, rib := range n.Item.(RibSlice) {
			if rib.Equal(ri) {
				found = rib
			} else {
				ribs = append(ribs, rib)
			}
		}
	}

	// Replace of the RIB.
	if found != nil {
		ptree.Release(n)
	}

	// Append a new RIB.
	rib := NewRib(p, ri)
	ribs = append(ribs, rib)
	n.Item = ribs

	// Redistribute check.
	if ri.Type == RIB_BGP {
		if conn := ri.Src.(net.Conn); conn != nil && ClientVersion(conn) == 2 {
			RedistIPv4Add(v.Index, p, rib, nil)
			rib.Redist = true
		}
	}

	// Process the rib.
	v.RibProcess(p, ribs, found, false)

	// Invoke wakler.
	v.RibWalker(p.AFI())
}

func (v *Vrf) RibDelete(p *netutil.Prefix, ri *Rib) {
	v.Mutex.Lock()
	defer v.Mutex.Unlock()
	p.ApplyMask()
	ptree := v.AfiPtree(p)

	n := ptree.Lookup(p.IP, p.Length)
	if n == nil {
		return
	}
	if n.Item == nil {
		fmt.Println("n.Item is nil, ", p, ri)
		return
	}

	var found *Rib
	var ribs RibSlice
	for _, rib := range n.Item.(RibSlice) {
		if rib.Equal(ri) {
			found = rib
		} else {
			ribs = append(ribs, rib)
		}
	}
	if found == nil {
		return
	}
	if len(ribs) == 0 {
		n.Item = nil
	} else {
		n.Item = ribs
	}
	ptree.Release(n)

	// Redistribute check.
	if found.Type == RIB_BGP {
		if conn := found.Src.(net.Conn); conn != nil && ClientVersion(conn) == 2 {
			RedistIPv4Delete(v.Index, p, found)
			found.Redist = true
		}
	}

	v.RibProcess(p, ribs, found, false)

	// Invoke wakler.
	v.RibWalker(p.AFI())
}

func (v *Vrf) RibClean(src interface{}) {
	v.Mutex.Lock()
	defer v.Mutex.Unlock()
	for _, ptree := range []*netutil.Ptree{v.ribTable[AFI_IP], v.ribTable[AFI_IP6]} {
		for n := ptree.Top(); n != nil; n = ptree.Next(n) {
			var found *Rib
			var ribs RibSlice

			for _, rib := range n.Item.(RibSlice) {
				if rib.Src == src {
					found = rib
				} else {
					ribs = append(ribs, rib)
				}
			}
			if found == nil {
				continue
			}
			if len(ribs) == 0 {
				n.Item = nil
			} else {
				n.Item = ribs
			}
			ptree.Release(n)

			p := netutil.PrefixFromIPPrefixlen(n.Key(), n.KeyLength())

			if conn := src.(net.Conn); conn != nil && ClientVersion(conn) == 2 && found.IsFib() {
				RedistIPv4Delete(v.Index, p, found)
				found.Redist = true
			}

			v.RibProcess(p, ribs, found, false)
		}
	}
}

func (v *Vrf) IsValid(rib *Rib) bool {
	if rib.Type == RIB_KERNEL {
		ifp := v.IfLookupByIndex(rib.Nexthop.Index)
		if ifp != nil && ifp.IsUp() {
			return true
		} else {
			return false
		}
	}
	return true
}

func (v *Vrf) RibSync(afi int) {
	ptree := v.ribTable[afi]
	for n := ptree.Top(); n != nil; n = ptree.Next(n) {
		var found *Rib
		var ribs RibSlice
		for _, rib := range n.Item.(RibSlice) {
			if !v.IsValid(rib) {
				found = rib
			} else {
				ribs = append(ribs, rib)
			}
			if found != nil {
				if len(ribs) == 0 {
					n.Item = nil
				} else {
					n.Item = ribs
				}
				ptree.Release(n)
			}
		}
	}
}

func RibAdd(index int, p *netutil.Prefix, ri *Rib) {
	if vrf := VrfLookupByIndex(index); vrf != nil {
		vrf.RibAdd(p, ri)
	}
}

func RibDelete(index int, p *netutil.Prefix, ri *Rib) {
	if vrf := VrfLookupByIndex(index); vrf != nil {
		vrf.RibDelete(p, ri)
	}
}

func RibClearSrc(src interface{}) {
	for _, vrf := range VrfMap {
		vrf.RibClean(src)
	}
}

func (v *Vrf) Resolve(ri *Rib) {
	if ri.Type == RIB_STATIC || ri.Type == RIB_BGP {
		ri.UnsetResolved()
		//fmt.Println("resovling nexthop for static/bgp")
		if ri.Nexthop != nil {
			//fmt.Println("resolving nexthop", ri.Nexthop)
			if ri.Nexthop.IP == nil {
				// TODO: case of interface nexthop
			} else {
				var ptree *netutil.Ptree
				len_ := len(ri.Nexthop.IP)
				if len_ == 4 {
					ptree = v.ribTable[AFI_IP]
				} else if len_ == 16 {
					ptree = v.ribTable[AFI_IP6]
				}
				if ptree != nil {
					n := ptree.Match(ri.Nexthop.IP, len_*8)
					if n != nil {
						// XXX self reference.
						ri.SetResolved()
					}
				}
			}
		}
	} else {
		ri.SetResolved()
	}
}

func (v *Vrf) RibWalker(af int) {
	//fmt.Println("Vrf RibWalker:", v.Name)
	//v.Walker = time.Timer()
	//GetInstance().eventChan <- Event{}
	if af == AFI_IP || af == AFI_IP6 {
		ptree := v.ribTable[af]
		for n := ptree.Top(); n != nil; n = ptree.Next(n) {
			if n.Item != nil {
				var ip []byte
				if af == AFI_IP {
					ip = make([]byte, 4)
				} else if af == AFI_IP6 {
					ip = make([]byte, 16)
				}
				copy(ip, n.Key())
				p := netutil.PrefixFromIPPrefixlen(ip, n.KeyLength())
				ribs := n.Item.(RibSlice)
				v.RibProcess(p, ribs, nil, true)
			}
		}
	}
}

func RibWalker() {
	fmt.Println("RibWalker for all Vrf")
	//GetInstance().eventChan <- Event{}
	for _, vrf := range VrfMap {
		vrf.RibWalker(AFI_IP)
		vrf.RibWalker(AFI_IP6)
	}
}
