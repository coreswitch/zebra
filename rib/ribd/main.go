package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/coreswitch/component"
	"github.com/coreswitch/zebra/fea"
	"github.com/coreswitch/zebra/fea/linux"
	"github.com/coreswitch/zebra/module"
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
		// rib.NoVRFDelete = true
	}
}

func main() {
	fmt.Println("Starting RIB daemon")

	module.Init()
	linux.Init()

	EsiHook()

	feaComponent := &fea.FeaComponent{}
	ribComponent := rib.NewServer()
	restComponent := rib.NewRestComponent()
	grpcComponent := rib.NewGrpcComponent()

	systemMap := component.SystemMap{
		"fea":  feaComponent,
		"rib":  component.ComponentWith(ribComponent, "fea"),
		"rest": component.ComponentWith(restComponent, "rib"),
		"grpc": component.ComponentWith(grpcComponent, "rib"),
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
