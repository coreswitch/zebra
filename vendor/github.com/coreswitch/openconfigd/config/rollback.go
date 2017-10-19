// Copyright 2017 OpenConfigd Project.
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

package config

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
)

var ConfigDir = "/usr/local/etc"
var ConfigRevisionMax = 50

type Rollback struct {
	Rev    int
	Detail string
}

type RollbackList []*Rollback

func (rlist RollbackList) Len() int {
	return len(rlist)
}

func (rlist RollbackList) Less(i, j int) bool {
	return rlist[i].Rev < rlist[j].Rev
}

func (rlist RollbackList) Swap(i, j int) {
	rlist[i], rlist[j] = rlist[j], rlist[i]
}

func RollbackDetailFetch(fileName string) string {
	fp, err := os.Open(ConfigDir + "/" + fileName)
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(fp)
	scanner.Scan()
	detail := scanner.Text()
	if len(detail) >= 2 && detail[0] == '#' {
		detail = detail[2:]
	} else {
		detail = ""
	}
	return detail
}

func RollbackListGet() ([]*Rollback, error) {
	rlist := RollbackList{}

	files, err := ioutil.ReadDir(ConfigDir)
	if err != nil {
		return nil, err
	}

	// File name match to ConfigFileBasename.<digit>.
	r := regexp.MustCompile(configFileBasename + "\\.(\\d+)")
	for _, file := range files {
		matches := r.FindAllStringSubmatch(file.Name(), -1)
		if matches != nil && len(matches) > 0 {
			match := matches[0]
			if len(match) > 1 {
				rev, _ := strconv.Atoi(match[1])
				rollback := &Rollback{
					Rev:    rev,
					Detail: RollbackDetailFetch(file.Name()),
				}
				rlist = append(rlist, rollback)
			}
		}
	}

	sort.Sort(rlist)

	return rlist, nil
}

func RollbackRevisionIncrement() error {
	err := os.Chdir(ConfigDir)
	if err != nil {
		return err
	}
	for rev := ConfigRevisionMax - 1; rev > 0; rev-- {
		os.Rename(configFileBasename+"."+strconv.Itoa(rev-1), configFileBasename+"."+strconv.Itoa(rev))
	}
	return nil
}

func RollbackCompletion(commands []string) []string {
	rlist, err := RollbackListGet()
	if err != nil {
		return []string{"0", "1"}
	}
	//item := []string{}
	for rollback := range rlist {
		fmt.Println(rollback)
	}
	return []string{"0", "1"}
}
