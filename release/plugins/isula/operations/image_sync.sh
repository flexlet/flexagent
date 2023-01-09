#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_SEARCH_IMAGES=1
ERROR_SYNC_IMAGES=2

# consts
CRYPTO="/usr/local/bin/crypto-tool"

# plugin base directory
BASE_DIR=$(readlink -f $(dirname $0)/..)
CATALOG_SEARCH="python ${BASE_DIR}/pyclient/registry_v2/catalog_search.py"

function fn_print_help() {
    echo "Help: $(basename $0) -r REGISTRY -k KEYWORDS [-t TAGS] [-u USER] [-i]
    Options:
      -r REGISTRY   required, Registry (<host>:<port>)
      -u USER       optional, Registry username (<host>:<port>)
      -k KEYWORDS   required, Repos keywords (comma seperate)
      -t TAGS       optional, Match tags (comma seperate)
      -i INSECURE   optional, http or not, default: no

    Envirments:
      IMAGE_REPO_PASS          optional, user password (plain text)
      IMAGE_REPO_CIPHER_PASS   optional, user password (cipher text)
    "

    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts r:u:k:t:i flag
    do
        case "${flag}" in
            r) REGISTRY=${OPTARG};;
            u) USER=${OPTARG};;
            k) KEYWORDS=${OPTARG};;
            t) TAGS=${OPTARG};;
            i) INSECURE=1;;
            ?) fn_print_help
        esac
    done

    if [ "${REGISTRY}" == "" ]; then fn_print_help; fi
    if [ "${KEYWORDS}" == "" ]; then fn_print_help; fi

    # convert cipher pass to plain pass
    if [ "${IMAGE_REPO_CIPHER_PASS}" != "" ]; then 
      IMAGE_REPO_PASS=$(printf ${IMAGE_REPO_CIPHER_PASS} | ${CRYPTO} --decrypt)
    fi

    if [ "${USER}" != "" -a "${IMAGE_REPO_PASS}" == "" ]; then fn_print_help; fi

    CATALOG_SEARCH="${CATALOG_SEARCH} --registry ${REGISTRY} --keywords ${KEYWORDS}"
    if [ "${USER}" != "" ]; then 
      export REGISTRY_PYCLIENT_USERNAME=${USER}
      export REGISTRY_PYCLIENT_PASSWORD=${IMAGE_REPO_PASS}
    fi

    if [ "${INSECURE}" == "1" ]; then CATALOG_SEARCH="${CATALOG_SEARCH} --insecure"; fi
    if [ "${TAGS}" != "" ]; then CATALOG_SEARCH="${CATALOG_SEARCH} --tags ${TAGS}"; fi

}

function fn_main() {
    # validate parameters
    fn_validate_params $@

  if ! MATCHED=$(${CATALOG_SEARCH}); then exit ${ERROR_SEARCH_IMAGES}; fi

  for image in ${MATCHED}; do
    if ! isula pull ${image}; then exit ${ERROR_PULL_IMAGES}; fi
  done
}

fn_main $@