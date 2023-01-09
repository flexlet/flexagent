#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1

# consts
DEFAULT_PORT=22
CRYPTO="/usr/local/bin/crypto-tool"

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -h HOST    required, ssh host
      -u USER    required, ipv4 address
      -p PORT    optional, ssh port, default ${DEFAULT_PORT}
      -f FILE    optional, file to transfer
      -d DIR     optional, directory to transfer
      -t TARGET  required, transfer target
      -m METHOD  required, method (GET/PUT)
    
    Envirments:
      SCP_PASS          optional, scp password (plain text)
      SCP_CIPHER_PASS   optional, scp password (cipher text)
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts h:u:p:f:d:t:m: flag
    do
        case "${flag}" in
            h) HOST=${OPTARG};;
            u) USER=${OPTARG};;
            p) PORT=${OPTARG};;
            f) FILE=${OPTARG};;
            d) DIR=${OPTARG};;
            t) TARGET=${OPTARG};;
            m) METHOD=${OPTARG};;
            ?) fn_print_help
        esac
    done

    if [ "${HOST}" == "" ]; then fn_print_help; fi
    if [ "${USER}" == "" ]; then fn_print_help; fi
    if [ "${PORT}" == "" ]; then PORT=${DEFAULT_PORT}; fi

    SOURCE=""
    ISDIR=0

    if [ "${FILE}" != "" -a "${DIR}" == "" ]; then
      SOURCE="${FILE}"
    fi

    if [ "${FILE}" == "" -a "${DIR}" != "" ]; then
      SOURCE="${DIR}"
      ISDIR=1
    fi

    if [ "${SOURCE}" == "" ]; then fn_print_help; fi
    if [ "${TARGET}" == "" ]; then fn_print_help; fi
    if [ "${METHOD}" == "" ]; then fn_print_help; fi
    if [ "${METHOD}" != "GET" -a "${METHOD}" != "PUT" ]; then fn_print_help; fi

    # convert cipher pass to plain pass
    if [ "${SCP_CIPHER_PASS}" != "" ]; then 
      SCP_PASS=$(printf ${SCP_CIPHER_PASS} | ${CRYPTO} --decrypt)
    fi
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    COMMAND="/usr/bin/scp -o StrictHostKeyChecking=no -P ${PORT}"

    if (( $ISDIR )); then
        COMMAND="${COMMAND} -r"
    fi

    if [ "${METHOD}" == "GET" ]; then
        COMMAND="${COMMAND} ${USER}@${HOST}:${SOURCE} ${TARGET}"
    else
        COMMAND="${COMMAND} ${SOURCE} ${USER}@${HOST}:${TARGET}"
    fi

    # execute scp via expect
    expect <<____END_EXPECT
      spawn ${COMMAND}
      expect {
        password: {
          send "${SCP_PASS}\r"
          exp_continue
        }
        eof exit
      }
____END_EXPECT

}

fn_main $@