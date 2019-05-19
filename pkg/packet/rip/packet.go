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
	"fmt"
	"net"
)

type Rte struct {
	Family  uint16
	Tag     uint16
	Prefix  net.IP
	Mask    net.IP
	Nexthop net.IP
	Metric  uint32
}

type Packet struct {
	Command  byte
	Version  byte
	Padding1 byte
	Padding2 byte
	Rtes     []*Rte
}

func (p *Packet) DecodeFromBytes(data []byte) error {
	fmt.Println("data len", len(data))
	if len(data) < RIP_PACKET_MINSIZE {
		return fmt.Errorf("Pakcet len %d is smaller than minimum size %d", len(data), RIP_PACKET_MINSIZE)
	}
	if (len(data)-RIP_PACKET_MINSIZE)%20 > 0 {
		return fmt.Errorf("packet size %d is wrong RIP packet alignment", len(data))
	}

	p.Command = data[0]
	p.Version = data[1]
	p.Padding1 = data[2]
	p.Padding2 = data[3]

	data = data[RIP_PACKET_MINSIZE:]
	rtenum := len(data) / 20
	fmt.Println("Num RTEs", rtenum)

	return nil
}

func (s *Server) PacketParse() error {
	p := &Packet{}
	err := p.DecodeFromBytes(s.Buffer)
	if err != nil {
		fmt.Println("Parse error")
		return err
	}
	fmt.Println("Packet:", p)
	return nil
}
