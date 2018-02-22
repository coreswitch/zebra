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

type WatcherRedist struct {
	typ [RIB_MAX]Watchers
	def Watchers
}

type Vrf struct {
	Name        string
	Id          uint32
	IfTable     *netutil.Ptree
	IfMap       map[string]*Interface
	ribTable    [AFI_MAX]*netutil.Ptree
	nhopTable   [AFI_MAX]*netutil.Ptree
	arpTable    [AFI_MAX]*netutil.Ptree
	staticTable [AFI_MAX]*netutil.Ptree
	redist      [AFI_MAX]WatcherRedist
	routerId    RouterId
	ZServer     *ZServer
	Mutex       sync.RWMutex
	IfMutex     sync.Mutex
	Watcher     map[*IfWatcher]bool // Will be merged to Watchers
	WMutex      sync.Mutex
	Watchers    []Watchers
}

// For all VRF redistribute.
var Redist [AFI_MAX]WatcherRedist

func VrfDefault() *Vrf {
	return VrfTable[0]
}

func VrfLookupByIndex(id uint32) *Vrf {
	if id < 0 || id >= VRF_ID_MAX {
		return nil
	}
	return VrfTable[id]
}

func VrfLookupByName(name string) *Vrf {
	if v, ok := VrfMap[name]; ok {
		return v
	}
	return nil
}

func NewVrf(name string, vrfId uint32) *Vrf {
	vrf := &Vrf{
		Name:     name,
		Id:       vrfId,
		IfTable:  netutil.NewPtree(32),
		IfMap:    make(map[string]*Interface),
		Watcher:  make(map[*IfWatcher]bool),
		Watchers: make([]Watchers, WATCH_TYPE_MAX),
	}
	vrf.ribTable[AFI_IP] = netutil.NewPtree(32)
	vrf.ribTable[AFI_IP6] = netutil.NewPtree(128)
	vrf.nhopTable[AFI_IP] = netutil.NewPtree(32)
	vrf.nhopTable[AFI_IP6] = netutil.NewPtree(128)
	vrf.arpTable[AFI_IP] = netutil.NewPtree(32)
	vrf.arpTable[AFI_IP6] = netutil.NewPtree(128)
	vrf.staticTable[AFI_IP] = netutil.NewPtree(32)
	vrf.staticTable[AFI_IP6] = netutil.NewPtree(128)

	vrf.routerId.Init()

	VrfTable[vrf.Id] = vrf
	VrfMap[vrf.Name] = vrf

	if vrfId != 0 {
		vrf.ZServer = ZServerStart("unix", fmt.Sprintf("/var/run/zserv-vrf%d.api", vrfId), vrfId)
	}

	return vrf
}

func VrfDefaultZservStart() {
	vrf := VrfDefault()
	if vrf != nil {
		vrf.ZServer = ZServerStart("unix", "/var/run/zserv.api", 0)
	}
}

func VrfAssignIndex() uint32 {
	for i := VRF_ID_MIN + 100; i <= VRF_ID_MAX; i++ {
		if VrfTable[i] == nil {
			return uint32(i)
		}
	}
	return 0
}

func VrfAssign(name string, id uint32) (*Vrf, error) {
	if id == 0 {
		id = VrfAssignIndex()
		if id == 0 {
			return nil, fmt.Errorf("Can't assing VRF index")
		}
	}
	v := NewVrf(name, id)

	return v, nil
}

func VrfExtractIndex(name string) uint32 {
	r := regexp.MustCompile("vrf(\\d+)")
	matches := r.FindAllStringSubmatch(name, -1)
	if matches != nil && len(matches) > 0 {
		match := matches[0]
		if len(match) >= 2 {
			fmt.Println("VrfExtractIndex", match[1])
			index, _ := strconv.Atoi(match[1])
			return uint32(index)
		}
	}
	return 0
}

func VrfStop() {
	for _, vrf := range VrfMap {
		if vrf.Id != 0 {
			server.VrfDelete(vrf.Name)
		}
	}
}
