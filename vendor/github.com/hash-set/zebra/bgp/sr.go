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
)

type TlvInterface interface {
	DecodeFromBytes([]byte) error
	// Serialize() ([]byte, error)
	Length() int
}

type AttrPrefixSid struct {
	AttrBase
	Tlvs []TlvInterface
}

const (
	PREFIX_SID_TYPE_LABEL_INDEX         uint8 = 1
	PREFIX_SID_TYPE_IPV6_TLV                  = 2
	PREFIX_SID_TYPE_ORIGINATOR_SRGB_TLV       = 3
	PREFIX_SID_TYPE_SRV6_VPN_SID_TLV          = 4 // TBD.
)

type TlvBase struct {
	TlvType   uint8
	TlvLength uint16
	Value     []byte
}

func (tlv TlvBase) DecodeFromBytes(data []byte) error {
	tlv.TlvType = data[0]
	tlv.TlvLength = binary.BigEndian.Uint16(data[1:3])
	if len(data) < int(tlv.TlvLength) {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	tlv.Value = data[3:(tlv.TlvLength - 3)]
	return nil
}

func (tlv TlvBase) Length() int {
	return int(tlv.TlvLength)
}

type LabelIndexTlv struct {
	TlvBase
	Flags      uint16
	LabelIndex uint32
}

func (tlv LabelIndexTlv) DecodeFromBytes(data []byte) error {
	err := tlv.TlvBase.DecodeFromBytes(data)
	if err != nil {
		return nil
	}
	if tlv.TlvLength != 7 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	tlv.Flags = binary.BigEndian.Uint16(data[1:3])
	tlv.LabelIndex = binary.BigEndian.Uint32(data[3:7])
	return nil
}

type Ipv6Tlv struct {
	TlvBase
	Flags uint16
}

func (tlv Ipv6Tlv) DecodeFromBytes(data []byte) error {
	err := tlv.TlvBase.DecodeFromBytes(data)
	if err != nil {
		return nil
	}
	if tlv.TlvLength != 3 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	tlv.Flags = binary.BigEndian.Uint16(data[1:3])
	return nil
}

type Srgb []byte

type OriginatorSrbgTlv struct {
	TlvBase
	Flags uint16
	Srgb  []Srgb
}

func (tlv OriginatorSrbgTlv) DecodeFromBytes(data []byte) error {
	err := tlv.TlvBase.DecodeFromBytes(data)
	if err != nil {
		return nil
	}
	if tlv.TlvLength < 2 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	if ((tlv.TlvLength - 2) % 6) != 0 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	tlv.Flags = binary.BigEndian.Uint16(data[0:2])
	data = data[2:]

	num := int((tlv.TlvLength - 2) / 6)
	for i := 0; i < num; i++ {
		srgb := make([]byte, 3)
		copy(srgb, data[0:3])
		data = data[3:]
		tlv.Srgb = append(tlv.Srgb, srgb)
	}
	return nil
}

type Srv6Sid struct {
	Type uint8  // 1 octet
	Sid  []byte // 16 octet
}

type Srv6VpnSidTlv struct {
	TlvBase
	Sids []*Srv6Sid
}

func (tlv Srv6VpnSidTlv) DecodeFromBytes(data []byte) error {
	err := tlv.TlvBase.DecodeFromBytes(data)
	if err != nil {
		return nil
	}
	if tlv.TlvLength < 1 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	if ((tlv.TlvLength - 1) % 17) != 0 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	data = data[1:]

	num := int((tlv.TlvLength - 1) / 17)
	for i := 0; i < num; i++ {
		sid := &Srv6Sid{}
		sid.Type = data[0]
		sid.Sid = make([]byte, 16)
		copy(sid.Sid, data[1:17])
		data = data[17:]
		tlv.Sids = append(tlv.Sids, sid)
	}
	return nil
}

func DecodeTlv(data []byte) (TlvInterface, error) {
	if len(data) < 3 {
		return nil, NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	var tlv TlvInterface
	switch uint8(data[0]) {
	case PREFIX_SID_TYPE_LABEL_INDEX:
		tlv = &LabelIndexTlv{}
	case PREFIX_SID_TYPE_IPV6_TLV:
		tlv = &Ipv6Tlv{}
	case PREFIX_SID_TYPE_ORIGINATOR_SRGB_TLV:
		tlv = &OriginatorSrbgTlv{}
	case PREFIX_SID_TYPE_SRV6_VPN_SID_TLV:
		tlv = &Srv6VpnSidTlv{}
	default:
		tlv = &TlvBase{}
	}
	err := tlv.DecodeFromBytes(data)
	return tlv, err
}

func (attr *AttrPrefixSid) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	for len(attr.Value) > 0 {
		tlv, err := DecodeTlv(data)
		if err != nil {
			return err
		}
		attr.Tlvs = append(attr.Tlvs, tlv)
		data = data[tlv.Length():]
	}
	return nil
}
