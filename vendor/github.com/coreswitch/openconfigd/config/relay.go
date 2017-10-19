// Copyright 2017 OpenConfigd Project.
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
	"os/exec"
	"sync"
	"time"

	"golang.org/x/net/context"
)

type RelayInstance struct {
	IfName   string
	Vrf      string
	Group    string
	ExitFunc func()
}

type RelayGroup struct {
	Server []string
}

var (
	RelayInstanceMap = map[string]*RelayInstance{}
	RelayGroupMap    = map[string]*RelayGroup{}
)

func RelayExitFunc() {
	for _, instance := range RelayInstanceMap {
		if instance.ExitFunc != nil {
			instance.ExitFunc()
			instance.ExitFunc = nil
		}
	}
}

func RelayGroupExec(instance *RelayInstance, group *RelayGroup, vrf string) func() {
	fmt.Println("RelayGroupExec")
	binary, err := exec.LookPath("dhcrelay")
	if err != nil {
		return nil
	}

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})
	var wg sync.WaitGroup

	pidFileName := fmt.Sprintf("/var/run/dhcrelay-%s.pid", instance.IfName)

	args := []string{"-d", "-4", "-pf", pidFileName, "-i", instance.IfName}
	args = append(args, group.Server...)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			cmd := exec.CommandContext(ctx, binary, args...)

			env := os.Environ()
			if vrf != "" {
				env = append(env, fmt.Sprintf("VRF=%s", vrf))
				env = append(env, "LD_PRELOAD=/usr/bin/vrf_socket.so")
			}
			fmt.Println("dhcrelay cmd:", cmd)
			cmd.Env = env
			startErr := cmd.Start()
			fmt.Println("dhcrelay", vrf, "cmd.Start()", startErr)
			err := cmd.Wait()
			fmt.Println("dhcrelay", vrf, "cmd.Wait():", err)

			retryTimer := time.NewTimer(time.Second * 1)
			select {
			case <-retryTimer.C:
				fmt.Println("dhcrelay retryTimer expired")
			case <-done:
				retryTimer.Stop()
				fmt.Println("dhcrelay retryTimer stop")
				return
			}
		}
	}()

	return func() {
		close(done)
		cancel()
		os.Remove(pidFileName)
		fmt.Println("WaitGroup wait")
		wg.Wait()
		fmt.Println("WaitGroup done")
	}
}

func RelayGroupUpdate(dhcp *Dhcp) {
	fmt.Println("[dhcp]RelayGroupUpdate")
	RelayExitFunc()

	RelayGroupMap = map[string]*RelayGroup{}
	for _, g := range dhcp.Relay.ServerGroupList {
		group := &RelayGroup{}
		for _, addr := range g.ServerAddressList {
			group.Server = append(group.Server, addr.Address)
		}
		RelayGroupMap[g.ServerGroupName] = group
	}
	for key, instance := range RelayInstanceMap {
		fmt.Println("RelayGroupUpdate: looking for instance", key)
		group := RelayGroupMap[instance.Group]
		if group != nil {
			instance.ExitFunc = RelayGroupExec(instance, group, instance.Vrf)
		}
	}
}

func RelayAdd(ifName, group string) {
	fmt.Println("RelayAdd", ifName, group)
	vrf := ConfigLookupVrf(ifName)
	fmt.Println("Vrf:", vrf)

	exists := RelayInstanceMap[ifName]
	if exists != nil && exists.ExitFunc != nil {
		exists.ExitFunc()
		exists.ExitFunc = nil
	}

	instance := &RelayInstance{
		IfName: ifName,
		Vrf:    vrf,
		Group:  group,
	}
	RelayInstanceMap[ifName] = instance

	g := RelayGroupMap[group]
	if g != nil {
		instance.ExitFunc = RelayGroupExec(instance, g, vrf)
	}
}

func RelayDelete(ifName, group string) {
	fmt.Println("RelayDelete", ifName, group)
	instance := RelayInstanceMap[ifName]
	if instance == nil {
		return
	}
	if instance.ExitFunc != nil {
		instance.ExitFunc()
		instance.ExitFunc = nil
	}
	delete(RelayInstanceMap, ifName)
}

func RelayApi(set bool, Args []interface{}) {
	if len(Args) != 2 {
		return
	}
	ifName := Args[0].(string)
	group := Args[1].(string)
	if set {
		RelayAdd(ifName, group)
	} else {
		RelayDelete(ifName, group)
	}
}
