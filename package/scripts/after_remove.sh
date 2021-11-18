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
    echo "Post uninstall script for Ubuntu 14.04"
    /usr/bin/supervisorctl reload
    return 0
}

_process_1604 ()
{
    echo "Post uninstall script for Ubuntu 16.04"
    /bin/systemctl daemon-reload
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
