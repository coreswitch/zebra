package quagga

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/coreswitch/cmd"
	rpc "github.com/coreswitch/openconfigd/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"io"
	"math/rand"
	"net"
	"strings"
	"time"
)

const (
	grpcConnRetryInterval = 5
)

var (
	showParser *cmd.Node
	execParser *cmd.Node
	grpcServer *grpc.Server
)

type grpcExecServer struct {
}

func (s *grpcExecServer) DoExec(_ context.Context, request *rpc.ExecRequest) (*rpc.ExecReply, error) {
	reply := &rpc.ExecReply{}
	switch request.Type {
	case rpc.ExecType_COMPLETE_DYNAMIC:
		arg := ""
		if len(request.Args) > 0 {
			arg = request.Args[0]
		}
		switch arg {
		case "neighbor":
			// XXX
		}
	case rpc.ExecType_EXEC:
		_, fn, _, _ := execParser.ParseLine(request.Line)
		exec := fn.(func(string) *string)
		out := exec(request.Line)
		reply.Code = rpc.ExecCode_SHOW
		reply.Lines = *out
	}
	return reply, nil
}

type grpcShowServer struct {
}

func (s *grpcShowServer) Show(request *rpc.ShowRequest, stream rpc.Show_ShowServer) error {
	reply := &rpc.ShowReply{}
	_, fn, _, _ := showParser.ParseLine(request.Line)
	show := fn.(func(string) *string)
	out := show(request.Line)
	lines := strings.Split(*out, "\n")
	for _, line := range lines {
		reply.Str = fmt.Sprintln(line)
		err := stream.Send(reply)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	return nil
}

func grpcRegisterModule(conn *grpc.ClientConn) error {
	client := rpc.NewRegisterClient(conn)
	request := &rpc.RegisterModuleRequest{
		Module: QUAGGAD_MODULE,
		Port:   fmt.Sprintf("%d", QUAGGAD_PORT),
	}
	_, err := client.DoRegisterModule(context.Background(), request)
	if err != nil {
		return err
	}
	return nil
}

func grpcRegisterCommands(conn *grpc.ClientConn) {
	client := rpc.NewRegisterClient(conn)

	var showCommands []rpc.RegisterRequest
	json.Unmarshal([]byte(showCmdSpec), &showCommands)

	showParser = cmd.NewParser()

	for _, command := range showCommands {
		command.Module = QUAGGAD_MODULE
		command.Privilege = 1
		command.Code = rpc.ExecCode_REDIRECT_SHOW

		_, err := client.DoRegister(context.Background(), &command)
		if err != nil {
			grpclog.Fatalf("client DoRegister failed: %v", err)
		}
		showParser.InstallLine(command.Line, showCmdMap[command.Name])
	}

	var execCommands []rpc.RegisterRequest
	json.Unmarshal([]byte(execCmdSpec), &execCommands)

	execParser = cmd.NewParser()

	for _, command := range execCommands {
		command.Module = QUAGGAD_MODULE
		command.Privilege = 1
		command.Code = rpc.ExecCode_REDIRECT

		_, err := client.DoRegister(context.Background(), &command)
		if err != nil {
			grpclog.Fatalf("client DoRegister failed: %v", err)
		}
		execParser.InstallLine(command.Line, execCmdMap[command.Name])
	}
}

func grpcSubscribe(conn *grpc.ClientConn) (rpc.Config_DoConfigClient, error) {
	client := rpc.NewConfigClient(conn)
	stream, err := client.DoConfig(context.Background())
	if err != nil {
		return nil, err
	}

	path := []string{"interfaces", "protocols", "policy"}
	request := &rpc.ConfigRequest{
		Type:   rpc.ConfigType_SUBSCRIBE_MULTI,
		Module: QUAGGAD_MODULE,
		Port:   QUAGGAD_PORT,
		Path:   path,
	}
	err = stream.Send(request)
	if err != nil {
		return nil, err
	}
	return stream, err
}

func grpcLoop() {
	for {
		conn, err := grpc.Dial(":2650",
			grpc.WithInsecure(),
			grpc.FailOnNonTempDialError(true),
			grpc.WithBlock(),
			grpc.WithTimeout(time.Second*grpcConnRetryInterval),
		)
		defer conn.Close()
		if err != nil {
			interval := (rand.Intn(grpcConnRetryInterval) + 1)
			select {
			case <-time.After(time.Second * time.Duration(interval)):
				// Wait timeout.
			}
			continue
		}
		grpcRegisterModule(conn)
		grpcRegisterCommands(conn)
		stream, err := grpcSubscribe(conn)
		if err != nil {
			interval := (rand.Intn(grpcConnRetryInterval) + 1)
			select {
			case <-time.After(time.Second * time.Duration(interval)):
				// Wait timeout.
			}
			continue
		}
		for {
			conf, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				break
			}
			switch conf.Type {
			case rpc.ConfigType_VALIDATE_START:
				fmt.Println("VALIDATE_START:", conf.Path)
				validating = true
				quaggaConfigStateSync()
				//fmt.Println("***")
				//fmt.Println("*** rpc.ConfigType_VALIDATE_START ***")
				//fmt.Println("***")
				//quaggaConfigStateDump()
			case rpc.ConfigType_VALIDATE_END:
				fmt.Println("VALIDATE_END:", conf.Path)

				b := bytes.NewBuffer(make([]byte, 0))
				f := bufio.NewWriter(b)

				request := &rpc.ConfigRequest{}
				if quaggaConfigValid(f) {
					request.Type = rpc.ConfigType_VALIDATE_SUCCESS
				} else {
					request.Type = rpc.ConfigType_VALIDATE_FAILED
				}
				f.Flush()
				fmt.Print(b.String())
				err := stream.Send(request)
				if err != nil {
					fmt.Println(err)
				}
				//fmt.Println("***")
				//fmt.Println("*** rpc.ConfigType_VALIDATE_END ***")
				//fmt.Println("***")
				//quaggaConfigStateDump()
				validating = false
			case rpc.ConfigType_COMMIT_START:
				fmt.Println("COMMIT_START:", conf.Path)
				commitedAccessList = false
				commitedAccessList6 = false
				commitedAsPathList = false
				commitedCommunityList = false
				commitedPrefixList = false
				commitedPrefixList6 = false
				commitedRouteMap = false
				//fmt.Println("***")
				//fmt.Println("*** rpc.ConfigType_COMMIT_START ***")
				//fmt.Println("***")
				//quaggaConfigStateDump()
			case rpc.ConfigType_COMMIT_END:
				fmt.Println("COMMIT_END:", conf.Path)
				//fmt.Println("***")
				//fmt.Println("*** rpc.ConfigType_COMMIT_END ***")
				//fmt.Println("***")
				//quaggaConfigStateDump()
			case rpc.ConfigType_SET, rpc.ConfigType_DELETE:
				quaggaConfig(int(conf.Type), conf.Path)
			default:
			}
		}
	}
}

func initGrpc() {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", QUAGGAD_PORT))
	if err != nil {
		grpclog.Fatalf("failed %v", err)
	}

	grpcServer = grpc.NewServer()

	// Openconfigd API
	rpc.RegisterExecServer(grpcServer, &grpcExecServer{})
	rpc.RegisterShowServer(grpcServer, &grpcShowServer{})

	go grpcServer.Serve(listen)
	go grpcLoop()
}
