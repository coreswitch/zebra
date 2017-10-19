# Quagga

All of Quagga modules such as `ospfd`, `bgpd`, `isisd` can be used as Zebra's
protocol module.

## Quagga protocol module configuration

In case of Quagga is build with default configuration. Basically no
configuration is necessary. Just simply launch Quagga module should work fine.

Here is a example of Quagga `ospfd` with Zebra 2.0 `ribd`.

```shel
$ sudo openconfigd -d
$ sudo ribd -d
$ sudo ospfd --user=root
```

When Quagga is built with specific configuration such as `--localstatedir`.
Zebra API path will be changed.

```shell
$ ./configure --localstatedir=/var/run/quagga
```

When Quagga is built under above configuration, Zebra API path will be
`/var/run/quagga/zserv.api`. In such case, Zebra 2.0 API path should be
specified by `--socket` option of Quagga module.

```shel
$ sudo openconfigd -d
$ sudo ribd -d
$ sudo ospfd --user=root --socket=/var/run/zserv.api
```
