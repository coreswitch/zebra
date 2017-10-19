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
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

type LeaseValue struct {
	Address        string `json:"address"`
	Hardware       string `json:"hardware"`
	ClientHostName string `json:"client-hostname"`
}

type LeaseMap map[string]*LeaseValue

func LoadLease(reader io.Reader) (*LeaseMap, error) {
	l := lex(reader)
	var leaseValue *LeaseValue
	leaseKeyword := false
	leaseMap := make(LeaseMap, 0)
	cmds := []string{}

	for {
		item := l.nextItem()
		switch item.typ {
		case itemEOF:
			return &leaseMap, nil
		case itemWhiteSpace:
			// Skip.
		case itemString:

			cmds = append(cmds, item.val)
		case itemIdentifier:
			if item.val == "lease" {
				leaseKeyword = true
				continue
			}
			if leaseKeyword {
				leaseValue = &LeaseValue{Address: item.val}
				leaseKeyword = false
				continue
			}
			cmds = append(cmds, item.val)
		case itemLeftBrace:
			// Start of the value, just ignore
		case itemRightBrace:
			// End of the value.
			if leaseValue != nil {
				leaseMap[leaseValue.Address] = leaseValue
			}
		case itemSemiColon:
			if leaseValue != nil {
				switch cmds[0] {
				case "hardware":
					leaseValue.Hardware = cmds[len(cmds)-1]
				case "client-hostname":
					leaseValue.ClientHostName = cmds[len(cmds)-1]
				}
			}
			cmds = cmds[:0]
		case itemError:
			return nil, fmt.Errorf("Parse error")
		}
	}
	return &leaseMap, nil
}

type DhcpRange struct {
	StartIp string `json:"start_ip"`
	EndIp   string `json:"end_ip"`
}

type DhcpLeaseStatus struct {
	LeaseTime uint32       `json:"lease-time"`
	Range     []DhcpRange  `json:"range"`
	Lease     []LeaseValue `json:"lease"`
}

func DhcpLeaseGet(poolConfig *DhcpIpPool) (string, error) {
	dhcpLease := &DhcpLeaseStatus{}
	dhcpLease.Lease = []LeaseValue{}

	leaseFileName := fmt.Sprintf("/var/lib/dhcp/dhcpd-%s.leases", poolConfig.Interface)
	fmt.Println("leaseFileName", leaseFileName)

	dhcpLease.LeaseTime = poolConfig.DefaultLeaseTime
	for _, r := range poolConfig.RangeList {
		ra := DhcpRange{}
		ra.StartIp = r.RangeStartIp
		ra.EndIp = r.RangeEndIp
		dhcpLease.Range = append(dhcpLease.Range, ra)
	}

	input, err := ioutil.ReadFile(leaseFileName)
	if err != nil {
		fmt.Println("ioutil.ReadFile err", err)
		return "", err
	}
	reader := bytes.NewBufferString(string(input))
	leaseMap, err := LoadLease(reader)
	if err != nil {
		fmt.Println("LoadLease err", err)
		return "", err
	}
	for _, lease := range *leaseMap {
		dhcpLease.Lease = append(dhcpLease.Lease, *lease)
	}
	byte, err := json.Marshal(dhcpLease)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return "", err
	}
	return string(byte), nil
}

var (
	etcdDhcpEndpoints = []string{"http://127.0.0.1:2379"}
	etcdDhcpPath      = "/state/services/port/dhcp"
)

func DhcpStatusUpdate() {
	fmt.Println("DhcpStatusUpdate")
	first := true
	fmt.Println("EtcdVrfMap length", len(EtcdVrfMap))
	str := ""

	for _, vrfConfig := range EtcdVrfMap {
		for _, poolConfig := range vrfConfig.Dhcp.Server.DhcpIpPoolList {
			dhcpLease, err := DhcpLeaseGet(&poolConfig)
			if err != nil {
				fmt.Println("DhcpLeaseGet err", err)
				return
			}
			if first {
				first = false
			} else {
				str += ","
			}
			str += fmt.Sprintf("\"%s\":%s", poolConfig.Interface, dhcpLease)
		}
	}
	str = "{" + str + "}"

	fmt.Println("Etcd DHCP status:", str)

	cfg := clientv3.Config{
		Endpoints:   etcdDhcpEndpoints,
		DialTimeout: 3 * time.Second,
	}
	conn, err := clientv3.New(cfg)
	if err != nil {
		fmt.Println("DhcpStatusUpdate clientv3.New:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Put(context.Background(), etcdDhcpPath, str)
	if err != nil {
		fmt.Println("DhcpStatusUpdate Put:", err)
		return
	}
}
