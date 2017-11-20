// Copyright 2016, 2017 zebra project.
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
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/coreswitch/netutil"
)

// ZAPI version 2.
//
// Header length is 6.
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           Length (2)          |  Marker (1)   |  Version (1)  |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |          Command (2)          |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

// ZAPI version 3.
//
// Header length is 8.
//
// 0                   1                   2                   3
// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           Length (2)          |  Marker (1)   |  Version (1)  |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
// |           VRF ID (2)          |          Command (2)          |
// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

// ZAPI version 2 message types.
//
// #define ZEBRA_INTERFACE_ADD                1
// #define ZEBRA_INTERFACE_DELETE             2
// #define ZEBRA_INTERFACE_ADDRESS_ADD        3
// #define ZEBRA_INTERFACE_ADDRESS_DELETE     4
// #define ZEBRA_INTERFACE_UP                 5
// #define ZEBRA_INTERFACE_DOWN               6
// #define ZEBRA_IPV4_ROUTE_ADD               7
// #define ZEBRA_IPV4_ROUTE_DELETE            8
// #define ZEBRA_IPV6_ROUTE_ADD               9
// #define ZEBRA_IPV6_ROUTE_DELETE           10
// #define ZEBRA_REDISTRIBUTE_ADD            11
// #define ZEBRA_REDISTRIBUTE_DELETE         12
// #define ZEBRA_REDISTRIBUTE_DEFAULT_ADD    13
// #define ZEBRA_REDISTRIBUTE_DEFAULT_DELETE 14
// #define ZEBRA_IPV4_NEXTHOP_LOOKUP         15
// #define ZEBRA_IPV6_NEXTHOP_LOOKUP         16
// #define ZEBRA_IPV4_IMPORT_LOOKUP          17
// #define ZEBRA_IPV6_IMPORT_LOOKUP          18
// #define ZEBRA_INTERFACE_RENAME            19
// #define ZEBRA_ROUTER_ID_ADD               20
// #define ZEBRA_ROUTER_ID_DELETE            21
// #define ZEBRA_ROUTER_ID_UPDATE            22
// #define ZEBRA_HELLO                       23
// #define ZEBRA_IPV4_NEXTHOP_LOOKUP_MRIB    24
// #define ZEBRA_MESSAGE_MAX                 25

type Client struct {
	Version   uint8
	VrfId     int
	RouterId  bool
	Interface bool
}

var (
	ClientMap   = map[net.Conn]*Client{}
	ClientMutex sync.RWMutex
)

func ClientRegister(conn net.Conn) {
	ClientMutex.Lock()
	defer ClientMutex.Unlock()
	fmt.Println("ClientRegister", conn)
	ClientMap[conn] = &Client{}
}

func ClientUnregister(conn net.Conn) {
	ClientMutex.Lock()
	defer ClientMutex.Unlock()
	fmt.Println("ClientUnregister", conn)
	delete(ClientMap, conn)
}

type Message struct {
	Header Header
	Body   Body
}

func (m *Message) Serialize() ([]byte, error) {
	var body []byte
	if m.Body != nil {
		var err error
		body, err = m.Body.Serialize()
		if err != nil {
			return nil, err
		}
	}
	m.Header.Length = uint16(len(body) + HeaderSize(m.Header.Version))
	hdr, err := m.Header.Serialize()
	if err != nil {
		return nil, err
	}
	return append(hdr, body...), nil
}

type Header struct {
	Length  uint16
	Marker  uint8
	Version uint8
	VrfId   uint16
	Command COMMAND_TYPE
}

const (
	HEADER_V2_LEN = 6
	HEADER_V3_LEN = 8
)

func HeaderSize(version uint8) int {
	switch version {
	case 2:
		return HEADER_V2_LEN
	case 3:
		return HEADER_V3_LEN
	default:
		return HEADER_V2_LEN
	}
}

func (h *Header) Serialize() ([]byte, error) {
	if h.Marker != HEADER_MARKER {
		h.Marker = HEADER_MARKER
	}
	if h.Version != 3 {
		h.Version = 2
	}
	buf := make([]byte, HeaderSize(h.Version))
	binary.BigEndian.PutUint16(buf[0:], h.Length)
	buf[2] = h.Marker
	buf[3] = h.Version
	switch h.Version {
	case 2:
		binary.BigEndian.PutUint16(buf[4:], uint16(h.Command))
	case 3:
		binary.BigEndian.PutUint16(buf[4:6], uint16(h.VrfId))
		binary.BigEndian.PutUint16(buf[6:], uint16(h.Command))
	}
	return buf, nil
}

func (h *Header) DecodeFromBytes(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("ZAPI message header length is too small")
	}

	h.Length = binary.BigEndian.Uint16(data[0:2])
	h.Marker = data[2]
	h.Version = data[3]

	if h.Version != 2 && h.Version != 3 {
		return fmt.Errorf("Unsupported ZAPI version")
	}

	switch h.Version {
	case 2:
		if len(data) < HEADER_V2_LEN {
			return fmt.Errorf("Header length %d is smaller than minium vesion 2 length", len(data))
		}
		h.Command = COMMAND_TYPE(binary.BigEndian.Uint16(data[4:6]))
	case 3:
		if len(data) < HEADER_V3_LEN {
			return fmt.Errorf("Header length %d is smaller than minium version 3 length", len(data))
		}
		h.VrfId = binary.BigEndian.Uint16(data[4:6])
		h.Command = COMMAND_TYPE(binary.BigEndian.Uint16(data[6:8]))
	}
	return nil
}

type Body interface {
	DecodeFromBytes([]byte) error
	Serialize() ([]byte, error)
	Process(net.Conn, *Header) error
}

// ROUTER_ID_ADD
type RouterIDUpdateBody struct {
	IP     net.IP
	Length uint8
}

func (b *RouterIDUpdateBody) Serialize() ([]byte, error) {
	buf := make([]byte, 1)
	buf[0] = syscall.AF_INET
	buf = append(buf, b.IP...)
	buf = append(buf, byte(b.Length))
	return buf, nil
}

func (b *RouterIDUpdateBody) DecodeFromBytes(data []byte) error {
	family := data[0]
	var addrlen int8
	switch family {
	case syscall.AF_INET:
		addrlen = net.IPv4len
	case syscall.AF_INET6:
		addrlen = net.IPv6len
	default:
		return fmt.Errorf("unknown address family: %d", family)
	}
	b.IP = data[1 : 1+addrlen]
	b.Length = data[1+addrlen]
	return nil
}

func (b *RouterIDUpdateBody) String() string {
	return fmt.Sprintf("id: %s/%d", b.IP, b.Length)
}

func (b *RouterIDUpdateBody) Process(net.Conn, *Header) error {
	return nil
}

func Hello(conn net.Conn, h *Header, data []byte) error {
	fmt.Println("[zapi]HELLO handler", len(data), data)

	ClientMutex.Lock()
	defer ClientMutex.Unlock()

	if client, ok := ClientMap[conn]; ok {
		fmt.Println("[zapi]Register version", h.Version, "and vrfId", h.VrfId)
		client.Version = h.Version
		client.VrfId = int(h.VrfId)
	}
	return nil
}

func RouterIdUpdateSend(conn net.Conn, version byte, routerId net.IP) {
	body := &RouterIDUpdateBody{IP: routerId, Length: 32}
	m := &Message{
		Header: Header{
			Marker:  HEADER_MARKER,
			Version: version,
			VrfId:   0,
			Command: ROUTER_ID_UPDATE,
		},
		Body: body,
	}
	s, _ := m.Serialize()
	conn.Write(s)
}

func RouterIdAdd(conn net.Conn, version byte, data []byte) {
	ClientMutex.Lock()
	defer ClientMutex.Unlock()

	if client, ok := ClientMap[conn]; ok {
		client.RouterId = true

		v := VrfLookupByIndex(client.VrfId)
		if v != nil {
			RouterIdUpdateSend(conn, version, v.RouterId())
		}
	}
}

func RouterIdUpdate(vrfId int, routerId net.IP) {
	ClientMutex.Lock()
	defer ClientMutex.Unlock()

	for conn, client := range ClientMap {
		if client.Version == 2 && client.VrfId == vrfId && client.RouterId {
			RouterIdUpdateSend(conn, client.Version, routerId)
		}
	}
}

type INTERFACE_STATUS uint8

const (
	INTERFACE_ACTIVE        = 0x01
	INTERFACE_SUB           = 0x02
	INTERFACE_LINKDETECTION = 0x04
)

func (t INTERFACE_STATUS) String() string {
	ss := make([]string, 0, 3)
	if t&INTERFACE_ACTIVE > 0 {
		ss = append(ss, "ACTIVE")
	}
	if t&INTERFACE_SUB > 0 {
		ss = append(ss, "SUB")
	}
	if t&INTERFACE_LINKDETECTION > 0 {
		ss = append(ss, "LINKDETECTION")
	}
	return strings.Join(ss, "|")
}

type InterfaceUpdateBody struct {
	Name      string
	Index     uint32
	Status    INTERFACE_STATUS
	Flags     uint64
	Metric    uint32
	Mtu       uint32
	Mtu6      uint32
	Bandwidth uint32
	HwAddr    net.HardwareAddr
}

func NewInterfaceUpdateBody(ifp *Interface) *InterfaceUpdateBody {
	body := &InterfaceUpdateBody{
		Name:   ifp.Name,
		Index:  uint32(ifp.Index),
		Status: INTERFACE_ACTIVE,
		Flags:  uint64(ifp.Flags),
		Metric: uint32(ifp.Metric),
		Mtu:    uint32(ifp.Mtu),
		Mtu6:   uint32(ifp.Mtu),
		HwAddr: ifp.HwAddr,
	}
	return body
}

func (b *InterfaceUpdateBody) Serialize() ([]byte, error) {
	buf := make([]byte, INTERFACE_NAMSIZ+37+len(b.HwAddr))
	copy(buf, b.Name)
	binary.BigEndian.PutUint32(buf[INTERFACE_NAMSIZ:INTERFACE_NAMSIZ+4], b.Index)
	buf[INTERFACE_NAMSIZ+4] = byte(b.Status)
	binary.BigEndian.PutUint64(buf[INTERFACE_NAMSIZ+5:INTERFACE_NAMSIZ+13], b.Flags)
	binary.BigEndian.PutUint32(buf[INTERFACE_NAMSIZ+13:INTERFACE_NAMSIZ+17], b.Metric)
	binary.BigEndian.PutUint32(buf[INTERFACE_NAMSIZ+17:INTERFACE_NAMSIZ+21], b.Mtu)
	binary.BigEndian.PutUint32(buf[INTERFACE_NAMSIZ+21:INTERFACE_NAMSIZ+25], b.Mtu6)
	binary.BigEndian.PutUint32(buf[INTERFACE_NAMSIZ+25:INTERFACE_NAMSIZ+29], b.Bandwidth)
	binary.BigEndian.PutUint32(buf[INTERFACE_NAMSIZ+29:INTERFACE_NAMSIZ+33], ZEBRA_LLT_ETHER)
	binary.BigEndian.PutUint32(buf[INTERFACE_NAMSIZ+33:INTERFACE_NAMSIZ+37], uint32(len(b.HwAddr)))
	if len(b.HwAddr) > 0 {
		hw := buf[INTERFACE_NAMSIZ+37:]
		copy(hw, b.HwAddr)
	}
	return buf, nil
}

func (b *InterfaceUpdateBody) DecodeFromBytes([]byte) error {
	return nil
}

func (b *InterfaceUpdateBody) Process(net.Conn, *Header) error {
	return nil
}

type InterfaceAddressUpdateBody struct {
	Index  uint32
	Flags  uint8
	Prefix net.IP
	Length uint8
}

func NewInterfaceAddrUpdateBody(addr *IfAddr) *InterfaceAddressUpdateBody {
	body := &InterfaceAddressUpdateBody{
		Prefix: addr.Prefix.IP,
		Length: uint8(addr.Prefix.Length),
	}
	return body
}

func (b *InterfaceAddressUpdateBody) DecodeFromBytes(data []byte) error {
	b.Index = binary.BigEndian.Uint32(data[:4])
	b.Flags = data[4]
	family := data[5]
	var addrlen int8
	switch family {
	case syscall.AF_INET:
		addrlen = net.IPv4len
	case syscall.AF_INET6:
		addrlen = net.IPv6len
	default:
		return fmt.Errorf("unknown address family: %d", family)
	}
	b.Prefix = data[6 : 6+addrlen]
	b.Length = data[6+addrlen]
	return nil
}

func (b *InterfaceAddressUpdateBody) Serialize() ([]byte, error) {
	buf := make([]byte, 6)
	binary.BigEndian.PutUint32(buf[:4], b.Index)
	buf[4] = b.Flags
	buf[5] = syscall.AF_INET
	buf = append(buf, b.Prefix...)
	buf = append(buf, b.Length)
	buf = append(buf, make([]byte, 4)...)
	return buf, nil
}

func (b *InterfaceAddressUpdateBody) Process(net.Conn, *Header) error {
	return nil
}

func interfaceAddressAdd(conn net.Conn, version byte, ifp *Interface, addr *IfAddr) {
	body := NewInterfaceAddrUpdateBody(addr)
	body.Index = uint32(ifp.Index)

	m := &Message{
		Header: Header{
			Marker:  HEADER_MARKER,
			Version: version,
			VrfId:   0,
			Command: INTERFACE_ADDRESS_ADD,
		},
		Body: body,
	}
	s, _ := m.Serialize()
	//fmt.Println(s)
	conn.Write(s)
	//fmt.Println(written, err)
}

func InterfaceAdd(conn net.Conn, h *Header, data []byte) {
	//fmt.Println("INTERFACE_ADD handler")

	v := VrfLookupByIndex(int(h.VrfId))
	if v == nil {
		return
	}

	for n := v.IfTable.Top(); n != nil; n = v.IfTable.Next(n) {
		ifp := n.Item.(*Interface)
		body := NewInterfaceUpdateBody(ifp)

		m := &Message{
			Header: Header{
				Marker:  HEADER_MARKER,
				Version: h.Version,
				VrfId:   0,
				Command: INTERFACE_ADD,
			},
			Body: body,
		}
		s, _ := m.Serialize()
		//fmt.Println(s)
		conn.Write(s)
		//fmt.Println(written, err)

		for _, addr := range ifp.Addrs[AFI_IP] {
			interfaceAddressAdd(conn, h.Version, ifp, addr)
		}
	}

	go func() {
		time.Sleep(time.Second * 5)
		RedistSync(int(h.VrfId), conn)
	}()
}

type IPRouteBody struct {
	Type      ROUTE_TYPE
	Flags     FLAG
	Message   uint8
	SAFI      SAFI
	Prefix    *netutil.Prefix
	Nexthop   *Nexthop
	Nexthops  []*Nexthop
	Ifindexes []uint32
	Distance  uint8
	Metric    uint32
	Api       COMMAND_TYPE
	Version   uint8
}

func RedistSyncVrf(vrf *Vrf, vrfId int, conn net.Conn) {
	ptree := vrf.ribTable[AFI_IP]
	for n := ptree.Top(); n != nil; n = ptree.Next(n) {
		if n.Item != nil {
			ip := make([]byte, 4)
			copy(ip, n.Key())
			p := netutil.PrefixFromIPPrefixlen(ip, n.KeyLength())
			ribs := n.Item.(RibSlice)
			for _, rib := range ribs {
				if rib.IsSelectedFib() {
					RedistIPv4Add(vrfId, p, rib, conn)
				}
			}
		}
	}
}

func RedistSync(vrfId int, conn net.Conn) {
	fmt.Println("[zapi]RedistSync", vrfId)

	if vrfId == 0 {
		// When VRF ID is zero (version 3).
		for _, vrf := range VrfMap {
			if vrf.Index != 0 {
				RedistSyncVrf(vrf, vrf.Index, conn)
			}
		}
	} else {
		// When VRF ID is specified just walk through the RIB.
		vrf := VrfLookupByIndex(vrfId)
		if vrf == nil {
			fmt.Println("[zapi]RedistSync can't find VRF", vrfId)
			return
		}
		RedistSyncVrf(vrf, vrfId, conn)
	}
}

var HubNodeMap = map[int]string{}

func VrfHubNodeAdd(vrf string, hubNode string) {
	vrfId := VrfExtractIndex(vrf)
	if vrfId == 0 {
		return
	}
	HubNodeMap[vrfId] = hubNode
}

func VrfHubNodeDelete(vrf string, hubNode string) {
	vrfId := VrfExtractIndex(vrf)
	if vrfId == 0 {
		return
	}
	delete(HubNodeMap, vrfId)
}

func RedistIPv4Route(command COMMAND_TYPE, conn net.Conn, version uint8, vrfId int, p *netutil.Prefix, rib *Rib) {
	fmt.Println("Redist IPv4 route", p, vrfId, command.String())

	if version == 3 {
		cmd := "add"
		if command == IPV4_ROUTE_DELETE {
			cmd = "del"
		}
		fmt.Println("Redist Executing /usr/bin/gobgp-vrf.sh", strconv.Itoa(vrfId), cmd, p.String(), rib.Metric)

		if hubNode, ok := HubNodeMap[vrfId]; ok {
			err := exec.Command("/usr/bin/gobgp-vrf.sh", strconv.Itoa(vrfId), cmd, p.String(), strconv.Itoa(int(rib.Metric)), hubNode).Run()
			if err != nil {
				fmt.Println("Redist IPv4 route script error:", err)
			}
		} else {
			err := exec.Command("/usr/bin/gobgp-vrf.sh", strconv.Itoa(vrfId), cmd, p.String(), strconv.Itoa(int(rib.Metric))).Run()
			if err != nil {
				fmt.Println("Redist IPv4 route script error:", err)
			}
		}
		return
	}

	body := &IPRouteBody{
		Type:    ROUTE_CONNECT,
		Flags:   0,
		Message: MESSAGE_NEXTHOP,
		SAFI:    SAFI_UNICAST,
		Version: version,
	}
	body.Prefix = p
	addr := netutil.ParseIPv4("0.0.0.0")
	if rib.Type == RIB_BGP && rib.Nexthop != nil {
		addr = rib.Nexthop.IP
	}
	body.Nexthops = append(body.Nexthops, NewNexthopAddr(addr))
	body.Ifindexes = append(body.Ifindexes, 0)
	m := &Message{
		Header: Header{
			Marker:  HEADER_MARKER,
			Version: version,
			VrfId:   uint16(vrfId),
			Command: command,
		},
		Body: body,
	}
	s, _ := m.Serialize()
	fmt.Println(s)
	conn.Write(s)
}

func ClientVersion(conn net.Conn) uint8 {
	ClientMutex.Lock()
	defer ClientMutex.Unlock()

	if client, ok := ClientMap[conn]; ok {
		return client.Version
	}
	return 2
}

func EsiConnected(p *netutil.Prefix) bool {
	if len(p.IP) > 0 {
		if p.IP[0] == 172 && p.Length == 12 {
			return true
		}
		if p.IP[0] == 198 && p.Length == 15 {
			return true
		}
	}
	return false
}

func RedistIPv4Add(vrfId int, p *netutil.Prefix, rib *Rib, conn net.Conn) {
	if rib.Type != RIB_CONNECTED && rib.Type != RIB_BGP && rib.Type != RIB_OSPF {
		return
	}
	if vrfId == 0 {
		//fmt.Println("RedistIPv4Add: do not perform redist for default Vrf")
		return
	}
	if len(p.IP) != 4 {
		//fmt.Println("RedistIPv4Add: non IPv4 length", len(p.IP))
		return
	}

	ClientMutex.Lock()
	defer ClientMutex.Unlock()

	if conn == nil {
		fmt.Println("RedistIPv4Add: client len", len(ClientMap))
		for conn, client := range ClientMap {
			if client.Version == 0 {
				fmt.Println("RedistIPv4Add: skip bogus client")
				continue
			}
			if rib.Src == conn {
				fmt.Println("RedistIPv4Add: same source RIB")
				continue
			}
			if client.Version == 2 && client.VrfId != vrfId {
				fmt.Println("RedistIPv4Add: version 2 and vrfId is different")
				continue
			}
			if client.Version == 2 && rib.Type == RIB_CONNECTED && EsiConnected(p) {
				fmt.Println("RedistIPv4Add: version 2 and backbone connected")
				continue
			}
			if client.Version == 2 && rib.Metric != 0 {
				fmt.Println("RedistIPv4Add: version 2 and metric is not 0")
				continue
			}
			if client.Version == 3 && rib.Type == RIB_CONNECTED {
				fmt.Println("RedistIPv4Add: version 3 do not redist connected")
				continue
			}
			RedistIPv4Route(IPV4_ROUTE_ADD, conn, client.Version, vrfId, p, rib)
		}
	} else {
		// Syncer part.
		if rib.Src != conn {
			var ver uint8
			if client, ok := ClientMap[conn]; ok {
				ver = client.Version
			} else {
				ver = 2
			}
			if ver == 0 {
				fmt.Println("RedistIPv4Add: skip bogus client")
				return
			}
			// Already checked.
			// if rib.Src == conn {
			// 	fmt.Println("RedistIPv4Add: same source RIB")
			// 	continue
			// }
			// Already checked at caller.
			// if ver == 2 && client.VrfId != vrfId {
			// 	fmt.Println("RedistIPv4Add: version 2 and vrfId is different")
			// 	continue
			// }
			if ver == 2 && rib.Type == RIB_CONNECTED && EsiConnected(p) {
				fmt.Println("RedistIPv4Add: version 2 and backbone connected")
				return
			}
			if ver == 2 && rib.Metric != 0 {
				fmt.Println("RedistIPv4Add: version 2 and metric is not 0")
				return
			}
			if ver == 3 && rib.Type == RIB_CONNECTED {
				fmt.Println("RedistIPv4Add: version 3 do not redist connected")
				return
			}
			RedistIPv4Route(IPV4_ROUTE_ADD, conn, ver, vrfId, p, rib)
		}
	}
}

func RedistIPv4Delete(vrfId int, p *netutil.Prefix, rib *Rib) {
	if rib.Type != RIB_CONNECTED && rib.Type != RIB_BGP && rib.Type != RIB_OSPF {
		return
	}
	if vrfId == 0 {
		//fmt.Println("RedistIPv4Delete: do not perform redist for default Vrf")
		return
	}
	if len(p.IP) != 4 {
		fmt.Println("RedistIPv4Delete: non IPv4 length", len(p.IP))
		return
	}

	ClientMutex.Lock()
	defer ClientMutex.Unlock()

	fmt.Println("RedistIPv4Delete: client len", len(ClientMap))
	for conn, client := range ClientMap {
		if client.Version == 0 {
			fmt.Println("RedistIPv4Delete: skip bogus client")
			continue
		}
		if rib.Src == conn {
			fmt.Println("RedistIPv4Delete: same source RIB")
			continue
		}
		if client.Version == 2 && client.VrfId != vrfId {
			fmt.Println("RedistIPv4Delete: version 2 and vrfId is different")
			continue
		}
		if client.Version == 2 && rib.Metric != 0 {
			fmt.Println("RedistIPv4Delete: version 2 and metric is not 0")
			continue
		}
		if client.Version == 3 && rib.Type == RIB_CONNECTED {
			fmt.Println("RedistIPv4Delete: version 3 do not redist connected")
			continue
		}
		RedistIPv4Route(IPV4_ROUTE_DELETE, conn, client.Version, vrfId, p, rib)
	}
}

func (b *IPRouteBody) SerializeV2() ([]byte, error) {
	buf := make([]byte, 3)
	buf[0] = uint8(b.Type)
	buf[1] = uint8(b.Flags)
	buf[2] = b.Message

	bitlen := byte(b.Prefix.Length)
	bytelen := (int(b.Prefix.Length) + 7) / 8
	bbuf := make([]byte, bytelen)

	copy(bbuf, b.Prefix.IP)
	buf = append(buf, bitlen)
	buf = append(buf, bbuf...)

	if b.Message&MESSAGE_NEXTHOP > 0 {
		if b.Flags&FLAG_BLACKHOLE > 0 {
			buf = append(buf, []byte{1, uint8(NEXTHOP_BLACKHOLE)}...)
		} else {
			buf = append(buf, uint8(len(b.Nexthops)+len(b.Ifindexes)))
		}
		for _, v := range b.Nexthops {
			buf = append(buf, v.IP...)
		}
	}

	if b.Message&MESSAGE_DISTANCE > 0 {
		buf = append(buf, b.Distance)
	}

	if b.Message&MESSAGE_METRIC > 0 {
		bbuf := make([]byte, 4)
		binary.BigEndian.PutUint32(bbuf, b.Metric)
		buf = append(buf, bbuf...)
	}

	return buf, nil
}

func (b *IPRouteBody) SerializeV3() ([]byte, error) {
	buf := make([]byte, 3)
	buf[0] = uint8(b.Type)
	buf[1] = uint8(b.Flags)
	buf[2] = b.Message

	bitlen := byte(b.Prefix.Length)
	bytelen := (int(b.Prefix.Length) + 7) / 8
	bbuf := make([]byte, bytelen)

	copy(bbuf, b.Prefix.IP)
	buf = append(buf, bitlen)
	buf = append(buf, bbuf...)

	if b.Message&MESSAGE_NEXTHOP > 0 {
		if b.Flags&FLAG_BLACKHOLE > 0 {
			buf = append(buf, []byte{1, uint8(NEXTHOP_BLACKHOLE)}...)
		} else {
			buf = append(buf, uint8(len(b.Nexthops)))
		}
		for _, v := range b.Nexthops {
			buf = append(buf, v.IP...)
		}

		for _, v := range b.Ifindexes {
			buf = append(buf, uint8(NEXTHOP_IFINDEX))
			bbuf := make([]byte, 4)
			binary.BigEndian.PutUint32(bbuf, v)
			buf = append(buf, bbuf...)
		}
	}

	if b.Message&MESSAGE_DISTANCE > 0 {
		buf = append(buf, b.Distance)
	}

	if b.Message&MESSAGE_METRIC > 0 {
		bbuf := make([]byte, 4)
		binary.BigEndian.PutUint32(bbuf, b.Metric)
		buf = append(buf, bbuf...)
	}

	return buf, nil
}

func (b *IPRouteBody) Serialize() ([]byte, error) {
	if b.Version == 3 {
		return b.SerializeV3()
	} else {
		return b.SerializeV2()
	}
}

func (b *IPRouteBody) DecodeFromBytes(data []byte) error {
	isV4 := false
	//addrLen := int(net.IPv6len)

	if b.Api == IPV4_ROUTE_ADD || b.Api == IPV4_ROUTE_DELETE {
		isV4 = true
	}

	b.Type = ROUTE_TYPE(data[0])
	b.Flags = FLAG(data[1])
	b.Message = data[2]
	b.SAFI = SAFI(binary.BigEndian.Uint16(data[3:5]))

	// Decode Prefix
	if isV4 {
		b.Prefix = netutil.NewPrefixAFI(netutil.AFI_IP)
	} else {
		b.Prefix = netutil.NewPrefixAFI(netutil.AFI_IP6)
	}
	b.Prefix.Length = int(data[5])
	byteLen := int((b.Prefix.Length + 7) / 8)
	pos := 6
	copy(b.Prefix.IP, data[pos:pos+byteLen])
	pos += byteLen

	if b.Message&MESSAGE_NEXTHOP > 0 {
		numNexthop := int(data[pos])
		pos += 1

		for i := 0; i < numNexthop; i++ {
			flag := NEXTHOP_FLAG(data[pos])
			pos += 1

			var addr net.IP
			var ifindex uint32
			var nexthop *Nexthop

			switch flag {
			case NEXTHOP_IFINDEX:
				ifindex = binary.BigEndian.Uint32(data[pos : pos+4])
				nexthop = NewNexthopIf(IfIndex(ifindex))
				pos += 4
			case NEXTHOP_IPV4:
				addr = data[pos : pos+4]
				nexthop = NewNexthopAddr(net.IP(addr).To4())
				pos += 4
			case NEXTHOP_IPV4_IFINDEX:
				addr = data[pos : pos+4]
				pos += 4
				ifindex = binary.BigEndian.Uint32(data[pos : pos+4])
				pos += 4
				nexthop = NewNexthopAddrIf(net.IP(addr).To4(), IfIndex(ifindex))
			}
			if numNexthop == 1 {
				b.Nexthop = nexthop
			} else {
				b.Nexthops = append(b.Nexthops, nexthop)
			}
		}
	}

	if b.Message&MESSAGE_DISTANCE > 0 {
		b.Distance = data[pos]
		pos += 1
	}

	if b.Message&MESSAGE_METRIC > 0 {
		b.Metric = binary.BigEndian.Uint32(data[pos : pos+4])
		pos += 4
		if OspfMetricFilter && b.Type == ROUTE_OSPF {
			b.Metric = 0
		}
	}

	return nil
}

func (b *IPRouteBody) Process(net.Conn, *Header) error {
	return nil
}

func IPv4Route(command COMMAND_TYPE, version uint8, conn net.Conn, data []byte, vrfId uint16) {
	// Parse IPv4Route.
	body := &IPRouteBody{Api: command}
	body.DecodeFromBytes(data)

	if DefaultVrfProtect && version == 3 && vrfId == 0 {
		return
	}

	if body.Nexthop != nil && body.Nexthop.IP.Equal(net.IPv4zero.To4()) {
		return
	}

	// Prepare RibInfo
	vrf := VrfLookupByIndex(int(vrfId))
	if vrf == nil {
		fmt.Println("Can't find VRF id:", vrfId)

		// Try to create VRF.
		server.VrfAdd(fmt.Sprintf("vrf%d", vrfId))
		vrf = VrfLookupByIndex(int(vrfId))
		if vrf == nil {
			fmt.Println("Couldn't create VRF id:", vrfId)
			return
		}
	}

	ri := &Rib{
		Type:     RouteType2RibType(body.Type),
		Nexthop:  body.Nexthop,
		Nexthops: body.Nexthops,
		Src:      conn,
		Metric:   body.Metric,
	}

	// Call RIB API.
	if command == IPV4_ROUTE_ADD {
		fmt.Println("Route add", body.Prefix)
		vrf.RibAdd(body.Prefix, ri)
	} else {
		fmt.Println("Route delete", body.Prefix)
		vrf.RibDelete(body.Prefix, ri)
	}
}

// Redistribute message for:
//
// ZEBRA_REDISTRIBUTE_ADD            11
// ZEBRA_REDISTRIBUTE_DELETE         12
//
type RedistributeBody struct {
	Type uint8
}

func (b *RedistributeBody) Serialize() ([]byte, error) {
	buf := make([]byte, 1)
	buf[0] = b.Type
	return buf, nil
}

func (b *RedistributeBody) DecodeFromBytes(data []byte) error {
	b.Type = data[0]
	return nil
}

func (b *RedistributeBody) Process(conn net.Conn, h *Header) error {
	fmt.Println("Processing", h.Command.String(), RouteTypeStringMap[ROUTE_TYPE(b.Type)])
	return nil
}

type IPv4NexthopLookupBody struct {
	Addr net.IP
}

func (b *IPv4NexthopLookupBody) Serialize() ([]byte, error) {
	buf := make([]byte, 4)
	copy(buf, b.Addr)
	return buf, nil
}

func (b *IPv4NexthopLookupBody) DecodeFromBytes(data []byte) error {
	b.Addr = make([]byte, 4)
	copy(b.Addr, data[:4])
	return nil
}

func (b *IPv4NexthopLookupBody) Process(net.Conn, *Header) error {
	return nil
}

type IPv4NexthopReplyBody struct {
	Addr       net.IP
	Metric     uint32
	NexthopNum uint8
	Nexthops   []*Nexthop
}

func (b *IPv4NexthopReplyBody) Serialize() ([]byte, error) {
	buf := make([]byte, 9)
	copy(buf, b.Addr)
	binary.BigEndian.PutUint32(buf[4:8], b.Metric)
	buf[8] = b.NexthopNum
	for _, nexthop := range b.Nexthops {
		nbuf := make([]byte, 5)
		nbuf[0] = byte(NEXTHOP_IPV4)
		copy(nbuf[1:], nexthop.IP)
		buf = append(buf, nbuf...)
	}
	// fmt.Println("buf len", len(buf))
	// for i := 0; i < len(buf); i++ {
	// 	fmt.Printf("%d ", buf[i])
	// }
	// fmt.Printf("\n")

	return buf, nil
}

func (b *IPv4NexthopReplyBody) DecodeFromBytes([]byte) error {
	return nil
}

func (b *IPv4NexthopReplyBody) Process(net.Conn, *Header) error {
	return nil
}

var NexthopLookupHook func(vrf *Vrf, nexthop net.IP) bool

func EsiNexthopLookup(vrf *Vrf, nexthop net.IP) bool {
	ptree := vrf.ribTable[AFI_IP]
	fmt.Println("EsiNexthopLookup", nexthop)
	n := ptree.Match(nexthop, 32)
	if n != nil {
		if n.Item != nil {
			fmt.Println("EsiNexthopLookup: n.Item is not nil")
			for _, rib := range n.Item.(RibSlice) {
				if rib.IsFib() && rib.Type == RIB_CONNECTED {
					if rib.Nexthop != nil && rib.Nexthop.IsIfOnly() {
						fmt.Println("EsiNexthopLookup: Interface only nexthop", rib.Nexthop.Index)
						ifp := IfLookupByIndex(rib.Nexthop.Index)
						fmt.Println("EsiNexthopLookup: ifp", ifp.Name)
						r := regexp.MustCompile(`sproute\d+`)
						if r.MatchString(ifp.Name) {
							fmt.Println("EsiNexthopLookup: sproute interface, perform tunnel check")
							result := DtlsNexthop(vrf.Index, nexthop)
							fmt.Println("EsiNexthopLookup: return", result)
							return result
						}
					}
				}
			}
		}
	}
	return true
}

func IPv4NexthopLookup(conn net.Conn, version byte, data []byte, vrfId uint16) {
	vrf := VrfLookupByIndex(int(vrfId))
	if vrf == nil {
		fmt.Println("[zapi]IPv4NexthopLookup: Can't find VRF with id", vrfId)
		return
	}

	body := &IPv4NexthopLookupBody{}
	body.DecodeFromBytes(data)
	fmt.Println("[zapi]Lookup nexthop:", body.Addr)

	reply := &IPv4NexthopReplyBody{
		Addr: body.Addr,
	}

	found := true
	if NexthopLookupHook != nil {
		found = NexthopLookupHook(vrf, body.Addr)
	}

	if found {
		nexthop := NewNexthopAddr(body.Addr)
		reply.Nexthops = append(reply.Nexthops, nexthop)
		reply.NexthopNum = uint8(len(reply.Nexthops))
	}

	m := &Message{
		Header: Header{
			Marker:  HEADER_MARKER,
			Version: version,
			VrfId:   vrfId,
			Command: IPV4_NEXTHOP_LOOKUP,
		},
		Body: reply,
	}
	s, _ := m.Serialize()

	conn.Write(s)
}

func HandleRequest(conn net.Conn, vrfId int) {
	defer conn.Close()

	var version byte
	for {
		// We don't know client version yet.
		data := make([]byte, HeaderSize(version))
		_, err := conn.Read(data)
		if err != nil {
			break
		}
		// Peek version information.
		if version == 0 {
			version = data[3]
			if version == 3 {
				d := make([]byte, 2)
				_, err = conn.Read(d)
				if err != nil {
					break
				}
				data = append(data, d...)
			}
		}
		h := Header{}
		h.DecodeFromBytes(data)

		payloadLen := int(h.Length) - HeaderSize(version)

		fmt.Printf("[zapi]%s(%d) payloadLen %d ver %d vrf %d (override %d)\n",
			h.Command.String(), h.Command, payloadLen, h.Version, h.VrfId, vrfId)

		if vrfId != 0 {
			h.VrfId = uint16(vrfId)
		}
		err = HandleMessage(conn, &h, payloadLen)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
	}
	fmt.Println("[zapi] disconnected", "vrf", vrfId, "version", version)
	RibClearSrc(conn)

	ClientUnregister(conn)
}

func HandleMessage(conn net.Conn, h *Header, payloadLen int) error {
	var data []byte

	if payloadLen > 0 {
		data = make([]byte, payloadLen)

		_, err := conn.Read(data)
		if err != nil {
			return err
		}
	}

	switch h.Command {
	case HELLO:
		err := Hello(conn, h, data)
		if err != nil {
			return err
		}
	case ROUTER_ID_ADD:
		RouterIdAdd(conn, h.Version, data)
	case INTERFACE_ADD:
		InterfaceAdd(conn, h, data)
	case IPV4_ROUTE_ADD, IPV4_ROUTE_DELETE:
		fmt.Println("IPv4 route add/delete", h.Command)
		IPv4Route(h.Command, h.Version, conn, data, h.VrfId)
	case IPV4_NEXTHOP_LOOKUP:
		IPv4NexthopLookup(conn, h.Version, data, h.VrfId)
	case REDISTRIBUTE_ADD:
		body := RedistributeBody{}
		err := body.DecodeFromBytes(data)
		if err != nil {
			return err
		}
		err = body.Process(conn, h)
		if err != nil {
			return err
		}
	case REDISTRIBUTE_DELETE:
		body := RedistributeBody{}
		err := body.DecodeFromBytes(data)
		if err != nil {
			return err
		}
		err = body.Process(conn, h)
		if err != nil {
			return err
		}
	case REDISTRIBUTE_DEFAULT_ADD:
	case REDISTRIBUTE_DEFAULT_DELETE:
	}
	return nil
}

type ZServer struct {
	Path   string
	Listen net.Listener
	VrfId  int
}

func ZServerStart(typ string, path string, vrfId int) *ZServer {
	var lis net.Listener
	var err error

	switch typ {
	case "tcp":
		// e.g. path: ":9000"
		tcpAddr, err := net.ResolveTCPAddr("tcp", path)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return nil
		}
		lis, err = net.ListenTCP("tcp", tcpAddr)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return nil
		}
	case "unix", "unix-writable":
		// e.g. path: "/var/run/zapi.serv"
		os.Remove(path)
		lis, err = net.Listen("unix", path)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return nil
		}
		if typ == "unix-writable" {
			err = os.Chmod(path, 0777)
			if err != nil {
				return nil
			}
		}
	default:
		fmt.Println("ZServerStart type is not unix nor tcp.")
		return nil
	}

	server := &ZServer{
		Path:   path,
		Listen: lis,
		VrfId:  vrfId,
	}

	go func() {
		fmt.Println("ZAPI Server started at", path)
		for {
			// Listen for an incoming connection.
			conn, err := lis.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				return
			}

			// Register client.
			ClientRegister(conn)

			// Handle connections in a new go routine.
			go HandleRequest(conn, vrfId)
		}
	}()

	return server
}

func ZServerStop(s *ZServer) {
	if s != nil {
		if s.Listen != nil {
			s.Listen.Close()
		}
		os.Remove(s.Path)
	}
}
