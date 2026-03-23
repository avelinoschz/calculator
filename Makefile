INSTALL_BIN_DIR := $(CURDIR)/bin
GOLANGCI_LINT_VERSION := latest

export PATH := $(INSTALL_BIN_DIR):$(PATH)
export GOBIN := $(INSTALL_BIN_DIR)

.PHONY: help \
	backend.setup \
	backend.run frontend.run run \
	backend.test frontend.test test \
	backend.lint frontend.lint lint \
	backend.build frontend.build build \
	backend.docker.build frontend.docker.build docker.build \
	up down

help: ## Show available make targets
	@grep -E '^[a-zA-Z_.]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  %-25s %s\n", $$1, $$2}'

# ── Setup ─────────────────────────────────────────────────────────────────────

backend.setup: ## Install backend tooling into bin/
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

# ── Run ───────────────────────────────────────────────────────────────────────

backend.run: ## Run the Go backend (port 8080)
	cd backend && go run ./cmd/server

frontend.run: ## Run the Vite dev server (port 5173)
	cd frontend && npm run dev

run: ## Run backend and frontend locally in parallel
	$(MAKE) -j2 backend.run frontend.run

# ── Test ──────────────────────────────────────────────────────────────────────

backend.test: ## Run Go tests
	cd backend && go test ./...

frontend.test: ## Run Vitest (single-run)
	cd frontend && npx vitest run

test: backend.test frontend.test ## Run all tests

# ── Lint ──────────────────────────────────────────────────────────────────────

backend.lint: ## Run golangci-lint
	cd backend && $(INSTALL_BIN_DIR)/golangci-lint run ./...

frontend.lint: ## Run ESLint
	cd frontend && npm run lint

lint: backend.lint frontend.lint ## Run all linters

# ── Build ─────────────────────────────────────────────────────────────────────

backend.build: ## Build Go binary → backend/bin/server
	cd backend && go build -o bin/server ./cmd/server

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
