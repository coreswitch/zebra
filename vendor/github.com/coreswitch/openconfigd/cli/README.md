cli
======

### openconfigd shell

`cli` is a openconfid package shell based upon `bash-4.3.30`. It takes user
input then execute command for a router and/or switch system. When it does not
match to router/switch command it execute normal command as same as standard
`bash`.

### Build & Install

To install `cli`, please enable Programmable Completion feature of `bash`.
Typical build and installation procedure is as follows.

```sh
$ cd cli
$ ./configure --enable-progcomp
$ sudo make install
```

This will build `cli` command and install under `/usr/local/bin`. To install the
executable to different location, please specify --prefix option to configure
script.

```shell
$ cd cli
$ ./configure --prefix=/usr --enable-progcomp
$ sudo make install
```

will install `cli` under `/usr/bin`.

### Completion Scripts

`cli` requires start up scripts to perform completion and execution of commands.
For system wide installation, the scripts will go under
`/etc/bash_completion.d`. For user local installation, the script will go under

```shell
$ cd 
```

* System Wide installation

```sh
$ sudo make install_scripts_system
```

* User Local Installation

```sh
$ make install_scripts_local
```

### Environmental variable

`cli` has *Environmental Variable* `CLI_MODE` which indicates that `cli` feature
is enabled. Usually this is enabled in Completion Script. User can disable
`cli` feature with unset the *Environmental Variable*

```
unset CLI_MODE
```

When `CLI_MODE` is disabled, `cli` works as normal `bash`.
