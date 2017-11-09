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
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/coreswitch/cmd"
	"github.com/coreswitch/component"
	"github.com/coreswitch/openconfigd/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	BGP_MODULE_NAME         = "bgpd"
	BGP_MODULE_PORT         = "2651"
	GRPC_SERVER_ADDRESS     = "" // This means localhost.
	GRPC_SERVER_PORT        = "2650"
	GRPC_CONNECT_RETRY_TIME = 5
)

type GrpcComponent struct {
	Server *Server
	conn   *grpc.ClientConn
	parser *cmd.Node
}

func GrpcDial() (*grpc.ClientConn, error) {
	return grpc.Dial(GRPC_SERVER_ADDRESS+":"+GRPC_SERVER_PORT,
		grpc.WithInsecure(),
		grpc.FailOnNonTempDialError(true),
		grpc.WithBlock(),
		grpc.WithTimeout(time.Second*GRPC_CONNECT_RETRY_TIME),
	)
}

func GrpcRegisterModule(conn *grpc.ClientConn, moduleName string, modulePort string) error {
	client := openconfig.NewRegisterClient(conn)
	req := &openconfig.RegisterModuleRequest{
		Module: moduleName,
		Port:   modulePort,
	}
	_, err := client.DoRegisterModule(context.Background(), req)
	if err != nil {
		return err
	}
	return nil
}

func GrpcRegisterCli(conn *grpc.ClientConn, moduleName string, parser *cmd.Node) error {
	client := openconfig.NewRegisterClient(conn)
	var clis []openconfig.RegisterRequest

	// Show commands.
	err := json.Unmarshal([]byte(cliShowJson), &clis)
	if err != nil {
		return err
	}
	for _, cli := range clis {
		cli.Module = moduleName
		cli.Privilege = 1
		cli.Code = openconfig.ExecCode_REDIRECT_SHOW

		_, err := client.DoRegister(context.Background(), &cli)
		if err != nil {
			return err
		}
		parser.InstallLine(cli.Line, cliShowFuncMap[cli.Name])
	}

	// Operational commands.
	err = json.Unmarshal([]byte(cliOperJson), &clis)
	if err != nil {
		return err
	}
	for _, cli := range clis {
		cli.Module = moduleName
		cli.Privilege = 1

		_, err := client.DoRegister(context.Background(), &cli)
		if err != nil {
			return err
		}
		parser.InstallLine(cli.Line, cliOperFuncMap[cli.Name])
	}
	return nil
}

func (s *GrpcComponent) Show(req *openconfig.ShowRequest, stream openconfig.Show_ShowServer) error {
	reply := &openconfig.ShowReply{}

	result, fn, args, _ := s.parser.ParseLine(req.Line)
	if result != cmd.ParseSuccess || fn == nil {
		reply.Str = "% Command can't find: \"" + req.Line + "\"\n"
		err := stream.Send(reply)
		if err != nil {
			fmt.Println(err)
		}
		return nil
	}

	show := fn.(func(*Server, *ShowTask, []interface{}))
	task := NewShowTask()
	task.Json = req.Json
	for {
		task.Continue = false
		show(s.Server, task, args)
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

func (s *GrpcComponent) DoExec(_ context.Context, req *openconfig.ExecRequest) (*openconfig.ExecReply, error) {
	reply := new(openconfig.ExecReply)
	if req.Type != openconfig.ExecType_EXEC {
		return reply, nil
	}
	ret, fn, args, _ := s.parser.ParseLine(req.Line)
	if ret == cmd.ParseSuccess {
		callback := fn.(func(*Server, []interface{}) string)
		reply.Code = openconfig.ExecCode_SHOW
		reply.Lines = callback(s.Server, args)
	}
	return reply, nil
}

func (s *GrpcComponent) StartServer() {
	lis, err := net.Listen("tcp", ":"+BGP_MODULE_PORT)
	if err != nil {
		// XXX failure recovery method needed.
		return
	}
	serv := grpc.NewServer()
	openconfig.RegisterShowServer(serv, s)
	openconfig.RegisterExecServer(serv, s)
	serv.Serve(lis)
}

func (s *GrpcComponent) StartClient() {
loop:
	conn, err := GrpcDial()
	if err != nil {
		fmt.Println("Failed to connect grpc server, retrying...")
		time.Sleep(time.Second * GRPC_CONNECT_RETRY_TIME)
		goto loop
	}
	fmt.Println("Connected to grpc server.")
	s.conn = conn
	err = GrpcRegisterModule(conn, BGP_MODULE_NAME, BGP_MODULE_PORT)
	if err != nil {
		// XXX should goto retry loop.
		fmt.Println(err)
		return
	}
	err = GrpcRegisterCli(conn, BGP_MODULE_NAME, s.parser)
	if err != nil {
		// XXX should goto retry loop.
		fmt.Println(err)
		return
	}
}

func (s *GrpcComponent) Start() component.Component {
	s.parser = cmd.NewParser()
	go s.StartServer()
	go s.StartClient()
	return s
}

func (s *GrpcComponent) Stop() component.Component {
	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
	return s
}
