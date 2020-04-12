PROJECT_ROOT=$(shell git rev-parse --show-toplevel)
install:
	@ go install $(PROJECT_ROOT)/*.go

build:
	@ go build -o $(PROJECT_ROOT)/bin/git-credential-substitute ./main.go 

test:
	@ go test $(PROJECT_ROOT)/*.go
	@ go test $(PROJECT_ROOT)/internal/*.go
