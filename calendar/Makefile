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

build: 
	@CGO_ENABLED=0 go build -v -o calendar_api ${SOURCE}
	@CGO_ENABLED=0 go build -v -o calendar_scheduler ./cmd/scheduler
	@CGO_ENABLED=0 go build -ldflags="-s -w" -v -o calendar_sender ./cmd/sender

lint:
	@goimports -w ${GO_SRC_DIRS}	
	@gofmt -s -w ${GO_SRC_DIRS}
	@golangci-lint run

run:
	@go run ${SOURCE} grpc_server --config=configs/prod/config.yaml

run-client:
	@go run ${SOURCE} grpc_client --server="0.0.0.0:50051"

run-web:
	@go run ${SOURCE} web_server --config=configs/prod/config.yaml

run-mq:
	@go run ./cmd/scheduler --config=configs/prod/config-scheduler.yaml

run-sender:
	@go run ./cmd/sender --config=configs/prod/config-sender.yaml

test:
	@go test -v $(GO_TEST_DIRS)

gen:
	protoc -I api/proto --go_out=plugins=grpc:pkg/calendar/api api/proto/calendar.proto

mod:
	go mod verify
	go mod tidy

.PHONY: docker-build
docker-build: 
	docker build -t puzanovma/calendar -f ./build/calendar/api_build.Dockerfile .

.PHONY: up
up:
	./build/docker-up.sh

.PHONY: down
down:
	./build/docker-down.sh

.PHONY: serv-up
# Запуск postgres (с тестовыми данными) и rabbitmq для тестирования
serv-up:
	docker-compose -f build/docker-compose-serv.yml up --build --detach

.PHONY: serv-up
serv-down:
	docker-compose  --file build/docker-compose-serv.yml down

release:
	rm -rf ${RELEASE_DIR}${APP}*
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o ${RELEASE_DIR}/${APP}.exe ${SOURCE}
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o ${RELEASE_DIR}/${APP} ${SOURCE}

.PHONY: build run release lint test gen mod