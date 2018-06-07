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
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreswitch/openconfigd/quagga"
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type OspfNeighbor struct {
	Neighbor  string `json:"neighbor"`
	Priority  int    `json:"priority"`
	State     string `json:"state"`
	Address   string `json:"address"`
	Interface string `json:"interface"`
}

type OspfNexthop struct {
	Address   string `mapstructure:"address" json:"address"`
	Interface string `mapstructure:"interface" json:"interface"`
}

type OspfRoute struct {
	Network  string        `json:"network"`
	Nexthops []OspfNexthop `json:"next-hop"`
}

type OspfVrfStat struct {
	Vrf       string         `json:"vrf"`
	Neighbors []OspfNeighbor `json:"neighbors"`
	Routes    []OspfRoute    `json:"routes"`
}

type OspfStat struct {
	Ospf []OspfVrfStat `json:"ospf"`
}

var (
	OspfEtcdEndpoints  = []string{"http://127.0.0.1:2379"}
	OspfEtcdStatusPath = "/state/services/ospf"
)

func VrfName(vrfId int) string {
	return fmt.Sprintf("vrf%d", vrfId)
}

var ospfVrfRouteCmd = `#! /bin/bash
source /etc/bash_completion.d/cli
SHOW_MODE=json
show ip route vrf %s ospf
`

// Only used for JSON parse.
type OspfJSONRoutes struct {
	Prefix   string        `mapstructure:"prefix"`
	Nexthops []OspfNexthop `mapstructure:"nexthops"`
	Metric   int           `mapstructure:"metric"`
	Distance int           `mapstructure:"distance"`
	Type     string        `mapstructure:"type"`
}

func OspfStatusRoute(vrfName string, stat *OspfVrfStat) error {
	cmdName := "/tmp/.ospf_vrf_route_show.sh"
	os.Remove(cmdName)

	// OSPF route via ribd JSON.
	cmdStr := fmt.Sprintf(ospfVrfRouteCmd, vrfName)

	err := ioutil.WriteFile(cmdName, []byte(cmdStr), 0755)
	if err != nil {
		fmt.Println("[ospf]WriteFile err:", err)
		return err
	}

	// Execute command.
	cmd := exec.Command(cmdName)
	out, err := cmd.Output()
	os.Remove(cmdName)
	if err != nil {
		fmt.Println("OspfStatusRoute err:", err)
		return err
	}

	// Parse JSON
	var jsonIntf interface{}
	err = json.Unmarshal(out, &jsonIntf)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return err
	}

	// Convert to OspfVrfStat
	for _, ent := range jsonIntf.([]interface{}) {
		var val OspfJSONRoutes
		err = mapstructure.Decode(ent, &val)
		if err != nil {
			fmt.Println("Decode err:", err)
			return err
		}
		r := OspfRoute{
			Network:  val.Prefix,
			Nexthops: val.Nexthops,
		}
		stat.Routes = append(stat.Routes, r)
	}
	return nil
}
func OspfStatusNeighbor(vrfName string, stat *OspfVrfStat) error {
    var in []string
    in = append(in, "show ip ospf neighbor\n")
    out, err := VrfQuaggaGet(vrfName, "ospfd", quagga.GetPasswd(), time.Second, in)
    if err != nil {
        log.Error("QuaggaStatusBgpSummary: VrfQuaggaGet()", err)
        return err
    }

	strs := strings.Split(out[0], "\n")
	for _, str := range strs {
		r := regexp.MustCompile(`^[0-9]`)
		if r.MatchString(str) {
			fields := strings.Fields(str)
			if len(fields) >= 9 {
				neigh := OspfNeighbor{}
				neigh.Neighbor = fields[0]
				neigh.Priority, _ = strconv.Atoi(fields[1])
				neigh.State = fields[2]
				neigh.Address = fields[4]
				neigh.Interface = fields[5]
				stat.Neighbors = append(stat.Neighbors, neigh)
			}
		}
	}
	return nil
}

func OspfStatusVrf(vrfName string) *OspfVrfStat {
	fmt.Println("OspfStatusVrf", vrfName)
	stat := &OspfVrfStat{
		Vrf:       vrfName,
		Neighbors: []OspfNeighbor{},
		Routes:    []OspfRoute{},
	}
	OspfStatusRoute(vrfName, stat)
	OspfStatusNeighbor(vrfName, stat)

	return stat
}

func OspfStatusUpdate() {
	stats := OspfStat{}

	for vrfId, _ := range OspfVrfMap {
		s := OspfStatusVrf(VrfName(vrfId))
		stats.Ospf = append(stats.Ospf, *s)
	}

	str := "{}"
	byte, err := json.Marshal(stats)
	if err == nil {
		str = string(byte)
	}

	cfg := clientv3.Config{
		Endpoints:   OspfEtcdEndpoints,
		DialTimeout: 3 * time.Second,
	}
	conn, err := clientv3.New(cfg)
	if err != nil {
		fmt.Println("OspfStatusUpdate clientv3.New:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Put(context.Background(), OspfEtcdStatusPath, str)
	if err != nil {
		fmt.Println("OpsfStatusUpdate Put:", err)
		return
	}
}
