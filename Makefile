BINARY_NAME=main
ZIP_NAME=${BINARY_NAME}.zip

all: build

env:
	@echo "go mod"
	go env -w GO111MODULE=auto
	go mod init devops_go

deps: env
	@echo "go get deps"
	go mod tidy
	go mod vendor

test: deps
	@echo "testing ..."
	go test -v ./...

build: deps
	@echo "build" $(BINARY_NAME)
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} main.go

package: build
	@echo "zip" $(BINARY_NAME)
	zip $(ZIP_NAME) $(BINARY_NAME)

docker-build:
	@echo "Docker build"
	docker build -t devops_golang . --no-cache
	@echo "success"

clean:
	@echo "Clean file..."
	rm -f $(BINARY_NAME) $(ZIP_NAME)

.PHONY:
	all
	deps
	build
	test
	docker-build
	clean
