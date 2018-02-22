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

const (
	EVPN_ROUTE_TYPE_ETHERNET_AUTO_DISCOVERY = 1
	EVPN_ROUTE_TYPE_MAC_IP_ADVERTISEMENT    = 2
	EVPN_INCLUSIVE_MULTICAST_ETHERNET_TAG   = 3
	EVPN_ETHERNET_SEGMENT_ROUTE             = 4
	EVPN_IP_PREFIX                          = 5
)

type EvpnRouteTypeInterface interface {
	DecodeFromBytes([]byte) error
}

type EvpnNlri struct {
	RouteType     uint8
	Length        uint8
	RouteTypeData EvpnRouteTypeInterface
}

func (n *EvpnNlri) DecodeFromBytes(data []byte) error {
	n.RouteType = data[0]
	n.Length = data[1]
	data = data[2:]

	var rdata EvpnRouteTypeInterface
	switch n.RouteType {
	case EVPN_ROUTE_TYPE_ETHERNET_AUTO_DISCOVERY:
		rdata = &EvpnEthernetAutoDiscoveryRoute{}
	case EVPN_ROUTE_TYPE_MAC_IP_ADVERTISEMENT:
	case EVPN_INCLUSIVE_MULTICAST_ETHERNET_TAG:
	case EVPN_ETHERNET_SEGMENT_ROUTE:
	case EVPN_IP_PREFIX:
	default:
	}
	rdata.DecodeFromBytes(data)
	return nil
}

func (n *EvpnNlri) Serialize() ([]byte, error) {
	return nil, nil
}

type RouteDistinguisherInterface interface {
	DecodeFromBytes([]byte) error
	Serialize() ([]byte, error)
	Len() int
	String() string
	MarshalJSON() ([]byte, error)
}

type EsiType uint8

const (
	ESI_ARBITRARY EsiType = iota
	ESI_LACP
	ESI_MSTP
	ESI_MAC
	ESI_ROUTERID
	ESI_AS
)

type EthernetSegmentIdentifier struct {
	Type  EsiType
	Value []byte
}

type EvpnEthernetAutoDiscoveryRoute struct {
	Rd    RouteDistinguisherInterface
	Esi   EthernetSegmentIdentifier
	Etag  uint32
	Label uint32
}

func (e *EvpnEthernetAutoDiscoveryRoute) DecodeFromBytes([]byte) error {
	return nil
}
