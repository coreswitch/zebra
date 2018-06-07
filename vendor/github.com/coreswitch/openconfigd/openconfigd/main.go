// Copyright 2016 OpenConfigd Project.
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
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/coreswitch/component"
	"github.com/coreswitch/openconfigd/config"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
)

func EsiHook() {
	if _, err := os.Stat("/etc/esi/env"); err == nil {
		config.RibdAsyncUpdate = true
	}
}

func main() {
	fmt.Println("Starting openconfigd.")

	EsiHook()

	var opts struct {
		ConfigActiveFile string `short:"c" long:"config-file" description:"active config file name" default:"openconfigd.conf"`
		ConfigFileDir    string `short:"p" long:"config-dir" description:"config file directory" default:"/usr/local/etc"`
		RpcEndPoint      string `short:"g" long:"grpc-endpoint" description:"Grpc End Point" default:"127.0.0.1:2650"`
		YangPaths        string `short:"y" long:"yang-paths" description:"comma separated YANG load path directories"`
		TwoPhaseCommit   bool   `short:"2" long:"two-phase" description:"two phase commit"`
		ZeroConfig       bool   `short:"z" long:"zero-config" description:"Do not save or load config other than openconfigd.conf"`
	}

	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	// Arguments are YANG module names. When no modules is specified "coreswitch"
	// is set as default YANG module name.
	if len(args) == 0 {
		args = []string{"coreswitch"}
	}

	// Set log output.
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})

	// Set maxprocs
	runtime.GOMAXPROCS(runtime.NumCPU())

	// YANG Component
	yangComponent := &config.YangComponent{
		YangPaths:   opts.YangPaths,
		YangModules: args,
	}

	// Config component
	configComponent := &config.ConfigComponent{
		ConfigFileDir:    opts.ConfigFileDir,
		ConfigActiveFile: opts.ConfigActiveFile,
		TwoPhaseCommit:   opts.TwoPhaseCommit,
		ZeroConfig:       opts.ZeroConfig,
	}

	// CLI component
	cliComponent := &config.CliComponent{}

	// Rpc Component
	rpcComponent := &config.RpcComponent{
		GrpcEndpoint: opts.RpcEndPoint,
	}

	systemMap := component.SystemMap{
		"yang":   yangComponent,
		"cli":    component.ComponentWith(cliComponent, "yang"),
		"config": component.ComponentWith(configComponent, "yang", "cli"),
		"rpc":    component.ComponentWith(rpcComponent, "config"),
	}

	// Register clearn up function.
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		systemMap.Stop()
		os.Exit(1)
	}()

	// Start components.
	systemMap.Start()
}
