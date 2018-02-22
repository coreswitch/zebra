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
	//"encoding/json"

	"net"
	"strconv"
	"strings"
)

var MaskBits = []byte{0x00, 0x80, 0xc0, 0xe0, 0xf0, 0xf8, 0xfc, 0xfe, 0xff}
var MaskReverseBits = []byte{0xff, 0x7f, 0x3f, 0x1f, 0x0f, 0x07, 0x03, 0x01, 0x00}

const (
	AFI_IP = iota
	AFI_IP6
	AFI_MAX
)

type Prefix struct {
	net.IP
	Length int
}

func CopyIP(ip net.IP) net.IP {
	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}

func (p *Prefix) AFI() int {
	if len(p.IP) == net.IPv4len {
		return AFI_IP
	}
	if len(p.IP) == net.IPv6len {
		return AFI_IP6
	}
	return AFI_MAX
}

func (p *Prefix) ByteLength() int {
	return (p.Length + 7) / 8
}

func NewPrefixAFI(afi int) *Prefix {
	switch afi {
	case AFI_IP:
		return &Prefix{IP: make(net.IP, net.IPv4len), Length: 0}
	case AFI_IP6:
		return &Prefix{IP: make(net.IP, net.IPv6len), Length: 0}
	default:
		return nil
	}
}

func ParsePrefix(s string) (*Prefix, error) {
	i := strings.IndexByte(s, '/')
	if i < 0 {
		return nil, &net.ParseError{Type: "Prefix address", Text: s}
	}

	addr, mask := s[:i], s[i+1:]

	ip := net.ParseIP(addr)
	if ip == nil {
		return nil, &net.ParseError{Type: "Prefix address", Text: s}
	}

	ip4 := ip.To4()
	if ip4 != nil {
		ip = ip4
	}

	length, err := strconv.Atoi(mask)
	if err != nil {
		return nil, err
	}
	return &Prefix{IP: ip, Length: length}, nil
}

func (p *Prefix) String() string {
	return p.IP.String() + "/" + strconv.Itoa(p.Length)
}

func (p *Prefix) MarshalJSON() ([]byte, error) {
	return []byte(`"` + p.String() + `"`), nil
}

func (p *Prefix) ApplyMask() *Prefix {
	i := p.Length / 8

	if i >= len(p.IP) {
		return p
	}

	offset := p.Length % 8
	p.IP[i] &= MaskBits[offset]
	i++

	for i < len(p.IP) {
		p.IP[i] = 0
		i++
	}
	return p
}

func (p *Prefix) ApplyReverseMask() *Prefix {
	i := p.Length / 8

	if i >= len(p.IP) {
		return p
	}

	offset := p.Length % 8
	p.IP[i] |= MaskReverseBits[offset]
	i++

	for i < len(p.IP) {
		p.IP[i] = 0
		i++
	}
	return p
}

func (p *Prefix) Copy() *Prefix {
	return &Prefix{IP: CopyIP(p.IP), Length: p.Length}
}

func (p *Prefix) Equal(x *Prefix) bool {
	if p.IP.Equal(x.IP) && p.Length == x.Length {
		return true
	} else {
		return false
	}
}

func PrefixFromIPNet(net net.IPNet) *Prefix {
	ip := net.IP.To4()
	if ip == nil {
		ip = net.IP
	}
	len, _ := net.Mask.Size()
	return &Prefix{IP: CopyIP(ip), Length: len}
}

func PrefixFromNode(n *PtreeNode) *Prefix {
	ip := make([]byte, len(n.Key()))
	copy(ip, n.Key())
	p := &Prefix{
		IP:     ip,
		Length: n.KeyLength(),
	}
	return p
}

func IPNetFromPrefix(p *Prefix) net.IPNet {
	if len(p.IP) == net.IPv4len {
		return net.IPNet{IP: p.IP, Mask: net.CIDRMask(p.Length, 32)}
	} else {
		return net.IPNet{IP: p.IP, Mask: net.CIDRMask(p.Length, 128)}
	}
}

func PrefixFromIPPrefixlen(ip net.IP, len int) *Prefix {
	return &Prefix{IP: CopyIP(ip), Length: len}
}

func ParseIPv4(s string) net.IP {
	ip := net.ParseIP(s)
	ip4 := ip.To4()
	if ip4 != nil {
		return ip4
	}
	return ip
}

func SameIp(a net.IP, b net.IP) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil && b != nil {
		return false
	}
	if a != nil && b == nil {
		return false
	}
	return a.Equal(b)
}

func (p *Prefix) IsDefault() bool {
	if p.Length != 0 {
		return false
	}
	for _, v := range p.IP {
		if v != 0 {
			return false
		}
	}
	return true
}

func (p *Prefix) Match(q *Prefix) bool {
	if p.Length > q.Length {
		return false
	}
	if len(p.IP) != len(q.IP) {
		return false
	}
	if len(p.IP) != 4 && len(p.IP) != 16 {
		return false
	}

	offset := p.Length / 8
	shift := p.Length % 8

	if shift != 0 {
		if (MaskBits[shift] & (p.IP[offset] ^ q.IP[offset])) != 0 {
			return false
		}
	}
	for offset != 0 {
		offset--
		if p.IP[offset] != q.IP[offset] {
			return false
		}
	}
	return true
}
