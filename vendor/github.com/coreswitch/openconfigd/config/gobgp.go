// Copyright 2016, 2017 OpenConfigd Project.
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
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/coreswitch/netutil"
	"github.com/coreswitch/process"
	"github.com/mitchellh/mapstructure"
	"github.com/osrg/gobgp/client"
	bgpconfig "github.com/osrg/gobgp/config"
	"github.com/osrg/gobgp/packet/bgp"
	"github.com/osrg/gobgp/table"
	log "github.com/sirupsen/logrus"
)

const defaultGrpcEndpoint = ":50051"

var (
	gobgpConfig       GobgpConfig
	gobgpRouterId     string
	gobgpGrpcEndpoint string = defaultGrpcEndpoint
)

type GobgpConfig struct {
	Global            bgpconfig.Global             `mapstructure:"global"`
	Neighbors         []bgpconfig.Neighbor         `mapstructure:"neighbors"`
	Zebra             bgpconfig.Zebra              `mapstructure:"zebra"`
	Vrfs              []VrfConfig                  `mapstructure:"vrfs"`
	DefinedSets       bgpconfig.DefinedSets        `mapstructure:"defined_sets"`
	PolicyDefinitions []bgpconfig.PolicyDefinition `mapstructure:"policy_definitions"`
	Ribs              []VrfRib                     `mapstructure:"ribs" json:"ribs,omitempty" json:"ribs"`
	Interfaces        Interfaces                   `mapstructure:"interfaces" json:"interfaces,omitempty"`
}

type VrfConfig struct {
	Name     string   `mapstructure:"name" json:"name,omitempty"`
	VrfId    int      `mapstructure:"vrf-id" json:"vrf-id,omitempty"`
	Rd       string   `mapstructure:"rd" json:"rd,omitempty"`
	RtImport string   `mapstructure:"rt-import" json:"rt-import,omitempty"`
	RtExport string   `mapstructure:"rt-export" json:"rt-export,omitempty"`
	RtBoth   string   `mapstructure:"rt-both" json:"rt-both,omitempty"`
	VrfRibs  []VrfRib `mapstructure:"ribs" json:"ribs,omitempty"`
	Hubs     []Hub    `mapstructure:"hubs" json:"hubs,omitempty"`
	HubNode  string   `mapstructure:"hub_node" json:"hub_node,omitempty"`
	Static   Static
}

type VrfRib struct {
	Prefix  string `mapstructure:"prefix" json:"prefix,omitempty"`
	NextHop string `mapstructure:"next-hop" json:"next-hop,omitempty"`
	Same    bool
}

func (lhs *VrfConfig) Equal(rhs *VrfConfig) bool {
	if lhs == nil || rhs == nil {
		return false
	}
	if lhs.Name != rhs.Name {
		return false
	}
	if lhs.VrfId != rhs.VrfId {
		return false
	}
	if lhs.Rd != rhs.Rd {
		return false
	}
	if lhs.RtImport != rhs.RtImport {
		return false
	}
	if lhs.RtExport != rhs.RtExport {
		return false
	}
	if lhs.RtBoth != rhs.RtBoth {
		return false
	}
	if len(lhs.VrfRibs) != len(rhs.VrfRibs) {
		return false
	}
	for pos, r := range lhs.VrfRibs {
		if !r.Equal(&lhs.VrfRibs[pos]) {
			return false
		}
	}
	if len(lhs.Hubs) != len(rhs.Hubs) {
		return false
	}
	for pos, r := range lhs.Hubs {
		if r.Address != rhs.Hubs[pos].Address {
			return false
		}
	}
	if len(lhs.Static.Route) != len(rhs.Static.Route) {
		return false
	}
	for pos, s := range lhs.Static.Route {
		if !s.Equal(&rhs.Static.Route[pos]) {
			return false
		}
	}
	return true
}

func (lhs *VrfRib) Equal(rhs *VrfRib) bool {
	if lhs.Prefix != rhs.Prefix {
		return false
	}
	if lhs.NextHop != rhs.NextHop {
		return false
	}
	return true
}

func NewGobgpClient() (*client.Client, error) {
	return client.New(gobgpGrpcEndpoint)
}

func GobgpRouterIdRegister(routerId string) {
	if gobgpRouterId == routerId {
		return
	}
	if routerId != "" {
		fmt.Println("[gobgp]Router Id", routerId)
		gobgpRouterId = routerId
	}
}

func GobgpVrfPath(c *VrfRib) (*table.Path, error) {
	attrs := table.PathAttrs(make([]bgp.PathAttributeInterface, 0, 1))

	// Origin.
	typ := bgp.BGP_ORIGIN_ATTR_TYPE_INCOMPLETE
	attrs = append(attrs, bgp.NewPathAttributeOrigin(uint8(typ)))

	// Nexthop.
	attrs = append(attrs, bgp.NewPathAttributeNextHop(c.NextHop))

	// Prefix to NLRI.
	p, err := netutil.ParsePrefix(c.Prefix)
	if err != nil {
		return nil, err
	}
	nlri := bgp.NewIPAddrPrefix(uint8(p.Length), p.IP.String())

	// Return a new Path.
	return table.NewPath(nil, nlri, false, attrs, time.Now(), false), nil
}

func GobgpStaticPath(s *Route) (*table.Path, error) {
	attrs := table.PathAttrs(make([]bgp.PathAttributeInterface, 0, 1))

	// Origin.
	typ := bgp.BGP_ORIGIN_ATTR_TYPE_INCOMPLETE
	attrs = append(attrs, bgp.NewPathAttributeOrigin(uint8(typ)))

	// Nexthop.
	attrs = append(attrs, bgp.NewPathAttributeNextHop("0.0.0.0"))

	// Community.
	attrs = append(attrs, bgp.NewPathAttributeCommunities([]uint32{
		stringToCommunityValue("101:0"),
	}))

	// Prefix to NLRI.
	p, err := netutil.ParsePrefix(s.Prefix)
	if err != nil {
		return nil, err
	}
	nlri := bgp.NewIPAddrPrefix(uint8(p.Length), p.IP.String())

	// Return a new Path.
	return table.NewPath(nil, nlri, false, attrs, time.Now(), false), nil
}

func GobgpClearVrfRib(c *VrfConfig) error {
	client, err := NewGobgpClient()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("GobgpClearVrfRib:client.New()")
		return err
	}
	defer client.Close()

	for _, r := range c.VrfRibs {
		path, err := GobgpVrfPath(&r)
		if err != nil {
			return err
		}
		err = client.DeleteVRFPath(c.Name, []*table.Path{path})
		if err != nil {
			fmt.Println("GobgpClearVrfRib DeleteVRFPath:", err)
		}
	}
	return nil
}

func GobgpAddVrf(client *client.Client, c *VrfConfig) error {
	fmt.Println("GobgpAddVrf", c.Name, "ribs num", len(c.VrfRibs))
	rd, err := bgp.ParseRouteDistinguisher(c.Rd)
	if err != nil {
		return err
	}

	importRt := make([]bgp.ExtendedCommunityInterface, 0)
	exportRt := make([]bgp.ExtendedCommunityInterface, 0)

	if c.RtBoth != "" {
		rt, err := bgp.ParseRouteTarget(c.RtBoth)
		if err != nil {
			fmt.Println("ParseRouteTarget err:", err)
			return err
		}
		importRt = append(importRt, rt)
		exportRt = append(exportRt, rt)
	}
	err = client.AddVRF(c.Name, c.VrfId, rd, importRt, exportRt)
	if err != nil {
		fmt.Println("GobgpAddVrf AddVRF:", err)
		return err
	}

	for _, r := range c.VrfRibs {
		path, err := GobgpVrfPath(&r)
		if err != nil {
			fmt.Println("GobgpVrfPath err", err)
			return err
		}
		_, err = client.AddVRFPath(c.Name, []*table.Path{path})
		if err != nil {
			fmt.Println("GobgpAddVrf AddVRFPath:", err)
		}
		fmt.Println("RIB is added", path)
	}

	for _, s := range c.Static.Route {
		path, err := GobgpStaticPath(&s)
		if err != nil {
			fmt.Println("GobgpStaticPath err", err)
			return err
		}
		_, err = client.AddVRFPath(c.Name, []*table.Path{path})
		if err != nil {
			fmt.Println("GobgpAddVrf AddVRFPath static:", err)
		}
		fmt.Println("Static is added", path)
	}

	return nil
}

func GobgpAddPolicyDefinition(client *client.Client, cfg bgpconfig.PolicyDefinition) {
	policy, err := table.NewPolicy(cfg)
	if err != nil {
		fmt.Println("GobgpAddPolicyDefinition NewPolicy():", err)
		return
	}
	for _, st := range policy.Statements {
		err = client.AddStatement(st)
		if err != nil {
			fmt.Println("GobgpAddPolicyDefinition AddStatement():", err)
		}
	}
	err = client.AddPolicy(policy, true)
	if err != nil {
		fmt.Println("GobgpAddPolicyDefinition AddPolicy():", err)
	}
}

func GobgpDeletePolicyDefinition(client *client.Client, cfg bgpconfig.PolicyDefinition) {
	policy, err := table.NewPolicy(cfg)
	if err != nil {
		fmt.Println("GobgpDeletePolicyDefinition NewPolicy():", err)
		return
	}
	for _, st := range policy.Statements {
		err = client.DeleteStatement(st, true)
		if err != nil {
			fmt.Println("GobgpDeletePolicyDefinition DeleteStatement():", err)
		}
	}
	err = client.DeletePolicy(policy, true, false)
	if err != nil {
		fmt.Println("GobgpDeletePolicyDefinition DeletePolicy():", err)
	}
}

func GobgpAddGlobalPolicy(client *client.Client, direction string, policyNames []string) error {
	assign := &table.PolicyAssignment{}

	switch direction {
	case "in":
		assign.Type = table.POLICY_DIRECTION_IN
	case "import":
		assign.Type = table.POLICY_DIRECTION_IMPORT
	case "export":
		assign.Type = table.POLICY_DIRECTION_EXPORT
	}

	ps := make([]*table.Policy, 0, len(policyNames))
	for _, name := range policyNames {
		ps = append(ps, &table.Policy{Name: name})
	}
	assign.Policies = ps
	assign.Default = table.ROUTE_TYPE_ACCEPT

	err := client.AddPolicyAssignment(assign)
	if err != nil {
		fmt.Println("GobgpAddGlobalPolicy:", err)
		return err
	}
	return nil
}

func GobgpDeleteGlobalPolicy(client *client.Client, direction string, policyNames []string) error {
	assign := &table.PolicyAssignment{}

	switch direction {
	case "in":
		assign.Type = table.POLICY_DIRECTION_IN
	case "import":
		assign.Type = table.POLICY_DIRECTION_IMPORT
	case "export":
		assign.Type = table.POLICY_DIRECTION_EXPORT
	}

	ps := make([]*table.Policy, 0, len(policyNames))
	for _, name := range policyNames {
		ps = append(ps, &table.Policy{Name: name})
	}
	assign.Policies = ps
	assign.Default = table.ROUTE_TYPE_ACCEPT

	err := client.DeletePolicyAssignment(assign, false)
	if err != nil {
		fmt.Println("GobgpDeleteGlobalPolicy", err)
		return err
	}
	return nil
}

func GobgpSetGlobal(client *client.Client, cfg *GobgpConfig) {
	fmt.Println("GobgpSetGlobal", cfg.Global.Config)
	err := client.StartServer(&cfg.Global)
	if err != nil {
		fmt.Println("GobgpSetGlobal:", err)
	}
}

func GobgpSetZebraRoutine() error {
	client, err := NewGobgpClient()
	if err != nil {
		return err
	}
	defer client.Close()

	zebra := &bgpconfig.Zebra{}
	zebra.Config.Enabled = true
	zebra.Config.Url = "unix:/var/run/zserv.api"
	zebra.Config.Version = 3
	err = client.EnableZebra(zebra)
	if err != nil {
		return err
	}
	return nil
}

var GobgpZebraRetry bool

func GobgpSetZebra(client *client.Client, cfg *GobgpConfig, version uint8) {
	zebra := &bgpconfig.Zebra{}
	zebra.Config.Enabled = true
	zebra.Config.Url = "unix:/var/run/zserv.api"
	zebra.Config.Version = version
	zebra.Config.RedistributeRouteTypeList = []bgpconfig.InstallProtocolType{"bgp", "ospf"}
	err := client.EnableZebra(zebra)
	if err != nil {
		if strings.Contains(err.Error(), "zserv") && !GobgpZebraRetry {
			GobgpZebraRetry = true
			go func() {
				defer func() {
					GobgpZebraRetry = false
				}()
				for {
					time.Sleep(time.Second * 3)
					err := GobgpSetZebraRoutine()
					if err == nil {
						fmt.Println("GobgpSetZebra: retry success!")
						GobgpReset(&gobgpConfig)
						return
					}
					if !strings.Contains(err.Error(), "zserv") {
						fmt.Println("GobgpSetZebra: retry other error", err)
						return
					}
				}
			}()
		}
	} else {
		fmt.Println("GobgpSetZebra: Success at first attempt!")
	}
}

func GobgpSetVrf(client *client.Client, cfg *GobgpConfig) {
	for _, v := range cfg.Vrfs {
		GobgpAddVrf(client, &v)
	}
}

func GobgpSetNeighbor(client *client.Client, cfg *GobgpConfig) {
	for _, n := range cfg.Neighbors {
		n.GracefulRestart.Config.RestartTime = 1
		err := client.AddNeighbor(&n)
		if err != nil {
			fmt.Println("GobgpSetNeighbor:", err)
		}
	}
}

func GobgpSoftresetNeighbor(client *client.Client, cfg *GobgpConfig) {
	err := client.SoftReset("", 0)
	if err != nil {
		fmt.Println("GobgpSoftresetNeighbor:", err)
	}
}

func GobgpSetDefinedSet(client *client.Client, cfg *GobgpConfig) {
	for _, nset := range cfg.DefinedSets.NeighborSets {
		tnset, _ := table.NewNeighborSet(nset)
		if tnset != nil {
			client.AddDefinedSet(tnset)
		}
	}
	for _, cset := range cfg.DefinedSets.BgpDefinedSets.CommunitySets {
		tcset, _ := table.NewCommunitySet(cset)
		if tcset != nil {
			client.AddDefinedSet(tcset)
		}
	}
	for _, eset := range cfg.DefinedSets.BgpDefinedSets.ExtCommunitySets {
		teset, _ := table.NewExtCommunitySet(eset)
		if teset != nil {
			client.AddDefinedSet(teset)
		}
	}
}

func GobgpSetPolicyDefinition(client *client.Client, cfg *GobgpConfig) {
	for _, p := range cfg.PolicyDefinitions {
		GobgpAddPolicyDefinition(client, p)
	}
}

func GobgpSetGlobalPolicy(client *client.Client, cfg *GobgpConfig) error {
	config := &cfg.Global.ApplyPolicy.Config
	if len(config.ExportPolicyList) > 0 {
		GobgpAddGlobalPolicy(client, "export", config.ExportPolicyList)
	}
	if len(config.ImportPolicyList) > 0 {
		GobgpAddGlobalPolicy(client, "import", config.ImportPolicyList)
	}
	if len(config.InPolicyList) > 0 {
		GobgpAddGlobalPolicy(client, "in", config.InPolicyList)
	}
	return nil
}

func GobgpClearGlobal(client *client.Client) {
	err := client.StopServer()
	if err != nil {
		fmt.Println("GobgpClearGlobal:", err)
		return
	}
}

func GobgpClearVrf(client *client.Client) {
	for _, cfg := range gobgpConfig.Vrfs {
		client.DeleteVRF(cfg.Name)
	}
}

func GobgpClearNeighbor(client *client.Client) {
	for _, cfg := range gobgpConfig.Neighbors {
		client.DeleteNeighbor(&cfg)
	}
}

func GobgpClearDefinedSet(client *client.Client) {
	cfg := &gobgpConfig.DefinedSets
	for _, nset := range cfg.NeighborSets {
		tnset, _ := table.NewNeighborSet(nset)
		if tnset != nil {
			client.DeleteDefinedSet(tnset, true)
		}
	}
	for _, cset := range cfg.BgpDefinedSets.CommunitySets {
		tcset, _ := table.NewCommunitySet(cset)
		if tcset != nil {
			client.DeleteDefinedSet(tcset, true)
		}
	}
	for _, eset := range cfg.BgpDefinedSets.ExtCommunitySets {
		teset, _ := table.NewExtCommunitySet(eset)
		if teset != nil {
			client.DeleteDefinedSet(teset, true)
		}
	}
}

func GobgpClearPolicyDefinition(client *client.Client) {
	for _, p := range gobgpConfig.PolicyDefinitions {
		GobgpDeletePolicyDefinition(client, p)
	}
}

func GobgpClearGlobalPolicy(client *client.Client) {
	cfg := &gobgpConfig.Global.ApplyPolicy.Config
	if len(cfg.ExportPolicyList) > 0 {
		GobgpDeleteGlobalPolicy(client, "export", cfg.ExportPolicyList)
	}
	if len(cfg.ImportPolicyList) > 0 {
		GobgpDeleteGlobalPolicy(client, "import", cfg.ImportPolicyList)
	}
	if len(cfg.InPolicyList) > 0 {
		GobgpDeleteGlobalPolicy(client, "in", cfg.InPolicyList)
	}
}

func GobgpClearAll() {
	client, err := NewGobgpClient()
	if err != nil {
		fmt.Println("GobgpStopServer:", err)
		return
	}
	defer client.Close()

	GobgpClearGlobalPolicy(client)
	GobgpClearPolicyDefinition(client)
	GobgpClearDefinedSet(client)
	GobgpClearNeighbor(client)
	GobgpClearVrf(client)
	GobgpClearGlobal(client)
}

func GobgpUpdateNeighbor(client *client.Client, cfg *GobgpConfig) {
	type GobgpNeighborCache struct {
		Same     bool
		Neighbor bgpconfig.Neighbor
	}

	// Cache existing neighbor.
	neighborCache := map[string]*GobgpNeighborCache{}
	for _, n := range gobgpConfig.Neighbors {
		neighborCache[n.Config.NeighborAddress] = &GobgpNeighborCache{Same: false, Neighbor: n}
	}

	// Mark same neighbor.
	for _, n := range cfg.Neighbors {
		if exist, ok := neighborCache[n.Config.NeighborAddress]; ok {
			if exist.Neighbor.Equal(&n) {
				exist.Same = true
			}
		}
	}

	// Purge not same neighbor.
	for _, v := range neighborCache {
		if !v.Same {
			client.DeleteNeighbor(&v.Neighbor)
		}
	}

	// Update neighbor configuration.
	for _, n := range cfg.Neighbors {
		if exist, ok := neighborCache[n.Config.NeighborAddress]; ok {
			if !exist.Same {
				client.AddNeighbor(&n)
			}
		} else {
			client.AddNeighbor(&n)
		}
	}
}

func GobgpUpdateVrf(client *client.Client, cfg *GobgpConfig) {
	type GobgpVrfCache struct {
		Same      bool
		VrfConfig VrfConfig
	}

	// Cache existing VRF.
	vrfCache := map[string]*GobgpVrfCache{}
	for _, v := range gobgpConfig.Vrfs {
		vrfCache[v.Name] = &GobgpVrfCache{Same: false, VrfConfig: v}
	}

	// Mark same VRF.
	for _, v := range cfg.Vrfs {
		if exist, ok := vrfCache[v.Name]; ok {
			if exist.VrfConfig.Equal(&v) {
				exist.Same = true
			}
		}
	}
	// Purge not same VRF.
	for _, v := range vrfCache {
		if !v.Same {
			client.DeleteVRF(v.VrfConfig.Name)
		}
	}

	// Update VRF configuration.
	for _, v := range cfg.Vrfs {
		if exist, ok := vrfCache[v.Name]; ok {
			if !exist.Same {
				GobgpAddVrf(client, &v)
			} else {
				fmt.Println("Skip adding same config VRF")
			}
		} else {
			GobgpAddVrf(client, &v)
		}
	}
}

func GobgpUpdate(cfg *GobgpConfig) error {
	fmt.Println("Updating configuration")
	client, err := NewGobgpClient()
	if err != nil {
		return err
	}
	defer client.Close()

	// Clear
	GobgpClearGlobalPolicy(client)
	GobgpClearPolicyDefinition(client)
	GobgpClearDefinedSet(client)

	// Set and update.
	GobgpUpdateVrf(client, cfg)
	GobgpUpdateNeighbor(client, cfg)
	GobgpSetDefinedSet(client, cfg)
	GobgpSetPolicyDefinition(client, cfg)
	GobgpSetGlobalPolicy(client, cfg)
	GobgpSetZebra(client, cfg, 3)

	// Soft reset all of neighbors to reflect policy change.
	GobgpSoftresetNeighbor(client, cfg)

	return nil
}

func GobgpReset(cfg *GobgpConfig) error {
	client, err := NewGobgpClient()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("GobgpReset:client.New()")
		return err
	}
	defer client.Close()

	// Clear all
	GobgpClearAll()

	// Set.
	GobgpSetGlobal(client, cfg)
	GobgpSetVrf(client, cfg)
	GobgpSetNeighbor(client, cfg)
	GobgpSetDefinedSet(client, cfg)
	GobgpSetPolicyDefinition(client, cfg)
	GobgpSetGlobalPolicy(client, cfg)
	GobgpSetZebra(client, cfg, 3)

	return nil
}

// Map for configured neighbor.
var (
	GobgpNeighborMap   = map[string]bgpconfig.Neighbor{}
	GobgpNeighborCache GobgpConfig
	GobgpNeighborMutex sync.Mutex
	GobgpNeighborDelay time.Duration = 3
	GobgpAfterFunc     *time.Timer
)

type NeighborConfig struct {
	Neighbor bgpconfig.Neighbor `json:"neighbor"`
}

func GobgpNeighborUpdate(ncfg *bgpconfig.Neighbor, add bool) {
	var cfg GobgpConfig
	if GobgpAfterFunc == nil {
		cfg = gobgpConfig
		cfg.Neighbors = make([]bgpconfig.Neighbor, 0)
		for _, n := range gobgpConfig.Neighbors {
			if n.Config.NeighborAddress != ncfg.Config.NeighborAddress {
				cfg.Neighbors = append(cfg.Neighbors, n)
			}
		}
	} else {
		cfg = GobgpNeighborCache
		cfg.Neighbors = make([]bgpconfig.Neighbor, 0)
		for _, n := range GobgpNeighborCache.Neighbors {
			if n.Config.NeighborAddress != ncfg.Config.NeighborAddress {
				cfg.Neighbors = append(cfg.Neighbors, n)
			}
		}
	}

	if add {
		cfg.Neighbors = append(cfg.Neighbors, *ncfg)
	}

	if cfg.Global.Equal(&gobgpConfig.Global) {
		GobgpNeighborCache = cfg
		if GobgpAfterFunc == nil {
			GobgpAfterFunc = time.AfterFunc(time.Second*GobgpNeighborDelay, func() {
				GobgpNeighborMutex.Lock()
				defer GobgpNeighborMutex.Unlock()

				GobgpUpdate(&GobgpNeighborCache)
				gobgpConfig = GobgpNeighborCache
				GobgpAfterFunc = nil
			})
		}
	} else {
		GobgpReset(&cfg)
		gobgpConfig = cfg
	}
}

func GobgpNeighborAdd(id string, jsonStr string) {
	GobgpNeighborMutex.Lock()
	defer GobgpNeighborMutex.Unlock()

	var jsonIntf interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonIntf)
	if err != nil {
		log.WithFields(log.Fields{
			"json":  jsonStr,
			"error": err,
		}).Error("GobgpNeighborAdd:json.Unmarshal()")
		return
	}

	var ncfgWrapper NeighborConfig
	err = mapstructure.Decode(jsonIntf, &ncfgWrapper)
	if err != nil {
		log.WithFields(log.Fields{
			"json-intf": jsonIntf,
			"error":     err,
		}).Error("GobgpNeighborAdd:mapstructure.Decode()")
		return
	}

	var ncfg bgpconfig.Neighbor
	ncfg = ncfgWrapper.Neighbor

	// Update neighbor map.
	GobgpNeighborMap[id] = ncfg
	GobgpNeighborUpdate(&ncfg, true)
}

func GobgpNeighborDelete(id string) {
	GobgpNeighborMutex.Lock()
	defer GobgpNeighborMutex.Unlock()

	ncfg, ok := GobgpNeighborMap[id]
	if !ok {
		fmt.Println("GobgpNeighborDelete: can't find neighbor with id", id)
		return
	}

	// Update neighbor map.
	GobgpNeighborUpdate(&ncfg, false)
	delete(GobgpNeighborMap, id)
}

func GobgpParse(jsonStr string) {
	var jsonIntf interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonIntf)
	if err != nil {
		log.WithFields(log.Fields{
			"json":  jsonStr,
			"error": err,
		}).Error("GobgpParse:json.Unmarshal()")
		return
	}

	var cfg GobgpConfig
	err = mapstructure.Decode(jsonIntf, &cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"json-intf": jsonIntf,
			"error":     err,
		}).Error("GobgpParse:mapstructure.Decode()")
		return
	}

	// Copy existing VRF config.
	for _, vrf := range gobgpConfig.Vrfs {
		cfg.Vrfs = append(cfg.Vrfs, vrf)
	}

	// Copy existing neighbor config.
	for _, nei := range GobgpNeighborMap {
		cfg.Neighbors = append(cfg.Neighbors, nei)
	}

	// Router ID register.
	GobgpRouterIdRegister(cfg.Global.Config.RouterId)

	if cfg.Global.Equal(&gobgpConfig.Global) {
		GobgpUpdate(&cfg)
	} else {
		GobgpReset(&cfg)
	}

	gobgpConfig = cfg
}

func GobgpVrfEqual(ribs []VrfRib, rib VrfRib) bool {
	for pos, _ := range ribs {
		fmt.Println("Comparing", ribs[pos], rib)
		if ribs[pos].Equal(&rib) {
			ribs[pos].Same = true
			fmt.Println("Equal", ribs[pos], rib)
			return true
		}
	}
	return false
}

func GobgpVrfRibSync(name string, old, new []VrfRib) {
	fmt.Println("------")
	for _, x := range old {
		fmt.Println("o:", x)
	}
	fmt.Println("------")
	for _, x := range new {
		fmt.Println("n:", x)
	}
	fmt.Println("------")

	client, err := NewGobgpClient()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("GobgpVrfRibSync:client.New()")
		return
	}
	defer client.Close()

	for pos, _ := range old {
		old[pos].Same = false
	}
	for _, nrib := range new {
		if !GobgpVrfEqual(old, nrib) {
			path, err := GobgpVrfPath(&nrib)
			if err != nil {
				fmt.Println("GobgpVrfRibSync err", err)
			} else {
				fmt.Println("Add:path", path)
				_, err = client.AddVRFPath(name, []*table.Path{path})
				if err != nil {
					fmt.Println("GobgpVrfRibSync err:", err)
				}
			}
		}
	}
	for pos, orib := range old {
		if !old[pos].Same {
			path, err := GobgpVrfPath(&orib)
			if err != nil {
				fmt.Println("GobgpVrfRibSync err", err)
			} else {
				fmt.Println("Del:path", path)
				err = client.DeleteVRFPath(name, []*table.Path{path})
				if err != nil {
					fmt.Println("GobgpVrfRibSync err:", err)
				}
			}
		}
	}
	fmt.Println("------")
	for _, x := range old {
		fmt.Println("o:", x)
	}
	fmt.Println("------")
	for _, x := range new {
		fmt.Println("n:", x)
	}
	fmt.Println("------")
}

func GobgpVrfUpdate(vrfConfig VrfConfig) {
	fmt.Println("GobgpVrfUpdate Start")
	var cfg GobgpConfig

	cfg = gobgpConfig
	cfg.Vrfs = make([]VrfConfig, 0)

	for _, vrf := range gobgpConfig.Vrfs {
		if vrf.VrfId == vrfConfig.VrfId {
			GobgpVrfRibSync(vrf.Name, vrf.VrfRibs, vrfConfig.VrfRibs)
		} else {
			cfg.Vrfs = append(cfg.Vrfs, vrf)
		}
	}
	cfg.Vrfs = append(cfg.Vrfs, vrfConfig)

	if cfg.Global.Equal(&gobgpConfig.Global) {
		GobgpUpdate(&cfg)
	} else {
		GobgpReset(&cfg)
	}

	gobgpConfig = cfg
	fmt.Println("GobgpVrfUpdate End")
}

func GobgpVrfDelete(vrfId int) {
	var cfg GobgpConfig

	cfg = gobgpConfig
	cfg.Vrfs = make([]VrfConfig, 0)
	for _, vrf := range gobgpConfig.Vrfs {
		if vrf.VrfId == vrfId {
			fmt.Println("GobgpVrfDelete: removing vrf", vrfId)
			//GobgpClearVrfRib(&vrf)
		} else {
			cfg.Vrfs = append(cfg.Vrfs, vrf)
		}
	}

	if cfg.Global.Equal(&gobgpConfig.Global) {
		GobgpUpdate(&cfg)
	} else {
		GobgpReset(&cfg)
	}

	gobgpConfig = cfg
}

var GobgpHubProcessList = map[int]*process.Process{}
var GobgpHubNode = map[int]string{}

func GobgpHubUpdate(vrfConfig VrfConfig) {
	GobgpHubDelete(vrfConfig.VrfId)
	if len(vrfConfig.Hubs) < 2 {
		return
	}
	for _, rib := range vrfConfig.VrfRibs {
		for _, h := range vrfConfig.Hubs {
			if rib.NextHop == h.Address {
				fmt.Println("This is Hub node so do not run Hub monitoring")
				if vrfConfig.HubNode != "" {
					GobgpHubNode[vrfConfig.VrfId] = vrfConfig.HubNode
					fmt.Println("Setting hub-node", vrfConfig.HubNode)
					ExecLine(fmt.Sprintf("set vrf name vrf%d hub-node %s", vrfConfig.VrfId, vrfConfig.HubNode))
					Commit()
				}
				return
			}
		}
	}
	var hubs []string
	for _, h := range vrfConfig.Hubs {
		hubs = append(hubs, h.Address)
	}
	args := []string{
		"-v", fmt.Sprintf("vrf%d", vrfConfig.VrfId),
		"-r", vrfConfig.Rd,
		"-u", strings.Join(hubs, ":"),
	}
	fmt.Println("gobgp_hub", args)

	proc := process.NewProcess("gobgp_hub", args...)
	proc.StartTimer = 3
	GobgpHubProcessList[vrfConfig.VrfId] = proc
	process.ProcessRegister(proc)
}

func GobgpHubDelete(vrfId int) {
	if proc, ok := GobgpHubProcessList[vrfId]; ok {
		process.ProcessUnregister(proc)
		delete(GobgpHubProcessList, vrfId)
	}
	if hubNode, ok := GobgpHubNode[vrfId]; ok {
		ExecLine(fmt.Sprintf("delete vrf name vrf%d hub-node %s", vrfId, hubNode))
		Commit()
	}

	delete(GobgpHubNode, vrfId)
}

// configure# clear gobgp
func GobgpClearApi(Args []string) (inst int, instStr string) {
	inst = CliSuccess
	GobgpClearAll()
	gobgpConfig = GobgpConfig{}
	return
}

// configure# reset gobgp
func GobgpResetApi(Args []string) (inst int, instStr string) {
	inst = CliSuccess
	GobgpReset(&gobgpConfig)
	return
}

// GoBGP WAN Process
var GobgpWanProcess *process.Process

// GoBGP WAN Config.
var gobgpWanConfigLocal GobgpConfig
var gobgpWanConfigGohan GobgpConfig

func GobgpSetRib(client *client.Client, cfg *GobgpConfig) {
	for _, r := range cfg.Ribs {
		path, err := GobgpVrfPath(&r)
		if err != nil {
			fmt.Println("GobgpSetRib GobgpVrfPath:", err)
			return
		}
		_, err = client.AddPath([]*table.Path{path})
		if err != nil {
			fmt.Println("GobgpSetRib AddPath:", err)
		}
	}
}

// GoBGP WAN configure
func GobgpWanConfig(cfg *GobgpConfig) {
	fmt.Println(cfg.Global.Config)

	args := []string{
		"--api-hosts=:50052",
		"--pprof-disable",
	}
	GobgpWanProcess = process.NewProcess("gobgpd", args...)
	process.ProcessRegister(GobgpWanProcess)

	var c *client.Client
	for i := 0; i < 10; i++ {
		var err error
		c, err = client.New(":50052")
		if err == nil {
			break
		}
		fmt.Println("GobgpWanConfig err:", err)
		time.Sleep(time.Second * 1)
	}
	if c == nil {
		fmt.Println("GobgpWanConfig retry count exceed")
		return
	}
	defer c.Close()

	GobgpSetGlobal(c, cfg)
	GobgpSetNeighbor(c, cfg)
	GobgpSetDefinedSet(c, cfg)
	GobgpSetPolicyDefinition(c, cfg)
	GobgpSetGlobalPolicy(c, cfg)
	GobgpSetRib(c, cfg)
	GobgpSetZebra(c, cfg, 2)
}

// GoBGP WAN
func GobgpWanParse(jsonStr string, local bool) {
	if GobgpWanProcess != nil {
		process.ProcessUnregister(GobgpWanProcess)
		GobgpWanProcess = nil
	}

	var jsonIntf interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonIntf)
	if err != nil {
		log.WithFields(log.Fields{
			"json":  jsonStr,
			"error": err,
		}).Error("GobgpWanParse:json.Unmarshal()")
		return
	}

	var cfg GobgpConfig
	err = mapstructure.Decode(jsonIntf, &cfg)
	if err != nil {
		log.WithFields(log.Fields{
			"json-intf": jsonIntf,
			"error":     err,
		}).Error("GobgpWanParse:mapstructure.Decode()")
		return
	}

	if local {
		gobgpWanConfigLocal = cfg
	} else {
		gobgpWanConfigGohan = cfg
	}

	if gobgpWanConfigGohan.Global.Config.As != 0 {
		GobgpWanConfig(&gobgpWanConfigGohan)
	} else if gobgpWanConfigLocal.Global.Config.As != 0 {
		GobgpWanConfig(&gobgpWanConfigLocal)
	}
}

func GobgpWanStop(local bool) {
	if GobgpWanProcess != nil {
		process.ProcessUnregister(GobgpWanProcess)
		GobgpWanProcess = nil
	}
	if local {
		gobgpWanConfigLocal.Global.Config.As = 0
	} else {
		gobgpWanConfigGohan.Global.Config.As = 0
	}

	if gobgpWanConfigGohan.Global.Config.As != 0 {
		GobgpWanConfig(&gobgpWanConfigGohan)
	} else if gobgpWanConfigLocal.Global.Config.As != 0 {
		GobgpWanConfig(&gobgpWanConfigLocal)
	}
}

func GobgpWanExit() {
	if GobgpWanProcess != nil {
		process.ProcessUnregister(GobgpWanProcess)
		GobgpWanProcess = nil
	}
}

func stringToCommunityValue(comStr string) uint32 {
	elem := strings.Split(comStr, ":")
	asn, _ := strconv.ParseUint(elem[0], 10, 16)
	val, _ := strconv.ParseUint(elem[1], 10, 16)
	return uint32(asn<<16 | val)
}

func ConfigureGobgpGrpcEndpointApi(set bool, args []interface{}) {
	if len(args) != 1 {
		return
	}
	grpcEndPoint := args[0].(string)
	if set {
		SetGobgpGrpcEndpoint(grpcEndPoint)
	} else {
		ClearGobgpGrpcEndpoint()
	}
}

func SetGobgpGrpcEndpoint(grpcEndpoint string) {
	gobgpGrpcEndpoint = grpcEndpoint
}

func ClearGobgpGrpcEndpoint() {
	gobgpGrpcEndpoint = defaultGrpcEndpoint
}
