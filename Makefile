# Disable cgo by default.
CGO_ENABLED ?= 0

GOBIN := $(if $(shell go env GOBIN),$(shell go env GOBIN),$(GOPATH)/bin)
PATH := $(GOBIN):$(PATH)

generate:
	buf generate
clean:
	rm -rf build

regenerate: clean generate

install:
	@go install github.com/bufbuild/buf/cmd/buf@latest