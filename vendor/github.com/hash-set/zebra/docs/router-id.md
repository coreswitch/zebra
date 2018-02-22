# Router ID

Router ID is unique identification ID for the router in the autonomous system.
Due to the historical reason one of the IPv4 address belong to the router is
used for the purpose.

## Router ID decision logic

When Router ID is configured manually. The value is used as Router ID.

When IPv4 address is assigned to the loopback interface (other than 127.0.0.1 or
IP address belongs to network 127.0.0.0/8), the address is used as a Router ID.
When loopback has multiple IPv4 adderesses, the biggest IPv4 address among
addresses is used as a Router ID.

Otherwise, biggest IP Address among all of the interfaces is used as a Router
ID.

Each VRF has separate Router ID. VRF's Router ID selection does not affect to
other VRF.

## Router ID 

You can check current Router ID with `show router-id` command in CLI.

```shell
zebra>show router-id 
Router ID: 192.168.55.2 (automatic)
```
