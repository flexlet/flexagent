#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_CERT_EXPIRED=1

# consts
SECONDS_PER_DAY=86400
DEFAULT_EXPIRE=30

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -f FILE     required, certificate file
      -e EXPIRE   optional, expire days to check, default ${DEFAULT_EXPIRE} days
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts f:e: flag
    do
        case "${flag}" in
            f) FILE=${OPTARG};;
            e) EXPIRE=${OPTARG};;
            ?) fn_print_help
        esac
    done

    if [ "${FILE}" == "" ]; then fn_print_help; fi

    if [ ! -f ${FILE} ]; then
      echo "cert file '${FILE}' does not exist"
      exit ${ERROR_INVALID_PARAM} 
    fi

    if [ "${EXPIRE}" == "" ]; then EXPIRE=${DEFAULT_EXPIRE}; fi

}


function fn_main() {
    # validate parameters
    fn_validate_params $@

    EXPIRE_SECONDS=$(python -c "print( ${SECONDS_PER_DAY} * ${EXPIRE})")

    COMMAND="openssl x509 -checkend ${EXPIRE_SECONDS} -in ${FILE}"

    # check expired or not
    if ${COMMAND}; then
      exit 0
    else
      exit ${ERROR_CERT_EXPIRED}
    fi
}

fn_main $@