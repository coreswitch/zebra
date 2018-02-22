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

package rib

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

var (
	etcdEndpoints = []string{"http://127.0.0.1:2379"}
	etcdPath      = "/state/services/port/status"
)

type EtcdIfStatus struct {
	Status         string             `json:"status"`
	ProtocolStatus string             `json:"protocol_status"`
	Address        string             `json:"address,omitempty"`
	Prefix         string             `json:"prefix,omitempty"`
	IPv6           []EtcdIfStatusIPv6 `json:"ipv6,omitempty"`
}

type EtcdIfStatusIPv6 struct {
	Address string `json:"address,omitempty"`
	Prefix  string `json:"prefix,omitempty"`
}

func EtcdIf(ifName string, ifp *Interface) string {
	eif := &EtcdIfStatus{}

	if ifp.IsUp() {
		eif.Status = "up"
	} else {
		eif.Status = "down"
	}
	if ifp.IsRunning() {
		eif.ProtocolStatus = "up"
	} else {
		eif.ProtocolStatus = "down"
	}
	if len(ifp.Addrs[AFI_IP]) > 0 {
		addr := ifp.Addrs[AFI_IP][0]
		eif.Address = addr.Prefix.IP.String()
		prefix := addr.Prefix.Copy()
		prefix.ApplyMask()
		eif.Prefix = prefix.String()
	}
	for _, addr := range ifp.Addrs[AFI_IP6] {
		if addr.Prefix.IP.IsGlobalUnicast() {
			var ipv6 EtcdIfStatusIPv6
			ipv6.Address = addr.Prefix.IP.String()
			prefix := addr.Prefix.Copy()
			prefix.ApplyMask()
			ipv6.Prefix = prefix.String()
			eif.IPv6 = append(eif.IPv6, ipv6)
		}
	}
	json, _ := json.Marshal(eif)

	return fmt.Sprintf("\"%s\":%s", ifName, string(json))
}

func EtcdSetIfStatus(ifname string, up bool, running bool) {
	//fmt.Println("Setting ifname", ifname, up, running)

	var str string
	var separator bool

	for ifName, ifp := range IfMap {
		if separator {
			str += ","
		} else {
			separator = true
		}
		str += EtcdIf(ifName, ifp)
	}
	str = "{" + str + "}"

	fmt.Println("Etcd if status:", str)

	cfg := clientv3.Config{
		Endpoints:   etcdEndpoints,
		DialTimeout: 3 * time.Second,
	}
	conn, err := clientv3.New(cfg)
	if err != nil {
		fmt.Println("EtcdSetIfStatus clientv3.New:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Put(context.Background(), etcdPath, str)
	if err != nil {
		fmt.Println("EtcdSetIfStatus Put:", err)
		return
	}
}
