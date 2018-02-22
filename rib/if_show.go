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

package rib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"text/template"
)

const IfShowTemplate = `Interface {{.Name}}
  Hardware is {{.Type}}{{if .HwAddr}}, address is {{.HwAddr}}{{end}}
{{if .Desc}}  Description: {{.Desc}}
{{end}}  index {{.Index}} metric {{.Metric}} mtu {{.Mtu}}
  <{{.Flags}}>
  VRF Binding: {{if .Vrf}}{{.Vrf}}{{else}}Not bound{{end}}
  Label switching is disabled
{{range $i, $v := .AddressIPv4}}  inet {{$v.Prefix}}
{{end}}{{range $i, $v := .AddressIPv6}}  inet6 {{$v.Prefix}}
{{end}}`

type IfShow struct {
	Name        string      `json:"name,omitempty"`
	Type        string      `json:"type,omitempty"`
	HwAddr      string      `json:"hardware-adddress,omitempty"`
	Desc        string      `json:"description,omitempty"`
	Index       IfIndex     `json:"index,omitempty"`
	Metric      uint32      `json:"metric,omitempty"`
	Mtu         uint32      `json:"mtu,omitempty"`
	Flags       string      `json:"flags,omitempty"`
	Vrf         string      `json:"vrf,omitempty"`
	AddressIPv4 IfAddrSlice `json:"ipv4-address,omitempty"`
	AddressIPv6 IfAddrSlice `json:"ipv6-address,omitempty"`
}

func IfVrfString(vrfId uint32) string {
	if vrfId == 0 {
		return ""
	}
	vrf := VrfLookupByIndex(vrfId)
	if vrf != nil {
		return vrf.Name
	}
	return fmt.Sprintf("VRF Interface Index %d", vrfId)
}

func (ifp *Interface) Show(jsonFlag bool) string {
	ifShow := &IfShow{
		Name:        ifp.Name,
		Type:        IfTypeStringMap[ifp.IfType],
		HwAddr:      net.HardwareAddr(ifp.HwAddr).String(),
		Desc:        ifp.Description,
		Index:       ifp.Index,
		Metric:      ifp.Metric,
		Mtu:         ifp.Mtu,
		Flags:       IfFlagsString(ifp.Flags),
		Vrf:         IfVrfString(ifp.VrfIndex),
		AddressIPv4: ifp.Addrs[AFI_IP],
		AddressIPv6: ifp.Addrs[AFI_IP6],
	}
	if jsonFlag {
		buf, err := json.Marshal(ifShow)
		if err != nil {
			return fmt.Sprintln(err)
		}
		return string(buf)
	} else {
		buf := new(bytes.Buffer)
		tmpl := template.Must(template.New("IfShow").Parse(IfShowTemplate))
		tmpl.Execute(buf, ifShow)
		str := buf.String()
		str += ifp.ShowStats()
		return str
	}
}

func InterfaceShow(jsonFlag bool, Args ...string) string {
	IfStatsUpdate()

	if len(Args) != 0 {
		ifp := IfLookupByName(Args[0])
		if ifp != nil {
			return ifp.Show(jsonFlag)
		}
	}

	str := ""
	first := true
	for n := IfTable.Top(); n != nil; n = IfTable.Next(n) {
		ifp := n.Item.(*Interface)

		if first {
			first = false
		} else {
			if jsonFlag {
				str += ","
			}
		}
		str += ifp.Show(jsonFlag)
	}
	if jsonFlag {
		str = `{"interfaces":[` + str + `]}`
	}
	return str
}

func (v *Vrf) InterfaceShow(jsonFlag bool, Args ...string) (line string) {
	IfStatsUpdate()

	if len(Args) != 0 {
		ifp := v.IfLookupByName(Args[0])
		if ifp != nil {
			return ifp.Show(jsonFlag)
		}
	}

	str := ""
	first := true
	for n := v.IfTable.Top(); n != nil; n = v.IfTable.Next(n) {
		ifp := n.Item.(*Interface)

		if first {
			first = false
		} else {
			if jsonFlag {
				str += ","
			}
		}
		str += ifp.Show(jsonFlag)
	}
	if jsonFlag {
		str = `{"interfaces":[` + str + `]}`
	}
	return str
}
