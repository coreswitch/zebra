// Copyright 2018 OpenConfigd Project.
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
	"fmt"
	"github.com/coreswitch/log"
	"github.com/jamesharr/expect"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"net"
	"os"
	"os/exec"
	"time"
)

const (
	passReq = `Password:`
)

type QuaggaAccess struct {
	Exp *expect.Expect

	vrf    string
	addr   net.IP
	prog   string
	prompt string
	fp     *os.File
}

func (qa *QuaggaAccess) Init(vrf string, addr net.IP, prog string, timeout time.Duration) error {
	var err error

	qa.vrf = vrf
	qa.addr = addr
	qa.prog = prog
	if out, err := exec.Command("hostname").Output(); err == nil {
		qa.prompt = string(out[:len(out)-1]) + "> "
	} else {
		log.Error("QuaggaAccess.Init(): hostname: ", err)
		return err
	}

	log.Info("QuaggaAccess.Init(): prompt: ", qa.prompt)

	if qa.fp, err = ioutil.TempFile("", "quagga-"+prog); err == nil {
		cmd := fmt.Sprintf("#!/bin/sh\nVRF=%s LD_PRELOAD=/usr/bin/vrf_socket.so ", vrf)
		cmd += fmt.Sprintf("telnet '%s' '%s'\n", addr.String(), prog)
		_, err = qa.fp.WriteString(cmd)
		qa.fp.Close()
		if err != nil {
			os.Remove(qa.fp.Name())
			log.Error("QuaggaAccess.Init(): ", err)
			return err
		}
		if err = os.Chmod(qa.fp.Name(), 0755); err != nil {
			log.Error("QuaggaAccess.Init(): os.Chmod(): ", err)
			os.Remove(qa.fp.Name())
			qa.fp = nil
			return err
		}
		if qa.Exp, err = expect.Spawn(qa.fp.Name()); err == nil {
			qa.Exp.SetTimeout(timeout)
		} else {
			log.Error("QuaggaAccess.Init(): expect.Spawn(): ", err)
			os.Remove(qa.fp.Name())
			qa.Exp = nil
			qa.fp = nil
		}
	} else {
		log.Error("QuaggaAccess.Init(): ioutil.Tempfile(): ", err)
		if qa.fp != nil {
			os.Remove(qa.fp.Name())
			qa.fp = nil
		}
	}
	return err
}

func (qa *QuaggaAccess) GetTempFileName() string {
	if qa.fp == nil {
		return ""
	}
	return qa.fp.Name()
}

func (qa *QuaggaAccess) Close() {
	os.Remove(qa.fp.Name())
	qa.Exp.Close()
	qa.Exp = nil
	qa.fp = nil
}

func (qa *QuaggaAccess) Auth(pass string) error {
	qa.Exp.Send("\n")
	if _, err := qa.Exp.Expect(passReq); err != nil {
		log.Errorf("Error: QuaggaAccess.Auth(): %s", passReq)
		return err
	}
	qa.Exp.Send(pass + "\n")
	if _, err := qa.Exp.Expect(qa.prompt); err != nil {
		log.Errorf("Error: QuaggaAccess.Auth(): password")
		return err
	}
	return nil
}

func (qa *QuaggaAccess) Batch(in []string) []string {
	var out []string

	for _, s := range in {
		qa.Exp.Send(s)
		if m, err := qa.Exp.Expect(qa.prompt); err == nil {
			log.Info("QuaggaAccess.Batch(): ", m.Before)
			out = append(out, m.Before)
		} else {
			out = append(out, "")
		}
	}
	return out
}

func VrfIdByName(vrf string) (int, error) {
	if l, err := netlink.LinkByName(vrf); err == nil {
		switch l := l.(type) {
		case *netlink.Vrf:
			return int(l.Table), nil
		default:
			return -1, fmt.Errorf("VrfIdByName(%s): not a VRF", vrf)
		}
	} else {
		return -1, err
	}
}

func VrfGetRoutes(vrf string, family int, tableType int) ([]netlink.Route, error) {
	if tid, err := VrfIdByName(vrf); err == nil {
		routeFilter := &netlink.Route{
			Table:    tid,
			Type:     tableType,
			Protocol: unix.RTPROT_KERNEL,
		}
		filterMask := netlink.RT_FILTER_TABLE |
			netlink.RT_FILTER_TYPE |
			netlink.RT_FILTER_PROTOCOL
		return netlink.RouteListFiltered(family, routeFilter, filterMask)
	} else {
		log.Errorf("VrfGetRoutes(%s): %v", vrf, err)
		return nil, err
	}
}

func VrfGetIPv4localRoutes(vrf string) ([]netlink.Route, error) {
	return VrfGetRoutes(vrf, netlink.FAMILY_V4, unix.RTN_LOCAL)
}

func VrfQuaggaAccess(
	vrf string, // VRF name
	prog string, // program name: bgpd, ospfd, etc.
	pass string, // passowrd
	timeout time.Duration) (*QuaggaAccess, error) {
	// timeout
	log.Infof("VrfQuaggaAccess(%v, %v, %v, %v)", vrf, prog, pass, timeout)

	if routes, err := VrfGetIPv4localRoutes(vrf); err == nil {
		for _, route := range routes {
			if route.Dst != nil {
				log.Infof("VrfQuaggaAccess(%v)", route.Dst.String())
				if ones, _ := route.Dst.Mask.Size(); ones == 32 {
					var qa QuaggaAccess

					if err := qa.Init(vrf, route.Dst.IP, prog, timeout); err == nil {
						//qa.Exp.SetLogger(expect.StderrLogger())
						if err := qa.Auth(pass); err == nil {
							return &qa, nil
						} else {
							log.Warn("VrfQuaggaAccess(): auth failure")
							qa.Close()
							return nil, fmt.Errorf("VrfQuaggaAccess(): auth failure")
						}
					} else {
						s := "VrfQuaggaAccess(): QuaggaAccess.Init(%v, %v, %v, %v)"
						s = fmt.Sprintf(s, vrf, route.Dst.IP, prog, timeout)
						log.Errorf(s)
						err = fmt.Errorf(s)
					}
				} else {
					log.Warnf(" mask is not 32: %v\n", route.Dst.Mask)
					continue
				}
			}
		}
		if err == nil {
			err = fmt.Errorf("failed to connect to %s on %s", prog, vrf)
		}
		return nil, err
	} else {
		log.Errorf("VrfQuaggaAccess(%v): VrfGetIPv4localRoutes(): %v", vrf, err)
		return nil, err
	}
}

func VrfQuaggaGet(
	vrf string, // VRF name
	prog string, // program name: bgpd, ospfd, etc.
	pass string, // passowrd
	timeout time.Duration, // timeout
	in []string) ([]string, error) {
	// array of input string

	if qa, err := VrfQuaggaAccess(vrf, prog, pass, timeout); err == nil {
		out := qa.Batch(in)
		qa.Close()
		return out, err
	} else {
		log.Error("VrfQuaggaGet(): VrfQuaggaAccess()", err)
		return nil, err
	}
}
