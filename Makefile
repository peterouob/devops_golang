BINARY_NAME=main
ZIP_NAME=${BINARY_NAME}.zip
GO_FILES=${wildcard *.go}

all: build

deps:
	go mod tidy

build: deps
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} main.go

.PHONY:
	all
	deps
	build
