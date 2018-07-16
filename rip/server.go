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

	"github.com/coreswitch/cfg"
	"github.com/coreswitch/component"
	"github.com/coreswitch/log"
)

type Server struct {
	Version    uint8 // RIPv1 and RIPv2
	Sock       int
	SyncCh     chan *Fn
	DispatCh   chan interface{}
	Done       chan interface{}
	Client     *Client
	Running    bool
	Interfaces *Interfaces
}

type Fn struct {
	fn  func() error
	err chan error
}

func NewServer() *Server {
	return &Server{
		Version:    RIPv2,
		Sock:       -1,
		SyncCh:     make(chan *Fn, 1024),
		DispatCh:   make(chan interface{}),
		Done:       make(chan interface{}),
		Interfaces: NewInterfaces(),
	}
}

func (s *Server) Start() {
	s.Client = NewClient(s.DispatCh)
	go s.Client.Start()

	for {
		select {
		case sync := <-s.SyncCh:
			sync.err <- sync.fn()
		case res := <-s.DispatCh:
			s.Dispatch(res)
		case <-s.Done:
			log.Info("Server Done")
			return
		}
	}
}

func (s *Server) Stop() {
	close(s.Done)
}

func (s *Server) api(fn func() error) error {
	err := make(chan error)
	s.SyncCh <- &Fn{fn: fn, err: err}
	return <-err
}

func (s *Server) RouterSet() error {
	return s.api(func() error {
		fmt.Println("RouterSet")
		s.Run()
		return nil
	})
}

func (s *Server) RouterUnset() error {
	return nil
}

func (s *Server) VersionSet(version int) error {
	return nil
}

func (s *Server) VersionUnset(version int) error {
	return nil
}

func (s *Server) EnableInterfaceSet(ifName string) error {
	return s.api(func() error {
		ifp := s.Interfaces.GetByName(ifName)
		if cfg.BoolVal(ifp.Enable) {
			return nil
		}
		s.EnableInterface(ifp)
		ifp.Enable = cfg.Bool(true)
		return nil
	})
}

func (s *Server) enableInterfaceUnset(ifName string) error {
	return nil
}

func (s *Server) EnableInterfaceUnset(ifName string) error {
	return s.api(func() error {
		return s.enableInterfaceUnset(ifName)
	})
}

func (s *Server) EnableNetworkSet() error {
	return nil
}

func (s *Server) EnableNetworkUnset() error {
	return nil
}

func (s *Server) NeighborSet() error {
	return nil
}

func (s *Server) NeighborUnset() error {
	return nil
}

func (s *Server) Run() {
	if s.Running {
		return
	}
	s.Sock = MakeSocket()
	if s.Sock < 0 {
		return
	}

	go s.Read()

	s.Running = true
}

func (s *Server) down() {
	s.Running = false
}

type ServerComponent struct {
	Server *Server
}

func (c *ServerComponent) Start() component.Component {
	go c.Server.Start()
	return c
}

func (c *ServerComponent) Stop() component.Component {
	c.Server.Stop()
	return c
}
