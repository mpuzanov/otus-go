SOURCE=./cmd/calendar
APP=calendar
GOBASE=$(shell pwd)
RELEASE_DIR=$(GOBASE)/bin
VERSION=1.0

.DEFAULT_GOAL = build 

GO_SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)
GO_TEST_DIRS := $(shell \
	find . -name "*_test.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)	

build: lint
	@go build -v -o ${APP} ${SOURCE}
	@go build -v -o scheduler ./cmd/scheduler
	@go build -v -o sender ./cmd/sender

lint:
	@goimports -w ${GO_SRC_DIRS}	
	@gofmt -s -w ${GO_SRC_DIRS}
	@golangci-lint run

run:
	@go run ${SOURCE} grpc_server --config=configs/config-prod.yaml

run-client:
	@go run ${SOURCE} grpc_client --server=":50051"

run-web:
	@go run ${SOURCE} web_server --config=configs/config-prod.yaml

run-mq:
	@go run ./cmd/scheduler --config=configs/config-dev.yaml

run-sender:
	@go run ./cmd/sender --config=configs/config-dev.yaml

test:
	@go test -v $(GO_TEST_DIRS)

gen:
	protoc -I api/proto --go_out=plugins=grpc:pkg/calendar/api api/proto/calendar.proto

mod:
	go mod verify
	go mod tidy

.PHONY: dockerbuild
dockerbuild: 
	docker build -t puzanovma/calendar -f ./deployments/calendar/Dockerfile .

.PHONY: docker-compose-up
docker-compose-up: 
	docker-compose -f ./deployments/docker-compose.yaml up --build

.PHONY: docker-compose-down
docker-compose-down: 
	docker-compose -f ./deployments/docker-compose.yaml down -v

release:
	rm -rf ${RELEASE_DIR}${APP}*
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o ${RELEASE_DIR}/${APP}.exe ${SOURCE}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ${RELEASE_DIR}/${APP} ${SOURCE}

.PHONY: build run release lint test gen mod