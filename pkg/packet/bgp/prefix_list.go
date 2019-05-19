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
	"github.com/coreswitch/netutil"
)

type PrefixListEntry struct {
	Seq    int             `json:"seq"`
	Eq     int             `json:"eq,omitempty"`
	Le     int             `json:"le,omitempty"`
	Ge     int             `json:"ge,omitempty"`
	Prefix *netutil.Prefix `json:"prefix"`
}

type PrefixList struct {
	Description string             `json:"description,omitempty"`
	Entries     []*PrefixListEntry `json:"prefix-list"`
}

type PrefixListMaster struct {
	PrefixLists map[string]*PrefixList `json:"prefix-list"`
}

func NewPrefixList() *PrefixList {
	return &PrefixList{}
}

func NewPrefixListMaster() *PrefixListMaster {
	return &PrefixListMaster{
		PrefixLists: map[string]*PrefixList{},
	}
}

func (m *PrefixListMaster) PrefixListGet(name string) *PrefixList {
	plist := m.PrefixLists[name]
	if plist == nil {
		plist = NewPrefixList()
		m.PrefixLists[name] = plist
	}
	return plist
}

func (m *PrefixListMaster) PrefixListLookup(name string) *PrefixList {
	return m.PrefixLists[name]
}

func (m *PrefixListMaster) PrefixListDelete(name string) {
	delete(m.PrefixLists, name)
}

func (plist *PrefixList) Add(entry *PrefixListEntry) {
	entries := []*PrefixListEntry{}
	added := false
	addOnce := func(entry *PrefixListEntry) {
		if !added {
			entries = append(entries, entry)
			added = true
		}
	}
	for _, e := range plist.Entries {
		switch {
		case e.Seq < entry.Seq:
			entries = append(entries, e)
		case e.Seq == entry.Seq:
			addOnce(entry)
		case e.Seq > entry.Seq:
			addOnce(entry)
			entries = append(entries, e)
		}
	}
	addOnce(entry)
	plist.Entries = entries
}

func (plist *PrefixList) Delete(entry *PrefixListEntry) {
	entries := []*PrefixListEntry{}
	for _, e := range plist.Entries {
		if e.Seq != entry.Seq {
			entries = append(entries, e)
		}
	}
	plist.Entries = entries
}

func (plist *PrefixList) NewSeq() int {
	maxseq := 0
	for _, entry := range plist.Entries {
		if maxseq < entry.Seq {
			maxseq = entry.Seq
		}
	}
	return ((maxseq / 5) * 5) + 5
}

func (m *PrefixListMaster) EntryAdd(name string, p *netutil.Prefix, seq, eq, le, ge int) {
	plist := m.PrefixListGet(name)
	if seq == 0 {
		seq = plist.NewSeq()
	}
	entry := &PrefixListEntry{
		Seq:    seq,
		Eq:     eq,
		Le:     le,
		Ge:     ge,
		Prefix: p,
	}
	plist.Add(entry)
}

func (m *PrefixListMaster) EntryDelete(name string, p *netutil.Prefix, seq, eq, le, ge int) {
	plist := m.PrefixListLookup(name)
	if plist == nil {
		return
	}
	entry := &PrefixListEntry{
		Seq:    seq,
		Eq:     eq,
		Le:     le,
		Ge:     ge,
		Prefix: p,
	}
	plist.Delete(entry)
}

func (m *PrefixListMaster) DescriptionSet(name string, desc string) {
	plist := m.PrefixListLookup(name)
	if plist == nil {
		return
	}
	plist.Description = desc
}

func (m *PrefixListMaster) DescriptionUnset(name string, desc string) {
	plist := m.PrefixListLookup(name)
	if plist == nil {
		return
	}
	plist.Description = ""
}
