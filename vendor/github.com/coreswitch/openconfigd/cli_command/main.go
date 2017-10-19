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

package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	rpc "github.com/coreswitch/openconfigd/proto"
	"google.golang.org/grpc"
)

func argString() string {
	if len(flag.Args()) != 0 {
		return strings.Join(flag.Args(), " ")
	} else {
		return ""
	}
}

func privilegeGet() (priv uint32) {
	priv = 1
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "CLI_PRIVILEGE" {
			i, _ := strconv.Atoi(pair[1])
			priv = uint32(i)
		}
	}
	return
}

func show(port int, json bool) {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	line := argString()
	args := flag.Args()
	if len(args) > 0 && args[0] == "run" {
		args = args[1:]
		line = strings.Join(args, " ")
	}

	client := rpc.NewShowClient(conn)
	req := &rpc.ShowRequest{
		Json: json,
		Line: line,
	}

	stream, err := client.Show(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Show")
	for {
		reply, err := stream.Recv()
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			return
		}
		fmt.Print(reply.Str)
	}
}

func output(reply *rpc.ExecReply) {
	switch reply.Code {
	case rpc.ExecCode_SHOW:
		fmt.Println("Show")
	}
	fmt.Println(reply.Lines)
}

func redirect(req *rpc.ExecRequest, port uint32) {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", port), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Fail to dial:", err)
		return
	}
	defer conn.Close()

	if len(req.Args) > 0 && req.Args[0] == "run" {
		req.Args = req.Args[1:]
		req.Line = strings.Join(req.Args, " ")
	}
	client := rpc.NewExecClient(conn)
	reply, err := client.DoExec(context.Background(), req)
	if err != nil {
		fmt.Println("client DoExec failed:", err)
		return
	}
	output(reply)
}

func exec(conn *grpc.ClientConn, modeFlag string, jsonFlag bool) {
	client := rpc.NewExecClient(conn)

	req := new(rpc.ExecRequest)
	req.Type = rpc.ExecType_EXEC
	req.Privilege = privilegeGet()
	req.Mode = modeFlag
	req.Line = argString()
	req.Args = flag.Args()

	reply, err := client.DoExec(context.Background(), req)
	if err != nil {
		fmt.Println("client DoExec failed:", err)
		return
	}

	switch reply.Code {
	case rpc.ExecCode_REDIRECT:
		redirect(req, reply.Port)
	case rpc.ExecCode_REDIRECT_SHOW:
		show(int(reply.Port), jsonFlag)
	default:
		output(reply)
	}
}

func completion(conn *grpc.ClientConn, compFlag bool, trailFlag bool, firstFlag bool, modeFlag string) {
	client := rpc.NewExecClient(conn)
	req := new(rpc.ExecRequest)
	req.Privilege = privilegeGet()
	req.Mode = modeFlag
	req.Line = argString()
	req.Args = flag.Args()

	switch {
	case compFlag:
		req.Type = rpc.ExecType_COMPLETE
	case trailFlag:
		req.Type = rpc.ExecType_COMPLETE_TRAILING_SPACE
	case firstFlag:
		req.Type = rpc.ExecType_COMPLETE_FIRST_COMMANDS
	}

	reply, err := client.DoExec(context.Background(), req)
	if err != nil {
		fmt.Println("client DoExec failed:", err)
		return
	}
	fmt.Println(reply.Lines)
}

func main() {
	var (
		compFlag  bool
		trailFlag bool
		firstFlag bool
		modeFlag  string
		showFlag  bool
		portFlag  int
		jsonFlag  bool
	)
	flag.BoolVar(&compFlag, "c", false, "Completion of the command")
	flag.BoolVar(&firstFlag, "f", false, "First commands list")
	flag.BoolVar(&trailFlag, "t", false, "Command has trailing space")
	flag.StringVar(&modeFlag, "m", "exec", "Current mode")
	flag.BoolVar(&showFlag, "s", false, "Show command flag")
	flag.IntVar(&portFlag, "p", 2650, "Show command port")
	flag.BoolVar(&jsonFlag, "j", false, "Show output in JSON format")
	flag.Parse()

	// Connect to service.
	conn, err := grpc.Dial(fmt.Sprintf(":%d", portFlag), grpc.WithInsecure())
	if err != nil {
		fmt.Println("Fail to dial:", err)
		return
	}
	defer conn.Close()

	// Figure out which gRPC service to use.
	switch {
	case showFlag:
		show(portFlag, jsonFlag)
	case compFlag || trailFlag || firstFlag:
		completion(conn, compFlag, trailFlag, firstFlag, modeFlag)
	default:
		exec(conn, modeFlag, jsonFlag)
	}
}
