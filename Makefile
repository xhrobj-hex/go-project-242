lint:
	golangci-lint run ./...

run:
	go run cmd/hexlet-path-size/main.go

test:
	go test -v ./...

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size