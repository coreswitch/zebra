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
	BGP_PORT = 179 // BGP default port
)

const (
	DEFAULT_HOLDTIME = 90
)

const (
	GR_RESTART_TIME    = 120
	GR_STALE_PATH_TIME = 360
)

const (
	AS_UNSPEC = 0     // Unspecifed AS number will be used in config.
	AS_TRANS  = 23456 // RFC6793 BGP Support for Four-Octet AS Number Space
)

type (
	Afi     uint16
	Safi    uint8
	AfiSafi uint32
)

const (
	AFI_IP     Afi = 1
	AFI_IP6        = 2
	AFI_L2VPN      = 25
	AFI_LS         = 16388
	AFI_OPAQUE     = 16397
)

const (
	SAFI_UNICAST            Safi = 1
	SAFI_MULTICAST               = 2
	SAFI_MPLS_LABEL              = 4
	SAFI_ENCAPSULATION           = 7
	SAFI_VPLS                    = 65
	SAFI_EVPN                    = 70
	SAFI_LS                      = 71
	SAFI_LS_VPN                  = 72
	SAFI_MPLS_VPN                = 128
	SAFI_MPLS_VPN_MULTICAST      = 129
	SAFI_RT_CONSTRTAINS          = 132
	SAFI_FLOW_SPEC_UNICAST       = 133
	SAFI_FLOW_SPEC_VPN           = 134
	SAFI_KEY_VALUE               = 241
)

var Afi2String = map[Afi]string{
	AFI_IP:     "ipv4",
	AFI_IP6:    "ipv6",
	AFI_L2VPN:  "l2vpn",
	AFI_LS:     "ls",
	AFI_OPAQUE: "opaque",
}

var Safi2String = map[Safi]string{
	SAFI_UNICAST:            "unicast",
	SAFI_MULTICAST:          "multicast",
	SAFI_MPLS_LABEL:         "label",
	SAFI_ENCAPSULATION:      "encap",
	SAFI_VPLS:               "vpls",
	SAFI_EVPN:               "evpn",
	SAFI_MPLS_VPN:           "vpn",
	SAFI_MPLS_VPN_MULTICAST: "vpn-multicast",
	SAFI_RT_CONSTRTAINS:     "rt-constraints",
	SAFI_FLOW_SPEC_UNICAST:  "flowspec-unicast",
	SAFI_FLOW_SPEC_VPN:      "flowspec-vpn",
	SAFI_KEY_VALUE:          "key",
}

func AfiSafiValue(afi Afi, safi Safi) AfiSafi {
	return AfiSafi(uint32(afi)<<16 | uint32(safi))
}

func (v AfiSafi) Afi() Afi {
	return Afi(v >> 16)
}

func (v AfiSafi) Safi() Safi {
	return Safi(v & 0xff)
}

func (afisafi AfiSafi) String() string {
	afistring, ok := Afi2String[afisafi.Afi()]
	if !ok {
		afistring = "unknown"
	}
	safistring, ok := Safi2String[afisafi.Safi()]
	if !ok {
		safistring = "unknown"
	}
	return afistring + "-" + safistring
}

// Notification error code  rfc 4271 4.5.
const (
	_ uint8 = iota
	BGP_ERR_MSG_HEADER_ERROR
	BGP_ERR_OPEN_MESSAGE_ERROR
	BGP_ERR_UPDATE_MESSAGE_ERROR
	BGP_ERR_HOLD_TIMER_EXPIRED
	BGP_ERR_FSM_ERROR
	BGP_ERR_CEASE
	BGP_ERR_ROUTE_REFRESH_MESSAGE_ERROR
)

// Notification Error Subcode for BGP_ERR_MESSAGE_HEADER_ERROR
const (
	_ uint8 = iota
	BGP_ERR_SUB_CONNECTION_NOT_SYNCHRONIZED
	BGP_ERR_SUB_BAD_MESSAGE_LENGTH
	BGP_ERR_SUB_BAD_MESSAGE_TYPE
)

// Notification Error Subcode for BGP_ERR_OPEN_MESSAGE_ERROR
const (
	_ uint8 = iota
	BGP_ERR_SUB_UNSUPPORTED_VERSION_NUMBER
	BGP_ERR_SUB_BAD_PEER_AS
	BGP_ERR_SUB_BAD_BGP_IDENTIFIER
	BGP_ERR_SUB_UNSUPPORTED_OPTIONAL_PARAMETER
	BGP_ERR_SUB_DEPRECATED_AUTHENTICATION_FAILURE
	BGP_ERR_SUB_UNACCEPTABLE_HOLD_TIME
	BGP_ERR_SUB_UNSUPPORTED_CAPABILITY
)

// Notification Error Subcode for BGP_ERR_UPDATE_MESSAGE_ERROR
const (
	_ uint8 = iota
	BGP_ERR_SUB_MALFORMED_ATTRIBUTE_LIST
	BGP_ERR_SUB_UNRECOGNIZED_WELL_KNOWN_ATTRIBUTE
	BGP_ERR_SUB_MISSING_WELL_KNOWN_ATTRIBUTE
	BGP_ERR_SUB_ATTRIBUTE_FLAGS_ERROR
	BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR
	BGP_ERR_SUB_INVALID_ORIGIN_ATTRIBUTE
	BGP_ERR_SUB_DEPRECATED_ROUTING_LOOP
	BGP_ERR_SUB_INVALID_NEXT_HOP_ATTRIBUTE
	BGP_ERR_SUB_OPTIONAL_ATTRIBUTE_ERROR
	BGP_ERR_SUB_INVALID_NETWORK_FIELD
	BGP_ERR_SUB_MALFORMED_AS_PATH
)

// Notification Error Subcode for BGP_ERR_HOLD_TIMER_EXPIRED
const (
	_ uint8 = iota
	BGP_ERR_SUB_HOLD_TIMER_EXPIRED
)

// Notification Error Subcode for BGP_ERR_FSM_ERROR
const (
	_ uint8 = iota
	BGP_ERR_SUB_RECEIVE_UNEXPECTED_MESSAGE_IN_OPENSENT_STATE
	BGP_ERR_SUB_RECEIVE_UNEXPECTED_MESSAGE_IN_OPENCONFIRM_STATE
	BGP_ERR_SUB_RECEIVE_UNEXPECTED_MESSAGE_IN_ESTABLISHED_STATE
)

// Notification Error Subcode for BGP_ERR_CEASE (RFC4486)
const (
	_ uint8 = iota
	BGP_ERR_SUB_MAXIMUM_NUMBER_OF_PREFIXES_REACHED
	BGP_ERR_SUB_ADMINISTRATIVE_SHUTDOWN
	BGP_ERR_SUB_PEER_DECONFIGURED
	BGP_ERR_SUB_ADMINISTRATIVE_RESET
	BGP_ERR_SUB_CONNECTION_REJECTED
	BGP_ERR_SUB_OTHER_CONFIGURATION_CHANGE
	BGP_ERR_SUB_CONNECTION_COLLISION_RESOLUTION
	BGP_ERR_SUB_OUT_OF_RESOURCES
	BGP_ERR_SUB_HARD_RESET //draft-ietf-idr-bgp-gr-notification-07
)

// Notification Error Subcode for BGP_ERR_ROUTE_REFRESH
const (
	_ uint8 = iota
	BGP_ERR_SUB_INVALID_MESSAGE_LENGTH
)
