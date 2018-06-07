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

	"encoding/json"
	"strings"
	"sync"

	"github.com/coreswitch/process"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
)

var QuaggaConfigDir = "/etc/quagga/"

type OpenBgpConfig struct {
	QuaggaInstanceConfig string `mapstructure:"quagga-config" json:"quagga-config,omitempty"`
	RoutingProtocol      string `mapstructure:"routing-protocol" json:"routing-protocol,omitempty"`
}

type QuaggaManager interface {
	SetConfig()
	DeleteConfig()
	Destroy()
}

type QuaggaInstanceConfig struct {
	QuaggaInstanceConfig string `mapstructure:"quagga-config" json:"quagga-config,omitempty"`
	RoutingProtocol      string `mapstructure:"routing-protocol" json:"routing-protocol,omitempty"`
}

type QuaggaInterfaceConfigs map[string][]QuaggaInstanceConfig

type QuaggaInstanceManager struct {
	QuaggaProcMap     map[int]map[string]process.ProcessSlice
	QuaggaConfigCache map[int]QuaggaInterfaceConfigs
	LocalCacheLock    *sync.Mutex
	QuaggaConfigDir   string
}

func NewQuaggaManager() *QuaggaInstanceManager {
	qim := &QuaggaInstanceManager{
		QuaggaProcMap:     map[int]map[string]process.ProcessSlice{},
		QuaggaConfigCache: map[int]QuaggaInterfaceConfigs{},
		LocalCacheLock:    new(sync.Mutex),
		QuaggaConfigDir:   QuaggaConfigDir,
	}
	return qim
}

func (qim QuaggaInstanceManager) LockLocalCache() {
	qim.LocalCacheLock.Lock()
}

func (qim QuaggaInstanceManager) UnLockLocalCache() {
	qim.LocalCacheLock.Unlock()
}

func (qim QuaggaInstanceManager) QuaggaCleanup() {
	qim.LocalCacheLock.Lock()
	defer qim.LocalCacheLock.Unlock()
	for vrfId, vrfProcMap := range qim.QuaggaProcMap {
		for interfaceName, interfaceProcList := range vrfProcMap {
			for _, proc := range interfaceProcList {
				process.ProcessUnregister(proc)
				err := os.Remove(proc.File)
				if err != nil {
					log.WithFields(log.Fields{
						"error": err,
					}).Error("Delete pid file Failed")
				}
			}
			qim.QuaggaProcMap[vrfId][interfaceName] = qim.QuaggaProcMap[vrfId][interfaceName][:0]
			delete(qim.QuaggaProcMap[vrfId], interfaceName)
		}
		delete(qim.QuaggaProcMap, vrfId)
	}
}

func (qim QuaggaInstanceManager) ProcessConfigDelete(vrfId int, quaggaConfigMap QuaggaInterfaceConfigs) {
	log.WithFields(log.Fields{
		"vrfId": vrfId,
	}).Debug("Process quagga config delete")
	_, ok := qim.QuaggaProcMap[vrfId]
	if ok {
		for interfaceName, _ := range quaggaConfigMap {
			for _, proc := range qim.QuaggaProcMap[vrfId][interfaceName] {
				if proc != nil {
					process.ProcessUnregister(proc)
					err := os.Remove(proc.File)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err,
						}).Error("Delete pid file Failed")
					}
				}
			}
			delete(qim.QuaggaProcMap[vrfId], interfaceName)
		}
	}
}

func (qim QuaggaInstanceManager) ProcessVrfQuaggaConfigDelete(vrfId int) {
	qim.LocalCacheLock.Lock()
	defer qim.LocalCacheLock.Unlock()
	log.WithFields(log.Fields{
		"vrfId": vrfId,
	}).Debug("Quagga vrf config delete")
	_, ok := qim.QuaggaProcMap[vrfId]
	if ok {
		for interfaceName, interfaceProcList := range qim.QuaggaProcMap[vrfId] {
			for _, proc := range interfaceProcList {
				if proc != nil {
					process.ProcessUnregister(proc)
					err := os.Remove(proc.File)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err,
						}).Error("Delete pid file Failed")
					}
					configFileName := fmt.Sprintf("%s/bgpd-vrf%d-%s.conf", qim.QuaggaConfigDir, vrfId, interfaceName)
					err = os.Remove(configFileName)
					if err != nil {
						log.WithFields(log.Fields{
							"error": err,
						}).Error("Deleting quagga bgp config file Failed")
					}
				}
			}
			qim.QuaggaProcMap[vrfId][interfaceName] = qim.QuaggaProcMap[vrfId][interfaceName][:0]
			delete(qim.QuaggaProcMap[vrfId], interfaceName)
		}
		delete(qim.QuaggaProcMap, vrfId)
	}
}

func (qim QuaggaInstanceManager) ProcessQuaggaConfigAdd(vrfId int, quaggaConfigMap QuaggaInterfaceConfigs) {
	log.WithFields(log.Fields{
		"vrfId": vrfId,
	}).Debug("QuaggaConfig Add")
	for interfaceName, quaggaConfig := range quaggaConfigMap {
		for _, quaggaInstanceConfig := range quaggaConfig {
			qim.SpawnQuagga(vrfId, interfaceName, quaggaInstanceConfig.QuaggaInstanceConfig)
		}
	}
}

func (qim QuaggaInstanceManager) SpawnQuagga(vrfId int, interfaceName string, configStr string) {
	log.WithFields(log.Fields{
		"vrfId":         vrfId,
		"interfaceName": interfaceName,
		"quaggaconfig":  configStr,
	}).Debug("Spawn Quagga")
	configFileName := fmt.Sprintf("%s/bgpd-vrf%d-%s.conf", qim.QuaggaConfigDir, vrfId, interfaceName)
	zapiSocketName := fmt.Sprintf("/var/run/zserv-vrf%d.api", vrfId)
	pidFileName := fmt.Sprintf("/var/run/bgpd-vrf%d-%s.pid", vrfId, interfaceName)

	err := ioutil.WriteFile(configFileName, []byte(configStr), 0644)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Quagga config file write error")
		return
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
	proc.StartTimer = 0
	_, ok := qim.QuaggaProcMap[vrfId][interfaceName]
	if !ok {
		qim.QuaggaProcMap[vrfId] = make(map[string]process.ProcessSlice)
	}
	qim.QuaggaProcMap[vrfId][interfaceName] = append(qim.QuaggaProcMap[vrfId][interfaceName], proc)
	process.ProcessRegister(proc)
}

func (qim QuaggaInstanceManager) QuaggaBgpConfigDiff(oldQuaggaConfig QuaggaInterfaceConfigs, newQuaggaConfig QuaggaInterfaceConfigs) (QuaggaInterfaceConfigs, QuaggaInterfaceConfigs, QuaggaInterfaceConfigs) {
	addedConfigs := make(QuaggaInterfaceConfigs)
	removedConfigs := make(QuaggaInterfaceConfigs)
	modifiedConfigs := make(QuaggaInterfaceConfigs)

	if &oldQuaggaConfig == &newQuaggaConfig {
		return addedConfigs, removedConfigs, modifiedConfigs
	}

	for interfaceName, quaggaConfig := range oldQuaggaConfig {
		newQuaggaConfig, ok := newQuaggaConfig[interfaceName]
		if ok {
			if len(quaggaConfig) != len(newQuaggaConfig) {
				//TODO:  In case of modified, it doesnt make sense to also assign the value
				modifiedConfigs[interfaceName] = quaggaConfig
			} else {
				for index, instanceConfig := range quaggaConfig {
					// Just do a value by value diff for now. We wont change the order today
					if strings.Compare(instanceConfig.QuaggaInstanceConfig, newQuaggaConfig[index].QuaggaInstanceConfig) != 0 {
						modifiedConfigs[interfaceName] = quaggaConfig
						break
					}
				}
			}
		} else {
			removedConfigs[interfaceName] = quaggaConfig
		}
	}
	for interfaceName, quaggaConfig := range newQuaggaConfig {
		_, ok := oldQuaggaConfig[interfaceName]
		if !ok {
			addedConfigs[interfaceName] = quaggaConfig
		}
	}
	return addedConfigs, removedConfigs, modifiedConfigs
}

func (qim QuaggaInstanceManager) QuaggaConfigSync(jsonStr string, vrfId int) {
	var jsonIntf interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonIntf)
	if err != nil {
		log.WithFields(log.Fields{
			"json":  jsonStr,
			"error": err,
		}).Error("QuaggaConfigSync:json.Unmarshal()")
		return
	}

	var quaggaConfigs QuaggaInterfaceConfigs
	err = mapstructure.Decode(jsonIntf, &quaggaConfigs)
	if err != nil {
		log.WithFields(log.Fields{
			"json-intf": jsonIntf,
			"error":     err,
		}).Error("QuaggaConfigSync:mapstructure.Decode()")
		return
	}

	log.WithFields(log.Fields{
		"vrfId": vrfId,
	}).Debug("QuaggaConfigSync for vrf")

	// TODO: Make locks more granular
	qim.LocalCacheLock.Lock()
	fmt.Printf("Got lock qim: %+v\n", qim)
	defer fmt.Println("Deferred unlocking")
	defer qim.LocalCacheLock.Unlock()
	_, ok := qim.QuaggaConfigCache[vrfId]

	if ok {
		// Handle config update for the vrf
		added, removed, modified := qim.QuaggaBgpConfigDiff(qim.QuaggaConfigCache[vrfId], quaggaConfigs)
		qim.ProcessConfigDelete(vrfId, removed)
		qim.ProcessConfigDelete(vrfId, modified)
		qim.ProcessQuaggaConfigAdd(vrfId, modified)
		qim.ProcessQuaggaConfigAdd(vrfId, added)
	} else {
		for interfaceName, quaggaConfig := range quaggaConfigs {
			for _, quaggaInstanceConfig := range quaggaConfig {
				qim.SpawnQuagga(vrfId, interfaceName, quaggaInstanceConfig.QuaggaInstanceConfig)
			}
		}
	}
	qim.QuaggaConfigCache[vrfId] = quaggaConfigs
}

func (qim QuaggaInstanceManager) ProcessQuaggaConfigDelete(vrfId int) {
	qim.ProcessVrfQuaggaConfigDelete(vrfId)
}

var (
	quaggaManagers      = make(map[string]QuaggaInstanceManager)
	quaggaManagersMutex sync.Mutex
)

func QuaggaConfigSync(jsonStr string, vrfId int, instanceName string) {
	log.WithFields(log.Fields{
		"vrfId":        vrfId,
		"instanceName": instanceName,
	}).Debug("QuaggaConfigSync")

	manager := GetOrCreateInstanceManager(instanceName)
	manager.QuaggaConfigSync(jsonStr, vrfId)
}

func ProcessQuaggaConfigDelete(vrfId int, instanceName string) {
	manager := GetInstanceManager(instanceName)
	if manager != nil {
		manager.ProcessQuaggaConfigDelete(vrfId)
	}
}

func GetOrCreateInstanceManager(instanceName string) *QuaggaInstanceManager {
	quaggaManagersMutex.Lock()
	defer quaggaManagersMutex.Unlock()

	if manager := getInstanceManager(instanceName); manager != nil {
		return manager
	}

	return createInstanceManager(instanceName)
}

func GetInstanceManager(instanceName string) *QuaggaInstanceManager {
	quaggaManagersMutex.Lock()
	defer quaggaManagersMutex.Unlock()

	return getInstanceManager(instanceName)
}

func getInstanceManager(instanceName string) *QuaggaInstanceManager {
	manager, ok := quaggaManagers[instanceName]
	if !ok {
		return nil
	} else {
		return &manager
	}
}

func createInstanceManager(instanceName string) *QuaggaInstanceManager {
	log.Debug("Create new quagga manager instance")
	manager := NewQuaggaManager()
	quaggaManagers[instanceName] = *manager
	return manager
}
