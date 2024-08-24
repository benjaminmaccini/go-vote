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

## Clean the previous and install the latest binary.
install:
	@echo "Installing..."
	@mkdir -p $(CONFIG_DIR)
    @cp -n .go-vote-template.yaml $(CONFIG_DIR)/config.yaml || true
	@go clean
	@go mod tidy
	@go install
	@echo Make sure to add alias go-vote=\$$GOPATH/bin/go-vote to your \~/.bashrc. Replacing GOPATH with your own

## Run tests
test:
	go test -failfast ./...
