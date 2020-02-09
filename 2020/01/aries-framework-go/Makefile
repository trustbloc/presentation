# Copyright SecureKey Technologies Inc.
#
# SPDX-License-Identifier: Apache-2.0

# Namespace for the agent images
DOCKER_OUTPUT_NS      ?= trustbloc
AGENT_REST_IMAGE_NAME ?= agent-rest
WEBHOOK_IMAGE_NAME    ?= sample-webhook
BOT_IMAGE_NAME        ?= sample-bot

# Tool commands (overridable)
ALPINE_VER ?= 3.11
GO_TAGS    ?=
GO_VER     ?= 1.13.7

.PHONY: docker
docker: sample-webhook-docker sample-bot-docker

.PHONY: sample-webhook
sample-webhook:
	@echo "Building sample webhook server"
	@mkdir -p ./build/bin
	@cd cmd/webhook && go build -o ../../build/bin/webhook-server

.PHONY: sample-webhook-docker
sample-webhook-docker:
	@echo "Building sample webhook server docker image"
	@docker build -f ./images/mocks/webhook/Dockerfile --no-cache -t $(DOCKER_OUTPUT_NS)/$(WEBHOOK_IMAGE_NAME):latest \
	--build-arg GO_VER=$(GO_VER) \
	--build-arg ALPINE_VER=$(ALPINE_VER) \
	--build-arg GO_TAGS=$(GO_TAGS) \
	--build-arg GOPROXY=$(GOPROXY) .

.PHONY: sample-bot
sample-bot:
	@echo "Building sample bot server"
	@mkdir -p ./build/bin
	@cd cmd/bot && go build -o ../../build/bin/bot-server

.PHONY: sample-bot-docker
sample-bot-docker:
	@echo "Building sample bot server docker image"
	@docker build -f ./images/mocks/bot/Dockerfile --no-cache -t $(DOCKER_OUTPUT_NS)/$(BOT_IMAGE_NAME):latest \
	--build-arg GO_VER=$(GO_VER) \
	--build-arg ALPINE_VER=$(ALPINE_VER) \
	--build-arg GO_TAGS=$(GO_TAGS) \
	--build-arg GOPROXY=$(GOPROXY) .
