# Requirements

## Goal

Build a small full-stack calculator application with a React frontend and a Go backend.

The solution should prioritize maintainability, clarity, and
correctness over extra features, while demonstrating production-minded
engineering practices.

## In Scope

### Core Functional Requirements

- Support the following operations:
  - Addition
  - Subtraction
  - Multiplication
  - Division

### Frontend

- Provide an intuitive UI to:
  - Enter input values
  - Select an operation
  - Submit a calculation request
  - Display results
- Validate user input before submitting
- Display clear error messages
- Support basic responsive behavior for mobile screens

### Backend

- Expose a REST API endpoint for calculator operations
- Validate request payloads
- Handle edge cases
- Return JSON responses
- Implement structured logging

## Quality Requirements

- Clean, readable, idiomatic code (frontend and backend)
- Separation of concerns (business logic vs transport layer)
- Unit tests for key backend and frontend behavior
- Structured logging in the backend (e.g., `log/slog`)
- Consistent error handling model
- Basic linting / static analysis
- Clear and concise documentation including:
  - Setup instructions
  - API usage
  - Design rationale

## API Contract Requirements

- Define the calculator API contract explicitly
- Provide a minimal OpenAPI specification
- Keep the API surface small and easy to understand
- Prefer a single calculation endpoint over multiple operation-specific
  endpoints unless there is a strong reason not to

### Proposed API Shape

- `POST /api/v1/calculations`

## Dev & Tooling Requirements

### Backend tooling

- Language: Go
- HTTP API: Go standard library (`net/http`)
- Testing: `testify`
- Logging: structured logging (`log/slog`)
- Linting: `golangci-lint`

### Frontend tooling

- Language: TypeScript
- Framework: React
- Build tool: Vite
- Testing: `vitest`, `React Testing Library`, `@testing-library/user-event`
- Styling: plain CSS (no UI framework)

## Developer Experience Requirements

- Provide a `Makefile` with common development commands
- Support both local development and Docker-based development flows
- Keep developer commands simple and discoverable
- Provide a small set of common targets for run, test, lint, build,
  and Docker-based workflows.

## Containerization

- Provide Docker support for running the application locally

### Backend container

- Use a **multi-stage Dockerfile**:
  - Stage 1: build the Go binary
  - Stage 2: minimal runtime image (e.g., distroless or alpine)

### Frontend container

- Dockerfile to build and serve the frontend

### Orchestration

- Use Docker Compose to orchestrate frontend + backend together for local execution

## CI (Continuous Integration)

- Provide a basic CI workflow (e.g., GitHub Actions) that:
  - Runs linters
  - Runs tests
  - Builds the project

Note: Full CD (deployment) is not required.

## Optional Scope

These items are lower priority and should only be implemented if the
core scope is complete:

- Exponentiation
- Square root
- Percentage
- Local observability support:
  - OpenTelemetry tracing
  - Jaeger integration
- Additional Make targets such as:
  - `make ci`
  - `make clean`

## Out of Scope

- Authentication or authorization
- Persistent storage or database
- Caching frequent calculations
- User accounts
- Calculation history
- Advanced styling or animations
- Complex state management
- Graphing or scientific calculator features
- Production deployment infrastructure
- Full CD / release automation
- Advanced observability platforms such as:
  - Datadog
  - Prometheus
  - Full metrics pipelines

## Acceptance Criteria

### Functional

- A user can enter valid numeric input(s) and select an operation
- A user can submit a calculation from the frontend
- The backend returns the correct result for core operations
- The frontend displays the result clearly
- Invalid input is rejected with a helpful error message
- Division by zero is handled safely and clearly

### Non-Functional

- Code is structured and easy to follow
- Business logic is testable independently of HTTP/UI layers
- Tests cover the most important behavior
- Logging provides useful debugging context
- Linting passes without critical issues
- The README explains setup and usage
- Common developer workflows are executable through the Makefile
- The API contract is documented in OpenAPI format
- The full stack can be started with Docker Compose

## Edge Cases

- Missing input
- Non-numeric input
- Empty strings
- Invalid operation
- Division by zero
- Malformed JSON request
- Extra unexpected fields in request
- Decimal values
- Negative values

## Prioritization

- P0:
  - Core operations
  - Validation and error handling
  - Unit tests
  - Structured logging
  - README
  - OpenAPI contract

- P1:
  - Docker support
  - Docker Compose orchestration
  - Basic CI workflow
  - Responsive UI polish
  - Makefile

- P2:
  - Optional operations
  - Observability (OpenTelemetry / Jaeger)
  - Extended developer tooling

## Project Notes

This project intentionally favors a clean, well-tested, and
well-documented core solution over feature expansion.
