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
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

type GobgpStatGlobal struct {
	As       string `json:"as"`
	RouterId string `json:"router-id"`
}

type GobgpStatNeighbor struct {
	Peer     string `json:"peer"`
	As       string `json:"as"`
	Age      string `json:"age"`
	State    string `json:"state"`
	Recevied string `json:"received"`
	Accepted string `json:"accepted"`
}

type GobgpStatRib struct {
	Network string `json:"network"`
	Nexthop string `json:"next-hop"`
	Metric  uint32 `json:"metric,omitempty"`
}

type GobgpStat struct {
	Global   GobgpStatGlobal     `json:"global"`
	Neighbor []GobgpStatNeighbor `json:"neighbors"`
	Ribs     []GobgpStatRib      `json:"ribs"`
	Vrf      string              `json:"vrf,omitempty"`
}

func GobgpStatusGlobal(stat *GobgpStat) error {
	out, err := exec.Command("gobgp", "-p", "50052", "global").Output()
	if err != nil {
		fmt.Println("GobgpStatus: err", err)
		return err
	}
	strs := strings.Split(string(out), "\n")
	for pos, str := range strs {
		switch pos {
		case 0:
			r := regexp.MustCompile(`^AS:\s+(\d+)`)
			matches := r.FindAllStringSubmatch(str, -1)
			if len(matches) > 0 && len(matches[0]) == 2 {
				stat.Global.As = matches[0][1]
				//fmt.Println(matches[0][1])
			}
		case 1:
			r := regexp.MustCompile(`^Router-ID:\s+([\d\.]+)`)
			matches := r.FindAllStringSubmatch(str, -1)
			if len(matches) > 0 && len(matches[0]) == 2 {
				stat.Global.RouterId = matches[0][1]
				//fmt.Println(matches[0][1])
			}
		}
	}
	return nil
}

func GobgpStatusNeighbor(stat *GobgpStat) error {
	out, err := exec.Command("gobgp", "-p", "50052", "neighbor").Output()
	if err != nil {
		fmt.Println("GobgpStatus: err", err)
	}
	strs := strings.Split(string(out), "\n")
	for pos, str := range strs {
		switch pos {
		case 0:
		default:
			fields := strings.Fields(str)
			if len(fields) == 7 {
				neigh := GobgpStatNeighbor{}
				neigh.Peer = fields[0]
				neigh.As = fields[1]
				neigh.Age = fields[2]
				neigh.State = fields[3]
				neigh.Recevied = fields[5]
				neigh.Accepted = fields[6]
				stat.Neighbor = append(stat.Neighbor, neigh)
			}
		}
	}
	return nil
}

func GobgpStatusRib(stat *GobgpStat) error {
	out, err := exec.Command("gobgp", "-p", "50052", "global", "rib").Output()
	if err != nil {
		fmt.Println("GobgpStatus: err", err)
	}
	strs := strings.Split(string(out), "\n")
	for pos, str := range strs {
		switch pos {
		case 0:
		default:
			fields := strings.Fields(str)
			if len(fields) == 7 {
				rib := GobgpStatRib{}
				rib.Network = fields[1]
				rib.Nexthop = fields[2]
				stat.Ribs = append(stat.Ribs, rib)
			}
		}
	}
	return nil
}

func GobgpStatus() (string, error) {
	stat := &GobgpStat{}
	stat.Neighbor = []GobgpStatNeighbor{}
	stat.Ribs = []GobgpStatRib{}

	err := GobgpStatusGlobal(stat)
	if err != nil {
		return "", err
	}
	err = GobgpStatusNeighbor(stat)
	if err != nil {
		return "", err
	}
	err = GobgpStatusRib(stat)
	if err != nil {
		return "", err
	}
	byte, err := json.Marshal(stat)
	if err != nil {
		fmt.Println("json.Marshal err", err)
		return "", err
	}
	return string(byte), nil
}

var (
	etcdBgpWanEndpoints = []string{"http://127.0.0.1:2379"}
	etcdBgpWanPath      = "/state/services/bgp_wan"
)

func GobgpStatusUpdate() {
	fmt.Println("GobgpStatusUpdate")
	str, err := GobgpStatus()
	if err != nil {
		str = "{}"
	}

	cfg := clientv3.Config{
		Endpoints:   etcdBgpWanEndpoints,
		DialTimeout: 3 * time.Second,
	}
	conn, err := clientv3.New(cfg)
	if err != nil {
		fmt.Println("GobgpStatusUpdate clientv3.New:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Put(context.Background(), etcdBgpWanPath, str)
	if err != nil {
		fmt.Println("GobgpStatusUpdate Put:", err)
		return
	}
}
