// Copyright 2018 zebra project.
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
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreswitch/component"
	"github.com/coreswitch/log"
	"github.com/coreswitch/zebra/rip"
)

func main() {
	log.Info("RIPd Starting")
	server := rip.NewServer()
	serverComponent := &rip.ServerComponent{
		Server: server,
	}
	rpcComponent := &rip.RpcComponent{
		Server: server,
	}
	systemMap := component.SystemMap{
		"server": serverComponent,
		"rpc":    component.ComponentWith(rpcComponent, "server"),
	}
	systemMap.Start()
	log.Info("RIPd Started")

	time.Sleep(time.Second * 3)
	log.Info("Adding enp0s6")
	server.EnableInterfaceSet("enp0s6")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Ignore(syscall.SIGWINCH)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		log.Info("RIPd Stopping")
		systemMap.Stop()
		log.Info("RIPd Stopped")
		done <- true
	}()

	<-done
}
