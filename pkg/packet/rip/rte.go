// Copyright 2019 zebra project.
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

package rip

import (
	"encoding/binary"
	"fmt"
	"net"
)

// RTEInterface is common interface for RTE, AuthInfo and AuthData.
type RTEInterface interface {
	DecodeFromBytes(data []byte) error
	Serialize() ([]byte, error)
}

// RTE Routing table entry
type RTE struct {
	Family  uint16
	Tag     uint16
	Prefix  net.IP
	Mask    net.IP
	Nexthop net.IP
	Metric  uint32
}

// AuthText is simple text authentication.
type AuthText struct {
	Family     uint16
	AuthType   uint16
	AuthString [16]byte
}

// AuthCrypto is authentication information.
type AuthCrypto struct {
	Family    uint16
	AuthType  uint16
	PacketLen uint16
	KeyID     byte
	AuthLen   byte
	Sequence  uint32
	Reserve1  uint32
	Reserve2  uint32
}

// AuthData is authentication data.
type AuthData struct {
	Family   uint16
	AuthType uint16
	Digest   [16]byte
}

// DecodeFromBytes return parsed RIP packet.
func (rte *RTE) DecodeFromBytes(data []byte) error {
	if len(data) < RTELen {
		return fmt.Errorf("RTE length is too small")
	}
	rte.Family = binary.BigEndian.Uint16(data[0:2])
	rte.Tag = binary.BigEndian.Uint16(data[2:4])
	rte.Prefix = data[4:8]
	rte.Mask = data[8:12]
	rte.Nexthop = data[12:16]
	rte.Metric = binary.BigEndian.Uint32(data[16:20])
	return nil
}

// Serialize RTE.
func (rte *RTE) Serialize() ([]byte, error) {
	buf := make([]byte, RTELen)
	binary.BigEndian.PutUint16(buf[0:2], rte.Family)
	binary.BigEndian.PutUint16(buf[2:4], rte.Tag)
	copy(buf[4:8], rte.Prefix)
	copy(buf[8:12], rte.Mask)
	copy(buf[12:16], rte.Nexthop)
	binary.BigEndian.PutUint32(buf[16:20], rte.Metric)
	return buf, nil
}

// DecodeFromBytes for AuthText.
func (rte *AuthText) DecodeFromBytes(data []byte) error {
	if len(data) < RTELen {
		return fmt.Errorf("RTE length is too small")
	}
	rte.Family = binary.BigEndian.Uint16(data[0:2])
	rte.AuthType = binary.BigEndian.Uint16(data[2:4])
	copy(rte.AuthString[:], data[4:20])
	return nil
}

// Serialize AutText.
func (rte *AuthText) Serialize() ([]byte, error) {
	buf := make([]byte, RTELen)
	binary.BigEndian.PutUint16(buf[0:2], rte.Family)
	binary.BigEndian.PutUint16(buf[2:4], rte.AuthType)
	copy(buf[4:20], rte.AuthString[:])
	return buf, nil
}

// DecodeFromBytes for AuthCrypto.
func (rte *AuthCrypto) DecodeFromBytes(data []byte) error {
	if len(data) < RTELen {
		return fmt.Errorf("RTE length is too small")
	}
	rte.Family = binary.BigEndian.Uint16(data[0:2])
	rte.AuthType = binary.BigEndian.Uint16(data[2:4])
	rte.PacketLen = binary.BigEndian.Uint16(data[4:6])
	rte.KeyID = data[7]
	rte.AuthLen = data[8]
	rte.Sequence = binary.BigEndian.Uint32(data[8:12])
	rte.Reserve1 = binary.BigEndian.Uint32(data[12:16])
	rte.Reserve2 = binary.BigEndian.Uint32(data[16:20])
	return nil
}

// Serialize for AuthCrypto.
func (rte *AuthCrypto) Serialize() ([]byte, error) {
	buf := make([]byte, RTELen)
	binary.BigEndian.PutUint16(buf[0:2], rte.Family)
	binary.BigEndian.PutUint16(buf[2:4], rte.AuthType)
	binary.BigEndian.PutUint16(buf[4:6], rte.PacketLen)
	buf[7] = rte.KeyID
	buf[8] = rte.AuthLen
	binary.BigEndian.PutUint32(buf[8:12], rte.Sequence)
	binary.BigEndian.PutUint32(buf[12:16], rte.Reserve1)
	binary.BigEndian.PutUint32(buf[16:20], rte.Reserve2)

	return buf, nil
}

// DecodeFromBytes for AuthData.
func (rte *AuthData) DecodeFromBytes(data []byte) error {
	if len(data) < RTELen {
		return fmt.Errorf("RTE length is too small")
	}
	rte.Family = binary.BigEndian.Uint16(data[0:2])
	rte.AuthType = binary.BigEndian.Uint16(data[2:4])
	copy(rte.Digest[:], data[4:20])
	return nil
}

// Serialize for AuthData.
func (rte *AuthData) Serialize() ([]byte, error) {
	buf := make([]byte, RTELen)
	binary.BigEndian.PutUint16(buf[0:2], rte.Family)
	binary.BigEndian.PutUint16(buf[2:4], rte.AuthType)
	copy(buf[4:20], rte.Digest[:])
	return buf, nil
}
