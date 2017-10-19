// Copyright 2016 Zebra Project
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

package rib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type IfStats struct {
	Rx_packets          uint32
	Rx_bytes            uint64
	Rx_errors           uint32
	Rx_dropped          uint32
	Rx_multicast        uint32
	Rx_compressed       uint32
	Rx_length_errors    uint32
	Rx_over_errors      uint32
	Rx_crc_errors       uint32
	Rx_frame_errors     uint32
	Rx_fifo_errors      uint32
	Rx_missed_errors    uint32
	Tx_packets          uint32
	Tx_bytes            uint64
	Tx_errors           uint32
	Tx_dropped          uint32
	Tx_compressed       uint32
	Tx_aborted_errors   uint32
	Tx_carrier_errors   uint32
	Tx_fifo_errors      uint32
	Tx_heartbeat_errors uint32
	Tx_window_errors    uint32
	Collisions          uint32
}

type IfStatsYang struct {
	date_and_time      uint64
	in_broadcast_pkts  uint64
	in_discards        uint32
	in_errors          uint32
	in_multicast_pkts  uint64
	in_octets          uint64
	in_unicast_pkts    uint64
	in_unknown_protos  uint32
	out_broadcast_pkts uint64
	out_discards       uint32
	out_errors         uint32
	out_multicast_pkts uint64
	out_octets         uint64
	out_unicast_pkts   uint64
}

func IfStatsScan(version int, buf string, stats *IfStats) {
	switch version {
	case 3:
		fmt.Sscanf(buf,
			"%d %d %d %d %d %d %d %d %d %d %d %d %d %d %d %d",
			&stats.Rx_bytes,
			&stats.Rx_packets,
			&stats.Rx_errors,
			&stats.Rx_dropped,
			&stats.Rx_fifo_errors,
			&stats.Rx_frame_errors,
			&stats.Rx_compressed,
			&stats.Rx_multicast,
			&stats.Tx_bytes,
			&stats.Tx_packets,
			&stats.Tx_errors,
			&stats.Tx_dropped,
			&stats.Tx_fifo_errors,
			&stats.Collisions,
			&stats.Tx_carrier_errors,
			&stats.Tx_compressed)
	case 2:
		fmt.Sscanf(buf, "%lld %d %d %d %d %d %lld %d %d %d %d %d %d",
			&stats.Rx_bytes,
			&stats.Rx_packets,
			&stats.Rx_errors,
			&stats.Rx_dropped,
			&stats.Rx_fifo_errors,
			&stats.Rx_frame_errors,
			&stats.Tx_bytes,
			&stats.Tx_packets,
			&stats.Tx_errors,
			&stats.Tx_dropped,
			&stats.Tx_fifo_errors,
			&stats.Collisions,
			&stats.Tx_carrier_errors)
		stats.Rx_multicast = 0
	case 1:
		fmt.Sscanf(buf, "%d %d %d %d %d %d %d %d %d %d %d",
			&stats.Rx_packets,
			&stats.Rx_errors,
			&stats.Rx_dropped,
			&stats.Rx_fifo_errors,
			&stats.Rx_frame_errors,
			&stats.Tx_packets,
			&stats.Tx_errors,
			&stats.Tx_dropped,
			&stats.Tx_fifo_errors,
			&stats.Collisions,
			&stats.Tx_carrier_errors)
		stats.Rx_bytes = 0
		stats.Tx_bytes = 0
		stats.Rx_multicast = 0
	}
}

func ifNameCut(str string) string {
	var i int
	for i = 0; i < len(str); i++ {
		if str[i] != ' ' {
			break
		}
	}
	str = str[i:]
	p := strings.IndexByte(str, ':')
	if p < 0 {
		return ""
	}
	return str[:p]
}

func IfStatsUpdate() {
	fp, err := os.Open("/proc/net/dev")
	if err != nil {
		return
	}
	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 4096)

	line, _, err := reader.ReadLine()
	line, _, err = reader.ReadLine()
	str := string(line)

	var version int
	if strings.Contains(str, "compressed") {
		version = 3
	} else if strings.Contains(str, "bytes") {
		version = 2
	} else {
		version = 1
	}

	for {
		line, _, err = reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			break
		}
		str = string(line)
		ifname := ifNameCut(str)

		p := strings.IndexByte(str, ':')
		if p < 0 {
			continue
		}
		str = str[p+1:]

		ifp := IfLookupByName(ifname)
		if ifp != nil {
			IfStatsScan(version, str, &ifp.Stats)
		}
	}
}

func (ifp *Interface) ShowStats() string {
	const IfStatTmpl = `    input packets {{.Rx_packets}}, bytes {{.Rx_bytes}}, dropped {{.Rx_dropped}}, multicast packets {{.Rx_multicast}}
    input errors {{.Rx_errors}}, length {{.Rx_length_errors}}, overrun {{.Rx_over_errors}}, CRC {{.Rx_crc_errors}}, frame {{.Rx_frame_errors}}, fifo {{.Rx_fifo_errors}}, missed {{.Rx_missed_errors}}
    output packets {{.Tx_packets}}, bytes {{.Tx_bytes}}, dropped {{.Tx_dropped}}
    output errors {{.Tx_errors}}, aborted {{.Tx_aborted_errors}}, carrier {{.Tx_carrier_errors}}, fifo {{.Tx_fifo_errors}}, heartbeat {{.Tx_heartbeat_errors}}, window {{.Tx_window_errors}}
    collisions {{.Collisions}}
`
	buf := new(bytes.Buffer)
	tmpl := template.Must(template.New("Stats").Parse(IfStatTmpl))
	tmpl.Execute(buf, ifp.Stats)
	return buf.String()
}
