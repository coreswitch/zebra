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
	"strings"
	"syscall"
)

func IfFlagsString(flags uint32) string {
	strs := make([]string, 0)
	if (flags & syscall.IFF_UP) != 0 {
		strs = append(strs, "UP")
	}
	if (flags & syscall.IFF_BROADCAST) != 0 {
		strs = append(strs, "BROADCAST")
	}
	if (flags & syscall.IFF_LOOPBACK) != 0 {
		strs = append(strs, "LOOPBACK")
	}
	if (flags & syscall.IFF_POINTOPOINT) != 0 {
		strs = append(strs, "POINTOPOINT")
	}
	if (flags & syscall.IFF_MULTICAST) != 0 {
		strs = append(strs, "MULTICAST")
	}
	return strings.Join(strs, ",")
}
