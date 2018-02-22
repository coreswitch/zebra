package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/coreswitch/netutil"
)

var VrfId = 1
var SyncFlag bool
var CliCommand = `#! /bin/bash
source /etc/bash_completion.d/cli
show ip route vrf vrf%d
`

var CliCommand2 = `#! /bin/bash
source /etc/bash_completion.d/cli
show ip route
`

var RibMap = map[string]string{}
var FibMap = map[string]string{}

func RibDump() error {
	cmd := fmt.Sprintf(CliCommand, VrfId)
	ribdump := "/tmp/ribdump.sh"

	ioutil.WriteFile(ribdump, []byte(cmd), os.ModePerm)

	err := os.Chmod(ribdump, 0777)
	if err != nil {
		return err
	}

	out, err := exec.Command(ribdump).Output()
	if err != nil {
		fmt.Println(err)
		return err
	}

	lines := strings.Split(string(out), "\n")
	//r := regexp.MustCompile(`^S\s`)
	r := regexp.MustCompile(`^B\s`)

	for _, line := range lines {
		// Filter non BGP routes
		if r.MatchString(line) {
			// fmt.Println(line)
			col := strings.Fields(line)
			if len(col) < 5 {
				continue
			}
			RibMap[col[1]] = col[4]
		}
	}

	// Dump Rib
	fmt.Println("Num of RIB:", len(RibMap))
	// fmt.Println("-----------")
	// for key, value := range RibMap {
	// 	fmt.Println(key, value)
	// }
	// fmt.Println("-----------")

	err = os.Remove(ribdump)
	if err != nil {
		return err
	}

	return nil
}

func FibDump() error {
	cmd := []string{"route", "show", "table", strconv.Itoa(VrfId)}

	out, err := exec.Command("ip", cmd...).Output()
	if err != nil {
		fmt.Println(err)
		return err
	}

	lines := strings.Split(string(out), "\n")
	r := regexp.MustCompile(`zebra`)

	for _, line := range lines {
		var prefix *netutil.Prefix
		// Filter non zebra routes.
		if !r.MatchString(line) {
			continue
		}

		// At least 3 fields exists.
		col := strings.Fields(line)
		if len(col) < 3 {
			continue
		}

		// Prefix parse.
		if col[0] == "default" {
			prefix = netutil.NewPrefixAFI(netutil.AFI_IP)
		} else {
			// /32 prefix.
			str := col[0]
			i := strings.IndexByte(str, '/')
			if i < 0 {
				str += "/32"
			}
			// Parse prefix.
			prefix, _ = netutil.ParsePrefix(str)
			if prefix == nil {
				continue
			}
			if prefix.AFI() != netutil.AFI_IP {
				continue
			}
		}
		if prefix == nil {
			continue
		}

		// via
		if col[1] != "via" {
			continue
		}

		FibMap[prefix.String()] = col[2]
	}

	// Dump FIB
	fmt.Println("Num of FIB:", len(FibMap))
	// fmt.Println("-----------")
	// for key, value := range FibMap {
	// 	fmt.Println(key, value)
	// }
	// fmt.Println("-----------")

	return nil
}

func main() {
	// --sync option.
	flag.BoolVar(&SyncFlag, "sync", false, "Sync FIB")

	flag.Parse()

	// "show ip route vrf VRF"
	RibDump()

	// "ip route show table ID"
	FibDump()

	// Compare
	for key, value := range RibMap {
		_, ok := FibMap[key]
		if !ok {
			fmt.Println("Missing FIB", key, value)
			if SyncFlag {
				args := []string{"route", "add", key, "via", value, "proto", "zebra", "table", strconv.Itoa(VrfId)}
				err := exec.Command("ip", args...).Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Sync FIB", key, value)
				}
			}
		}
	}
}
