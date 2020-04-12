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
	@echo "Компилилируем сервисы"
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

gen:
	protoc -I api/proto --go_out=plugins=grpc:pkg/calendar/api api/proto/calendar.proto

mod:
	go mod verify
	go mod tidy

.PHONY: docker-build
docker-build: 
	docker build -t puzanovma/calendar -f ./build/docker/calendar/api_build.Dockerfile .

up: 
	./build/docker-up.sh

down:
	./build/docker-down.sh

.PHONY: test
# Поднимаем postgres, rabbitmq, тестируем, убираем postgres
test:
	docker-compose -f build/deploy/docker-compose-db.yml up --build --detach ;\
	go test -v -count=1 ./internal/... ;\
	docker-compose  --file build/deploy/docker-compose-db.yml down ;\

.PHONY: integration-tests
# Запуск Integration tests
integration-tests:
	set -e ;\
	docker-compose -f build/deploy/docker-compose.test.yml up --build -d ;\
	test_status_code=0 ;\
	docker-compose -f build/deploy/docker-compose.test.yml run integration_tests go test ./tests/integration || test_status_code=$$? ;\
	docker-compose -f build/deploy/docker-compose.test.yml down --volumes;\
	printf "Return code is $$test_status_code\n" ;\
	exit $$test_status_code ;\

release:
	rm -rf ${RELEASE_DIR}${APP}*
	GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o ${RELEASE_DIR}/${APP}.exe ${SOURCE}
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o ${RELEASE_DIR}/${APP} ${SOURCE}

.PHONY: build run release lint test gen mod up down