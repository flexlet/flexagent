#!/bin/bash

# build directory
BUILD_DIR=$(dirname $0)

# workspace directory
BASE_DIR=$(readlink -f ${BUILD_DIR}/..)

# release directory
RELEASE_DIR="${BASE_DIR}/release"

# read build profile
source ${BUILD_DIR}/profile

# build target
TARGET_DIR="${BASE_DIR}/build/target"
CERTS_DIR="${TARGET_DIR}/certs"

# clean certs
rm -rf ${CERTS_DIR}

# create certs directories
mkdir -p ${CERTS_DIR}

cd ${CERTS_DIR}

# create root cert key
openssl genrsa -out ca.key 2048

# create root cert request
openssl req -new -out ca.csr -key ca.key -subj "${ISSUER}"

# create root cert
openssl x509 -req -in ca.csr -out ca.crt -signkey ca.key -CAcreateserial -days 3650

# create server cert key
openssl genrsa -out server.key 2048

# create server cert request
openssl req -new -out server.csr -key server.key -subj "${ISSUER}"

# create server cert
openssl x509 -req -in server.csr -out server.crt -signkey server.key -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650

# create client cert key
openssl genrsa -out client.key 2048

# create client cert request
openssl req -new -out client.csr -key client.key -subj "${ISSUER}"

# create client cert
openssl x509 -req -in client.csr -out client.crt -signkey client.key -CA ca.crt -CAkey ca.key -CAcreateserial -days 3650

# copy to release dir
cp ca.crt server.crt server.key ${RELEASE_DIR}/certs/

cd - 