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
	"fmt"
)

type Node struct {
	Name      string
	Help      string
	Type      NodeType
	Privilege uint32
	Fn        Callback
	Hook      Hook
	Nodes     *NodeSlice
	Min       uint64
	Max       uint64
	Module    string
	Dynamic   []string
	Paren     bool
}

type Args []interface{}

type Callback interface{}

type Hook interface{}

type NodeSlice []*Node

type NodeType int

const (
	NodeEOL NodeType = iota
	NodeKeyword
	NodeIPv4
	NodeIPv4Prefix
	NodeIPv6
	NodeIPv6Prefix
	NodeWord
	NodeLine
	NodeRange
	NodeDynamic
	nodeWhiteSpace
	nodeParenOpen
	nodeParenClose
	nodeBraceOpen
	nodeBraceClose
	nodeCBraceOpen
	nodeCBraceClose
	nodeSeparator
	nodeAmpersand
	nodeUnknown
)

var nodeType2StirngMap = map[NodeType]string{
	NodeKeyword:    "Keyword",
	NodeIPv4:       "A.B.C.D",
	NodeIPv4Prefix: "A.B.C.D/M",
	NodeIPv6:       "X:X::X:X",
	NodeIPv6Prefix: "X:X::X:X/M",
	NodeWord:       "WORD",
	NodeLine:       "LINE",
	NodeDynamic:    "Dynamic",
}

var string2NodeType = map[string]NodeType{
	"A.B.C.D":    NodeIPv4,
	"A.B.C.D/M":  NodeIPv4Prefix,
	"X:X::X:X":   NodeIPv6,
	"X:X::X:X/M": NodeIPv6Prefix,
	"WORD":       NodeWord,
	"LINE":       NodeLine,
}

var String2NodeTypeMap = map[string]NodeType{}

func NewNode() *Node {
	return &Node{Nodes: &NodeSlice{}}
}

func (n *Node) Lookup(name string) *Node {
	for _, node := range *n.Nodes {
		if node.Name == name {
			return node
		}
	}
	return nil
}

func (n *Node) LinkNodes(node *Node) {
	n.Nodes = node.Nodes
}

func (n *Node) String() string {
	switch n.Type {
	case NodeKeyword, NodeDynamic:
		return n.Name
	case NodeRange:
		return "<>"
	default:
		return nodeType2StirngMap[n.Type]
	}
}

func (n *Node) DumpNode(depth int) {
	if depth > 0 {
		fmt.Printf("%*s", depth*4, " ")
	}
	fmt.Printf("%s", n.String())
	fmt.Printf("[%s]", nodeType2StirngMap[n.Type])
	if n.Fn != nil {
		fmt.Printf(" (*)")
	}
	fmt.Printf("\n")
	for _, node := range *n.Nodes {
		node.DumpNode(depth + 1)
	}
}

func (p *Node) Dump() {
	for _, node := range *p.Nodes {
		node.DumpNode(0)
	}
}
