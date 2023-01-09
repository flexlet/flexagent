#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_MOUNT_FAILED=1
ERROR_INJECT_FAILED=1

# consts
FSTAB="/etc/fstab"
DATETIME=$(date +%Y%m%d%H%M%S)

# plugin base directory
BASE_DIR=$(readlink -f $(dirname $0)/..)
TEMP_DIR="${BASE_DIR}/temp"
TEMP_FSTAB="${TEMP_DIR}/fstab-${DATETIME}"

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -d DEVICE   required, device or nfs path
      -m MOUNT    required, mount point
      -t FSTYPE   required, fs type
      -o OPTIONS  optional, mount options, default: 'defaults'
      -D          optional, dump or not, default: 0
      -C          optional, check or not, default: 0
    
    Enviorments:
      MOUNT_INJECT_DATA   optional, inject data after mount, format: <file1>:<base64 text>,<file2>:<base64 text>...
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts d:m:t:o:DC flag
    do
        case "${flag}" in
            d) DEVICE=${OPTARG};;
            m) MOUNT=${OPTARG};;
            t) FSTYPE=${OPTARG};;
            o) OPTIONS=${OPTARG};;
            D) DUMP=1;;
            C) CHECK=1;;
            ?) fn_print_help
        esac
    done

    if [ "${DEVICE}" == "" ]; then fn_print_help; fi
    if [ "${MOUNT}" == "" ]; then fn_print_help; fi
    if [ "${FSTYPE}" == "" ]; then fn_print_help; fi
    if [ "${OPTIONS}" == "" ]; then OPTIONS="defaults"; fi
    if [ "${DUMP}" != "1" ]; then DUMP=0; fi
    if [ "${CHECK}" != "1" ]; then CHECK=0; fi
}

function fn_mount() {
    # backup fstab
    cp ${FSTAB} ${TEMP_FSTAB}

    # remove old mount points
    cat ${TEMP_FSTAB} | awk -v M=${MOUNT} '{if($2!=M){print}}' > ${FSTAB}

    # add new mount points
    echo "${DEVICE} ${MOUNT} ${FSTYPE} ${OPTIONS} ${DUMP} ${CHECK}" >> ${FSTAB}

    # create mount point directory
    if [ ! -f "${MOUNT}" ]; then mkdir -p ${MOUNT}; fi

    # mount
    if mount -a; then
        # mount succeeded, remove backup
        rm ${TEMP_FSTAB}
    else
        # mount failed, restore backup
        mv ${TEMP_FSTAB} ${FSTAB}
        mount -a
        echo "mount failed"
        exit ${ERROR_MOUNT_FAILED}
    fi
}

function fn_inject_data() {
    if [ "${MOUNT_INJECT_DATA}" != "" ]; then
        for KV in $(echo "${MOUNT_INJECT_DATA}" | sed 's/,/ /g' ); do
            KV=($(echo "${KV}" | sed 's/:/ /g'))
            printf "${KV[1]}" | base64 -d > ${MOUNT}/${KV[0]}
            if [ "$?" != "0" ]; then
                echo "inject data failed: ${KV[0]}"
                exit ${ERROR_INJECT_FAILED}
            fi
        done
    fi
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    fn_mount

    fn_inject_data

}

fn_main $@