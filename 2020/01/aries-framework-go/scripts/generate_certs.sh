#!/bin/sh
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e


echo "Generating TrustBloc local PKI"

cd /opt/workspace/trustbloc

trustblocSSLConf=$(mktemp)
echo "subjectKeyIdentifier=hash
authorityKeyIdentifier = keyid,issuer
extendedKeyUsage = serverAuth
keyUsage = Digital Signature, Key Encipherment
subjectAltName = @alt_names
[alt_names]
DNS.1 = *.trustbloc.local" >> "$trustblocSSLConf"

CERT_CA="fixtures/certs/trustbloc-dev-ca.crt"
if [ ! -f "$CERT_CA" ]; then
    echo "Generating CA cert"
    openssl ecparam -name prime256v1 -genkey -noout -out fixtures/certs/trustbloc-dev-ca.key
    openssl req -new -x509 -key fixtures/certs/trustbloc-dev-ca.key -subj "/C=CA/ST=ON/O=TrustBloc/OU=TrustBloc Dev CA" -out fixtures/certs/trustbloc-dev-ca.crt -days 1095
else
    echo "Skipping CA generation - already exists"
fi

# Create TLS certs
openssl ecparam -name prime256v1 -genkey -noout -out fixtures/certs/trustbloc.local.key
openssl req -new -key fixtures/certs/trustbloc.local.key -subj "/C=CA/ST=ON/O=TrustBloc/OU=trustbloc-edge-sandbox/CN=trustbloc.local" -out fixtures/certs/trustbloc.local.csr
openssl x509 -req -in fixtures/certs/trustbloc.local.csr -CA fixtures/certs/trustbloc-dev-ca.crt -CAkey fixtures/certs/trustbloc-dev-ca.key -CAcreateserial -extfile "$trustblocSSLConf" -out fixtures/certs/trustbloc.local.crt -days 365

echo "done generating TrustBloc local PKI"
