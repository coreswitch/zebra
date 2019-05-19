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

	"github.com/coreswitch/zebra/policy"
)

func TestCommunityList(t *testing.T) {
	clist := NewCommunityListMap()
	clist.CommunityListAdd("test", CommunityListStandard, policy.Permit, "100:1")
	clist.CommunityListAdd("test", CommunityListStandard, policy.Permit, "100:2")
	for _, list := range clist {
		fmt.Println(list)
	}
	clist.CommunityListDelete("test", CommunityListStandard, policy.Permit, "100:1")
	for _, list := range clist {
		fmt.Println(list)
	}
}
