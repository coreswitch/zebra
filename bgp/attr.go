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

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"strings"

	"github.com/coreswitch/zebra/pkg/packet/bgp"
)

type BGPAttrFlag uint8

const (
	BGP_ATTR_FLAG_EXTENDED_LENGTH BGPAttrFlag = 1 << 4
	BGP_ATTR_FLAG_PARTIAL         BGPAttrFlag = 1 << 5
	BGP_ATTR_FLAG_TRANSITIVE      BGPAttrFlag = 1 << 6
	BGP_ATTR_FLAG_OPTIONAL        BGPAttrFlag = 1 << 7
)

func (f BGPAttrFlag) IsExtendedLength() bool {
	return (f & BGP_ATTR_FLAG_EXTENDED_LENGTH) != 0
}

func (f BGPAttrFlag) IsPartial() bool {
	return (f & BGP_ATTR_FLAG_PARTIAL) != 0
}

func (f BGPAttrFlag) IsTransitive() bool {
	return (f & BGP_ATTR_FLAG_TRANSITIVE) != 0
}

func (f BGPAttrFlag) IsOptional() bool {
	return (f & BGP_ATTR_FLAG_OPTIONAL) != 0
}

func (f BGPAttrFlag) Info() string {
	info := []string{}
	if f.IsExtendedLength() {
		info = append(info, "extended_length")
	}
	if f.IsPartial() {
		info = append(info, "partial")
	}
	if f.IsTransitive() {
		info = append(info, "transitive")
	}
	if f.IsOptional() {
		info = append(info, "optional")
	}
	return strings.Join(info, ",")
}

type BGPAttrType uint8

const (
	_                              BGPAttrType = iota
	BGP_ATTR_TYPE_ORIGIN                       // 1
	BGP_ATTR_TYPE_AS_PATH                      // 2
	BGP_ATTR_TYPE_NEXT_HOP                     // 3
	BGP_ATTR_TYPE_MED                          // 4
	BGP_ATTR_TYPE_LOCAL_PREF                   // 5
	BGP_ATTR_TYPE_ATOMIC_AGGREGATE             // 6
	BGP_ATTR_TYPE_AGGREGATOR                   // 7
	BGP_ATTR_TYPE_COMMUNITIES                  // 8
	BGP_ATTR_TYPE_ORIGINATOR_ID                // 9
	BGP_ATTR_TYPE_CLUSTER_LIST                 // 10
	_
	_
	_
	BGP_ATTR_TYPE_MP_REACH_NLRI                    // 14
	BGP_ATTR_TYPE_MP_UNREACH_NLRI                  // 15
	BGP_ATTR_TYPE_EXTENDED_COMMUNITIES             // 16
	BGP_ATTR_TYPE_AS4_PATH                         // 17
	BGP_ATTR_TYPE_AS4_AGGREGATOR                   // 18
	BGP_ATTR_TYPE_PREFIX_SID           BGPAttrType = 40
)

var BGPAttrTypeString = map[BGPAttrType]string{
	BGP_ATTR_TYPE_ORIGIN:               "origin",
	BGP_ATTR_TYPE_AS_PATH:              "aspath",
	BGP_ATTR_TYPE_NEXT_HOP:             "nexthop",
	BGP_ATTR_TYPE_MED:                  "med",
	BGP_ATTR_TYPE_LOCAL_PREF:           "local preference",
	BGP_ATTR_TYPE_ATOMIC_AGGREGATE:     "atmic aggregate",
	BGP_ATTR_TYPE_AGGREGATOR:           "aggregator",
	BGP_ATTR_TYPE_COMMUNITIES:          "communities",
	BGP_ATTR_TYPE_ORIGINATOR_ID:        "originator id",
	BGP_ATTR_TYPE_CLUSTER_LIST:         "cluster list",
	BGP_ATTR_TYPE_MP_REACH_NLRI:        "mp reach nlri",
	BGP_ATTR_TYPE_MP_UNREACH_NLRI:      "mp unreach nlri",
	BGP_ATTR_TYPE_EXTENDED_COMMUNITIES: "ext communities",
	BGP_ATTR_TYPE_AS4_PATH:             "as4 path",
	BGP_ATTR_TYPE_AS4_AGGREGATOR:       "as4 aggregator",
	BGP_ATTR_TYPE_PREFIX_SID:           "prefix sid",
}

type AttrInterface interface {
	DecodeFromBytes([]byte) error
	Serialize() ([]byte, error)
	TotalLength() int
}

func NewAttrByType(typ byte) AttrInterface {
	switch BGPAttrType(typ) {
	case BGP_ATTR_TYPE_ORIGIN:
		return &AttrOrigin{}
	case BGP_ATTR_TYPE_AS_PATH:
		return &AttrAsPath{}
	case BGP_ATTR_TYPE_NEXT_HOP:
		return &AttrNexthop{}
	case BGP_ATTR_TYPE_MED:
		return &AttrMed{}
	case BGP_ATTR_TYPE_LOCAL_PREF:
		return &AttrLocalPref{}
	case BGP_ATTR_TYPE_ATOMIC_AGGREGATE:
		return &AttrAtomicAggregate{}
	case BGP_ATTR_TYPE_AGGREGATOR:
		return &AttrAggregator{}
	case BGP_ATTR_TYPE_COMMUNITIES:
		return &AttrCommunity{}
	case BGP_ATTR_TYPE_ORIGINATOR_ID:
		return &AttrOriginatorId{}
	case BGP_ATTR_TYPE_CLUSTER_LIST:
		return &AttrClusterList{}
		// case BGP_ATTR_TYPE_MP_REACH_NLRI:
		// case BGP_ATTR_TYPE_MP_UNREACH_NLRI:
		// case BGP_ATTR_TYPE_EXTENDED_COMMUNITIES:
		// case BGP_ATTR_TYPE_AS4_PATH:
		// case BGP_ATTR_TYPE_AS4_AGGREGATOR:
	case BGP_ATTR_TYPE_PREFIX_SID:
		return &AttrPrefixSid{}
	default:
		return &AttrUnknown{}
	}
}

type AttrBase struct {
	Flags  BGPAttrFlag
	Type   BGPAttrType
	Length uint16
	Value  []byte
}

func (attr *AttrBase) TotalLength() int {
	if attr.Flags.IsExtendedLength() {
		return int(attr.Length + 2 + 2)
	} else {
		return int(attr.Length + 2 + 1)
	}
}

func (attr *AttrBase) DecodeFromBytes(data []byte) error {
	code := BGP_ERR_UPDATE_MESSAGE_ERROR
	subCode := BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR
	if len(data) < 2 {
		return NewBgpError(code, subCode, nil, "")
	}
	attr.Flags = BGPAttrFlag(data[0])
	attr.Type = BGPAttrType(data[1])
	data = data[2:]
	fmt.Printf("Flag: %02X (%s)\n", attr.Flags, attr.Flags.Info())
	fmt.Println("Type:", BGPAttrTypeString[attr.Type])

	if attr.Flags.IsExtendedLength() {
		if len(data) < 2 {
			return NewBgpError(code, subCode, nil, "")
		}
		attr.Length = binary.BigEndian.Uint16(data[:2])
		data = data[2:]
	} else {
		if len(data) < 1 {
			return NewBgpError(code, subCode, nil, "")
		}
		attr.Length = uint16(data[0])
		data = data[1:]
	}

	if len(data) < int(attr.Length) {
		return NewBgpError(code, subCode, nil, "")
	}

	if attr.Length > 0 {
		attr.Value = data[:attr.Length]
	}

	return nil
}

func (attr *AttrBase) Serialize() ([]byte, error) {
	return nil, nil
}

// Origin.
type AttrOrigin struct {
	AttrBase
}

const (
	BGP_ORIGIN_IGP        = 0
	BGP_ORIGIN_EGP        = 1
	BGP_ORIGIN_INCOMPLETE = 2
)

func (attr *AttrOrigin) String() string {
	switch attr.Value[0] {
	case BGP_ORIGIN_IGP:
		return "i"
	case BGP_ORIGIN_EGP:
		return "e"
	case BGP_ORIGIN_INCOMPLETE:
		return "?"
	default:
		return "?"
	}
}

func (attr *AttrOrigin) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Origin string `json:"origin"`
	}{
		Origin: attr.String(),
	})
}

// AS path.
type AttrAsPath struct {
	AttrBase
	AsPath AsPath
}

func (attr *AttrAsPath) String() string {
	return attr.AsPath.String()
}

func (attr *AttrAsPath) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Origin string `json:"aspath"`
	}{
		Origin: attr.String(),
	})
}

func (attr *AttrAsPath) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}

	if attr.Value == nil {
		return nil
	}

	v := attr.Value
	remaining := len(v)
	for remaining > 0 {
		var seg AsSegmentInterface

		seg = &As4Segment{}
		err := seg.DecodeFromBytes(v)
		if err != nil {
			return err
		}
		remaining -= seg.EncodeLength()
		v = v[seg.EncodeLength():]

		attr.AsPath = append(attr.AsPath, seg)
	}
	return nil
}

// Nexthop.
type AttrNexthop struct {
	AttrBase
	Nexthop net.IP
}

func (attr *AttrNexthop) String() string {
	return attr.Nexthop.String()
}

func (attr *AttrNexthop) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nexthop string `json:"nexthop"`
	}{
		Nexthop: attr.String(),
	})
}

func (attr *AttrNexthop) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	if len(attr.Value) != net.IPv4len && len(attr.Value) != net.IPv6len {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	attr.Nexthop = attr.AttrBase.Value
	return nil
}

// MED.
type AttrMed struct {
	AttrBase
	Med uint32
}

func (attr *AttrMed) String() string {
	return fmt.Sprintf("%d", attr.Med)
}

func (attr *AttrMed) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Med uint32 `json:"med"`
	}{
		Med: attr.Med,
	})
}

func (attr *AttrMed) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	if len(attr.Value) != 4 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	attr.Med = binary.BigEndian.Uint32(data)

	return nil
}

// Local preference.
type AttrLocalPref struct {
	AttrBase
	LocalPref uint32
}

func (attr *AttrLocalPref) String() string {
	return fmt.Sprintf("%d", attr.LocalPref)
}

func (attr *AttrLocalPref) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		LocalPref uint32 `json:"local-preference"`
	}{
		LocalPref: attr.LocalPref,
	})
}

func (attr *AttrLocalPref) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	if len(attr.Value) != 4 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	attr.LocalPref = binary.BigEndian.Uint32(data)

	return nil
}

// Atomic Aggregate.
type AttrAtomicAggregate struct {
	AttrBase
}

func (attr *AttrAtomicAggregate) String() string {
	return ""
}

func (attr *AttrAtomicAggregate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Flag bool `json:"atomic-aggregate"`
	}{
		Flag: true,
	})
}

func (attr *AttrAtomicAggregate) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	if len(attr.Value) != 0 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}

	return nil
}

// Aggregator.
type AttrAggregator struct {
	AttrBase
	As      uint32
	Address net.IP
}

func (attr *AttrAggregator) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	switch len(attr.Value) {
	case 6:
		attr.As = uint32(binary.BigEndian.Uint16(attr.Value[0:2]))
		attr.Address = attr.Value[2:]
	case 8:
		attr.As = binary.BigEndian.Uint32(attr.Value[0:4])
		attr.Address = attr.Value[4:]
	default:
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	return nil
}

// Community.
type AttrCommunity struct {
	AttrBase
	Community bgp.Community
}

func (attr *AttrCommunity) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	if len(attr.Value)%4 != 0 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	v := attr.Value
	for len(v) > 0 {
		attr.Community = append(attr.Community, binary.BigEndian.Uint32(v))
		v = v[4:]
	}
	return nil
}

func (attr *AttrCommunity) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Community bgp.Community `json:"community"`
	}{
		Community: attr.Community,
	})
}

// Originator ID.
type AttrOriginatorId struct {
	AttrBase
	RouterId net.IP
}

func (attr *AttrOriginatorId) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	if len(attr.Value) != 4 && len(attr.Value) != 16 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	attr.RouterId = attr.Value
	return nil
}

// Cluster list.
type AttrClusterList struct {
	AttrBase
	ClusterList []net.IP
}

func (attr *AttrClusterList) DecodeFromBytes(data []byte) error {
	err := attr.AttrBase.DecodeFromBytes(data)
	if err != nil {
		return err
	}
	value := attr.Value
	if len(value)%4 != 0 {
		return NewBgpError(BGP_ERR_UPDATE_MESSAGE_ERROR, BGP_ERR_SUB_ATTRIBUTE_LENGTH_ERROR, nil, "")
	}
	for len(value) >= 4 {
		attr.ClusterList = append(attr.ClusterList, value[:4])
		value = value[4:]
	}
	return nil
}

// Unknown.
type AttrUnknown struct {
	AttrBase
}
