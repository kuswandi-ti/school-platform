.DEFAULT_GOAL := help

GO_MODULES := $(shell find services packages -name go.mod -not -path "*/vendor/*" -exec dirname {} \; 2>/dev/null)
SERVICE_PATH := services/$(service)

.PHONY: help setup up down restart logs ps test lint fmt build clean proto openapi-check event-schema-check

help:
	@echo "school-platform local commands"
	@echo ""
	@echo "Usage:"
	@echo "  make <target>"
	@echo "  make <target> service=<service-name>"
	@echo ""
	@echo "Local stack:"
	@echo "  setup                 Prepare local env files when missing"
	@echo "  up                    Start Docker Compose local dependencies"
	@echo "  down                  Stop Docker Compose local dependencies"
	@echo "  restart               Restart Docker Compose local dependencies"
	@echo "  ps                    Show Docker Compose service status"
	@echo "  logs                  Show Docker Compose logs"
	@echo "  logs service=<name>   Show logs for one Docker Compose service"
	@echo ""
	@echo "Go:"
	@echo "  test                  Run go test in all Go modules"
	@echo "  test service=<name>   Run go test in services/<name>"
	@echo "  lint                  Run go vet in all Go modules"
	@echo "  lint service=<name>   Run go vet in services/<name>"
	@echo "  fmt                   Run gofmt in all Go modules"
	@echo "  build                 Build all Go service modules"
	@echo "  build service=<name>  Build services/<name>"
	@echo ""
	@echo "Contracts:"
	@echo "  proto                 Placeholder proto generation/check target"
	@echo "  openapi-check         Validate OpenAPI file presence"
	@echo "  event-schema-check    Validate event schema file presence"
	@echo ""
	@echo "Maintenance:"
	@echo "  clean                 Remove local bin output and Go test cache"

setup:
	@if [ ! -f .env ] && [ -f .env.example ]; then \
		cp .env.example .env; \
		echo "Created .env from .env.example"; \
	else \
		echo ".env already exists or .env.example is missing"; \
	fi

up:
	docker compose up -d

down:
	docker compose down

restart: down up

ps:
	docker compose ps

logs:
	@if [ -n "$(service)" ]; then \
		docker compose logs -f "$(service)"; \
	else \
		docker compose logs -f; \
	fi

test:
	@$(MAKE) go-each command='go test ./...'

lint:
	@$(MAKE) go-each command='go vet ./...'

fmt:
	@$(MAKE) go-each command='go fmt ./...'

build:
	@if [ -n "$(service)" ]; then \
		if [ ! -f "$(SERVICE_PATH)/go.mod" ]; then \
			echo "No go.mod found in $(SERVICE_PATH)"; \
			exit 1; \
		fi; \
		mkdir -p bin; \
		(cd "$(SERVICE_PATH)" && go build -o "../../bin/$(service)" ./cmd/server); \
	else \
		set -e; \
		for module in $(GO_MODULES); do \
			if [ -d "$$module/cmd/server" ]; then \
				name=$$(basename "$$module"); \
				mkdir -p bin; \
				echo "Building $$module"; \
				(cd "$$module" && go build -o "../../bin/$$name" ./cmd/server); \
			else \
				echo "Skipping $$module: no cmd/server"; \
			fi; \
		done; \
	fi

clean:
	rm -rf bin
	@$(MAKE) go-each command='go clean -testcache'

proto:
	@echo "Proto generation/check tooling is not configured yet."
	@echo "Contract placeholders are under packages/proto."
	@test -d packages/proto/common/v1
	@test -d packages/proto/identity/v1
	@test -d packages/proto/schoolcore/v1
	@test -d packages/proto/admission/v1
	@test -d packages/proto/academic/v1
	@test -d packages/proto/finance/v1
	@test -d packages/proto/communication/v1
	@test -d packages/proto/reporting/v1
	@echo "Proto placeholder directories are available."

openapi-check:
	@test -f packages/openapi/api-gateway.v1.yaml
	@echo "OpenAPI skeleton exists: packages/openapi/api-gateway.v1.yaml"

event-schema-check:
	@test -f packages/events/envelope.schema.json
	@echo "Event envelope schema exists: packages/events/envelope.schema.json"

.PHONY: go-each
go-each:
	@if [ -n "$(service)" ]; then \
		if [ ! -f "$(SERVICE_PATH)/go.mod" ]; then \
			echo "No go.mod found in $(SERVICE_PATH)"; \
			exit 1; \
		fi; \
		echo "Running in $(SERVICE_PATH): $(command)"; \
		(cd "$(SERVICE_PATH)" && $(command)); \
	else \
		if [ -z "$(GO_MODULES)" ]; then \
			echo "No Go modules found"; \
			exit 0; \
		fi; \
		set -e; \
		for module in $(GO_MODULES); do \
			echo "Running in $$module: $(command)"; \
			(cd "$$module" && $(command)); \
		done; \
	fi
