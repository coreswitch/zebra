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

package policy

import (
	"fmt"
	"testing"

	pb "github.com/hash-set/zebra/proto"
)

func TestRouteMapTag(t *testing.T) {
	rm := NewRouteMapMaster()

	// Route-map entry get.
	ent := rm.RouteMapEntryGet("test", Permit, 10)
	fmt.Println(ent)

	// Route-map match add.
	ent.MatchAdd("tag", "100", "200")

	// Make Test Rib.
	rib := &pb.Rib{}

	// Execute match.
	rmap := rm.RouteMapLookup("test")
	if rmap == nil {
		fmt.Errorf("RouteMapLookup failed")
		return
	}

	rib.Tag = 1
	if rmap.Match(nil, rib) != false {
		fmt.Errorf("match to tag 1 should failed")
	}
	rib.Tag = 100
	if rmap.Match(nil, rib) != true {
		fmt.Errorf("match to tag 1 should success")
	}
}
