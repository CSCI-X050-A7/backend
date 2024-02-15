.PHONY: clean critic security lint test swag build run

APP_NAME = backend
BUILD_DIR = ./build
MIGRATIONS_FOLDER = ./platform/migrations
DB_NAME = fiber_go_api
DB_USER = dev
DB_PASS = dev
DATABASE_URL = postgres://$(DB_USER):$(DB_PASS)@localhost/$(DB_NAME)?sslmode=disable

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out

critic:
	gocritic check -enableAll ./...

security:
	gosec -quiet ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
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

docker.run: docker.setup docker.postgres docker.fiber
	@echo "\n===========FGB==========="
	@echo "App is running...\nVisit: http://localhost:5000 OR http://localhost:5000/swagger/"

docker.setup:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network
	docker volume create fibergb-pgdata

docker.fiber.build: swag
	docker build -t fibergb:latest .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name fibergb-api \
		--network dev-network \
		-p 5000:5000 \
		fibergb

docker.postgres:
	docker run --rm -d \
		--name fibergb-postgres \
		--network dev-network \
		-e POSTGRES_USER=dev \
		-e POSTGRES_PASSWORD=dev \
		-e POSTGRES_DB=fiber_go_api \
		-v fibergb-pgdata:/var/lib/postgresql/data \
		-p 5432:5432 \
		postgres

docker.stop: docker.stop.fiber docker.stop.postgres

docker.stop.fiber:
	docker stop fibergb-api || true

docker.stop.postgres:
	docker stop fibergb-postgres || true

docker.dev:
	docker-compose up
