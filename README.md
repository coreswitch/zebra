# zebra

zebra is an open source implementation as a successor of GNU Zebra and Quagga
project. zebra is implemented in Go programming language to fit in container
or virtual machine environment.

## Getting started

zebra depends upon [OpenConfigd](https://github.com/coreswitch/openconfigd) so
before installing zebra, openconfigd should be installed.

### Install

Please build ribd as a first zebra module.

``` bash
$ go get github.com/coreswitch/zebra/rib/ribd
```

Then execute ribd under root privilege.

``` bash
$ sudo ${GOPATH}/bin/ribd
```

## Using zebra

 * [Integration with Quagga](https://github.com/coreswitch/zebra/blob/master/docs/quagga.md)
 * [Integration with GoBGP](https://github.com/coreswitch/zebra/blob/master/docs/gobgp.md)
 * [Router ID](https://github.com/coreswitch/zebra/blob/master/docs/router-id.md)
