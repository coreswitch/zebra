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

	"github.com/coreswitch/component"
	"github.com/coreswitch/log"
	pb "github.com/coreswitch/zebra/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type rpcServer struct {
	Mutex  sync.RWMutex
	server *Server
	peers  map[peer.Peer]*rpcPeer
}

type rpcPeer struct {
	server           *Server
	interfaceStream  pb.Zebra_InterfaceServiceServer
	routerIdStream   pb.Zebra_RouterIdServiceServer
	redistIPv4Stream pb.Zebra_RedistributeIPv4ServiceServer
	redistIPv6Stream pb.Zebra_RedistributeIPv6ServiceServer
	routeIPv4Stream  pb.Zebra_RouteIPv4ServiceServer
	routeIPv6Stream  pb.Zebra_RouteIPv6ServiceServer
	dispatCh         chan interface{}
	done             chan interface{}
}

func NewRpcPeer(s *Server) *rpcPeer {
	p := &rpcPeer{
		server:   s,
		dispatCh: make(chan interface{}),
		done:     make(chan interface{}),
	}
	go p.Dispatch()
	return p
}

func (p *rpcPeer) Dispatch() {
	for {
		select {
		case mes := <-p.dispatCh:
			switch mes.(type) {
			case *pb.InterfaceRequest:
				req := mes.(*pb.InterfaceRequest)
				log.With("VrfId", req.VrfId).Info("InterfaceRequest:", req.Op)
				p.server.InterfaceSubscribe(p, req.VrfId)
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

func (p *rpcPeer) Notify(mes interface{}) {
	switch mes.(type) {
	case *pb.InterfaceUpdate:
		p.interfaceStream.Send(mes.(*pb.InterfaceUpdate))
	}
}

func (r *rpcServer) PeerGet(p *peer.Peer) *rpcPeer {
	r.Mutex.Lock()
	peer, ok := r.peers[*p]
	if !ok {
		fmt.Println("New Peer", p)
		peer = NewRpcPeer(r.server)
		r.peers[*p] = peer
	} else {
		fmt.Println("Existing Peer", p)
	}
	r.Mutex.Unlock()
	return peer
}

func (r *rpcServer) PeerDelete(p *peer.Peer) {
	r.Mutex.Lock()
	peer, ok := r.peers[*p]
	if ok {
		close(peer.done)
		delete(r.peers, *p)
	}
	r.Mutex.Unlock()
}

func NewRpcServer(s *Server) *rpcServer {
	rs := &rpcServer{
		server: s,
		peers:  map[peer.Peer]*rpcPeer{},
	}
	return rs
}

func (r *rpcServer) InterfaceService(stream pb.Zebra_InterfaceServiceServer) error {
	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := r.PeerGet(p)
	peer.interfaceStream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			r.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (r *rpcServer) RouterIdService(stream pb.Zebra_RouterIdServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := r.PeerGet(p)
	peer.routerIdStream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			r.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (r *rpcServer) RedistributeIPv4Service(stream pb.Zebra_RedistributeIPv4ServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := r.PeerGet(p)
	peer.redistIPv4Stream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			r.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (r *rpcServer) RedistributeIPv6Service(stream pb.Zebra_RedistributeIPv6ServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := r.PeerGet(p)
	peer.redistIPv6Stream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			r.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (r *rpcServer) RouteIPv4Service(stream pb.Zebra_RouteIPv4ServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := r.PeerGet(p)
	peer.routeIPv4Stream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			r.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (r *rpcServer) RouteIPv6Service(stream pb.Zebra_RouteIPv6ServiceServer) error {
	fmt.Println("goroutine", runtime.NumGoroutine())

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := r.PeerGet(p)
	peer.routeIPv6Stream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			r.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

type RpcComponent struct {
	s    *Server
	port int
	gs   *grpc.Server
}

func NewRpcComponent(server *Server, port int) *RpcComponent {
	return &RpcComponent{
		s:    server,
		port: port,
	}
}

func (r *RpcComponent) Start() component.Component {
	log.Infof("Starting RPC component at port %d", r.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", r.port))
	if err != nil {
		log.Errorf("Starting RPC component failed: %s", err)
		return r
	}
	r.gs = grpc.NewServer()
	pb.RegisterZebraServer(r.gs, NewRpcServer(r.s))
	go r.gs.Serve(lis)
	return r
}

func (r *RpcComponent) Stop() component.Component {
	log.Info("Stopping RPC component")
	r.gs.Stop()
	return r
}
