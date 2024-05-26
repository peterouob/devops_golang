BINARY_NAME=main
ZIP_NAME=${BINARY_NAME}.zip

all: build

deps:
	@echo "Running go mod tidy"
	go mod tidy

env:
	@echo "Setting up Go environment"
	@if [ ! -f go.mod ]; then go mod init devops_go; fi
	go env -w GO111MODULE=auto

test: deps
	@echo "testing ..."
	go test -v ./...

build: deps
	@echo "build" $(BINARY_NAME)
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/${BINARY_NAME} main.go

deploy_prod: build
	@echo "deploy ..."
	serverless deploy --stage prod

package: build
	@echo "zip" $(BINARY_NAME)
	zip $(ZIP_NAME) bin/$(BINARY_NAME)

docker-build:
	@echo "Docker build"
	docker build -t devops_golang . --no-cache
	@echo "success"

clean:
	@echo "Clean file..."
	rm -f $(BINARY_NAME) $(ZIP_NAME)

.PHONY:
	all
	env
	deps
	build
	test
	docker-build
	clean
