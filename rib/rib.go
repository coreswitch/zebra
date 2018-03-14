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

	"github.com/coreswitch/netutil"
	pb "github.com/coreswitch/zebra/proto"
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
	RIB_MAX
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
	RIB_FLAG_SELECTED RibFlag = 1 << iota
	RIB_FLAG_FIB
	RIB_FLAG_DISTANCE
	RIB_FLAG_METRIC
	RIB_FLAG_RESOLVED  // Static route and iBGP (multipath eBGP) next hop will be resolved.
	RIB_FLAG_BLACKHOLE // This is black hole route.
	RIB_FLAG_DELETE
)

type Rib struct {
	pb.Rib
	Flags    RibFlag
	Prefix   *netutil.Prefix
	Type     uint8
	SubType  uint8
	Distance uint8
	Metric   uint32
	PathId   uint32
	Nexthops []*Nexthop
	Color    []string
	Aux      []byte
	Src      interface{}
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

	return json.Marshal(ribJSON)
}

func (r *Rib) String() string {
	return fmt.Sprintf("%s [%d/%d]", r.Prefix, r.Distance, r.Metric)
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
	return r.CheckFlag(RIB_FLAG_FIB)
}

func (r *Rib) SetFib() {
	r.SetFlag(RIB_FLAG_FIB)
}

func (r *Rib) UnsetFib() {
	r.UnsetFlag(RIB_FLAG_FIB)
}

func (r *Rib) IsSelected() bool {
	return r.CheckFlag(RIB_FLAG_SELECTED)
}

func (r *Rib) SetSelected() {
	r.SetFlag(RIB_FLAG_SELECTED)
}

func (r *Rib) UnsetSelected() {
	r.UnsetFlag(RIB_FLAG_SELECTED)
}

func (r *Rib) IsSelectedFib() bool {
	return r.IsSelected() && r.IsFib()
}

func (r *Rib) IsSystem() bool {
	return r.Type == RIB_KERNEL || r.Type == RIB_CONNECTED
}

func (r *Rib) HasDistance() bool {
	return r.CheckFlag(RIB_FLAG_DISTANCE)
}

func (r *Rib) HasMetric() bool {
	return r.CheckFlag(RIB_FLAG_METRIC)
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

func (r *Rib) IsDelete() bool {
	return r.CheckFlag(RIB_FLAG_DELETE)
}

func (r *Rib) SetDelete() {
	r.SetFlag(RIB_FLAG_DELETE)
}

func (r *Rib) UnsetDelete() {
	r.UnsetFlag(RIB_FLAG_DELETE)
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
		Nexthops: ri.Nexthops,
		Metric:   ri.Metric,
		PathId:   ri.PathId,
		Src:      ri.Src,
		Aux:      ri.Aux,
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
		if rib.Src == ri.Src && len(rib.Nexthops) == 1 && len(ri.Nexthops) == 1 {
			nhops := rib.Nexthops[0]
			nhopi := ri.Nexthops[0]
			if nhops.Equal(nhopi) {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	case RIB_STATIC:
		return true
	case RIB_BGP:
		if rib.Src == ri.Src && rib.PathId == ri.PathId {
			return true
		} else {
			return false
		}
	default:
		// Other type of routes are considered as implicit withdraw.
		return true
	}
}

var AddPathDefault bool

// Current Selected (only one)
// Current FIB

// New Selected
// New FIB

func (v *Vrf) FibDelete(p *netutil.Prefix, rib *Rib) {
	if !rib.IsSystem() {
		NetlinkRouteDelete(p, rib, v.Id)
	}
	rib.UnsetFib()
}

func (v *Vrf) FibAdd(p *netutil.Prefix, rib *Rib) {
	if !rib.IsSystem() {
		NetlinkRouteAdd(p, rib, v.Id)
	}
	rib.SetFib()
}

func (v *Vrf) RibProcess(p *netutil.Prefix, ribs RibSlice, dels []*Rib, resolve bool) {
	// Find existing FIBs.
	var oFibs []*Rib
	var nFibs []*Rib
	var oSelected *Rib
	var nSelected *Rib

	var def *Rib
	for _, rib := range ribs {
		if rib.IsFib() {
			oFibs = append(oFibs, rib)
		}
		if rib.IsSelected() {
			oSelected = rib
		}
	}
	for _, del := range dels {
		if del.IsFib() {
			oFibs = append(oFibs, del)
		}
		if del.IsSelected() {
			oSelected = del
		}
	}

	// New FIBs and new selected.
	for _, rib := range ribs {
		if resolve {
			v.Resolve(rib)
		}
		if !rib.IsResolved() {
			continue
		}
		if AddPathDefault && p.IsDefault() && rib.Type == RIB_KERNEL {
			def = rib
			continue
		}

		if nSelected == nil {
			nSelected = rib
			nFibs = []*Rib{rib}
		} else {
			switch {
			case rib.Distance < nSelected.Distance:
				nSelected = rib
				nFibs = []*Rib{rib}
			case rib.Distance == nSelected.Distance:
				if rib.Metric < nSelected.Metric {
					nSelected = rib
				}
				nFibs = append(nFibs, rib)
			case rib.Distance > nSelected.Distance:
			}
		}
	}
	// Special default route add path treatment.
	if def != nil {
		if len(nFibs) > 0 {
			nSelected = nFibs[0]
		} else {
			nSelected = def
		}
		nFibs = append(nFibs, def)
	}

	// Old FIB to be removed.
	for _, ofib := range oFibs {
		keep := false
		for _, nfib := range nFibs {
			if ofib == nfib {
				keep = true
			}
		}
		if !keep {
			v.FibDelete(p, ofib)
		}
	}

	// New FIB to be added.
	for _, nfib := range nFibs {
		if !nfib.IsFib() {
			v.FibAdd(p, nfib)
		} else {
			// Redundant?
			if nfib.Type == RIB_STATIC || nfib.Type == RIB_BGP {
				v.FibAdd(p, nfib)
			}
		}
	}

	// Sync Selected for redistribute.
	if oSelected == nSelected {
		return
	}
	if oSelected != nil && nSelected != nil {
		if oSelected.Type == nSelected.Type {
			oSelected.UnsetSelected()
			nSelected.SetSelected()
			v.RedistAdd(p, nSelected)
			return
		}
	}
	if oSelected != nil {
		oSelected.UnsetSelected()
		v.RedistDelete(p, oSelected)
	}
	if nSelected != nil {
		nSelected.SetSelected()
		v.RedistAdd(p, nSelected)
	}
	return
}

func NexthopEncode(rib *Rib) []*pb.Nexthop {
	nhops := []*pb.Nexthop{}
	for _, nhop := range rib.Nexthops {
		nhops = append(nhops, &pb.Nexthop{
			Addr:    nhop.IP,
			Ifindex: uint32(nhop.Index),
		})
	}
	return nhops
}

func (vrf *Vrf) RedistSync(w Watcher, afi int, typ uint8) {
	ptree := vrf.ribTable[afi]
	for n := ptree.Top(); n != nil; n = ptree.Next(n) {
		if n.Item != nil {
			ribs := n.Item.(RibSlice)
			for _, rib := range ribs {
				if rib.IsSelectedFib() && rib.Type == typ {
					ip := make([]byte, 4)
					copy(ip, n.Key())
					p := netutil.PrefixFromIPPrefixlen(ip, n.KeyLength())
					if !p.IsDefault() {
						w.Notify(vrf.Rib2Route(pb.Op_RouteAdd, p, rib))
					}
				}
			}
		}
	}
}

func (vrf *Vrf) RedistDefaultSync(w Watcher, afi int) {
	ptree := vrf.ribTable[afi]
	p := netutil.NewPrefixAFI(afi)
	n := ptree.Lookup(p.IP, p.Length)
	if n == nil || n.Item == nil {
		return
	}
	for _, rib := range n.Item.(RibSlice) {
		if rib.IsSelectedFib() {
			if LocalPolicy && rib.Type == RIB_KERNEL {
				return
			}
			w.Notify(vrf.Rib2Route(pb.Op_RouteAdd, p, rib))
		}
	}
}

func (vrf *Vrf) Rib2Route(op pb.Op, p *netutil.Prefix, rib *Rib) *pb.Route {
	nhops := NexthopEncode(rib)
	return &pb.Route{
		Op:    op,
		VrfId: uint32(vrf.Id),
		Prefix: &pb.Prefix{
			Addr:   p.IP,
			Length: uint32(p.Length),
		},
		Type:     pb.RouteType(rib.Type),
		SubType:  pb.RouteSubType(rib.SubType),
		Distance: uint32(rib.Distance),
		Metric:   rib.Metric,
		Tag:      rib.Tag,
		Nexthops: nhops,
		Color:    rib.Color,
		Aux:      rib.Aux,
	}
}

func (v *Vrf) Redistribute(op pb.Op, p *netutil.Prefix, rib *Rib) {
	if LocalPolicy {
		if p.IsDefault() && rib.Type == RIB_KERNEL {
			return
		}
	}
	notifyFunc := func(wr WatcherRedist) {
		var watchers Watchers
		if p.IsDefault() {
			watchers = wr.def
		} else {
			watchers = wr.typ[rib.Type]
		}
		for _, w := range watchers {
			if rib.Src != w {
				w.Notify(v.Rib2Route(op, p, rib))
			}
		}
	}
	notifyFunc(Redist[p.AFI()])
	notifyFunc(v.redist[p.AFI()])
}

func (v *Vrf) RedistAdd(p *netutil.Prefix, rib *Rib) {
	v.Redistribute(pb.Op_RouteAdd, p, rib)
}

func (v *Vrf) RedistDelete(p *netutil.Prefix, rib *Rib) {
	v.Redistribute(pb.Op_RouteDelete, p, rib)
}

func (v *Vrf) RibAdd(p *netutil.Prefix, ri *Rib) {
	v.Mutex.Lock()
	defer v.Mutex.Unlock()
	p.ApplyMask()
	ptree := v.AfiPtree(p)
	n := ptree.Acquire(p.IP, p.Length)

	// Resolve nexthop.
	v.Resolve(ri)

	// Updated ribs.
	var dels []*Rib
	var ribs RibSlice

	// Check is this rib same source and same nexthop route.
	if n.Item != nil {
		for _, rib := range n.Item.(RibSlice) {
			if ri.Type == rib.Type {
				if ri.Src == rib.Src {
					if len(rib.Nexthops) == 1 && len(ri.Nexthops) == 1 {
						nhops := rib.Nexthops[0]
						nhopi := ri.Nexthops[0]
						if nhops.Equal(nhopi) {
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
				dels = append(dels, rib)
			} else {
				ribs = append(ribs, rib)
			}
		}
	}

	// Replace of the RIB.
	if len(dels) != 0 {
		ptree.Release(n)
	}

	// Append a new RIB.
	rib := NewRib(p, ri)
	ribs = append(ribs, rib)
	n.Item = ribs

	// Process the rib.
	v.RibProcess(p, ribs, dels, false)

	// Invoke walker.
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

	// Process RIB.
	v.RibProcess(p, ribs, []*Rib{found}, false)

	// Invoke walker.
	v.RibWalker(p.AFI())
}

func (v *Vrf) RibClean(src interface{}) {
	v.Mutex.Lock()
	defer v.Mutex.Unlock()
	for afi := AFI_IP; afi < AFI_MAX; afi++ {
		ptree := v.ribTable[afi]
		for n := ptree.Top(); n != nil; n = ptree.Next(n) {
			var dels []*Rib
			var ribs RibSlice

			for _, rib := range n.Item.(RibSlice) {
				if rib.Src == src {
					dels = append(dels, rib)
				} else {
					ribs = append(ribs, rib)
				}
			}
			if len(dels) == 0 {
				continue
			}
			if len(ribs) == 0 {
				n.Item = nil
			} else {
				n.Item = ribs
			}
			for num, _ := range dels {
				fmt.Println("RibClean ptree release", num)
				ptree.Release(n)
			}
			var ip []byte
			if afi == AFI_IP {
				ip = make([]byte, 4)
			} else if afi == AFI_IP6 {
				ip = make([]byte, 16)
			}
			copy(ip, n.Key())
			p := netutil.PrefixFromIPPrefixlen(ip, n.KeyLength())

			v.RibProcess(p, ribs, dels, false)
		}
	}
}

func (v *Vrf) IsValid(rib *Rib) bool {
	if rib.Type == RIB_KERNEL {
		if len(rib.Nexthops) != 1 {
			return false
		}
		ifp := v.IfLookupByIndex(rib.Nexthops[0].Index)
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

func RibAdd(vrfId uint32, p *netutil.Prefix, ri *Rib) {
	if vrf := VrfLookupByIndex(vrfId); vrf != nil {
		vrf.RibAdd(p, ri)
	}
}

func RibDelete(vrfId uint32, p *netutil.Prefix, ri *Rib) {
	if vrf := VrfLookupByIndex(vrfId); vrf != nil {
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
		if len(ri.Nexthops) == 1 {
			nhop := ri.Nexthops[0]
			//fmt.Println("resolving nexthop", ri.Nexthop)
			if nhop.IP == nil {
				// TODO: case of interface nexthop
			} else {
				var ptree *netutil.Ptree
				len_ := len(nhop.IP)
				if len_ == 4 {
					ptree = v.ribTable[AFI_IP]
				} else if len_ == 16 {
					ptree = v.ribTable[AFI_IP6]
				}
				if ptree != nil {
					n := ptree.Match(nhop.IP, len_*8)
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
