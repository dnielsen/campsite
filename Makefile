#!/usr/bin/env bash

GO ?= go1.15.3
MAKE ?= make
DOCKER ?= docker

SERVICES = services

SERVICE_LANG ?= go1.15.3
SERVICE_PATH = services

# Available services
ALL_SERVICES = $(shell cd $(SERVICES) && git ls-tree -d --name-only HEAD)

ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

help: ;
	$(info Help for Campsite Services `make`)
	@echo
	@echo Commands:
	@echo vendor
	@echo \ \ Tidy and vendorize Go packages
	@echo local
	@echo \ \ Build a service\'s docker image and run it locally

.PHONY: vendor
vendor:
	-$(GO) mod tidy
	$(GO) mod vendor
	@-rm .vendor

.vendor:
	$(DOCKER) build --tag vendor -f Dockerfile.vendor .
	touch .vendor

ifdef SERVICE_NAME

.PHONY: build
build: .vendor
	@echo "Building service $(SERVICE_NAME).."
ifeq ($(SERVICE_TYPE), jobs)
	$(DOCKER) build --tag $(REPO_URL) -f ./$(SERVICE_PATH)/../Dockerfile.$(SERVICE_LANG).$(SERVICE_TYPE)  --build-arg RPAAS_SERVICE=$(SERVICE_NAME) ./$(SERVICE_PATH)/$(SERVICE_NAME)/.
else
	$(DOCKER) build --tag $(REPO_URL) ./$(SERVICE_PATH)/$(SERVICE_NAME)/. -f ./$(SERVICE_PATH)/../../Dockerfile.sandbox --build-arg RPAAS_SERVICE=$(SERVICE_NAME)
endif

endif

.PHONY: local
local:
	@echo "Doing a local build/run of service $(SERVICE_NAME)..."
	cd $(SERVICE_PATH)/$(SERVICE_NAME) && $(GO) build . && ./$(SERVICE_NAME)

.PHONY: $(ALL_SERVICES)
$(ALL_SERVICES):
	@$(MAKE) --no-print-directory -C $(SERVICE_PATH)/$@ $(call ARGS,deploy)

% ::
	@: