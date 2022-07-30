BINARY_NAME=similar
GOARCH=$(shell go env GOARCH)
GOOS=$(shell go env GOOS)

all: build install

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o bin/${BINARY_NAME}-$(GOOS)-${GOARCH} ./cmd/...

install: build
	cp bin/${BINARY_NAME}-$(GOOS)-${GOARCH} /usr/local/bin/similar || \
	mkdir -p ~/bin && cp bin/${BINARY_NAME}-$(GOOS)-${GOARCH} ~/bin/similar


build_all:
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME}-linux-amd64 ./cmd/...
	GOOS=linux GOARCH=arm64 go build -o bin/${BINARY_NAME}-linux-arm64 ./cmd/...
	GOOS=linux GOARCH=arm go build -o bin/${BINARY_NAME}-linux-arm ./cmd/...
	GOOS=darwin GOARCH=amd64 go build -o bin/${BINARY_NAME}-darwin-amd64 ./cmd/...
	GOOS=darwin GOARCH=arm64 go build -o bin/${BINARY_NAME}-darwin-arm64 ./cmd/...
	GOOS=windows GOARCH=amd64 go build -o bin/${BINARY_NAME}-windows-amd64 ./cmd/...

test:
	go test ./...

clean:
	go clean
	rm bin/${BINARY_NAME}-*
