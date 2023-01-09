#!/bin/bash

SERVICE="sshd"
COMMAND="journalctl -u ${COMMAND}"

function fn_print_help() {
    echo "Help: $(basename $0) -s SINCE -u UNTIL
    Options:
      -s SINCE     optional, print logs since time, format: 'N hours ago' or 'yyyy-mm-dd HH24:MM:SS'
      -u UNTIL     optional, print logs until time, format: 'N hours ago' or 'yyyy-mm-dd HH24:MM:SS'
    "

    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts s:u: flag
    do
        case "${flag}" in
            s) COMMAND="${COMMAND} --since '${OPTARG}'";;
            u) COMMAND="${COMMAND} --until '${OPTARG}'";;
            ?) fn_print_help
        esac
    done
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    ${COMMAND}
}

fn_main $@