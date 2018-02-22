// Copyright 2017 OpenConfigd Project.
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

package config

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/coreswitch/cmd"
	"github.com/coreswitch/goyang/pkg/yang"
)

func YangSetUp(yangFileName string) (*yang.Entry, error) {
	yang.AddPath(Env("GOPATH") + "/src/github.com/coreswitch/openconfigd/yang")

	// Load YANG.
	ms := yang.NewModules()
	err := ms.Read(yangFileName)
	if err != nil {
		return nil, err
	}

	// Process YANG.
	ms.Process()

	// Avoid duplication.
	mods := map[string]*yang.Module{}
	var names []string
	for _, m := range ms.Modules {
		if mods[m.Name] == nil {
			mods[m.Name] = m
			names = append(names, m.Name)
		}
	}

	// Alloc top entry.
	entry := &yang.Entry{
		Kind: yang.DirectoryEntry,
		Dir:  map[string]*yang.Entry{},
	}

	// Convert to Module to Entry.
	for _, name := range names {
		e := yang.ToEntry(mods[name])
		for key, value := range e.Dir {
			entry.Dir[key] = value
		}
	}

	return entry, nil
}

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

// List.
func TestList(t *testing.T) {
	// Load multiple key YANG.
	entry, err := YangSetUp("test-list.yang")
	if err != nil {
		fmt.Println(err)
	}

	// Parse success test.
	config := &Config{}
	path := []string{"top", "key", "key-value", "value", "leaf-value"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// JSON test.
	jsonStr := config.JsonMarshal()
	targetStr := `{"top":{"key": [{"first":"key-value","value":"leaf-value"}]}}`
	var result bool
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Parse incomplete test.
	config = &Config{}
	path = []string{"top", "key"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseIncomplete {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// Parse incomplete test.
	config = &Config{}
	path = []string{"top", "key", "key-value", "value"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseIncomplete {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// Delete leaf preparation.
	config = &Config{}
	path = []string{"top", "key", "key1", "value", "value1"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	path = []string{"top", "key", "key2", "value", "value2"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// JSON test.
	jsonStr = config.JsonMarshal()
	targetStr = `{"top":{"key": [{"first":"key1","value":"value1"},{"first":"key2","value":"value2"}]}}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Delete key.
	path = []string{"top", "key", "key1"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Presence parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{"top":{"key": [{"first":"key2","value":"value2"}]}}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Delete leaf node.
	path = []string{"top", "key", "key2", "value"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Presence parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{"top":{"key": [{"first":"key2"}]}}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr, err, result)
	}

	// Delete key node.
	path = []string{"top", "key", "key2"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Presence parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}
}

// Container.
func TestContainer(t *testing.T) {
}

// Leaf.
func TestLeaf(t *testing.T) {
	// Load multiple key YANG.
	entry, err := YangSetUp("test-leaf.yang")
	if err != nil {
		fmt.Println(err)
	}

	// Empty value. Set test.
	config := &Config{}
	path := []string{"leaf-empty"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// Empty value. JSON test.
	jsonStr := config.JsonMarshal()
	targetStr := `{"leaf-empty":true}`
	var result bool
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Empty value. Delete test.
	path = []string{"leaf-empty"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List leaf failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// uint16 value.
	config = &Config{}
	path = []string{"leaf-uint16-range", "68"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{"leaf-uint16-range":68}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	config = &Config{}
	path = []string{"leaf-uint16-range", "1"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseNoMatch {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// Boolean value.
	config = &Config{}
	path = []string{"leaf-boolean", "true"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	path = []string{"leaf-boolean", "false"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{"leaf-boolean":false}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}
	path = []string{"leaf-boolean", "true"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseNoMatch {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	path = []string{"leaf-boolean", "false"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Enum value.
	config = &Config{}
	path = []string{"leaf-enum", "grape"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseNoMatch {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	path = []string{"leaf-enum", "orange"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{"leaf-enum":"orange"}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}
	config = &Config{}
	path = []string{"leaf-enum", "or"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseIncomplete {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
}

func TestLeafList(t *testing.T) {
	// Load leaf-list YANG.
	entry, err := YangSetUp("test-leaf-list.yang")
	if err != nil {
		fmt.Println(err)
	}

	// Empty value. Set test.
	config := &Config{}
	path := []string{"top", "segments", "a::1", "b::1"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Leaf parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// Empty value. JSON test.
	jsonStr := config.JsonMarshal()
	targetStr := `{"top":{"segments":["a::1","b::1"]}}`
	var result bool
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}
}

// Multiple key.
func TestMultiKey(t *testing.T) {
	// Load multiple key YANG.
	entry, err := YangSetUp("test-multikey.yang")
	if err != nil {
		fmt.Println(err)
	}

	// Parse success test.
	config := &Config{}
	path := []string{"multikey", "first", "second"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Multikey parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// JSON test.
	jsonStr := config.JsonMarshal()
	targetStr := `{"multikey": [{"first":"first","second":"second"}]}`
	var result bool
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	config = &Config{}
	path = []string{"multikey", "first", "second", "third", "value"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Multikey parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// JSON test.
	jsonStr = config.JsonMarshal()
	targetStr = `{"multikey": [{"first":"first","second":"second","third":"value"}]}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Parse incomplete test.
	config = &Config{}
	path = []string{"multikey", "first"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseIncomplete {
		t.Error("Multikey parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	config = &Config{}
	path = []string{"multikey", "first", "second", "third"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseIncomplete {
		t.Error("Multikey parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// Parse failure test.
	config = &Config{}
	path = []string{"multikey", "first", "second", "other"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseNoMatch {
		t.Error("Multikey parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// Commit set test.
	// Command test.

	// Delete failure test.
	// Delete incomplete test.
	// Delete success test.

	// Commit delete test.
}

func TestPresence(t *testing.T) {
	// Load multiple key YANG.
	entry, err := YangSetUp("test-presence.yang")
	if err != nil {
		fmt.Println(err)
	}

	// Parse success.
	config := &Config{}
	path := []string{"presence-container"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Presence parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// JSON.
	jsonStr := config.JsonMarshal()
	targetStr := `{"presence-container": {}}`
	var result bool
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Delete success.
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Presence parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	jsonStr = config.JsonMarshal()
	targetStr = `{}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Parse failure.
	config = &Config{}
	path = []string{"normal-container"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseIncomplete {
		t.Error("Presence parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
}

func TestTypeDef(t *testing.T) {
	// Load typedef YANG.
	entry, err := YangSetUp("test-typedef.yang")
	if err != nil {
		fmt.Println(err)
	}

	// Parse success.
	config := &Config{}
	path := []string{"completed", "100"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// Parse failure.
	config = &Config{}
	path = []string{"completed", "101"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseNoMatch {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	path = []string{"completed"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
}

func TestChoice(t *testing.T) {
	// Load choice YANG.
	entry, err := YangSetUp("test-choice.yang")
	if err != nil {
		fmt.Println(err)
	}

	// Parse success.
	config := &Config{}
	path := []string{"protocol", "tcp"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	jsonStr := config.JsonMarshal()
	targetStr := `{"protocol":{"tcp":true}}`
	var result bool
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Replace case.
	path = []string{"protocol", "udp"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// JSON.
	jsonStr = config.JsonMarshal()
	targetStr = `{"protocol":{"udp":true}}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	// Parse success.
	config = &Config{}
	path = []string{"food", "pretzel"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	path = []string{"food", "beer", "sapporo"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{"food":{"beer":"sapporo","pretzel":true}}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	path = []string{"food", "chocolate", "milk"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	jsonStr = config.JsonMarshal()
	targetStr = `{"food":{"chocolate":"milk"}}`
	result, err = EqualJSON(jsonStr, targetStr)
	if err != nil || !result {
		t.Error("List JSON error for", config.JsonMarshal(), targetStr)
	}

	config = &Config{}
	path = []string{"food", "pretzel"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("typedef parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
}

func TestUnion(t *testing.T) {
	// Load union YANG.
	entry, err := YangSetUp("test-union.yang")
	if err != nil {
		t.Error(err)
		return
	}

	// Parse success.
	config := &Config{}
	path := []string{"union-address", "10.0.0.1"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("mandatory parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	config = &Config{}
	path = []string{"union-address", "2001::1"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("mandatory parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	config = &Config{}
	path = []string{"union-address", "nomatch"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseNoMatch {
		t.Error("mandatory parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
}

func TestMandatory(t *testing.T) {
	// Load multiple key YANG.
	entry, err := YangSetUp("test-list.yang")
	if err != nil {
		t.Error("Parse error", err)
		return
	}

	// Parse success test.
	config := &Config{}
	path := []string{"lists", "key", "key-value", "two", "two-value"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// MandatoryCheck.
	err = config.MandatoryCheck()
	if err != nil {
		t.Error("Must be mandatory true", err)
	}

	// Delete mandatory leaf.
	path = []string{"lists", "key", "key-value", "two", "two-value"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List delete failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// MandatoryCheck.
	err = config.MandatoryCheck()
	if err == nil {
		t.Error("Must be mandatory error")
	}

	// jsonStr = config.JsonMarshal()
	// fmt.Println(jsonStr)
}

func TestMandatoryMultiKey(t *testing.T) {
	// Load multiple key YANG.
	entry, err := YangSetUp("test-multikey.yang")
	if err != nil {
		t.Error("Parse error", err)
		return
	}

	// Parse success test.
	config := &Config{}
	path := []string{"multikey2", "first", "second"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// MandatoryCheck.
	err = config.MandatoryCheck()
	if err == nil {
		t.Error("Must be mandatory error")
	}

	// // Delete mandatory leaf.
	path = []string{"multikey2", "first", "second", "third", "third-value"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Mandatory node add failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// // MandatoryCheck.
	err = config.MandatoryCheck()
	if err != nil {
		t.Error("Mandatory error should not happen", err)
	}

	// jsonStr = config.JsonMarshal()
	// fmt.Println(jsonStr)
}

func TestMandatoryNestKey(t *testing.T) {
	// Load multiple key YANG.
	entry, err := YangSetUp("test-mandatory.yang")
	if err != nil {
		t.Error("Parse error", err)
		return
	}

	// Parse success test.
	config := &Config{}
	path := []string{"static", "route", "10.0.0.0/24", "nexthop", "10.0.0.1"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	// jsonStr = config.JsonMarshal()
	// fmt.Println(jsonStr)

	err = config.MandatoryCheck()
	if err != nil {
		t.Error("Mandatory error should not happen", err)
	}

	path = []string{"static", "route", "10.0.0.0/24", "nexthop", "10.0.0.1"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("Delete failed for", path, "result", cmd.ParseResult2String(ret))
	}

	// jsonStr = config.JsonMarshal()
	// fmt.Println(jsonStr)

	err = config.MandatoryCheck()
	if err == nil {
		t.Error("Mandatory error should happen")
	}
}

func TestContainerMandatory(t *testing.T) {
	// Load mandatory YANG.
	entry, err := YangSetUp("test-mandatory.yang")
	if err != nil {
		t.Error("Parse error", err)
		return
	}

	// Parse success test with missing mandatory.
	config := &Config{}
	path := []string{"top", "mandatory", "value"}
	ret, _, _, _ := Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	err = config.MandatoryCheck()
	if err != nil {
		t.Error("Mandatory error", err)
	}

	path = []string{"top", "value", "node"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	path = []string{"top", "mandatory", "value"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	err = config.MandatoryCheck()
	if err == nil {
		t.Error("Mandatory error should occur")
	}

	// presence container check.
	config = &Config{}
	path = []string{"top-presence", "value", "node"}
	ret, _, _, _ = Parse(path, entry, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List parse failed for", path, "result", cmd.ParseResult2String(ret))
	}
	err = config.MandatoryCheck()
	if err != nil {
		t.Error("Mandatory should not occur for presence container", err)
	}
}
