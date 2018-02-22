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
	"fmt"
	"io"
	"net"
	"os"

	"github.com/coreswitch/cmd"
)

type Vlan struct {
	VlanId   uint16
	IfName   string
	IfConfig *InterfaceConfig
	File     *os.File
	Ports    []*InterfaceConfig
}

type vlanDB map[uint16]*Vlan

var VlanDB = &vlanDB{}

var (
	VlanAddHook    func(uint16, *InterfaceConfig) int
	VlanDeleteHook func(uint16, *InterfaceConfig) int
	VlanTapHook    func(uint16, []byte, int)
)

func (vlanDB *vlanDB) Lookup(vlanId uint16) *Vlan {
	return (*vlanDB)[vlanId]
}

func (vlanDB *vlanDB) Add(vlanId uint16) {
	vlan := VlanDB.Lookup(vlanId)
	if vlan != nil {
		return
	}

	ifname := fmt.Sprintf("vlan%d", vlanId)

	vlan = &Vlan{VlanId: vlanId, IfName: ifname}
	(*vlanDB)[vlanId] = vlan

	// Bind interface config.
	ifc := InterfaceConfigGet(ifname)
	ifc.VlanId = int(vlanId)
	vlan.File, _ = TapAdd(ifname)
	vlan.IfConfig = ifc
	go vlan.Reader()

	/* We have special hook.  */
	hwaddr, err := net.ParseMAC(fmt.Sprintf("a0:aa:aa:aa:aa:%02x", vlanId))
	if err != nil {
		fmt.Println(err)
	}
	ifc.HwAddr = hwaddr

	if VlanAddHook != nil {
		VlanAddHook(vlanId, ifc)
	}
}

func (vlan *Vlan) Reader() {
	buf := make([]byte, 2048)
	for {
		n, err := vlan.File.Read(buf)
		if err == io.EOF {
			return
		}
		// fmt.Println("Read from vlan", vlan.VlanId, "bytes", n)
		if VlanTapHook != nil {
			VlanTapHook(vlan.VlanId, buf, n)
		}
	}
}

func (vlanDB *vlanDB) Dump() {
	for key, val := range *vlanDB {
		fmt.Println(key, val)
	}
}

func (vlanDB *vlanDB) Delete(vlanId uint16) {
	//vlanDB.Dump()
	vlan := vlanDB.Lookup(vlanId)
	if vlan == nil {
		return
	}

	ifname := fmt.Sprintf("vlan%d", vlanId)

	if VlanDeleteHook != nil {
		VlanDeleteHook(vlanId, nil)
	}

	ifp := vlan.IfConfig.Interface
	LinkSetDown(ifp)
	TapDelete(ifname)
}

func VlanApi(Cmd int, Args cmd.Args) int {
	vlanId := uint16(Args[0].(uint64))
	if Cmd == cmd.Set {
		VlanDB.Add(vlanId)
	} else {
		VlanDB.Delete(vlanId)
	}
	return cmd.Success
}
