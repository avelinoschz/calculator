VERSION ?= $(shell git rev-parse HEAD 2>/dev/null || echo "dev")
INSTALL_BIN_DIR := $(CURDIR)/bin
GOLANGCI_LINT_VERSION := v1.64.8

GO_BUILD_FLAGS := -ldflags "-X main.version=$(VERSION)"
GO_TEST_FLAGS := -shuffle=on -count=1
GO_COVERAGE_FLAGS := -coverprofile=coverage.out -covermode=atomic

export PATH := $(INSTALL_BIN_DIR):$(PATH)
export GOBIN := $(INSTALL_BIN_DIR)

.PHONY: help \
	backend.setup frontend.setup docs.setup setup \
	backend.run frontend.run run \
	backend.test frontend.test test \
	backend.coverage backend.coverage.html frontend.coverage coverage \
	backend.lint frontend.lint docs.lint lint \
	backend.format frontend.format format \
	backend.build frontend.build build \
	backend.clean frontend.clean clean \
	backend.docker.build frontend.docker.build docker.build \
	up down

help: ## Show available make targets
	@grep -E '^[a-zA-Z_.]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  %-25s %s\n", $$1, $$2}'

# ── Setup ─────────────────────────────────────────────────────────────────────

backend.setup: ## Install backend tooling and download Go module dependencies
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	cd backend && go mod download

frontend.setup: ## Install Node dependencies (npm ci)
	cd frontend && npm ci

docs.setup: ## Install docs tooling (markdownlint-cli)
	npm ci

setup: backend.setup frontend.setup docs.setup ## Bootstrap local environment (tools + dependencies)

# ── Run ───────────────────────────────────────────────────────────────────────

backend.run: ## Run the Go backend (port 8080)
	cd backend && go run $(GO_BUILD_FLAGS) ./cmd/server

frontend.run: ## Run the Vite dev server (port 5173)
	cd frontend && npm run dev

run: ## Run backend and frontend locally in parallel
	$(MAKE) -j2 backend.run frontend.run

# ── Test ──────────────────────────────────────────────────────────────────────

backend.test: ## Run Go tests
	cd backend && go test $(GO_TEST_FLAGS) ./...

frontend.test: ## Run Vitest (single-run)
	cd frontend && npx vitest run

test: backend.test frontend.test ## Run all tests

# ── Coverage ──────────────────────────────────────────────────────────────────

backend.coverage: ## Run Go tests and print per-function coverage summary
	cd backend && go test $(GO_TEST_FLAGS) $(GO_COVERAGE_FLAGS) ./... && go tool cover -func=coverage.out

backend.coverage.html: ## Run Go tests and generate coverage.html for browser viewing
	cd backend && go test $(GO_TEST_FLAGS) $(GO_COVERAGE_FLAGS) ./... && go tool cover -html=coverage.out -o coverage.html

frontend.coverage: ## Run Vitest with coverage report (text + lcov)
	cd frontend && npx vitest run --coverage

coverage: backend.coverage frontend.coverage ## Run all tests with coverage reports

# ── Lint ──────────────────────────────────────────────────────────────────────

backend.lint: ## Run golangci-lint
	cd backend && $(INSTALL_BIN_DIR)/golangci-lint run ./...

frontend.lint: ## Run ESLint
	cd frontend && npm run lint

docs.lint: ## Run markdownlint on all markdown files
	./node_modules/.bin/markdownlint **/*.md

lint: backend.lint frontend.lint docs.lint ## Run all linters

# ── Format ────────────────────────────────────────────────────────────────────

backend.format: ## Auto-fix Go lint issues
	cd backend && $(INSTALL_BIN_DIR)/golangci-lint run --fix ./...

frontend.format: ## Auto-fix frontend lint issues
	cd frontend && npm run lint -- --fix

format: backend.format frontend.format ## Auto-fix all lint issues

# ── Build ─────────────────────────────────────────────────────────────────────

backend.build: ## Build Go binary → backend/bin/server
	cd backend && go build $(GO_BUILD_FLAGS) -o bin/server ./cmd/server

frontend.build: ## Build frontend static assets → frontend/dist/
	cd frontend && npm run build

build: backend.build frontend.build ## Build backend binary and frontend assets

# ── Docker ────────────────────────────────────────────────────────────────────

backend.docker.build: ## Build the backend Docker image
	docker build -t calculator-backend ./backend

frontend.docker.build: ## Build the frontend Docker image
	docker build -t calculator-frontend ./frontend

docker.build: backend.docker.build frontend.docker.build ## Build all Docker images

# ── Compose ───────────────────────────────────────────────────────────────────

up: ## Start the full stack with Docker Compose
	docker compose up --build

down: ## Stop the full stack
	docker compose down

# ── Clean ─────────────────────────────────────────────────────────────────────

backend.clean: ## Remove backend build artifacts (backend/bin/)
	rm -rf backend/bin

frontend.clean: ## Remove frontend build artifacts (frontend/dist/)
	rm -rf frontend/dist

clean: backend.clean frontend.clean ## Remove all build artifacts
