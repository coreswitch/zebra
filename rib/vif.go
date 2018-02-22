// Copyright 2017 zebra project.
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

import "fmt"

type VIF struct {
	Id uint64
}

func NewVIF(vlanId uint64) *VIF {
	vif := &VIF{
		Id: vlanId,
	}
	return vif
}

func (ifp *Interface) VIFLookup(vlanId uint64) *VIF {
	for _, vif := range ifp.VIFs {
		if vif.Id == vlanId {
			return vif
		}
	}
	return nil
}

func (ifp *Interface) UnregisterVIF(vlanId uint64) *VIF {
	var del *VIF
	var vifs []*VIF
	for _, vif := range ifp.VIFs {
		if vif.Id == vlanId {
			del = vif
		} else {
			vifs = append(vifs, vif)
		}
	}
	ifp.VIFs = vifs
	return del
}

func VIFClean() {
	type VifList struct {
		name string
		id   int
	}
	var vifList []*VifList

	for _, ifp := range IfMap {
		for _, vif := range ifp.VIFs {
			vifList = append(vifList, &VifList{
				name: fmt.Sprintf("%s.%d", ifp.Name, vif.Id),
				id:   int(vif.Id),
			})
		}
	}
	for _, vif := range vifList {
		NetlinkVlanDelete(vif.name, vif.id)
	}
}
