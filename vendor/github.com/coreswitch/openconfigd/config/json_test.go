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
	"testing"

	"github.com/coreswitch/cmd"
)

func TestJson(t *testing.T) {
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

	// Delete mandatory leaf.
	path = []string{"lists", "key", "key-value", "two", "two-value"}
	ret, _, _, _ = ParseDelete(path, config, nil)
	if ret != cmd.ParseSuccess {
		t.Error("List delete failed for", path, "result", cmd.ParseResult2String(ret))
	}
}
