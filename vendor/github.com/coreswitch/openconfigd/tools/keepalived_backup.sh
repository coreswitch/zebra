#! /bin/bash

TYPE=$1
NAME=$2
STATE=$3

IFS='_. ' read -r -a array <<<${0}
VRF=${array[2]}
CIDR=`awk -F\- '{print $NF}' <<<$NAME`

case $STATE in
    "MASTER") gobgp_local 120 ${NAME} ${STATE} ${VRF} ${CIDR}
              gobgp neighbor all softresetout
              exit 0
              ;;
    "BACKUP") gobgp_local 50 ${NAME} ${STATE} ${VRF} ${CIDR}
              gobgp neighbor all softresetout
              exit 0
              ;;
    "FAULT")  gobgp_local 50 ${NAME} ${STATE} ${VRF} ${CIDR}
              gobgp neighbor all softresetout
              exit 0
              ;;
    *)        echo "unknown state"
              exit 1
              ;;
esac
