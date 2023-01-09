#!/bin/bash

# build directory
BUILD_DIR=$(dirname $0)

# workspace directory
BASE_DIR=$(readlink -f ${BUILD_DIR}/..)

# release directory
RELEASE_DIR="${BASE_DIR}/release"

# agent name
AGENT="flexagent"

# crypto tool
CRYPTO="crypto-tool"

# read build profile
source ${BUILD_DIR}/profile

# build target
TARGET_DIR="${BASE_DIR}/build/target"
PKG_DIR="${TARGET_DIR}/${AGENT}-${ARCH}-${VERSION}"

# clean target
rm -rf ${PKG_DIR}

# create target directories
mkdir -p ${PKG_DIR}
mkdir -p ${PKG_DIR}/bin
mkdir -p ${PKG_DIR}/plugins

# copy config files and install scripts
cp -r ${RELEASE_DIR}/etc ${PKG_DIR}
cp -r ${RELEASE_DIR}/certs ${PKG_DIR}
cp -r ${RELEASE_DIR}/systemd ${PKG_DIR}
cp    ${RELEASE_DIR}/install.sh ${PKG_DIR}

# copy plugins
plugin_list=""
for plugin in ${PLUGINS}; do
  plugin_list="${plugin_list},\"${plugin}\""
  cp -r ${RELEASE_DIR}/plugins/${plugin} ${PKG_DIR}/plugins/
done

# set plugins in config file
sed -i "s/^plugins: .*$/plugins: [${plugin_list:1}]/" ${PKG_DIR}/etc/config.yaml

# set plugins in install script
sed -i "s/^PLUGINS=.*$/PLUGINS=\"${PLUGINS}\"/" ${PKG_DIR}/install.sh

# build binary
GOOS=linux GOARCH=${ARCH} CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o ${PKG_DIR}/bin/${AGENT} ${BASE_DIR}/cmd/agent-server/main.go
GOOS=linux GOARCH=${ARCH} CGO_ENABLED=0 go build -ldflags="-extldflags=-static" -o ${PKG_DIR}/bin/${CRYPTO} ${BASE_DIR}/cmd/crypto/main.go

# pack target directory
cd ${TARGET_DIR}
tar -zcf ${AGENT}-${ARCH}-${VERSION}.tar.gz ${AGENT}-${ARCH}-${VERSION}

echo "${TARGET_DIR}/${AGENT}-${ARCH}-${VERSION}.tar.gz"