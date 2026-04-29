# Load environment variables from .env
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Database Connection String
DB_URL="host=localhost user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) port=$(POSTGRES_PORT) sslmode=disable"

## Help: Help: show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

## Dev: Dev: Run the application with live reload (Air)
run:
	air

## Docker: Docker: Start infrastructure (Postgres, pgAdmin)
docker-up:
	docker compose up -d

docker-down:
	docker compose down

## Migrations: Create a new migration (usage: make migrate-create name=add_users_table)
migrate-create:
	goose -dir migrations create $(name) sql

## Migrations: Apply all pending migrations
migrate-up:
	goose -dir migrations postgres $(DB_URL) up

## Migrations: Rollback the last migration
migrate-down:
	goose -dir migrations postgres $(DB_URL) down

## Migrations: Check migration status
migrate-status:
	goose -dir migrations postgres $(DB_URL) status

## Build: Build: Compile the binary
build:
	go build -o bin/server main.go

## Test: Test: Run stress test (oha)
stress-test:
	oha -z 10s -c 50 http://localhost:8080/api/health

.PHONY: dev docker-up docker-down migrate-create migrate-up migrate-down migrate-status build stress-test help