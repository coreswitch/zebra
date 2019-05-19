package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/coreswitch/component"
	"github.com/coreswitch/log"
	"github.com/coreswitch/zebra/fea"
	"github.com/coreswitch/zebra/fea/linux"
	"github.com/coreswitch/zebra/pkg/server/module"
	"github.com/coreswitch/zebra/rib"
)

func EsiHook() {
	if _, err := os.Stat("/etc/esi/env"); err == nil {
		rib.IfAddHook = rib.EsiIfAddHook
		rib.IfAddrAddHook = rib.EsiIfAddrAddHook
		rib.IfAddrDeleteHook = rib.EsiIfAddrDeleteHook
		rib.IfStatusChangeHook = rib.IfStatusNotifyEtcd
		rib.NetlinkDoneHook = rib.EsiNetlinkDoneHook
		rib.DefaultVrfProtect = true
		rib.NexthopLookupHook = rib.EsiNexthopLookup
		rib.ShutdownSkipHook = rib.EsiShutdownSkipHook
		rib.AddPathDefault = true
		rib.IfForceUpFlag = true
		rib.LanInterfaceMonitorStart()
		// rib.NoVRFDelete = true
		rib.NetlinkIPv4AddressEnsure = true
		rib.LocalPolicy = true
		rib.NewServerHook = rib.EsiNewServerHook
	}
	rib.NewServerHook = rib.EsiNewServerHook
}

func main() {
	fmt.Println("Starting RIB daemon")

	log.FuncField = false
	log.SourceField = false
	log.SetTextFormatter()

	module.Init()
	linux.Init()

	EsiHook()

	// rib.PlistTest()

	feaComponent := &fea.FeaComponent{}
	ribComponent := rib.NewServer()
	grpcComponent := rib.NewGrpcComponent()
	rpcComponent := rib.NewRpcComponent(ribComponent, 2699)

	systemMap := component.SystemMap{
		"fea":  feaComponent,
		"rib":  component.ComponentWith(ribComponent, "fea"),
		"grpc": component.ComponentWith(grpcComponent, "rib"),
		"rpc":  component.ComponentWith(rpcComponent, "rib"),
	}
	systemMap.Start()

	// Register clearn up function.
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Ignore(syscall.SIGWINCH)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigs
		systemMap.Stop()
		done <- true
	}()

	<-done
}
