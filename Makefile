# Load .env file
include .env
export

ODIR 	:= build/_output
UNAME := $(shell uname)

download:
	go mod download
	go mod tidy

build: generate format
	go build -o ./$(ODIR)/account-service ./cmd

run: build
	./$(ODIR)/account-service api

generate: tool-mockgen
	go generate ./...

format:
	go fmt ./...

bin:
	@mkdir -p bin

install-go-migrate-tool: bin
ifneq (,$(wildcard bin/migrate))
    # do not download again
else ifeq ($(UNAME), Linux)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.linux-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else ifeq ($(UNAME), Darwin)
	@curl -sSfL https://github.com/golang-migrate/migrate/releases/download/v4.18.2/migrate.darwin-amd64.tar.gz | tar zxf - --directory /tmp \
	&& cp /tmp/migrate bin/
else
	@echo "Your OS is not supported."
endif

migrate: install-go-migrate-tool
	@bin/migrate -source file://db/migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" $(MIGRATE_ARGS) $(N)

create-db-migration: install-go-migrate-tool
	@bin/migrate create -ext sql -dir db/migrate $(MIGRATE_NAME)

tool-mockgen:
	@go install github.com/golang/mock/mockgen@v1.6.0
