.DEFAULT_GOAL = build

APP=gocopy
SOURSE=. 
RELEASE_DIR=bin
pkgs=$(shell go list ./... | fgrep -v /vendor)

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)	

build:
	go build -o ${APP} ${SOURSE}

run:
	go run $(SOURSE) -from=files/file1 -to=files/file2
	go run $(SOURSE) -from=files/file1 -to=files/file_limit -limit=5
	go run $(SOURSE) -from=files/file1 -to=files/file_offset -offset=4	
	go run $(SOURSE) -from=files/file1 -to=files/file_offset_limit -offset=5 -limit=10
	@#go run $(SOURSE) -from=file_not_exists -to=out_file_err
	@#go run $(SOURSE) -from=/dev/urandom -to=out_file_err

test: lint
	go test -v $(GO_TEST_DIRS)

lint:
	golangci-lint run

release:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ${RELEASE_DIR}/win/${APP}.exe main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${RELEASE_DIR}/linux/${APP} main.go

coverage:
	go test -coverprofile=coverage.out.tmp $(GO_TEST_DIRS)
	go tool cover -func=coverage.out.tmp
	go tool cover -html=coverage.out.tmp

.PHONY: build run test coverage release lint