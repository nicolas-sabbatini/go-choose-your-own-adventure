PATH_TO_BINARY=./cmd/
.DEFAULT_GOAL := run

run:
	go run ${PATH_TO_BINARY}$(BIN)/$(BIN).go

build-linux:
	GOARCH=amd64 GOOS=linux go build -o target/$(BIN)-linux-amd64 ${PATH_TO_BINARY}$(BIN)/$(BIN).go

build-windows:
	GOARCH=amd64 GOOS=windows go build -o target/$(BIN)-windows-amd64.exe ${PATH_TO_BINARY}$(BIN)/$(BIN).go

build-mac:
	GOARCH=amd64 GOOS=darwin go build -o target/$(BIN)-darwin-amd64 ${PATH_TO_BINARY}$(BIN)/$(BIN).go

build-all: build-linux build-windows build-mac

clear:
	rm -rf target
	go clean

tidy:
	go mod tidy

air:
	go build -o ./tmp/main ${PATH_TO_BINARY}$(BIN)/$(BIN).go
