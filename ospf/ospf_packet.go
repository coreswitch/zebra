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

import (
	"encoding/binary"
	"net"
)

const (
	OSPF_VERSION          = 2
	OSPF_AUTH_SIMPLE_SIZE = 8
)

const (
	OSPF_PACKET_HELLO    = 1 // OSPF Hello.
	OSPF_PACKET_DB_DESC  = 2 // OSPF Database Description.
	OSPF_PACKET_LS_REQ   = 3 // OSPF Link State Request.
	OSPF_PACKET_LS_UPD   = 4 // OSPF Link State Update.
	OSPF_PACKET_LS_ACK   = 5 // OSPF Link State Acknoledgement.
	OSPF_PACKET_TYPE_MAX = 6
)

func ospfPacketHeader(oi *OspfInterface, typ byte) []byte {
	data := make([]byte, 24)
	data[0] = OSPF_VERSION
	data[1] = typ
	// data[2:3] = Packet length.
	copy(data[4:8], oi.Top.RouterId)
	copy(data[8:12], oi.Area.AreaId)
	// data[12:14] = Check SUm
	if oi.Network != nil {
		data[14] = oi.Network.InstanceId
	} else {
		data[14] = 0
	}
	data[15] = oi.AuthType()
	// data[16:24]     // Authentication.
	return data
}

func ospfPacketHello(oi *OspfInterface, resetNeighbor bool) []byte {
	data := make([]byte, 20)

	if oi.IsUnnumbered() || oi.Type == OspfIfTypeVirtualLink {
		// Make is empty 0.0.0.0.
	} else {
		copy(data, net.CIDRMask(oi.Ident.Address.Length, 32))
	}
	binary.BigEndian.PutUint16(data[4:6], oi.HelloInterval())
	data[6] = oi.Options()
	data[7] = oi.Priority()
	binary.BigEndian.PutUint32(data[8:12], oi.DeadInterval())

	if resetNeighbor {
	}

	copy(data[12:16], oi.Ident.DRouter)
	copy(data[16:20], oi.Ident.BDRouter)

	for n := oi.Nbrs.Top(); n != nil; n = oi.Nbrs.Next(n) {
		nbr := n.Item.(*OspfNeighbor)
		if nbr.State == NFSM_Attempt || nbr.State == NFSM_Down {
			continue
		}
		//data = append(data, nbr.Ident.RouterID)
	}
	return data
}

func ospfPacketHandle() {

}
