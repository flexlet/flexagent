#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_INVALID_OS=-2
ERROR_FAILED_TRUST_CERT=1
ERROR_FAILED_CONFIG_HOST=2
ERROR_FAILED_LOGIN=3
ERROR_CONFIG_ISULAD=4
ERROR_FAILED_LOCK=5

# consts
CRYPTO="/usr/local/bin/crypto-tool"
ISULAD_CONFIG="/etc/isulad/daemon.json"
SECONDS_PER_DAY=86400
CERT_VALID_DAYS=90
CERT_VALID_SECS=$(python -c "print( ${SECONDS_PER_DAY} * ${CERT_VALID_DAYS})")
DATETIME=$(date +%Y%m%d%H%M%S)

# plugin base directory
BASE_DIR=$(readlink -f $(dirname $0)/..)
TEMP_DIR="${BASE_DIR}/temp"

# lock before config
LOCK_FILE="${TEMP_DIR}/registry_cfg.lock"

# backup /etc/hosts
HOSTS_CONFIG_BKUP="${TEMP_DIR}/host.${DATETIME}"

# backup /etc/isulad/daemon.json
ISULAD_CONFIG_BKUP="${TEMP_DIR}/daemon.json.${DATETIME}"

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -a ADDR      required, repo server ip address
      -p PORT      optional, repo server port
      -n NAME      optional, repo server DNS name
      -u USER      optional, user name
      -i INSECURE  optional, http or not, default: no
    
    Envirments:
      IMAGE_REPO_CERT          required for https, repo cert (base64 text)
      IMAGE_REPO_PASS          optional, user password (plain text)
      IMAGE_REPO_CIPHER_PASS   optional, user password (cipher text)
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts a:p:n:u:i flag
    do
        case "${flag}" in
            a) ADDR=${OPTARG};;
            p) PORT=${OPTARG};;
            n) NAME=${OPTARG};;
            u) USER=${OPTARG};;
            i) INSECURE=1;;
            ?) fn_print_help
        esac
    done

    if [ "${ADDR}" == "" ]; then fn_print_help; fi
    if [ "${NAME}" != "" ]; then REPO="${NAME}"; else REPO="${ADDR}"; fi
    if [ "${PORT}" != "" ]; then REPO="${REPO}:${PORT}"; fi

    if [ "${INSECURE}" != "1" -a "${IMAGE_REPO_CERT}" == "" ]; then fn_print_help; fi

    # validate cert (at least valid in 90 days)
    if [ "${IMAGE_REPO_CERT}" != "" ]; then
        if ! printf ${IMAGE_REPO_CERT} | base64 -d | openssl x509 -checkend ${CERT_VALID_SECS} >/dev/null 2>&1; then
            echo "cert invalid or will expire in ${CERT_VALID_DAYS} days"
            exit ${ERROR_INVALID_PARAM}
        fi
    fi

    # convert cipher pass to plain pass
    if [ "${IMAGE_REPO_CIPHER_PASS}" != "" ]; then 
      IMAGE_REPO_PASS=$(printf ${IMAGE_REPO_CIPHER_PASS} | ${CRYPTO} --decrypt)
    fi

    if [ "${USER}" != "" -a "${IMAGE_REPO_PASS}" == "" ]; then fn_print_help; fi
}

# trust cert in os level
function fn_trust_cert() {
    # os dependent consts
    OS=$(awk -F '"' '/^ID=/{print $2}' /etc/os-release)
    if [ "${OS}" == "euleros" -o "${OS}" == "openEuler" -o "${OS}" == "centos" -o "${OS}" == "rhel" ]; then
        CA_TRUST_ANCHORS="/etc/pki/ca-trust/source/anchors"
        UPDATE_CA_TRUST="update-ca-trust extract"
    elif [ "${OS}" == "sles" ]; then
        CA_TRUST_ANCHORS="/etc/pki/trust/anchors"
        UPDATE_CA_TRUST="update-ca-certificates"
    else
        echo "OS '${OS}' not supported yet"
        exit ${ERROR_INVALID_OS}
    fi

    # mkdir anchors dir if not exist
    if [ ! -d "${CA_TRUST_ANCHORS}" ]; then mkdir -p ${CA_TRUST_ANCHORS}; fi

    # put cert to anchors dir and trust it
    TEMP_CA_TRUST_CERT="${CA_TRUST_ANCHORS}/${REPO}.crt"
    printf ${IMAGE_REPO_CERT} | base64 -d | openssl x509 -out ${TEMP_CA_TRUST_CERT} >/dev/null 2>&1
    ${UPDATE_CA_TRUST}
}

# set host resolv entry
function fn_config_host_resolv() {
    # backup file
    cp /etc/hosts ${HOSTS_CONFIG_BKUP}

    # replace host entry
    sed -i "/${NAME}/d" /etc/hosts
    echo "${ADDR} ${NAME}" >> /etc/hosts
}

# login repo
function fn_login_repo() {
    printf "${IMAGE_REPO_PASS}" | isula login -u ${USER} --password-stdin ${REPO}
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
if '${INSECURE}' != '1' and '${REPO}' not in cfg['registry-mirrors']:
    cfg['registry-mirrors'].append('${REPO}')
    modified = True
if '${INSECURE}' == '1' and '${REPO}' not in cfg['insecure-registries']:
    cfg['insecure-registries'].append('${REPO}')
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
    if [ -f "${HOSTS_CONFIG_BKUP}" ]; then cp ${HOSTS_CONFIG_BKUP} /etc/hosts; fi
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
    if [ -f "${TEMP_CA_TRUST_CERT}" ]; then rm ${TEMP_CA_TRUST_CERT}; fi
    if [ -f "${HOSTS_CONFIG_BKUP}" ]; then rm ${HOSTS_CONFIG_BKUP}; fi
    if [ -f "${ISULAD_CONFIG_BKUP}" ]; then rm ${ISULAD_CONFIG_BKUP}; fi 
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    # lock before config
    fn_lock

    # trust cert
    if [ "${INSECURE}" != "1" ]; then
        if ! fn_trust_cert; then
            fn_clean
            exit ERROR_FAILED_TRUST_CERT
        fi
    fi

    # config host resolv
    if [ "${NAME}" != "" ]; then
        if ! fn_config_host_resolv; then 
            fn_rollback
            fn_clean
            exit ERROR_FAILED_CONFIG_HOST
        fi
    fi

    # login repo
    if [ "${USER}" != "" ]; then
        if ! fn_login_repo; then
            fn_rollback
            fn_clean
            exit ERROR_FAILED_LOGIN
        fi   
    fi

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