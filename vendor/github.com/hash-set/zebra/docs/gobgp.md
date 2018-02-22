# GoBGP

GoBGP can be used as a Zebra's BGP protocol module. To make it work, GoBGP
should be configured with appropriate configuration.

## GoBGP Configuration

By default, Zebra listen to UNIX domain socket `/var/run/zserv.api` for Zebra
API. So GoBGP Global configuration look like:

```toml
[zebra]
    [zebra.config]
        enabled = true
        url = "unix:/var/run/zserv.api"
````

No configuration is necessary to Zebra side.
