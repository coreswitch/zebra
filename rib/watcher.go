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

import (
	pb "github.com/coreswitch/zebra/proto"
)

const (
	WATCH_TYPE_INTERFACE      = 0
	WATCH_TYPE_ROUTER_ID      = 1
	WATCH_TYPE_REDIST         = 2
	WATCH_TYPE_REDIST_DEFAULT = 3
	WATCH_TYPE_MAX            = 4
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
		Flags:  ifp.Flags,
		Mtu:    ifp.Mtu,
		Metric: ifp.Metric,
		HwAddr: &pb.HwAddr{Addr: ifp.HwAddr},
	}
}

func NewAddress(addr *IfAddr) *pb.Address {
	return &pb.Address{
		Addr: &pb.Prefix{
			Addr:   addr.Prefix.IP,
			Length: uint32(addr.Prefix.Length),
		},
	}
}

func NewInterfaceUpdateFull(op pb.Op, ifp *Interface) *pb.InterfaceUpdate {
	update := NewInterfaceUpdate(op, ifp)
	for _, addr := range ifp.Addrs[AFI_IP] {
		update.AddrIpv4 = append(update.AddrIpv4, NewAddress(addr))
	}
	for _, addr := range ifp.Addrs[AFI_IP6] {
		update.AddrIpv6 = append(update.AddrIpv6, NewAddress(addr))
	}
	return update
}

func NotifyInterfaces(w Watcher, vrf *Vrf) {
	for _, ifp := range vrf.IfMap {
		w.Notify(NewInterfaceUpdateFull(pb.Op_InterfaceAdd, ifp))
	}
}

func (ifp *Interface) NotifyInterface(op pb.Op) {
	for _, w := range ifp.Vrf.Watchers[WATCH_TYPE_INTERFACE] {
		w.Notify(NewInterfaceUpdate(op, ifp))
	}
}

func (ifp *Interface) NotifyInterfaceAdd() {
	ifp.NotifyInterface(pb.Op_InterfaceAdd)
}

func (ifp *Interface) NotifyInterfaceDelete() {
	ifp.NotifyInterface(pb.Op_InterfaceDelete)
}

func (ifp *Interface) NotifyInterfaceNameChange() {
	ifp.NotifyInterface(pb.Op_InterfaceNameChange)
}

func (ifp *Interface) NotifyInterfaceMtuChange() {
	ifp.NotifyInterface(pb.Op_InterfaceMtuChange)
}

func (ifp *Interface) NotifyInterfaceUp() {
	ifp.NotifyInterface(pb.Op_InterfaceUp)
}

func (ifp *Interface) NotifyInterfaceDown() {
	ifp.NotifyInterface(pb.Op_InterfaceDown)
}

func (ifp *Interface) NotifyInterfaceFlagChange() {
	ifp.NotifyInterface(pb.Op_InterfaceFlagChange)
}

func NewInterfaceAddrUpdate(op pb.Op, ifp *Interface, addr *IfAddr) *pb.InterfaceUpdate {
	update := NewInterfaceUpdate(op, ifp)
	switch addr.Prefix.AFI() {
	case AFI_IP:
		update.AddrIpv4 = append(update.AddrIpv4, NewAddress(addr))
	case AFI_IP6:
		update.AddrIpv6 = append(update.AddrIpv6, NewAddress(addr))
	}
	return update
}

func watcherNotifyIfAddr(op pb.Op, ifp *Interface, addr *IfAddr) {
	for _, w := range ifp.Vrf.Watchers[WATCH_TYPE_INTERFACE] {
		w.Notify(NewInterfaceAddrUpdate(op, ifp, addr))
	}
}

func WatcherNotifyAddressAdd(ifp *Interface, addr *IfAddr) {
	watcherNotifyIfAddr(pb.Op_InterfaceAddrAdd, ifp, addr)
}

func WatcherNotifyAddressDelete(ifp *Interface, addr *IfAddr) {
	watcherNotifyIfAddr(pb.Op_InterfaceAddrDelete, ifp, addr)
}

func NotifyRouterId(w Watcher, vrf *Vrf) {
	w.Notify(&pb.RouterIdUpdate{
		VrfId:    uint32(vrf.Id),
		RouterId: vrf.RouterId(),
	})
}

func (vrf *Vrf) NotifyRouterId() {
	for _, w := range vrf.Watchers[WATCH_TYPE_ROUTER_ID] {
		NotifyRouterId(w, vrf)
	}
}
