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

package fea

import (
	pb "github.com/coreswitch/zebra/api"
)

type Interface struct {
	Name     string     `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Index    uint32     `protobuf:"varint,4,opt,name=index" json:"index,omitempty"`
	Flags    uint32     `protobuf:"varint,5,opt,name=flags" json:"flags,omitempty"`
	Mtu      uint32     `protobuf:"varint,6,opt,name=mtu" json:"mtu,omitempty"`
	Metric   uint32     `protobuf:"varint,7,opt,name=metric" json:"metric,omitempty"`
	HwAddr   []byte     `protobuf:"bytes,8,opt,name=hw_addr,json=hwAddr" json:"hw_addr,omitempty"`
	AddrIpv4 []*Address `protobuf:"bytes,9,rep,name=addr_ipv4,json=addrIpv4" json:"addr_ipv4,omitempty"`
	AddrIpv6 []*Address `protobuf:"bytes,10,rep,name=addr_ipv6,json=addrIpv6" json:"addr_ipv6,omitempty"`
	Info     interface{}
}

func InterfaceFromPb(from *pb.InterfaceUpdate) *Interface {
	ifp := &Interface{
		Name:   from.Name,
		Index:  from.Index,
		Flags:  from.Flags,
		Mtu:    from.Mtu,
		Metric: from.Metric,
	}
	if from.HwAddr != nil {
		ifp.HwAddr = from.HwAddr.Addr
	}
	for _, addr := range from.AddrIpv4 {
		ifp.AddrIpv4 = append(ifp.AddrIpv4, AddressFromPb(addr))
	}
	for _, addr := range from.AddrIpv6 {
		ifp.AddrIpv6 = append(ifp.AddrIpv6, AddressFromPb(addr))
	}

	return ifp
}
