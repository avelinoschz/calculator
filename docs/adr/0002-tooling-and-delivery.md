# ADR 0002: Tooling and Delivery

## Status

Accepted

## Context

The project needs lightweight tooling that keeps local development, CI,
and container builds aligned without adding unnecessary complexity.

## Decisions

### 1. Use structured JSON logging

The backend uses `log/slog` with a JSON handler.

Why:

- keeps logging dependency-free
- works well in local, CI, and container environments
- gives enough operational context for this scope

### 2. Standardize workflows through the Makefile

Common tasks are exposed through documented Make targets for setup, run,
test, coverage, lint, build, `ci`, and Docker flows.

Why:

- reduces reviewer guesswork
- keeps CI and local usage aligned
- avoids hidden one-off commands

### 3. Pin backend tooling locally

Backend tools such as `golangci-lint` are installed into the
project-local `bin/` directory through `GOBIN`.

Why:

- avoids reliance on global tool installs
- makes local and CI usage more repeatable

### 3a. Pin the recommended local toolchain

The repository includes a root `.tool-versions` file for the preferred
Go and Node.js versions used in development.

Why:

- makes local setup more reproducible for contributors using `asdf`
- reduces version drift across backend and frontend workflows
- complements, rather than replaces, language-native manifests

### 4. Use multi-stage Docker images

- backend: `golang:alpine` -> `distroless/static`
- frontend: `node:20-alpine` -> `nginx:alpine`

Why:

- keeps runtime images small
- separates build tooling from runtime
- keeps the container path close to production-minded defaults

### 5. Embed version metadata in backend builds

The backend binary receives `main.version` through ldflags in local
builds, CI builds, and Docker builds.

Why:

- makes startup logs traceable to a specific revision
- keeps versioning cheap and visible

### 6. Keep CI narrow and useful

CI runs lint, test, and build using the documented Make targets, and the
same validation gate is exposed locally as `make ci`.

Why:

- covers the main handoff risks
- avoids duplicating workflow logic outside the Makefile

### 7. Support graceful shutdown

The backend handles `SIGINT` and `SIGTERM` with a 10-second shutdown timeout.

Why:

- avoids abrupt request interruption
- matches Docker and local runtime expectations

## Consequences

### Positive

- local, CI, and Docker paths stay close together
- review and handoff workflows are easier to follow
- tooling remains proportionate to project size
- repo-local AI workflows can be documented without changing the
  product's runtime model

### Trade-offs

- less automation than a larger production system
- operational visibility is intentionally basic
