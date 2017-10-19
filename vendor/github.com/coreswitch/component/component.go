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

package component

import (
	"github.com/coreswitch/dependency"
)

type Component interface {
	Start() Component
	Stop() Component
}

type SystemMap map[string]Component

type ComponentItem struct {
	Component
	depends []string
}

func (this *ComponentItem) Start() Component {
	return this.Component.Start()
}

func (this *ComponentItem) Stop() Component {
	return this.Component.Stop()
}

func ComponentWith(c Component, deps ...string) Component {
	return &ComponentItem{Component: c, depends: deps}
}

func (this *SystemMap) Graph() *dependency.MapDependencyGraph {
	graph := dependency.NewGraph()
	for key, value := range *this {
		if item, ok := value.(*ComponentItem); ok {
			for _, dep := range item.depends {
				graph.Depend(key, dep)
			}
		}
	}
	return graph
}

func (this *SystemMap) Start() {
	graph := this.Graph()
	for _, key := range graph.TopoSort() {
		if component, ok := (*this)[key]; ok {
			component.Start()
		}
	}
}

func (this *SystemMap) Stop() {
	graph := this.Graph()
	topo := graph.TopoSort()
	for i := len(topo) - 1; i >= 0; i-- {
		if component, ok := (*this)[topo[i]]; ok {
			component.Stop()
		}
	}
}
