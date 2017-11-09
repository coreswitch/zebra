package quagga

import (
	"fmt"
	"io"
)

func quaggaConfigValidOspfAccessList(f io.Writer, accessList string) bool {
	//protocols ospf access-list WORD
	valid := true
	exports := configCandidate.values(
		[]string{"protocols", "ospf", "access-list", accessList, "export"})
	if len(exports) == 0 {
		fmt.Fprintln(f, "protocols ospf access-list", accessList)
		fmt.Fprintln(f, "must add protocol to filter.")
		valid = false
	}
	for _, export := range exports {
		if !validatorInclude(export, []string{"bgp", "connected", "kernel", "rip", "static"}) {
			fmt.Fprintln(f, "protocols ospf access-list", accessList, "export", export)
			fmt.Fprintln(f, "Must be (bgp, connected, kernel, rip, or static).")
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidOspfAccessLists(f io.Writer) bool {
	valid := true
	accessLists := configCandidate.values(
		[]string{"protocols", "ospf", "access-list"})
	for _, accessList := range accessLists {
		if !quaggaConfigValidOspfAccessList(f, accessList) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidOspfAreaNetwork(f io.Writer, area, network string) bool {
	//protocols ospf area WORD network A.B.C.D/M
	valid := true
	if !validatorIPv4CIDR(network) {
		fmt.Fprintln(f, "protocols ospf area", area, "network", network)
		fmt.Fprintln(f, "format error.")
		valid = false
	}
	return valid
}

func quaggaConfigValidOspfAreaRange(f io.Writer, area, range_ string) bool {
	//protocols ospf area WORD range A.B.C.D/M
	valid := true
	if !validatorIPv4CIDR(range_) {
		fmt.Fprintln(f, "protocols ospf area", area, "range", range_)
		fmt.Fprintln(f, "format error.")
		valid = false
		return valid
	}
	cost := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "range", range_, "cost"})
	if cost != nil && !validatorRange(*cost, 0, 16777215) {
		fmt.Fprintln(f, "protocols ospf area", area, "range", range_, "cost", *cost)
		fmt.Fprintln(f, "Metric must be between 0-16777215.")
		valid = false
	}
	substitute := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "range", range_, "substitute"})
	if substitute != nil && !validatorIPv4CIDR(*substitute) {
		fmt.Fprintln(f, "protocols ospf area", area, "range", range_, "substitute", *substitute)
		fmt.Fprintln(f, "format error.")
		valid = false
	}
	notAdvertise := configCandidate.lookup(
		[]string{"protocols", "ospf", "area", area, "range", range_, "not-advertise"}) != nil
	if notAdvertise && (cost != nil || substitute != nil) {
		if cost != nil {
			fmt.Fprintln(f, "protocols ospf area", area, "range", range_, "cost", *cost)
		}
		if substitute != nil {
			fmt.Fprintln(f, "protocols ospf area", area, "range", range_, "substitute", *substitute)
		}
		fmt.Fprintln(f, "protocols ospf area", area, "range", range_, "not-advertise")
		fmt.Fprintln(f, "Remove 'not-advertise' before setting cost or substitue.")
		valid = false
	}
	return valid
}

func quaggaConfigValidOspfAreaVirtualLink(f io.Writer, area, virtualLink string) bool {
	//protocols ospf area WORD virtual-link A.B.C.D
	valid := true
	if !validatorIPv4Address(virtualLink) {
		fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink)
		fmt.Fprintln(f, "format error.")
		valid = false
		return valid
	}
	if area == "0" || area == "0.0.0.0" {
		fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink)
		fmt.Fprintln(f, "Can't configure VL over area", area, ".")
		valid = false
		return valid
	}
	keyIds := configCandidate.values(
		[]string{"protocols", "ospf", "area", area, "virtual-link", virtualLink,
			"authentication", "md5", "key-id"})
	for _, keyId := range keyIds {
		if !validatorRange(keyId, 1, 255) {
			fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
				" authentication md5 key-id ", keyId)
			fmt.Fprintln(f, "ID must be between (1-255).")
			valid = false
		}
		md5Mey := configCandidate.value(
			[]string{"protocols", "ospf", "area", area, "virtual-link", virtualLink,
				"authentication", "md5", "key-id", keyId, "md5-key"})
		if md5Mey == nil {
			fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
				"authentication md5 key-id", keyId)
			fmt.Fprintln(f, "Must add the md5-key for key-id", keyId, ".")
			valid = false
		} else if len(*md5Mey) > 16 {
			fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
				"authentication md5 key-id", keyId, "md5-key", *md5Mey)
			fmt.Fprintln(f, "MD5 key must be 16 characters or less.")
			valid = false
		}
	}
	plaintextPassword := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "virtual-link", virtualLink,
			"authentication", "plaintext-password"})
	if plaintextPassword != nil && len(*plaintextPassword) > 8 {
		fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
			"authentication plaintext-password", plaintextPassword)
		fmt.Fprintln(f, "Password must be 8 characters or less.")
		valid = false
	}
	if len(keyIds) > 0 && plaintextPassword != nil {
		for _, keyId := range keyIds {
			fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
				"authentication md5 key-id", keyId)
		}
		fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
			"authentication plaintext-password", plaintextPassword)
		fmt.Fprintln(f, "authentication already set.")
		valid = false
	}
	deadInterval := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "virtual-link", virtualLink, "dead-interval"})
	if deadInterval != nil && !validatorRange(*deadInterval, 1, 65535) {
		fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
			"dead-interval", *deadInterval)
		fmt.Fprintln(f, "Must be between 1-65535.")
		valid = false
	}
	helloInterval := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "virtual-link", virtualLink, "hello-interval"})
	if helloInterval != nil && !validatorRange(*helloInterval, 1, 65535) {
		fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
			"hello-interval", *helloInterval)
		fmt.Fprintln(f, "Must be between 1-65535.")
		valid = false
	}
	retransmitInterval := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "virtual-link", virtualLink, "retransmit-interval"})
	if retransmitInterval != nil && !validatorRange(*retransmitInterval, 1, 65535) {
		fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
			"retransmit-interval", *retransmitInterval)
		fmt.Fprintln(f, "Must be between 1-65535.")
		valid = false
	}
	transmitDelay := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "virtual-link", virtualLink, "transmit-delay"})
	if transmitDelay != nil && !validatorRange(*transmitDelay, 1, 65535) {
		fmt.Fprintln(f, "protocols ospf area", area, "virtual-link", virtualLink,
			"transmit-delay", *transmitDelay)
		fmt.Fprintln(f, "Must be between 1-65535.")
		valid = false
	}
	return valid
}

func quaggaConfigValidOspfArea(f io.Writer, area string) bool {
	//protocols ospf area WORD
	valid := true
	if !validatorIPv4Address(area) && !validatorRange(area, 0, 4294967295) {
		fmt.Fprintln(f, "protocols ospf area", area)
		fmt.Fprintln(f, "format error.")
		valid = false
		return valid
	}
	areaTypeNormal := configCandidate.lookup(
		[]string{"protocols", "ospf", "area", area, "area-type", "normal"}) != nil
	areaTypeNssa := configCandidate.lookup(
		[]string{"protocols", "ospf", "area", area, "area-type", "nssa"}) != nil
	areaTypeStub := configCandidate.lookup(
		[]string{"protocols", "ospf", "area", area, "area-type", "stub"}) != nil
	areaTypeCount := 0
	if areaTypeNormal {
		areaTypeCount++
	}
	if areaTypeNssa {
		areaTypeCount++
	}
	if areaTypeStub {
		areaTypeCount++
	}
	if areaTypeCount > 1 {
		fmt.Fprintln(f, "protocols ospf area", area)
		fmt.Fprintln(f, "you may only define one area type.")
		valid = false
	}
	areaTypeNssaDefaultCost := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "area-type", "nssa", "default-cost"})
	if areaTypeNssaDefaultCost != nil && !validatorRange(*areaTypeNssaDefaultCost, 0, 16777215) {
		fmt.Fprintln(f, "protocols ospf area", area,
			"area-type nssa default-cost", *areaTypeNssaDefaultCost)
		fmt.Fprintln(f, "Cost must be between 0-16777215.")
		valid = false
	}
	areaTypeNssaTranslate := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "area-type", "nssa", "translate"})
	if areaTypeNssaTranslate != nil && !validatorInclude(*areaTypeNssaTranslate,
		[]string{"always", "candidate", "never"}) {
		fmt.Fprintln(f, "protocols ospf area", area,
			"area-type nssa translate", *areaTypeNssaTranslate)
		fmt.Fprintln(f, "Must be (always, candidate, or never).")
		valid = false
	}
	areaTypeStubDefaultCost := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "area-type", "stub", "default-cost"})
	if areaTypeStubDefaultCost != nil && !validatorRange(*areaTypeStubDefaultCost, 0, 16777215) {
		fmt.Fprintln(f, "protocols ospf area", area,
			"area-type stub default-cost", *areaTypeStubDefaultCost)
		fmt.Fprintln(f, "Cost must be between 0-16777215.")
		valid = false
	}
	authentication := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "authentication"})
	if authentication != nil && !validatorInclude(*authentication, []string{"plaintext-password", "md5"}) {
		fmt.Fprintln(f, "protocols ospf area", area, "authentication", *authentication)
		fmt.Fprintln(f, "Must be either plaintext-password or md5.")
		valid = false
	}
	networks := configCandidate.values(
		[]string{"protocols", "ospf", "area", area, "network"})
	for _, network := range networks {
		if !quaggaConfigValidOspfAreaNetwork(f, area, network) {
			valid = false
		}
	}
	ranges := configCandidate.values(
		[]string{"protocols", "ospf", "area", area, "range"})
	for _, range_ := range ranges {
		if !quaggaConfigValidOspfAreaRange(f, area, range_) {
			valid = false
		}
	}
	shortcut := configCandidate.value(
		[]string{"protocols", "ospf", "area", area, "shortcut"})
	if shortcut != nil && !validatorInclude(*shortcut, []string{"default", "disable", "enable"}) {
		fmt.Fprintln(f, "protocols ospf area", area, "shortcut", *shortcut)
		fmt.Fprintln(f, "Must be (default, disable, enable).")
		valid = false
	}
	virtualLinks := configCandidate.values(
		[]string{"protocols", "ospf", "area", area, "virtual-link"})
	for _, virtualLink := range virtualLinks {
		if !quaggaConfigValidOspfAreaVirtualLink(f, area, virtualLink) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidOspfAreas(f io.Writer) bool {
	valid := true
	areas := configCandidate.values(
		[]string{"protocols", "ospf", "area"})
	for _, area := range areas {
		if !quaggaConfigValidOspfArea(f, area) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidOspfNeighbor(f io.Writer, neighbor string) bool {
	//protocols ospf neighbor A.B.C.D
	valid := true
	if !validatorIPv4Address(neighbor) {
		fmt.Fprintln(f, "protocols ospf neighbor", neighbor)
		fmt.Fprintln(f, "format error.")
		valid = false
		return valid
	}
	pollInterval := configCandidate.value(
		[]string{"protocols", "ospf", "neighbor", neighbor, "poll-interval"})
	if pollInterval != nil && !validatorRange(*pollInterval, 1, 65535) {
		fmt.Fprintln(f, "protocols ospf neighbor", neighbor, "poll-interval", *pollInterval)
		fmt.Fprintln(f, "Must be between 1-65535 seconds.")
		valid = false
	}
	priority := configCandidate.value(
		[]string{"protocols", "ospf", "neighbor", neighbor, "priority"})
	if priority != nil && !validatorRange(*priority, 0, 255) {
		fmt.Fprintln(f, "protocols ospf neighbor", neighbor, "priority", *priority)
		fmt.Fprintln(f, "Priority must be between 0-255.")
		valid = false
	}
	return valid
}

func quaggaConfigValidOspfNeighbors(f io.Writer) bool {
	valid := true
	neighbors := configCandidate.values(
		[]string{"protocols", "ospf", "neighbor"})
	for _, neighbor := range neighbors {
		if !quaggaConfigValidOspfNeighbor(f, neighbor) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidOspf(f io.Writer) bool {
	valid := true
	if !quaggaConfigValidOspfAccessLists(f) {
		valid = false
	}
	if !quaggaConfigValidOspfAreas(f) {
		valid = false
	}
	autoCostReferenceBandwidth := configCandidate.value(
		[]string{"protocols", "ospf", "auto-cost", "reference-bandwidth"})
	if autoCostReferenceBandwidth != nil && !validatorRange(*autoCostReferenceBandwidth, 1, 4294967) {
		fmt.Fprintln(f, "protocols ospf auto-cost reference-bandwidth", *autoCostReferenceBandwidth)
		fmt.Fprintln(f, "Must be between 1-4294967.")
		valid = false
	}
	defaultInformationOriginateMetricType := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric-type"})
	if defaultInformationOriginateMetricType != nil &&
		!validatorInclude(*defaultInformationOriginateMetricType, []string{"1", "2"}) {
		fmt.Fprintln(f, "protocols ospf default-information originate metric-type",
			*defaultInformationOriginateMetricType)
		fmt.Fprintln(f, "metric must be either 1 or 2.")
		valid = false
	}
	defaultInformationOriginateMetric := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric"})
	if defaultInformationOriginateMetric != nil &&
		!validatorRange(*defaultInformationOriginateMetric, 0, 16777214) {
		fmt.Fprintln(f, "protocols ospf default-information originate metric", *defaultInformationOriginateMetric)
		fmt.Fprintln(f, "must be between 0-16777214.")
		valid = false
	}
	defaultInformationOriginateRouteMap := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "route-map"})
	if defaultInformationOriginateRouteMap != nil &&
		!validatorExistsRouteMap(*defaultInformationOriginateRouteMap) {
		fmt.Fprintln(f, "protocols ospf default-information originate route-map",
			*defaultInformationOriginateRouteMap)
		fmt.Fprintln(f, "route-map not found.")
		valid = false
	}
	defaultMetric := configCandidate.value(
		[]string{"protocols", "ospf", "default-metric"})
	if defaultMetric != nil && !validatorRange(*defaultMetric, 0, 16777214) {
		fmt.Fprintln(f, "protocols ospf default-metric", *defaultMetric)
		fmt.Fprintln(f, "Must be between 0-16777214.")
		valid = false
	}
	distanceGlobal := configCandidate.value(
		[]string{"protocols", "ospf", "distance", "global"})
	if distanceGlobal != nil && !validatorRange(*distanceGlobal, 1, 255) {
		fmt.Fprintln(f, "protocols ospf distance global", *distanceGlobal)
		fmt.Fprintln(f, "Must be between 1-255.")
		valid = false
	}
	distanceOspfExternal := configCandidate.value(
		[]string{"protocols", "ospf", "distance", "ospf", "external"})
	if distanceOspfExternal != nil && !validatorRange(*distanceOspfExternal, 1, 255) {
		fmt.Fprintln(f, "protocols ospf distance ospf external", *distanceOspfExternal)
		fmt.Fprintln(f, "Must be between 1-255.")
		valid = false
	}
	distanceOspfInterArea := configCandidate.value(
		[]string{"protocols", "ospf", "distance", "ospf", "inter-area"})
	if distanceOspfInterArea != nil && !validatorRange(*distanceOspfInterArea, 1, 255) {
		fmt.Fprintln(f, "protocols ospf distance ospf inter-area", *distanceOspfInterArea)
		fmt.Fprintln(f, "Must be between 1-255.")
		valid = false
	}
	distanceOspfIntraArea := configCandidate.value(
		[]string{"protocols", "ospf", "distance", "ospf", "intra-area"})
	if distanceOspfIntraArea != nil && !validatorRange(*distanceOspfIntraArea, 1, 255) {
		fmt.Fprintln(f, "protocols ospf distance ospf intra-area", *distanceOspfIntraArea)
		fmt.Fprintln(f, "Must be between 1-255.")
		valid = false
	}
	maxMetricRouterLsaOnShutdown := configCandidate.value(
		[]string{"protocols", "ospf", "max-metric", "router-lsa", "on-shutdown"})
	if maxMetricRouterLsaOnShutdown != nil && !validatorRange(*maxMetricRouterLsaOnShutdown, 5, 86400) {
		fmt.Fprintln(f, "protocols ospf max-metric router-lsa on-shutdown", *maxMetricRouterLsaOnShutdown)
		fmt.Fprintln(f, "must be between 5-86400 seconds.")
		valid = false
	}
	maxMetricRouterLsaOnStartup := configCandidate.value(
		[]string{"protocols", "ospf", "max-metric", "router-lsa", "on-startup"})
	if maxMetricRouterLsaOnStartup != nil && !validatorRange(*maxMetricRouterLsaOnStartup, 5, 86400) {
		fmt.Fprintln(f, "protocols ospf max-metric router-lsa on-startup", *maxMetricRouterLsaOnStartup)
		fmt.Fprintln(f, "must be between 5-86400 seconds.")
		valid = false
	}
	mplsTeRouterAddress := configCandidate.value(
		[]string{"protocols", "ospf", "mpls-te", "router-address"})
	if mplsTeRouterAddress != nil && !validatorIPv4Address(*mplsTeRouterAddress) {
		fmt.Fprintln(f, "protocols ospf mpls-te router-address", *mplsTeRouterAddress)
		fmt.Fprintln(f, "format error.")
		valid = false
	}
	if !quaggaConfigValidOspfNeighbors(f) {
		valid = false
	}
	parametersAbrType := configCandidate.value(
		[]string{"protocols", "ospf", "parameters", "abr-type"})
	if parametersAbrType != nil && !validatorInclude(*parametersAbrType,
		[]string{"cisco", "ibm", "shortcut", "standard"}) {
		fmt.Fprintln(f, "protocols ospf parameters abr-type", *parametersAbrType)
		fmt.Fprintln(f, "Must be (cisco, ibm, shortcut, standard).")
		valid = false
	}
	parametersRouterId := configCandidate.value(
		[]string{"protocols", "ospf", "parameters", "router-id"})
	if parametersRouterId != nil && !validatorIPv4Address(*parametersRouterId) {
		fmt.Fprintln(f, "protocols ospf parameters router-id", *parametersRouterId)
		fmt.Fprintln(f, "format error.")
		valid = false
	}
	// XXX:
	redistributeBgpMetricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric-type"})
	if redistributeBgpMetricType != nil && !validatorInclude(*redistributeBgpMetricType, []string{"1", "2"}) {
		fmt.Fprintln(f, "protocols ospf redistribute bgp metric-type", *redistributeBgpMetricType)
		fmt.Fprintln(f, "metric-type must be either 1 or 2.")
		valid = false
	}
	redistributeBgpMetric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric"})
	if redistributeBgpMetric != nil && !validatorRange(*redistributeBgpMetric, 1, 16) {
		fmt.Fprintln(f, "protocols ospf redistribute bgp metric", *redistributeBgpMetric)
		fmt.Fprintln(f, "metric must be between 1 and 16.")
		valid = false
	}
	redistributeBgpRouteMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "route-map"})
	if redistributeBgpRouteMap != nil && !validatorExistsRouteMap(*redistributeBgpRouteMap) {
		fmt.Fprintln(f, "protocols ospf redistribute bgp route-map", *redistributeBgpRouteMap)
		fmt.Fprintln(f, "route-map", *redistributeBgpRouteMap, "doesn't exist.")
		valid = false
	}
	// XXX:
	redistributeConnectedMetricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric-type"})
	if redistributeConnectedMetricType != nil &&
		!validatorInclude(*redistributeConnectedMetricType, []string{"1", "2"}) {
		fmt.Fprintln(f, "protocols ospf redistribute connected metric-type", *redistributeConnectedMetricType)
		fmt.Fprintln(f, "metric-type must be either 1 or 2.")
		valid = false
	}
	redistributeConnectedMetric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric"})
	if redistributeConnectedMetric != nil && !validatorRange(*redistributeConnectedMetric, 1, 16) {
		fmt.Fprintln(f, "protocols ospf redistribute connected metric", *redistributeConnectedMetric)
		fmt.Fprintln(f, "metric must be between 1 and 16.")
		valid = false
	}
	redistributeConnectedRouteMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "route-map"})
	if redistributeConnectedRouteMap != nil && !validatorExistsRouteMap(*redistributeConnectedRouteMap) {
		fmt.Fprintln(f, "protocols ospf redistribute connected route-map", *redistributeConnectedRouteMap)
		fmt.Fprintln(f, "route-map", *redistributeConnectedRouteMap, "doesn't exist.")
		valid = false
	}
	// XXX:
	redistributeKernelMetricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric-type"})
	if redistributeKernelMetricType != nil && !validatorInclude(*redistributeKernelMetricType, []string{"1", "2"}) {
		fmt.Fprintln(f, "protocols ospf redistribute kernel metric-type", *redistributeKernelMetricType)
		fmt.Fprintln(f, "metric-type must be either 1 or 2.")
		valid = false
	}
	redistributeKernelMetric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric"})
	if redistributeKernelMetric != nil && !validatorRange(*redistributeKernelMetric, 1, 16) {
		fmt.Fprintln(f, "protocols ospf redistribute kernel metric", *redistributeKernelMetric)
		fmt.Fprintln(f, "metric must be between 1 and 16.")
		valid = false
	}
	redistributeKernelRouteMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "route-map"})
	if redistributeKernelRouteMap != nil && !validatorExistsRouteMap(*redistributeKernelRouteMap) {
		fmt.Fprintln(f, "protocols ospf redistribute kernel route-map", *redistributeKernelRouteMap)
		fmt.Fprintln(f, "route-map", *redistributeKernelRouteMap, "doesn't exist.")
		valid = false
	}
	// XXX:
	redistributeRipMetricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "metric-type"})
	if redistributeRipMetricType != nil && !validatorInclude(*redistributeRipMetricType, []string{"1", "2"}) {
		fmt.Fprintln(f, "protocols ospf redistribute rip metric-type", *redistributeRipMetricType)
		fmt.Fprintln(f, "metric-type must be either 1 or 2.")
		valid = false
	}
	redistributeRipMetric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "metric"})
	if redistributeRipMetric != nil && !validatorRange(*redistributeRipMetric, 1, 16) {
		fmt.Fprintln(f, "protocols ospf redistribute rip metric", *redistributeRipMetric)
		fmt.Fprintln(f, "metric must be between 1 and 16.")
		valid = false
	}
	redistributeRipRouteMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "route-map"})
	if redistributeRipRouteMap != nil && !validatorExistsRouteMap(*redistributeRipRouteMap) {
		fmt.Fprintln(f, "protocols ospf redistribute rip route-map", *redistributeRipRouteMap)
		fmt.Fprintln(f, "route-map", *redistributeRipRouteMap, "doesn't exist.")
		valid = false
	}
	// XXX:
	redistributeStaticMetricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "metric-type"})
	if redistributeStaticMetricType != nil && !validatorInclude(*redistributeStaticMetricType, []string{"1", "2"}) {
		fmt.Fprintln(f, "protocols ospf redistribute static metric-type", *redistributeStaticMetricType)
		fmt.Fprintln(f, "metric-type must be either 1 or 2.")
		valid = false
	}
	redistributeStaticMetric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "metric"})
	if redistributeStaticMetric != nil && !validatorRange(*redistributeStaticMetric, 1, 16) {
		fmt.Fprintln(f, "protocols ospf redistribute static metric", *redistributeStaticMetric)
		fmt.Fprintln(f, "metric must be between 1 and 16.")
		valid = false
	}
	redistributeStaticRouteMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "route-map"})
	if redistributeStaticRouteMap != nil && !validatorExistsRouteMap(*redistributeStaticRouteMap) {
		fmt.Fprintln(f, "protocols ospf redistribute static route-map", *redistributeStaticRouteMap)
		fmt.Fprintln(f, "route-map", *redistributeStaticRouteMap, "doesn't exist.")
		valid = false
	}
	// XXX:
	refreshTimers := configCandidate.value(
		[]string{"protocols", "ospf", "refresh", "timers"})
	if refreshTimers != nil && !validatorRange(*refreshTimers, 10, 1800) {
		fmt.Fprintln(f, "protocols ospf refresh timers", *refreshTimers)
		fmt.Fprintln(f, "must be between 10-1800.")
		valid = false
	}
	timersThrottleSpfDelay := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "delay"})
	if timersThrottleSpfDelay != nil && !validatorRange(*timersThrottleSpfDelay, 0, 600000) {
		fmt.Fprintln(f, "protocols ospf timers throttle spf delay", *timersThrottleSpfDelay)
		fmt.Fprintln(f, "must be between 0-600000.")
		valid = false
	}
	timersThrottleSpfInitialHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "initial-holdtime"})
	if timersThrottleSpfInitialHoldtime != nil && !validatorRange(*timersThrottleSpfInitialHoldtime, 0, 600000) {
		fmt.Fprintln(f, "protocols ospf timers throttle spf initial-holdtime", *timersThrottleSpfInitialHoldtime)
		fmt.Fprintln(f, "must be between 0-600000.")
		valid = false
	}
	timersThrottleSpfMaxHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "max-holdtime"})
	if timersThrottleSpfMaxHoldtime != nil && !validatorRange(*timersThrottleSpfMaxHoldtime, 0, 600000) {
		fmt.Fprintln(f, "protocols ospf timers throttle spf max-holdtime", *timersThrottleSpfMaxHoldtime)
		fmt.Fprintln(f, "must be between 0-600000.")
		valid = false
	}
	return valid
}

// XXX:
func quaggaConfigValidInterfaceOspf(f io.Writer, interface_ string) bool {
	valid := true
	authenticationMd5KeyIds := configCandidate.values(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "authentication", "md5", "key-id"})
	for _, authenticationMd5KeyId := range authenticationMd5KeyIds {
		if !validatorRange(authenticationMd5KeyId, 1, 255) {
			fmt.Fprintln(f, "interfaces interface", interface_,
				"ipv4 ospf authentication md5 key-id", authenticationMd5KeyId)
			fmt.Fprintln(f, "ID must be between (1-255).")
			valid = false
		}
		authenticationMd5KeyIdMd5Key := configCandidate.value(
			[]string{"interfaces", "interface", interface_,
				"ipv4", "ospf", "authentication", "md5", "key-id", authenticationMd5KeyId, "md5-key"})
		if authenticationMd5KeyIdMd5Key == nil {
			fmt.Fprintln(f, "interfaces interface", interface_,
				"ipv4 ospf authentication md5 key-id", authenticationMd5KeyId)
			fmt.Fprintln(f, "Must add the md5-key for key-id", authenticationMd5KeyId, ".")
			valid = false
		} else if len(*authenticationMd5KeyIdMd5Key) < 1 || len(*authenticationMd5KeyIdMd5Key) > 16 {
			fmt.Fprintln(f, "interfaces interface", interface_,
				"ipv4 ospf authentication md5 key-id", authenticationMd5KeyId,
				" md5-key ", *authenticationMd5KeyIdMd5Key)
			fmt.Fprintln(f, "MD5 key must be 16 characters or less.")
			valid = false
		}
	}
	authenticationPlaintextPassword := configCandidate.value(
		[]string{"interfaces", "interface", interface_,
			"ipv4", "ospf", "authentication", "plaintext-password"})
	if authenticationPlaintextPassword != nil &&
		(len(*authenticationPlaintextPassword) < 1 || len(*authenticationPlaintextPassword) > 8) {
		fmt.Fprintln(f, "interfaces interface", interface_,
			"ipv4 ospf authentication plaintext-password", *authenticationPlaintextPassword)
		fmt.Fprintln(f, "Password must be 8 characters or less.")
		valid = false
	}
	bandwidth := configCandidate.value(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "bandwidth"})
	if bandwidth != nil && !validatorRange(*bandwidth, 1, 10000000) {
		fmt.Fprintln(f, "interfaces interface", interface_, "ipv4 ospf bandwidth", *bandwidth)
		fmt.Fprintln(f, "Must be between 1-10000000.")
		valid = false
	}
	cost := configCandidate.value(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "cost"})
	if cost != nil && !validatorRange(*cost, 1, 65535) {
		fmt.Fprintln(f, "interfaces interface", interface_, "ipv4 ospf cost", *cost)
		fmt.Fprintln(f, "Must be between 1-65535.")
		valid = false
	}
	deadInterval := configCandidate.value(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "dead-interval"})
	if deadInterval != nil && !validatorRange(*deadInterval, 1, 65535) {
		fmt.Fprintln(f, "interfaces interface", interface_, "ipv4 ospf dead-interval", *deadInterval)
		fmt.Fprintln(f, "Must be between 1-65535.")
		valid = false
	}
	helloInterval := configCandidate.value(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "hello-interval"})
	if helloInterval != nil && !validatorRange(*helloInterval, 1, 65535) {
		fmt.Fprintln(f, "interfaces interface", interface_, "ipv4 ospf hello-interval", *helloInterval)
		fmt.Fprintln(f, "Must be between 1-65535.")
		valid = false
	}
	network := configCandidate.value(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "network"})
	if network != nil && !validatorInclude(*network,
		[]string{"broadcast", "non-broadcast", "point-to-multipoint", "point-to-point"}) {
		fmt.Fprintln(f, "interfaces interface", interface_, "ipv4 ospf network", *network)
		fmt.Fprintln(f, "Must be (broadcast|non-broadcast|point-to-multipoint|point-to-point).")
		valid = false
	}
	priority := configCandidate.value(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "priority"})
	if priority != nil && !validatorRange(*priority, 0, 255) {
		fmt.Fprintln(f, "interfaces interface", interface_, "ipv4 ospf priority", *priority)
		fmt.Fprintln(f, "Must be between 0-255.")
		valid = false
	}
	retransmitInterval := configCandidate.value(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "retransmit-interval"})
	if retransmitInterval != nil && !validatorRange(*retransmitInterval, 3, 65535) {
		fmt.Fprintln(f, "interfaces interface", interface_, "ipv4 ospf retransmit-interval", *retransmitInterval)
		fmt.Fprintln(f, "Must be between 3-65535.")
		valid = false
	}
	transmitDelay := configCandidate.value(
		[]string{"interfaces", "interface", interface_, "ipv4", "ospf", "transmit-delay"})
	if transmitDelay != nil && !validatorRange(*transmitDelay, 1, 65535) {
		fmt.Fprintln(f, "interfaces interface", interface_, "ipv4 ospf transmit-delay", *transmitDelay)
		fmt.Fprintln(f, "Must be between 1-65535.")
		valid = false
	}
	return valid
}

func quaggaConfigValidInterfacesOspf(f io.Writer) bool {
	valid := true
	interfaces := configCandidate.values([]string{"interfaces", "interface"})
	for _, interface_ := range interfaces {
		if !quaggaConfigValidInterfaceOspf(f, interface_) {
			valid = false
		}
	}
	return valid
}
