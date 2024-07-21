.PHONY: all build build/all build/datagenerator install fmt fmt/check lint lint/govet test

# Install binaries into ./bin rather than $GOPATH/bin to avoid conflicts
GOBIN=$(shell pwd)/bin

GOFMT_DIRS=cmd/ internal/

export PATH := $(GOBIN):$(PATH)

all: install

build: install build/all

build/all: install build/datagenerator

build/datagenerator: install
	go build -o csv-data-generator cmd/csv-data-generator/main.go

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
	go test -v -mod=vendor ./...
