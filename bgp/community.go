// Copyright 2017 zebra project
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

package bgp

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Community []uint32

const (
	COMMUNITY_INTERNET     uint32 = 0x00000000
	COMMUNITY_GSHUT               = 0xffff0000
	COMMUNITY_NO_EXPORT           = 0xffffff01
	COMMUNITY_NO_ADVERTISE        = 0xffffff02
	COMMUNITY_LOCAL_AS            = 0xffffff03
	COMMUNITY_NO_PEER             = 0xffffff04
)

var WellKnownCommunityMap = map[uint32]string{
	COMMUNITY_INTERNET:     "internet",
	COMMUNITY_GSHUT:        "gshut",
	COMMUNITY_NO_EXPORT:    "no-export",
	COMMUNITY_NO_ADVERTISE: "no-advertise",
	COMMUNITY_LOCAL_AS:     "local-AS",
	COMMUNITY_NO_PEER:      "nopeer",
}

var WellKnownCommunityStrMap = map[string]uint32{
	"internet":     COMMUNITY_INTERNET,
	"gshut":        COMMUNITY_GSHUT,
	"no-export":    COMMUNITY_NO_EXPORT,
	"no-advertise": COMMUNITY_NO_ADVERTISE,
	"local-AS":     COMMUNITY_LOCAL_AS,
	"nopeer":       COMMUNITY_NO_PEER,
}

func (lhs Community) Equal(rhs Community) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for pos, _ := range lhs {
		if lhs[pos] != rhs[pos] {
			return false
		}
	}
	return true
}

func (c Community) String() string {
	str := []string{}
	for _, v := range c {
		s, ok := WellKnownCommunityMap[v]
		if ok {
			str = append(str, s)
		} else {
			str = append(str, fmt.Sprintf("%d:%d", v>>16, 0x0000ffff&v))
		}
	}
	return strings.Join(str, " ")
}

func (c Community) MarshalJSON() ([]byte, error) {
	str := []string{}
	for _, v := range c {
		s, ok := WellKnownCommunityMap[v]
		if ok {
			str = append(str, strconv.Quote(s))
		} else {
			str = append(str, strconv.Quote(fmt.Sprintf("%d:%d", v>>16, 0x0000ffff&v)))
		}
	}
	return []byte("[" + strings.Join(str, ",") + "]"), nil
}

func CommunityValParse(s string) (uint32, error) {
	str := strings.Split(s, ":")
	if len(str) == 2 {
		as, err := strconv.ParseUint(str[0], 10, 16)
		if err != nil {
			return 0, err
		}
		val, err := strconv.ParseUint(str[1], 10, 16)
		if err != nil {
			return 0, err
		}
		return uint32(as<<16 | val), nil
	} else {
		val, err := strconv.ParseUint(str[0], 10, 32)
		if err != nil {
			return 0, err
		}
		return uint32(val), nil
	}
}

func CommunityParse(s string) (Community, error) {
	c := Community{}
	for _, str := range strings.Split(s, " ") {
		v, ok := WellKnownCommunityStrMap[str]
		if ok {
			c = append(c, v)
		} else {
			v, err := CommunityValParse(str)
			if err != nil {
				return nil, err
			}
			c = append(c, v)
		}
	}
	return c, nil
}

func (c Community) SortUnique() Community {
	sort.Slice(c, func(i, j int) bool { return c[i] < c[j] })
	d := Community{}
	var prev uint32
	for i := 0; i < len(c); i++ {
		if i == 0 || c[i] != prev {
			d = append(d, c[i])
			prev = c[i]
		}
	}
	return d
}
