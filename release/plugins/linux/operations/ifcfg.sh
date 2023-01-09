#!/bin/bash

# error codes
ERROR_INVALID_PARAM=-1
ERROR_FAILED_IFUP=1

function fn_print_help() {
    echo "Help: $(basename $0) [options]
    Options:
      -d DEVICE   optional, network device, must specifi DEVICE or MAC
      -M MAC      optional, mac address, must specifi DEVICE or MAC
      -a IPADDR   required, ipv4 address
      -m NETMASK  required, network mask (ipv4)
      -g GATEWAY  optional, gateway (ipv4)
      -r ROUTES   optional, routes (multiple), format: '<network>/<prefix>:<router>'
    "
    exit ${ERROR_INVALID_PARAM} 
}

function fn_validate_params() {
    while getopts d:M:a:m:g:r: flag
    do
        case "${flag}" in
            d) DEVICE=${OPTARG};;
            M) MAC=${OPTARG};;
            a) IPADDR=${OPTARG};;
            m) NETMASK=${OPTARG};;
            g) GATEWAY=${OPTARG};;
            r) ROUTES="${ROUTES}${OPTARG}\n";;
            ?) fn_print_help
        esac
    done

    if [ "${MAC}" != "" -a "${DEVICE}" == "" ]; then
      for f in /sys/class/net/*/address; do
        if [ "$(cat $f)" == "${MAC}" ]; then 
          DEVICE=$(basename $(dirname $f))
          break
        fi
      done
    fi

    if [ "${DEVICE}" == "" ]; then fn_print_help; fi
    if [ "${IPADDR}" == "" ]; then fn_print_help; fi
    if [ "${NETMASK}" == "" ]; then fn_print_help; fi
    
    # check device exist or not
    /usr/sbin/ip a s ${DEVICE} > /dev/null 2>&1
    if (( $? == 1 )); then
      printf "DEVICE %s not exist\n" ${DEVICE}
      exit ${ERROR_INVALID_PARAM} 
    fi

    local IPV4_PATTERN='^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$'
    
    # check ipaddr pattern
    local MATCH=$(echo ${IPADDR} | awk "{if(/${IPV4_PATTERN}/){print 1}else{print 0}}")    
    if (( ${MATCH} == 0 )); then
      printf "IPADDR %s format error\n" ${IPADDR}
      exit ${ERROR_INVALID_PARAM} 
    fi

    # check netmask pattern
    MATCH=$(echo ${NETMASK} | awk "{if(/${IPV4_PATTERN}/){print 1}else{print 0}}")    
    if (( ${MATCH} == 0 )); then
      printf "NETMASK %s format error\n" ${NETMASK}
      exit ${ERROR_INVALID_PARAM} 
    fi

    if [ "${GATEWAY}" != "" ]; then
        # check gateway pattern
        MATCH=$(echo ${GATEWAY} | awk "{if(/${IPV4_PATTERN}/){print 1}else{print 0}}")    
        if (( ${MATCH} == 0 )); then
          printf "GATEWAY %s format error\n" ${GATEWAY}
          exit ${ERROR_INVALID_PARAM} 
        fi
    fi
}

function fn_main() {
    # validate parameters
    fn_validate_params $@

    # write to network config
    IFCFG="/etc/sysconfig/network-scripts/ifcfg-${DEVICE}"
    IFROUTE="/etc/sysconfig/network-scripts/route-${DEVICE}"

    cat <<____EOF > ${IFCFG}
TYPE=Ethernet
BOOTPROTO=static
ONBOOT=yes
IPV6INIT=no
NAME=${DEVICE}
DEVICE=${DEVICE}
IPADDR=${IPADDR}
NETMASK=${NETMASK}
____EOF

    if [ "${GATEWAY}" != "" ]; then
        echo "GATEWAY=${GATEWAY}" >> ${IFCFG}
    fi

    if [ "${ROUTES}" != "" ]; then
        printf "${ROUTES}" | sed 's/:/ via /g' > ${IFROUTE}
    fi

    # bring up interface
    /usr/sbin/ifdown ${DEVICE} && /usr/sbin/ifup ${DEVICE}

    # check ip is up or not
    IFUP=$(ip a s ${DEVICE} | grep "${IPADDR}" | wc -l)

    if (( ${IFUP} == 0 )); then
        printf "IPADDR %s can not up\n" ${IPADDR}
        exit ${ERROR_FAILED_IFUP}
    fi

    # show ip address
    /usr/sbin/ip a s ${DEVICE}

    exit 0
}

fn_main $@