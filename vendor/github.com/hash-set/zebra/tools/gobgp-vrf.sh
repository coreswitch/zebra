#! /bin/bash

if [ $# != 4 -a $# != 5 ];then
    echo "gobgp-vrf.sh VRF_ID {add,del} A.B.C.D/M MED [COMMUNITY]"
    exit 1
fi

VRFS=$(gobgp vrf)

OIFS=${IFS}
IFS=''
while read line; do
    IFS=${OIFS}
    array=(${line})
    IFS=''

    if [ ${array[4]} = ${1} ];then
        if [ $# == 4 ];then
            if [ ${4} == 0 ];then
                gobgp vrf ${array[0]} rib ${2} ${3}
            else
                gobgp vrf ${array[0]} rib ${2} ${3} med ${4}
            fi
        else
            if [ ${4} == 0 ];then
                gobgp vrf ${array[0]} rib ${2} ${3} community ${5}
            else
                gobgp vrf ${array[0]} rib ${2} ${3} med ${4} community ${5}
            fi
        fi
    fi
done <<<${VRFS}
IFS=${OIFS}
