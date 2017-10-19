// Copyright 2016 CoreSwitch
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
	"strconv"
	"strings"
)

type Param struct {
	Command       []string
	Helps         []string
	HelpIndex     int
	Privilege     uint32
	Complete      bool
	Dynamic       bool
	TrailingSpace bool
	Args          []interface{}
	Hook          Hook
	Paren         bool
}

func (ns NodeSlice) setCallback(fn Callback) {
	for _, n := range ns {
		n.Fn = fn
	}
}

func (ns NodeSlice) setHook(hook Hook) {
	for _, n := range ns {
		n.Hook = hook
	}
}

func (ns NodeSlice) lookup(typ NodeType, name string) *Node {
	for _, n := range ns {
		for _, m := range *n.Nodes {
			if m.Type == typ && m.Name == name {
				return m
			}
		}
	}
	return nil
}

func (ns NodeSlice) add(node *Node) {
	for _, n := range ns {
		*n.Nodes = append(*n.Nodes, node)
	}
}

func NewNodeType(typ NodeType, lit string, paren bool) *Node {
	node := NewNode()
	node.Type = typ
	node.Name = lit
	node.Paren = paren

	if typ == NodeRange {
		p := strings.IndexByte(lit, '-')
		if p < 0 || (p+2) >= len(lit) {
			return node
		}
		min_str := lit[1:p]
		max_str := lit[p+1 : len(lit)-1]
		node.Min, _ = strconv.ParseUint(min_str, 10, 64)
		node.Max, _ = strconv.ParseUint(max_str, 10, 64)
	}

	if typ == NodeDynamic {
		lit = lit[1:]

		dynamic := strings.Split(lit, ":")
		if len(dynamic) == 0 {
			return node
		}
		if len(dynamic) == 1 {
			node.Module = dynamic[0]
		}
		if len(dynamic) > 1 {
			node.Module = dynamic[0]
			dynamic = dynamic[1:]
			node.Dynamic = dynamic
		}
	}

	return node
}

func Build(s Scanner, fn Callback, parent NodeSlice, head *NodeSlice, tail *NodeSlice, param *Param) NodeType {
	headline := true
	for {
		typ, lit := s.Scan()
		switch typ {
		case NodeEOL:
			if param.Hook != nil {
				parent.setHook(param.Hook)
			} else {
				parent.setCallback(fn)
			}
			return NodeEOL
		case nodeWhiteSpace:
			// Ignore white space.
		case nodeParenOpen, nodeBraceOpen, nodeCBraceOpen:
			head = &NodeSlice{}
			tail = &NodeSlice{}
			param.Paren = true
			for {
				p := make(NodeSlice, len(parent), len(parent))
				copy(p, parent)
				token := Build(s, fn, p, head, tail, param)
				if token == nodeParenClose || token == nodeBraceClose || token == nodeCBraceClose {
					break
				}
			}
			param.Paren = false
			parent = make(NodeSlice, len(*tail), len(*tail))
			copy(parent, *tail)
		case nodeParenClose, nodeBraceClose, nodeCBraceClose, nodeSeparator:
			if tail != nil {
				for _, n := range parent {
					*tail = append(*tail, n)
				}
			}
			return typ
		case nodeAmpersand:
			for _, n := range parent {
				parent.add(n)
			}
		default:
			if typ == NodeDynamic && !param.Dynamic {
				typ = NodeWord
			}
			node := parent.lookup(typ, lit)
			if node == nil {
				node = NewNodeType(typ, lit, param.Paren)
				parent.add(node)
			}
			if param.HelpIndex < len(param.Helps) && len(param.Helps[param.HelpIndex]) > 0 {
				node.Help = param.Helps[param.HelpIndex]
			}
			if headline && head != nil {
				*head = append(*head, node)
				headline = false
			}
			param.HelpIndex++
			parent = NodeSlice{node}
		}
	}
}

func (p *Node) Install(s Scanner, fn Callback, args []*Param) {
	param := Param{}
	if len(args) > 0 {
		param = *args[0]
	}
	Build(s, fn, NodeSlice{p}, nil, nil, &param)
}

func (p *Node) InstallCmd(cmd []string, fn Callback, args ...*Param) {
	s := NewScannerCmd(cmd)
	p.Install(s, fn, args)
}

func (p *Node) InstallLine(line string, fn Callback, args ...*Param) {
	s := NewScannerLine(line)
	p.Install(s, fn, args)
}
