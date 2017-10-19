package quagga

import (
	"fmt"
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
	remoteAs                      string
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
	remoteAs                      string
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
	asNum                                string
	aggregateAddress                     map[string]*quaggaBgpAggregateAddress
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

func quaggaConfigValidBgpAsNum(asNum *string) bool {
	if asNum == nil {
		fmt.Println("AS number required.")
		return false
	}
	if !validatorRange(*asNum, 1, 4294967294) {
		fmt.Println("protocols bgp", *asNum)
		fmt.Println("AS number must be between 1 and 4294967294.")
		return false
	}
	return true
}

func quaggaConfigValidBgpNeighbor(asNum *string, neigh string) bool {
	if asNum == nil {
		return false
	}
	if !validatorIPv4Address(neigh) && !validatorIPv6Address(neigh) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh)
		fmt.Println("neighbor format error.")
		return false
	}
	remoteAs := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "remote-as"})
	if remoteAs == nil {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh)
		fmt.Println("remote-as required.")
		return false
	}
	if !validatorRange(*remoteAs, 1, 4294967294) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "remote-as", *remoteAs)
		fmt.Println("AS number must be between 1 and 4294967294.")
		return false
	}
	// advertisement-interval
	advertisementInterval := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "advertisement-interval"})
	if advertisementInterval != nil && !validatorRange(*advertisementInterval, 0, 600) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"advertisement-interval", *advertisementInterval)
		fmt.Println("must be between 0 and 600.")
		return false
	}
	// allowas-in number [2]
	allowasInNumber := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "allowas-in", "number"})
	if allowasInNumber != nil && !validatorRange(*allowasInNumber, 1, 10) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "allowas-in number", *allowasInNumber)
		fmt.Println("allowas-in number must be between 1 and 10.")
		return false
	}
	allowasInNumber6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "allowas-in", "number"})
	if allowasInNumber6 != nil && !validatorRange(*allowasInNumber6, 1, 10) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast allowas-in number", *allowasInNumber6)
		fmt.Println("allowas-in number must be between 1 and 10.")
		return false
	}
	// default-originate route-map [2]
	defaultOriginateRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "default-originate", "route-map"})
	if defaultOriginateRouteMap != nil && !validatorExistsRouteMap(*defaultOriginateRouteMap) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"default-originate route-map", *defaultOriginateRouteMap)
		fmt.Println("route-map ", *defaultOriginateRouteMap, " doesn't exist.")
		return false
	}
	defaultOriginateRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "default-originate", "route-map"})
	if defaultOriginateRouteMap6 != nil && !validatorExistsRouteMap(*defaultOriginateRouteMap6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast default-originate route-map", *defaultOriginateRouteMap6)
		fmt.Println("route-map ", *defaultOriginateRouteMap6, " doesn't exist.")
		return false
	}
	// distribute-list export [2]
	distributeListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "distribute-list", "export"})
	if distributeListExport != nil && !validatorExistsAccessList(*distributeListExport) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"distribute-list export", *distributeListExport)
		fmt.Println("access-list ", *distributeListExport, " doesn't exist")
		return false
	}
	distributeListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "distribute-list", "export"})
	if distributeListExport6 != nil && !validatorExistsAccessList6(*distributeListExport6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast distribute-list export", *distributeListExport6)
		fmt.Println("access-list6 ", *distributeListExport6, " doesn't exist")
		return false
	}
	// distribute-list import [2]
	distributeListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "distribute-list", "import"})
	if distributeListImport != nil && !validatorExistsAccessList(*distributeListImport) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"distribute-list import", *distributeListImport)
		fmt.Println("access-list ", *distributeListImport, " doesn't exist")
		return false
	}
	distributeListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "distribute-list", "import"})
	if distributeListImport6 != nil && !validatorExistsAccessList6(*distributeListImport6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast distribute-list import", *distributeListImport6)
		fmt.Println("access-list6 ", *distributeListImport6, " doesn't exist")
		return false
	}
	// ebgp-multihop
	ebgpMultihop := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "ebgp-multihop"})
	if ebgpMultihop != nil && !validatorRange(*ebgpMultihop, 1, 255) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "ebgp-multihop", *ebgpMultihop)
		fmt.Println("ebgp-multihop must be between 1 and 255.")
		return false
	}
	// filter-list export [2]
	filterListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "filter-list", "export"})
	if filterListExport != nil && !validatorExistsAsPathList(*filterListExport) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "filter-list export", *filterListExport)
		fmt.Println("as-path-list ", *filterListExport, " doesn't exist.")
		return false
	}
	filterListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "filter-list", "export"})
	if filterListExport6 != nil && !validatorExistsAsPathList(*filterListExport6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast filter-list export", *filterListExport6)
		fmt.Println("as-path-list ", *filterListExport6, " doesn't exist.")
		return false
	}
	// filter-list import [2]
	filterListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "filter-list", "import"})
	if filterListImport != nil && !validatorExistsAsPathList(*filterListImport) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "filter-list import", *filterListImport)
		fmt.Println("as-path-list ", *filterListImport, " doesn't exist.")
		return false
	}
	filterListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "filter-list", "import"})
	if filterListImport6 != nil && !validatorExistsAsPathList(*filterListImport6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast filter-list import", *filterListImport6)
		fmt.Println("as-path-list ", *filterListImport6, " doesn't exist.")
		return false
	}
	// local-as
	localAs := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "local-as"})
	if localAs != nil && !validatorRange(*localAs, 1, 4294967294) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "local-as", *localAs)
		fmt.Println("local-as must be between 1 and 4294967294.")
		return false
	}
	if localAs != nil && *localAs == *asNum {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "local-as", *localAs)
		fmt.Println("you can't set local-as the same as the router AS.")
		return false
	}
	// maximum-prefix [2]
	maximumPrefix := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "maximum-prefix"})
	if maximumPrefix != nil && !validatorRange(*maximumPrefix, 1, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "maximum-prefix", *maximumPrefix)
		fmt.Println("maximum-prefix must be between 1 and 4294967295.")
		return false
	}
	maximumPrefix6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "maximum-prefix"})
	if maximumPrefix6 != nil && !validatorRange(*maximumPrefix6, 1, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast maximum-prefix", *maximumPrefix6)
		fmt.Println("maximum-prefix must be between 1 and 4294967295.")
		return false
	}
	// password
	password := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "password"})
	if password != nil && len(*password) < 1 && len(*password) > 80 {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "password")
		fmt.Println("password must be 80 characters or less.")
		return false
	}
	// peer-group [2]
	peerGroup := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "peer-group"})
	if peerGroup != nil && !validatorExistsPeerGroup(*asNum, *peerGroup) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "peer-group", *peerGroup)
		fmt.Println("protocols bgp ", *asNum, " peer-group ", *peerGroup, " doesn't exist.")
		return false
	}
	peerGroup6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "peer-group"})
	if peerGroup6 != nil && !validatorExistsPeerGroup(*asNum, *peerGroup6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast peer-group", *peerGroup6)
		fmt.Println("protocols bgp ", *asNum, " peer-group ", *peerGroup6, " doesn't exist.")
		return false
	}
	// port
	port := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "port"})
	if port != nil && !validatorRange(*port, 1, 65535) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "port", *port)
		fmt.Println("port must be between 1 and 65535.")
		return false
	}
	// prefix-list export [2]
	prefixListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "prefix-list", "export"})
	if prefixListExport != nil && !validatorExistsPrefixList(*prefixListExport) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "prefix-list export", *prefixListExport)
		fmt.Println("prefix-list ", *prefixListExport, " doesn't exist.")
		return false
	}
	prefixListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "prefix-list", "export"})
	if prefixListExport6 != nil && !validatorExistsPrefixList6(*prefixListExport6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast prefix-list export", *prefixListExport6)
		fmt.Println("prefix-list6 ", *prefixListExport6, " doesn't exist.")
		return false
	}
	// prefix-list import [2]
	prefixListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "prefix-list", "import"})
	if prefixListImport != nil && !validatorExistsPrefixList(*prefixListImport) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "prefix-list import", *prefixListImport)
		fmt.Println("prefix-list ", *prefixListImport, " doesn't exist.")
		return false
	}
	prefixListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "prefix-list", "import"})
	if prefixListImport6 != nil && !validatorExistsPrefixList6(*prefixListImport6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast prefix-list import", *prefixListImport6)
		fmt.Println("prefix-list6 ", *prefixListImport6, " doesn't exist.")
		return false
	}
	// route-map export [2]
	routeMapExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "route-map", "export"})
	if routeMapExport != nil && !validatorExistsRouteMap(*routeMapExport) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "route-map export", *routeMapExport)
		fmt.Println("route-map ", *routeMapExport, " doesn't exist.")
		return false
	}
	routeMapExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "route-map", "export"})
	if routeMapExport6 != nil && !validatorExistsRouteMap(*routeMapExport6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast route-map export", *routeMapExport6)
		fmt.Println("route-map ", *routeMapExport6, " doesn't exist.")
		return false
	}
	// route-map import [2]
	routeMapImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "route-map", "import"})
	if routeMapImport != nil && !validatorExistsRouteMap(*routeMapImport) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "route-map import", *routeMapImport)
		fmt.Println("route-map ", *routeMapImport, " doesn't exist.")
		return false
	}
	routeMapImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "route-map", "import"})
	if routeMapImport6 != nil && !validatorExistsRouteMap(*routeMapImport6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast route-map import", *routeMapImport6)
		fmt.Println("route-map ", *routeMapImport6, " doesn't exist.")
		return false
	}
	// timers connect
	timersConnect := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "timers", "connect"})
	if timersConnect != nil && !validatorRange(*timersConnect, 1, 65535) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "timers connect", *timersConnect)
		fmt.Println("BGP connect timer must be between 0 and 65535.")
		return false
	}
	// timers holdtime
	timersHoldtime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "timers", "holdtime"})
	if timersHoldtime != nil && !validatorRange(*timersHoldtime, 1, 65535) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "timers holdtime", *timersHoldtime)
		fmt.Println("Holdtime interval must be 0 or between 4 and 65535.")
		return false
	}
	// timers keepalive
	timersKeepalive := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "timers", "keepalive"})
	if timersKeepalive != nil && !validatorRange(*timersKeepalive, 1, 65535) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "timers keepalive", *timersKeepalive)
		fmt.Println("Keepalive interval must be between 1 and 65535.")
		return false
	}
	// ttl-security hops
	ttlSecurityHops := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "ttl-security", "hops"})
	if ttlSecurityHops != nil && !validatorRange(*ttlSecurityHops, 1, 254) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "ttl-security hops", *ttlSecurityHops)
		fmt.Println("ttl-security hops must be between 1 and 254.")
		return false
	}
	// unsuppress-map [2]
	unsuppressMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "unsuppress-map"})
	if unsuppressMap != nil && !validatorExistsRouteMap(*unsuppressMap) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "unsuppress-map", *unsuppressMap)
		fmt.Println("route-map ", *unsuppressMap, " doesn't exist.")
		return false
	}
	unsuppressMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh,
			"address-family", "ipv6-unicast", "unsuppress-map"})
	if unsuppressMap6 != nil && !validatorExistsRouteMap(*unsuppressMap6) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast unsuppress-map", *unsuppressMap6)
		fmt.Println("route-map ", *unsuppressMap6, " doesn't exist.")
		return false
	}
	// update-source
	updateSource := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "update-source"})
	if updateSource != nil && !validatorSource(*updateSource) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "update-source", *updateSource)
		fmt.Println("update-source format error.")
		return false
	}
	// weight
	weight := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "neighbor", neigh, "weight"})
	if weight != nil && !validatorRange(*weight, 1, 65535) {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "weight", *weight)
		fmt.Println("weight must be between 1 and 65535.")
		return false
	}
	// distribute-list export & prefix-list export [2]
	if distributeListExport != nil && prefixListExport != nil {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"distribute-list export", *distributeListExport)
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "prefix-list export", *prefixListExport)
		fmt.Println("you can't set both a prefix-list and a distribute list.")
		return false
	}
	if distributeListExport6 != nil && prefixListExport6 != nil {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast distribute-list export", *distributeListExport6)
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast prefix-list export", *prefixListExport6)
		fmt.Println("you can't set both a prefix-list and a distribute list.")
		return false
	}
	// distribute-list import & prefix-list import [2]
	if distributeListImport != nil && prefixListImport != nil {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"distribute-list import", *distributeListImport)
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "prefix-list import", *prefixListImport)
		fmt.Println("you can't set both a prefix-list and a distribute list.")
		return false
	}
	if distributeListImport6 != nil && prefixListImport6 != nil {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast distribute-list import", *distributeListImport6)
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh,
			"address-family ipv6-unicast prefix-list import", *prefixListImport6)
		fmt.Println("you can't set both a prefix-list and a distribute list.")
		return false
	}
	// ebgp-multihop & ttl-security hops
	if ebgpMultihop != nil && ttlSecurityHops != nil {
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "ebgp-multihop", *ebgpMultihop)
		fmt.Println("protocols bgp", *asNum, "neighbor", neigh, "ttl-security hops", *ttlSecurityHops)
		fmt.Println("you can't set both ebgp-multihop and ttl-security hops.")
		return false
	}
	return true
}

func quaggaConfigValidBgpNeighbors(asNum *string) bool {
	if asNum == nil {
		return false
	}
	neighs := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "neighbor"})
	for _, neigh := range neighs {
		if !quaggaConfigValidBgpNeighbor(asNum, neigh) {
			return false
		}
	}
	return true
}

func quaggaConfigValidBgpPeerGroup(asNum *string, peerGroup string) bool {
	remoteAs := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "remote-as"})
	if remoteAs == nil {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup)
		fmt.Println("remote-as required.")
		return false
	}
	if !validatorRange(*remoteAs, 1, 4294967294) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "remote-as", *remoteAs)
		fmt.Println("AS number must be between 1 and 4294967294.")
		return false
	}
	// allowas-in number [2]
	allowasInNumber := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "allowas-in", "number"})
	if allowasInNumber != nil && !validatorRange(*allowasInNumber, 1, 10) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "allowas-in number", *allowasInNumber)
		fmt.Println("allowas-in number must be between 1 and 10.")
		return false
	}
	allowasInNumber6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "allowas-in", "number"})
	if allowasInNumber6 != nil && !validatorRange(*allowasInNumber6, 1, 10) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast allowas-in number", *allowasInNumber6)
		fmt.Println("allowas-in number must be between 1 and 10.")
		return false
	}

	// default-originate route-map [2]
	defaultOriginateRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "default-originate", "route-map"})
	if defaultOriginateRouteMap != nil && !validatorExistsRouteMap(*defaultOriginateRouteMap) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"default-originate route-map", *defaultOriginateRouteMap)
		fmt.Println("route-map ", *defaultOriginateRouteMap, " doesn't exist.")
		return false
	}
	defaultOriginateRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "default-originate", "route-map"})
	if defaultOriginateRouteMap6 != nil && !validatorExistsRouteMap(*defaultOriginateRouteMap6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast default-originate route-map", *defaultOriginateRouteMap6)
		fmt.Println("route-map ", *defaultOriginateRouteMap6, " doesn't exist.")
		return false
	}
	// distribute-list export [2]
	distributeListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "distribute-list", "export"})
	if distributeListExport != nil && !validatorExistsAccessList(*distributeListExport) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"distribute-list export", *distributeListExport)
		fmt.Println("access-list ", *distributeListExport, " doesn't exist")
		return false
	}
	distributeListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "distribute-list", "export"})
	if distributeListExport6 != nil && !validatorExistsAccessList6(*distributeListExport6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast distribute-list export", *distributeListExport6)
		fmt.Println("access-list6 ", *distributeListExport6, " doesn't exist")
		return false
	}
	// distribute-list import [2]
	distributeListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "distribute-list", "import"})
	if distributeListImport != nil && !validatorExistsAccessList(*distributeListImport) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"distribute-list import", *distributeListImport)
		fmt.Println("access-list ", *distributeListImport, " doesn't exist")
		return false
	}
	distributeListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "distribute-list", "import"})
	if distributeListImport6 != nil && !validatorExistsAccessList6(*distributeListImport6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast distribute-list import", *distributeListImport6)
		fmt.Println("access-list6 ", *distributeListImport6, " doesn't exist")
		return false
	}
	// ebgp-multihop
	ebgpMultihop := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "ebgp-multihop"})
	if ebgpMultihop != nil && !validatorRange(*ebgpMultihop, 1, 255) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "ebgp-multihop", *ebgpMultihop)
		fmt.Println("ebgp-multihop must be between 1 and 255.")
		return false
	}
	// filter-list export [2]
	filterListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "filter-list", "export"})
	if filterListExport != nil && !validatorExistsAsPathList(*filterListExport) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "filter-list export", *filterListExport)
		fmt.Println("as-path-list ", *filterListExport, " doesn't exist.")
		return false
	}
	filterListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "filter-list", "export"})
	if filterListExport6 != nil && !validatorExistsAsPathList(*filterListExport6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast filter-list export", *filterListExport6)
		fmt.Println("as-path-list ", *filterListExport6, " doesn't exist.")
		return false
	}
	// filter-list import [2]
	filterListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "filter-list", "import"})
	if filterListImport != nil && !validatorExistsAsPathList(*filterListImport) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "filter-list import", *filterListImport)
		fmt.Println("as-path-list ", *filterListImport, " doesn't exist.")
		return false
	}
	filterListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "filter-list", "import"})
	if filterListImport6 != nil && !validatorExistsAsPathList(*filterListImport6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast filter-list import", *filterListImport6)
		fmt.Println("as-path-list ", *filterListImport6, " doesn't exist.")
		return false
	}
	// local-as
	localAs := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "local-as"})
	if localAs != nil && !validatorRange(*localAs, 1, 4294967294) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "local-as", *localAs)
		fmt.Println("local-as must be between 1 and 4294967294.")
		return false
	}
	if localAs != nil && *localAs == *asNum {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "local-as", *localAs)
		fmt.Println("you can't set local-as the same as the router AS.")
		return false
	}
	// maximum-prefix [2]
	maximumPrefix := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "maximum-prefix"})
	if maximumPrefix != nil && !validatorRange(*maximumPrefix, 1, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "maximum-prefix", *maximumPrefix)
		fmt.Println("maximum-prefix must be between 1 and 4294967295.")
		return false
	}
	maximumPrefix6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "maximum-prefix"})
	if maximumPrefix6 != nil && !validatorRange(*maximumPrefix6, 1, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast maximum-prefix", *maximumPrefix6)
		fmt.Println("maximum-prefix must be between 1 and 4294967295.")
		return false
	}
	// password
	password := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "password"})
	if password != nil && len(*password) < 1 && len(*password) > 80 {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "password")
		fmt.Println("password must be 80 characters or less.")
		return false
	}
	// prefix-list export [2]
	prefixListExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "prefix-list", "export"})
	if prefixListExport != nil && !validatorExistsPrefixList(*prefixListExport) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "prefix-list export", *prefixListExport)
		fmt.Println("prefix-list ", *prefixListExport, " doesn't exist.")
		return false
	}
	prefixListExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "prefix-list", "export"})
	if prefixListExport6 != nil && !validatorExistsPrefixList6(*prefixListExport6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast prefix-list export", *prefixListExport6)
		fmt.Println("prefix-list6 ", *prefixListExport6, " doesn't exist.")
		return false
	}
	// prefix-list import [2]
	prefixListImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "prefix-list", "import"})
	if prefixListImport != nil && !validatorExistsPrefixList(*prefixListImport) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "prefix-list import", *prefixListImport)
		fmt.Println("prefix-list ", *prefixListImport, " doesn't exist.")
		return false
	}
	prefixListImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "prefix-list", "import"})
	if prefixListImport6 != nil && !validatorExistsPrefixList6(*prefixListImport6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast prefix-list import", *prefixListImport6)
		fmt.Println("prefix-list6 ", *prefixListImport6, " doesn't exist.")
		return false
	}
	// route-map export [2]
	routeMapExport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "route-map", "export"})
	if routeMapExport != nil && !validatorExistsRouteMap(*routeMapExport) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "route-map export", *routeMapExport)
		fmt.Println("route-map ", *routeMapExport, " doesn't exist.")
		return false
	}
	routeMapExport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "route-map", "export"})
	if routeMapExport6 != nil && !validatorExistsRouteMap(*routeMapExport6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast route-map export", *routeMapExport6)
		fmt.Println("route-map ", *routeMapExport6, " doesn't exist.")
		return false
	}
	// route-map import [2]
	routeMapImport := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "route-map", "import"})
	if routeMapImport != nil && !validatorExistsRouteMap(*routeMapImport) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "route-map import", *routeMapImport)
		fmt.Println("route-map ", *routeMapImport, " doesn't exist.")
		return false
	}
	routeMapImport6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "route-map", "import"})
	if routeMapImport6 != nil && !validatorExistsRouteMap(*routeMapImport6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast route-map import", *routeMapImport6)
		fmt.Println("route-map ", *routeMapImport6, " doesn't exist.")
		return false
	}
	// ttl-security hops
	ttlSecurityHops := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "ttl-security", "hops"})
	if ttlSecurityHops != nil && !validatorRange(*ttlSecurityHops, 1, 254) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "ttl-security hops", *ttlSecurityHops)
		fmt.Println("ttl-security hops must be between 1 and 254.")
		return false
	}
	// unsuppress-map [2]
	unsuppressMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "unsuppress-map"})
	if unsuppressMap != nil && !validatorExistsRouteMap(*unsuppressMap) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "unsuppress-map", *unsuppressMap)
		fmt.Println("route-map ", *unsuppressMap, " doesn't exist.")
		return false
	}
	unsuppressMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup,
			"address-family", "ipv6-unicast", "unsuppress-map"})
	if unsuppressMap6 != nil && !validatorExistsRouteMap(*unsuppressMap6) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast unsuppress-map", *unsuppressMap6)
		fmt.Println("route-map ", *unsuppressMap6, " doesn't exist.")
		return false
	}
	// update-source
	updateSource := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "update-source"})
	if updateSource != nil && !validatorSource(*updateSource) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "update-source", *updateSource)
		fmt.Println("update-source format error.")
		return false
	}
	// weight
	weight := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "peer-group", peerGroup, "weight"})
	if weight != nil && !validatorRange(*weight, 1, 65535) {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "weight", *weight)
		fmt.Println("weight must be between 1 and 65535.")
		return false
	}
	// distribute-list export & prefix-list export [2]
	if distributeListExport != nil && prefixListExport != nil {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"distribute-list export", *distributeListExport)
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "prefix-list export", *prefixListExport)
		fmt.Println("you can't set both a prefix-list and a distribute list.")
		return false
	}
	if distributeListExport6 != nil && prefixListExport6 != nil {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast distribute-list export", *distributeListExport6)
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast prefix-list export", *prefixListExport6)
		fmt.Println("you can't set both a prefix-list and a distribute list.")
		return false
	}
	// distribute-list import & prefix-list import [2]
	if distributeListImport != nil && prefixListImport != nil {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"distribute-list import", *distributeListImport)
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "prefix-list import", *prefixListImport)
		fmt.Println("you can't set both a prefix-list and a distribute list.")
		return false
	}
	if distributeListImport6 != nil && prefixListImport6 != nil {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast distribute-list import", *distributeListImport6)
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup,
			"address-family ipv6-unicast prefix-list import", *prefixListImport6)
		fmt.Println("you can't set both a prefix-list and a distribute list.")
		return false
	}
	// ebgp-multihop & ttl-security hops
	if ebgpMultihop != nil && ttlSecurityHops != nil {
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "ebgp-multihop", *ebgpMultihop)
		fmt.Println("protocols bgp", *asNum, "peer-group", peerGroup, "ttl-security hops", *ttlSecurityHops)
		fmt.Println("you can't set both ebgp-multihop and ttl-security hops.")
		return false
	}
	return true
}

func quaggaConfigValidBgpPeerGroups(asNum *string) bool {
	if asNum == nil {
		return false
	}
	peerGroups := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "peer-group"})
	for _, peerGroup := range peerGroups {
		if !quaggaConfigValidBgpPeerGroup(asNum, peerGroup) {
			return false
		}
	}
	return true
}

func quaggaConfigValidBgpAggregateAddresses(asNum *string) bool {
	if asNum == nil {
		return false
	}
	aggregateAddresses := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "aggregate-address"})
	for _, aggregateAddress := range aggregateAddresses {
		if !validatorIPv4Address(aggregateAddress) {
			return false
		}
	}
	aggregateAddress6s := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast", "aggregate-address"})
	for _, aggregateAddress6 := range aggregateAddress6s {
		if !validatorIPv6Address(aggregateAddress6) {
			return false
		}
	}
	return true
}

func quaggaConfigValidBgpNetworks(asNum *string) bool {
	if asNum == nil {
		return false
	}
	networks := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "network"})
	for _, network := range networks {
		if !validatorIPv4CIDR(network) {
			return false
		}
		routeMap := configCandidate.value(
			[]string{"protocols", "bgp", *asNum, "network", network, "route-map"})
		if routeMap != nil && !validatorExistsRouteMap(*routeMap) {
			return false
		}
	}
	network6s := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast", "network"})
	for _, network6 := range network6s {
		if !validatorIPv6CIDR(network6) {
			return false
		}
		routeMap6 := configCandidate.value(
			[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
				"network", network6, "route-map"})
		if routeMap6 != nil && !validatorExistsRouteMap(*routeMap6) {
			return false
		}
	}
	return true
}

func quaggaConfigValidBgpParameterses(asNum *string) bool {
	if asNum == nil {
		return false
	}
	//protocols bgp WORD parameters cluster-id A.B.C.D
	clusterId := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "cluster-id"})
	if clusterId != nil && !validatorIPv4Address(*clusterId) {
		fmt.Println("protocols bgp", *asNum, "parameters cluster-id", *clusterId)
		fmt.Println("cluster-id format error.")
		return false
	}
	//protocols bgp WORD parameters confederation identifier WORD
	confederationIdentifier := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "confederation", "identifier"})
	if confederationIdentifier != nil && !validatorRange(*confederationIdentifier, 1, 4294967294) {
		fmt.Println("protocols bgp", *asNum, "parameters confederation identifier", *confederationIdentifier)
		fmt.Println("confederation AS id must be between 1 and 4294967294.")
		return false
	}
	//protocols bgp WORD parameters confederation peers WORD
	confederationPeers := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "confederation", "peers"})
	if confederationPeers != nil && !validatorRange(*confederationPeers, 1, 4294967294) {
		fmt.Println("protocols bgp", *asNum, "parameters confederation peers", *confederationPeers)
		fmt.Println("confederation AS id must be between 1 and 4294967294.")
		return false
	}
	if confederationPeers != nil && !validatorConfediBGPASNCheck(*confederationPeers, *asNum) {
		fmt.Println("protocols bgp", *asNum, "parameters confederation peers", *confederationPeers)
		fmt.Println("Can't set confederation peers ASN to ", *confederationPeers, ".")
		fmt.Println("Delete any neighbors with remote-as ", *confederationPeers,
			" and/or change the local ASN first.")
		return false
	}
	//protocols bgp WORD parameters dampening half-life WORD
	dampeningHalfLife := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "dampening", "half-life"})
	if dampeningHalfLife != nil && !validatorRange(*dampeningHalfLife, 1, 45) {
		fmt.Println("protocols bgp", *asNum, "parameters dampening half-life", *dampeningHalfLife)
		fmt.Println("Half-life penalty must be between 1 and 45.")
		return false
	}
	//protocols bgp WORD parameters dampening max-suppress-time WORD
	dampeningMaxSuppressTime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "dampening", "max-suppress-time"})
	if dampeningMaxSuppressTime != nil && !validatorRange(*dampeningMaxSuppressTime, 1, 255) {
		fmt.Println("protocols bgp", *asNum,
			"parameters dampening max-suppress-time", *dampeningMaxSuppressTime)
		fmt.Println("Max-suppress-time must be between 1 and 255.")
		return false
	}
	//protocols bgp WORD parameters dampening re-use WORD
	dampeningReUse := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "dampening", "re-use"})
	if dampeningReUse != nil && !validatorRange(*dampeningReUse, 1, 20000) {
		fmt.Println("protocols bgp", *asNum, "parameters dampening re-use", *dampeningReUse)
		fmt.Println("Re-use value must be between 1 and 20000.")
		return false
	}
	//protocols bgp WORD parameters dampening start-suppress-time WORD
	dampeningStartSuppressTime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "dampening", "start-suppress-time"})
	if dampeningStartSuppressTime != nil && !validatorRange(*dampeningStartSuppressTime, 1, 20000) {
		fmt.Println("protocols bgp", *asNum,
			"parameters dampening start-suppress-time", *dampeningStartSuppressTime)
		fmt.Println("Start-suppress-time must be between 1 and 20000.")
		return false
	}
	//protocols bgp WORD parameters dampening
	if (dampeningHalfLife != nil || dampeningMaxSuppressTime != nil ||
		dampeningReUse != nil || dampeningStartSuppressTime != nil) &&
		(dampeningHalfLife == nil || dampeningMaxSuppressTime == nil ||
			dampeningReUse == nil || dampeningStartSuppressTime == nil) {
		fmt.Println("protocols bgp", *asNum, "parameters dampening")
		fmt.Println("you must set a half-life, max-suppress-time, re-use, start-suppress-time.")
		return false
	}
	//protocols bgp WORD parameters default local-pref WORD
	defaultLocalPref := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "default", "local-pref"})
	if defaultLocalPref != nil && !validatorRange(*defaultLocalPref, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "parameters default local-pref", *defaultLocalPref)
		fmt.Println("default local-pref format error.")
		return false
	}
	//protocols bgp WORD parameters distance global external WORD
	distanceGlobalExternal := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "distance", "global", "external"})
	if distanceGlobalExternal != nil && !validatorRange(*distanceGlobalExternal, 1, 255) {
		fmt.Println("protocols bgp", *asNum, "parameters distance global external", *distanceGlobalExternal)
		fmt.Println("Must be between 1-255")
		return false
	}
	//protocols bgp WORD parameters distance global internal WORD
	distanceGlobalInternal := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "distance", "global", "internal"})
	if distanceGlobalInternal != nil && !validatorRange(*distanceGlobalInternal, 1, 255) {
		fmt.Println("protocols bgp", *asNum, "parameters distance global internal", *distanceGlobalInternal)
		fmt.Println("Must be between 1-255")
		return false
	}
	//protocols bgp WORD parameters distance global local WORD
	distanceGlobalLocal := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "distance", "global", "local"})
	if distanceGlobalLocal != nil && !validatorRange(*distanceGlobalLocal, 1, 255) {
		fmt.Println("protocols bgp", *asNum, "parameters distance global local", *distanceGlobalLocal)
		fmt.Println("Must be between 1-255")
		return false
	}
	//protocols bgp WORD parameters distance global
	if (distanceGlobalExternal != nil || distanceGlobalInternal != nil || distanceGlobalLocal != nil) &&
		(distanceGlobalExternal == nil || distanceGlobalInternal == nil || distanceGlobalLocal == nil) {
		fmt.Println("protocols bgp", *asNum, "parameters distance global")
		fmt.Println("you must set an external, internal, local distance.")
		return false
	}
	//protocols bgp WORD parameters distance prefix A.B.C.D/M
	//protocols bgp WORD parameters distance prefix A.B.C.D/M distance WORD
	distancePrefixes := configCandidate.values(
		[]string{"protocols", "bgp", *asNum, "parameters", "distance", "prefix"})
	for _, distancePrefix := range distancePrefixes {
		if !validatorIPv4CIDR(distancePrefix) {
			fmt.Println("protocols bgp", *asNum, "parameters distance prefix", distancePrefix)
			fmt.Println("format error.")
			return false
		}
		distance := configCandidate.value(
			[]string{"protocols", "bgp", *asNum, "parameters", "distance", "prefix", distancePrefix,
				"distance"})
		if distance == nil || !validatorRange(*distance, 1, 255) {
			t := ""
			if distance != nil {
				t = *distance
			}
			fmt.Println("protocols bgp", *asNum, "parameters distance prefix", distancePrefix,
				"distance", t)
			fmt.Println("Must be between 1-255.")
			return false
		}
	}
	//protocols bgp WORD parameters graceful-restart stalepath-time WORD
	gracefulRestartStalepathTime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "graceful-restart", "stalepath-time"})
	if gracefulRestartStalepathTime != nil && !validatorRange(*gracefulRestartStalepathTime, 1, 3600) {
		fmt.Println("protocols bgp", *asNum, "parameters graceful-restart stalepath-time",
			*gracefulRestartStalepathTime)
		fmt.Println("stalepath-time must be between 1 and 3600.")
		return false
	}
	//protocols bgp WORD parameters router-id A.B.C.D
	routerId := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "router-id"})
	if routerId != nil && !validatorIPv4Address(*routerId) {
		fmt.Println("protocols bgp", *asNum, "parameters router-id", *routerId)
		fmt.Println("format error.")
		return false
	}
	//protocols bgp WORD parameters scan-time WORD
	scanTime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "parameters", "scan-time"})
	if scanTime != nil && !validatorRange(*scanTime, 5, 60) {
		fmt.Println("protocols bgp", *asNum, "parameters scan-time", *scanTime)
		fmt.Println("scan-time must be between 5 and 60 seconds.")
		return false
	}
	return true
}

func quaggaConfigValidBgpRedistributes(asNum *string) bool {
	if asNum == nil {
		return false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute connected metric WORD
	connectedMetric6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "connected", "metric"})
	if connectedMetric6 != nil && !validatorRange(*connectedMetric6, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute connected metric", *connectedMetric6)
		fmt.Println("metric format error.")
		return false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute connected route-map WORD
	connectedRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "connected", "route-map"})
	if connectedRouteMap6 != nil && !validatorExistsRouteMap(*connectedRouteMap6) {
		fmt.Println("protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute connected route-map", *connectedRouteMap6)
		fmt.Println("route-map not found.")
		return false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute kernel metric WORD
	kernelMetric6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "kernel", "metric"})
	if kernelMetric6 != nil && !validatorRange(*kernelMetric6, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute kernel metric", *kernelMetric6)
		fmt.Println("metric format error.")
		return false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute kernel route-map WORD
	kernelRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "kernel", "route-map"})
	if kernelRouteMap6 != nil && !validatorExistsRouteMap(*kernelRouteMap6) {
		fmt.Println("protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute kernel route-map", *kernelRouteMap6)
		fmt.Println("route-map not found.")
		return false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute ospfv3 metric WORD
	ospfv3Metric6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "ospfv3", "metric"})
	if ospfv3Metric6 != nil && !validatorRange(*ospfv3Metric6, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute ospfv3 metric", *ospfv3Metric6)
		fmt.Println("metric format error.")
		return false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute ospfv3 route-map WORD
	ospfv3RouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "ospfv3", "route-map"})
	if ospfv3RouteMap6 != nil && !validatorExistsRouteMap(*ospfv3RouteMap6) {
		fmt.Println("protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute ospfv3 route-map", *ospfv3RouteMap6)
		fmt.Println("route-map not found.")
		return false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute ripng metric WORD
	//protocols bgp WORD address-family ipv6-unicast redistribute ripng route-map WORD
	//protocols bgp WORD address-family ipv6-unicast redistribute static metric WORD
	staticMetric6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "static", "metric"})
	if staticMetric6 != nil && !validatorRange(*staticMetric6, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute static metric", *staticMetric6)
		fmt.Println("metric format error.")
		return false
	}
	//protocols bgp WORD address-family ipv6-unicast redistribute static route-map WORD
	staticRouteMap6 := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "address-family", "ipv6-unicast",
			"redistribute", "static", "route-map"})
	if staticRouteMap6 != nil && !validatorExistsRouteMap(*staticRouteMap6) {
		fmt.Println("protocols bgp", *asNum,
			"address-family ipv6-unicast redistribute static route-map", *staticRouteMap6)
		fmt.Println("route-map not found.")
		return false
	}
	//protocols bgp WORD redistribute connected metric WORD
	connectedMetric := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "connected", "metric"})
	if connectedMetric != nil && !validatorRange(*connectedMetric, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "redistribute connected metric", *connectedMetric)
		fmt.Println("metric format error.")
		return false
	}
	//protocols bgp WORD redistribute connected route-map WORD
	connectedRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "connected", "route-map"})
	if connectedRouteMap != nil && !validatorExistsRouteMap(*connectedRouteMap) {
		fmt.Println("protocols bgp", *asNum, "redistribute connected route-map", *connectedRouteMap)
		fmt.Println("route-map not found.")
		return false
	}
	//protocols bgp WORD redistribute kernel metric WORD
	kernelMetric := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "kernel", "metric"})
	if kernelMetric != nil && !validatorRange(*kernelMetric, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "redistribute kernel metric", *kernelMetric)
		fmt.Println("metric format error.")
		return false
	}
	//protocols bgp WORD redistribute kernel route-map WORD
	kernelRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "kernel", "route-map"})
	if kernelRouteMap != nil && !validatorExistsRouteMap(*kernelRouteMap) {
		fmt.Println("protocols bgp", *asNum, "redistribute kernel route-map", *kernelRouteMap)
		fmt.Println("route-map not found.")
		return false
	}
	//protocols bgp WORD redistribute ospf metric WORD
	ospfMetric := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "ospf", "metric"})
	if ospfMetric != nil && !validatorRange(*ospfMetric, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "redistribute ospf metric", *ospfMetric)
		fmt.Println("metric format error.")
		return false
	}
	//protocols bgp WORD redistribute ospf route-map WORD
	ospfRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "ospf", "route-map"})
	if ospfRouteMap != nil && !validatorExistsRouteMap(*ospfRouteMap) {
		fmt.Println("protocols bgp", *asNum, "redistribute ospf route-map", *ospfRouteMap)
		fmt.Println("route-map not found.")
		return false
	}
	//protocols bgp WORD redistribute rip metric WORD
	//protocols bgp WORD redistribute rip route-map WORD
	//protocols bgp WORD redistribute static metric WORD
	staticMetric := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "static", "metric"})
	if staticMetric != nil && !validatorRange(*staticMetric, 0, 4294967295) {
		fmt.Println("protocols bgp", *asNum, "redistribute static metric", *staticMetric)
		fmt.Println("metric format error.")
		return false
	}
	//protocols bgp WORD redistribute static route-map WORD
	staticRouteMap := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "redistribute", "static", "route-map"})
	if staticRouteMap != nil && !validatorExistsRouteMap(*staticRouteMap) {
		fmt.Println("protocols bgp", *asNum, "redistribute static route-map", *staticRouteMap)
		fmt.Println("route-map not found.")
		return false
	}
	return true
}

func quaggaConfigValidBgpMaximumPathses(asNum *string) bool {
	if asNum == nil {
		return false
	}
	//protocols bgp WORD maximum-paths ebgp WORD
	ebgp := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "maximum-paths", "ebgp"})
	if ebgp != nil && !validatorRange(*ebgp, 1, 255) {
		fmt.Println("protocols bgp", *asNum, "maximum-paths ebgp", *ebgp)
		fmt.Println("Must be between (1-255).")
		return false
	}
	//protocols bgp WORD maximum-paths ibgp WORD
	ibgp := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "maximum-paths", "ibgp"})
	if ibgp != nil && !validatorRange(*ibgp, 1, 255) {
		fmt.Println("protocols bgp", *asNum, "maximum-paths ibgp", *ibgp)
		fmt.Println("Must be between (1-255).")
		return false
	}
	return true
}

func quaggaConfigValidBgpTimerses(asNum *string) bool {
	if asNum == nil {
		return false
	}
	//protocols bgp WORD timers holdtime WORD
	holdtime := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "timers", "holdtime"})
	if holdtime != nil && !validatorRange(*holdtime, 0, 0) && !validatorRange(*holdtime, 4, 65535) {
		fmt.Println("protocols bgp", *asNum, "timers holdtime", *holdtime)
		fmt.Println("hold-time interval must be 0 or between 4 and 65535.")
		return false
	}
	//protocols bgp WORD timers keepalive WORD
	keepalive := configCandidate.value(
		[]string{"protocols", "bgp", *asNum, "timers", "keepalive"})
	if keepalive != nil && !validatorRange(*keepalive, 1, 65535) {
		fmt.Println("protocols bgp", *asNum, "timers keepalive", *keepalive)
		fmt.Println("Keepalive interval must be between 1 and 65535.")
		return false
	}
	return true
}

func quaggaConfigValidBgp() bool {
	asNums := configCandidate.values([]string{"protocols", "bgp"})
	if len(asNums) == 0 {
		return true
	}
	if len(asNums) > 1 {
		for _, asNum := range asNums {
			fmt.Println("protocols bgp", asNum)
		}
		fmt.Println("Multiple BGP can not be set.")
		return false
	}
	asNum := configCandidate.value([]string{"protocols", "bgp"})
	if !quaggaConfigValidBgpAsNum(asNum) {
		return false
	}
	if !quaggaConfigValidBgpNeighbors(asNum) {
		return false
	}
	if !quaggaConfigValidBgpPeerGroups(asNum) {
		return false
	}
	if !quaggaConfigValidBgpAggregateAddresses(asNum) {
		return false
	}
	if !quaggaConfigValidBgpNetworks(asNum) {
		return false
	}
	if !quaggaConfigValidBgpParameterses(asNum) {
		return false
	}
	if !quaggaConfigValidBgpRedistributes(asNum) {
		return false
	}
	if !quaggaConfigValidBgpMaximumPathses(asNum) {
		return false
	}
	if !quaggaConfigValidBgpTimerses(asNum) {
		return false
	}
	return true
}
