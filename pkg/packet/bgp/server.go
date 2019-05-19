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
	"sync"

	"github.com/coreswitch/component"
)

const (
	BGPServerReset = iota
)

type ServerConfig struct {
	As              uint32
	RouterId        string
	GracefulRestart bool
	GrRestartTime   *int
	GrStalePathTime *int
}

type Server struct {
	Config     ServerConfig
	Neighbors  map[string]*Neighbor
	as         uint32
	port       *uint16
	routerId   net.IP
	event      chan int
	fn         chan *Fn
	done       chan interface{}
	wg         sync.WaitGroup
	restarting bool
	preserved  bool
}

var (
	DefaultServer *Server
)

func NewServer(as uint32) *Server {
	server := &Server{
		as:        as,
		Neighbors: make(map[string]*Neighbor),
		event:     make(chan int, 1024),
		fn:        make(chan *Fn, 1024),
		done:      make(chan interface{}),
	}
	return server
}

func ServerStart(As uint32) *Server {
	return nil
}

func ServerStop() {
}

func (s *Server) PortSet(port uint16) {
	if port == BGP_PORT {
		s.PortUnset()
		return
	}
	if s.port != nil && *s.port == port {
		return
	}
	s.port = &port
	s.event <- BGPServerReset
}

func (s *Server) PortUnset() {
	if s.port == nil {
		return
	}
	s.port = nil
	s.event <- BGPServerReset
}

func (s *Server) Port() uint16 {
	if s.port == nil {
		return BGP_PORT
	}
	return *s.port
}

func (s *Server) RouterId() net.IP {
	return s.routerId
}

type Fn struct {
	fn  func() error
	err chan error
}

func (s *Server) api(fn func() error) error {
	err := make(chan error)
	s.fn <- &Fn{fn: fn, err: err}
	return <-err
}

func ParseIP(str string) net.IP {
	ip := net.ParseIP(str)
	if ip == nil {
		return nil
	}
	ip4 := ip.To4()
	if ip4 != nil {
		return ip4
	}
	return ip
}

func (s *Server) Runnable() bool {
	if s.routerId != nil {
		return true
	} else {
		return false
	}
}

func (s *Server) routerIdSet(routerIdStr string) error {
	routerId := ParseIP(routerIdStr)
	if routerId == nil {
		return fmt.Errorf("router ID format error: %s", routerIdStr)
	}
	if len(routerId) != net.IPv4len {
		return fmt.Errorf("router ID is not IPv4 format: %s", routerIdStr)
	}
	if s.routerId != nil && s.routerId.Equal(routerId) {
		return nil
	}
	s.routerId = routerId
	s.event <- BGPServerReset
	return nil
}

func (s *Server) routerIdUnset() error {
	if s.routerId == nil {
		return nil
	}
	s.routerId = nil
	s.event <- BGPServerReset
	return nil
}

func (s *Server) RouterIdSet(routerIdStr string) error {
	return s.api(func() error {
		return s.routerIdSet(routerIdStr)
	})
}

func (s *Server) RouterIdUnset() error {
	return s.api(func() error {
		return s.routerIdUnset()
	})
}

func (s *Server) gracefulRestartEnable() error {
	s.Config.GracefulRestart = true
	return nil
}

func (s *Server) GracefulRestartEnable() error {
	return s.api(func() error {
		return s.gracefulRestartEnable()
	})
}

func (s *Server) gracefulRestartDisable() error {
	s.Config.GracefulRestart = false
	return nil
}

func (s *Server) GracefulRestartDisable() error {
	return s.api(func() error {
		return s.gracefulRestartDisable()
	})
}

func (s *Server) GracefulRestartTime() int {
	if s.Config.GrRestartTime != nil {
		return *s.Config.GrRestartTime
	} else {
		return GR_RESTART_TIME
	}
}

type ResetType int

const (
	HARD_RESET = iota
	SOFT_RESET_BOTH
	SOFT_RESET_IN
	SOFT_RESET_OUT
)

func (n *Neighbor) Reset(resetType ResetType) error {
	switch resetType {
	case HARD_RESET:
	case SOFT_RESET_BOTH:
	case SOFT_RESET_IN:
	case SOFT_RESET_OUT:
	}
	return nil
}

func (s *Server) Reset(addrStr string, resetType ResetType) error {
	// Special neighbor "*" means all neighbor.
	if addrStr == "*" {
		for _, n := range s.Neighbors {
			n.Reset(resetType)
		}
	} else {
		n, err := s.neighborLookup(addrStr)
		if err != nil {
			return err
		}
		n.Reset(resetType)
	}
	return nil
}

func (s *Server) ResetAll() error {
	fmt.Println("BGP server reset")

	// Stop all of neighbors.
	close(s.done)
	s.wg.Wait()

	// Check if this server is runnable, start neighbor FSM.
	s.done = make(chan interface{})
	if s.Runnable() {
		for _, n := range s.Neighbors {
			n.fsm.EventLoop(s.done)
		}
	}
	return nil
}

func (s *Server) Serv() {
	for {
		select {
		case e := <-s.event:
			switch e {
			case BGPServerReset:
				s.ResetAll()
			}
		case f := <-s.fn:
			f.err <- f.fn()
		}
	}
}

type ServerComponent struct {
	Server *Server
}

func (s *ServerComponent) Start() component.Component {
	go s.Server.Serv()
	return s
}

func (s *ServerComponent) Stop() component.Component {
	return s
}
