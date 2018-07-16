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
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/coreswitch/log"
	"github.com/coreswitch/zebra/fea"
	pb "github.com/coreswitch/zebra/proto"
	"google.golang.org/grpc"
)

type Client struct {
	running         bool
	conn            *grpc.ClientConn
	serv            pb.ZebraClient
	dispatCh        chan interface{}
	interfaceStream pb.Zebra_InterfaceServiceClient
	routerIdStream  pb.Zebra_RouterIdServiceClient
	wg              sync.WaitGroup
	//done            chan interface{}
}

func NewClient(dispatCh chan interface{}) *Client {
	return &Client{
		dispatCh: dispatCh,
	}
}

func (c *Client) InterfaceSubscribe(vrfId uint32) error {
	stream, err := c.serv.InterfaceService(context.Background())
	if err != nil {
		return err
	}
	c.interfaceStream = stream

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				fmt.Println("XXX interface stream.Recv error", err)
				return
			}
			c.dispatCh <- res
		}
	}()

	req := &pb.InterfaceRequest{
		Op:    pb.Op_InterfaceSubscribe,
		VrfId: vrfId,
	}
	err = stream.Send(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) RouterIdSubscribe(vrfId uint32) error {
	stream, err := c.serv.RouterIdService(context.Background())
	if err != nil {
		return err
	}
	c.routerIdStream = stream

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			res, err := stream.Recv()
			if err != nil {
				fmt.Println("XXX router id stream.Recv error", err)
				return
			}
			c.dispatCh <- res
		}
	}()

	req := &pb.RouterIdRequest{
		Op:    pb.Op_RouterIdSubscribe,
		VrfId: vrfId,
	}
	err = stream.Send(req)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Start() {
	log.Info("Client start")
	c.running = true
	for {
		for {
			var err error
			c.conn, err = grpc.Dial(":2699", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second))
			if err == nil {
				log.Info("Client conn success")
				break
			}
			log.Info("Client start err", err)
			timer := time.NewTimer(time.Second * 3)
			select {
			// case <-c.done:
			// 	timer.Stop()
			// 	break
			case <-timer.C:
				log.Info("Client conn retry!")
			}
		}
		c.serv = pb.NewZebraClient(c.conn)

		err := c.InterfaceSubscribe(VRF_DEFAULT)
		if err != nil {
			c.conn.Close()
			goto Retry
		}
		err = c.RouterIdSubscribe(VRF_DEFAULT)
		if err != nil {
			c.conn.Close()
			goto Retry
		}

	Retry:
		c.wg.Wait()
		c.conn.Close()

		// Clean up interfaces.

		// If client is shutting down, return here.
		if !c.running {
			return
		}
	}
}

func (s *Server) Dispatch(res interface{}) {
	switch res.(type) {
	case *pb.InterfaceUpdate:
		mes := res.(*pb.InterfaceUpdate)
		dev := fea.InterfaceFromPb(mes)
		switch mes.Op {
		case pb.Op_InterfaceAdd:
			log.Info("ZAPI: ifp ", dev)
			s.Interfaces.Register(dev)
		case pb.Op_InterfaceDelete:
			s.Interfaces.Unregister(dev)
		}
	case *pb.RouterIdUpdate:
	case *pb.Route:
	}
}

func (c *Client) Stop() {
	log.Info("Stop")
	c.running = false
	if c.conn != nil {
		c.conn.Close()
	}
}
