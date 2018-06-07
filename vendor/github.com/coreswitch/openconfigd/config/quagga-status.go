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
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreswitch/openconfigd/quagga"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type QuaggaStat struct {
	Stats []GobgpStat `json:"bgp_lan"`
}

func QuaggaStatusBgpSummary(vrf string, stat *GobgpStat) error {
	var in []string
	in = append(in, "show ip bgp sum\n")
	out, err := VrfQuaggaGet(vrf, "bgpd", quagga.GetPasswd(), time.Second, in)
	if err != nil {
		log.Error("QuaggaStatusBgpSummary: VrfQuaggaGet()", err)
		return err
	}
	neighbor := false

	strs := strings.Split(out[0], "\n")
	for _, str := range strs {
		if !neighbor {
			r := regexp.MustCompile(`^BGP[a-z\s]+([\d\.]+).*\s([\d\.]+)$`)
			matches := r.FindAllStringSubmatch(str, -1)
			if len(matches) > 0 && len(matches[0]) >= 3 {
				stat.Global.RouterId = matches[0][1]
				stat.Global.As = matches[0][2]
			}
			r = regexp.MustCompile(`^Neighbor`)
			if r.MatchString(str) {
				neighbor = true
			}
		} else {
			r := regexp.MustCompile(`^[0-9]`)
			if r.MatchString(str) {
				fields := strings.Fields(str)
				if len(fields) >= 10 {
					neigh := GobgpStatNeighbor{}
					neigh.Peer = fields[0]
					neigh.As = fields[2]
					neigh.Accepted = fields[3]
					neigh.Recevied = fields[3]
					neigh.Age = fields[8]
					if r = regexp.MustCompile(`\d+`); r.MatchString(fields[9]) {
						neigh.State = "Established"
					} else {
						neigh.State = fields[9]
					}
					stat.Neighbor = append(stat.Neighbor, neigh)
				}
			}
		}
	}
	return nil
}

func QuaggaStatusRib(vrf string, stat *GobgpStat) error {
	if len(stat.Neighbor) == 0 {
		log.Info("QuaggaStatusRib(): no neighbors")
		return nil
	} else {
		log.Infof("QuaggaStatusRib(): No. of neighbors: %d", len(stat.Neighbor))
	}
	var in []string
	fmt.Println("QuaggaStatusRib(): neighbor", stat.Neighbor[0].Peer)

	in = append(in, fmt.Sprintf("show ip bgp neighbor %s received-routes\n", stat.Neighbor[0].Peer))
	out, err := VrfQuaggaGet(vrf, "bgpd", quagga.GetPasswd(), time.Second, in)
	if err != nil {
		log.Error("QuaggaStatusBgpRib: VrfQuaggaGet()", err)
		return err
	}
	network := false

	strs := strings.Split(out[0], "\n")
	for _, str := range strs {
		if !network {
			r := regexp.MustCompile(`^\s+Network`)
			if r.MatchString(str) {
				network = true
				fmt.Println("network is on")
			}
		} else {
			r := regexp.MustCompile(`^\*`)
			if r.MatchString(str) {
				fields := strings.Fields(str)
				fmt.Println(len(fields))
				if len(fields) > 3 {
					rib := GobgpStatRib{}
					rib.Network = fields[1]
					rib.Nexthop = fields[2]
					metric, err := strconv.Atoi(fields[3])
					if err == nil {
						rib.Metric = uint32(metric)
					}
					stat.Ribs = append(stat.Ribs, rib)
				}
			}
		}
	}
	return nil
}

func QuaggaStatusVrf(vrf string, stats *QuaggaStat) {
	fmt.Println("QuaggaStatusVrf", vrf)
	stat := GobgpStat{}
	stat.Vrf = vrf
	stat.Neighbor = []GobgpStatNeighbor{}
	stat.Ribs = []GobgpStatRib{}

	err := QuaggaStatusBgpSummary(vrf, &stat)
	if err != nil {
		return
	}
	err = QuaggaStatusRib(vrf, &stat)
	if err != nil {
		return
	}
	stats.Stats = append(stats.Stats, stat)
}

var (
	etcdBgpLanEndpoints = []string{"http://127.0.0.1:2379"}
	etcdBgpLanPath      = "/state/services/bgp_lan"
)

func QuaggaStatusUpdate() {
	fmt.Println("QuaggaStatusUpdate")
	stats := QuaggaStat{}

	for vrfId, _ := range QuaggaProc {
		vrf := fmt.Sprintf("vrf%d", vrfId)
		QuaggaStatusVrf(vrf, &stats)
	}

	str := ""
	byte, err := json.Marshal(stats)
	if err != nil {
		str = "{}"
	}
	str = string(byte)
	cfg := clientv3.Config{
		Endpoints:   etcdBgpLanEndpoints,
		DialTimeout: 3 * time.Second,
	}
	conn, err := clientv3.New(cfg)
	if err != nil {
		fmt.Println("QuaggaStatusUpdate clientv3.New:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Put(context.Background(), etcdBgpLanPath, str)
	if err != nil {
		fmt.Println("QuaggaStatusUpdate Put:", err)
		return
	}
}
