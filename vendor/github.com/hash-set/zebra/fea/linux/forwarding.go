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

package linux

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

const (
	FORWARDING_IPV4_PATH = "/proc/sys/net/ipv4/ip_forward"
	FORWARDING_IPV6_PATH = "/proc/sys/net/ipv6/conf/all/forwarding"
)

func forwardingGet(path string) bool {
	fp, err := os.Open(path)
	if err != nil {
		return false
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	if scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return false
		}
		if val == 0 {
			return false
		} else {
			return true
		}
	}

	return false
}

func forwardingSet(path string, value bool) {
	if value {
		ioutil.WriteFile(path, []byte("1"), 0644)
	} else {
		ioutil.WriteFile(path, []byte("0"), 0644)
	}
}

func ForwardingIPv4Set(value bool) {
	forwardingSet(FORWARDING_IPV4_PATH, value)
}

func ForwardingIPv4Get() bool {
	return forwardingGet(FORWARDING_IPV4_PATH)
}

func ForwardingIPv6Set(value bool) {
	forwardingSet(FORWARDING_IPV6_PATH, value)
}

func ForwardingIPv6Get() bool {
	return forwardingGet(FORWARDING_IPV6_PATH)
}
