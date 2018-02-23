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
	"github.com/coreswitch/netutil"
	pb "github.com/coreswitch/zebra/proto"
)

func PrefixFromPb(from *pb.Prefix) *netutil.Prefix {
	p := &netutil.Prefix{
		IP:     from.Addr,
		Length: int(from.Length),
	}
	return p
}

type Address struct {
	Address *netutil.Prefix
	Flags   uint32
}

func AddressFromPb(from *pb.Address) *Address {
	addr := &Address{
		Address: PrefixFromPb(from.Addr),
		Flags:   from.Flags,
	}
	return addr
}
