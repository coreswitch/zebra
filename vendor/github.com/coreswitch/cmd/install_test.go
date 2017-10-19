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
	"fmt"
	"net"
	"testing"
)

func Hostname(Cmd int, Args Args) int {
	fmt.Println("Hostname arg len", len(Args))
	hostname := Args[0].(string)
	fmt.Println("Arg", hostname)
	return 0
}

func RouterId(Cmd int, Args Args) int {
	id := Args[0].(net.IP)
	fmt.Println("RouterId arg len", len(Args), id)
	return 0
}

func TestInstall(t *testing.T) {
	p := NewParser()
	if p == nil {
		t.Errorf("NewParser() failed")
	}
	p.InstallLine("system host-name WORD", Hostname)
	p.InstallLine("system host-name :fea:ifname", Hostname)
	p.InstallLine("routing-option router-id A.B.C.D", RouterId)
	p.InstallLine("show (ip) route", RouterId)
	p.Dump()

	ret, fn, args, _ := p.ParseCmd([]string{"system", "host-name", "newhostname"})
	if ret == ParseSuccess {
		callback := fn.(func(int, Args) int)
		callback(0, args)
	} else {
		fmt.Println("Parse failed")
	}

	ret, fn, args, _ = p.ParseCmd([]string{"routing-option", "router-id", "1.1.1.1"})
	if ret == ParseSuccess {
		callback := fn.(func(int, Args) int)
		callback(0, args)
	} else {
		fmt.Println("Parse failed")
	}
}
