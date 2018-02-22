# openconfigd

openconfigd is software which manages [OpenConfig](http://www.openconfig.net/)
common data models for networking. It handles networking protocol configuration
of switch, router, DNS, DHCP, NAT and Firewall.

openconfigd reads YANG model definition then generate configuration schema from
it.

### Install

Below command install `openconfigd` and `cli_command` to $GOPATH/bin.

``` bash
$ go get github.com/coreswitch/openconfigd/openconfigd
$ go get github.com/coreswitch/openconfigd/cli_command
```

CLI command build and set up.

``` bash
$ cd $GOPATH/src/github.com/coreswitch/openconfigd/cli
$ ./configure
$ make
$ sudo make install
```

will install `cli` command to /usr/local/bin.

CLI completion file another `cli` file exists under bash_completion.d. On
Ubuntu, this file should be installed under `/etc/bach_completion.d`

``` bash
$ cd $GOPATH/src/github.com/coreswitch/openconfigd/bash_completion.d
$ sudo cp cli /etc/bash_completion.d/
```

### Quick Start

Invoke openconfigd, then start cli.  "show version" display version information.

``` bash
$ openconfigd &
$ cli
ubuntu>
ubuntu> show version
Developer Preview version of openconfigd
ubuntu>
```

### Options

`openconfigd` takes YANG module names as arguments.  When no YANG module is specified, default `coreswitch.yang` is used.  '.yang' suffix is optional.  Use can specify multiple YANG file.  So

``` bash
$ openconfigd lagopus ietf-ip
```

will load both `lagopus.yang` and `ietf-ip.yang` modules.

There are several other options.

*  -c, --config-file= active config file name (default: coreswitch.conf)
*  -p, --config-dir=  config file directory (default: /usr/local/etc)
*  -y, --yang-paths=  colon separated YANG load path directories
*  -2, --two-phase    enable two phase commit
*  -z, --zero-config  do not save or load config other than openconfigd.conf
*  -h, --help         Show this help message

`-c` option specify active config file name.  `-p` option specify config file save directory.  When full path is specified to `-c` option's base directory overrides the `-p` option config file directory.

`-y` option specify YANG file load path.  Use can specify multiple YANG load path with colon separated list.

`-2` option enables two phase commit.  It send validate start and end message to protocol modules.

When `-z` option is specified, only openconfigd.conf file is loaded on start up.  Configuraion is never saved to the config file.

``` bash
$ openconfigd -y /usr/shared/yang:/opt/yang
```

will search both `/usr/shared/yang` and `/opt/yang` directory.  Default YANG load path `$GOPATH/src/github.com/coreswitch/openconfigd/yang` is automatically added.

### openconfigd scripting

openconfigd support CLI scripting. All operational and configuration mode
commands can run from script.

Here is an example:

``` shell
#! /bin/bash

source /etc/bash_completion.d/cli

show version
```

