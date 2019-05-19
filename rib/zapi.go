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
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/coreswitch/log"
	"github.com/coreswitch/netutil"
	pb "github.com/coreswitch/zebra/api"
	"github.com/coreswitch/zebra/policy"
)

type Client struct {
	conn      net.Conn
	version   uint8
	allVrf    bool
	vrfId     uint32
	routeType RouteType
}

var (
	ClientMap   = map[net.Conn]*Client{}
	ClientMutex sync.RWMutex
)

func ClientRegister(conn net.Conn) *Client {
	ClientMutex.Lock()
	defer ClientMutex.Unlock()

	log.Info("zapi:ClientRegister", conn)
	client := &Client{conn: conn}
	ClientMap[conn] = client
	return client
}

func ClientUnregister(conn net.Conn) {
	ClientMutex.Lock()
	defer ClientMutex.Unlock()

	log.Info("zapi:ClientUnregister", conn)
	delete(ClientMap, conn)
}

func (c *Client) Notify(mes interface{}) {
	switch mes.(type) {
	case *pb.InterfaceUpdate:
		c.InterfaceNotify(mes.(*pb.InterfaceUpdate))
	case *pb.RouterIdUpdate:
		c.RouterIdUpdateNotify(mes.(*pb.RouterIdUpdate))
	case *pb.Route:
		c.RouteNotify(mes.(*pb.Route))
	}
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

func (m *Message) Send(conn net.Conn) {
	s, err := m.Serialize()
	if err != nil {
		log.Error(err)
		return
	}
	conn.Write(s)
}

type Header struct {
	Length  uint16
	Marker  uint8
	Version uint8
	VrfId   uint16
	Command CommandType
}

const (
	HEADER_V2_LEN    = 6
	HEADER_V3_V4_LEN = 8
)

func HeaderSize(version uint8) int {
	switch version {
	case 2:
		return HEADER_V2_LEN
	case 3, 4:
		return HEADER_V3_V4_LEN
	default:
		return HEADER_V2_LEN
	}
}

func NewMessage(version uint8, command CommandType, body Body) *Message {
	return &Message{
		Header: Header{
			Marker:  HEADER_MARKER,
			Version: version,
			VrfId:   0,
			Command: command,
		},
		Body: body,
	}
}

func (h *Header) Serialize() ([]byte, error) {
	buf := make([]byte, HeaderSize(h.Version))
	binary.BigEndian.PutUint16(buf[0:], h.Length)
	buf[2] = h.Marker
	buf[3] = h.Version
	switch h.Version {
	case 2:
		binary.BigEndian.PutUint16(buf[4:], uint16(h.Command))
	case 3, 4:
		binary.BigEndian.PutUint16(buf[4:6], uint16(h.VrfId))
		binary.BigEndian.PutUint16(buf[6:], uint16(h.Command))
	}
	return buf, nil
}

func (h *Header) DecodeFromBytes(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("ZAPI message header length %d is too small", len(data))
	}

	h.Length = binary.BigEndian.Uint16(data[0:2])
	h.Marker = data[2]
	h.Version = data[3]

	switch h.Version {
	case 2:
		if len(data) < HEADER_V2_LEN {
			return fmt.Errorf("Header length %d is smaller than minium vesion 2 length", len(data))
		}
		h.Command = CommandType(binary.BigEndian.Uint16(data[4:6]))
	case 3, 4:
		if len(data) < HEADER_V3_V4_LEN {
			return fmt.Errorf("Header length %d is smaller than minium version 3 or 4 length", len(data))
		}
		h.VrfId = binary.BigEndian.Uint16(data[4:6])
		h.Command = CommandType(binary.BigEndian.Uint16(data[6:8]))
	default:
		return fmt.Errorf("Unsupported ZAPI version: %d", h.Version)
	}
	return nil
}

type Body interface {
	DecodeFromBytes(CommandType, []byte) error
	Serialize() ([]byte, error)
	Process(*Client, *Header) error
}

type HelloBody struct {
	RouteType RouteType `json:"route-type"`
}

func (b *HelloBody) MarshalJSON() ([]byte, error) {
	helloJSON := struct {
		RouteType string `json:"route-type"`
	}{
		RouteType: b.RouteType.String(),
	}
	return json.Marshal(helloJSON)
}

func (b *HelloBody) DecodeFromBytes(command CommandType, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	if len(data) == 1 {
		b.RouteType = RouteType(data[0])
	}
	bytes, _ := json.Marshal(b)
	log.Infof("zapi:HELLO %s", string(bytes))
	return nil
}

func (b *HelloBody) Serialize() ([]byte, error) {
	return nil, nil
}

func (b *HelloBody) Process(client *Client, h *Header) error {
	client.version = h.Version
	client.vrfId = uint32(h.VrfId)
	client.routeType = b.RouteType
	return nil
}

// Router ID update.
type RouterIdUpdateBody struct {
	RouterId net.IP
	Length   uint8
}

func (b *RouterIdUpdateBody) Serialize() ([]byte, error) {
	buf := make([]byte, 1)
	buf[0] = syscall.AF_INET
	buf = append(buf, b.RouterId...)
	buf = append(buf, byte(b.Length))
	return buf, nil
}

func (b *RouterIdUpdateBody) DecodeFromBytes(command CommandType, data []byte) error {
	if len(data) == 0 {
		return nil
	}
	family := data[0]
	var addrlen int8
	switch family {
	case syscall.AF_INET:
		addrlen = net.IPv4len
	case syscall.AF_INET6:
		addrlen = net.IPv6len
	default:
		return fmt.Errorf("Unknown address family: %d", family)
	}
	b.RouterId = data[1 : 1+addrlen]
	b.Length = data[1+addrlen]
	return nil
}

func (b *RouterIdUpdateBody) Process(client *Client, h *Header) error {
	return server.RouterIdSubscribe(client, client.vrfId)
}

func (c *Client) RouterIdUpdateNotify(mes *pb.RouterIdUpdate) {
	log.Info("zapi:SEND ", mes)
	body := &RouterIdUpdateBody{
		RouterId: mes.RouterId,
		Length:   32,
	}
	m := NewMessage(c.version, ROUTER_ID_UPDATE, body)
	s, _ := m.Serialize()
	c.conn.Write(s)
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
	HwAddr    []byte
}

func NewInterfaceUpdateBodyPb(mes *pb.InterfaceUpdate) *InterfaceUpdateBody {
	body := &InterfaceUpdateBody{
		Name:   mes.Name,
		Index:  uint32(mes.Index),
		Status: INTERFACE_ACTIVE,
		Flags:  uint64(mes.Flags),
		Metric: uint32(mes.Metric),
		Mtu:    uint32(mes.Mtu),
		Mtu6:   uint32(mes.Mtu),
	}
	if mes.HwAddr != nil {
		body.HwAddr = (*mes.HwAddr).Addr
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

func (b *InterfaceUpdateBody) DecodeFromBytes(CommandType, []byte) error {
	return nil
}

func (b *InterfaceUpdateBody) Process(client *Client, h *Header) error {
	server.InterfaceSubscribe(client, client.vrfId)

	if LocalPolicy {
		if client.version == 2 && client.routeType == ROUTE_BGP {
			log.Info("Force register BGP for LAN redistribute vrf Id ", h.VrfId)
			server.RedistSubscribe(client, false, uint32(h.VrfId), AFI_IP, RIB_STATIC)
			server.RedistSubscribe(client, false, uint32(h.VrfId), AFI_IP, RIB_CONNECTED)
			server.RedistSubscribe(client, false, uint32(h.VrfId), AFI_IP, RIB_OSPF)
			server.RedistSubscribe(client, false, uint32(h.VrfId), AFI_IP, RIB_BGP)
			server.RedistDefaultSubscribe(client, false, uint32(h.VrfId), AFI_IP)
		}
		if client.version == 3 && client.routeType == ROUTE_BGP {
			server.RedistDefaultSubscribe(client, true, 0, AFI_IP)
		}
	}

	return nil
}

func NewInterfaceAddrUpdateBodyPb(ifIndex uint32, addr *pb.Address) *InterfaceAddressUpdateBody {
	body := &InterfaceAddressUpdateBody{
		Index:  ifIndex,
		Prefix: addr.Addr.Addr,
		Length: uint8(addr.Addr.Length),
	}
	return body
}

func (c *Client) InterfaceNotify(mes *pb.InterfaceUpdate) {
	log.Info("zapi:SEND ", mes)
	m := &Message{
		Header: Header{
			Marker:  HEADER_MARKER,
			Version: c.version,
			VrfId:   0,
		},
		Body: NewInterfaceUpdateBodyPb(mes),
	}
	switch mes.Op {
	case pb.Op_InterfaceAdd, pb.Op_InterfaceNameChange, pb.Op_InterfaceMtuChange, pb.Op_InterfaceFlagChange:
		m.Header.Command = INTERFACE_ADD
		m.Send(c.conn)
	case pb.Op_InterfaceDelete:
		m.Header.Command = INTERFACE_DELETE
		m.Send(c.conn)
	case pb.Op_InterfaceUp:
		m.Header.Command = INTERFACE_UP
		m.Send(c.conn)
	case pb.Op_InterfaceDown:
		m.Header.Command = INTERFACE_DOWN
		m.Send(c.conn)
	}

	for _, addr := range mes.AddrIpv4 {
		m := &Message{
			Header: Header{
				Marker:  HEADER_MARKER,
				Version: c.version,
				VrfId:   0,
			},
			Body: NewInterfaceAddrUpdateBodyPb(mes.Index, addr),
		}
		switch mes.Op {
		case pb.Op_InterfaceAdd, pb.Op_InterfaceAddrAdd:
			m.Header.Command = INTERFACE_ADDRESS_ADD
			m.Send(c.conn)
		case pb.Op_InterfaceAddrDelete:
			m.Header.Command = INTERFACE_ADDRESS_DELETE
			m.Send(c.conn)
		}
	}
	for _, addr := range mes.AddrIpv6 {
		m := &Message{
			Header: Header{
				Marker:  HEADER_MARKER,
				Version: c.version,
				VrfId:   0,
			},
			Body: NewInterfaceAddrUpdateBodyPb(mes.Index, addr),
		}
		switch mes.Op {
		case pb.Op_InterfaceAdd, pb.Op_InterfaceAddrAdd:
			m.Header.Command = INTERFACE_ADDRESS_ADD
			m.Send(c.conn)
		case pb.Op_InterfaceAddrDelete:
			m.Header.Command = INTERFACE_ADDRESS_DELETE
			m.Send(c.conn)
		}
	}
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

func (b *InterfaceAddressUpdateBody) DecodeFromBytes(command CommandType, data []byte) error {
	if len(data) == 0 {
		return nil
	}
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

func (b *InterfaceAddressUpdateBody) Process(client *Client, h *Header) error {
	return nil
}

type RouteUpdateBody struct {
	Message  uint8
	Type     RouteType
	Flags    FLAG
	Prefix   *netutil.Prefix
	Nexthops []*Nexthop
	Distance uint8
	Metric   uint32
	PathId   uint32
	Aux      []byte
}

var HubNodeMap = map[uint32]string{}

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

func NewNexthopFromPb(n *pb.Nexthop) *Nexthop {
	return &Nexthop{
		IP:    n.Addr,
		Index: IfIndex(n.Ifindex),
	}
}

func RouteTypeFromPb(afi int, typ pb.RouteType) RouteType {
	switch typ {
	case pb.RIB_UNKNOWN:
		return ROUTE_SYSTEM
	case pb.RIB_KERNEL:
		return ROUTE_KERNEL
	case pb.RIB_CONNECTED:
		return ROUTE_CONNECT
	case pb.RIB_STATIC:
		return ROUTE_STATIC
	case pb.RIB_RIP:
		if afi == AFI_IP {
			return ROUTE_RIP
		} else {
			return ROUTE_RIPNG
		}
	case pb.RIB_OSPF:
		if afi == AFI_IP {
			return ROUTE_OSPF
		} else {
			return ROUTE_OSPF6
		}
	case pb.RIB_ISIS:
		return ROUTE_ISIS
	case pb.RIB_BGP:
		return ROUTE_BGP
	}
	return ROUTE_SYSTEM
}

var LocalPolicy = false

func PrefixFilter(p *netutil.Prefix, r *pb.Route) bool {
	if r.Type != pb.RIB_CONNECTED {
		return false
	}
	plist := server.PrefixListOut()
	if plist != nil {
		ret := plist.Match(p)
		if ret == policy.Deny {
			return true
		}
	}
	vrf := VrfLookupByIndex(r.VrfId)
	if vrf != nil && len(r.Nexthops) == 1 && r.Nexthops[0].Ifindex != 0 {
		ifp := vrf.IfLookupByIndex(IfIndex(r.Nexthops[0].Ifindex))
		if ifp != nil {
			r := regexp.MustCompile(`^veth`)
			if r.MatchString(ifp.Name) {
				return true
			}
		}
	}
	return false
}

func VrfFilter(c *Client, r *pb.Route) bool {
	if r.VrfId == 0 {
		return true
	}
	return false
}

func (c *Client) RouteNotifyGobgp(p *netutil.Prefix, r *pb.Route) {
	cmd := "add"
	if r.Op == pb.Op_RouteDelete {
		cmd = "del"
	}
	log.Infof("Executing /usr/bin/gobgp-vrf.sh vrf %d %s %s %d", r.VrfId, cmd, p.String(), r.Metric)

	aspathStr := ""
	if r.Op == pb.Op_RouteAdd && r.Aux != nil {
		aspath := &policy.ASPath{}
		aspath.DecodeFromBytes(r.Aux)
		aspathStr = aspath.String()
	}

	if hubNode, ok := HubNodeMap[r.VrfId]; ok {
		exec.Command("/usr/bin/gobgp-vrf.sh", fmt.Sprint(r.VrfId), cmd, p.String(), strconv.Itoa(int(r.Metric)), aspathStr, hubNode).Run()
	} else {
		exec.Command("/usr/bin/gobgp-vrf.sh", fmt.Sprint(r.VrfId), cmd, p.String(), strconv.Itoa(int(r.Metric)), aspathStr).Run()
	}
}

func (c *Client) RouteNotify(r *pb.Route) {
	log.Info("zapi:SEND ", r)
	p := NewPrefixFromPb(r.Prefix)

	if LocalPolicy {
		// XXX must be replaced by route-map.
		if PrefixFilter(p, r) {
			log.Info("zapi:Prefix is filtered by prefix policy")
			return
		}
		// XXX must be replaced by route-map.
		if VrfFilter(c, r) {
			log.Info("zapi:Prefix is filtered by VRF policy")
			return
		}
		if c.version == 3 {
			c.RouteNotifyGobgp(p, r)
			return
		}
	}

	body := &RouteUpdateBody{
		Type:   RouteTypeFromPb(p.AFI(), r.Type),
		Prefix: p,
	}

	for _, n := range r.Nexthops {
		body.Message = (MESSAGE_NEXTHOP | MESSAGE_IFINDEX)
		body.Nexthops = append(body.Nexthops, NewNexthopFromPb(n))
		break
	}

	if r.Op == pb.Op_RouteAdd {
		body.Message |= (MESSAGE_DISTANCE | MESSAGE_METRIC)
		body.Distance = uint8(r.Distance)
		body.Metric = r.Metric

		if r.Aux != nil {
			body.Message |= MESSAGE_ASPATH
			body.Aux = r.Aux
		}
	}

	var command CommandType
	switch r.Op {
	case pb.Op_RouteAdd:
		if p.AFI() == AFI_IP {
			command = IPV4_ROUTE_ADD
		} else {
			command = IPV6_ROUTE_ADD
		}
	case pb.Op_RouteDelete:
		if p.AFI() == AFI_IP {
			command = IPV4_ROUTE_DELETE
		} else {
			command = IPV6_ROUTE_DELETE
		}
	}

	m := &Message{
		Header: Header{
			Marker:  HEADER_MARKER,
			Version: c.version,
			VrfId:   uint16(c.vrfId),
			Command: command,
		},
		Body: body,
	}
	s, _ := m.Serialize()
	c.conn.Write(s)
}

func (b *RouteUpdateBody) Serialize() ([]byte, error) {
	if len(b.Nexthops) == 0 {
		log.Error("zapi:Serialize RouteUpdateBody does not have nexthop")
		return nil, fmt.Errorf("zapi:Serialize RouteUpdateBody does not have nexthop")
	}
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
		buf = append(buf, uint8(len(b.Nexthops)))
		if b.Nexthops[0].IP == nil {
			var len int
			if b.Prefix.AFI() == AFI_IP6 {
				len = net.IPv6len
			} else {
				len = net.IPv4len
			}
			bbuf := make([]byte, len)
			buf = append(buf, bbuf...)
		} else {
			buf = append(buf, b.Nexthops[0].IP...)
		}
	}

	if b.Message&MESSAGE_IFINDEX > 0 {
		buf = append(buf, uint8(len(b.Nexthops)))
		ifindex := make([]byte, 4)
		binary.BigEndian.PutUint32(ifindex, uint32(b.Nexthops[0].Index))
		buf = append(buf, ifindex...)
	}

	if b.Message&MESSAGE_DISTANCE > 0 {
		buf = append(buf, b.Distance)
	}

	if b.Message&MESSAGE_METRIC > 0 {
		metric := make([]byte, 4)
		binary.BigEndian.PutUint32(metric, b.Metric)
		buf = append(buf, metric...)
	}

	if b.Message&MESSAGE_ASPATH > 0 {
		aspath := &policy.ASPath{}
		aspath.DecodeFromBytes(b.Aux)
		aspath.Replace(23456, 64512)
		aux, _ := aspath.Serialize()
		bbuf := make([]byte, 4)
		binary.BigEndian.PutUint32(bbuf, uint32(len(aux)))
		buf = append(buf, bbuf...)
		buf = append(buf, aux...)
	}

	return buf, nil
}

func (b *RouteUpdateBody) DecodeFromBytes(command CommandType, data []byte) error {
	afi := AFI_IP
	if command == IPV6_ROUTE_ADD || command == IPV6_ROUTE_DELETE {
		afi = AFI_IP6
	}
	b.Type = RouteType(data[0])
	b.Flags = FLAG(data[1])
	b.Message = data[2]
	// binary.BigEndian.Uint16(data[3:5]) -- SAFI is not used.

	b.Prefix = netutil.NewPrefixAFI(afi)
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
			b.Nexthops = append(b.Nexthops, nexthop)
		}
	}

	if b.Message&MESSAGE_DISTANCE > 0 {
		b.Distance = data[pos]
		pos += 1
	}

	if b.Message&MESSAGE_METRIC > 0 {
		b.Metric = binary.BigEndian.Uint32(data[pos : pos+4])
		pos += 4
	}

	if b.Message&MESSAGE_PATH_ID > 0 {
		b.PathId = binary.BigEndian.Uint32(data[pos : pos+4])
		pos += 4
	}

	if b.Message&MESSAGE_ASPATH > 0 {
		data = data[pos:]
		aspathLen := binary.BigEndian.Uint32(data[:4])
		data = data[4:]
		if int(aspathLen) == len(data) {
			aspath := &policy.ASPath{}
			aspath.DecodeFromBytes(data)
			aspath.Replace(64512, 23456)
			aux, _ := aspath.Serialize()
			b.Aux = aux
		} else {
			log.Warnf("zapi:ASPath len %d is different with data len %d", aspathLen, len(data))
		}
	}

	return nil
}

func DistributeListOspf(dlistName string, p *netutil.Prefix, ri *Rib) policy.Action {
	// Lookup primary
	primary := server.PrefixListLookup(dlistName + "-primary")
	fmt.Println("plist primary", dlistName+"-primary", primary)
	if primary != nil && primary.Match(p) {
		ri.SetFlag(RIB_FLAG_DISTANCE)
		ri.Distance = 180
		return policy.Permit
	}

	backup := server.PrefixListLookup(dlistName + "-backup")
	fmt.Println("plist backup", dlistName+"-backup", backup)
	if backup != nil && backup.Match(p) {
		ri.SetFlag(RIB_FLAG_DISTANCE)
		ri.Distance = 180
		//ri.Metric +=
		return policy.Permit
	}

	return policy.Deny
}

func (b *RouteUpdateBody) Process(client *Client, h *Header) error {
	if DefaultVrfProtect && h.Version == 3 && h.VrfId == 0 {
		return nil
	}

	if len(b.Nexthops) == 1 && b.Nexthops[0].IP.Equal(net.IPv4zero.To4()) {
		return nil
	}

	// Prepare RibInfo
	vrf := VrfLookupByIndex(uint32(h.VrfId))
	if vrf == nil {
		fmt.Println("Can't find VRF id:", h.VrfId)

		// Try to create VRF.
		server.VrfAdd(fmt.Sprintf("vrf%d", h.VrfId))
		vrf = VrfLookupByIndex(uint32(h.VrfId))
		if vrf == nil {
			fmt.Println("Couldn't create VRF id:", h.VrfId)
			return nil
		}
	}

	ri := &Rib{
		Type:     RouteType2RibType(b.Type),
		Nexthops: b.Nexthops,
		Src:      client,
		Metric:   b.Metric,
		Distance: b.Distance,
		Aux:      b.Aux,
		PathId:   b.PathId,
	}
	if ri.Distance != 0 {
		ri.SetFlag(RIB_FLAG_DISTANCE)
	}

	// OSPF route-map.
	// if LocalPolicy {
	// 	OspfRouteMap(ri)
	// }
	if ri.Type == RIB_OSPF && vrf.DListOspf != "" {
		action := DistributeListOspf(vrf.DListOspf, b.Prefix, ri)
		if action == policy.Deny {
			fmt.Println("OSPF route filtered by distribute-list", b.Prefix)
			return nil
		}
	}

	// Call RIB API.
	if h.Command == IPV4_ROUTE_ADD {
		fmt.Println("Route add", b.Prefix, b.Nexthops)
		vrf.RibAdd(b.Prefix, ri)
	} else {
		fmt.Println("Route delete", b.Prefix)
		vrf.RibDelete(b.Prefix, ri)
	}
	return nil
}

// Redistribute routes.
type RedistributeBody struct {
	Type uint8
}

func (b *RedistributeBody) Serialize() ([]byte, error) {
	buf := make([]byte, 1)
	buf[0] = b.Type
	return buf, nil
}

func (b *RedistributeBody) DecodeFromBytes(command CommandType, data []byte) error {
	b.Type = data[0]
	return nil
}

func RibTypeFromRouteType(typ uint8) (uint8, int) {
	switch RouteType(typ) {
	case ROUTE_KERNEL:
		return RIB_KERNEL, AFI_IP
	case ROUTE_CONNECT:
		return RIB_CONNECTED, AFI_IP
	case ROUTE_STATIC:
		return RIB_STATIC, AFI_IP
	case ROUTE_RIP:
		return RIB_RIP, AFI_IP
	case ROUTE_RIPNG:
		return RIB_RIP, AFI_IP6
	case ROUTE_OSPF:
		return RIB_OSPF, AFI_IP
	case ROUTE_OSPF6:
		return RIB_OSPF, AFI_IP6
	case ROUTE_ISIS:
		return RIB_ISIS, AFI_IP
	case ROUTE_BGP:
		return RIB_BGP, AFI_IP
	default:
		return RIB_UNKNOWN, AFI_IP
	}
	return RIB_UNKNOWN, AFI_IP
}

func (b *RedistributeBody) Process(client *Client, h *Header) error {
	log.Info("zapi:", h.Command.String(), " ", RouteTypeStringMap[RouteType(b.Type)])

	typ, afi := RibTypeFromRouteType(b.Type)
	if typ == RIB_UNKNOWN {
		return fmt.Errorf("zapi:Unknown redistribute route type")
	}
	// When client version is 3, assumes allVrf true.
	allVrf := false
	if client.version == 3 {
		allVrf = true
	}

	// AFI is determined.
	if typ == RIB_RIP || typ == RIB_OSPF {
		server.RedistSubscribe(client, allVrf, uint32(h.VrfId), afi, typ)
		return nil
	}
	// Guess client's address family from route type in Hello message.
	switch client.routeType {
	case ROUTE_RIP, ROUTE_OSPF:
		server.RedistSubscribe(client, allVrf, uint32(h.VrfId), AFI_IP, typ)
	case ROUTE_RIPNG, ROUTE_OSPF6:
		server.RedistSubscribe(client, allVrf, uint32(h.VrfId), AFI_IP6, typ)
	case ROUTE_BGP, ROUTE_ISIS:
		server.RedistSubscribe(client, allVrf, uint32(h.VrfId), AFI_IP, typ)
		server.RedistSubscribe(client, allVrf, uint32(h.VrfId), AFI_IP6, typ)
	default:
		// Do nothing for unknown client.
	}
	return nil
}

// Redistribute default message. Message body is empty. When redistribute
// default is on, it effective to both IPv4 and IPv6.
type RedistributeDefaultBody struct {
}

func (b *RedistributeDefaultBody) Serialize() ([]byte, error) {
	return nil, nil
}

func (b *RedistributeDefaultBody) DecodeFromBytes(command CommandType, data []byte) error {
	return nil
}

func (b *RedistributeDefaultBody) Process(client *Client, h *Header) error {
	switch h.Command {
	case REDISTRIBUTE_DEFAULT_ADD:
		if client.routeType == ROUTE_RIP || client.routeType == ROUTE_OSPF || client.routeType == ROUTE_BGP {
			server.RedistDefaultSubscribe(client, client.allVrf, client.vrfId, AFI_IP)
		}
		if client.routeType == ROUTE_RIPNG || client.routeType == ROUTE_OSPF6 || client.routeType == ROUTE_BGP {
			server.RedistDefaultSubscribe(client, client.allVrf, client.vrfId, AFI_IP6)
		}
	case REDISTRIBUTE_DEFAULT_DELETE:
		if client.routeType == ROUTE_RIP || client.routeType == ROUTE_OSPF || client.routeType == ROUTE_BGP {
			server.RedistDefaultUnsubscribe(client, client.allVrf, client.vrfId, AFI_IP)
		}
		if client.routeType == ROUTE_RIPNG || client.routeType == ROUTE_OSPF6 || client.routeType == ROUTE_BGP {
			server.RedistDefaultUnsubscribe(client, client.allVrf, client.vrfId, AFI_IP6)
		}
	}
	return nil
}

// Nexthop lookup for IPv4 routes.
type IPv4NexthopLookupBody struct {
	Addr net.IP
}

func (b *IPv4NexthopLookupBody) Serialize() ([]byte, error) {
	buf := make([]byte, 4)
	copy(buf, b.Addr)
	return buf, nil
}

func (b *IPv4NexthopLookupBody) DecodeFromBytes(command CommandType, data []byte) error {
	b.Addr = make([]byte, 4)
	copy(b.Addr, data[:4])
	return nil
}

func (b *IPv4NexthopLookupBody) Process(client *Client, h *Header) error {
	vrf := VrfLookupByIndex(uint32(h.VrfId))
	if vrf == nil {
		fmt.Println("[zapi]IPv4NexthopLookup: Can't find VRF with id", h.VrfId)
		return nil
	}

	reply := &IPv4NexthopReplyBody{
		Addr: b.Addr,
	}

	found := true
	if NexthopLookupHook != nil {
		found = NexthopLookupHook(vrf, b.Addr)
	}

	if found {
		nexthop := NewNexthopAddr(b.Addr)
		reply.Nexthops = append(reply.Nexthops, nexthop)
		reply.NexthopNum = uint8(len(reply.Nexthops))
	}

	m := &Message{
		Header: Header{
			Marker:  HEADER_MARKER,
			Version: h.Version,
			VrfId:   uint16(h.VrfId),
			Command: IPV4_NEXTHOP_LOOKUP,
		},
		Body: reply,
	}
	s, _ := m.Serialize()

	client.conn.Write(s)
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
	return buf, nil
}

func (b *IPv4NexthopReplyBody) DecodeFromBytes(CommandType, []byte) error {
	return nil
}

func (b *IPv4NexthopReplyBody) Process(*Client, *Header) error {
	return nil
}

var NexthopLookupHook func(vrf *Vrf, nexthop net.IP) bool

func EsiNexthopLookup(vrf *Vrf, nexthop net.IP) bool {
	ptree := vrf.ribTable[AFI_IP]
	fmt.Println("EsiNexthopLookup", nexthop)
	n := ptree.Match(nexthop, 32)
	if n != nil {
		if n.Item != nil {
			for _, rib := range n.Item.(RibSlice) {
				if rib.IsFib() && rib.Type == RIB_CONNECTED {
					if len(rib.Nexthops) == 1 && rib.Nexthops[0].IsIfOnly() {
						nhop := rib.Nexthops[0]
						ifp := IfLookupByIndex(nhop.Index)
						fmt.Println("EsiNexthopLookup: ifp", ifp.Name)
						r := regexp.MustCompile(`sproute\d+`)
						if r.MatchString(ifp.Name) {
							result := DtlsNexthop(vrf.Id, nexthop)
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

func (c *Client) HandleRequest(conn net.Conn, vrfId uint32) {
	defer conn.Close()

	var version byte
	for {
		data := make([]byte, HeaderSize(version))
		_, err := conn.Read(data)
		if err != nil {
			break
		}

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
		len := int(h.Length) - HeaderSize(version)

		if vrfId != 0 {
			h.VrfId = uint16(vrfId)
		}

		err = HandleBody(c, &h, len)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			break
		}
	}
	fmt.Println("[zapi] disconnected", "vrf", vrfId, "version", version)
	server.WatcherUnsubscribe(c)
	ClientUnregister(conn)
}

func HandleBody(client *Client, h *Header, len int) error {
	log.Infof("zapi:%s version %d vrf %d len %d", h.Command.String(), h.Version, h.VrfId, len)

	var data []byte
	if len > 0 {
		data = make([]byte, len)
		_, err := client.conn.Read(data)
		if err != nil {
			return err
		}
	}

	var body Body
	switch h.Command {
	case HELLO:
		body = &HelloBody{}
	case ROUTER_ID_ADD:
		body = &RouterIdUpdateBody{}
	case INTERFACE_ADD:
		body = &InterfaceUpdateBody{}
	case IPV4_ROUTE_ADD, IPV4_ROUTE_DELETE:
		body = &RouteUpdateBody{}
	case IPV4_NEXTHOP_LOOKUP:
		body = &IPv4NexthopLookupBody{}
	case REDISTRIBUTE_ADD, REDISTRIBUTE_DELETE:
		body = &RedistributeBody{}
	case REDISTRIBUTE_DEFAULT_ADD, REDISTRIBUTE_DEFAULT_DELETE:
		body = &RedistributeDefaultBody{}
	default:
		log.Infof("zapi:Unhandled command %s, skipping", h.Command.String())
		return nil
	}

	err := body.DecodeFromBytes(h.Command, data)
	if err != nil {
		return err
	}
	err = body.Process(client, h)
	if err != nil {
		return err
	}
	return nil
}

type ZServer struct {
	Path   string
	Listen net.Listener
	VrfId  uint32
}

func ZServerStart(typ string, path string, vrfId uint32) *ZServer {
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
		log.Infof("zapi:Server started at %s", path)
		for {
			// Listen for an incoming connection.
			conn, err := lis.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				return
			}

			// Register client.
			client := ClientRegister(conn)

			// Handle connections in a new go routine.
			go client.HandleRequest(conn, vrfId)
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
