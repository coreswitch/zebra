// Copyright 2016 Zebra 2.0 Project
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

package ldp

import (
	"encoding/binary"
	"net"
)

const (
	LDP_HEADER_LEN = 10
)

const (
	LDP_TYPE_NULL             = 0x0000
	LDP_TYPE_NOTIFICATION     = 0x0001
	LDP_TYPE_HELLO            = 0x0100
	LDP_TYPE_INITIALIZATION   = 0x0200
	LDP_TYPE_KEEPALIVE        = 0x0201
	LDP_TYPE_ADDRESS          = 0x0300
	LDP_TYPE_ADDRESS_WITHDRAW = 0x0301
	LDP_TYPE_LABEL_MAPPING    = 0x0400
	LDP_TYPE_LABEL_REQUEST    = 0x0401
	LDP_TYPE_LABEL_WITHDRAW   = 0x0402
	LDP_TYPE_LABEL_RELEASE    = 0x0403
	LDP_TYPE_REQUEST_ABORT    = 0x0404
)

type LdpHeader struct {
	Version    uint16
	Length     uint16
	RouterId   net.IP
	LabelSpace uint16
}

func (msg *LdpHeader) DecodeFromBytes(data []byte) error {
	if len(data) < LDP_HEADER_LEN {
		return nil
	}
	msg.Version = binary.BigEndian.Uint16(data[0:2])
	msg.Length = binary.BigEndian.Uint16(data[2:4])
	if int(msg.Length) < LDP_HEADER_LEN {
		return nil
	}
	return nil
}

func (msg *LdpHeader) Serialize() ([]byte, error) {
	buf := make([]byte, LDP_HEADER_LEN)
	binary.BigEndian.PutUint16(buf[0:2], msg.Version)
	binary.BigEndian.PutUint16(buf[2:4], msg.Length)
	copy(buf[4:8], msg.RouterId)
	binary.BigEndian.PutUint16(buf[8:10], msg.LabelSpace)
	return buf, nil
}

type LdpBody interface {
	DecodeFromBytes([]byte) error
	Serialize() ([]byte, error)
}

type LdpMessage struct {
	Header LdpHeader
	Body   LdpBody
}

type LdpTlv struct {
	Ubit   bool
	Fbit   bool
	Type   uint16
	Length uint16
}

type LdpHello struct {
	Tlv LdpTlv
}

func (msg *LdpHello) DecodeFromBytes([]byte) error {
	return nil
}

func (msg *LdpHello) Serialize() ([]byte, error) {
	return nil, nil
}

// type LdpNotification struct {
// }

// type LdpKeepalive struct {
// }
