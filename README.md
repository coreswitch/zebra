# Zebra 2.0

Zebra 2.0 is an open source implementation as a successor of GNU Zebra and Quagga
project. Zebra 2.0 is implemented in Go programming language to fit in container
or virtual machine environment.

## Getting started

Zebra 2.0 depends upon [OpenConfigd](https://github.com/coreswitch/openconfigd) so
before installing Zebra, openconfigd should be installed.

### Install

Please build ribd as first zebra module.

``` bash
$ go get github.com/coreswitch/zebra/rib/ribd
```

Then execute ribd under root privilege.

``` bash
$ sudo ${GOPATH}/bin/ribd
```

## Using Zebra 2.0

 * [Integration with Quagga](https://github.com/coreswitch/zebra/blob/master/docs/quagga.md)
 * [Integration with GoBGP](https://github.com/coreswitch/zebra/blob/master/docs/gobgp.md)
 * [Router ID](https://github.com/coreswitch/zebra/blob/master/docs/router-id.md)
