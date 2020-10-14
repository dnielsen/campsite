#!/usr/bin/env bash

SERVICES = packages

ALL_SERVICES = $(shell cd $(SERVICES) && git ls-tree -d --name-only HEAD)

ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

.PHONY: local
local:
	cd packages/$(SERVICE_NAME) && go run cmd/main.go


.PHONY: $(ALL_SERVICES)
$(ALL_SERVICES):
        @$(MAKE) --no-print-directory -C $(SERVICE_PATH)/$@ $(call ARGS,deploy)
% ::
        @:
