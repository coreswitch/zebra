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

package rib

const (
	IF_TYPE_UNKNOWN uint8 = iota
	IF_TYPE_ETHERNET
	IF_TYPE_LOOPBACK
	IF_TYPE_HDLC
	IF_TYPE_PPP
	IF_TYPE_ATM
	IF_TYPE_FRELAY
	IF_TYPE_VLAN
	IF_TYPE_IPIP
	IF_TYPE_IPGRE
	IF_TYPE_IP6GRE
	IF_TYPE_IP6IP
)

var IfTypeStringMap = map[uint8]string{
	IF_TYPE_UNKNOWN:  "Unknown",
	IF_TYPE_ETHERNET: "Ethernet",
	IF_TYPE_LOOPBACK: "Loopback",
	IF_TYPE_HDLC:     "HDLC",
	IF_TYPE_PPP:      "PPP",
	IF_TYPE_ATM:      "ATM",
	IF_TYPE_FRELAY:   "Frame Relay",
	IF_TYPE_VLAN:     "VLAN",
	IF_TYPE_IPIP:     "IP over IP Tunnel",
	IF_TYPE_IPGRE:    "IP over GRE Tunnel",
	IF_TYPE_IP6GRE:   "IPv6 over GRE Tunnel",
	IF_TYPE_IP6IP:    "IPv6 over IP Tunnel",
}
