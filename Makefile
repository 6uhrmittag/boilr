SHELL := /bin/bash
.PHONY: help
GIT_HASH := $(shell eval git rev-parse --short HEAD || echo 'snapshot')

.DEFAULT_GOAL := help

help: ## This help message
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | column -c2 -t -s :


test: ## run tests
	go test ./...

build: ## build linux_amd64
	sed -i 's/Version = "0.3.0"/Version = "${GIT_HASH}"/g' pkg/boilr/configuration.go
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build boilr.go

build-all: ## build binaries and tar
	rm -fr *.tgz
	sed -i 's/Version = "0.3.0"/Version = "${GIT_HASH}"/g' pkg/boilr/configuration.go
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build boilr.go  && tar czvf boilr-${GIT_HASH}-linux_amd64.tgz boilr
	GOARCH=386   GOOS=linux CGO_ENABLED=0 go build boilr.go  && tar czvf boilr-${GIT_HASH}-linux_386.tgz boilr
	# GOARCH=amd64 GOOS=darwin CGO_ENABLED=0 go build boilr.go && tar czvf boilr-${GIT_HASH}-darwin_amd64.tgz boilr
	# GOARCH=386   GOOS=darwin CGO_ENABLED=0 go build boilr.go && tar czvf boilr-${GIT_HASH}-darwin_386.tgz boilr

gofmt: ## format code
	go fmt -x ./...

govet: ## reports suspicious constructs
	go vet ./...

clean: ## delete builds
	rm -fr *.tgz
