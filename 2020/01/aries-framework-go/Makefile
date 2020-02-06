# Copyright SecureKey Technologies Inc.
#
# SPDX-License-Identifier: Apache-2.0

.PHONY: clean
clean: clean-build

.PHONY: clean-build
clean-build:
	@rm -Rf ./build
	@rm -Rf ./fixtures/certs

.PHONY: generate-test-keys
generate-test-keys: clean
	@scripts/generate_fixtures.sh

.PHONY: demo-start
demo-start: clean generate-test-keys
	@scripts/docker-start.sh