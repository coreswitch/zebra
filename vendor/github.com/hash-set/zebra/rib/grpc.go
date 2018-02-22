// Copyright 2016 Zebra Project.
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
	"io"
	"math/rand"
	"time"

	"github.com/coreswitch/cmd"
	"github.com/coreswitch/component"
	rpc "github.com/hash-set/openconfigd/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var cmdSpec = `
[
    {
        "name": "show_interface",
        "line": "show interface (:ribd:interface|)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Interface information",
            "Interface name"
        ]
    },
    {
        "name": "show_interface_vrf",
        "line": "show interface vrf :ribd:vrf (:ribd:vrf:$3:interface|)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Interface information",
            "VRF",
            "VRF name",
            "Interface name"
        ]
    },
    {
        "name": "show_ip_interface_brief",
        "line": "show ip interface brief",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Internet Protocol (IP)",
            "Interface information",
            "Brief"
        ]
    },
    {
        "name": "show_ipv6_interface_brief",
        "line": "show ipv6 interface brief",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Internet Protocol verion 6 (IPv6)",
            "Interface information",
            "Brief"
        ]
    },
    {
        "name": "show_ip_route",
        "line": "show ip route",
        "mode": "exec",
        "helps": [
            "Show running system information",
			      "Internet Protocol (IP)",
			      "IP routing table"
        ]
    },
    {
        "name": "show_ip_route_type",
        "line": "show ip route (kernel|connected|rip|ospf|bgp|isis)",
        "mode": "exec",
        "helps": [
            "Show running system information",
			"Internet Protocol (IP)",
			"IP routing table",
            "Kernel routes",
            "Connected routes",
            "RIP routes",
            "OSPF routes",
            "BGP routes",
            "IS-IS routes"
        ]
    },
    {
        "name": "show_ip_route_database",
        "line": "show ip route database",
        "mode": "exec",
        "helps": [
            "Show running system information",
			"Internet Protocol (IP)",
			"IP routing table",
            "IP routing table database"
        ]
    },
    {
        "name": "show_ip_route_vrf",
        "line": "show ip route vrf :ribd:vrf",
        "mode": "exec",
        "helps": [
            "Show running system information",
			"Internet Protocol (IP)",
			"IP routing table",
            "VRF",
            "VRF"
        ]
    },
    {
        "name": "show_ip_route_vrf_type",
        "line": "show ip route vrf :ribd:vrf (kernel|connected|rip|ospf|bgp|isis)",
        "mode": "exec",
        "helps": [
            "Show running system information",
			"Internet Protocol (IP)",
			"IP routing table",
            "VRF",
            "VRF",
            "Kernel routes",
            "Connected routes",
            "RIP routes",
            "OSPF routes",
            "BGP routes",
            "IS-IS routes"
        ]
    },
    {
        "name": "show_ip_route_vrf_database",
        "line": "show ip route vrf :ribd:vrf database",
        "mode": "exec",
        "helps": [
            "Show running system information",
			"Internet Protocol (IP)",
			"IP routing table",
            "VRF",
            "VRF",
            "IP routing table database"
        ]
    },
    {
        "name": "show_ipv6_route",
        "line": "show ipv6 route",
        "mode": "exec",
        "helps": [
            "Show running system information",
			      "Internet Protocol version 6 (IPv6)",
			      "IP routing table"
        ]
    },
    {
        "name": "show_ipv6_route_database",
        "line": "show ipv6 route database",
        "mode": "exec",
        "helps": [
            "Show running system information",
			      "Internet Protocol version 6 (IPv6)",
			      "IP routing table",
			      "IP routing table database"
        ]
    },
    {
        "name": "show_ipv6_route_vrf",
        "line": "show ipv6 route vrf :ribd:vrf",
        "mode": "exec",
        "helps": [
            "Show running system information",
			"Internet Protocol version 6 (IPv6)",
			"IP routing table",
            "VRF",
            "VRF"
        ]
    },
    {
        "name": "show_ipv6_route_vrf_database",
        "line": "show ipv6 route vrf :ribd:vrf database",
        "mode": "exec",
        "helps": [
            "Show running system information",
			"Internet Protocol version 6 (IPv6)",
			"IP routing table",
            "VRF",
            "VRF",
			"IP routing table database"
        ]
    },
    {
        "name": "show_router_id",
        "line": "show router-id",
        "mode": "exec",
        "helps": [
            "Show running system information",
			      "Router ID"
        ]
    }
]
`

var cmdNameMap = map[string]func(*ShowTask, []interface{}){
	"show_interface":               ShowInterface,
	"show_interface_vrf":           ShowInterfaceVrf,
	"show_ip_interface_brief":      ShowIpInterfaceBrief,
	"show_ipv6_interface_brief":    ShowIpv6InterfaceBrief,
	"show_ip_route":                ShowIpRoute,
	"show_ip_route_type":           ShowIpRouteType,
	"show_ip_route_database":       ShowIpRouteDatabase,
	"show_ip_route_vrf":            ShowIpRouteVrf,
	"show_ip_route_vrf_type":       ShowIpRouteVrfType,
	"show_ip_route_vrf_database":   ShowIpRouteVrfDatabase,
	"show_ipv6_route":              ShowIpv6Route,
	"show_ipv6_route_database":     ShowIpv6RouteDatabase,
	"show_ipv6_route_vrf":          ShowIpv6RouteVrf,
	"show_ipv6_route_vrf_database": ShowIpv6RouteVrfDatabase,
	"show_router_id":               ShowRouterId,
}

const (
	CliSuccess = iota
	CliSuccessExec
	CliSuccessShow
)

var ShowParser *cmd.Node

func rpcRegisterModule(conn *grpc.ClientConn) error {
	client := rpc.NewRegisterClient(conn)
	req := &rpc.RegisterModuleRequest{
		Module: "ribd",
		Port:   "2601",
	}
	_, err := client.DoRegisterModule(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func ShowInterface(t *ShowTask, Args []interface{}) {
	if len(Args) > 0 {
		ifname := Args[0].(string)
		t.Str = InterfaceShow(t.Json, ifname)
	} else {
		t.Str = InterfaceShow(t.Json)
	}
}

func ShowInterfaceVrf(t *ShowTask, Args []interface{}) {
	name := Args[0].(string)
	vrf := VrfLookupByName(name)
	if vrf == nil {
		t.Str = "% VRF does not exists"
		return
	}
	if len(Args) > 1 {
		ifname := Args[1].(string)
		t.Str = vrf.InterfaceShow(t.Json, ifname)
	} else {
		t.Str = vrf.InterfaceShow(t.Json)
	}
}

func ShowIpInterfaceBrief(t *ShowTask, Args []interface{}) {
	t.Str = InterfaceShowBrief(AFI_IP)
	return
}

func ShowIpv6InterfaceBrief(t *ShowTask, Args []interface{}) {
	t.Str = InterfaceShowBrief(AFI_IP6)
}

func ShowRouterId(t *ShowTask, Args []interface{}) {
	t.Str = VrfDefault().RouterIdShow()
}

func rpcRegisterCli(conn *grpc.ClientConn) {
	client := rpc.NewRegisterClient(conn)

	var clis []rpc.RegisterRequest
	json.Unmarshal([]byte(cmdSpec), &clis)

	ShowParser = cmd.NewParser()

	for _, cli := range clis {
		cli.Module = "ribd"
		cli.Privilege = 1
		cli.Code = rpc.ExecCode_REDIRECT_SHOW

		_, err := client.DoRegister(context.Background(), &cli)
		if err != nil {
			grpclog.Fatalf("client DoRegister failed: %v", err)
		}
		ShowParser.InstallLine(cli.Line, cmdNameMap[cli.Name])
	}
}

func RpcConfig(conf *rpc.ConfigReply) {
	Config(int(conf.Type), conf.Path)
}

var Stream rpc.Config_DoConfigClient

func ConfigPush(path []string) {
	if Stream != nil {
		msg := &rpc.ConfigRequest{
			Type: rpc.ConfigType_SET,
			Path: path,
		}
		err := Stream.Send(msg)
		if err != nil {
			return
		}
	}
}

func ConfigPull(path []string) {
	if Stream != nil {
		fmt.Println("ConfigPull", path)
		msg := &rpc.ConfigRequest{
			Type: rpc.ConfigType_DELETE,
			Path: path,
		}
		err := Stream.Send(msg)
		if err != nil {
			return
		}
	}
}

func rpcSubscribe(conn *grpc.ClientConn) error {
	defer func() {
		Stream = nil
		conn.Close()
	}()

	client := rpc.NewConfigClient(conn)
	stream, err := client.DoConfig(context.Background())
	if err != nil {
		return err
	}
	Stream = stream

	path := []string{"vrf", "vlans", "interfaces", "routing-options"}
	msg := &rpc.ConfigRequest{
		Type:   rpc.ConfigType_SUBSCRIBE_MULTI,
		Module: "ribd",
		Port:   2601,
		Path:   path,
	}
	err = stream.Send(msg)
	if err != nil {
		return err
	}

	InterfaceConfigPush()

	validating := false
loop:
	for {
		conf, err := stream.Recv()
		if err == io.EOF {
			break loop
		}
		if err != nil {
			break loop
		}
		switch conf.Type {
		case rpc.ConfigType_VALIDATE_START:
			validating = true
		case rpc.ConfigType_VALIDATE_END:
			// To avoid commit fail with timeout,
			// when openconfigd is configured with two-phase-commit option.
			result := &rpc.ConfigRequest{}
			result.Type = rpc.ConfigType_VALIDATE_SUCCESS
			stream.Send(result)
			validating = false
		case rpc.ConfigType_COMMIT_END:
			GrpcCommitCount++
			if GrpcCommitCount == 1 {
				go func() {
					fmt.Println("Going to sync interface status")
					time.Sleep(time.Second)
					InterfaceSyncWithConfig()
				}()
			}
		case rpc.ConfigType_SET, rpc.ConfigType_DELETE:
			if !validating {
				RpcConfig(conf)
			}
		default:
		}
	}
	return err
}

const (
	GrpcConnRetryInterval = 5
)

var (
	GrpcConn        *grpc.ClientConn
	GrpcCommitCount int
)

func RpcRegister() {
	InitAPI()

	for {
		GrpcCommitCount = 0
		conn, err := grpc.Dial(":2650",
			grpc.WithInsecure(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithBlock(),
			grpc.WithTimeout(time.Second*GrpcConnRetryInterval),
		)
		if err == nil {
			GrpcConn = conn
			rpcRegisterModule(conn)
			rpcRegisterCli(conn)
			rpcSubscribe(conn)
			GrpcConn = nil
		} else {
			interval := (rand.Intn(GrpcConnRetryInterval) + 1)
			select {
			case <-time.After(time.Second * time.Duration(interval)):
				// Wait timeout.
			}
		}
	}
}

type GrpcComponent struct {
}

func NewGrpcComponent() *GrpcComponent {
	return &GrpcComponent{}
}

func (g *GrpcComponent) Start() component.Component {
	go RpcRegister()
	go RpcServer()
	return g
}

func (g *GrpcComponent) Stop() component.Component {
	return g
}
