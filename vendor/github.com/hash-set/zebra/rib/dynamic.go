// Copyright 2016 Zebra Project
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

	"github.com/coreswitch/cmd"
	rpc "github.com/hash-set/openconfigd/proto"
	zebra "github.com/hash-set/zebra/proto"
	//"github.com/hash-set/zebra/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func (s *execServer) DoExec(_ context.Context, req *rpc.ExecRequest) (*rpc.ExecReply, error) {
	reply := new(rpc.ExecReply)

	if req.Type == rpc.ExecType_COMPLETE_DYNAMIC {
		arg := ""
		if len(req.Args) > 0 {
			arg = req.Args[0]
		}
		switch arg {
		case "interface":
			InterfaceIterate(
				func(ifp *Interface) {
					reply.Candidates = append(reply.Candidates, ifp.Name)
				})
		case "vrf":
			if len(req.Args) > 1 && len(req.Commands) > 3 {
				if vrf := VrfLookupByName(req.Commands[3]); vrf != nil {
					vrf.InterfaceIterate(
						func(ifp *Interface) {
							reply.Candidates = append(reply.Candidates, ifp.Name)
						})
				}
			} else {
				for key, _ := range VrfMap {
					if key != "" {
						reply.Candidates = append(reply.Candidates, key)
					}
				}
			}
		}
	}

	return reply, nil
}

type execModuleServer struct{}

func newExecModuleServer() *execModuleServer {
	return &execModuleServer{}
}

func (s *execModuleServer) DoExecModule(_ context.Context, req *rpc.ExecModuleRequest) (*rpc.ExecModuleReply, error) {
	reply := new(rpc.ExecModuleReply)
	return reply, nil
}

type execServer struct{}

func newExecServer() *execServer {
	s := &execServer{}
	return s
}

// Show framework.
type ShowTask struct {
	Json     bool
	First    bool
	Continue bool
	Str      string
	Index    interface{}
}

type ShowServer struct {
}

func NewShowServer() *ShowServer {
	return &ShowServer{}
}

func NewShowTask() *ShowTask {
	return &ShowTask{
		First: true,
	}
}

type ShowFunc func(*ShowTask, []interface{})

func (s *ShowServer) Show(req *rpc.ShowRequest, stream rpc.Show_ShowServer) error {
	reply := &rpc.ShowReply{}

	result, fn, args, _ := ShowParser.ParseLine(req.Line)
	if result != cmd.ParseSuccess || fn == nil {
		reply.Str = "% Command can't find: \"" + req.Line + "\"\n"
		err := stream.Send(reply)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}

	show := fn.(func(*ShowTask, []interface{}))
	task := NewShowTask()
	task.Json = req.Json
	for {
		task.Str = ""
		task.Continue = false
		show(task, args)
		task.First = false

		reply.Str = task.Str
		err := stream.Send(reply)
		if err != nil {
			fmt.Println(err)
			break
		}
		if !task.Continue {
			break
		}
	}
	return nil
}

type ZebraServer struct{}

func NewZebraServer() *ZebraServer {
	return &ZebraServer{}
}

func (z *ZebraServer) RouteIPv4Add(_ context.Context, req *zebra.RouteIPv4) (*zebra.RouteIPv4Response, error) {
	return nil, nil
}

func (z *ZebraServer) RouteIPv4Delete(_ context.Context, req *zebra.RouteIPv4) (*zebra.RouteIPv4Response, error) {
	return nil, nil
}

func (z *ZebraServer) RouteIPv6Add(_ context.Context, req *zebra.RouteIPv6) (*zebra.RouteIPv6Response, error) {
	return nil, nil
}

func (z *ZebraServer) RouteIPv6Delete(_ context.Context, req *zebra.RouteIPv6) (*zebra.RouteIPv6Response, error) {
	return nil, nil
}

func RpcServer() {
	lis, err := net.Listen("tcp", ":2601")
	if err != nil {
		grpclog.Fatalf("failed %v", err)
	}

	grpcServer := grpc.NewServer()

	// Openconfigd API
	rpc.RegisterExecServer(grpcServer, newExecServer())
	rpc.RegisterExecModuleServer(grpcServer, newExecModuleServer())
	rpc.RegisterShowServer(grpcServer, NewShowServer())

	// Zebra API
	zebra.RegisterZebraApiServer(grpcServer, NewZebraServer())

	grpcServer.Serve(lis)
}
