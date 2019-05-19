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
	"sort"
	"strconv"

	pb "github.com/coreswitch/zebra/api"
)

type RouteMapMaster struct {
	Map map[string]*RouteMap
}

type RouteMap struct {
	name    string
	entries []*RouteMapEntry
}

type RouteMapEntry struct {
	typ     Action
	seq     uint32
	matches []RouteMapMatch
	actions []RouteMapAction
}

type RouteMapAction interface {
	Action(interface{}, interface{})
}

func NewRouteMapMaster() *RouteMapMaster {
	return &RouteMapMaster{
		Map: map[string]*RouteMap{},
	}
}

func MatchVersionCompile(args []string) interface{} {
	return nil
}

func MatchVersion(rule interface{}, val interface{}) bool {
	var version = rule.(int)

	fmt.Println(version)

	return false
}

func NewRouteMapEntry(typ Action, seq uint32) *RouteMapEntry {
	return &RouteMapEntry{
		typ: typ,
		seq: seq,
	}
}

func (rmap *RouteMap) RouteMapEntryAdd(ent *RouteMapEntry) {
	rmap.entries = append(rmap.entries, ent)
	sort.Slice(rmap.entries, func(i, j int) bool {
		return rmap.entries[i].seq < rmap.entries[j].seq
	})
}

func (m *RouteMapMaster) RouteMapLookup(name string) *RouteMap {
	return m.Map[name]
}

func (m *RouteMapMaster) RouteMapEntryGet(name string, typ Action, seq uint32) *RouteMapEntry {
	rmap, ok := m.Map[name]
	if !ok {
		rmap = &RouteMap{name: name}
		m.Map[name] = rmap
	}
	for _, ent := range rmap.entries {
		if ent.seq == seq {
			// Force overwrite type when existing type is different with argument.
			if ent.typ != typ {
				ent.typ = typ
			}
			return ent
		}
	}
	ent := NewRouteMapEntry(typ, seq)
	rmap.RouteMapEntryAdd(ent)

	return ent
}

type RouteMapMatch interface {
	Match(interface{}, interface{}) bool
	MatchAdd(args ...string)
}

// MatchTag for route-map match statement such as "match tag 100".
type MatchTag struct {
	tags []uint32
}

func (match *MatchTag) Match(p, r interface{}) bool {
	rib := r.(*pb.Rib)
	for _, t := range match.tags {
		fmt.Println(t, "<->", rib.Tag)
		if t == rib.Tag {
			return true
		}
	}
	return false
}

func (match *MatchTag) MatchAdd(args ...string) {
	for _, tagStr := range args {
		tag, err := strconv.ParseUint(tagStr, 10, 32)
		if err == nil {
			match.tags = append(match.tags, uint32(tag))
		}
	}
}

func (rmap *RouteMapEntry) MatchAdd(name string, args ...string) {
	var match RouteMapMatch
	switch name {
	case "tag":
		match = &MatchTag{}
		match.MatchAdd(args...)
		rmap.matches = append(rmap.matches, match)
	}
}

func (ent *RouteMapEntry) Match(p, r interface{}) bool {
	for _, match := range ent.matches {
		if match.Match(p, r) {
			return true
		}
	}
	return false
}

func (ent *RouteMapEntry) Action(p, r interface{}) {
	for _, action := range ent.actions {
		action.Action(p, r)
	}
}

func (rmap *RouteMap) Match(p, r interface{}) bool {
	for _, ent := range rmap.entries {
		fmt.Println("ent", ent)
		if ent.Match(p, r) {
			return true
		}
	}
	return false
}
