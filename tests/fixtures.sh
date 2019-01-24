#! /bin/bash

set -eu

DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

mkdir -p "${DIR}/fixtures"
pushd "${DIR}/fixtures"

# Create the CA certificate
openssl req -new -x509 -extensions v3_ca -keyout ca.key -out ca.crt -days 1825 -nodes -subj "/C=US/ST=Denial/L=Springfield/O=Sofa/CN=www.example.com"

# Create a key with no password & generate a certificate from it
openssl genrsa -out passwordless.key 2048
openssl req -new -key passwordless.key -out passwordless.csr -subj "/C=US/ST=Denial/L=Springfield/O=Sofa/CN=passwordless"
openssl x509 -req -days 1825 -in passwordless.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out passwordless.crt

# Create a key with a password & generate a certificate from it
openssl genrsa -aes256 -passout pass:Pa55word -out password.key 2048
openssl req -new -passin pass:Pa55word -key password.key -out password.csr -subj "/C=US/ST=Denial/L=Springfield/O=Sofa/CN=password"
openssl x509 -req -days 1825 -in password.csr -CA ca.crt -CAkey ca.key -set_serial 02 -out password.crt

popd
