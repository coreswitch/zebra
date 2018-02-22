# Zebra 2.0

Zebra 2.0 is an open source implementation as a successor of GNU Quagga/Zebra
project. Zebra 2.0 is implemented in Go programming language to fit in container
or virtual machine environment.

## Getting started

Zebra 2.0 depends upon [OpenConfigd](https://github.com/hash-set/openconfigd) so
before installing Zebra, openconfigd should be installed. Here is a most
simplest way of doing it.

### Install

Please checkout esi-ribd as github.com/hash-set/zebra.

``` bash
$ cd $GOPATH
$ mkdir -p src/github.com/hash-set
$ git clone git@bitbucket.org:ntti3/esi-ribd.git zebra
$ go get github.com/hash-set/zebra/rib/ribd
```

### Building Debian package

To build debian package, please goto debian directory then execute `make`.

``` bash
$ cd debian
$ make
```

This will make package under out directory.

## Using Zebra 2.0

 * [Getting Started](https://github.com/hash-set/zebra/blob/master/docs/getting-started.md)
 * [Router ID](https://github.com/hash-set/zebra/blob/master/docs/router-id.md)
 * [Integration with Quagga](https://github.com/hash-set/zebra/blob/master/docs/quagga.md)
 * [Integration with GoBGP](https://github.com/hash-set/zebra/blob/master/docs/gobgp.md)
