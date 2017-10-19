// Copyright 2016 CoreSwitch
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

package dependency

import (
	"reflect"
	"testing"
)

// building a graph like:
//
//       :a
//      / |
//    :b  |
//      \ |
//       :c
//        |
//       :d
//
func G1() *MapDependencyGraph {
	g1 := NewGraph()
	g1.Depend("b", "a")
	g1.Depend("c", "b")
	g1.Depend("c", "a")
	g1.Depend("d", "c")
	return g1
}

//      'one    'five
//        |       |
//      'two      |
//       / \      |
//      /   \     |
//     /     \   /
// 'three   'four
//    |      /
//  'six    /
//    |    /
//    |   /
//    |  /
//  'seven
//
func G2() *MapDependencyGraph {
	g2 := NewGraph()
	g2.Depend("two", "one")
	g2.Depend("three", "two")
	g2.Depend("four", "two")
	g2.Depend("four", "five")
	g2.Depend("six", "three")
	g2.Depend("seven", "six")
	g2.Depend("seven", "four")
	return g2
}

//               :level0
//               / | |  \
//          -----  | |   -----
//         /       | |        \
// :level1a :level1b :level1c :level1d
//         \       | |        /
//          -----  | |   -----
//               \ | |  /
//               :level2
//               / | |  \
//          -----  | |   -----
//         /       | |        \
// :level3a :level3b :level3c :level3d
//         \       | |        /
//          -----  | |   -----
//               \ | |  /
//               :level4
//
// ... and so on in a repeating pattern like that, up to :level26

func G3() *MapDependencyGraph {
	g3 := NewGraph()
	g3.Depend("level1a", "level0")
	g3.Depend("level1b", "level0")
	g3.Depend("level1c", "level0")
	g3.Depend("level1d", "level0")
	g3.Depend("level2", "level1a")
	g3.Depend("level2", "level1b")
	g3.Depend("level2", "level1c")
	g3.Depend("level2", "level1d")

	g3.Depend("level3a", "level2")
	g3.Depend("level3b", "level2")
	g3.Depend("level3c", "level2")
	g3.Depend("level3d", "level2")
	g3.Depend("level4", "level3a")
	g3.Depend("level4", "level3b")
	g3.Depend("level4", "level3c")
	g3.Depend("level4", "level3d")

	g3.Depend("level5a", "level4")
	g3.Depend("level5b", "level4")
	g3.Depend("level5c", "level4")
	g3.Depend("level5d", "level4")
	g3.Depend("level6", "level5a")
	g3.Depend("level6", "level5b")
	g3.Depend("level6", "level5c")
	g3.Depend("level6", "level5d")

	g3.Depend("level7a", "level6")
	g3.Depend("level7b", "level6")
	g3.Depend("level7c", "level6")
	g3.Depend("level7d", "level6")
	g3.Depend("level8", "level7a")
	g3.Depend("level8", "level7b")
	g3.Depend("level8", "level7c")
	g3.Depend("level8", "level7d")

	g3.Depend("level9a", "level8")
	g3.Depend("level9b", "level8")
	g3.Depend("level9c", "level8")
	g3.Depend("level9d", "level8")
	g3.Depend("level10", "level9a")
	g3.Depend("level10", "level9b")
	g3.Depend("level10", "level9c")
	g3.Depend("level10", "level9d")

	g3.Depend("level11a", "level10")
	g3.Depend("level11b", "level10")
	g3.Depend("level11c", "level10")
	g3.Depend("level11d", "level10")
	g3.Depend("level12", "level11a")
	g3.Depend("level12", "level11b")
	g3.Depend("level12", "level11c")
	g3.Depend("level12", "level11d")

	g3.Depend("level13a", "level12")
	g3.Depend("level13b", "level12")
	g3.Depend("level13c", "level12")
	g3.Depend("level13d", "level12")
	g3.Depend("level14", "level13a")
	g3.Depend("level14", "level13b")
	g3.Depend("level14", "level13c")
	g3.Depend("level14", "level13d")

	g3.Depend("level15a", "level14")
	g3.Depend("level15b", "level14")
	g3.Depend("level15c", "level14")
	g3.Depend("level15d", "level14")
	g3.Depend("level16", "level15a")
	g3.Depend("level16", "level15b")
	g3.Depend("level16", "level15c")
	g3.Depend("level16", "level15d")

	g3.Depend("level17a", "level16")
	g3.Depend("level17b", "level16")
	g3.Depend("level17c", "level16")
	g3.Depend("level17d", "level16")
	g3.Depend("level18", "level17a")
	g3.Depend("level18", "level17b")
	g3.Depend("level18", "level17c")
	g3.Depend("level18", "level17d")

	g3.Depend("level19a", "level18")
	g3.Depend("level19b", "level18")
	g3.Depend("level19c", "level18")
	g3.Depend("level19d", "level18")
	g3.Depend("level20", "level19a")
	g3.Depend("level20", "level19b")
	g3.Depend("level20", "level19c")
	g3.Depend("level20", "level19d")

	g3.Depend("level21a", "level20")
	g3.Depend("level21b", "level20")
	g3.Depend("level21c", "level20")
	g3.Depend("level21d", "level20")
	g3.Depend("level22", "level21a")
	g3.Depend("level22", "level21b")
	g3.Depend("level22", "level21c")
	g3.Depend("level22", "level21d")

	g3.Depend("level23a", "level22")
	g3.Depend("level23b", "level22")
	g3.Depend("level23c", "level22")
	g3.Depend("level23d", "level22")
	g3.Depend("level24", "level23a")
	g3.Depend("level24", "level23b")
	g3.Depend("level24", "level23c")
	g3.Depend("level24", "level23d")

	g3.Depend("level25a", "level24")
	g3.Depend("level25b", "level24")
	g3.Depend("level25c", "level24")
	g3.Depend("level25d", "level24")
	g3.Depend("level26", "level25a")
	g3.Depend("level26", "level25b")
	g3.Depend("level26", "level25c")
	g3.Depend("level26", "level25d")
	return g3
}

func TestTransitiveDependencies(t *testing.T) {
	g1 := G1()
	g1val := DependencyMap{
		"a": struct{}{},
		"b": struct{}{},
		"c": struct{}{},
	}
	if !reflect.DeepEqual(g1val, g1.TransitiveDependencies("d")) {
		t.Errorf("transitive dependencies error")
	}

	g2 := G2()
	g2val := DependencyMap{
		"one":   struct{}{},
		"two":   struct{}{},
		"three": struct{}{},
		"four":  struct{}{},
		"five":  struct{}{},
		"six":   struct{}{},
	}
	if !reflect.DeepEqual(g2val, g2.TransitiveDependencies("seven")) {
		t.Errorf("transitive dependencies error")
	}
}

func TestTransitiveDependencies2(t *testing.T) {
	g3 := G3()
	g3val := DependencyMap{
		"level0":   struct{}{},
		"level1a":  struct{}{},
		"level1b":  struct{}{},
		"level1c":  struct{}{},
		"level1d":  struct{}{},
		"level2":   struct{}{},
		"level3a":  struct{}{},
		"level3b":  struct{}{},
		"level3c":  struct{}{},
		"level3d":  struct{}{},
		"level4":   struct{}{},
		"level5a":  struct{}{},
		"level5b":  struct{}{},
		"level5c":  struct{}{},
		"level5d":  struct{}{},
		"level6":   struct{}{},
		"level7a":  struct{}{},
		"level7b":  struct{}{},
		"level7c":  struct{}{},
		"level7d":  struct{}{},
		"level8":   struct{}{},
		"level9a":  struct{}{},
		"level9b":  struct{}{},
		"level9c":  struct{}{},
		"level9d":  struct{}{},
		"level10":  struct{}{},
		"level11a": struct{}{},
		"level11b": struct{}{},
		"level11c": struct{}{},
		"level11d": struct{}{},
		"level12":  struct{}{},
		"level13a": struct{}{},
		"level13b": struct{}{},
		"level13c": struct{}{},
		"level13d": struct{}{},
		"level14":  struct{}{},
		"level15a": struct{}{},
		"level15b": struct{}{},
		"level15c": struct{}{},
		"level15d": struct{}{},
		"level16":  struct{}{},
		"level17a": struct{}{},
		"level17b": struct{}{},
		"level17c": struct{}{},
		"level17d": struct{}{},
		"level18":  struct{}{},
		"level19a": struct{}{},
		"level19b": struct{}{},
		"level19c": struct{}{},
		"level19d": struct{}{},
		"level20":  struct{}{},
		"level21a": struct{}{},
		"level21b": struct{}{},
		"level21c": struct{}{},
		"level21d": struct{}{},
		"level22":  struct{}{},
		"level23a": struct{}{},
		"level23b": struct{}{},
		"level23c": struct{}{},
		"level23d": struct{}{},
	}
	if !reflect.DeepEqual(g3val, g3.TransitiveDependencies("level24")) {
		t.Errorf("transitive dependencies error")
	}

	g3val = DependencyMap{
		"level0":   struct{}{},
		"level1a":  struct{}{},
		"level1b":  struct{}{},
		"level1c":  struct{}{},
		"level1d":  struct{}{},
		"level2":   struct{}{},
		"level3a":  struct{}{},
		"level3b":  struct{}{},
		"level3c":  struct{}{},
		"level3d":  struct{}{},
		"level4":   struct{}{},
		"level5a":  struct{}{},
		"level5b":  struct{}{},
		"level5c":  struct{}{},
		"level5d":  struct{}{},
		"level6":   struct{}{},
		"level7a":  struct{}{},
		"level7b":  struct{}{},
		"level7c":  struct{}{},
		"level7d":  struct{}{},
		"level8":   struct{}{},
		"level9a":  struct{}{},
		"level9b":  struct{}{},
		"level9c":  struct{}{},
		"level9d":  struct{}{},
		"level10":  struct{}{},
		"level11a": struct{}{},
		"level11b": struct{}{},
		"level11c": struct{}{},
		"level11d": struct{}{},
		"level12":  struct{}{},
		"level13a": struct{}{},
		"level13b": struct{}{},
		"level13c": struct{}{},
		"level13d": struct{}{},
		"level14":  struct{}{},
		"level15a": struct{}{},
		"level15b": struct{}{},
		"level15c": struct{}{},
		"level15d": struct{}{},
		"level16":  struct{}{},
		"level17a": struct{}{},
		"level17b": struct{}{},
		"level17c": struct{}{},
		"level17d": struct{}{},
		"level18":  struct{}{},
		"level19a": struct{}{},
		"level19b": struct{}{},
		"level19c": struct{}{},
		"level19d": struct{}{},
		"level20":  struct{}{},
		"level21a": struct{}{},
		"level21b": struct{}{},
		"level21c": struct{}{},
		"level21d": struct{}{},
		"level22":  struct{}{},
		"level23a": struct{}{},
		"level23b": struct{}{},
		"level23c": struct{}{},
		"level23d": struct{}{},
		"level24":  struct{}{},
		"level25a": struct{}{},
		"level25b": struct{}{},
		"level25c": struct{}{},
		"level25d": struct{}{},
	}
	if !reflect.DeepEqual(g3val, g3.TransitiveDependencies("level26")) {
		t.Errorf("transitive dependencies error")
	}
}

func TestTransitiveDependents(t *testing.T) {
	g2 := G2()

	g2five := DependencyMap{
		"four":  struct{}{},
		"seven": struct{}{},
	}
	if !reflect.DeepEqual(g2five, g2.TransitiveDependents("five")) {
		t.Errorf("transitive dependents error")
	}
	g2two := DependencyMap{
		"three": struct{}{},
		"four":  struct{}{},
		"six":   struct{}{},
		"seven": struct{}{},
	}
	if !reflect.DeepEqual(g2two, g2.TransitiveDependents("two")) {
		t.Errorf("transitive dependents error")
	}
}

func Before(l DependencyList, x, y string) bool {
	for _, val := range l {
		if x == val {
			return true
		}
		if y == val {
			return false
		}
	}
	return true
}

func TestBefore(t *testing.T) {
	if Before(DependencyList{"a", "b", "c", "d"}, "a", "b") != true {
		t.Errorf("before error")
	}
	if Before(DependencyList{"a", "b", "c", "d"}, "b", "c") != true {
		t.Errorf("before error")
	}
	if Before(DependencyList{"a", "b", "c", "d"}, "d", "c") != false {
		t.Errorf("before error")
	}
	if Before(DependencyList{"a", "b", "c", "d"}, "c", "a") != false {
		t.Errorf("before error")
	}
}

func TestTopoComparator(t *testing.T) {
	dep := DependencyList{"d", "a", "b", "foo"}
	g1 := G1()
	g1.Sort(dep)
	data := [][]string{
		{"a", "b"},
		{"a", "d"},
		{"a", "foo"},
		{"b", "d"},
		{"b", "foo"},
		{"d", "foo"},
	}
	for _, d := range data {
		if Before(dep, d[0], d[1]) != true {
			t.Errorf("before error %s %s", d[0], d[1])
		}
	}
}

func TestTopoComparator2(t *testing.T) {
	dep := DependencyList{"three", "seven", "nine", "eight", "five"}
	g2 := G2()
	g2.Sort(dep)
	data := [][]string{
		{"three", "seven"},
		{"three", "eight"},
		{"three", "nine"},
		{"five", "eight"},
		{"five", "nine"},
		{"seven", "eight"},
		{"seven", "nine"},
	}
	for _, d := range data {
		if Before(dep, d[0], d[1]) != true {
			t.Errorf("before error %s %s", d[0], d[1])
		}
	}
}

func TestTopoSort(t *testing.T) {
	g2 := G2()
	nodes := g2.Nodes()
	dep := DependencyList{}
	for k, _ := range nodes {
		dep = append(dep, k)
	}
	g2.Sort(dep)
	data := [][]string{
		{"one", "two"},
		{"one", "three"},
		{"one", "four"},
		{"one", "six"},
		{"one", "seven"},
		{"two", "three"},
		{"two", "four"},
		{"two", "six"},
		{"two", "seven"},
		{"three", "six"},
		{"three", "seven"},
		{"four", "seven"},
		{"five", "four"},
		{"five", "seven"},
		{"six", "seven"},
	}
	for _, d := range data {
		if Before(dep, d[0], d[1]) != true {
			t.Errorf("before error %s %s", d[0], d[1])
		}
	}
}

func TestNoCycles(t *testing.T) {
	g := NewGraph()
	err := g.Depend("a", "b")
	if err != nil {
		t.Errorf("cycle check error: a b")
	}
	err = g.Depend("b", "c")
	if err != nil {
		t.Errorf("cycle check error: b c")
	}
	err = g.Depend("c", "a")
	if err == nil {
		t.Errorf("cycle check error: c a")
	}
}

func TestNoSelfCycles(t *testing.T) {
	g := NewGraph()
	err := g.Depend("a", "b")
	if err != nil {
		t.Errorf("cycle check error: a b")
	}
	err = g.Depend("a", "a")
	if err == nil {
		t.Errorf("cycle check error: a a")
	}
}
