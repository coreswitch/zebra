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
	"reflect"
	"sort"
	"strings"
	"time"
	//"net"

	log "github.com/sirupsen/logrus"
	//"github.com/coreswitch/netutil"
	"github.com/mitchellh/mapstructure"
)

type IPv4Addr struct {
	Ip string `mapstructure:"ip" json:"ip,omitempty"`
}

type IPv4 struct {
	Address []IPv4Addr `mapstructure:"address" json:"address,omitempty"`
}

type Vlan struct {
	VlanId int `mapstructure:"vlan-id" json:"vlan-id,omitempty"`
}

type Hub struct {
	Address string `mapstructure:"address" json:"address,omitempty"`
}

type Interface struct {
	IPv4           IPv4   `mapstructure:"ipv4" json:"ipv4,omitempty"`
	Vlans          []Vlan `mapstructure:"vlans" json:"vlans,omitempty"`
	Name           string `mapstructure:"name" json:"name,omitempty"`
	Duplex         string `mapstructure:"duplex" json:"duplex,omitempty"`
	Speed          string `mapstructure:"speed" json:"speed,omitempty"`
	DhcpRelayGroup string `mapstructure:"dhcp-relay-group" json:"dhcp-relay-group,omitempty"`
}

type Interfaces struct {
	Interface []Interface `mapstructure:"interface" json:"interface,omitempty"`
}

func (i Interfaces) Len() int {
	return len(i.Interface)
}

func (i Interfaces) Swap(j, k int) {
	i.Interface[j], i.Interface[k] = i.Interface[k], i.Interface[j]
}

func (i Interfaces) Less(j, k int) bool {
	return i.Interface[j].Name < i.Interface[k].Name
}

type Static struct {
	Route []Route `mapstructure:"route" json:"route,omitempty"`
}

type QuaggaBgp struct {
	CiscoConfig string `mapstructure:"cisco-config" json:"cisco-config,omitempty"`
	Interface   string `mapstructure:"interface" json:"interface,omitempty"`
}

type PriorityConfig struct {
	Priority string `mapstructure:"priority" json:"priority,omitempty"`
}

type VrfsConfig struct {
	Name       string         `mapstructure:"name" json:"name,omitempty"`
	Id         int            `mapstructure:"vrf_id" json:"vrf_id,omitempty"`
	Rd         string         `mapstructure:"rd" json:"rd,omitempty"`
	RtImport   string         `mapstructure:"rt_import" json:"rt_import,omitempty"`
	RtExport   string         `mapstructure:"rt_export" json:"rt_export,omitempty"`
	RtBoth     string         `mapstructure:"rt_both" json:"rt_both,omitempty"`
	VrfRibs    []VrfRib       `mapstructure:"ribs" json:"ribs,omitempty"`
	Hubs       []Hub          `mapstructure:"hubs" json:"hubs,omitempty"`
	HubNode    string         `mapstructure:"hub_node" json:"hub_node,omitempty"`
	Interfaces Interfaces     `mapstructure:"interfaces" json:"interfaces,omitempty"`
	Vrrp       []Vrrp         `mapstructure:"vrrp" json:"vrrp,omitempty"`
	Dhcp       Dhcp           `mapstructure:"dhcp" json:"dhcp,omitempty"`
	Static     Static         `mapstructure:"static" json:"static,omitempty"`
	Bgp        []QuaggaBgp    `mapstructure:"bgp" json:"bgp,omitempty"`
	Ospf       OspfArray      `mapstructure:"ospf" json:"ospf,omitempty"`
	Pair       PriorityConfig `mapstructure:"pair" json:"pair,omitempty"`
	Gateway    PriorityConfig `mapstructure:"gateway" json:"gateway,omitempty"`
}

var EtcdVrfMap = map[int]VrfsConfig{}

func EtcdVrfVlanSubinterfacesDelete(vrf *VrfsConfig) {
	for _, ifp := range vrf.Interfaces.Interface {
		for _, Vlan := range ifp.Vlans {
			ExecLine(fmt.Sprintf("delete interfaces interface %s vlans %d", ifp.Name, Vlan.VlanId))
		}
	}
	Commit()
}

func EtcdVrfAddressClear(vrfId int, vrf *VrfsConfig) {
	for _, ifp := range vrf.Interfaces.Interface {
		ExecLine(fmt.Sprintf("delete interfaces interface %s ipv4", ifp.Name))
		Commit()
		ExecLine(fmt.Sprintf("delete interfaces interface %s vrf vrf%d", ifp.Name, vrfId))
		if ifp.DhcpRelayGroup != "" {
			ExecLine(fmt.Sprintf("delete interfaces interface %s dhcp-relay-group %s", ifp.Name, ifp.DhcpRelayGroup))
		}
	}
}

func EtcdVrfVlanSubinterfacesAdd(vrf *VrfsConfig) {
	for _, ifp := range vrf.Interfaces.Interface {
		for _, Vlan := range ifp.Vlans {
			ExecLine(fmt.Sprintf("set interfaces interface %s vlans %d", ifp.Name, Vlan.VlanId))
		}
	}

	Commit()
}

// TODO: Return failure if we finally fail.
// TODO: Generic failure recovery mechanism for any configuration set
func ExecLineWaitIfNoMatch(command string) {
	loopCount := 0
	for {
		result := ExecLine(command)
		if strings.TrimRight(result, "\n") != "NoMatch" {
			break
		} else {
			if loopCount == 10 {
				fmt.Println("Execution of command failed", command)
				break
			}
			loopCount++
			time.Sleep(1 * time.Second)
		}
	}
}

func EtcdVrfAddressAdd(vrfId int, vrf *VrfsConfig) {
	for _, ifp := range vrf.Interfaces.Interface {
		// Wait for Ribd to create the interface in case we get NoMatch
		ExecLineWaitIfNoMatch(fmt.Sprintf("set interfaces interface %s", ifp.Name))
		ExecLine(fmt.Sprintf("set interfaces interface %s vrf vrf%d", ifp.Name, vrfId))
		for _, addr := range ifp.IPv4.Address {
			ExecLine(fmt.Sprintf("set interfaces interface %s ipv4 address %s", ifp.Name, addr.Ip))
		}
		if ifp.DhcpRelayGroup != "" {
			ExecLine(fmt.Sprintf("set interfaces interface %s dhcp-relay-group %s", ifp.Name, ifp.DhcpRelayGroup))
		}
		Commit()
	}
}

func ProcessInterfacesAdd(vrfId int, ifaces []Interface) {
	for _, ifp := range ifaces {
		for _, Vlan := range ifp.Vlans {
			ExecLine(fmt.Sprintf("set interfaces interface %s vlans %d", ifp.Name, Vlan.VlanId))
		}
	}
	Commit()

	for _, ifp := range ifaces {
		// Wait for Ribd to create the interface in case we get NoMatch
		ExecLineWaitIfNoMatch(fmt.Sprintf("set interfaces interface %s", ifp.Name))
		ExecLine(fmt.Sprintf("set interfaces interface %s vrf vrf%d", ifp.Name, vrfId))
		for _, addr := range ifp.IPv4.Address {
			ExecLine(fmt.Sprintf("set interfaces interface %s ipv4 address %s", ifp.Name, addr.Ip))
		}
		if ifp.DhcpRelayGroup != "" {
			ExecLine(fmt.Sprintf("set interfaces interface %s dhcp-relay-group %s", ifp.Name, ifp.DhcpRelayGroup))
		}
	}
}

func ProcessInterfacesDelete(vrfId int, ifaces []Interface) {
	for _, ifp := range ifaces {
		ExecLine(fmt.Sprintf("delete interfaces interface %s ipv4", ifp.Name))
		Commit()
		ExecLine(fmt.Sprintf("delete interfaces interface %s vrf vrf%d", ifp.Name, vrfId))
		if ifp.DhcpRelayGroup != "" {
			ExecLine(fmt.Sprintf("delete interfaces interface %s dhcp-relay-group %s", ifp.Name, ifp.DhcpRelayGroup))
		}
	}

	for _, ifp := range ifaces {
		for _, Vlan := range ifp.Vlans {
			ExecLine(fmt.Sprintf("delete interfaces interface %s vlans %d", ifp.Name, Vlan.VlanId))
		}
	}
	Commit()
}

func CalculateVlanChanges(oldVlans []Vlan, newVlans []Vlan) ([]Vlan, []Vlan) {
	var found bool
	var addedVlans []Vlan
	var removedVlans []Vlan
	for _, oldVlan := range oldVlans {
		found = false
		for _, newVlan := range newVlans {
			if oldVlan.VlanId == newVlan.VlanId {
				found = true
				break
			}
		}
		if found == false {
			removedVlans = append(removedVlans, oldVlan)
		}
	}

	for _, newVlan := range newVlans {
		found = false
		for _, oldVlan := range oldVlans {
			if oldVlan.VlanId == newVlan.VlanId {
				found = true
				break
			}
		}
		if found == false {
			addedVlans = append(addedVlans, newVlan)
		}
	}
	return addedVlans, removedVlans
}

func RemoveVlans(ifp Interface, vlans []Vlan) {
	for _, Vlan := range vlans {
		ExecLine(fmt.Sprintf("delete interfaces interface %s vlans %d", ifp.Name, Vlan.VlanId))
	}
	Commit()
}

func AddVlans(ifp Interface, vlans []Vlan) {
	for _, Vlan := range vlans {
		ExecLine(fmt.Sprintf("set interfaces interface %s vlans %d", ifp.Name, Vlan.VlanId))
	}
	// Vlan creation requires to be processed by ribd before next operation on interface
	Commit()
}

func CalculateAddressChanges(oldAddresses []IPv4Addr, newAddresses []IPv4Addr) ([]IPv4Addr, []IPv4Addr) {
	var found bool
	var addedIPs []IPv4Addr
	var removedIPs []IPv4Addr
	for _, oldIP := range oldAddresses {
		found = false
		for _, newIP := range newAddresses {
			if oldIP.Ip == newIP.Ip {
				found = true
				break
			}
		}
		if found == false {
			removedIPs = append(removedIPs, oldIP)
		}
	}

	for _, oldIP := range newAddresses {
		found = false
		for _, newIP := range oldAddresses {
			if oldIP.Ip == newIP.Ip {
				found = true
				break
			}
		}
		if found == false {
			addedIPs = append(addedIPs, oldIP)
		}
	}
	return addedIPs, removedIPs
}

func RemoveIPAddress(ifp Interface, ipaddrs []IPv4Addr) {
	for _, ip := range ipaddrs {
		ExecLine(fmt.Sprintf("delete interfaces interface %s ipv4 address %s", ifp.Name, ip.Ip))
	}
}

func AddIPAddress(ifp Interface, ipaddrs []IPv4Addr) {
	for _, ip := range ipaddrs {
		ExecLine(fmt.Sprintf("set interfaces interface %s ipv4 address %s", ifp.Name, ip.Ip))
	}
}

func ProcessVrfInterfaceChanges(vrfId int, currentConfig *VrfsConfig, newConfig *VrfsConfig) {
	var addedInterfaces []Interface
	var deletedInterfaces []Interface
	//var updatedInterfaces []Interface
	var found bool
	for _, oldifp := range currentConfig.Interfaces.Interface {
		found = false
		for _, newifp := range newConfig.Interfaces.Interface {
			if oldifp.Name == newifp.Name {
				// Check IPV4 diff
				addedIPs, removedIPs := CalculateAddressChanges(oldifp.IPv4.Address, newifp.IPv4.Address)
				fmt.Println("Interface : ", oldifp.Name, "Added ips : ", addedIPs, "Removed ips: ", removedIPs)

				// Check Vlans diff
				addedVlans, removedVlans := CalculateVlanChanges(oldifp.Vlans, newifp.Vlans)
				fmt.Println("Interface : ", oldifp.Name, "Added vlans : ", addedVlans, "Removed Vlans: ", removedVlans)

				RemoveIPAddress(oldifp, removedIPs)
				RemoveVlans(oldifp, removedVlans)

				AddVlans(newifp, addedVlans)
				AddIPAddress(newifp, addedIPs)

				// Check DhcpRelayGroup
				if oldifp.DhcpRelayGroup != newifp.DhcpRelayGroup {
					if oldifp.DhcpRelayGroup != "" {
						ExecLine(fmt.Sprintf("delete interfaces interface %s dhcp-relay-group %s", oldifp.Name, oldifp.DhcpRelayGroup))
					}
					if newifp.DhcpRelayGroup != "" {
						ExecLine(fmt.Sprintf("set interfaces interface %s dhcp-relay-group %s", newifp.Name, newifp.DhcpRelayGroup))
					}

				}
				found = true
				break
			}
		}
		if found == false {
			deletedInterfaces = append(deletedInterfaces, oldifp)
		}
	}

	for _, newifp := range newConfig.Interfaces.Interface {
		found = false
		for _, oldifp := range currentConfig.Interfaces.Interface {
			if oldifp.Name == newifp.Name {
				found = true
				break
			}
		}
		if found == false {
			addedInterfaces = append(addedInterfaces, newifp)
		}
	}

	fmt.Println("Added interfaces: ", addedInterfaces, "Removed Interfaces: ", deletedInterfaces)
	ProcessInterfacesAdd(vrfId, addedInterfaces)
	ProcessInterfacesDelete(vrfId, deletedInterfaces)
}

var EtcdVrfStaticMap = map[int]*EtcdVrfStaticRoute{}

type EtcdVrfStaticRoute map[string]string

func EtcdVrfStaticUpdate(vrfId int, vrf *VrfsConfig) {
	EtcdVrfStaticClear(vrfId)
	if len(vrf.Static.Route) > 0 {
		staticRoutes := EtcdVrfStaticRoute{}
		for _, route := range vrf.Static.Route {
			if len(route.NexthopList) > 0 {
				prefix := route.Prefix
				nexthop := route.NexthopList[0].Address
				ExecLine(fmt.Sprintf("set vrf name vrf%d static route %s nexthop %s", vrfId, prefix, nexthop))
				staticRoutes[prefix] = nexthop
			}
		}
		EtcdVrfStaticMap[vrfId] = &staticRoutes
	}
	Commit()
}

func EtcdVrfStaticClear(vrfId int) {
	if staticRoutes, ok := EtcdVrfStaticMap[vrfId]; ok {
		for prefix, nexthop := range *staticRoutes {
			ExecLine(fmt.Sprintf("delete vrf name vrf%d static route %s nexthop %s", vrfId, prefix, nexthop))
		}
		delete(EtcdVrfStaticMap, vrfId)
	}
}

func ProcessVrfUpdate(vrfId int, vrf *VrfsConfig) {
	if vrfConfig, ok := EtcdVrfMap[vrfId]; ok {
		if !reflect.DeepEqual(*vrf, EtcdVrfMap[vrfId]) {
			fmt.Println("Vrf config updated for : ", vrfId)
			ProcessVrfInterfaceChanges(vrfId, &vrfConfig, vrf)
		} else {
			fmt.Println("No change in vrf config for vrfId : ", vrfId, ". Force Resync")
			// EtcdVrfDelete(vrfId, false)
			// EtcdVrfVlanSubinterfacesAdd(vrf)
			// EtcdVrfAddressAdd(vrfId, vrf)
		}
	} else {
		EtcdVrfVlanSubinterfacesAdd(vrf)
		EtcdVrfAddressAdd(vrfId, vrf)
	}
}

func EtcdVrfSync(vrfId int, vrf *VrfsConfig) {
	ExecLine(fmt.Sprintf("set vrf name vrf%d", vrfId))
	if vrf.Pair.Priority == "backup" {
		ExecLine(fmt.Sprintf("set vrf name vrf%d pair-backup", vrfId))
	} else {
		ExecLine(fmt.Sprintf("delete vrf name vrf%d pair-backup", vrfId))
	}
	if vrf.Gateway.Priority == "backup" {
		ExecLine(fmt.Sprintf("set vrf name vrf%d region-backup", vrfId))
	} else {
		ExecLine(fmt.Sprintf("delete vrf name vrf%d region-backup", vrfId))
	}
	Commit()
	ProcessVrfUpdate(vrfId, vrf)
	Commit()
}

var VrfIfDeleteCache = map[string]bool{}

func EtcdVrfCommand(cmd *Command) bool {
	if len(cmd.cmds) == 3 && cmd.cmds[0] == "vrf" && cmd.cmds[1] == "name" {
		return true
	} else {
		return false
	}
}

func EtcdVrfDeleteCacheRegister(cmd *Command) {
	if EtcdVrfCommand(cmd) {
		str := strings.Join(cmd.cmds, " ")
		fmt.Println("EtcdVrfDeleteCacheRegister", str)
		VrfIfDeleteCache[str] = true
	}
}

func EtcdVrfDeleteCacheCheck(str string) bool {
	return VrfIfDeleteCache[str]
}

func EtcdVrfDelete(vrfId int, vrfIfDelete bool) {
	fmt.Println("EtcdVrfDelete:", vrfId)
	if vrfConfig, ok := EtcdVrfMap[vrfId]; ok {
		EtcdVrfAddressClear(vrfId, &vrfConfig)
		EtcdVrfVlanSubinterfacesDelete(&vrfConfig)
		EtcdVrfStaticClear(vrfId)
		if !vrfIfDelete {
			str := fmt.Sprintf("vrf name vrf%d", vrfId)
			delete(VrfIfDeleteCache, str)
		}
		ExecLine(fmt.Sprintf("delete vrf name vrf%d", vrfId))
		Commit()
	} else {
		fmt.Println("EtcdVrfDelete: can't find vrf cache for", vrfId)
	}
}

func (vrf *VrfsConfig) Copy() VrfConfig {
	var vrfConfig VrfConfig
	vrfConfig.Name = vrf.Name
	vrfConfig.VrfId = vrf.Id
	vrfConfig.Rd = vrf.Rd
	vrfConfig.RtImport = vrf.RtImport
	vrfConfig.RtExport = vrf.RtExport
	vrfConfig.RtBoth = vrf.RtBoth
	for _, rib := range vrf.VrfRibs {
		vrfConfig.VrfRibs = append(vrfConfig.VrfRibs, rib)
	}
	for _, hub := range vrf.Hubs {
		vrfConfig.Hubs = append(vrfConfig.Hubs, hub)
	}
	vrfConfig.HubNode = vrf.HubNode
	vrfConfig.Static = vrf.Static
	return vrfConfig
}

func VrfParse(vrfId int, jsonStr string) {
	var jsonIntf interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonIntf)
	if err != nil {
		log.WithFields(log.Fields{
			"json":  jsonStr,
			"error": err,
		}).Error("VrfParse:json.Unmarshal()")
		return
	}

	var vrf VrfsConfig
	err = mapstructure.Decode(jsonIntf, &vrf)
	if err != nil {
		log.WithFields(log.Fields{
			"json-intf": jsonIntf,
			"error":     err,
		}).Error("VrfParse:mapstructure.Decode()")
		return
	}

	// Sort interfce and sub interface by name.
	sort.Sort(vrf.Interfaces)

	// Vrf Sync.
	EtcdVrfSync(vrfId, &vrf)
	DhcpVrfSync(vrfId, &vrf)
	VrrpVrfSync(vrfId, &vrf)
	QuaggaVrfSync(vrfId, &vrf)
	OspfVrfSync(vrfId, &vrf)
	DistributeListSync(vrfId, &vrf)

	// GoBGP VrfConfig
	vrfConfig := vrf.Copy()
	GobgpVrfUpdate(vrfConfig)
	GobgpHubUpdate(vrfConfig)

	// Static route update after GoBGP VRF config.
	EtcdVrfStaticUpdate(vrfId, &vrf)

	EtcdVrfMap[vrfId] = vrf

	fmt.Println("VrfParse ends here")
}

func VrfDelete(vrfId int, vrfIfDelete bool) {
	// GoBGP VRF
	GobgpVrfDelete(vrfId)
	GobgpHubDelete(vrfId)

	// Vrf Sync.
	DhcpVrfDelete(vrfId)
	VrrpVrfDelete(vrfId)
	QuaggaVrfDelete(vrfId)
	OspfVrfDelete(vrfId)
	DistributeListDelete(vrfId)
	EtcdVrfDelete(vrfId, vrfIfDelete)

	delete(EtcdVrfMap, vrfId)

	fmt.Println("VrfDelete ends here")
}

func ClearVrfCache() {
	fmt.Println("Clearing vrf cache")
	for _, vrfConfig := range EtcdVrfMap {
		VrfDelete(vrfConfig.Id, false)
	}
	EtcdVrfMap = map[int]VrfsConfig{}
}
