# Default db migration settings
export DB_USERNAME ?= account_domain_rw_dev
export DB_PASSWORD ?= passdev
export DB_HOST ?= localhost
export DB_PORT ?= 3306
export DB_NAME ?= accountdb

UNAME := $(shell uname)

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

generate-db-migration: install-go-migrate-tool
	@bin/migrate create -ext sql -dir db/migrate $(MIGRATE_NAME)
