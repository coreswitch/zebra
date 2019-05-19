// Copyright 2017 zebra project
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

package bgp

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"

	"github.com/coreswitch/netutil"
)

const (
	BGP_MARKER_LEN     = 16
	BGP_HEADER_LEN     = 19
	BGP_MAX_PACKET_LEN = 4096
)

const (
	BGP_MSG_OPEN              = 1
	BGP_MSG_UPDATE            = 2
	BGP_MSG_NOTIFICATION      = 3
	BGP_MSG_KEEPALIVE         = 4
	BGP_MSG_ROUTE_REFRESH     = 5
	BGP_MSG_CAPABILITY        = 6
	BGP_MSG_ROUTE_REFRESH_OLD = 128
)

type BgpHeader struct {
	Marker []byte
	Len    uint16
	Type   uint8
}

func (msg *BgpHeader) DecodeFromBytes(data []byte) error {
	if len(data) < BGP_HEADER_LEN {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	for i := 0; i < BGP_MARKER_LEN; i++ {
		if data[i] != 0xff {
			return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_CONNECTION_NOT_SYNCHRONIZED, nil, "")
		}
	}
	msg.Len = binary.BigEndian.Uint16(data[16:18])
	if int(msg.Len) < BGP_HEADER_LEN {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	msg.Type = data[18]
	switch msg.Type {
	case BGP_MSG_OPEN:
	case BGP_MSG_UPDATE:
	case BGP_MSG_NOTIFICATION:
	case BGP_MSG_KEEPALIVE:
		if msg.Len != BGP_HEADER_LEN {
			return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
		}
	case BGP_MSG_ROUTE_REFRESH:
	case BGP_MSG_CAPABILITY:
	case BGP_MSG_ROUTE_REFRESH_OLD:
	default:
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_TYPE, nil, "")
	}
	return nil
}

func (msg *BgpHeader) Serialize() ([]byte, error) {
	buf := make([]byte, BGP_HEADER_LEN)
	for i, _ := range buf[:16] {
		buf[i] = 0xff
	}
	binary.BigEndian.PutUint16(buf[16:18], msg.Len)
	buf[18] = msg.Type
	return buf, nil
}

type BgpBody interface {
	DecodeFromBytes([]byte) error
	Serialize() ([]byte, error)
}

type BgpMessage struct {
	Header BgpHeader
	Body   BgpBody
}

func (msg *BgpMessage) Type() int {
	return int(msg.Header.Type)
}

type BgpOpen struct {
	Version     uint8
	As          uint32
	HoldTime    uint16
	RouterId    net.IP
	OptParamLen uint8
	OptParams   []OptParamInterface
}

const BGP_OPT_CAPABILITY = 2

type OptParamInterface interface {
	Serialize() ([]byte, error)
}

type OptParamCapability struct {
	OptParamType uint8
	OptParamLen  uint8
	OptCaps      []CapabilityInterface
}

func (o *OptParamCapability) Serialize() ([]byte, error) {
	buf := make([]byte, 2)
	buf[0] = o.OptParamType
	for _, cap := range o.OptCaps {
		cbuf, err := cap.Serialize()
		if err != nil {
			return nil, err
		}
		buf = append(buf, cbuf...)
	}
	o.OptParamLen = uint8(len(buf) - 2)
	buf[1] = o.OptParamLen
	return buf, nil
}

const (
	BGP_CAP_MULTIPROTOCOL                   uint8 = 1  // RFC2858
	BGP_CAP_ROUTE_REFRESH                         = 2  // RFC2918
	BGP_CAP_ORF                                   = 3  // RFC5291
	BGP_CAP_CARRYING_LABEL_INFO                   = 4  // RFC3107
	BGP_CAP_EXT_NEXTHOP                           = 5  // RFC5549
	BGP_CAP_GRACEFUL_RESTART                      = 64 // RFC4724
	BGP_CAP_4OCTET_AS_NUMBER                      = 65 // RFC6793
	BGP_CAP_DYNAMIC_CAP                           = 67 // draft-ietf-idr-dynamic-cap
	BGP_CAP_MULTI_SESSION                         = 68 // draft-ietf-idr-bgp-multisession
	BGP_CAP_ADD_PATH                              = 69 // RFC7911
	BGP_CAP_ENHANCED_ROUTE_REFRESH                = 70 // RFC7313
	BGP_CAP_LONG_LIVED_GRACEFUL_RESTART           = 71 // draft-uttaro-idr-bgp-persistence
	BGP_CAP_ROUTE_REFRESH_CISCO                   = 128
	BGP_CAP_LONG_LIVED_GRACEFUL_RESTART_OLD       = 129 // draft-uttaro-idr-bgp-persistence
)

var CapabilityCode2String = map[uint8]string{
	BGP_CAP_MULTIPROTOCOL:               "Muitiprotocol",
	BGP_CAP_ROUTE_REFRESH:               "RouteRefresh",
	BGP_CAP_CARRYING_LABEL_INFO:         "Carrying Label Info",
	BGP_CAP_GRACEFUL_RESTART:            "Graceful Restart",
	BGP_CAP_4OCTET_AS_NUMBER:            "4 Octet AS",
	BGP_CAP_ADD_PATH:                    "Add Path",
	BGP_CAP_ENHANCED_ROUTE_REFRESH:      "Enhanced Route Refresh",
	BGP_CAP_ROUTE_REFRESH_CISCO:         "Route Refresh Cisco",
	BGP_CAP_LONG_LIVED_GRACEFUL_RESTART: "Long Lived Graceful Restart",
}

func CapabilityCodeString(code uint8) string {
	str, ok := CapabilityCode2String[code]
	if !ok {
		return "Unknown"
	}
	return str
}

type CapabilityInterface interface {
	DecodeFromBytes([]byte) error
	Serialize() ([]byte, error)
	Length() int
}

type CapabilityBase struct {
	CapCode  uint8
	CapLen   uint8
	CapValue []byte
}

func (c *CapabilityBase) Length() int {
	return int(c.CapLen + 2)
}

func (c *CapabilityBase) DecodeFromBytes(data []byte) error {
	if len(data) < 2 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	c.CapCode = data[0]
	c.CapLen = data[1]
	totalLen := int(c.CapLen) + 2
	if len(data) < totalLen {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	if c.CapLen > 0 {
		c.CapValue = data[2:totalLen]
	}
	return nil
}

func (c *CapabilityBase) Serialize() ([]byte, error) {
	c.CapLen = uint8(len(c.CapValue))
	buf := make([]byte, 2)
	buf[0] = c.CapCode
	buf[1] = c.CapLen
	buf = append(buf, c.CapValue...)
	return buf, nil
}

type CapMultiProtocol struct {
	CapabilityBase
	AfiSafi AfiSafi
}

func (c *CapMultiProtocol) DecodeFromBytes(data []byte) error {
	err := c.CapabilityBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	if c.CapLen != 4 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	data = data[2:]
	if len(data) < 4 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	c.AfiSafi = AfiSafiValue(Afi(binary.BigEndian.Uint16(data[0:2])), Safi(data[3]))
	return nil
}

func (c *CapMultiProtocol) Serialize() ([]byte, error) {
	c.CapCode = BGP_CAP_MULTIPROTOCOL
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf[0:], uint32(c.AfiSafi))
	c.CapValue = buf
	return c.CapabilityBase.Serialize()
}

type CapRouteRefresh struct {
	CapabilityBase
}

func (c *CapRouteRefresh) Serialize() ([]byte, error) {
	c.CapCode = BGP_CAP_ROUTE_REFRESH
	return c.CapabilityBase.Serialize()
}

type CapCarryingLabelInfo struct {
	CapabilityBase
}

type CapGracefulRestartTuple struct {
	Afi   Afi
	Safi  Safi
	Flags uint8
}

const RESTART_FLAG = 0x08
const FORWARDING_PRESERVED_FLAG = 0xf0

type CapGracefulRestart struct {
	CapabilityBase
	Flags  uint8  // 4bit
	Time   uint16 // 12bit
	Tuples []*CapGracefulRestartTuple
}

func (c *CapGracefulRestart) DecodeFromBytes(data []byte) error {
	err := c.CapabilityBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	data = data[2:]
	if len(data) < 2 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	value := binary.BigEndian.Uint16(data[0:2])
	c.Flags = uint8(value >> 12)
	c.Time = value & 0xfff
	data = data[2:]

	remaining := int(c.CapLen) - 2
	if (remaining % 4) != 0 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	if remaining > len(data) {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	for i := 0; i < remaining/4; i++ {
		tuple := &CapGracefulRestartTuple{
			Afi:   Afi(binary.BigEndian.Uint16(data[0:2])),
			Safi:  Safi(data[2]),
			Flags: data[3],
		}
		c.Tuples = append(c.Tuples, tuple)
		data = data[4:]
	}
	return nil
}

func (c *CapGracefulRestart) Serialize() ([]byte, error) {
	c.CapCode = BGP_CAP_GRACEFUL_RESTART
	val := uint16(c.Flags)
	val = val << 12
	val |= c.Time
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf[0:], val)
	for _, tuple := range c.Tuples {
		tbuf := make([]byte, 4)
		binary.BigEndian.PutUint16(tbuf[0:2], uint16(tuple.Afi))
		tbuf[2] = uint8(tuple.Safi)
		tbuf[3] = tuple.Flags
		buf = append(buf, tbuf...)
	}
	c.CapValue = buf
	return c.CapabilityBase.Serialize()
}

type Cap4OctetAsNumber struct {
	CapabilityBase
	As uint32
}

func (c *Cap4OctetAsNumber) DecodeFromBytes(data []byte) error {
	err := c.CapabilityBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	data = data[2:]
	if c.CapLen != 4 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	if len(data) < 4 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	c.As = binary.BigEndian.Uint32(data[0:4])
	fmt.Println("    4 Octet As:", c.As)
	return nil
}

func (c *Cap4OctetAsNumber) Serialize() ([]byte, error) {
	c.CapCode = BGP_CAP_4OCTET_AS_NUMBER
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf[0:], c.As)
	c.CapValue = buf
	return c.CapabilityBase.Serialize()
}

const (
	BGP_ADD_PATH_RECEIVE uint8 = 1
	BGP_ADD_PATH_SEND          = 2
	BGP_ADD_PATH_BOTH          = 3
)

type CapAddPath struct {
	CapabilityBase
	AfiSafi     AfiSafi
	SendRecieve uint8
}

func (c *CapAddPath) DecodeFromBytes(data []byte) error {
	err := c.CapabilityBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	data = data[2:]
	if c.CapLen != 4 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	if len(data) < 4 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	c.AfiSafi = AfiSafiValue(Afi(binary.BigEndian.Uint16(data[0:2])), Safi(data[2]))
	c.SendRecieve = data[3]
	fmt.Println("AfiSafi:", c.AfiSafi, "SendRecieve", c.SendRecieve)
	return nil
}

type CapEnhancedRouteRefresh struct {
	CapabilityBase
}

type CapRouteRefreshCisco struct {
	CapabilityBase
}

type CapLongLivedGracefulRestartTuple struct {
	Afi         Afi
	Safi        Safi
	Flags       uint8
	RestartTime uint32
}

type CapLongLivedGracefulRestart struct {
	CapabilityBase
	Tuples []*CapLongLivedGracefulRestartTuple
}

func (c *CapLongLivedGracefulRestart) DecodeFromBytes(data []byte) error {
	err := c.CapabilityBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	data = data[2:]
	if len(data) < 2 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	remaining := int(c.CapLen) - 2
	if (remaining % 7) != 0 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	if remaining > len(data) {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	for i := 0; i < remaining/7; i++ {
		tuple := &CapLongLivedGracefulRestartTuple{
			Afi:         Afi(binary.BigEndian.Uint16(data[0:2])),
			Safi:        Safi(data[2]),
			Flags:       data[3],
			RestartTime: uint32(data[4])<<16 | uint32(data[5])<<8 | uint32(data[6]),
		}
		c.Tuples = append(c.Tuples, tuple)
		data = data[7:]
	}
	return nil
}

type CapUnknown struct {
	CapabilityBase
}

type Capability struct {
	OptType    uint8
	OptLen     uint8
	Capability []CapabilityInterface
}

func DecodeCapability(data []byte) (CapabilityInterface, error) {
	if len(data) < 2 {
		return nil, NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	var c CapabilityInterface
	fmt.Printf("  Cap code: %s(%d)\n", CapabilityCodeString(data[0]), data[0])
	switch uint8(data[0]) {
	case BGP_CAP_MULTIPROTOCOL:
		c = &CapMultiProtocol{}
	case BGP_CAP_ROUTE_REFRESH:
		c = &CapRouteRefresh{}
	case BGP_CAP_CARRYING_LABEL_INFO:
		c = &CapCarryingLabelInfo{}
	case BGP_CAP_GRACEFUL_RESTART:
		c = &CapGracefulRestart{}
	case BGP_CAP_4OCTET_AS_NUMBER:
		c = &Cap4OctetAsNumber{}
	case BGP_CAP_ADD_PATH:
		c = &CapAddPath{}
	case BGP_CAP_ENHANCED_ROUTE_REFRESH:
		c = &CapEnhancedRouteRefresh{}
	case BGP_CAP_ROUTE_REFRESH_CISCO:
		c = &CapRouteRefreshCisco{}
	case BGP_CAP_LONG_LIVED_GRACEFUL_RESTART:
		c = &CapLongLivedGracefulRestart{}
	default:
		c = &CapUnknown{}
	}
	err := c.DecodeFromBytes(data)
	return c, err
}

func (c *Capability) DecodeFromBytes(data []byte) error {
	for len(data) > 0 {
		cap, err := DecodeCapability(data)
		if err != nil {
			return err
		}
		c.Capability = append(c.Capability, cap)
		data = data[cap.Length():]
	}
	return nil
}

func (c *Capability) Serialize() ([]byte, error) {
	return nil, nil
}

type OptParamUnknown struct {
	OptType  uint8
	OptLen   uint8
	OptValue []byte
}

func (o *OptParamUnknown) Serialize() ([]byte, error) {
	return nil, nil
}

func (msg *BgpOpen) DecodeFromBytes(data []byte) error {
	msg.Version = data[0]
	msg.As = uint32(binary.BigEndian.Uint16(data[1:3]))
	msg.HoldTime = binary.BigEndian.Uint16(data[3:5])
	msg.RouterId = net.IP(data[5:9]).To4()
	fmt.Println("  Version:", msg.Version)
	fmt.Println("  As:", msg.As)
	fmt.Println("  HoldTime:", msg.HoldTime)
	fmt.Println("  Router Id:", msg.RouterId)

	if msg.HoldTime != 0 && msg.HoldTime < 3 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_UNACCEPTABLE_HOLD_TIME, nil, "")
	}

	msg.OptParamLen = data[9]
	data = data[10:]
	if len(data) < int(msg.OptParamLen) {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	remaining := msg.OptParamLen
	for remaining > 0 {
		if remaining < 2 {
			return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
		}
		typ := data[0]
		len := data[1]
		if remaining < len+2 {
			return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
		}
		remaining -= len + 2
		if typ == BGP_OPT_CAPABILITY {
			cap := &Capability{}
			cap.OptType = typ
			cap.OptLen = len
			err := cap.DecodeFromBytes(data[2 : 2+len])
			if err != nil {
				return err
			}
			msg.OptParams = append(msg.OptParams, cap)
		} else {
			opt := &OptParamUnknown{}
			opt.OptType = typ
			opt.OptLen = len
			opt.OptValue = data[2 : 2+len]
			msg.OptParams = append(msg.OptParams, opt)
		}
		data = data[2+len:]
	}
	return nil
}

func (msg *BgpOpen) Serialize() ([]byte, error) {
	buf := make([]byte, 10) // Version(1) + AS(2) + HoldTime(2) + RouterId(4) + OptParamLen(1)
	buf[0] = msg.Version
	var as uint16 = AS_TRANS
	if msg.As <= 65535 {
		as = uint16(msg.As)
	}
	binary.BigEndian.PutUint16(buf[1:3], as)
	binary.BigEndian.PutUint16(buf[3:5], msg.HoldTime)
	copy(buf[5:9], msg.RouterId)
	obuf := make([]byte, 0)
	for _, p := range msg.OptParams {
		pbuf, err := p.Serialize()
		if err != nil {
			return nil, err
		}
		obuf = append(obuf, pbuf...)
	}
	msg.OptParamLen = uint8(len(obuf))
	fmt.Println("msg.OptParamLen", msg.OptParamLen)
	buf[9] = msg.OptParamLen
	return append(buf, obuf...), nil
}

type BgpUpdate struct {
	WithdrawLen  uint16
	Withdraws    []*PrefixNLRI
	TotalAttrLen uint16
	Attrs        []AttrInterface `json:"attrs"`
	NLRI         []*PrefixNLRI
}

type PrefixNLRI struct {
	netutil.Prefix
}

func (p *PrefixNLRI) TotalLength() int {
	return p.ByteLength() + 1
}

func (p *PrefixNLRI) DecodeFromBytes(data []byte) error {
	if len(data) < 1 {
		NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_INVALID_NETWORK_FIELD, nil, "")
	}
	p.Length = int(data[0])
	data = data[1:]
	byteLen := p.ByteLength()
	if len(data) < byteLen {
		NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_INVALID_NETWORK_FIELD, nil, "")
	}
	p.IP = make([]byte, 4)
	copy(p.IP, data[:byteLen])

	return nil
}

func (msg *BgpUpdate) DecodeFromBytes(data []byte) error {
	code := BGP_ERR_UPDATE_MESSAGE_ERROR
	subCode := BGP_ERR_SUB_MALFORMED_ATTRIBUTE_LIST

	// Withdraw.
	if len(data) < 2 {
		return NewBgpError(code, subCode, nil, "")
	}
	msg.WithdrawLen = binary.BigEndian.Uint16(data[0:2])
	data = data[2:]
	if len(data) < int(msg.WithdrawLen) {
		return NewBgpError(code, subCode, nil, "")
	}
	remaining := int(msg.WithdrawLen)
	for remaining > 0 {
		w := &PrefixNLRI{}
		err := w.DecodeFromBytes(data)
		if err != nil {
			return err
		}
		remaining -= w.TotalLength()
		data = data[w.TotalLength():]

		fmt.Println("Withdraw", w.Prefix)
		msg.Withdraws = append(msg.Withdraws, w)
	}

	// Attributes.
	if len(data) < 2 {
		return NewBgpError(code, subCode, nil, "")
	}
	msg.TotalAttrLen = binary.BigEndian.Uint16(data[0:2])
	data = data[2:]
	if len(data) < int(msg.TotalAttrLen) {
		return NewBgpError(code, subCode, nil, "")
	}
	remaining = int(msg.TotalAttrLen)
	for remaining > 0 {
		if len(data) < 2 {
			return NewBgpError(code, subCode, nil, "")
		}
		attr := NewAttrByType(data[1])
		err := attr.DecodeFromBytes(data)
		if err != nil {
			return err
		}
		//fmt.Println(attr)
		remaining -= attr.TotalLength()
		data = data[attr.TotalLength():]
		msg.Attrs = append(msg.Attrs, attr)
	}
	byte, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
	} else {
		fmt.Println(string(byte))
	}

	// NLRI.
	remaining = len(data)
	for remaining > 0 {
		n := &PrefixNLRI{}
		err := n.DecodeFromBytes(data)
		if err != nil {
			return err
		}
		remaining -= n.TotalLength()
		data = data[n.TotalLength():]

		//fmt.Println("NLRI", n.Prefix)
		msg.NLRI = append(msg.NLRI, n)
	}

	return nil
}

func (msg *BgpUpdate) Serialize() ([]byte, error) {
	return nil, nil
}

type BgpNotification struct {
	Code    uint8
	SubCode uint8
	Data    []byte
	Message string
}

func (e *BgpNotification) Error() string {
	return e.Message
}

func NewBgpError(code, subCode uint8, data []byte, message string) error {
	return &BgpNotification{
		Code:    code,
		SubCode: subCode,
		Data:    data,
		Message: message,
	}
}

func NewBgpNotificationMsg(notify *BgpNotification) *BgpMessage {
	msg := &BgpMessage{
		Header: BgpHeader{Type: BGP_MSG_NOTIFICATION},
		Body:   notify,
	}
	return msg
}

func (msg *BgpNotification) DecodeFromBytes(data []byte) error {
	// Then minimum length is already checked at BgpHeader DecodeFromBytes().
	// if len(data) < 2 {
	// 	return nil
	// }
	msg.Code = data[0]
	msg.SubCode = data[1]
	if len(data) > 2 {
		msg.Data = data[2:]
	}
	return nil
}

func (msg *BgpNotification) Serialize() ([]byte, error) {
	buf := make([]byte, 2)
	buf[0] = msg.Code
	buf[1] = msg.SubCode
	buf = append(buf, msg.Data...)
	return buf, nil
}

type BgpKeepAlive struct {
}

func (msg *BgpKeepAlive) DecodeFromBytes(data []byte) error {
	return nil
}

func (msg *BgpKeepAlive) Serialize() ([]byte, error) {
	return make([]byte, 0), nil
}

type BgpRouteRefresh struct {
	Afi         uint16
	Demarcation uint8
	Safi        uint8
}

func (msg *BgpRouteRefresh) DecodeFromBytes(data []byte) error {
	// XXX This may not happen.
	if len(data) < 4 {
		return NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_LENGTH, nil, "")
	}
	msg.Afi = binary.BigEndian.Uint16(data[0:2])
	msg.Demarcation = data[2]
	msg.Safi = data[3]
	return nil
}

func (msg *BgpRouteRefresh) Serialize() ([]byte, error) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint16(buf[0:2], msg.Afi)
	buf[2] = msg.Demarcation
	buf[3] = msg.Safi
	return buf, nil
}

type BgpCapability struct {
}

func (msg *BgpCapability) DecodeFromBytes(data []byte) error {
	return nil
}

func (msg *BgpCapability) Serialize() ([]byte, error) {
	return nil, nil
}

func ParseBgpBody(header *BgpHeader, buf []byte) (*BgpMessage, error) {
	msg := &BgpMessage{Header: *header}
	switch msg.Header.Type {
	case BGP_MSG_OPEN:
		fmt.Println("[Open:Recv]")
		msg.Body = &BgpOpen{}
	case BGP_MSG_UPDATE:
		fmt.Println("[Update:Recv]")
		msg.Body = &BgpUpdate{}
	case BGP_MSG_NOTIFICATION:
		fmt.Println("[Notification:Recv]")
		msg.Body = &BgpNotification{}
	case BGP_MSG_KEEPALIVE:
		fmt.Println("[Keepalive:Recv]")
		msg.Body = &BgpKeepAlive{}
	case BGP_MSG_ROUTE_REFRESH, BGP_MSG_ROUTE_REFRESH_OLD:
		fmt.Println("[Route Refresh:Recv]")
		msg.Body = &BgpRouteRefresh{}
	case BGP_MSG_CAPABILITY:
		fmt.Println("[Capability:Recv]")
		msg.Body = &BgpCapability{}
	default:
		// This won't happen because BgpHeader DecodeFromBytes() already check the type.
		return nil, NewBgpError(BGP_ERR_MSG_HEADER_ERROR, BGP_ERR_SUB_BAD_MESSAGE_TYPE, nil, "")
	}
	err := msg.Body.DecodeFromBytes(buf)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (msg *BgpMessage) Serialize() ([]byte, error) {
	body, err := msg.Body.Serialize()
	if err != nil {
		return nil, err
	}
	if msg.Header.Len == 0 {
		len := len(body) + BGP_HEADER_LEN
		if len > BGP_MAX_PACKET_LEN {
			// XXX BGP Internal Error.
			return nil, err
		}
		msg.Header.Len = uint16(len)
	}
	header, err := msg.Header.Serialize()
	if err != nil {
		// XXX BGP Internal Error.
		return nil, err
	}
	return append(header, body...), nil
}

func NewBgpOpenMsg(n *Neighbor) *BgpMessage {
	body := &BgpOpen{
		Version:  4,
		As:       n.LocalAs(),
		HoldTime: n.HoldTime(),
		RouterId: n.server.RouterId(),
	}
	if !n.dontCapAll {
		optParam := &OptParamCapability{OptParamType: BGP_OPT_CAPABILITY}
		body.OptParams = append(body.OptParams, optParam)
		for afiSafi, _ := range n.afiSafi {
			cap := &CapMultiProtocol{AfiSafi: afiSafi}
			optParam.OptCaps = append(optParam.OptCaps, cap)
		}
		if !n.dontCap4As {
			cap := &Cap4OctetAsNumber{As: n.LocalAs()}
			optParam.OptCaps = append(optParam.OptCaps, cap)
		}
		if !n.dontCapRefresh {
			cap := &CapRouteRefresh{}
			optParam.OptCaps = append(optParam.OptCaps, cap)
		}
		if n.server.Config.GracefulRestart {
			cap := &CapGracefulRestart{}
			if n.server.restarting {
				cap.Flags = RESTART_FLAG
			}
			cap.Time = uint16(n.server.GracefulRestartTime())
			for afiSafi, _ := range n.afiSafi {
				afi := afiSafi.Afi()
				safi := afiSafi.Safi()
				tuple := &CapGracefulRestartTuple{
					Afi:  afi,
					Safi: safi,
				}
				if n.server.preserved {
					tuple.Flags = FORWARDING_PRESERVED_FLAG
				}
				cap.Tuples = append(cap.Tuples, tuple)
			}
			optParam.OptCaps = append(optParam.OptCaps, cap)
		}
		n.localCaps = optParam.OptCaps
	}
	msg := &BgpMessage{
		Header: BgpHeader{Type: BGP_MSG_OPEN},
		Body:   body,
	}
	return msg
}

func NewBgpKeepAliveMsg() *BgpMessage {
	msg := &BgpMessage{
		Header: BgpHeader{Type: BGP_MSG_KEEPALIVE},
		Body:   &BgpKeepAlive{},
	}
	return msg
}
