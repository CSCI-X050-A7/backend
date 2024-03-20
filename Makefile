.PHONY: all clean critic security lint test swag build run

APP_NAME = backend
BUILD_DIR = ./build

all:
	air

clean:
	rm -rf $(BUILD_DIR)/*
	rm -f *.out

lint:
	golangci-lint run

test: clean lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

swag:
	swag init

BUILT_AT := $(shell date +'%F %T %z')
GO_VERSION := $(shell go version)
GIT_AUTHOR := $(shell git show -s --format='format:%aN <%ae>' HEAD)
GIT_COMMIT := $(shell git log --pretty=format:"%H" -1)
VERSION := dev
FLAGS := "-s -w \
-X 'github.com/CSCI-X050-A7/backend/pkg/config.BuiltAt=$(BUILT_AT)' \
-X 'github.com/CSCI-X050-A7/backend/pkg/config.GoVersion=$(GO_VERSION)' \
-X 'github.com/CSCI-X050-A7/backend/pkg/config.GitAuthor=$(GIT_AUTHOR)' \
-X 'github.com/CSCI-X050-A7/backend/pkg/config.GitCommit=$(GIT_COMMIT)' \
-X 'github.com/CSCI-X050-A7/backend/pkg/config.Version=$(VERSION)'"
build: swag clean
	go mod tidy
	CGO_ENABLED=1 go build -ldflags=$(FLAGS) -o $(BUILD_DIR)/$(APP_NAME) main.go

run: build
	$(BUILD_DIR)/$(APP_NAME)

run-without-build:
	$(BUILD_DIR)/$(APP_NAME)
