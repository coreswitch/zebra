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
	"fmt"

	"github.com/coreswitch/netutil"
	"github.com/coreswitch/zebra/fea"
)

func (s *Server) ifLookupByName(ifName string) *fea.Interface {
	return s.IfMap[ifName]
}

func (s *Server) enableIfLookup(ifName string) bool {
	return s.EnableIfMap[ifName]
}

func (s *Server) enableIfAdd(ifName string) {
	s.EnableIfMap[ifName] = true
}

func (s *Server) enableIfDelete(ifName string) {
	delete(s.EnableIfMap, ifName)
}

func (s *Server) interfaceWakeUp(ifp *fea.Interface) {
	maddr := netutil.ParseIPv4(INADDR_RIP_GROUP)
	for _, ifAddr := range ifp.AddrIpv4 {
		multicastJoin(s.Sock, maddr, ifAddr.Address.IP, ifp.Index)
	}
}

func (s *Server) prefixUpdate(ifi *InterfaceInfo) {
}

func (s *Server) triggeredUpdateAll() {
}

func (s *Server) enableApply(ifp *fea.Interface) {
	ifi := s.InterfaceInfoGet(ifp)

	s.passiveIfApply(ifi)
	s.prefixUpdate(ifi)

	s.up()

	if !ifi.up {
		ifi.up = true
		s.interfaceWakeUp(ifp)
	}

	s.triggeredUpdateAll()
}

func (s *Server) enableApplyAll() {
}

func (s *Server) passiveIfLookup(ifName string) bool {
	return s.PassiveIfMap[ifName]
}

func (s *Server) passiveIfAdd(ifName string) {
	s.PassiveIfMap[ifName] = true
}

func (s *Server) passiveIfDelete(ifName string) {
	delete(s.PassiveIfMap, ifName)
}

func (s *Server) passiveIfApply(ifi *InterfaceInfo) {
	if s.passiveIfLookup(ifi.ifp.Name) {
		ifi.passive = true
	} else {
		ifi.passive = false
	}
}

func (s *Server) multicastJoin(ifp *fea.Interface) bool {
	// s.sock,
	multicast := netutil.ParseIPv4(INADDR_RIP_GROUP)
	fmt.Println(multicast)
	return true
}
