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
	"encoding/json"
	"fmt"
	"testing"
)

func TestPacketLength(t *testing.T) {
	p := &Packet{}

	// Packet minimum length test.
	data := []byte{0, 0, 0}
	err := p.DecodeFromBytes(data)
	if err == nil {
		t.Errorf("Packet length less than %d must return error", RIP_HEADER_LEN)
	}

	// Packet size alignment test.  (Length - RIP_HEADER_LEN) % RIP_RTE_LEN == 0.
	data = []byte{RIP_REQUEST, RIPv2, 0, 0, 0}
	err = p.DecodeFromBytes(data)
	if err == nil {
		t.Errorf("Packet length does not aligned must return error")
	}

	// Packet size alignment test.  (Length - RIP_HEADER_LEN) % RIP_RTE_LEN == 0.
	data = []byte{RIP_REQUEST, RIPv2, 0, 0}
	err = p.DecodeFromBytes(data)
	if err != nil {
		t.Errorf("Header only packet decode must success")
	}
}

func TestPacketMarshal(t *testing.T) {
	p := &Packet{}

	data := []byte{RIP_REQUEST, RIPv2, 0, 0}
	err := p.DecodeFromBytes(data)
	if err != nil {
		t.Errorf("Decode must success")
	}

	byte, err := json.Marshal(p)
	fmt.Println(string(byte), err)
}
