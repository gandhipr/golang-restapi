MODULE = $(shell go list -m)
PACKAGES := $(shell go list ./... | grep -v /vendor/)


.PHONY: default
default: help

.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:  ## build the API server binary
	go mod tidy
	go build -a -o apiserver $(MODULE)/

.PHONY: run
run:  ## run the API server binary
	go run apiserver

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)

.PHONY: test
test:  ## run tests
	go test $(MODULE)/./...

.PHONY: generate-mock
generate-mock:  ## run unit-tests
	go generate $(MODULE)/./...
