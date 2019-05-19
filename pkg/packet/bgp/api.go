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

func BgpAsApi(As uint32) {
	// If As == 0 return invalid argument error.

	// If As is configured.
	// If As != Config.As, return BGP instance is already configured.
	// If As == Config.As nothing to do.

	// If As is not configured (that means BGP Server is not up).
	// StartBgp Server.
}

func BgpPeerApi(As uint32) {
	//bgpServer := BgpServerGet(As)
}
