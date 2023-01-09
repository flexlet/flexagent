#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1

# consts
DEFAULT_INTERVAL=1
DEFAULT_COUNT=1

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -a IPADDR   required, ipv4 address
      -i INTERVAL optional, interval, default ${DEFAULT_INTERVAL}
      -c COUNT    optional, count, default ${DEFAULT_COUNT}
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts a:i:c: flag
    do
        case "${flag}" in
            a) IPADDR=${OPTARG};;
            i) INTERVAL=${OPTARG};;
            c) COUNT=${OPTARG};;
            ?) fn_print_help
        esac
    done

    if [ "${IPADDR}" == "" ]; then fn_print_help; fi

    local IPV4_PATTERN='^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$'
    
    # check ipaddr pattern
    local MATCH=$(echo ${IPADDR} | awk "{if(/${IPV4_PATTERN}/){print 1}else{print 0}}")    
    if (( ${MATCH} == 0 )); then fn_print_help; fi

    if [ "${INTERVAL}" == "" ]; then INTERVAL=${DEFAULT_INTERVAL}; fi
    if [ "${COUNT}" == "" ]; then COUNT=${DEFAULT_COUNT}; fi
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    # ping
    /usr/bin/ping ${IPADDR} -i ${INTERVAL} -c ${COUNT}
}

fn_main $@