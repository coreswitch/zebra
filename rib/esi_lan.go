// Copyright 2017 zebra Project
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

package rib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/howeyc/fsnotify"
	"github.com/mitchellh/mapstructure"
)

type LanForwardingTable struct {
	InterfaceTable struct {
		Tenant struct {
			Interfaces []string `mapstructure:"local-interface"`
		} `mapstructure:"tenant"`
	} `mapstructure:"interface-table"`
}

var (
	LanInterfaceTable LanForwardingTable
	LanInterfaceFound bool
	LanInterfaceMutex sync.RWMutex
)

func lanFileRead(fileName string) error {
	LanInterfaceMutex.Lock()
	defer LanInterfaceMutex.Unlock()

	// Init variable
	LanInterfaceTable = LanForwardingTable{}

	// Read file.
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	// Parse JSON.
	var jsonIntf interface{}
	err = json.Unmarshal(bytes, &jsonIntf)
	if err != nil {
		return err
	}

	// Map JSON to struct.
	err = mapstructure.Decode(jsonIntf, &LanInterfaceTable)
	if err != nil {
		return err
	}

	// Set LAN is found
	LanInterfaceFound = true

	// LAN File read success
	fmt.Println("--lan list--")
	for _, lan := range LanInterfaceTable.InterfaceTable.Tenant.Interfaces {
		fmt.Println(lan)
		ifp := IfLookupByName(lan)
		if ifp != nil && !ifp.IsUp() {
			fmt.Println(lan, "should be up")
			LinkSetUp(ifp)
		}
	}

	return nil
}

func LanFileMonitor(fileName string) {
	for {
		err := lanFileRead(fileName)

		// File read success.
		if err == nil {
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				goto Retry
			}

			err = watcher.Watch(fileName)
			if err != nil {
				goto Retry
			}

			for {
				select {
				case <-watcher.Event:
					err = lanFileRead(fileName)
					if err != nil {
						watcher.Close()
						goto Retry
					}
				case <-watcher.Error:
					watcher.Close()
					goto Retry
				}
			}

		}
	Retry:
		time.Sleep(time.Second * 10)
	}
}

func IsLanInterface(ifName string) bool {
	LanInterfaceMutex.Lock()
	defer LanInterfaceMutex.Unlock()

	// When "forwarder.json" is not yet created, make interface up by default.
	if !LanInterfaceFound {
		return true
	}

	for _, lan := range LanInterfaceTable.InterfaceTable.Tenant.Interfaces {
		if ifName == lan {
			return true
		}
	}
	return false
}

func LanInterfaceMonitorStart() {
	go LanFileMonitor("/persist/etc/esi/forwarder.json")
}
