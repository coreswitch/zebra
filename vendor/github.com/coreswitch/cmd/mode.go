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

type Mode struct {
	Name     string
	TopLevel bool
	Prompt   string
	Header   string
	Parent   *Mode
	Modes    []*Mode
	Parser   *Node
}

func NewMode(name string, header string, prompt string) *Mode {
	return &Mode{Name: name, Header: header, Prompt: prompt, Parser: NewParser()}
}

func (m *Mode) InstallLine(line string, fn Callback, args ...*Param) {
	s := NewScannerLine(line)
	m.Parser.Install(s, fn, args)
}

func (m *Mode) InstallHook(line string, hook Hook, args ...*Param) {
	param := Param{}
	if len(args) > 0 {
		param = *args[0]
	}
	param.Hook = hook
	m.Parser.InstallLine(line, nil, &param)
}
