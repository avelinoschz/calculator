# ADR 0002: Tooling and Delivery

## Status

Accepted

## Context

The project should demonstrate production-minded engineering practices
while remaining lightweight and focused.

## Decisions

### 1. Use structured logging

The backend will use Go's `log/slog` for structured logging.

Logs should be emitted in JSON format to remain compatible with common
log aggregation and observability platforms such as AWS, GCP, and
Datadog.

Rationale:

- Improves observability
- Keeps logging simple and dependency-free
- Makes logs easier to ingest, parse, and query in centralized logging systems

### 2. Use lightweight tooling for quality gates

Rationale:

- Improves code quality
- Avoids heavy setup

### 9. Pin tool binaries locally in bin/

Developer tools (e.g. `golangci-lint`) are installed into a project-local
`bin/` directory via `GOBIN`. The Makefile exports `GOBIN := $(CURDIR)/bin`
so `go install` always writes there, not to the system `$GOPATH/bin`.

A `make backend.setup` target installs all backend tooling at the pinned
version. CI runs the same target, ensuring local and CI environments are
identical.

Tool versions are pinned to exact semver in the Makefile (e.g.
`GOLANGCI_LINT_VERSION := v1.64.8`) rather than `latest`.

Rationale:

- No dependency on system-installed tool versions
- Reproducible across machines and CI runners
- Avoids version mismatch between the project's Go version and pre-built
  tool binaries (e.g. `golangci-lint` built for an older Go release)
- Using `latest` can silently introduce breaking lint rules between runs

### 10. Embed git SHA in the binary via ldflags

The `VERSION` variable is captured from `git rev-parse HEAD` at
build and run time:

```makefile
VERSION ?= $(shell git rev-parse HEAD 2>/dev/null || echo "dev")
GO_BUILD_FLAGS := -ldflags "-X main.version=$(VERSION)"
```

The value is injected into a package-level `var version = "dev"` in
`cmd/server/main.go` and logged at startup:

```json
{"level":"INFO","msg":"starting","version":"7880150..."}
```

Outside a git repository the version falls back to `"dev"`.

Rationale:

- Zero-cost traceability — no extra package or runtime overhead
- Visible immediately in startup logs; no need to query a separate API
- Makes it trivial to correlate a running instance with the exact commit

### 3. Use Docker pragmatically

- Backend uses multi-stage Docker build
- Frontend has its own Dockerfile
- Docker Compose orchestrates full stack

Rationale:

- Reproducible environment
- Easy reviewer setup

### 4. Provide a Makefile

The project will include a Makefile with common commands.

Goals:

- Simplify workflows
- Standardize commands

### 5. Provide CI readiness (not full CD)

A minimal GitHub Actions workflow will run:

- lint
- tests
- build

Rationale:

- Demonstrates good practices
- Avoids unnecessary complexity

### 6. Prioritize core over optional features

Core features must be completed before optional enhancements.

Rationale:

- Ensures a complete and stable submission
- Avoids partially implemented features

### 7. Treat observability as a production-minded consideration (not a core dependency)

Advanced observability is valuable, but should not displace core
implementation quality in a small project.

Current approach:

- Structured logging is in scope
- Logs are emitted in JSON format for compatibility with aggregation systems
- Advanced tracing and metrics are optional

Rationale:

- Keeps implementation focused on core functionality
- Demonstrates awareness of production practices without overengineering
- Aligns with a focused implementation while still signaling maturity

### 8. Implement graceful shutdown

The server handles `SIGINT` and `SIGTERM` signals. On receipt, it calls
`http.Server.Shutdown` with a 10-second context timeout. In-flight requests
are allowed to complete; new connections are rejected immediately.

Rationale:

- Prevents abrupt connection drops during restarts or container stops
- Aligns with production expectations without adding complexity
- Required for correct behaviour in Docker and orchestrated environments

## Consequences

### Positive

- Strong developer experience
- Easy to run and evaluate
- Balanced use of tooling

### Trade-offs

- Limited automation compared to production systems
- Observability remains minimal

## Notes

Structured JSON logging provides a lightweight operational baseline
without expanding the initial scope.
