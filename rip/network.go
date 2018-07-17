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
	"unsafe"

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
	// XXX VRF binding.
	if err = SocketBroadcast(sock, 1); err != nil {
		log.Warn(err)
	}
	if err = SocketReusePort(sock, 1); err != nil {
		log.Warn(err)
	}
	if err = SocketReuseAddress(sock, 1); err != nil {
		log.Warn(err)
	}
	if err = SocketPktInfo(sock, 1); err != nil {
		log.Warn(err)
	}
	// recvbuf

	// sockaddr for port bind
	addr := &unix.SockaddrInet4{}
	addr.Port = RIP_PORT_DEFAULT

	// bind.
	err = unix.Bind(sock, addr)
	fmt.Println("Bind:", err)

	return sock
}

func SocketBroadcast(sock int, value int) error {
	return unix.SetsockoptInt(sock, unix.SOL_SOCKET, unix.SO_BROADCAST, value)
}

func SocketReusePort(sock int, value int) error {
	return unix.SetsockoptInt(sock, unix.SOL_SOCKET, unix.SO_REUSEPORT, value)
}

func SocketReuseAddress(sock int, value int) error {
	return unix.SetsockoptInt(sock, unix.SOL_SOCKET, unix.SO_REUSEADDR, value)
}

func SocketIPv4MulticastLoop(sock int, value int) error {
	return unix.SetsockoptInt(sock, unix.IPPROTO_IP, unix.IP_MULTICAST_LOOP, value)
}

func SocketPktInfo(sock int, value int) error {
	return unix.SetsockoptInt(sock, unix.IPPROTO_IP, unix.IP_PKTINFO, value)
}

func SendMulticastPacket(ifp *Interface, p *Packet) error {
	log.Info("SendMulticastPacket")
	sock, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, 0)
	if err != nil {
		return err
	}
	defer unix.Close(sock)

	// XXX VRF binding.
	if err = SocketBroadcast(sock, 1); err != nil {
		log.Warn(err)
	}
	if err = SocketReusePort(sock, 1); err != nil {
		log.Warn(err)
	}
	if err = SocketReuseAddress(sock, 1); err != nil {
		log.Warn(err)
	}
	if err = SocketIPv4MulticastLoop(sock, 0); err != nil {
		log.Warn(err)
	}
	if err = InterfaceMulticastIf(sock, ifp.dev); err != nil {
		log.Warn(err)
	}

	sin := &unix.SockaddrInet4{}
	sin.Port = RIP_PORT_DEFAULT
	copy(sin.Addr[:], RIP_GROUP_ADDR)

	data, err := p.Serialize()
	if err != nil {
		log.Error(err)
		return err
	}
	unix.Sendto(sock, data, 0, sin)

	return nil
}

func multicastIf(sock int, ifAddr []byte, ifIndex uint32) error {
	var mr unix.IPMreqn
	copy(mr.Address[:], ifAddr)
	mr.Ifindex = int32(ifIndex)
	return unix.SetsockoptIPMreqn(sock, unix.IPPROTO_IP, unix.IP_MULTICAST_IF, &mr)
}

func multicastJoin(sock int, mcAddr []byte, ifAddr []byte, ifIndex uint32) error {
	var mr unix.IPMreqn
	copy(mr.Multiaddr[:], mcAddr)
	copy(mr.Address[:], ifAddr)
	mr.Ifindex = int32(ifIndex)
	return unix.SetsockoptIPMreqn(sock, unix.IPPROTO_IP, unix.IP_ADD_MEMBERSHIP, &mr)
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

func (s *Server) ResponseRecv(p *Packet) {
	log.Info("RESPONSE packet")
	// Auth
}

func (s *Server) RequestRecv(p *Packet) {
	log.Info("REQUEST packet")
	// Auth
}

type PktInfo struct {
	ifIndex int32
	ifAddr  net.IP
	dstAddr net.IP
}

func NewPktInfo() *PktInfo {
	return &PktInfo{
		ifAddr:  make([]byte, net.IPv4len),
		dstAddr: make([]byte, net.IPv4len),
	}
}

func Recv(sock int, buf []byte) (int, *unix.SockaddrInet4, *PktInfo, error) {
	var info *PktInfo
	oob := make([]byte, SOCKET_CONTROLL_BUFLEN)

	n, oobn, _, from, err := unix.Recvmsg(sock, buf, oob, 0)
	if err != nil {
		return -1, nil, nil, err
	}
	oob = oob[:oobn]

	fromAddr, ok := from.(*unix.SockaddrInet4)
	if !ok {
		return -1, nil, nil, fmt.Errorf("Recvmsg's from address is not SockaddrInet4")
	}

	if oobn > 0 {
		cmsgs, err := unix.ParseSocketControlMessage(oob)
		if err != nil {
			return -1, nil, nil, err
		}
		for _, cmsg := range cmsgs {
			if cmsg.Header.Level == unix.IPPROTO_IP && cmsg.Header.Type == unix.IP_PKTINFO {
				cmsgInfo := (*unix.Inet4Pktinfo)(unsafe.Pointer(&cmsg.Data[0]))
				info = NewPktInfo()
				info.ifIndex = cmsgInfo.Ifindex
				copy(info.ifAddr, cmsgInfo.Spec_dst[:])
				copy(info.dstAddr, cmsgInfo.Addr[:])
			}
		}
	}
	return n, fromAddr, info, nil
}

func (s *Server) Read() {
	for {
		// Read packet.
		buf := make([]byte, RIP_PACKET_MAXLEN)
		nbytes, from, info, err := Recv(s.Sock, buf)
		if err != nil {
			log.Error(err)
			return
		}
		buf = buf[:nbytes]
		fmt.Println("from:", from)
		fmt.Println("info:", info)

		// Decode packet.
		p := &Packet{}
		err = p.DecodeFromBytes(buf)
		if err != nil {
			log.Info("Parse error")
			continue
		}
		log.Info("Packet:", p)

		// Validate packet.
		err = p.Validate()
		if err != nil {
			continue
		}

		// Process Packet.
		switch p.Command {
		case RIP_REQUEST:
			s.RequestRecv(p)
		case RIP_RESPONSE:
			s.ResponseRecv(p)
		case RIP_TRACEON, RIP_TRACEOFF, RIP_POLL, RIP_POLL_ENTRY:
			// peer_bad_packet()
			log.Warnf("RECV[%s] Obsolete RIP command %s received", Command2Str(p.Command))
		default:
			// peer_bad_packet()
			log.Warnf("RECV[%s] Unknown RIP command %d received", p.Command)
		}
	}
}
