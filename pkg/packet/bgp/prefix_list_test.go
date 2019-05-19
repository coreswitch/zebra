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
	"encoding/json"
	"fmt"
	"github.com/coreswitch/netutil"
	"testing"
)

func TestPrefixList(t *testing.T) {
	fmt.Println("Testing prefix-list")
	plist := NewPrefixListMaster()

	p, _ := netutil.ParsePrefix("10.0.0.0/8")
	plist.EntryAdd("hoge", p, 0, 10, 0, 0)

	p, _ = netutil.ParsePrefix("11.0.0.0/8")
	plist.EntryAdd("hoge", p, 0, 0, 0, 0)

	p, _ = netutil.ParsePrefix("11.0.0.0/8")
	plist.EntryAdd("hoge", p, 10, 0, 0, 0)

	plist.DescriptionSet("hoge", "")
	fmt.Println(plist)

	byte, err := json.Marshal(plist)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
	} else {
		fmt.Println(string(byte))
	}
}
