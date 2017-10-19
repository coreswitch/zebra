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
	"bytes"
	"strings"
)

// ScannerCmd start.
type ScannerCmd struct {
	cmd   []string
	index int
}

func (s *ScannerCmd) Scan() (typ NodeType, lit string) {
	if s.index >= len(s.cmd) {
		return NodeEOL, ""
	}

	lit = s.cmd[s.index]
	s.index++

	if len(lit) > 0 {
		switch lit[0] {
		case '<':
			typ = NodeRange
			return
		case '&':
			typ = nodeAmpersand
			return
		}
	}

	typ, ok := string2NodeType[lit]
	if !ok {
		typ = NodeKeyword
	}
	return
}

func NewScannerCmd(cmd []string) Scanner {
	return &ScannerCmd{cmd: cmd, index: 0}
}

// ScannerCmd ends here.

type Scanner interface {
	Scan() (NodeType, string)
}

type ScannerLine struct {
	r *strings.Reader
}

func NewScannerLine(line string) Scanner {
	return &ScannerLine{strings.NewReader(line)}
}

var eol = byte(0)

func (s *ScannerLine) read() byte {
	c, err := s.r.ReadByte()
	if err != nil {
		return eol
	}
	return c
}

func (s *ScannerLine) scanDynamic() (typ NodeType, lit string) {
	var buf bytes.Buffer
	buf.WriteByte(s.read())

	for {
		c := s.read()

		if c == eol {
			break
		}
		if !isDynamicChar(c) {
			s.r.UnreadByte()
			break
		}
		buf.WriteByte(c)
	}
	return NodeDynamic, buf.String()
}

func (s *ScannerLine) scanRange() (typ NodeType, lit string) {
	var buf bytes.Buffer
	buf.WriteByte(s.read())

	for {
		c := s.read()

		if c == eol {
			break
		}
		if !isRangeChar(c) {
			s.r.UnreadByte()
			break
		}
		buf.WriteByte(c)
	}
	return NodeRange, buf.String()
}

func (s *ScannerLine) scanParameter() (typ NodeType, lit string) {
	var buf bytes.Buffer
	buf.WriteByte(s.read())

	for {
		c := s.read()

		if c == eol {
			break
		}
		if !isParameterChar(c) {
			s.r.UnreadByte()
			break
		}
		buf.WriteByte(c)
	}

	parameter := buf.String()

	typ, ok := string2NodeType[parameter]
	if !ok {
		return NodeKeyword, parameter
	}
	return typ, parameter
}

func (s *ScannerLine) scanKeyword() (typ NodeType, lit string) {
	var buf bytes.Buffer
	buf.WriteByte(s.read())

	for {
		c := s.read()

		if c == eol {
			break
		}
		if !isKeywordChar(c) {
			s.r.UnreadByte()
			break
		}
		buf.WriteByte(c)
	}

	return NodeKeyword, buf.String()
}

func (s *ScannerLine) Scan() (typ NodeType, lit string) {
	c := s.read()

	switch {
	case c == eol:
		return NodeEOL, lit
	case isWhiteSpaceChar(c):
		return nodeWhiteSpace, lit
	case c == '(':
		return nodeParenOpen, lit
	case c == ')':
		return nodeParenClose, lit
	case c == '{':
		return nodeCBraceOpen, lit
	case c == '}':
		return nodeCBraceClose, lit
	case c == '[':
		return nodeBraceOpen, lit
	case c == ']':
		return nodeBraceClose, lit
	case c == '|':
		return nodeSeparator, lit
	case c == '&':
		return nodeAmpersand, lit
	case c == ':':
		s.r.UnreadByte()
		return s.scanDynamic()
	case c == '<':
		s.r.UnreadByte()
		return s.scanRange()
	case isUpperChar(c):
		s.r.UnreadByte()
		return s.scanParameter()
	case isLowerChar(c):
		s.r.UnreadByte()
		return s.scanKeyword()
	}
	return nodeUnknown, lit
}
