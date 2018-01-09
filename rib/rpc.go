// Copyright 2017, 2018 zebra project.
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
	"sync"

	"github.com/coreswitch/component"
	"github.com/coreswitch/log"
	"github.com/coreswitch/netutil"
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
	server          *Server
	interfaceStream pb.Zebra_InterfaceServiceServer
	routerIdStream  pb.Zebra_RouterIdServiceServer
	redistStream    pb.Zebra_RedistServiceServer
	routeStream     pb.Zebra_RouteServiceServer
	dispatCh        chan interface{}
	done            chan interface{}
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

func NewPrefix(p *pb.Prefix) *netutil.Prefix {
	return &netutil.Prefix{
		IP:     p.Addr,
		Length: int(p.Length),
	}
}

func NewNexthop(nhop *pb.Nexthop) *Nexthop {
	return &Nexthop{
		IP:    nhop.Addr,
		Index: IfIndex(nhop.Ifindex),
	}
}

func NewRoute(r *pb.Route, p *netutil.Prefix, src *rpcPeer) *Rib {
	rib := &Rib{
		Prefix:  p,
		Type:    uint8(r.Type),
		SubType: uint8(r.SubType),
		Metric:  r.Metric,
		Src:     src,
	}
	if r.Distance != 0 {
		rib.Distance = uint8(r.Distance)
	} else {
		rib.Distance = DistanceCalc(rib.Type, rib.SubType)
	}
	if len(r.Nexthops) == 1 {
		rib.Nexthop = NewNexthop(r.Nexthops[0])
	}
	return rib
}

func (p *rpcPeer) Dispatch() {
	for {
		select {
		case mes := <-p.dispatCh:
			switch mes.(type) {
			case *pb.RedistRequest:
				log.Info("RedistRequest:", mes)
			case *pb.Route:
				req := mes.(*pb.Route)
				vrf := VrfLookupByIndex(int(req.VrfId))
				if vrf == nil {
					log.Errorf("VRF can't find with VRF ID %d", req.VrfId)
					continue
				}
				if len(req.Nexthops) == 0 {
					log.Errorf("No nexthop in request")
					continue
				}
				prefix := NewPrefix(req.Prefix)
				rib := NewRoute(req, prefix, p)
				log.Info(req.Op, p, rib)
				switch req.Op {
				case pb.Op_RouteAdd:
					vrf.RibAdd(prefix, rib)
				case pb.Op_RouteDelete:
					vrf.RibDelete(prefix, rib)
				}
			}
		case <-p.done:
			log.Info("Peer dispatch is done.  Exiting goroutine")
			return
		}
	}
}

func (p *rpcPeer) Notify(mes interface{}) {
	switch mes.(type) {
	case *pb.InterfaceUpdate:
		p.interfaceStream.Send(mes.(*pb.InterfaceUpdate))
	case *pb.RouterIdUpdate:
		p.routerIdStream.Send(mes.(*pb.RouterIdUpdate))
	}
}

func (r *rpcServer) PeerGet(p *peer.Peer) *rpcPeer {
	r.Mutex.Lock()
	peer, ok := r.peers[*p]
	if !ok {
		peer = NewRpcPeer(r.server)
		r.peers[*p] = peer
	}
	r.Mutex.Unlock()
	return peer
}

func (r *rpcServer) PeerDelete(p *peer.Peer) {
	r.Mutex.Lock()
	peer, ok := r.peers[*p]
	if ok {
		r.server.WatcherUnsubscribe(peer)
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
	logGoroutine()
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
			log.Info("InterfaceService Exit")
			return nil
		}

		log.With("VrfId", req.VrfId).Info(req.Op)

		switch req.Op {
		case pb.Op_InterfaceSubscribe:
			r.server.InterfaceSubscribe(peer, req.VrfId)
		case pb.Op_InterfaceUnsubscribe:
			r.server.InterfaceUnsubscribe(peer, req.VrfId)
		}
	}
}

func (r *rpcServer) RouterIdService(stream pb.Zebra_RouterIdServiceServer) error {
	logGoroutine()
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
			log.Info("RouterIdService Exit")
			return nil
		}
		log.With("VrfId", req.VrfId).Info(req.Op)
		r.server.RouterIdSubscribe(peer, req.VrfId)
	}
}

func (r *rpcServer) RedistService(stream pb.Zebra_RedistServiceServer) error {
	logGoroutine()
	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := r.PeerGet(p)
	peer.redistStream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			r.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func (r *rpcServer) RouteService(stream pb.Zebra_RouteServiceServer) error {
	logGoroutine()
	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("Can't get peer from context")
	}
	peer := r.PeerGet(p)
	peer.routeStream = stream

	for {
		req, err := stream.Recv()
		if err != nil {
			r.PeerDelete(p)
			return nil
		}
		peer.dispatCh <- req
	}
}

func logGoroutine() {
	// log.Infof("goroutine %d", runtime.NumGoroutine())
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
