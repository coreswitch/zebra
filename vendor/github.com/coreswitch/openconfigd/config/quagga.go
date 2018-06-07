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
	"io/ioutil"
	"os"
	"regexp"

	"github.com/coreswitch/netutil"
	"github.com/coreswitch/openconfigd/quagga"
	"github.com/coreswitch/process"
)

var QuaggaProc = map[int]process.ProcessSlice{}

func QuaggaExit() {
	for _, vrfProcList := range QuaggaProc {
		for _, proc := range vrfProcList {
			process.ProcessUnregister(proc)
			err := os.Remove(proc.File)
			if err != nil {
				fmt.Println("Delete pid file:", err)
			}
		}
	}
}

func QuaggaDelete(vrfId int) {
	fmt.Println("[quagga]delete: vrfId, Processes: %+v", vrfId, QuaggaProc[vrfId])
	_, ok := QuaggaProc[vrfId]
	if ok {
		for _, proc := range QuaggaProc[vrfId] {
			if proc != nil {
				process.ProcessUnregister(proc)
				err := os.Remove(proc.File)
				if err != nil {
					fmt.Println("Delete pid file:", err)
				}
			}
		}
		QuaggaProc[vrfId] = QuaggaProc[vrfId][:0]
		delete(QuaggaProc, vrfId)
	}
}

func QuaggaExec(vrfId int, interfaceName string, configStr string) {
	re := regexp.MustCompile(`password 8 ([0-9A-Za-z]{13})`)
	configStr = re.ReplaceAllString(configStr, "password 8 "+quagga.GetHash())

	fmt.Println("[quagga]config: vrfId", vrfId, "Interface Name: ", interfaceName, configStr)
	configFileName := fmt.Sprintf("/etc/quagga/bgpd-vrf%d-%s.conf", vrfId, interfaceName)
	zapiSocketName := fmt.Sprintf("/var/run/zserv-vrf%d.api", vrfId)
	pidFileName := fmt.Sprintf("/var/run/bgpd-vrf%d-%s.pid", vrfId, interfaceName)

	err := ioutil.WriteFile(configFileName, []byte(configStr), 0644)
	if err != nil {
		fmt.Println("[quagga]WriteFile err:", err)
	}

	args := []string{
		"-u", "root",
		"-g", "root",
		"-p", "0",
		"-f", configFileName,
		"-z", zapiSocketName,
		//"-l", LocalAddrLookup(interfaceName), -l implies -n and does not install kernel route
		"-i", pidFileName,
	}

	proc := process.NewProcess("bgpd", args...)
	proc.Vrf = fmt.Sprintf("vrf%d", vrfId)
	proc.File = pidFileName
	proc.StartTimer = 3
	_, ok := QuaggaProc[vrfId]
	if ok {
		QuaggaProc[vrfId] = append(QuaggaProc[vrfId], proc)
	} else {
		QuaggaProc[vrfId] = process.ProcessSlice{}
		QuaggaProc[vrfId] = append(QuaggaProc[vrfId], proc)
	}
	process.ProcessRegister(proc)
}

func LocalAddrLookup(ifName string) string {
	addrConfig := configActive.LookupByPath([]string{"interfaces", "interface", ifName, "ipv4", "address"})
	if addrConfig != nil && len(addrConfig.Keys) > 0 {
		prefix, _ := netutil.ParsePrefix(addrConfig.Keys[0].Name)
		return prefix.IP.String()
	}
	return ""
}

func QuaggaVrfSync(vrfId int, cfg *VrfsConfig) {
	// if len(cfg.Bgp) == 0 {
	// 	return
	// }
	fmt.Println("QuaggaVrfSync", vrfId)
	QuaggaDelete(vrfId)
	for _, bgpConfig := range cfg.Bgp {
		QuaggaExec(vrfId, bgpConfig.Interface, bgpConfig.CiscoConfig)
	}
}

func QuaggaVrfDelete(vrfId int) {
	QuaggaDelete(vrfId)
}

func NexthopWalkerUpdate() {
	numProc := len(QuaggaProc) + len(OspfProcessMap)
	if numProc == 0 {
		ExecLine(fmt.Sprintf("delete routing-options nexthop-walker"))
	} else {
		ExecLine(fmt.Sprintf("set routing-options nexthop-walker"))
	}
	Commit()
}
