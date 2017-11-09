package quagga

import (
	"fmt"
	"io"
)

type quaggaBgpAggregateAddress struct {
	asSet       bool
	summaryOnly bool
}

func newQuaggaBgpAggregateAddress() *quaggaBgpAggregateAddress {
	aggregateAddress := &quaggaBgpAggregateAddress{}
	return aggregateAddress
}

func quaggaBgpAggregateAddressKey(arg interface{}) (string, bool) {
	if bgpConfigState == nil {
		return "", false
	}
	if s, ok := arg.(string); ok {
		return s, true
	}
	if stringer, ok := arg.(fmt.Stringer); ok {
		return stringer.String(), true
	}
	return "", false
}

func quaggaBgpAggregateAddressLookup(arg interface{}) (string, *quaggaBgpAggregateAddress, bool) {
	aggrAddrKey, ok := quaggaBgpAggregateAddressKey(arg)
	if !ok {
		return "", nil, false
	}
	aggrAddrValue, ok := bgpConfigState.aggregateAddress[aggrAddrKey]
	if !ok {
		return "", nil, false
	}
	return aggrAddrKey, aggrAddrValue, true
}

func quaggaBgpAggregateAddressCreate(arg interface{}) (*quaggaBgpAggregateAddress, bool) {
	aggrAddrKey, ok := quaggaBgpAggregateAddressKey(arg)
	if !ok {
		return nil, false
	}
	aggrAddrValue, ok := bgpConfigState.aggregateAddress[aggrAddrKey]
	if !ok {
		aggrAddrValue = newQuaggaBgpAggregateAddress()
		bgpConfigState.aggregateAddress[aggrAddrKey] = aggrAddrValue
		return aggrAddrValue, true
	}
	return aggrAddrValue, false
}

func quaggaBgpAggregateAddressDelete(arg interface{}) bool {
	aggrAddrKey, ok := quaggaBgpAggregateAddressKey(arg)
	if !ok {
		return false
	}
	_, ok = bgpConfigState.aggregateAddress[aggrAddrKey]
	if ok {
		delete(bgpConfigState.aggregateAddress, aggrAddrKey)
		return true
	}
	return false
}

type quaggaBgpNeighbor struct {
	//	remoteAs                      string
	ipv4AttributeUnchanged        bool
	ipv4AttributeUnchangedAsPath  bool
	ipv4AttributeUnchangedMed     bool
	ipv4AttributeUnchangedNextHop bool
	ipv6AttributeUnchanged        bool
	ipv6AttributeUnchangedAsPath  bool
	ipv6AttributeUnchangedMed     bool
	ipv6AttributeUnchangedNextHop bool
	timers                        bool
	timersKeepalive               string
	timersHoldtime                string
}

func newQuaggaBgpNeighbor() *quaggaBgpNeighbor {
	bgpNeighbor := &quaggaBgpNeighbor{}
	return bgpNeighbor
}

func quaggaBgpNeighborKey(arg interface{}) (string, bool) {
	if bgpConfigState == nil {
		return "", false
	}
	if s, ok := arg.(string); ok {
		return s, true
	}
	if stringer, ok := arg.(fmt.Stringer); ok {
		return stringer.String(), true
	}
	return "", false
}

func quaggaBgpNeighborLookup(arg interface{}) (string, *quaggaBgpNeighbor, bool) {
	neighKey, ok := quaggaBgpNeighborKey(arg)
	if !ok {
		return "", nil, false
	}
	neighValue, ok := bgpConfigState.neighbor[neighKey]
	if !ok {
		return "", nil, false
	}
	return neighKey, neighValue, true
}

func quaggaBgpNeighborCreate(arg interface{}) (*quaggaBgpNeighbor, bool) {
	neighKey, ok := quaggaBgpNeighborKey(arg)
	if !ok {
		return nil, false
	}
	neighValue, ok := bgpConfigState.neighbor[neighKey]
	if !ok {
		neighValue = newQuaggaBgpNeighbor()
		bgpConfigState.neighbor[neighKey] = neighValue
		return neighValue, true
	}
	return neighValue, false
}

func quaggaBgpNeighborDelete(arg interface{}) bool {
	neighKey, ok := quaggaBgpNeighborKey(arg)
	if !ok {
		return false
	}
	_, ok = bgpConfigState.neighbor[neighKey]
	if ok {
		delete(bgpConfigState.neighbor, neighKey)
		return true
	}
	return false
}

type quaggaBgpPeerGroup struct {
	//	remoteAs                      string
	ipv4AttributeUnchanged        bool
	ipv4AttributeUnchangedAsPath  bool
	ipv4AttributeUnchangedMed     bool
	ipv4AttributeUnchangedNextHop bool
	ipv6AttributeUnchanged        bool
	ipv6AttributeUnchangedAsPath  bool
	ipv6AttributeUnchangedMed     bool
	ipv6AttributeUnchangedNextHop bool
	//	timers                        bool
	//	timersKeepalive               string
	//	timersHoldtime                string
}

func newQuaggaBgpPeerGroup() *quaggaBgpPeerGroup {
	bgpPeerGroup := &quaggaBgpPeerGroup{}
	return bgpPeerGroup
}

func quaggaBgpPeerGroupKey(arg interface{}) (string, bool) {
	if bgpConfigState == nil {
		return "", false
	}
	if s, ok := arg.(string); ok {
		return s, true
	}
	if stringer, ok := arg.(fmt.Stringer); ok {
		return stringer.String(), true
	}
	return "", false
}

func quaggaBgpPeerGroupLookup(arg interface{}) (string, *quaggaBgpPeerGroup, bool) {
	peerGrpKey, ok := quaggaBgpPeerGroupKey(arg)
	if !ok {
		return "", nil, false
	}
	peerGrpValue, ok := bgpConfigState.peerGroup[peerGrpKey]
	if !ok {
		return "", nil, false
	}
	return peerGrpKey, peerGrpValue, true
}

func quaggaBgpPeerGroupCreate(arg interface{}) (*quaggaBgpPeerGroup, bool) {
	peerGrpKey, ok := quaggaBgpPeerGroupKey(arg)
	if !ok {
		return nil, false
	}
	peerGrpValue, ok := bgpConfigState.peerGroup[peerGrpKey]
	if !ok {
		peerGrpValue = newQuaggaBgpPeerGroup()
		bgpConfigState.peerGroup[peerGrpKey] = peerGrpValue
		return peerGrpValue, true
	}
	return peerGrpValue, false
}

func quaggaBgpPeerGroupDelete(arg interface{}) bool {
	peerGrpKey, ok := quaggaBgpPeerGroupKey(arg)
	if !ok {
		return false
	}
	_, ok = bgpConfigState.peerGroup[peerGrpKey]
	if ok {
		delete(bgpConfigState.peerGroup, peerGrpKey)
		return true
	}
	return false
}

type quaggaBgpNetwork struct {
	backdoor bool
	routeMap string
}

func newQuaggaBgpNetwork() *quaggaBgpNetwork {
	bgpNetwork := &quaggaBgpNetwork{}
	return bgpNetwork
}

func quaggaBgpNetworkKey(arg interface{}) (string, bool) {
	if bgpConfigState == nil {
		return "", false
	}
	if s, ok := arg.(string); ok {
		return s, true
	}
	if stringer, ok := arg.(fmt.Stringer); ok {
		return stringer.String(), true
	}
	return "", false
}

func quaggaBgpNetworkLookup(arg interface{}) (string, *quaggaBgpNetwork, bool) {
	netKey, ok := quaggaBgpNetworkKey(arg)
	if !ok {
		return "", nil, false
	}
	netValue, ok := bgpConfigState.network[netKey]
	if !ok {
		return "", nil, false
	}
	return netKey, netValue, true
}

func quaggaBgpNetworkCreate(arg interface{}) (*quaggaBgpNetwork, bool) {
	netKey, ok := quaggaBgpNetworkKey(arg)
	if !ok {
		return nil, false
	}
	netValue, ok := bgpConfigState.network[netKey]
	if !ok {
		netValue = newQuaggaBgpNetwork()
		bgpConfigState.network[netKey] = netValue
		return netValue, true
	}
	return netValue, false
}

func quaggaBgpNetworkDelete(arg interface{}) bool {
	netKey, ok := quaggaBgpNetworkKey(arg)
	if !ok {
		return false
	}
	_, ok = bgpConfigState.network[netKey]
	if ok {
		delete(bgpConfigState.network, netKey)
		return true
	}
	return false
}

type quaggaBgp struct {
	asNum            string
	aggregateAddress map[string]*quaggaBgpAggregateAddress
	/*
		ipv4RedistributeConnectedMetric      string
		ipv4RedistributeConnectedRouteMap    string
		ipv4RedistributeKernelMetric         string
		ipv4RedistributeKernelRouteMap       string
		ipv4RedistributeOspfMetric           string
		ipv4RedistributeOspfRouteMap         string
		ipv4RedistributeRipMetric            string
		ipv4RedistributeRipRouteMap          string
		ipv4RedistributeStaticMetric         string
		ipv4RedistributeStaticRouteMap       string
		ipv6RedistributeConnectedMetric      string
		ipv6RedistributeConnectedRouteMap    string
		ipv6RedistributeKernelMetric         string
		ipv6RedistributeKernelRouteMap       string
		ipv6RedistributeOspfv3Metric         string
		ipv6RedistributeOspfv3RouteMap       string
		ipv6RedistributeRipngMetric          string
		ipv6RedistributeRipngRouteMap        string
		ipv6RedistributeStaticMetric         string
		ipv6RedistributeStaticRouteMap       string
	*/
	neighbor                             map[string]*quaggaBgpNeighbor
	peerGroup                            map[string]*quaggaBgpPeerGroup
	network                              map[string]*quaggaBgpNetwork
	parametersDampening                  bool
	parametersDampeningHalfLife          string
	parametersDampeningReUse             string
	parametersDampeningStartSuppressTime string
	parametersDampeningMaxSuppressTime   string
	parametersDistanceGlobal             bool
	parametersDistanceGlobalExternal     string
	parametersDistanceGlobalInternal     string
	parametersDistanceGlobalLocal        string
	timers                               bool
	timersKeepalive                      string
	timersHoldtime                       string
}

func newQuaggaBgp() *quaggaBgp {
	bgp := &quaggaBgp{}
	bgp.aggregateAddress = make(map[string]*quaggaBgpAggregateAddress)
	bgp.neighbor = make(map[string]*quaggaBgpNeighbor)
	bgp.peerGroup = make(map[string]*quaggaBgpPeerGroup)
	bgp.network = make(map[string]*quaggaBgpNetwork)
	return bgp
}

var bgpConfigState *quaggaBgp

func quaggaConfigValidBgpAsNum(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		fmt.Fprintln(f, "AS number required.")
		valid = false
		return valid
	}
	if !validatorRange(*asNum, 1, 4294967294) {
		fmt.Fprintln(f, "protocols bgp", *asNum)
		fmt.Fprintln(f, "AS number must be between 1 and 4294967294.")
		valid = false
	}
	return valid
}

func quaggaConfigValidBgpNeighbor(f io.Writer, asNum *string, neigh string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	if !validatorIPv4Address(neigh) && !validatorIPv6Address(neigh) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh)
		fmt.Fprintln(f, "neighbor format error.")
		valid = false
		return valid
	}
	remoteAs := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "remote-as"})
	if remoteAs == nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh)
		fmt.Fprintln(f, "remote-as required.")
		valid = false
	} else if !validatorRange(*remoteAs, 1, 4294967294) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "remote-as", *remoteAs)
		fmt.Fprintln(f, "AS number must be between 1 and 4294967294.")
		valid = false
	}
	// advertisement-interval
	advertisementInterval := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "advertisement-interval"})
	if advertisementInterval != nil && !validatorRange(*advertisementInterval, 0, 600) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"advertisement-interval", *advertisementInterval)
		fmt.Fprintln(f, "must be between 0 and 600.")
		valid = false
	}
	// allowas-in number [2]
	allowasInNumber := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "allowas-in", "number"})
	if allowasInNumber != nil && !validatorRange(*allowasInNumber, 1, 10) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "allowas-in number", *allowasInNumber)
		fmt.Fprintln(f, "allowas-in number must be between 1 and 10.")
		valid = false
	}
	allowasInNumber6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "allowas-in", "number"})
	if allowasInNumber6 != nil && !validatorRange(*allowasInNumber6, 1, 10) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast allowas-in number", *allowasInNumber6)
		fmt.Fprintln(f, "allowas-in number must be between 1 and 10.")
		valid = false
	}
	// default-originate route-map [2]
	defaultOriginateRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "default-originate", "route-map"})
	if defaultOriginateRouteMap != nil && !validatorExistsRouteMap(*defaultOriginateRouteMap) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"default-originate route-map", *defaultOriginateRouteMap)
		fmt.Fprintln(f, "route-map", *defaultOriginateRouteMap, "doesn't exist.")
		valid = false
	}
	defaultOriginateRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "default-originate", "route-map"})
	if defaultOriginateRouteMap6 != nil && !validatorExistsRouteMap(*defaultOriginateRouteMap6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast default-originate route-map", *defaultOriginateRouteMap6)
		fmt.Fprintln(f, "route-map", *defaultOriginateRouteMap6, "doesn't exist.")
		valid = false
	}
	// distribute-list export [2]
	distributeListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "distribute-list", "export"})
	if distributeListExport != nil && !validatorExistsAccessList(*distributeListExport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"distribute-list export", *distributeListExport)
		fmt.Fprintln(f, "access-list", *distributeListExport, "doesn't exist.")
		valid = false
	}
	distributeListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "distribute-list", "export"})
	if distributeListExport6 != nil && !validatorExistsAccessList6(*distributeListExport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast distribute-list export", *distributeListExport6)
		fmt.Fprintln(f, "access-list6", *distributeListExport6, "doesn't exist.")
		valid = false
	}
	// distribute-list import [2]
	distributeListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "distribute-list", "import"})
	if distributeListImport != nil && !validatorExistsAccessList(*distributeListImport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"distribute-list import", *distributeListImport)
		fmt.Fprintln(f, "access-list", *distributeListImport, "doesn't exist.")
		valid = false
	}
	distributeListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "distribute-list", "import"})
	if distributeListImport6 != nil && !validatorExistsAccessList6(*distributeListImport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast distribute-list import", *distributeListImport6)
		fmt.Fprintln(f, "access-list6", *distributeListImport6, "doesn't exist.")
		valid = false
	}
	// ebgp-multihop
	ebgpMultihop := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "ebgp-multihop"})
	if ebgpMultihop != nil && !validatorRange(*ebgpMultihop, 1, 255) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "ebgp-multihop", *ebgpMultihop)
		fmt.Fprintln(f, "ebgp-multihop must be between 1 and 255.")
		valid = false
	}
	// filter-list export [2]
	filterListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "filter-list", "export"})
	if filterListExport != nil && !validatorExistsAsPathList(*filterListExport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "filter-list export", *filterListExport)
		fmt.Fprintln(f, "as-path-list", *filterListExport, "doesn't exist.")
		valid = false
	}
	filterListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "filter-list", "export"})
	if filterListExport6 != nil && !validatorExistsAsPathList(*filterListExport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast filter-list export", *filterListExport6)
		fmt.Fprintln(f, "as-path-list", *filterListExport6, "doesn't exist.")
		valid = false
	}
	// filter-list import [2]
	filterListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "filter-list", "import"})
	if filterListImport != nil && !validatorExistsAsPathList(*filterListImport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "filter-list import", *filterListImport)
		fmt.Fprintln(f, "as-path-list", *filterListImport, "doesn't exist.")
		valid = false
	}
	filterListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "filter-list", "import"})
	if filterListImport6 != nil && !validatorExistsAsPathList(*filterListImport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast filter-list import", *filterListImport6)
		fmt.Fprintln(f, "as-path-list", *filterListImport6, "doesn't exist.")
		valid = false
	}
	// local-as
	localAs := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "local-as"})
	if localAs != nil && !validatorRange(*localAs, 1, 4294967294) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "local-as", *localAs)
		fmt.Fprintln(f, "local-as must be between 1 and 4294967294.")
		valid = false
	}
	if localAs != nil && *localAs == *asNum {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "local-as", *localAs)
		fmt.Fprintln(f, "you can't set local-as the same as the router AS.")
		valid = false
	}
	// maximum-prefix [2]
	maximumPrefix := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "maximum-prefix"})
	if maximumPrefix != nil && !validatorRange(*maximumPrefix, 1, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "maximum-prefix", *maximumPrefix)
		fmt.Fprintln(f, "maximum-prefix must be between 1 and 4294967295.")
		valid = false
	}
	maximumPrefix6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "maximum-prefix"})
	if maximumPrefix6 != nil && !validatorRange(*maximumPrefix6, 1, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast maximum-prefix", *maximumPrefix6)
		fmt.Fprintln(f, "maximum-prefix must be between 1 and 4294967295.")
		valid = false
	}
	// password
	password := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "password"})
	if password != nil && len(*password) < 1 && len(*password) > 80 {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "password")
		fmt.Fprintln(f, "password must be 80 characters or less.")
		valid = false
	}
	// peer-group [2]
	peerGroup := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "peer-group"})
	if peerGroup != nil && !validatorExistsPeerGroup(*asNum, *peerGroup) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "peer-group", *peerGroup)
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", *peerGroup, "doesn't exist.")
		valid = false
	}
	peerGroup6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "peer-group"})
	if peerGroup6 != nil && !validatorExistsPeerGroup(*asNum, *peerGroup6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast peer-group", *peerGroup6)
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", *peerGroup6, "doesn't exist.")
		valid = false
	}
	// port
	port := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "port"})
	if port != nil && !validatorRange(*port, 1, 65535) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "port", *port)
		fmt.Fprintln(f, "port must be between 1 and 65535.")
		valid = false
	}
	// prefix-list export [2]
	prefixListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "prefix-list", "export"})
	if prefixListExport != nil && !validatorExistsPrefixList(*prefixListExport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "prefix-list export", *prefixListExport)
		fmt.Fprintln(f, "prefix-list", *prefixListExport, "doesn't exist.")
		valid = false
	}
	prefixListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "prefix-list", "export"})
	if prefixListExport6 != nil && !validatorExistsPrefixList6(*prefixListExport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast prefix-list export", *prefixListExport6)
		fmt.Fprintln(f, "prefix-list6", *prefixListExport6, "doesn't exist.")
		valid = false
	}
	// prefix-list import [2]
	prefixListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "prefix-list", "import"})
	if prefixListImport != nil && !validatorExistsPrefixList(*prefixListImport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "prefix-list import", *prefixListImport)
		fmt.Fprintln(f, "prefix-list", *prefixListImport, "doesn't exist.")
		valid = false
	}
	prefixListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "prefix-list", "import"})
	if prefixListImport6 != nil && !validatorExistsPrefixList6(*prefixListImport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast prefix-list import", *prefixListImport6)
		fmt.Fprintln(f, "prefix-list6", *prefixListImport6, "doesn't exist.")
		valid = false
	}
	// route-map export [2]
	routeMapExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "route-map", "export"})
	if routeMapExport != nil && !validatorExistsRouteMap(*routeMapExport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "route-map export", *routeMapExport)
		fmt.Fprintln(f, "route-map", *routeMapExport, "doesn't exist.")
		valid = false
	}
	routeMapExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "route-map", "export"})
	if routeMapExport6 != nil && !validatorExistsRouteMap(*routeMapExport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast route-map export", *routeMapExport6)
		fmt.Fprintln(f, "route-map", *routeMapExport6, "doesn't exist.")
		valid = false
	}
	// route-map import [2]
	routeMapImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "route-map", "import"})
	if routeMapImport != nil && !validatorExistsRouteMap(*routeMapImport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "route-map import", *routeMapImport)
		fmt.Fprintln(f, "route-map", *routeMapImport, "doesn't exist.")
		valid = false
	}
	routeMapImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "route-map", "import"})
	if routeMapImport6 != nil && !validatorExistsRouteMap(*routeMapImport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast route-map import", *routeMapImport6)
		fmt.Fprintln(f, "route-map", *routeMapImport6, "doesn't exist.")
		valid = false
	}
	// timers connect
	timersConnect := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "timers", "connect"})
	if timersConnect != nil && !validatorRange(*timersConnect, 1, 65535) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "timers connect", *timersConnect)
		fmt.Fprintln(f, "BGP connect timer must be between 0 and 65535.")
		valid = false
	}
	// timers holdtime
	timersHoldtime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "timers", "holdtime"})
	if timersHoldtime != nil && !validatorRange(*timersHoldtime, 1, 65535) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "timers holdtime", *timersHoldtime)
		fmt.Fprintln(f, "Holdtime interval must be 0 or between 4 and 65535.")
		valid = false
	}
	// timers keepalive
	timersKeepalive := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "timers", "keepalive"})
	if timersKeepalive != nil && !validatorRange(*timersKeepalive, 1, 65535) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "timers keepalive", *timersKeepalive)
		fmt.Fprintln(f, "Keepalive interval must be between 1 and 65535.")
		valid = false
	}
	// ttl-security hops
	ttlSecurityHops := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "ttl-security", "hops"})
	if ttlSecurityHops != nil && !validatorRange(*ttlSecurityHops, 1, 254) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "ttl-security hops", *ttlSecurityHops)
		fmt.Fprintln(f, "ttl-security hops must be between 1 and 254.")
		valid = false
	}
	// unsuppress-map [2]
	unsuppressMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "unsuppress-map"})
	if unsuppressMap != nil && !validatorExistsRouteMap(*unsuppressMap) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "unsuppress-map", *unsuppressMap)
		fmt.Fprintln(f, "route-map", *unsuppressMap, "doesn't exist.")
		valid = false
	}
	unsuppressMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "unsuppress-map"})
	if unsuppressMap6 != nil && !validatorExistsRouteMap(*unsuppressMap6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast unsuppress-map", *unsuppressMap6)
		fmt.Fprintln(f, "route-map", *unsuppressMap6, "doesn't exist.")
		valid = false
	}
	// update-source
	updateSource := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "update-source"})
	if updateSource != nil && !validatorSource(*updateSource) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "update-source", *updateSource)
		fmt.Fprintln(f, "update-source format error.")
		valid = false
	}
	// weight
	weight := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "weight"})
	if weight != nil && !validatorRange(*weight, 1, 65535) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "weight", *weight)
		fmt.Fprintln(f, "weight must be between 1 and 65535.")
		valid = false
	}
	// distribute-list export & prefix-list export [2]
	if distributeListExport != nil && prefixListExport != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"distribute-list export", *distributeListExport)
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "prefix-list export", *prefixListExport)
		fmt.Fprintln(f, "you can't set both a prefix-list and a distribute list.")
		valid = false
	}
	if distributeListExport6 != nil && prefixListExport6 != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast distribute-list export", *distributeListExport6)
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast prefix-list export", *prefixListExport6)
		fmt.Fprintln(f, "you can't set both a prefix-list and a distribute list.")
		valid = false
	}
	// distribute-list import & prefix-list import [2]
	if distributeListImport != nil && prefixListImport != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"distribute-list import", *distributeListImport)
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "prefix-list import", *prefixListImport)
		fmt.Fprintln(f, "you can't set both a prefix-list and a distribute list.")
		valid = false
	}
	if distributeListImport6 != nil && prefixListImport6 != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast distribute-list import", *distributeListImport6)
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast prefix-list import", *prefixListImport6)
		fmt.Fprintln(f, "you can't set both a prefix-list and a distribute list.")
		valid = false
	}
	// ebgp-multihop & ttl-security hops
	if ebgpMultihop != nil && ttlSecurityHops != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "ebgp-multihop", *ebgpMultihop)
		fmt.Fprintln(f, "protocols bgp", *asNum, "neighbor", neigh, "ttl-security hops", *ttlSecurityHops)
		fmt.Fprintln(f, "you can't set both ebgp-multihop and ttl-security hops.")
		valid = false
	}
	return valid
}

func quaggaConfigValidBgpNeighbors(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	neighs := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "neighbor"})
	for _, neigh := range neighs {
		if !quaggaConfigValidBgpNeighbor(f, asNum, neigh) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidBgpPeerGroup(f io.Writer, asNum *string, peerGroup string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	remoteAs := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "remote-as"})
	if remoteAs == nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup)
		fmt.Fprintln(f, "remote-as required.")
		valid = false
	} else if !validatorRange(*remoteAs, 1, 4294967294) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "remote-as", *remoteAs)
		fmt.Fprintln(f, "AS number must be between 1 and 4294967294.")
		valid = false
	}
	// allowas-in number [2]
	allowasInNumber := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "allowas-in", "number"})
	if allowasInNumber != nil && !validatorRange(*allowasInNumber, 1, 10) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "allowas-in number", *allowasInNumber)
		fmt.Fprintln(f, "allowas-in number must be between 1 and 10.")
		valid = false
	}
	allowasInNumber6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "allowas-in", "number"})
	if allowasInNumber6 != nil && !validatorRange(*allowasInNumber6, 1, 10) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast allowas-in number", *allowasInNumber6)
		fmt.Fprintln(f, "allowas-in number must be between 1 and 10.")
		valid = false
	}

	// default-originate route-map [2]
	defaultOriginateRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "default-originate", "route-map"})
	if defaultOriginateRouteMap != nil && !validatorExistsRouteMap(*defaultOriginateRouteMap) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"default-originate route-map", *defaultOriginateRouteMap)
		fmt.Fprintln(f, "route-map", *defaultOriginateRouteMap, "doesn't exist.")
		valid = false
	}
	defaultOriginateRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "default-originate", "route-map"})
	if defaultOriginateRouteMap6 != nil && !validatorExistsRouteMap(*defaultOriginateRouteMap6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast default-originate route-map", *defaultOriginateRouteMap6)
		fmt.Fprintln(f, "route-map", *defaultOriginateRouteMap6, "doesn't exist.")
		valid = false
	}
	// distribute-list export [2]
	distributeListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "distribute-list", "export"})
	if distributeListExport != nil && !validatorExistsAccessList(*distributeListExport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"distribute-list export", *distributeListExport)
		fmt.Fprintln(f, "access-list", *distributeListExport, "doesn't exist.")
		valid = false
	}
	distributeListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "distribute-list", "export"})
	if distributeListExport6 != nil && !validatorExistsAccessList6(*distributeListExport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast distribute-list export", *distributeListExport6)
		fmt.Fprintln(f, "access-list6", *distributeListExport6, "doesn't exist.")
		valid = false
	}
	// distribute-list import [2]
	distributeListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "distribute-list", "import"})
	if distributeListImport != nil && !validatorExistsAccessList(*distributeListImport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"distribute-list import", *distributeListImport)
		fmt.Fprintln(f, "access-list", *distributeListImport, "doesn't exist.")
		valid = false
	}
	distributeListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "distribute-list", "import"})
	if distributeListImport6 != nil && !validatorExistsAccessList6(*distributeListImport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast distribute-list import", *distributeListImport6)
		fmt.Fprintln(f, "access-list6", *distributeListImport6, "doesn't exist.")
		valid = false
	}
	// ebgp-multihop
	ebgpMultihop := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "ebgp-multihop"})
	if ebgpMultihop != nil && !validatorRange(*ebgpMultihop, 1, 255) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "ebgp-multihop", *ebgpMultihop)
		fmt.Fprintln(f, "ebgp-multihop must be between 1 and 255.")
		valid = false
	}
	// filter-list export [2]
	filterListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "filter-list", "export"})
	if filterListExport != nil && !validatorExistsAsPathList(*filterListExport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "filter-list export", *filterListExport)
		fmt.Fprintln(f, "as-path-list", *filterListExport, "doesn't exist.")
		valid = false
	}
	filterListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "filter-list", "export"})
	if filterListExport6 != nil && !validatorExistsAsPathList(*filterListExport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast filter-list export", *filterListExport6)
		fmt.Fprintln(f, "as-path-list", *filterListExport6, "doesn't exist.")
		valid = false
	}
	// filter-list import [2]
	filterListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "filter-list", "import"})
	if filterListImport != nil && !validatorExistsAsPathList(*filterListImport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "filter-list import", *filterListImport)
		fmt.Fprintln(f, "as-path-list", *filterListImport, "doesn't exist.")
		valid = false
	}
	filterListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "filter-list", "import"})
	if filterListImport6 != nil && !validatorExistsAsPathList(*filterListImport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast filter-list import", *filterListImport6)
		fmt.Fprintln(f, "as-path-list", *filterListImport6, "doesn't exist.")
		valid = false
	}
	// local-as
	localAs := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "local-as"})
	if localAs != nil && !validatorRange(*localAs, 1, 4294967294) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "local-as", *localAs)
		fmt.Fprintln(f, "local-as must be between 1 and 4294967294.")
		valid = false
	}
	if localAs != nil && *localAs == *asNum {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "local-as", *localAs)
		fmt.Fprintln(f, "you can't set local-as the same as the router AS.")
		valid = false
	}
	// maximum-prefix [2]
	maximumPrefix := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "maximum-prefix"})
	if maximumPrefix != nil && !validatorRange(*maximumPrefix, 1, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "maximum-prefix", *maximumPrefix)
		fmt.Fprintln(f, "maximum-prefix must be between 1 and 4294967295.")
		valid = false
	}
	maximumPrefix6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "maximum-prefix"})
	if maximumPrefix6 != nil && !validatorRange(*maximumPrefix6, 1, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast maximum-prefix", *maximumPrefix6)
		fmt.Fprintln(f, "maximum-prefix must be between 1 and 4294967295.")
		valid = false
	}
	// password
	password := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "password"})
	if password != nil && len(*password) < 1 && len(*password) > 80 {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "password")
		fmt.Fprintln(f, "password must be 80 characters or less.")
		valid = false
	}
	// prefix-list export [2]
	prefixListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "prefix-list", "export"})
	if prefixListExport != nil && !validatorExistsPrefixList(*prefixListExport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "prefix-list export", *prefixListExport)
		fmt.Fprintln(f, "prefix-list", *prefixListExport, "doesn't exist.")
		valid = false
	}
	prefixListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "prefix-list", "export"})
	if prefixListExport6 != nil && !validatorExistsPrefixList6(*prefixListExport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast prefix-list export", *prefixListExport6)
		fmt.Fprintln(f, "prefix-list6", *prefixListExport6, "doesn't exist.")
		valid = false
	}
	// prefix-list import [2]
	prefixListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "prefix-list", "import"})
	if prefixListImport != nil && !validatorExistsPrefixList(*prefixListImport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "prefix-list import", *prefixListImport)
		fmt.Fprintln(f, "prefix-list", *prefixListImport, "doesn't exist.")
		valid = false
	}
	prefixListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "prefix-list", "import"})
	if prefixListImport6 != nil && !validatorExistsPrefixList6(*prefixListImport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast prefix-list import", *prefixListImport6)
		fmt.Fprintln(f, "prefix-list6", *prefixListImport6, "doesn't exist.")
		valid = false
	}
	// route-map export [2]
	routeMapExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "route-map", "export"})
	if routeMapExport != nil && !validatorExistsRouteMap(*routeMapExport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "route-map export", *routeMapExport)
		fmt.Fprintln(f, "route-map", *routeMapExport, "doesn't exist.")
		valid = false
	}
	routeMapExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "route-map", "export"})
	if routeMapExport6 != nil && !validatorExistsRouteMap(*routeMapExport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast route-map export", *routeMapExport6)
		fmt.Fprintln(f, "route-map", *routeMapExport6, "doesn't exist.")
		valid = false
	}
	// route-map import [2]
	routeMapImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "route-map", "import"})
	if routeMapImport != nil && !validatorExistsRouteMap(*routeMapImport) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "route-map import", *routeMapImport)
		fmt.Fprintln(f, "route-map", *routeMapImport, "doesn't exist.")
		valid = false
	}
	routeMapImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "route-map", "import"})
	if routeMapImport6 != nil && !validatorExistsRouteMap(*routeMapImport6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast route-map import", *routeMapImport6)
		fmt.Fprintln(f, "route-map", *routeMapImport6, "doesn't exist.")
		valid = false
	}
	// ttl-security hops
	ttlSecurityHops := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "ttl-security", "hops"})
	if ttlSecurityHops != nil && !validatorRange(*ttlSecurityHops, 1, 254) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "ttl-security hops", *ttlSecurityHops)
		fmt.Fprintln(f, "ttl-security hops must be between 1 and 254.")
		valid = false
	}
	// unsuppress-map [2]
	unsuppressMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "unsuppress-map"})
	if unsuppressMap != nil && !validatorExistsRouteMap(*unsuppressMap) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "unsuppress-map", *unsuppressMap)
		fmt.Fprintln(f, "route-map", *unsuppressMap, "doesn't exist.")
		valid = false
	}
	unsuppressMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "unsuppress-map"})
	if unsuppressMap6 != nil && !validatorExistsRouteMap(*unsuppressMap6) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast unsuppress-map", *unsuppressMap6)
		fmt.Fprintln(f, "route-map", *unsuppressMap6, "doesn't exist.")
		valid = false
	}
	// update-source
	updateSource := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "update-source"})
	if updateSource != nil && !validatorSource(*updateSource) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "update-source", *updateSource)
		fmt.Fprintln(f, "update-source format error.")
		valid = false
	}
	// weight
	weight := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "weight"})
	if weight != nil && !validatorRange(*weight, 1, 65535) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "weight", *weight)
		fmt.Fprintln(f, "weight must be between 1 and 65535.")
		valid = false
	}
	// distribute-list export & prefix-list export [2]
	if distributeListExport != nil && prefixListExport != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"distribute-list export", *distributeListExport)
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "prefix-list export", *prefixListExport)
		fmt.Fprintln(f, "you can't set both a prefix-list and a distribute list.")
		valid = false
	}
	if distributeListExport6 != nil && prefixListExport6 != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast distribute-list export", *distributeListExport6)
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast prefix-list export", *prefixListExport6)
		fmt.Fprintln(f, "you can't set both a prefix-list and a distribute list.")
		valid = false
	}
	// distribute-list import & prefix-list import [2]
	if distributeListImport != nil && prefixListImport != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"distribute-list import", *distributeListImport)
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "prefix-list import", *prefixListImport)
		fmt.Fprintln(f, "you can't set both a prefix-list and a distribute list.")
		valid = false
	}
	if distributeListImport6 != nil && prefixListImport6 != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast distribute-list import", *distributeListImport6)
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast prefix-list import", *prefixListImport6)
		fmt.Fprintln(f, "you can't set both a prefix-list and a distribute list.")
		valid = false
	}
	// ebgp-multihop & ttl-security hops
	if ebgpMultihop != nil && ttlSecurityHops != nil {
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "ebgp-multihop", *ebgpMultihop)
		fmt.Fprintln(f, "protocols bgp", *asNum, "peer-group", peerGroup, "ttl-security hops", *ttlSecurityHops)
		fmt.Fprintln(f, "you can't set both ebgp-multihop and ttl-security hops.")
		valid = false
	}
	return valid
}

func quaggaConfigValidBgpPeerGroups(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	peerGroups := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "peer-group"})
	for _, peerGroup := range peerGroups {
		if !quaggaConfigValidBgpPeerGroup(f, asNum, peerGroup) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidBgpAggregateAddresses(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	aggregateAddresses := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "aggregate-address"})
	for _, aggregateAddress := range aggregateAddresses {
		if !validatorIPv4Address(aggregateAddress) {
			valid = false
		}
	}
	aggregateAddress6s := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast", "aggregate-address"})
	for _, aggregateAddress6 := range aggregateAddress6s {
		if !validatorIPv6Address(aggregateAddress6) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidBgpNetworks(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	networks := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "network"})
	for _, network := range networks {
		if !validatorIPv4CIDR(network) {
			valid = false
		}
		routeMap := configCandidate.value(
			[]string{"protocols", "bgp", *asNum, "network", network, "route-map"})
		if routeMap != nil && !validatorExistsRouteMap(*routeMap) {
			valid = false
		}
	}
	network6s := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast", "network"})
	for _, network6 := range network6s {
		if !validatorIPv6CIDR(network6) {
			valid = false
		}
		routeMap6 := configCandidate.value(
			[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
				"network", network6, "route-map"})
		if routeMap6 != nil && !validatorExistsRouteMap(*routeMap6) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidBgpParameterses(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	//protocols bgp WORD parameters cluster-id A.B.C.D
	clusterId := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "cluster-id"})
	if clusterId != nil && !validatorIPv4Address(*clusterId) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters cluster-id", *clusterId)
		fmt.Fprintln(f, "cluster-id format error.")
		valid = false
	}
	//protocols bgp WORD parameters confederation identifier WORD
	confederationIdentifier := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "confederation", "identifier"})
	if confederationIdentifier != nil && !validatorRange(*confederationIdentifier, 1, 4294967294) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters confederation identifier", *confederationIdentifier)
		fmt.Fprintln(f, "confederation AS id must be between 1 and 4294967294.")
		valid = false
	}
	//protocols bgp WORD parameters confederation peers WORD
	confederationPeers := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "parameters", "confederation", "peers"})
	for _, confederationPeer := range confederationPeers {
		if !validatorRange(confederationPeer, 1, 4294967294) {
			fmt.Fprintln(f, "protocols bgp", *asNum, "parameters confederation peers", confederationPeer)
			fmt.Fprintln(f, "confederation AS id must be between 1 and 4294967294.")
			valid = false
		}
		if !validatorConfediBGPASNCheck(confederationPeer, *asNum) {
			fmt.Fprintln(f, "protocols bgp", *asNum, "parameters confederation peers", confederationPeer)
			fmt.Fprintln(f, "Can't set confederation peers ASN to", confederationPeer, ".")
			fmt.Fprintln(f, "Delete any neighbors with remote-as", confederationPeer,
				"and/or change the local ASN first.")
			valid = false
		}
	}
	//protocols bgp WORD parameters dampening half-life WORD
	dampeningHalfLife := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "dampening", "half-life"})
	if dampeningHalfLife != nil && !validatorRange(*dampeningHalfLife, 1, 45) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters dampening half-life", *dampeningHalfLife)
		fmt.Fprintln(f, "Half-life penalty must be between 1 and 45.")
		valid = false
	}
	//protocols bgp WORD parameters dampening max-suppress-time WORD
	dampeningMaxSuppressTime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "dampening", "max-suppress-time"})
	if dampeningMaxSuppressTime != nil && !validatorRange(*dampeningMaxSuppressTime, 1, 255) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"parameters dampening max-suppress-time", *dampeningMaxSuppressTime)
		fmt.Fprintln(f, "Max-suppress-time must be between 1 and 255.")
		valid = false
	}
	//protocols bgp WORD parameters dampening re-use WORD
	dampeningReUse := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "dampening", "re-use"})
	if dampeningReUse != nil && !validatorRange(*dampeningReUse, 1, 20000) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters dampening re-use", *dampeningReUse)
		fmt.Fprintln(f, "Re-use value must be between 1 and 20000.")
		valid = false
	}
	//protocols bgp WORD parameters dampening start-suppress-time WORD
	dampeningStartSuppressTime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "dampening", "start-suppress-time"})
	if dampeningStartSuppressTime != nil && !validatorRange(*dampeningStartSuppressTime, 1, 20000) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"parameters dampening start-suppress-time", *dampeningStartSuppressTime)
		fmt.Fprintln(f, "Start-suppress-time must be between 1 and 20000.")
		valid = false
	}
	//protocols bgp WORD parameters dampening
	if (dampeningHalfLife != nil || dampeningMaxSuppressTime != nil ||
		dampeningReUse != nil || dampeningStartSuppressTime != nil) &&
		(dampeningHalfLife == nil || dampeningMaxSuppressTime == nil ||
			dampeningReUse == nil || dampeningStartSuppressTime == nil) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters dampening")
		fmt.Fprintln(f, "you must set a half-life, max-suppress-time, re-use, start-suppress-time.")
		valid = false
	}
	//protocols bgp WORD parameters default local-pref WORD
	defaultLocalPref := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "default", "local-pref"})
	if defaultLocalPref != nil && !validatorRange(*defaultLocalPref, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters default local-pref", *defaultLocalPref)
		fmt.Fprintln(f, "default local-pref format error.")
		valid = false
	}
	//protocols bgp WORD parameters distance global external WORD
	distanceGlobalExternal := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "distance", "global", "external"})
	if distanceGlobalExternal != nil && !validatorRange(*distanceGlobalExternal, 1, 255) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters distance global external", *distanceGlobalExternal)
		fmt.Fprintln(f, "Must be between 1-255.")
		valid = false
	}
	//protocols bgp WORD parameters distance global internal WORD
	distanceGlobalInternal := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "distance", "global", "internal"})
	if distanceGlobalInternal != nil && !validatorRange(*distanceGlobalInternal, 1, 255) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters distance global internal", *distanceGlobalInternal)
		fmt.Fprintln(f, "Must be between 1-255.")
		valid = false
	}
	//protocols bgp WORD parameters distance global local WORD
	distanceGlobalLocal := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "distance", "global", "local"})
	if distanceGlobalLocal != nil && !validatorRange(*distanceGlobalLocal, 1, 255) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters distance global local", *distanceGlobalLocal)
		fmt.Fprintln(f, "Must be between 1-255.")
		valid = false
	}
	//protocols bgp WORD parameters distance global
	if (distanceGlobalExternal != nil || distanceGlobalInternal != nil || distanceGlobalLocal != nil) &&
		(distanceGlobalExternal == nil || distanceGlobalInternal == nil || distanceGlobalLocal == nil) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters distance global")
		fmt.Fprintln(f, "you must set an external, internal, local distance.")
		valid = false
	}
	//protocols bgp WORD parameters distance prefix A.B.C.D/M
	//protocols bgp WORD parameters distance prefix A.B.C.D/M distance WORD
	distancePrefixes := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "parameters", "distance", "prefix"})
	for _, distancePrefix := range distancePrefixes {
		if !validatorIPv4CIDR(distancePrefix) {
			fmt.Fprintln(f, "protocols bgp", *asNum, "parameters distance prefix", distancePrefix)
			fmt.Fprintln(f, "format error.")
			valid = false
		}
		distance := configCandidate.value(
			[]string{"protocols", "bgp", *asNum, "parameters", "distance", "prefix", distancePrefix,
				"distance"})
		if distance == nil || !validatorRange(*distance, 1, 255) {
			t := ""
			if distance != nil {
				t = *distance
			}
			fmt.Fprintln(f, "protocols bgp", *asNum, "parameters distance prefix", distancePrefix,
				"distance", t)
			fmt.Fprintln(f, "Must be between 1-255.")
			valid = false
		}
	}
	//protocols bgp WORD parameters graceful-restart stalepath-time WORD
	gracefulRestartStalepathTime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "graceful-restart", "stalepath-time"})
	if gracefulRestartStalepathTime != nil && !validatorRange(*gracefulRestartStalepathTime, 1, 3600) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters graceful-restart stalepath-time",
			*gracefulRestartStalepathTime)
		fmt.Fprintln(f, "stalepath-time must be between 1 and 3600.")
		valid = false
	}
	//protocols bgp WORD parameters router-id A.B.C.D
	routerId := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "router-id"})
	if routerId != nil && !validatorIPv4Address(*routerId) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters router-id", *routerId)
		fmt.Fprintln(f, "format error.")
		valid = false
	}
	//protocols bgp WORD parameters scan-time WORD
	scanTime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "scan-time"})
	if scanTime != nil && !validatorRange(*scanTime, 5, 60) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "parameters scan-time", *scanTime)
		fmt.Fprintln(f, "scan-time must be between 5 and 60 seconds.")
		valid = false
	}
	return valid
}

func quaggaConfigValidBgpRedistributes(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute connected metric WORD
	connectedMetric6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "connected", "metric"})
	if connectedMetric6 != nil && !validatorRange(*connectedMetric6, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute connected metric", *connectedMetric6)
		fmt.Fprintln(f, "metric format error.")
		valid = false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute connected route-map WORD
	connectedRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "connected", "route-map"})
	if connectedRouteMap6 != nil && !validatorExistsRouteMap(*connectedRouteMap6) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute connected route-map", *connectedRouteMap6)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute kernel metric WORD
	kernelMetric6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "kernel", "metric"})
	if kernelMetric6 != nil && !validatorRange(*kernelMetric6, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute kernel metric", *kernelMetric6)
		fmt.Fprintln(f, "metric format error.")
		valid = false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute kernel route-map WORD
	kernelRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "kernel", "route-map"})
	if kernelRouteMap6 != nil && !validatorExistsRouteMap(*kernelRouteMap6) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute kernel route-map", *kernelRouteMap6)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute ospfv3 metric WORD
	ospfv3Metric6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "ospfv3", "metric"})
	if ospfv3Metric6 != nil && !validatorRange(*ospfv3Metric6, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute ospfv3 metric", *ospfv3Metric6)
		fmt.Fprintln(f, "metric format error.")
		valid = false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute ospfv3 route-map WORD
	ospfv3RouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "ospfv3", "route-map"})
	if ospfv3RouteMap6 != nil && !validatorExistsRouteMap(*ospfv3RouteMap6) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute ospfv3 route-map", *ospfv3RouteMap6)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute ripng metric WORD
	//protocols bgp WORD address-family ipv6-unicast redistribute ripng route-map WORD
	//protocols bgp WORD address-family ipv6-unicast redistribute static metric WORD
	staticMetric6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "static", "metric"})
	if staticMetric6 != nil && !validatorRange(*staticMetric6, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute static metric", *staticMetric6)
		fmt.Fprintln(f, "metric format error.")
		valid = false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute static route-map WORD
	staticRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "static", "route-map"})
	if staticRouteMap6 != nil && !validatorExistsRouteMap(*staticRouteMap6) {
		fmt.Fprintln(f, "protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute static route-map", *staticRouteMap6)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	//protocols bgp WORD redistribute connected metric WORD
	connectedMetric := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "connected", "metric"})
	if connectedMetric != nil && !validatorRange(*connectedMetric, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "redistribute connected metric", *connectedMetric)
		fmt.Fprintln(f, "metric format error.")
		valid = false
	}
	//protocols bgp WORD redistribute connected route-map WORD
	connectedRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "connected", "route-map"})
	if connectedRouteMap != nil && !validatorExistsRouteMap(*connectedRouteMap) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "redistribute connected route-map", *connectedRouteMap)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	//protocols bgp WORD redistribute kernel metric WORD
	kernelMetric := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "kernel", "metric"})
	if kernelMetric != nil && !validatorRange(*kernelMetric, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "redistribute kernel metric", *kernelMetric)
		fmt.Fprintln(f, "metric format error.")
		valid = false
	}
	//protocols bgp WORD redistribute kernel route-map WORD
	kernelRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "kernel", "route-map"})
	if kernelRouteMap != nil && !validatorExistsRouteMap(*kernelRouteMap) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "redistribute kernel route-map", *kernelRouteMap)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	//protocols bgp WORD redistribute ospf metric WORD
	ospfMetric := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "ospf", "metric"})
	if ospfMetric != nil && !validatorRange(*ospfMetric, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "redistribute ospf metric", *ospfMetric)
		fmt.Fprintln(f, "metric format error.")
		valid = false
	}
	//protocols bgp WORD redistribute ospf route-map WORD
	ospfRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "ospf", "route-map"})
	if ospfRouteMap != nil && !validatorExistsRouteMap(*ospfRouteMap) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "redistribute ospf route-map", *ospfRouteMap)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	//protocols bgp WORD redistribute rip metric WORD
	//protocols bgp WORD redistribute rip route-map WORD
	//protocols bgp WORD redistribute static metric WORD
	staticMetric := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "static", "metric"})
	if staticMetric != nil && !validatorRange(*staticMetric, 0, 4294967295) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "redistribute static metric", *staticMetric)
		fmt.Fprintln(f, "metric format error.")
		valid = false
	}
	//protocols bgp WORD redistribute static route-map WORD
	staticRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "static", "route-map"})
	if staticRouteMap != nil && !validatorExistsRouteMap(*staticRouteMap) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "redistribute static route-map", *staticRouteMap)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	return valid
}

func quaggaConfigValidBgpMaximumPathses(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	//protocols bgp WORD maximum-paths ebgp WORD
	ebgp := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "maximum-paths", "ebgp"})
	if ebgp != nil && !validatorRange(*ebgp, 1, 255) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "maximum-paths ebgp", *ebgp)
		fmt.Fprintln(f, "Must be between (1-255).")
		valid = false
	}
	//protocols bgp WORD maximum-paths ibgp WORD
	ibgp := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "maximum-paths", "ibgp"})
	if ibgp != nil && !validatorRange(*ibgp, 1, 255) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "maximum-paths ibgp", *ibgp)
		fmt.Fprintln(f, "Must be between (1-255).")
		valid = false
	}
	return valid
}

func quaggaConfigValidBgpTimerses(f io.Writer, asNum *string) bool {
	valid := true
	if asNum == nil {
		valid = false
		return valid
	}
	//protocols bgp WORD timers holdtime WORD
	holdtime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "timers", "holdtime"})
	if holdtime != nil && !validatorRange(*holdtime, 0, 0) && !validatorRange(*holdtime, 4, 65535) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "timers holdtime", *holdtime)
		fmt.Fprintln(f, "hold-time interval must be 0 or between 4 and 65535.")
		valid = false
	}
	//protocols bgp WORD timers keepalive WORD
	keepalive := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "timers", "keepalive"})
	if keepalive != nil && !validatorRange(*keepalive, 1, 65535) {
		fmt.Fprintln(f, "protocols bgp", *asNum, "timers keepalive", *keepalive)
		fmt.Fprintln(f, "Keepalive interval must be between 1 and 65535.")
		valid = false
	}
	return valid
}

func quaggaConfigValidBgp(f io.Writer) bool {
	valid := true
	asNums := configCandidate.values([]string{"protocols", "bgp"})
	if len(asNums) == 0 {
		return valid
	}
	if len(asNums) > 1 {
		for _, asNum := range asNums {
			fmt.Fprintln(f, "protocols bgp", asNum)
		}
		fmt.Fprintln(f, "Multiple BGP can not be set.")
		valid = false
		return valid
	}
	asNum := configCandidate.value([]string{"protocols", "bgp"})
	if !quaggaConfigValidBgpAsNum(f, asNum) {
		valid = false
	}
	if !quaggaConfigValidBgpNeighbors(f, asNum) {
		valid = false
	}
	if !quaggaConfigValidBgpPeerGroups(f, asNum) {
		valid = false
	}
	if !quaggaConfigValidBgpAggregateAddresses(f, asNum) {
		valid = false
	}
	if !quaggaConfigValidBgpNetworks(f, asNum) {
		valid = false
	}
	if !quaggaConfigValidBgpParameterses(f, asNum) {
		valid = false
	}
	if !quaggaConfigValidBgpRedistributes(f, asNum) {
		valid = false
	}
	if !quaggaConfigValidBgpMaximumPathses(f, asNum) {
		valid = false
	}
	if !quaggaConfigValidBgpTimerses(f, asNum) {
		valid = false
	}
	return valid
}
