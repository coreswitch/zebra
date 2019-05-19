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

const (
	RIPv1               uint8 = 1
	RIPv2               uint8 = 2
	RIP_PORT_DEFAULT          = 520
	RIP_PACKET_MINSIZE        = 4
	INADDR_RIP_GROUP          = "224.0.0.9"
	RIP_METRIC_INFINITY       = 16
)

const (
	RIP_REQUEST     = 1
	RIP_RESPONSE    = 2
	RIP_TRACEON     = 3 // Obsolete
	RIP_TRACEOFF    = 4 // Obsolete
	RIP_POLL        = 5
	RIP_POLL_ENTRY  = 6
	RIP_COMMAND_MAX = 7
)

const (
	VRF_DEFAULT = 0
)
