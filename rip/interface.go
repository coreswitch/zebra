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

package rip

import (
	"github.com/coreswitch/log"
	"github.com/coreswitch/netutil"
	"github.com/coreswitch/zebra/fea"
)

type Interface struct {
	dev     *fea.Interface
	Enabled bool
}

type Interfaces struct {
	IfMap   map[string]*Interface
	IfTable map[uint32]*Interface
}

func NewInterfaces() *Interfaces {
	ifdb := &Interfaces{
		IfTable: map[uint32]*Interface{},
		IfMap:   map[string]*Interface{},
	}
	return ifdb
}

func (ifdb *Interfaces) GetByName(ifName string) *Interface {
	ifp, ok := ifdb.IfMap[ifName]
	if !ok {
		ifp = &Interface{}
		ifdb.IfMap[ifName] = ifp
	}
	return ifp
}

func (ifdb *Interfaces) LookupByName(ifName string) *Interface {
	return ifdb.IfMap[ifName]
}

func (ifdb *Interfaces) Register(dev *fea.Interface) {
	ifp := ifdb.GetByName(dev.Name)
	ifp.dev = dev
	ifdb.IfTable[dev.Index] = ifp
}

func (ifdb *Interfaces) Unregister(dev *fea.Interface) {
	ifp := ifdb.GetByName(dev.Name)
	ifp.dev = nil
	delete(ifdb.IfTable, dev.Index)
}

func InterfaceMulticastJoin(sock int, dev *fea.Interface) {
	maddr := netutil.ParseIPv4(INADDR_RIP_GROUP)
	for _, ifAddr := range dev.AddrIpv4 {
		multicastJoin(sock, maddr, ifAddr.Address.IP, dev.Index)
	}
}

func (s *Server) EnableInterface(ifp *Interface) {
	if ifp.dev == nil {
		log.Info("Do not enable interface since ifp.dev is nil")
		return
	}

	s.Run()

	// s.passiveIfApply(ifp)
	// s.prefixUpdate(ifp)

	InterfaceMulticastJoin(s.Sock, ifp.dev)

	//s.triggeredUpdateAll()

	ifp.Enabled = true
}
