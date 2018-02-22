// Copyright 2017 zebra project
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

package bgp

import (
	"fmt"
	"net"
	"time"
)

type MessageCount struct {
	open         uint64
	update       uint64
	keepalive    uint64
	notification uint64
	refresh      uint64
	capability   uint64
	prefix       uint64
}

func (m MessageCount) Sum() uint64 {
	return m.open + m.update + m.keepalive + m.notification + m.refresh + m.capability
}

type NeighborConfig struct {
	holdTime *uint16
}

type Neighbor struct {
	Config           NeighborConfig
	server           *Server
	addr             net.IP
	as               *uint32
	afiSafi          map[AfiSafi]bool
	passive          bool
	fsm              *Fsm
	running          bool
	connRetryTime    int
	connRetryCounter uint64
	dontCapAll       bool
	dontCap4As       bool
	dontCapRefresh   bool
	adminShutdown    bool
	localCaps        []CapabilityInterface
	remoteCaps       []CapabilityInterface
	in               MessageCount
	out              MessageCount
	uptime           time.Time
	reflectorClient  bool
}

func NewNeighbor(server *Server, addr net.IP) *Neighbor {
	n := &Neighbor{
		server:        server,
		addr:          addr,
		afiSafi:       map[AfiSafi]bool{},
		connRetryTime: 3,
	}
	n.fsm = NewFsm(n)
	n.uptime = time.Now()
	return n
}

func (n *Neighbor) LocalAs() uint32 {
	return n.server.as
}

func (n *Neighbor) RemoteAs() uint32 {
	if n.as != nil {
		return *n.as
	}
	return AS_UNSPEC
}

func (n *Neighbor) HoldTime() uint16 {
	if n.Config.holdTime != nil {
		return *n.Config.holdTime
	}
	return DEFAULT_HOLDTIME
}

func (n *Neighbor) Address() net.IP {
	return n.addr
}

func (n *Neighbor) Start() {
	fmt.Println("Start Neighbor", n.addr)
	if n.passive {
		n.fsm.SendEvent(ManualStart_with_Passive)
	} else {
		n.fsm.SendEvent(ManualStart)
	}
}

func (n *Neighbor) Stop() {
	close(n.fsm.doneCh)
}

func (n *Neighbor) startOnCondition() {
	if !n.running {
		if n.as != nil && len(n.afiSafi) > 0 {
			n.Start()
			n.running = true
		}
	}
}

func (n *Neighbor) RemoteAsSet(as uint32) error {
	if n.as != nil {
		return fmt.Errorf("AS already set as %d", *n.as)
	}
	n.as = &as
	n.startOnCondition()
	return nil
}

func (n *Neighbor) AfiSafiSet(afi Afi, safi Safi) error {
	afiSafi := AfiSafiValue(afi, safi)
	if _, ok := n.afiSafi[afiSafi]; ok {
		return fmt.Errorf("afi safi is already configured")
	}
	n.afiSafi[afiSafi] = true
	n.startOnCondition()
	return nil
}

func (s *Server) neighborAdd(addrStr string) error {
	ip := ParseIP(addrStr)
	if ip == nil {
		return fmt.Errorf("address format error: %s", addrStr)
	}
	if _, ok := s.Neighbors[ip.String()]; ok {
		return fmt.Errorf("neighbor %s already exists", addrStr)
	}

	n := NewNeighbor(s, ip)
	s.Neighbors[addrStr] = n

	if s.Runnable() {
		n.fsm.EventLoop(s.done)
	}

	return nil
}

func (s *Server) neighborLookup(addrStr string) (*Neighbor, error) {
	ip := ParseIP(addrStr)
	if ip == nil {
		return nil, fmt.Errorf("address format error: %s", addrStr)
	}
	n := s.Neighbors[ip.String()]
	if n == nil {
		return nil, fmt.Errorf("can't find neighbor with address %s", addrStr)
	}
	return n, nil
}

func (s *Server) neighborRemoteAsSet(addrStr string, as uint32) error {
	n, err := s.neighborLookup(addrStr)
	if err != nil {
		return err
	}
	return n.RemoteAsSet(as)
}

func (s *Server) neighborAfiSafiSet(addrStr string, afi Afi, safi Safi) error {
	n, err := s.neighborLookup(addrStr)
	if err != nil {
		return err
	}
	return n.AfiSafiSet(afi, safi)
}

func (s *Server) neighborReflectorClientEnable(addrStr string) error {
	n, err := s.neighborLookup(addrStr)
	if err != nil {
		return err
	}
	n.reflectorClient = true
	return nil
}

func (s *Server) neighborReflectorClientDisable(addrStr string) error {
	n, err := s.neighborLookup(addrStr)
	if err != nil {
		return err
	}
	n.reflectorClient = false
	return nil
}

func (s *Server) NeighborAdd(addrStr string) error {
	return s.api(func() error {
		return s.neighborAdd(addrStr)
	})
}

func (s *Server) NeighborRemoteAsSet(addrStr string, as uint32) error {
	return s.api(func() error {
		return s.neighborRemoteAsSet(addrStr, as)
	})
}

func (s *Server) NeighborAfiSafiSet(addrStr string, afi Afi, safi Safi) error {
	return s.api(func() error {
		return s.neighborAfiSafiSet(addrStr, afi, safi)
	})
}

func (s *Server) NeighborReflectorClientEnable(addrStr string) error {
	return s.api(func() error {
		return s.neighborReflectorClientEnable(addrStr)
	})
}

func (s *Server) NeighborReflectorClientDisable(addrStr string) error {
	return s.api(func() error {
		return s.neighborReflectorClientDisable(addrStr)
	})
}
