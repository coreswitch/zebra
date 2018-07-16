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
	"github.com/coreswitch/cfg"
	"github.com/coreswitch/log"
	"github.com/coreswitch/zebra/fea"
	"golang.org/x/sys/unix"
)

type Interface struct {
	dev         *fea.Interface
	Enable      *bool
	SendVersion *byte
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

func (ifdb *Interfaces) Register(dev *fea.Interface) *Interface {
	ifp := ifdb.GetByName(dev.Name)
	ifp.dev = dev
	ifdb.IfTable[dev.Index] = ifp
	return ifp
}

func (ifdb *Interfaces) Unregister(dev *fea.Interface) {
	ifp := ifdb.GetByName(dev.Name)
	ifp.dev = nil
	delete(ifdb.IfTable, dev.Index)
}

func InterfaceMulticastJoin(sock int, dev *fea.Interface) {
	for _, ifAddr := range dev.AddrIpv4 {
		multicastJoin(sock, RIP_GROUP_ADDR, ifAddr.Address.IP, dev.Index)
	}
}

func InterfaceMulticastIf(sock int, dev *fea.Interface) error {
	for _, ifAddr := range dev.AddrIpv4 {
		err := multicastIf(sock, ifAddr.Address.IP, dev.Index)
		if err != nil {
			return err
		}
		addr := &unix.SockaddrInet4{}
		addr.Port = RIP_PORT_DEFAULT
		copy(addr.Addr[:], ifAddr.Address.IP)
		err = unix.Bind(sock, addr)
		if err != nil {
			log.Warn(err)
		}
	}
	return nil
}

func RequestSendPacket(ifp *Interface, version byte) {
	log.Info("RequestSendPacket")
	// RFC2453 3.9.1 Request Messages: If there is exactly one entry in the
	// request, and it has an address family identifier of zero and a metric of
	// infinity (i.e., 16), then this is a request to send the entire routing
	// table.
	p := &Packet{}
	p.Command = RIP_REQUEST
	p.Version = version
	rte := &RTE{}
	rte.Metric = RIP_METRIC_INFINITY
	p.RTEs = append(p.RTEs, rte)
	SendMulticastPacket(ifp, p)
}

func RequestSendInterface(ifp *Interface, version byte) {
	if version == RIPv2 {
		RequestSendPacket(ifp, version)
	} else {
	}
}

func RequestSend(ifp *Interface) {
	// Figure out version.
	version := RIPv2
	if ifp.SendVersion != nil {
		version = cfg.ByteVal(ifp.SendVersion)
	}
	RequestSendInterface(ifp, version)
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

	RequestSend(ifp)
}
