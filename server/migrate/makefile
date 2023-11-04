# Makefile for Bun migrations server

.PHONY: init migrate rollback status create-go create-sql help

# Run init db
init:
	@go run . db init

# Run migrations
migrate:
	@BUNDEBUG=2 go run . db migrate

# Rollback migrations
rollback:
	@go run . db rollback

# View status of migrations
status:
	@go run . db status

# Create a Go migration
create-go:
ifdef name
	@go run . db create_go $(name)
else
	@echo "使用法: make create-go name=<マイグレーション名>"
endif

# Create a SQL migration
create-sql:
ifdef name
	@go run . db create_sql $(name)
else
	@echo "使用法: make create-sql name=<マイグレーション名>"
endif

# Display help for db command
help:
	@go run . db

# Default target
default: help

