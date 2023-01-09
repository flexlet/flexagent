#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_CONFIG_FLEXLB=1

# consts
FLEXLB_CONFIG="/etc/flexlb/config.yaml"

# plugin base directory
BASE_DIR=$(readlink -f $(dirname $0)/..)
TEMP_DIR="${BASE_DIR}/temp"

# backup /etc/flexlb/config.yaml
FLEXLB_CONFIG_BKUP="${TEMP_DIR}config.yaml.${DATETIME}"

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -m MEMBERS   optional, member advertize endpoints, comma seperate (<node1_private_ip>:<port>,<node1_private_ip>:<port>,...)
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts m: flag
    do
        case "${flag}" in
            m) MEMBERS=${OPTARG};;
            ?) fn_print_help
        esac
    done

    if [ "${MEMBERS}" == "" ]; then fn_print_help; fi
}

# set members
function fn_config_flexlb() {
    # backup config
    cp ${FLEXLB_CONFIG} ${FLEXLB_CONFIG_BKUP}
    
    # update config
    sed -i "s/^  member: .*$/  member: \"${MEMBERS}\"/" ${FLEXLB_CONFIG}

    # restart service
    printf "Restart flexlb ... "
    systemctl restart flexlb
    active=$(systemctl is-active flexlb)
    echo "${active}"
    if [ "${active}" != "active" ]; then
        return ${ERROR_CONFIG_FLEXLB}
    fi
}

function fn_rollback() {
    if [ -f "${FLEXLB_CONFIG_BKUP}" ]; then 
        # restore config file
        cp ${FLEXLB_CONFIG_BKUP} ${FLEXLB_CONFIG}
        printf "Restart flexlb ... "
        systemctl restart flexlb
        systemctl is-active flexlb
    fi 
}

function fn_clean() {
    if [ -f "${FLEXLB_CONFIG_BKUP}" ]; then rm ${FLEXLB_CONFIG_BKUP}; fi 
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    # config FLEXLB
    if ! fn_config_flexlb; then
        fn_rollback
        fn_clean
        exit ERROR_CONFIG_FLEXLB
    fi

    # clean temp file
    fn_clean
}

fn_main $@