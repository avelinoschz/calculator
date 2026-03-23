# ADR 0003: Frontend Architecture

## Status

Accepted

## Context

The frontend should stay small, easy to review, and closely aligned with
the backend contract.

## Decisions

### 1. Use Vite with React and TypeScript

Why:

- fast local feedback
- minimal configuration
- strict typing without extra framework weight

### 2. Keep state local with `useState`

Why:

- the UI only manages form input, loading, result, and error state
- external state management would be unnecessary complexity

### 3. Isolate the API layer in `src/api/`

Components do not call `fetch` directly.

Why:

- keeps network behavior testable without rendering React
- keeps UI code focused on interaction and rendering

### 4. Validate in the UI, but treat the backend as authoritative

The frontend blocks obviously invalid input before submission, but
backend validation still defines the contract.

Why:

- improves UX
- preserves a single source of truth

### 5. Test in three layers

- API client tests
- isolated component tests
- full-app integration tests

Why:

- keeps concerns independently testable
- matches the project’s separation-of-concerns goals

### 6. Use plain CSS

Why:

- sufficient for a single-form application
- avoids framework or utility-class overhead

### 7. Use proxy-based API routing

- Vite proxies `/api/*` to `:8080` in development
- nginx proxies `/api/*` to the backend in the containerized flow

Why:

- avoids hard-coded backend URLs in browser code
- keeps local and Docker networking straightforward

## Consequences

### Positive

- small dependency surface
- easy-to-review component tree
- clear separation between UI and data access

### Trade-offs

- no routing or global state primitives if scope grows later
- responsive behavior remains hand-authored CSS
