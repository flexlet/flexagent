#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_DEVICE_NOT_EXIST=1
ERROR_VG_EXIST=2
ERROR_CREATE_PV=3
ERROR_CREATE_VG=4
ERROR_EXTEND_VG=5
ERROR_PV_NOT_EXIST=6

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -g VGNAME   required, vg name
      -d DEVICES  required, block device (multiple)
      -f          optional, force or not, default: false
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts g:d:f flag
    do
        case "${flag}" in
            g) VGNAME=${OPTARG};;
            d) DEVICES="${DEVICES}${OPTARG} ";;
            f) FORCE="y";;
            ?) fn_print_help
        esac
    done

    if [ "${VGNAME}" == "" ]; then fn_print_help; fi
    if [ "${DEVICES}" == "" ]; then fn_print_help; fi
    if [ "${FORCE}" != "y" ]; then FORCE="n"; fi

    # check device exist or not
    for device in ${DEVICES}; do
        if [ ! -e "${device}" ]; then 
            echo "device not exist: ${device}"
            exit ${ERROR_DEVICE_NOT_EXIST}
        fi
    done
}

function fn_create_pv() {
    BLOCK_DEVICES=""
    for device in ${DEVICES}; do
        EXIST=$(pvs ${device} 2>&1 | awk "{if(\$2==\"${VGNAME}\"){print 1}}" )
        # bypass exist pv
        if [ "${EXIST}" == "1" ]; then
            continue
        fi
        
        # create pv
        echo "${force}" | pvcreate ${device}
        if [ "$?" != "0" ]; then exit ${ERROR_CREATE_PV}; fi

        # get block device
        BLOCK_DEVICE=$(pvdisplay ${device} 2>&1 | awk '/PV Name/{print $NF}')
        if [ "$?" == "" ]; then exit ${ERROR_PV_NOT_EXIST}; fi
        BLOCK_DEVICES="${BLOCK_DEVICES}${BLOCK_DEVICE} "
    done
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    # create non exit pv
    fn_create_pv

    if ! vgs ${VGNAME} > /dev/null 2>&1; then
        # create vg
        if ! vgcreate ${VGNAME} ${BLOCK_DEVICES}; then
            echo "create vg failed"
            exit ${ERROR_CREATE_VG}
        fi
    else
        # extend vg
        for device in ${BLOCK_DEVICES}; do
            EXIST=$(pvs ${device} 2>&1 | awk "{if(\$2==\"${VGNAME}\"){print 1}}" )
            if [ "${EXIST}" != "1" ]; then
                if ! vgextend ${VGNAME} ${device}; then
                    echo "extend vg failed"
                    exit ${ERROR_EXTEND_VG}
                fi
            fi
        done
    fi
}

fn_main $@