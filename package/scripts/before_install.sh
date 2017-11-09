#! /bin/bash

_ubuntu_version ()
{
    if /bin/grep 16.04 /etc/lsb-release>/dev/null;then
        U1604=1
    else
        U1604=0
    fi
}

_process_1404 ()
{
    echo "Pre install script for Ubuntu 14.04"
    /bin/mkdir -p /var/log/ribd
    /usr/bin/supervisorctl stop ribd
    return 0
}

_process_1604 ()
{
    echo "Pre install script for Ubuntu 16.04"
    /usr/bin/supervisorctl stop ribd
    return 0
}

_main ()
{
    _ubuntu_version

    if [ ${U1604} == 1 ]; then
        _process_1604
    else
        _process_1404
    fi
}

_main
