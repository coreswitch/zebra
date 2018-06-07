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
	"os"
	"text/template"

	"github.com/coreswitch/openconfigd/quagga"
	"github.com/coreswitch/process"
)

type Ospf struct {
	Network      string                `mapstructure:"network" json:"network,omitempty"`
	Area         uint32                `mapstructure:"area" json:"area,omitempty"`
	InterfaceIps []InterfaceIp         `mapstructure:"interfaces" json:"interfaces,omitempty"`
	Interface    string                `mapstructure:"interface" json:"interface,omitempty"` // Deplicated from 2.4
	PrimaryList  []DistributeListEntry `mapstructure:"distribute-list" json:"distribute-list,omitempty"`
	BackupList   []DistributeListEntry `mapstructure:"backup-distribute-list" json:"backup-distribute-list,omitempty"`
}

type OspfArray []Ospf

var OspfVrfMap = map[int]OspfArray{}
var OspfProcessMap = map[int]*process.Process{}

func (lhs *Ospf) Equal(rhs *Ospf) bool {
	if lhs.Network != rhs.Network {
		return false
	}
	if lhs.Area != rhs.Area {
		return false
	}
	if len(lhs.InterfaceIps) != len(rhs.InterfaceIps) {
		return false
	}
	{
		lmap := make(map[string]*InterfaceIp)
		for i, l := range lhs.InterfaceIps {
			lmap[mapkey(i, string(l.Name))] = &lhs.InterfaceIps[i]
		}
		for i, r := range rhs.InterfaceIps {
			if l, y := lmap[mapkey(i, string(r.Name))]; !y {
				return false
			} else if !r.Equal(l) {
				return false
			}
		}
	}
	return true
}

func (lhs *OspfArray) Equal(rhs *OspfArray) bool {
	if len(*lhs) != len(*rhs) {
		return false
	}
	for pos, o := range *lhs {
		if !o.Equal(&(*rhs)[pos]) {
			return false
		}
	}
	return true
}

var ospfTemplate = `!
password 8 {{passwdHash}}
service password-encryption
!
{{range $i, $v := .OspfArray}}{{interfaceIp2Config $v}}{{end}}
!
router ospf
  redistribute bgp metric-type 1
  redistribute connected metric-type 1
  default-information originate metric-type 1
{{areaAuthentication .OspfArray}}
{{range $i, $v := .OspfArray}}  network {{$v.Network}} area {{$v.Area}}
{{end}}
!
`

func areaAuthentication(ospfArray *OspfArray) string {
	var areaAuth bool
	for _, ospf := range *ospfArray {
		for _, ifps := range ospf.InterfaceIps {
			if ifps.Ip.OspfIp.AuthenticationKey != "" {
				areaAuth = true
			}
		}
	}
	if areaAuth {
		return "  area 0 authentication"
	}
	return ""
}

func interface2Config(ifp InterfaceIp) string {
	cfg := ifp.Ip.OspfIp
	// if cfg.AuthenticationKey == "" && cfg.Cost == 0 && cfg.DeadInterval == 0 &&
	// 	cfg.HelloInterval == 0 && cfg.Priority == "" && cfg.RetransmitInterval == 0 && cfg.TransmitDelay == 0 {
	// 	return ""
	// }
	str := fmt.Sprintf("interface %s\n", ifp.Name)
	str += " ip ospf mtu-ignore\n"
	if cfg.AuthenticationKey != "" {
		str += fmt.Sprintf(" ip ospf authentication-key %s\n", cfg.AuthenticationKey)
	}
	if cfg.Cost != 0 {
		str += fmt.Sprintf(" ip ospf cost %d\n", cfg.Cost)
	}
	if cfg.DeadInterval != 0 {
		str += fmt.Sprintf(" ip ospf dead-interval %d\n", cfg.DeadInterval)
	}
	if cfg.HelloInterval != 0 {
		str += fmt.Sprintf(" ip ospf hello-interval %d\n", cfg.HelloInterval)
	}
	if cfg.Priority != "" {
		str += fmt.Sprintf(" ip ospf priority %s\n", cfg.Priority)
	}
	if cfg.RetransmitInterval != 0 {
		str += fmt.Sprintf(" ip ospf retransmit-interval %d\n", cfg.RetransmitInterval)
	}
	if cfg.TransmitDelay != 0 {
		str += fmt.Sprintf(" ip ospf transmit-delay %d\n", cfg.TransmitDelay)
	}
	return str
}

//func interfaceIp2Config(ifps []InterfaceIp) string {
func interfaceIp2Config(ospf Ospf) string {
	str := ""
	for _, ifp := range ospf.InterfaceIps {
		str += interface2Config(ifp)
	}
	return str
}

func OspfExec(vrfId int, ospfArray *OspfArray) {
	// Config file name.
	configFile := fmt.Sprintf("/etc/quagga/ospfd-vrf%d.conf", vrfId)

	// Open config file.
	f, err := os.Create(configFile)
	if err != nil {
		fmt.Println("Create file:", err)
		return
	}

	// Generate config from template.
	type TemplValue struct {
		OspfArray *OspfArray
	}
	tmpl := template.Must(template.New("ospfTmpl").Funcs(template.FuncMap{
		"areaAuthentication": areaAuthentication,
		"interfaceIp2Config": interfaceIp2Config,
		"passwdHash":         quagga.GetHash,
	}).Parse(ospfTemplate))
	tmpl.Execute(f, &TemplValue{OspfArray: ospfArray})

	// Set args.
	args := []string{
		"-u", "root",
		"-g", "root",
		"-f", configFile,
		"-V", fmt.Sprintf("vrf%d", vrfId),
		"-z", fmt.Sprintf("/var/run/zserv-vrf%d.api", vrfId),
	}

	// Prepare ospfd process.
	proc := process.NewProcess("ospfd", args...)
	proc.StartTimer = 3
	OspfProcessMap[vrfId] = proc

	// Register and start process.
	process.ProcessRegister(proc)
}

func OspfVrfStop(vrfId int) {
	if proc, ok := OspfProcessMap[vrfId]; ok {
		process.ProcessUnregister(proc)
		delete(OspfProcessMap, vrfId)
	}
}

func OspfVrfSync(vrfId int, cfg *VrfsConfig) {
	fmt.Println("OSPF cfg", cfg.Ospf)

	// Compare ospf config with previous one.  If it is same just return.
	o := OspfVrfMap[vrfId]
	n := &cfg.Ospf
	fmt.Println("XXX old", o, "new", n)
	// if n.Equal(&o) {
	// 	fmt.Println("XXX old and new is same return")
	// 	return
	// }

	// Stop ospfd if it is running.
	OspfVrfStop(vrfId)

	// Update stored config.
	OspfVrfMap[vrfId] = cfg.Ospf

	// Exit if no need of ospfd to be run.
	if len(cfg.Ospf) == 0 {
		fmt.Println("XX Empty ospf config returning")
		return
	}

	// Launch process.
	OspfExec(vrfId, &cfg.Ospf)
}

func OspfVrfDelete(vrfId int) {
	// Stop ospfd if it is running.
	OspfVrfStop(vrfId)

	// Delete stored config.
	delete(OspfVrfMap, vrfId)
}

func OspfVrfExit() {
	// Foreach ospf vrf config.
	for vrfId, _ := range OspfVrfMap {
		OspfVrfDelete(vrfId)
	}
}
