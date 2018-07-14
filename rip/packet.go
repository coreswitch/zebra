// Copyright 2018 zebra project.
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

type Packet struct {
	Command  byte   `json:"command"`
	Version  byte   `json:"version"`
	Padding1 byte   `json:"padding1,omitempty"`
	Padding2 byte   `json:"padding2,omitempty"`
	RTEs     []*RTE `json:"rtes,omitempty"`
}

type RTE struct {
	Family  uint16
	Tag     uint16
	Prefix  net.IP
	Mask    net.IP
	Nexthop net.IP
	Metric  uint32
}

type DecodeOption struct {
}

func (p *Packet) DecodeFromBytes(data []byte, opts ...DecodeOption) error {
	if len(data) < RIP_HEADER_LEN {
		return fmt.Errorf("Pakcet len %d is smaller than minimum size %d", len(data), RIP_HEADER_LEN)
	}
	if len(opts) > 0 && len(data) > RIP_PACKET_MAXLEN {
		return fmt.Errorf("Pakcet len %d is larger than maximum size %d", len(data), RIP_PACKET_MAXLEN)
	}
	if (len(data)-RIP_HEADER_LEN)%RIP_RTE_LEN > 0 {
		return fmt.Errorf("Packet len %d is wrong RIP packet alignment", len(data))
	}

	p.Command = data[0]
	p.Version = data[1]
	p.Padding1 = data[2]
	p.Padding2 = data[3]

	data = data[RIP_HEADER_LEN:]
	rteNum := len(data) / 20

	for i := 0; i < rteNum; i++ {
		rte := &RTE{}
		rte.Family = binary.BigEndian.Uint16(data[0:2])
		rte.Tag = binary.BigEndian.Uint16(data[2:4])
		rte.Prefix = data[4:8]
		rte.Mask = data[8:12]
		rte.Nexthop = data[12:16]
		rte.Metric = binary.BigEndian.Uint32(data[16:20])

		p.RTEs = append(p.RTEs, rte)
		data = data[RIP_RTE_LEN:]
	}
	return nil
}

func (p *Packet) String() string {
	str := ""
	str += fmt.Sprintf("command:%s ", Command2Str(p.Command))
	str += fmt.Sprintf("version:%d ", p.Version)
	for _, rte := range p.RTEs {
		str += fmt.Sprintf("family:%d ", rte.Family)
		str += fmt.Sprintf("tag:%d ", rte.Tag)
		str += fmt.Sprintf("prefix:%s ", rte.Prefix)
		str += fmt.Sprintf("mask:%s ", rte.Mask)
		str += fmt.Sprintf("nexthop:%s ", rte.Nexthop)
		str += fmt.Sprintf("metrid:%d ", rte.Metric)
	}
	return str
}

func (p *Packet) Serialize() ([]byte, error) {
	buf := make([]byte, RIP_HEADER_LEN+(len(p.RTEs)*RIP_RTE_LEN))
	buf[0] = p.Command
	buf[1] = p.Version

	i := RIP_HEADER_LEN
	for _, rte := range p.RTEs {
		binary.BigEndian.PutUint16(buf[i:], rte.Family)
		binary.BigEndian.PutUint16(buf[i+2:], rte.Tag)
		copy(buf[i+4:], rte.Prefix)
		copy(buf[i+8:], rte.Mask)
		copy(buf[i+12:], rte.Nexthop)
		binary.BigEndian.PutUint32(buf[i+16:], rte.Metric)
		i += RIP_RTE_LEN
	}
	return buf, nil
}
