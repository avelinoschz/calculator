# ADR 0003: Frontend Architecture

## Status

Accepted

## Context

Phase 2 adds a React + TypeScript frontend to the existing Go backend.
The goal is to keep the frontend consistent with the project's principles:
simplicity, maintainability, and correctness over feature breadth.

## Decisions

### 1. Use Vite as the build tool

The frontend uses Vite instead of Create React App or Next.js.

Rationale:

- Fast dev server with native ESM and HMR
- Lightweight configuration with no framework overhead
- Next.js adds SSR complexity not required for a single-page calculator
- Create React App is unmaintained and adds unnecessary defaults

### 2. Use no CSS framework

Plain CSS is used without Tailwind, Bootstrap, or any utility library.

Rationale:

- Sufficient for a single-form layout
- Avoids adding a build-time dependency for minimal visual scope
- Keeps styles explicit and easy to audit

### 3. Use React useState only

No external state management library (Redux, Zustand, Context) is used.

Rationale:

- The app has three state values: result, error, loading
- Component tree is shallow; prop drilling is not a concern
- Adding a state library would be overengineering at this scope

### 4. Isolate the API layer in `src/api/`

All fetch calls are in `src/api/calculator.ts`. Components never import
fetch or call the backend directly.

Rationale:

- Separates data fetching from UI rendering
- `src/api/` is testable without rendering components
- Components are testable without mocking fetch
- Mirrors the backend's separation of domain logic from transport logic

### 5. Use Vitest and React Testing Library

Tests use Vitest as the test runner, React Testing Library for component
rendering, and `@testing-library/user-event` for realistic interaction simulation.
`@testing-library/jest-dom` provides DOM matchers; despite the name it does not
require Jest and works with Vitest's global test API.

Tests are structured in three layers that mirror the architectural separation:

- **API layer** (`src/api/`): plain TypeScript tests; `fetch` is stubbed with
  `vi.stubGlobal`. No React rendering involved.
- **Component layer** (`src/components/`): components rendered in isolation via
  React Testing Library; no network calls.
- **Integration** (`App`): full component tree rendered with mocked fetch;
  asserts on visible output after async state updates.

Rationale:

- Vitest runs inside the Vite pipeline; no separate Jest configuration needed
- React Testing Library encourages testing user-visible behaviour, not internals
- Three-layer structure keeps each concern independently testable
- Consistent with the project's preference for lightweight, well-integrated tooling

### 6. Use no React Router

The application is a single page with no navigation.

Rationale:

- No multi-page or URL-driven use case exists in the current scope
- Adding routing would be unnecessary complexity

### 7. Use the Vite dev proxy for API requests

The Vite dev server proxies `/api` to `http://localhost:8080`.

Rationale:

- Avoids hardcoding the backend URL in application code
- Eliminates CORS issues during local development without backend changes
- In production, the proxy is replaced by the nginx container's network routing

### 9. Use CSS custom properties with fluid `clamp()` values for responsive design

All design tokens (colors, font sizes, spacing, layout dimensions) are declared
as CSS custom properties in the `:root` block of `App.css`. Fluid values use
`clamp(minimum, preferred-vw, maximum)` so the layout scales continuously with
the viewport without breakpoints.

Rationale:

- All tuning knobs are in one place; changing a value updates every element that
  references it
- `clamp()` produces smooth, continuous scaling instead of discrete breakpoint
  jumps
- No JavaScript or framework feature needed — pure CSS
- Meets the requirements' basic responsive design criterion with minimal complexity

### 8. Use TypeScript strict mode

`tsconfig.json` enables `"strict": true`.

Rationale:

- Catches null/undefined errors at compile time
- Keeps type definitions explicit
- Aligns with the project's emphasis on correctness

## Consequences

### Positive

- Minimal dependency surface
- Each concern is independently testable
- Dev setup is fast and has no framework magic to debug

### Trade-offs

- No server-side rendering or code splitting for larger apps
- Plain CSS with no utility framework; all responsive behaviour is hand-authored
- No global state management if scope expands significantly

## Notes

These decisions are intentionally scoped to a single-form calculator.
They should be revisited if the application grows in page count,
data complexity, or team size.
