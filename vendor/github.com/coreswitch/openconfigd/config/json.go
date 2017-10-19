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
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

func JsonFlatten(paths []string, v reflect.Value, f func([]string)) {
	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			JsonFlatten(append(paths, key.Interface().(string)), v.MapIndex(key), f)
		}
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			JsonFlatten(paths, v.Index(i), f)
		}
	case reflect.Interface:
		JsonFlatten(paths, v.Elem(), f)
	case reflect.String:
		f(append(paths, v.String()))
	case reflect.Bool:
		f(append(paths, strconv.FormatBool(v.Bool())))
	case reflect.Float64:
		f(append(paths, strconv.FormatFloat(v.Float(), 'f', -1, 64)))
	default:
	}
}

func JsonParse(jsonString string) ([]string, error) {
	var jsonIntf interface{}

	err := json.Unmarshal([]byte(jsonString), &jsonIntf)
	if err != nil {
		return nil, err
	}
	JsonFlatten(nil, reflect.ValueOf(jsonIntf),
		func(path []string) {
			fmt.Println("Path:", path)
		},
	)
	fmt.Println()
	return nil, nil
}

var jsonStr2 = `
{
    "dhcp": {
        "server": {
            "default-lease-time": 600,
            "dhcp-ip-pool": [
                {
                    "default-lease-time": 456,
                    "gateway-ip": "192.168.10.1",
                    "host": [
                        {
                            "host-name": "h0",
                            "ip-address": "192.168.10.23",
                            "mac-address": "00:1c:42:83:e5:ac"
                        }
                    ],
                    "interface": "lan-1",
                    "ip-pool-name": "904cd99a-f447-4bc0-ac6c-d151606be5bd",
                    "max-lease-time": 34567,
                    "option": {
                        "domain-name": "ntti3.com",
                        "domain-name-servers": [
                            {
                                "server": "8.8.8.8"
                            },
                            {
                                "server": "4.4.8.8"
                            }
                        ],
                        "ntp-servers": [
                            {
                                "server": "192.168.10.2"
                            }
                        ]
                    },
                    "range": [
                        {
                            "range-end-ip": "192.168.10.200",
                            "range-index": 1,
                            "range-start-ip": "192.168.10.100"
                        }
                    ],
                    "subnet": "192.168.10.0/24"
                }
            ],
            "max-lease-time": 7200
        }
    }
}
`

var jsonStr = `
{
    "dhcp": {
        "server": {
            "default-lease-time": 600,
            "max-lease-time": 7200
        }
    }
}
`

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'g', -1, 64)
	default:
		return v.Type().String() + " value"
	}
}

func parse(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			parse(fmt.Sprintf("%s[%s]", path, key), v.MapIndex(key))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			parse(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			switch v.Elem().Kind() {
			case reflect.String, reflect.Float64, reflect.Bool:
				fmt.Printf("%s.value = %s (%s)\n", path, formatAtom(v.Elem()), v.Elem().Type())
			default:
				// fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
				parse(path, v.Elem())
			}
		}
	default:
		fmt.Printf("%s\n", formatAtom(v))
	}
}

// Yang based json parser.
func YangJsonParse() {
	fmt.Println("YangJsonParse")

	var jsonIntf interface{}
	err := json.Unmarshal([]byte(jsonStr2), &jsonIntf)
	if err != nil {
		fmt.Println("json error:", err)
		return
	}
	parse("", reflect.ValueOf(jsonIntf))
}
