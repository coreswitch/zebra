// Copyright 2017 CoreSwitch
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

package netutil

import (
	"encoding/binary"
	"fmt"
	"net"
	"testing"
)

func TestPrefixRange(t *testing.T) {
	fmt.Println("Prefx range")
	p, _ := ParsePrefix("20.0.0.1/24")
	fmt.Println(p)

	pmin := p.Copy()
	pmax := p.Copy()
	pmin.ApplyMask()
	pmax.ApplyReverseMask()
	fmt.Println(pmin)
	fmt.Println(pmax)
	//
	val := binary.BigEndian.Uint32(pmin.IP)
	val++
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, val)
	ip := net.IP(bytes)
	fmt.Println(ip)
}

var prefixDefaltTests = []struct {
	str    string
	result bool
}{
	{"0.0.0.0/0", true},
	{"0.0.0.0/8", false},
	{"0.0.0.10/0", false},
	{"10.0.0.0/8", false},
	{"::/0", true},
	{"::/128", false},
	{"2001::/64", false},
	{"00:ff::/64", false},
	{"00:ff::/0", false},
}

func TestPrefixIsDefault(t *testing.T) {
	for _, tt := range prefixDefaltTests {
		p, err := ParsePrefix(tt.str)
		if err != nil {
			t.Errorf("Parse error %v", err)
		}
		if p.IsDefault() != tt.result {
			t.Errorf("IsDefault() for %s failed", tt.str)
		}
	}
}
