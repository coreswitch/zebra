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
	"fmt"
	"testing"
)

func TestAsPathDecode(t *testing.T) {
	buf := make([]byte, 2+8)
	buf[0] = AS_SEQUENCE
	buf[1] = 2
	binary.BigEndian.PutUint32(buf[2:], uint32(1))
	binary.BigEndian.PutUint32(buf[6:], uint32(2))

	seg := &As4Segment{}
	err := seg.DecodeFromBytes(buf)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println(seg)
	}
	aspath := AsPath{}
	aspath = append(aspath, seg)

	fmt.Println("string:", seg.String())
	fmt.Println("aspath string:", aspath.String())
}

func TestAsPathParse(t *testing.T) {
	str := "(3)1{600} 2 [101 123] [1]"

	aspath, err := AsPathParse(str)
	if err != nil {
		fmt.Errorf("AsPathParse err: %s", err)
	}
	if aspath.String() != "(3) 1 {600} 2 [101 123] [1]" {
		fmt.Errorf("AsPathParse err: %s", err)
	}
}

func TestAsPathParse256(t *testing.T) {
	aspath := AsPath{}
	for i := 0; i < 255; i++ {
		aspath = aspath.Append(uint32(i))
	}
	if len(aspath) != 1 {
		fmt.Errorf("AS_PATH segment number is not 1")
	}
	aspath = aspath.Append(256)
	if len(aspath) != 2 {
		fmt.Errorf("AS_PATH segment number is not 2")
	}
}
