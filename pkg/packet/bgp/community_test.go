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
	"fmt"
	"testing"
)

func TestCommunity(t *testing.T) {
	comStr := "100:10 100:20 100:30"
	com, err := CommunityParse(comStr)
	if err != nil {
		t.Errorf("parse error")
		return
	}
	if com == nil {
		t.Errorf("community string can't be parsed")
		return
	}
	comVal := uint32(100<<16 | 10)
	if com[0] != comVal {
		t.Errorf("community value must be 100:10")
		return
	}
}

func TestCommunityString(t *testing.T) {
	comStr := "no-export 100:1"
	com, err := CommunityParse(comStr)
	if err != nil {
		t.Errorf("parse error")
		return
	}
	str := fmt.Sprint(com)
	if str != "no-export 100:1" {
		t.Errorf("community string is not no-export 100:1")
		return
	}
}

func TestCommunityString2(t *testing.T) {
	comStr := "no-export 100:1 65537"
	com, err := CommunityParse(comStr)
	if err != nil {
		t.Errorf("parse error")
		return
	}
	str := fmt.Sprint(com)
	if str != "no-export 100:1 1:1" {
		t.Errorf("community string mismatch %s", str)
		return
	}
}

func TestCommunityVal(t *testing.T) {
	comStr := "100"
	comVal, err := CommunityValParse(comStr)
	if err != nil {
		t.Errorf("community val parse error")
		return
	}
	if comVal != 100 {
		t.Errorf("community val must be 100")
		return
	}
}

func TestCommunityVal2(t *testing.T) {
	comStr := "100:1"
	comVal, err := CommunityValParse(comStr)
	if err != nil {
		t.Errorf("community val parse error")
		return
	}
	if comVal != uint32(100<<16|1) {
		t.Errorf("community val must be 100:1")
		return
	}
}

func TestCommunityVal3(t *testing.T) {
	comStr := "test"
	_, err := CommunityValParse(comStr)
	if err == nil {
		t.Errorf("community val parse must generate error")
		return
	}
}

func TestCommunityVal4(t *testing.T) {
	comStr := "65537"
	comVal, err := CommunityValParse(comStr)
	if err != nil {
		t.Errorf("community val parse error")
		return
	}
	if comVal != uint32(65537) {
		t.Errorf("community val mismatch %d", comVal)
		return
	}
}

func TestCommunityParse(t *testing.T) {
	com1 := Community{4, 3, 2, 1, 2, 3, 1, 5, 4}
	com1 = com1.SortUnique()
	fmt.Println(com1.String())
	com2, err := CommunityParse("0:1 0:2 0:3 0:4 0:5")
	if err != nil {
		t.Errorf("Can't parse string")
	}
	if !com1.Equal(com2) {
		t.Errorf("Community SortUnique error")
	}
}
