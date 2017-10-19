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

package netutil

import (
	"encoding/binary"
)

type Ptree struct {
	top          *PtreeNode
	maxKeyBits   int
	maxKeyOctets int
	reverse      bool
}

type PtreeNode struct {
	parent    *PtreeNode
	left      *PtreeNode
	right     *PtreeNode
	Item      interface{}
	key       []byte
	keyLength int
	refcnt    uint32
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func bitToOctet(maxKeyBits int) int {
	return max((maxKeyBits+7)/8, 1)
}

func NewPtree(maxKeyBits int) *Ptree {
	return &Ptree{
		maxKeyBits:   maxKeyBits,
		maxKeyOctets: bitToOctet(maxKeyBits),
		//maskBits:     []byte(maskBits),
	}
}

func NewPtreeNode(key []byte, keyLength int) *PtreeNode {
	node := &PtreeNode{keyLength: keyLength}
	node.key = make([]byte, len(key), len(key))
	copy(node.key, key)
	return node
}

func (this *Ptree) ReverseOrderSet() {
	this.reverse = true
}

func (n *PtreeNode) Key() []byte {
	return n.key
}

func (n *PtreeNode) KeyLength() int {
	return n.keyLength
}

func (this *Ptree) keyMatch(key1 []byte, key1Length int, key2 []byte, key2Length int) bool {
	if key1Length > key2Length {
		return false
	}

	offset := min(key1Length, key2Length) / 8
	shift := min(key1Length, key2Length) % 8

	if shift != 0 {
		if (MaskBits[shift] & (key1[offset] ^ key2[offset])) != 0 {
			return false
		}
	}

	for offset > 0 {
		offset--
		if key1[offset] != key2[offset] {
			return false
		}
	}

	return true
}

func (node *PtreeNode) Refcnt() uint32 {
	return node.refcnt
}

func addReference(node *PtreeNode) *PtreeNode {
	node.refcnt++
	return node
}

func (this *Ptree) nodeRemove(node *PtreeNode) {
	if node.left != nil && node.right != nil {
		return
	}

	var child *PtreeNode
	if node.left != nil {
		child = node.left
	} else {
		child = node.right
	}

	parent := node.parent

	if child != nil {
		child.parent = parent
	}

	if parent != nil {
		if parent.left == node {
			parent.left = child
		} else {
			parent.right = child
		}
	} else {
		this.top = child
	}

	if parent != nil && parent.refcnt == 0 {
		this.nodeRemove(parent)
	}
}

func (this *Ptree) delReference(node *PtreeNode) {
	if node.refcnt > 0 {
		node.refcnt--
	}

	if node.refcnt == 0 {
		this.nodeRemove(node)
	}
}

func bitCheck(key []byte, keyLength int, reverse bool) bool {
	offset := keyLength / 8
	shift := 7 - (keyLength % 8)

	bit := ((key[offset] >> uint(shift)) & 1) == 1

	if reverse {
		return !bit
	}

	return bit
}

func (this *Ptree) nodeCommon(node *PtreeNode, key []byte, keyLength int) *PtreeNode {
	var i int

	limit := min(node.keyLength, keyLength) / 8

	for i = 0; i < limit; i++ {
		if node.key[i] != key[i] {
			break
		}
	}

	commonLength := i * 8

	boundary := false

	if commonLength != keyLength {
		diff := node.key[i] ^ key[i]
		mask := byte(0x80)

		for commonLength < keyLength && ((mask & diff) == 0) {
			mask >>= 1
			commonLength++
			boundary = true
		}
	}

	commonKey := make([]byte, bitToOctet(commonLength))

	for j := 0; j < i; j++ {
		commonKey[j] = node.key[j]
	}

	if boundary {
		commonKey[i] = (node.key[i] & MaskBits[commonLength%8])
	}

	return NewPtreeNode(commonKey, commonLength)
}

func (this *Ptree) nodeLink(node *PtreeNode, add *PtreeNode) {
	bit := bitCheck(add.key, node.keyLength, this.reverse)

	if bit {
		node.right = add
	} else {
		node.left = add
	}

	add.parent = node
}

func (this *Ptree) Acquire(key []byte, keyLength int) *PtreeNode {
	var match *PtreeNode
	var add *PtreeNode

	if keyLength > this.maxKeyBits {
		return nil
	}

	node := this.top

	for node != nil && node.keyLength <= keyLength && this.keyMatch(node.key, node.keyLength, key, keyLength) {
		if node.keyLength == keyLength {
			return addReference(node)
		}

		match = node

		if bitCheck(key, node.keyLength, this.reverse) {
			node = node.right
		} else {
			node = node.left
		}
	}

	if node == nil {
		add = NewPtreeNode(key, keyLength)

		if match != nil {
			this.nodeLink(match, add)
		} else {
			this.top = add
		}
	} else {
		add = this.nodeCommon(node, key, keyLength)

		if match != nil {
			this.nodeLink(match, add)
		} else {
			this.top = add
		}

		this.nodeLink(add, node)

		if add.keyLength != keyLength {
			match = add
			add = NewPtreeNode(key, keyLength)
			this.nodeLink(match, add)
		}
	}

	return addReference(add)
}

func (this *Ptree) Lookup(key []byte, keyLength int) *PtreeNode {
	if keyLength > this.maxKeyBits {
		return nil
	}

	node := this.top

	for node != nil && node.keyLength <= keyLength && this.keyMatch(node.key, node.keyLength, key, keyLength) {
		if node.keyLength == keyLength {
			if node.refcnt > 0 {
				return node
			} else {
				return nil
			}
		}

		if bitCheck(key, node.keyLength, this.reverse) {
			node = node.right
		} else {
			node = node.left
		}
	}

	return nil
}

func (this *Ptree) Match(key []byte, keyLength int) *PtreeNode {
	if keyLength > this.maxKeyBits {
		return nil
	}

	node := this.top
	var matched *PtreeNode

	for node != nil && node.keyLength <= keyLength && this.keyMatch(node.key, node.keyLength, key, keyLength) {
		if node.refcnt > 0 {
			matched = node
		}
		if node.keyLength == keyLength {
			break
		}
		if bitCheck(key, node.keyLength, this.reverse) {
			node = node.right
		} else {
			node = node.left
		}
	}
	return matched
}

func (this *Ptree) MatchIPv4(key []byte) *PtreeNode {
	return this.Match(key, 32)
}

func (this *Ptree) MatchIPv6(key []byte) *PtreeNode {
	return this.Match(key, 128)
}

func (this *Ptree) Release(node *PtreeNode) {
	this.delReference(node)
}

func (this *Ptree) LookupByMaxBits(key []byte) *PtreeNode {
	return this.Lookup(key, this.maxKeyBits)
}

func (this *Ptree) LookupByIPv4(key []byte) *PtreeNode {
	return this.Lookup(key, 32)
}

func (this *Ptree) LookupByIPv6(key []byte) *PtreeNode {
	return this.Lookup(key, 128)
}

func (this *Ptree) LookupByUint32(keyInt uint32) *PtreeNode {
	key := make([]byte, 4)
	binary.BigEndian.PutUint32(key, keyInt)
	return this.Lookup(key, 32)
}

func (this *Ptree) AcquireByMaxBits(key []byte) *PtreeNode {
	return this.Acquire(key, this.maxKeyBits)
}

func (this *Ptree) AcquireByIPv4(key []byte) *PtreeNode {
	return this.Acquire(key, 32)
}

func (this *Ptree) AcquireByIPv6(key []byte) *PtreeNode {
	return this.Acquire(key, 128)
}

func (this *Ptree) AcquireByUint32(keyInt uint32) *PtreeNode {
	key := make([]byte, 4)
	binary.BigEndian.PutUint32(key, keyInt)
	return this.Acquire(key, 32)
}

func (this *Ptree) AcquireWithItem(key []byte, keyLength int, v interface{}) *PtreeNode {
	node := this.Acquire(key, keyLength)
	if node.Item != nil {
		this.Release(node)
		return nil
	}
	node.Item = v
	return node
}

func (this *Ptree) ReleaseByUint32(keyInt uint32) {
	key := make([]byte, 4)
	binary.BigEndian.PutUint32(key, keyInt)

	node := this.Lookup(key, 32)
	if node == nil {
		return
	}
	node.Item = nil
	this.Release(node)
}

func (this *Ptree) ReleaseWithItem(key []byte, keyLength int, v interface{}) {
	node := this.Lookup(key, keyLength)
	if node == nil {
		return
	}
	if node.Item != v {
		return
	}
	node.Item = nil
	this.Release(node)
}

func (this *Ptree) Top() *PtreeNode {
	node := this.top

	if node == nil {
		return nil
	} else {
		addReference(node)

		if node.Item == nil {
			return this.Next(node)
		} else {
			return node
		}
	}
}

func (this *Ptree) LookupTop() *PtreeNode {
	node := this.Top()
	if node == nil {
		return nil
	}
	this.delReference(node)
	return node
}

func (this *Ptree) Next(node *PtreeNode) *PtreeNode {
	if node.left != nil {
		next := node.left
		addReference(next)
		this.delReference(node)
		if next.Item == nil {
			return this.Next(next)
		}
		return next
	}
	if node.right != nil {
		next := node.right
		addReference(next)
		this.delReference(node)
		if next.Item == nil {
			return this.Next(next)
		}
		return next
	}

	start := node

	for node.parent != nil {
		if node.parent.left == node && node.parent.right != nil {
			next := node.parent.right
			addReference(next)
			this.delReference(start)
			if next.Item == nil {
				return this.Next(next)
			}
			return next
		}
		node = node.parent
	}
	this.delReference(start)

	return nil
}
