// Copyright 2016 OpenConfigd Project.
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

package config

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	"github.com/coreswitch/cmd"
	"github.com/coreswitch/component"
	rpc "github.com/coreswitch/openconfigd/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	CliSuccess = iota
	CliSuccessExec
	CliSuccessShow
	CliSuccessModule
	CliSuccessRedirect
	CliSuccessRedirectShow
)

var (
	RIBD_SYNCHRONIZED bool
)

var GrpcModuleMap = map[string]string{}

type openconfigServer struct {
	cmd *cmd.Cmd
}

type registerServer struct {
	cmd *cmd.Cmd
}

type configServer struct {
	cmd *cmd.Cmd
}

func execute(Cmd *cmd.Cmd, mode string, args []string, line string, reply *rpc.ExecReply) (ret string) {
	result, fn, pargs, _ := Cmd.ParseLine(mode, line, &cmd.Param{Command: args})

	fmt.Println(args)
	if RIBD_SYNCHRONIZED {
		if len(args) >= 4 {
			if args[0] == "set" && args[1] == "interfaces" && args[2] == "interface" {
				conf := configActive.LookupByPath([]string{"interfaces", "interface", args[3]})
				if conf == nil {
					ret = "NoMatch\n"
					return
				}
			}
		}
	}
	switch result {
	case cmd.ParseSuccess:
		cb, ok := fn.(func([]string) (int, string))
		if ok {
			inst, instStr := cb(cmd.Interface2String(pargs))
			switch inst {
			case CliSuccess:
				ret = "Success"
			case CliSuccessExec:
				ret = "SuccessExec\n"
				ret += instStr
			case CliSuccessShow:
				reply.Code = rpc.ExecCode_SHOW
				//ret = "SuccessShow\n"
				ret = instStr
			case CliSuccessModule:
				ret = "SuccessModule\n"
				ret += instStr
			case CliSuccessRedirect, CliSuccessRedirectShow:
				if inst == CliSuccessRedirectShow {
					reply.Code = rpc.ExecCode_REDIRECT_SHOW
				} else {
					reply.Code = rpc.ExecCode_REDIRECT
				}
				port, _ := strconv.ParseUint(instStr, 10, 32)
				reply.Port = uint32(port)
				ret = "SuccessRedirect\n"
				ret += instStr
			}
		}
	case cmd.ParseIncomplete:
		ret = "Incomplete\n"
	case cmd.ParseNoMatch:
		ret = "NoMatch\n"
	case cmd.ParseAmbiguous:
		ret = "Ambiguous\n"
	}
	return
}

func complete(Cmd *cmd.Cmd, mode string, args []string, line string, trailing bool) (ret string) {
	result, _, _, comps := Cmd.ParseLine(mode, line, &cmd.Param{Command: args, Complete: true, TrailingSpace: trailing})

	switch result {
	case cmd.ParseSuccess, cmd.ParseIncomplete:
		ret = "Success\n"
	case cmd.ParseNoMatch:
		ret = "NoMatch\n"
	case cmd.ParseAmbiguous:
		ret = "Ambiguous\n"
	}

	for _, comp := range comps {
		var pre string
		if !comp.Dir && !comp.Additive {
			pre = "--"
		} else {
			if comp.Additive {
				pre = "+"
			} else {
				pre = " "
			}
			if comp.Dir {
				pre += ">"
			} else {
				pre += " "
			}
		}
		ret += fmt.Sprintf("%s\t%s\t%s\n", comp.Name, pre, comp.Help)
		//ret += comp.Name + "\t" + "" + "\t" + comp.Help + "\n"
	}

	if result == cmd.ParseSuccess {
		ret += "<cr>\t--\t\n"
	}

	return
}

func unquote(req *rpc.ExecRequest) {
	if req.Type == rpc.ExecType_COMPLETE || req.Type == rpc.ExecType_COMPLETE_TRAILING_SPACE {
		for pos, arg := range req.Args {
			arg, err := strconv.Unquote(arg)
			if err == nil {
				req.Args[pos] = arg
			}
		}
	}
}

func ExecLine(line string) string {
	reply := new(rpc.ExecReply)
	args := strings.Split(line, " ")
	return execute(TopCmd, "configure", args, line, reply)
}

func (s *openconfigServer) DoExec(_ context.Context, req *rpc.ExecRequest) (*rpc.ExecReply, error) {
	reply := new(rpc.ExecReply)

	unquote(req)

	switch req.Type {
	case rpc.ExecType_EXEC:
		reply.Lines = execute(s.cmd, req.Mode, req.Args, req.Line, reply)
	case rpc.ExecType_COMPLETE:
		reply.Lines = complete(s.cmd, req.Mode, req.Args, req.Line, false)
	case rpc.ExecType_COMPLETE_TRAILING_SPACE:
		reply.Lines = complete(s.cmd, req.Mode, req.Args, req.Line+" ", true)
	case rpc.ExecType_COMPLETE_FIRST_COMMANDS:
		reply.Lines = s.cmd.FirstCommands(req.Mode, req.Privilege)
	case rpc.ExecType_COMPLETE_DYNAMIC:
		// Ignore dynamic completion in openconfigd.
	}
	return reply, nil
}

func (s *registerServer) DoRegister(_ context.Context, req *rpc.RegisterRequest) (*rpc.RegisterReply, error) {
	reply := new(rpc.RegisterReply)
	port := GrpcModuleMap[req.Module]
	if port == "" {
		port = "2601"
	} else {
		//fmt.Println("Set port", port)
	}

	inst := CliSuccessRedirect
	if req.Code == rpc.ExecCode_REDIRECT_SHOW {
		inst = CliSuccessRedirectShow
	}

	if mode := s.cmd.LookupMode(req.Mode); mode != nil {
		mode.InstallLine(req.Line,
			func(Args []string) (int, string) {
				return inst, port
			},
			&cmd.Param{Helps: req.Helps, Privilege: req.Privilege, Dynamic: true})
	}
	return reply, nil
}

func (s *registerServer) DoRegisterModule(_ context.Context, req *rpc.RegisterModuleRequest) (*rpc.RegisterModuleReply, error) {
	reply := new(rpc.RegisterModuleReply)
	GrpcModuleMap[req.Module] = req.Port
	return reply, nil
}

func (s *configServer) DoConfig(stream rpc.Config_DoConfigServer) error {
loop:
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("EOF")
			break loop
		}
		if err != nil {
			fmt.Println(err)
			break loop
		}
		switch msg.Type {
		case rpc.ConfigType_SUBSCRIBE:
			go SubscribeRemoteAdd(stream, msg)
		case rpc.ConfigType_SUBSCRIBE_MULTI:
			go SubscribeRemoteAddMulti(stream, msg)
		case rpc.ConfigType_SUBSCRIBE_REQUEST:
			go SubscribeAdd(stream, msg)
		case rpc.ConfigType_SET:
			YangConfigPush(msg.Path)
		case rpc.ConfigType_DELETE:
			YangConfigPull(msg.Path)
		case rpc.ConfigType_VALIDATE_SUCCESS:
			//fmt.Println("Validate Success")
			SubscribeValidateProcess(stream, msg.Type)
		case rpc.ConfigType_VALIDATE_FAILED:
			//fmt.Println("Validate Failed")
			SubscribeValidateProcess(stream, msg.Type)
		}
	}
	if stream != nil {
		SubscribeDelete(stream)
	}

	return nil
}

func DynamicCompletionLocal(commands []string, module string, args []string) []string {
	if len(args) > 0 {
		switch args[0] {
		case "rollback":
			return RollbackCompletion(commands)
		default:
		}
	}
	return []string{}
}

func DynamicCompletion(commands []string, module string, args []string) []string {
	// Local completion check.
	// if module == "local" {
	// 	return DynamicCompletionLocal(commands, module, args)
	// }

	// XXX Need to leverage stream connection. (No need of make a new connection)
	host := ":2601" // Default port.
	port := SubscribePortLookup(module)
	if port != 0 {
		host = fmt.Sprintf(":%d", port)
	}
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		fmt.Println("DynamicCompletion: Fail to dial", err)
		return []string{}
	}
	defer conn.Close()

	client := rpc.NewExecClient(conn)
	req := &rpc.ExecRequest{
		Type:     rpc.ExecType_COMPLETE_DYNAMIC,
		Mode:     module,
		Commands: commands,
		Args:     args,
	}

	reply, err := client.DoExec(context.Background(), req)
	if err != nil {
		fmt.Println("client DoExec COMPLETE_DYNAMIC failed:", err)
		return []string{}
	}
	return reply.Candidates
}

// RPC component.
type RpcComponent struct {
	GrpcEndpoint string
}

// RPC component start method.
func (this *RpcComponent) Start() component.Component {
	lis, err := net.Listen("tcp", this.GrpcEndpoint)
	if err != nil {
		grpclog.Fatalf("failed %v", err)
	}
	grpcServer := grpc.NewServer()
	rpc.RegisterExecServer(grpcServer, &openconfigServer{TopCmd})
	rpc.RegisterRegisterServer(grpcServer, &registerServer{TopCmd})
	rpc.RegisterConfigServer(grpcServer, &configServer{TopCmd})
	grpcServer.Serve(lis)
	return this
}

func (this *RpcComponent) Stop() component.Component {
	// fmt.Println("rpc component stop")
	return this
}
