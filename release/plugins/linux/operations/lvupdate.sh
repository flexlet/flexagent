#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_VG_NOT_EXIST=1
ERROR_CREATE_LV=2
ERROR_FORMAT_FS=3
ERROR_EXTEND_LV=4

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -g VGNAME       required, vg name
      -n LVNAME       required, block device
      -i STRIPE       optional, number of stripes, default: number of active pv
      -I STRIPESIZE   optional, stripe size, default: 64
      -L SIZE         optional, lv size, default: all free space
      -t FSTYPE       optional, format filesystem
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts g:n:i:I:L:t: flag
    do
        case "${flag}" in
            g) VGNAME=${OPTARG};;
            n) LVNAME=${OPTARG};;
            i) STRIPE=${OPTARG};;
            I) STRIPESIZE=${OPTARG};;
            L) SIZE=${OPTARG};;
            t) FSTYPE=${OPTARG};;
            ?) fn_print_help
        esac
    done

    if [ "${VGNAME}" == "" ]; then fn_print_help; fi
    if [ "${LVNAME}" == "" ]; then fn_print_help; fi
    if [ "${FORCE}" != "y" ]; then FORCE="n"; fi

    # check vg exist or not
    if ! vgs ${VGNAME} > /dev/null 2>&1; then
        echo "vg does not exist"
        exit ${ERROR_VG_NOT_EXIST}
    fi

    NumPVs=$(vgdisplay ${VGNAME} | awk '/Act PV/{print $NF}')
    if [ "${STRIPE}" == "" ]; then 
         STRIPE=${NumPVs}
    elif (( ${STRIPE} > ${NumPVs} )); then
        STRIPE=${NumPVs}
    fi

    if [ "${STRIPESIZE}" == "" ]; then STRIPESIZE=64; fi

    FreePE=$(vgdisplay ${VGNAME} | awk '/Free  PE/{print $(NF-3)}')
    if [ "${SIZE}" == "" ]; then
        SizePE=${FreePE}
    else
        SizePE=$((${SIZE}*256))
    fi
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    FreePE=$(vgdisplay ${VGNAME} | awk '/Free  PE/{print $(NF-3)}')
    
    if ! lvs /dev/${VGNAME}/${LVNAME} > /dev/null 2>&1; then
        # too large size
        if (( ${SizePE} > ${FreePE} )); then
            SizePE=${FreePE}
        fi
        
        # create lv if not exist
        if ! lvcreate -n ${LVNAME} -i ${STRIPE} -I ${STRIPESIZE} -l ${SizePE} ${VGNAME} ; then
            echo "create lv failed"
            exit ${ERROR_CREATE_LV}
        fi
        
        # format fs
        if [ "${FSTYPE}" != "" ]; then
            if ! mkfs -t ${FSTYPE} /dev/${VGNAME}/${LVNAME}; then
                echo "create fs failed"
                exit ${ERROR_FORMAT_FS}
            fi
        fi
    else
        # extend lv and fs
        CurLE=$(lvdisplay /dev/${VGNAME}/${LVNAME} | awk '/Current LE/{print $NF}')
        NewPE=$((${SizePE}-${CurLE}))

        # too large size
        if (( ${NewPE} > ${FreePE} )); then
            NewPE=${FreePE}
        fi

        if (( ${NewPE} > 0 )); then
            if ! lvextend -i ${STRIPE} -I ${STRIPESIZE} -l +${NewPE} -r /dev/${VGNAME}/${LVNAME} ; then
                echo "extend lv failed"
                exit ${ERROR_EXTEND_LV}
            fi
        fi
    fi
}

fn_main $@