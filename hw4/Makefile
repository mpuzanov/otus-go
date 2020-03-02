build:
	go build .

run:
	go run .

test:
	go test -timeout 30s ./doublelinkedlist/

lint:
	@goimports -w .
	@gofmt -s -w .

.PHONY: build run test