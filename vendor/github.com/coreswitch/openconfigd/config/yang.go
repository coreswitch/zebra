// Copyright 2016 OpenConfigd Project.
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
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"

	"github.com/coreswitch/cmd"
	"github.com/coreswitch/component"
	"github.com/coreswitch/goyang/pkg/yang"
	"github.com/coreswitch/openconfigd/modules/ntp"
)

var (
	rootEntry *yang.Entry
)

func DirMatch(e *yang.Entry, s string) *yang.Entry {
	for _, ent := range e.Dir {
		if s == ent.Name {
			return ent
		}
	}
	return nil
}

func IsEmptyLeaf(e *yang.Entry) bool {
	if e.Kind == yang.LeafEntry && e.Type != nil && e.Type.Kind == yang.Yempty {
		return true
	}
	return false
}

func IsPresenceContainer(e *yang.Entry) bool {
	n, ok := e.Node.(*yang.Container)
	if !ok {
		return false
	}
	return n.Presence != nil
}

func IsMandatory(e *yang.Entry) bool {
	if n, ok := e.Node.(*yang.Leaf); ok {
		if n.Mandatory != nil {
			return true
		}
	}
	if n, ok := e.Node.(*yang.Choice); ok {
		if n.Mandatory != nil {
			return true
		}
	}
	return false
}

func HasKey(e *yang.Entry) bool {
	if e.Key != "" {
		return true
	} else {
		return false
	}
}

func KeySlice(key string) []string {
	return strings.Split(key, " ")
}

func KeyLength(ent *yang.Entry) int {
	return len(KeySlice(ent.Key))
}

func KeyEntry(ent *yang.Entry, index int) *yang.Entry {
	return ent.Dir[KeyIndexString(ent.Key, index)]
}

func KeyIncludeValue(key, value string) bool {
	slice := KeySlice(key)
	for _, val := range slice {
		if val == value {
			return true
		}
	}
	return false
}

func KeyIndexString(key string, index int) string {
	slice := KeySlice(key)
	if index < len(slice) {
		return slice[index]
	} else {
		return ""
	}
}

func ExtHelp(e *yang.Entry) string {
	if e.Exts != nil {
		for _, ext := range e.Exts {
			if ext.Keyword == "ext:help" {
				return ext.Argument
			}
		}
	}
	return ""
}

func YangSet(Args []string) (inst int, instStr string) {
	SubscribeMutex.Lock()
	defer SubscribeMutex.Unlock()
	inst = CliSuccess
	instStr = ""
	Parse(Args, rootEntry, configCandidate, nil)
	return
}

func YangConfigPush(Args []string) {
	SubscribeMutex.Lock()
	defer SubscribeMutex.Unlock()
	Parse(Args, rootEntry, configActive, nil)
	Parse(Args, rootEntry, configCandidate, nil)
}

func YangConfigPull(Args []string) {
	SubscribeMutex.Lock()
	defer SubscribeMutex.Unlock()

	c := configActive.LookupByPath(Args)
	if c != nil {
		Delete(c, true)
	}
	c = configCandidate.LookupByPath(Args)
	if c != nil {
		Delete(c, true)
	}
}

func YParseSet(param *cmd.Param) (int, cmd.Callback, []interface{}, cmd.CompSlice) {
	SubscribeMutex.Lock()
	defer SubscribeMutex.Unlock()

	// Trim "set"
	if len(param.Command) > 0 {
		param.Command = param.Command[1:]
	}
	if len(param.Command) == 0 {
		param.Command = []string{""}
	}
	status := &YMatchState{
		state:    StateDir,
		complete: true,
		trailing: param.TrailingSpace,
	}
	ret, callback, _, comps := Parse(param.Command, rootEntry, configCandidate, status)
	return ret, callback, cmd.String2Interface(param.Command), comps
}

func ProcessDelete(config *Config) {
	//fmt.Println("SubscribeMutex.Lock ProcessDelete")
	Delete(config, true)
}

func YangDelete(Args []string) (inst int, instStr string) {
	inst = CliSuccess
	instStr = ""
	return
}

///// New code

const (
	StateDir = iota
	StateDirMatched
	StateKey
	StateKeyMatched
	StateLeaf
	StateLeafMatched
	StateLeafList
	StateLeafListMatched
)

var StateStr = map[int]string{
	StateDir:             "StateDir",
	StateDirMatched:      "StateDirMatched",
	StateKey:             "StateKey",
	StateKeyMatched:      "StateKeyMatched",
	StateLeaf:            "StateLeaf",
	StateLeafMatched:     "StateLeafMatched",
	StateLeafList:        "StateLeafList",
	StateLeafListMatched: "StateLeafListMatched",
}

type YMatchState struct {
	complete bool
	match    cmd.MatchType
	count    int
	pos      int
	entry    *yang.Entry
	comps    cmd.CompSlice
	state    int
	trailing bool
	index    int
	config   *Config
}

const (
	YMatchTypeKeyword = iota
	YMatchTypeNumber
	YMatchTypeString
)

func MatchNumber(str string, e *yang.Entry) (pos int, match cmd.MatchType) {
	num, err := yang.ParseNumber(str)
	if err != nil {
		match = cmd.MatchTypeNone
		return
	}
	for _, r := range e.Type.Range {
		if num.Less(r.Min) || r.Max.Less(num) {
			match = cmd.MatchTypeNone
			return
		}
	}
	pos = len(str)
	match = cmd.MatchTypeRange
	return
}

func MatchString(str string, e *yang.Entry, typ *yang.YangType) (pos int, match cmd.MatchType) {
	if len(typ.Pattern) > 0 {
		regex, err := regexp.Compile("^" + typ.Pattern[0] + "$")
		if err != nil {
			match = cmd.MatchTypeNone
			return
		}
		if regex.MatchString(str) {
			match = cmd.MatchTypeExact
		} else {
			match = cmd.MatchTypeNone
		}
		return
	} else {
		match = cmd.MatchTypeExact
		return
	}
}

func MatchEntry(matchType int, ent *yang.Entry, typ *yang.YangType, name string, str string, state *YMatchState) {
	var pos int
	var match cmd.MatchType

	switch matchType {
	case YMatchTypeKeyword:
		pos, match = cmd.MatchKeyword(str, name)
	case YMatchTypeNumber:
		pos, match = MatchNumber(str, ent)
	case YMatchTypeString:
		pos, match = MatchString(str, ent, typ)
	}

	if match == cmd.MatchTypeNone {
		return
	}

	if state.complete {
		switch matchType {
		case YMatchTypeKeyword:
			state.comps = append(state.comps, &cmd.Comp{Name: name, Help: ExtHelp(ent), Dir: ent.IsDir()})
		case YMatchTypeNumber, YMatchTypeString:
			state.comps = append(state.comps, &cmd.Comp{Name: "<" + ent.Name + ">"})
		}
	}
	if match > state.match {
		state.match = match
		state.pos = pos
		state.entry = ent
		state.count = 1
	} else if match == state.match {
		state.count++
	}
}

func MatchType(e *yang.Entry, typ *yang.YangType, str string, state *YMatchState) {
	switch typ.Kind {
	case yang.Yint8, yang.Yint16, yang.Yint32, yang.Yint64,
		yang.Yuint8, yang.Yuint16, yang.Yuint32, yang.Yuint64:
		MatchEntry(YMatchTypeNumber, e, typ, "", str, state)
	case yang.Ystring:
		MatchEntry(YMatchTypeString, e, typ, "", str, state)
	case yang.Ybool:
		MatchEntry(YMatchTypeKeyword, e, typ, "true", str, state)
		MatchEntry(YMatchTypeKeyword, e, typ, "false", str, state)
	case yang.Yenum:
		for name, _ := range typ.Enum.NameMap() {
			MatchEntry(YMatchTypeKeyword, e, typ, name, str, state)
		}
	}
}

func MatchLeaf(e *yang.Entry, str string, state *YMatchState) {
	if e.Type == nil {
		return
	}
	if e.Type.Kind == yang.Yleafref {
		e = e.Find(e.Type.Path)
		if e == nil {
			return
		}
	}
	if e.Type.Kind == yang.Yunion {
		for _, typ := range e.Type.Type {
			MatchType(e, typ, str, state)
		}
	} else {
		MatchType(e, e.Type, str, state)
	}
}

func MatchDir(ent *yang.Entry, str string, status *YMatchState) {
	for _, e := range ent.Dir {
		if e.IsChoice() {
			MatchChoice(e, str, status)
		} else {
			MatchEntry(YMatchTypeKeyword, e, e.Type, e.Name, str, status)
		}
	}
}

func MatchChoice(ent *yang.Entry, str string, status *YMatchState) {
	for _, e := range ent.Dir {
		MatchDir(e, str, status)
	}
}

func MatchKey(ent *yang.Entry, str string, status *YMatchState) {
	key := KeyEntry(ent, status.index)
	if key != nil {
		MatchLeaf(key, str, status)
	}
}

func MatchKeyMatched(ent *yang.Entry, str string, status *YMatchState) {
	for _, e := range ent.Dir {
		if !KeyIncludeValue(ent.Key, e.Name) {
			if e.IsChoice() {
				MatchChoice(e, str, status)
			} else {
				MatchEntry(YMatchTypeKeyword, e, e.Type, e.Name, str, status)
			}
		}
	}
}

func CompChoice(ent *yang.Entry, comps cmd.CompSlice) cmd.CompSlice {
	for _, e := range ent.Dir {
		comps = CompDir(e, comps)
	}
	return comps
}

func IsCompDir(e *yang.Entry) bool {
	if e == nil {
		return false
	}
	if e.IsDir() {
		if e.IsList() {
			if len(e.Dir) > KeyLength(e) {
				return true
			}
		} else {
			if len(e.Dir) > 0 {
				return true
			}
		}
	}
	return false
}

func IsAdditive(e *yang.Entry) bool {
	if e == nil {
		return false
	}
	if e.IsList() && len(e.Key) > 0 {
		return true
	}
	return false
}

func CompDir(ent *yang.Entry, comps cmd.CompSlice) cmd.CompSlice {
	for _, e := range ent.Dir {
		if e.IsChoice() {
			comps = CompChoice(e, comps)
		} else {
			comps = append(comps, &cmd.Comp{Name: e.Name, Dir: IsCompDir(e), Additive: IsAdditive(e)})
		}
	}
	return comps
}

func CompLeaf(ent *yang.Entry) cmd.CompSlice {
	comps := cmd.CompSlice{}
	if EntryExpandable(ent) {
		for _, name := range EntryExpanded(ent) {
			comps = append(comps, &cmd.Comp{Name: name})
		}
	} else {
		comps = append(comps, &cmd.Comp{Name: "<" + ent.Name + ">"})
	}
	return comps
}

func CompKey(ent *yang.Entry, index int) cmd.CompSlice {
	comps := cmd.CompSlice{}
	key := KeyEntry(ent, index)
	if key != nil {
		comps = append(comps, &cmd.Comp{Name: "<" + key.Name + ">", Dir: IsCompDir(ent), Additive: IsAdditive(ent)})
	}
	return comps
}

func CompKeyMatched(ent *yang.Entry, comps cmd.CompSlice) cmd.CompSlice {
	for _, e := range ent.Dir {
		if !KeyIncludeValue(ent.Key, e.Name) {
			if e.IsChoice() {
				comps = CompChoice(e, comps)
			} else {
				comps = append(comps, &cmd.Comp{Name: e.Name, Dir: IsCompDir(e), Additive: IsAdditive(e)})
			}
		}
	}
	return comps
}

func EntryExpandable(e *yang.Entry) bool {
	if e.Type.Kind == yang.Ybool || e.Type.Kind == yang.Yenum {
		return true
	} else {
		return false
	}
}

func EntryExpanded(e *yang.Entry) []string {
	if e.Type.Kind == yang.Ybool {
		return []string{"false", "true"}
	}
	if e.Type.Kind == yang.Yenum {
		comp := []string{}
		for key, _ := range e.Type.Enum.NameMap() {
			comp = append(comp, key)
		}
		return comp
	}
	return nil
}

func EntryNextState(e *yang.Entry) int {
	switch e.Kind {
	case yang.DirectoryEntry:
		if HasKey(e) {
			return StateKey
		} else if IsPresenceContainer(e) {
			return StateDirMatched
		} else {
			return StateDir
		}
	case yang.LeafEntry:
		if e.IsLeafList() {
			return StateLeafList
		} else {
			return StateLeaf
		}
	}
	return StateDir
}

func CompHasName(comps cmd.CompSlice, name string) bool {
	for _, comp := range comps {
		if comp.Name == name {
			return true
		}
	}
	return false
}

func Parse(cmds []string, ent *yang.Entry, config *Config, s *YMatchState) (int, cmd.Callback, []interface{}, cmd.CompSlice) {
	str := cmds[0]

	if s == nil {
		s = &YMatchState{state: StateDir}
	}
	s.count = 0
	s.match = 0
	s.comps = s.comps[:0]

	switch s.state {
	case StateDir, StateDirMatched:
		MatchDir(ent, str, s)
	case StateKey:
		MatchKey(ent, str, s)
	case StateKeyMatched:
		MatchKeyMatched(ent, str, s)
	case StateLeaf:
		MatchLeaf(ent, str, s)
	case StateLeafMatched:
		// Nothing to do.
	case StateLeafList, StateLeafListMatched:
		MatchLeaf(ent, str, s)
	}

	sort.Sort(s.comps)

	// Set completion.
	cs := &YMatchState{complete: true}
	if config != nil && s.complete {
		if config.HasDir() {
			MatchConfigDir(config, str, cs)
		} else {
			MatchConfigValue(config, str, cs)
			cs.index++
		}
		for _, comp := range cs.comps {
			if !CompHasName(s.comps, comp.Name) {
				s.comps = append(s.comps, comp)
			}
		}
		if cs.count == 1 && cs.match == cmd.MatchTypeExact {
			config = cs.config
		} else {
			config = nil
		}
	}

	if s.count == 0 {
		return cmd.ParseNoMatch, nil, nil, s.comps
	}
	if s.count > 1 {
		return cmd.ParseAmbiguous, nil, nil, s.comps
	}

	matched := s.entry
	next := s.state

	switch s.state {
	case StateDir, StateDirMatched, StateKeyMatched:
		ent = matched
		next = EntryNextState(matched)
		if next == StateKey {
			s.index = 0
		}
	case StateKey:
		s.index++
		if s.index >= KeyLength(ent) {
			next = StateKeyMatched
		}
	case StateLeaf:
		next = StateLeafMatched
	case StateLeafList:
		next = StateLeafListMatched
	case StateLeafMatched, StateLeafListMatched:
		// Keep current state.
	}

	// Config set mode.
	if config != nil && !s.complete {
		switch next {
		case StateDir, StateDirMatched:
			config = config.Set(matched)
		case StateKey:
			if s.state != StateKey {
				config = config.Set(matched)
			} else {
				config = config.SetKey(matched, str, ent.Name, false)
			}
		case StateKeyMatched:
			config = config.SetKey(matched, str, ent.Name, true)
		case StateLeaf:
			config = config.Set(matched)
		case StateLeafMatched:
			config = config.SetValue(str)
		case StateLeafList:
			config = config.SetLeafList(matched)
		case StateLeafListMatched:
			config = config.SetListValue(str)
		}
	}

	if next == StateLeaf && IsEmptyLeaf(matched) {
		next = StateLeafMatched
	}

	s.state = next

	cmds = cmds[1:]
	if len(cmds) == 0 {
		if s.complete && s.trailing {
			s.comps = s.comps[:0]
			switch next {
			case StateDir, StateDirMatched:
				s.comps = CompDir(ent, s.comps)
			case StateLeaf, StateLeafList, StateLeafListMatched:
				s.comps = CompLeaf(ent)
			case StateLeafMatched:
				// Match is done.  No need to have completion.
			case StateKey:
				s.comps = CompKey(ent, s.index)
			case StateKeyMatched:
				s.comps = CompKeyMatched(ent, s.comps)
			}

			sort.Sort(s.comps)

			if config != nil {
				cs.comps = cs.comps[:0]
				if config.HasDir() {
					cs.comps = CompConfig(config, cs.comps)
				} else {
					cs.comps = CompValue(config, cs.comps, cs.index)
				}
				for _, comp := range cs.comps {
					if !CompHasName(s.comps, comp.Name) {
						s.comps = append(s.comps, comp)
					}
				}
			}
		}

		if s.match == cmd.MatchTypePartial {
			return cmd.ParseIncomplete, nil, nil, s.comps
		}

		if next == StateLeafMatched || next == StateLeafListMatched || next == StateKeyMatched || next == StateDirMatched {
			return cmd.ParseSuccess, YangSet, nil, s.comps
		} else {
			return cmd.ParseIncomplete, nil, nil, s.comps
		}
	}

	return Parse(cmds, ent, config, s)
}

// For ParseDelete

func MatchConfig(config *Config, val string, str string, s *YMatchState) {
	pos, match := cmd.MatchKeyword(str, val)
	if match == cmd.MatchTypeNone {
		return
	}

	if s.complete {
		s.comps = append(s.comps, &cmd.Comp{Name: val})
	}
	if match > s.match {
		s.match = match
		s.pos = pos
		s.config = config
		s.count = 1
	} else if match == s.match {
		s.count++
	}
}

func MatchConfigDir(config *Config, str string, s *YMatchState) {
	for _, c := range config.Configs {
		MatchConfig(c, c.Name, str, s)
	}
	for _, c := range config.Keys {
		MatchConfig(c, c.Name, str, s)
	}
}

func MatchConfigValue(config *Config, str string, s *YMatchState) {
	if config.Entry.IsLeafList() {
		if s.index < len(config.ValueList) {
			MatchConfig(config, config.ValueList[s.index], str, s)
		}
	} else {
		if s.index == 0 {
			MatchConfig(config, config.Value, str, s)
		}
	}
}

func CompConfig(config *Config, comps cmd.CompSlice) cmd.CompSlice {
	for _, c := range config.Configs {
		comps = append(comps, &cmd.Comp{Name: c.Name})
	}
	for _, c := range config.Keys {
		comps = append(comps, &cmd.Comp{Name: c.Name})
	}
	return comps

}

func CompValue(config *Config, comps cmd.CompSlice, index int) cmd.CompSlice {
	if config.Entry.IsLeafList() {
		if index < len(config.ValueList) {
			comps = append(comps, &cmd.Comp{Name: config.ValueList[index]})
		}
	} else {
		if index == 0 && !IsEmptyLeaf(config.Entry) {
			comps = append(comps, &cmd.Comp{Name: config.Value})
		}
	}
	return comps
}

func (c *Config) HasDir() bool {
	return len(c.Configs) > 0 || len(c.Keys) > 0 || c.KeyConfig || c.Entry == nil
}

func ParseDelete(cmds []string, config *Config, s *YMatchState) (int, cmd.Callback, []interface{}, cmd.CompSlice) {
	str := cmds[0]

	if s == nil {
		s = &YMatchState{state: StateDir}
	}
	s.count = 0
	s.match = 0
	s.comps = s.comps[:0]

	if config.HasDir() {
		MatchConfigDir(config, str, s)
	} else {
		MatchConfigValue(config, str, s)
		s.index++
	}

	if s.count == 0 {
		return cmd.ParseNoMatch, nil, nil, s.comps
	}
	if s.count > 1 {
		return cmd.ParseAmbiguous, nil, nil, s.comps
	}

	config = s.config

	cmds = cmds[1:]
	if len(cmds) == 0 {
		if s.complete && s.trailing {
			s.comps = s.comps[:0]
			if config.HasDir() {
				s.comps = CompConfig(config, s.comps)
			} else {
				s.comps = CompValue(config, s.comps, s.index)
			}
		}
		if s.match != cmd.MatchTypeExact {
			return cmd.ParseIncomplete, nil, nil, s.comps
		}
		if config.IsValueLeaf() || config.IsContainer() {
			// Execute delete.
			if !s.complete {
				// We need sumarter treatment of read only config.
				if config.Entry.ReadOnlyConfig {
					return cmd.ParseNoMatch, nil, nil, nil
				} else {
					ProcessDelete(config)
				}
			}
			return cmd.ParseSuccess, YangDelete, nil, s.comps
		} else {
			return cmd.ParseIncomplete, nil, nil, s.comps
		}
	}

	return ParseDelete(cmds, config, s)
}

func YParseDelete(param *cmd.Param) (int, cmd.Callback, []interface{}, cmd.CompSlice) {
	SubscribeMutex.Lock()
	defer SubscribeMutex.Unlock()

	// Trim "delete"
	if len(param.Command) > 0 {
		param.Command = param.Command[1:]
	}
	if len(param.Command) == 0 {
		param.Command = []string{""}
	}
	status := &YMatchState{
		state:    StateDir,
		complete: param.Complete,
		trailing: param.TrailingSpace,
	}
	return ParseDelete(param.Command, configCandidate, status)
}

func EntryDump(e *yang.Entry, depth int) {
	if depth != 0 {
		fmt.Printf("%*s", depth*2, " ")
	}
	fmt.Printf("%s ", e.Name)
	if len(e.Key) != 0 {
		fmt.Printf("[%s] ", e.Key)
	}
	if e.Kind != yang.LeafEntry {
		fmt.Printf("%s\n", yang.EntryKindToName[e.Kind])
	} else {
		if e.IsLeafList() {
			fmt.Printf("%s\n", "LeafList")
		} else {
			fmt.Printf("%s\n", "Leaf")
		}
	}
	if e.Kind == yang.ChoiceEntry {
		fmt.Println(e)
	}
	if e.Kind == yang.DirectoryEntry || e.Kind == yang.ChoiceEntry {
		for _, ent := range e.Dir {
			EntryDump(ent, depth+1)
		}
	}
}

// Yang entry lookup.
func EntryLookup(e *yang.Entry, p []string) *yang.Entry {
	for _, path := range p {
		ent := DirMatch(e, path)
		if ent == nil {
			return nil
		}
		e = ent
	}
	return e
}

// Yang component.
type YangComponent struct {
	YangPaths   string
	YangModules []string
}

// Yang component start method.
func (this *YangComponent) Start() component.Component {
	// Set up Yang file load path. Append GOPATH + openconfigd's source yang
	// directory as well.
	yang.AddPath(this.YangPaths)
	yang.AddPath(Env("GOPATH") + "/src/github.com/coreswitch/openconfigd/yang")

	// Initialize YANG modules
	ms := yang.NewModules()

	// Read YANG modules.
	for _, name := range this.YangModules {
		if err := ms.Read(name); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
	ms.Process()

	// Avoid duplicated module name.
	mods := map[string]*yang.Module{}
	var names []string
	for _, m := range ms.Modules {
		if mods[m.Name] == nil {
			mods[m.Name] = m
			names = append(names, m.Name)
		}
	}

	// Unique list of module names.
	rootEntry = &yang.Entry{
		Kind: yang.DirectoryEntry,
		Dir:  map[string]*yang.Entry{},
	}
	for _, name := range names {
		e := yang.ToEntry(mods[name])
		for key, value := range e.Dir {
			rootEntry.Dir[key] = value
		}
	}

	// Add local subscription.
	SubscribeLocalAdd([]string{"system"}, nil)
	SubscribeLocalAdd([]string{"system", "ntp"}, ntp.Configure)
	SubscribeLocalAdd([]string{"protocols"}, nil)
	SubscribeLocalAdd([]string{"vrrp"}, VrrpJsonConfig)
	SubscribeLocalAdd([]string{"dhcp"}, DhcpJsonConfig)
	SubscribeLocalAdd([]string{"vrf", "name", "*", "vrrp"}, VrrpJsonConfig)
	SubscribeLocalAdd([]string{"vrf", "name", "*", "dhcp"}, DhcpJsonConfig)
	SubscribeLocalAdd([]string{"vrf", "name", "*", "ntp"}, ntp.Configure)
	SubscribeLocalAdd([]string{"interfaces", "interface", "*", "dhcp-relay-group"}, nil)

	// ReadOnlyConfig
	ent := EntryLookup(rootEntry, []string{"interfaces", "interface", "name"})
	if ent != nil {
		ent.ReadOnlyConfig = true
	}

	return this
}

// Yang component stop method.
func (this *YangComponent) Stop() component.Component {
	yang.Path = nil
	return this
}
