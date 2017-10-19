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

	"github.com/coreswitch/netutil"
)

const (
	ParseNoMatch = iota
	ParseAmbiguous
	ParseIncomplete
	ParseSuccess
)

func ParseResult2String(result int) string {
	switch result {
	case ParseNoMatch:
		return "ParseNoMatch"
	case ParseAmbiguous:
		return "ParseAmbiguous"
	case ParseIncomplete:
		return "ParseIncomplete"
	case ParseSuccess:
		return "ParseSuccess"
	default:
		return "Unknown"
	}
}

func NewParser() *Node {
	return NewNode()
}

func String2Interface(strs []string) []interface{} {
	dest := []interface{}{}
	for _, str := range strs {
		dest = append(dest, str)
	}
	return dest
}

func Interface2String(srcs []interface{}) []string {
	dest := []string{}
	for _, src := range srcs {
		if ip, ok := src.(net.IP); ok {
			dest = append(dest, ip.String())
		} else if p, ok := src.(*netutil.Prefix); ok {
			dest = append(dest, p.String())
		} else {
			dest = append(dest, src.(string))
		}
	}
	return dest
}

func (n *Node) ParseArgSet(str string, param *Param) {
	switch n.Type {
	case NodeKeyword:
		if n.Paren {
			param.Args = append(param.Args, n.Name)
		}
	case NodeWord, NodeDynamic, NodeLine:
		param.Args = append(param.Args, str)
	case NodeRange:
		u, _ := strconv.ParseUint(str, 10, 64)
		param.Args = append(param.Args, u)
	case NodeIPv4:
		param.Args = append(param.Args, netutil.ParseIPv4(str))
	case NodeIPv4Prefix:
		prefix, _ := netutil.ParsePrefix(str)
		param.Args = append(param.Args, prefix)
	case NodeIPv6:
		param.Args = append(param.Args, net.ParseIP(str))
	case NodeIPv6Prefix:
		prefix, _ := netutil.ParsePrefix(str)
		param.Args = append(param.Args, prefix)
	}
}

func (n *Node) ParseMatch(line string, param *Param, state *MatchState) {
	for _, node := range *n.Nodes {
		if param.Privilege > 0 && node.Privilege < param.Privilege {
			continue
		}
		if node.Type == NodeDynamic {
			node.MatchDynamic(line, param.Command, state)
		} else {
			node.MatchNode(line, "", state)
		}
	}
}

func (n *Node) Parse(line string, param *Param) (int, Callback, []interface{}, CompSlice) {
	state := &MatchState{complete: param.Complete}

	n.ParseMatch(line, param, state)

	if state.count == 0 {
		return ParseNoMatch, nil, nil, nil
	}
	if state.count > 1 {
		return ParseAmbiguous, nil, nil, state.comps
	}

	if !param.Complete {
		state.node.ParseArgSet(line[0:state.pos], param)
	}

	// Skip trailing white space.
	pos := state.pos
	match := state.match
	node := state.node
	for ; state.pos < len(line); state.pos++ {
		if !isWhiteSpace(line, state.pos) {
			break
		}
	}
	if pos != state.pos {
		pos = state.pos
		if state.node.Hook != nil {
			if hook, ok := state.node.Hook.(func(*Param) (int, Callback, []interface{}, CompSlice)); ok {
				return hook(param)
			}
		}
		if param.Complete {
			state.comps = state.comps[:0]
			state.node.ParseMatch(line[pos:], param, state)
		}
	}
	line = line[pos:]

	if len(line) == 0 {
		if node.Fn == nil || match == MatchTypeIncomplete {
			return ParseIncomplete, nil, nil, state.comps
		} else {
			return ParseSuccess, node.Fn, param.Args, state.comps
		}
	}
	return node.Parse(line, param)
}

func (n *Node) ParseCmd(cmd []string, args ...*Param) (int, Callback, []interface{}, CompSlice) {
	var param *Param
	if len(args) > 0 {
		param = args[0]
	} else {
		param = &Param{}
	}
	param.Command = cmd
	return n.Parse(strings.Join(cmd, " "), param)
}

func (n *Node) ParseLine(line string, args ...*Param) (int, Callback, []interface{}, CompSlice) {
	var param *Param
	if len(args) > 0 {
		param = args[0]
	} else {
		param = &Param{}
	}
	return n.Parse(line, param)
}
