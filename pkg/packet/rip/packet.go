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
)

// Packet RIP packet.
type Packet struct {
	Command    byte
	Version    byte
	Padding1   byte
	Padding2   byte
	RTEs       []*RTE
	AuthText   *AuthText
	AuthCrypto *AuthCrypto
	AuthData   *AuthData
}

// DecodeFromBytes return parsed RIP packet.
func (p *Packet) DecodeFromBytes(data []byte) error {
	// Packet length check.
	if len(data) < HeaderLen {
		return fmt.Errorf("Packet length is too small")
	}
	if (len(data)-HeaderLen)%RTELen != 0 {
		return fmt.Errorf("Packet length is not aligned with RTE length 20")
	}
	// Parse header.
	p.Command = data[0]
	p.Version = data[1]
	p.Padding1 = data[2]
	p.Padding2 = data[3]
	data = data[HeaderLen:]

	// Parse RTEs.
	for len(data) > 0 {
		var rteif RTEInterface
		family := binary.BigEndian.Uint16(data[0:2])
		tag := binary.BigEndian.Uint16(data[2:4])
		if family == 0xffff {
			switch tag {
			case AuthTypeData:
				p.AuthData = &AuthData{}
				rteif = p.AuthData
			case AuthTypeText:
				p.AuthText = &AuthText{}
				rteif = p.AuthText
			case AuthTypeCrypto:
				p.AuthCrypto = &AuthCrypto{}
				rteif = p.AuthCrypto
			default:
				return fmt.Errorf("Unknown authentication type")
			}
		} else {
			rte := &RTE{}
			p.RTEs = append(p.RTEs, rte)
			rteif = rte
		}
		err := rteif.DecodeFromBytes(data)
		if err != nil {
			return err
		}
		data = data[RTELen:]
	}
	return nil
}

// Serialize packet.
func (p *Packet) Serialize() ([]byte, error) {
	buf := make([]byte, HeaderLen)
	buf[0] = p.Command
	buf[1] = p.Version
	buf[2] = p.Padding1
	buf[3] = p.Padding2
	for _, rte := range p.RTEs {
		rbuf, err := rte.Serialize()
		if err != nil {
			return nil, err
		}
		buf = append(buf, rbuf...)
	}
	return buf, nil
}
