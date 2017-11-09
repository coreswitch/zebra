// Copyright 2016, 2017 OpenConfigd Project.
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

package config

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type item struct {
	typ itemType
	pos int
	val string
}

type itemType int

const (
	itemError itemType = iota
	itemEOF
	itemString
	itemIdentifier
	itemBlockComment
	itemLineComment
	itemWhiteSpace
	itemLeftBrace
	itemRightBrace
	itemLeftBracket
	itemRightBracket
	itemSemiColon
)

func (t itemType) String() string {
	switch t {
	case itemError:
		return "error"
	case itemEOF:
		return "EOF"
	case itemString:
		return "string"
	case itemIdentifier:
		return "identifier"
	case itemBlockComment:
		return "blockComment"
	case itemLineComment:
		return "lineComment"
	case itemWhiteSpace:
		return "whiteSpace"
	case itemLeftBrace:
		return "leftBrace"
	case itemRightBrace:
		return "rightBrace"
	case itemLeftBracket:
		return "leftBracket"
	case itemRightBracket:
		return "rightBracket"
	case itemSemiColon:
		return "semiColon"
	default:
		return "unknown"
	}
}

const eof = -1

type stateFn func(*lexer) stateFn

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	}
	return fmt.Sprintf("%s:%q", i.typ, i.val)
}

type lexer struct {
	input  *bufio.Reader
	buffer bytes.Buffer
	state  stateFn
	pos    int
	start  int
	items  chan item
}

func lex(input io.Reader) *lexer {
	l := &lexer{
		input: bufio.NewReader(input),
		items: make(chan item),
	}
	go l.run()
	return l
}

func (l *lexer) nextItem() item {
	item := <-l.items
	return item
}

func (l *lexer) run() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
}

func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{itemError, l.start, fmt.Sprintf(format, args...)}
	return nil
}

func (l *lexer) next() rune {
	r, w, err := l.input.ReadRune()
	if err == io.EOF {
		return eof
	}
	l.pos += w
	l.buffer.WriteRune(r)
	return r
}

func (l *lexer) emit(t itemType) {
	l.items <- item{t, l.start, l.buffer.String()}
	l.start = l.pos
	l.buffer.Truncate(0)
}

func (l *lexer) peek() rune {
	lead, err := l.input.Peek(1)
	if err == io.EOF {
		return eof
	} else if err != nil {
		l.errorf("%s", err.Error())
		return 0
	}

	p, err := l.input.Peek(runeLen(lead[0]))
	if err == io.EOF {
		return eof
	} else if err != nil {
		l.errorf("%s", err.Error())
		return 0
	}
	r, _ := utf8.DecodeRune(p)
	return r
}

func lexText(l *lexer) stateFn {
Loop:
	for {
		r := l.peek()
		switch r {
		case '"':
			return lexString
		case '{':
			l.next()
			l.emit(itemLeftBrace)
		case '}':
			l.next()
			l.emit(itemRightBrace)
		case ';':
			l.next()
			l.emit(itemSemiColon)
		case '#':
			return lexCommentLine
		default:
			if unicode.IsSpace(r) {
				return lexWhiteSpace
			} else if r == eof {
				l.next()
				break Loop
			} else {
				return lexIdentifier
			}
		}
	}
	l.emit(itemEOF)
	return nil
}

func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.peek()) >= 0 {
		l.next()
		return true
	}
	return false
}

const (
	hexdigit  = "0123456789ABDEFabcdef"
	digit     = "0123456789"
	digit1To9 = "123456789"
)

func lexString(l *lexer) stateFn {
	l.next()
	for {
		switch r := l.next(); {
		case r == '"':
			l.emit(itemString)
			return lexText
		case r == '\\':
			if l.accept(`"\/bfnrt`) {
				break
			} else if r := l.next(); r == 'u' {
				for i := 0; i < 4; i++ {
					if !l.accept(hexdigit) {
						return l.errorf("expected 4 hexadecimal digits")
					}
				}
			} else {
				// Accept non unicode escape character for DHCP lease file parse.
			}
		case unicode.IsControl(r):
			return l.errorf("cannot contain control characters in strings")
		case r == eof:
			return l.errorf("unclosed string")
		}
	}
}

func isIdentifierPart(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' || r == '/' || r == '-' || r == ':' || r == '_'
}

func lexIdentifier(l *lexer) stateFn {
	r := l.peek()
	if !isIdentifierPart(r) {
		return l.errorf("identifier expected")
	}

	for r = l.peek(); isIdentifierPart(r); {
		l.next()
		r = l.peek()
	}

	l.emit(itemIdentifier)

	return lexText
}

func lexWhiteSpace(l *lexer) stateFn {
	for unicode.IsSpace(l.peek()) {
		l.next()
	}
	l.emit(itemWhiteSpace)
	return lexText
}

func lexCommentLine(l *lexer) stateFn {
	for r := l.peek(); r != '\n'; {
		l.next()
		r = l.peek()
	}
	l.emit(itemWhiteSpace)
	return lexText
}

func runeLen(lead byte) int {
	if lead < 0xC0 {
		return 1
	} else if lead < 0xE0 {
		return 2
	} else if lead < 0xF0 {
		return 3
	} else {
		return 4
	}
}

type stack [][]string

func pop(s stack) stack {
	if len(s) > 0 {
		return s[:len(s)-1]
	}
	return s
}

func flatten(stack stack) []string {
	ret := []string{}
	for _, a := range stack {
		ret = append(ret, a...)
	}
	return ret
}

func LoadConfig(reader io.Reader) error {
	s := stack{}
	a := []string{}
	l := lex(reader)

	for {
		switch item := l.nextItem(); item.typ {
		case itemEOF:
			return nil
		case itemIdentifier:
			a = append(a, item.val)
		case itemString:
			val, err := strconv.Unquote(item.val)
			if err == nil {
				item.val = val
			}
			a = append(a, item.val)
		case itemLeftBrace:
			s = append(s, a)
			a = make([]string, 0)
		case itemRightBrace:
			s = pop(s)
		case itemSemiColon:
			s = append(s, a)
			a = make([]string, 0)
			Parse(flatten(s), rootEntry, configCandidate, nil)
			s = pop(s)
		case itemError:
			return fmt.Errorf("Parse error")
		default:
		}
	}
}

func Load(path string) error {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	reader := bytes.NewBufferString(string(input))
	err = LoadConfig(reader)
	if err != nil {
		return err
	}
	return nil
}
