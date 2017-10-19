package quagga

import (
	"fmt"
	"regexp"
	"strconv"
)

/*
 * access-list
 */

type quaggaAccessListRule struct {
	action         *string
	dstAny         bool
	dstHost        *string
	dstInverseMask *string
	dstNetwork     *string
	srcAny         bool
	srcHost        *string
	srcInverseMask *string
	srcNetwork     *string
}

func (config *quaggaConfigStateNode) makeQuaggaAccessListRule(accessList, rule string) *quaggaAccessListRule {
	base := []string{"policy", "access-list", accessList, "rule", rule}
	if config.lookup(base) == nil {
		return nil
	}
	r := &quaggaAccessListRule{}
	r.action = config.value(append(base, "action"))
	r.dstAny = config.lookup(append(base, "destination", "any")) != nil
	r.dstHost = config.value(append(base, "destination", "host"))
	r.dstInverseMask = config.value(append(base, "destination", "inverse-mask"))
	r.dstNetwork = config.value(append(base, "destination", "network"))
	r.srcAny = config.lookup(append(base, "source", "any")) != nil
	r.srcHost = config.value(append(base, "source", "host"))
	r.srcInverseMask = config.value(append(base, "source", "inverse-mask"))
	r.srcNetwork = config.value(append(base, "source", "network"))
	return r
}

func quaggaConfigValidAccessListRule(accessList, rule string) bool {
	if !validatorRange(rule, 1, 65535) {
		fmt.Println("policy access-list", accessList, "rule", rule)
		fmt.Println("rule number must be between 1 and 65535.")
		return false
	}
	r := configCandidate.makeQuaggaAccessListRule(accessList, rule)
	if r == nil {
		fmt.Println("policy access-list", accessList, "rule", rule)
		fmt.Println("rule not found.")
		return false
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Println("policy access-list", accessList, "rule", rule, "action", action)
		fmt.Println("action must be permit or deny.")
		return false
	}
	if r.srcHost != nil && !validatorIPv4Address(*r.srcHost) {
		fmt.Println("policy access-list", accessList, "rule", rule, "source host", *r.srcHost)
		fmt.Println("source host format error.")
		return false
	}
	if r.srcInverseMask != nil && !validatorIPv4Address(*r.srcInverseMask) {
		fmt.Println("policy access-list", accessList, "rule", rule, "source inverse-mask", *r.srcInverseMask)
		fmt.Println("source inverse-mask format error.")
		return false
	}
	if r.srcNetwork != nil && !validatorIPv4Address(*r.srcNetwork) {
		fmt.Println("policy access-list", accessList, "rule", rule, "source network", *r.srcNetwork)
		fmt.Println("source network format error.")
		return false
	}
	srcMatches := 0
	if r.srcAny == true {
		srcMatches++
	}
	if r.srcHost != nil {
		srcMatches++
	}
	if r.srcNetwork != nil {
		srcMatches++
	}
	if srcMatches == 0 {
		fmt.Println("policy access-list", accessList, "rule", rule, "source")
		fmt.Println("you may only define one filter type (host|network|any).")
		return false
	}
	if srcMatches > 1 {
		fmt.Println("policy access-list", accessList, "rule", rule, "source")
		fmt.Println("you may only define one filter type (host|network|any).")
		return false
	}
	if r.srcNetwork != nil && r.srcInverseMask == nil {
		fmt.Println("policy access-list", accessList, "rule", rule, "source")
		fmt.Println("you must specify an inverse-mask if you configure a network.")
		return false
	}
	if r.srcNetwork == nil && r.srcInverseMask != nil {
		fmt.Println("policy access-list", accessList, "rule", rule, "source")
		fmt.Println("you must specify a network if you configure an inverse mask.")
		return false
	}
	if r.dstHost != nil && !validatorIPv4Address(*r.dstHost) {
		fmt.Println("policy access-list", accessList, "rule", rule, "destination host", *r.dstHost)
		fmt.Println("destination host format error.")
		return false
	}
	if r.dstInverseMask != nil && !validatorIPv4Address(*r.dstInverseMask) {
		fmt.Println("policy access-list", accessList, "rule", rule,
			"destination inverse-mask", *r.dstInverseMask)
		fmt.Println("destination inverse-mask format error.")
		return false
	}
	if r.dstNetwork != nil && !validatorIPv4Address(*r.dstNetwork) {
		fmt.Println("policy access-list", accessList, "rule", rule, "destination network", *r.dstNetwork)
		fmt.Println("destination network format error.")
		return false
	}
	dstMatches := 0
	if r.dstAny == true {
		dstMatches++
	}
	if r.dstHost != nil {
		dstMatches++
	}
	if r.dstNetwork != nil {
		dstMatches++
	}
	if dstMatches > 0 && !validatorRange(accessList, 100, 199) && !validatorRange(accessList, 2000, 2699) {
		fmt.Println("policy access-list", accessList, "rule", rule, "destination")
		fmt.Println("access-list number must be <100-199> or <2000-2699> to set destination matches.")
		return false
	}
	if dstMatches == 0 && (validatorRange(accessList, 100, 199) || validatorRange(accessList, 2000, 2699)) {
		fmt.Println("policy access-list", accessList, "rule", rule, "destination")
		fmt.Println("you may only define one filter type (host|network|any).")
		return false
	}
	if dstMatches > 1 {
		fmt.Println("policy access-list", accessList, "rule", rule, "destination")
		fmt.Println("you may only define one filter type (host|network|any).")
		return false
	}
	if r.dstNetwork != nil && r.dstInverseMask == nil {
		fmt.Println("policy access-list", accessList, "rule", rule, "destination")
		fmt.Println("you must specify an inverse-mask if you configure a network.")
		return false
	}
	if r.dstNetwork == nil && r.dstInverseMask != nil {
		fmt.Println("policy access-list", accessList, "rule", rule, "destination")
		fmt.Println("you must specify a network if you configure an inverse mask.")
		return false
	}
	return true
}

func quaggaConfigValidAccessList(accessList string) bool {
	if !validatorRange(accessList, 1, 199) && !validatorRange(accessList, 1300, 2699) {
		fmt.Println("policy access-list", accessList)
		fmt.Println("Access list number must be:")
		fmt.Println("<1-99>      IP standard access list")
		fmt.Println("<100-199>   IP extended access list")
		fmt.Println("<1300-1999> IP standard access list (expanded range)")
		fmt.Println("<2000-2699> IP extended access list (expanded range)")
		return false
	}
	rules := configCandidate.values([]string{"policy", "access-list", accessList, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidAccessListRule(accessList, rule) {
			return false
		}
	}
	return true
}

func quaggaConfigValidAccessLists() bool {
	accessLists := configCandidate.values([]string{"policy", "access-list"})
	for _, accessList := range accessLists {
		if !quaggaConfigValidAccessList(accessList) {
			return false
		}
	}
	return true
}

func quaggaConfigCommitAccessList(accessList string) {
	quaggaVtysh("configure terminal",
		fmt.Sprint("no access-list ", accessList))
	rules := configCandidate.values([]string{"policy", "access-list", accessList, "rule"})
	quaggaRuleSort(rules)
	for _, rule := range rules {
		r := configCandidate.makeQuaggaAccessListRule(accessList, rule)
		if r == nil {
			continue
		}
		action := ""
		ip := ""
		src := ""
		dst := ""
		if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
			continue
		} else {
			action = " " + *r.action
		}
		if r.srcAny {
			src = " any"
		} else if r.srcHost != nil {
			src = " host " + *r.srcHost
		} else if r.srcNetwork != nil && r.srcInverseMask != nil {
			src = " " + *r.srcNetwork + " " + *r.srcInverseMask
		}
		if validatorRange(accessList, 100, 199) || validatorRange(accessList, 2000, 2699) {
			ip = " ip"
			if r.dstAny {
				dst = " any"
			} else if r.dstHost != nil {
				dst = " host " + *r.dstHost
			} else if r.dstNetwork != nil && r.dstInverseMask != nil {
				dst = " " + *r.dstNetwork + " " + *r.dstInverseMask
			}
		}
		quaggaVtysh("configure terminal",
			fmt.Sprint("access-list ", accessList, action, ip, src, dst))
	}
}

func quaggaConfigCommitAccessLists() {
	for _, accessList := range configRunning.values([]string{"policy", "access-list"}) {
		if configCandidate.lookup([]string{"policy", "access-list", accessList}) == nil {
			quaggaVtysh("configure terminal",
				fmt.Sprint("no access-list ", accessList))
		}
	}
	for _, accessList := range configCandidate.values([]string{"policy", "access-list"}) {
		if !quaggaConfigStateDiff([]string{"policy", "access-list", accessList}) {
			continue
		}
		quaggaConfigCommitAccessList(accessList)
	}
}

func quaggaUpdateCheckAccessList() {
	if !commitedAccessList && quaggaConfigStateDiff([]string{"policy", "access-list"}) {
		commitedAccessList = true
		quaggaConfigCommitAccessLists()
	}
}

/*
 * access-list6
 */

type quaggaAccessList6Rule struct {
	action        *string
	srcAny        bool
	srcExactMatch bool
	srcNetwork    *string
}

func (config *quaggaConfigStateNode) makeQuaggaAccessList6Rule(accessList, rule string) *quaggaAccessList6Rule {
	base := []string{"policy", "access-list6", accessList, "rule", rule}
	if config.lookup(base) == nil {
		return nil
	}
	r := &quaggaAccessList6Rule{}
	r.action = config.value(append(base, "action"))
	r.srcAny = config.lookup(append(base, "source", "any")) != nil
	r.srcExactMatch = config.lookup(append(base, "source", "exact-match")) != nil
	r.srcNetwork = config.value(append(base, "source", "network"))
	return r
}

func quaggaConfigValidAccessList6Rule(accessList6, rule string) bool {
	if !validatorRange(rule, 1, 65535) {
		fmt.Println("policy access-list6", accessList6, "rule", rule)
		fmt.Println("rule number must be between 1 and 65535.")
		return false
	}
	r := configCandidate.makeQuaggaAccessList6Rule(accessList6, rule)
	if r == nil {
		fmt.Println("policy access-list6", accessList6, "rule", rule)
		fmt.Println("rule not found.")
		return false
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Println("policy access-list6", accessList6, "rule", rule, "action", action)
		fmt.Println("action must be permit or deny.")
		return false
	}
	if r.srcNetwork != nil && !validatorIPv6CIDR(*r.srcNetwork) {
		fmt.Println("policy access-list6", accessList6, "rule", rule, "source network", *r.srcNetwork)
		fmt.Println("source network format error.")
		return false
	}
	srcMatches := 0
	if r.srcAny == true {
		srcMatches++
	}
	if r.srcNetwork != nil {
		srcMatches++
	}
	if srcMatches == 0 {
		fmt.Println("policy access-list6", accessList6, "rule", rule, "source")
		fmt.Println("you may only define one filter type (network|any).")
		return false
	}
	if srcMatches > 1 {
		fmt.Println("policy access-list6", accessList6, "rule", rule, "source")
		fmt.Println("you may only define one filter type (network|any).")
		return false
	}
	return true
}

func quaggaConfigValidAccessList6(accessList6 string) bool {
	if len(accessList6) < 1 || len(accessList6) > 64 {
		fmt.Println("policy access-list6", accessList6)
		fmt.Println("access-list name must be 64 characters or less.")
		return false
	}
	if accessList6[0] == '-' {
		fmt.Println("policy access-list6", accessList6)
		fmt.Println("access-list name cannot start with \"-\".")
		return false
	}
	nameRegexp := regexp.MustCompile(`^[^|;&$<>]*$`)
	if !nameRegexp.MatchString(accessList6) {
		fmt.Println("policy access-list6", accessList6)
		fmt.Println("access-list name cannot contain shell punctuation.")
		return false
	}
	rules := configCandidate.values([]string{"policy", "access-list6", accessList6, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidAccessList6Rule(accessList6, rule) {
			return false
		}
	}
	return true
}

func quaggaConfigValidAccessList6s() bool {
	accessList6s := configCandidate.values([]string{"policy", "access-list6"})
	for _, accessList6 := range accessList6s {
		if !quaggaConfigValidAccessList6(accessList6) {
			return false
		}
	}
	return true
}

func quaggaConfigCommitAccessList6(accessList6 string) {
	quaggaVtysh("configure terminal",
		fmt.Sprint("no ipv6 access-list ", accessList6))
	rules := configCandidate.values([]string{"policy", "access-list6", accessList6, "rule"})
	quaggaRuleSort(rules)
	for _, rule := range rules {
		r := configCandidate.makeQuaggaAccessList6Rule(accessList6, rule)
		if r == nil {
			continue
		}
		action := ""
		ip := ""
		src := ""
		if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
			continue
		} else {
			action = " " + *r.action
		}
		if r.srcAny {
			src = " any"
		} else if r.srcNetwork != nil {
			src = " " + *r.srcNetwork
			if r.srcExactMatch {
				src += " exact-match"
			}
		}
		quaggaVtysh("configure terminal",
			fmt.Sprint("ipv6 access-list ", accessList6, action, ip, src))
	}
}

func quaggaConfigCommitAccessList6s() {
	for _, accessList6 := range configRunning.values([]string{"policy", "access-list6"}) {
		if configCandidate.lookup([]string{"policy", "access-list6", accessList6}) == nil {
			quaggaVtysh("configure terminal",
				fmt.Sprint("no ipv6 access-list ", accessList6))
		}
	}
	for _, accessList6 := range configCandidate.values([]string{"policy", "access-list6"}) {
		if !quaggaConfigStateDiff([]string{"policy", "access-list6", accessList6}) {
			continue
		}
		quaggaConfigCommitAccessList6(accessList6)
	}
}

func quaggaUpdateCheckAccessList6() {
	if !commitedAccessList6 && quaggaConfigStateDiff([]string{"policy", "access-list6"}) {
		commitedAccessList6 = true
		quaggaConfigCommitAccessList6s()
	}
}

/*
 * as-path-list
 */

type quaggaAsPathListRule struct {
	action *string
	regex  *string
}

func (config *quaggaConfigStateNode) makeQuaggaAsPathListRule(asPathList, rule string) *quaggaAsPathListRule {
	base := []string{"policy", "as-path-list", asPathList, "rule", rule}
	if config.lookup(base) == nil {
		return nil
	}
	r := &quaggaAsPathListRule{}
	r.action = config.value(append(base, "action"))
	r.regex = config.value(append(base, "regex"))
	return r
}

func quaggaConfigValidAsPathListRule(asPathList, rule string) bool {
	if !validatorRange(rule, 1, 65535) {
		fmt.Println("policy as-path-list", asPathList, "rule", rule)
		fmt.Println("rule number must be between 1 and 65535.")
		return false
	}
	r := configCandidate.makeQuaggaAsPathListRule(asPathList, rule)
	if r == nil {
		fmt.Println("policy as-path-list", asPathList, "rule", rule)
		fmt.Println("rule not found.")
		return false
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Println("policy as-path-list", asPathList, "rule", rule, "action", action)
		fmt.Println("action must be permit or deny.")
		return false
	}
	if r.regex == nil {
		fmt.Println("policy as-path-list", asPathList, "rule", rule, "regex")
		fmt.Println("you must specify a regex.")
		return false
	}
	return true
}

func quaggaConfigValidAsPathList(asPathList string) bool {
	nameRegexp := regexp.MustCompile(`^[-a-zA-Z0-9.]+$`)
	if !nameRegexp.MatchString(asPathList) {
		fmt.Println("policy as-path-list", asPathList)
		fmt.Println("as-path-list name must be alpha-numeric.")
		return false
	}
	rules := configCandidate.values([]string{"policy", "as-path-list", asPathList, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidAsPathListRule(asPathList, rule) {
			return false
		}
	}
	return true
}

func quaggaConfigValidAsPathLists() bool {
	asPathLists := configCandidate.values([]string{"policy", "as-path-list"})
	for _, asPathList := range asPathLists {
		if !quaggaConfigValidAsPathList(asPathList) {
			return false
		}
	}
	return true
}

func quaggaConfigCommitAsPathList(asPathList string) {
	quaggaVtysh("configure terminal",
		fmt.Sprint("no ip as-path access-list ", asPathList))
	rules := configCandidate.values([]string{"policy", "as-path-list", asPathList, "rule"})
	quaggaRuleSort(rules)
	for _, rule := range rules {
		r := configCandidate.makeQuaggaAsPathListRule(asPathList, rule)
		if r == nil {
			continue
		}
		action := ""
		regex := ""
		if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
			continue
		} else {
			action = " " + *r.action
		}
		if r.regex == nil {
			continue
		} else {
			regex = " " + *r.regex
		}
		quaggaVtysh("configure terminal",
			fmt.Sprint("ip as-path access-list ", asPathList, action, regex))
	}
}

func quaggaConfigCommitAsPathLists() {
	for _, asPathList := range configRunning.values([]string{"policy", "as-path-list"}) {
		if configCandidate.lookup([]string{"policy", "as-path-list", asPathList}) == nil {
			quaggaVtysh("configure terminal",
				fmt.Sprint("no ip as-path access-list ", asPathList))
		}
	}
	for _, asPathList := range configCandidate.values([]string{"policy", "as-path-list"}) {
		if !quaggaConfigStateDiff([]string{"policy", "as-path-list", asPathList}) {
			continue
		}
		quaggaConfigCommitAsPathList(asPathList)
	}
}

func quaggaUpdateCheckAsPathList() {
	if !commitedAsPathList && quaggaConfigStateDiff([]string{"policy", "as-path-list"}) {
		commitedAsPathList = true
		quaggaConfigCommitAsPathLists()
	}
}

/*
 * community-list
 */

type quaggaCommunityListRule struct {
	action *string
	regex  *string
}

func (config *quaggaConfigStateNode) makeQuaggaCommunityListRule(communityList, rule string) *quaggaCommunityListRule {
	base := []string{"policy", "community-list", communityList, "rule", rule}
	if config.lookup(base) == nil {
		return nil
	}
	r := &quaggaCommunityListRule{}
	r.action = config.value(append(base, "action"))
	r.regex = config.value(append(base, "regex"))
	return r
}

func quaggaConfigValidCommunityListRule(communityList, rule string) bool {
	if !validatorRange(rule, 1, 65535) {
		fmt.Println("policy community-list", communityList, "rule", rule)
		fmt.Println("rule number must be between 1 and 65535.")
		return false
	}
	r := configCandidate.makeQuaggaCommunityListRule(communityList, rule)
	if r == nil {
		fmt.Println("policy community-list", communityList, "rule", rule)
		fmt.Println("rule not found.")
		return false
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Println("policy community-list", communityList, "rule", rule, "action", action)
		fmt.Println("action must be permit or deny.")
		return false
	}
	if r.regex == nil {
		fmt.Println("policy community-list", communityList, "rule", rule, "regex")
		fmt.Println("you must specify a regex.")
		return false
	}
	if r.regex != nil && validatorRange(rule, 1, 99) {
		stdRegexp := regexp.MustCompile(`^(internet|local-AS|no-advertise|no-export|\d+:\d+)$`)
		if !stdRegexp.MatchString(*r.regex) {
			fmt.Println("policy community-list", communityList, "rule", rule, "regex", *r.regex)
			fmt.Println("regex ", *r.regex, " is invalid for a standard community list")
			return false
		}
	}
	return true
}

func quaggaConfigValidCommunityList(communityList string) bool {
	if !validatorRange(communityList, 1, 500) {
		fmt.Println("policy community-list", communityList)
		fmt.Println("community-list must be:")
		fmt.Println("<1-99>    BGP community list (standard)")
		fmt.Println("<100-500> BGP community list (expanded)")
		return false
	}
	rules := configCandidate.values([]string{"policy", "community-list", communityList, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidCommunityListRule(communityList, rule) {
			return false
		}
	}
	return true
}

func quaggaConfigValidCommunityLists() bool {
	communityLists := configCandidate.values([]string{"policy", "community-list"})
	for _, communityList := range communityLists {
		if !quaggaConfigValidCommunityList(communityList) {
			return false
		}
	}
	return true
}

func quaggaConfigCommitCommunityList(communityList string) {
	quaggaVtysh("configure terminal",
		fmt.Sprint("no ip community-list ", communityList))
	rules := configCandidate.values([]string{"policy", "community-list", communityList, "rule"})
	quaggaRuleSort(rules)
	for _, rule := range rules {
		r := configCandidate.makeQuaggaCommunityListRule(communityList, rule)
		if r == nil {
			continue
		}
		action := ""
		regex := ""
		if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
			continue
		} else {
			action = " " + *r.action
		}
		if r.regex == nil {
			continue
		} else {
			regex = " " + *r.regex
		}
		quaggaVtysh("configure terminal",
			fmt.Sprint("ip community-list ", communityList, action, regex))
	}
}

func quaggaConfigCommitCommunityLists() {
	for _, communityList := range configRunning.values([]string{"policy", "community-list"}) {
		if configCandidate.lookup([]string{"policy", "community-list", communityList}) == nil {
			quaggaVtysh("configure terminal",
				fmt.Sprint("no ip community-list ", communityList))
		}
	}
	for _, communityList := range configCandidate.values([]string{"policy", "community-list"}) {
		if !quaggaConfigStateDiff([]string{"policy", "community-list", communityList}) {
			continue
		}
		quaggaConfigCommitCommunityList(communityList)
	}
}

func quaggaUpdateCheckCommunityList() {
	if !commitedCommunityList && quaggaConfigStateDiff([]string{"policy", "community-list"}) {
		commitedCommunityList = true
		quaggaConfigCommitCommunityLists()
	}
}

/*
 * prefix-list
 */

type quaggaPrefixListRule struct {
	action *string
	ge     *string
	le     *string
	prefix *string
}

func (config *quaggaConfigStateNode) makeQuaggaPrefixListRule(prefixList, rule string) *quaggaPrefixListRule {
	base := []string{"policy", "prefix-list", prefixList, "rule", rule}
	if config.lookup(base) == nil {
		return nil
	}
	r := &quaggaPrefixListRule{}
	r.action = config.value(append(base, "action"))
	r.ge = config.value(append(base, "ge"))
	r.le = config.value(append(base, "le"))
	r.prefix = config.value(append(base, "prefix"))
	return r
}

func quaggaConfigValidPrefixListRule(prefixList, rule string) bool {
	if !validatorRange(rule, 1, 65535) {
		fmt.Println("policy prefix-list", prefixList, "rule", rule)
		fmt.Println("rule number must be between 1 and 65535.")
		return false
	}
	r := configCandidate.makeQuaggaPrefixListRule(prefixList, rule)
	if r == nil {
		fmt.Println("policy prefix-list", prefixList, "rule", rule)
		fmt.Println("rule not found.")
		return false
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Println("policy prefix-list", prefixList, "rule", rule, "action", action)
		fmt.Println("action must be permit or deny.")
		return false
	}
	if r.le != nil && !validatorRange(*r.le, 0, 32) {
		le := ""
		if r.le != nil {
			le = *r.le
		}
		fmt.Println("policy prefix-list", prefixList, "rule", rule, "le", le)
		fmt.Println("le must be between 0 and 32.")
		return false
	}
	if r.ge != nil && !validatorRange(*r.ge, 0, 32) {
		ge := ""
		if r.ge != nil {
			ge = *r.ge
		}
		fmt.Println("policy prefix-list", prefixList, "rule", rule, "ge", ge)
		fmt.Println("ge must be between 0 and 32.")
		return false
	}
	if r.prefix == nil || !validatorIPv4CIDR(*r.prefix) {
		prefix := ""
		if r.prefix != nil {
			prefix = *r.prefix
		}
		fmt.Println("policy prefix-list", prefixList, "rule", rule, "prefix", prefix)
		fmt.Println("you must specify a prefix.")
		return false
	}
	return true
}

func quaggaConfigValidPrefixList(prefixList string) bool {
	nameRegexp := regexp.MustCompile(`^[-a-zA-Z0-9.]+$`)
	if !nameRegexp.MatchString(prefixList) {
		fmt.Println("policy prefix-list", prefixList)
		fmt.Println("prefix-list name must be alpha-numeric.")
		return false
	}
	rules := configCandidate.values([]string{"policy", "prefix-list", prefixList, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidPrefixListRule(prefixList, rule) {
			return false
		}
	}
	return true
}

func quaggaConfigValidPrefixLists() bool {
	prefixLists := configCandidate.values([]string{"policy", "prefix-list"})
	for _, prefixList := range prefixLists {
		if !quaggaConfigValidPrefixList(prefixList) {
			return false
		}
	}
	return true
}

func quaggaConfigCommitPrefixList(prefixList string) {
	quaggaVtysh("configure terminal",
		fmt.Sprint("no ip prefix-list ", prefixList))
	rules := configCandidate.values([]string{"policy", "prefix-list", prefixList, "rule"})
	quaggaRuleSort(rules)
	for _, rule := range rules {
		r := configCandidate.makeQuaggaPrefixListRule(prefixList, rule)
		if r == nil {
			continue
		}
		action := ""
		prefix := ""
		cond := ""
		if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
			continue
		} else {
			action = " " + *r.action
		}
		if r.prefix == nil || !validatorIPv4CIDR(*r.prefix) {
			continue
		} else {
			prefix = " " + *r.prefix
		}
		if r.ge != nil && !validatorRange(*r.ge, 0, 32) {
			continue
		} else if r.ge != nil {
			cond += " ge " + *r.ge
		}
		if r.le != nil && !validatorRange(*r.le, 0, 32) {
			continue
		} else if r.le != nil {
			cond += " le " + *r.le
		}
		quaggaVtysh("configure terminal",
			fmt.Sprint("ip prefix-list ", prefixList, " seq ", rule, action, prefix, cond))
	}
}

func quaggaConfigCommitPrefixLists() {
	for _, prefixList := range configRunning.values([]string{"policy", "prefix-list"}) {
		if configCandidate.lookup([]string{"policy", "prefix-list", prefixList}) == nil {
			quaggaVtysh("configure terminal",
				fmt.Sprint("no ip prefix-list ", prefixList))
		}
	}
	for _, prefixList := range configCandidate.values([]string{"policy", "prefix-list"}) {
		if !quaggaConfigStateDiff([]string{"policy", "prefix-list", prefixList}) {
			continue
		}
		quaggaConfigCommitPrefixList(prefixList)
	}
}

func quaggaUpdateCheckPrefixList() {
	if !commitedPrefixList && quaggaConfigStateDiff([]string{"policy", "prefix-list"}) {
		commitedPrefixList = true
		quaggaConfigCommitPrefixLists()
	}
}

/*
 * prefix-list6
 */

type quaggaPrefixList6Rule struct {
	action *string
	ge     *string
	le     *string
	prefix *string
}

func (config *quaggaConfigStateNode) makeQuaggaPrefixList6Rule(prefixList6, rule string) *quaggaPrefixList6Rule {
	base := []string{"policy", "prefix-list6", prefixList6, "rule", rule}
	if config.lookup(base) == nil {
		return nil
	}
	r := &quaggaPrefixList6Rule{}
	r.action = config.value(append(base, "action"))
	r.ge = config.value(append(base, "ge"))
	r.le = config.value(append(base, "le"))
	r.prefix = config.value(append(base, "prefix"))
	return r
}

func quaggaConfigValidPrefixList6Rule(prefixList6, rule string) bool {
	if !validatorRange(rule, 1, 65535) {
		fmt.Println("policy prefix-list6", prefixList6, "rule", rule)
		fmt.Println("rule number must be between 1 and 65535.")
		return false
	}
	r := configCandidate.makeQuaggaPrefixList6Rule(prefixList6, rule)
	if r == nil {
		fmt.Println("policy prefix-list6", prefixList6, "rule", rule)
		fmt.Println("rule not found.")
		return false
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Println("policy prefix-list6", prefixList6, "rule", rule, "action", action)
		fmt.Println("action must be permit or deny.")
		return false
	}
	if r.le != nil && !validatorRange(*r.le, 0, 32) {
		le := ""
		if r.le != nil {
			le = *r.le
		}
		fmt.Println("policy prefix-list6", prefixList6, "rule", rule, "le", le)
		fmt.Println("le must be between 0 and 32.")
		return false
	}
	if r.ge != nil && !validatorRange(*r.ge, 0, 32) {
		ge := ""
		if r.ge != nil {
			ge = *r.ge
		}
		fmt.Println("policy prefix-list6", prefixList6, "rule", rule, "ge", ge)
		fmt.Println("ge must be between 0 and 32.")
		return false
	}
	if r.prefix == nil || !validatorIPv6CIDR(*r.prefix) {
		prefix := ""
		if r.prefix != nil {
			prefix = *r.prefix
		}
		fmt.Println("policy prefix-list6", prefixList6, "rule", rule, "prefix", prefix)
		fmt.Println("you must specify a prefix.")
		return false
	}
	return true
}

func quaggaConfigValidPrefixList6(prefixList6 string) bool {
	nameRegexp := regexp.MustCompile(`^[-a-zA-Z0-9.]+$`)
	if !nameRegexp.MatchString(prefixList6) {
		fmt.Println("policy prefix-list6", prefixList6)
		fmt.Println("prefix-list6 name must be alpha-numeric.")
		return false
	}
	rules := configCandidate.values([]string{"policy", "prefix-list6", prefixList6, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidPrefixList6Rule(prefixList6, rule) {
			return false
		}
	}
	return true
}

func quaggaConfigValidPrefixList6s() bool {
	prefixList6s := configCandidate.values([]string{"policy", "prefix-list6"})
	for _, prefixList6 := range prefixList6s {
		if !quaggaConfigValidPrefixList6(prefixList6) {
			return false
		}
	}
	return true
}

func quaggaConfigCommitPrefixList6(prefixList6 string) {
	quaggaVtysh("configure terminal",
		fmt.Sprint("no ipv6 prefix-list ", prefixList6))
	rules := configCandidate.values([]string{"policy", "prefix-list6", prefixList6, "rule"})
	quaggaRuleSort(rules)
	for _, rule := range rules {
		r := configCandidate.makeQuaggaPrefixList6Rule(prefixList6, rule)
		if r == nil {
			continue
		}
		action := ""
		prefix := ""
		cond := ""
		if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
			continue
		} else {
			action = " " + *r.action
		}
		if r.prefix == nil || !validatorIPv6CIDR(*r.prefix) {
			continue
		} else {
			prefix = " " + *r.prefix
		}
		if r.ge != nil && !validatorRange(*r.ge, 0, 128) {
			continue
		} else if r.ge != nil {
			cond += " ge " + *r.ge
		}
		if r.le != nil && !validatorRange(*r.le, 0, 128) {
			continue
		} else if r.le != nil {
			cond += " le " + *r.le
		}
		quaggaVtysh("configure terminal",
			fmt.Sprint("ipv6 prefix-list ", prefixList6, " seq ", rule, action, prefix, cond))
	}
}

func quaggaConfigCommitPrefixList6s() {
	for _, prefixList6 := range configRunning.values([]string{"policy", "prefix-list6"}) {
		if configCandidate.lookup([]string{"policy", "prefix-list6", prefixList6}) == nil {
			quaggaVtysh("configure terminal",
				fmt.Sprint("no ipv6 prefix-list ", prefixList6))
		}
	}
	for _, prefixList6 := range configCandidate.values([]string{"policy", "prefix-list6"}) {
		if !quaggaConfigStateDiff([]string{"policy", "prefix-list6", prefixList6}) {
			continue
		}
		quaggaConfigCommitPrefixList6(prefixList6)
	}
}

func quaggaUpdateCheckPrefixList6() {
	if !commitedPrefixList6 && quaggaConfigStateDiff([]string{"policy", "prefix-list6"}) {
		commitedPrefixList6 = true
		quaggaConfigCommitPrefixList6s()
	}
}

/*
 * route-map
 */

type quaggaRouteMapRule struct {
	action                       *string
	call                         *string
	continue_                    *string
	matchAsPath                  *string
	matchCommunityCommunityList  *string
	matchCommunityExactMatch     bool
	matchInterface               *string
	matchIpAddressAccessList     *string
	matchIpAddressPrefixList     *string
	matchIpNexthopAccessList     *string
	matchIpNexthopPrefixList     *string
	matchIpRouteSourceAccessList *string
	matchIpRouteSourcePrefixList *string
	matchIpv6AddressAccessList   *string
	matchIpv6AddressPrefixList   *string
	matchIpv6NexthopAccessList   *string
	matchIpv6NexthopPrefixList   *string
	matchMetric                  *string
	matchOrigin                  *string
	matchPeer                    *string
	matchTag                     *string
	onMatchGoto                  *string
	onMatchNext                  bool
	setAggregatorAs              *string
	setAggregatorIp              *string
	setAsPathPrepend             *string
	setAtomicAggregate           bool
	setCommListCommList          *string
	setCommListDelete            bool
	setCommunity                 *string
	setIpNextHop                 *string
	setIpv6NextHopGlobal         *string
	setIpv6NextHopLocal          *string
	setLocalPreference           *string
	setMetricType                *string
	setMetric                    *string
	setOrigin                    *string
	setOriginatorId              *string
	setTag                       *string
	setWeight                    *string
}

func (config *quaggaConfigStateNode) makeQuaggaRouteMapRule(routeMap, rule string) *quaggaRouteMapRule {
	base := []string{"policy", "route-map", routeMap, "rule", rule}
	if config.lookup(base) == nil {
		return nil
	}
	r := &quaggaRouteMapRule{}

	r.action = config.value(append(base, "action"))
	r.call = config.value(append(base, "call"))
	r.continue_ = config.value(append(base, "continue"))
	r.matchAsPath = config.value(append(base, "match", "as-path"))
	r.matchCommunityCommunityList = config.value(append(base, "match", "community", "community-list"))
	r.matchCommunityExactMatch = config.lookup(append(base, "match", "community", "exact-match")) != nil
	r.matchInterface = config.value(append(base, "match", "interface"))
	r.matchIpAddressAccessList = config.value(append(base, "match", "ip", "address", "access-list"))
	r.matchIpAddressPrefixList = config.value(append(base, "match", "ip", "address", "prefix-list"))
	r.matchIpNexthopAccessList = config.value(append(base, "match", "ip", "nexthop", "access-list"))
	r.matchIpNexthopPrefixList = config.value(append(base, "match", "ip", "nexthop", "prefix-list"))
	r.matchIpRouteSourceAccessList = config.value(append(base, "match", "ip", "route-source", "access-list"))
	r.matchIpRouteSourcePrefixList = config.value(append(base, "match", "ip", "route-source", "prefix-list"))
	r.matchIpv6AddressAccessList = config.value(append(base, "match", "ipv6", "address", "access-list"))
	r.matchIpv6AddressPrefixList = config.value(append(base, "match", "ipv6", "address", "prefix-list"))
	r.matchIpv6NexthopAccessList = config.value(append(base, "match", "ipv6", "nexthop", "access-list"))
	r.matchIpv6NexthopPrefixList = config.value(append(base, "match", "ipv6", "nexthop", "prefix-list"))
	r.matchMetric = config.value(append(base, "match", "metric"))
	r.matchOrigin = config.value(append(base, "match", "origin"))
	r.matchPeer = config.value(append(base, "match", "peer"))
	r.matchTag = config.value(append(base, "match", "tag"))
	r.onMatchGoto = config.value(append(base, "on-match", "goto"))
	r.onMatchNext = config.lookup(append(base, "on-match", "next")) != nil
	r.setAggregatorAs = config.value(append(base, "set", "aggregator", "as"))
	r.setAggregatorIp = config.value(append(base, "set", "aggregator", "ip"))
	r.setAsPathPrepend = config.value(append(base, "set", "as-path-prepend"))
	r.setAtomicAggregate = config.lookup(append(base, "set", "atomic-aggregate")) != nil
	r.setCommListCommList = config.value(append(base, "set", "comm-list", "comm-list"))
	r.setCommListDelete = config.lookup(append(base, "set", "comm-list", "delete")) != nil
	r.setCommunity = config.value(append(base, "set", "community"))
	r.setIpNextHop = config.value(append(base, "set", "ip-next-hop"))
	r.setIpv6NextHopGlobal = config.value(append(base, "set", "ipv6-next-hop", "global"))
	r.setIpv6NextHopLocal = config.value(append(base, "set", "ipv6-next-hop", "local"))
	r.setLocalPreference = config.value(append(base, "set", "local-preference"))
	r.setMetricType = config.value(append(base, "set", "metric-type"))
	r.setMetric = config.value(append(base, "set", "metric"))
	r.setOrigin = config.value(append(base, "set", "origin"))
	r.setOriginatorId = config.value(append(base, "set", "originator-id"))
	r.setTag = config.value(append(base, "set", "tag"))
	r.setWeight = config.value(append(base, "set", "weight"))

	return r
}

func quaggaConfigValidRouteMapRule(routeMap, rule string) bool {
	if !validatorRange(rule, 1, 65535) {
		fmt.Println("policy route-map", routeMap, "rule", rule)
		fmt.Println("rule number must be between 1 and 65535.")
		return false
	}
	r := configCandidate.makeQuaggaRouteMapRule(routeMap, rule)
	if r == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule)
		fmt.Println("rule not found.")
		return false
	}

	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Println("policy route-map", routeMap, "rule", rule, "action", action)
		fmt.Println("action must be permit or deny.")
		return false
	}

	if r.call != nil && configCandidate.lookup([]string{"policy", "route-map", *r.call}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule, "call", *r.call)
		fmt.Println("called route-map ", *r.call, " doesn't exist.")
		return false
	}
	if r.continue_ != nil {
		if !validatorRange(*r.continue_, 1, 65535) {
			fmt.Println("policy route-map", routeMap, "rule", rule, "continue", *r.continue_)
			fmt.Println("continue must be between 1 and 65535..")
			return false
		}
		from, _ := strconv.Atoi(rule)
		to, _ := strconv.Atoi(*r.continue_)
		if !(to > from) {
			fmt.Println("policy route-map", routeMap, "rule", rule, "continue", *r.continue_)
			fmt.Println("you may only continue forward in the route-map..")
			return false
		}
	}
	if r.matchAsPath != nil &&
		configCandidate.lookup([]string{"policy", "as-path-list", *r.matchAsPath}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule, "match as-path", *r.matchAsPath)
		fmt.Println("match as-path: AS path list ", *r.matchAsPath, " doesn't exist.")
		return false
	}
	if r.matchCommunityCommunityList != nil &&
		configCandidate.lookup([]string{"policy", "community-list", *r.matchCommunityCommunityList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match community community-list", *r.matchCommunityCommunityList)
		fmt.Println("community-list ", *r.matchCommunityCommunityList, " doesn't exist.")
		return false
	}
	/*
		if r.matchCommunityExactMatch != nil && !validator(*r.matchCommunityExactMatch) {
			fmt.Println("policy route-map", routeMap, "rule", rule,
				"matchCommunityExactMatch", *r.matchCommunityExactMatch)
			fmt.Println("matchCommunityExactMatch format error.")
			return false
		}
	*/
	/*
		XXX:
		if r.matchInterface != nil && !validator(*r.matchInterface) {
			fmt.Println("policy route-map", routeMap, "rule", rule,
				"matchInterface", *r.matchInterface)
			fmt.Println("matchInterface format error.")
			return false
		}
	*/
	if r.matchIpAddressAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list", *r.matchIpAddressAccessList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip address access-list", *r.matchIpAddressAccessList)
		fmt.Println("access-list ", *r.matchIpAddressAccessList, " does not exist.")
		return false
	}
	if r.matchIpAddressPrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list", *r.matchIpAddressPrefixList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip address prefix-list", *r.matchIpAddressPrefixList)
		fmt.Println("prefix-list ", *r.matchIpAddressPrefixList, " does not exist.")
		return false
	}
	if r.matchIpAddressAccessList != nil && r.matchIpAddressPrefixList != nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip address access-list", *r.matchIpAddressAccessList)
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip address prefix-list", *r.matchIpAddressPrefixList)
		fmt.Println("you may only specify a prefix-list or access-list")
		return false
	}
	if r.matchIpNexthopAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list", *r.matchIpNexthopAccessList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip nexthop access-list", *r.matchIpNexthopAccessList)
		fmt.Println("access-list ", *r.matchIpNexthopAccessList, " does not exist.")
		return false
	}
	if r.matchIpNexthopPrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list", *r.matchIpNexthopPrefixList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip nexthop prefix-list", *r.matchIpNexthopPrefixList)
		fmt.Println("prefix-list ", *r.matchIpNexthopPrefixList, " does not exist.")
		return false
	}
	if r.matchIpNexthopAccessList != nil && r.matchIpNexthopPrefixList != nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip nexthop access-list", *r.matchIpNexthopAccessList)
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip nexthop prefix-list", *r.matchIpNexthopPrefixList)
		fmt.Println("you may only specify a prefix-list or access-list")
		return false
	}
	if r.matchIpRouteSourceAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list", *r.matchIpRouteSourceAccessList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip route-source access-list", *r.matchIpRouteSourceAccessList)
		fmt.Println("access-list ", *r.matchIpRouteSourceAccessList, " does not exist.")
		return false
	}
	if r.matchIpRouteSourcePrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list", *r.matchIpRouteSourcePrefixList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip route-source prefix-list", *r.matchIpRouteSourcePrefixList)
		fmt.Println("prefix-list ", *r.matchIpRouteSourcePrefixList, " does not exist.")
		return false
	}
	if r.matchIpRouteSourceAccessList != nil && r.matchIpRouteSourcePrefixList != nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip route-source access-list", *r.matchIpRouteSourceAccessList)
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ip route-source prefix-list", *r.matchIpRouteSourcePrefixList)
		fmt.Println("you may only specify a prefix-list or access-list")
		return false
	}
	if r.matchIpv6AddressAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list6", *r.matchIpv6AddressAccessList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ipv6 address access-list", *r.matchIpv6AddressAccessList)
		fmt.Println("access-list6 ", *r.matchIpv6AddressAccessList, " does not exist.")
		return false
	}
	if r.matchIpv6AddressPrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list6", *r.matchIpv6AddressPrefixList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ipv6 address prefix-list", *r.matchIpv6AddressPrefixList)
		fmt.Println("prefix-list6 ", *r.matchIpv6AddressPrefixList, " does not exist.")
		return false
	}
	if r.matchIpv6AddressAccessList != nil && r.matchIpv6AddressPrefixList != nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ipv6 address access-list", *r.matchIpv6AddressAccessList)
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ipv6 address prefix-list", *r.matchIpv6AddressPrefixList)
		fmt.Println("you may only specify a prefix-list or access-list")
		return false
	}
	if r.matchIpv6NexthopAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list6", *r.matchIpv6NexthopAccessList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ipv6 nexthop access-list", *r.matchIpv6NexthopAccessList)
		fmt.Println("access-list6 ", *r.matchIpv6NexthopAccessList, " does not exist.")
		return false
	}
	if r.matchIpv6NexthopPrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list6", *r.matchIpv6NexthopPrefixList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ipv6 nexthop prefix-list", *r.matchIpv6NexthopPrefixList)
		fmt.Println("prefix-list6 ", *r.matchIpv6NexthopPrefixList, " does not exist.")
		return false
	}
	if r.matchIpv6NexthopAccessList != nil && r.matchIpv6NexthopPrefixList != nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ipv6 nexthop access-list", *r.matchIpv6NexthopAccessList)
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"match ipv6 nexthop prefix-list", *r.matchIpv6NexthopPrefixList)
		fmt.Println("you may only specify a prefix-list or access-list")
		return false
	}
	if r.matchMetric != nil && !validatorRange(*r.matchMetric, 1, 65535) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "match metric", *r.matchMetric)
		fmt.Println("metric must be between 1 and 65535.")
		return false
	}
	if r.matchOrigin != nil && !validatorInclude(*r.matchOrigin, []string{"egp", "igp", "incomplete"}) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "match origin", *r.matchOrigin)
		fmt.Println("origin must be egp, igp, or incomplete.")
		return false
	}
	if r.matchPeer != nil && !validatorPeer(*r.matchPeer) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "match peer", *r.matchPeer)
		fmt.Println("peer must be either an IP or local.")
		return false
	}
	if r.matchTag != nil && !validatorRange(*r.matchTag, 1, 65535) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "match tag", *r.matchTag)
		fmt.Println("tag must be between 1 and 65535.")
		return false
	}
	if r.onMatchGoto != nil {
		if !validatorRange(*r.onMatchGoto, 1, 65535) {
			fmt.Println("policy route-map", routeMap, "rule", rule, "on-match goto", *r.onMatchGoto)
			fmt.Println("goto must be a rule number between 1 and 65535.")
			return false
		}
		from, _ := strconv.Atoi(rule)
		to, _ := strconv.Atoi(*r.onMatchGoto)
		if !(to > from) {
			fmt.Println("policy route-map", routeMap, "rule", rule, "on-match goto", *r.onMatchGoto)
			fmt.Println("you may only go forward in the route-map.")
			return false
		}
	}
	/*
		if r.onMatchNext != nil && !validator(*r.onMatchNext) {
			fmt.Println("policy route-map", routeMap, "rule", rule, "on-match next", *r.onMatchNext)
			fmt.Println("onMatchNext format error.")
			return false
		}
	*/
	if r.onMatchGoto != nil && r.onMatchNext {
		fmt.Println("policy route-map", routeMap, "rule", rule, "on-match goto", *r.onMatchGoto)
		fmt.Println("policy route-map", routeMap, "rule", rule, "on-match next")
		fmt.Println("you may set only goto or next")
		return false
	}
	if r.setAggregatorAs != nil && !validatorRange(*r.setAggregatorAs, 1, 65535) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "set aggregator as", *r.setAggregatorAs)
		fmt.Println("BGP AS number must be between 1 and 4294967294.")
		return false
	}
	if r.setAggregatorIp != nil && !validatorIPv4Address(*r.setAggregatorIp) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "set aggregator ip", *r.setAggregatorIp)
		fmt.Println("aggregator IP format error.")
		return false
	}
	if r.setAggregatorAs != nil && r.setAggregatorIp == nil ||
		r.setAggregatorAs == nil && r.setAggregatorIp != nil {
		setAggregatorAs := ""
		if r.setAggregatorAs != nil {
			setAggregatorAs = *r.setAggregatorAs
		}
		setAggregatorIp := ""
		if r.setAggregatorIp != nil {
			setAggregatorIp = *r.setAggregatorIp
		}
		fmt.Println("policy route-map", routeMap, "rule", rule, "set aggregator as", setAggregatorAs)
		fmt.Println("policy route-map", routeMap, "rule", rule, "set aggregator ip", setAggregatorIp)
		fmt.Println("you must configure both as and ip")
		return false
	}
	if r.setAsPathPrepend != nil && !validatorAsPathPrepend(*r.setAsPathPrepend) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "setAsPathPrepend", *r.setAsPathPrepend)
		fmt.Println("invalid AS path string.")
		return false
	}
	/*
		if r.setAtomicAggregate != nil && !validator(*r.setAtomicAggregate) {
			fmt.Println("policy route-map", routeMap, "rule", rule,
				"setAtomicAggregate", *r.setAtomicAggregate)
			fmt.Println("setAtomicAggregate format error.")
			return false
		}
	*/
	if r.setCommListCommList != nil &&
		configCandidate.lookup([]string{"policy", "community-list", *r.setCommListCommList}) == nil {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"set comm-list comm-list", *r.setCommListCommList)
		fmt.Println("community list ", *r.setCommListCommList, " does not exist.")
		return false
	}
	/*
		if r.setCommListDelete != nil && !validator(*r.setCommListDelete) {
			fmt.Println("policy route-map", routeMap, "rule", rule,
				"setCommListDelete", *r.setCommListDelete)
			fmt.Println("setCommListDelete format error.")
			return false
		}
	*/
	if r.setCommunity != nil && !validatorCommunity(*r.setCommunity) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "set community", *r.setCommunity)
		fmt.Println("community format error.")
		return false
	}
	if r.setIpNextHop != nil && !validatorIPv4Address(*r.setIpNextHop) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "set ip-next-hop", *r.setIpNextHop)
		fmt.Println("ip-next-hop format error.")
		return false
	}
	if r.setIpv6NextHopGlobal != nil && !validatorIPv6Address(*r.setIpv6NextHopGlobal) {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"set ipv6-next-hop global", *r.setIpv6NextHopGlobal)
		fmt.Println("ipv6-next-hop global format error.")
		return false
	}
	if r.setIpv6NextHopLocal != nil && !validatorIPv6Address(*r.setIpv6NextHopLocal) {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"set ipv6-next-hop local", *r.setIpv6NextHopLocal)
		fmt.Println("ipv6-next-hop local format error.")
		return false
	}
	if r.setLocalPreference != nil && !validatorRange(*r.setLocalPreference, 0, 4294967295) {
		fmt.Println("policy route-map", routeMap, "rule", rule,
			"set local-preference", *r.setLocalPreference)
		fmt.Println("local-preference format error.")
		return false
	}
	if r.setMetricType != nil && !validatorInclude(*r.setMetricType, []string{"type-1", "type-2"}) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "set metric-type", *r.setMetricType)
		fmt.Println("Must be (type-1, type-2).")
		return false
	}
	if r.setMetric != nil && !validatorRange(*r.setMetric, -4294967295, 4294967295) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "set metric", *r.setMetric)
		fmt.Println("metric must be an integer with an optional +/- prepend.")
		return false
	}
	if r.setOrigin != nil && !validatorInclude(*r.setOrigin, []string{"igp", "egp", "incomplete"}) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "set origin", *r.setOrigin)
		fmt.Println("origin must be one of igp, egp, or incomplete.")
		return false
	}
	if r.setOriginatorId != nil && !validatorIPv4Address(*r.setOriginatorId) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "set originator-id", *r.setOriginatorId)
		fmt.Println("setOriginatorId format error.")
		return false
	}
	if r.setTag != nil && !validatorRange(*r.setTag, 1, 65535) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "setTag", *r.setTag)
		fmt.Println("tag must be between 1 and 65535.")
		return false
	}
	if r.setWeight != nil && !validatorRange(*r.setWeight, 0, 4294967295) {
		fmt.Println("policy route-map", routeMap, "rule", rule, "setWeight", *r.setWeight)
		fmt.Println("weight format error.")
		return false
	}

	return true
}

func quaggaConfigValidRouteMap(routeMap string) bool {
	nameRegexp := regexp.MustCompile(`^[-a-zA-Z0-9.]+$`)
	if !nameRegexp.MatchString(routeMap) {
		fmt.Println("policy route-map", routeMap)
		fmt.Println("route-map name must be alpha-numeric.")
		return false
	}
	rules := configCandidate.values([]string{"policy", "route-map", routeMap, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidRouteMapRule(routeMap, rule) {
			return false
		}
	}
	return true
}

func quaggaConfigValidRouteMaps() bool {
	routeMaps := configCandidate.values([]string{"policy", "route-map"})
	for _, routeMap := range routeMaps {
		if !quaggaConfigValidRouteMap(routeMap) {
			return false
		}
	}
	return true
}

func quaggaConfigCommitRouteMap(routeMap string) {
	/*
		quaggaVtysh("configure terminal",
			fmt.Sprint("no ipv6 route-map ", routeMap))
	*/
	for _, rule := range configRunning.values([]string{"policy", "route-map", routeMap, "rule"}) {
		if configCandidate.lookup([]string{"policy", "route-map", routeMap, "rule", rule}) == nil {
			action := configRunning.value(
				[]string{"policy", "route-map", routeMap, "rule", rule, "action"})
			if action == nil {
				continue
			}
			quaggaVtysh("configure terminal",
				fmt.Sprint("no route-map ", routeMap, " ", *action, " ", rule))
		}
	}
	rules := configCandidate.values([]string{"policy", "route-map", routeMap, "rule"})
	quaggaRuleSort(rules)
	for _, rule := range rules {
		rRunning := configRunning.makeQuaggaRouteMapRule(routeMap, rule)
		rCandidate := configCandidate.makeQuaggaRouteMapRule(routeMap, rule)
		if rCandidate == nil {
			continue
		}
		if rCandidate.action == nil || !validatorInclude(*rCandidate.action, []string{"permit", "deny"}) {
			continue
		}
		actionChanged := false
		if rRunning != nil && rRunning.action != nil && *rCandidate.action != *rRunning.action {
			actionChanged = true
		}
		quaggaVtysh("configure terminal",
			fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule))
		/* call */
		if !actionChanged && rRunning != nil && rRunning.call != nil && rCandidate.call != nil &&
			*rRunning.call == *rCandidate.call {
			// not changed
		} else if rRunning != nil && rRunning.call != nil && rCandidate.call == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no call")
		} else if rCandidate.call != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("call ", *rCandidate.call))
		}
		/* continue */
		if !actionChanged && rRunning != nil && rRunning.continue_ != nil && rCandidate.continue_ != nil &&
			*rRunning.continue_ == *rCandidate.continue_ {
			// not changed
		} else if rRunning != nil && rRunning.continue_ != nil && rCandidate.continue_ == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no continue")
		} else if rCandidate.continue_ != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("continue ", *rCandidate.continue_))
		}
		/* match as-path */
		if !actionChanged && rRunning != nil && rRunning.matchAsPath != nil && rCandidate.matchAsPath != nil &&
			*rRunning.matchAsPath == *rCandidate.matchAsPath {
			// not changed
		} else if rRunning != nil && rRunning.matchAsPath != nil && rCandidate.matchAsPath == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match as-path")
		} else if rCandidate.matchAsPath != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match as-path ", *rCandidate.matchAsPath))
		}
		/* match community */
		if !actionChanged && rRunning != nil &&
			rRunning.matchCommunityExactMatch == rCandidate.matchCommunityExactMatch &&
			(rRunning.matchCommunityCommunityList != nil &&
				rCandidate.matchCommunityCommunityList != nil &&
				*rRunning.matchCommunityCommunityList == *rCandidate.matchCommunityCommunityList ||
				rRunning.matchCommunityCommunityList == nil &&
					rCandidate.matchCommunityCommunityList == nil) {
			// not changed
		} else {
			if rRunning != nil && rRunning.matchCommunityCommunityList != nil {
				quaggaVtysh("configure terminal",
					fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
					"no match community")
			}
			cond := ""
			if rCandidate.matchCommunityExactMatch {
				cond = " exact-match"
			}
			if rCandidate.matchCommunityCommunityList != nil {
				quaggaVtysh("configure terminal",
					fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
					fmt.Sprint("match community ", *rCandidate.matchCommunityCommunityList, cond))
			}
		}
		/* match interface */
		if !actionChanged && rRunning != nil &&
			rRunning.matchInterface != nil && rCandidate.matchInterface != nil &&
			*rRunning.matchInterface == *rCandidate.matchInterface {
			// not changed
		} else if rRunning != nil && rRunning.matchInterface != nil && rCandidate.matchInterface == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match interface")
		} else if rCandidate.matchInterface != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match interface ", *rCandidate.matchInterface))
		}
		/* match ip address access-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpAddressAccessList != nil && rCandidate.matchIpAddressAccessList != nil &&
			*rRunning.matchIpAddressAccessList == *rCandidate.matchIpAddressAccessList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpAddressAccessList != nil && rCandidate.matchIpAddressAccessList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ip address")
		} else if rCandidate.matchIpAddressAccessList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ip address ", *rCandidate.matchIpAddressAccessList))
		}
		/* match ip address prefix-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpAddressPrefixList != nil && rCandidate.matchIpAddressPrefixList != nil &&
			*rRunning.matchIpAddressPrefixList == *rCandidate.matchIpAddressPrefixList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpAddressPrefixList != nil && rCandidate.matchIpAddressPrefixList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ip address prefix-list")
		} else if rCandidate.matchIpAddressPrefixList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ip address prefix-list ", *rCandidate.matchIpAddressPrefixList))
		}
		/* match ip nexthop access-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpNexthopAccessList != nil && rCandidate.matchIpNexthopAccessList != nil &&
			*rRunning.matchIpNexthopAccessList == *rCandidate.matchIpNexthopAccessList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpNexthopAccessList != nil && rCandidate.matchIpNexthopAccessList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ip nexthop")
		} else if rCandidate.matchIpNexthopAccessList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ip nexthop ", *rCandidate.matchIpNexthopAccessList))
		}
		/* match ip nexthop prefix-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpNexthopPrefixList != nil && rCandidate.matchIpNexthopPrefixList != nil &&
			*rRunning.matchIpNexthopPrefixList == *rCandidate.matchIpNexthopPrefixList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpNexthopPrefixList != nil && rCandidate.matchIpNexthopPrefixList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ip nexthop prefix-list")
		} else if rCandidate.matchIpNexthopPrefixList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ip nexthop prefix-ist ", *rCandidate.matchIpNexthopPrefixList))
		}
		/* match ip route-source access-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpRouteSourceAccessList != nil &&
			rCandidate.matchIpRouteSourceAccessList != nil &&
			*rRunning.matchIpRouteSourceAccessList == *rCandidate.matchIpRouteSourceAccessList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpRouteSourceAccessList != nil &&
			rCandidate.matchIpRouteSourceAccessList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ip route-source")
		} else if rCandidate.matchIpRouteSourceAccessList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ip route-source ", *rCandidate.matchIpRouteSourceAccessList))
		}
		/* match ip route-source prefix-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpRouteSourcePrefixList != nil &&
			rCandidate.matchIpRouteSourcePrefixList != nil &&
			*rRunning.matchIpRouteSourcePrefixList == *rCandidate.matchIpRouteSourcePrefixList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpRouteSourcePrefixList != nil &&
			rCandidate.matchIpRouteSourcePrefixList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ip route-source prefix-ist")
		} else if rCandidate.matchIpRouteSourcePrefixList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ip route-source prefix-list ",
					*rCandidate.matchIpRouteSourcePrefixList))
		}
		/* match ipv6 address access-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpv6AddressAccessList != nil && rCandidate.matchIpv6AddressAccessList != nil &&
			*rRunning.matchIpv6AddressAccessList == *rCandidate.matchIpv6AddressAccessList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpv6AddressAccessList != nil && rCandidate.matchIpv6AddressAccessList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ipv6 address")
		} else if rCandidate.matchIpv6AddressAccessList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ipv6 address ", *rCandidate.matchIpv6AddressAccessList))
		}
		/* match ipv6 address prefix-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpv6AddressPrefixList != nil && rCandidate.matchIpv6AddressPrefixList != nil &&
			*rRunning.matchIpv6AddressPrefixList == *rCandidate.matchIpv6AddressPrefixList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpv6AddressPrefixList != nil && rCandidate.matchIpv6AddressPrefixList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ipv6 address prefix-list")
		} else if rCandidate.matchIpv6AddressPrefixList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ipv6 address prefix-list ", *rCandidate.matchIpv6AddressPrefixList))
		}
		/* match ipv6 nexthop access-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpv6NexthopAccessList != nil && rCandidate.matchIpv6NexthopAccessList != nil &&
			*rRunning.matchIpv6NexthopAccessList == *rCandidate.matchIpv6NexthopAccessList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpv6NexthopAccessList != nil && rCandidate.matchIpv6NexthopAccessList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ipv6 nexthop")
		} else if rCandidate.matchIpv6NexthopAccessList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ipv6 nexthop ", *rCandidate.matchIpv6NexthopAccessList))
		}
		/* match ipv6 nexthop prefix-list */
		if !actionChanged && rRunning != nil &&
			rRunning.matchIpv6NexthopPrefixList != nil && rCandidate.matchIpv6NexthopPrefixList != nil &&
			*rRunning.matchIpv6NexthopPrefixList == *rCandidate.matchIpv6NexthopPrefixList {
			// not changed
		} else if rRunning != nil &&
			rRunning.matchIpv6NexthopPrefixList != nil && rCandidate.matchIpv6NexthopPrefixList == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match ipv6 nexthop prefix-list")
		} else if rCandidate.matchIpv6NexthopPrefixList != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match ipv6 nexthop prefix-list ", *rCandidate.matchIpv6NexthopPrefixList))
		}
		/* match metric */
		if !actionChanged && rRunning != nil && rRunning.matchMetric != nil && rCandidate.matchMetric != nil &&
			*rRunning.matchMetric == *rCandidate.matchMetric {
			// not changed
		} else if rRunning != nil && rRunning.matchMetric != nil && rCandidate.matchMetric == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match metric")
		} else if rCandidate.matchMetric != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match metric ", *rCandidate.matchMetric))
		}
		/* match origin */
		if !actionChanged && rRunning != nil && rRunning.matchOrigin != nil && rCandidate.matchOrigin != nil &&
			*rRunning.matchOrigin == *rCandidate.matchOrigin {
			// not changed
		} else if rRunning != nil && rRunning.matchOrigin != nil && rCandidate.matchOrigin == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match origin")
		} else if rCandidate.matchOrigin != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match origin ", *rCandidate.matchOrigin))
		}
		/* match peer */
		if !actionChanged && rRunning != nil && rRunning.matchPeer != nil && rCandidate.matchPeer != nil &&
			*rRunning.matchPeer == *rCandidate.matchPeer {
			// not changed
		} else if rRunning != nil && rRunning.matchPeer != nil && rCandidate.matchPeer == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match peer")
		} else if rCandidate.matchPeer != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match peer ", *rCandidate.matchPeer))
		}
		/* match tag */
		if !actionChanged && rRunning != nil && rRunning.matchTag != nil && rCandidate.matchTag != nil &&
			*rRunning.matchTag == *rCandidate.matchTag {
			// not changed
		} else if rRunning != nil && rRunning.matchTag != nil && rCandidate.matchTag == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no match tag")
		} else if rCandidate.matchTag != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("match tag ", *rCandidate.matchTag))
		}
		/* on-match goto */
		if !actionChanged && rRunning != nil && rRunning.onMatchGoto != nil && rCandidate.onMatchGoto != nil &&
			*rRunning.onMatchGoto == *rCandidate.onMatchGoto {
			// not changed
		} else if rRunning != nil && rRunning.onMatchGoto != nil && rCandidate.onMatchGoto == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no on-match goto")
		} else if rCandidate.onMatchGoto != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("on-match goto ", *rCandidate.onMatchGoto))
		}
		/* on-match next */
		if !actionChanged && rRunning != nil && rRunning.onMatchNext == rCandidate.onMatchNext {
			// not changed
		} else if rRunning != nil && rRunning.onMatchNext && !rCandidate.onMatchNext {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no on-match next")
		} else if rCandidate.onMatchNext {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"on-match next")
		}
		/* set aggregator */
		if !actionChanged && rRunning != nil &&
			(rRunning.setAggregatorAs != nil && rCandidate.setAggregatorAs != nil &&
				*rRunning.setAggregatorAs == *rCandidate.setAggregatorAs ||
				rRunning.setAggregatorAs == nil && rCandidate.setAggregatorAs == nil) &&
			(rRunning.setAggregatorIp != nil && rCandidate.setAggregatorIp != nil &&
				*rRunning.setAggregatorIp == *rCandidate.setAggregatorIp ||
				rRunning.setAggregatorIp == nil && rCandidate.setAggregatorIp == nil) {
			// not changed
		} else {
			if rRunning != nil && rRunning.setAggregatorAs != nil && rRunning.setAggregatorIp != nil {
				quaggaVtysh("configure terminal",
					fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
					"no set aggregator as")
			}
			if rCandidate.setAggregatorAs != nil && rCandidate.setAggregatorIp != nil {
				as := " " + *rCandidate.setAggregatorAs
				ip := " " + *rCandidate.setAggregatorIp
				quaggaVtysh("configure terminal",
					fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
					fmt.Sprint("set aggregator as", as, ip))
			}
		}
		/* set as-path-prepend */
		if !actionChanged && rRunning != nil &&
			rRunning.setAsPathPrepend != nil && rCandidate.setAsPathPrepend != nil &&
			*rRunning.setAsPathPrepend == *rCandidate.setAsPathPrepend {
			// not changed
		} else if rRunning != nil && rRunning.setAsPathPrepend != nil && rCandidate.setAsPathPrepend == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set as-path prepend")
		} else if rCandidate.setAsPathPrepend != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set as-path prepend ", *rCandidate.setAsPathPrepend))
		}
		/* set atomic-aggregate */
		if !actionChanged && rRunning != nil && rRunning.setAtomicAggregate == rCandidate.setAtomicAggregate {
			// not changed
		} else if rRunning != nil && rRunning.setAtomicAggregate && !rCandidate.setAtomicAggregate {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set atomic-aggregate")
		} else if rCandidate.setAtomicAggregate {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"set atomic-aggregate")
		}
		/* set comm-list */
		if !actionChanged && rRunning != nil &&
			rRunning.setCommListDelete == rCandidate.setCommListDelete &&
			(rRunning.setCommListCommList != nil &&
				rCandidate.setCommListCommList != nil &&
				*rRunning.setCommListCommList == *rCandidate.setCommListCommList ||
				rRunning.setCommListCommList == nil &&
					rCandidate.setCommListCommList == nil) {
			// not changed
		} else {
			if rRunning != nil && rRunning.setCommListCommList != nil {
				quaggaVtysh("configure terminal",
					fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
					"no set comm-list")
			}
			cond := ""
			if rCandidate.setCommListDelete {
				cond = " delete"
			}
			if rCandidate.setCommListCommList != nil {
				quaggaVtysh("configure terminal",
					fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
					fmt.Sprint("set comm-list ", *rCandidate.setCommListCommList, cond))
			}
		}
		/* set community */
		if !actionChanged && rRunning != nil &&
			rRunning.setCommunity != nil && rCandidate.setCommunity != nil &&
			*rRunning.setCommunity == *rCandidate.setCommunity {
			// not changed
		} else if rRunning != nil && rRunning.setCommunity != nil && rCandidate.setCommunity == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set community")
		} else if rCandidate.setCommunity != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set community ", *rCandidate.setCommunity))
		}
		/* set ip-next-hop */
		if !actionChanged && rRunning != nil &&
			rRunning.setIpNextHop != nil && rCandidate.setIpNextHop != nil &&
			*rRunning.setIpNextHop == *rCandidate.setIpNextHop {
			// not changed
		} else if rRunning != nil && rRunning.setIpNextHop != nil && rCandidate.setIpNextHop == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set ip next-hop")
		} else if rCandidate.setIpNextHop != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set ip next-hop ", *rCandidate.setIpNextHop))
		}
		/* set ipv6-next-hop global */
		if !actionChanged && rRunning != nil &&
			rRunning.setIpv6NextHopGlobal != nil && rCandidate.setIpv6NextHopGlobal != nil &&
			*rRunning.setIpv6NextHopGlobal == *rCandidate.setIpv6NextHopGlobal {
			// not changed
		} else if rRunning != nil &&
			rRunning.setIpv6NextHopGlobal != nil && rCandidate.setIpv6NextHopGlobal == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set ipv6 next-hop global")
		} else if rCandidate.setIpv6NextHopGlobal != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set ipv6 next-hop global ", *rCandidate.setIpv6NextHopGlobal))
		}
		/* set ipv6-next-hop local */
		if !actionChanged && rRunning != nil &&
			rRunning.setIpv6NextHopLocal != nil && rCandidate.setIpv6NextHopLocal != nil &&
			*rRunning.setIpv6NextHopLocal == *rCandidate.setIpv6NextHopLocal {
			// not changed
		} else if rRunning != nil &&
			rRunning.setIpv6NextHopLocal != nil && rCandidate.setIpv6NextHopLocal == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set ipv6 next-hop local")
		} else if rCandidate.setIpv6NextHopLocal != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set ipv6 next-hop local ", *rCandidate.setIpv6NextHopLocal))
		}
		/* set local-preference */
		if !actionChanged && rRunning != nil &&
			rRunning.setLocalPreference != nil && rCandidate.setLocalPreference != nil &&
			*rRunning.setLocalPreference == *rCandidate.setLocalPreference {
			// not changed
		} else if rRunning != nil &&
			rRunning.setLocalPreference != nil && rCandidate.setLocalPreference == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set local-preference")
		} else if rCandidate.setLocalPreference != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set local-preference ", *rCandidate.setLocalPreference))
		}
		/* set metric-type */
		if !actionChanged && rRunning != nil &&
			rRunning.setMetricType != nil && rCandidate.setMetricType != nil &&
			*rRunning.setMetricType == *rCandidate.setMetricType {
			// not changed
		} else if rRunning != nil &&
			rRunning.setMetricType != nil && rCandidate.setMetricType == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set metric-type")
		} else if rCandidate.setMetricType != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set metric-type ", *rCandidate.setMetricType))
		}
		/* set metric */
		if !actionChanged && rRunning != nil && rRunning.setMetric != nil && rCandidate.setMetric != nil &&
			*rRunning.setMetric == *rCandidate.setMetric {
			// not changed
		} else if rRunning != nil && rRunning.setMetric != nil && rCandidate.setMetric == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set metric")
		} else if rCandidate.setMetric != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set metric ", *rCandidate.setMetric))
		}
		/* set origin */
		if !actionChanged && rRunning != nil && rRunning.setOrigin != nil && rCandidate.setOrigin != nil &&
			*rRunning.setOrigin == *rCandidate.setOrigin {
			// not changed
		} else if rRunning != nil && rRunning.setOrigin != nil && rCandidate.setOrigin == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set origin")
		} else if rCandidate.setOrigin != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set origin ", *rCandidate.setOrigin))
		}
		/* set originator-id */
		if !actionChanged && rRunning != nil &&
			rRunning.setOriginatorId != nil && rCandidate.setOriginatorId != nil &&
			*rRunning.setOriginatorId == *rCandidate.setOriginatorId {
			// not changed
		} else if rRunning != nil &&
			rRunning.setOriginatorId != nil && rCandidate.setOriginatorId == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set originator-id")
		} else if rCandidate.setOriginatorId != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set originator-id ", *rCandidate.setOriginatorId))
		}
		/* set tag */
		if !actionChanged && rRunning != nil && rRunning.setTag != nil && rCandidate.setTag != nil &&
			*rRunning.setTag == *rCandidate.setTag {
			// not changed
		} else if rRunning != nil && rRunning.setTag != nil && rCandidate.setTag == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set tag")
		} else if rCandidate.setTag != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set tag ", *rCandidate.setTag))
		}
		/* set weight */
		if !actionChanged && rRunning != nil && rRunning.setWeight != nil && rCandidate.setWeight != nil &&
			*rRunning.setWeight == *rCandidate.setWeight {
			// not changed
		} else if rRunning != nil && rRunning.setWeight != nil && rCandidate.setWeight == nil {
			// deleted
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				"no set weight")
		} else if rCandidate.setWeight != nil {
			// updated
			quaggaVtysh("configure terminal",
				fmt.Sprint("route-map ", routeMap, " ", *rCandidate.action, " ", rule),
				fmt.Sprint("set weight ", *rCandidate.setWeight))
		}
	}
}

func quaggaConfigCommitRouteMaps() {
	for _, routeMap := range configRunning.values([]string{"policy", "route-map"}) {
		if configCandidate.lookup([]string{"policy", "route-map", routeMap}) == nil {
			quaggaVtysh("configure terminal",
				fmt.Sprint("no route-map ", routeMap))
		}
	}
	for _, routeMap := range configCandidate.values([]string{"policy", "route-map"}) {
		if !quaggaConfigStateDiff([]string{"policy", "route-map", routeMap}) {
			continue
		}
		quaggaConfigCommitRouteMap(routeMap)
	}
}

func quaggaUpdateCheckRouteMap() {
	if !commitedRouteMap && quaggaConfigStateDiff([]string{"policy", "route-map"}) {
		commitedRouteMap = true
		quaggaConfigCommitRouteMaps()
	}
}

/*
 *
 */

func quaggaConfigValidPolicy() bool {
	if !quaggaConfigValidAccessLists() {
		return false
	}
	if !quaggaConfigValidAccessList6s() {
		return false
	}
	if !quaggaConfigValidAsPathLists() {
		return false
	}
	if !quaggaConfigValidCommunityLists() {
		return false
	}
	if !quaggaConfigValidPrefixLists() {
		return false
	}
	if !quaggaConfigValidPrefixList6s() {
		return false
	}
	if !quaggaConfigValidRouteMaps() {
		return false
	}
	return true
}
