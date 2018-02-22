// Copyright 2016 zebra project
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

package main

import (
	"fmt"

	"github.com/coreswitch/component"
	"github.com/hash-set/zebra/bgp"
	"github.com/hash-set/zebra/module"
)

func main() {
	fmt.Println("Starting bgpd")
	module.Init()

	server := bgp.NewServer(65100)

	serverComponent := &bgp.ServerComponent{
		Server: server,
	}
	grpcComponent := &bgp.GrpcComponent{
		Server: server,
	}

	systemMap := component.SystemMap{
		"server": serverComponent,
		"grpc":   component.ComponentWith(grpcComponent, "server"),
	}
	systemMap.Start()

	err := server.RouterIdSet("192.168.55.1")
	if err != nil {
		fmt.Println(err)
	}

	err = server.NeighborAdd("192.168.55.2")
	if err != nil {
		fmt.Println(err)
	}
	err = server.NeighborRemoteAsSet("192.168.55.2", 65520)
	if err != nil {
		fmt.Println(err)
	}
	err = server.NeighborAfiSafiSet("192.168.55.2", bgp.AFI_IP, bgp.SAFI_UNICAST)
	if err != nil {
		fmt.Println(err)
	}
	err = server.NeighborAfiSafiSet("192.168.55.2", bgp.AFI_IP6, bgp.SAFI_UNICAST)
	if err != nil {
		fmt.Println(err)
	}

	ch := make(chan struct{})
	<-ch
}
