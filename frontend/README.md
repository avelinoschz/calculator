# Frontend

React + TypeScript UI for the calculator project.

Built with Vite, plain CSS (no UI framework), and tested with Vitest and
React Testing Library.

## Prerequisites

- Node.js 20+
- npm

## Quick start

```sh
make frontend.run
# or
cd frontend && npm install && npm run dev
```

The dev server starts on `http://localhost:5173`. The backend must be
running on `:8080` for calculations to work — see
[`backend/README.md`](../backend/README.md).

## Project structure

```text
frontend/
  src/
    api/
      calculator.ts           ← typed fetch wrapper, no React imports
      calculator.test.ts
    components/
      CalculatorForm.tsx       ← form with client-side validation
      CalculatorForm.test.tsx
    App.tsx                    ← root component; result, error, loading state
    App.test.tsx
    main.tsx                   ← ReactDOM.createRoot entry point
    test/
      setup.ts                 ← Vitest + Testing Library global setup
  index.html
  vite.config.ts               ← build config, dev proxy (/api → :8080), Vitest config
  Dockerfile                   ← node:20-alpine build → nginx:alpine serve
```

Components never import `fetch` directly. All network calls go through
`src/api/`, keeping UI logic and data fetching independently testable.

## Dev proxy

In development, Vite proxies all `/api/*` requests to `http://localhost:8080`.
This means the frontend uses relative paths (`/api/v1/calculations`) and
never hard-codes the backend address.

In Docker Compose, there is no Vite — nginx takes over that role via
the `nginx.conf` at the repository root, which proxies `/api/` to the
`backend` service. See [`../nginx.conf`](../nginx.conf) for details.

## Makefile targets

| Target | Description |
| --- | --- |
| `make frontend.run` | Run the Vite dev server (port 5173) |
| `make frontend.test` | Run Vitest (single-run) |
| `make frontend.lint` | Run ESLint |
| `make frontend.build` | Build static assets → `frontend/dist/` |
| `make frontend.docker.build` | Build the frontend Docker image |

## Testing

```sh
make frontend.test
# or
cd frontend && npx vitest run
```

Tests are organised across three layers:

- **API layer** (`src/api/`) — plain TypeScript, mocks `fetch` with
  `vi.stubGlobal`; no React or DOM involved
- **Component layer** (`src/components/`) — renders components in
  isolation with React Testing Library; no network calls
- **Integration** (`App.test.tsx`) — full component tree with mocked
  fetch; asserts on visible output

## Linting

```sh
make frontend.lint
# or
cd frontend && npm run lint
```

Uses ESLint with React and TypeScript plugins (configured in
`eslint.config.js`).

## Build

```sh
make frontend.build
# or
cd frontend && npm run build
```

Output: `frontend/dist/`. TypeScript is compiled first, then Vite
bundles the result.

## Docker

```sh
make frontend.docker.build
# or
docker build -t calculator-frontend ./frontend
```

The Dockerfile uses a two-stage build:

- **Stage 1 (`build`)** — `node:20-alpine`; installs dependencies and
  runs `npm run build`
- **Stage 2 (`serve`)** — `nginx:alpine`; serves the compiled assets
  from `/usr/share/nginx/html` on port 80

Run the image standalone (note: API calls will not work without a
backend and nginx proxy config):

```sh
docker run -p 3000:80 calculator-frontend
```

For the full stack with API proxying, use Docker Compose from the
repository root:

```sh
make up
```
