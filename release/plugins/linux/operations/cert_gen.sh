#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_FAILED_GEN=1

# consts
DEFAULT_EXPIRE=3650
DEFAULT_KEYLEN=2048
SECONDS_PER_DAY=86400
CRYPTO="/usr/local/bin/crypto-tool"

# clean temp files
function fn_clean_temp() {
    if [ -f "${TEMP_ENV_CA_CERT}" ]; then rm ${TEMP_ENV_CA_CERT}; fi
    if [ -f "${TEMP_ENV_CA_KEY}" ]; then rm ${TEMP_ENV_CA_KEY}; fi
    if [ -f "${TEMP_KEY}" ]; then rm ${TEMP_KEY}; fi
    if [ -f "${TEMP_CSR}" ]; then rm ${TEMP_CSR}; fi
    if [ -f "${TEMP_CA_KEY}" ]; then rm ${TEMP_CA_KEY}; fi
}

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -n CN       required, CN name
      -c CERT     required, cert file path
      -k KEY      required, key file path (will be encrypted)
      -e EXPIRE   optional, expire days, default ${DEFAULT_EXPIRE}
      -l KEYLEN   optional, key length, default ${DEFAULT_KEYLEN}
      -C CA_CERT  optional, ca cert
      -K CA_KEY   optional, ca key (encrypted)
    
    Enviorments:
      CERT_ENV_CA_CERT              optional, base64 encoded ca cert
      CERT_ENV_CA_KEY               optional, base64 encoded ca key
      CERT_ENV_CA_KEY_PASS          optional, base64 encoded ca key passphase (plain text)
      CERT_ENV_CA_KEY_CIPHER_PASS   optional, base64 encoded ca key passphase (cipher text)
    "
    exit ${ERROR_INVALID_PARAM}
}

function fn_validate_params() {
    while getopts n:c:k:e:l:C:K: flag
    do
        case "${flag}" in
            n) CN=${OPTARG};;
            c) CERT=${OPTARG};;
            k) KEY=${OPTARG};;
            e) EXPIRE=${OPTARG};;
            l) KEYLEN=${OPTARG};;
            C) CA_CERT=${OPTARG};;
            K) CA_KEY=${OPTARG};;
            ?) fn_print_help
        esac
    done

    if [ "${CN}" == "" ]; then fn_print_help; fi
    if [ "${CERT}" == "" ]; then fn_print_help; fi

    CERT_DIR=$(dirname "${CERT}")
    if [ ! -d ${CERT_DIR} ]; then 
        echo "directory ${CERT_DIR} does not exit"
        exit ${ERROR_INVALID_PARAM}
    fi

    if [ "${KEY}" == "" ]; then fn_print_help; fi
    
    KEY_DIR=$(dirname "${KEY}")
    if [ ! -d ${KEY_DIR} ]; then 
        echo "directory ${KEY_DIR} does not exit"
        exit ${ERROR_INVALID_PARAM}
    fi

    if [ "${EXPIRE}" == "" ]; then EXPIRE=${DEFAULT_EXPIRE}; fi
    if [ "${KEYLEN}" == "" ]; then KEYLEN=${DEFAULT_KEYLEN}; fi

    # ca cert param not exist, use env
    if [ "${CA_CERT}" == "" -a "${CERT_ENV_CA_CERT}" != "" ]; then
        # create temp ca cert file
        TEMP_ENV_CA_CERT=$(mktemp); printf "${CERT_ENV_CA_CERT}" | base64 -d > ${TEMP_ENV_CA_CERT}
        CA_CERT=${TEMP_ENV_CA_CERT}
    fi

    # validate ca cert (at least valid in one day)
    if [ "${CA_CERT}" != "" ]; then
        if ! openssl x509 -checkend ${SECONDS_PER_DAY} -in ${CA_CERT} > /dev/null 2>&1; then
            echo "invalid or expired ca cert"
            fn_clean_temp
            exit ${ERROR_INVALID_PARAM}
        fi
    fi

    # ca key param not exist, use env
    if [ "${CA_KEY}" == "" -a "${CERT_ENV_CA_KEY}" != "" ]; then
        # create temp ca key file (cipher)
        TEMP_ENV_CA_KEY=$(mktemp);

        # convert cipher pass to plain pass
        if [ "${CERT_ENV_CA_KEY_CIPHER_PASS}" != "" ]; then 
            CERT_ENV_CA_KEY_PASS=$(printf ${CERT_ENV_CA_KEY_CIPHER_PASS} | ${CRYPTO} --decrypt)
            export CERT_ENV_CA_KEY_PASS
        fi
        
        if [ "${CERT_ENV_CA_KEY_PASS}" == "" ]; then
            # no passphase, encrypt and save to temp ca key
            printf "${CERT_ENV_CA_KEY}" | ${CRYPTO} -encrypt -format base64 | base64 -d > ${TEMP_ENV_CA_KEY}
        else
            # with passphase, convert to plain text, then encrypt and save to temp ca key
            printf "${CERT_ENV_CA_KEY}" | base64 -d | openssl rsa -passin env:CERT_ENV_CA_KEY_PASS 2>/dev/null | ${CRYPTO} -encrypt -sourcefile /dev/stdin -targetfile ${TEMP_ENV_CA_KEY}
        fi
    fi

    # validate ca key
    if [ "${CA_KEY}" != "" ]; then
        # validate ca key
        if ! ${CRYPTO} -decrypt -sourcefile ${CA_KEY} | openssl req -new -key /dev/stdin -subj "/CN=${CN}" > /dev/null 2>&1; then
            echo "invalid ca key"
            fn_clean_temp
            exit ${ERROR_INVALID_PARAM}
        fi
    fi
}


# generate temp key file encrypted with passphase
function fn_gen_temp_key() {
    # random passphase
    export TEMP_PASS=$(openssl rand -base64 32)

    # create temp key file with passphase
    TEMP_KEY=$(mktemp); openssl genrsa -out ${TEMP_KEY} -passout env:TEMP_PASS ${KEYLEN} 2>/dev/null
}


# convert temp key file to crypto-tool encrypted key file
function fn_gen_cipher_key() {
    openssl rsa -in ${TEMP_KEY} -passout env:TEMP_PASS 2>/dev/null | ${CRYPTO} -encrypt -sourcefile /dev/stdin -targetfile ${KEY}
}


# generate temp csr file
function fn_gen_csr() {
    TEMP_CSR=$(mktemp); openssl req -new -key ${TEMP_KEY} -passin env:TEMP_PASS -out ${TEMP_CSR} -subj "/CN=${CN}" 2>/dev/null
}


# generate cert
function fn_gen_cert() {
    if [ "${CA_CERT}" == "" -o "${CA_KEY}" == "" ]; then
        # generate cert without ca
        openssl x509 -req -in ${TEMP_CSR} -out ${CERT} -signkey ${TEMP_KEY} -passin env:TEMP_PASS -CAcreateserial -days ${EXPIRE} 2>/dev/null
        return
    else
        # create temp ca key file encrypted with passphase
        TEMP_CA_KEY=$(mktemp); ${CRYPTO} -decrypt -sourcefile ${CA_KEY} | openssl rsa -passout env:TEMP_PASS -out ${TEMP_CA_KEY} 2>/dev/null
        # generate cert with ca
        openssl x509 -req -in ${TEMP_CSR} -out ${CERT} -signkey ${TEMP_KEY} -passin env:TEMP_PASS -CA ${CA_CERT} -CAkey ${TEMP_CA_KEY} -CAcreateserial -days ${EXPIRE} 2>/dev/null
    fi
}

function fn_exit_failed() {
    fn_clean_temp
    exit ${ERROR_FAILED_GEN}
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    # generate temp key file 
    if ! fn_gen_temp_key; then fn_exit_failed; fi

    # conver to key file
    if ! fn_gen_cipher_key; then fn_exit_failed; fi

    # generate csr file
    if ! fn_gen_csr; then fn_exit_failed; fi

    # generate cert file
    if ! fn_gen_cert; then fn_exit_failed; fi

    # clean temp files
    fn_clean_temp
}

fn_main $@