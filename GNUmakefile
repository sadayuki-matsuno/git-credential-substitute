PROJECT_ROOT=$(shell git rev-parse --show-toplevel)
install:
	@ go install $(PROJECT_ROOT)/cmd/*.go

build:
	@ go build -o $(PROJECT_ROOT)/bin/git-credential-substitute ./cmd/main.go 

test:
	@ go test $(PROJECT_ROOT)/cmd/*.go
	@ go test $(PROJECT_ROOT)/internal/*.go
