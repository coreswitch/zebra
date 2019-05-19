// Copyright 2016 Zebra Project
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

package ospf

const (
	IFSM_DependUpon = iota
	IFSM_Down
	IFSM_Loopback
	IFSM_Waiting
	IFSM_PointToPoint
	IFSM_DROther
	IFSM_Backup
	IFSM_DR
	IFSM_StateMax
)

const (
	IFSM_NoEvent = iota
	IFSM_InterfaceUp
	IFSM_WaitTimer
	IFSM_BackupSeen
	IFSM_NeighborChange
	IFSM_LoopInd
	IFSM_UnloopInd
	IFSM_InterfaceDown
	IFSM_EventMax
)
