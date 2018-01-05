// Copyright 2017 Zebra Project
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
	"strings"
	//"github.com/hash-set/zebra/rib"
)

const (
	HEADER_MARKER    = 255
	INTERFACE_NAMSIZ = 20
)

var (
	DefaultVrfProtect = false
	OspfMetricFilter  = false
)

type COMMAND_TYPE uint16

const (
	_ COMMAND_TYPE = iota
	INTERFACE_ADD
	INTERFACE_DELETE
	INTERFACE_ADDRESS_ADD
	INTERFACE_ADDRESS_DELETE
	INTERFACE_UP
	INTERFACE_DOWN
	IPV4_ROUTE_ADD
	IPV4_ROUTE_DELETE
	IPV6_ROUTE_ADD
	IPV6_ROUTE_DELETE
	REDISTRIBUTE_ADD
	REDISTRIBUTE_DELETE
	REDISTRIBUTE_DEFAULT_ADD
	REDISTRIBUTE_DEFAULT_DELETE
	IPV4_NEXTHOP_LOOKUP
	IPV6_NEXTHOP_LOOKUP
	IPV4_IMPORT_LOOKUP
	IPV6_IMPORT_LOOKUP
	INTERFACE_RENAME
	ROUTER_ID_ADD
	ROUTER_ID_DELETE
	ROUTER_ID_UPDATE
	HELLO
	IPV4_NEXTHOP_LOOKUP_MRIB
	VRF_UNREGISTER
	INTERFACE_LINK_PARAMS
	NEXTHOP_REGISTER
	NEXTHOP_UNREGISTER
	NEXTHOP_UPDATE
	MESSAGE_MAX
)

var CommandTypeMap = map[COMMAND_TYPE]string{
	INTERFACE_ADD:               "INTERFACE_ADD",
	INTERFACE_DELETE:            "INTERFACE_DELETE",
	INTERFACE_ADDRESS_ADD:       "INTERFACE_ADDRESS_ADD",
	INTERFACE_ADDRESS_DELETE:    "INTERFACE_ADDRESS_DELETE",
	INTERFACE_UP:                "INTERFACE_UP",
	INTERFACE_DOWN:              "INTERFACE_DOWN",
	IPV4_ROUTE_ADD:              "IPV4_ROUTE_ADD",
	IPV4_ROUTE_DELETE:           "IPV4_ROUTE_DELETE",
	IPV6_ROUTE_ADD:              "IPV6_ROUTE_ADD",
	IPV6_ROUTE_DELETE:           "IPV6_ROUTE_DELETE",
	REDISTRIBUTE_ADD:            "REDISTRIBUTE_ADD",
	REDISTRIBUTE_DELETE:         "REDISTRIBUTE_DELETE",
	REDISTRIBUTE_DEFAULT_ADD:    "REDISTRIBUTE_DEFAULT_ADD",
	REDISTRIBUTE_DEFAULT_DELETE: "REDISTRIBUTE_DEFAULT_DELETE",
	IPV4_NEXTHOP_LOOKUP:         "IPV4_NEXTHOP_LOOKUP",
	IPV6_NEXTHOP_LOOKUP:         "IPV6_NEXTHOP_LOOKUP",
	IPV4_IMPORT_LOOKUP:          "IPV4_IMPORT_LOOKUP",
	IPV6_IMPORT_LOOKUP:          "IPV6_IMPORT_LOOKUP",
	INTERFACE_RENAME:            "INTERFACE_RENAME",
	ROUTER_ID_ADD:               "ROUTER_ID_ADD",
	ROUTER_ID_DELETE:            "ROUTER_ID_DELETE",
	ROUTER_ID_UPDATE:            "ROUTER_ID_UPDATE",
	HELLO:                       "HELLO",
	INTERFACE_LINK_PARAMS: "INTERFACE_LINK_PARAMS",
	NEXTHOP_REGISTER:      "NEXTHOP_REGISTER",
	NEXTHOP_UNREGISTER:    "NEXTHOP_UNREGISTER",
	NEXTHOP_UPDATE:        "NEXTHOP_UPDATE",
}

func (c COMMAND_TYPE) String() string {
	if str, ok := CommandTypeMap[c]; ok {
		return str
	} else {
		return "UNKNOWN_COMMAND"
	}
}

type ROUTE_TYPE uint8

const (
	ROUTE_SYSTEM ROUTE_TYPE = iota
	ROUTE_KERNEL
	ROUTE_CONNECT
	ROUTE_STATIC
	ROUTE_RIP
	ROUTE_RIPNG
	ROUTE_OSPF
	ROUTE_OSPF6
	ROUTE_ISIS
	ROUTE_BGP
	ROUTE_HSLS
	ROUTE_OLSR
	ROUTE_BABEL
	ROUTE_MAX
)

var RouteTypeStringMap = map[ROUTE_TYPE]string{
	ROUTE_SYSTEM:  "system",
	ROUTE_KERNEL:  "kernel",
	ROUTE_CONNECT: "connect",
	ROUTE_STATIC:  "static",
	ROUTE_RIP:     "rip",
	ROUTE_RIPNG:   "ripng",
	ROUTE_OSPF:    "ospf",
	ROUTE_OSPF6:   "ospf3",
	ROUTE_ISIS:    "isis",
	ROUTE_BGP:     "bgp",
	ROUTE_HSLS:    "hsls",
	ROUTE_OLSR:    "olsr",
	ROUTE_BABEL:   "babel",
}

type NEXTHOP_FLAG uint8

const (
	_ NEXTHOP_FLAG = iota
	NEXTHOP_IFINDEX
	NEXTHOP_IFNAME
	NEXTHOP_IPV4
	NEXTHOP_IPV4_IFINDEX
	NEXTHOP_IPV4_IFNAME
	NEXTHOP_IPV6
	NEXTHOP_IPV6_IFINDEX
	NEXTHOP_IPV6_IFNAME
	NEXTHOP_BLACKHOLE
)

var NexthopFlagMap = map[NEXTHOP_FLAG]string{
	NEXTHOP_IFINDEX:      "NEXTHOP_IFINDEX",
	NEXTHOP_IFNAME:       "NEXTHOP_IFNAME",
	NEXTHOP_IPV4:         "NEXTHOP_IPV4",
	NEXTHOP_IPV4_IFINDEX: "NEXTHOP_IPV4_IFINDEX",
	NEXTHOP_IPV4_IFNAME:  "NEXTHOP_IPV4_IFNAME",
	NEXTHOP_IPV6:         "NEXTHOP_IPV6",
	NEXTHOP_IPV6_IFINDEX: "NEXTHOP_IPV6_IFINDEX",
	NEXTHOP_IPV6_IFNAME:  "NEXTHOP_IPV6_IFNAME",
	NEXTHOP_BLACKHOLE:    "NEXTHOP_BLACKHOLE",
}

func (n NEXTHOP_FLAG) String() string {
	if str, ok := NexthopFlagMap[n]; ok {
		return str
	} else {
		return "NEXTHOP_UNKNOWN"
	}
}

// Message Flags
type FLAG uint64

const (
	FLAG_INTERNAL  FLAG = 0x01
	FLAG_SELFROUTE FLAG = 0x02
	FLAG_BLACKHOLE FLAG = 0x04
	FLAG_IBGP      FLAG = 0x08
	FLAG_SELECTED  FLAG = 0x10
	FLAG_CHANGED   FLAG = 0x20
	FLAG_STATIC    FLAG = 0x40
	FLAG_REJECT    FLAG = 0x80
)

func (t FLAG) String() string {
	var ss []string
	if t&FLAG_INTERNAL > 0 {
		ss = append(ss, "FLAG_INTERNAL")
	}
	if t&FLAG_SELFROUTE > 0 {
		ss = append(ss, "FLAG_SELFROUTE")
	}
	if t&FLAG_BLACKHOLE > 0 {
		ss = append(ss, "FLAG_BLACKHOLE")
	}
	if t&FLAG_IBGP > 0 {
		ss = append(ss, "FLAG_IBGP")
	}
	if t&FLAG_SELECTED > 0 {
		ss = append(ss, "FLAG_SELECTED")
	}
	if t&FLAG_CHANGED > 0 {
		ss = append(ss, "FLAG_CHANGED")
	}
	if t&FLAG_STATIC > 0 {
		ss = append(ss, "FLAG_STATIC")
	}
	if t&FLAG_REJECT > 0 {
		ss = append(ss, "FLAG_REJECT")
	}
	return strings.Join(ss, "|")
}

// Subsequent Address Family Identifier.
type SAFI uint8

const (
	_ SAFI = iota
	SAFI_UNICAST
	SAFI_MULTICAST
	SAFI_RESERVED_3
	SAFI_MPLS_VPN
	SAFI_MAX
)

const (
	MESSAGE_NEXTHOP  = 0x01
	MESSAGE_IFINDEX  = 0x02
	MESSAGE_DISTANCE = 0x04
	MESSAGE_METRIC   = 0x08
)

func MessageString(m uint8) string {
	var ss []string
	if m&MESSAGE_NEXTHOP > 0 {
		ss = append(ss, "MESSAGE_NEXTHOP")
	}
	if m&MESSAGE_IFINDEX > 0 {
		ss = append(ss, "MESSAGE_IFINDEX")
	}
	if m&MESSAGE_DISTANCE > 0 {
		ss = append(ss, "MESSAGE_DISTANCE")
	}
	if m&MESSAGE_METRIC > 0 {
		ss = append(ss, "MESSAGE_METRIC")
	}
	return strings.Join(ss, "|")
}

func RouteType2RibType(t ROUTE_TYPE) uint8 {
	switch t {
	case ROUTE_KERNEL:
		return RIB_KERNEL
	case ROUTE_CONNECT:
		return RIB_CONNECTED
	case ROUTE_STATIC:
		return RIB_STATIC
	case ROUTE_RIP:
		return RIB_RIP
	case ROUTE_RIPNG:
		return RIB_RIP
	case ROUTE_OSPF:
		return RIB_OSPF
	case ROUTE_OSPF6:
		return RIB_OSPF
	case ROUTE_ISIS:
		return RIB_ISIS
	case ROUTE_BGP:
		return RIB_BGP
	case ROUTE_HSLS:
		return RIB_UNKNOWN
	case ROUTE_OLSR:
		return RIB_UNKNOWN
	case ROUTE_BABEL:
		return RIB_UNKNOWN
	default:
		return RIB_UNKNOWN
	}
}

func RibType2RouteType(t uint8) ROUTE_TYPE {
	switch t {
	case RIB_KERNEL:
		return ROUTE_KERNEL
	case RIB_CONNECTED:
		return ROUTE_CONNECT
	case RIB_STATIC:
		return ROUTE_STATIC
	case RIB_RIP:
		return ROUTE_RIP
	case RIB_OSPF:
		return ROUTE_OSPF
	case RIB_ISIS:
		return ROUTE_ISIS
	case RIB_BGP:
		return ROUTE_BGP
	default:
		return ROUTE_SYSTEM
	}
}
