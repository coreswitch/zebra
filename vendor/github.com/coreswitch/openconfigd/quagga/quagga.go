package quagga

import (
	"fmt"
	"github.com/coreswitch/cmd"
	"io"
	"os/exec"
	"sort"
	"strconv"
)

const (
	QUAGGAD_MODULE = "quaggad"
	QUAGGAD_PORT   = 2699
)

var (
	validating            bool
	commitedAccessList    bool
	commitedAccessList6   bool
	commitedAsPathList    bool
	commitedCommunityList bool
	commitedPrefixList    bool
	commitedPrefixList6   bool
	commitedRouteMap      bool
	configRunning         *quaggaConfigStateNode
	configCandidate       *quaggaConfigStateNode
)

type RuleSlice []string

func (rules RuleSlice) Len() int {
	return len(rules)
}

func (rules RuleSlice) Less(i, j int) bool {
	stri := rules[i]
	strj := rules[j]
	inti, erri := strconv.Atoi(stri)
	intj, errj := strconv.Atoi(strj)
	if erri != nil || errj != nil {
		return false
	}
	return inti < intj
}

func (rules RuleSlice) Swap(i, j int) {
	rules[i], rules[j] = rules[j], rules[i]
}

func quaggaRuleSort(rules []string) {
	sort.Sort(RuleSlice(rules))
}

func quaggaVtysh(cmds ...string) *string {
	cmdArgs := []string{}
	cmdArgs = append(cmdArgs, "vtysh")
	for _, c := range cmds {
		cmdArgs = append(cmdArgs, "-c", c)
		fmt.Println("quaggaVtysh: ", c)
	}
	out, err := exec.Command("sudo", cmdArgs...).CombinedOutput()
	if err != nil {
		s := fmt.Sprint(err)
		fmt.Println("quaggaVtysh: ", s)
		return &s
	}
	s := string(out)
	fmt.Println("quaggaVtysh: ", s)
	return &s
}

type quaggaConfigStateNode struct {
	parent   *quaggaConfigStateNode
	children map[string]*quaggaConfigStateNode
}

func makeQuaggaConfigStateNode(parent *quaggaConfigStateNode) *quaggaConfigStateNode {
	n := &quaggaConfigStateNode{}
	n.parent = parent
	n.children = make(map[string]*quaggaConfigStateNode)
	return n
}

func (config *quaggaConfigStateNode) quaggaConfigStateSet(path []string) {
	fmt.Println(path)
	c := config
	for _, n := range path {
		ct, ok := c.children[n]
		if !ok {
			ct = makeQuaggaConfigStateNode(c)
			c.children[n] = ct
		}
		c = ct
	}
}

func (config *quaggaConfigStateNode) quaggaConfigStateDelete(path []string) {
	fmt.Println(path)
	c := config
	l := ""
	for _, n := range path {
		ct, ok := c.children[n]
		if !ok {
			return
		}
		c = ct
		l = n
	}
	if c.parent == nil || l == "" {
		return
	}
	delete(c.parent.children, l)
}

func quaggaConfigStateUpdate(command int, path []string) {
	config := configRunning
	if validating {
		config = configCandidate
	}
	switch command {
	case cmd.Set:
		config.quaggaConfigStateSet(path)
	case cmd.Delete:
		config.quaggaConfigStateDelete(path)
	}
}

func quaggaConfigStateSync1(src, dst *quaggaConfigStateNode) {
	for k, v := range src.children {
		configNew := makeQuaggaConfigStateNode(dst)
		dst.children[k] = configNew
		quaggaConfigStateSync1(v, configNew)
	}
}

func quaggaConfigStateSync() {
	configNew := makeQuaggaConfigStateNode(nil)
	quaggaConfigStateSync1(configRunning, configNew)
	configCandidate = configNew
}

func quaggaConfigStateDiff1(l, r *quaggaConfigStateNode) bool {
	if l == nil && r == nil {
		return false
	}
	if l != nil && r == nil {
		return true
	}
	if l == nil && r != nil {
		return true
	}
	if len(l.children) != len(r.children) {
		return true
	}
	for label, lc := range l.children {
		rc, ok := r.children[label]
		if !ok {
			return true
		}
		if quaggaConfigStateDiff1(lc, rc) {
			return true
		}
	}
	return false
}

func quaggaConfigStateDiff(path []string) bool {
	candidate := configCandidate.lookup(path)
	running := configRunning.lookup(path)
	return quaggaConfigStateDiff1(candidate, running)
}

func quaggaConfigStateDump1(l, r *quaggaConfigStateNode, indent int) {
	labels := make([]string, 0)
	if l != nil {
		for label, _ := range l.children {
			labels = append(labels, label)
		}
	}
	if r != nil {
		for label, _ := range r.children {
			if l != nil {
				if _, ok := l.children[label]; ok {
					continue
				}
			}
			labels = append(labels, label)
		}
	}
	sort.Strings(labels)
	for _, label := range labels {
		le := false
		if l != nil {
			if _, ok := l.children[label]; ok {
				le = true
			}
		}
		re := false
		if r != nil {
			if _, ok := r.children[label]; ok {
				re = true
			}
		}
		var lc *quaggaConfigStateNode
		var rc *quaggaConfigStateNode
		s := ""
		if le && re {
			s += "  "
			lc = l.children[label]
			rc = r.children[label]
		} else if le {
			s += "- "
			lc = l.children[label]
		} else if re {
			s += "+ "
			rc = r.children[label]
		} else {
			s += "? "
		}
		for i := 0; i < indent; i++ {
			s += "    "
		}
		s += label
		fmt.Println(s)
		quaggaConfigStateDump1(lc, rc, indent+1)
	}
}

func quaggaConfigStateDump() {
	quaggaConfigStateDump1(configRunning, configCandidate, 0)
}

func (config *quaggaConfigStateNode) lookup(path []string) *quaggaConfigStateNode {
	if len(path) == 0 {
		return nil
	}
	c := config
	for _, n := range path {
		ct, ok := c.children[n]
		if !ok {
			return nil
		}
		c = ct
	}
	return c
}

func (config *quaggaConfigStateNode) value(path []string) *string {
	c := config.lookup(path)
	if c == nil {
		return nil
	}
	if len(c.children) == 1 {
		for label, _ := range c.children {
			return &label
		}
	} else if len(c.children) > 1 {
		fmt.Println("*** BUG? quaggaConfigStateValue multipul values ", path)
		for label, _ := range c.children {
			fmt.Println("*** label: ", label)
		}
	}
	return nil
}

func (config *quaggaConfigStateNode) values(path []string) []string {
	values := make([]string, 0)
	c := config.lookup(path)
	if c == nil {
		return values
	}
	for label, _ := range c.children {
		values = append(values, label)
	}
	return values
}

func quaggaConfigValid(f io.Writer) bool {
	valid := true

	if !quaggaConfigValidInterfacesOspf(f) {
		valid = false
	}

	if !quaggaConfigValidPolicy(f) {
		valid = false
	}

	if !quaggaConfigValidBgp(f) {
		valid = false
	}

	if !quaggaConfigValidOspf(f) {
		valid = false
	}

	return valid
}

func quaggaConfig(command int, path []string) {
	switch command {
	case cmd.Set:
		fmt.Println("[cmd] add", path)
	case cmd.Delete:
		fmt.Println("[cmd] del", path)
	}
	if !validating {
		ret, fn, args, _ := configParser.ParseCmd(path)
		if ret == cmd.ParseSuccess {
			fn.(func(int, cmd.Args) int)(command, args)
		}
	}
	quaggaConfigStateUpdate(command, path)
}

func Main() {
	configRunning = makeQuaggaConfigStateNode(nil)
	configCandidate = makeQuaggaConfigStateNode(nil)
	initConfig()
	initGrpc()
}
