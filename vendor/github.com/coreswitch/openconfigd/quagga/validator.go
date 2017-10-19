package quagga

import (
	"net"
	"regexp"
	"strconv"
)

func validatorRange(s string, min, max int) bool {
	i, err := strconv.Atoi(s)
	if err != nil {
		return false
	}
	if i < min || i > max {
		return false
	}
	return true
}

func validatorInclude(s string, list []string) bool {
	for _, item := range list {
		if s == item {
			return true
		}
	}
	return false
}

func validatorIPv4Address(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}
	return true
}

func validatorIPv4CIDR(s string) bool {
	ip, _, err := net.ParseCIDR(s)
	if err != nil {
		return false
	}
	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}
	return true
}

func validatorIPv6Address(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		return false
	}
	ip4 := ip.To4()
	if ip4 != nil {
		return false
	}
	return true
}

func validatorIPv6CIDR(s string) bool {
	ip, _, err := net.ParseCIDR(s)
	if err != nil {
		return false
	}
	ip4 := ip.To4()
	if ip4 != nil {
		return false
	}
	return true
}

func validatorPeer(s string) bool {
	return s == "local" || validatorIPv4Address(s) || validatorIPv6Address(s)
}

func validatorAsPathPrepend(s string) bool {
	r := regexp.MustCompile(`\s+`)
	ss := r.Split(s, -1)
	count := 0
	for _, asNum := range ss {
		if asNum == "" {
			continue
		}
		if !validatorRange(asNum, 1, 4294967294) {
			return false
		}
		count++
	}
	if count > 24 {
		return false
	}
	return true
}

func validatorCommunity(s string) bool {
	r := regexp.MustCompile(`^(additive|internet|local-AS|no-advertise|no-export|none|\d+:\d+)$`)
	return r.MatchString(s)
}

func validatorSource(s string) bool {
	// XXX:
	return validatorIPv4Address(s) || validatorIPv6Address(s)
}

func validatorConfediBGPASNCheck(s, asNum string) bool {
	if s == asNum {
		return false
	}
	neighs := configCandidate.values(
		[]string{"protocols", "bgp", asNum, "neighbor"})
	for _, neigh := range neighs {
		remoteAs := configCandidate.value(
			[]string{"protocols", "bgp", asNum, "neighbor", neigh, "remote-as"})
		if remoteAs != nil && s == *remoteAs {
			return false
		}
	}
	return true
}

func validatorExistsPeerGroup(asNum, peerGroup string) bool {
	return configCandidate.lookup(
		[]string{"protocols", "bgp", asNum, "peer-group", peerGroup}) != nil
}

func validatorExistsAccessList(s string) bool {
	return configCandidate.lookup(
		[]string{"policy", "access-list", s}) != nil
}

func validatorExistsAccessList6(s string) bool {
	return configCandidate.lookup(
		[]string{"policy", "access-list6", s}) != nil
}

func validatorExistsAsPathList(s string) bool {
	return configCandidate.lookup(
		[]string{"policy", "as-path-list", s}) != nil
}

func validatorExistsCommunityList(s string) bool {
	return configCandidate.lookup(
		[]string{"policy", "community-list", s}) != nil
}

func validatorExistsPrefixList(s string) bool {
	return configCandidate.lookup(
		[]string{"policy", "prefix-list", s}) != nil
}

func validatorExistsPrefixList6(s string) bool {
	return configCandidate.lookup(
		[]string{"policy", "prefix-list6", s}) != nil
}

func validatorExistsRouteMap(s string) bool {
	return configCandidate.lookup(
		[]string{"policy", "route-map", s}) != nil
}
