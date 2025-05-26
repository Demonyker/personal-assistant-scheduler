ifneq ($(wildcard .env),)
include .env
export
else
$(warning WARNING: .env file not found! Using .env.example)
include .env.example
export
endif

LOCAL_BIN:=$(CURDIR)/bin

deps: ### deps tidy + verify
	go mod tidy && go mod verify
.PHONY: deps

deps-audit: ### check dependencies vulnerabilities
	govulncheck ./...
.PHONY: deps-audit

format: ### Run code formatter
	gofumpt -l -w .
	gci write . --skip-generated -s standard -s default
.PHONY: format

linter-golangci: ### check by golangci linter
	golangci-lint run
.PHONY: linter-golangci

linter-hadolint: ### check by hadolint linter
	git ls-files --exclude='Dockerfile*' --ignored | xargs hadolint
.PHONY: linter-hadolint

migrate-create:  ### create new migration
	migrate create -ext sql -dir migrations '$(word 2,$(MAKECMDGOALS))'
.PHONY: migrate-create

migrate-up: ### migration up
	migrate -path migrations -database '$(DB_URL)' up
.PHONY: migrate-up

bin-deps: ### install tools
	GOBIN=$(LOCAL_BIN) go install tool
.PHONY: bin-deps

pre-commit: format linter-golangci ### run pre-commit
.PHONY: pre-commit
