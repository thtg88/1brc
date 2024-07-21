.PHONY: all build install fmt fmt/check lint lint/govet test

# Install binaries into ./bin rather than $GOPATH/bin to avoid conflicts
GOBIN=$(shell pwd)/bin

GOFMT_DIRS=cmd/ internal/

export PATH := $(GOBIN):$(PATH)

all: install

# TODO: build each cmd package into separate executables
build: install

install:
	go install -mod=vendor -v ./...

fmt:
	gofmt -w $(GOFMT_DIRS)

fmt/check:
	if [ "$$(gofmt -l ${GOFMT_DIRS} | wc -l)" -gt 0 ]; then exit 1; fi

lint: lint/govet

lint/govet:
	go vet ./...

test:
	go test -mod=vendor ./...
