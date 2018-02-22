# Getting Started

To install Zebra 2.0 you need netlink, netns lib for Go cloned locally.

* netlink / netns library for go
   * https://github.com/vishvananda/netlink
   * https://github.com/vishvananda/netns

* Sample directory structure:

```sh
~/go/src/github.com/hash-set/zebra/
~/go/src/github.com/vishvananda/netlink
~/go/src/github.com/vishvananda/netns
```

* Sample clone / installation log

You can clone all dependencies using `go get` and install.
```sh
$ export GOPATH=/home/$USER/go
$ go install github.com/hash-set/zebra/rib/ribd
```

Instead, you could also git clone one by one.

```sh
~/go/src/github.com$ mkdir vishvananda
~/go/src/github.com$ git clone git://github.com/vishvananda/netlink vishvananda/netlink/
~/go/src/github.com$ git clone git://github.com/vishvananda/netns vishvananda/netns/
~/go/src/github.com$ ls vishvananda/
netlink  netns

~/go/src/github.com$ go install github.com/hash-set/zebra/rib/ribd
~/go/src/github.com$ go install github.com/hash-set/zebra/fea
```
