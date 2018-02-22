package quagga

import (
	"fmt"
	"io"
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

func quaggaConfigValidAccessListRule(f io.Writer, accessList, rule string) bool {
	valid := true
	if !validatorRange(rule, 1, 65535) {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule)
		fmt.Fprintln(f, "rule number must be between 1 and 65535.")
		valid = false
		return valid
	}
	r := configCandidate.makeQuaggaAccessListRule(accessList, rule)
	if r == nil {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule)
		fmt.Fprintln(f, "rule not found.")
		valid = false
		return valid
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "action", action)
		fmt.Fprintln(f, "action must be permit or deny.")
		valid = false
	}
	if r.srcHost != nil && !validatorIPv4Address(*r.srcHost) {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "source host", *r.srcHost)
		fmt.Fprintln(f, "source host format error.")
		valid = false
	}
	if r.srcInverseMask != nil && !validatorIPv4Address(*r.srcInverseMask) {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "source inverse-mask", *r.srcInverseMask)
		fmt.Fprintln(f, "source inverse-mask format error.")
		valid = false
	}
	if r.srcNetwork != nil && !validatorIPv4Address(*r.srcNetwork) {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "source network", *r.srcNetwork)
		fmt.Fprintln(f, "source network format error.")
		valid = false
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
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "source")
		fmt.Fprintln(f, "you may only define one filter type (host|network|any).")
		valid = false
	}
	if srcMatches > 1 {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "source")
		fmt.Fprintln(f, "you may only define one filter type (host|network|any).")
		valid = false
	}
	if r.srcNetwork != nil && r.srcInverseMask == nil {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "source")
		fmt.Fprintln(f, "you must specify an inverse-mask if you configure a network.")
		valid = false
	}
	if r.srcNetwork == nil && r.srcInverseMask != nil {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "source")
		fmt.Fprintln(f, "you must specify a network if you configure an inverse mask.")
		valid = false
	}
	if r.dstHost != nil && !validatorIPv4Address(*r.dstHost) {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "destination host", *r.dstHost)
		fmt.Fprintln(f, "destination host format error.")
		valid = false
	}
	if r.dstInverseMask != nil && !validatorIPv4Address(*r.dstInverseMask) {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule,
			"destination inverse-mask", *r.dstInverseMask)
		fmt.Fprintln(f, "destination inverse-mask format error.")
		valid = false
	}
	if r.dstNetwork != nil && !validatorIPv4Address(*r.dstNetwork) {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "destination network", *r.dstNetwork)
		fmt.Fprintln(f, "destination network format error.")
		valid = false
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
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "destination")
		fmt.Fprintln(f, "access-list number must be <100-199> or <2000-2699> to set destination matches.")
		valid = false
	}
	if dstMatches == 0 && (validatorRange(accessList, 100, 199) || validatorRange(accessList, 2000, 2699)) {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "destination")
		fmt.Fprintln(f, "you may only define one filter type (host|network|any).")
		valid = false
	}
	if dstMatches > 1 {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "destination")
		fmt.Fprintln(f, "you may only define one filter type (host|network|any).")
		valid = false
	}
	if r.dstNetwork != nil && r.dstInverseMask == nil {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "destination")
		fmt.Fprintln(f, "you must specify an inverse-mask if you configure a network.")
		valid = false
	}
	if r.dstNetwork == nil && r.dstInverseMask != nil {
		fmt.Fprintln(f, "policy access-list", accessList, "rule", rule, "destination")
		fmt.Fprintln(f, "you must specify a network if you configure an inverse mask.")
		valid = false
	}
	return valid
}

func quaggaConfigValidAccessList(f io.Writer, accessList string) bool {
	valid := true
	if !validatorRange(accessList, 1, 199) && !validatorRange(accessList, 1300, 2699) {
		fmt.Fprintln(f, "policy access-list", accessList)
		fmt.Fprintln(f, "Access list number must be:")
		fmt.Fprintln(f, "<1-99>      IP standard access list")
		fmt.Fprintln(f, "<100-199>   IP extended access list")
		fmt.Fprintln(f, "<1300-1999> IP standard access list (expanded range)")
		fmt.Fprintln(f, "<2000-2699> IP extended access list (expanded range)")
		valid = false
		return valid
	}
	rules := configCandidate.values([]string{"policy", "access-list", accessList, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidAccessListRule(f, accessList, rule) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidAccessLists(f io.Writer) bool {
	valid := true
	accessLists := configCandidate.values([]string{"policy", "access-list"})
	for _, accessList := range accessLists {
		if !quaggaConfigValidAccessList(f, accessList) {
			valid = false
		}
	}
	return valid
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

func quaggaConfigValidAccessList6Rule(f io.Writer, accessList6, rule string) bool {
	valid := true
	if !validatorRange(rule, 1, 65535) {
		fmt.Fprintln(f, "policy access-list6", accessList6, "rule", rule)
		fmt.Fprintln(f, "rule number must be between 1 and 65535.")
		valid = false
		return valid
	}
	r := configCandidate.makeQuaggaAccessList6Rule(accessList6, rule)
	if r == nil {
		fmt.Fprintln(f, "policy access-list6", accessList6, "rule", rule)
		fmt.Fprintln(f, "rule not found.")
		valid = false
		return valid
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Fprintln(f, "policy access-list6", accessList6, "rule", rule, "action", action)
		fmt.Fprintln(f, "action must be permit or deny.")
		valid = false
	}
	if r.srcNetwork != nil && !validatorIPv6CIDR(*r.srcNetwork) {
		fmt.Fprintln(f, "policy access-list6", accessList6, "rule", rule, "source network", *r.srcNetwork)
		fmt.Fprintln(f, "source network format error.")
		valid = false
	}
	srcMatches := 0
	if r.srcAny == true {
		srcMatches++
	}
	if r.srcNetwork != nil {
		srcMatches++
	}
	if srcMatches == 0 {
		fmt.Fprintln(f, "policy access-list6", accessList6, "rule", rule, "source")
		fmt.Fprintln(f, "you may only define one filter type (network|any).")
		valid = false
	}
	if srcMatches > 1 {
		fmt.Fprintln(f, "policy access-list6", accessList6, "rule", rule, "source")
		fmt.Fprintln(f, "you may only define one filter type (network|any).")
		valid = false
	}
	return valid
}

func quaggaConfigValidAccessList6(f io.Writer, accessList6 string) bool {
	valid := true
	if len(accessList6) < 1 || len(accessList6) > 64 {
		fmt.Fprintln(f, "policy access-list6", accessList6)
		fmt.Fprintln(f, "access-list name must be 64 characters or less.")
		valid = false
	}
	if accessList6[0] == '-' {
		fmt.Fprintln(f, "policy access-list6", accessList6)
		fmt.Fprintln(f, "access-list name cannot start with \"-\".")
		valid = false
	}
	nameRegexp := regexp.MustCompile(`^[^|;&$<>]*$`)
	if !nameRegexp.MatchString(accessList6) {
		fmt.Fprintln(f, "policy access-list6", accessList6)
		fmt.Fprintln(f, "access-list name cannot contain shell punctuation.")
		valid = false
	}
	rules := configCandidate.values([]string{"policy", "access-list6", accessList6, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidAccessList6Rule(f, accessList6, rule) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidAccessList6s(f io.Writer) bool {
	valid := true
	accessList6s := configCandidate.values([]string{"policy", "access-list6"})
	for _, accessList6 := range accessList6s {
		if !quaggaConfigValidAccessList6(f, accessList6) {
			valid = false
		}
	}
	return valid
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

func quaggaConfigValidAsPathListRule(f io.Writer, asPathList, rule string) bool {
	valid := true
	if !validatorRange(rule, 1, 65535) {
		fmt.Fprintln(f, "policy as-path-list", asPathList, "rule", rule)
		fmt.Fprintln(f, "rule number must be between 1 and 65535.")
		valid = false
		return valid
	}
	r := configCandidate.makeQuaggaAsPathListRule(asPathList, rule)
	if r == nil {
		fmt.Fprintln(f, "policy as-path-list", asPathList, "rule", rule)
		fmt.Fprintln(f, "rule not found.")
		valid = false
		return valid
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Fprintln(f, "policy as-path-list", asPathList, "rule", rule, "action", action)
		fmt.Fprintln(f, "action must be permit or deny.")
		valid = false
	}
	if r.regex == nil {
		fmt.Fprintln(f, "policy as-path-list", asPathList, "rule", rule, "regex")
		fmt.Fprintln(f, "you must specify a regex.")
		valid = false
	}
	return valid
}

func quaggaConfigValidAsPathList(f io.Writer, asPathList string) bool {
	valid := true
	nameRegexp := regexp.MustCompile(`^[-a-zA-Z0-9.]+$`)
	if !nameRegexp.MatchString(asPathList) {
		fmt.Fprintln(f, "policy as-path-list", asPathList)
		fmt.Fprintln(f, "as-path-list name must be alpha-numeric.")
		valid = false
	}
	rules := configCandidate.values([]string{"policy", "as-path-list", asPathList, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidAsPathListRule(f, asPathList, rule) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidAsPathLists(f io.Writer) bool {
	valid := true
	asPathLists := configCandidate.values([]string{"policy", "as-path-list"})
	for _, asPathList := range asPathLists {
		if !quaggaConfigValidAsPathList(f, asPathList) {
			valid = false
		}
	}
	return valid
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

func quaggaConfigValidCommunityListRule(f io.Writer, communityList, rule string) bool {
	valid := true
	if !validatorRange(rule, 1, 65535) {
		fmt.Fprintln(f, "policy community-list", communityList, "rule", rule)
		fmt.Fprintln(f, "rule number must be between 1 and 65535.")
		valid = false
		return valid
	}
	r := configCandidate.makeQuaggaCommunityListRule(communityList, rule)
	if r == nil {
		fmt.Fprintln(f, "policy community-list", communityList, "rule", rule)
		fmt.Fprintln(f, "rule not found.")
		valid = false
		return valid
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Fprintln(f, "policy community-list", communityList, "rule", rule, "action", action)
		fmt.Fprintln(f, "action must be permit or deny.")
		valid = false
	}
	if r.regex == nil {
		fmt.Fprintln(f, "policy community-list", communityList, "rule", rule, "regex")
		fmt.Fprintln(f, "you must specify a regex.")
		valid = false
	}
	if r.regex != nil && validatorRange(rule, 1, 99) {
		stdRegexp := regexp.MustCompile(`^(internet|local-AS|no-advertise|no-export|\d+:\d+)$`)
		if !stdRegexp.MatchString(*r.regex) {
			fmt.Fprintln(f, "policy community-list", communityList, "rule", rule, "regex", *r.regex)
			fmt.Fprintln(f, "regex", *r.regex, "is invalid for a standard community list.")
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidCommunityList(f io.Writer, communityList string) bool {
	valid := true
	if !validatorRange(communityList, 1, 500) {
		fmt.Fprintln(f, "policy community-list", communityList)
		fmt.Fprintln(f, "community-list must be:")
		fmt.Fprintln(f, "<1-99>    BGP community list (standard)")
		fmt.Fprintln(f, "<100-500> BGP community list (expanded)")
		valid = false
	}
	rules := configCandidate.values([]string{"policy", "community-list", communityList, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidCommunityListRule(f, communityList, rule) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidCommunityLists(f io.Writer) bool {
	valid := true
	communityLists := configCandidate.values([]string{"policy", "community-list"})
	for _, communityList := range communityLists {
		if !quaggaConfigValidCommunityList(f, communityList) {
			valid = false
		}
	}
	return valid
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

func quaggaConfigValidPrefixListRule(f io.Writer, prefixList, rule string) bool {
	valid := true
	if !validatorRange(rule, 1, 65535) {
		fmt.Fprintln(f, "policy prefix-list", prefixList, "rule", rule)
		fmt.Fprintln(f, "rule number must be between 1 and 65535.")
		valid = false
		return valid
	}
	r := configCandidate.makeQuaggaPrefixListRule(prefixList, rule)
	if r == nil {
		fmt.Fprintln(f, "policy prefix-list", prefixList, "rule", rule)
		fmt.Fprintln(f, "rule not found.")
		valid = false
		return valid
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Fprintln(f, "policy prefix-list", prefixList, "rule", rule, "action", action)
		fmt.Fprintln(f, "action must be permit or deny.")
		valid = false
	}
	if r.le != nil && !validatorRange(*r.le, 0, 32) {
		le := ""
		if r.le != nil {
			le = *r.le
		}
		fmt.Fprintln(f, "policy prefix-list", prefixList, "rule", rule, "le", le)
		fmt.Fprintln(f, "le must be between 0 and 32.")
		valid = false
	}
	if r.ge != nil && !validatorRange(*r.ge, 0, 32) {
		ge := ""
		if r.ge != nil {
			ge = *r.ge
		}
		fmt.Fprintln(f, "policy prefix-list", prefixList, "rule", rule, "ge", ge)
		fmt.Fprintln(f, "ge must be between 0 and 32.")
		valid = false
	}
	if r.prefix == nil || !validatorIPv4CIDR(*r.prefix) {
		prefix := ""
		if r.prefix != nil {
			prefix = *r.prefix
		}
		fmt.Fprintln(f, "policy prefix-list", prefixList, "rule", rule, "prefix", prefix)
		fmt.Fprintln(f, "you must specify a prefix.")
		valid = false
	}
	return valid
}

func quaggaConfigValidPrefixList(f io.Writer, prefixList string) bool {
	valid := true
	nameRegexp := regexp.MustCompile(`^[-a-zA-Z0-9.]+$`)
	if !nameRegexp.MatchString(prefixList) {
		fmt.Fprintln(f, "policy prefix-list", prefixList)
		fmt.Fprintln(f, "prefix-list name must be alpha-numeric.")
		valid = false
	}
	rules := configCandidate.values([]string{"policy", "prefix-list", prefixList, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidPrefixListRule(f, prefixList, rule) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidPrefixLists(f io.Writer) bool {
	valid := true
	prefixLists := configCandidate.values([]string{"policy", "prefix-list"})
	for _, prefixList := range prefixLists {
		if !quaggaConfigValidPrefixList(f, prefixList) {
			valid = false
		}
	}
	return valid
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

func quaggaConfigValidPrefixList6Rule(f io.Writer, prefixList6, rule string) bool {
	valid := true
	if !validatorRange(rule, 1, 65535) {
		fmt.Fprintln(f, "policy prefix-list6", prefixList6, "rule", rule)
		fmt.Fprintln(f, "rule number must be between 1 and 65535.")
		valid = false
		return valid
	}
	r := configCandidate.makeQuaggaPrefixList6Rule(prefixList6, rule)
	if r == nil {
		fmt.Fprintln(f, "policy prefix-list6", prefixList6, "rule", rule)
		fmt.Fprintln(f, "rule not found.")
		valid = false
		return valid
	}
	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Fprintln(f, "policy prefix-list6", prefixList6, "rule", rule, "action", action)
		fmt.Fprintln(f, "action must be permit or deny.")
		valid = false
	}
	if r.le != nil && !validatorRange(*r.le, 0, 32) {
		le := ""
		if r.le != nil {
			le = *r.le
		}
		fmt.Fprintln(f, "policy prefix-list6", prefixList6, "rule", rule, "le", le)
		fmt.Fprintln(f, "le must be between 0 and 32.")
		valid = false
	}
	if r.ge != nil && !validatorRange(*r.ge, 0, 32) {
		ge := ""
		if r.ge != nil {
			ge = *r.ge
		}
		fmt.Fprintln(f, "policy prefix-list6", prefixList6, "rule", rule, "ge", ge)
		fmt.Fprintln(f, "ge must be between 0 and 32.")
		valid = false
	}
	if r.prefix == nil || !validatorIPv6CIDR(*r.prefix) {
		prefix := ""
		if r.prefix != nil {
			prefix = *r.prefix
		}
		fmt.Fprintln(f, "policy prefix-list6", prefixList6, "rule", rule, "prefix", prefix)
		fmt.Fprintln(f, "you must specify a prefix.")
		valid = false
	}
	return valid
}

func quaggaConfigValidPrefixList6(f io.Writer, prefixList6 string) bool {
	valid := true
	nameRegexp := regexp.MustCompile(`^[-a-zA-Z0-9.]+$`)
	if !nameRegexp.MatchString(prefixList6) {
		fmt.Fprintln(f, "policy prefix-list6", prefixList6)
		fmt.Fprintln(f, "prefix-list6 name must be alpha-numeric.")
		valid = false
	}
	rules := configCandidate.values([]string{"policy", "prefix-list6", prefixList6, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidPrefixList6Rule(f, prefixList6, rule) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidPrefixList6s(f io.Writer) bool {
	valid := true
	prefixList6s := configCandidate.values([]string{"policy", "prefix-list6"})
	for _, prefixList6 := range prefixList6s {
		if !quaggaConfigValidPrefixList6(f, prefixList6) {
			valid = false
		}
	}
	return valid
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

func quaggaConfigValidRouteMapRule(f io.Writer, routeMap, rule string) bool {
	valid := true
	if !validatorRange(rule, 1, 65535) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule)
		fmt.Fprintln(f, "rule number must be between 1 and 65535.")
		valid = false
		return valid
	}
	r := configCandidate.makeQuaggaRouteMapRule(routeMap, rule)
	if r == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule)
		fmt.Fprintln(f, "rule not found.")
		valid = false
		return valid
	}

	if r.action == nil || !validatorInclude(*r.action, []string{"permit", "deny"}) {
		action := ""
		if r.action != nil {
			action = *r.action
		}
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "action", action)
		fmt.Fprintln(f, "action must be permit or deny.")
		valid = false
	}

	if r.call != nil && configCandidate.lookup([]string{"policy", "route-map", *r.call}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "call", *r.call)
		fmt.Fprintln(f, "called route-map", *r.call, "doesn't exist.")
		valid = false
	}
	if r.continue_ != nil {
		if !validatorRange(*r.continue_, 1, 65535) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "continue", *r.continue_)
			fmt.Fprintln(f, "continue must be between 1 and 65535.")
			valid = false
		}
		from, _ := strconv.Atoi(rule)
		to, _ := strconv.Atoi(*r.continue_)
		if !(to > from) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "continue", *r.continue_)
			fmt.Fprintln(f, "you may only continue forward in the route-map.")
			valid = false
		}
	}
	if r.matchAsPath != nil &&
		configCandidate.lookup([]string{"policy", "as-path-list", *r.matchAsPath}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "match as-path", *r.matchAsPath)
		fmt.Fprintln(f, "match as-path: AS path list", *r.matchAsPath, "doesn't exist.")
		valid = false
	}
	if r.matchCommunityCommunityList != nil &&
		configCandidate.lookup([]string{"policy", "community-list", *r.matchCommunityCommunityList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match community community-list", *r.matchCommunityCommunityList)
		fmt.Fprintln(f, "community-list", *r.matchCommunityCommunityList, "doesn't exist.")
		valid = false
	}
	/*
		if r.matchCommunityExactMatch != nil && !validator(*r.matchCommunityExactMatch) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
				"matchCommunityExactMatch", *r.matchCommunityExactMatch)
			fmt.Fprintln(f, "matchCommunityExactMatch format error.")
			valid = false
		}
	*/
	/*
		XXX:
		if r.matchInterface != nil && !validator(*r.matchInterface) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
				"matchInterface", *r.matchInterface)
			fmt.Fprintln(f, "matchInterface format error.")
			valid = false
		}
	*/
	if r.matchIpAddressAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list", *r.matchIpAddressAccessList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip address access-list", *r.matchIpAddressAccessList)
		fmt.Fprintln(f, "access-list", *r.matchIpAddressAccessList, "does not exist.")
		valid = false
	}
	if r.matchIpAddressPrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list", *r.matchIpAddressPrefixList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip address prefix-list", *r.matchIpAddressPrefixList)
		fmt.Fprintln(f, "prefix-list", *r.matchIpAddressPrefixList, "does not exist.")
		valid = false
	}
	if r.matchIpAddressAccessList != nil && r.matchIpAddressPrefixList != nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip address access-list", *r.matchIpAddressAccessList)
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip address prefix-list", *r.matchIpAddressPrefixList)
		fmt.Fprintln(f, "you may only specify a prefix-list or access-list.")
		valid = false
	}
	if r.matchIpNexthopAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list", *r.matchIpNexthopAccessList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip nexthop access-list", *r.matchIpNexthopAccessList)
		fmt.Fprintln(f, "access-list", *r.matchIpNexthopAccessList, "does not exist.")
		valid = false
	}
	if r.matchIpNexthopPrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list", *r.matchIpNexthopPrefixList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip nexthop prefix-list", *r.matchIpNexthopPrefixList)
		fmt.Fprintln(f, "prefix-list", *r.matchIpNexthopPrefixList, "does not exist.")
		valid = false
	}
	if r.matchIpNexthopAccessList != nil && r.matchIpNexthopPrefixList != nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip nexthop access-list", *r.matchIpNexthopAccessList)
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip nexthop prefix-list", *r.matchIpNexthopPrefixList)
		fmt.Fprintln(f, "you may only specify a prefix-list or access-list.")
		valid = false
	}
	if r.matchIpRouteSourceAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list", *r.matchIpRouteSourceAccessList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip route-source access-list", *r.matchIpRouteSourceAccessList)
		fmt.Fprintln(f, "access-list", *r.matchIpRouteSourceAccessList, "does not exist.")
		valid = false
	}
	if r.matchIpRouteSourcePrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list", *r.matchIpRouteSourcePrefixList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip route-source prefix-list", *r.matchIpRouteSourcePrefixList)
		fmt.Fprintln(f, "prefix-list", *r.matchIpRouteSourcePrefixList, "does not exist.")
		valid = false
	}
	if r.matchIpRouteSourceAccessList != nil && r.matchIpRouteSourcePrefixList != nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip route-source access-list", *r.matchIpRouteSourceAccessList)
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ip route-source prefix-list", *r.matchIpRouteSourcePrefixList)
		fmt.Fprintln(f, "you may only specify a prefix-list or access-list.")
		valid = false
	}
	if r.matchIpv6AddressAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list6", *r.matchIpv6AddressAccessList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ipv6 address access-list", *r.matchIpv6AddressAccessList)
		fmt.Fprintln(f, "access-list6", *r.matchIpv6AddressAccessList, "does not exist.")
		valid = false
	}
	if r.matchIpv6AddressPrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list6", *r.matchIpv6AddressPrefixList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ipv6 address prefix-list", *r.matchIpv6AddressPrefixList)
		fmt.Fprintln(f, "prefix-list6", *r.matchIpv6AddressPrefixList, "does not exist.")
		valid = false
	}
	if r.matchIpv6AddressAccessList != nil && r.matchIpv6AddressPrefixList != nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ipv6 address access-list", *r.matchIpv6AddressAccessList)
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ipv6 address prefix-list", *r.matchIpv6AddressPrefixList)
		fmt.Fprintln(f, "you may only specify a prefix-list or access-list.")
		valid = false
	}
	if r.matchIpv6NexthopAccessList != nil &&
		configCandidate.lookup([]string{"policy", "access-list6", *r.matchIpv6NexthopAccessList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ipv6 nexthop access-list", *r.matchIpv6NexthopAccessList)
		fmt.Fprintln(f, "access-list6", *r.matchIpv6NexthopAccessList, "does not exist.")
		valid = false
	}
	if r.matchIpv6NexthopPrefixList != nil &&
		configCandidate.lookup([]string{"policy", "prefix-list6", *r.matchIpv6NexthopPrefixList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ipv6 nexthop prefix-list", *r.matchIpv6NexthopPrefixList)
		fmt.Fprintln(f, "prefix-list6", *r.matchIpv6NexthopPrefixList, "does not exist.")
		valid = false
	}
	if r.matchIpv6NexthopAccessList != nil && r.matchIpv6NexthopPrefixList != nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ipv6 nexthop access-list", *r.matchIpv6NexthopAccessList)
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"match ipv6 nexthop prefix-list", *r.matchIpv6NexthopPrefixList)
		fmt.Fprintln(f, "you may only specify a prefix-list or access-list.")
		valid = false
	}
	if r.matchMetric != nil && !validatorRange(*r.matchMetric, 1, 65535) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "match metric", *r.matchMetric)
		fmt.Fprintln(f, "metric must be between 1 and 65535.")
		valid = false
	}
	if r.matchOrigin != nil && !validatorInclude(*r.matchOrigin, []string{"egp", "igp", "incomplete"}) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "match origin", *r.matchOrigin)
		fmt.Fprintln(f, "origin must be egp, igp, or incomplete.")
		valid = false
	}
	if r.matchPeer != nil && !validatorPeer(*r.matchPeer) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "match peer", *r.matchPeer)
		fmt.Fprintln(f, "peer must be either an IP or local.")
		valid = false
	}
	if r.matchTag != nil && !validatorRange(*r.matchTag, 1, 65535) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "match tag", *r.matchTag)
		fmt.Fprintln(f, "tag must be between 1 and 65535.")
		valid = false
	}
	if r.onMatchGoto != nil {
		if !validatorRange(*r.onMatchGoto, 1, 65535) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "on-match goto", *r.onMatchGoto)
			fmt.Fprintln(f, "goto must be a rule number between 1 and 65535.")
			valid = false
		}
		from, _ := strconv.Atoi(rule)
		to, _ := strconv.Atoi(*r.onMatchGoto)
		if !(to > from) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "on-match goto", *r.onMatchGoto)
			fmt.Fprintln(f, "you may only go forward in the route-map.")
			valid = false
		}
	}
	/*
		if r.onMatchNext != nil && !validator(*r.onMatchNext) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "on-match next", *r.onMatchNext)
			fmt.Fprintln(f, "onMatchNext format error.")
			valid = false
		}
	*/
	if r.onMatchGoto != nil && r.onMatchNext {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "on-match goto", *r.onMatchGoto)
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "on-match next")
		fmt.Fprintln(f, "you may set only goto or next.")
		valid = false
	}
	if r.setAggregatorAs != nil && !validatorRange(*r.setAggregatorAs, 1, 65535) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set aggregator as", *r.setAggregatorAs)
		fmt.Fprintln(f, "BGP AS number must be between 1 and 4294967294.")
		valid = false
	}
	if r.setAggregatorIp != nil && !validatorIPv4Address(*r.setAggregatorIp) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set aggregator ip", *r.setAggregatorIp)
		fmt.Fprintln(f, "aggregator IP format error.")
		valid = false
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
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set aggregator as", setAggregatorAs)
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set aggregator ip", setAggregatorIp)
		fmt.Fprintln(f, "you must configure both as and ip.")
		valid = false
	}
	if r.setAsPathPrepend != nil && !validatorAsPathPrepend(*r.setAsPathPrepend) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "setAsPathPrepend", *r.setAsPathPrepend)
		fmt.Fprintln(f, "invalid AS path string.")
		valid = false
	}
	/*
		if r.setAtomicAggregate != nil && !validator(*r.setAtomicAggregate) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
				"setAtomicAggregate", *r.setAtomicAggregate)
			fmt.Fprintln(f, "setAtomicAggregate format error.")
			valid = false
		}
	*/
	if r.setCommListCommList != nil &&
		configCandidate.lookup([]string{"policy", "community-list", *r.setCommListCommList}) == nil {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"set comm-list comm-list", *r.setCommListCommList)
		fmt.Fprintln(f, "community list", *r.setCommListCommList, "does not exist.")
		valid = false
	}
	/*
		if r.setCommListDelete != nil && !validator(*r.setCommListDelete) {
			fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
				"setCommListDelete", *r.setCommListDelete)
			fmt.Fprintln(f, "setCommListDelete format error.")
			valid = false
		}
	*/
	if r.setCommunity != nil && !validatorCommunity(*r.setCommunity) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set community", *r.setCommunity)
		fmt.Fprintln(f, "community format error.")
		valid = false
	}
	if r.setIpNextHop != nil && !validatorIPv4Address(*r.setIpNextHop) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set ip-next-hop", *r.setIpNextHop)
		fmt.Fprintln(f, "ip-next-hop format error.")
		valid = false
	}
	if r.setIpv6NextHopGlobal != nil && !validatorIPv6Address(*r.setIpv6NextHopGlobal) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"set ipv6-next-hop global", *r.setIpv6NextHopGlobal)
		fmt.Fprintln(f, "ipv6-next-hop global format error.")
		valid = false
	}
	if r.setIpv6NextHopLocal != nil && !validatorIPv6Address(*r.setIpv6NextHopLocal) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"set ipv6-next-hop local", *r.setIpv6NextHopLocal)
		fmt.Fprintln(f, "ipv6-next-hop local format error.")
		valid = false
	}
	if r.setLocalPreference != nil && !validatorRange(*r.setLocalPreference, 0, 4294967295) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule,
			"set local-preference", *r.setLocalPreference)
		fmt.Fprintln(f, "local-preference format error.")
		valid = false
	}
	if r.setMetricType != nil && !validatorInclude(*r.setMetricType, []string{"type-1", "type-2"}) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set metric-type", *r.setMetricType)
		fmt.Fprintln(f, "Must be (type-1, type-2).")
		valid = false
	}
	if r.setMetric != nil && !validatorRange(*r.setMetric, -4294967295, 4294967295) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set metric", *r.setMetric)
		fmt.Fprintln(f, "metric must be an integer with an optional +/- prepend.")
		valid = false
	}
	if r.setOrigin != nil && !validatorInclude(*r.setOrigin, []string{"igp", "egp", "incomplete"}) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set origin", *r.setOrigin)
		fmt.Fprintln(f, "origin must be one of igp, egp, or incomplete.")
		valid = false
	}
	if r.setOriginatorId != nil && !validatorIPv4Address(*r.setOriginatorId) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "set originator-id", *r.setOriginatorId)
		fmt.Fprintln(f, "setOriginatorId format error.")
		valid = false
	}
	if r.setTag != nil && !validatorRange(*r.setTag, 1, 65535) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "setTag", *r.setTag)
		fmt.Fprintln(f, "tag must be between 1 and 65535.")
		valid = false
	}
	if r.setWeight != nil && !validatorRange(*r.setWeight, 0, 4294967295) {
		fmt.Fprintln(f, "policy route-map", routeMap, "rule", rule, "setWeight", *r.setWeight)
		fmt.Fprintln(f, "weight format error.")
		valid = false
	}

	return valid
}

func quaggaConfigValidRouteMap(f io.Writer, routeMap string) bool {
	valid := true
	nameRegexp := regexp.MustCompile(`^[-a-zA-Z0-9.]+$`)
	if !nameRegexp.MatchString(routeMap) {
		fmt.Fprintln(f, "policy route-map", routeMap)
		fmt.Fprintln(f, "route-map name must be alpha-numeric.")
		valid = false
	}
	rules := configCandidate.values([]string{"policy", "route-map", routeMap, "rule"})
	for _, rule := range rules {
		if !quaggaConfigValidRouteMapRule(f, routeMap, rule) {
			valid = false
		}
	}
	return valid
}

func quaggaConfigValidRouteMaps(f io.Writer) bool {
	valid := true
	routeMaps := configCandidate.values([]string{"policy", "route-map"})
	for _, routeMap := range routeMaps {
		if !quaggaConfigValidRouteMap(f, routeMap) {
			valid = false
		}
	}
	return valid
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

func quaggaConfigValidPolicy(f io.Writer) bool {
	valid := true
	if !quaggaConfigValidAccessLists(f) {
		valid = false
	}
	if !quaggaConfigValidAccessList6s(f) {
		valid = false
	}
	if !quaggaConfigValidAsPathLists(f) {
		valid = false
	}
	if !quaggaConfigValidCommunityLists(f) {
		valid = false
	}
	if !quaggaConfigValidPrefixLists(f) {
		valid = false
	}
	if !quaggaConfigValidPrefixList6s(f) {
		valid = false
	}
	if !quaggaConfigValidRouteMaps(f) {
		valid = false
	}
	return valid
}
