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

func isWhiteSpace(str string, i int) bool {
	return str[i] == ' ' || str[i] == '\n' || str[i] == '\t'
}

func isDigit(str string, i int) bool {
	return '0' <= str[i] && str[i] <= '9'
}

func isDelimiter(str string, i int) bool {
	return (len(str) == i || isWhiteSpace(str, i))
}

func longestCommonPrefix(a, b string) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			break
		}
	}
	return
}

func isWhiteSpaceChar(c byte) bool {
	return (c == ' ' || c == '\t' || c == '\n')
}

func isUpperChar(c byte) bool {
	return ('A' <= c && c <= 'Z')
}

func isLowerChar(c byte) bool {
	return ('a' <= c && c <= 'z')
}

func isNumberChar(c byte) bool {
	return ('0' <= c && c <= '9')
}

func isAlnumChar(c byte) bool {
	return isUpperChar(c) || isLowerChar(c) || isNumberChar(c)
}

func isDynamicChar(c byte) bool {
	return isAlnumChar(c) || c == '-' || c == ':' || c == '$'
}

func isRangeChar(c byte) bool {
	return isNumberChar(c) || c == '-' || c == '<' || c == '>'
}

func isParameterChar(c byte) bool {
	return isAlnumChar(c) || c == '-' || c == ':' || c == '.' || c == '/'
}

func isKeywordChar(c byte) bool {
	return isAlnumChar(c) || c == '-'
}
