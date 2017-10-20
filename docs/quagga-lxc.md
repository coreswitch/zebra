# Running Zebra 2.0 and Quagga protocol module on LXC

Zebra 2.0 and Quagga protocol modules can run on LXC container with some additional configurations.

This document describes an example steps to run Quagga bgpd and ospfd with Zebra 2.0 / OpenConfigd on LXC container.

> Note: Steps are verified using Ubuntu 16.04.03 and Quagga 0.99.24.1

### Topology

```
 +-------+   +-------+
 | host1 |   | host2 |
 +---+---+   +---+---+
     |           |
 +---+-----------+----+
 | lxcbr0 10.0.3.1/24 |
 +--------------------+
```

* host1 IP address
    * eth0 : 10.0.3.61/24
    * lo: 10.10.0.1/32, 10.10.10.1/31
* host2 IP address
    * eth0 : 10.0.3.62/24
    * lo: 10.10.0.2/32, 10.10.10.2/31

### Steps on the host

Install OpenConfigd with CLI and Zebra 2.0.

* https://github.com/coreswitch/openconfigd/blob/master/README.md
* https://github.com/coreswitch/zebra/blob/master/README.md

Install quaggad, an openconfigd configuration gateway to quagga.

* https://github.com/coreswitch/openconfigd/tree/master/quagga

```
$ go get github.com/coreswitch/openconfigd/quagga/quaggad
```

Install Quagga.

```
# apt install quagga
```

Install LXC and create 2 containers.

```
# apt install lxc
# lxc-create -t ubuntu -n host1
# lxc-create -t ubuntu -n host2
```

Container Network will be automatically created.

```
$ ip a
 ... snip ...
3: lxcbr0: <NO-CARRIER,BROADCAST,MULTICAST,UP> mtu 1500 qdisc noqueue state DOWN group default qlen 1000
    link/ether 00:16:3e:00:00:00 brd ff:ff:ff:ff:ff:ff
    inet 10.0.3.1/24 scope global lxcbr0
       valid_lft forever preferred_lft forever
```

Copy `vtysh` under `/usr/local/bin/`
> Note: `/usr/bin` will be initialized after container restart, and symbolic link is not accessbile from container.

```
$ cp /usr/bin/vtysh /usr/local/bin/
```

Create `/var/lib/lxc/<host>/fstab` and mount required directories.

```
> Example for host1
# vi /var/lib/lxc/host1/fstab
/home/<user>/go home/ubuntu/go none bind,create=dir
/usr/lib/quagga usr/lib/quagga none bind,create=dir
/usr/local/bin usr/local/bin none bind,create=dir
/usr/lib usr/lib none bind,create=dir
```

Edit `/var/lib/lxc/<host>/config` to mount and assign static address to the hosts.

```
# /var/lib/lxc/host1/config
lxc.mount= /var/lib/lxc/host1/fstab
lxc.network.ipv4 = 10.0.3.61/24
lxc.network.ipv4.gateway = auto
```

### Steps inside the containers.

Start and log into the contatiner. (default user/password is 'ubuntu/ubuntu')

```
$ sudo lxc-start -n host1 -d
$ ssh ubuntu@10.0.3.61
```

You can also use console to log into container.

```
# lxc-console -n host1
> Type `Ctrl+A Q` to exit from console.
```

Set $PATH for `vtysh` and $GOPATH

```
$ vi ~/.bashrc
export GOPATH=$HOME/go
export PATH=$GOPATH/bin:$PATH
export PATH=/usr/local/bin:$PATH
$ source ~/.bashrc
```

Change `eth0` to manual. (Since you statically set address in `/var/lib/lxc/host1/config`)
```
# vi /etc/network/interfaces
auto eth0
iface eth0 inet manual
```

Install bash-completion for CLI.

```
# apt install bash-completion
# cp $GOPATH/src/github.com/coreswitch/openconfigd/bash_completion.d/cli /etc/bash_completion.d/
```

Add group and user.

```
# groupadd quaggavty
# groupadd quagga
# useradd -g quagga -s /bin/false quagga
```

Create `/etc/pam.d/quagga`

```
$ sudo vi /etc/pam.d/quagga
# Any user may call vtysh but only those belonging to the group quaggavty can
# actually connect to the socket and use the program.
auth    sufficient      pam_permit.so
```


### Run

Create and run below script to start bgpd, ospfd, openconfigd and ribd in the container.

```
#!/bin/bash

# run with sudo

mkdir /var/run/quagga
chown quagga:quagga /var/run/quagga/

openconfigd -y ${GOPATH}/src/github.com/coreswitch/openconfigd/yang &
ribd &

sleep 10
chmod 777 /var/run/zserv.api

/usr/lib/quagga/bgpd --config_file /dev/null --pid_file /var/run/quagga/bgpd.pid --socket /var/run/zserv.api &
/usr/lib/quagga/ospfd --config_file /dev/null --pid_file /var/run/quagga/ospfd.pid --socket /var/run/zserv.api &

quaggad &
```

### config and show examples

Open new terminal and start CLI. (to avoid logs to show up in terminal running CLI)

```
ubuntu@host1:~$ cli
host1> configure
```

Configure ospf/bgpd on both host.

```
> host1
set interfaces interface lo ipv4 address 10.10.0.1/32
set interfaces interface lo ipv4 address 10.10.10.1/32
set protocols ospf parameters router-id 10.10.0.1
set protocols ospf area 0.0.0.0 network 10.0.3.0/24
set protocols ospf area 0.0.0.0 network 10.10.10.0/24
set protocols bgp 65001 parameters router-id 10.0.3.61
set protocols bgp 65001 neighbor 10.0.3.62 remote-as 65002
set protocols bgp 65001 network 10.10.0.1/32
commit

> host2
set interfaces interface lo ipv4 address 10.10.0.2/32
set interfaces interface lo ipv4 address 10.10.10.2/32
set protocols ospf parameters router-id 10.10.0.2
set protocols ospf area 0.0.0.0 network 10.0.3.0/24
set protocols ospf area 0.0.0.0 network 10.10.10.0/24
set protocols bgp 65002 parameters router-id 10.0.3.62
set protocols bgp 65002 neighbor 10.0.3.61 remote-as 65001
set protocols bgp 65002 network 10.10.0.2/32
commit
```

Show routes.

```
host1>show ip ospf route
============ OSPF network routing table ============
N    10.0.3.0/24           [10] area: 0.0.0.0
                           directly attached to eth0
N    10.10.10.1/32         [10] area: 0.0.0.0
                           directly attached to lo

============ OSPF router routing table =============

============ OSPF external routing table ===========

host1>show ip bgp
BGP table version is 0, local router ID is 10.0.3.61
Status codes: s suppressed, d damped, h history, * valid, > best, = multipath,
              i internal, r RIB-failure, S Stale, R Removed
Origin codes: i - IGP, e - EGP, ? - incomplete

   Network          Next Hop            Metric LocPrf Weight Path
*> 10.10.0.1/32     0.0.0.0                  0         32768 i
*> 10.10.0.2/32     10.0.3.62                0             0 65002 i

Total number of prefixes 2
host1>show ip route
Codes: K - kernel, C - connected, S - static, R - RIP, B - BGP
       O - OSPF, IA - OSPF inter area
       N1 - OSPF NSSA external type 1, N2 - OSPF NSSA external type 2
       E1 - OSPF external type 1, E2 - OSPF external type 2
       i - IS-IS, L1 - IS-IS level-1, L2 - IS-IS level-2, ia - IS-IS inter area

K        0.0.0.0/0 via 10.0.3.1, eth0
C        10.0.3.0/24 is directly connected eth0
C        10.10.0.1/32 is directly connected lo
B        10.10.0.2/32 [200/0] via 10.0.3.62
C        10.10.10.1/32 is directly connected lo
C        127.0.0.0/8 is directly connected lo
```
