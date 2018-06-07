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
	"runtime"
	"strconv"
	"strings"

	"github.com/coreswitch/cmd"
	"github.com/coreswitch/component"
	"github.com/coreswitch/openconfigd/quagga"
	"github.com/coreswitch/process"
)

var TopCmd *cmd.Cmd
var Parser *cmd.Node

func showVersion(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = "Developer Preview version of openconfigd\n"
	return
}

func showProcess(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = process.ProcessListShow()
	return
}

func startProcess(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	if len(Args) == 1 {
		arg := Args[0]
		num, err := strconv.Atoi(arg)
		if err == nil {
			process.ProcessStart(num)
		}
	}
	instStr = ""
	return
}

func stopProcess(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	if len(Args) == 1 {
		arg := Args[0]
		num, err := strconv.Atoi(arg)
		if err == nil {
			process.ProcessStop(num)
		}
	}
	instStr = ""
	return
}

func showNumGoroutine(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = fmt.Sprintf(`Number of goroutine: %v`, runtime.NumGoroutine())
	return
}

func showIpBgp(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "gobgp global"
	return
}

func showIpBgpRoute(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "gobgp global rib"
	return
}

func showIpBgpNeighbor(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "gobgp neighbor"
	return
}

func showQuaggaPassword(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = fmt.Sprintf("quagga password is: %s", quagga.GetPasswd())
	return
}

func enableFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "CLI_PRIVILEGE=15;_cli_refresh"
	return
}

func disableFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "CLI_PRIVILEGE=1;_cli_refresh"
	return
}

func exitFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "exit"
	return
}

func helpFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "echo help function"
	return
}

func logoutFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "exit"
	return
}

func quitFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "exit"
	return
}

func configureTerminal(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "CLI_MODE=config;CLI_MODE_STR=Configure;CLI_MODE_PROMPT=\"(config)\";_cli_refresh"
	return
}

func configure(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "CLI_MODE=configure;CLI_MODE_STR=Configure;CLI_PRIVILEGE=15;_cli_refresh"
	return
}

func configureDiscardFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	diff := ConfigDiscard()
	if diff {
		instStr = "All changes has been discarded."
	} else {
		instStr = "No changes has been discarded."
	}
	return
}

func configureExitFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessExec
	instStr = "CLI_MODE=exec;CLI_PRIVILEGE=1;_cli_refresh"
	return
}

func configureShowFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	//instStr = TopCandidate.ConfigString()
	instStr = Compare()
	return
}

func configureJsonFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = JsonMarshal()
	return
}

func configureCommitFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	err := Commit()
	if err != nil {
		instStr = err.Error()
	}
	return
}

func configureUpFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = ""
	return
}

func configureEditFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = ""
	return
}

func configureCompareFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = CompareCommand()
	return
}

func configureCommandsFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = Commands()
	return
}

func configureRollbackFunc(Args []string) (inst int, instStr string) {
	inst = CliSuccessShow
	instStr = ""
	return
}

type Command struct {
	set  bool
	cmds []string
}

func ExecCmd(c *Command) {
	ret, fn, args, _ := Parser.ParseCmd(c.cmds)
	if ret == cmd.ParseSuccess {
		if cb, ok := fn.(func(bool, []interface{})); ok {
			cb(c.set, args)
		}
	}
}

func UnQuote(args []string) {
	for pos, arg := range args {
		arg, err := strconv.Unquote(arg)
		if err == nil {
			args[pos] = arg
		}
	}
}

func NewCommand(line string) *Command {
	if line == "" {
		return nil
	}
	set := false
	switch line[0] {
	case '+':
		set = true
	case '-':
		set = false
	default:
		return nil
	}
	c := &Command{
		set:  set,
		cmds: strings.Fields(line),
	}
	c.cmds = c.cmds[1:]
	UnQuote(c.cmds)

	return c
}

// CLI component.
type CliComponent struct{}

// CLI component start method.
func (this *CliComponent) Start() component.Component {
	Cmd := cmd.NewCmd()
	Parser = cmd.NewParser()

	// Operational mode.
	mode := Cmd.InstallMode("exec", "Exec", "")
	mode.InstallLine("exit", exitFunc,
		&cmd.Param{Helps: []string{"End current mode and down to previous mode"}})
	mode.InstallLine("help", helpFunc,
		&cmd.Param{Helps: []string{"Description of the interactive help system"}})
	mode.InstallLine("logout", logoutFunc,
		&cmd.Param{Helps: []string{"Exit from EXEC"}})
	mode.InstallLine("quit", quitFunc,
		&cmd.Param{Helps: []string{"End current mode and down to previous mode"}})
	mode.InstallLine("show version", showVersion,
		&cmd.Param{Helps: []string{"Show running system information", "Display openconfigd version"}})
	mode.InstallLine("show ip bgp", showIpBgp,
		&cmd.Param{Helps: []string{"Show running system information", "IP", "BGP"}})
	mode.InstallLine("show ip bgp route", showIpBgpRoute,
		&cmd.Param{Helps: []string{"Show running system information", "IP", "BGP", "Route"}})
	mode.InstallLine("show ip bgp neighbors", showIpBgpNeighbor,
		&cmd.Param{Helps: []string{"Show running system information", "IP", "BGP", "Neighbor"}})
	mode.InstallLine("configure", configure,
		&cmd.Param{Helps: []string{"Manipulate software configuration information"}})
	mode.InstallLine("show system etcd", showSystemEtcd,
		&cmd.Param{Helps: []string{"", "System Information", "etcd endpoints and status"}})
	mode.InstallLine("show process", showProcess,
		&cmd.Param{Helps: []string{"", "Process Information"}})
	mode.InstallLine("show numgoroutine", showNumGoroutine,
		&cmd.Param{Helps: []string{"Show running system information", "Number of goroutine"}})

	opNode := mode.Parser

	cmd.DynamicFunc = DynamicCompletion

	// Configure mode.
	mode = Cmd.InstallMode("configure", "Configure", "")
	mode.InstallLine("help", helpFunc,
		&cmd.Param{Helps: []string{"Provide hellp information"}})
	mode.InstallHook("set", YParseSet,
		&cmd.Param{Helps: []string{"Set a parameter"}})
	mode.InstallHook("delete", YParseDelete,
		&cmd.Param{Helps: []string{"Delete a parameter"}})
	mode.InstallLine("discard", configureDiscardFunc,
		&cmd.Param{Helps: []string{"Discard candidate configuration"}})
	mode.InstallLine("exit", configureExitFunc,
		&cmd.Param{Helps: []string{"Exit from this level"}})
	mode.InstallLine("quit", configureExitFunc,
		&cmd.Param{Helps: []string{"Quit from this level"}})
	mode.InstallLine("show", configureShowFunc,
		&cmd.Param{Helps: []string{"Show a parameter"}})
	mode.InstallLine("json", configureJsonFunc,
		&cmd.Param{Helps: []string{"Show a JSON format configuration"}})

	mode.InstallLine("etcd json", configureEtcdJsonFunc,
		&cmd.Param{Helps: []string{"Show a etcd JSON configuration"}})
	mode.InstallLine("etcd bgp-body", configureEtcdBodyFunc,
		&cmd.Param{Helps: []string{"Show a etcd JSON configuration"}})
	mode.InstallLine("etcd bgp-version", configureEtcdVersionFunc,
		&cmd.Param{Helps: []string{"Show a etcd JSON configuration"}})
	mode.InstallLine("etcd bgp-config", configureEtcdBgpConfigFunc,
		&cmd.Param{Helps: []string{"Show a etcd BGP configuration"}})

	mode.InstallLine("etcd vrf-body", configureEtcdBodyFunc2,
		&cmd.Param{Helps: []string{"Show a etcd JSON configuration"}})
	mode.InstallLine("etcd vrf-version", configureEtcdVersionFunc2,
		&cmd.Param{Helps: []string{"Show a etcd JSON configuration"}})

	mode.InstallLine("etcd bgp-wan-body", configureEtcdBgpWanBodyFunc,
		&cmd.Param{Helps: []string{"Show a etcd JSON configuration"}})

	mode.InstallLine("clear gobgp", GobgpClearApi,
		&cmd.Param{Helps: []string{"Clear", "GoBGP configuration"}})
	mode.InstallLine("reset gobgp", GobgpResetApi,
		&cmd.Param{Helps: []string{"Reset", "GoBGP configuration"}})

	mode.InstallLine("commit", configureCommitFunc,
		&cmd.Param{Helps: []string{"Commit current set of changes"}})
	mode.InstallLine("up", configureUpFunc,
		&cmd.Param{Helps: []string{"Exit one level of configuration"}})
	mode.InstallLine("edit", configureEditFunc,
		&cmd.Param{Helps: []string{"Edit a sub-element"}})
	mode.InstallLine("compare", configureCompareFunc,
		&cmd.Param{Helps: []string{"Compare configuration tree"}})
	mode.InstallLine("commands", configureCommandsFunc,
		&cmd.Param{Helps: []string{"Show configuration commands"}})
	mode.InstallLine("run", nil,
		&cmd.Param{Helps: []string{"Run an operational-mode command"}})

	mode.InstallLine("rollback", configureRollbackFunc,
		&cmd.Param{Helps: []string{"Rollback configuration"}})
	mode.InstallLine("rollback :local:rollback", configureRollbackFunc,
		&cmd.Param{Helps: []string{"Rollback configuration"}})

	mode.InstallLine("start process WORD", startProcess,
		&cmd.Param{Helps: []string{"Start", "Process"}})
	mode.InstallLine("stop process WORD", stopProcess,
		&cmd.Param{Helps: []string{"Stop", "Process"}})
	mode.InstallLine("show quagga password", showQuaggaPassword,
		&cmd.Param{Helps: []string{"Show running system information", "quagga infromation", "Show password"}})

	// Link "run" command to operational node.
	run := mode.Parser.Lookup("run")
	run.LinkNodes(opNode)

	Parser.InstallLine("system host-name WORD", HostnameApi)
	Parser.InstallLine("system etcd endpoints WORD", EtcdEndpointsApi)
	Parser.InstallLine("system etcd path WORD", EtcdPathApi)
	Parser.InstallLine("system gobgp grpcendpoint WORD", ConfigureGobgpGrpcEndpointApi)
	Parser.InstallLine("interfaces interface WORD dhcp-relay-group WORD", RelayApi)

	TopCmd = Cmd

	return this
}

func (this *CliComponent) Stop() component.Component {
	return this
}
