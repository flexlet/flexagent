#!/bin/bash
BASE_DIR=$(readlink -f $(dirname $0))

AGENT="flexagent"
CRYPTO="crypto-tool"
PLUGINS="linux"


BIN="/usr/local/bin"
OPT="/opt/${AGENT}"
ETC="/etc/${AGENT}"
SVC="/usr/lib/systemd/system/${AGENT}.service"

function fn_clean_service() {
  if [ -f ${SVC} ]; then
    STARTED=$(ps -ef | grep ${AGENT} | grep -v grep | wc -l)
    if (( STARTED )); then systemctl stop ${AGENT}; fi
    systemctl disable ${AGENT}
    rm ${SVC}
    systemctl daemon-reload
  fi
}

function fn_clean() {
  echo "clean ..."
  fn_clean_service
  rm -rf ${BIN}/${AGENT} ${BIN}/${CRYPTO}
  rm -rf ${OPT} ${ETC}
}

function fn_install() {
  echo "install ..."
  mkdir -p ${OPT} ${ETC}
  cp -r ${BASE_DIR}/bin/* ${BIN}
  cp -r ${BASE_DIR}/etc/* ${ETC}
  cp -r ${BASE_DIR}/certs ${OPT}
  cp -r ${BASE_DIR}/plugins ${OPT}
  cp ${BASE_DIR}/systemd/${AGENT}.service ${SVC}

  chmod 500 ${BIN}/${AGENT} ${BIN}/${CRYPTO}
  chmod 600 ${ETC}/*
  chmod 400 ${OPT}/certs/*

  for plugin in ${PLUGINS}; do
    chmod 500 ${OPT}/plugins/${plugin}/probes/*.sh
    chmod 500 ${OPT}/plugins/${plugin}/operations/*.sh
    chmod 400 ${OPT}/plugins/${plugin}/plugin.yaml
  done

}

function fn_init_keystore() {
  echo "check and init keystore ..."
  echo | ${BIN}/${CRYPTO} -encrypt > /dev/null 2>&1
  if [ "$?" != "0" ]; then
    ${BIN}/${CRYPTO} -init
  fi
  if [ "$?" != "0" ]; then
    echo "WARN: keystore init failed"
    exit -1
  fi
}

function fn_encrypt_key(){
  echo "encrypt server key ..."
  ${BIN}/${CRYPTO} -encrypt -sourcefile "${OPT}/certs/server.key" -targetfile "${OPT}/certs/server.key"
}

function fn_start() {
  echo "enalbe and start service"
  systemctl daemon-reload
  systemctl enable ${AGENT}
  systemctl start ${AGENT}; sleep 0.1
  systemctl status ${AGENT}
}

function fn_main() {
  fn_clean
  fn_install
  fn_init_keystore
  fn_encrypt_key
  fn_start
}

fn_main
