#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -f FILE   required, file to write
      -m MODE   required, write mode: CREATE or APPEND
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts f:m: flag
    do
        case "${flag}" in
            f) FILE=${OPTARG};;
            m) MODE=${OPTARG};;
            ?) fn_print_help
        esac
    done

    if [ "${FILE}" == "" ]; then fn_print_help; fi
    if [ "${MODE}" == "" ]; then fn_print_help; fi
    if [ "${MODE}" != "CREATE" -a "${MODE}" != "APPEND" ]; then fn_print_help; fi

    return 0
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    # remove if exist
    if [ "${MODE}" == "CREATE" -a -f ${FILE} ]; then
        rm ${FILE}
    fi

    # read from input, write to file
    while true; do
        printf "\n> "
        read LINE
        if [ "$LINE" == "EOF" ]; then
            break
        fi
        echo ${LINE} >> ${FILE}
    done
}

fn_main $@