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
	"os"
	"os/exec"
	"regexp"
	"strings"
	"syscall"
	"unsafe"

	"github.com/coreswitch/netutil"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
	"github.com/vishvananda/netns"
)

var native = nl.NativeEndian()

func ARPHRD2IfType(ARPHRD uint16) uint8 {
	switch ARPHRD {
	case syscall.ARPHRD_ETHER:
		return IF_TYPE_ETHERNET
	case syscall.ARPHRD_LOOPBACK:
		return IF_TYPE_LOOPBACK
	default:
		return IF_TYPE_UNKNOWN
	}
}

func ifInfoDeserialize(m syscall.NetlinkMessage) (*IfInfo, error) {
	ifmsg := nl.DeserializeIfInfomsg(m.Data)

	attrs, err := nl.ParseRouteAttr(m.Data[ifmsg.Len():])
	if err != nil {
		return nil, err
	}

	ifi := &IfInfo{
		MsgType: m.Header.Type,
		IfType:  ARPHRD2IfType(ifmsg.Type),
		Index:   uint32(ifmsg.Index),
		Flags:   ifmsg.Flags,
	}

	for _, attr := range attrs {
		switch attr.Attr.Type {
		case syscall.IFLA_IFNAME:
			ifi.Name = string(attr.Value[:len(attr.Value)-1])
		case syscall.IFLA_MTU:
			ifi.Mtu = native.Uint32(attr.Value[0:4])
		case syscall.IFLA_LINK:
			//base.ParentIndex = int(native.Uint32(attr.Value[0:4]))
		case syscall.IFLA_MASTER:
			ifi.Master = int(native.Uint32(attr.Value[0:4]))
			//fmt.Println("Master:", ifi.Master)
		case syscall.IFLA_TXQLEN:
			//base.TxQLen = int(native.Uint32(attr.Value[0:4]))
		case syscall.IFLA_IFALIAS:
			//base.Alias = string(attr.Value[:len(attr.Value)-1])
		case syscall.IFLA_STATS:
			//base.Statistics = parseLinkStats(attr.Value[:])
		case syscall.IFLA_LINKINFO:
			linkType := ""
			infos, err := nl.ParseRouteAttr(attr.Value)
			if err != nil {
				return nil, err
			}
			for _, info := range infos {
				switch info.Attr.Type {
				case nl.IFLA_INFO_KIND:
					linkType = string(info.Value[:len(info.Value)-1])
				case nl.IFLA_INFO_DATA:
					data, err := nl.ParseRouteAttr(info.Value)
					if err != nil {
						return nil, err
					}
					switch linkType {
					case "vrf":
						for _, datum := range data {
							switch datum.Attr.Type {
							case nl.IFLA_VRF_TABLE:
								ifi.Table = int(native.Uint32(datum.Value[0:4]))
								//fmt.Println("Vrf binding", ifi.Table)
							default:
							}
						}
					default:
					}
				default:
				}
			}
		case syscall.IFLA_ADDRESS:
			var nonzero bool
			for _, b := range attr.Value {
				if b != 0 {
					nonzero = true
				}
			}
			if nonzero {
				ifi.HwAddr = attr.Value[:]
			}
		}
	}
	return ifi, nil
}

func ifAddrDeserialize(m syscall.NetlinkMessage) (*IfAddrInfo, error) {
	msg := nl.DeserializeIfAddrmsg(m.Data)

	attrs, err := nl.ParseRouteAttr(m.Data[msg.Len():])
	if err != nil {
		return nil, err
	}

	ifaddr := &IfAddrInfo{MsgType: m.Header.Type, Index: IfIndex(msg.Index), Family: int(msg.Family)}
	var local, dst *netutil.Prefix
	for _, attr := range attrs {
		switch attr.Attr.Type {
		case syscall.IFA_ADDRESS:
			dst = netutil.PrefixFromIPPrefixlen(attr.Value, int(msg.Prefixlen))
		case syscall.IFA_LOCAL:
			local = netutil.PrefixFromIPPrefixlen(attr.Value, int(msg.Prefixlen))
		case syscall.IFA_LABEL:
			ifaddr.Label = string(attr.Value[:len(attr.Value)-1])
		case netlink.IFA_FLAGS:
			//ifaddr.Flags = int(native.Uint32(attr.Value[0:4]))
		}
	}
	// IFA_LOCAL should be there but if not, fall back to IFA_ADDRESS
	if local != nil {
		ifaddr.Prefix = local
	} else {
		ifaddr.Prefix = dst
	}
	return ifaddr, nil
}

// Route represents a netlink route.
type RouteInfo struct {
	MsgType uint16
	Rib
	Table     int
	MultiPath []*NexthopInfo
}

func (route RouteInfo) String() string {
	strs := []string{}
	strs = append(strs, fmt.Sprintf("%s", route.Rib.Prefix))
	if route.Nexthop != nil {
		switch route.Nexthop.EncapType {
		case nl.LWTUNNEL_ENCAP_SEG6:
			strs = append(strs, fmt.Sprintf("encap seg6 %s", route.Nexthop.EncapSeg6.String()))
		}
	}
	return fmt.Sprintf("%s", strings.Join(strs, " "))
	//return route.Prefix.String() + " " + route.Rib.String()
}

type NexthopInfo struct {
	LinkIndex int
	Hops      int
	Gateway   net.IP
	EncapType int
	EncapSeg6 EncapSEG6
}

func deserializeRoute(m syscall.NetlinkMessage) (*RouteInfo, error) {
	msg := nl.DeserializeRtMsg(m.Data)

	if msg.Protocol == syscall.RTPROT_REDIRECT {
		//fmt.Println("RTPROT_REDIRECT")
		return nil, nil
	}

	if msg.Type != syscall.RTN_UNICAST {
		//fmt.Println("RTN_UNICAST", route.Dest)
		return nil, nil
	}
	if msg.Flags&syscall.RTM_F_CLONED != 0 {
		//fmt.Println("RTM_F_CLONED")
		return nil, nil
	}
	//fmt.Println("Table:", msg.Table)

	attrs, err := nl.ParseRouteAttr(m.Data[msg.Len():])
	if err != nil {
		return nil, err
	}

	route := RouteInfo{MsgType: m.Header.Type, Rib: Rib{Type: RIB_KERNEL}, Table: int(msg.Table)}
	nexthop := new(Nexthop)
	var encap, encapType syscall.NetlinkRouteAttr
	for _, attr := range attrs {
		switch attr.Attr.Type {
		case syscall.RTA_GATEWAY:
			nexthop.IP = net.IP(attr.Value)
		case syscall.RTA_PREFSRC:
			//route.Src = net.IP(attr.Value)
		case syscall.RTA_DST:
			route.Prefix = netutil.PrefixFromIPPrefixlen(attr.Value, int(msg.Dst_len))
		case syscall.RTA_OIF:
			//route.LinkIndex = int(native.Uint32(attr.Value[0:4]))
			nexthop.Index = IfIndex(native.Uint32(attr.Value[0:4]))
		case syscall.RTA_IIF:
			//route.ILinkIndex = int(native.Uint32(attr.Value[0:4]))
		case syscall.RTA_PRIORITY:
			//route.Priority = int(native.Uint32(attr.Value[0:4]))
		case syscall.RTA_TABLE:
			route.Table = int(native.Uint32(attr.Value[0:4]))
		case syscall.RTA_MULTIPATH:
			parseRtNexthop := func(value []byte) (*NexthopInfo, []byte, error) {
				if len(value) < syscall.SizeofRtNexthop {
					return nil, nil, fmt.Errorf("Lack of bytes")
				}
				nh := nl.DeserializeRtNexthop(value)
				if len(value) < int(nh.RtNexthop.Len) {
					return nil, nil, fmt.Errorf("Lack of bytes")
				}
				info := &NexthopInfo{
					LinkIndex: int(nh.RtNexthop.Ifindex),
					Hops:      int(nh.RtNexthop.Hops),
				}
				attrs, err := nl.ParseRouteAttr(value[syscall.SizeofRtNexthop:int(nh.RtNexthop.Len)])
				if err != nil {
					return nil, nil, err
				}
				var encap, encapType syscall.NetlinkRouteAttr
				for _, attr := range attrs {
					switch attr.Attr.Type {
					case syscall.RTA_GATEWAY:
						info.Gateway = net.IP(attr.Value)
					case nl.RTA_ENCAP_TYPE:
						encapType = attr
					case nl.RTA_ENCAP:
						encap = attr
					}
				}
				if len(encap.Value) != 0 && len(encapType.Value) != 0 {
					typ := int(native.Uint16(encapType.Value[0:2]))
					switch typ {
					// List more LWTUNNEL_ENCAP_XXX here
					case nl.LWTUNNEL_ENCAP_SEG6:
						seg6 := &netlink.SEG6Encap{}
						if err := seg6.Decode(encap.Value); err != nil {
							fmt.Println("ERROR: failed to Decode seg6 RTA")
							return nil, nil, err
						}
						info.EncapType = nl.LWTUNNEL_ENCAP_SEG6
						info.EncapSeg6.Mode = seg6.Mode
						info.EncapSeg6.Segments = seg6.Segments
					}
				}
				return info, value[int(nh.RtNexthop.Len):], nil
			}
			rest := attr.Value
			for len(rest) > 0 {
				info, buf, err := parseRtNexthop(rest)
				if err != nil {
					return nil, err
				}
				route.MultiPath = append(route.MultiPath, info)
				n := NewNexthopAddrIf(info.Gateway, IfIndex(info.LinkIndex))
				route.Nexthops = append(route.Nexthops, n)
				n.EncapType = info.EncapType
				n.EncapSeg6 = info.EncapSeg6
				rest = buf
			}
		case nl.RTA_ENCAP_TYPE:
			encapType = attr
		case nl.RTA_ENCAP:
			encap = attr
		}
	}

	if len(encap.Value) != 0 && len(encapType.Value) != 0 {
		typ := int(native.Uint16(encapType.Value[0:2]))
		switch typ {
		// List more LWTUNNEL_ENCAP_XXX here
		case nl.LWTUNNEL_ENCAP_SEG6:
			seg6 := &netlink.SEG6Encap{}
			if err := seg6.Decode(encap.Value); err != nil {
				return nil, err
			}
			nexthop.EncapType = nl.LWTUNNEL_ENCAP_SEG6
			nexthop.EncapSeg6.Mode = seg6.Mode
			nexthop.EncapSeg6.Segments = seg6.Segments
		}
	}

	if msg.Protocol == syscall.RTPROT_KERNEL {
		fmt.Println("RTPROT_KERNEL", route)
		return nil, nil
	}
	if msg.Protocol == syscall.RTPROT_ZEBRA {
		fmt.Println("RTPROT_ZEBRA", route)
		return nil, nil
	}

	// Ajust main table id to 0.
	if route.Table == syscall.RT_TABLE_MAIN {
		route.Table = 0
	}

	// Empty prefix means default route, create instance at here.
	if route.Prefix == nil {
		if msg.Family == syscall.AF_INET {
			route.Prefix = netutil.NewPrefixAFI(AFI_IP)
		} else {
			route.Prefix = netutil.NewPrefixAFI(AFI_IP6)
		}
	}

	// Skip multicast route.
	if route.Prefix.IsMulticast() {
		return nil, nil
	}

	// Make nexthop information.
	if route.MultiPath == nil {
		route.Nexthop = nexthop
	}

	return &route, nil
}

// VRF binding require two phase parse.
var linkResolvedMap = map[int]*IfInfo{}
var linkUnresolvedMap = map[int]*IfInfo{}

func ifMsgParse(m syscall.NetlinkMessage) error {
	ifi, err := ifInfoDeserialize(m)
	if err != nil {
		return err
	}

	ifi.Boot = true

	if ifi.Master != 0 {
		if _, ok := linkResolvedMap[ifi.Master]; !ok {
			linkUnresolvedMap[int(ifi.Index)] = ifi
			return nil
		}
	}
	linkResolvedMap[int(ifi.Index)] = ifi

	if ifi.MsgType == syscall.RTM_NEWLINK {
		fmt.Println("If add (boot):", ifi)
		IfUpdate(ifi)
	} else {
		fmt.Println("Del if (boot):", ifi)
		IfDelete(ifi)
	}
	return nil
}

func ifMsgSync() {
	for _, ifi := range linkUnresolvedMap {
		if ifi.MsgType == syscall.RTM_NEWLINK {
			fmt.Println("If add (boot):", ifi)
			IfUpdate(ifi)
		} else {
			fmt.Println("Del if (boot):", ifi)
			IfDelete(ifi)
		}
	}
}

func ifAddrMsgParse(m syscall.NetlinkMessage) error {
	ifaddr, err := ifAddrDeserialize(m)
	if err != nil {
		return err
	}
	ifaddr.Flags = ifaddr.Flags.SetFlag(IFADDR_SOURCE_SYSTEM)
	if ifaddr.MsgType == syscall.RTM_NEWADDR {
		fmt.Println("Addr add (boot):", ifaddr)
		IfAddrAdd(ifaddr)
	} else {
		fmt.Println("Addr del (boot):", ifaddr)
		IfAddrDelete(ifaddr)
	}
	return nil
}

func routeMsgParse(m syscall.NetlinkMessage) error {
	route, err := deserializeRoute(m)
	if err != nil {
		return err
	}
	if route == nil {
		return nil
	}
	if route.MsgType == syscall.RTM_NEWROUTE {
		fmt.Println("Route add (boot):", route, route.Table)
		RibAdd(route.Table, route.Prefix, &route.Rib)
	} else {
		fmt.Println("Route del (boot):", route, route.Table)
		RibDelete(route.Table, route.Prefix, &route.Rib)
	}
	return nil
}

const (
	NDA_UNSPEC = iota
	NDA_DST
	NDA_LLADDR
	NDA_CACHEINFO
	NDA_PROBES
	NDA_VLAN
	NDA_PORT
	NDA_VNI
	NDA_IFINDEX
	NDA_MAX = NDA_IFINDEX
)

// Neighbor Cache Entry States.
const (
	NUD_NONE       = 0x00
	NUD_INCOMPLETE = 0x01
	NUD_REACHABLE  = 0x02
	NUD_STALE      = 0x04
	NUD_DELAY      = 0x08
	NUD_PROBE      = 0x10
	NUD_FAILED     = 0x20
	NUD_NOARP      = 0x40
	NUD_PERMANENT  = 0x80
)

// Neighbor Flags
const (
	NTF_USE    = 0x01
	NTF_SELF   = 0x02
	NTF_MASTER = 0x04
	NTF_PROXY  = 0x08
	NTF_ROUTER = 0x80
)

type Ndmsg struct {
	Family uint8
	Index  uint32
	State  uint16
	Flags  uint8
	Type   uint8
}

func (msg *Ndmsg) Len() int {
	return int(unsafe.Sizeof(*msg))
}

func NeighStateString(state uint16) string {
	str := ""
	if (state & NUD_INCOMPLETE) != 0 {
		str += "incomplete "
	}
	if (state & NUD_REACHABLE) != 0 {
		str += "reachable "
	}
	if (state & NUD_STALE) != 0 {
		str += "stale "
	}
	if (state & NUD_DELAY) != 0 {
		str += "delay "
	}
	if (state & NUD_PROBE) != 0 {
		str += "probe "
	}
	if (state & NUD_FAILED) != 0 {
		str += "failed "
	}
	if (state & NUD_NOARP) != 0 {
		str += "noarp "
	}
	if (state & NUD_PERMANENT) != 0 {
		str += "permanent "
	}

	return str
}

func NeighFlagString(flag uint8) string {
	str := ""

	if (flag & NTF_USE) != 0 {
		str += "use "
	}
	if (flag & NTF_SELF) != 0 {
		str += "self "
	}
	if (flag & NTF_MASTER) != 0 {
		str += "master "
	}
	if (flag & NTF_PROXY) != 0 {
		str += "proxy "
	}
	if (flag & NTF_ROUTER) != 0 {
		str += "router "
	}

	return str
}

// Neigh represents a link layer neighbor from netlink.
type Neigh struct {
	LinkIndex    int
	Family       int
	State        uint16
	Type         int
	Flags        uint8
	IP           net.IP
	HardwareAddr net.HardwareAddr
}

func (n *Neigh) IsReachable() bool {
	if (n.Flags & NUD_FAILED) != 0 {
		return false
	}
	if (n.Flags & NUD_INCOMPLETE) != 0 {
		return false
	}
	if len(n.HardwareAddr) == 0 {
		return false
	}
	return true
}

func deserializeNdmsg(b []byte) *Ndmsg {
	var dummy Ndmsg
	return (*Ndmsg)(unsafe.Pointer(&b[0:unsafe.Sizeof(dummy)][0]))
}

func deserializeNeigh(m syscall.NetlinkMessage) (*Neigh, error) {
	msg := deserializeNdmsg(m.Data)
	neigh := &Neigh{
		LinkIndex: int(msg.Index),
		Family:    int(msg.Family),
		State:     msg.State,
		Type:      int(msg.Type),
		Flags:     msg.Flags,
	}

	attrs, err := nl.ParseRouteAttr(m.Data[msg.Len():])
	if err != nil {
		return nil, err
	}

	for _, attr := range attrs {
		switch attr.Attr.Type {
		case NDA_DST:
			neigh.IP = net.IP(attr.Value)
		case NDA_LLADDR:
			neigh.HardwareAddr = net.HardwareAddr(attr.Value)
		}
	}

	// Filter none.
	if neigh.State == NUD_NONE {
		return nil, nil
	}
	// Filter noarp.
	if (neigh.State & NUD_NOARP) != 0 {
		return nil, nil
	}

	//fmt.Println(neigh.IP, len(neigh.IP), neigh.HardwareAddr, neigh.Flags, NeighFlagString(neigh.Flags), neigh.State, NeighStateString(neigh.State))

	return neigh, nil
}

func neighMsgParse(m syscall.NetlinkMessage) error {
	neigh, err := deserializeNeigh(m)
	if err != nil {
		return err
	}
	if neigh == nil {
		return nil
	}
	// flag is failed or len(neigh.HardwareAddr) == 0 -> Delete
	if neigh.IsReachable() {
		// vrf.ArpAdd(neigh)
	} else {
		// vrf.ArpDelete(neigh)
	}
	return nil
}

func linkSubscribe(inst *Server, newNs, curNs netns.NsHandle, done <-chan struct{}) error {
	s, err := nl.SubscribeAt(newNs, curNs, syscall.NETLINK_ROUTE, syscall.RTNLGRP_LINK)
	if err != nil {
		return err
	}
	if done != nil {
		go func() {
			<-done
			s.Close()
		}()
	}
	go func() {
		defer close(inst.ifChan)
		for {
			msgs, err := s.Receive()
			if err != nil {
				return
			}
			for _, m := range msgs {
				ifi, err := ifInfoDeserialize(m)
				if err != nil {
					return
				}
				inst.ifChan <- *ifi
			}
		}
	}()

	return nil
}

func addrSubscribe(inst *Server, newNs, curNs netns.NsHandle, done <-chan struct{}) error {
	s, err := nl.SubscribeAt(newNs, curNs, syscall.NETLINK_ROUTE, syscall.RTNLGRP_IPV4_IFADDR, syscall.RTNLGRP_IPV6_IFADDR)
	if err != nil {
		return err
	}
	if done != nil {
		go func() {
			<-done
			s.Close()
		}()
	}
	go func() {
		defer close(inst.ifaddrChan)
		for {
			msgs, err := s.Receive()
			if err != nil {
				fmt.Printf("AddrSubscribe: Receive() error: %v", err)
				return
			}
			for _, m := range msgs {
				msgType := m.Header.Type
				if msgType != syscall.RTM_NEWADDR && msgType != syscall.RTM_DELADDR {
					fmt.Printf("AddrSubscribe: bad message type: %d", msgType)
					continue
				}
				ai, err := ifAddrDeserialize(m)
				if err != nil {
					fmt.Printf("Addr infor parse error")
					continue
				}
				inst.ifaddrChan <- *ai
			}
		}
	}()

	return nil
}

func routeSubscribe(inst *Server, newNs, curNs netns.NsHandle, done <-chan struct{}) error {
	s, err := nl.SubscribeAt(newNs, curNs, syscall.NETLINK_ROUTE, syscall.RTNLGRP_IPV4_ROUTE, syscall.RTNLGRP_IPV6_ROUTE)
	if err != nil {
		return err
	}
	if done != nil {
		go func() {
			<-done
			s.Close()
		}()
	}
	go func() {
		defer close(inst.routeChan)
		for {
			msgs, err := s.Receive()
			if err != nil {
				return
			}
			for _, m := range msgs {
				route, err := deserializeRoute(m)
				if err != nil {
					continue
				}
				if route == nil {
					continue
				}
				fmt.Println("Route", *route)
				inst.routeChan <- *route
			}
		}
	}()
	return nil
}

func neighSubscribe(inst *Server, newNs, curNs netns.NsHandle, done <-chan struct{}) error {
	s, err := nl.SubscribeAt(newNs, curNs, syscall.NETLINK_ROUTE, syscall.RTNLGRP_NEIGH)
	if err != nil {
		return err
	}
	if done != nil {
		go func() {
			<-done
			s.Close()
		}()
	}
	go func() {
		defer close(inst.routeChan)
		for {
			msgs, err := s.Receive()
			if err != nil {
				return
			}
			for _, m := range msgs {
				neigh, err := deserializeNeigh(m)
				if err != nil {
					continue
				}
				if neigh == nil {
					continue
				}
				if neigh.IsReachable() {
					//v.ArpAdd(neigh)
				} else {
					//v.ArpDelete(neigh)
				}
			}
		}
	}()
	return nil
}

type callbackFunc func(syscall.NetlinkMessage) error

func netlinkDump(s *nl.NetlinkSocket, proto int, family int, callback callbackFunc) error {
	req := nl.NewNetlinkRequest(proto, syscall.NLM_F_DUMP)
	msg := nl.NewIfInfomsg(family)
	req.AddData(msg)

	err := netlinkExec(s, req, callback)

	return err
}

func netlinkExec(s *nl.NetlinkSocket, req *nl.NetlinkRequest, callback callbackFunc) error {
	if err := s.Send(req); err != nil {
		return err
	}
	pid, err := s.GetPid()
	if err != nil {
		return err
	}
done:
	for {
		msgs, err := s.Receive()
		if err != nil {
			return err
		}
		for _, m := range msgs {
			if m.Header.Seq != req.Seq {
				return fmt.Errorf("Wrong Seq number %d, expected %d", m.Header.Seq, req.Seq)
			}
			if m.Header.Pid != pid {
				return fmt.Errorf("Wrong pid %d, expected %d", m.Header.Pid, pid)
			}
			if m.Header.Type == syscall.NLMSG_DONE {
				break done
			}
			if m.Header.Type == syscall.NLMSG_ERROR {
				errno := int32(native.Uint32(m.Data[0:4]))
				if errno == 0 {
					break done
				}
				return syscall.Errno(-errno)
			}
			// if resType != 0 && m.Header.Type != resType {
			// 	continue
			// }
			err = callback(m)
			if err != nil {
				continue
			}

			if m.Header.Flags&syscall.NLM_F_MULTI == 0 {
				break done
			}
		}
	}
	return nil
}

var NetlinkDoneHook func()

func EsiNetlinkDoneHook() {
	fmt.Println("EsiNetlinkDoneHook")
	regex, err := regexp.Compile(`lan\-\d+|ens\d+`)
	if err != nil {
		fmt.Println("Regex compile error:", err)
		return
	}

	for _, ifp := range IfMap {
		if regex.MatchString(ifp.Name) {
			fmt.Println("Flushing interface", ifp.Name)
			err := exec.Command("ip", "addr", "flush", "dev", ifp.Name).Run()
			if err != nil {
				fmt.Println("Flushing interface err", err)
			}
		}
	}
}

func NetlinkDumpAndSubscribe(inst *Server) error {
	nl, err := nl.GetNetlinkSocketAt(netns.None(), netns.None(), syscall.NETLINK_ROUTE)
	if err != nil {
		return err
	}
	type dumpList struct {
		protocol int
		family   int
		callback callbackFunc
	}
	for _, l := range []dumpList{
		{syscall.RTM_GETLINK, syscall.AF_UNSPEC, ifMsgParse},
	} {
		err = netlinkDump(nl, l.protocol, l.family, l.callback)
		if err != nil {
			fmt.Println(err)
		}
	}
	ifMsgSync()

	for _, l := range []dumpList{
		{syscall.RTM_GETADDR, syscall.AF_INET, ifAddrMsgParse},
		{syscall.RTM_GETADDR, syscall.AF_INET6, ifAddrMsgParse},
		{syscall.RTM_GETROUTE, syscall.AF_INET, routeMsgParse},
		{syscall.RTM_GETROUTE, syscall.AF_INET6, routeMsgParse},
		{syscall.RTM_GETNEIGH, syscall.AF_INET, neighMsgParse},
		{syscall.RTM_GETNEIGH, syscall.AF_INET6, neighMsgParse},
	} {
		err = netlinkDump(nl, l.protocol, l.family, l.callback)
		if err != nil {
			fmt.Println(err)
		}
	}

	type subscribeFunc func(inst *Server, newNs, curNs netns.NsHandle, done <-chan struct{}) error

	for _, s := range []subscribeFunc{
		linkSubscribe,
		addrSubscribe,
		routeSubscribe,
		neighSubscribe,
	} {
		err := s(inst, netns.None(), netns.None(), nil)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("Netlink boot dump finished")

	if NetlinkDoneHook != nil {
		NetlinkDoneHook()
	}

	VrfDefaultZservStart()

	return nil
}

func NetlinkRouteAdd(p *netutil.Prefix, rib *Rib, vrfId int) error {
	fmt.Printf("NetlinkRouteAdd(): %s/%d\n", p.IP, p.Length)
	len_ := len(p.IP) * 8
	route := &netlink.Route{
		Dst: &net.IPNet{
			IP:   p.IP,
			Mask: net.CIDRMask(p.Length, len_),
		},
		Protocol: syscall.RTPROT_ZEBRA,
	}
	if vrfId != 0 {
		route.Table = vrfId
	}
	if rib.Nexthop != nil {
		route.Gw = rib.Nexthop.IP
		route.LinkIndex = int(rib.Nexthop.Index)
		switch rib.Nexthop.EncapType {
		case nl.LWTUNNEL_ENCAP_SEG6:
			seg6 := &netlink.SEG6Encap{}
			seg6.Mode = rib.Nexthop.EncapSeg6.Mode
			seg6.Segments = rib.Nexthop.EncapSeg6.Segments
			route.Encap = seg6
		}
	} else {
		var multiPath []*netlink.NexthopInfo
		for _, nexthop := range rib.Nexthops {
			multiPath = append(multiPath, &netlink.NexthopInfo{LinkIndex: int(nexthop.Index), Gw: nexthop.IP, Hops: 0})
		}
		route.MultiPath = multiPath
	}
	err := netlink.RouteAdd(route)
	if err != nil {
		return err
	}
	return nil
}

func NetlinkRouteDelete(p *netutil.Prefix, rib *Rib, vrfId int) error {
	fmt.Printf("NetlinkRouteDelete(): %s/%d\n", p.IP, p.Length)
	len_ := len(p.IP) * 8
	route := &netlink.Route{
		Dst: &net.IPNet{
			IP:   p.IP,
			Mask: net.CIDRMask(p.Length, len_),
		},
		Protocol: syscall.RTPROT_ZEBRA,
	}
	if vrfId != 0 {
		route.Table = vrfId
	}
	if rib.Nexthop != nil {
		route.Gw = rib.Nexthop.IP
		route.LinkIndex = int(rib.Nexthop.Index)
		switch rib.Nexthop.EncapType {
		case nl.LWTUNNEL_ENCAP_SEG6:
			seg6 := &netlink.SEG6Encap{}
			seg6.Mode = rib.Nexthop.EncapSeg6.Mode
			seg6.Segments = rib.Nexthop.EncapSeg6.Segments
			route.Encap = seg6
		}
	} else {
		var multiPath []*netlink.NexthopInfo
		for _, nexthop := range rib.Nexthops {
			multiPath = append(multiPath, &netlink.NexthopInfo{LinkIndex: int(nexthop.Index), Gw: nexthop.IP, Hops: 0})
		}
		route.MultiPath = multiPath
	}

	err := netlink.RouteDel(route)
	if err != nil {
		return err
	}
	return nil
}

const (
	SizeOfIfReq = 40
	IFNAMSIZ    = 16
	TUN_DEV     = "/dev/net/tun"
)

type ifReq struct {
	Name  [IFNAMSIZ]byte
	Flags uint16
	pad   [SizeOfIfReq - IFNAMSIZ - 2]byte
}

func TapAdd(name string) (*os.File, error) {
	file, err := os.OpenFile(TUN_DEV, os.O_RDWR, 0)
	if err != nil {
		return nil, fmt.Errorf("Can't open tunnel device %s, err %v", TUN_DEV, err)
	}

	var req ifReq
	req.Flags = uint16(syscall.IFF_TAP | syscall.IFF_NO_PI)
	copy(req.Name[:15], name)
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(&req)))
	if errno != 0 {
		return nil, fmt.Errorf("Tuntap ioctl() TUNSETIFF failed, errno %v", errno)
	}

	// Do not make persistent for now.
	persistent := false
	if persistent {
		_, _, errno = syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETPERSIST), 1)
		if errno != 0 {
			return nil, fmt.Errorf("Tuntap ioctl() TUNSETPERSIST failed, errno %v", errno)
		}
	}

	return file, nil
}

func TapDelete(name string) error {
	ifp := VrfDefault().IfLookupByName(name)
	if ifp == nil {
		return fmt.Errorf("Can't find tap interface name: %s", name)
	}
	link := &netlink.Dummy{netlink.LinkAttrs{Index: int(ifp.Index)}}
	err := netlink.LinkDel(link)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func HardwareAddrSet(ifp *Interface, hwaddr net.HardwareAddr) error {
	link := &netlink.Dummy{netlink.LinkAttrs{Index: int(ifp.Index)}}
	err := netlink.LinkSetHardwareAddr(link, hwaddr)
	if err != nil {
		return err
	}
	return nil
}

// Netlink VRF
const (
	RULE_SELECTOR_IIF = iota
	RULE_SELECTOR_OIF
)

func ruleGet(name string, table int, family int, selector int) *netlink.Rule {
	rule := netlink.NewRule()
	rule.Table = table
	rule.Family = family
	switch selector {
	case RULE_SELECTOR_IIF:
		rule.IifName = name
	case RULE_SELECTOR_OIF:
		rule.OifName = name
	}
	return rule
}

func ruleAdd(name string, table int, family int, selector int) error {
	rule := ruleGet(name, table, family, selector)
	err := netlink.RuleAdd(rule)
	if err != nil {
		return err
	}
	return nil
}

func ruleDelete(name string, table int, family int, selector int) error {
	rule := ruleGet(name, table, family, selector)
	err := netlink.RuleDel(rule)
	if err != nil {
		return err
	}
	return nil
}

func linkVrfAdd(name string, table int) error {
	err := netlink.LinkAdd(&netlink.Vrf{
		LinkAttrs: netlink.LinkAttrs{Name: name},
		Table:     uint32(table),
	})
	return err
}

func linkVrfDelete(name string) error {
	err := netlink.LinkDel(&netlink.Dummy{
		LinkAttrs: netlink.LinkAttrs{Name: name},
	})
	return err
}

func NetlinkVrfAdd(name string, table int) {
	linkVrfAdd(name, table)
	// Below is not necessary Linux 4.8 and later.
	ruleAdd(name, table, syscall.AF_INET, RULE_SELECTOR_IIF)
	ruleAdd(name, table, syscall.AF_INET, RULE_SELECTOR_OIF)
	ruleAdd(name, table, syscall.AF_INET6, RULE_SELECTOR_IIF)
	ruleAdd(name, table, syscall.AF_INET6, RULE_SELECTOR_OIF)
}

func NetlinkVrfDelete(name string, table int) {
	// Below is not necessary Linux 4.8 and later.
	ruleDelete(name, table, syscall.AF_INET, RULE_SELECTOR_OIF)
	ruleDelete(name, table, syscall.AF_INET, RULE_SELECTOR_IIF)
	ruleDelete(name, table, syscall.AF_INET6, RULE_SELECTOR_OIF)
	ruleDelete(name, table, syscall.AF_INET6, RULE_SELECTOR_IIF)
	linkVrfDelete(name)
}

func NetlinkVlanAdd(name string, vlanId int, parentIndex int) {
	fmt.Println("[netlink]NetlinkVlanAdd:", name, vlanId, parentIndex)
	err := netlink.LinkAdd(&netlink.Vlan{netlink.LinkAttrs{Name: name, ParentIndex: parentIndex}, vlanId})
	if err != nil {
		fmt.Println("[netlink]NetlinkVlanAdd:", err)
	}
}

func NetlinkVlanDelete(name string, vlanId int) {
	fmt.Println("[netlink]NetlinkVlanDelete:", name, vlanId)
	err := netlink.LinkDel(&netlink.Dummy{
		LinkAttrs: netlink.LinkAttrs{Name: name},
	})
	if err != nil {
		fmt.Println("[netlink]NetlinkVlanDelete error:", err)
	}
}

func NetlinkVrfBindInterface(ifname string, ifindex IfIndex, master IfIndex) {
	link := &netlink.Dummy{
		netlink.LinkAttrs{
			Name:  ifname,
			Index: int(ifindex),
		},
	}
	netlink.LinkSetMasterByIndex(link, int(master))
}

func NetlinkVrfUnbindInterface(ifname string, ifindex IfIndex) {
	link := &netlink.Dummy{
		netlink.LinkAttrs{
			Name:  ifname,
			Index: int(ifindex),
		},
	}
	netlink.LinkSetMasterByIndex(link, 0)
}

func NetlinkIpAddrAdd(ifp *Interface, p *netutil.Prefix) {
	ipnet := netutil.IPNetFromPrefix(p)
	addr := &netlink.Addr{IPNet: &ipnet}

	// Specify broadcast address when the address is IPv4 and interface is
	// broadcast. Generate boradcast address when prefix length is less than or
	// equal 30 otherwise use the address as broadcast.
	if p.AFI() == netutil.AFI_IP && ifp.IsBroadcast() {
		if p.Length <= 30 {
			broadcast := p.Copy().ApplyReverseMask()
			addr.Broadcast = broadcast.IP
		} else {
			addr.Broadcast = p.IP
		}
	}

	link := &netlink.Dummy{
		netlink.LinkAttrs{
			Name:  ifp.Name,
			Index: int(ifp.Index),
		},
	}
	err := netlink.AddrAdd(link, addr)
	if err != nil {
		fmt.Println("IpAddrAdd() AddrAdd:", err)
	}
}

func NetlinkIpAddrDelete(ifp *Interface, p *netutil.Prefix) {
	ipnet := netutil.IPNetFromPrefix(p)
	addr := &netlink.Addr{IPNet: &ipnet}
	link := &netlink.Dummy{
		netlink.LinkAttrs{
			Name:  ifp.Name,
			Index: int(ifp.Index),
		},
	}
	err := netlink.AddrDel(link, addr)
	if err != nil {
		fmt.Println("IpAddrDeelte() AddrDel:", err)
	}
}

func LinkSetMtu(ifp *Interface, mtu uint32) error {
	return netlink.LinkSetMTU(&netlink.Dummy{netlink.LinkAttrs{Index: int(ifp.Index)}}, int(mtu))
}

func LinkSetUp(ifp *Interface) error {
	return netlink.LinkSetUp(&netlink.Dummy{netlink.LinkAttrs{Index: int(ifp.Index)}})
}

func LinkSetDown(ifp *Interface) error {
	return netlink.LinkSetDown(&netlink.Dummy{netlink.LinkAttrs{Index: int(ifp.Index)}})
}

// Encap Helper functions for LWT (light weight tunnel)
func EncapTypeString(typ int) string {
	switch typ {
	case nl.LWTUNNEL_ENCAP_NONE:
		return "none"
	case nl.LWTUNNEL_ENCAP_MPLS:
		return "mpls"
	case nl.LWTUNNEL_ENCAP_IP:
		return "ip"
	case nl.LWTUNNEL_ENCAP_ILA:
		return "ila"
	case nl.LWTUNNEL_ENCAP_IP6:
		return "ip6"
	case nl.LWTUNNEL_ENCAP_SEG6:
		return "seg6"
	case nl.LWTUNNEL_ENCAP_BPF:
		return "bpf"
	}
	return "unknown"
}
