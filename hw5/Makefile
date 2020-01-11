GOBASE=$(shell pwd)
GO_TEST_DIRS = $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
PACKAGES := $(shell go list ./... | grep -v /vendor/)

build:
	go build .

run:
	go run . 

test:
	go vet $(PACKAGES)
	go fmt $(PACKAGES) 
	go test -v -count=1 $(GO_TEST_DIRS)

coverage:
	go test -coverprofile=coverage.out.tmp $(GO_TEST_DIRS)
	go tool cover -func=coverage.out.tmp
	go tool cover -html=coverage.out.tmp

.PHONY: build run test coverage