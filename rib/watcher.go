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

package rib

import pb "github.com/coreswitch/zebra/proto"

const (
	WATCH_TYPE_INTERFACE   = 0
	WATCH_TYPE_ROUTER_ID   = 1
	WATCH_TYPE_RIB         = 2
	WATCH_TYPE_RIB_DEFAULT = 3
	WATCH_TYPE_MAX         = 4
)

type Watcher interface {
	Notify(interface{})
}

type Watchers []Watcher

func NewInterfaceUpdate(op pb.Op, ifp *Interface) *pb.InterfaceUpdate {
	return &pb.InterfaceUpdate{
		Op:     op,
		VrfId:  uint32(ifp.VrfIndex),
		Name:   ifp.Name,
		Index:  uint32(ifp.Index),
		Mtu:    ifp.Mtu,
		Metric: ifp.Metric,
	}
}

func WatcherNotifyAllInterface(w Watcher, vrf *Vrf) {
	for _, ifp := range vrf.IfMap {
		w.Notify(NewInterfaceUpdate(pb.Op_InterfaceAdd, ifp))
	}
}
