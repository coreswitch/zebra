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
	NFSM_DependUpon = iota
	NFSM_Down
	NFSM_Attempt
	NFSM_Init
	NFSM_TwoWay
	NFSM_ExStart
	NFSM_Exchange
	NFSM_Loading
	NFSM_Full
	NFSM_StateMax
)

const (
	NFSM_NoEvent = iota
	NFSM_HelloReceived
	NFSM_Start
	NFSM_TwoWayReceived
	NFSM_NegotiationDone
	NFSM_ExchangeDone
	NFSM_BadLSReq
	NFSM_LoadingDone
	NFSM_AdjOK
	NFSM_SeqNumberMismatch
	NFSM_OneWayReceived
	NFSM_KillNbr
	NFSM_InactivityTimer
	NFSM_LLDown
	NFSM_EventMax
)
