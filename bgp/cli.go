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

var cliShowJson = `
[
    {
        "name": "showIpBgp",
        "line": "show ip bgp",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Display IP information",
            "Display BGP status and configuration"
        ]
    },
    {
        "name": "showIpBgpSummary",
        "line": "show ip bgp summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Display IP information",
            "Display BGP status and configuration",
            "Display summarized information of BGP state"
        ]
    },
    {
        "name": "showIpBgpNeighbors",
        "line": "show ip bgp neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Display IP information",
            "Display BGP status and configuration",
            "Display all configured BGP neighbors"
        ]
    }
]
`
var cliOperJson = `
[
    {
        "name": "clearIpBgp",
        "line": "clear ip bgp all",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP",
            "Clear BGP connections",
            "All of neighbors"
        ]
    }
]
`

var cliShowFuncMap = map[string]func(*Server, *ShowTask, []interface{}){
	"showIpBgp":          showIpBgp,
	"showIpBgpSummary":   showIpBgpSummary,
	"showIpBgpNeighbors": showIpBgpNeighbors,
}

var cliOperFuncMap = map[string]func(*Server, []interface{}) string{
	"clearIpBgp": clearIpBgp,
}
