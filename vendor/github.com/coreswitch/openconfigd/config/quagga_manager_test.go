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

package config_test

import (
	"github.com/coreswitch/openconfigd/config"
	"strconv"
	"testing"
)

// TODO: Mock out Process Spawning
func TestProcessQuaggaConfigSync(t *testing.T) {

	testInstanceName := "local"
	completeNotify := make(chan int)
	config.QuaggaConfigDir = "/tmp/"
	for i := 0; i < 10; i++ {
		testJsonString := `{"lan-1": [{"quagga-config": "test1", "routing-protocol": "bgp"}, {"quagga-config": "test2", "routing-protocol": "rip"}],
		"lan-2": [{"quagga-config": "test1", "routing-protocol": "bgp"}, {"quagga-config": "test2", "routing-protocol": "rip"}]}`
		testVrfId := i
		go func(loopcount int, jsonstring string, vrfid int, instName string) {
			config.QuaggaConfigSync(jsonstring, vrfid, instName)
			completeNotify <- loopcount
		}(i, testJsonString, testVrfId, testInstanceName)
	}

	testInstanceName = "remote"
	for i := 0; i < 10; i++ {
		testJsonString := `{"lan-1": [{"quagga-config": "test1", "routing-protocol": "bgp"}, {"quagga-config": "test2", "routing-protocol": "rip"}],
		"lan-2": [{"quagga-config": "test1", "routing-protocol": "bgp"}, {"quagga-config": "test2", "routing-protocol": "rip"}]}`
		testVrfId := i
		go func(loopcount int, jsonstring string, vrfid int, instName string) {
			config.QuaggaConfigSync(jsonstring, vrfid, instName)
			completeNotify <- loopcount
		}(i, testJsonString, testVrfId, testInstanceName)
	}

	spawnCount := 0
	for range completeNotify {
		spawnCount++
		if spawnCount >= 20 {
			break
		}
	}

	localManagerInstance := config.GetInstanceManager("local")
	if localManagerInstance == nil {
		t.Error("Local Instance Manger is nil")
	}
	if len(localManagerInstance.QuaggaProcMap) != 10 {
		t.Errorf("Expected 10 entry in cache. But found %d", len(localManagerInstance.QuaggaProcMap))
	}
	for vrfId, vrfProcMap := range localManagerInstance.QuaggaProcMap {
		for _, interfaceProcList := range vrfProcMap {
			for _, proc := range interfaceProcList {
				if proc.Vrf != "vrf"+strconv.Itoa(vrfId) {
					t.Errorf("Expected vrf %s for process. But got %s", "vrf"+strconv.Itoa(vrfId), proc.Vrf)
				}
			}
		}
	}

	remoteManagerInstance := config.GetInstanceManager("remote")
	if remoteManagerInstance == nil {
		t.Error("Remote Instance Manger is nil")
	}
	if len(remoteManagerInstance.QuaggaProcMap) != 10 {
		t.Errorf("Expected 10 entry in cache. But found %d", len(remoteManagerInstance.QuaggaProcMap))
	}
	for vrfId, vrfProcMap := range remoteManagerInstance.QuaggaProcMap {
		for _, interfaceProcList := range vrfProcMap {
			for _, proc := range interfaceProcList {
				if proc.Vrf != "vrf"+strconv.Itoa(vrfId) {
					t.Errorf("Expected vrf %s for process. But got %s", "vrf"+strconv.Itoa(vrfId), proc.Vrf)
				}
			}
		}
	}
}
