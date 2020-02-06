#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#
set -e

echo "Running $0"

cd fixtures
dockerComposeFiles='-f docker-compose.yml'
(source .env && docker-compose $dockerComposeFiles down && docker-compose $dockerComposeFiles up --force-recreate)

