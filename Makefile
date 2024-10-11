SHELL := /bin/bash

CONFIG_DIR = ${HOME}/.config/go-vote

default: help

## This help screen. Requires targets to have comments with "##".
help:
	@printf "Available targets:\n\n"
	@awk '/^[a-zA-Z\-\0-9%:\\]+/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = $$1; \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
	gsub("\\\\", "", helpCommand); \
	gsub(":+$$", "", helpCommand); \
			printf "  \x1b[32;01m%-35s\x1b[0m %s\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST) | sort -u
	@printf "\n"


## Run migrations on the database
db/migrate:
	dbmate --url "sqlite:${DATABASE_URL}" up
	dbmate --url "sqlite:test.sqlite" up

## Rollback the previous migration
db/rollback:
	dbmate --url "sqlite:${DATABASE_URL}" down
	dbmate --url "sqlite:test.sqlite" down

## Generate the structs and queries using sqlc. Results can be found in pkg/db
db/generate:
	sqlc generate

## Run HTTP server for the docs
docs/serve:
	@python -m http.server -d docs/

## Clean the previous and install the latest binary.
install:
	@echo "Cleaning..."
	@go clean
	@go mod tidy
	@echo "Installing..."
	@go install
	@echo Make sure to add alias go-vote=\$$GOPATH/bin/go-vote to your \~/.bashrc. Replacing GOPATH with your own

## Run CI linting
lint:
	@golangci-lint run

## Run tests
test: test/go test/http

## Run the tests for the Go codebase
test/go:
	@touch test.sqlite
	@dbmate --url "sqlite:test.sqlite" up
	@echo "Testing Go library..."
	@go test -failfast ./...

## Run the tests for the API endpoints based on the OpenAPI specification
test/http:
	@echo "Testing API..."
	@openapi-to-hurl docs/schema.yaml --out-dir tests/ --validation body --grouping path

## Start hot-reloading for the binary, useful for iterative development
watch:
	ag -l --go | entr -s 'make install'
