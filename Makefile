SHELL := /bin/bash
BASEDIR = $(shell pwd)
export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

fmt:
	gofmt -w .
mod:
	go mod tidy
utest:
	go test -coverpkg=./... -coverprofile=coverage.data ./...
help:
	@echo "make fmt - format code"
	@echo "make mod - mod tidy"
	@echo "make utest - unit test"