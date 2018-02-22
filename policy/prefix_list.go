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

package policy

import (
	"fmt"

	"github.com/coreswitch/netutil"
)

type PrefixListEntry struct {
	Seq    int             `json:"seq"`
	Action Action          `json:"action"`
	Prefix *netutil.Prefix `json:"prefix"`
	Le     *int            `json:"le,omitempty"`
	Ge     *int            `json:"ge,omitempty"`
	Eq     *int            `json:"eq,omitempty"`
}

type PrefixListEntrySlice []*PrefixListEntry

type PrefixList struct {
	Name        string               `json:"name"`
	Description string               `json:"description,omitempty"`
	Entries     PrefixListEntrySlice `json:"prefix-list-entries"`
}

type PrefixListMaster struct {
	PrefixLists map[string]*PrefixList `json:"prefix-lists"`
}

func NewPrefixList(name string) *PrefixList {
	return &PrefixList{
		Name: name,
	}
}

func NewPrefixListMaster() *PrefixListMaster {
	return &PrefixListMaster{
		PrefixLists: map[string]*PrefixList{},
	}
}

func (m *PrefixListMaster) Get(name string) *PrefixList {
	plist := m.Lookup(name)
	if plist != nil {
		return plist
	}
	plist = NewPrefixList(name)
	m.PrefixLists[name] = plist
	return plist
}

func (m *PrefixListMaster) Lookup(name string) *PrefixList {
	return m.PrefixLists[name]
}

func (m *PrefixListMaster) Delete(name string) {
	delete(m.PrefixLists, name)
}

func (entries PrefixListEntrySlice) Add(entry *PrefixListEntry) PrefixListEntrySlice {
	var i int
	for i = 0; i < len(entries); i++ {
		v := entries[i]
		if entry.Seq < v.Seq {
			break
		}
		if entry.Seq == v.Seq {
			entries[i] = entry
			return entries
		}
	}
	if i < len(entries) {
		if i == 0 {
			return append(PrefixListEntrySlice{entry}, entries...)
		} else {
			return append(entries[:i], append(PrefixListEntrySlice{entry}, entries[i:]...)...)
		}
	} else {
		return append(entries, entry)
	}
}

func (plist *PrefixList) EntryAdd(entry *PrefixListEntry) {
	plist.Entries = plist.Entries.Add(entry)
}

func (plist *PrefixList) AddWithClosure(entry *PrefixListEntry) {
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

func (plist *PrefixList) EntryDelete(entry *PrefixListEntry) {
	for i, e := range plist.Entries {
		if e.Seq == entry.Seq {
			plist.Entries = append(plist.Entries[:i], plist.Entries[i+1:]...)
			break
		}
	}
}

func (entry *PrefixListEntry) Validate() error {
	if entry.Eq != nil {
		// When 'eq' is specified neither 'le' nor 'ge' can be specified.
		if entry.Le != nil || entry.Ge != nil {
			return fmt.Errorf("Can't specify eq-value with le-value or ge-value")
		}
		// 'eq' must be bigger than prefix length.
		if *entry.Eq <= entry.Prefix.Length {
			return fmt.Errorf("Invalid mask length entered: Make sure len < eq-value")
		}
		return nil
	}
	// Following is check for the rule len < ge-value <= le-value.
	if entry.Ge != nil {
		if *entry.Ge <= entry.Prefix.Length {
			return fmt.Errorf("Invalid mask length entered: Make sure len < ge-value <= le-value")
		}
	}
	if entry.Le != nil {
		if *entry.Le <= entry.Prefix.Length {
			return fmt.Errorf("Invalid mask length entered: Make sure len < ge-value <= le-value")
		}
	}
	if entry.Le != nil && entry.Ge != nil {
		if *entry.Le < *entry.Ge {
			return fmt.Errorf("Invalid mask length entered: Make sure len < ge-value <= le-value")
		}
	}
	return nil
}

func NewPrefixListEntry(seq int, action Action, p *netutil.Prefix, opts ...PrefixListOption) *PrefixListEntry {
	if seq == 0 {
		return nil
	}
	entry := &PrefixListEntry{
		Seq:    seq,
		Action: action,
		Prefix: p,
	}
	for _, opt := range opts {
		opt(entry)
	}
	return entry
}

func (m *PrefixListMaster) EntryAdd(name string, entry *PrefixListEntry) error {
	if entry == nil {
		return fmt.Errorf("Empty entry addition")
	}
	if err := entry.Validate(); err != nil {
		return err
	}
	plist := m.Get(name)
	plist.EntryAdd(entry)
	return nil
}

func (m *PrefixListMaster) EntryDelete(name string, entry *PrefixListEntry) {
	if entry == nil {
		return
	}
	plist := m.Lookup(name)
	if plist == nil {
		return
	}
	plist.EntryDelete(entry)
}

func (m *PrefixListMaster) DescriptionSet(name string, desc string) {
	plist := m.Lookup(name)
	if plist == nil {
		return
	}
	plist.Description = desc
}

func (m *PrefixListMaster) DescriptionUnset(name string, desc string) {
	plist := m.Lookup(name)
	if plist == nil {
		return
	}
	plist.Description = ""
}

type PrefixListOption func(*PrefixListEntry)

func WithGe(val int) PrefixListOption {
	return func(entry *PrefixListEntry) {
		entry.Ge = &val
	}
}

func WithLe(val int) PrefixListOption {
	return func(entry *PrefixListEntry) {
		entry.Le = &val
	}
}

func WithEq(val int) PrefixListOption {
	return func(entry *PrefixListEntry) {
		entry.Eq = &val
	}
}

func (entry *PrefixListEntry) Match(p *netutil.Prefix) bool {
	if !entry.Prefix.Match(p) {
		return false
	}
	if entry.Eq == nil && entry.Le == nil && entry.Ge == nil {
		if entry.Prefix.Length != p.Length {
			return false
		} else {
			return true
		}
	}
	if entry.Eq != nil {
		if p.Length != *entry.Eq {
			return false
		}
	}
	if entry.Le != nil {
		if p.Length > *entry.Le {
			return false
		}
	}
	if entry.Ge != nil {
		if p.Length < *entry.Ge {
			return false
		}
	}
	return true
}

func (plist *PrefixList) Match(p *netutil.Prefix) Action {
	for _, entry := range plist.Entries {
		if entry.Match(p) {
			return entry.Action
		}
	}
	return Deny
}
