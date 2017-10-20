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
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"
)

type ShowTask struct {
	Json     bool
	First    bool
	Continue bool
	Str      string
	Index    interface{}
}

func NewShowTask() *ShowTask {
	return &ShowTask{
		First: true,
	}
}

func showIpBgp(s *Server, t *ShowTask, Args []interface{}) {
	t.Str = "show ip bgp"
}

const showBgpSummaryTemplate = `BGP summary information for VRF {{.VrfName}}, address family {{.AfiSafi}}
BGP router identifier {{.RouterId}}, local AS number {{.LocalAs}}
BGP table version is {{.TblVer}}, IPv4 Unicast config peers {{.ConfigPeers}}, capable peers {{.CapPeers}}
`

func uptimeString(cur, last time.Time) string {
	dur := cur.Sub(last)

	hours0 := int(dur.Hours())
	days := hours0 / 24
	hours := hours0 % 24
	mins := int(dur.Minutes()) % 60
	secs := int(dur.Seconds()) % 60

	if days == 0 {
		return fmt.Sprintf("%02d:%02d:%02d", hours, mins, secs)
	} else if days < 7 {
		return fmt.Sprintf("%dd%02dh%02dm", days, hours, mins)
	} else {
		return fmt.Sprintf("%02dw%dd%02dh", days/7, days-((days/7)*7), hours)
	}
}

func stateOrPrefixString(n *Neighbor) string {
	if n.fsm.state != BGP_FSM_ESTABLISHED {
		return BgpStateString(n.fsm.state)
	} else {
		return strconv.FormatUint(n.in.prefix, 10)
	}
}

func showIpBgpSummary(s *Server, t *ShowTask, Args []interface{}) {
	afiSafi := AfiSafiValue(AFI_IP, SAFI_UNICAST)

	tmpl := template.Must(template.New("showIpBgpSummary").Parse(showBgpSummaryTemplate))

	type TmplValue struct {
		VrfName     string
		AfiSafi     string
		RouterId    string
		LocalAs     string
		TblVer      string
		ConfigPeers int
		CapPeers    int
	}
	value := &TmplValue{
		VrfName:     "default",
		AfiSafi:     "IPv4 Unicast",
		RouterId:    s.RouterId().String(),
		LocalAs:     strconv.Itoa(int(s.as)),
		TblVer:      "1",
		ConfigPeers: 0,
		CapPeers:    0,
	}

	var doc bytes.Buffer
	tmpl.Execute(&doc, value)
	str := doc.String()

	str += "\nNeighbor        V    AS MsgRcvd MsgSent   TblVer  InQ OutQ  Up/Down State/PfxRcd\n"

	cur := time.Now()

	for _, n := range s.Neighbors {
		if _, ok := n.afiSafi[afiSafi]; ok {
			str += fmt.Sprintf("%-16s4%6d%8d%8d%9d%5d%5d%9s %s\n",
				n.addr.String(), n.RemoteAs(), n.in.Sum(), n.out.Sum(), 0, 0, 0,
				uptimeString(cur, n.uptime), stateOrPrefixString(n))
		}
	}

	t.Str = str
}

func showIpBgpNeighbors(s *Server, t *ShowTask, Args []interface{}) {
	t.Str = "show ip bgp neighbors"
}
