// Copyright 2017 coreswitch.
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

package component

import (
	"reflect"
	"testing"
)

type Record struct {
	entry []string
}

type BaseComponent struct {
	record *Record
}

func (this *BaseComponent) Start() Component {
	this.record.entry = append(this.record.entry, "BaseComponent Start")
	return this
}

func (this *BaseComponent) Stop() Component {
	this.record.entry = append(this.record.entry, "BaseComponent Stop")
	return this
}

type FirstComponent struct {
	record *Record
}

func (this *FirstComponent) Start() Component {
	this.record.entry = append(this.record.entry, "FirstComponent Start")
	return this
}

func (this *FirstComponent) Stop() Component {
	this.record.entry = append(this.record.entry, "FirstComponent Stop")
	return this
}

type SecondComponent struct {
	record *Record
}

func (this *SecondComponent) Start() Component {
	this.record.entry = append(this.record.entry, "SecondComponent Start")
	return this
}

func (this *SecondComponent) Stop() Component {
	this.record.entry = append(this.record.entry, "SecondComponent Stop")
	return this
}

type ThirdComponent struct {
	record *Record
}

func (this *ThirdComponent) Start() Component {
	this.record.entry = append(this.record.entry, "ThirdComponent Start")
	return this
}

func (this *ThirdComponent) Stop() Component {
	this.record.entry = append(this.record.entry, "ThirdComponent Stop")
	return this
}

//
// <BaseComponent>
//        |
// <FirstComponent>
//
func TestSystemMapSimple(t *testing.T) {
	record := &Record{}
	baseComponent := &BaseComponent{
		record: record,
	}
	firstComponent := &FirstComponent{
		record: record,
	}

	systemMap := SystemMap{
		"base":  baseComponent,
		"first": ComponentWith(firstComponent, "base"),
	}

	systemMap.Start()
	startRecord := &Record{
		entry: []string{
			"BaseComponent Start",
			"FirstComponent Start",
		},
	}
	if !reflect.DeepEqual(record, startRecord) {
		t.Error("Component start order error", record, "must be", startRecord)
	}

	systemMap.Stop()
	stopRecord := &Record{
		entry: []string{
			"BaseComponent Start",
			"FirstComponent Start",
			"FirstComponent Stop",
			"BaseComponent Stop",
		},
	}
	if !reflect.DeepEqual(record, stopRecord) {
		t.Error("Component stop order error", record, "must be", stopRecord)
	}
}

//
//          <BaseComponent>
//           /         |
//          /          |
// <FirstComponent>    |
//          \          |
//           \         |
//         <SecondComponent>
//
func TestSystemMapTriangle(t *testing.T) {
	record := &Record{}
	baseComponent := &BaseComponent{
		record: record,
	}
	firstComponent := &FirstComponent{
		record: record,
	}
	secondComponent := &SecondComponent{
		record: record,
	}

	systemMap := SystemMap{
		"base":   baseComponent,
		"first":  ComponentWith(firstComponent, "base"),
		"second": ComponentWith(secondComponent, "base", "first"),
	}

	systemMap.Start()
	startRecord := &Record{
		entry: []string{
			"BaseComponent Start",
			"FirstComponent Start",
			"SecondComponent Start",
		},
	}
	if !reflect.DeepEqual(record, startRecord) {
		t.Error("Component start order error", record, "must be", startRecord)
	}

	systemMap.Stop()
	stopRecord := &Record{
		entry: []string{
			"BaseComponent Start",
			"FirstComponent Start",
			"SecondComponent Start",
			"SecondComponent Stop",
			"FirstComponent Stop",
			"BaseComponent Stop",
		},
	}
	if !reflect.DeepEqual(record, stopRecord) {
		t.Error("Component stop order error", record, "must be", stopRecord)
	}
}

//
//          <BaseComponent>
//           /         \
//          /           \
// <FirstComponent> <SecondComponent>
//          \           /
//           \         /
//         <ThirdComponent>
//
func TestSystemMapSquare(t *testing.T) {
	record := &Record{}
	baseComponent := &BaseComponent{
		record: record,
	}
	firstComponent := &FirstComponent{
		record: record,
	}
	secondComponent := &SecondComponent{
		record: record,
	}
	thirdComponent := &ThirdComponent{
		record: record,
	}

	systemMap := SystemMap{
		"base":   baseComponent,
		"first":  ComponentWith(firstComponent, "base"),
		"second": ComponentWith(secondComponent, "base"),
		"third":  ComponentWith(thirdComponent, "first", "second"),
	}

	systemMap.Start()
	startRecord1 := &Record{
		entry: []string{
			"BaseComponent Start",
			"FirstComponent Start",
			"SecondComponent Start",
			"ThirdComponent Start",
		},
	}
	startRecord2 := &Record{
		entry: []string{
			"BaseComponent Start",
			"SecondComponent Start",
			"FirstComponent Start",
			"ThirdComponent Start",
		},
	}
	if !reflect.DeepEqual(record, startRecord1) && !reflect.DeepEqual(record, startRecord2) {
		t.Error("Component start order error", record, "must be", startRecord1, "or", startRecord2)
	}

	record.entry = record.entry[:0]
	systemMap.Stop()
	stopRecord1 := &Record{
		entry: []string{
			"ThirdComponent Stop",
			"SecondComponent Stop",
			"FirstComponent Stop",
			"BaseComponent Stop",
		},
	}
	stopRecord2 := &Record{
		entry: []string{
			"ThirdComponent Stop",
			"FirstComponent Stop",
			"SecondComponent Stop",
			"BaseComponent Stop",
		},
	}
	if !reflect.DeepEqual(record, stopRecord1) && !reflect.DeepEqual(record, stopRecord2) {
		t.Error("Component stop order error", record, "must be", stopRecord1, "or", stopRecord2)
	}
}
