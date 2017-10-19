// Copyright 2016 OpenConfigd Project.
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
	"bytes"
	"os"
	"os/exec"
)

func Compare() string {
	configActive.WriteTo("/tmp/config.1")
	configCandidate.WriteTo("/tmp/config.2")

	var config string
	out, err := exec.Command("diff", "-U", "-1", "/tmp/config.1", "/tmp/config.2").Output()
	if err != nil {
		lines := bytes.Split(out, []byte{'\n'})
		if len(lines) > 3 {
			lines = lines[3:]
		}
		config = ""
		for _, s := range lines {
			config = config + string(s) + "\n"
		}
	} else {
		config = configCandidate.String()
	}

	os.Remove("/tmp/config.1")
	os.Remove("/tmp/config.2")

	return config
}

func JsonMarshal() string {
	return configCandidate.JsonMarshal()
}

func CompareCommand() string {
	configActive.WriteCommandTo("/tmp/config.1")
	configCandidate.WriteCommandTo("/tmp/config.2")

	var config string
	out, err := exec.Command("diff", "-U", "-1", "/tmp/config.1", "/tmp/config.2").Output()
	if err != nil {
		lines := bytes.Split(out, []byte{'\n'})
		if len(lines) > 3 {
			lines = lines[3:]
		}
		config = ""
		for _, s := range lines {
			config = config + string(s) + "\n"
		}
	}

	os.Remove("/tmp/config.1")
	os.Remove("/tmp/config.2")

	return config
}

func Commands() string {
	return configActive.CommandString()
}
