#!/bin/sh
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e

mkdir -p fixtures/certs
cp ~/.trustbloc-local/sandbox/certs/trustbloc-dev-ca.* fixtures/certs 2>/dev/null || :
docker run -i --rm \
    -v $(pwd):/opt/workspace/trustbloc \
    --entrypoint "/opt/workspace/trustbloc/scripts/generate_certs.sh" \
    frapsoft/openssl
