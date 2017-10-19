// Copyright 2016, 2017 CoreSwitch
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

package cmd

import (
	"net"
	"strconv"
	"strings"
)

type MatchType int

const (
	MatchTypeNone MatchType = iota
	MatchTypeIncomplete
	MatchTypeLine
	MatchTypeWord
	MatchTypeIPv6
	MatchTypeIPv6Prefix
	MatchTypeIPv4
	MatchTypeIPv4Prefix
	MatchTypeRange
	MatchTypePartial
	MatchTypeExact
)

type MatchState struct {
	complete bool
	match    MatchType
	count    int
	pos      int
	node     *Node
	comps    CompSlice
}

func MatchKeyword(str, name string) (pos int, match MatchType) {
	pos = longestCommonPrefix(str, name)

	if !isDelimiter(str, pos) {
		match = MatchTypeNone
		return
	}

	if isDelimiter(name, pos) {
		match = MatchTypeExact
	} else {
		match = MatchTypePartial
	}
	return
}

func MatchWord(str string) (pos int, match MatchType) {
	for pos = 0; pos < len(str); pos++ {
		if isWhiteSpace(str, pos) {
			break
		}
	}
	if pos == 0 {
		return pos, MatchTypeNone
	}
	return pos, MatchTypeWord
}

func MatchLine(str string) (pos int, match MatchType) {
	for pos = 0; pos < len(str); pos++ {
	}
	if pos == 0 {
		return pos, MatchTypeNone
	}
	return pos, MatchTypeLine
}

func MatchIPv4(str string) (pos int, match MatchType) {
	dots := 0
	numsNotSeen := true
	nums := 0
	cp := 0

	match = MatchTypeNone

	for pos = 0; pos < len(str); pos++ {
		if str[pos] == '.' {
			if dots > 3 {
				return
			}
			cp = pos + 1
			dots++
			numsNotSeen = true
			continue
		}
		if isWhiteSpace(str, pos) {
			break
		}
		if !isDigit(str, pos) {
			return
		}
		// Digit
		digitstr := str[cp : pos+1]
		digit, err := strconv.Atoi(digitstr)
		if err != nil || digit > 255 {
			return
		}
		// `nums' must be sequence of digit.
		if numsNotSeen == true {
			numsNotSeen = false
			nums++
		}
	}
	if nums > 4 || dots > 3 {
		return
	}
	if nums < 4 || dots < 3 {
		match = MatchTypeIncomplete
		return
	}

	match = MatchTypeIPv4
	return
}

func MatchIPv4Prefix(line string) (pos int, match MatchType) {
	match = MatchTypeNone
	numsNotSeen := true

	p := strings.IndexByte(line, '/')
	if p < 0 {
		pos, match = MatchIPv4(line)
		if match == MatchTypeNone {
			return
		} else {
			match = MatchTypeIncomplete
			return
		}
	}

	pos, match = MatchIPv4(line[0:p])
	if match != MatchTypeIPv4 {
		return
	}

	pos = p + 1
	for ; pos < len(line); pos++ {
		if isWhiteSpace(line, pos) {
			break
		}
		numsNotSeen = false
		digitstr := line[p+1 : pos+1]
		digit, err := strconv.Atoi(digitstr)
		if err != nil || digit > 32 {
			match = MatchTypeNone
			return
		}
	}

	if numsNotSeen == true {
		match = MatchTypeIncomplete
		return
	}

	match = MatchTypeIPv4Prefix
	return
}

func MatchIPv6(line string) (pos int, match MatchType) {
	match = MatchTypeNone
	const IPV6_ADDRSTRLEN = 46

	if len(line) == 0 {
		match = MatchTypeIncomplete
		return
	}
	for pos = 0; pos < len(line); pos++ {
		p := strings.IndexByte("0123456789abcdefABCDEF:.%", line[pos])
		if p < 0 {
			break
		}
	}
	str := line[:pos]
	if len(str) == 0 || len(str) > IPV6_ADDRSTRLEN {
		return
	}
	ip := net.ParseIP(str)
	if ip == nil {
		match = MatchTypeIncomplete
		return
	}
	ipv4 := ip.To4()
	if ipv4 != nil {
		match = MatchTypeNone
		return
	}

	match = MatchTypeIPv6
	return
}

func MatchIPv6Prefix(line string) (pos int, match MatchType) {
	match = MatchTypeNone
	numsNotSeen := true

	p := strings.IndexByte(line, '/')
	if p < 0 {
		pos, match = MatchIPv6(line)
		if match == MatchTypeNone {
			return
		} else {
			// Even function returns matchIPv6, convert it to incomplete match.
			match = MatchTypeIncomplete
			return
		}
	}

	pos, match = MatchIPv6(line[0:p])
	if match != MatchTypeIPv6 {
		return
	}

	pos = p + 1
	for ; pos < len(line); pos++ {
		if isDigit(line, pos) {
			numsNotSeen = false
			digitstr := line[p+1 : pos+1]
			digit, err := strconv.Atoi(digitstr)
			if err != nil || digit > 128 {
				match = MatchTypeNone
				return
			}
		} else {
			break
		}
	}

	if numsNotSeen == true {
		match = MatchTypeIncomplete
		return
	}

	match = MatchTypeIPv6Prefix
	return
}

func MatchRange(line string, min, max uint64) (pos int, match MatchType) {
	match = MatchTypeNone

	for pos = 0; pos < len(line); pos++ {
		if !isDigit(line, pos) {
			break
		}
	}
	if pos == 0 {
		return
	}
	digistr := line[0:pos]

	val, ok := strconv.ParseUint(digistr, 10, 64)
	if ok != nil {
		return
	}

	if min <= val && val <= max {
		match = MatchTypeRange
		return
	}

	return
}

func (n *Node) Match(str string, name string) (pos int, match MatchType) {
	switch n.Type {
	case NodeKeyword:
		return MatchKeyword(str, n.Name)
	case NodeDynamic:
		return MatchKeyword(str, name)
	case NodeWord:
		return MatchWord(str)
	case NodeLine:
		return MatchLine(str)
	case NodeIPv4:
		return MatchIPv4(str)
	case NodeIPv4Prefix:
		return MatchIPv4Prefix(str)
	case NodeIPv6:
		return MatchIPv6(str)
	case NodeIPv6Prefix:
		return MatchIPv6Prefix(str)
	case NodeRange:
		return MatchRange(str, n.Min, n.Max)
	default:
		return 0, MatchTypeNone
	}
}

func (n *Node) MatchNode(str string, name string, state *MatchState) {
	pos, match := n.Match(str, name)
	if match == MatchTypeNone {
		return
	}
	if state.complete {
		if n.Type == NodeDynamic {
			state.comps = append(state.comps, &Comp{Name: name, Help: "", Dir: false, Additive: false})
		} else {
			state.comps = append(state.comps, &Comp{Name: n.Name, Help: n.Help, Dir: false, Additive: false})
		}
	}
	if match > state.match {
		state.match = match
		state.pos = pos
		state.node = n
		state.count = 1
	} else if match == state.match {
		state.count++
	}
}

var (
	DynamicFunc func([]string, string, []string) []string
)

func (n *Node) MatchDynamic(line string, command []string, state *MatchState) {
	if DynamicFunc != nil {
		list := DynamicFunc(command, n.Module, n.Dynamic)
		for _, name := range list {
			n.MatchNode(line, name, state)
		}
	}
}
