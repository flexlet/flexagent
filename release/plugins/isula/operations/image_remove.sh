#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_REMOVE_IMAGE=1

function fn_print_help() {
    echo "Help: $(basename $0) <image> <image> ..."
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    if [ "$#" == "0" ]; then fn_print_help; fi
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    STATUS=0
    # pull images
    for IMAGE in $@; do
        if ! isula rmi ${IMAGE}; then STATUS=${ERROR_REMOVE_IMAGE}; fi
    done
    
    exit ${STATUS}
}

fn_main $@