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

package zebra

import (
	fmt "fmt"
	"net"

	"github.com/coreswitch/netutil"
)

type Rib struct {
	Tag uint32
}

func (v Address) MarshalText() (text []byte, err error) {
	p := netutil.PrefixFromIPPrefixlen(v.Addr.Addr, int(v.Addr.Length))
	return []byte(p.String()), nil
}

func (v Prefix) MarshalText() (text []byte, err error) {
	p := netutil.PrefixFromIPPrefixlen(v.Addr, int(v.Length))
	return []byte(p.String()), nil
}

func (v HwAddr) MarshalText() (text []byte, err error) {
	//haddr := net.HardwareAddr{}
	haddr := net.HardwareAddr(v.Addr)
	return []byte(haddr.String()), nil
}

func (v RouterIdUpdate) MarshalText() (text []byte, err error) {
	str := fmt.Sprintf(`vrf_id %d router_id:"%s"`, v.VrfId, net.IP(v.RouterId).String())
	return []byte(str), nil
}
