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

package policy

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/coreswitch/netutil"
)

func EqualJSON(s1, s2 string) (bool, error) {
	var i1 interface{}
	var i2 interface{}

	var err error
	err = json.Unmarshal([]byte(s1), &i1)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string1 %s", err.Error())
	}
	err = json.Unmarshal([]byte(s2), &i2)
	if err != nil {
		return false, fmt.Errorf("Error mashalling string2 %s", err.Error())
	}
	return reflect.DeepEqual(i1, i2), nil
}

func OrderEnsure(t *testing.T, plist *PrefixList) {
	seq := 0
	for _, v := range plist.Entries {
		if v.Seq <= seq {
			t.Error("OrderEnsure error")
		}
		seq = v.Seq
	}
}

func TestPrefixListAdd(t *testing.T) {
	var jsonStr string
	var targetStr string

	p1, _ := netutil.ParsePrefix("10.0.0.0/8")
	p2, _ := netutil.ParsePrefix("11.0.0.0/8")
	p3, _ := netutil.ParsePrefix("12.0.0.0/8")

	// Empty -> Add one.
	pm := NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1))
	plist := pm.Lookup("plist")
	if len(plist.Entries) != 1 {
		t.Error("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ := json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"}]}`
	result, err := EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Replace
	pm = NewPrefixListMaster()
	plist = pm.Get("plist")
	plist.EntryAdd(NewPrefixListEntry(5, Permit, p1))
	plist.EntryAdd(NewPrefixListEntry(5, Deny, p1))

	if len(plist.Entries) != 1 {
		t.Errorf("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"deny","prefix":"10.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Top. 10 -> 5.
	pm = NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(10, Permit, p1))
	pm.EntryAdd("plist", NewPrefixListEntry(5, Deny, p2))

	plist = pm.Lookup("plist")
	if len(plist.Entries) != 2 {
		t.Errorf("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"deny","prefix":"11.0.0.0/8"},{"seq":10,"action":"permit","prefix":"10.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Top. 5 -> 10.
	pm = NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1))
	pm.EntryAdd("plist", NewPrefixListEntry(10, Permit, p2))

	plist = pm.Lookup("plist")
	if len(plist.Entries) != 2 {
		t.Errorf("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"},{"seq":10,"action":"permit","prefix":"11.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Middle
	pm = NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1))
	pm.EntryAdd("plist", NewPrefixListEntry(10, Permit, p2))
	pm.EntryAdd("plist", NewPrefixListEntry(8, Permit, p3))

	plist = pm.Lookup("plist")
	if len(plist.Entries) != 3 {
		t.Errorf("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"},{"seq":8,"action":"permit","prefix":"12.0.0.0/8"},{"seq":10,"action":"permit","prefix":"11.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Bottom
	pm.EntryAdd("plist", NewPrefixListEntry(12, Deny, p3))
	plist = pm.Lookup("plist")
	if len(plist.Entries) != 4 {
		t.Errorf("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"},{"seq":8,"action":"permit","prefix":"12.0.0.0/8"},{"seq":10,"action":"permit","prefix":"11.0.0.0/8"},{"seq":12,"action":"deny","prefix":"12.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// seq number out of range (0) - nothing to do.
	pm.EntryAdd("plist", NewPrefixListEntry(0, Permit, p3))
	plist = pm.Lookup("plist")
	if len(plist.Entries) != 4 {
		t.Errorf("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"},{"seq":8,"action":"permit","prefix":"12.0.0.0/8"},{"seq":10,"action":"permit","prefix":"11.0.0.0/8"},{"seq":12,"action":"deny","prefix":"12.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}
}

func TestPrefixListAddRange(t *testing.T) {
	var jsonStr string
	var targetStr string

	p1, _ := netutil.ParsePrefix("10.0.0.0/8")
	pm := NewPrefixListMaster()
	err := pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1, WithEq(8)))
	if err == nil {
		t.Error("There must be eq error.")
	}

	pm = NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1, WithLe(32), WithGe(9)))
	plist := pm.Lookup("plist")
	byte, _ := json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8","le":32,"ge":9}]}`
	result, err := EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}
}

func TestPrefixListDelete(t *testing.T) {
	var jsonStr string
	var targetStr string

	p1, _ := netutil.ParsePrefix("10.0.0.0/8")
	p2, _ := netutil.ParsePrefix("11.0.0.0/8")
	p3, _ := netutil.ParsePrefix("12.0.0.0/8")

	// Delete
	pm := NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1))
	pm.EntryAdd("plist", NewPrefixListEntry(10, Permit, p2))
	pm.EntryDelete("plist", NewPrefixListEntry(10, Permit, p2))
	plist := pm.Lookup("plist")
	if len(plist.Entries) != 1 {
		t.Error("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ := json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"}]}`
	result, err := EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Delete existing top one.
	pm = NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1))
	pm.EntryAdd("plist", NewPrefixListEntry(10, Permit, p2))
	pm.EntryDelete("plist", NewPrefixListEntry(5, Permit, p1))
	plist = pm.Lookup("plist")
	if len(plist.Entries) != 1 {
		t.Error("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":10,"action":"permit","prefix":"11.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Delete existing middle one.
	pm = NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1))
	pm.EntryAdd("plist", NewPrefixListEntry(10, Permit, p2))
	pm.EntryAdd("plist", NewPrefixListEntry(15, Permit, p3))
	pm.EntryDelete("plist", NewPrefixListEntry(10, Permit, p2))
	plist = pm.Lookup("plist")
	if len(plist.Entries) != 2 {
		t.Error("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"},{"seq":15,"action":"permit","prefix":"12.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Delete existing last.
	pm = NewPrefixListMaster()
	pm.EntryAdd("plist", NewPrefixListEntry(5, Permit, p1))
	pm.EntryAdd("plist", NewPrefixListEntry(10, Permit, p2))
	pm.EntryAdd("plist", NewPrefixListEntry(15, Permit, p3))
	pm.EntryDelete("plist", NewPrefixListEntry(15, Permit, p3))
	plist = pm.Lookup("plist")
	if len(plist.Entries) != 2 {
		t.Error("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"},{"seq":10,"action":"permit","prefix":"11.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}

	// Delete non existing.
	pm.EntryDelete("plist", NewPrefixListEntry(20, Permit, p3))
	plist = pm.Lookup("plist")
	if len(plist.Entries) != 2 {
		t.Error("PrefixList entry length error")
	}
	OrderEnsure(t, plist)
	byte, _ = json.Marshal(plist)
	jsonStr = string(byte)
	targetStr = `{"name":"plist","prefix-list-entries":[{"seq":5,"action":"permit","prefix":"10.0.0.0/8"},{"seq":10,"action":"permit","prefix":"11.0.0.0/8"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("PrefixList error", jsonStr, "should be", targetStr)
	}
}

func TestPrefixListMatch(t *testing.T) {
	p1, _ := netutil.ParsePrefix("172.0.0.0/8")
	p2, _ := netutil.ParsePrefix("172.16.0.0/12")
	p3, _ := netutil.ParsePrefix("173.16.0.0/12")
	p4, _ := netutil.ParsePrefix("10.0.0.0/8")
	any, _ := netutil.ParsePrefix("0.0.0.0/0")

	pm := NewPrefixListMaster()
	plist := pm.Get("plist")
	plist.EntryAdd(NewPrefixListEntry(5, Deny, p1, WithEq(12)))
	plist.EntryAdd(NewPrefixListEntry(10, Permit, any, WithLe(32)))

	// Match & deny.
	ret := plist.Match(p2)
	if ret != Deny {
		t.Error("Must be deny match")
	}

	// Match & permit.
	ret = plist.Match(p3)
	if ret != Permit {
		t.Error("Must be any match permit")
	}

	// No match.
	plist = pm.Get("plist_empty")
	ret = plist.Match(p1)
	if ret != Deny {
		t.Error("Must be default deny")
	}

	// Exact match.
	plist = pm.Get("default")
	plist.EntryAdd(NewPrefixListEntry(5, Permit, any))
	ret = plist.Match(any)
	if ret != Permit {
		t.Error("Must be permit")
	}
	ret = plist.Match(p1)
	if ret != Deny {
		t.Error("Must be no match deny")
	}

	// LE
	plist = pm.Get("le")
	plist.EntryAdd(NewPrefixListEntry(5, Permit, p4, WithLe(30)))
	le1, _ := netutil.ParsePrefix("10.0.0.0/30")
	ret = plist.Match(le1)
	if ret != Permit {
		t.Error("Must be permit")
	}
	le2, _ := netutil.ParsePrefix("10.0.0.0/31")
	ret = plist.Match(le2)
	if ret != Deny {
		t.Error("Must be deny")
	}
	le3, _ := netutil.ParsePrefix("10.0.0.0/8")
	ret = plist.Match(le3)
	if ret != Permit {
		t.Error("Must be permit")
	}
	le4, _ := netutil.ParsePrefix("10.0.0.0/7")
	ret = plist.Match(le4)
	if ret != Deny {
		t.Error("Must be deny")
	}

	// GE
	plist = pm.Get("ge")
	plist.EntryAdd(NewPrefixListEntry(5, Permit, p4, WithGe(30)))
	ge1, _ := netutil.ParsePrefix("10.0.0.0/30")
	ret = plist.Match(ge1)
	if ret != Permit {
		t.Error("Must be permit")
	}
	ge2, _ := netutil.ParsePrefix("11.0.0.0/30")
	ret = plist.Match(ge2)
	if ret != Deny {
		t.Error("Must be deny")
	}
	ge3, _ := netutil.ParsePrefix("10.0.0.0/29")
	ret = plist.Match(ge3)
	if ret != Deny {
		t.Error("Must be deny")
	}
	ge4, _ := netutil.ParsePrefix("10.0.0.0/8")
	ret = plist.Match(ge4)
	if ret != Deny {
		t.Error("Must be deny")
	}

	// 0.0.0.0/0 ge 1 => match to all of non default route.
	plist = pm.Get("non-default")
	plist.EntryAdd(NewPrefixListEntry(5, Permit, any, WithGe(1)))
	ret = plist.Match(any)
	if ret != Deny {
		t.Error("Must be permit")
	}
	ret = plist.Match(p4)
	if ret != Permit {
		t.Error("Must be no match deny")
	}

	// LE & GE
	plist = pm.Get("lege")
	plist.EntryAdd(NewPrefixListEntry(5, Permit, p4, WithGe(29), WithLe(30)))
	lege1, _ := netutil.ParsePrefix("10.0.0.0/30")
	ret = plist.Match(lege1)
	if ret != Permit {
		t.Error("Must be permit")
	}
	lege2, _ := netutil.ParsePrefix("11.0.0.0/30")
	ret = plist.Match(lege2)
	if ret != Deny {
		t.Error("Must be deny")
	}
	lege3, _ := netutil.ParsePrefix("10.0.0.0/29")
	ret = plist.Match(lege3)
	if ret != Permit {
		t.Error("Must be deny")
	}
	lege4, _ := netutil.ParsePrefix("10.0.0.0/8")
	ret = plist.Match(lege4)
	if ret != Deny {
		t.Error("Must be deny")
	}

	// EQ
	plist = pm.Get("eq")
	plist.EntryAdd(NewPrefixListEntry(5, Deny, p4, WithEq(16)))
	plist.EntryAdd(NewPrefixListEntry(10, Permit, any, WithLe(32)))
	eq1, _ := netutil.ParsePrefix("10.0.0.0/30")
	ret = plist.Match(eq1)
	if ret != Permit {
		t.Error("Must be permit")
	}
	eq2, _ := netutil.ParsePrefix("10.11.0.0/16")
	ret = plist.Match(eq2)
	if ret != Deny {
		t.Error("Must be deny")
	}
	eq3, _ := netutil.ParsePrefix("10.0.0.0/29")
	ret = plist.Match(eq3)
	if ret != Permit {
		t.Error("Must be permit")
	}
	eq4, _ := netutil.ParsePrefix("11.11.11.0/16")
	ret = plist.Match(eq4)
	if ret != Permit {
		t.Error("Must be permit")
	}
}
