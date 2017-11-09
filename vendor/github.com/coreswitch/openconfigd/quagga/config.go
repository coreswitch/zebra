package quagga

import (
	"fmt"
	"github.com/coreswitch/cmd"
)

var (
	configParser *cmd.Node
)

/*
	tag:
	type: u32
	help: MD5 key id

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "ID must be between (1-255)"
	val_help: u32:1-255; MD5 key id

	commit:expression: $VAR(md5-key/) != ""; \
	       "Must add the md5-key for key-id $VAR(@)"

*/
func quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5KeyId(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication md5 key-id WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: MD5 key
	syntax:expression: pattern $VAR(@) "^[^[:space:]]{1,16}$"; "MD5 key must be 16 characters or less"
	val_help: MD5 Key (16 characters or less)

	# If this node is created
	create:
		vtysh -c "configure terminal" -c "interface $VAR(../../../../../../@)" \
	             -c "ip ospf message-digest-key $VAR(../@) md5 $VAR(@)"

	# If the value of this node is changed
	update:
		vtysh -c "configure terminal" -c "interface $VAR(../../../../../../@)" \
	             -c "no ip ospf message-digest-key $VAR(../@)" \
	             -c "ip ospf message-digest-key $VAR(../@) md5 $VAR(@)"

	# If this node is deleted
	delete:
	        vtysh -c "configure terminal" -c "interface $VAR(../../../../../../@)" \
	             -c "no ip ospf message-digest-key $VAR(../@)"
*/
func quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5KeyIdMd5Key(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication md5 key-id WORD md5-key WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf message-digest-key ", Args[1], " md5 ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("no ip ospf message-digest-key ", Args[1]))
	}
	return cmd.Success
}

/*
	help: MD5 parameters

	create: vtysh -c "configure terminal" \
		-c "interface $VAR(../../../../@)" \
		-c "no ip ospf authentication" \
		-c "ip ospf authentication message-digest"

	delete: vtysh -c "configure terminal" \
		-c "interface $VAR(../../../../@)" \
	        -c "no ip ospf authentication"
*/
func quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication md5
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf authentication",
			"ip ospf authentication message-digest")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf authentication")
	}
	return cmd.Success
}

/*
	help: OSPF interface authentication

*/
func quaggaInterfacesInterfaceIpv4OspfAuthentication(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Plain text password
	syntax:expression: pattern $VAR(@) "^[^[:space:]]{1,8}$" ; "Password must be 8 characters or less"
	val_help: Plain text password (8 characters or less)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../../@)" \
		 -c "no ip ospf authentication " -c "ip ospf authentication " \
		 -c "ip ospf authentication-key $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../../@)" \
		 -c "no ip ospf authentication " -c "no ip ospf authentication-key"
*/
func quaggaInterfacesInterfaceIpv4OspfAuthenticationPlaintextPassword(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf authentication plaintext-password WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf authentication",
			"ip ospf authentication",
			fmt.Sprint("ip ospf authentication-key ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf authentication",
			"no ip ospf authentication-key")
	}
	return cmd.Success
}

/*
	type: u32
	help: Bandwidth of interface (kilobits/sec)
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 10000000; "Must be between 1-10000000"
	val_help: u32:1-10000000; Bandwidth in kilobits/sec (for calculating OSPF cost)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "bandwidth $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no bandwidth"
*/
func quaggaInterfacesInterfaceIpv4OspfBandwidth(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf bandwidth WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("bandwidth ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no bandwidth ")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interface cost
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; OSPF interface cost

	update:vtysh -c "configure terminal" \
		-c "interface $VAR(../../../@)" \
		-c "ip ospf cost $VAR(@)"
	delete:vtysh -c "configure terminal" \
		-c "interface $VAR(../../../@)" \
		-c "no ip ospf cost"
*/
func quaggaInterfacesInterfaceIpv4OspfCost(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf cost WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf cost ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf cost")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval after which neighbor is dead
	default: 40
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; OSPF dead interval in seconds (default 40)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf dead-interval $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf dead-interval"
*/
func quaggaInterfacesInterfaceIpv4OspfDeadInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf dead-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf dead-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf dead-interval")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between hello packets
	default: 10
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Interval between OSPF hello packets in seconds (default 10)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf hello-interval $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf hello-interval"
*/
func quaggaInterfacesInterfaceIpv4OspfHelloInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf hello-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf hello-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf hello-interval")
	}
	return cmd.Success
}

/*
	help: Disable Maximum Transmission Unit (MTU) mismatch detection
	create:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf mtu-ignore"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf mtu-ignore"
*/
func quaggaInterfacesInterfaceIpv4OspfMtuIgnore(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf mtu-ignore
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ip ospf mtu-ignore")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf mtu-ignore")
	}
	return cmd.Success
}

/*
	type: txt
	help: Network type
	syntax:expression: $VAR(@) in "broadcast", "non-broadcast", "point-to-multipoint", "point-to-point"; \
	       "Must be (broadcast|non-broadcast|point-to-multipoint|point-to-point)"
	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf network $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf network"

	val_help: broadcast; Broadcast network type
	val_help: non-broadcast; Non-broadcast network type
	val_help: point-to-multipoint; Point-to-multipoint network type
	val_help: point-to-point; Point-to-point network type
*/
func quaggaInterfacesInterfaceIpv4OspfNetwork(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf network WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf network ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf network")
	}
	return cmd.Success
}

/*
	help: Open Shortest Path First (OSPF) parameters
*/
func quaggaInterfacesInterfaceIpv4Ospf(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Router priority
	default: 1
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "Must be between 0-255"
	val_help: u32:0-255; Priority (default 1)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf priority $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf priority"
*/
func quaggaInterfacesInterfaceIpv4OspfPriority(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf priority WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf priority ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf priority")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between retransmitting lost link state advertisements
	default: 5
	syntax:expression: $VAR(@) >= 3 && $VAR(@) <= 65535; "Must be between 3-65535"
	val_help: u32: 3-65535; Retransmit interval in seconds (default 5)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
		-c "ip ospf retransmit-interval $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
		-c "no ip ospf retransmit-interval"
*/
func quaggaInterfacesInterfaceIpv4OspfRetransmitInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf retransmit-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf retransmit-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf retransmit-interval")
	}
	return cmd.Success
}

/*
	type: u32
	help: Link state transmit delay
	default: 1
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Transmit delay in seconds (default 1)

	update:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ip ospf transmit-delay $VAR(@)"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ip ospf transmit-delay"
*/
func quaggaInterfacesInterfaceIpv4OspfTransmitDelay(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv4 ospf transmit-delay WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ip ospf transmit-delay ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ip ospf transmit-delay")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interface cost
	default: 1
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; OSPFv3 cost

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 cost $VAR(@)"

*/
func quaggaInterfacesInterfaceIpv6Ospfv3Cost(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 cost WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 cost ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ipv6 ospf6 cost")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval after which neighbor is declared dead
	default: 40
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Neighbor dead interval in seconds (default 40)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 dead-interval $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 dead-interval 40"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3DeadInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 dead-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 dead-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 dead-interval 40")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between hello packets
	default: 10
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Interval between OSPFv3 hello packets in seconds (default 10)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 hello-interval $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 hello-interval 10"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3HelloInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 hello-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 hello-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 hello-interval 10")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interface MTU
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Interface MTU

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 ifmtu $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ipv6 ospf6 ifmtu"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3Ifmtu(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 ifmtu WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 ifmtu ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ipv6 ospf6 ifmtu")
	}
	return cmd.Success
}

/*
	type: u32
	help: Instance-id
	default: 0
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "Must be between 0-255"
	val_help: u32:0-255; Instance Id (default 0)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 instance-id $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 instance-id 0"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3InstanceId(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 instance-id WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 instance-id ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 instance-id 0")
	}
	return cmd.Success
}

/*
	help: Disable Maximum Transmission Unit mismatch detection
	create:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 mtu-ignore"
	delete:vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ipv6 ospf6 mtu-ignore"

*/
func quaggaInterfacesInterfaceIpv6Ospfv3MtuIgnore(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 mtu-ignore
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 mtu-ignore")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ipv6 ospf6 mtu-ignore")
	}
	return cmd.Success
}

/*
	help: IPv6 Open Shortest Path First (OSPFv3)
*/
func quaggaInterfacesInterfaceIpv6Ospfv3(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Disable forming of adjacency
	create: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 passive"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "no ipv6 ospf6 passive"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3Passive(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 passive
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 passive")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"no ipv6 ospf6 passive")
	}
	return cmd.Success
}

/*
	type: u32
	help: Router priority
	default: 1
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "Must be between 0-255"
	val_help: u32:0-255; Priority (default 1)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 priority $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" -c "ipv6 ospf6 priority 1"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3Priority(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 priority WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 priority ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 priority 1")
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between retransmitting lost link state advertisements
	default: 5
	syntax:expression: $VAR(@) >= 3 && $VAR(@) <= 65535; "Must be between 3-65535"
	val_help: u32:3-65535; Retransmit interval in seconds (default 5)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 retransmit-interval $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 retransmit-interval 5"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3RetransmitInterval(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 retransmit-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 retransmit-interval ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 retransmit-interval 5")
	}
	return cmd.Success
}

/*
	type: u32
	help: Link state transmit delay
	default: 1
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"
	val_help: u32:1-65535; Link state transmit delay (default 1)

	update: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 transmit-delay $VAR(@)"
	delete: vtysh -c "configure terminal" -c "interface $VAR(../../../@)" \
	          -c "ipv6 ospf6 transmit-delay 1"
*/
func quaggaInterfacesInterfaceIpv6Ospfv3TransmitDelay(Cmd int, Args cmd.Args) int {
	//interfaces interface WORD ipv6 ospfv3 transmit-delay WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			fmt.Sprint("ipv6 ospf6 transmit-delay ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("interface ", Args[0]),
			"ipv6 ospf6 transmit-delay 1")
	}
	return cmd.Success
}

/*
	tag:
	priority: 470
	type: u32
	help: IP access-list filter
	syntax:expression: ($VAR(@) >= 1 && $VAR(@) <= 199) || ($VAR(@) >= 1300 && $VAR(@) <= 2699); \
	"Access list number must be
	  <1-99>\tIP standard access list
	  <100-199>\tIP extended access list
	  <1300-1999>\tIP standard access list (expanded range)
	  <2000-2699>\tIP extended access list (expanded range)"

	val_help: u32:1-99; IP standard access list
	val_help: u32:100-199; IP extended access list
	val_help: u32:1300-1999; IP standard access list (expanded range)
	val_help: u32:2000-2699; IP extended access list (expanded range)

	end: /opt/vyatta/sbin/vyatta-policy.pl --update-access-list $VAR(@)
*/
func quaggaPolicyAccessList(Cmd int, Args cmd.Args) int {
	//policy access-list WORD
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this access-list
*/
func quaggaPolicyAccessListDescription(Cmd int, Args cmd.Args) int {
	//policy access-list WORD description WORD
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: Rule for this access-list
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "rule number must be between 1 and 65535"
	val_help: u32:1-65535; Access-list rule number
*/
func quaggaPolicyAccessListRule(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Action to take on networks matching this rule [REQUIRED]
	syntax:expression: $VAR(@) in "permit", "deny"; "action must be permit or deny"
	val_help: permit; Permit matching networks
	val_help: deny; Deny matching networks
*/
func quaggaPolicyAccessListRuleAction(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD action WORD
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this rule
*/
func quaggaPolicyAccessListRuleDescription(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD description WORD
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Any IP address to match
	commit:expression: ($VAR(../../../@) >= 100 && $VAR(../../../@) <= 199) || ($VAR(../../../@) >= 2000 && $VAR(../../../@) <= 2699); "access-list number must be <100-199> or <2000-2699> to set destination matches"
	commit:expression: ($VAR(../host/) == "") && ($VAR(../network/) == ""); "you may only define one filter type. (host|network|any)"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessListRuleDestinationAny(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD destination any
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Single host IP address to match
	val_help: Host address to match
	commit:expression: ($VAR(../../../@) >= 100 && $VAR(../../../@) <= 199) || ($VAR(../../../@) >= 2000 && $VAR(../../../@) <= 2699); "access-list number must be <100-199> or <2000-2699> to set destination matches"
	commit:expression: ($VAR(../any/) == "") && ($VAR(../network/) == ""); "you may only define one filter type. (host|network|any)"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessListRuleDestinationHost(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD destination host A.B.C.D
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Network/netmask to match (requires network be defined)
	val_help: Inverse-mask to match
	commit:expression: ($VAR(../../../@) >= 100 && $VAR(../../../@) <= 199) || ($VAR(../../../@) >= 2000 && $VAR(../../../@) <= 2699); "access-list number must be <100-199> or <2000-2699> to set destination matches"
	commit:expression: ($VAR(../any/) == "") && ($VAR(../host/) == ""); "you may only define one filter type. (host|network|any)"
	commit:expression: $VAR(../network/) != ""; "you must specify a network if you configure an inverse mask."
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessListRuleDestinationInverseMask(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD destination inverse-mask A.B.C.D
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Network/netmask to match (requires inverse-mask be defined)
	val_help: Network to match
	commit:expression: ($VAR(../../../@) >= 100 && $VAR(../../../@) <= 199) || ($VAR(../../../@) >= 2000 && $VAR(../../../@) <= 2699); "access-list number must be <100-199> or <2000-2699> to set destination matches"
	commit:expression: ($VAR(../host/) == "") && ($VAR(../any/) == ""); "you may only define one filter type. (host|network|any)"
	commit:expression: $VAR(../inverse-mask/) != ""; "you must specify an inverse-mask if you configure a network"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessListRuleDestinationNetwork(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD destination network A.B.C.D
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Destination network or address
*/
func quaggaPolicyAccessListRuleDestination(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD destination
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Any IP address to match
	commit:expression: ($VAR(../host/) == "") && ($VAR(../network/) == ""); "you may only define one filter type. (host|network|any)"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessListRuleSourceAny(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD source any
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Single host IP address to match
	val_help: Host address to match
	commit:expression: ($VAR(../any/) == "") && ($VAR(../network/) == ""); "you may only define one filter type. (host|network|any)"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessListRuleSourceHost(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD source host A.B.C.D
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Network/netmask to match (requires network be defined)
	val_help: Inverse-mast to match
	commit:expression: ($VAR(../any/) == "") && ($VAR(../host/) == ""); "you may only define one filter type. (host|network|any)"
	commit:expression: $VAR(../network/) != ""; "you must specify a network if you configure an inverse-mask"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessListRuleSourceInverseMask(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD source inverse-mask A.B.C.D
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Network/netmask to match (requires inverse-mask be defined)
	val_help: Network to match

	commit:expression: ($VAR(../host/) == "") && ($VAR(../any/) == ""); "you may only define one filter type.  (host|network|any)"
	commit:expression: $VAR(../inverse-mask/) != ""; "you must specify an inverse-mask if you configure a network"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessListRuleSourceNetwork(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD source network A.B.C.D
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Source network or address to match
*/
func quaggaPolicyAccessListRuleSource(Cmd int, Args cmd.Args) int {
	//policy access-list WORD rule WORD source
	quaggaUpdateCheckAccessList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	priority: 470
	type: txt
	help: IPv6 access-list filter
	val_help: Name of IPv6 access-list

	syntax:expression: pattern $VAR(@) "^[[:graph:]]{1,64}$" ; \
	                   "access-list name must be 64 characters or less"
	syntax:expression: pattern $VAR(@) "^[^-]" ; \
	                   "access-list name cannot start with \"-\""
	syntax:expression: pattern $VAR(@) "^[^|;&$<>]*$" ; \
	                   "access-list name cannot contain shell punctuation"

	end: /opt/vyatta/sbin/vyatta-policy.pl --update-access-list6 "$VAR(@)"

*/
func quaggaPolicyAccessList6(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this IPv6 access-list
*/
func quaggaPolicyAccessList6Description(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD description WORD
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: Rule for this access-list6
	val_help: u32:1-65535;  Access-list6 rule number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; \
	                   "rule number must be between 1 and 65535"
*/
func quaggaPolicyAccessList6Rule(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD rule WORD
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Action to take on networks matching this rule [REQUIRED]
	val_help: permit; Permit matching networks
	val_help: deny; Deny matching networks

	syntax:expression: $VAR(@) in "permit", "deny"; "action must be permit or deny"
*/
func quaggaPolicyAccessList6RuleAction(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD rule WORD action WORD
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this IPv6 access-list rule
*/
func quaggaPolicyAccessList6RuleDescription(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD rule WORD description WORD
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Any IPv6 address to match
	commit:expression: ($VAR(../network/) == ""); "you may only define one filter type. (network|any)"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessList6RuleSourceAny(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD rule WORD source any
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Exact match of the network prefixes
	commit:expression: ($VAR(../any/) == ""); "exact-match can only be used with a network filter "
*/
func quaggaPolicyAccessList6RuleSourceExactMatch(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD rule WORD source exact-match
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv6net
	help: Network/netmask to match (requires inverse-mask be defined)
	val_help: IPv6 address and prefix length
	commit:expression: ($VAR(../any/) == ""); "you may only define one filter type.  (network|any)"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
*/
func quaggaPolicyAccessList6RuleSourceNetwork(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD rule WORD source network X:X::X:X/M
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Source IPv6 network to match
*/
func quaggaPolicyAccessList6RuleSource(Cmd int, Args cmd.Args) int {
	//policy access-list6 WORD rule WORD source
	quaggaUpdateCheckAccessList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	priority: 470
	type: txt
	help: Border Gateway Protocol (BGP) autonomous system path filter
	val_help: AS path list name

	syntax:expression: pattern $VAR(@) "^[-a-zA-Z0-9.]+$" ; "as-path-list name must be alpha-numeric"

	end: /opt/vyatta/sbin/vyatta-policy.pl --update-aspath-list $VAR(@)

*/
func quaggaPolicyAsPathList(Cmd int, Args cmd.Args) int {
	//policy as-path-list WORD
	quaggaUpdateCheckAsPathList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this as-path-list
*/
func quaggaPolicyAsPathListDescription(Cmd int, Args cmd.Args) int {
	//policy as-path-list WORD description WORD
	quaggaUpdateCheckAsPathList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: Rule for this as-path-list
	val_help: u32:1-65535; AS path list rule number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "rule number must be between 1 and 65535"

*/
func quaggaPolicyAsPathListRule(Cmd int, Args cmd.Args) int {
	//policy as-path-list WORD rule WORD
	quaggaUpdateCheckAsPathList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Action to take on AS paths matching this rule [REQUIRED]
	val_help: permit; Permit matching as-paths
	val_help: deny; Deny matching as-paths

	syntax:expression: $VAR(@) in "permit", "deny"; "action must be permit or deny"
*/
func quaggaPolicyAsPathListRuleAction(Cmd int, Args cmd.Args) int {
	//policy as-path-list WORD rule WORD action WORD
	quaggaUpdateCheckAsPathList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this rule
*/
func quaggaPolicyAsPathListRuleDescription(Cmd int, Args cmd.Args) int {
	//policy as-path-list WORD rule WORD description WORD
	quaggaUpdateCheckAsPathList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Regular expression to match against an AS path
	val_help: <aa:nn>; AS path regular expression (ex: "50:1 6553:1201")
	# TODO: check regex syntax; \
	#       "invalid chars in regex syntax"

	commit:expression: $VAR(../action/@) != ""; "You must specify an action"
*/
func quaggaPolicyAsPathListRuleRegex(Cmd int, Args cmd.Args) int {
	//policy as-path-list WORD rule WORD regex WORD
	quaggaUpdateCheckAsPathList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	priority: 470
	type: u32
	help: Border Gateway Protocol (BGP) community-list filter

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 500; "
	community-list must be
	  <1-99>\tBGP community list (standard)
	  <100-500>\tBGP community list (expanded) "

	val_help: u32:1-99; BGP community list (standard)
	val_help: u32:100-500; BGP community list (expanded)

	end: /opt/vyatta/sbin/vyatta-policy.pl --update-community-list $VAR(@)
*/
func quaggaPolicyCommunityList(Cmd int, Args cmd.Args) int {
	//policy community-list WORD
	quaggaUpdateCheckCommunityList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this community list
*/
func quaggaPolicyCommunityListDescription(Cmd int, Args cmd.Args) int {
	//policy community-list WORD description WORD
	quaggaUpdateCheckCommunityList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: create a rule for this BGP community list
	val_help: u32:1-65535; Community-list rule number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "rule number must be between 1 and 65535"
*/
func quaggaPolicyCommunityListRule(Cmd int, Args cmd.Args) int {
	//policy community-list WORD rule WORD
	quaggaUpdateCheckCommunityList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Action to take on communities matching this rule [REQUIRED]
	val_help: permit; Permit matching communities
	val_help: deny; Deny matching communities

	syntax:expression: $VAR(@) in "permit", "deny"; "action must be permit or deny"
*/
func quaggaPolicyCommunityListRuleAction(Cmd int, Args cmd.Args) int {
	//policy community-list WORD rule WORD action WORD
	quaggaUpdateCheckCommunityList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this rule
*/
func quaggaPolicyCommunityListRuleDescription(Cmd int, Args cmd.Args) int {
	//policy community-list WORD rule WORD description WORD
	quaggaUpdateCheckCommunityList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Regular expression to match against a community list
	val_help: Community list regular expression

	syntax:expression: exec " \
	if [ $VAR(../../@) -ge 1 ] && [ $VAR(../../@) -le 99 ]; then \
	  if [ -n \"`echo $VAR(@) | sed 's/[0-9]*:[0-9]* //g' | sed -e 's/internet//g' -e 's/local-AS//g' -e 's/no-advertise//g' -e 's/no-export//g'`\" ]; then \
	    echo regex $VAR(@) is invalid for a standard community list; \
	    exit 1 ; \
	  fi ; \
	fi ; "

	commit:expression: $VAR(../action/@) != ""; "You must specify an action"
*/
func quaggaPolicyCommunityListRuleRegex(Cmd int, Args cmd.Args) int {
	//policy community-list WORD rule WORD regex WORD
	quaggaUpdateCheckCommunityList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	priority: 470
	type: txt
	help: IP prefix-list filter
	val_help: Prefix list name

	syntax:expression: pattern $VAR(@) "^[-a-zA-Z0-9.]+$" ; "prefix-list name must be alpha-numeric"
*/
func quaggaPolicyPrefixList(Cmd int, Args cmd.Args) int {
	//policy prefix-list WORD
	quaggaUpdateCheckPrefixList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this prefix-list
*/
func quaggaPolicyPrefixListDescription(Cmd int, Args cmd.Args) int {
	//policy prefix-list WORD description WORD
	quaggaUpdateCheckPrefixList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: Rule for this prefix-list
	val_help: u32:1-65535; Prefix-list rule number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "rule number must be between 1 and 65535"

	commit:expression: $VAR(./prefix/) != ""; "You must specify a prefix"
	commit:expression: $VAR(./action/) != ""; "You must specify an action"

	delete:  touch /tmp/protocols-$VAR(../@)-$VAR(@).$PPID ;
	         len=`echo $VAR(@) | awk -F/ '{ print $2 }'` ;
	         if [ -n "$VAR(./ge/@)" ]; then
	           cond="ge $VAR(./ge/@) ";
	         fi;
	         if [ -n "$VAR(./le/@)" ]; then
	           cond="$cond le $VAR(./le/@) ";
	         fi;
	         vtysh -c "configure terminal"  \
	           -c "no ip prefix-list $VAR(../@) seq $VAR(@) $VAR(./action/@) $VAR(./prefix/@) $cond "

	end:  len=`echo $VAR(./prefix/@) | awk -F/ '{ print $2 }'` ;
	      if [ -n "$VAR(./ge/@)" ]; then
	        if [ $len -ge $VAR(./ge/@) ]; then
	          echo "ge must be greater than prefix length";
	          exit 1 ;
	        fi ;
	        cond="ge $VAR(./ge/@) ";
	      fi;
	      if [ -n "$VAR(./le/@)" ]; then
	        if [ $VAR(./le/@) -ne 32 ] && [ -n "$VAR(./ge/@)" ] && [ $VAR(./le/@) -le $VAR(./ge/@) ]; then
	          echo "le must be greater than or equal to ge";
	          exit 1 ;
	        fi ;
	        cond="$cond le $VAR(./le/@) ";
	      fi;
	      if [ -f "/tmp/protocols-$VAR(../@)-$VAR(@).$PPID" ]; then
	        rm -f "protocols-$VAR(../@)-$VAR(@).$PPID" ;
	      else
	        vtysh -c "configure terminal" \
	          -c "ip prefix-list $VAR(../@) seq $VAR(@) $VAR(./action/@) $VAR(./prefix/@) $cond " ;
	      fi ;
	      exit 0 ;
*/
func quaggaPolicyPrefixListRule(Cmd int, Args cmd.Args) int {
	//policy prefix-list WORD rule WORD
	quaggaUpdateCheckPrefixList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Action to take on prefixes matching this rule [REQUIRED]
	val_help: permit; Permit matching prefixes
	val_help: deny; Deny matching prefixes

	syntax:expression: $VAR(@) in "permit", "deny"; "action must be permit or deny"
*/
func quaggaPolicyPrefixListRuleAction(Cmd int, Args cmd.Args) int {
	//policy prefix-list WORD rule WORD action WORD
	quaggaUpdateCheckPrefixList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this rule
*/
func quaggaPolicyPrefixListRuleDescription(Cmd int, Args cmd.Args) int {
	//policy prefix-list WORD rule WORD description WORD
	quaggaUpdateCheckPrefixList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Prefix length to match a netmask greater than or equal to it
	val_help: u32:0-32; Netmask greater than length

	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 32; "ge must be between 0 and 32"
*/
func quaggaPolicyPrefixListRuleGe(Cmd int, Args cmd.Args) int {
	//policy prefix-list WORD rule WORD ge WORD
	quaggaUpdateCheckPrefixList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Prefix length to match a netmask less than or equal to it
	val_help: u32:0-32; Netmask less than length

	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 32; "le must be between 0 and 32"
*/
func quaggaPolicyPrefixListRuleLe(Cmd int, Args cmd.Args) int {
	//policy prefix-list WORD rule WORD le WORD
	quaggaUpdateCheckPrefixList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4net
	help: Prefix to match
	val_help: Prefix to match against
*/
func quaggaPolicyPrefixListRulePrefix(Cmd int, Args cmd.Args) int {
	//policy prefix-list WORD rule WORD prefix A.B.C.D/M
	quaggaUpdateCheckPrefixList()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	priority: 470
	type: txt
	help: IPv6 prefix-list filter
	val_help: Prefix list name

	syntax:expression: pattern $VAR(@) "^[-a-zA-Z0-9.]+$" ; "prefix-list6 name must be alpha-numeric"
*/
func quaggaPolicyPrefixList6(Cmd int, Args cmd.Args) int {
	//policy prefix-list6 WORD
	quaggaUpdateCheckPrefixList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this prefix-list6

*/
func quaggaPolicyPrefixList6Description(Cmd int, Args cmd.Args) int {
	//policy prefix-list6 WORD description WORD
	quaggaUpdateCheckPrefixList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: Rule for this prefix-list6
	val_help: u32:1-65535; Prefix list rule number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "rule number must be between 1 and 65535"

	commit:expression: $VAR(./prefix/) != ""; "You must specify a prefix"

	commit:expression: $VAR(./action/) != ""; "You must specify an action"

	delete:  len=`echo $VAR(@) | awk -F/ '{ print $2 }'` ;
	         if [ -n "$VAR(./ge/@)" ]; then
	           cond="ge $VAR(./ge/@) ";
	         fi;
	         if [ -n "$VAR(./le/@)" ]; then
	           cond="$cond le $VAR(./le/@) ";
	         fi;
	         vtysh -c "configure terminal"  \
	           -c "no ipv6 prefix-list $VAR(../@) seq $VAR(@) $VAR(./action/@) $VAR(./prefix/@) $cond "

	end:  len=`echo $VAR(./prefix/@) | awk -F/ '{ print $2 }'` ;
	      if [ -n "$VAR(./ge/@)" ]; then
	        if [ $len -ge $VAR(./ge/@) ]; then
	          echo "ge must be greater than prefix length";
	          exit 1 ;
	        fi ;
	        cond="ge $VAR(./ge/@) ";
	      fi;
	      if [ -n "$VAR(./le/@)" ]; then
	        if [ $VAR(./le/@) -ne 128 ] && [ -n "$VAR(./ge/@)" ] && [ $VAR(./le/@) -le $VAR(./ge/@) ]; then
	          echo "le must be greater than or equal to ge";
	          exit 1 ;
	        fi ;
	        cond="$cond le $VAR(./le/@) ";
	      fi;

	      if [ ${COMMIT_ACTION} = 'SET' -o ${COMMIT_ACTION} = 'ACTIVE' ]; then
	        vtysh -c "configure terminal" \
	          -c "ipv6 prefix-list $VAR(../@) seq $VAR(@) $VAR(./action/@) $VAR(./prefix/@) $cond " ;
	      fi;
	      exit 0 ;
*/
func quaggaPolicyPrefixList6Rule(Cmd int, Args cmd.Args) int {
	//policy prefix-list6 WORD rule WORD
	quaggaUpdateCheckPrefixList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Action to take on prefixes matching this rule
	val_help: permit; Permit matching prefixes
	val_help: deny; Deny matching prefixes

	syntax:expression: $VAR(@) in "permit", "deny"; "action must be permit or deny"
*/
func quaggaPolicyPrefixList6RuleAction(Cmd int, Args cmd.Args) int {
	//policy prefix-list6 WORD rule WORD action WORD
	quaggaUpdateCheckPrefixList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this rule
*/
func quaggaPolicyPrefixList6RuleDescription(Cmd int, Args cmd.Args) int {
	//policy prefix-list6 WORD rule WORD description WORD
	quaggaUpdateCheckPrefixList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Prefix length to match a netmask greater than or equal to it
	val_help: u32:0-128; Netmask greater than length

	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 128; "ge must be between 0 and 128"
*/
func quaggaPolicyPrefixList6RuleGe(Cmd int, Args cmd.Args) int {
	//policy prefix-list6 WORD rule WORD ge WORD
	quaggaUpdateCheckPrefixList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Prefix length to match a netmask less than or equal to it
	val_help: u32:0-128; Netmask less than length

	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 128; "le must be between 0 and 128"
*/
func quaggaPolicyPrefixList6RuleLe(Cmd int, Args cmd.Args) int {
	//policy prefix-list6 WORD rule WORD le WORD
	quaggaUpdateCheckPrefixList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv6net
	help: Prefix to match
	val_help: IPv6 prefix
*/
func quaggaPolicyPrefixList6RulePrefix(Cmd int, Args cmd.Args) int {
	//policy prefix-list6 WORD rule WORD prefix X:X::X:X/M
	quaggaUpdateCheckPrefixList6()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	priority: 470
	type: txt
	help: IP route-map
	val_help: Route map name

	syntax:expression: pattern $VAR(@) "^[-a-zA-Z0-9.]+$" ; "route-map $VAR(@): name must be alpha-numeric"
*/
func quaggaPolicyRouteMap(Cmd int, Args cmd.Args) int {
	//policy route-map WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this route-map
*/
func quaggaPolicyRouteMapDescription(Cmd int, Args cmd.Args) int {
	//policy route-map WORD description WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: Rule for this route-map
	val_help: u32:1-65535; Route-map rule number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "rule number must be between 1 and 65535"

	delete: if [ -f /tmp/route-map-$VAR(../@)-rule-$VAR(@)-action.$PPID ]; then
	             vtysh -c "configure terminal" -c "no route-map $VAR(../@) $VAR(./@/action/@) $VAR(@)";
	             rm -f /tmp/route-map-$VAR(../@)-rule-$VAR(@)-action.$PPID;
	        fi;
*/
func quaggaPolicyRouteMapRule(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Action to take on prefixes matching this rule [REQUIRED]
	val_help: permit; Permit matching prefixes
	val_help: deny; Deny matching prefixes

	syntax:expression: $VAR(@) in "permit", "deny"; "action must be permit or deny"

	update: /opt/vyatta/sbin/vyatta-policy.pl --check-routemap-action "policy route-map $VAR(../../@) rule $VAR(../@) action";
	        if [ $? -eq 0 ]; then
	          vtysh -c "configure terminal" -c "route-map $VAR(../../@) $VAR(@) $VAR(../@)";
	        else
	          echo    "You can not change the action.";
	          echo    "  To change the action you must first delete the rule ";
	          echo -e "  \"delete route-map $VAR(../../@) rule $VAR(../@)\" and commit it. \\n";
	          exit 1;
	        fi;

	delete: /opt/vyatta/sbin/vyatta-policy.pl --check-delete-routemap-action "policy route-map $VAR(../../@) rule $VAR(../@)";
	        if [ $? -eq 0 ]; then
	          touch /tmp/route-map-$VAR(../../@)-rule-$VAR(../@)-action.$PPID ;
	        else
	          echo    "Action is a required parameter. ";
	          echo -e "  To delete that node you must delete \"route-map $VAR(../../@) rule $VAR(../@)\". \\n";
	          exit 1;
	        fi;
*/
func quaggaPolicyRouteMapRuleAction(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD action WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Call another route-map on match
	val_help: Route map name

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" "; "called route-map $VAR(@) doesn't exist"
	commit:expression: $VAR(../action/) != ""; "you must define an action"

	update: vtysh -c "configure terminal" -c "route-map $VAR(../../@) $VAR(../action/@) $VAR(../@)" \
	           -c "call $VAR(@)"

	delete: vtysh -c "configure terminal" -c "route-map $VAR(../../@) $VAR(../action/@) $VAR(../@)" \
	           -c "no call "
*/
func quaggaPolicyRouteMapRuleCall(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD call WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Jump to a different rule in this route-map on a match
	val_help: u32:1-65535; Rule number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "continue must be between 1 and 65535"

	commit:expression: $VAR(@) > $VAR(../@); "you may only continue forward in the route-map"
	commit:expression: $VAR(../action/) != ""; "you must define an action"

	update: vtysh -c "configure terminal" \
	   	   -c "route-map $VAR(../../@) $VAR(../action/@) $VAR(../@)" \
	           -c "continue $VAR(@)"

	delete: vtysh -c "configure terminal" \
		   -c "route-map $VAR(../../@) $VAR(../action/@) $VAR(../@)" \
	           -c "no continue "
*/
func quaggaPolicyRouteMapRuleContinue(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD continue WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this rule
*/
func quaggaPolicyRouteMapRuleDescription(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD description WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: BGP as-path-list to match
	val_help: AS path list name
	allowed: cli-shell-api listActiveNodes policy as-path-list

	commit:expression: $VAR(../../action/) != ""; "You must specify an action"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" "; "match as-path: AS path list $VAR(@) doesn't exist"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "match as-path $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no match as-path $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchAsPath(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match as-path WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: BGP community-list to match
	val_help: u32:1-99; BGP community list (standard)
	val_help: u32:100-500; BGP community list (expanded)

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy community-list $VAR(@)\" "; "community-list $VAR(@) doesn't exist"
*/
func quaggaPolicyRouteMapRuleMatchCommunityCommunityList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match community community-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Community-list to exactly match
*/
func quaggaPolicyRouteMapRuleMatchCommunityExactMatch(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match community exact-match
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: BGP community-list to match
	delete: echo route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@) >>  /tmp/delete-policy-route-map-$VAR(../../../@)-$VAR(../../@)-match-community.$PPID
	end:  if [ -f /tmp/delete-policy-route-map-$VAR(../../../@)-$VAR(../../@)-match-community.$PPID ]; then
	        routemap=`cat /tmp/delete-policy-route-map-$VAR(../../../@)-$VAR(../../@)-match-community.$PPID`
	        rm -f /tmp/delete-policy-route-map-$VAR(../../../@)-$VAR(../../@)-match-community.$PPID;
	        vtysh --noerror -c "configure terminal" -c "$routemap " -c "no match community " ;
	        exit 0;
	      else
	        if [ -z "$VAR(./community-list/@)" ]; then
	          echo route-map $VAR(../../../@) rule $VAR(../../@) match community: you must configure a community-list ;
	          exit 1 ;
	        fi ;
	        if [ -z "$VAR(../../action/@)" ]; then
	          echo route-map $VAR(../../../@) rule $VAR(../../@): you must configure an action ;
	          exit 1 ;
	        fi ;
	        routemap='route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)';

		# uncomment and replace the call to vyatta-check-typeless-node.pl pending bug 2525
	        #if [ -n "$VAR(./exact-match/)" ]; then
	        #  cond="exact-match ";
	        #fi ;
	        ${vyatta_sbindir}/vyatta-check-typeless-node.pl "policy route-map $VAR(../../../@) rule $VAR(../../@) match community exact-match";
	        if [ $? -eq 0 ]; then
	          cond="exact-match ";
	        fi ;

	        vtysh --noerror -c "configure terminal" -c "$routemap " -c "no match community " ;
	        vtysh -c "configure terminal" -c "$routemap " -c "match community $VAR(./community-list/@) $cond" ;
	      fi
*/
func quaggaPolicyRouteMapRuleMatchCommunity(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match community
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: First hop interface of a route to match
	val_help: Interface name

	commit:expression: $VAR(../../action/) != ""; "You must specify an action"
	commit:expression: exec " \
	        if [ -z \"`ip addr | grep $VAR(@) `\" ]; then \
	          echo interface $VAR(@) doesn\\'t exist on this system ; \
	          exit 1 ; \
	        fi ; "

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "match interface $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no match interface $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchInterface(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match interface WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: IP access-list to match
	val_help: u32:1-99; IP standard access list number
	val_help: u32:100-199; IP extended access list number
	val_help: u32:1300-1999; IP standard access list number (expanded range)
	val_help: u32:2000-2699; IP extended access list number (expanded range)

	allowed: cli-shell-api listActiveNodes policy access-list

	commit:expression: $VAR(../prefix-list/) == ""; "you may only specify a prefix-list or access-list"

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list $VAR(@)\" "; "access-list $VAR(@) does not exist"

	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ip address $VAR(@) "

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ip address $VAR(@) "
*/
func quaggaPolicyRouteMapRuleMatchIpAddressAccessList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip address access-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: IP address of route to match
*/
func quaggaPolicyRouteMapRuleMatchIpAddress(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip address
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: IP prefix-list to match
	val_help: Prefix list name

	allowed: cli-shell-api listActiveNodes policy prefix-list

	commit:expression: $VAR(../access-list/) == ""; "you may only specify a prefix-list or access-list"

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list $VAR(@)\" "; "prefix-list $VAR(@) does not exist"

	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ip address prefix-list $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ip address prefix-list $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchIpAddressPrefixList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip address prefix-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: IP access-list to match
	val_help: u32:1-99; IP standard access list number
	val_help: u32:100-199; IP extended access list number
	val_help: u32:1300-1999; IP standard access list number (expanded range)
	val_help: u32:2000-2699; IP extended access list number (expanded range)

	commit:expression: $VAR(../prefix-list/) == ""; "you may only specify a prefix-list or access-list"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list $VAR(@)\" "; "access-list $VAR(@) does not exist"
	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ip next-hop $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ip next-hop $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchIpNexthopAccessList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip nexthop access-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: IP next-hop of route to match
*/
func quaggaPolicyRouteMapRuleMatchIpNexthop(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip nexthop
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: IP prefix-list to match
	val_help: Prefix list name

	commit:expression: $VAR(../access-list/) == ""; "you can only specify a prefix-list or access-list"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list $VAR(@)\" "; "prefix-list $VAR(@) does not exist"
	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ip next-hop prefix-list $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ip next-hop prefix-list $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchIpNexthopPrefixList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip nexthop prefix-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: IP prefix parameters to match
*/
func quaggaPolicyRouteMapRuleMatchIp(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: IP access-list to match
	val_help: u32:1-99; IP standard access list number
	val_help: u32:100-199; IP extended access list number
	val_help: u32:1300-1999; IP standard access list number (expanded range)
	val_help: u32:2000-2699; IP extended access list number (expanded range)

	commit:expression: $VAR(../prefix-list/) == ""; "you may only specify a prefix-list or access-list"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list $VAR(@)\" "; "access-list $VAR(@) does not exist"
	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal"  \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ip route-source $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ip route-source $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchIpRouteSourceAccessList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip route-source access-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: IP route-source to match
*/
func quaggaPolicyRouteMapRuleMatchIpRouteSource(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip route-source
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: IP prefix-list to match
	val_help: Prefix list name

	commit:expression: $VAR(../access-list/) == ""; "you can only specify a prefix-list or access-list"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list $VAR(@)\" "; "prefix-list $VAR(@) does not exist"
	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	                     -c "route-map $VAR(../../../../../@) \
	                         $VAR(../../../../action/@) $VAR(../../../../@)" \
	                     -c "match ip route-source prefix-list $VAR(@)"

	delete: vtysh -c "configure terminal" \
	                     -c "route-map $VAR(../../../../../@) \
	                         $VAR(../../../../action/@) $VAR(../../../../@)" \
	                     -c "no match ip route-source prefix-list $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchIpRouteSourcePrefixList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ip route-source prefix-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: IPv6 access-list6 to match
	val_help: IPV6 access list name

	allowed: cli-shell-api listActiveNodes policy access-list6

	commit:expression: $VAR(../prefix-list/) == ""; "you may only specify a prefix-list or access-list"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list6 $VAR(@)\" "; "access-list6 $VAR(@) does not exist"
	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ipv6 address $VAR(@) "

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ipv6 address $VAR(@) "
*/
func quaggaPolicyRouteMapRuleMatchIpv6AddressAccessList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ipv6 address access-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: IPv6 address of route to match
*/
func quaggaPolicyRouteMapRuleMatchIpv6Address(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ipv6 address
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: IPv6 prefix-list to match
	val_help: IPv6 prefix list name

	allowed: cli-shell-api listActiveNodes policy prefix-list6

	commit:expression: $VAR(../access-list/) == ""; "you may only specify a prefix-list or access-list"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list6 $VAR(@)\" "; "prefix-list6 $VAR(@) does not exist"
	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ipv6 address prefix-list $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ipv6 address prefix-list $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchIpv6AddressPrefixList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ipv6 address prefix-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: IPv6 access-list6 to match
	val_help: IPv6 access list

	commit:expression: $VAR(../prefix-list/) == ""; "you may only specify a prefix-list or access-list"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list6 $VAR(@)\" "; "access-list6 $VAR(@) does not exist"
	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ipv6 next-hop $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ipv6 next-hop $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchIpv6NexthopAccessList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ipv6 nexthop access-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: IP next-hop of route to match
*/
func quaggaPolicyRouteMapRuleMatchIpv6Nexthop(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ipv6 nexthop
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: IPv6 prefix-list to match
	val_help: IPv6 prefix list name

	commit:expression: $VAR(../access-list/) == ""; "you can only specify a prefix-list or access-list"
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list $VAR(@)\" "; "prefix-list $VAR(@) does not exist"
	commit:expression: $VAR(../../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "match ipv6 next-hop prefix-list $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../../@) $VAR(../../../../action/@) $VAR(../../../../@)" \
	         -c "no match ipv6 next-hop prefix-list $VAR(@)"
*/
func quaggaPolicyRouteMapRuleMatchIpv6NexthopPrefixList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ipv6 nexthop prefix-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: IPv6 prefix parameters to match
*/
func quaggaPolicyRouteMapRuleMatchIpv6(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match ipv6
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric of route to match
	val_help: u32:1-65535; Rrute metric

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "metric must be between 1 and 65535"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "match metric $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no match metric $VAR(@)"

*/
func quaggaPolicyRouteMapRuleMatchMetric(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match metric WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Route parameters to match
*/
func quaggaPolicyRouteMapRuleMatch(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Border Gateway Protocol (BGP) origin code to match
	val_help: egp; Exterior gateway protocol origin
	val_help: igp; Interior gateway protocol origin
	val_help: incomplete; Incomplete origin

	syntax:expression: $VAR(@) in "egp", "igp", "incomplete"; "origin must be egp, igp, or incomplete"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "match origin $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no match origin $VAR(@)"

*/
func quaggaPolicyRouteMapRuleMatchOrigin(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match origin WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Peer address to match
	val_help: ipv4; Peer IP address
	val_help: local: Static or redistributed routes

	syntax:expression: exec "/opt/vyatta/sbin/vyatta-policy.pl --check-peer-syntax $VAR(@)"; "peer must be either an IP or local"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "match peer $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no match peer "
*/
func quaggaPolicyRouteMapRuleMatchPeer(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match peer WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Route tag to match
	val_help: u32:1-65535; Route tag

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "tag must be between 1 and 65535"
	commit:expression: $VAR(../../action) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "match tag $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no match tag $VAR(@)"

*/
func quaggaPolicyRouteMapRuleMatchTag(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD match tag WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Rule number to goto on match
	val_help: u32:1-65535; Rule number

	syntax:expression: $VAR(../next/) == ""; "you may set only goto or next"
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "goto must be a rule number between 1 and 65535"
	commit:expression: $VAR(@) > $VAR(../../@); "you may only go forward in the route-map"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "on-match goto $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no on-match goto "
*/
func quaggaPolicyRouteMapRuleOnMatchGoto(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD on-match goto WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Next sequence number to goto on match
	syntax:expression: $VAR(../goto/) == ""; "you may set only goto or next"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "on-match next "
	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no on-match next "

*/
func quaggaPolicyRouteMapRuleOnMatchNext(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD on-match next
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Exit policy on matches
*/
func quaggaPolicyRouteMapRuleOnMatch(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD on-match
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: AS number of an aggregation
	val_help: u32:1-65535; BGP AS number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967294; "BGP AS number must be between 1 and 4294967294"
*/
func quaggaPolicyRouteMapRuleSetAggregatorAs(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set aggregator as WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: IP address of an aggregation
	val_help: IP address
*/
func quaggaPolicyRouteMapRuleSetAggregatorIp(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set aggregator ip A.B.C.D
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Border Gateway Protocol (BGP) aggregator attribute
	commit:expression: $VAR(./as/) != "" && $VAR(./ip/) != ""; "you must configure both as and ip"
	commit:expression: $VAR(../../action/) != ""; "you must configure an action"
	delete: echo $VAR(./as/@) $VAR(./ip/@) > /tmp/policy-route-map-$VAR(../../../@)-$VAR(../../action/@)-$VAR(../../@)-set-aggregator.$PPID
	end: if [ -f "/tmp/policy-route-map-$VAR(../../../@)-$VAR(../../action/@)-$VAR(../../@)-set-aggregator.$PPID" ]; then
	        as=$(cat /tmp/policy-route-map-$VAR(../../../@)-$VAR(../../action/@)-$VAR(../../@)-set-aggregator.$PPID);
	        vtysh -c "configure terminal" \
	          -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	          -c "no set aggregator as $as" ;
	        rm -rf /tmp/policy-route-map-$VAR(../../../@)-$VAR(../../action/@)-$VAR(../../@)-set-aggregator.$PPID;
	      else
		as="$VAR(./as/@) $VAR(./ip/@)";
	        vtysh -c "configure terminal" \
	          -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	          -c "set aggregator as $as" ;
	      fi ;
*/
func quaggaPolicyRouteMapRuleSetAggregator(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set aggregator
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Prepend string for a Border Gateway Protocol (BGP) AS-path attribute
	val_help: BGP AS path prepend string (ex: "456 64500 45001")

	syntax:expression: exec "/opt/vyatta/sbin/vyatta-check-as-prepend.pl \"$VAR(@)\" "; "invalid AS path string"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	           -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	           -c "set as-path prepend $VAR(@)"

	delete: vtysh -c "configure terminal" \
	           -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	           -c "no set as-path prepend "
*/
func quaggaPolicyRouteMapRuleSetAsPathPrepend(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set as-path-prepend WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Border Gateway Protocol (BGP) atomic aggregate attribute
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"
	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set atomic-aggregate"
	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set atomic-aggregate"
*/
func quaggaPolicyRouteMapRuleSetAtomicAggregate(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set atomic-aggregate
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: BGP communities with a community-list
	val_help: u32:1-99; BGP community list (standard)
	val_help: u32:100-500; BGP community list (expanded)

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy community-list $VAR(@)\""; "community list $VAR(@) does not exist"
*/
func quaggaPolicyRouteMapRuleSetCommListCommList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set comm-list comm-list WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Delete BGP communities matching the community-list
*/
func quaggaPolicyRouteMapRuleSetCommListDelete(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set comm-list delete
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Border Gateway Protocol (BGP) communities matching a community-list
	delete: touch /tmp/policy-route-map-$VAR(../../../@)-$VAR(../../action/@)-$VAR(../../@)-set-comm-list.$PPID
	end: if [ -z "$VAR(./comm-list/)" ]; then
	        echo You must configure a comm-list ;
	        exit 1 ;
	      fi ;
	      vtysh --noerror -c "configure terminal" \
	        -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	        -c "no set comm-list " ;
	      if [ -f "/tmp/policy-route-map-$VAR(../../../@)-$VAR(../../action/@)-$VAR(../../@)-set-comm-list.$PPID" ]; then
	        rm -rf /tmp/policy-route-map-$VAR(../../../@)-$VAR(../../action/@)-$VAR(../../@)-set-comm-list.$PPID;
	      else
	        # uncomment this when 2525 is fixed and comment out the subsequent call
	        #if [ -n "$VAR(./delete/)" ]; then
	        #  cond="delete" ;
	        #fi ; \
		${vyatta_sbindir}/vyatta-check-typeless-node.pl "policy route-map $VAR(../../../@) rule $VAR(../../@) set comm-list delete";
	        if [ $? -eq 0 ]; then
	          cond="delete ";
	        fi ;
	        vtysh -c "configure terminal" \
	          -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	          -c "set comm-list $VAR(./comm-list/@) $cond" ;
	      fi;
*/
func quaggaPolicyRouteMapRuleSetCommList(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set comm-list
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Border Gateway Protocl (BGP) community attribute
	val_help: <AA:NN>; Community in 4 octet AS:value format
	val_help: local-AS; Advertise communities in local AS only (NO_EXPORT_SUBCONFED)
	val_help: no-advertise; Don't advertise this route to any peer (NO_ADVERTISE)
	val_help: no-export; Don't advertise outside of this AS of confederation boundry (NO_EXPORT)
	val_help: internet; Symbolic Internet community 0

	allowed:echo "none local-AS no-advertise no-export internet"

	syntax:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --check-community $VAR(@)"
	commit:expression: $VAR(../../action/) != "" ; "You must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set community $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set community "
*/
func quaggaPolicyRouteMapRuleSetCommunity(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set community WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Nexthop IP address
	val_help: IP address

	# TODO: can also set to peer for BGP
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set ip next-hop $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set ip next-hop "
*/
func quaggaPolicyRouteMapRuleSetIpNextHop(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set ip-next-hop A.B.C.D
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv6
	help: Nexthop IPv6 global address
	val_help: IPv6 address

	# TODO: can also set to peer for BGP
	commit:expression: $VAR(../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../@) $VAR(../../../action/@) $VAR(../../../@)" \
	         -c "set ipv6 next-hop global $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../@) $VAR(../../../action/@) $VAR(../../../@)" \
	         -c "no set ipv6 next-hop global"
*/
func quaggaPolicyRouteMapRuleSetIpv6NextHopGlobal(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set ipv6-next-hop global X:X::X:X
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv6
	help: Nexthop IPv6 local address
	val_help: IPv6 address

	# TODO: can also set to peer for BGP
	commit:expression: $VAR(../../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../@) $VAR(../../../action/@) $VAR(../../../@)" \
	         -c "set ipv6 next-hop local $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../../@) $VAR(../../../action/@) $VAR(../../../@)" \
	         -c "no set ipv6 next-hop local"
*/
func quaggaPolicyRouteMapRuleSetIpv6NextHopLocal(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set ipv6-next-hop local X:X::X:X
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Nexthop IPv6 address

*/
func quaggaPolicyRouteMapRuleSetIpv6NextHop(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set ipv6-next-hop
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Border Gateway Protocol (BGP) local preference attribute
	val_help: Local preference value

	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set local-preference $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set local-preference "
*/
func quaggaPolicyRouteMapRuleSetLocalPreference(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set local-preference WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Open Shortest Path First (OSPF) external metric-type
	val_help: type-1; OSPF external type 1 metric
	val_help: type-2; OSPF external type 2 metric

	syntax:expression: $VAR(@) in "type-1", "type-2"; "Must be (type-1, type-2)"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set metric-type $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set metric-type "
*/
func quaggaPolicyRouteMapRuleSetMetricType(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set metric-type WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Destination routing protocol metric
	val_help: <+/-metric>; Add or subtract metric
	val_help: u32:0-4294967295; Metric value

	syntax:expression: exec "if [ -n \"$(echo $VAR(@) | sed 's/^[+-]*[0123456789]* //')\" ]; then exit 1; fi; "; "metric must be an integer with an optional +/- prepend"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set metric $VAR(@)"

	delete: vtysh --noerror -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set metric "
*/
func quaggaPolicyRouteMapRuleSetMetric(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set metric WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Route parameters
*/
func quaggaPolicyRouteMapRuleSet(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Border Gateway Protocl (BGP) origin code
	val_help: igp; Interior gateway protocol origin
	val_help: egp; Exterior gateway protocol origin
	val_help: incomplete; Incomplete origin
	allowed: echo "igp egp incomplete"

	syntax:expression: $VAR(@) in "igp", "egp", "incomplete"; "origin must be one of igp, egp, or incomplete"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set origin  $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set origin "
*/
func quaggaPolicyRouteMapRuleSetOrigin(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set origin WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Border Gateway Protocol (BGP) originator ID attribute
	val_help: Orignator IP address

	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set originator-id  $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set originator-id "
*/
func quaggaPolicyRouteMapRuleSetOriginatorId(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set originator-id A.B.C.D
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Tag value for routing protocol
	val_help: u32:1-65535; Tag value

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "tag must be between 1 and 65535"
	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set tag $VAR(@)"

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set tag "
*/
func quaggaPolicyRouteMapRuleSetTag(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set tag WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Border Gateway Protocol (BGP) weight attribute
	val_help: BGP weight

	commit:expression: $VAR(../../action/) != ""; "you must specify an action"

	update: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "set weight $VAR(@) "

	delete: vtysh -c "configure terminal" \
	         -c "route-map $VAR(../../../@) $VAR(../../action/@) $VAR(../../@)" \
	         -c "no set weight "
*/
func quaggaPolicyRouteMapRuleSetWeight(Cmd int, Args cmd.Args) int {
	//policy route-map WORD rule WORD set weight WORD
	quaggaUpdateCheckRouteMap()
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:1
	priority: 730
	type: u32
	help: Border Gateway Protocol (BGP) parameters
	val_help: u32:1-4294967294; AS number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967294 ; \
			   "AS number must be between 1 and 4294967294"

	end: if [ -z "$VAR(.)" ] || [ "$COMMIT_ACTION" != DELETE ]; then
	       /opt/vyatta/sbin/vyatta-bgp.pl --main
	       vtysh -d bgpd -c 'sh run' > /opt/vyatta/etc/quagga/bgpd.conf
	     else
	       rm -f /opt/vyatta/etc/quagga/bgpd.conf
	     fi

SET: router bgp #3
DEL: no router bgp #3
*/
func quaggaProtocolsBgp(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD
	switch Cmd {
	case cmd.Set:
		if bgpConfigState != nil {
			return cmd.Success
		}
		bgpConfigState = newQuaggaBgp()
		bgpConfigState.asNum = fmt.Sprint(Args[0])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]))
	case cmd.Delete:
		if bgpConfigState == nil {
			return cmd.Success
		}
		quaggaVtysh("configure terminal",
			fmt.Sprint("no router bgp ", Args[0]))
		bgpConfigState = nil
	}
	return cmd.Success
}

/*
	tag:
	type: ipv6net
	help: BGP IPv6 aggregate network
	val_help: IPv6 aggregate network
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"

SET: router bgp #3 ; address-family ipv6 ; aggregate-address #7 ?summary-only
DEL: router bgp #3 ; address-family ipv6 ; no aggregate-address #7
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastAggregateAddress(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast aggregate-address X:X::X:X/M
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaBgpAggregateAddressCreate(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("aggregate-address ", Args[1]))
	case cmd.Delete:
		quaggaBgpAggregateAddressDelete(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no aggregate-address ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Announce the aggregate summary network only
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastAggregateAddressSummaryOnly(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast aggregate-address X:X::X:X/M summary-only
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, aggrAddr, ok := quaggaBgpAggregateAddressLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		aggrAddr.summaryOnly = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no aggregate-address ", Args[1]),
			fmt.Sprint("aggregate-address ", Args[1], " summary-only"))
	case cmd.Delete:
		aggrAddr.summaryOnly = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no aggregate-address ", Args[1], " summary-only"),
			fmt.Sprint("aggregate-address ", Args[1]))
	}
	return cmd.Success
}

/*
	tag:
	type: ipv6net
	help: BGP IPv6 network
	val_help: IPv6 network
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"

SET: router bgp #3 ; address-family ipv6 ; network #7
DEL: router bgp #3 ; address-family ipv6 ; no network #7
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastNetwork(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast network X:X::X:X/M
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaBgpNetworkCreate(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("network ", Args[1]))
	case cmd.Delete:
		quaggaBgpNetworkDelete(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no network ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: AS-path hopcount limit
	val_help: u32:0-255; AS path hop count limit

	commit:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "path-limit must be between 0-255."
	commit:expression: $VAR(../route-map/) == ""; "you can't set path-limit and route-map for network"

SET: router bgp #3 ; address-family ipv6 ; network #7 pathlimit #9
DEL: router bgp #3 ; address-family ipv6 ; no network #7 pathlimit #9
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastNetworkPathLimit(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast network X:X::X:X/M path-limit WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNetworkLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("network ", Args[1], " pathlimit ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no network ", Args[1], " pathlimit ", Args[2]))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to modify route attributes
	val_help: Route map name

	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
	commit:expression: $VAR(../path-limit/) == ""; "you can't set route-map and path-limit for network"

SET: router bgp #3 ; address-family ipv6 ; network #7 route-map #9
DEL: router bgp #3 ; address-family ipv6 ; no network #7 route-map #9
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastNetworkRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast network X:X::X:X/M route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNetworkLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("network ", Args[1], " route-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no network ", Args[1], " route-map ", Args[2]))
	}
	return cmd.Success
}

/*
	help: BGP IPv6 settings

SET:
DEL:
*/
func quaggaProtocolsBgpAddressFamilyIpv6Unicast(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
	val_help: Metric for redistributed routes

SET: router bgp #3 ; address-family ipv6 ; redistribute connected metric #9
DEL: router bgp #3 ; address-family ipv6 ; no redistribute connected metric #9
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeConnectedMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute connected metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute connected metric ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute connected metric ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute connected routes into BGP

SET: router bgp #3 ; address-family ipv6 ; redistribute connected
DEL: router bgp #3 ; address-family ipv6 ; no redistribute connected
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeConnected(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute connected
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"redistribute connected")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"no redistribute connected")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	val_help: Route map name

	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; redistribute connected route-map #9
DEL: router bgp #3 ; address-family ipv6 ; no redistribute connected route-map #9
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeConnectedRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute connected route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute connected route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute connected route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeKernelMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute kernel metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute kernel metric ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute kernel metric ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute kernel routes into BGP


SET: router bgp #3 ; address-family ipv6 ; no redistribute kernel ; redistribute kernel ?route-map ?metric
DEL: router bgp #3 ; address-family ipv6 ; no redistribute kernel
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeKernel(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute kernel
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"redistribute kernel")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"no redistribute kernel")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeKernelRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute kernel route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute kernel route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute kernel route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute routes from other protocols into BGP

SET:
DEL:
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistribute(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeOspfv3Metric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute ospfv3 metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute ospfv3 metric ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute ospfv3 metric ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute OSPFv3 routes into BGP

SET: router bgp #3 ; address-family ipv6 ; no redistribute ospf6 ; redistribute ospf6 ?route-map ?metric
DEL: router bgp #3 ; address-family ipv6 ; no redistribute ospf6
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeOspfv3(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute ospfv3
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"redistribute ospfv3")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"no redistribute ospfv3")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeOspfv3RouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute ospfv3 route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute ospfv3 route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute ospfv3 route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeRipngMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute ripng metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute ripng route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute ripng route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute RIPng routes into BGP

SET: router bgp #3 ; address-family ipv6 ; no redistribute ripng ; redistribute ripng ?route-map ?metric
DEL: router bgp #3 ; address-family ipv6 ; no redistribute ripng
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeRipng(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute ripng
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"redistribute ripng")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"no redistribute ripng")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeRipngRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute ripng route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute ripng route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute ripng route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeStaticMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute static metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute static metric ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute static metric ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute static routes into BGP

SET: router bgp #3 ; address-family ipv6 ; no redistribute static ; redistribute static ?route-map ?metric
DEL: router bgp #3 ; address-family ipv6 ; no redistribute static
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeStatic(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute static
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"redistribute static")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			"no redistribute static")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeStaticRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family ipv6-unicast redistribute static route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("redistribute static route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no redistribute static route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	help: BGP address-family parameters

SET:
DEL:
*/
func quaggaProtocolsBgpAddressFamily(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD address-family
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: ipv4net
	help: BGP aggregate network
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"

SET: router bgp #3 ; aggregate-address #5 ?as-set ?summary-only
DEL: router bgp #3 ; no aggregate-address #5 ?as-set ?summary-only
*/
func quaggaProtocolsBgpAggregateAddress(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD aggregate-address A.B.C.D/M
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaBgpAggregateAddressCreate(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("aggregate-address ", Args[1]))
	case cmd.Delete:
		quaggaBgpAggregateAddressDelete(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no aggregate-address ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Generate AS-set path information for this aggregate address
*/
func quaggaProtocolsBgpAggregateAddressAsSet(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD aggregate-address A.B.C.D/M as-set
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, aggrAddr, ok := quaggaBgpAggregateAddressLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	summaryOnly := ""
	if aggrAddr.summaryOnly {
		summaryOnly = " summary-only"
	}
	switch Cmd {
	case cmd.Set:
		aggrAddr.asSet = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no aggregate-address ", Args[1]),
			fmt.Sprint("aggregate-address ", Args[1], " as-set", summaryOnly))
	case cmd.Delete:
		aggrAddr.asSet = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no aggregate-address ", Args[1]),
			fmt.Sprint("aggregate-address ", Args[1], summaryOnly))
	}
	return cmd.Success
}

/*
	help: Announce the aggregate summary network only
*/
func quaggaProtocolsBgpAggregateAddressSummaryOnly(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD aggregate-address A.B.C.D/M summary-only
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, aggrAddr, ok := quaggaBgpAggregateAddressLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	asSet := ""
	if aggrAddr.asSet {
		asSet = " as-set"
	}
	switch Cmd {
	case cmd.Set:
		aggrAddr.summaryOnly = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no aggregate-address ", Args[1]),
			fmt.Sprint("aggregate-address ", Args[1], asSet, " summary-only"))
	case cmd.Delete:
		aggrAddr.summaryOnly = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no aggregate-address ", Args[1]),
			fmt.Sprint("aggregate-address ", Args[1], asSet))
	}
	return cmd.Success
}

/*
	type: u32
	help: Maximum ebgp multipaths
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between (1-255)"
	val_help: u32:1-255; EBGP multipaths

SET: router bgp #3 ; maximum-paths #6
DEL: router bgp #3 ; no maximum-paths #6
*/
func quaggaProtocolsBgpMaximumPathsEbgp(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD maximum-paths ebgp WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("maximum-paths ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no maximum-paths ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Maximum ibgp multipaths
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between (1-255)"
	val_help: u32:1-255; IBGP multipaths

SET: router bgp #3 ; maximum-paths ibgp #6
DEL: router bgp #3 ; no maximum-paths ibgp #6
*/
func quaggaProtocolsBgpMaximumPathsIbgp(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD maximum-paths ibgp WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("maximum-paths ibgp ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no maximum-paths ibgp ", Args[1]))
	}
	return cmd.Success
}

/*
	help: BGP multipaths

SET:
DEL:
*/
func quaggaProtocolsBgpMaximumPaths(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD maximum-paths
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: ipv4, ipv6
	help: BGP neighbor
	val_help: ipv4; BGP neighbor IP address
	val_help: ipv6; BGP neighbor IPv6 address

	syntax:expression: exec "/opt/vyatta/sbin/vyatta-bgp.pl \
	                           --check-neighbor-ip --neighbor $VAR(@)"

SET:
DEL: router bgp #3 ; no neighbor #5
*/
func quaggaProtocolsBgpNeighbor(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaBgpNeighborCreate(Args[1])
	case cmd.Delete:
		quaggaBgpNeighborDelete(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Accept a route that contains the local-AS in the as-path

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 allowas-in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 allowas-in
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAllowasIn(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast allowas-in
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " allowas-in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " allowas-in"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Number of occurrences of AS number
	val_help: u32:1-10; Number of times AS is allowed in path
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 10; "allowas-in number must be between 1 and 10"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 allowas-in #10
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 allowas-in ; neighbor #5 allowas-in
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAllowasInNumber(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast allowas-in number WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " allowas-in ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " allowas-in"),
			fmt.Sprint("neighbor ", Args[1], " allowas-in"))
	}
	return cmd.Success
}

/*
	help: Send AS path unchanged
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAttributeUnchangedAsPath(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast attribute-unchanged as-path
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !neigh.ipv6AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.ipv6AttributeUnchangedAsPath = true
	case cmd.Delete:
		neigh.ipv6AttributeUnchangedAsPath = false
	}
	asPath := ""
	if neigh.ipv6AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if neigh.ipv6AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if neigh.ipv6AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"address-family ipv6",
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send multi-exit discriminator unchanged
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAttributeUnchangedMed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast attribute-unchanged med
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !neigh.ipv6AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.ipv6AttributeUnchangedMed = true
	case cmd.Delete:
		neigh.ipv6AttributeUnchangedMed = false
	}
	asPath := ""
	if neigh.ipv6AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if neigh.ipv6AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if neigh.ipv6AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"address-family ipv6",
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send nexthop unchanged
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAttributeUnchangedNextHop(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast attribute-unchanged next-hop
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !neigh.ipv6AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.ipv6AttributeUnchangedNextHop = true
	case cmd.Delete:
		neigh.ipv6AttributeUnchangedNextHop = false
	}
	asPath := ""
	if neigh.ipv6AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if neigh.ipv6AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if neigh.ipv6AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"address-family ipv6",
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send BGP attributes unchanged

SET: router bgp #3 ; address-family ipv6 ; no neighbor #5 attribute-unchanged ; neighbor #5 attribute-unchanged ?as-path ?med ?next-hop
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 attribute-unchanged
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAttributeUnchanged(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast attribute-unchanged
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.ipv6AttributeUnchanged = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " attribute-unchanged"))
	case cmd.Delete:
		neigh.ipv6AttributeUnchanged = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"))
	}
	return cmd.Success
}

/*
	help: Advertise dynamic capability to this neighbor

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 capability dynamic
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 capability dynamic
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityDynamic(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast capability dynamic
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " capability dynamic"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " capability dynamic"))
	}
	return cmd.Success
}

/*
	help: Advertise capabilities to this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapability(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast capability
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise ORF capability to this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityOrf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast capability orf
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise prefix-list ORF capability to this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityOrfPrefixList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast capability orf prefix-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Capability to receive the ORF

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 capability orf prefix-list receive
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 capability orf prefix-list receive
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityOrfPrefixListReceive(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast capability orf prefix-list receive
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " capability orf prefix-list receive"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " capability orf prefix-list receive"))
	}
	return cmd.Success
}

/*
	help: Capability to send the ORF

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 capability orf prefix-list send
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 capability orf prefix-list send
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityOrfPrefixListSend(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast capability orf prefix-list send
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " capability orf prefix-list send"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " capability orf prefix-list send"))
	}
	return cmd.Success
}

/*
	help: Send default route to this neighbor

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 default-originate
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 default-originate
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDefaultOriginate(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast default-originate
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " default-originate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " default-originate"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to specify criteria of the default
	allowed: local -a params
	        params=$(/opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map)
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" " ; "route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 default-originate route-map #10
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 default-originate route-map #10
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDefaultOriginateRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast default-originate route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " default-originate route-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " default-originate route-map ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Disable sending extended community attributes to this neighbor

SET: router bgp #3 ; address-family ipv6 ; no neighbor #5 send-community extended
DEL: router bgp #3 ; address-family ipv6 ; neighbor #5 send-community extended
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDisableSendCommunityExtended(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast disable-send-community extended
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " send-community extended"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " send-community extended"))
	}
	return cmd.Success
}

/*
	help: Disable sending community attributes to this neighbor
	commit:expression: ($VAR(./extended/) != "") || ($VAR(./standard/) != ""); "you must specify the type of community"

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDisableSendCommunity(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast disable-send-community
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Disable sending standard community attributes to this neighbor

SET: router bgp #3 ; address-family ipv6 ; no neighbor #5 send-community standard
DEL: router bgp #3 ; address-family ipv6 ; neighbor #5 send-community standard
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDisableSendCommunityStandard(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast disable-send-community standard
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " send-community standard"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " send-community standard"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Access-list to filter outgoing route updates to this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy access-list6 )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list6 $VAR(@)\" "; "access-list6 $VAR(@) doesn't exist"
	commit:expression: $VAR(../../prefix-list/export/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 distribute-list #10 out
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 distribute-list #10 out
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDistributeListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast distribute-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " distribute-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " distribute-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Access-list to filter incoming route updates from this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy access-list6 )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list6 $VAR(@)\" "; "access-list6 $VAR(@) doesn't exist"
	commit:expression: $VAR(../../prefix-list/import/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 distribute-list #10 in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 distribute-list #10 in
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDistributeListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast distribute-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " distribute-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " distribute-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Access-list to filter route updates to/from this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDistributeList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast distribute-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: As-path-list to filter outgoing route updates to this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy as-path-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" ";"as-path-list $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 filter-list #10 out
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 filter-list #10 out
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastFilterListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast filter-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " filter-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " filter-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: As-path-list to filter incoming route updates from this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy as-path-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" " ; "as-path-list $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 filter-list #10 in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 filter-list #10 in
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastFilterListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast filter-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " filter-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " filter-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: As-path-list to filter route updates to/from this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastFilterList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast filter-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Maximum number of prefixes to accept from this neighbor
	val_help: u32:1-4294967295; Prefix limit
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967295; "maximum-prefix must be between 1 and 4294967295"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 maximum-prefix #9
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 maximum-prefix #9
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastMaximumPrefix(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast maximum-prefix WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " maximum-prefix ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " maximum-prefix ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Nexthop attributes

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 nexthop-local unchanged
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 nexthop-local unchanged
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastNexthopLocal(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast nexthop-local
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Leave link-local nexthop unchanged for this peer
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastNexthopLocalUnchanged(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast nexthop-local unchanged
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " nexthop-local unchanged"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " nexthop-local unchanged"))
	}
	return cmd.Success
}

/*
	help: Nexthop for routes sent to this neighbor to be the local router

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 next-hop-self
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 next-hop-self
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastNexthopSelf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast nexthop-self
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " next-hop-self"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " next-hop-self"))
	}
	return cmd.Success
}

/*
	help: BGP neighbor IPv6 parameters

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 activate
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 activate
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6Unicast(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " activate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " activate"))
	}
	return cmd.Success
}

/*
	type: txt
	help: IPv6 peer group for this peer
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-bgp.pl --list-peer-groups --as $VAR(../../../../@))
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"protocols bgp $VAR(../../../../@) peer-group $VAR(@)\" "; "protocols bgp $VAR(../../../../@) peer-group $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 peer-group #9
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 peer-group #9 ; neighbor #5 activate
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastPeerGroup(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast peer-group WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " peer-group ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " peer-group ", Args[2]),
			fmt.Sprint("neighbor ", Args[1], " activate"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Prefix-list to filter outgoing route updates to this neighbor

	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy prefix-list6 )
	        echo -n ${params[@]##* /}

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list6 $VAR(@)\" "; "prefix-list6 $VAR(@) doesn't exist"
	commit:expression: $VAR(../../distribute-list/export/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 prefix-list #10 out
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 prefix-list #10 out
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastPrefixListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast prefix-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " prefix-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " prefix-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Prefix-list to filter incoming route updates from this neighbor

	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy prefix-list6 )
	        echo -n ${params[@]##* /}

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list6 $VAR(@)\" "; "prefix-list6 $VAR(@) doesn't exist"
	commit:expression: $VAR(../../distribute-list/import/) == ""; "you can't set both a prefix-list and a distribute list"


SET: router bgp #3 ; address-family ipv6 ; neighbor #5 prefix-list #10 in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 prefix-list #10 in
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastPrefixListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast prefix-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " prefix-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " prefix-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Prefix-list to filter route updates to/from this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastPrefixList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast prefix-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Remove private AS numbers from AS path in outbound route updates
	commit:expression: $VAR(../../../remote-as/@) != $VAR(../../../../@); "you can't set remove-private-as for an iBGP peer"


SET: router bgp #3 ; address-family ipv6 ; neighbor #5 remove-private-AS
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 remove-private-AS
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRemovePrivateAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast remove-private-as
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " remove-private-AS"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " remove-private-AS"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to filter outgoing route updates to this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 route-map #10 out
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 route-map #10 out
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteMapExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast route-map export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " route-map ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " route-map ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to filter incoming route updates from this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 route-map #10 in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 route-map #10 in
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteMapImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast route-map import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " route-map ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " route-map ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Route-map to filter route updates to/from this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast route-map
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Neighbor as a route reflector client
	commit:expression: $VAR(../../../../@) == $VAR(../../../remote-as/@); "protocols bgp $VAR(../../../../@) neighbor $VAR(../../../@) route-reflector-client: remote-as must equal local-as"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 route-reflector-client
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 route-reflector-client
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteReflectorClient(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast route-reflector-client
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " route-reflector-client"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " route-reflector-client"))
	}
	return cmd.Success
}

/*
	help: Neighbor as route server client

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 route-server-client
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 route-server-client
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteServerClient(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast route-server-client
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " route-server-client"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " route-server-client"))
	}
	return cmd.Success
}

/*
	help: Inbound soft reconfiguration for this neighbor [REQUIRED]

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 soft-reconfiguration inbound
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 soft-reconfiguration inbound
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastSoftReconfigurationInbound(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast soft-reconfiguration inbound
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " soft-reconfiguration inbound"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " soft-reconfiguration inbound"))
	}
	return cmd.Success
}

/*
	help: Soft reconfiguration for neighbor
	commit:expression: $VAR(./inbound/) != ""; "you must specify the type of soft-reconfiguration"

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastSoftReconfiguration(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast soft-reconfiguration
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to selectively unsuppress suppressed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 unsuppress-map #9
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 unsuppress-map #9
*/
func quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastUnsuppressMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family ipv6-unicast unsuppress-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " unsuppress-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " unsuppress-map ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Parameters relating to IPv4 or IPv6 routes

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborAddressFamily(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD address-family
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Minimum interval for sending routing updates
	val_help: u32:0-600; Advertisement interval in seconds
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 600; "must be between 0 and 600"

SET: router bgp #3 ; neighbor #5 advertisement-interval #7
DEL: router bgp #3 ; no neighbor #5 advertisement-interval
*/
func quaggaProtocolsBgpNeighborAdvertisementInterval(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD advertisement-interval WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " advertisement-interval ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " advertisement-interval"))
	}
	return cmd.Success
}

/*
	help: Accept a route that contains the local-AS in the as-path

SET: router bgp #3 ; neighbor #5 allowas-in
DEL: router bgp #3 ; no neighbor #5 allowas-in
*/
func quaggaProtocolsBgpNeighborAllowasIn(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD allowas-in
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " allowas-in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " allowas-in"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Number of occurrences of AS number
	val_help: u32:1-10; Number of times AS is allowed in path
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 10; "allowas-in number must be between 1 and 10"

SET: router bgp #3 ; neighbor #5 allowas-in #8
DEL: router bgp #3 ; no neighbor #5 allowas-in ; neighbor #5 allowas-in
*/
func quaggaProtocolsBgpNeighborAllowasInNumber(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD allowas-in number WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " allowas-in ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " allowas-in"),
			fmt.Sprint("neighbor ", Args[1], " allowas-in"))
	}
	return cmd.Success
}

/*
	help: Send AS path unchanged
*/
func quaggaProtocolsBgpNeighborAttributeUnchangedAsPath(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD attribute-unchanged as-path
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !neigh.ipv4AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.ipv4AttributeUnchangedAsPath = true
	case cmd.Delete:
		neigh.ipv4AttributeUnchangedAsPath = false
	}
	asPath := ""
	if neigh.ipv4AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if neigh.ipv4AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if neigh.ipv4AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send multi-exit discriminator unchanged
*/
func quaggaProtocolsBgpNeighborAttributeUnchangedMed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD attribute-unchanged med
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !neigh.ipv4AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.ipv4AttributeUnchangedMed = true
	case cmd.Delete:
		neigh.ipv4AttributeUnchangedMed = false
	}
	asPath := ""
	if neigh.ipv4AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if neigh.ipv4AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if neigh.ipv4AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send nexthop unchanged
*/
func quaggaProtocolsBgpNeighborAttributeUnchangedNextHop(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD attribute-unchanged next-hop
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !neigh.ipv4AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.ipv4AttributeUnchangedNextHop = true
	case cmd.Delete:
		neigh.ipv4AttributeUnchangedNextHop = false
	}
	asPath := ""
	if neigh.ipv4AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if neigh.ipv4AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if neigh.ipv4AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: BGP attributes are sent unchanged

SET: router bgp #3 ; no neighbor #5 attribute-unchanged ; neighbor #5 attribute-unchanged ?as-path ?med ?next-hop
DEL: router bgp #3 ; no neighbor #5 attribute-unchanged ?as-path ?med ?next-hop
*/
func quaggaProtocolsBgpNeighborAttributeUnchanged(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD attribute-unchanged
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.ipv4AttributeUnchanged = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " attribute-unchanged"))
	case cmd.Delete:
		neigh.ipv4AttributeUnchanged = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"))
	}
	return cmd.Success
}

/*
	help: Advertise dynamic capability to this neighbor

SET: router bgp #3 ; neighbor #5 capability dynamic
DEL: router bgp #3 ; no neighbor #5 capability dynamic
*/
func quaggaProtocolsBgpNeighborCapabilityDynamic(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD capability dynamic
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " capability dynamic"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " capability dynamic"))
	}
	return cmd.Success
}

/*
	help: Advertise capabilities to this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborCapability(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD capability
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise ORF capability to this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborCapabilityOrf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD capability orf
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise prefix-list ORF capability to this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborCapabilityOrfPrefixList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD capability orf prefix-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Capability to receive the ORF

SET: router bgp #3 ; neighbor #5 capability orf prefix-list receive
DEL: router bgp #3 ; no neighbor #5 capability orf prefix-list receive
*/
func quaggaProtocolsBgpNeighborCapabilityOrfPrefixListReceive(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD capability orf prefix-list receive
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " capability orf prefix-list receive"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " capability orf prefix-list receive"))
	}
	return cmd.Success
}

/*
	help: Capability to send the ORF

SET: router bgp #3 ; neighbor #5 capability orf prefix-list send
DEL: router bgp #3 ; no neighbor #5 capability orf prefix-list send
*/
func quaggaProtocolsBgpNeighborCapabilityOrfPrefixListSend(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD capability orf prefix-list send
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " capability orf prefix-list send"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " capability orf prefix-list send"))
	}
	return cmd.Success
}

/*
	help: Send default route to this neighbor

SET: router bgp #3 ; neighbor #5 default-originate
DEL: router bgp #3 ; no neighbor #5 default-originate
*/
func quaggaProtocolsBgpNeighborDefaultOriginate(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD default-originate
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " default-originate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " default-originate"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to specify criteria of the default
	allowed: local -a params
	        params=$(/opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map)
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" " ; "route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 default-originate route-map #8
DEL: router bgp #3 ; no neighbor #5 default-originate route-map #8
*/
func quaggaProtocolsBgpNeighborDefaultOriginateRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD default-originate route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " default-originate route-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " default-originate route-map ", Args[2]))
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this neighbor
*/
func quaggaProtocolsBgpNeighborDescription(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD description WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Disable capability negotiation with this neighbor

SET: router bgp #3 ; neighbor #5 dont-capability-negotiate
DEL: router bgp #3 ; no neighbor #5 dont-capability-negotiate
*/
func quaggaProtocolsBgpNeighborDisableCapabilityNegotiation(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD disable-capability-negotiation
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " dont-capability-negotiate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " dont-capability-negotiate"))
	}
	return cmd.Success
}

/*
	help: Disable check to see if EBGP peer's address is a connected route

SET: router bgp #3 ; neighbor #5 disable-connected-check
DEL: router bgp #3 ; no neighbor #5 disable-connected-check
*/
func quaggaProtocolsBgpNeighborDisableConnectedCheck(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD disable-connected-check
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " disable-connected-check"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " disable-connected-check"))
	}
	return cmd.Success
}

/*
	help: Disable sending extended community attributes to this neighbor

SET: router bgp #3 ; no neighbor #5 send-community extended
DEL: router bgp #3 ; neighbor #5 send-community extended
*/
func quaggaProtocolsBgpNeighborDisableSendCommunityExtended(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD disable-send-community extended
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " send-community extended"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " send-community extended"))
	}
	return cmd.Success
}

/*
	help: Disable sending community attributes to this neighbor
	commit:expression: ($VAR(./extended/) != "") || ($VAR(./standard/) != ""); "you must specify the type of community"

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborDisableSendCommunity(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD disable-send-community
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Disable sending standard community attributes to this neighbor

SET: router bgp #3 ; no neighbor #5 send-community standard
DEL: router bgp #3 ; neighbor #5 send-community standard
*/
func quaggaProtocolsBgpNeighborDisableSendCommunityStandard(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD disable-send-community standard
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " send-community standard"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " send-community standard"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Access-list to filter outgoing route updates to this neighbor
	val_help: u32:1-65535; Access list number

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Access list must be between 1 and 65535"
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy access-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list $VAR(@)\" "; "access-list $VAR(@) doesn't exist"
	commit:expression: $VAR(../../prefix-list/export/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; neighbor #5 distribute-list #8 out
DEL: router bgp #3 ; no neighbor #5 distribute-list #8 out
*/
func quaggaProtocolsBgpNeighborDistributeListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD distribute-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " distribute-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " distribute-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Access-list to filter incoming route updates from this neighbor
	val_help: u32:1-65535; Access-list number
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Access list must be between 1 and 65535"

	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy access-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list $VAR(@)\" "; "access-list $VAR(@) doesn't exist"
	commit:expression: $VAR(../../prefix-list/import/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; neighbor #5 distribute-list #8 in
DEL: router bgp #3 ; no neighbor #5 distribute-list #8 in
*/
func quaggaProtocolsBgpNeighborDistributeListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD distribute-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " distribute-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " distribute-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Access-list to filter route updates to/from this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborDistributeList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD distribute-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Allow this EBGP neighbor to not be on a directly connected network
	val_help: u32:1-255; Number of hops

	syntax:expression: $VAR(@) >=1 && $VAR(@) <= 255; "ebgp-multihop must be between 1 and 255"
	commit:expression: $VAR(../ttl-security/hops/) == ""; "you can't set both ebgp-multihop and ttl-security hops"

SET: router bgp #3 ; neighbor #5 ebgp-multihop #7
DEL: router bgp #3 ; no neighbor #5 ebgp-multihop
*/
func quaggaProtocolsBgpNeighborEbgpMultihop(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD ebgp-multihop WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " ebgp-multihop ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " ebgp-multihop"))
	}
	return cmd.Success
}

/*
	type: txt
	help: As-path-list to filter outgoing route updates to this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy as-path-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" ";"as-path-list $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 filter-list #8 out
DEL: router bgp #3 ; no neighbor #5 filter-list #8 out
*/
func quaggaProtocolsBgpNeighborFilterListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD filter-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " filter-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " filter-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: As-path-list to filter incoming route updates from this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy as-path-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" ";"as-path-list $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 filter-list #8 in
DEL: router bgp #3 ; no neighbor #5 filter-list #8 in
*/
func quaggaProtocolsBgpNeighborFilterListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD filter-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " filter-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " filter-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: As-path-list to filter route updates to/from this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborFilterList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD filter-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:1
	type: u32
	help: Local AS number
	val_help: u32:1-4294967294; Local AS number
	syntax:expression: $VAR(@) >=1 && $VAR(@) <= 4294967294; "local-as must be between 1 and 4294967294"
	commit:expression: $VAR(@) != $VAR(../../@); "you can't set local-as the same as the router AS"
	commit:expression: exec "/opt/vyatta/sbin/vyatta-bgp.pl --is-iBGP --neighbor $VAR(../@) --as $VAR(../../@)"; "local-as can't be set on iBGP peers"

SET: router bgp #3 ; no neighbor #5 local-as #7 ; neighbor #5 local-as #7
DEL: router bgp #3 ; no neighbor #5 local-as
*/
func quaggaProtocolsBgpNeighborLocalAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD local-as WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " local-as ", Args[2]),
			fmt.Sprint("neighbor ", Args[1], " local-as ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " local-as"))
	}
	return cmd.Success
}

/*
	help: Do not prepend local-as to updates from EBGP peers

SET: router bgp #3 ; no neighbor #5 local-as #7 ; neighbor #5 local-as #7 no-prepend
DEL: router bgp #3 ; no neighbor #5 local-as #7 no-prepend ; neighbor #5 local-as #7
*/
func quaggaProtocolsBgpNeighborLocalAsNoPrepend(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD local-as WORD no-prepend
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " local-as ", Args[2]),
			fmt.Sprint("neighbor ", Args[1], " local-as ", Args[2], " no-prepend"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " local-as ", Args[2], " no-prepend"),
			fmt.Sprint("neighbor ", Args[1], " local-as ", Args[2]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Maximum number of prefixes to accept from this neighbor
	val_help: u32:1-4294967295; Prefix limit
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967295; "maximum-prefix must be between 1 and 4294967295"

SET: router bgp #3 ; neighbor #5 maximum-prefix #7
DEL: router bgp #3 ; no neighbor #5 maximum-prefix
*/
func quaggaProtocolsBgpNeighborMaximumPrefix(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD maximum-prefix WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " maximum-prefix ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " maximum-prefix"))
	}
	return cmd.Success
}

/*
	help: Nexthop for routes sent to this neighbor to be the local router

SET: router bgp #3 ; neighbor #5 next-hop-self
DEL: router bgp #3 ; no neighbor #5 next-hop-self
*/
func quaggaProtocolsBgpNeighborNexthopSelf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD nexthop-self
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " next-hop-self"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " next-hop-self"))
	}
	return cmd.Success
}

/*
	help: Ignore capability negotiation with specified neighbor
	commit:expression: $VAR(../strict-capability/) == ""; "you can't set both strict-capability and override-capability"

SET: router bgp #3 ; neighbor #5 override-capability
DEL: router bgp #3 ; no neighbor #5 override-capability
*/
func quaggaProtocolsBgpNeighborOverrideCapability(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD override-capability
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " override-capability"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " override-capability"))
	}
	return cmd.Success
}

/*
	help: Do not initiate a session with this neighbor

SET: router bgp #3 ; neighbor #5 passive
DEL: router bgp #3 ; no neighbor #5 passive
*/
func quaggaProtocolsBgpNeighborPassive(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD passive
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " passive"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " passive"))
	}
	return cmd.Success
}

/*
	type: txt
	help: BGP MD5 password
	syntax:expression: exec "			      \
	        if [ `echo -n '$VAR(@)' | wc -c` -gt 80 ]; then   \
	          echo Password must be 80 characters or less;\
	          exit 1 ;                                     \
	        fi ; "

SET: router bgp #3 ; neighbor #5 password #7
DEL: router bgp #3 ; no neighbor #5 password
*/
func quaggaProtocolsBgpNeighborPassword(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD password WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " password ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " password"))
	}
	return cmd.Success
}

/*
	type: txt
	help: IPv4 peer group for this peer
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-bgp.pl --list-peer-groups --as $VAR(../../@) )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"protocols bgp $VAR(../../@) peer-group $VAR(@)\" "; "protocols bgp $VAR(../../@) peer-group $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 peer-group #7
DEL: router bgp #3 ; no neighbor #5 peer-group #7 ; neighbor #5 activate
*/
func quaggaProtocolsBgpNeighborPeerGroup(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD peer-group WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " peer-group ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " peer-group ", Args[2]),
			fmt.Sprint("neighbor ", Args[1], " activate"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Neighbor's BGP port
	val_help: u32: 1-65535; Neighbor BGP port number
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; \
	       "port must be between 1 and 65535"

SET: router bgp #3 ; neighbor #5 port #7
DEL: router bgp #3 ; no neighbor #5 port
*/
func quaggaProtocolsBgpNeighborPort(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD port WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " port ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " port"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Prefix-list to filter outgoing route updates to this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy prefix-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list $VAR(@)\" "; "prefix-list $VAR(@) doesn't exist"
	commit:expression: $VAR(../../distribute-list/export/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; neighbor #5 prefix-list #8 out
DEL: router bgp #3 ; no neighbor #5 prefix-list #8 out
*/
func quaggaProtocolsBgpNeighborPrefixListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD prefix-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " prefix-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " prefix-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Prefix-list to filter incoming route updates from this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy prefix-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list $VAR(@)\" "; "prefix-list $VAR(@) doesn't exist"
	commit:expression: $VAR(../../distribute-list/import/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; neighbor #5 prefix-list #8 in
DEL: router bgp #3 ; no neighbor #5 prefix-list #8 in
*/
func quaggaProtocolsBgpNeighborPrefixListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD prefix-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " prefix-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " prefix-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Prefix-list to filter route updates to/from this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborPrefixList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD prefix-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Neighbor BGP AS number [REQUIRED]
	val_help: u32: 1-4294967294; Neighbor AS number
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967294; \
	                   "remote-as must be between 1 and 4294967294"

SET: router bgp #3 ; neighbor #5 remote-as #7 ; neighbor #5 activate
DEL: router bgp #3 ; no neighbor #5 remote-as #7
*/
func quaggaProtocolsBgpNeighborRemoteAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD remote-as WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " remote-as ", Args[2]),
			fmt.Sprint("neighbor ", Args[1], " activate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " remote-as ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Remove private AS numbers from AS path in outbound route updates
	commit:expression: $VAR(../remote-as/@) != $VAR(../../@); "you can't set remove-private-as for an iBGP peer"

SET: router bgp #3 ; neighbor #5 remove-private-AS
DEL: router bgp #3 ; no neighbor #5 remove-private-AS
*/
func quaggaProtocolsBgpNeighborRemovePrivateAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD remove-private-as
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " remove-private-AS"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " remove-private-AS"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to filter outgoing route updates to this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" "; "route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 route-map #8 out
DEL: router bgp #3 ; no neighbor #5 route-map #8 out
*/
func quaggaProtocolsBgpNeighborRouteMapExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD route-map export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " route-map ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " route-map ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to filter incoming route updates from this neighbor
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" "; "route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 route-map #8 in
DEL: router bgp #3 ; no neighbor #5 route-map #8 in
*/
func quaggaProtocolsBgpNeighborRouteMapImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD route-map import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " route-map ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " route-map ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Route-map to filter route updates to/from this neighbor

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD route-map
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Neighbor as a route reflector client
	commit:expression: $VAR(../../@) == $VAR(../remote-as/@); "remote-as must equal local-as"

SET: router bgp #3 ; neighbor #5 route-reflector-client
DEL: router bgp #3 ; no neighbor #5 route-reflector-client
*/
func quaggaProtocolsBgpNeighborRouteReflectorClient(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD route-reflector-client
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " route-reflector-client"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " route-reflector-client"))
	}
	return cmd.Success
}

/*
	help: Neighbor is route server client

SET: router bgp #3 ; neighbor #5 route-server-client
DEL: router bgp #3 ; no neighbor #5 route-server-client
*/
func quaggaProtocolsBgpNeighborRouteServerClient(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD route-server-client
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " route-server-client"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " route-server-client"))
	}
	return cmd.Success
}

/*
	help: Administratively shut down neighbor

SET: router bgp #3 ; neighbor #5 shutdown
DEL: router bgp #3 ; no neighbor #5 shutdown
*/
func quaggaProtocolsBgpNeighborShutdown(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD shutdown
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " shutdown"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " shutdown"))
	}
	return cmd.Success
}

/*
	help: Inbound soft reconfiguration for this neighbor [REQUIRED]

SET: router bgp #3 ; neighbor #5 soft-reconfiguration inbound
DEL: router bgp #3 ; no neighbor #5 soft-reconfiguration inbound
*/
func quaggaProtocolsBgpNeighborSoftReconfigurationInbound(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD soft-reconfiguration inbound
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " soft-reconfiguration inbound"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " soft-reconfiguration inbound"))
	}
	return cmd.Success
}

/*
	help: Soft reconfiguration for neighbor
	commit:expression: $VAR(./inbound/) != ""; "you must specify the type of soft-reconfiguration"

SET:
DEL:
*/
func quaggaProtocolsBgpNeighborSoftReconfiguration(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD soft-reconfiguration
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Enable strict capability negotiation
	commit:expression: $VAR(../override-capability/) == ""; "you can't set both strict-capability and override-capability"

SET: router bgp #3 ; neighbor #5 strict-capability-match
DEL: router bgp #3 ; no neighbor #5 strict-capability-match
*/
func quaggaProtocolsBgpNeighborStrictCapabilityMatch(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD strict-capability-match
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " strict-capability-match"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " strict-capability-match"))
	}
	return cmd.Success
}

/*
	type: u32
	help: BGP connect timer for this neighbor
	val_help: u32:1-65535; Connect timer in seconds
	val_help: 0; Disable connect timer
	syntax:expression: $VAR(@) >=0 && $VAR(@) <= 65535; "BGP connect timer must be between 0 and 65535"

SET: router bgp #3 ; neighbor #5 timers connect #8
DEL: router bgp #3 ; no neighbor #5 timers connect
*/
func quaggaProtocolsBgpNeighborTimersConnect(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD timers connect WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " timers connect ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " timers connect"))
	}
	return cmd.Success
}

/*
	type: u32
	default: 180
	help: BGP hold timer for this neighbor
	val_help: u32:1-65535; Hold timer in seconds
	val_help: 0; Disable hold timer
	syntax:expression: $VAR(@) == 0 || ($VAR(@) >= 4 && $VAR(@) <= 65535); "Holdtime interval must be 0 or between 4 and 65535"
*/
func quaggaProtocolsBgpNeighborTimersHoldtime(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD timers holdtime WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !neigh.timers {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.timersHoldtime = fmt.Sprint(Args[2])
	case cmd.Delete:
		neigh.timersHoldtime = ""
	}
	if neigh.timersKeepalive != "" && neigh.timersHoldtime != "" {
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " timers ", neigh.timersKeepalive, " ", neigh.timersHoldtime))
	} else {
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " timers"))
	}
	return cmd.Success
}

/*
	type: u32
	default: 60
	help: BGP keepalive interval for this neighbor
	val_help: u32:1-65535; Keepalive interval in seconds (default 60)
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Keepalive interval must be between 1 and 65535"
*/
func quaggaProtocolsBgpNeighborTimersKeepalive(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD timers keepalive WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !neigh.timers {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.timersKeepalive = fmt.Sprint(Args[2])
	case cmd.Delete:
		neigh.timersKeepalive = ""
	}
	if neigh.timersKeepalive != "" && neigh.timersHoldtime != "" {
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " timers ", neigh.timersKeepalive, " ", neigh.timersHoldtime))
	} else {
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " timers"))
	}
	return cmd.Success
}

/*
	help: Neighbor timers
	# TODO: fix this.  Can set connect &&|| (keepalive && holdtime)
	commit:expression: $VAR(./keepalive/) != ""; "you must set a keepalive interval"
	commit:expression: $VAR(./holdtime/) != ""; "you must set a holdtime interval"


SET: router bgp #3 ; neighbor #5 timers @keepalive @holdtime
DEL: router bgp #3 ; no neighbor #5 timers
*/
func quaggaProtocolsBgpNeighborTimers(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD timers
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, neigh, ok := quaggaBgpNeighborLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		neigh.timers = true
	case cmd.Delete:
		neigh.timers = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " timers"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Number of the maximum number of hops to the BGP peer
	val_help: u32:1-254; Number of hops

	syntax:expression: $VAR(@) >=1 && $VAR(@) <= 254; "ttl-security hops must be between 1 and 254"
	commit:expression: $VAR(../../ebgp-multihop/) == ""; "you can't set both ebgp-multihop and ttl-security hops"

SET: router bgp #3 ; neighbor #5 ttl-security hops #8
DEL: router bgp #3 ; no neighbor #5 ttl-security hops #8
*/
func quaggaProtocolsBgpNeighborTtlSecurityHops(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD ttl-security hops WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " ttl-security hops ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " ttl-security hops ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Ttl security mechanism for this BGP peer


SET:
DEL:
*/
func quaggaProtocolsBgpNeighborTtlSecurity(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD ttl-security
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to selectively unsuppress suppressed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 unsuppress-map #7
DEL: router bgp #3 ; no neighbor #5 unsuppress-map #7
*/
func quaggaProtocolsBgpNeighborUnsuppressMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD unsuppress-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " unsuppress-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " unsuppress-map ", Args[2]))
	}
	return cmd.Success
}

/*
	type: txt
	help: Source IP of routing updates
	val_help: ipv4; IP address of route source
	val_help: <interface>;  Interface as route source
	commit:expression: exec "/opt/vyatta/sbin/vyatta-bgp.pl --check-source $VAR(@)"

SET: router bgp #3 ; neighbor #5 update-source #7
DEL: router bgp #3 ; no neighbor #5 update-source
*/
func quaggaProtocolsBgpNeighborUpdateSource(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD update-source WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " update-source ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " update-source"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Default weight for routes from this neighbor
	val_help: u32: 1-65535; Weight for routes from this neighbor
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "weight must be between 1 and 65535"

SET: router bgp #3 ; neighbor #5 weight #7
DEL: router bgp #3 ; no neighbor #5 weight
*/
func quaggaProtocolsBgpNeighborWeight(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD neighbor WORD weight WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpNeighborLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " weight ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " weight"))
	}
	return cmd.Success
}

/*
	tag:
	type: ipv4net
	help: BGP network
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"
	commit:expression: !($VAR(./backdoor/) != "" && $VAR(./route-map/) != ""); "you may specify route-map or backdoor but not both"

SET: router bgp #3 ; network #5 ?backdoor
DEL: router bgp #3 ; no network #5
*/
func quaggaProtocolsBgpNetwork(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD network A.B.C.D/M
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaBgpNetworkCreate(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("network ", Args[1]))
	case cmd.Delete:
		quaggaBgpNetworkDelete(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no network ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Network as a backdoor route
*/
func quaggaProtocolsBgpNetworkBackdoor(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD network A.B.C.D/M backdoor
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, net, ok := quaggaBgpNetworkLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		net.backdoor = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("network ", Args[1], " backdoor"))
	case cmd.Delete:
		net.backdoor = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no network ", Args[1], " backdoor"),
			fmt.Sprint("network ", Args[1]))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to modify route attributes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; network #5 route-map #7
DEL: router bgp #3 ; no network #5 route-map #7 ; network #5
*/
func quaggaProtocolsBgpNetworkRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD network A.B.C.D/M route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	_, net, ok := quaggaBgpNetworkLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		net.routeMap = fmt.Sprint(Args[2])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("network ", Args[1], " route-map ", Args[2]))
	case cmd.Delete:
		net.routeMap = ""
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no network ", Args[1], " route-map ", Args[2]),
			fmt.Sprint("network ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Always compare MEDs from different neighbors

SET: router bgp #3 ; bgp always-compare-med
DEL: router bgp #3 ; no bgp always-compare-med
*/
func quaggaProtocolsBgpParametersAlwaysCompareMed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters always-compare-med
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp always-compare-med")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp always-compare-med")
	}
	return cmd.Success
}

/*
	help: Compare AS-path lengths including confederation sets & sequences

SET: router bgp #3 ; bgp bestpath as-path confed
DEL: router bgp #3 ; no bgp bestpath as-path confed
*/
func quaggaProtocolsBgpParametersBestpathAsPathConfed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters bestpath as-path confed
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp bestpath as-path confed")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp bestpath as-path confed")
	}
	return cmd.Success
}

/*
	help: Ignore AS-path length in selecting a route

SET: router bgp #3 ; bgp bestpath as-path ignore
DEL: router bgp #3 ; no bgp bestpath as-path ignore
*/
func quaggaProtocolsBgpParametersBestpathAsPathIgnore(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters bestpath as-path ignore
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp bestpath as-path ignore")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp bestpath as-path ignore")
	}
	return cmd.Success
}

/*
	help: AS-path attribute comparison parameters

SET:
DEL:
*/
func quaggaProtocolsBgpParametersBestpathAsPath(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters bestpath as-path
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Compare the router-id for identical EBGP paths

SET: router bgp #3 ; bgp bestpath compare-routerid
DEL: router bgp #3 ; no bgp bestpath compare-routerid
*/
func quaggaProtocolsBgpParametersBestpathCompareRouterid(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters bestpath compare-routerid
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp bestpath compare-routerid")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp bestpath compare-routerid")
	}
	return cmd.Success
}

/*
	help: Compare MEDs among confederation paths

SET: router bgp #3 ; bgp bestpath med confed
DEL: router bgp #3 ; no bgp bestpath med confed
*/
func quaggaProtocolsBgpParametersBestpathMedConfed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters bestpath med confed
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp bestpath med confed")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp bestpath med confed")
	}
	return cmd.Success
}

/*
	help: Treat missing route as a MED as the least preferred one

SET: router bgp #3 ; bgp bestpath med missing-as-worst
DEL: router bgp #3 ; no bgp bestpath med missing-as-worst
*/
func quaggaProtocolsBgpParametersBestpathMedMissingAsWorst(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters bestpath med missing-as-worst
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp bestpath med missing-as-worst")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp bestpath med missing-as-worst")
	}
	return cmd.Success
}

/*
	help: MED attribute comparison parameters

SET:
DEL:
*/
func quaggaProtocolsBgpParametersBestpathMed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters bestpath med
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Default bestpath selection mechanism

SET:
DEL:
*/
func quaggaProtocolsBgpParametersBestpath(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters bestpath
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Route-reflector cluster-id

SET: router bgp #3 ; bgp cluster-id #6
DEL: router bgp #3 ; no bgp cluster-id #6
*/
func quaggaProtocolsBgpParametersClusterId(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters cluster-id A.B.C.D
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("bgp cluster-id ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no bgp cluster-id ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Confederation AS identifier [REQUIRED]
	val_help: u32:1-4294967294; Confederation AS id
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967294; "confederation AS id must be between 1 and 4294967294"

SET: router bgp #3 ; bgp confederation identifier #7
DEL: router bgp #3 ; no bgp confederation identifier #7
*/
func quaggaProtocolsBgpParametersConfederationIdentifier(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters confederation identifier WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("bgp confederation identifier ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no bgp confederation identifier ", Args[1]))
	}
	return cmd.Success
}

/*
	help: AS confederation parameters

SET:
DEL:
*/
func quaggaProtocolsBgpParametersConfederation(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters confederation
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	multi:
	type: u32
	help: Peer ASs in the BGP confederation
	val_help: u32:1-4294967294; Peer AS number
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967294; "confederation AS id must be between 1 and 4294967294"
	commit:expression: exec "/opt/vyatta/sbin/vyatta-bgp.pl --confed-iBGP-ASN-check $VAR(@) --as $VAR(../../../@)"; "Can't set confederation peers ASN to $VAR(@).  Delete any neighbors with remote-as $VAR(@) and/or change the local ASN first."

SET: router bgp #3 ; bgp confederation peers #7
DEL: router bgp #3 ; no bgp confederation peers #7
*/
func quaggaProtocolsBgpParametersConfederationPeers(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters confederation peers WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("bgp confederation peers ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no bgp confederation peers ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Half-life time for dampening [REQUIRED]
	val_help: u32:1-45; Half-life penalty in seconds
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 45; "Half-life penalty must be between 1 and 45"
*/
func quaggaProtocolsBgpParametersDampeningHalfLife(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters dampening half-life WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.parametersDampening {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDampeningHalfLife = fmt.Sprint(Args[1])
	case cmd.Delete:
		bgpConfigState.parametersDampeningHalfLife = ""
	}
	halfLife := ""
	reUse := ""
	startSuppressTime := ""
	maxSuppressTime := ""
	halfLife = fmt.Sprint(" ", bgpConfigState.parametersDampeningHalfLife)
	if bgpConfigState.parametersDampeningHalfLife != "" &&
		bgpConfigState.parametersDampeningReUse != "" &&
		bgpConfigState.parametersDampeningStartSuppressTime != "" &&
		bgpConfigState.parametersDampeningMaxSuppressTime != "" {
		reUse = fmt.Sprint(" ", bgpConfigState.parametersDampeningReUse)
		startSuppressTime = fmt.Sprint(" ", bgpConfigState.parametersDampeningStartSuppressTime)
		maxSuppressTime = fmt.Sprint(" ", bgpConfigState.parametersDampeningMaxSuppressTime)
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"no bgp dampening",
		fmt.Sprint("bgp dampening ", halfLife, reUse, startSuppressTime, maxSuppressTime))
	return cmd.Success
}

/*
	type: u32
	help: Maximum duration to suppress a stable route [REQUIRED]
	val_help: u32:1-255; Maximum suppress duration in seconds

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Max-suppress-time must be between 1 and 255"
	commit:expression: $VAR(../re-use/) != ""; "you must set a re-use time"
	commit:expression: $VAR(../start-suppress-time/) != ""; "you must set a start-suppress-time"
*/
func quaggaProtocolsBgpParametersDampeningMaxSuppressTime(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters dampening max-suppress-time WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.parametersDampening {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDampeningMaxSuppressTime = fmt.Sprint(Args[1])
	case cmd.Delete:
		bgpConfigState.parametersDampeningMaxSuppressTime = ""
	}
	halfLife := ""
	reUse := ""
	startSuppressTime := ""
	maxSuppressTime := ""
	halfLife = fmt.Sprint(" ", bgpConfigState.parametersDampeningHalfLife)
	if bgpConfigState.parametersDampeningHalfLife != "" &&
		bgpConfigState.parametersDampeningReUse != "" &&
		bgpConfigState.parametersDampeningStartSuppressTime != "" &&
		bgpConfigState.parametersDampeningMaxSuppressTime != "" {
		reUse = fmt.Sprint(" ", bgpConfigState.parametersDampeningReUse)
		startSuppressTime = fmt.Sprint(" ", bgpConfigState.parametersDampeningStartSuppressTime)
		maxSuppressTime = fmt.Sprint(" ", bgpConfigState.parametersDampeningMaxSuppressTime)
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"no bgp dampening",
		fmt.Sprint("bgp dampening ", halfLife, reUse, startSuppressTime, maxSuppressTime))
	return cmd.Success
}

/*
	help: Enable route-flap dampening
	# Note that there is a bug in quagga here.  If bgpd gets two 'no bgp dampening'
	# commands in a row it will crash

	commit:expression: $VAR(./half-life/) != "" || $VAR(./max-suppress-time/) != "" || \
	                   $VAR(./re-use/) != "" || $VAR(./start-suppress-time/) != "" ; \
	                   "To define dampening, all parameters must be set"

SET: router bgp #3 ; no bgp dampening ; bgp dampening @half-life @re-use @start-suppress-time @max-suppress-time
DEL: router bgp #3 ; no bgp dampening
*/
func quaggaProtocolsBgpParametersDampening(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters dampening
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDampening = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp dampening",
			"bgp dampening")
	case cmd.Delete:
		bgpConfigState.parametersDampening = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp dampening")
	}
	return cmd.Success
}

/*
	type: u32
	help: Time to start reusing a route [REQUIRED]
	val_help: u32:1-20000; Re-use time in seconds

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 20000; "Re-use value must be between 1 and 20000"
	commit:expression: $VAR(../start-suppress-time/) != ""; "you must set start-suppress-time"
	commit:expression: $VAR(../max-suppress-time/) != ""; "you must set max-suppress-time"
*/
func quaggaProtocolsBgpParametersDampeningReUse(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters dampening re-use WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.parametersDampening {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDampeningReUse = fmt.Sprint(Args[1])
	case cmd.Delete:
		bgpConfigState.parametersDampeningReUse = ""
	}
	halfLife := ""
	reUse := ""
	startSuppressTime := ""
	maxSuppressTime := ""
	halfLife = fmt.Sprint(" ", bgpConfigState.parametersDampeningHalfLife)
	if bgpConfigState.parametersDampeningHalfLife != "" &&
		bgpConfigState.parametersDampeningReUse != "" &&
		bgpConfigState.parametersDampeningStartSuppressTime != "" &&
		bgpConfigState.parametersDampeningMaxSuppressTime != "" {
		reUse = fmt.Sprint(" ", bgpConfigState.parametersDampeningReUse)
		startSuppressTime = fmt.Sprint(" ", bgpConfigState.parametersDampeningStartSuppressTime)
		maxSuppressTime = fmt.Sprint(" ", bgpConfigState.parametersDampeningMaxSuppressTime)
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"no bgp dampening",
		fmt.Sprint("bgp dampening ", halfLife, reUse, startSuppressTime, maxSuppressTime))
	return cmd.Success
}

/*
	type: u32
	help: When to start suppressing a route [REQUIRED]
	val_help: u32:1-20000; Start-suppress-time
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 20000; "Start-suppress-time must be between 1 and 20000"
	commit:expression: $VAR(../re-use/) != ""; "you must set re-use"
	commit:expression: $VAR(../max-suppress-time/) != ""; "you must set max-suppress-time"
*/
func quaggaProtocolsBgpParametersDampeningStartSuppressTime(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters dampening start-suppress-time WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.parametersDampening {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDampeningStartSuppressTime = fmt.Sprint(Args[1])
	case cmd.Delete:
		bgpConfigState.parametersDampeningStartSuppressTime = ""
	}
	halfLife := ""
	reUse := ""
	startSuppressTime := ""
	maxSuppressTime := ""
	halfLife = fmt.Sprint(" ", bgpConfigState.parametersDampeningHalfLife)
	if bgpConfigState.parametersDampeningHalfLife != "" &&
		bgpConfigState.parametersDampeningReUse != "" &&
		bgpConfigState.parametersDampeningStartSuppressTime != "" &&
		bgpConfigState.parametersDampeningMaxSuppressTime != "" {
		reUse = fmt.Sprint(" ", bgpConfigState.parametersDampeningReUse)
		startSuppressTime = fmt.Sprint(" ", bgpConfigState.parametersDampeningStartSuppressTime)
		maxSuppressTime = fmt.Sprint(" ", bgpConfigState.parametersDampeningMaxSuppressTime)
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"no bgp dampening",
		fmt.Sprint("bgp dampening ", halfLife, reUse, startSuppressTime, maxSuppressTime))
	return cmd.Success
}

/*
	type: u32
	help: Default local preference (higher=more preferred)
	val_help: u32:0-4294967295; Local preference

SET: router bgp #3 ; bgp default local-preference #7
DEL: router bgp #3 ; no bgp default local-preference #7
*/
func quaggaProtocolsBgpParametersDefaultLocalPref(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters default local-pref WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("bgp default local-preference ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no bgp default local-preference ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Deactivate IPv4 unicast for a peer by default

SET: router bgp #3 ; no bgp default ipv4-unicast
DEL: router bgp #3 ; bgp default ipv4-unicast
*/
func quaggaProtocolsBgpParametersDefaultNoIpv4Unicast(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters default no-ipv4-unicast
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp default ipv4-unicast")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp default ipv4-unicast")
	}
	return cmd.Success
}

/*
	help: BGP defaults

SET:
DEL:
*/
func quaggaProtocolsBgpParametersDefault(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters default
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Compare MEDs between different peers in the same AS

SET: router bgp #3 ; bgp deterministic-med
DEL: router bgp #3 ; no bgp deterministic-med
*/
func quaggaProtocolsBgpParametersDeterministicMed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters deterministic-med
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp deterministic-med")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp deterministic-med")
	}
	return cmd.Success
}

/*
	help: Disable IGP route check for network statements

SET: router bgp #3 ; no bgp network import-check
DEL: router bgp #3 ; bgp network import-check
*/
func quaggaProtocolsBgpParametersDisableNetworkImportCheck(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters disable-network-import-check
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp network import-check")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp network import-check")
	}
	return cmd.Success
}

/*
	type: u32
	help: Administrative distance for external BGP routes
	val_help: u32:1-255; Administrative distance for external BGP routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between 1-255"
*/
func quaggaProtocolsBgpParametersDistanceGlobalExternal(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters distance global external WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.parametersDistanceGlobal {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDistanceGlobalExternal = fmt.Sprint(Args[1])
	case cmd.Delete:
		bgpConfigState.parametersDistanceGlobalExternal = ""
	}
	external := ""
	internal := ""
	local := ""
	if bgpConfigState.parametersDistanceGlobalExternal != "" &&
		bgpConfigState.parametersDistanceGlobalInternal != "" &&
		bgpConfigState.parametersDistanceGlobalLocal != "" {
		external = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalExternal)
		internal = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalInternal)
		local = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalLocal)
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"no distance bgp",
		fmt.Sprint("distance bgp ", external, internal, local))
	return cmd.Success
}

/*
	type: u32
	help: Administrative distance for internal BGP routes
	val_help: u32:1-255; Administrative distance for internal BGP routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between 1-255"
*/
func quaggaProtocolsBgpParametersDistanceGlobalInternal(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters distance global internal WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.parametersDistanceGlobal {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDistanceGlobalInternal = fmt.Sprint(Args[1])
	case cmd.Delete:
		bgpConfigState.parametersDistanceGlobalInternal = ""
	}
	external := ""
	internal := ""
	local := ""
	if bgpConfigState.parametersDistanceGlobalExternal != "" &&
		bgpConfigState.parametersDistanceGlobalInternal != "" &&
		bgpConfigState.parametersDistanceGlobalLocal != "" {
		external = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalExternal)
		internal = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalInternal)
		local = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalLocal)
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"no distance bgp",
		fmt.Sprint("distance bgp ", external, internal, local))
	return cmd.Success
}

/*
	type: u32
	help: Administrative distance for local BGP routes
	val_help: u32:1-255; Administrative distance for local BGP routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between 1-255"
*/
func quaggaProtocolsBgpParametersDistanceGlobalLocal(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters distance global local WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.parametersDistanceGlobal {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDistanceGlobalLocal = fmt.Sprint(Args[1])
	case cmd.Delete:
		bgpConfigState.parametersDistanceGlobalLocal = ""
	}
	external := ""
	internal := ""
	local := ""
	if bgpConfigState.parametersDistanceGlobalExternal != "" &&
		bgpConfigState.parametersDistanceGlobalInternal != "" &&
		bgpConfigState.parametersDistanceGlobalLocal != "" {
		external = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalExternal)
		internal = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalInternal)
		local = fmt.Sprint(" ", bgpConfigState.parametersDistanceGlobalLocal)
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"no distance bgp",
		fmt.Sprint("distance bgp ", external, internal, local))
	return cmd.Success
}

/*
	help: Global administratives distances for BGP routes
	commit:expression: $VAR(./external/) != ""; "you must set an external route distance"
	commit:expression: $VAR(./internal/) != ""; "you must set an internal route distance"
	commit:expression: $VAR(./local/) != ""; "you must set a local route distance"

SET: router bgp #3 ; distance bgp @external @internal @local
DEL: router bgp #3 ; no distance bgp
*/
func quaggaProtocolsBgpParametersDistanceGlobal(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters distance global
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.parametersDistanceGlobal = true
	case cmd.Delete:
		bgpConfigState.parametersDistanceGlobal = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no distance bgp")

	}
	return cmd.Success
}

/*
	help: Administratives distances for BGP routes

SET:
DEL:
*/
func quaggaProtocolsBgpParametersDistance(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters distance
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: ipv4net
	help: Administrative distance for a specific BGP prefix
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"
	commit:expression: $VAR(./distance/) != ""; "you must set a route distance for this prefix"

SET:
DEL:
*/
func quaggaProtocolsBgpParametersDistancePrefix(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters distance prefix A.B.C.D/M
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Administrative distance for prefix
	val_help: u32:1-255; Administrative distance for external BGP routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between 1-255"

SET: router bgp #3 ; distance #9 #7
DEL: router bgp #3 ; no distance #9 #7
*/
func quaggaProtocolsBgpParametersDistancePrefixDistance(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters distance prefix A.B.C.D/M distance WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("distance ", Args[2], Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no distance ", Args[2], Args[1]))
	}
	return cmd.Success
}

/*
	help: Require first AS in the path to match peer's AS

SET: router bgp #3 ; bgp enforce-first-as
DEL: router bgp #3 ; no bgp enforce-first-as
*/
func quaggaProtocolsBgpParametersEnforceFirstAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters enforce-first-as
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp enforce-first-as")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp enforce-first-as")
	}
	return cmd.Success
}

/*
	help: Graceful restart capability parameters

SET:
DEL:
*/
func quaggaProtocolsBgpParametersGracefulRestart(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters graceful-restart
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Maximum time to hold onto restarting peer's stale paths
	val_help: u32:1-3600; Hold time in seconds
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 3600; "stalepath-time must be between 1 and 3600"

SET: router bgp #3 ; bgp graceful-restart stalepath-time #7
DEL: router bgp #3 ; no bgp graceful-restart stalepath-time #7
*/
func quaggaProtocolsBgpParametersGracefulRestartStalepathTime(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters graceful-restart stalepath-time WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("bgp graceful-restart stalepath-time ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no bgp graceful-restart stalepath-time ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Log neighbor up/down changes and reset reason

SET: router bgp #3 ; bgp log-neighbor-changes
DEL: router bgp #3 ; no bgp log-neighbor-changes
*/
func quaggaProtocolsBgpParametersLogNeighborChanges(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters log-neighbor-changes
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp log-neighbor-changes")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp log-neighbor-changes")
	}
	return cmd.Success
}

/*
	help: Disable client to client route reflection

SET: router bgp #3 ; no bgp client-to-client reflection
DEL: router bgp #3 ; bgp client-to-client reflection
*/
func quaggaProtocolsBgpParametersNoClientToClientReflection(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters no-client-to-client-reflection
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp client-to-client reflection")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp client-to-client reflection")
	}
	return cmd.Success
}

/*
	help: Disable immediate sesison reset if peer's connected link goes down

SET: router bgp #3 ; no bgp fast-external-failover
DEL: router bgp #3 ; bgp fast-external-failover
*/
func quaggaProtocolsBgpParametersNoFastExternalFailover(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters no-fast-external-failover
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no bgp fast-external-failover")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"bgp fast-external-failover")
	}
	return cmd.Success
}

/*
	help: BGP parameters

SET:
DEL:
*/
func quaggaProtocolsBgpParameters(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: BGP router id

SET: router bgp #3 ; bgp router-id #6
DEL: router bgp #3 ; no bgp router-id #6
*/
func quaggaProtocolsBgpParametersRouterId(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters router-id A.B.C.D
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("bgp router-id ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no bgp router-id ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: BGP route scanner interval
	val_help: u32:5-60; Scan interval in seconds
	syntax:expression: $VAR(@) >= 5 && $VAR(@) <= 60; "scan-time must be between 5 and 60 seconds"

SET: router bgp #3 ; bgp scan-time #6
DEL: router bgp #3 ; no bgp scan-time #6
*/
func quaggaProtocolsBgpParametersScanTime(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD parameters scan-time WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("bgp scan-time ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no bgp scan-time ", Args[1]))
	}
	return cmd.Success
}

/*
	tag:
	type: txt
	help: BGP peer-group
	syntax:expression: exec "/opt/vyatta/sbin/vyatta-bgp.pl \
	                           --check-peergroup-name $VAR(@)"
	delete:expression: exec "/opt/vyatta/sbin/vyatta-bgp.pl \
	                           --check-peer-groups --peergroup $VAR(@) --as $VAR(../@)"

SET: router bgp #3 ; neighbor #5 peer-group
DEL: router bgp #3 ; no neighbor #5 peer-group
*/
func quaggaProtocolsBgpPeerGroup(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaBgpPeerGroupCreate(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " peer-group"))
	case cmd.Delete:
		quaggaBgpPeerGroupDelete(Args[1])
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " peer-group"))
	}
	return cmd.Success
}

/*
	help: Accept a route that contains the local-AS in the as-path

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 allowas-in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 allowas-in
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAllowasIn(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast allowas-in
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " allowas-in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " allowas-in"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Number of occurrences of AS number
	val_help: u32:1-10; Number of times AS is allowed in path
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 10; "allowas-in number must be between 1 and 10"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 allowas-in #10
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 allowas-in ; neighbor #5 allowas-in
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAllowasInNumber(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast allowas-in number WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " allowas-in ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " allowas-in"),
			fmt.Sprint("neighbor ", Args[1], " allowas-in"))
	}
	return cmd.Success
}

/*
	help: Send AS path unchanged
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAttributeUnchangedAsPath(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast attribute-unchanged as-path
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, peerGrp, ok := quaggaBgpPeerGroupLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !peerGrp.ipv6AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		peerGrp.ipv6AttributeUnchangedAsPath = true
	case cmd.Delete:
		peerGrp.ipv6AttributeUnchangedAsPath = false
	}
	asPath := ""
	if peerGrp.ipv6AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if peerGrp.ipv6AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if peerGrp.ipv6AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"address-family ipv6",
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send multi-exit discriminator unchanged
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAttributeUnchangedMed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast attribute-unchanged med
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, peerGrp, ok := quaggaBgpPeerGroupLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !peerGrp.ipv6AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		peerGrp.ipv6AttributeUnchangedMed = true
	case cmd.Delete:
		peerGrp.ipv6AttributeUnchangedMed = false
	}
	asPath := ""
	if peerGrp.ipv6AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if peerGrp.ipv6AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if peerGrp.ipv6AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"address-family ipv6",
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send nexthop unchanged
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAttributeUnchangedNextHop(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast attribute-unchanged next-hop
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, peerGrp, ok := quaggaBgpPeerGroupLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !peerGrp.ipv6AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		peerGrp.ipv6AttributeUnchangedNextHop = true
	case cmd.Delete:
		peerGrp.ipv6AttributeUnchangedNextHop = false
	}
	asPath := ""
	if peerGrp.ipv6AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if peerGrp.ipv6AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if peerGrp.ipv6AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		"address-family ipv6",
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send BGP attributes unchanged

SET: router bgp #3 ; address-family ipv6 ; no neighbor #5 attribute-unchanged ; neighbor #5 attribute-unchanged ?as-path ?med ?next-hop
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 attribute-unchanged
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAttributeUnchanged(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast attribute-unchanged
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, peerGrp, ok := quaggaBgpPeerGroupLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		peerGrp.ipv6AttributeUnchanged = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " attribute-unchanged"))
	case cmd.Delete:
		peerGrp.ipv6AttributeUnchanged = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"))
	}
	return cmd.Success
}

/*
	help: Advertise dynamic capability to this peer-group

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 capability dynamic
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 capability dynamic
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityDynamic(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast capability dynamic
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " capability dynamic"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " capability dynamic"))
	}
	return cmd.Success
}

/*
	help: Advertise capabilities to this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapability(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast capability
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise ORF capability to this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityOrf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast capability orf
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise prefix-list ORF capability to this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityOrfPrefixList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast capability orf prefix-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Capability to receive the ORF

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 capability orf prefix-list receive
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 capability orf prefix-list receive
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityOrfPrefixListReceive(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast capability orf prefix-list receive
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " capability orf prefix-list receive"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " capability orf prefix-list receive"))
	}
	return cmd.Success
}

/*
	help: Capability to send the ORF

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 capability orf prefix-list send
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 capability orf prefix-list send
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityOrfPrefixListSend(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast capability orf prefix-list send
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " capability orf prefix-list send"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " capability orf prefix-list send"))
	}
	return cmd.Success
}

/*
	help: Send default route to this peer-group

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 default-originate
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 default-originate
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDefaultOriginate(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast default-originate
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " default-originate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " default-originate"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to specify criteria of the default
	allowed: local -a params
	        params=$(/opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map)
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" " ; "route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 default-originate route-map #10
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 default-originate route-map #10
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDefaultOriginateRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast default-originate route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " default-originate route-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " default-originate route-map ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Disable sending extended community attributes to this peer-group

SET: router bgp #3 ; address-family ipv6 ; no neighbor #5 send-community extended
DEL: router bgp #3 ; address-family ipv6 ; neighbor #5 send-community extended
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDisableSendCommunityExtended(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast disable-send-community extended
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " send-community extended"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " send-community extended"))
	}
	return cmd.Success
}

/*
	help: Disable sending community attributes to this peer-group
	commit:expression: ($VAR(./extended/) != "") || ($VAR(./standard/) != ""); "you must specify the type of community"

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDisableSendCommunity(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast disable-send-community
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Disable sending standard community attributes to this peer-group

SET: router bgp #3 ; address-family ipv6 ; no neighbor #5 send-community standard
DEL: router bgp #3 ; address-family ipv6 ; neighbor #5 send-community standard
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDisableSendCommunityStandard(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast disable-send-community standard
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " send-community standard"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " send-community standard"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Access-list to filter outgoing route updates to this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy access-list6 )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list6 $VAR(@)\" "; "access-list6 $VAR(@) doesn't exist"
	commit:expression: $VAR(../../prefix-list/export/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 distribute-list #10 out
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 distribute-list #10 out
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDistributeListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast distribute-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " distribute-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " distribute-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Access-list to filter incoming route updates from this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy access-list6 )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list6 $VAR(@)\" "; "access-list6 $VAR(@) doesn't exist"
	commit:expression: $VAR(../../prefix-list/import/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 distribute-list #10 in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 distribute-list #10 in
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDistributeListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast distribute-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " distribute-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " distribute-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Access-list to filter route updates to/from this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDistributeList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast distribute-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: As-path-list to filter outgoing route updates to this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy as-path-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" ";"as-path-list $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 filter-list #10 out
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 filter-list #10 out
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastFilterListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast filter-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " filter-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " filter-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: As-path-list to filter incoming route updates from this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy as-path-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" ";"as-path-list $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 filter-list #10 in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 filter-list #10 in
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastFilterListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast filter-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " filter-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " filter-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: As-path-list to filter route updates to/from this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastFilterList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast filter-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Maximum number of prefixes to accept from this peer-group
	val_help: u32:1-4294967295;  Prefix limit
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967295; "maximum-prefix must be between 1 and 4294967295"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 maximum-prefix #9
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 maximum-prefix #9
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastMaximumPrefix(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast maximum-prefix WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " maximum-prefix ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " maximum-prefix ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Nexthop attributes

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 nexthop-local unchanged
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 nexthop-local unchanged
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastNexthopLocal(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast nexthop-local
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " nexthop-local unchanged"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " nexthop-local unchanged"))
	}
	return cmd.Success
}

/*
	help: Leave link-local nexthop unchanged for this peer
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastNexthopLocalUnchanged(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast nexthop-local unchanged
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Nexthop for routes sent to this peer-group to be the local router

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 next-hop-self
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 next-hop-self
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastNexthopSelf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast nexthop-self
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " next-hop-self"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " next-hop-self"))
	}
	return cmd.Success
}

/*
	help: BGP peer-group IPv6 parameters
	delete:expression: exec "/opt/vyatta/sbin/vyatta-bgp.pl \
	                           --check-peer-groups-6 --peergroup $VAR(../../@) --as $VAR(../../../@)"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 activate
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 activate
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6Unicast(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " activate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " activate"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Prefix-list to filter outgoing route updates to this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy prefix-list6 )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list6 $VAR(@)\" "; "prefix-list6 $VAR(@) doesn't exist"
	commit:expression: $VAR(../../distribute-list/export/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 prefix-list #10 out
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 prefix-list #10 out
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastPrefixListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast prefix-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " prefix-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " prefix-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Prefix-list to filter incoming route updates from this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy prefix-list6 )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list6 $VAR(@)\" "; "prefix-list6 $VAR(@) doesn't exist"
	commit:expression: $VAR(../../distribute-list/import/) == ""; "you can't set both a prefix-list and a distribute list"


SET: router bgp #3 ; address-family ipv6 ; neighbor #5 prefix-list #10 in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 prefix-list #10 in
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastPrefixListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast prefix-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " prefix-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " prefix-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Prefix-list to filter route updates to/from this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastPrefixList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast prefix-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Remove private AS numbers from AS path in outbound route updates

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 remove-private-AS
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 remove-private-AS
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRemovePrivateAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast remove-private-as
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " remove-private-AS"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " remove-private-AS"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to filter outgoing route updates to this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 route-map #10 out
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 route-map #10 out
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteMapExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast route-map export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " route-map ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " route-map ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to filter incoming route updates from this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 route-map #10 in
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 route-map #10 in
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteMapImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast route-map import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " route-map ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " route-map ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Route-map to filter route updates to/from this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast route-map
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Peer-group as a route reflector client
	commit:expression: $VAR(../../../../@) == $VAR(../../../remote-as/@); "remote-as must equal local-as"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 route-reflector-client
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 route-reflector-client
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteReflectorClient(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast route-reflector-client
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " route-reflector-client"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " route-reflector-client"))
	}
	return cmd.Success
}

/*
	help: Peer-group as route server client

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 route-server-client
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 route-server-client
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteServerClient(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast route-server-client
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " route-server-client"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " route-server-client"))
	}
	return cmd.Success
}

/*
	help: Inbound soft reconfiguration for this peer-group [REQUIRED]

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 soft-reconfiguration inbound
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 soft-reconfiguration inbound
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastSoftReconfigurationInbound(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast soft-reconfiguration inbound
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " soft-reconfiguration inbound"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " soft-reconfiguration inbound"))
	}
	return cmd.Success
}

/*
	help: Soft reconfiguration for peer-group
	commit:expression: $VAR(./inbound/) != ""; "you must specify the type of soft-reconfiguration"

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastSoftReconfiguration(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast soft-reconfiguration
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to selectively unsuppress suppressed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; address-family ipv6 ; neighbor #5 unsuppress-map #9
DEL: router bgp #3 ; address-family ipv6 ; no neighbor #5 unsuppress-map #9
*/
func quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastUnsuppressMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family ipv6-unicast unsuppress-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("neighbor ", Args[1], " unsuppress-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"address-family ipv6",
			fmt.Sprint("no neighbor ", Args[1], " unsuppress-map ", Args[2]))
	}
	return cmd.Success
}

/*
	help: BGP peer-group address-family parameters

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupAddressFamily(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD address-family
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Accept a route that contains the local-AS in the as-path

SET: router bgp #3 ; neighbor #5 allowas-in
DEL: router bgp #3 ; no neighbor #5 allowas-in
*/
func quaggaProtocolsBgpPeerGroupAllowasIn(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD allowas-in
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " allowas-in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " allowas-in"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Number of occurrences of AS number
	val_help: u32:1-10; Number of times AS is allowed in path
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 10; "allowas-in number must be between 1 and 10"

SET: router bgp #3 ; neighbor #5 allowas-in #8
DEL: router bgp #3 ; no neighbor #5 allowas-in ; neighbor #5 allowas-in
*/
func quaggaProtocolsBgpPeerGroupAllowasInNumber(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD allowas-in number WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " allowas-in ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " allowas-in"),
			fmt.Sprint("neighbor ", Args[1], " allowas-in"))
	}
	return cmd.Success
}

/*
	help: Send AS path unchanged
*/
func quaggaProtocolsBgpPeerGroupAttributeUnchangedAsPath(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD attribute-unchanged as-path
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, peerGrp, ok := quaggaBgpPeerGroupLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !peerGrp.ipv4AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		peerGrp.ipv4AttributeUnchangedAsPath = true
	case cmd.Delete:
		peerGrp.ipv4AttributeUnchangedAsPath = false
	}
	asPath := ""
	if peerGrp.ipv4AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if peerGrp.ipv4AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if peerGrp.ipv4AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send multi-exit discriminator unchanged
*/
func quaggaProtocolsBgpPeerGroupAttributeUnchangedMed(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD attribute-unchanged med
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, peerGrp, ok := quaggaBgpPeerGroupLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !peerGrp.ipv4AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		peerGrp.ipv4AttributeUnchangedMed = true
	case cmd.Delete:
		peerGrp.ipv4AttributeUnchangedMed = false
	}
	asPath := ""
	if peerGrp.ipv4AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if peerGrp.ipv4AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if peerGrp.ipv4AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: Send nexthop unchanged
*/
func quaggaProtocolsBgpPeerGroupAttributeUnchangedNextHop(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD attribute-unchanged next-hop
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, peerGrp, ok := quaggaBgpPeerGroupLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	if !peerGrp.ipv4AttributeUnchanged {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		peerGrp.ipv4AttributeUnchangedNextHop = true
	case cmd.Delete:
		peerGrp.ipv4AttributeUnchangedNextHop = false
	}
	asPath := ""
	if peerGrp.ipv4AttributeUnchangedAsPath {
		asPath = " as-path"
	}
	med := ""
	if peerGrp.ipv4AttributeUnchangedMed {
		med = " med"
	}
	nextHop := ""
	if peerGrp.ipv4AttributeUnchangedNextHop {
		nextHop = " next-hop"
	}
	quaggaVtysh("configure terminal",
		fmt.Sprint("router bgp ", Args[0]),
		fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"),
		fmt.Sprint("neighbor ", Args[1], " attribute-unchanged", asPath, med, nextHop))
	return cmd.Success
}

/*
	help: BGP attributes are sent unchanged

SET: router bgp #3 ; no neighbor #5 attribute-unchanged ; neighbor #5 attribute-unchanged ?as-path ?med ?next-hop
DEL: router bgp #3 ; no neighbor #5 attribute-unchanged ?as-path ?med ?next-hop
*/
func quaggaProtocolsBgpPeerGroupAttributeUnchanged(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD attribute-unchanged
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	_, peerGrp, ok := quaggaBgpPeerGroupLookup(Args[1])
	if !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		peerGrp.ipv4AttributeUnchanged = true
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " attribute-unchanged"))
	case cmd.Delete:
		peerGrp.ipv4AttributeUnchanged = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " attribute-unchanged"))

	}
	return cmd.Success
}

/*
	help: Advertise dynamic capability to this peer-group

SET: router bgp #3 ; neighbor #5 capability dynamic
DEL: router bgp #3 ; no neighbor #5 capability dynamic
*/
func quaggaProtocolsBgpPeerGroupCapabilityDynamic(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD capability dynamic
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " capability dynamic"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " capability dynamic"))
	}
	return cmd.Success
}

/*
	help: Advertise capabilities to this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupCapability(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD capability
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise ORF capability to this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupCapabilityOrf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD capability orf
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise prefix-list ORF capability to this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupCapabilityOrfPrefixList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD capability orf prefix-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Capability to receive the ORF

SET: router bgp #3 ; neighbor #5 capability orf prefix-list receive
DEL: router bgp #3 ; no neighbor #5 capability orf prefix-list receive
*/
func quaggaProtocolsBgpPeerGroupCapabilityOrfPrefixListReceive(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD capability orf prefix-list receive
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " capability orf prefix-list receive"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " capability orf prefix-list receive"))
	}
	return cmd.Success
}

/*
	help: Capability to send the ORF

SET: router bgp #3 ; neighbor #5 capability orf prefix-list send
DEL: router bgp #3 ; no neighbor #5 capability orf prefix-list send
*/
func quaggaProtocolsBgpPeerGroupCapabilityOrfPrefixListSend(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD capability orf prefix-list send
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " capability orf prefix-list send"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " capability orf prefix-list send"))
	}
	return cmd.Success
}

/*
	help: Send default route to this peer-group

SET: router bgp #3 ; neighbor #5 activate ; neighbor #5 default-originate
DEL: router bgp #3 ; no neighbor #5 default-originate
*/
func quaggaProtocolsBgpPeerGroupDefaultOriginate(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD default-originate
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " activate"),
			fmt.Sprint("neighbor ", Args[1], " default-originate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " default-originate"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to specify criteria of the default
	allowed: local -a params
	        params=$(/opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map)
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" " ; "route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 activate ; neighbor #5 default-originate route-map #8
DEL: router bgp #3 ; no neighbor #5 default-originate route-map #8
*/
func quaggaProtocolsBgpPeerGroupDefaultOriginateRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD default-originate route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " activate"),
			fmt.Sprint("neighbor ", Args[1], " default-originate route-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " default-originate route-map ", Args[2]))
	}
	return cmd.Success
}

/*
	type: txt
	help: Description for this peer-group
*/
func quaggaProtocolsBgpPeerGroupDescription(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD description WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Disable capability negotiation with this peer-group

SET: router bgp #3 ; neighbor #5 dont-capability-negotiate
DEL: router bgp #3 ; no neighbor #5 dont-capability-negotiate
*/
func quaggaProtocolsBgpPeerGroupDisableCapabilityNegotiation(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD disable-capability-negotiation
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " dont-capability-negotiate"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " dont-capability-negotiate"))
	}
	return cmd.Success
}

/*
	help: Disable check to see if EBGP peer's address is a connected route

SET: router bgp #3 ; neighbor #5 disable-connected-check
DEL: router bgp #3 ; no neighbor #5 disable-connected-check
*/
func quaggaProtocolsBgpPeerGroupDisableConnectedCheck(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD disable-connected-check
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " disable-connected-check"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " disable-connected-check"))
	}
	return cmd.Success
}

/*
	help: Disable sending extended community attributes to this peer-group

SET: router bgp #3 ; no neighbor #5 send-community extended
DEL: router bgp #3 ; neighbor #5 send-community extended
*/
func quaggaProtocolsBgpPeerGroupDisableSendCommunityExtended(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD disable-send-community extended
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " send-community extended"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " send-community extended"))
	}
	return cmd.Success
}

/*
	help: Disable sending community attributes to this peer-group
	commit:expression: ($VAR(./extended/) != "") || ($VAR(./standard/) != ""); "you must specify the type of community"

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupDisableSendCommunity(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD disable-send-community
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Disable sending standard community attributes to this peer-group

SET: router bgp #3 ; no neighbor #5 send-community standard
DEL: router bgp #3 ; neighbor #5 send-community standard
*/
func quaggaProtocolsBgpPeerGroupDisableSendCommunityStandard(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD disable-send-community standard
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " send-community standard"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " send-community standard"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Access-list to filter outgoing route updates to this peer-group
	val_help: u32:1-65535; Access list number
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Access list must be between 1 and 65535"

	allowed: local -a params
	 	params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy access-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list $VAR(@)\" "; "access-list $VAR(@) doesn't exist"
	commit:expression: $VAR(../../prefix-list/export/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; neighbor #5 distribute-list #8 out
DEL: router bgp #3 ; no neighbor #5 distribute-list #8 out
*/
func quaggaProtocolsBgpPeerGroupDistributeListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD distribute-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " distribute-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " distribute-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Access-list to filter incoming route updates from this peer-group
	val_help: u32:1-65535; Access list number
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Access list must be between 1 and 65535"

	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy access-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy access-list $VAR(@)\" "; "access-list $VAR(@) doesn't exist"
	commit:expression: $VAR(../../prefix-list/import/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; neighbor #5 distribute-list #8 in
DEL: router bgp #3 ; no neighbor #5 distribute-list #8 in
*/
func quaggaProtocolsBgpPeerGroupDistributeListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD distribute-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " distribute-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " distribute-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Access-list to filter route updates to/from this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupDistributeList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD distribute-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Allow this EBGP peer-group to not be on a directly connected network
	val_help: u32:1-255; Number of hops

	syntax:expression: $VAR(@) >=1 && $VAR(@) <= 255; "ebgp-multihop must be between 1 and 255"
	commit:expression: $VAR(../ttl-security/hops/) == ""; "you can't set both ebgp-multihop and ttl-security hops"

SET: router bgp #3 ; neighbor #5 ebgp-multihop #7
DEL: router bgp #3 ; no neighbor #5 ebgp-multihop #7
*/
func quaggaProtocolsBgpPeerGroupEbgpMultihop(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD ebgp-multihop WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " ebgp-multihop ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " ebgp-multihop ", Args[2]))
	}
	return cmd.Success
}

/*
	type: txt
	help: As-path-list to filter outgoing route updates to this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy as-path-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" ";"as-path-list $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 filter-list #8 out
DEL: router bgp #3 ; no neighbor #5 filter-list #8 out
*/
func quaggaProtocolsBgpPeerGroupFilterListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD filter-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " filter-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " filter-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: As-path-list to filter incoming route updates from this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy as-path-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy as-path-list $VAR(@)\" ";"as-path-list $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 filter-list #8 in
DEL: router bgp #3 ; no neighbor #5 filter-list #8 in
*/
func quaggaProtocolsBgpPeerGroupFilterListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD filter-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " filter-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " filter-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: As-path-list to filter route updates to/from this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupFilterList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD filter-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: Local AS number [REQUIRED]
	val_help: u32: 1-4294967294; Local AS number
	syntax:expression: $VAR(@) >=1 && $VAR(@) <= 4294967294; "local-as must be between 1 and 4294967294"
	commit:expression: $VAR(@) != $VAR(../../@); "you can't set local-as the same as the router AS"

SET: router bgp #3 ; no neighbor #5 local-as ; neighbor #5 local-as #7
DEL: router bgp #3 ; no neighbor #5 local-as #7
*/
func quaggaProtocolsBgpPeerGroupLocalAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD local-as WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " local-as"),
			fmt.Sprint("neighbor ", Args[1], " local-as ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " local-as ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Disable prepending local-as to updates from EBGP peers

SET: router bgp #3 ; no neighbor #5 local-as #7 ; neighbor #5 local-as #7i no-prepend
DEL: router bgp #3 ; no neighbor #5 local-as #7 no-prepend ; neighbor #5 local-as #7
*/
func quaggaProtocolsBgpPeerGroupLocalAsNoPrepend(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD local-as WORD no-prepend
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " local-as ", Args[2]),
			fmt.Sprint("neighbor ", Args[1], " local-as ", Args[2], " no-prepend"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " local-as ", Args[2], " no-prepend"),
			fmt.Sprint("neighbor ", Args[1], " local-as ", Args[2]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Maximum number of prefixes to accept from this peer-group
	val_help: u32:1-4294967295;  Prefix limit
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967295; "maximum-prefix must be between 1 and 4294967295"

SET: router bgp #3 ; neighbor #5 maximum-prefix #7
DEL: router bgp #3 ; no neighbor #5 maximum-prefix #7
*/
func quaggaProtocolsBgpPeerGroupMaximumPrefix(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD maximum-prefix WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " maximum-prefix ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " maximum-prefix ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Nexthop for routes sent to this peer-group to be the local router

SET: router bgp #3 ; neighbor #5 next-hop-self
DEL: router bgp #3 ; no neighbor #5 next-hop-self
*/
func quaggaProtocolsBgpPeerGroupNexthopSelf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD nexthop-self
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " next-hop-self"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " next-hop-self"))
	}
	return cmd.Success
}

/*
	help: Ignore capability negotiation with specified peer-group
	commit:expression: $VAR(../strict-capability/) == ""; "you can't set both strict-capability and override-capability"

SET: router bgp #3 ; neighbor #5 override-capability
DEL: router bgp #3 ; no neighbor #5 override-capability
*/
func quaggaProtocolsBgpPeerGroupOverrideCapability(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD override-capability
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " override-capability"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " override-capability"))
	}
	return cmd.Success
}

/*
	help: Don not intiate a session with this peer-group

SET: router bgp #3 ; neighbor #5 passive
DEL: router bgp #3 ; no neighbor #5 passive
*/
func quaggaProtocolsBgpPeerGroupPassive(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD passive
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " passive"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " passive"))
	}
	return cmd.Success
}

/*
	type: txt
	help: BGP MD5 password
	syntax:expression: exec "			      \
	        if [ `echo -n '$VAR(@)' | wc -c` -gt 80 ]; then   \
	          echo Password must be 80 characters or less;\
	          exit 1 ;                                     \
	        fi ; "


SET: router bgp #3 ; neighbor #5 password #7
DEL: router bgp #3 ; no neighbor #5 password
*/
func quaggaProtocolsBgpPeerGroupPassword(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD password WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " password ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " password"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Prefix-list to filter outgoing route updates to this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy prefix-list )
	        echo -n ${params[@]##* /}

	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list $VAR(@)\" "; "prefix-list $VAR(@) doesn't exist"
	commit:expression: $VAR(../../distribute-list/export/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; neighbor #5 prefix-list #8 out
DEL: router bgp #3 ; no neighbor #5 prefix-list #8 out
*/
func quaggaProtocolsBgpPeerGroupPrefixListExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD prefix-list export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " prefix-list ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " prefix-list ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Prefix-list to filter incoming route updates from this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy prefix-list )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy prefix-list $VAR(@)\" "; "prefix-list $VAR(@) doesn't exist"
	commit:expression: $VAR(../../distribute-list/import/) == ""; "you can't set both a prefix-list and a distribute list"

SET: router bgp #3 ; neighbor #5 prefix-list #8 in
DEL: router bgp #3 ; no neighbor #5 prefix-list #8 in
*/
func quaggaProtocolsBgpPeerGroupPrefixListImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD prefix-list import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " prefix-list ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " prefix-list ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Prefix-list to filter route updates to/from this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupPrefixList(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD prefix-list
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Peer-group BGP AS number [REQUIRED]
	val_help: u32:1-4294967294; AS number
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967294; \
	                   "remote-as must be between 1 and 4294967294"

SET: router bgp #3 ; neighbor #5 peer-group ; neighbor #5 remote-as #7
DEL: router bgp #3 ; no neighbor #5 remote-as #7
*/
func quaggaProtocolsBgpPeerGroupRemoteAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD remote-as WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " peer-group"),
			fmt.Sprint("neighbor ", Args[1], " remote-as ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " remote-as ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Remove private AS numbers from AS path in outbound route updates

SET: router bgp #3 ; neighbor #5 remove-private-AS
DEL: router bgp #3 ; no neighbor #5 remove-private-AS
*/
func quaggaProtocolsBgpPeerGroupRemovePrivateAs(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD remove-private-as
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " remove-private-AS"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " remove-private-AS"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to filter outgoing route updates to this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 route-map #8 out
DEL: router bgp #3 ; no neighbor #5 route-map #8 out
*/
func quaggaProtocolsBgpPeerGroupRouteMapExport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD route-map export WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " route-map ", Args[2], " out"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " route-map ", Args[2], " out"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to filter incoming route updates from this peer-group
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 route-map #8 in
DEL: router bgp #3 ; no neighbor #5 route-map #8 in
*/
func quaggaProtocolsBgpPeerGroupRouteMapImport(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD route-map import WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " route-map ", Args[2], " in"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " route-map ", Args[2], " in"))
	}
	return cmd.Success
}

/*
	help: Route-map to filter route updates to/from this peer-group

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD route-map
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Peer-group as a route reflector client
	commit:expression: $VAR(../../@) == $VAR(../remote-as/@); "remote-as must equal local-as"

SET: router bgp #3 ; neighbor #5 route-reflector-client
DEL: router bgp #3 ; no neighbor #5 route-reflector-client
*/
func quaggaProtocolsBgpPeerGroupRouteReflectorClient(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD route-reflector-client
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " route-reflector-client"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " route-reflector-client"))
	}
	return cmd.Success
}

/*
	help: Peer-group as route server client

SET: router bgp #3 ; neighbor #5 route-server-client
DEL: router bgp #3 ; no neighbor #5 route-server-client
*/
func quaggaProtocolsBgpPeerGroupRouteServerClient(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD route-server-client
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " route-server-client"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " route-server-client"))
	}
	return cmd.Success
}

/*
	help: Administratively shut down peer-group

SET: router bgp #3 ; neighbor #5 shutdown
DEL: router bgp #3 ; no neighbor #5 shutdown
*/
func quaggaProtocolsBgpPeerGroupShutdown(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD shutdown
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " shutdown"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " shutdown"))
	}
	return cmd.Success
}

/*
	help: Inbound soft reconfiguration for this peer-group [REQUIRED]

SET: router bgp #3 ; neighbor #5 soft-reconfiguration inbound
DEL: router bgp #3 ; no neighbor #5 soft-reconfiguration inbound
*/
func quaggaProtocolsBgpPeerGroupSoftReconfigurationInbound(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD soft-reconfiguration inbound
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " soft-reconfiguration inbound"))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " soft-reconfiguration inbound"))
	}
	return cmd.Success
}

/*
	help: Soft reconfiguration for peer-group
	commit:expression: $VAR(./inbound/) != ""; "you must specify the type of soft-reconfiguration"

SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupSoftReconfiguration(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD soft-reconfiguration
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Number of the maximum number of hops to the BGP peer
	val_help: u32:1-254; Number of hops

	syntax:expression: $VAR(@) >=1 && $VAR(@) <= 254; "ttl-security hops must be between 1 and 254"
	commit:expression: $VAR(../../ebgp-multihop/) == ""; "you can't set both ebgp-multihop and ttl-security hops"

SET: router bgp #3 ; neighbor #5 ttl-security hops #8
DEL: router bgp #3 ; no neighbor #5 ttl-security hops #8
*/
func quaggaProtocolsBgpPeerGroupTtlSecurityHops(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD ttl-security hops WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " ttl-security hops ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " ttl-security hops ", Args[2]))
	}
	return cmd.Success
}

/*
	help: Ttl security mechanism


SET:
DEL:
*/
func quaggaProtocolsBgpPeerGroupTtlSecurity(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD ttl-security
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route-map to selectively unsuppress suppressed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; neighbor #5 unsuppress-map #7
DEL: router bgp #3 ; no neighbor #5 unsuppress-map #7
*/
func quaggaProtocolsBgpPeerGroupUnsuppressMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD unsuppress-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " unsuppress-map ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " unsuppress-map ", Args[2]))
	}
	return cmd.Success
}

/*
	type: txt
	help: Source IP of routing updates
	val_help: ipv4;IP address of route source
	val_help: <interface>; Interface as route source

	commit:expression: exec "/opt/vyatta/sbin/vyatta-bgp.pl --check-source $VAR(@)"

SET: router bgp #3 ; neighbor #5 update-source #7
DEL: router bgp #3 ; no neighbor #5 update-source
*/
func quaggaProtocolsBgpPeerGroupUpdateSource(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD update-source WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " update-source ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " update-source"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Default weight for routes from this peer-group
	val_help: u32:1-65535; Weight for routes from this peer-group
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "weight must be between 1 and 65535"

SET: router bgp #3 ; neighbor #5 weight #7
DEL: router bgp #3 ; no neighbor #5 weight #7
*/
func quaggaProtocolsBgpPeerGroupWeight(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD peer-group WORD weight WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if _, _, ok := quaggaBgpPeerGroupLookup(Args[1]); !ok {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("neighbor ", Args[1], " weight ", Args[2]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no neighbor ", Args[1], " weight ", Args[2]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes

SET: router bgp #3 ; redistribute connected metric #7
DEL: router bgp #3 ; no redistribute connected metric #7
*/
func quaggaProtocolsBgpRedistributeConnectedMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute connected metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute connected metric ", Args[1]))

	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute connected metric ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute connected routes into BGP

SET: router bgp #3 ; redistribute connected ?route-map ?metric
DEL: router bgp #3 ; no redistribute connected
*/
func quaggaProtocolsBgpRedistributeConnected(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute connected
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"redistribute connected")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no redistribute connected")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

SET: router bgp #3 ; redistribute connected route-map #7
DEL: router bgp #3 ; no redistribute connected route-map #7
*/
func quaggaProtocolsBgpRedistributeConnectedRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute connected route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute connected route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute connected route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
*/
func quaggaProtocolsBgpRedistributeKernelMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute kernel metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute kernel metric ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute kernel metric ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute kernel routes into BGP

SET: router bgp #3 ; no redistribute kernel ; redistribute kernel ?route-map ?metric
DEL: router bgp #3 ; no redistribute kernel
*/
func quaggaProtocolsBgpRedistributeKernel(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute kernel
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"redistribute kernel")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no redistribute kernel")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsBgpRedistributeKernelRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute kernel route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute kernel route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute kernel route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute routes from other protocols into BGP

SET:
DEL:
*/
func quaggaProtocolsBgpRedistribute(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
*/
func quaggaProtocolsBgpRedistributeOspfMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute ospf metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute ospf metric ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute ospf metric ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute OSPF routes into BGP

SET: router bgp #3 ; no redistribute ospf ; redistribute ospf ?route-map ?metric
DEL: router bgp #3 ; no redistribute ospf
*/
func quaggaProtocolsBgpRedistributeOspf(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute ospf
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"redistribute ospf")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no redistribute ospf")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsBgpRedistributeOspfRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute ospf route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute ospf route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute ospf route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
*/
func quaggaProtocolsBgpRedistributeRipMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute rip metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute rip route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute rip route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute RIP routes into BGP

SET: router bgp #3 ; no redistribute rip ; redistribute rip ?route-map ?metric
DEL: router bgp #3 ; no redistribute rip
*/
func quaggaProtocolsBgpRedistributeRip(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute rip
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"redistribute rip")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no redistribute rip")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsBgpRedistributeRipRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute rip route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute rip route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute rip route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
*/
func quaggaProtocolsBgpRedistributeStaticMetric(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute static metric WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute static metric ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute static metric ", Args[1]))
	}
	return cmd.Success
}

/*
	help: Redistribute static routes into BGP

SET: router bgp #3 ; no redistribute static ; redistribute static ?route-map ?metric
DEL: router bgp #3 ; no redistribute static
*/
func quaggaProtocolsBgpRedistributeStatic(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute static
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"redistribute static")
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no redistribute static")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map to filter redistributed routes
	allowed: local -a params
	        params=$( /opt/vyatta/sbin/vyatta-policy.pl --list-policy route-map )
	        echo -n ${params[@]##* /}
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsBgpRedistributeStaticRouteMap(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD redistribute static route-map WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("redistribute static route-map ", Args[1]))
	case cmd.Delete:
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("no redistribute static route-map ", Args[1]))
	}
	return cmd.Success
}

/*
	type: u32
	help: BGP holdtime interval
	val_help: u32:4-65535; Hold-time in seconds (default 180)
	val_help: 0; Don't hold routes

	default: 180
	syntax:expression: $VAR(@) == 0 || ($VAR(@) >= 4 && $VAR(@) <= 65535); \
	       "hold-time interval must be 0 or between 4 and 65535"
*/
func quaggaProtocolsBgpTimersHoldtime(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD timers holdtime WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.timers {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.timersHoldtime = fmt.Sprint(Args[2])
	case cmd.Delete:
		bgpConfigState.timersHoldtime = ""
	}
	if bgpConfigState.timersKeepalive != "" && bgpConfigState.timersHoldtime != "" {
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("timers bgp ", bgpConfigState.timersKeepalive, " ", bgpConfigState.timersHoldtime))
	} else {
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no timers bgp")
	}
	return cmd.Success
}

/*
	type: u32
	default: 60
	help: Keepalive interval
	val_help: u32:1-65535; Keep-alive time in seconds (default 60)
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; \
	       "Keepalive interval must be between 1 and 65535"
*/
func quaggaProtocolsBgpTimersKeepalive(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD timers keepalive WORD
	if bgpConfigState == nil {
		return cmd.Success
	}
	if bgpConfigState == nil || !bgpConfigState.timers {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.timersKeepalive = fmt.Sprint(Args[2])
	case cmd.Delete:
		bgpConfigState.timersKeepalive = ""
	}
	if bgpConfigState.timersKeepalive != "" && bgpConfigState.timersHoldtime != "" {
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			fmt.Sprint("timers bgp ", bgpConfigState.timersKeepalive, " ", bgpConfigState.timersHoldtime))
	} else {
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no timers bgp")
	}
	return cmd.Success
}

/*
	help: BGP protocol timers
	commit:expression: $VAR(./keepalive/) != ""; "you must set a keepalive interval"
	commit:expression: $VAR(./holdtime/) != ""; "you must set a holdtime interval"

SET: router bgp #3 ; timers bgp @keepalive @holdtime
DEL: router bgp #3 ; no timers bgp
*/
func quaggaProtocolsBgpTimers(Cmd int, Args cmd.Args) int {
	//protocols bgp WORD timers
	if bgpConfigState == nil {
		return cmd.Success
	}
	// XXX
	if bgpConfigState == nil {
		return cmd.Success
	}
	switch Cmd {
	case cmd.Set:
		bgpConfigState.timers = true
	case cmd.Delete:
		bgpConfigState.timers = false
		quaggaVtysh("configure terminal",
			fmt.Sprint("router bgp ", Args[0]),
			"no timers bgp")
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	commit:expression: $VAR(./export/) != ""; "must add protocol to filter"
	help: Access list to filter networks in routing updates
*/
func quaggaProtocolsOspfAccessList(Cmd int, Args cmd.Args) int {
	//protocols ospf access-list WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	multi:
	type: txt
	help: Filter for outgoing routing updates [REQUIRED]
	syntax:expression: $VAR(@) in "bgp", "connected", "kernel", "rip", "static"; "Must be (bgp, connected, kernel, rip, or static)"

	val_help: bgp                Filter bgp routes;
	val_help: connected          Filter connected routes;
	val_help: kernel             Filter kernel routes;
	val_help: rip                Filter rip routes;
	val_help: static             Filter static routes;

	create: vtysh -c "configure terminal" \
	            -c "router ospf"                                      \
	            -c "distribute-list $VAR(../@) out $VAR(@)";

	delete: vtysh -c "configure terminal" \
	            -c "router ospf"                                      \
	            -c "no distribute-list $VAR(../@) out $VAR(@)";

*/
func quaggaProtocolsOspfAccessListExport(Cmd int, Args cmd.Args) int {
	//protocols ospf access-list WORD export WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("distribute-list ", Args[0], " out ", Args[1]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "access-list", fmt.Sprint(Args[0]),
				"export", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no distribute-list ", Args[0], " out ", Args[1]))
		}
	}
	return cmd.Success
}

/*
	tag:
	type: txt
	help: OSPF Area
	syntax:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --check-ospf-area $VAR(@)"
	val_help: u32; OSPF area in decimal notation
	val_help: ipv4; OSPF area in dotted decimal notation
*/
func quaggaProtocolsOspfArea(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Area type
	val_help: normal; Normal area type
	val_help: nssa; Not so stubby area type
	val_help: stub; Stub Area type
*/
func quaggaProtocolsOspfAreaAreaType(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Normal OSPF area
	syntax:expression: $VAR(../stub/) == "" ; "Must delete stub area type first"
	syntax:expression: $VAR(../nssa/) == "" ; "Must delete nssa area type first"
	create:expression: "                                                            \
	      if [ x$VAR(../../@) != x0.0.0.0 ] && [ x$VAR(../../@) != x0 ]; then       \
	         vtysh -c \"configure terminal\"            \
	           -c \"router ospf\"                                                   \
	           -c \"no area $VAR(../../@) stub\" -c \"no area $VAR(../../@) nssa\"; \
	      fi; "
*/
func quaggaProtocolsOspfAreaAreaTypeNormal(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type normal
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no area ", Args[0], " stub"),
			fmt.Sprint("no area ", Args[0], " nssa"))
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Summary-default cost of nssa area
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 16777215; "Cost must be between 0-16777215"
	val_help: u32:0-16777215; Summary default cost

	create: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "area $VAR(../../../@) nssa"                         \
	          -c "area $VAR(../../../@) default-cost $VAR(@)";

	update: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "area $VAR(../../../@) default-cost $VAR(@)";

	delete: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "no area $VAR(../../../@) default-cost $VAR(@)";
*/
func quaggaProtocolsOspfAreaAreaTypeNssaDefaultCost(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type nssa default-cost WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " nssa"),
			fmt.Sprint("area ", Args[0], " default-cost ", Args[1]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"default-cost", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " default-cost ", Args[1]))
		}
	}
	return cmd.Success
}

/*
	help: Do not inject inter-area routes into stub
*/
func quaggaProtocolsOspfAreaAreaTypeNssaNoSummary(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type nssa no-summary
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	translate := configCandidate.value(
		[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]), "area-type", "nssa", "translate"})
	noSummary := configCandidate.lookup(
		[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]), "area-type", "nssa", "no-summary"})
	if translate != nil {
		parm += " translate-" + *translate
	}
	if noSummary != nil {
		parm += " no-summary"
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no area ", Args[0], " nssa"),
			fmt.Sprint("area ", Args[0], " nssa", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no area ", Args[0], " nssa"))
	}
	return cmd.Success
}

/*
	help: Nssa OSPF area
	syntax:expression: ! $VAR(../../@) in "0", "0.0.0.0"; "Backbone can't be NSSA"
	syntax:expression: $VAR(../normal/) == "" ; "Must delete normal area type first"
	syntax:expression: $VAR(../stub/) == "" ; "Must delete stub area type first"

	delete: touch /tmp/ospf-area-nssa.$PPID
	end: if [ -f "/tmp/ospf-area-nssa.$PPID" ]; then
	        vtysh -c "configure terminal" \
	          -c "router ospf" -c "no area $VAR(../../@) nssa";
		rm /tmp/ospf-area-nssa.$PPID;
	     else
	        if [ -n "$VAR(translate/@)" ]; then
	           PARM="translate-$VAR(translate/@)";
	        fi;
	        # using workaround pending bug 2525
	        #
	        # if [ -n "$VAR(no-summary/)" ]; then
	        #   PARM="$PARM no-summary";
	        # fi;
		${vyatta_sbindir}/vyatta-check-typeless-node.pl           \
	          "protocols ospf area $VAR(../../@) area-type nssa no-summary";
	        if [ $? -eq 0 ] ; then
	           PARM="$PARM no-summary";
	        fi;
	        vtysh -c "configure terminal" \
	          -c "router ospf" -c "area $VAR(../../@) nssa $PARM";
	     fi;
*/
func quaggaProtocolsOspfAreaAreaTypeNssa(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type nssa
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]), "area-type", "nssa"}) != nil {
		}
	}
	parm := ""
	translate := configCandidate.value(
		[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]), "area-type", "nssa", "translate"})
	noSummary := configCandidate.lookup(
		[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]), "area-type", "nssa", "no-summary"})
	if translate != nil {
		parm += " translate-" + *translate
	}
	if noSummary != nil {
		parm += " no-summary"
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no area ", Args[0], " nssa"),
			fmt.Sprint("area ", Args[0], " nssa", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no area ", Args[0], " nssa"))
	}
	return cmd.Success
}

/*
	type: txt
	help: Nssa-abr
	default: "candidate"
	syntax:expression: $VAR(@) in "always", "candidate", "never"; "Must be (always, candidate, or never)"

	val_help: always; NSSA-ABR to always translate
	val_help: candidate; NSSA-ABR for translate election (default)
	val_help: never; NSSA-ABR to never translate
*/
func quaggaProtocolsOspfAreaAreaTypeNssaTranslate(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type nssa translate WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	translate := configCandidate.value(
		[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]), "area-type", "nssa", "translate"})
	noSummary := configCandidate.lookup(
		[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]), "area-type", "nssa", "no-summary"})
	if translate != nil {
		parm += " translate-" + *translate
	}
	if noSummary != nil {
		parm += " no-summary"
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no area ", Args[0], " nssa"),
			fmt.Sprint("area ", Args[0], " nssa", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no area ", Args[0], " nssa"))
	}
	return cmd.Success
}

/*
	type: u32
	help: Summary-default cost of stub area
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 16777215; "Cost must be between 0-16777215"
	val_help: u32:0-16777215; Summary default cost of stub area

	create: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "area $VAR(../../../@) stub"                         \
	          -c "area $VAR(../../../@) default-cost $VAR(@)";

	update: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "area $VAR(../../../@) default-cost $VAR(@)";

	delete: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "no area $VAR(../../../@) default-cost $VAR(@)";
*/
func quaggaProtocolsOspfAreaAreaTypeStubDefaultCost(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type stub default-cost WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " stub"),
			fmt.Sprint("area ", Args[0], " default-cost ", Args[1]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"area-type", "stub", "default-cost", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " default-cost ", Args[1]))
		}
	}
	return cmd.Success
}

/*
	help: Do not inject inter-area routes into stub

	create:
		vtysh -c "configure terminal" \
		    -c "router ospf" \
		    -c "area $VAR(../../../@) stub no-summary "

	delete:
		vtysh -c "configure terminal" \
	            -c "router ospf" \
	            -c "no area $VAR(../../../@) stub no-summary "
*/
func quaggaProtocolsOspfAreaAreaTypeStubNoSummary(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type stub no-summary
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " stub no-summary"))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"area-type", "stub", "no-summary"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " stub no-summary"))
		}
	}
	return cmd.Success
}

/*
	help: Stub OSPF area

	syntax:expression: ! $VAR(../../@) in "0", "0.0.0.0"; "Backbone can't be stub"

	syntax:expression: $VAR(../nssa/) == "" ; "Must delete nssa area type first"

	syntax:expression: $VAR(../normal/) == "" ; "Must delete normal area type first"

	create:
		vtysh -c "configure terminal" \
		    -c "router ospf" \
		    -c "area $VAR(../../@) stub"

	delete:
		vtysh -c "configure terminal" \
		    -c "router ospf" \
		    -c "no area $VAR(../../@) stub"
*/
func quaggaProtocolsOspfAreaAreaTypeStub(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD area-type stub
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " stub"))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"area-type", "stub"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " stub"))
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: OSPF area authentication type
	allowed: echo "plaintext-password md5"
	syntax:expression: $VAR(@) in "plaintext-password", "md5"; \
	       "Must be either plaintext-password or md5"
	val_help: plaintext-password; Use plain-text authentication
	val_help: md5; Use md5 authentication

	update:expression: "\
	        if [ x$VAR(@) == xplaintext-password ]; then               \
	           vtysh                       \
	            -c \"configure terminal\"                              \
	            -c \"router ospf \"                                    \
	            -c \"no area $VAR(../@) authentication \"              \
	            -c \"area $VAR(../@) authentication \" ;               \
	         else                                                      \
	           vtysh                       \
	            -c \"configure terminal\"                              \
	            -c \"router ospf \"                                    \
	            -c \"no area $VAR(../@) authentication \"              \
	            -c \"area $VAR(../@) authentication message-digest\" ; \
	         fi; "

	delete:expression: "vtysh -c \"configure terminal\" \
	            -c \"router ospf \"                                                 \
	            -c \"no area $VAR(../@) authentication \" "
*/
func quaggaProtocolsOspfAreaAuthentication(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD authentication WORD
	switch Cmd {
	case cmd.Set:
		authentication := fmt.Sprint(Args[1])
		if authentication == "plaintext-password" {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " authentication"),
				fmt.Sprint("area ", Args[0], " authentication"))
		} else {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " authentication"),
				fmt.Sprint("area ", Args[0], " authentication message-digest"))
		}
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"authentication", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " authentication"))
		}
	}
	return cmd.Success
}

/*
	multi:
	type: ipv4net
	help: OSPF network [REQUIRED]
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"
	create:vtysh -c "configure terminal" \
	       -c "router ospf" -c "network $VAR(@) area $VAR(../@)"
	delete:vtysh -c "configure terminal" \
	       -c "router ospf" -c "no network $VAR(@) area $VAR(../@)"
*/
func quaggaProtocolsOspfAreaNetwork(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD network A.B.C.D/M
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("network ", Args[1], " area ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"network", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no network ", Args[1], " area ", Args[0]))
		}
	}
	return cmd.Success
}

/*
	tag:
	type: ipv4net
	help: Summarize routes matching prefix (border routers only)
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"

	delete: touch /tmp/ospf-range.$PPID

	end: if [ -f /tmp/ospf-range.$PPID ]; then
	        vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "no area $VAR(../@) range $VAR(@)";
	        rm /tmp/ospf-range.$PPID;
	     else
	        ${vyatta_sbindir}/vyatta-check-typeless-node.pl           \
	          "protocols ospf area $VAR(../@) range $VAR(@) not-advertise";
	        if [ $? -eq 0 ] ; then
	           if [ -n "$VAR(cost/@)" ] || [ -n "$VAR(substitute/@)" ]; then
	              echo "Remove 'not-advertise' before setting cost or substitue";
		      exit 1;
	           fi;
	           vtysh --noerror -c "configure terminal" \
	             -c "router ospf"                                               \
	             -c "no area $VAR(../@) range $VAR(@)";
	           vtysh -c "configure terminal" \
	             -c "router ospf"                                        \
	             -c "area $VAR(../@) range $VAR(@) not-advertise";
	        else
	           vtysh --noerror -c "configure terminal" \
	             -c "router ospf"                                               \
	             -c "no area $VAR(../@) range $VAR(@)";
	           vtysh -c "configure terminal" \
	             -c "router ospf"                                        \
	             -c "area $VAR(../@) range $VAR(@)";
	           if [ -n "$VAR(cost/@)" ]; then
	              vtysh -c "configure terminal" \
	                -c "router ospf"                                        \
	                -c "area $VAR(../@) range $VAR(@) cost $VAR(cost/@)";
	           fi;
	           if [ -n "$VAR(substitute/@)" ]; then
	              vtysh -c "configure terminal" \
	                -c "router ospf"                                        \
	                -c "area $VAR(../@) range $VAR(@) substitute $VAR(substitute/@)";
	           fi;
	        fi;
	     fi;
*/
func quaggaProtocolsOspfAreaRange(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD range A.B.C.D/M
	switch Cmd {
	case cmd.Set:
		area := fmt.Sprint(Args[0])
		range_ := fmt.Sprint(Args[1])
		notAdvertise := configCandidate.lookup(
			[]string{"protocols", "ospf", "area", area, "range", range_, "not-advertise"})
		if notAdvertise != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " range ", Args[1]))
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("area ", Args[0], " range ", Args[1], " not-advertise"))
		} else {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " range ", Args[1]))
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("area ", Args[0], " range ", Args[1]))
			cost := configCandidate.value(
				[]string{"protocols", "ospf", "area", area, "range", range_, "cost"})
			if cost != nil {
				quaggaVtysh("configure terminal",
					"router ospf",
					fmt.Sprint("area ", Args[0], " range ", Args[1], " cost ", *cost))
			}
			substitute := configCandidate.value(
				[]string{"protocols", "ospf", "area", area, "range", range_, "substitute"})
			if substitute != nil {
				quaggaVtysh("configure terminal",
					"router ospf",
					fmt.Sprint("area ", Args[0], " range ", Args[1], " substitute ", *substitute))
			}
		}
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"range", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " range ", Args[1]))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for this range
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 16777215; "Metric must be between 0-16777215"
	val_help: u32: 0-16777215; Metric for this range
*/
func quaggaProtocolsOspfAreaRangeCost(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD range A.B.C.D/M cost WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " range ", Args[1], " cost ", Args[2]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"range", fmt.Sprint(Args[1]), "cost", fmt.Sprint(Args[2])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " range ", Args[1], " cost ", Args[2]))
		}
	}
	return cmd.Success
}

/*
	help: Do not advertise this range
	create:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"area $VAR(../../@) range $VAR(../@) not-advertise\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no area $VAR(../../@) range $VAR(../@) not-advertise\"; "
*/
func quaggaProtocolsOspfAreaRangeNotAdvertise(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD range A.B.C.D/M not-advertise
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " range ", Args[1], " not-advertise"))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"range", fmt.Sprint(Args[1]), "not-advertise"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " range ", Args[1], " not-advertise"))
		}
	}
	return cmd.Success
}

/*
	type: ipv4net
	help: Announce area range as another prefix
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"
*/
func quaggaProtocolsOspfAreaRangeSubstitute(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD range A.B.C.D/M substitute A.B.C.D/M
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " range ", Args[1], " substitute ", Args[2]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"range", fmt.Sprint(Args[1]), "substitute", fmt.Sprint(Args[2])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " range ", Args[1], " substitute ", Args[2]))
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: Area's shortcut mode
	allowed: echo "default disable enable"
	syntax:expression: $VAR(@) in "default", "disable", "enable"; "Must be (default, disable, enable)"
	val_help: default; Set default;
	val_help: disable; Disable shortcutting mode;
	val_help: enable; Enable  shortcutting mode;

	update:expression: "vtysh -c \"configure terminal\" \
	          -c \"router ospf\" \
	          -c \"area $VAR(../@) shortcut $VAR(@)\"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	          -c \"router ospf\" \
	          -c \"no area $VAR(../@) shortcut $VAR(@)\"; "
*/
func quaggaProtocolsOspfAreaShortcut(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD shortcut WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " shortcut ", Args[1]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"shortcut", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " shortcut ", Args[1]))
		}
	}
	return cmd.Success
}

/*
	tag:
	type: ipv4
	help: Virtual link
	syntax:expression: ! $VAR(../@) in "0", "0.0.0.0"; "Can't configure VL over area $VAR(../@)"
	create:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"area $VAR(../@) virtual-link $VAR(@)\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no area $VAR(../@) virtual-link $VAR(@)\"; "
*/
func quaggaProtocolsOspfAreaVirtualLink(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " virtual-link ", Args[1]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"virtual-link", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " virtual-link ", Args[1]))
		}
	}
	return cmd.Success
}

/*
	tag:
	type: u32
	help: MD5 key id
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "ID must be between (1-255)"
	val_help: u32:1-255; MD5 key id

	commit:expression: $VAR(md5-key/) != ""; "Must add the md5-key for key-id $VAR(@)"

	delete:expression: "touch /tmp/ospf-md5.$PPID"

	end:expression: "\
	      if [ -f \"/tmp/ospf-md5.$PPID\" ]; then                          \
	         vtysh -c \"configure terminal\"   \
	           -c \"router ospf\"                                          \
	           -c \"no area $VAR(../../../../@)                            \
	           virtual-link $VAR(../../../@) message-digest-key $VAR(@)\"; \
	         rm /tmp/ospf-md5.$PPID;                                       \
	      else                                                             \
	         vtysh -c \"configure terminal\"   \
	           -c \"router ospf\"                                          \
	           -c \"area $VAR(../../../../@) virtual-link $VAR(../../../@) \
	           message-digest-key $VAR(@) md5 $VAR(md5-key/@)\";           \
	      fi; "
*/
func quaggaProtocolsOspfAreaVirtualLinkAuthenticationMd5KeyId(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D authentication md5 key-id WORD
	switch Cmd {
	case cmd.Set:
		area := fmt.Sprint(Args[0])
		virtualLink := fmt.Sprint(Args[1])
		keyId := fmt.Sprint(Args[2])
		md5Key := configCandidate.value(
			[]string{"protocols", "ospf", "area", area, "virtual-link", virtualLink,
				"authentication", "md5", "key-id", keyId, "md5-key"})
		if md5Key != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("area ", Args[0], " virtual-link ", Args[1],
					" message-digest-key ", Args[2], " md5 ", *md5Key))
		}
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"virtual-link", fmt.Sprint(Args[1]),
				"authentication", "md5", "key-id", fmt.Sprint(Args[2])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " virtual-link ", Args[1], " message-digest-key ", Args[2]))
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: MD5 key
	val_help: MD5 Key (16 characters or less)

	syntax:expression: exec "                              \
	        if [ `echo -n '$VAR(@)' | wc -c` -gt 16 ]; then  \
	          echo MD5 key must be 16 characters or less ; \
	          exit 1 ;                                     \
	        fi ; "
*/
func quaggaProtocolsOspfAreaVirtualLinkAuthenticationMd5KeyIdMd5Key(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D authentication md5 key-id WORD md5-key WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: MD5 key id
	commit:expression: $VAR(../plaintext-password/) == "" ; "plaintext-password already set"

	create: vtysh -c "configure terminal"                              \
	          -c "router ospf"                                         \
	          -c "no area $VAR(../../../@) virtual-link $VAR(../../@)  \
	            authentication-key"                                    \
	          -c "area $VAR(../../../@) virtual-link $VAR(../../@)     \
	            authentication message-digest";

	delete: vtysh -c "configure terminal"                              \
	          -c "router ospf"                                         \
	          -c "area $VAR(../../../@) virtual-link $VAR(../../@)     \
	            authentication null";
*/
func quaggaProtocolsOspfAreaVirtualLinkAuthenticationMd5(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D authentication md5
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no area ", Args[0], " virtual-link ", Args[1], " authentication-key"),
			fmt.Sprint("area ", Args[0], " virtual-link ", Args[1], " authentication message-digest"))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"virtual-link", fmt.Sprint(Args[1]),
				"authentication", "md5"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("area ", Args[0], " virtual-link ", Args[1], " authentication null"))
		}
	}
	return cmd.Success
}

/*
	help: Authentication
*/
func quaggaProtocolsOspfAreaVirtualLinkAuthentication(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D authentication
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Plain text password
	val_help: Plain text password (8 characters or less)

	syntax:expression: exec "                              \
	        if [ `echo -n '$VAR(@)' | wc -c` -gt 8 ]; then   \
	          echo Password must be 8 characters or less ; \
	          exit 1 ;                                     \
	        fi ; "

	commit:expression: $VAR(../md5/) == "" ; "md5 password already set"

	update: vtysh -c "configure terminal" -c "router ospf"         \
	          -c "area $VAR(../../../@) virtual-link $VAR(../../@) \
	            authentication authentication-key $VAR(@) "

	delete: vtysh  -c "configure terminal" -c "router ospf"           \
	          -c "no area $VAR(../../../@) virtual-link $VAR(../../@) \
	            authentication authentication-key";
*/
func quaggaProtocolsOspfAreaVirtualLinkAuthenticationPlaintextPassword(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D authentication plaintext-password WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " virtual-link ", Args[1],
				" authentication authentication-key ", Args[2]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"virtual-link", fmt.Sprint(Args[1]),
				"authentication", "plaintext-password", fmt.Sprint(Args[2])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " virtual-link ", Args[1],
					" authentication authentication-key"))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval after which a neighbor is declared dead
	val_help: u32:1-65535; Neighbor dead interval (seconds)

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"

	update:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"area $VAR(../../@) virtual-link $VAR(../@) dead-interval $VAR(@)\"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no area $VAR(../../@) virtual-link $VAR(../@) dead-interval \"; "
*/
func quaggaProtocolsOspfAreaVirtualLinkDeadInterval(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D dead-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " virtual-link ", Args[1],
				" dead-interval ", Args[2]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"virtual-link", fmt.Sprint(Args[1]),
				"dead-interval", fmt.Sprint(Args[2])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " virtual-link ", Args[1],
					" dead-interval"))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between hello packets
	val_help: u32:1-65535; Hello interval (seconds)

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"

	update:expression: "vtysh -c \"configure terminal\" \
	      -c \"router ospf\" \
	      -c \"area $VAR(../../@) virtual-link $VAR(../@) hello-interval $VAR(@)\"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	      -c \"router ospf\" \
	      -c \"no area $VAR(../../@) virtual-link $VAR(../@) hello-interval \"; "
*/
func quaggaProtocolsOspfAreaVirtualLinkHelloInterval(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D hello-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " virtual-link ", Args[1],
				" hello-interval ", Args[2]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"virtual-link", fmt.Sprint(Args[1]),
				"hello-interval", fmt.Sprint(Args[2])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " virtual-link ", Args[1],
					" hello-interval"))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: Interval between retransmitting lost link state advertisements
	val_help: u32:1-65535; Retransmit interval (seconds)

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"

	update:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"area $VAR(../../@) virtual-link $VAR(../@) \
	       retransmit-interval $VAR(@)\"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no area $VAR(../../@) virtual-link $VAR(../@) \
	       retransmit-interval \"; "
*/
func quaggaProtocolsOspfAreaVirtualLinkRetransmitInterval(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D retransmit-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " virtual-link ", Args[1],
				" retransmit-interval ", Args[2]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"virtual-link", fmt.Sprint(Args[1]),
				"etransmit-interval", fmt.Sprint(Args[2])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " virtual-link ", Args[1],
					" retransmit-interval"))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: Link state transmit delay
	val_help: u32:1-65535; Link state transmit delay (seconds)

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535"

	update:expression: "vtysh -c \"configure terminal\" \
	      -c \"router ospf\" \
	      -c \"area $VAR(../../@) virtual-link $VAR(../@) transmit-delay $VAR(@)\"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	      -c \"router ospf\" \
	      -c \"no area $VAR(../../@) virtual-link $VAR(../@) transmit-delay \"; "
*/
func quaggaProtocolsOspfAreaVirtualLinkTransmitDelay(Cmd int, Args cmd.Args) int {
	//protocols ospf area WORD virtual-link A.B.C.D transmit-delay WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("area ", Args[0], " virtual-link ", Args[1],
				" transmit-delay ", Args[2]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "area", fmt.Sprint(Args[0]),
				"virtual-link", fmt.Sprint(Args[1]),
				"transmit-delay", fmt.Sprint(Args[2])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no area ", Args[0], " virtual-link ", Args[1],
					" transmit-delay"))
		}
	}
	return cmd.Success
}

/*
	help: Calculate OSPF interface cost according to bandwidth
*/
func quaggaProtocolsOspfAutoCost(Cmd int, Args cmd.Args) int {
	//protocols ospf auto-cost
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Reference bandwidth method to assign OSPF cost
	default: 100
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 4294967; \
	       "Must be between 1-4294967"
	val_help: u32:1-4294967; Reference bandwidth cost in Mbits/sec (default 100)

	update:expression: "vtysh --noerror \
	       -c \"configure terminal\"                       \
	       -c \"router ospf\"                              \
	       -c \"auto-cost reference-bandwidth $VAR(@) \";  \
	       echo 'OSPF: Reference bandwidth is changed.';   \
	       echo '      Please ensure reference bandwidth is consistent across all routers'; "

	delete:expression: "vtysh --noerror \
	       -c \"configure terminal\"                       \
	       -c \"router ospf\"                              \
	       -c \"no auto-cost reference-bandwidth \";       \
	       echo 'OSPF: Reference bandwidth is changed.';   \
	       echo '      Please ensure reference bandwidth is consistent across all routers'; "
*/
func quaggaProtocolsOspfAutoCostReferenceBandwidth(Cmd int, Args cmd.Args) int {
	//protocols ospf auto-cost reference-bandwidth WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("auto-cost reference-bandwidth ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "auto-cost", "reference-bandwidth", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no auto-cost reference-bandwidth")
		}
	}
	return cmd.Success
}

/*
	help: Control distribution of default information
*/
func quaggaProtocolsOspfDefaultInformation(Cmd int, Args cmd.Args) int {
	//protocols ospf default-information
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Always advertise default route
*/
func quaggaProtocolsOspfDefaultInformationOriginateAlways(Cmd int, Args cmd.Args) int {
	//protocols ospf default-information originate always
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	always := configCandidate.lookup(
		[]string{"protocols", "ospf", "default-information", "originate", "always"})
	if always != nil {
		parm += " always"
	}
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate",
			fmt.Sprint("default-information originate", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate")
	}
	return cmd.Success
}

/*
	type: u32
	help: OSPF metric type for default routes
	default: 2
	syntax:expression: $VAR(@) in 1, 2 ; "metric must be either 1 or 2"
	val_help: u32:1-2; Metric type for default routes (default 2)
*/
func quaggaProtocolsOspfDefaultInformationOriginateMetricType(Cmd int, Args cmd.Args) int {
	//protocols ospf default-information originate metric-type WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	always := configCandidate.lookup(
		[]string{"protocols", "ospf", "default-information", "originate", "always"})
	if always != nil {
		parm += " always"
	}
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate",
			fmt.Sprint("default-information originate", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate")
	}
	return cmd.Success
}

/*
	type: u32
	help: OSPF default metric
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 16777214; "must be between 0-16777214"
	val_help: u32:0-16777214; Default metric
*/
func quaggaProtocolsOspfDefaultInformationOriginateMetric(Cmd int, Args cmd.Args) int {
	//protocols ospf default-information originate metric WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	always := configCandidate.lookup(
		[]string{"protocols", "ospf", "default-information", "originate", "always"})
	if always != nil {
		parm += " always"
	}
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate",
			fmt.Sprint("default-information originate", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate")
	}
	return cmd.Success
}

/*
	help: Distribute a default route
	delete: touch /tmp/ospf-default-info.$PPID
	end: if [ -f "/tmp/ospf-default-info.$PPID" ]; then
	        vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "no default-information originate";
	     else
	        # uncomment and remove script pending bug 2525
	        #
	        # if [ -n "$VAR(./always/)" ]; then
	        #  PARM="always";
	        # fi;
	        ${vyatta_sbindir}/vyatta-check-typeless-node.pl           \
	          "protocols ospf default-information originate always";
	        if [ $? -eq 0 ] ; then
	           PARM="always";
	        fi;
	        if [ -n "$VAR(./metric/@)" ]; then
	           PARM="$PARM metric $VAR(./metric/@)";
	        fi;
	        if [ -n "$VAR(./metric-type/@)" ]; then
	           PARM="$PARM metric-type $VAR(./metric-type/@)";
	        fi;
	        if [ -n "$VAR(./route-map/@)" ]; then
	           PARM="$PARM route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "default-information originate $PARM";
	     fi;
*/
func quaggaProtocolsOspfDefaultInformationOriginate(Cmd int, Args cmd.Args) int {
	//protocols ospf default-information originate
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	always := configCandidate.lookup(
		[]string{"protocols", "ospf", "default-information", "originate", "always"})
	if always != nil {
		parm += " always"
	}
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate",
			fmt.Sprint("default-information originate", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate")
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
*/
func quaggaProtocolsOspfDefaultInformationOriginateRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospf default-information originate route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	always := configCandidate.lookup(
		[]string{"protocols", "ospf", "default-information", "originate", "always"})
	if always != nil {
		parm += " always"
	}
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "default-information", "originate", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate",
			fmt.Sprint("default-information originate", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no default-information originate")
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric of redistributed routes
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 16777214; "Must be between 0-16777214"
	val_help: u32:0-16777214; Metric of redistributed routes

	update:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"default-metric $VAR(@) \"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no default-metric $VAR(@) \"; "
*/
func quaggaProtocolsOspfDefaultMetric(Cmd int, Args cmd.Args) int {
	//protocols ospf default-metric WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"default-metric ", fmt.Sprint(Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "default-metric", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no default-metric ", fmt.Sprint(Args[0]))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: OSPF administrative distance
	val_help: u32:1-255; Administrative distance

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between 1-255"

	update:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"distance $VAR(@) \"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no distance $VAR(@) \"; "
*/
func quaggaProtocolsOspfDistanceGlobal(Cmd int, Args cmd.Args) int {
	//protocols ospf distance global WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"distance ", fmt.Sprint(Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "distance", "global", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no distance ", fmt.Sprint(Args[0]))
		}
	}
	return cmd.Success
}

/*
	help: Administrative distance
*/
func quaggaProtocolsOspfDistance(Cmd int, Args cmd.Args) int {
	//protocols ospf distance
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Distance for external routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between 1-255"
	val_help: u32: 1-255; Distance for external routes
*/
func quaggaProtocolsOspfDistanceOspfExternal(Cmd int, Args cmd.Args) int {
	//protocols ospf distance ospf external WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	intraArea := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "intra-area"})
	if intraArea != nil {
		parm += " intra-area " + *intraArea
	}
	interArea := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "inter-area"})
	if interArea != nil {
		parm += " inter-area " + *interArea
	}
	external := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "external"})
	if external != nil {
		parm += " external " + *external
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no distance ospf",
			fmt.Sprint("distance ospf", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no distance ospf")
	}
	return cmd.Success
}

/*
	type: u32
	help: Distance for inter-area routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between 1-255"
	val_help: u32:1-255; Distance for inter-area routes
*/
func quaggaProtocolsOspfDistanceOspfInterArea(Cmd int, Args cmd.Args) int {
	//protocols ospf distance ospf inter-area WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	intraArea := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "intra-area"})
	if intraArea != nil {
		parm += " intra-area " + *intraArea
	}
	interArea := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "inter-area"})
	if interArea != nil {
		parm += " inter-area " + *interArea
	}
	external := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "external"})
	if external != nil {
		parm += " external " + *external
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no distance ospf",
			fmt.Sprint("distance ospf", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no distance ospf")
	}
	return cmd.Success
}

/*
	type: u32
	help: Distance for intra-area routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 255; "Must be between 1-255"
	val_help: u32:1-255; Distance for intra-area routes
*/
func quaggaProtocolsOspfDistanceOspfIntraArea(Cmd int, Args cmd.Args) int {
	//protocols ospf distance ospf intra-area WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	intraArea := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "intra-area"})
	if intraArea != nil {
		parm += " intra-area " + *intraArea
	}
	interArea := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "inter-area"})
	if interArea != nil {
		parm += " inter-area " + *interArea
	}
	external := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "external"})
	if external != nil {
		parm += " external " + *external
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no distance ospf",
			fmt.Sprint("distance ospf", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no distance ospf")
	}
	return cmd.Success
}

/*
	help: OSPF administrative distance
	delete:expression: "touch /tmp/ospf-distance.$PPID"
	end:expression: "\
	      if [ -f \"/tmp/ospf-distance.$PPID\" ]; then                   \
	         vtysh -c \"configure terminal\" \
	           -c \"router ospf\"                                        \
	           -c \"no distance ospf\";                                  \
		 rm /tmp/ospf-distance.$PPID;                                \
	      else                                                           \
	         if [ -n \"$VAR(./intra-area/@)\" ]; then                    \
	           PARM=\"intra-area $VAR(./intra-area/@)\";                 \
	         fi;                                                         \
	         if [ -n \"$VAR(./inter-area/@)\" ]; then                    \
	           PARM=\"$PARM inter-area $VAR(./inter-area/@)\";           \
	         fi;                                                         \
	         if [ -n \"$VAR(./external/@)\" ]; then                      \
	           PARM=\"$PARM external $VAR(./external/@)\";               \
	         fi;                                                         \
	         vtysh -c \"configure terminal\" \
	           -c \"router ospf\"                                        \
	           -c \"no distance ospf\" -c \"distance ospf $PARM\";       \
	      fi; "
*/
func quaggaProtocolsOspfDistanceOspf(Cmd int, Args cmd.Args) int {
	//protocols ospf distance ospf
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	intraArea := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "intra-area"})
	if intraArea != nil {
		parm += " intra-area " + *intraArea
	}
	interArea := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "inter-area"})
	if interArea != nil {
		parm += " inter-area " + *interArea
	}
	external := configCandidate.value([]string{"protocols", "ospf", "distance", "ospf", "external"})
	if external != nil {
		parm += " external " + *external
	}
	if parm != "" {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no distance ospf",
			fmt.Sprint("distance ospf", parm))
	} else {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no distance ospf")
	}
	return cmd.Success
}

/*
	help: Log all state changes
	create:expression: "vtysh -c \"configure terminal\" -c \"router ospf\" \
	         -c \"log-adjacency-changes detail\"; "
	delete:expression: "vtysh -c \"configure terminal\" -c \"router ospf\" \
	         -c \"no log-adjacency-changes detail\"; "

*/
func quaggaProtocolsOspfLogAdjacencyChangesDetail(Cmd int, Args cmd.Args) int {
	//protocols ospf log-adjacency-changes detail
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"log-adjacency-changes detail")
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "log-adjacency-changes", "detail"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no log-adjacency-changes detail")
		}
	}
	return cmd.Success
}

/*
	help: Log changes in adjacency state
	create:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"log-adjacency-changes\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no log-adjacency-changes\"; "
*/
func quaggaProtocolsOspfLogAdjacencyChanges(Cmd int, Args cmd.Args) int {
	//protocols ospf log-adjacency-changes
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"log-adjacency-changes")
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "log-adjacency-changes"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no log-adjacency-changes")
		}
	}
	return cmd.Success
}

/*
	help: OSPF maximum/infinite-distance metric
*/
func quaggaProtocolsOspfMaxMetric(Cmd int, Args cmd.Args) int {
	//protocols ospf max-metric
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Administratively apply, for an indefinite period
	create:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"max-metric router-lsa administrative\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no max-metric router-lsa administrative \"; "
*/
func quaggaProtocolsOspfMaxMetricRouterLsaAdministrative(Cmd int, Args cmd.Args) int {
	//protocols ospf max-metric router-lsa administrative
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"max-metric router-lsa administrative")
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "max-metric", "router-lsa", "administrative"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no max-metric router-lsa administrative")
		}
	}
	return cmd.Success
}

/*
	help: Advertise own Router-LSA with infinite distance (stub router)
*/
func quaggaProtocolsOspfMaxMetricRouterLsa(Cmd int, Args cmd.Args) int {
	//protocols ospf max-metric router-lsa
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Advertise stub-router prior to full shutdown of OSPF
	syntax:expression: $VAR(@) >= 5 && $VAR(@) <= 86400; "must be between 5-86400 seconds"
	val_help: u32:5-86400; Time (seconds) to advertise self as stub-router

	update:expression: "vtysh -c \"configure terminal\" \
	        -c \"router ospf\" \
	        -c \"max-metric router-lsa on-shutdown $VAR(@)\"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	        -c \"router ospf\" \
	        -c \"no max-metric router-lsa on-shutdown \"; "
*/
func quaggaProtocolsOspfMaxMetricRouterLsaOnShutdown(Cmd int, Args cmd.Args) int {
	//protocols ospf max-metric router-lsa on-shutdown WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("max-metric router-lsa on-shutdown ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "max-metric", "router-lsa",
				"on-shutdown", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no max-metric router-lsa on-shutdown")
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: Automatically advertise stub Router-LSA on startup of OSPF
	syntax:expression: $VAR(@) >= 5 && $VAR(@) <= 86400; "must be between 5-86400 seconds"
	val_help: u32:5-86400; Time (seconds) to advertise self as stub-router

	update:expression: "vtysh -c \"configure terminal\" \
	        -c \"router ospf\" \
	        -c \"max-metric router-lsa on-startup $VAR(@)\"; "

	delete:expression: "vtysh -c \"configure terminal\" \
	        -c \"router ospf\" \
	        -c \"no max-metric router-lsa on-startup \"; "
*/
func quaggaProtocolsOspfMaxMetricRouterLsaOnStartup(Cmd int, Args cmd.Args) int {
	//protocols ospf max-metric router-lsa on-startup WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("router-lsa on-startup ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "max-metric", "router-lsa",
				"on-startup", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no max-metric router-lsa on-startup")
		}
	}
	return cmd.Success
}

/*
	help: Enable MPLS-TE functionality
	create:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"mpls-te on\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no mpls-te\"; "
*/
func quaggaProtocolsOspfMplsTeEnable(Cmd int, Args cmd.Args) int {
	//protocols ospf mpls-te enable
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"mpls-te on")
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "mpls-te", "enable"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no mpls-te")
		}
	}
	return cmd.Success
}

/*
	help: MultiProtocol Label Switching-Traffic Engineering (MPLS-TE) parameters
*/
func quaggaProtocolsOspfMplsTe(Cmd int, Args cmd.Args) int {
	//protocols ospf mpls-te
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Stable IP address of the advertising router
	update:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"mpls-te router-address $VAR(@)\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no mpls-te\"; "
*/
func quaggaProtocolsOspfMplsTeRouterAddress(Cmd int, Args cmd.Args) int {
	//protocols ospf mpls-te router-address A.B.C.D
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("mpls-te router-address ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "mpls-te", "router-address", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no mpls-te")
		}
	}
	return cmd.Success
}

/*
	tag:
	type: ipv4
	help: Neighbor IP address
	create:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"neighbor $VAR(@)\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf\" \
	       -c \"no neighbor $VAR(@)\"; "
*/
func quaggaProtocolsOspfNeighbor(Cmd int, Args cmd.Args) int {
	//protocols ospf neighbor A.B.C.D
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("neighbor ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "neighbor", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no neighbor ", Args[0]))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: Dead neighbor polling interval
	default: 60
	val_help: u32:1-65535; Seconds between dead neighbor polling interval (default 60)

	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 65535; "Must be between 1-65535 seconds"

	update: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "neighbor $VAR(../@) poll-interval $VAR(@)";

	delete: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "neighbor $VAR(../@) poll-interval 60";
*/
func quaggaProtocolsOspfNeighborPollInterval(Cmd int, Args cmd.Args) int {
	//protocols ospf neighbor A.B.C.D poll-interval WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("neighbor ", Args[0], " poll-interval ", Args[1]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "neighbor", fmt.Sprint(Args[0]),
				"poll-interval", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("neighbor ", Args[0], " poll-interval 60"))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: Neighbor priority in seconds
	default: 0
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 255; "Priority must be between 0-255"
	val_help: u32:0-255; Neighbor priority (default 0)

	update: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "neighbor $VAR(../@) priority $VAR(@)";

	delete: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "neighbor $VAR(../@) priority 0";
*/
func quaggaProtocolsOspfNeighborPriority(Cmd int, Args cmd.Args) int {
	//protocols ospf neighbor A.B.C.D priority WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("neighbor ", Args[0], " priority ", Args[1]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "neighbor", fmt.Sprint(Args[0]),
				"priority", fmt.Sprint(Args[1])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("neighbor ", Args[0], " priority 0"))
		}
	}
	return cmd.Success
}

/*
	priority: 620
	help: Open Shortest Path First protocol (OSPF) parameters
	begin: if [ "$COMMIT_ACTION" != DELETE ]; then
	         if [ -n "$VAR(parameters/router-id/@)" ]; then
	           vtysh -c "configure terminal" -c "router ospf" \
	                 -c "ospf router-id $VAR(parameters/router-id/@)"
	         else
	           vtysh -c "configure terminal" -c "router ospf" \
	                 -c "no ospf router-id"
	         fi
	       fi
	end: if [ "$COMMIT_ACTION" == DELETE ]; then
	       vtysh -c "configure terminal" -c "router ospf" -c "no ospf router-id"
	       vtysh -c "configure terminal" -c "no router ospf"
	       rm -f /opt/vyatta/etc/quagga/ospfd.conf
	     else
	       vtysh -d ospfd -c 'sh run' > /opt/vyatta/etc/quagga/ospfd.conf
	     fi
*/
func quaggaProtocolsOspf(Cmd int, Args cmd.Args) int {
	//protocols ospf
	switch Cmd {
	case cmd.Set:
		routerId := configCandidate.value(
			[]string{"protocols", "ospf", "parameters", "router-id"})
		if routerId != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("ospf router-id ", *routerId))
		} else {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no ospf router-id")
		}
	case cmd.Delete:
		if configRunning.lookup([]string{"protocols", "ospf"}) != nil {
			quaggaVtysh("configure terminal", "no router ospf")
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: OSPF ABR type
	default: "cisco"
	syntax:expression: $VAR(@) in "cisco", "ibm", "shortcut", "standard"; "Must be (cisco, ibm, shortcut, standard)"
	val_help: cisco; Cisco ABR type (default)
	val_help: ibm; Ibm ABR type
	val_help: shortcut; Shortcut ABR type
	val_help: standard; Standard ABR type

	update: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "ospf abr-type $VAR(@)";

	delete: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "ospf abr-type cisco";

*/
func quaggaProtocolsOspfParametersAbrType(Cmd int, Args cmd.Args) int {
	//protocols ospf parameters abr-type WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("ospf abr-type ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "parameters", "abr-type", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"ospf abr-type cisco")
		}
	}
	return cmd.Success
}

/*
	help: OSPF specific parameters
*/
func quaggaProtocolsOspfParameters(Cmd int, Args cmd.Args) int {
	//protocols ospf parameters
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Enable the Opaque-LSA capability (rfc2370)
	create:expression: "vtysh -c \"configure terminal\" -c \"router ospf\" \
	          -c \"ospf opaque-lsa \"; "
	delete:expression: "vtysh -c \"configure terminal\" -c \"router ospf\" \
	          -c \"no ospf opaque-lsa \"; "

*/
func quaggaProtocolsOspfParametersOpaqueLsa(Cmd int, Args cmd.Args) int {
	//protocols ospf parameters opaque-lsa
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"ospf opaque-lsa")
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "parameters", "opaque-lsa"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no ospf opaque-lsa")
		}

	}
	return cmd.Success
}

/*
	help: Enable rfc1583 criteria for handling AS external routes
	create:expression: "vtysh -c \"configure terminal\" -c \"router ospf\" \
	          -c \"ospf rfc1583compatibility \"; "
	delete:expression: "vtysh -c \"configure terminal\" -c \"router ospf\" \
	          -c \"no ospf rfc1583compatibility \"; "
*/
func quaggaProtocolsOspfParametersRfc1583Compatibility(Cmd int, Args cmd.Args) int {
	//protocols ospf parameters rfc1583-compatibility
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"ospf rfc1583compatibility")
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "parameters", "rfc1583compatibility"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no ospf rfc1583compatibility")
		}
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Override the default router identifier
*/
func quaggaProtocolsOspfParametersRouterId(Cmd int, Args cmd.Args) int {
	//protocols ospf parameters router-id A.B.C.D
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("ospf router-id ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "parameters", "router-id", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no ospf router-id")
		}
	}
	return cmd.Success
}

/*
	multi:
	type: txt
	help: Interface to exclude when using 'passive-interface default'
	val_help:<interface>; Interface to exclude from 'passive-interface default'

	allowed: ${vyatta_sbindir}/vyatta-interfaces.pl --show all

	syntax:expression: $VAR(../passive-interface/@) == "default"; \
	  "passive-interface-excluded can only be used with 'passive-interface default'"

	commit:expression: exec "/opt/vyatta/sbin/vyatta-interfaces.pl --dev=$VAR(@) --warn"

	create: if [ -z $VAR(@) ] ; then
	          echo "Error: must include interface";
	          exit 1;
	     	else
	          vtysh -c "configure terminal" -c "router ospf" \
	             -c "no passive-interface $VAR(@)"
		fi;

	delete: vtysh -c "configure terminal" -c "router ospf" \
	          -c "passive-interface $VAR(@)";
*/
func quaggaProtocolsOspfPassiveInterfaceExclude(Cmd int, Args cmd.Args) int {
	//protocols ospf passive-interface-exclude WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("no passive-interface ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "passive-interface-exclude", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("passive-interface ", Args[0]))
		}
	}
	return cmd.Success
}

/*
	multi:
	type: txt
	help: Suppress routing updates on an interface
	allowed: ${vyatta_sbindir}/vyatta-interfaces.pl --show all && echo default
	val_help:<interface>; Interface to be passive (i.e. suppress routing updates)
	val_help:default; Default to suppress routing updates on all interfaces

	create: sudo /opt/vyatta/sbin/vyatta_quagga_utils.pl \
	           --check-ospf-passive="$VAR(@)"
	        if [ $? != 0 ] ; then
	           exit 1;
	        fi
	        if [ -z $VAR(@) ] || [ "$VAR(@)" == "default" ] ; then
	           vtysh -c "configure terminal" \
	                 -c "router ospf"        \
	                 -c "passive-interface default";
	        else
	           vtysh -c "configure terminal" \
	                 -c "router ospf"        \
	                 -c "passive-interface $VAR(@)"
	        fi

	delete: if [ -z $VAR(@) ]
		then
	           vtysh -c "configure terminal" \
	                 -c "router ospf"        \
	                 -c "no passive-interface default"
		else
	           if [ "$VAR(@)" == "default" ]
	           then
	              if [ $VAR(../passive-interface-exclude/@) ]
	              then
	                 echo "Error: delete passive-interface-exclude before deleting passive-interface default";
	                 exit 1;
	              fi
	           fi
	           vtysh -c "configure terminal" \
	                 -c "router ospf"        \
	                 -c "no passive-interface $VAR(@)"
	 	fi
*/
func quaggaProtocolsOspfPassiveInterface(Cmd int, Args cmd.Args) int {
	//protocols ospf passive-interface WORD
	switch Cmd {
	case cmd.Set:
		if fmt.Sprint(Args[0]) == "default" {
			quaggaVtysh("configure terminal",
				"router ospf",
				"passive-interface default")
		} else {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("passive-interface ", Args[0]))
		}
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "passive-interface", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no passive-interface ", Args[0]))
		}
	}
	return cmd.Success
}

/*
	type: u32
	help: OSPF metric type
	default: 2
	syntax:expression: $VAR(@) in 1, 2 ; "metric-type must be either 1 or 2"
	val_help: u32:1-2; Metric type (default 2)
*/
func quaggaProtocolsOspfRedistributeBgpMetricType(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute bgp metric-type WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute bgp", parm))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 16; "metric must be between 1 and 16"
	val_help: u32:1-16; Metric for redistributed routes
*/
func quaggaProtocolsOspfRedistributeBgpMetric(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute bgp metric WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute bgp", parm))
	}
	return cmd.Success
}

/*
	help: Redistribute BGP routes
	delete:expression: "touch /tmp/ospf-redist-bgp.$PPID"
	end: vtysh -c "configure terminal"      \
	        -c "router ospf"                                             \
	        -c "no redistribute bgp";
	     if [ -f "/tmp/ospf-redist-bgp.$PPID" ]; then
	        rm -f /tmp/ospf-redist-bgp.$PPID;
	     else
	        if [ -n "$VAR(./metric/@)" ]; then
	           COND="metric $VAR(./metric/@)";
	        fi;
		if [ -n "$VAR(./metric-type/@)" ]; then
		   COND="$COND metric-type $VAR(./metric-type/@)";
	        fi;
	        if [ -n "$VAR(./route-map/@)" ]; then
	           COND="$COND route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	           -c "router ospf"                                       \
	           -c "redistribute bgp $COND";
	     fi;
*/
func quaggaProtocolsOspfRedistributeBgp(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute bgp
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"redistribute bgp")
	case cmd.Delete:
		if configRunning.lookup([]string{"protocols", "ospf", "redistribute", "bgp"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no redistribute bgp")
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

*/
func quaggaProtocolsOspfRedistributeBgpRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute bgp route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "bgp", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute bgp", parm))
	}
	return cmd.Success
}

/*
	type: u32
	help: OSPF metric type
	default: 2
	syntax:expression: $VAR(@) in 1, 2 ; "metric-type must be either 1 or 2"
	val_help: u32:1-2; Metric type (default 2)
*/
func quaggaProtocolsOspfRedistributeConnectedMetricType(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute connected metric-type WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute connected", parm))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 16; "metric must be between 1 and 16"
	val_help: u32:1-16; Metric for redistributed routes
*/
func quaggaProtocolsOspfRedistributeConnectedMetric(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute connected metric WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute connected", parm))
	}
	return cmd.Success
}

/*
	help: Redistribute connected routes
	delete:expression: "touch /tmp/ospf-redist-connected.$PPID"
	end: vtysh -c "configure terminal"    \
	        -c "router ospf"                                          \
	        -c "no redistribute connected";
	     if [ -f "/tmp/ospf-redist-connected.$PPID" ]; then
	        rm -f /tmp/ospf-redist-connected.$PPID;
	     else
	        if [ -n "$VAR(./metric/@)" ]; then
	           COND="metric $VAR(./metric/@)";
	        fi;
		if [ -n "$VAR(./metric-type/@)" ]; then
		   COND="$COND metric-type $VAR(./metric-type/@)";
	        fi;
	        if [ -n "$VAR(./route-map/@)" ]; then
	           COND="$COND route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	           -c "router ospf"                                       \
	           -c "redistribute connected $COND";
	     fi;
*/
func quaggaProtocolsOspfRedistributeConnected(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute connected
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"redistribute connected")
	case cmd.Delete:
		if configRunning.lookup([]string{"protocols", "ospf", "redistribute", "connected"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no redistribute connected")
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsOspfRedistributeConnectedRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute connected route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "connected", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute connected", parm))
	}
	return cmd.Success
}

/*
	type: u32
	help: OSPF metric type
	default: 2
	syntax:expression: $VAR(@) in 1, 2 ; "metric-type must be either 1 or 2"
	val_help: u32:1-2; Metric type (default 2)
*/
func quaggaProtocolsOspfRedistributeKernelMetricType(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute kernel metric-type WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute kernel", parm))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 16; "metric must be between 1 and 16"
	val_help: u32:1-16; Metric for redistributed routes
*/
func quaggaProtocolsOspfRedistributeKernelMetric(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute kernel metric WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute kernel", parm))
	}
	return cmd.Success
}

/*
	help: Redistribute kernel routes
	delete:expression: "touch /tmp/ospf-redist-kernel.$PPID"
	end: vtysh -c "configure terminal" \
	                  -c "router ospf"        \
	                  -c "no redistribute kernel";
	     if [ -f "/tmp/ospf-redist-kernel.$PPID" ]; then
	        rm -f /tmp/ospf-redist-kernel.$PPID;
	     else
	        if [ -n "$VAR(./metric/@)" ]; then
	           COND="metric $VAR(./metric/@)";
	        fi;
	        if [ -n "$VAR(./metric-type/@)" ]; then
		   COND="$COND metric-type $VAR(./metric-type/@)";
	        fi;
	        if [ -n "$VAR(./route-map/@)" ]; then
	           COND="$COND route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	                     -c "router ospf"        \
	                     -c "redistribute kernel $COND";
	     fi;
*/
func quaggaProtocolsOspfRedistributeKernel(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute kernel
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"redistribute kernel")
	case cmd.Delete:
		if configRunning.lookup([]string{"protocols", "ospf", "redistribute", "kernel"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no redistribute kernel")
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsOspfRedistributeKernelRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute kernel route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "kernel", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute kernel", parm))
	}
	return cmd.Success
}

/*
	help: Redistribute information from another routing protocol
*/
func quaggaProtocolsOspfRedistribute(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: OSPF metric type
	default: 2
	syntax:expression: $VAR(@) in 1, 2 ; "metric-type must be either 1 or 2"
	val_help: u32:1-2; Metric type (default 2)
*/
func quaggaProtocolsOspfRedistributeRipMetricType(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute rip metric-type WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute rip", parm))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 16; "metric must be between 1 and 16"
	val_help: u32:1-16; Metric for redistributed routes
*/
func quaggaProtocolsOspfRedistributeRipMetric(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute rip metric WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute rip", parm))
	}
	return cmd.Success
}

/*
	help: Redistribute RIP routes
	delete:expression: "touch /tmp/ospf-redist-rip.$PPID"
	end: vtysh -c "configure terminal"    \
	        -c "router ospf"                                          \
	        -c "no redistribute rip";
	     if [ -f "/tmp/ospf-redist-rip.$PPID" ]; then
	        rm -f /tmp/ospf-redist-rip.$PPID;
	     else
	        if [ -n "$VAR(./metric/@)" ]; then
	           COND="metric $VAR(./metric/@)";
	        fi;
		if [ -n "$VAR(./metric-type/@)" ]; then
		   COND="$COND metric-type $VAR(./metric-type/@)";
	        fi;
	        if [ -n "$VAR(./route-map/@)" ]; then
	           COND="$COND route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "redistribute rip $COND";
	     fi;
*/
func quaggaProtocolsOspfRedistributeRip(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute rip
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"redistribute rip")
	case cmd.Delete:
		if configRunning.lookup([]string{"protocols", "ospf", "redistribute", "rip"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no redistribute rip")
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsOspfRedistributeRipRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute rip route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "rip", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute rip", parm))
	}
	return cmd.Success
}

/*
	type: u32
	help: OSPF metric type
	default: 2
	syntax:expression: $VAR(@) in 1, 2 ; "metric-type must be either 1 or 2"
	val_help: u32:1-2; Metric type (default 2)
*/
func quaggaProtocolsOspfRedistributeStaticMetricType(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute static metric-type WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute static", parm))
	}
	return cmd.Success
}

/*
	type: u32
	help: Metric for redistributed routes
	syntax:expression: $VAR(@) >= 1 && $VAR(@) <= 16; "metric must be between 1 and 16"
	val_help: u32:1-16; Metric for redistributed routes
*/
func quaggaProtocolsOspfRedistributeStaticMetric(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute static metric WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute static", parm))
	}
	return cmd.Success
}

/*
	help: Redistribute static routes
	delete:expression: "touch /tmp/ospf-redist-static.$PPID"
	end: vtysh -c "configure terminal"    \
	        -c "router ospf"                                          \
	        -c "no redistribute static";
	     if [ -f "/tmp/ospf-redist-static.$PPID" ]; then
	        rm -f /tmp/ospf-redist-static.$PPID;
	     else
	        if [ -n "$VAR(./metric/@)" ]; then
	           COND="metric $VAR(./metric/@)";
	        fi;
		if [ -n "$VAR(./metric-type/@)" ]; then
		   COND="$COND metric-type $VAR(./metric-type/@)";
	         fi;
	        if [ -n "$VAR(./route-map/@)" ]; then
	           COND="$COND route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	           -c "router ospf"                                       \
	           -c "redistribute static $COND";
	     fi;
*/
func quaggaProtocolsOspfRedistributeStatic(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute static
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"redistribute static")
	case cmd.Delete:
		if configRunning.lookup([]string{"protocols", "ospf", "redistribute", "static"}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				"no redistribute static")
		}
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsOspfRedistributeStaticRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospf redistribute static route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	parm := ""
	metric := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "metric"})
	if metric != nil {
		parm += " metric " + *metric
	}
	metricType := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "metric-type"})
	if metricType != nil {
		parm += " metric-type " + *metricType
	}
	routeMap := configCandidate.value(
		[]string{"protocols", "ospf", "redistribute", "static", "route-map"})
	if routeMap != nil {
		parm += " route-map " + *routeMap
	}
	if metric != nil || metricType != nil || routeMap != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("redistribute static", parm))
	}
	return cmd.Success
}

/*
	help: Adjust refresh parameters
*/
func quaggaProtocolsOspfRefresh(Cmd int, Args cmd.Args) int {
	//protocols ospf refresh
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Refresh timer
	syntax:expression: $VAR(@) >= 10 && $VAR(@) <= 1800; "must be between 10-1800"
	val_help: u32:10-1800; Timer value in seconds

	update: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "no refresh timer" -c "refresh timer $VAR(@)";

	delete: vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "no refresh timer $VAR(@)";
*/
func quaggaProtocolsOspfRefreshTimers(Cmd int, Args cmd.Args) int {
	//protocols ospf refresh timers WORD
	switch Cmd {
	case cmd.Set:
		quaggaVtysh("configure terminal",
			"router ospf",
			"no refresh timer",
			fmt.Sprint("refresh timer ", Args[0]))
	case cmd.Delete:
		if configRunning.lookup(
			[]string{"protocols", "ospf", "refresh", "timers", fmt.Sprint(Args[0])}) != nil {
			quaggaVtysh("configure terminal",
				"router ospf",
				fmt.Sprint("no refresh timer ", Args[0]))
		}
	}
	return cmd.Success
}

/*
	help: Adjust routing timers
*/
func quaggaProtocolsOspfTimers(Cmd int, Args cmd.Args) int {
	//protocols ospf timers
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Throttling adaptive timers
*/
func quaggaProtocolsOspfTimersThrottle(Cmd int, Args cmd.Args) int {
	//protocols ospf timers throttle
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: u32
	help: Delay (msec) from first change received till SPF calculation
	default: 200
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 600000; "must be between 0-600000"
	val_help: u32:0-600000; Delay in msec (default 200)
*/
func quaggaProtocolsOspfTimersThrottleSpfDelay(Cmd int, Args cmd.Args) int {
	//protocols ospf timers throttle spf delay WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	delay := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "delay"})
	initialHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "initial-holdtime"})
	maxHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "max-holdtime"})
	if delay != nil && initialHoldtime != nil && maxHoldtime != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("timers throttle spf ", *delay, " ", *initialHoldtime, " ", *maxHoldtime))
	} else if configRunning.lookup([]string{"protocols", "ospf"}) != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no timers throttle spf")
	}
	return cmd.Success
}

/*
	type: u32
	help: Initial hold time(msec) between consecutive SPF calculations
	default: 1000
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 600000; "must be between 0-600000"
	val_help: u32:0-600000; Initial hold time in msec (default 1000)
*/
func quaggaProtocolsOspfTimersThrottleSpfInitialHoldtime(Cmd int, Args cmd.Args) int {
	//protocols ospf timers throttle spf initial-holdtime WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	delay := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "delay"})
	initialHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "initial-holdtime"})
	maxHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "max-holdtime"})
	if delay != nil && initialHoldtime != nil && maxHoldtime != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("timers throttle spf ", *delay, " ", *initialHoldtime, " ", *maxHoldtime))
	} else if configRunning.lookup([]string{"protocols", "ospf"}) != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no timers throttle spf")
	}
	return cmd.Success
}

/*
	type: u32
	help: Maximum hold time (msec)
	default: 10000
	syntax:expression: $VAR(@) >= 0 && $VAR(@) <= 600000; "must be between 0-600000"
	val_help: u32:0-600000; Max hold time in msec (default 10000)
*/
func quaggaProtocolsOspfTimersThrottleSpfMaxHoldtime(Cmd int, Args cmd.Args) int {
	//protocols ospf timers throttle spf max-holdtime WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	delay := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "delay"})
	initialHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "initial-holdtime"})
	maxHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "max-holdtime"})
	if delay != nil && initialHoldtime != nil && maxHoldtime != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("timers throttle spf ", *delay, " ", *initialHoldtime, " ", *maxHoldtime))
	} else if configRunning.lookup([]string{"protocols", "ospf"}) != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no timers throttle spf")
	}
	return cmd.Success
}

/*
	help: OSPF SPF timers
	delete: touch /tmp/ospf-timer.$PPID
	end: if [ -f "/tmp/ospf-timer.$PPID" ]; then
	        vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "no timers throttle spf";
	        rm /tmp/ospf-timer.$PPID;
	     else
	        vtysh -c "configure terminal" \
	          -c "router ospf"                                        \
	          -c "timers throttle spf $VAR(delay/@) $VAR(initial-holdtime/@) $VAR(max-holdtime/@)";
	     fi;
*/
func quaggaProtocolsOspfTimersThrottleSpf(Cmd int, Args cmd.Args) int {
	//protocols ospf timers throttle spf
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	delay := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "delay"})
	initialHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "initial-holdtime"})
	maxHoldtime := configCandidate.value(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "max-holdtime"})
	if delay != nil && initialHoldtime != nil && maxHoldtime != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			fmt.Sprint("timers throttle spf ", *delay, " ", *initialHoldtime, " ", *maxHoldtime))
	} else if configRunning.lookup([]string{"protocols", "ospf"}) != nil {
		quaggaVtysh("configure terminal",
			"router ospf",
			"no timers throttle spf")
	}
	return cmd.Success
}

/*
	tag:
	type: txt
	help: OSPFv3 Area
	syntax:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --check-ospf-area $VAR(@)"; "Invalid OSFPv3 area \"$VAR(@)\" "
	val_help: u32; OSPFv3 area in decimal notation
	val_help: ipv4; OSPFv3 area in dotted decimal notation
*/
func quaggaProtocolsOspfv3Area(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 area WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Name of export-list
	create:expression: "vtysh -c \"configure terminal\" \
	        -c \"router ospf6 \" \
	        -c \"area $VAR(../@) export-list $VAR(@) \"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	        -c \"router ospf6 \" \
	        -c \"no area $VAR(../@) export-list $VAR(@) \"; "
*/
func quaggaProtocolsOspfv3AreaExportList(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 area WORD export-list WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Name of import-list
	create:expression: "vtysh -c \"configure terminal\" \
	        -c \"router ospf6 \" \
	        -c \"area $VAR(../@) import-list $VAR(@) \"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	        -c \"router ospf6 \" \
	        -c \"no area $VAR(../@) import-list $VAR(@) \"; "
*/
func quaggaProtocolsOspfv3AreaImportList(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 area WORD import-list WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	multi:
	type: txt
	help: OSPFv3 area interface

	create: vtysh -c "configure terminal" -c "router ospf6" \
	          -c "interface $VAR(@) area $VAR(../@)"

	delete: vtysh -c "configure terminal" -c "router ospf6" \
	           -c "no interface $VAR(@) area $VAR(../@)"

	allowed: ${vyatta_sbindir}/vyatta-interfaces.pl --show all
*/
func quaggaProtocolsOspfv3AreaInterface(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 area WORD interface WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	tag:
	type: ipv6net
	help: Specify IPv6 prefix (border routers only)
	syntax:expression: exec "${vyatta_sbindir}/check_prefix_boundary $VAR(@)"

	delete: touch /tmp/ospf6-range.$PPID

	end: if [ -f /tmp/ospf6-range.$PPID ]; then
	        vtysh -c "configure terminal" \
	          -c "router ospf6"                                        \
	          -c "no area $VAR(../@) range $VAR(@)";
	        rm /tmp/ospf6-range.$PPID;
	     else
	        vtysh --noerror -c "configure terminal" \
	          -c "router ospf6"                                               \
	          -c "no area $VAR(../@) range $VAR(@)";
	        vtysh -c "configure terminal" \
	          -c "router ospf6"                                        \
	          -c "area $VAR(../@) range $VAR(@)";
	     fi;
*/
func quaggaProtocolsOspfv3AreaRange(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 area WORD range X:X::X:X/M
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Advertise this range
	create:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf6\" \
	       -c \"area $VAR(../../@) range $VAR(../@) advertise\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf6\" \
	       -c \"no area $VAR(../../@) range $VAR(../@) advertise\"; "
*/
func quaggaProtocolsOspfv3AreaRangeAdvertise(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 area WORD range X:X::X:X/M advertise
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Do not advertise this range
	create:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf6\" \
	       -c \"area $VAR(../../@) range $VAR(../@) not-advertise\"; "
	delete:expression: "vtysh -c \"configure terminal\" \
	       -c \"router ospf6\" \
	       -c \"no area $VAR(../../@) range $VAR(../@) not-advertise\"; "
*/
func quaggaProtocolsOspfv3AreaRangeNotAdvertise(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 area WORD range X:X::X:X/M not-advertise
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	priority: 640
	help: IPv6 Open Shortest Path First protocol (OSPFv3) parameters
	begin: if [ "$COMMIT_ACTION" != DELETE ]; then
	         if [ -n "$VAR(parameters/router-id/@)" ]; then
	           vtysh -c "configure terminal" -c "router ospf6" \
	                 -c "router-id $VAR(parameters/router-id/@)"
	         else
	           vtysh -c "configure terminal" -c "router ospf6" \
	                 -c "no router-id"
	         fi
	         vtysh -d ospf6d -c 'sh run' > /opt/vyatta/etc/quagga/ospf6d.conf
	       fi
	end: if [ "$COMMIT_ACTION" == DELETE ]; then
	       vtysh -c "configure terminal" -c "router ospf6" -c "no router-id"
	       vtysh -c "configure terminal" -c "no router ospf6"
	     fi
*/
func quaggaProtocolsOspfv3(Cmd int, Args cmd.Args) int {
	//protocols ospfv3
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: OSPFv3 specific parameters
*/
func quaggaProtocolsOspfv3Parameters(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 parameters
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: ipv4
	help: Router identifier
*/
func quaggaProtocolsOspfv3ParametersRouterId(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 parameters router-id A.B.C.D
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Redistribute bgp routes

	end: vtysh -c "configure terminal" \
	       -c "router ospf6"                                       \
	       -c "no redistribute bgp";
	     if [ "$COMMIT_ACTION" = "SET" -o "$COMMIT_ACTION" = "ACTIVE" ]; then
	        if [ -n "$VAR(./route-map/@)" ]; then
	          COND="route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	          -c "router ospf6"                                         \
	          -c "redistribute bgp $COND";
	     fi;
*/
func quaggaProtocolsOspfv3RedistributeBgp(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute bgp
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"

*/
func quaggaProtocolsOspfv3RedistributeBgpRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute bgp route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Redistribute connected routes

	end: vtysh -c "configure terminal" \
	       -c "router ospf6"                                       \
	       -c "no redistribute connected";
	     if [ "$COMMIT_ACTION" = "SET" -o "$COMMIT_ACTION" = "ACTIVE" ]; then
	        if [ -n "$VAR(./route-map/@)" ]; then
	          COND="route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	          -c "router ospf6"                                         \
	          -c "redistribute connected $COND";
	     fi;
*/
func quaggaProtocolsOspfv3RedistributeConnected(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute connected
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsOspfv3RedistributeConnectedRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute connected route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Redistribute kernel routes

	end: vtysh -c "configure terminal" \
	       -c "router ospf6"                                       \
	       -c "no redistribute kernel";
	     if [ "$COMMIT_ACTION" = "SET" -o "$COMMIT_ACTION" = "ACTIVE" ]; then
	        if [ -n "$VAR(./route-map/@)" ]; then
	          COND="route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	          -c "router ospf6"                                         \
	          -c "redistribute kernel $COND";
	     fi;
*/
func quaggaProtocolsOspfv3RedistributeKernel(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute kernel
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsOspfv3RedistributeKernelRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute kernel route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Redistribute information from another routing protocol
*/
func quaggaProtocolsOspfv3Redistribute(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Redistribute RIPNG routes

	end: vtysh -c "configure terminal" \
	       -c "router ospf6"                                       \
	       -c "no redistribute ripng";
	     if [ "$COMMIT_ACTION" = "SET" -o "$COMMIT_ACTION" = "ACTIVE" ]; then
	        if [ -n "$VAR(./route-map/@)" ]; then
	          COND="route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	          -c "router ospf6"                                         \
	          -c "redistribute ripng $COND";
	     fi;
*/
func quaggaProtocolsOspfv3RedistributeRipng(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute ripng
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsOspfv3RedistributeRipngRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute ripng route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	help: Redistribute static routes

	end: vtysh -c "configure terminal" \
	       -c "router ospf6"                                       \
	       -c "no redistribute static";
	     if [ "$COMMIT_ACTION" = "SET" -o "$COMMIT_ACTION" = "ACTIVE" ]; then
	        if [ -n "$VAR(./route-map/@)" ]; then
	          COND="route-map $VAR(./route-map/@)";
	        fi;
	        vtysh -c "configure terminal" \
	          -c "router ospf6"                                         \
	          -c "redistribute static $COND";
	     fi;
*/
func quaggaProtocolsOspfv3RedistributeStatic(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute static
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

/*
	type: txt
	help: Route map reference
	commit:expression: exec "/opt/vyatta/sbin/vyatta_quagga_utils.pl --exists \"policy route-map $VAR(@)\" ";"route-map $VAR(@) doesn't exist"
*/
func quaggaProtocolsOspfv3RedistributeStaticRouteMap(Cmd int, Args cmd.Args) int {
	//protocols ospfv3 redistribute static route-map WORD
	switch Cmd {
	case cmd.Set:
	case cmd.Delete:
	}
	return cmd.Success
}

func initConfig() {
	configParser = cmd.NewParser()
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication", "md5", "key-id", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5KeyId)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication", "md5", "key-id", "WORD", "md5-key", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5KeyIdMd5Key)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication", "md5"},
		quaggaInterfacesInterfaceIpv4OspfAuthenticationMd5)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication"},
		quaggaInterfacesInterfaceIpv4OspfAuthentication)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "authentication", "plaintext-password", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfAuthenticationPlaintextPassword)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "bandwidth", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfBandwidth)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "cost", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfCost)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "dead-interval", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfDeadInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "hello-interval", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfHelloInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "mtu-ignore"},
		quaggaInterfacesInterfaceIpv4OspfMtuIgnore)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "network", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfNetwork)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf"},
		quaggaInterfacesInterfaceIpv4Ospf)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "priority", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfPriority)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "retransmit-interval", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfRetransmitInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv4", "ospf", "transmit-delay", "WORD"},
		quaggaInterfacesInterfaceIpv4OspfTransmitDelay)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "cost", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3Cost)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "dead-interval", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3DeadInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "hello-interval", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3HelloInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "ifmtu", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3Ifmtu)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "instance-id", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3InstanceId)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "mtu-ignore"},
		quaggaInterfacesInterfaceIpv6Ospfv3MtuIgnore)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3"},
		quaggaInterfacesInterfaceIpv6Ospfv3)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "passive"},
		quaggaInterfacesInterfaceIpv6Ospfv3Passive)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "priority", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3Priority)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "retransmit-interval", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3RetransmitInterval)
	configParser.InstallCmd(
		[]string{"interfaces", "interface", "WORD", "ipv6", "ospfv3", "transmit-delay", "WORD"},
		quaggaInterfacesInterfaceIpv6Ospfv3TransmitDelay)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD"},
		quaggaPolicyAccessList)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "description", "WORD"},
		quaggaPolicyAccessListDescription)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD"},
		quaggaPolicyAccessListRule)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "action", "WORD"},
		quaggaPolicyAccessListRuleAction)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "description", "WORD"},
		quaggaPolicyAccessListRuleDescription)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "destination", "any"},
		quaggaPolicyAccessListRuleDestinationAny)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "destination", "host", "A.B.C.D"},
		quaggaPolicyAccessListRuleDestinationHost)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "destination", "inverse-mask", "A.B.C.D"},
		quaggaPolicyAccessListRuleDestinationInverseMask)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "destination", "network", "A.B.C.D"},
		quaggaPolicyAccessListRuleDestinationNetwork)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "destination"},
		quaggaPolicyAccessListRuleDestination)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "source", "any"},
		quaggaPolicyAccessListRuleSourceAny)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "source", "host", "A.B.C.D"},
		quaggaPolicyAccessListRuleSourceHost)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "source", "inverse-mask", "A.B.C.D"},
		quaggaPolicyAccessListRuleSourceInverseMask)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "source", "network", "A.B.C.D"},
		quaggaPolicyAccessListRuleSourceNetwork)
	configParser.InstallCmd(
		[]string{"policy", "access-list", "WORD", "rule", "WORD", "source"},
		quaggaPolicyAccessListRuleSource)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD"},
		quaggaPolicyAccessList6)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD", "description", "WORD"},
		quaggaPolicyAccessList6Description)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD", "rule", "WORD"},
		quaggaPolicyAccessList6Rule)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD", "rule", "WORD", "action", "WORD"},
		quaggaPolicyAccessList6RuleAction)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD", "rule", "WORD", "description", "WORD"},
		quaggaPolicyAccessList6RuleDescription)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD", "rule", "WORD", "source", "any"},
		quaggaPolicyAccessList6RuleSourceAny)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD", "rule", "WORD", "source", "exact-match"},
		quaggaPolicyAccessList6RuleSourceExactMatch)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD", "rule", "WORD", "source", "network", "X:X::X:X/M"},
		quaggaPolicyAccessList6RuleSourceNetwork)
	configParser.InstallCmd(
		[]string{"policy", "access-list6", "WORD", "rule", "WORD", "source"},
		quaggaPolicyAccessList6RuleSource)
	configParser.InstallCmd(
		[]string{"policy", "as-path-list", "WORD"},
		quaggaPolicyAsPathList)
	configParser.InstallCmd(
		[]string{"policy", "as-path-list", "WORD", "description", "WORD"},
		quaggaPolicyAsPathListDescription)
	configParser.InstallCmd(
		[]string{"policy", "as-path-list", "WORD", "rule", "WORD"},
		quaggaPolicyAsPathListRule)
	configParser.InstallCmd(
		[]string{"policy", "as-path-list", "WORD", "rule", "WORD", "action", "WORD"},
		quaggaPolicyAsPathListRuleAction)
	configParser.InstallCmd(
		[]string{"policy", "as-path-list", "WORD", "rule", "WORD", "description", "WORD"},
		quaggaPolicyAsPathListRuleDescription)
	configParser.InstallCmd(
		[]string{"policy", "as-path-list", "WORD", "rule", "WORD", "regex", "WORD"},
		quaggaPolicyAsPathListRuleRegex)
	configParser.InstallCmd(
		[]string{"policy", "community-list", "WORD"},
		quaggaPolicyCommunityList)
	configParser.InstallCmd(
		[]string{"policy", "community-list", "WORD", "description", "WORD"},
		quaggaPolicyCommunityListDescription)
	configParser.InstallCmd(
		[]string{"policy", "community-list", "WORD", "rule", "WORD"},
		quaggaPolicyCommunityListRule)
	configParser.InstallCmd(
		[]string{"policy", "community-list", "WORD", "rule", "WORD", "action", "WORD"},
		quaggaPolicyCommunityListRuleAction)
	configParser.InstallCmd(
		[]string{"policy", "community-list", "WORD", "rule", "WORD", "description", "WORD"},
		quaggaPolicyCommunityListRuleDescription)
	configParser.InstallCmd(
		[]string{"policy", "community-list", "WORD", "rule", "WORD", "regex", "WORD"},
		quaggaPolicyCommunityListRuleRegex)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list", "WORD"},
		quaggaPolicyPrefixList)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list", "WORD", "description", "WORD"},
		quaggaPolicyPrefixListDescription)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list", "WORD", "rule", "WORD"},
		quaggaPolicyPrefixListRule)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list", "WORD", "rule", "WORD", "action", "WORD"},
		quaggaPolicyPrefixListRuleAction)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list", "WORD", "rule", "WORD", "description", "WORD"},
		quaggaPolicyPrefixListRuleDescription)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list", "WORD", "rule", "WORD", "ge", "WORD"},
		quaggaPolicyPrefixListRuleGe)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list", "WORD", "rule", "WORD", "le", "WORD"},
		quaggaPolicyPrefixListRuleLe)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list", "WORD", "rule", "WORD", "prefix", "A.B.C.D/M"},
		quaggaPolicyPrefixListRulePrefix)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list6", "WORD"},
		quaggaPolicyPrefixList6)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list6", "WORD", "description", "WORD"},
		quaggaPolicyPrefixList6Description)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list6", "WORD", "rule", "WORD"},
		quaggaPolicyPrefixList6Rule)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list6", "WORD", "rule", "WORD", "action", "WORD"},
		quaggaPolicyPrefixList6RuleAction)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list6", "WORD", "rule", "WORD", "description", "WORD"},
		quaggaPolicyPrefixList6RuleDescription)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list6", "WORD", "rule", "WORD", "ge", "WORD"},
		quaggaPolicyPrefixList6RuleGe)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list6", "WORD", "rule", "WORD", "le", "WORD"},
		quaggaPolicyPrefixList6RuleLe)
	configParser.InstallCmd(
		[]string{"policy", "prefix-list6", "WORD", "rule", "WORD", "prefix", "X:X::X:X/M"},
		quaggaPolicyPrefixList6RulePrefix)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD"},
		quaggaPolicyRouteMap)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "description", "WORD"},
		quaggaPolicyRouteMapDescription)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD"},
		quaggaPolicyRouteMapRule)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "action", "WORD"},
		quaggaPolicyRouteMapRuleAction)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "call", "WORD"},
		quaggaPolicyRouteMapRuleCall)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "continue", "WORD"},
		quaggaPolicyRouteMapRuleContinue)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "description", "WORD"},
		quaggaPolicyRouteMapRuleDescription)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "as-path", "WORD"},
		quaggaPolicyRouteMapRuleMatchAsPath)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "community", "community-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchCommunityCommunityList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "community", "exact-match"},
		quaggaPolicyRouteMapRuleMatchCommunityExactMatch)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "community"},
		quaggaPolicyRouteMapRuleMatchCommunity)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "interface", "WORD"},
		quaggaPolicyRouteMapRuleMatchInterface)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "address", "access-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpAddressAccessList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "address"},
		quaggaPolicyRouteMapRuleMatchIpAddress)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "address", "prefix-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpAddressPrefixList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "nexthop", "access-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpNexthopAccessList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "nexthop"},
		quaggaPolicyRouteMapRuleMatchIpNexthop)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "nexthop", "prefix-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpNexthopPrefixList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip"},
		quaggaPolicyRouteMapRuleMatchIp)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "route-source", "access-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpRouteSourceAccessList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "route-source"},
		quaggaPolicyRouteMapRuleMatchIpRouteSource)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ip", "route-source", "prefix-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpRouteSourcePrefixList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ipv6", "address", "access-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpv6AddressAccessList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ipv6", "address"},
		quaggaPolicyRouteMapRuleMatchIpv6Address)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ipv6", "address", "prefix-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpv6AddressPrefixList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ipv6", "nexthop", "access-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpv6NexthopAccessList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ipv6", "nexthop"},
		quaggaPolicyRouteMapRuleMatchIpv6Nexthop)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ipv6", "nexthop", "prefix-list", "WORD"},
		quaggaPolicyRouteMapRuleMatchIpv6NexthopPrefixList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "ipv6"},
		quaggaPolicyRouteMapRuleMatchIpv6)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "metric", "WORD"},
		quaggaPolicyRouteMapRuleMatchMetric)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match"},
		quaggaPolicyRouteMapRuleMatch)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "origin", "WORD"},
		quaggaPolicyRouteMapRuleMatchOrigin)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "peer", "WORD"},
		quaggaPolicyRouteMapRuleMatchPeer)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "match", "tag", "WORD"},
		quaggaPolicyRouteMapRuleMatchTag)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "on-match", "goto", "WORD"},
		quaggaPolicyRouteMapRuleOnMatchGoto)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "on-match", "next"},
		quaggaPolicyRouteMapRuleOnMatchNext)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "on-match"},
		quaggaPolicyRouteMapRuleOnMatch)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "aggregator", "as", "WORD"},
		quaggaPolicyRouteMapRuleSetAggregatorAs)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "aggregator", "ip", "A.B.C.D"},
		quaggaPolicyRouteMapRuleSetAggregatorIp)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "aggregator"},
		quaggaPolicyRouteMapRuleSetAggregator)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "as-path-prepend", "WORD"},
		quaggaPolicyRouteMapRuleSetAsPathPrepend)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "atomic-aggregate"},
		quaggaPolicyRouteMapRuleSetAtomicAggregate)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "comm-list", "comm-list", "WORD"},
		quaggaPolicyRouteMapRuleSetCommListCommList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "comm-list", "delete"},
		quaggaPolicyRouteMapRuleSetCommListDelete)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "comm-list"},
		quaggaPolicyRouteMapRuleSetCommList)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "community", "WORD"},
		quaggaPolicyRouteMapRuleSetCommunity)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "ip-next-hop", "A.B.C.D"},
		quaggaPolicyRouteMapRuleSetIpNextHop)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "ipv6-next-hop", "global", "X:X::X:X"},
		quaggaPolicyRouteMapRuleSetIpv6NextHopGlobal)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "ipv6-next-hop", "local", "X:X::X:X"},
		quaggaPolicyRouteMapRuleSetIpv6NextHopLocal)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "ipv6-next-hop"},
		quaggaPolicyRouteMapRuleSetIpv6NextHop)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "local-preference", "WORD"},
		quaggaPolicyRouteMapRuleSetLocalPreference)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "metric-type", "WORD"},
		quaggaPolicyRouteMapRuleSetMetricType)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "metric", "WORD"},
		quaggaPolicyRouteMapRuleSetMetric)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set"},
		quaggaPolicyRouteMapRuleSet)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "origin", "WORD"},
		quaggaPolicyRouteMapRuleSetOrigin)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "originator-id", "A.B.C.D"},
		quaggaPolicyRouteMapRuleSetOriginatorId)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "tag", "WORD"},
		quaggaPolicyRouteMapRuleSetTag)
	configParser.InstallCmd(
		[]string{"policy", "route-map", "WORD", "rule", "WORD", "set", "weight", "WORD"},
		quaggaPolicyRouteMapRuleSetWeight)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD"},
		quaggaProtocolsBgp)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "aggregate-address", "X:X::X:X/M"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastAggregateAddress)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "aggregate-address", "X:X::X:X/M", "summary-only"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastAggregateAddressSummaryOnly)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "network", "X:X::X:X/M"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastNetwork)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "network", "X:X::X:X/M", "path-limit", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastNetworkPathLimit)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "network", "X:X::X:X/M", "route-map", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastNetworkRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast"},
		quaggaProtocolsBgpAddressFamilyIpv6Unicast)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "connected", "metric", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeConnectedMetric)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "connected"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeConnected)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "connected", "route-map", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeConnectedRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "kernel", "metric", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeKernelMetric)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "kernel"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeKernel)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "kernel", "route-map", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeKernelRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistribute)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "ospfv3", "metric", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeOspfv3Metric)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "ospfv3"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeOspfv3)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "ospfv3", "route-map", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeOspfv3RouteMap)
	/*
		configParser.InstallCmd(
			[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "ripng", "metric", "WORD"},
			quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeRipngMetric)
		configParser.InstallCmd(
			[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "ripng"},
			quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeRipng)
		configParser.InstallCmd(
			[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "ripng", "route-map", "WORD"},
			quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeRipngRouteMap)
	*/
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "static", "metric", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeStaticMetric)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "static"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeStatic)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family", "ipv6-unicast", "redistribute", "static", "route-map", "WORD"},
		quaggaProtocolsBgpAddressFamilyIpv6UnicastRedistributeStaticRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "address-family"},
		quaggaProtocolsBgpAddressFamily)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "aggregate-address", "A.B.C.D/M"},
		quaggaProtocolsBgpAggregateAddress)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "aggregate-address", "A.B.C.D/M", "as-set"},
		quaggaProtocolsBgpAggregateAddressAsSet)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "aggregate-address", "A.B.C.D/M", "summary-only"},
		quaggaProtocolsBgpAggregateAddressSummaryOnly)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "maximum-paths", "ebgp", "WORD"},
		quaggaProtocolsBgpMaximumPathsEbgp)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "maximum-paths", "ibgp", "WORD"},
		quaggaProtocolsBgpMaximumPathsIbgp)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "maximum-paths"},
		quaggaProtocolsBgpMaximumPaths)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD"},
		quaggaProtocolsBgpNeighbor)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "allowas-in"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAllowasIn)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "allowas-in", "number", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAllowasInNumber)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "attribute-unchanged", "as-path"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAttributeUnchangedAsPath)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "attribute-unchanged", "med"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAttributeUnchangedMed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "attribute-unchanged", "next-hop"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAttributeUnchangedNextHop)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "attribute-unchanged"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastAttributeUnchanged)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "capability", "dynamic"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityDynamic)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "capability"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapability)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "capability", "orf"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityOrf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "capability", "orf", "prefix-list"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityOrfPrefixList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "capability", "orf", "prefix-list", "receive"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityOrfPrefixListReceive)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "capability", "orf", "prefix-list", "send"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastCapabilityOrfPrefixListSend)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "default-originate"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDefaultOriginate)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "default-originate", "route-map", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDefaultOriginateRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "disable-send-community", "extended"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDisableSendCommunityExtended)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "disable-send-community"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDisableSendCommunity)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "disable-send-community", "standard"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDisableSendCommunityStandard)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "distribute-list", "export", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDistributeListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "distribute-list", "import", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDistributeListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "distribute-list"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastDistributeList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "filter-list", "export", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastFilterListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "filter-list", "import", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastFilterListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "filter-list"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastFilterList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "maximum-prefix", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastMaximumPrefix)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "nexthop-local"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastNexthopLocal)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "nexthop-local", "unchanged"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastNexthopLocalUnchanged)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "nexthop-self"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastNexthopSelf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6Unicast)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "peer-group", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastPeerGroup)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "prefix-list", "export", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastPrefixListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "prefix-list", "import", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastPrefixListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "prefix-list"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastPrefixList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "remove-private-as"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRemovePrivateAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "route-map", "export", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteMapExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "route-map", "import", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteMapImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "route-map"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "route-reflector-client"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteReflectorClient)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "route-server-client"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastRouteServerClient)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "soft-reconfiguration", "inbound"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastSoftReconfigurationInbound)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "soft-reconfiguration"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastSoftReconfiguration)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family", "ipv6-unicast", "unsuppress-map", "WORD"},
		quaggaProtocolsBgpNeighborAddressFamilyIpv6UnicastUnsuppressMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "address-family"},
		quaggaProtocolsBgpNeighborAddressFamily)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "advertisement-interval", "WORD"},
		quaggaProtocolsBgpNeighborAdvertisementInterval)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "allowas-in"},
		quaggaProtocolsBgpNeighborAllowasIn)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "allowas-in", "number", "WORD"},
		quaggaProtocolsBgpNeighborAllowasInNumber)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "attribute-unchanged", "as-path"},
		quaggaProtocolsBgpNeighborAttributeUnchangedAsPath)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "attribute-unchanged", "med"},
		quaggaProtocolsBgpNeighborAttributeUnchangedMed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "attribute-unchanged", "next-hop"},
		quaggaProtocolsBgpNeighborAttributeUnchangedNextHop)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "attribute-unchanged"},
		quaggaProtocolsBgpNeighborAttributeUnchanged)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "capability", "dynamic"},
		quaggaProtocolsBgpNeighborCapabilityDynamic)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "capability"},
		quaggaProtocolsBgpNeighborCapability)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "capability", "orf"},
		quaggaProtocolsBgpNeighborCapabilityOrf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "capability", "orf", "prefix-list"},
		quaggaProtocolsBgpNeighborCapabilityOrfPrefixList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "capability", "orf", "prefix-list", "receive"},
		quaggaProtocolsBgpNeighborCapabilityOrfPrefixListReceive)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "capability", "orf", "prefix-list", "send"},
		quaggaProtocolsBgpNeighborCapabilityOrfPrefixListSend)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "default-originate"},
		quaggaProtocolsBgpNeighborDefaultOriginate)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "default-originate", "route-map", "WORD"},
		quaggaProtocolsBgpNeighborDefaultOriginateRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "description", "WORD"},
		quaggaProtocolsBgpNeighborDescription)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "disable-capability-negotiation"},
		quaggaProtocolsBgpNeighborDisableCapabilityNegotiation)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "disable-connected-check"},
		quaggaProtocolsBgpNeighborDisableConnectedCheck)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "disable-send-community", "extended"},
		quaggaProtocolsBgpNeighborDisableSendCommunityExtended)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "disable-send-community"},
		quaggaProtocolsBgpNeighborDisableSendCommunity)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "disable-send-community", "standard"},
		quaggaProtocolsBgpNeighborDisableSendCommunityStandard)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "distribute-list", "export", "WORD"},
		quaggaProtocolsBgpNeighborDistributeListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "distribute-list", "import", "WORD"},
		quaggaProtocolsBgpNeighborDistributeListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "distribute-list"},
		quaggaProtocolsBgpNeighborDistributeList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "ebgp-multihop", "WORD"},
		quaggaProtocolsBgpNeighborEbgpMultihop)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "filter-list", "export", "WORD"},
		quaggaProtocolsBgpNeighborFilterListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "filter-list", "import", "WORD"},
		quaggaProtocolsBgpNeighborFilterListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "filter-list"},
		quaggaProtocolsBgpNeighborFilterList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "local-as", "WORD"},
		quaggaProtocolsBgpNeighborLocalAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "local-as", "WORD", "no-prepend"},
		quaggaProtocolsBgpNeighborLocalAsNoPrepend)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "maximum-prefix", "WORD"},
		quaggaProtocolsBgpNeighborMaximumPrefix)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "nexthop-self"},
		quaggaProtocolsBgpNeighborNexthopSelf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "override-capability"},
		quaggaProtocolsBgpNeighborOverrideCapability)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "passive"},
		quaggaProtocolsBgpNeighborPassive)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "password", "WORD"},
		quaggaProtocolsBgpNeighborPassword)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "peer-group", "WORD"},
		quaggaProtocolsBgpNeighborPeerGroup)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "port", "WORD"},
		quaggaProtocolsBgpNeighborPort)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "prefix-list", "export", "WORD"},
		quaggaProtocolsBgpNeighborPrefixListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "prefix-list", "import", "WORD"},
		quaggaProtocolsBgpNeighborPrefixListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "prefix-list"},
		quaggaProtocolsBgpNeighborPrefixList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "remote-as", "WORD"},
		quaggaProtocolsBgpNeighborRemoteAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "remove-private-as"},
		quaggaProtocolsBgpNeighborRemovePrivateAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "route-map", "export", "WORD"},
		quaggaProtocolsBgpNeighborRouteMapExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "route-map", "import", "WORD"},
		quaggaProtocolsBgpNeighborRouteMapImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "route-map"},
		quaggaProtocolsBgpNeighborRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "route-reflector-client"},
		quaggaProtocolsBgpNeighborRouteReflectorClient)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "route-server-client"},
		quaggaProtocolsBgpNeighborRouteServerClient)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "shutdown"},
		quaggaProtocolsBgpNeighborShutdown)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "soft-reconfiguration", "inbound"},
		quaggaProtocolsBgpNeighborSoftReconfigurationInbound)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "soft-reconfiguration"},
		quaggaProtocolsBgpNeighborSoftReconfiguration)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "strict-capability-match"},
		quaggaProtocolsBgpNeighborStrictCapabilityMatch)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "timers", "connect", "WORD"},
		quaggaProtocolsBgpNeighborTimersConnect)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "timers", "holdtime", "WORD"},
		quaggaProtocolsBgpNeighborTimersHoldtime)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "timers", "keepalive", "WORD"},
		quaggaProtocolsBgpNeighborTimersKeepalive)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "timers"},
		quaggaProtocolsBgpNeighborTimers)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "ttl-security", "hops", "WORD"},
		quaggaProtocolsBgpNeighborTtlSecurityHops)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "ttl-security"},
		quaggaProtocolsBgpNeighborTtlSecurity)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "unsuppress-map", "WORD"},
		quaggaProtocolsBgpNeighborUnsuppressMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "update-source", "WORD"},
		quaggaProtocolsBgpNeighborUpdateSource)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "neighbor", "WORD", "weight", "WORD"},
		quaggaProtocolsBgpNeighborWeight)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "network", "A.B.C.D/M"},
		quaggaProtocolsBgpNetwork)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "network", "A.B.C.D/M", "backdoor"},
		quaggaProtocolsBgpNetworkBackdoor)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "network", "A.B.C.D/M", "route-map", "WORD"},
		quaggaProtocolsBgpNetworkRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "always-compare-med"},
		quaggaProtocolsBgpParametersAlwaysCompareMed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "bestpath", "as-path", "confed"},
		quaggaProtocolsBgpParametersBestpathAsPathConfed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "bestpath", "as-path", "ignore"},
		quaggaProtocolsBgpParametersBestpathAsPathIgnore)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "bestpath", "as-path"},
		quaggaProtocolsBgpParametersBestpathAsPath)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "bestpath", "compare-routerid"},
		quaggaProtocolsBgpParametersBestpathCompareRouterid)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "bestpath", "med", "confed"},
		quaggaProtocolsBgpParametersBestpathMedConfed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "bestpath", "med", "missing-as-worst"},
		quaggaProtocolsBgpParametersBestpathMedMissingAsWorst)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "bestpath", "med"},
		quaggaProtocolsBgpParametersBestpathMed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "bestpath"},
		quaggaProtocolsBgpParametersBestpath)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "cluster-id", "A.B.C.D"},
		quaggaProtocolsBgpParametersClusterId)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "confederation", "identifier", "WORD"},
		quaggaProtocolsBgpParametersConfederationIdentifier)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "confederation"},
		quaggaProtocolsBgpParametersConfederation)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "confederation", "peers", "WORD"},
		quaggaProtocolsBgpParametersConfederationPeers)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "dampening", "half-life", "WORD"},
		quaggaProtocolsBgpParametersDampeningHalfLife)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "dampening", "max-suppress-time", "WORD"},
		quaggaProtocolsBgpParametersDampeningMaxSuppressTime)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "dampening"},
		quaggaProtocolsBgpParametersDampening)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "dampening", "re-use", "WORD"},
		quaggaProtocolsBgpParametersDampeningReUse)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "dampening", "start-suppress-time", "WORD"},
		quaggaProtocolsBgpParametersDampeningStartSuppressTime)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "default", "local-pref", "WORD"},
		quaggaProtocolsBgpParametersDefaultLocalPref)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "default", "no-ipv4-unicast"},
		quaggaProtocolsBgpParametersDefaultNoIpv4Unicast)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "default"},
		quaggaProtocolsBgpParametersDefault)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "deterministic-med"},
		quaggaProtocolsBgpParametersDeterministicMed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "disable-network-import-check"},
		quaggaProtocolsBgpParametersDisableNetworkImportCheck)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "distance", "global", "external", "WORD"},
		quaggaProtocolsBgpParametersDistanceGlobalExternal)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "distance", "global", "internal", "WORD"},
		quaggaProtocolsBgpParametersDistanceGlobalInternal)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "distance", "global", "local", "WORD"},
		quaggaProtocolsBgpParametersDistanceGlobalLocal)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "distance", "global"},
		quaggaProtocolsBgpParametersDistanceGlobal)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "distance"},
		quaggaProtocolsBgpParametersDistance)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "distance", "prefix", "A.B.C.D/M"},
		quaggaProtocolsBgpParametersDistancePrefix)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "distance", "prefix", "A.B.C.D/M", "distance", "WORD"},
		quaggaProtocolsBgpParametersDistancePrefixDistance)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "enforce-first-as"},
		quaggaProtocolsBgpParametersEnforceFirstAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "graceful-restart"},
		quaggaProtocolsBgpParametersGracefulRestart)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "graceful-restart", "stalepath-time", "WORD"},
		quaggaProtocolsBgpParametersGracefulRestartStalepathTime)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "log-neighbor-changes"},
		quaggaProtocolsBgpParametersLogNeighborChanges)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "no-client-to-client-reflection"},
		quaggaProtocolsBgpParametersNoClientToClientReflection)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "no-fast-external-failover"},
		quaggaProtocolsBgpParametersNoFastExternalFailover)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters"},
		quaggaProtocolsBgpParameters)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "router-id", "A.B.C.D"},
		quaggaProtocolsBgpParametersRouterId)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "parameters", "scan-time", "WORD"},
		quaggaProtocolsBgpParametersScanTime)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD"},
		quaggaProtocolsBgpPeerGroup)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "allowas-in"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAllowasIn)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "allowas-in", "number", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAllowasInNumber)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "attribute-unchanged", "as-path"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAttributeUnchangedAsPath)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "attribute-unchanged", "med"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAttributeUnchangedMed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "attribute-unchanged", "next-hop"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAttributeUnchangedNextHop)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "attribute-unchanged"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastAttributeUnchanged)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "capability", "dynamic"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityDynamic)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "capability"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapability)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "capability", "orf"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityOrf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "capability", "orf", "prefix-list"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityOrfPrefixList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "capability", "orf", "prefix-list", "receive"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityOrfPrefixListReceive)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "capability", "orf", "prefix-list", "send"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastCapabilityOrfPrefixListSend)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "default-originate"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDefaultOriginate)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "default-originate", "route-map", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDefaultOriginateRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "disable-send-community", "extended"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDisableSendCommunityExtended)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "disable-send-community"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDisableSendCommunity)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "disable-send-community", "standard"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDisableSendCommunityStandard)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "distribute-list", "export", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDistributeListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "distribute-list", "import", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDistributeListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "distribute-list"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastDistributeList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "filter-list", "export", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastFilterListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "filter-list", "import", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastFilterListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "filter-list"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastFilterList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "maximum-prefix", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastMaximumPrefix)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "nexthop-local"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastNexthopLocal)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "nexthop-local", "unchanged"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastNexthopLocalUnchanged)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "nexthop-self"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastNexthopSelf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6Unicast)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "prefix-list", "export", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastPrefixListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "prefix-list", "import", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastPrefixListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "prefix-list"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastPrefixList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "remove-private-as"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRemovePrivateAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "route-map", "export", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteMapExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "route-map", "import", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteMapImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "route-map"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "route-reflector-client"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteReflectorClient)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "route-server-client"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastRouteServerClient)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "soft-reconfiguration", "inbound"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastSoftReconfigurationInbound)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "soft-reconfiguration"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastSoftReconfiguration)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family", "ipv6-unicast", "unsuppress-map", "WORD"},
		quaggaProtocolsBgpPeerGroupAddressFamilyIpv6UnicastUnsuppressMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "address-family"},
		quaggaProtocolsBgpPeerGroupAddressFamily)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "allowas-in"},
		quaggaProtocolsBgpPeerGroupAllowasIn)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "allowas-in", "number", "WORD"},
		quaggaProtocolsBgpPeerGroupAllowasInNumber)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "attribute-unchanged", "as-path"},
		quaggaProtocolsBgpPeerGroupAttributeUnchangedAsPath)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "attribute-unchanged", "med"},
		quaggaProtocolsBgpPeerGroupAttributeUnchangedMed)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "attribute-unchanged", "next-hop"},
		quaggaProtocolsBgpPeerGroupAttributeUnchangedNextHop)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "attribute-unchanged"},
		quaggaProtocolsBgpPeerGroupAttributeUnchanged)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "capability", "dynamic"},
		quaggaProtocolsBgpPeerGroupCapabilityDynamic)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "capability"},
		quaggaProtocolsBgpPeerGroupCapability)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "capability", "orf"},
		quaggaProtocolsBgpPeerGroupCapabilityOrf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "capability", "orf", "prefix-list"},
		quaggaProtocolsBgpPeerGroupCapabilityOrfPrefixList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "capability", "orf", "prefix-list", "receive"},
		quaggaProtocolsBgpPeerGroupCapabilityOrfPrefixListReceive)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "capability", "orf", "prefix-list", "send"},
		quaggaProtocolsBgpPeerGroupCapabilityOrfPrefixListSend)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "default-originate"},
		quaggaProtocolsBgpPeerGroupDefaultOriginate)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "default-originate", "route-map", "WORD"},
		quaggaProtocolsBgpPeerGroupDefaultOriginateRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "description", "WORD"},
		quaggaProtocolsBgpPeerGroupDescription)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "disable-capability-negotiation"},
		quaggaProtocolsBgpPeerGroupDisableCapabilityNegotiation)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "disable-connected-check"},
		quaggaProtocolsBgpPeerGroupDisableConnectedCheck)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "disable-send-community", "extended"},
		quaggaProtocolsBgpPeerGroupDisableSendCommunityExtended)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "disable-send-community"},
		quaggaProtocolsBgpPeerGroupDisableSendCommunity)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "disable-send-community", "standard"},
		quaggaProtocolsBgpPeerGroupDisableSendCommunityStandard)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "distribute-list", "export", "WORD"},
		quaggaProtocolsBgpPeerGroupDistributeListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "distribute-list", "import", "WORD"},
		quaggaProtocolsBgpPeerGroupDistributeListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "distribute-list"},
		quaggaProtocolsBgpPeerGroupDistributeList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "ebgp-multihop", "WORD"},
		quaggaProtocolsBgpPeerGroupEbgpMultihop)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "filter-list", "export", "WORD"},
		quaggaProtocolsBgpPeerGroupFilterListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "filter-list", "import", "WORD"},
		quaggaProtocolsBgpPeerGroupFilterListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "filter-list"},
		quaggaProtocolsBgpPeerGroupFilterList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "local-as", "WORD"},
		quaggaProtocolsBgpPeerGroupLocalAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "local-as", "WORD", "no-prepend"},
		quaggaProtocolsBgpPeerGroupLocalAsNoPrepend)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "maximum-prefix", "WORD"},
		quaggaProtocolsBgpPeerGroupMaximumPrefix)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "nexthop-self"},
		quaggaProtocolsBgpPeerGroupNexthopSelf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "override-capability"},
		quaggaProtocolsBgpPeerGroupOverrideCapability)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "passive"},
		quaggaProtocolsBgpPeerGroupPassive)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "password", "WORD"},
		quaggaProtocolsBgpPeerGroupPassword)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "prefix-list", "export", "WORD"},
		quaggaProtocolsBgpPeerGroupPrefixListExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "prefix-list", "import", "WORD"},
		quaggaProtocolsBgpPeerGroupPrefixListImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "prefix-list"},
		quaggaProtocolsBgpPeerGroupPrefixList)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "remote-as", "WORD"},
		quaggaProtocolsBgpPeerGroupRemoteAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "remove-private-as"},
		quaggaProtocolsBgpPeerGroupRemovePrivateAs)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "route-map", "export", "WORD"},
		quaggaProtocolsBgpPeerGroupRouteMapExport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "route-map", "import", "WORD"},
		quaggaProtocolsBgpPeerGroupRouteMapImport)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "route-map"},
		quaggaProtocolsBgpPeerGroupRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "route-reflector-client"},
		quaggaProtocolsBgpPeerGroupRouteReflectorClient)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "route-server-client"},
		quaggaProtocolsBgpPeerGroupRouteServerClient)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "shutdown"},
		quaggaProtocolsBgpPeerGroupShutdown)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "soft-reconfiguration", "inbound"},
		quaggaProtocolsBgpPeerGroupSoftReconfigurationInbound)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "soft-reconfiguration"},
		quaggaProtocolsBgpPeerGroupSoftReconfiguration)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "ttl-security", "hops", "WORD"},
		quaggaProtocolsBgpPeerGroupTtlSecurityHops)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "ttl-security"},
		quaggaProtocolsBgpPeerGroupTtlSecurity)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "unsuppress-map", "WORD"},
		quaggaProtocolsBgpPeerGroupUnsuppressMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "update-source", "WORD"},
		quaggaProtocolsBgpPeerGroupUpdateSource)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "peer-group", "WORD", "weight", "WORD"},
		quaggaProtocolsBgpPeerGroupWeight)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "connected", "metric", "WORD"},
		quaggaProtocolsBgpRedistributeConnectedMetric)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "connected"},
		quaggaProtocolsBgpRedistributeConnected)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "connected", "route-map", "WORD"},
		quaggaProtocolsBgpRedistributeConnectedRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "kernel", "metric", "WORD"},
		quaggaProtocolsBgpRedistributeKernelMetric)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "kernel"},
		quaggaProtocolsBgpRedistributeKernel)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "kernel", "route-map", "WORD"},
		quaggaProtocolsBgpRedistributeKernelRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute"},
		quaggaProtocolsBgpRedistribute)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "ospf", "metric", "WORD"},
		quaggaProtocolsBgpRedistributeOspfMetric)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "ospf"},
		quaggaProtocolsBgpRedistributeOspf)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "ospf", "route-map", "WORD"},
		quaggaProtocolsBgpRedistributeOspfRouteMap)
	/*
		configParser.InstallCmd(
			[]string{"protocols", "bgp", "WORD", "redistribute", "rip", "metric", "WORD"},
			quaggaProtocolsBgpRedistributeRipMetric)
		configParser.InstallCmd(
			[]string{"protocols", "bgp", "WORD", "redistribute", "rip"},
			quaggaProtocolsBgpRedistributeRip)
		configParser.InstallCmd(
			[]string{"protocols", "bgp", "WORD", "redistribute", "rip", "route-map", "WORD"},
			quaggaProtocolsBgpRedistributeRipRouteMap)
	*/
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "static", "metric", "WORD"},
		quaggaProtocolsBgpRedistributeStaticMetric)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "static"},
		quaggaProtocolsBgpRedistributeStatic)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "redistribute", "static", "route-map", "WORD"},
		quaggaProtocolsBgpRedistributeStaticRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "timers", "holdtime", "WORD"},
		quaggaProtocolsBgpTimersHoldtime)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "timers", "keepalive", "WORD"},
		quaggaProtocolsBgpTimersKeepalive)
	configParser.InstallCmd(
		[]string{"protocols", "bgp", "WORD", "timers"},
		quaggaProtocolsBgpTimers)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "access-list", "WORD"},
		quaggaProtocolsOspfAccessList)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "access-list", "WORD", "export", "WORD"},
		quaggaProtocolsOspfAccessListExport)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD"},
		quaggaProtocolsOspfArea)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type"},
		quaggaProtocolsOspfAreaAreaType)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type", "normal"},
		quaggaProtocolsOspfAreaAreaTypeNormal)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type", "nssa", "default-cost", "WORD"},
		quaggaProtocolsOspfAreaAreaTypeNssaDefaultCost)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type", "nssa", "no-summary"},
		quaggaProtocolsOspfAreaAreaTypeNssaNoSummary)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type", "nssa"},
		quaggaProtocolsOspfAreaAreaTypeNssa)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type", "nssa", "translate", "WORD"},
		quaggaProtocolsOspfAreaAreaTypeNssaTranslate)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type", "stub", "default-cost", "WORD"},
		quaggaProtocolsOspfAreaAreaTypeStubDefaultCost)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type", "stub", "no-summary"},
		quaggaProtocolsOspfAreaAreaTypeStubNoSummary)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "area-type", "stub"},
		quaggaProtocolsOspfAreaAreaTypeStub)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "authentication", "WORD"},
		quaggaProtocolsOspfAreaAuthentication)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "network", "A.B.C.D/M"},
		quaggaProtocolsOspfAreaNetwork)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "range", "A.B.C.D/M"},
		quaggaProtocolsOspfAreaRange)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "range", "A.B.C.D/M", "cost", "WORD"},
		quaggaProtocolsOspfAreaRangeCost)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "range", "A.B.C.D/M", "not-advertise"},
		quaggaProtocolsOspfAreaRangeNotAdvertise)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "range", "A.B.C.D/M", "substitute", "A.B.C.D/M"},
		quaggaProtocolsOspfAreaRangeSubstitute)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "shortcut", "WORD"},
		quaggaProtocolsOspfAreaShortcut)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D"},
		quaggaProtocolsOspfAreaVirtualLink)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "authentication", "md5", "key-id", "WORD"},
		quaggaProtocolsOspfAreaVirtualLinkAuthenticationMd5KeyId)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "authentication", "md5", "key-id", "WORD", "md5-key", "WORD"},
		quaggaProtocolsOspfAreaVirtualLinkAuthenticationMd5KeyIdMd5Key)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "authentication", "md5"},
		quaggaProtocolsOspfAreaVirtualLinkAuthenticationMd5)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "authentication"},
		quaggaProtocolsOspfAreaVirtualLinkAuthentication)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "authentication", "plaintext-password", "WORD"},
		quaggaProtocolsOspfAreaVirtualLinkAuthenticationPlaintextPassword)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "dead-interval", "WORD"},
		quaggaProtocolsOspfAreaVirtualLinkDeadInterval)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "hello-interval", "WORD"},
		quaggaProtocolsOspfAreaVirtualLinkHelloInterval)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "retransmit-interval", "WORD"},
		quaggaProtocolsOspfAreaVirtualLinkRetransmitInterval)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "area", "WORD", "virtual-link", "A.B.C.D", "transmit-delay", "WORD"},
		quaggaProtocolsOspfAreaVirtualLinkTransmitDelay)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "auto-cost"},
		quaggaProtocolsOspfAutoCost)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "auto-cost", "reference-bandwidth", "WORD"},
		quaggaProtocolsOspfAutoCostReferenceBandwidth)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "default-information"},
		quaggaProtocolsOspfDefaultInformation)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "default-information", "originate", "always"},
		quaggaProtocolsOspfDefaultInformationOriginateAlways)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "default-information", "originate", "metric-type", "WORD"},
		quaggaProtocolsOspfDefaultInformationOriginateMetricType)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "default-information", "originate", "metric", "WORD"},
		quaggaProtocolsOspfDefaultInformationOriginateMetric)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "default-information", "originate"},
		quaggaProtocolsOspfDefaultInformationOriginate)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "default-information", "originate", "route-map", "WORD"},
		quaggaProtocolsOspfDefaultInformationOriginateRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "default-metric", "WORD"},
		quaggaProtocolsOspfDefaultMetric)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "distance", "global", "WORD"},
		quaggaProtocolsOspfDistanceGlobal)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "distance"},
		quaggaProtocolsOspfDistance)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "distance", "ospf", "external", "WORD"},
		quaggaProtocolsOspfDistanceOspfExternal)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "distance", "ospf", "inter-area", "WORD"},
		quaggaProtocolsOspfDistanceOspfInterArea)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "distance", "ospf", "intra-area", "WORD"},
		quaggaProtocolsOspfDistanceOspfIntraArea)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "distance", "ospf"},
		quaggaProtocolsOspfDistanceOspf)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "log-adjacency-changes", "detail"},
		quaggaProtocolsOspfLogAdjacencyChangesDetail)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "log-adjacency-changes"},
		quaggaProtocolsOspfLogAdjacencyChanges)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "max-metric"},
		quaggaProtocolsOspfMaxMetric)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "max-metric", "router-lsa", "administrative"},
		quaggaProtocolsOspfMaxMetricRouterLsaAdministrative)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "max-metric", "router-lsa"},
		quaggaProtocolsOspfMaxMetricRouterLsa)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "max-metric", "router-lsa", "on-shutdown", "WORD"},
		quaggaProtocolsOspfMaxMetricRouterLsaOnShutdown)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "max-metric", "router-lsa", "on-startup", "WORD"},
		quaggaProtocolsOspfMaxMetricRouterLsaOnStartup)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "mpls-te", "enable"},
		quaggaProtocolsOspfMplsTeEnable)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "mpls-te"},
		quaggaProtocolsOspfMplsTe)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "mpls-te", "router-address", "A.B.C.D"},
		quaggaProtocolsOspfMplsTeRouterAddress)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "neighbor", "A.B.C.D"},
		quaggaProtocolsOspfNeighbor)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "neighbor", "A.B.C.D", "poll-interval", "WORD"},
		quaggaProtocolsOspfNeighborPollInterval)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "neighbor", "A.B.C.D", "priority", "WORD"},
		quaggaProtocolsOspfNeighborPriority)
	configParser.InstallCmd(
		[]string{"protocols", "ospf"},
		quaggaProtocolsOspf)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "parameters", "abr-type", "WORD"},
		quaggaProtocolsOspfParametersAbrType)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "parameters"},
		quaggaProtocolsOspfParameters)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "parameters", "opaque-lsa"},
		quaggaProtocolsOspfParametersOpaqueLsa)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "parameters", "rfc1583-compatibility"},
		quaggaProtocolsOspfParametersRfc1583Compatibility)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "parameters", "router-id", "A.B.C.D"},
		quaggaProtocolsOspfParametersRouterId)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "passive-interface-exclude", "WORD"},
		quaggaProtocolsOspfPassiveInterfaceExclude)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "passive-interface", "WORD"},
		quaggaProtocolsOspfPassiveInterface)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric-type", "WORD"},
		quaggaProtocolsOspfRedistributeBgpMetricType)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "bgp", "metric", "WORD"},
		quaggaProtocolsOspfRedistributeBgpMetric)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "bgp"},
		quaggaProtocolsOspfRedistributeBgp)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "bgp", "route-map", "WORD"},
		quaggaProtocolsOspfRedistributeBgpRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric-type", "WORD"},
		quaggaProtocolsOspfRedistributeConnectedMetricType)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "connected", "metric", "WORD"},
		quaggaProtocolsOspfRedistributeConnectedMetric)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "connected"},
		quaggaProtocolsOspfRedistributeConnected)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "connected", "route-map", "WORD"},
		quaggaProtocolsOspfRedistributeConnectedRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric-type", "WORD"},
		quaggaProtocolsOspfRedistributeKernelMetricType)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "kernel", "metric", "WORD"},
		quaggaProtocolsOspfRedistributeKernelMetric)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "kernel"},
		quaggaProtocolsOspfRedistributeKernel)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "kernel", "route-map", "WORD"},
		quaggaProtocolsOspfRedistributeKernelRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute"},
		quaggaProtocolsOspfRedistribute)
	/*
		configParser.InstallCmd(
			[]string{"protocols", "ospf", "redistribute", "rip", "metric-type", "WORD"},
			quaggaProtocolsOspfRedistributeRipMetricType)
		configParser.InstallCmd(
			[]string{"protocols", "ospf", "redistribute", "rip", "metric", "WORD"},
			quaggaProtocolsOspfRedistributeRipMetric)
		configParser.InstallCmd(
			[]string{"protocols", "ospf", "redistribute", "rip"},
			quaggaProtocolsOspfRedistributeRip)
		configParser.InstallCmd(
			[]string{"protocols", "ospf", "redistribute", "rip", "route-map", "WORD"},
			quaggaProtocolsOspfRedistributeRipRouteMap)
	*/
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "static", "metric-type", "WORD"},
		quaggaProtocolsOspfRedistributeStaticMetricType)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "static", "metric", "WORD"},
		quaggaProtocolsOspfRedistributeStaticMetric)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "static"},
		quaggaProtocolsOspfRedistributeStatic)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "redistribute", "static", "route-map", "WORD"},
		quaggaProtocolsOspfRedistributeStaticRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "refresh"},
		quaggaProtocolsOspfRefresh)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "refresh", "timers", "WORD"},
		quaggaProtocolsOspfRefreshTimers)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "timers"},
		quaggaProtocolsOspfTimers)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "timers", "throttle"},
		quaggaProtocolsOspfTimersThrottle)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "delay", "WORD"},
		quaggaProtocolsOspfTimersThrottleSpfDelay)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "initial-holdtime", "WORD"},
		quaggaProtocolsOspfTimersThrottleSpfInitialHoldtime)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "timers", "throttle", "spf", "max-holdtime", "WORD"},
		quaggaProtocolsOspfTimersThrottleSpfMaxHoldtime)
	configParser.InstallCmd(
		[]string{"protocols", "ospf", "timers", "throttle", "spf"},
		quaggaProtocolsOspfTimersThrottleSpf)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "area", "WORD"},
		quaggaProtocolsOspfv3Area)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "area", "WORD", "export-list", "WORD"},
		quaggaProtocolsOspfv3AreaExportList)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "area", "WORD", "import-list", "WORD"},
		quaggaProtocolsOspfv3AreaImportList)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "area", "WORD", "interface", "WORD"},
		quaggaProtocolsOspfv3AreaInterface)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "area", "WORD", "range", "X:X::X:X/M"},
		quaggaProtocolsOspfv3AreaRange)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "area", "WORD", "range", "X:X::X:X/M", "advertise"},
		quaggaProtocolsOspfv3AreaRangeAdvertise)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "area", "WORD", "range", "X:X::X:X/M", "not-advertise"},
		quaggaProtocolsOspfv3AreaRangeNotAdvertise)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3"},
		quaggaProtocolsOspfv3)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "parameters"},
		quaggaProtocolsOspfv3Parameters)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "parameters", "router-id", "A.B.C.D"},
		quaggaProtocolsOspfv3ParametersRouterId)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute", "bgp"},
		quaggaProtocolsOspfv3RedistributeBgp)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute", "bgp", "route-map", "WORD"},
		quaggaProtocolsOspfv3RedistributeBgpRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute", "connected"},
		quaggaProtocolsOspfv3RedistributeConnected)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute", "connected", "route-map", "WORD"},
		quaggaProtocolsOspfv3RedistributeConnectedRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute", "kernel"},
		quaggaProtocolsOspfv3RedistributeKernel)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute", "kernel", "route-map", "WORD"},
		quaggaProtocolsOspfv3RedistributeKernelRouteMap)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute"},
		quaggaProtocolsOspfv3Redistribute)
	/*
		configParser.InstallCmd(
			[]string{"protocols", "ospfv3", "redistribute", "ripng"},
			quaggaProtocolsOspfv3RedistributeRipng)
		configParser.InstallCmd(
			[]string{"protocols", "ospfv3", "redistribute", "ripng", "route-map", "WORD"},
			quaggaProtocolsOspfv3RedistributeRipngRouteMap)
	*/
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute", "static"},
		quaggaProtocolsOspfv3RedistributeStatic)
	configParser.InstallCmd(
		[]string{"protocols", "ospfv3", "redistribute", "static", "route-map", "WORD"},
		quaggaProtocolsOspfv3RedistributeStaticRouteMap)
}
