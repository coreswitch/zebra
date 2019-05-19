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

package policy

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"strings"
)

const (
	AS_SET             = 1
	AS_SEQUENCE        = 2
	AS_CONFED_SEQUENCE = 3
	AS_CONFED_SET      = 4
)

type AsPath []AsSegmentInterface

//////////////////////////////
type ASPath struct {
	segs []*As4Segment
}

func NewASPath() *ASPath {
	return &ASPath{}
}

type DecodeOption func() int

func With2Octet() DecodeOption {
	return func() int {
		return 2
	}
}

func (aspath *ASPath) DecodeFromBytes(data []byte, opts ...DecodeOption) error {
	aslen := 4
	for _, opt := range opts {
		aslen = opt()
	}

	for len(data) > 0 {
		seg := &As4Segment{}

		if len(data) < 2 {
			return fmt.Errorf("Invalid segment length")
		}

		seg.Type = data[0]
		if seg.Type < AS_SET || seg.Type > AS_CONFED_SET {
			return fmt.Errorf("Invalid segment type")
		}
		seg.Length = data[1]
		if len(data) < int(seg.Length)*aslen {
			return fmt.Errorf("Invalid segment length")
		}
		data = data[2:]

		for i := 0; i < int(seg.Length); i++ {
			switch aslen {
			case 4:
				seg.As = append(seg.As, binary.BigEndian.Uint32(data))
				data = data[4:]
			case 2:
				seg.As = append(seg.As, uint32(binary.BigEndian.Uint16(data)))
				data = data[2:]
			}
		}
		aspath.segs = append(aspath.segs, seg)
	}

	return nil
}

func (aspath *ASPath) String() string {
	strs := []string{}
	for _, seg := range aspath.segs {
		strs = append(strs, seg.String())
	}
	return strings.Join(strs, " ")
}

func (aspath *ASPath) Replace(from, to uint32) *ASPath {
	for _, seg := range aspath.segs {
		seg.Replace(from, to)
	}
	return aspath
}

//////////////////////////////

func (aspath AsPath) String() string {
	strs := []string{}
	for _, asseg := range aspath {
		strs = append(strs, asseg.String())
	}
	return strings.Join(strs, " ")
}

func (aspath AsPath) PathLength() int {
	len := 0
	for _, asseg := range aspath {
		len += asseg.PathLength()
	}
	return len
}

type AsSegmentInterface interface {
	String() string
	Serialize() ([]byte, error)
	DecodeFromBytes([]byte) error
	EncodeLength() int
	PathLength() int
	Append(uint32)
	GetLength() int
	GetType() uint8
}

type As4Segment struct {
	Type   uint8
	Length uint8
	As     []uint32
}

func AsSegmentDelimiter(typ uint8) (string, string) {
	switch typ {
	case AS_SEQUENCE:
		return "", ""
	case AS_SET:
		return "{", "}"
	case AS_CONFED_SEQUENCE:
		return "(", ")"
	case AS_CONFED_SET:
		return "[", "]"
	default:
		return "", ""
	}
}

func (seg *As4Segment) String() string {
	head, tail := AsSegmentDelimiter(seg.Type)
	strs := []string{}
	for _, as := range seg.As {
		strs = append(strs, strconv.FormatUint(uint64(as), 10))
	}
	return head + strings.Join(strs, " ") + tail
}

func (aspath *ASPath) Serialize() ([]byte, error) {
	var data []byte
	for _, seg := range aspath.segs {
		buf, err := seg.Serialize()
		if err != nil {
			return nil, err
		}
		data = append(data, buf...)
	}
	return data, nil
}

func (seg *As4Segment) GetLength() int {
	return int(seg.Length)
}

func (seg *As4Segment) GetType() uint8 {
	return seg.Type
}

func (seg *As4Segment) DecodeFromBytes(data []byte) error {
	// code := BGP_ERR_UPDATE_MESSAGE_ERROR
	// subCode := BGP_ERR_SUB_MALFORMED_AS_PATH
	if len(data) < 2 {
		// return NewBgpError(code, subCode, nil, "")
		return fmt.Errorf("")
	}
	seg.Type = data[0]
	if seg.Type < AS_SET || seg.Type > AS_CONFED_SET {
		// return NewBgpError(code, subCode, nil, "")
		return fmt.Errorf("")
	}
	seg.Length = data[1]
	data = data[2:]
	if len(data) < int(seg.Length*4) {
		// return NewBgpError(code, subCode, nil, "")
		return fmt.Errorf("")
	}
	for i := 0; i < int(seg.Length); i++ {
		seg.As = append(seg.As, binary.BigEndian.Uint32(data))
		data = data[4:]
	}
	return nil
}

func (seg *As4Segment) Serialize() ([]byte, error) {
	buf := make([]byte, 2+len(seg.As)*4)
	buf[0] = seg.Type
	buf[1] = seg.Length
	for pos, as := range seg.As {
		binary.BigEndian.PutUint32(buf[2+pos*4:], as)
	}
	return buf, nil
}

func (seg *As4Segment) EncodeLength() int {
	return int(2 + (seg.Length * 4))
}

func (seg *As4Segment) PathLength() int {
	switch seg.Type {
	case AS_SET, AS_CONFED_SET:
		return int(seg.Length)
	case AS_SEQUENCE, AS_CONFED_SEQUENCE:
		return 1
	default:
		return 0
	}
}

func (seg *As4Segment) Replace(from, to uint32) {
	for pos, as := range seg.As {
		if as == from {
			seg.As[pos] = to
		}
	}
}

type As2Segment struct {
	Type   uint8
	Length uint8
	As     []uint16
}

func (seg *As2Segment) String() string {
	head, tail := AsSegmentDelimiter(seg.Type)
	strs := []string{}
	for _, as := range seg.As {
		strs = append(strs, strconv.FormatUint(uint64(as), 10))
	}
	return head + strings.Join(strs, " ") + tail
}

func (seg *As2Segment) Serialize() ([]byte, error) {
	return nil, nil
}

func (seg *As2Segment) DecodeFromBytes(data []byte) error {
	return nil
}

func (seg *As2Segment) EncodeLength() int {
	return int(2 + (seg.Length * 2))
}

func (seg *As2Segment) PathLength() int {
	switch seg.Type {
	case AS_SET, AS_CONFED_SET:
		return int(seg.Length)
	case AS_SEQUENCE, AS_CONFED_SEQUENCE:
		return 1
	default:
		return 0
	}
}

const (
	asTokenError = iota
	asTokenEOF
	asTokenNumber
	asTokenSetStart
	asTokenSetEnd
	asTokenConfedSequenceStart
	asTokenConfedSequenceEnd
	asTokenConfedSetStart
	asTokenConfedSetEnd
)

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func AsPathTokenGet(str string) (string, int, uint32) {
	// Skip seperators ' ' for sequences ',' for sets.
	for len(str) > 0 && (str[0] == ' ' || str[0] == ',') {
		str = str[1:]
	}
	if len(str) == 0 {
		return "", asTokenEOF, 0
	}
	switch str[0] {
	case '{':
		return str[1:], asTokenSetStart, '{'
	case '}':
		return str[1:], asTokenSetEnd, '}'
	case '(':
		return str[1:], asTokenConfedSequenceStart, '('
	case ')':
		return str[1:], asTokenConfedSequenceEnd, ')'
	case '[':
		return str[1:], asTokenConfedSetStart, '['
	case ']':
		return str[1:], asTokenConfedSetEnd, ']'
	}

	if isDigit(str[0]) {
		asval := uint32(str[0] - '0')
		str = str[1:]

		for len(str) > 0 && isDigit(str[0]) {
			asval *= 10
			asval += uint32(str[0] - '0')
			str = str[1:]
		}

		return str, asTokenNumber, asval
	}

	return str, asTokenError, 0
}

func NewAs4Segment(typ uint8) *As4Segment {
	return &As4Segment{Type: typ}
}

func (seg *As4Segment) Append(asnum uint32) {
	seg.Length++
	seg.As = append(seg.As, asnum)
}

func (aspath AsPath) Append(asnum uint32) AsPath {
	var seg AsSegmentInterface
	if len(aspath) == 0 {
		seg = NewAs4Segment(AS_SEQUENCE)
		aspath = append(aspath, seg)
	} else {
		seg = aspath[len(aspath)-1]
	}
	// Add a new segment if existing segment is full.
	if seg.GetLength() == 255 {
		seg = NewAs4Segment(seg.GetType())
		aspath = append(aspath, seg)
	}
	seg.Append(asnum)

	return aspath
}

func (aspath *ASPath) Append(asnum uint32) *ASPath {
	var seg *As4Segment
	if len(aspath.segs) == 0 {
		seg = NewAs4Segment(AS_SEQUENCE)
		aspath.segs = append(aspath.segs, seg)
	} else {
		seg = aspath.segs[len(aspath.segs)-1]
	}
	// Add a new segment if existing segment is full.
	if seg.GetLength() == 255 {
		seg = NewAs4Segment(seg.GetType())
		aspath.segs = append(aspath.segs, seg)
	}
	seg.Append(asnum)

	return aspath
}

// AsPath Prepend.
func (aspath AsPath) Prepend(path AsPath) AsPath {
	aspath = append(path, aspath...)
	return aspath
}

func AsPathParse(str string) (AsPath, error) {
	aspath := AsPath{}

	token := 0
	segType := uint8(AS_SEQUENCE)
	needSegment := true
	asnum := uint32(0)
	for {
		str, token, asnum = AsPathTokenGet(str)
		switch token {
		case asTokenEOF:
			return aspath, nil
		case asTokenNumber:
			if needSegment {
				aspath = append(aspath, NewAs4Segment(segType))
				needSegment = false
			}
			aspath = aspath.Append(asnum)
		case asTokenSetStart:
			segType = AS_SET
			aspath = append(aspath, NewAs4Segment(AS_SET))
			needSegment = false
		case asTokenConfedSequenceStart:
			segType = AS_CONFED_SEQUENCE
			aspath = append(aspath, NewAs4Segment(AS_CONFED_SEQUENCE))
			needSegment = false
		case asTokenConfedSetStart:
			segType = AS_CONFED_SET
			aspath = append(aspath, NewAs4Segment(AS_CONFED_SET))
			needSegment = false
		case asTokenSetEnd, asTokenConfedSequenceEnd, asTokenConfedSetEnd:
			segType = AS_SEQUENCE
			needSegment = true
		case asTokenError:
			return nil, fmt.Errorf("AS_PATH parse error")
		}
	}

	return aspath, nil
}
