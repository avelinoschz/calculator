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
