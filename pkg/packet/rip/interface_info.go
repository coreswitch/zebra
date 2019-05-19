// Copyright 2018 zebra project.
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

package rip

import "github.com/coreswitch/zebra/fea"

type InterfaceInfo struct {
	ifp       *fea.Interface
	ifEnabled bool
	passive   bool
	up        bool
}

func (s *Server) InterfaceInfoGet(ifp *fea.Interface) *InterfaceInfo {
	if ifp.Info != nil {
		return ifp.Info.(*InterfaceInfo)
	}
	ifi := &InterfaceInfo{
		ifp: ifp,
	}
	ifp.Info = ifi
	return ifi
}
