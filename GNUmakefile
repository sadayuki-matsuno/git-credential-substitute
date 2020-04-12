PROJECT_ROOT=$(shell git rev-parse --show-toplevel)
install:
	@ go install $(PROJECT_ROOT)/*.go

build:
	@ go build

test:
	@ go test $(PROJECT_ROOT)/*.go
	@ go test $(PROJECT_ROOT)/internal/*.go
