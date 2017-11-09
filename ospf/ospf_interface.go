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

package ospf

import (
	"net"

	"github.com/coreswitch/netutil"
	"github.com/coreswitch/zebra/rib"
)

const (
	OspfIfTypeNone = iota
	OspfIfTypeBroadcast
	OspfIfTypeNBMA
	OspfIfTypePointoPoint
	OspfIfTypePointoPointNBMA
	OspfIfTypePointoMultipoint
	OspfIfTypePointoMultipointNBMA
	OspfIfTypeVirtualLink
	OspfIfTypeLoopback
	OspfIfTypeHost
	OspfIfTypeMax
)

const (
	OSPF_OPTION_T  = (1 << 0) // TOS.
	OSPF_OPTION_E  = (1 << 1)
	OSPF_OPTION_MC = (1 << 2)
	OSPF_OPTION_NP = (1 << 3)
	OSPF_OPTION_EA = (1 << 4)
	OSPF_OPTION_DC = (1 << 5)
	OSPF_OPTION_O  = (1 << 6)
	OSPF_OPTION_L  = OSPF_OPTION_EA // LLS.
	OSPF_OPTION_DN = (1 << 7)       // OSPF Down-bit
)

const (
	OSPF_HELLO_INTERVAL_DEFAULT      = 10
	OSPF_HELLO_INTERVAL_NBMA_DEFAULT = 30
)

type OspfIdent struct {
	Address  *netutil.Prefix
	RouterId net.IP
	DRouter  net.IP
	BDRouter net.IP
	Priority byte
}

type OspfInterface struct {
	Top     *Ospf
	Type    int
	Ifp     *rib.Interface
	Ident   OspfIdent
	Nbrs    *netutil.Ptree
	Area    *OspfArea
	Network *OspfNetwork
}

func NewOspfInterface() *OspfInterface {
	return &OspfInterface{}
}

func (oi *OspfInterface) IsUnnumbered() bool {
	// return oi.Type == OspfIfTypePointoPoint && oi.Ifp.IsUnnumbered()
	return oi.Type == OspfIfTypePointoPoint
}

func (oi *OspfInterface) Options() byte {
	var options byte

	if oi.Area.IsDefault() {
		options = OSPF_OPTION_E
	} else if oi.Area.IsNSSA() {
		options = OSPF_OPTION_NP
	}
	if oi.Top.RestartEnabled() {
		options |= OSPF_OPTION_L
	}
	return options
}

func (oi *OspfInterface) Priority() byte {
	return oi.Ident.Priority
}

func (oi *OspfInterface) HelloInterval() uint16 {
	switch oi.Type {
	case OspfIfTypeNBMA, OspfIfTypePointoMultipoint, OspfIfTypePointoMultipointNBMA:
		return OSPF_HELLO_INTERVAL_NBMA_DEFAULT
	default:
		return OSPF_HELLO_INTERVAL_DEFAULT
	}
}

func (oi *OspfInterface) DeadInterval() uint32 {
	return 0
}

func (oi *OspfInterface) AuthType() byte {
	return 0
}
