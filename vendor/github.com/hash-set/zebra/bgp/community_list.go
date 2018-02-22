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

// ip community-list standard WORD {permit, deny} 100:1 100:2 no-export
// ip community-list expanded WORD {permit, deny} _100:1_

import (
	"fmt"

	"github.com/hash-set/zebra/policy"
)

type CommunityListType int

const (
	CommunityListStandard CommunityListType = iota
	CommunityListExpanded
)

type CommunityListMap map[string]*CommunityList

type CommunityList struct {
	Type  CommunityListType
	Name  string
	Entry []*CommunityListEntry
}

type CommunityListEntry struct {
	Type      CommunityListType
	Policy    policy.Type
	Community Community
	Regstr    string
}

func (typ CommunityListType) String() string {
	switch typ {
	case CommunityListStandard:
		return "standard"
	case CommunityListExpanded:
		return "expanded"
	default:
		return "unknown"
	}
}

func (list *CommunityList) String() string {
	str := ""
	for _, e := range list.Entry {
		str += fmt.Sprintf("ip community-list %s %s %s ",
			list.Name, list.Type.String(), e.Policy.String())
		switch e.Type {
		case CommunityListStandard:
			str += e.Community.String()
		case CommunityListExpanded:
			str += e.Regstr
		}
		str += "\n"
	}
	return str
}

func NewCommunityListMap() CommunityListMap {
	return CommunityListMap{}
}

func (clist CommunityListMap) CommunityListGet(name string) *CommunityList {
	if list, ok := clist[name]; ok {
		return list
	}
	list := &CommunityList{Name: name}
	clist[name] = list
	return list
}

func (lhs *CommunityListEntry) Equal(rhs *CommunityListEntry) bool {
	if lhs.Policy != rhs.Policy {
		return false
	}
	if lhs.Type != rhs.Type {
		return false
	}
	switch lhs.Type {
	case CommunityListStandard:
		if !lhs.Community.Equal(rhs.Community) {
			return false
		}
	case CommunityListExpanded:
		if lhs.Regstr != rhs.Regstr {
			return false
		}
	}
	return true
}

func NewCommunityListEntry(typ CommunityListType, policy policy.Type, str string) (*CommunityListEntry, error) {
	entry := &CommunityListEntry{
		Type:   typ,
		Policy: policy,
	}
	switch typ {
	case CommunityListStandard:
		com, err := CommunityParse(str)
		if err != nil {
			return nil, err
		}
		entry.Community = com
	case CommunityListExpanded:
		entry.Regstr = str
	}
	return entry, nil
}

func (clist CommunityListMap) CommunityListAdd(name string, typ CommunityListType, policy policy.Type, str string) error {
	list := clist.CommunityListGet(name)
	if list.Type != typ {
		return fmt.Errorf("community-list type mismatch")
	}
	entry, err := NewCommunityListEntry(typ, policy, str)
	if err != nil {
		return err
	}
	for _, e := range list.Entry {
		if e.Equal(entry) {
			return nil
		}
	}
	list.Entry = append(list.Entry, entry)
	return nil
}

func (clist CommunityListMap) CommunityListDelete(name string, typ CommunityListType, policy policy.Type, str string) error {
	list := clist[name]
	if list == nil {
		return fmt.Errorf("community-list name %s dose not exist", name)
	}
	if list.Type != typ {
		return fmt.Errorf("community-list type mismatch")
	}
	entry, err := NewCommunityListEntry(typ, policy, str)
	if err != nil {
		return err
	}
	elist := []*CommunityListEntry{}
	for _, e := range list.Entry {
		if !e.Equal(entry) {
			elist = append(elist, e)
		}
	}
	list.Entry = elist
	return nil
}
