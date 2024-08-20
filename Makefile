SHELL := /bin/bash

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

## Clean the previous and install the latest binary
install:
	go clean
	go mod tidy
	go install 
	@echo Make sure to add alias go-vote=\$$GOPATH/bin/go-vote to your \~/.bashrc

## Run tests
test:
	go test -failfast ./...

