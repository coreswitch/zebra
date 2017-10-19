package quagga

var showCmdMap = map[string]func(string) *string{
	"quagga_show": quaggaShow,
}

var execCmdMap = map[string]func(string) *string{
	"quagga_exec": quaggaExec,
}

func quaggaShow(line string) *string {
	return quaggaVtysh(line)
}

func quaggaExec(line string) *string {
	return quaggaVtysh(line)
}

const showCmdSpec = `
[
    {
        "name": "quagga_show",
        "line": "show running-config",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "running configuration"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp neighbors (A.B.C.D|X:X::X:X) routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show route-map [WORD]",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "route-map information",
            "route-map name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities",
            "",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the communities",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community-list (<1-500>|WORD)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the community-list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community-list (<1-500>|WORD)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the community-list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp community-list (<1-500>|WORD) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the community-list",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 community-list (<1-500>|WORD) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the community-list",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp filter-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes conforming to the filter-list",
            "Regular expression access list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 filter-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes conforming to the filter-list",
            "Regular expression access list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv4 (unicast|multicast) rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv4 (unicast|multicast) rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 (unicast|multicast) rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 (unicast|multicast) rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 (unicast|multicast) summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors (A.B.C.D|X:X::X:X) prefix-counts",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display detailed prefix count information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 (unicast|multicast)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 (unicast|multicast) X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 (unicast|multicast) X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 (unicast|multicast) summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp memory",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Global BGP memory statistics"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors (A.B.C.D|X:X::X:X) received prefix-filter",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display information received from a BGP neighbor",
            "Display the prefixlist filter"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp neighbors (A.B.C.D|X:X::X:X) received prefix-filter",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display information received from a BGP neighbor",
            "Display the prefixlist filter"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "IPv6 prefix <network>/<length>"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "IPv6 prefix <network>/<length>"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes conforming to the prefix-list",
            "IPv6 prefix-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes conforming to the prefix-list",
            "IPv6 prefix-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp X:X::X:X/M longer-prefixes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "IPv6 prefix <network>/<length>",
            "Display route and more specific routes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 X:X::X:X/M longer-prefixes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "IPv6 prefix <network>/<length>",
            "Display route and more specific routes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 regexp .LINE",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the AS path regular expression",
            "A regular-expression to match the BGP AS paths"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp regexp .LINE",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the AS path regular expression",
            "A regular-expression to match the BGP AS paths"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 route-map WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Display routes matching the route-map",
            "A route-map to match on"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp route-map WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Display routes matching the route-map",
            "A route-map to match on"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp (ipv4) (vpnv4) statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "",
            "",
            "BGP RIB advertisement statistics"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp (ipv4|ipv6) (unicast|multicast) statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "",
            "",
            "BGP RIB advertisement statistics"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD (ipv4) (vpnv4) statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "Address family"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD (ipv4|ipv6) (unicast|multicast) statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "Address family"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD (ipv4|ipv6) (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD (ipv4|ipv6) (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD (ipv4|ipv6) (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD (ipv4|ipv6) (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD (ipv4|ipv6) (unicast|multicast) community",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD (ipv4|ipv6) (unicast|multicast) neighbors (A.B.C.D|X:X::X:X) (advertised-routes|received-routes)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "",
            "",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv4 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Client"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv4 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Client"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv4 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X) A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Client",
            "",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv4 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X) A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Client",
            "",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv4 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X) A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Client",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv4 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X) A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Client",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Client"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Client"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X) X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Client",
            "",
            "IP prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X) X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Client",
            "",
            "IP prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X) X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Client",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 (unicast|multicast) rsclient (A.B.C.D|X:X::X:X) X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Client",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors (A.B.C.D|X:X::X:X) dampened-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the dampened routes received from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp neighbors (A.B.C.D|X:X::X:X) dampened-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the dampened routes received from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 neighbors (A.B.C.D|X:X::X:X) dampened-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the dampened routes received from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD neighbors (A.B.C.D|X:X::X:X) dampened-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the dampened routes received from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors (A.B.C.D|X:X::X:X) flap-statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display flap statistics of the routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp neighbors (A.B.C.D|X:X::X:X) flap-statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display flap statistics of the routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 neighbors (A.B.C.D|X:X::X:X) flap-statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display flap statistics of the routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD neighbors (A.B.C.D|X:X::X:X) flap-statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display flap statistics of the routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 neighbors (A.B.C.D|X:X::X:X) received prefix-filter",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display information received from a BGP neighbor",
            "Display the prefixlist filter"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD neighbors (A.B.C.D|X:X::X:X) received prefix-filter",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display information received from a BGP neighbor",
            "Display the prefixlist filter"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors (A.B.C.D|X:X::X:X) routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp neighbors (A.B.C.D|X:X::X:X) routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 neighbors (A.B.C.D|X:X::X:X) routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD neighbors (A.B.C.D|X:X::X:X) routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp neighbors (A.B.C.D|X:X::X:X) routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "IPv6 prefix <network>/<length>"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "IPv6 prefix <network>/<length>"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp rsclient (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Information about Route Server Client"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD rsclient (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Information about Route Server Client"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp rsclient (A.B.C.D|X:X::X:X) X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Information about Route Server Client",
            "",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD rsclient (A.B.C.D|X:X::X:X) X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Information about Route Server Client",
            "",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp rsclient (A.B.C.D|X:X::X:X) X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Information about Route Server Client",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD rsclient (A.B.C.D|X:X::X:X) X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Information about Route Server Client",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp views",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Show the defined BGP views"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show debugging bgp",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Debugging functions (see also 'undebug')",
            "BGP information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show debugging ospf",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Debugging functions (see also 'undebug')",
            "OSPF information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show history",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Display the session command history"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip access-list",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "List IP access lists"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip access-list (<1-99>|<100-199>|<1300-1999>|<2000-2699>|WORD)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "List IP access lists"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip as-path-access-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "List AS path access lists",
            "AS path access list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip as-path-access-list",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "List AS path access lists"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp attribute-info",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "List all bgp attribute information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp cidr-only",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display only routes with non-natural netmasks"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community-info",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "List all bgp community information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community-list (<1-500>|WORD)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the community-list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp community-list (<1-500>|WORD) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the community-list",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp dampened-paths",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display paths suppressed due to dampening"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp filter-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes conforming to the filter-list",
            "Regular expression access list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics cidr-only",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes",
            "Display only routes with non-natural netmasks"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics filter-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes",
            "Display routes conforming to the filter-list",
            "Regular expression access list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes",
            "Display routes conforming to the prefix-list",
            "IP prefix-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics A.B.C.D/M longer-prefixes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8",
            "Display route and more specific routes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics regexp .LINE",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes",
            "Display routes matching the AS path regular expression",
            "A regular-expression to match the BGP AS paths"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics route-map WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes",
            "Display routes matching the route-map",
            "A route-map to match on"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp flap-statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display flap statistics of routes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD ipv4 (unicast|multicast) rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv4 (unicast|multicast) summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD ipv4 (unicast|multicast) summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD ipv6 neighbors (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp view WORD neighbors (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD neighbors (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv4 (unicast|multicast)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) cidr-only",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display only routes with non-natural netmasks"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities",
            "",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the communities",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community-list (<1-500>|WORD)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the community-list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) community-list (<1-500>|WORD) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the community-list",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) filter-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes conforming to the filter-list",
            "Regular expression access list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) neighbors (A.B.C.D|X:X::X:X) prefix-counts",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display detailed prefix count information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) neighbors (A.B.C.D|X:X::X:X) received prefix-filter",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display information received from a BGP neighbor",
            "Display the prefixlist filter"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) neighbors (A.B.C.D|X:X::X:X) routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) paths",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Path information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv4 (unicast|multicast) A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes conforming to the prefix-list",
            "IP prefix-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) A.B.C.D/M longer-prefixes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8",
            "Display route and more specific routes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) regexp .LINE",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the AS path regular expression",
            "A regular-expression to match the BGP AS paths"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv4 (unicast|multicast) A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) route-map WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Display routes matching the route-map",
            "A route-map to match on"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv4 (unicast|multicast) summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors (A.B.C.D|X:X::X:X) dampened-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the dampened routes received from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors (A.B.C.D|X:X::X:X) flap-statistics",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display flap statistics of the routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors (A.B.C.D|X:X::X:X) prefix-counts",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display detailed prefix count information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors (A.B.C.D|X:X::X:X) received prefix-filter",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display information received from a BGP neighbor",
            "Display the prefixlist filter"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors (A.B.C.D|X:X::X:X) routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn neighbors",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information for a route distinguisher",
            "VPN Route Distinguisher",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp ipv6 neighbors (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Address family",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show bgp neighbors (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp ipv4 (unicast|multicast) neighbors (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all neighbors A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "Detailed information on TCP and BGP neighbor connections",
            "Neighbor to display information about"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn neighbors A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "Detailed information on TCP and BGP neighbor connections",
            "Neighbor to display information about"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp paths",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Path information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes conforming to the prefix-list",
            "IP prefix-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp A.B.C.D/M longer-prefixes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8",
            "Display route and more specific routes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp regexp .LINE",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the AS path regular expression",
            "A regular-expression to match the BGP AS paths"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp route-map WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the route-map",
            "A route-map to match on"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp rsclient summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Information about Route Server Clients",
            "Summary of all Route Server Clients"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp scan",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP scan status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp scan detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP scan status",
            "More detailed output"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD neighbors (A.B.C.D|X:X::X:X) advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD neighbors (A.B.C.D|X:X::X:X) received-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Detailed information on TCP and BGP neighbor connections",
            "",
            "Display the received routes from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp rsclient (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Information about Route Server Client"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD rsclient (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Information about Route Server Client"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp rsclient (A.B.C.D|X:X::X:X) A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Information about Route Server Client",
            "",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD rsclient (A.B.C.D|X:X::X:X) A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Information about Route Server Client",
            "",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp rsclient (A.B.C.D|X:X::X:X) A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Information about Route Server Client",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp view WORD rsclient (A.B.C.D|X:X::X:X) A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "BGP view",
            "View name",
            "Information about Route Server Client",
            "",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all neighbors A.B.C.D advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "Detailed information on TCP and BGP neighbor connections",
            "Neighbor to display information about",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all neighbors A.B.C.D routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "Detailed information on TCP and BGP neighbor connections",
            "Neighbor to display information about",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all tags",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information about all VPNv4 NLRIs",
            "Display BGP tags for prefixes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 all neighbors (A.B.C.D|X:X::X:X) prefix-counts",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Address family",
            "Address Family modifier",
            "Address Family modifier",
            "",
            "Neighbor to display information about"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information for a route distinguisher",
            "VPN Route Distinguisher"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn neighbors A.B.C.D advertised-routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information for a route distinguisher",
            "VPN Route Distinguisher",
            "Detailed information on TCP and BGP neighbor connections",
            "Neighbor to display information about",
            "Display the routes advertised to a BGP neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn neighbors A.B.C.D routes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information for a route distinguisher",
            "VPN Route Distinguisher",
            "Detailed information on TCP and BGP neighbor connections",
            "Neighbor to display information about",
            "Display routes learned from neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information for a route distinguisher",
            "VPN Route Distinguisher",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information for a route distinguisher",
            "VPN Route Distinguisher",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information for a route distinguisher",
            "VPN Route Distinguisher",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip bgp vpnv4 rd ASN:nn_or_IP-address:nn tags",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display VPNv4 NLRI specific information",
            "Display information for a route distinguisher",
            "VPN Route Distinguisher",
            "Display BGP tags for prefixes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip community-list",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "List community-list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip community-list (<1-500>|WORD)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "List community-list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip extcommunity-list",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "List extended-community list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip extcommunity-list (<1-500>|WORD)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "List extended-community list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf border-routers",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "show all the ABR's and ASBR's",
            "for this area"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf database",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Database summary"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf database (asbr-summary|external|network|router|summary|nssa-external|opaque-link|opaque-area|opaque-as) A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Database summary",
            "",
            "Link State ID (as an IP address)"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf database (asbr-summary|external|network|router|summary|nssa-external|opaque-link|opaque-area|opaque-as) A.B.C.D (self-originate|)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Database summary",
            "",
            "Link State ID (as an IP address)"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf database (asbr-summary|external|network|router|summary|nssa-external|opaque-link|opaque-area|opaque-as) A.B.C.D adv-router A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Database summary",
            "",
            "Link State ID (as an IP address)",
            "Advertising Router link states",
            "Advertising Router (as an IP address)"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf database (asbr-summary|external|network|router|summary|nssa-external|opaque-link|opaque-area|opaque-as|max-age|self-originate)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Database summary"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf database (asbr-summary|external|network|router|summary|nssa-external|opaque-link|opaque-area|opaque-as) (self-originate|)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Database summary"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf database (asbr-summary|external|network|router|summary|nssa-external|opaque-link|opaque-area|opaque-as) adv-router A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Database summary",
            "",
            "Advertising Router link states",
            "Advertising Router (as an IP address)"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf interface [INTERFACE]",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Interface information",
            "Interface name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf neighbor",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Neighbor list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf neighbor all",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Neighbor list",
            "include down status neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf neighbor detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Neighbor list",
            "detail of all neighbors"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf neighbor detail all",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Neighbor list",
            "detail of all neighbors",
            "include down status neighbor"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf neighbor A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Neighbor list",
            "Neighbor ID"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf neighbor IFNAME",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Neighbor list",
            "Interface name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf neighbor IFNAME detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "Neighbor list",
            "Interface name",
            "detail of all neighbors"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip ospf route",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "OSPF information",
            "OSPF routing table"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Detail of prefix lists"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list detail WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Detail of prefix lists",
            "Name of a prefix list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list WORD seq <1-4294967295>",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list",
            "sequence number of an entry",
            "Sequence number"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list WORD A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list WORD A.B.C.D/M first-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8",
            "First matched prefix"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list WORD A.B.C.D/M longer",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8",
            "Lookup longer prefix"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Summary of prefix lists"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ip prefix-list summary WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "Build a prefix list",
            "Summary of prefix lists",
            "Name of a prefix list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 access-list",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "List IPv6 access lists"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 access-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "List IPv6 access lists",
            "IPv6 zebra access-list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the communities",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the community-list",
            "community-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp community-list WORD exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the community-list",
            "community-list name",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp filter-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes conforming to the filter-list",
            "Regular expression access list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Display routes matching the prefix-list",
            "IPv6 prefix-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp X:X::X:X/M longer-prefixes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16",
            "Display route and more specific routes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp regexp .LINE",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the AS path regular expression",
            "A regular-expression to match the BGP AS paths"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Network in the BGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 bgp summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "BGP information",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "MBGP information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities",
            "",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities",
            "",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community (AA:NN|local-AS|no-advertise|no-export) (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities",
            "",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community (AA:NN|local-AS|no-advertise|no-export) exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the communities",
            "",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the community-list",
            "community-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp community-list WORD exact-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the community-list",
            "community-list name",
            "Exact match of the communities"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp filter-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes conforming to the filter-list",
            "Regular expression access list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "MBGP information",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Display routes matching the prefix-list",
            "IPv6 prefix-list name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp X:X::X:X/M longer-prefixes",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16",
            "Display route and more specific routes"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp regexp .LINE",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "BGP information",
            "Display routes matching the AS path regular expression",
            "A regular-expression to match the MBGP AS paths"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp X:X::X:X",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IP information",
            "MBGP information",
            "Network in the MBGP routing table to display"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 mbgp summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "MBGP information",
            "Summary of BGP neighbor status"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 area A.B.C.D spf tree",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Area information",
            "Area ID (as an IPv4 notation)",
            "Shortest Path First caculation",
            "Show SPF tree"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 border-routers",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display routing table for ABR and ASBR"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 border-routers (A.B.C.D|detail)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display routing table for ABR and ASBR"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database adv-router A.B.C.D linkstate-id A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Search by Advertising Router",
            "Specify Advertising Router as IPv4 address notation",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database adv-router A.B.C.D linkstate-id A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Search by Advertising Router",
            "Specify Advertising Router as IPv4 address notation",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database * A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Any Link state Type",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database * A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Any Link state Type",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database linkstate-id A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database linkstate-id A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database * A.B.C.D A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Any Link state Type",
            "Specify Link state ID as IPv4 address notation",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database * A.B.C.D A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Any Link state Type",
            "Specify Link state ID as IPv4 address notation",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database * * A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Any Link state Type",
            "Any Link state ID",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database * * A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Any Link state Type",
            "Any Link state ID",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database adv-router A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Search by Advertising Router",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database adv-router A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "Search by Advertising Router",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database self-originated",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Self-originated LSAs"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database self-originated (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Self-originated LSAs",
            "Display details of LSAs"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) adv-router A.B.C.D linkstate-id A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Search by Advertising Router",
            "Specify Advertising Router as IPv4 address notation",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) adv-router A.B.C.D linkstate-id A.B.C.D (dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Search by Advertising Router",
            "Specify Advertising Router as IPv4 address notation",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) linkstate-id A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) linkstate-id A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) A.B.C.D A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Specify Link state ID as IPv4 address notation",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) A.B.C.D A.B.C.D (dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Specify Link state ID as IPv4 address notation",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) A.B.C.D self-originated",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Specify Link state ID as IPv4 address notation",
            "Display Self-originated LSAs"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) A.B.C.D self-originated (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Display Self-originated LSAs",
            "Search by Link state ID"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) * A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Any Link state ID",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) * A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Any Link state ID",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) adv-router A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Search by Advertising Router",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) adv-router A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Search by Advertising Router",
            "Specify Advertising Router as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) self-originated",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Display Self-originated LSAs"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) self-originated (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Display Self-originated LSAs"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) self-originated linkstate-id A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Display Self-originated LSAs",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 database (router|network|inter-prefix|inter-router|as-external|group-membership|type-7|link|intra-prefix) self-originated linkstate-id A.B.C.D (detail|dump|internal)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display Link state database",
            "",
            "Display Self-originated LSAs",
            "Search by Link state ID",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 interface",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Interface infomation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 interface IFNAME",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Interface infomation",
            "Interface name(e.g. ep0)"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 interface IFNAME prefix",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Interface infomation",
            "Interface name(e.g. ep0)",
            "Display connected prefixes to advertise"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 interface IFNAME prefix (X:X::X:X|X:X::X:X/M|detail)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Interface infomation",
            "Interface name(e.g. ep0)",
            "Display connected prefixes to advertise"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 interface IFNAME prefix X:X::X:X/M (match|detail)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Interface infomation",
            "Interface name(e.g. ep0)",
            "Display connected prefixes to advertise",
            "Display the route"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 interface prefix",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Interface infomation",
            "Display connected prefixes to advertise"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 interface prefix (X:X::X:X|X:X::X:X/M|detail)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Interface infomation",
            "Display connected prefixes to advertise"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 interface prefix X:X::X:X/M (match|detail)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Interface infomation",
            "Display connected prefixes to advertise",
            "Display the route"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 linkstate",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display linkstate routing table"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 linkstate network A.B.C.D A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display linkstate routing table",
            "Display Network Entry",
            "Specify Router ID as IPv4 address notation",
            "Specify Link state ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 linkstate router A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display linkstate routing table",
            "Display Router Entry",
            "Specify Router ID as IPv4 address notation"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 linkstate detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Display linkstate routing table"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 neighbor",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Neighbor list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 neighbor (detail|drchoice)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Neighbor list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 redistribute",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "redistributing External information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 route",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Routing Table"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 route (X:X::X:X|X:X::X:X/M|detail|summary)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Routing Table"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 route (intra-area|inter-area|external-1|external-2)",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Routing Table"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 route X:X::X:X/M longer",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Routing Table",
            "Specify IPv6 prefix",
            "Display routes longer than the specified route"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 route X:X::X:X/M match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Routing Table",
            "Specify IPv6 prefix",
            "Display routes which match the specified route"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 route X:X::X:X/M longer detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Routing Table",
            "Specify IPv6 prefix",
            "Display routes longer than the specified route",
            "Detailed information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 route X:X::X:X/M match detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Routing Table",
            "Specify IPv6 prefix",
            "Display routes which match the specified route",
            "Detailed information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 route (intra-area|inter-area|external-1|external-2) detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Routing Table",
            "",
            "Detailed information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 simulate spf-tree A.B.C.D area A.B.C.D",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Shortest Path First caculation",
            "Show SPF tree",
            "Specify root's router-id to calculate another router's SPF tree"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 ospf6 spf tree",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 Information",
            "Open Shortest Path First (OSPF) for IPv6",
            "Shortest Path First caculation",
            "Show SPF tree"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list detail",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Detail of prefix lists"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list detail WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Detail of prefix lists",
            "Name of a prefix list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list WORD seq <1-4294967295>",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list",
            "sequence number of an entry",
            "Sequence number"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list WORD X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list WORD X:X::X:X/M first-match",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16",
            "First matched prefix"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list WORD X:X::X:X/M longer",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16",
            "Lookup longer prefix"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list summary",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Summary of prefix lists"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show ipv6 prefix-list summary WORD",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "IPv6 information",
            "Build a prefix list",
            "Summary of prefix lists",
            "Name of a prefix list"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show logging",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Show current logging configuration"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory all",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "All memory statistics"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory babel",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "Babel memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory bgp",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "BGP memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory isis",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "ISIS memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory lib",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "Library memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory ospf",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "OSPF memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory ospf6",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "OSPF6 memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory pim",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "PIM memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory rip",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "RIP memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory ripng",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "RIPng memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show memory zebra",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Memory statistics",
            "Zebra memory"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show mpls-te interface [INTERFACE]",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "MPLS-TE information",
            "Interface information",
            "Interface name"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show mpls-te router",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "MPLS-TE information",
            "Router information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show startup-config",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Contentes of startup configuration"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show thread cpu [FILTER]",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Thread information",
            "Thread CPU usage",
            "Display filter (rwtexb)"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show version",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Displays zebra version"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show version ospf6",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Displays ospf6d version"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show work-queues",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Work Queue information"
        ]
    },
    {
        "name": "quagga_show",
        "line": "show zebra",
        "mode": "exec",
        "helps": [
            "Show running system information",
            "Zebra information"
        ]
    }
]
`

const execCmdSpec = `
[
    {
        "name": "quagga_exec",
        "line": "clear bgp * in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 * in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 view WORD * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp view WORD * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp view WORD * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp * in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 * in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp view WORD * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp * out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 * out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp view WORD * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp <1-4294967295> in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 <1-4294967295> in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp <1-4294967295> soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 <1-4294967295> soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp <1-4294967295> in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp <1-4294967295> soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 <1-4294967295> in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 <1-4294967295> soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp <1-4294967295> out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp <1-4294967295> soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 <1-4294967295> out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 <1-4294967295> soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp external in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 external in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp external soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 external soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp external in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp external soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 external WORD in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 external soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp external out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp external soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 external WORD out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 external soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 peer-group WORD in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp peer-group WORD in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 peer-group WORD soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp peer-group WORD soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 peer-group WORD in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 peer-group WORD soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp peer-group WORD in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp peer-group WORD soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 peer-group WORD out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 peer-group WORD soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp peer-group WORD out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp peer-group WORD soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig inbound update",
            "Push out the existing ORF prefix-list"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out the existing ORF prefix-list"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 view WORD (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "BGP view",
            "view name",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp view WORD (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all peers"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all peers"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp view WORD *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD *",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * vpnv4 unicast soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * vpnv4 unicast in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * vpnv4 unicast soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * vpnv4 unicast out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp * vpnv4 unicast soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp <1-4294967295>",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear peers with the AS number"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 <1-4294967295>",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear peers with the AS number"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295>",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family modifier",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family modifier",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family modifier",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp <1-4294967295> vpnv4 unicast soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear peers with the AS number",
            "Address family",
            "Address Family modifier",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp dampening",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear route flap dampening information"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp dampening A.B.C.D",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear route flap dampening information",
            "Network to clear damping information"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp dampening A.B.C.D A.B.C.D",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear route flap dampening information",
            "Network to clear damping information",
            "Network mask"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp dampening A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear route flap dampening information",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp external",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all external peers"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 external",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all external peers"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp external soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all external peers",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all peers",
            "Address family",
            "Address Family modifier",
            "Address Family modifier"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD * ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "Clear all peers",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp (A.B.C.D|X:X::X:X)",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp ipv6 peer-group WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Address family",
            "Clear all members of peer-group",
            "BGP peer-group name"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear bgp peer-group WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out prefix-list ORF and do inbound soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp peer-group WORD soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "Clear all members of peer-group",
            "BGP peer-group name",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig inbound update",
            "Push out the existing ORF prefix-list"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) in prefix-filter",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig inbound update",
            "Push out the existing ORF prefix-list"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D ipv4 (unicast|multicast) soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp view WORD (A.B.C.D|X:X::X:X) rsclient",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP view",
            "view name",
            "",
            "Soft reconfig for rsclient RIB"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast soft",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast soft in",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig",
            "Soft reconfig inbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip bgp A.B.C.D vpnv4 unicast soft out",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "BGP information",
            "BGP neighbor address to clear",
            "Address family",
            "Address Family Modifier",
            "Soft reconfig",
            "Soft reconfig outbound update"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip prefix-list",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "Build a prefix list"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ip prefix-list WORD A.B.C.D/M",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IP information",
            "Build a prefix list",
            "Name of a prefix list",
            "IP prefix <network>/<length>, e.g., 35.0.0.0/8"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ipv6 prefix-list",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IPv6 information",
            "Build a prefix list"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ipv6 prefix-list WORD",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear ipv6 prefix-list WORD X:X::X:X/M",
        "mode": "exec",
        "helps": [
            "Reset functions",
            "IPv6 information",
            "Build a prefix list",
            "Name of a prefix list",
            "IPv6 prefix <network>/<length>, e.g., 3ffe::/16"
        ]
    },
    {
        "name": "quagga_exec",
        "line": "clear thread cpu [FILTER]",
        "mode": "exec",
        "helps": [
            "Clear stored data",
            "Thread information",
            "Thread CPU usage",
            "Display filter (rwtexb)"
        ]
    }
]
`
