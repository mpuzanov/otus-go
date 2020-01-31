.DEFAULT_GOAL = build

APP=go-envdir
SOURSE=. 
RELEASE_DIR=bin

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
	
run: build
	@#./go-envdir myEnvExample ls -l
	@#./go-envdir myEnvExample env
	./go-envdir myEnvExample printenv EP1 

test: lint
	go test -v $(GO_TEST_DIRS)

lint:
	goimports -w ${GO_SRC_DIRS}
	golint ${GO_SRC_DIRS}
	@#golangci-lint run

release:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ${RELEASE_DIR}/win/${APP}.exe main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ${RELEASE_DIR}/linux/${APP} main.go

coverage:
	go test -coverprofile=coverage.out.tmp $(GO_TEST_DIRS)
	go tool cover -func=coverage.out.tmp
	go tool cover -html=coverage.out.tmp

.PHONY: build run test coverage release lint