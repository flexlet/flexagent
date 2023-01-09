#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_LIST_IMAGES=1
ERROR_REMOVE_IMAGES=2

# consts
DEFAULT_DAYS_NOT_USED=30

# plugin base directory
BASE_DIR=$(readlink -f $(dirname $0)/..)
DEFAULT_EXCEPT_FILE="${BASE_DIR}/config/image_whitelist.txt"

function fn_print_help() {
    echo "Help: $(basename $0) -d DAYS_NOT_USED -E EXCEPT_FILE -e <EXCEPT_IMAGE> -e <EXCEPT_IMAGE> ...
    Options:
      -d DAYS_NOT_USED     optional, days not used, default: ${DEFAULT_DAYS_NOT_USED}
      -E EXCEPT_FILE       optional, file with except image names
      -e EXCEPT_IMAGE      optional, except image
    "

    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts d:E:e: flag
    do
        case "${flag}" in
            d) DAYS_NOT_USED=${OPTARG};;
            E) EXCEPT_FILE=${OPTARG};;
            e) EXCEPT_IMAGES="${EXCEPT_IMAGES}\"${OPTARG}\",";;
            ?) fn_print_help
        esac
    done

    if [ "${DAYS_NOT_USED}" == "" ]; then DAYS_NOT_USED=${DEFAULT_DAYS_NOT_USED}; fi
    
    if [ "${EXCEPT_FILE}" == "" ]; then EXCEPT_FILE="${DEFAULT_EXCEPT_FILE}"; fi
    for EXCEPT in $(cat ${EXCEPT_FILE}); do EXCEPT_IMAGES="${EXCEPT_IMAGES}\"${EXCEPT:0:-1}\","; done
    
    if [ "${EXCEPT_IMAGES}" != "" ]; then EXCEPT_IMAGES="[${EXCEPT_IMAGES:0:-1}]"; fi
}

function fn_image_not_used() {
    IMAGES_ALL=$(isula images | grep -v '^REPOSITORY' | awk '//{printf "{\"name\":\"%s:%s\",\"id\":\"%s\"},",$1,$2,$3}')
    IMAGES_ALL="[${IMAGES_ALL:0:-1}]"
    HISTORY=$(isula ps -a --no-trunc --format '{"id":"{{.ID}}","image":"{{.Image}}","created":"{{.Created}}","status":"{{.Status}}"},')
    HISTORY="[${HISTORY:0:-1}]"
    python -c "
import json

IMAGES_ALL = json.loads('${IMAGES_ALL}')
EXCEPT_IMAGES = json.loads('${EXCEPT_IMAGES}')
HISTORY = json.loads('$(echo ${HISTORY})')

IMAGES_IN_USE = []
for ITEM in HISTORY:
    if ITEM['image'] in IMAGES_IN_USE:
        continue
    if ITEM['status'].startswith('Up'):
        IMAGES_IN_USE.append(ITEM['image'])
        continue
    if 'days' not in ITEM['created']:
        IMAGES_IN_USE.append(ITEM['image'])
        continue
    CREATED_DAYS = int(ITEM['created'].split(' ')[0])
    if CREATED_DAYS <= ${DAYS_NOT_USED}:
        IMAGES_IN_USE.append(ITEM['image'])

IMAGES_NOT_EXCEPTED = []
for IMAGE in IMAGES_ALL:
    IS_EXCEPTED = False
    for IMAGE_EXCEPT in EXCEPT_IMAGES:
        if IMAGE_EXCEPT in IMAGE['name']:
            IS_EXCEPTED = True
            break
    if not IS_EXCEPTED:
        IMAGES_NOT_EXCEPTED.append(IMAGE)

for IMAGE in IMAGES_NOT_EXCEPTED:
    IS_IN_USE = False
    for IMAGE_IN_USE in IMAGES_IN_USE:
        if IMAGE['name'] == IMAGE_IN_USE or IMAGE['id'] in IMAGE_IN_USE:
            IS_IN_USE = True
            break
    if not IS_IN_USE:
        print(IMAGE['id'], end=' ')
"
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    if ! IMAGES_NOT_USED=$(fn_image_not_used); then exit ${ERROR_LIST_IMAGES}; fi

    if [ "${IMAGES_NOT_USED}" == "" ]; then exit 0; fi

    if ! isula rmi ${IMAGES_NOT_USED}; then exit ${ERROR_REMOVE_IMAGES}; fi
}

fn_main $@