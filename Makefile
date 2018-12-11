IMPORT_PATH := kkn.fi/dotenv

SRC := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
GOMETALINTER := $(GOPATH)/bin/gometalinter

.SHELLFLAGS := -c # Run commands in a -c flag
.ONESHELL: ; # scripts execute in same shell
.SILENT: ; # no need for @
.NOTPARALLEL: ; # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.DEFAULT_GOAL := dotenv

.PHONY: help
help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.PHONY: unit
unit: ## Run unit tests
	go test -v ./...

.PHONY: dotenv
dotenv: ## Build package (default target)
	go build $(IMPORT_PATH)

.PHONY: metalint
metalint: $(GOMETALINTER) ## Run gometalinter
	gometalinter --vendor --exclude=vendor ./...

$(GOMETALINTER): ## Install gometalinter and dependencies
	GO111MODULE=off go get -u github.com/alecthomas/gometalinter
	gometalinter --install

.PHONY: fmt
fmt: ## Run gofmt with simplify and write results to files
	gofmt -s -w $(SRC)

.PHONY: clean
clean: ## Clean all generated files
	go clean -i -r $(IMPORT_PATH)

COVERFILE := $(shell mktemp)
.PHONY: cover
cover: ## Run coverage report
	go test -coverprofile=$$COVERFILE $(shell go list ./...)
	go tool cover -func=$$COVERFILE
	go tool cover -html=$$COVERFILE
	rm -f $$COVERFILE

.PHONY: dev-tools
dev-tools: ## Install Go developer tools
	GO111MODULE=off go get -u golang.org/x/lint/golint
