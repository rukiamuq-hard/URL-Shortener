# Deploy & Migration Makefile (Run from inside 'deploy/' directory)
# Usage:
#   cd deploy
#   make up              - start docker containers
#   make migrate-up      - run migrations
#   make down            - stop docker containers

DB_DSN ?= postgres://postgres:postgres@127.0.0.1:5432/url_shortener?sslmode=disable
MIGRATIONS_DIR := ../internal/platform/postgres/migration

.PHONY: up down restart migrate-up migrate-down migrate-status

up:
	docker compose up -d

down:
	docker compose down

restart:
	docker compose restart

migrate-up:
	@echo "Running migrations UP..."
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" up

migrate-down:
	@echo "Running migrations DOWN..."
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down

migrate-status:
	@echo "Checking migration status..."
	goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" status
