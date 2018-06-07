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
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"os/user"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/coreswitch/component"
	"github.com/coreswitch/goyang/pkg/yang"
)

var (
	configActiveFile   string  // Currently active configuration file name.
	configFileDir      string  // Config save path.
	configActive       *Config // Currently active configuration.
	configCandidate    *Config // Candidate configuration.
	configFileBasename string  // Configuration base filename.
	twoPhaseCommit     bool
	zeroConfig         bool
)

type Config struct {
	Name          string
	Entry         *yang.Entry
	Parent        *Config
	KeyConfig     bool
	KeyOnlyConfig bool
	HasValue      bool
	HasListValue  bool
	Configs       ConfigSlice
	Keys          ConfigSlice
	Value         string
	ValueList     []string
	Prefix        string
	Case          *yang.Entry
}

type ConfigSlice []*Config

func (c *Config) IsKeyConfig() bool {
	return len(c.Keys) > 0 && c.KeyConfig
}

func (c *Config) IsLeaf() bool {
	return (len(c.Keys) == 0 && len(c.Configs) == 0) || (c.Entry != nil && c.Entry.Kind == yang.LeafEntry)
}

func (c *Config) IsContainer() bool {
	return c.Entry != nil && c.Entry.IsContainer()
}

func (c *Config) IsValueLeaf() bool {
	return c.IsLeaf() && !c.IsKeyConfig()
}

func (c *Config) IsPresenceContainer() bool {
	return c.Entry != nil && IsPresenceContainer(c.Entry)
}

func (c *Config) lookup(name string) *Config {
	for _, n := range c.Configs {
		if n.Name == name {
			return n
		}
	}
	return nil
}

func (c *Config) lookupKeyShallow(key string) (*Config, *Config) {
	for _, n := range c.Keys {
		if n.Name == key {
			return c, n
		}
	}
	return c, nil
}

func (c *Config) lookupKey(key string) (*Config, *Config) {
	if c.IsKeyConfig() {
		for _, p := range c.Parent.Keys {
			_, n := p.lookupKeyShallow(key)
			if n != nil {
				return p, n
			}
		}
		return c, nil
	}
	return c.lookupKeyShallow(key)
}

func ConfigLookupVrf(ifName string) string {
	vrf := ""
	c := configCandidate.LookupByPath([]string{"interfaces", "interface", ifName, "vrf"})
	if c != nil {
		vrf = c.Value
	}
	return vrf
}

func (c *Config) LookupByPath(path []string) *Config {
	var next *Config
	for _, p := range path {
		// fmt.Println("LookupByPath: ", p)
		next = c.lookup(p)
		if next == nil {
			_, next = c.lookupKey(p)
			if next == nil {
				// fmt.Println("LookupByPath: can't find", p)
				return nil
			}
		}
		c = next
	}
	return c
}

func (c *Config) Empty() bool {
	if len(c.Configs) == 0 && len(c.Keys) == 0 {
		return true
	}
	return false
}

// Static config priority until we add priority to YANG entry.
func (c *Config) Priority() int {
	if c.Entry.Name == "vrf" {
		return 150
	}
	if c.Entry.Name == "vlans" {
		return 100
	}
	if c.Entry.Name == "interfaces" {
		return 50
	}
	if c.Entry.Name == "interface" {
		return 10
	}
	if c.Entry.Name == "subnet" {
		return 5
	}
	if c.Entry.Name == "section-start-ip" {
		return 1
	}
	if c.Entry.Name == "range-start-ip" {
		return 1
	}
	return 0
}

// ConfigSlice sort
func (configs ConfigSlice) Len() int {
	return len(configs)
}

func (configs ConfigSlice) Less(i, j int) bool {
	pi := configs[i].Priority()
	pj := configs[j].Priority()
	if pi != pj {
		return pi > pj
	} else {
		inum, ierr := strconv.Atoi(configs[i].Name)
		jnum, jerr := strconv.Atoi(configs[j].Name)
		if ierr == nil && jerr == nil {
			return inum < jnum
		} else {
			return configs[i].Name < configs[j].Name
		}
	}
}

func (configs ConfigSlice) Swap(i, j int) {
	configs[i], configs[j] = configs[j], configs[i]
}

func ConfigDumpTool(c *Config) {
	ConfigDump(c, -1, false)
}

func ConfigDump(c *Config, depth int, keynode bool) {
	if c.Name != "" {
		if depth != 0 {
			fmt.Printf("%*s", depth*2, " ")
		}
		if keynode {
			fmt.Printf("-> ")
		}
		fmt.Printf("%s", c.Name)
		if HasKey(c.Entry) {
			fmt.Printf("[hasKey]")
		}
		if c.KeyConfig {
			fmt.Printf("[key]")
		}
		if c.KeyOnlyConfig {
			fmt.Printf("[keyOnly]")
		}
		if c.Prefix != "" {
			fmt.Printf("(%s)", c.Prefix)
		}
		if c.Value != "" {
			fmt.Printf(": %s", c.Value)
		}
		if len(c.ValueList) > 0 {
			fmt.Printf(": [")
			for pos, v := range c.ValueList {
				if pos != 0 {
					fmt.Printf(",")
				}
				fmt.Printf("%s", v)
			}
			fmt.Printf("]")
		}
		fmt.Printf("\n")
	}
	for _, key := range c.Keys {
		ConfigDump(key, depth+1, true)
	}
	for _, cfg := range c.Configs {
		ConfigDump(cfg, depth+1, false)
	}
}

func ConfigDumpCandidate() {
	return
	fmt.Println("-------")
	ConfigDump(configCandidate, -1, false)
}

func CaseEntry(e *yang.Entry) *yang.Entry {
	if e.Parent != nil && e.Parent.IsCase() {
		return e.Parent
	}
	return nil
}

func exclude(before ConfigSlice, c *Config) ConfigSlice {
	var after ConfigSlice

	for _, n := range before {
		if n.Case == nil || n.Case == c.Case {
			after = append(after, n)
		} else {
			// fmt.Println("Exclude case", n.Case.Name)
		}
	}
	return after
}

func (c *Config) Set(e *yang.Entry) *Config {
	// if e.Parent != nil && e.Parent.IsCase() {
	// 	fmt.Println("Case", e.Parent.Name)
	// }
	n := c.lookup(e.Name)
	if n == nil {
		n = &Config{Name: e.Name, Entry: e, Case: CaseEntry(e)}
		n.Parent = c
		c.Configs = append(c.Configs, n)
		if n.Case != nil {
			c.Configs = exclude(c.Configs, n)
		}
		sort.Sort(c.Configs)
	}
	ConfigDumpCandidate()
	return n
}

func (c *Config) SetKey(e *yang.Entry, key string, prefix string, last bool) *Config {
	c, n := c.lookupKey(key)
	if n == nil {
		n = &Config{Name: key, Entry: e, KeyConfig: true, Prefix: prefix}
		n.Parent = c
		if c.KeyConfig {
			c.KeyOnlyConfig = true
		}
		c.Keys = append(c.Keys, n)
		sort.Sort(c.Keys)
	}
	ConfigDumpCandidate()
	return n
}

func (c *Config) SetLeafList(e *yang.Entry) *Config {
	n := c.Set(e)
	n.ValueList = []string{}
	return n
}

func (c *Config) SetValue(value string) *Config {
	c.Value = value
	c.HasValue = true
	ConfigDumpCandidate()
	return c
}

func (c *Config) SetListValue(value string) *Config {
	c.ValueList = append(c.ValueList, value)
	c.HasValue = true
	ConfigDumpCandidate()
	return c
}

func (c *Config) Delete(n *Config) {
	// fmt.Println("[c.Delete] ", c.Name, n.Name)
	configs := []*Config{}
	for _, conf := range c.Configs {
		if conf != n {
			configs = append(configs, conf)
		}
	}
	c.Configs = configs
}

func (c *Config) DeleteKey(n *Config) {
	configs := []*Config{}
	for _, conf := range c.Keys {
		if conf != n {
			configs = append(configs, conf)
		}
	}
	c.Keys = configs
}

func Delete(c *Config, leaf bool) {
	// fmt.Println("[Delete]", c.Name, leaf)
	if c.Entry != nil {
		if c.Entry.Kind == yang.LeafEntry {
			// fmt.Println("[Delete] leafEntry", c.Name, c.Value)
			if leaf {
				if (c.HasValue || IsEmptyLeaf(c.Entry)) && c.Parent != nil {
					c.Parent.Delete(c)
				}
				if c.KeyConfig && c.Parent != nil {
					c.Parent.DeleteKey(c)
				}
			} else {
				if len(c.Configs) == 0 && len(c.Keys) == 0 {
					if c.KeyOnlyConfig {
						c.Parent.DeleteKey(c)
					}
				}
			}
		}
		if c.Entry.Kind == yang.DirectoryEntry {
			// fmt.Println("[Delete] directoryEntry:", c.Name)
			if leaf {
				c.Configs = c.Configs[:0]
				c.Keys = c.Keys[:0]
			}
			if len(c.Configs) == 0 && len(c.Keys) == 0 && c.Parent != nil {
				// nfmt.Println("Removing directory")
				c.Parent.Delete(c)
			}
		}
	}
	if c.Parent != nil && !c.Parent.IsPresenceContainer() {
		Delete(c.Parent, false)
	}
}

func (c *Config) quote() bool {
	if c.Entry.Type.Kind == yang.Ystring {
		if len(c.Entry.Type.Pattern) == 0 {
			return true
		}
	}
	return false
}

func (c *Config) CommandList(list []*Config) []*Config {
	if c.IsValueLeaf() || c.IsPresenceContainer() {
		list = append(list, c)
	}
	for _, n := range c.Keys {
		list = n.CommandList(list)
	}
	for _, n := range c.Configs {
		list = n.CommandList(list)
	}
	return list
}

func (c *Config) CommandPath() []string {
	if c.Parent != nil {
		ret := append(c.Parent.CommandPath(), c.Name)
		if c.Value != "" {
			ret = append(ret, c.Value)
		}
		if len(c.ValueList) > 0 {
			for _, v := range c.ValueList {
				ret = append(ret, v)
			}
		}
		return ret
	}
	return []string{}
}

func (c *Config) Command() *Command {
	cmd := &Command{
		set:  true,
		cmds: c.CommandPath(),
	}
	return cmd
}

func (c *Config) Copy(parent *Config) *Config {
	config := &Config{
		Name:      c.Name,
		Entry:     c.Entry,
		Parent:    parent,
		KeyConfig: c.KeyConfig,
		HasValue:  c.HasValue,
		Configs:   make([]*Config, len(c.Configs)),
		Keys:      make([]*Config, len(c.Keys)),
		Value:     c.Value,
		ValueList: c.ValueList,
		Prefix:    c.Prefix,
	}
	for pos, subConfig := range c.Configs {
		config.Configs[pos] = subConfig.Copy(config)
	}
	for pos, subKey := range c.Keys {
		config.Keys[pos] = subKey.Copy(config)
	}
	return config
}

func isConfigKeyNode(c *Config) bool {
	if c.Entry.Key == "" {
		return false
	} else {
		return true
	}
}

func (c *Config) CommandLine() []string {
	if c.Parent != nil {
		strs := append(c.Parent.CommandLine(), c.Name)
		if c.Value != "" {
			if c.quote() {
				strs = append(strs, "\""+c.Value+"\"")
			} else {
				strs = append(strs, c.Value)
			}
		}
		if len(c.ValueList) > 0 {
			for _, v := range c.ValueList {
				if c.quote() {
					strs = append(strs, "\""+v+"\"")
				} else {
					strs = append(strs, v)
				}
			}
		}
		return strs
	}
	return []string{}
}

func (c *Config) writeCommand(out io.Writer) {
	if c.IsValueLeaf() || c.IsPresenceContainer() {
		fmt.Fprintf(out, "set "+strings.Join(c.CommandLine(), " ")+"\n")
	}
	for _, n := range c.Keys {
		n.writeCommand(out)
	}
	for _, n := range c.Configs {
		n.writeCommand(out)
	}
}

func (c *Config) CommandString() string {
	buf := new(bytes.Buffer)
	c.writeCommand(buf)
	return buf.String()
}

func (c *Config) WriteCommandTo(path string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("File can't be created")
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, config := range c.Configs {
		config.writeCommand(w)
	}
	w.Flush()
}

func (c *Config) DisplayEntry() bool {
	if len(c.Keys) > 0 {
		return false
	}
	return true
}

func (c *Config) PrefixWrite(out io.Writer) {
	if c.Prefix == "" {
		return
	}
	if c.Parent != nil && c.Parent.Prefix != "" {
		c.Parent.PrefixWrite(out)
	}
	if c.Prefix != "" {
		fmt.Fprintf(out, "%s ", c.Parent.Name)
	}
}

func (c *Config) write(out io.Writer, depth int) {
	brace := true
	keyFirst := false

	if len(c.Keys) == 0 && len(c.Configs) == 0 {
		brace = false
	}
	if len(c.Keys) > 0 && c.KeyConfig {
		keyFirst = true
		brace = false
	}

	if c.DisplayEntry() {
		if depth != 0 {
			fmt.Fprintf(out, "%*s", depth*4, " ")
		}
		c.PrefixWrite(out)
		fmt.Fprintf(out, "%s", c.Name)

		if c.Value != "" {
			if c.quote() {
				fmt.Fprintf(out, " \"%s\"", c.Value)
			} else {
				fmt.Fprintf(out, " %s", c.Value)
			}
		}

		if len(c.ValueList) > 0 {
			for _, v := range c.ValueList {
				if c.quote() {
					fmt.Fprintf(out, " \"%s\"", v)
				} else {
					fmt.Fprintf(out, " %s", v)
				}
			}
		}
		if brace {
			fmt.Fprintf(out, " {\n")
		} else {
			if !keyFirst {
				fmt.Fprintf(out, ";\n")
			}
		}
	}

	for _, n := range c.Keys {
		n.write(out, depth)
	}
	for _, n := range c.Configs {
		n.write(out, depth+1)
	}

	if c.DisplayEntry() {
		if brace {
			if depth != 0 {
				fmt.Fprintf(out, "%*s", depth*4, " ")
			}
			fmt.Fprintf(out, "}\n")
		}
	}
}

func (c *Config) hasPrefix() bool {
	return c.Prefix != ""
}

func (c *Config) needQuote() bool {
	if c.Entry != nil && c.Entry.Type != nil {
		switch c.Entry.Type.Kind {
		case yang.Yint8, yang.Yint16, yang.Yint32, yang.Yint64,
			yang.Yuint8, yang.Yuint16, yang.Yuint32, yang.Yuint64,
			yang.Ybool:
			return false
		default:
			return true
		}
	}
	return true
}

func mandatoryFindList(c *Config, e *yang.Entry, depth int) error {
	if depth == 0 {
		for _, cfg := range c.Configs {
			if cfg.Entry == e {
				// Found mandatory node.
				return nil
			}
		}
		// Couldn't find mandatory node.
		return fmt.Errorf("Mandatory node '%s' is necessary under %s.", e.Name, c.Name)
	} else {
		for _, cfg := range c.Keys {
			err := mandatoryFindList(cfg, e, depth-1)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func mandatoryCheck(c *Config, e *yang.Entry) error {
	// Check list only.
	if e.IsList() {
		for _, ent := range e.Dir {
			if ent.IsList() {
				// Need to check the List has mandatory leaf key.
				for _, leaf := range ent.Dir {
					if IsMandatory(leaf) && KeyIncludeValue(ent.Key, leaf.Name) {
						err := mandatoryFindList(c, ent, KeyLength(c.Entry))
						if err != nil {
							return err
						}
					}
				}
			}
			if ent.IsLeaf() {
				// Need to check non key mandatory leaf.
				if IsMandatory(ent) && !KeyIncludeValue(e.Key, ent.Name) {
					err := mandatoryFindList(c, ent, KeyLength(c.Entry))
					if err != nil {
						return err
					}
				}
			}
		}
	}
	if e.IsContainer() && !IsPresenceContainer(e) {
		for _, ent := range e.Dir {
			if ent.IsLeaf() {
				if IsMandatory(ent) {
					err := mandatoryFindList(c, ent, 0)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (c *Config) MandatoryCheck() error {
	if c.Entry != nil {
		err := mandatoryCheck(c, c.Entry)
		if err != nil {
			return err
		}
	}
	for _, k := range c.Keys {
		err := k.MandatoryCheck()
		if err != nil {
			return err
		}
	}
	for _, c := range c.Configs {
		err := c.MandatoryCheck()
		if err != nil {
			return err
		}
	}
	return nil
}

func YEntryJson(e *yang.Entry) string {
	if e.Type == nil {
		return ""
	}
	if e.Type.Kind == yang.Yempty {
		return "true"
	}
	return ""
}

func (c *Config) jsonQuotedString(name, value string) string {
	if c.needQuote() {
		return fmt.Sprintf(`"%s":"%s"`, name, value)
	} else {
		return fmt.Sprintf(`"%s":%s`, name, value)
	}
}

func (c *Config) jsonMarshalLeafList() string {
	str := ""
	for pos, val := range c.ValueList {
		if pos != 0 {
			str += ","
		}
		if c.needQuote() {
			str += fmt.Sprintf(`"%s"`, val)
		} else {
			str += fmt.Sprintf(`%s`, val)
		}
	}
	return fmt.Sprintf(`"%s":[%s]`, c.Name, str)
}

func (c *Config) jsonMarshal(pos int) (str string) {
	if pos != 0 {
		str += ","
	}

	if isConfigKeyNode(c) {
		str += `"` + c.Name + `": [`
	} else {
		if c.hasPrefix() {
			if c.Parent != nil && !c.Parent.KeyOnlyConfig {
				str += `{`
			}
			str += c.jsonQuotedString(c.Entry.Name, c.Name)
		} else {
			if c.Value != "" {
				str += c.jsonQuotedString(c.Name, c.Value)
			} else if len(c.ValueList) > 0 {
				str += c.jsonMarshalLeafList()
			} else {
				str += fmt.Sprintf(`"%s":%s`, c.Name, YEntryJson(c.Entry))
			}
		}
	}

	if len(c.Keys) != 0 {
		for pos, n := range c.Keys {
			if c.KeyOnlyConfig {
				str += n.jsonMarshal(pos + 1)
			} else {
				str += n.jsonMarshal(pos)
			}
		}
	}

	if len(c.Configs) != 0 {
		if !c.hasPrefix() {
			str += "{"
		}
		for pos, n := range c.Configs {
			if c.hasPrefix() {
				str += n.jsonMarshal(pos + 1)
			} else {
				str += n.jsonMarshal(pos)
			}
		}
		str += "}"
	} else {
		if c.hasPrefix() {
			if c.Parent != nil && c.Parent.KeyOnlyConfig {
				str += "}"
			} else if len(c.Keys) == 0 {
				str += "}"
			}
		}
	}

	if len(c.Keys) == 0 && len(c.Configs) == 0 && c.IsPresenceContainer() {
		str += "{}"
	}

	if isConfigKeyNode(c) {
		str += "]"
	}

	return
}

func (c *Config) JsonMarshal() string {
	var str string

	if len(c.Keys) > 0 {
		for pos, config := range c.Keys {
			if pos != 0 {
				str += ","
			}
			str += config.jsonMarshal(0)
		}
		return "[" + str + "]"
	}
	for pos, config := range c.Configs {
		if pos != 0 {
			str += ","
		}
		str += config.jsonMarshal(0)
	}
	return "{" + str + "}"
}

func (c *Config) String() string {
	buf := new(bytes.Buffer)
	for _, config := range c.Configs {
		config.write(buf, 0)
	}
	return buf.String()
}

func (c *Config) Signature(out io.Writer, via string) {
	const layout = "2006-01-02 15:04:05 MST"
	username := "anonymous"
	user, _ := user.Current()
	if user != nil {
		username = user.Username
	}
	fmt.Fprintf(out, "# %s by %s via %s\n", time.Now().Format(layout), username, via)
}

func (c *Config) WriteTo(path string, by ...string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("File can't be created")
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	if len(by) > 0 {
		c.Signature(w, by[0])
	}
	for _, config := range c.Configs {
		config.write(w, 0)
	}
	w.Flush()
}

func ConfigDiscard() bool {
	SubscribeMutex.Lock()
	defer SubscribeMutex.Unlock()

	diff := CompareCommand()
	if diff != "" {
		configCandidate = configActive.Copy(nil)
		return true
	} else {
		return false
	}
}

// Config component.
type ConfigComponent struct {
	ConfigActiveFile string
	ConfigFileDir    string
	TwoPhaseCommit   bool
	ZeroConfig       bool
}

// Config component start method.
func (this *ConfigComponent) Start() component.Component {
	configActive = &Config{}
	configCandidate = &Config{}

	// When active file is absolute path, it overwrite opts.ConfigFileDir.
	if path.IsAbs(this.ConfigActiveFile) {
		this.ConfigFileDir = path.Dir(this.ConfigActiveFile)
	} else {
		this.ConfigActiveFile = this.ConfigFileDir + "/" + this.ConfigActiveFile
	}

	configFileDir = this.ConfigFileDir
	configActiveFile = this.ConfigActiveFile
	configFileBasename = path.Base(configActiveFile)
	twoPhaseCommit = this.TwoPhaseCommit
	zeroConfig = this.ZeroConfig

	// Load saved configuration.
	if zeroConfig {
		err := Load(configActiveFile)
		if err != nil {
			fmt.Println("Can't load config:", err)
			return this
		}
	} else {
		err := Load(configActiveFile + ".0")
		if err != nil {
			err = Load(configActiveFile)
			if err != nil {
				fmt.Println("Can't load config:", err)
				return this
			}
		}
	}

	// Check configCandidate is properly loaded.
	if configCandidate.Empty() {
		fmt.Println("Loaded config is empty")
	} else {
		Commit()
	}

	return this
}

// Config component stop method.
func (this *ConfigComponent) Stop() component.Component {
	DiscardConfigChange()
	DhcpExitFunc()
	VrrpServerStopAll()
	RelayExitFunc()
	QuaggaExit()
	OspfVrfExit()
	DistributeListExit()
	GobgpWanExit()

	return this
}
