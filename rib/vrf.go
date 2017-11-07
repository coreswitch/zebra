// Copyright 2016 Zebra Project
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

package rib

import (
	"fmt"
	"regexp"
	"strconv"
	"sync"

	"github.com/coreswitch/netutil"
)

const (
	VRF_ID_MIN     = 1
	VRF_ID_MAX     = 253
	VRF_TABLE_SIZE = 254
)

var (
	VrfTable = [VRF_TABLE_SIZE]*Vrf{}
	VrfMap   = map[string]*Vrf{}
	VrfMutex sync.RWMutex
)

type Vrf struct {
	Name        string
	Index       int
	IfTable     *netutil.Ptree
	IfMap       map[string]*Interface
	ribTable    [AFI_MAX]*netutil.Ptree
	nhopTable   [AFI_MAX]*netutil.Ptree
	arpTable    [AFI_MAX]*netutil.Ptree
	staticTable [AFI_MAX]*netutil.Ptree
	routerId    RouterId
	ZServer     *ZServer
	Mutex       sync.RWMutex
	IfMutex     sync.Mutex
	Watcher     map[*IfWatcher]bool
	WMutex      sync.Mutex
}

func VrfDefault() *Vrf {
	return VrfTable[0]
}

func VrfLookupByIndex(index int) *Vrf {
	if index < 0 || index >= VRF_ID_MAX {
		return nil
	}
	return VrfTable[index]
}

func VrfLookupByName(name string) *Vrf {
	if v, ok := VrfMap[name]; ok {
		return v
	}
	return nil
}

func NewVrf(name string, index int) *Vrf {
	v := &Vrf{
		Name:    name,
		Index:   index,
		IfTable: netutil.NewPtree(32),
		IfMap:   make(map[string]*Interface),
		Watcher: make(map[*IfWatcher]bool),
	}
	v.ribTable[AFI_IP] = netutil.NewPtree(32)
	v.ribTable[AFI_IP6] = netutil.NewPtree(128)
	v.nhopTable[AFI_IP] = netutil.NewPtree(32)
	v.nhopTable[AFI_IP6] = netutil.NewPtree(128)
	v.arpTable[AFI_IP] = netutil.NewPtree(32)
	v.arpTable[AFI_IP6] = netutil.NewPtree(128)
	v.staticTable[AFI_IP] = netutil.NewPtree(32)
	v.staticTable[AFI_IP6] = netutil.NewPtree(128)

	v.routerId.Init()

	VrfTable[v.Index] = v
	VrfMap[v.Name] = v

	if index != 0 {
		v.ZServer = ZServerStart("unix", fmt.Sprintf("/var/run/zserv-vrf%d.api", index), index)
	}

	return v
}

func VrfDefaultZservStart() {
	vrf := VrfDefault()
	if vrf != nil {
		vrf.ZServer = ZServerStart("unix", "/var/run/zserv.api", 0)
	}
}

func VrfAssignIndex() int {
	for i := VRF_ID_MIN + 100; i <= VRF_ID_MAX; i++ {
		if VrfTable[i] == nil {
			return i
		}
	}
	return 0
}

func VrfAssign(name string, index int) (*Vrf, error) {
	if index == 0 {
		index = VrfAssignIndex()
		if index == 0 {
			return nil, fmt.Errorf("Can't assing VRF index")
		}
	}
	v := NewVrf(name, index)

	return v, nil
}

func VrfExtractIndex(name string) int {
	r := regexp.MustCompile("vrf(\\d+)")
	matches := r.FindAllStringSubmatch(name, -1)
	if matches != nil && len(matches) > 0 {
		match := matches[0]
		if len(match) >= 2 {
			fmt.Println("VrfExtractIndex", match[1])
			index, _ := strconv.Atoi(match[1])
			return index
		}
	}
	return 0
}

func VrfStop() {
	for _, vrf := range VrfMap {
		if vrf.Index != 0 {
			server.VrfDelete(vrf.Name)
		}
	}
}
