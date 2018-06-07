// Copyright 2018 OpenConfigd Project.
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
    "io/ioutil"
    "net"
    "testing"
    "time"
)

func TestQuaggaInit(t *testing.T) {
    var (
        qa   QuaggaAccess
        vrf  string = "foobar"
        addr net.IP = net.ParseIP("127.0.0.1")
        prog string = "bgpd"
        tm   time.Duration = time.Second
        cmd  string = `#!/bin/sh
VRF=%s LD_PRELOAD=/usr/bin/vrf_socket.so telnet '%s' '%s'
`
    )

    cmd = fmt.Sprintf(cmd, vrf, addr.String(), prog)

    if err := qa.Init(vrf, addr, prog, tm); err != nil {
        t.Fatalf("QuaggaAccess.Init(): %v\n", err)
    }
    if dat, err := ioutil.ReadFile(qa.GetTempFileName()); err == nil {
        qa.Close()
        file := string(dat)
        if cmd != file {
            t.Fail()
            t.Logf("cmd:\n%v", cmd)
            t.Fatalf("\nfile:\n%v", file)
        }
    } else {
        t.Fatalf("ReadFile(%s)\n", qa.GetTempFileName())
    }
}
