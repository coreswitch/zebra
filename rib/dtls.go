// Copyright 2017 Zebra Project
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
	//"os"
	"net"
	"os/exec"
)

func TunnelProbe(vrfId string, peerId string) bool {
	jsonStr, err := exec.Command("curl", "-s", "http://localhost:9999/peers").Output()
	if err != nil {
		fmt.Println("File read error", err)
		return false
	}

	var jsonIntf interface{}
	err = json.Unmarshal([]byte(jsonStr), &jsonIntf)
	if err != nil {
		fmt.Println("json Unmarshal error", err)
		return false
	}

	if jsonMap, ok := jsonIntf.(map[string]interface{}); ok {
		if vrf, ok := jsonMap[vrfId]; ok {
			if vrfMap, ok := vrf.(map[string]interface{}); ok {
				if peer, ok := vrfMap[peerId]; ok {
					if peerMap, ok := peer.(map[string]interface{}); ok {
						for key, value := range peerMap {
							fmt.Println("key:", key)
							if valueMap, ok := value.(map[string]interface{}); ok {
								if state, ok := valueMap["state"]; ok {
									fmt.Println("state", state)
									if state == "ESTABLISHED" {
										return true
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return false
}

func PeerId(vrfId string, nexthop string) string {
	jsonStr, err := exec.Command("curl", "-s", "http://localhost:9999/config").Output()
	if err != nil {
		fmt.Println("File read error", err)
		return ""
	}

	var jsonIntf interface{}
	err = json.Unmarshal([]byte(jsonStr), &jsonIntf)
	if err != nil {
		fmt.Println("json Unmarshal error", err)
		return ""
	}

	if jsonMap, ok := jsonIntf.(map[string]interface{}); ok {
		if links, ok := jsonMap["links"]; ok {
			if vrfs, ok := links.(map[string]interface{}); ok {
				if vrf, ok := vrfs[vrfId]; ok {
					if peers, ok := vrf.(map[string]interface{}); ok {
						if peer, ok := peers["peers"]; ok {
							if entry, ok := peer.(map[string]interface{}); ok {
								for peerId, item := range entry {
									if peer, ok := item.(map[string]interface{}); ok {
										if nhop, ok := peer["nexthop-id"]; ok {
											if nhop == nexthop {
												return peerId
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return ""
}

func DtlsNexthop(vrfId uint32, nexthop net.IP) bool {
	vrf := fmt.Sprintf("vrf%d", vrfId)
	peerId := PeerId(vrf, nexthop.String())
	fmt.Println("Peer ID:", peerId)
	if peerId == "" {
		fmt.Println("Peer does not exists")
		return false
	}
	result := TunnelProbe(vrf, peerId)
	if result {
		fmt.Println("Tunnel is UP")
		return true
	} else {
		fmt.Println("Tunnel is DOWN")
		return false
	}
}
