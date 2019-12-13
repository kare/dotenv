import_path := kkn.fi/dotenv
golint := $(GOPATH)/bin/golint

.SHELLFLAGS := -c # Run commands in a -c flag
.ONESHELL: ; # scripts execute in same shell
.SILENT: ; # no need for @
.NOTPARALLEL: ; # wait for this target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell
.DEFAULT_GOAL := dotenv

.PHONY: help
help: ## Show Help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.PHONY: test
test: ## Run unit tests
	go test -v ./...

.PHONY: benchmark
benchmark: ## Run benchmarks
	go test -run=XXX -bench=.

.PHONY: dotenv
dotenv: ## Build package (default target)
	go build $(import_path)

.PHONY: fmt
fmt: ## Run gofmt with simplify and write results to files
	gofmt -s -w .

.PHONY: clean
clean: ## Clean all generated files
	go clean -i -r

coverfile := $(shell mktemp)
.PHONY: cover
cover: ## Run coverage report
	go test -coverprofile=$$coverfile $(shell go list ./...)
	go tool cover -func=$$coverfile
	go tool cover -html=$$coverfile
	rm -f $$coverfile

$(golint): ## Install golint
	GO111MODULE=off go get -u golang.org/x/lint/golint

.PHONY: golint
golint: $(golint) ## Run golint
	golint
