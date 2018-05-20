# SRv6 Local Functions

Since Kernel version 4.14, Linux supports (part of) End.* functions in SRv6 Network Programming [1].
> [1] draft-filsfils-spring-srv6-network-programming

This document descrives how to show/set/delete End.* functions from Zebra and Openconfigd running on Linux dataplane.

## CLI commands

set commands (configure mode)
```
# set routing-options ipv6 localsid <prefix> nexthop <address> action <action> <parameters>
<action> : SRv6 End.* functions.
<parameters> : parameters required for each End.* functions.
```

delete commands (configure mode)
```
# delete routing-options ipv6 localsid <prefix>
```

show commands
```
> show ipv4 route
> show ipv6 route
```

## Examples

Make sure you have "interface" and "route to next hop" you want to configure.

```
$ sudo ip -6 addr add 2001:db8::100/64 dev enp0s3
$ ip a
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host
       valid_lft forever preferred_lft forever
2: enp0s3: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether 08:00:27:db:e8:12 brd ff:ff:ff:ff:ff:ff
    inet 10.0.2.15/24 brd 10.0.2.255 scope global enp0s3
       valid_lft forever preferred_lft forever
    inet6 2001:db8::100/64 scope global
       valid_lft forever preferred_lft forever
    inet6 fe80::a00:27ff:fedb:e812/64 scope link
       valid_lft forever preferred_lft forever
3: lxcbr0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 00:16:3e:00:00:00 brd ff:ff:ff:ff:ff:ff
    inet 10.0.3.1/24 scope global lxcbr0
       valid_lft forever preferred_lft forever
```

Set configuration via CLI and commit.

```
coreswitch# 
set routing-options ipv6 localsid fc00::a1/128 nexthop 2001:db8::1 action End
set routing-options ipv6 localsid fc00::a2/128 nexthop 2001:db8::1 action End.X nh6 fc00::1:1
set routing-options ipv6 localsid fc00::a3/128 nexthop 2001:db8::1 action End.T table 100
set routing-options ipv6 localsid fc00::a4/128 nexthop 2001:db8::1 action End.DX2 oif lxcbr0
set routing-options ipv6 localsid fc00::a5/128 nexthop 2001:db8::1 action End.DX6 nh6 fc00::1:1
set routing-options ipv6 localsid fc00::a6/128 nexthop 2001:db8::1 action End.DX4 nh4 10.0.3.254
set routing-options ipv6 localsid fc00::a7/128 nexthop 2001:db8::1 action End.DT6 table 200
set routing-options ipv6 localsid fc00::a8/128 nexthop 2001:db8::1 action End.B6 segments beaf::1 beaf::2
set routing-options ipv6 localsid fc00::a9/128 nexthop 2001:db8::1 action End.B6.Encaps segments beaf::1 beaf::2
coreswitch# commit
```

Check routes are correctly applyed

```
coreswitch#run show ipv6 route
[run show ipv6 route]
Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area

C     ::1/128 [0/0]
      via lo, directly connected
C     2001:db8::/64 [0/0]
      via enp0s3, directly connected
S     fc00::a1/128 [1/0]
      encap seg6local action End
      via 2001:db8::1
S     fc00::a2/128 [1/0]
      encap seg6local action End.X nh6 fc00::1:1
      via 2001:db8::1
S     fc00::a3/128 [1/0]
      encap seg6local action End.T table 100
      via 2001:db8::1
S     fc00::a4/128 [1/0]
      encap seg6local action End.DX2 oif lxcbr0
      via 2001:db8::1
S     fc00::a5/128 [1/0]
      encap seg6local action End.DX6 nh6 fc00::1:1
      via 2001:db8::1
S     fc00::a6/128 [1/0]
      encap seg6local action End.DX4 nh4 10.0.3.254
      via 2001:db8::1
S     fc00::a7/128 [1/0]
      encap seg6local action End.DT6 table 200
      via 2001:db8::1
S     fc00::a8/128 [1/0]
      encap seg6local action End.B6 segs 3 [ beaf::1 beaf::2 :: ]
      via 2001:db8::1
S     fc00::a9/128 [1/0]
      encap seg6local action End.B6.Encaps segs 2 [ beaf::1 beaf::2 ]
      via 2001:db8::1
C     fe80::/64 [0/0]
      via enp0s3, directly connected
```

Note: above routes corresponds to below routes shown via iproute2 command.

```
$ ip -6 route
2001:db8::/64 dev enp0s3 proto kernel metric 256 pref medium
fc00::a1  encap seg6local action End via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fc00::a2  encap seg6local action End.X nh6 fc00::1:1 via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fc00::a3  encap seg6local action End.T table 100 via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fc00::a4  encap seg6local action End.DX2 oif lxcbr0 via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fc00::a5  encap seg6local action End.DX6 nh6 fc00::1:1 via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fc00::a6  encap seg6local action End.DX4 nh4 10.0.3.254 via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fc00::a7  encap seg6local action End.DT6 table 200 via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fc00::a8  encap seg6local action End.B6 srh segs 3 [ beaf::1 beaf::2 :: ] via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fc00::a9  encap seg6local action End.B6.Encaps srh segs 2 [ beaf::1 beaf::2 ] via 2001:db8::1 dev enp0s3 proto zebra metric 1024 pref medium
fe80::/64 dev enp0s3 proto kernel metric 256 pref medium
```

Delete configuration and commit.
```
coreswitch#
delete routing-options ipv6 localsid fc00::a1/128
delete routing-options ipv6 localsid fc00::a2/128
delete routing-options ipv6 localsid fc00::a3/128
delete routing-options ipv6 localsid fc00::a4/128
delete routing-options ipv6 localsid fc00::a5/128
delete routing-options ipv6 localsid fc00::a6/128
delete routing-options ipv6 localsid fc00::a7/128
delete routing-options ipv6 localsid fc00::a8/128
delete routing-options ipv6 localsid fc00::a9/128
coreswitch# commit
```
