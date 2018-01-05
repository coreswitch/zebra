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

func NewAddressIPv4(addr *IfAddr) *pb.AddressIPv4 {
	return &pb.AddressIPv4{
		Addr: &pb.PrefixIPv4{
			Addr:   addr.Prefix.IP,
			Length: uint32(addr.Prefix.Length),
		},
	}
}

func NewAddressIPv6(addr *IfAddr) *pb.AddressIPv6 {
	return &pb.AddressIPv6{
		Addr: &pb.PrefixIPv6{
			Addr:   addr.Prefix.IP,
			Length: uint32(addr.Prefix.Length),
		},
	}
}

func NewInterfaceUpdateFull(op pb.Op, ifp *Interface) *pb.InterfaceUpdate {
	update := NewInterfaceUpdate(op, ifp)
	for _, addr := range ifp.Addrs[AFI_IP] {
		update.AddrIpv4 = append(update.AddrIpv4, NewAddressIPv4(addr))
	}
	for _, addr := range ifp.Addrs[AFI_IP6] {
		update.AddrIpv6 = append(update.AddrIpv6, NewAddressIPv6(addr))
	}
	return update
}

func WatcherNotifyAllInterfaces(w Watcher, vrf *Vrf) {
	for _, ifp := range vrf.IfMap {
		w.Notify(NewInterfaceUpdateFull(pb.Op_InterfaceAdd, ifp))
	}
}

func watcherNotifyInterface(ifp *Interface, op pb.Op) {
	for _, w := range ifp.Vrf.Watchers[WATCH_TYPE_INTERFACE] {
		w.Notify(NewInterfaceUpdate(op, ifp))
	}
}

func WatcherNotifyInterfaceAdd(ifp *Interface) {
	watcherNotifyInterface(ifp, pb.Op_InterfaceAdd)
}

func WatcherNotifyInterfaceDelete(ifp *Interface) {
	watcherNotifyInterface(ifp, pb.Op_InterfaceDelete)
}

func WatcherNotifyInterfaceNameChange(ifp *Interface) {
	watcherNotifyInterface(ifp, pb.Op_InterfaceNameChange)
}

func WatcherNotifyInterfaceMtuChange(ifp *Interface) {
	watcherNotifyInterface(ifp, pb.Op_InterfaceMtuChange)
}

func WatcherNotifyInterfaceUp(ifp *Interface) {
	watcherNotifyInterface(ifp, pb.Op_InterfaceUp)
}

func WatcherNotifyInterfaceDown(ifp *Interface) {
	watcherNotifyInterface(ifp, pb.Op_InterfaceDown)
}

func WatcherNotifyInterfaceFlagChange(ifp *Interface) {
	watcherNotifyInterface(ifp, pb.Op_InterfaceFlagChange)
}

func WatcherNotifyAddressAdd() {
}

func WatcherNotifyAddressDelete() {
}
