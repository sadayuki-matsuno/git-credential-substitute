PROJECT_ROOT=$(shell git rev-parse --show-toplevel)

VERSION := $(shell git describe --tags)
LDFLAGS=-ldflags "-s -w -X=main.version=$(VERSION)"

install:
	@ go install $(LDFLAGS)

build:
	@ go build $(LDFLAGS)

test:
	@ go test $(PROJECT_ROOT)/*.go
	@ go test $(PROJECT_ROOT)/internal/*.go
