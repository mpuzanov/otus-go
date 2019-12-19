build:
	go build .

run:
	go run . a4bc2d5e

test:
	go test -timeout 30s ./unpack

bench:
	go test ./unpack -bench=. -benchmem

.PHONY: build run test bench