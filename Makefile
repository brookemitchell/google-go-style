.PHONY: lint test build vet

default: lint test vet

test:
	go test -v -cover ./...

lint:
	golangci-lint run

build:
	go build ./cmd/google-go-style

vet: build
	go vet -vettool=./google-go-style ./...
