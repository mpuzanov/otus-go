
build:
	go build .

run:
	go run .
	
test:
	go test -timeout 30s -v .

check:
	go fmt .
	go vet .
	~/go/bin/golint -set_exit_status .

.PHONY: build test run check