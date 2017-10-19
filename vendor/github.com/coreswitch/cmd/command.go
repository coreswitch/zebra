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

const (
	Set int = iota
	Delete
	SetValidate
	DeleteValidate
	Subscribe
)

const (
	Success int = iota
)

type Cmd struct {
	Modes map[string]*Mode
}

func NewCmd() *Cmd {
	return &Cmd{Modes: map[string]*Mode{}}
}

func (c *Cmd) LookupMode(name string) *Mode {
	return c.Modes[name]
}

func (c *Cmd) InstallMode(name string, header string, prompt string) *Mode {
	mode := NewMode(name, header, prompt)
	mode.TopLevel = true
	c.Modes[name] = mode
	return mode
}

func (c *Cmd) InstallSubMode(mode *Mode, name string, header string, prompt string) *Mode {
	m := NewMode(name, header, prompt)
	m.Parent = mode
	c.Modes[name] = m
	mode.Modes = append(mode.Modes, m)
	return mode
}

func (c *Cmd) FirstCommands(mode string, privilege uint32) (cmds string) {
	m := c.LookupMode(mode)
	if m == nil {
		return
	}
	for _, node := range *m.Parser.Nodes {
		if privilege < node.Privilege {
			continue
		}
		cmds += node.Name
		cmds += "\n"
	}
	return
}

func (c *Cmd) ParseLine(mode string, line string, args ...*Param) (int, Callback, []interface{}, CompSlice) {
	m := c.LookupMode(mode)
	if m == nil {
		return ParseNoMatch, nil, nil, nil
	}

	var param *Param
	if len(args) > 0 {
		param = args[0]
	} else {
		param = &Param{}
	}
	return m.Parser.Parse(line, param)
}
