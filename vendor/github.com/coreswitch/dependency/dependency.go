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
	"fmt"
	"math"
	"sort"
)

type Dependency map[string]DependencyMap
type DependencyMap map[string]struct{}
type DependencyList []string

type MapDependencyGraph struct {
	dependencies Dependency
	dependents   Dependency
}

func updateIn(dep *Dependency, key string, value string) {
	if _, ok := (*dep)[key]; !ok {
		(*dep)[key] = DependencyMap{}
	}
	(*dep)[key][value] = struct{}{}
}

func expand(neighbors *Dependency, expanded DependencyMap, node string) DependencyMap {
	unexpanded := (*neighbors)[node]
	for n, _ := range unexpanded {
		if _, ok := expanded[n]; !ok {
			expanded[n] = struct{}{}
			expand(neighbors, expanded, n)
		}
	}
	return expanded
}

func NewGraph() *MapDependencyGraph {
	return &MapDependencyGraph{
		dependencies: Dependency{},
		dependents:   Dependency{},
	}
}

func (g *MapDependencyGraph) Copy() *MapDependencyGraph {
	graph := NewGraph()
	for key, value := range g.dependencies {
		(*graph).dependencies[key] = DependencyMap{}
		for k, _ := range value {
			(*graph).dependencies[key][k] = struct{}{}
		}
	}
	for key, value := range g.dependents {
		(*graph).dependents[key] = DependencyMap{}
		for k, _ := range value {
			(*graph).dependents[key][k] = struct{}{}
		}
	}
	return graph
}

func (g *MapDependencyGraph) RemoveEdge(node, dep string) *MapDependencyGraph {
	graph := g.Copy()
	delete(graph.dependencies[node], dep)
	delete(graph.dependents[dep], node)
	return graph
}

func (g *MapDependencyGraph) RemoveNode(node string) *MapDependencyGraph {
	graph := g.Copy()
	delete(graph.dependencies, node)
	return graph
}

func (g *MapDependencyGraph) IsDepend(node, dep string) bool {
	expanded := g.TransitiveDependencies(node)
	if _, ok := expanded[dep]; !ok {
		return false
	} else {
		return true
	}
}

func (g *MapDependencyGraph) Depend(node, dep string) error {
	if node == dep || g.IsDepend(dep, node) {
		return fmt.Errorf("Circular dependency between %s and %s", node, dep)
	}
	updateIn(&g.dependencies, node, dep)
	updateIn(&g.dependents, dep, node)
	return nil
}

func (g *MapDependencyGraph) Nodes() DependencyMap {
	nodes := DependencyMap{}
	for key, _ := range g.dependencies {
		nodes[key] = struct{}{}
	}
	for key, _ := range g.dependents {
		nodes[key] = struct{}{}
	}
	return nodes
}

func Transitive(neighbors *Dependency, node string) DependencyMap {
	return expand(neighbors, DependencyMap{}, node)
}

func (g *MapDependencyGraph) TransitiveDependencies(node string) DependencyMap {
	return Transitive(&g.dependencies, node)
}

func (g *MapDependencyGraph) TransitiveDependents(node string) DependencyMap {
	return Transitive(&g.dependents, node)
}

func (g *MapDependencyGraph) ImmediateDependents(node string) DependencyMap {
	if val, ok := g.dependents[node]; ok {
		return val
	} else {
		return DependencyMap{}
	}
}

func (g *MapDependencyGraph) ImmediateDependencies(node string) DependencyMap {
	if val, ok := g.dependencies[node]; ok {
		return val
	} else {
		return DependencyMap{}
	}
}

func (g *MapDependencyGraph) TopoSort() DependencyList {
	sorted := DependencyList{}
	graph := g.Copy()
	todo := DependencyList{}
	for n, _ := range g.Nodes() {
		d := g.ImmediateDependents(n)
		if len(d) == 0 {
			todo = append(todo, n)
		}
	}
Loop:
	if len(todo) == 0 {
		return sorted
	}
	node := todo[0]
	more := todo[1:]
	deps := graph.ImmediateDependencies(node)

	add := DependencyList{}
	for d, _ := range deps {
		graph = graph.RemoveEdge(node, d)
		if len(graph.ImmediateDependents(d)) == 0 {
			add = append(add, d)
		}
	}
	sorted = append(DependencyList{node}, sorted...)
	graph = graph.RemoveNode(node)
	todo = append(more, add...)
	goto Loop
}

func (l *DependencyList) Pos(s string) int {
	for pos, val := range *l {
		if val == s {
			return pos
		}
	}
	return math.MaxInt32
}

type By func(x, y string) bool

type TopoSorter struct {
	dep DependencyList
	by  By
}

func (by By) Sort(dep DependencyList) {
	d := TopoSorter{
		dep: dep,
		by:  by,
	}
	sort.Sort(d)
}

func (t TopoSorter) Len() int {
	return len(t.dep)
}

func (t TopoSorter) Swap(i, j int) {
	t.dep[i], t.dep[j] = t.dep[j], t.dep[i]
}

func (t TopoSorter) Less(i, j int) bool {
	return t.by(t.dep[i], t.dep[j])
}

func (g *MapDependencyGraph) Sort(dep DependencyList) {
	topo := g.TopoSort()
	comparator := func(x, y string) bool {
		if topo.Pos(x) < topo.Pos(y) {
			return true
		} else {
			return false
		}
	}
	By(comparator).Sort(dep)
}
