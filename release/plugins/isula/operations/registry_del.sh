#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_INVALID_OS=-2
ERROR_CONFIG_ISULAD=4
ERROR_FAILED_LOCK=5

# consts
ISULAD_CONFIG="/etc/isulad/daemon.json"

# plugin base directory
BASE_DIR=$(readlink -f $(dirname $0)/..)
TEMP_DIR="${BASE_DIR}/temp"
DATETIME=$(date +%Y%m%d%H%M%S)

# lock before config
LOCK_FILE="${TEMP_DIR}/registry_cfg.lock"

# backup /etc/isulad/daemon.json
ISULAD_CONFIG_BKUP="${TEMP_DIR}/daemon.json.${DATETIME}"

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -r REPO      required, repo uri
      -i INSECURE  optional, http or not, default: no
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts r:i flag
    do
        case "${flag}" in
            r) REPO=${OPTARG};;
            i) INSECURE=1;;
            ?) fn_print_help
        esac
    done

    if [ "${REPO}" == "" ]; then fn_print_help; fi
}

# set repo to isulad config file
function fn_config_isulad() {
    # backup config
    cp ${ISULAD_CONFIG} ${ISULAD_CONFIG_BKUP}
    
    # update config
    python -c "
import json
with open('${ISULAD_CONFIG}') as cfgFIle:
    cfg = json.load(cfgFIle)

modified = False
if '${INSECURE}' != '1' and '${REPO}' in cfg['registry-mirrors']:
    cfg['registry-mirrors'].remove('${REPO}')
    modified = True
if '${INSECURE}' == '1' and '${REPO}' in cfg['insecure-registries']:
    cfg['insecure-registries'].remove('${REPO}')
    modified = True

if modified:
    cfgJson = json.dumps(cfg, indent=4)
    with open('${ISULAD_CONFIG}', 'w') as cfgFIle:
        cfgFIle.write(cfgJson)
    exit(1)
"
    # restart service
    if [ "$?" == "1" ]; then 
        printf "Restart isulad ... "
        systemctl restart isulad
        systemctl is-active isulad
    fi
}

function fn_lock() {
    if [ -f "${LOCK_FILE}" ]; then exit ${ERROR_FAILED_LOCK}; fi
    if ! touch ${LOCK_FILE}; then exit ${ERROR_FAILED_LOCK}; fi
}

function fn_rollback() {
    if [ -f "${ISULAD_CONFIG_BKUP}" ]; then 
        # restore config file
        cp ${ISULAD_CONFIG_BKUP} ${ISULAD_CONFIG}
        printf "Restart isulad ... "
        systemctl restart isulad
        systemctl is-active isulad
    fi 
}

function fn_clean() {
    if [ -f "${LOCK_FILE}" ]; then rm ${LOCK_FILE}; fi
    if [ -f "${ISULAD_CONFIG_BKUP}" ]; then rm ${ISULAD_CONFIG_BKUP}; fi 
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    # lock before config
    fn_lock

    # logout repo
    isula logout ${REPO}

    # config isulad
    if ! fn_config_isulad; then
        fn_rollback
        fn_clean
        exit ERROR_CONFIG_ISULAD
    fi

    # clean temp file
    fn_clean
}

fn_main $@