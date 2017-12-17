// Copyright 2017 zebra Project
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
	"net"
	"runtime"
	"sync"
	"time"

	pb "github.com/coreswitch/zebra/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type grpcServer struct {
	Mutex sync.RWMutex
	Peers map[peer.Peer]*grpcPeer
}

type grpcPeer struct {
	interfaceStream  pb.Zebra_InterfaceServiceServer
	routerIdStream   pb.Zebra_RouterIdServiceServer
	redistIPv4Stream pb.Zebra_RedistributeIPv4ServiceServer
	redistIPv6Stream pb.Zebra_RedistributeIPv6ServiceServer
	routeIPv4Stream  pb.Zebra_RouteIPv4ServiceServer
	routeIPv6Stream  pb.Zebra_RouteIPv6ServiceServer
	dispatCh         chan interface{}
	done             chan interface{}
}

func NewGrpcPeer() *grpcPeer {
	p := &grpcPeer{
		dispatCh: make(chan interface{}),
		done:     make(chan interface{}),
	}
	go p.Dispatch()
	return p
}

func (p *grpcPeer) Dispatch() {
	for {
		select {
		case mes := <-p.dispatCh:
			switch mes.(type) {
			case *pb.InterfaceRequest:
				fmt.Println("InterfaceRequest", mes)
			case *pb.RouterIdRequest:
				fmt.Println("RouterIdRequest:", mes)
			case *pb.RedistributeIPv4Request:
				fmt.Println("RedistributeIPv4Request", mes)
			case *pb.RedistributeIPv6Request:
				fmt.Println("RedistributeIPv6Request", mes)
			case *pb.RouteIPv4:
				fmt.Println("RouteIPv4", mes)
			case *pb.RouteIPv6:
				fmt.Println("RouteIPv6", mes)
			}
		case <-p.done:
			fmt.Println("peer Dispatch: done")
			return
		}
	}
}

func (s *grpcServer) PeerGet(p *peer.Peer) *grpcPeer {
	s.Mutex.Lock()
	peer, ok := s.Peers[*p]
	if !ok {
		fmt.Println("New Peer", p)
		peer = NewGrpcPeer()
		s.Peers[*p] = peer
	} else {
		fmt.Println("Existing Peer", p)
	}
	s.Mutex.Unlock()
	return peer
}

func (s *grpcServer) PeerDelete(p *peer.Peer) {
	s.Mutex.Lock()
	peer, ok := s.Peers[*p]
	if ok {
		close(peer.done)
		delete(s.Peers, *p)
	}
	s.Mutex.Unlock()
}

func NewGrpcServer() *grpcServer {
	s := &grpcServer{
		Peers: map[peer.Peer]*grpcPeer{},
	}
	return s
}

func (s *grpcServer) InterfaceService(stream pb.Zebra_InterfaceServiceServer) error {
	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := s.PeerGet(p)
	peer.interfaceStream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			s.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (s *grpcServer) RouterIdService(stream pb.Zebra_RouterIdServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := s.PeerGet(p)
	peer.routerIdStream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			s.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (s *grpcServer) RedistributeIPv4Service(stream pb.Zebra_RedistributeIPv4ServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := s.PeerGet(p)
	peer.redistIPv4Stream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			s.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (s *grpcServer) RedistributeIPv6Service(stream pb.Zebra_RedistributeIPv6ServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := s.PeerGet(p)
	peer.redistIPv6Stream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			s.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (s *grpcServer) RouteIPv4Service(stream pb.Zebra_RouteIPv4ServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := s.PeerGet(p)
	peer.routeIPv4Stream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			s.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (s *grpcServer) RouteIPv6Service(stream pb.Zebra_RouteIPv6ServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := s.PeerGet(p)
	peer.routeIPv6Stream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			s.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func Start() {
	fmt.Println("Server start")
	fmt.Println("goroutine", runtime.NumGoroutine())

	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	s := grpc.NewServer()

	fmt.Println("goroutine", runtime.NumGoroutine())

	gserver := NewGrpcServer()

	pb.RegisterZebraServer(s, gserver)

	fmt.Println("goroutine", runtime.NumGoroutine())

	go s.Serve(lis)

	for {
		fmt.Println("goroutine", runtime.NumGoroutine())
		time.Sleep(time.Second)
	}
}

func Stop() {
	fmt.Println("Server stop")
}
