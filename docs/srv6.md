# SRv6 (Segment Routing for IPv6)

SRv6 show/set/delete commands are supported on Zebra and Openconfigd using following dataplane.

* Linux Kernel version 4.10 or later
	* SRv6 routes are implemented as LWT (Light Weight Tunnel) on Linux.
	* Kernel version 4.10 supports T.Insert and T.Encaps
	* Kernel version 4.14 supports some of End.* functions in SRv6 Network Programming [1]
	* [1] draft-filsfils-spring-srv6-network-programming

This document descrives how to show/set/delete T.Insert and T.Encaps.

See "srv6local.md" for description of other End.* functions.

## Prerequisites

Below sysctl keys must be enabled (1) to make Linux Kernel to actually handle SRv6 packets. It's not required to just test CLI commands to show/set/delete SRv6 routes.

```
net.ipv6.conf.*.seg6_enabled = 0
net.ipv6.conf.*.seg6_require_hmac = 0
```

## CLI commands

set commands (configure mode)
```
# set routing-options ipv6 route-srv6 <prefix> nexthop <address> seg6 <mode> segments <segments>
# set routing-options ipv4 route-srv6 <prefix> nexthop <address> seg6 <mode> segments <segments>
<mode> : encap or inline
```

delete commands (configure mode)
```
# delete routing-options ipv6 route-srv6 <prefix>
# delete routing-options ipv4 route-srv6 <prefix>
```

show commands
```
> show ipv4 route
> show ipv6 route
```

## Examples

CLI output example of set/show commands.

```
> configure 
# set routing-options ipv6 route-srv6 beef::/64 nexthop 2001:db8::1 seg6 encap segments c0be::1 c0be::2 c0be::3
# set routing-options ipv6 route-srv6 c0be::/64 nexthop 2001:db8::1 seg6 inline segments c0be::1 c0be::2 c0be::3
# commit
```

```
# show
... snip ...
routing-options {
    ipv6 {
        route-srv6 beef::/64 {
            nexthop 2001:db8::1 {
                seg6 encap {
                    segments c0be::1 c0be::2 c0be::3;
                }
            }
        }
        route-srv6 c0be::/64 {
            nexthop 2001:db8::1 {
                seg6 inline {
                    segments c0be::1 c0be::2 c0be::3;
                }
            }
        }
    }
}
```

```
> show ipv6 route
Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area

C     ::1/128 [0/0]
      via lo, directly connected
C     2001:db8::/64 [0/0]
      via enp0s8, directly connected
S     beef::/64 [1/0]
      encap seg6 mode encap segs 3 [ c0be::1 c0be::2 c0be::3 ]
      via 2001:db8::1
S     c0be::/64 [1/0]
      encap seg6 mode inline segs 4 [ c0be::1 c0be::2 c0be::3 :: ]
      via 2001:db8::1
C     fe80::/64 [0/0]
      via enp0s3, directly connected
```
