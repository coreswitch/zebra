// Copyright 2018 zebra project.
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

package rip

import (
	"fmt"
	"net"

	"github.com/coreswitch/log"
	"golang.org/x/sys/unix"
)

func MakeSocket() int {
	// socket.
	sock, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)

	log.Info("makeSocket", sock, err)
	if err != nil {
		return -1
	}

	// VRF?

	// broadcast
	// reuseaddr
	// reuseport

	// recvif

	// ipv4_dstaddr

	// recvbuf

	// sockaddr for port bind
	addr := &unix.SockaddrInet4{}
	addr.Port = 520

	// bind.
	err = unix.Bind(sock, addr)
	fmt.Println("Bind:", err)

	return sock
}

func multicastJoin(sock int, mcAddr []byte, ifAddr []byte, ifIndex uint32) {
	var mr unix.IPMreqn
	copy(mr.Multiaddr[:], mcAddr)
	copy(mr.Address[:], ifAddr)
	mr.Ifindex = int32(ifIndex)
	fmt.Println(mr)
	err := unix.SetsockoptIPMreqn(sock, unix.IPPROTO_IP, unix.IP_ADD_MEMBERSHIP, &mr)
	fmt.Println("multicastJoin", err)
}

func IsClassA(ip net.IP) bool {
	if len(ip) == net.IPv4len {
		return ip[0]&0x80 == 0
	}
	return false
}

func IsClassABroadcast(ip net.IP) bool {
	if len(ip) == net.IPv4len {
		return ip[1] == 0xff && ip[2] == 0xff && ip[3] == 0xff
	}
	return false
}

func IsClassB(ip net.IP) bool {
	if len(ip) == net.IPv4len {
		return ip[0]&0xc0 == 0x80
	}
	return false
}

func IsClassBBroadcast(ip net.IP) bool {
	if len(ip) == net.IPv4len {
		return ip[2] == 0xff && ip[3] == 0xff
	}
	return false
}

func IsClassC(ip net.IP) bool {
	if len(ip) == net.IPv4len {
		return ip[0]&0xe0 == 0xc0
	}
	return false
}

func IsClassCBroadcast(ip net.IP) bool {
	if len(ip) == net.IPv4len {
		return ip[3] == 0xff
	}
	return false
}

// Check if the destination address is valid (unicast; not net 0 or 127)
// (RFC2453 Section 3.9.2 - Page 26). But we don't check net 0 because we accept
// default route. And if the destination address is CLASS A, B, C broadcast
// address(RIPv1), reject this route.
func destinationCheck(addr net.IP) bool {
	if len(addr) != 4 {
		return false
	}
	if addr.IsLoopback() {
		return false
	}
	if addr[0] == 0 {
		if addr.IsUnspecified() {
			return true
		} else {
			return false
		}
	}
	if IsClassA(addr) {
		if IsClassABroadcast(addr) {
			return false
		} else {
			return true
		}
	}
	if IsClassB(addr) {
		if IsClassBBroadcast(addr) {
			return false
		} else {
			return true
		}
	}
	if IsClassC(addr) {
		if IsClassCBroadcast(addr) {
			return false
		} else {
			return true
		}
	}
	return false
}

func (s *Server) PacketRecv() {
	s.Buffer = make([]byte, 1024)
	nbytes, err := unix.Read(s.Sock, s.Buffer)
	fmt.Println("XXX Read", nbytes, err)
	s.Buffer = s.Buffer[0:nbytes]

	// unix.Recvmsg(s.Sock, p []byte, oob []byte, flags int)
}

func (s *Server) Read() {
	for {
		log.Info("Start Read")
		s.PacketRecv()
		s.PacketParse()
	}
}
