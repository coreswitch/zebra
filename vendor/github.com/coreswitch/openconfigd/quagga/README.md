# quaggad

`quaggad` is openconfigd configuration gateway to quagga.

```
$ go get github.com/coreswitch/zebra/rib/ribd
$ go get github.com/coreswitch/openconfigd/openconfigd
$ go get github.com/coreswitch/openconfigd/quagga/quaggad
```

```
$ sudo ${GOPATH}/bin/openconfigd -y ${GOPATH}/src/github.com/coreswitch/openconfigd/yang &
$ sudo ${GOPATH}/bin/ribd &
$ sudo /usr/lib/quagga/bgpd --config_file /dev/null --pid_file /tmp/bgpd.pid --socket /var/run/zserv.api &
$ sudo ${GOPATH}/bin/quaggad &
```
