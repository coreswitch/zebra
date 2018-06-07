// Copyright 2018 openconfigd project.
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

import "fmt"

type DistributeListEntry struct {
	Action string `mapstructure:"action" json:"action,omitempty"`
	Ge     *int   `mapstructure:"ge" json:"ge,omitempty"`
	Le     *int   `mapstructure:"le" json:"le,omitempty"`
	Prefix string `mapstructure:"prefix" json:"prefix,omitempty"`
}

type DistributeList struct {
	Primary []DistributeListEntry
	Backup  []DistributeListEntry
}

var DistributeListMap = map[int]DistributeList{}

func DistributeListSync(vrfId int, cfg *VrfsConfig) {
	// Delete existing distribute-list.
	ExecLine(fmt.Sprintf("delete prefix-list distribute-list-vrf%d", vrfId))
	ExecLine(fmt.Sprintf("delete vrf name vrf%d distribute-list-ospf distribute-list-vrf%d", vrfId, vrfId))
	if len(cfg.Ospf) == 0 {
		fmt.Println("DistributeListSync Empty ospf config returning")
		return
	}

	dlist := DistributeList{}

	for _, ospf := range cfg.Ospf {
		for _, entry := range ospf.PrimaryList {
			dlist.Primary = append(dlist.Primary, entry)
		}
	}

	// Add distribute-list.
	ExecLine(fmt.Sprintf("set vrf name vrf%d distribute-list-ospf distribute-list-vrf%d", vrfId, vrfId))

	for pos, entry := range dlist.Primary {
		if entry.Action == "" {
			entry.Action = "permit"
		}
		ExecLine(fmt.Sprintf("set prefix-list distribute-list-vrf%d seq %d action %s", vrfId, (pos+1)*5, entry.Action))
		ExecLine(fmt.Sprintf("set prefix-list distribute-list-vrf%d seq %d prefix %s", vrfId, (pos+1)*5, entry.Prefix))
		if entry.Le != nil {
			ExecLine(fmt.Sprintf("set prefix-list distribute-list-vrf%d seq %d le %d", vrfId, (pos+1)*5, *entry.Le))
		}
		if entry.Ge != nil {
			ExecLine(fmt.Sprintf("set prefix-list distribute-list-vrf%d seq %d ge %d", vrfId, (pos+1)*5, *entry.Ge))
		}
	}
/*
	for pos, entry := range dlist.Backup {
		if entry.Action == "" {
			entry.Action = "permit"
		}
		ExecLine(fmt.Sprintf("set prefix-list distribute-list-vrf%d-backup seq %d action %s", vrfId, (pos+1)*5, entry.Action))
		ExecLine(fmt.Sprintf("set prefix-list distribute-list-vrf%d-backup seq %d prefix %s", vrfId, (pos+1)*5, entry.Prefix))
		if entry.Le != nil {
			ExecLine(fmt.Sprintf("set prefix-list distribute-list-vrf%d-backup seq %d le %d", vrfId, (pos+1)*5, *entry.Le))
		}
		if entry.Ge != nil {
			ExecLine(fmt.Sprintf("set prefix-list distribute-list-vrf%d-backup seq %d ge %d", vrfId, (pos+1)*5, *entry.Ge))
		}
	}
*/
	Commit()

	DistributeListMap[vrfId] = dlist
}

func DistributeListDelete(vrfId int) {
	ExecLine(fmt.Sprintf("delete vrf name vrf%d distribute-list-ospf distribute-list-vrf%d", vrfId, vrfId))
	ExecLine(fmt.Sprintf("delete prefix-list distribute-list-vrf%d", vrfId))
	//ExecLine(fmt.Sprintf("delete prefix-list distribute-list-vrf%d-backup", vrfId))
	delete(DistributeListMap, vrfId)
}

func DistributeListExit() {
	for vrfId, _ := range DistributeListMap {
		ExecLine(fmt.Sprintf("delete vrf name vrf%d distribute-list-ospf distribute-list-vrf%d", vrfId, vrfId))
		ExecLine(fmt.Sprintf("delete prefix-list distribute-list-vrf%d", vrfId))
		//ExecLine(fmt.Sprintf("delete prefix-list distribute-list-vrf%d-backup", vrfId))
	}
	DistributeListMap = map[int]DistributeList{}
}
