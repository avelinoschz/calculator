# Frontend

React + TypeScript UI for the calculator project.

## Run

```sh
make frontend.setup
make frontend.run
```

The Vite dev server runs on `http://localhost:5173` and proxies `/api/*` to `http://localhost:8080`.

## Structure

```text
frontend/
  src/api/           typed API client
  src/components/    form UI and validation
  src/App.tsx        result, error, and loading state
  vite.config.ts     Vite, dev proxy, and Vitest config
  nginx.conf         production nginx config baked into the image
  Dockerfile         multi-stage frontend image
```

## Configuration

Optional Vite environment variables:

| Variable | Description |
| --- | --- |
| `VITE_CALC_MIN` | Minimum operand value for client-side validation |
| `VITE_CALC_MAX` | Maximum operand value for client-side validation |

`make frontend.run` sources the repository `.env` file automatically
when present.

These values are build-time inputs for the browser bundle. They improve
UX only; the backend remains the authoritative validator.

## UI Behavior

- API calls are isolated in `src/api/calculator.ts`
- components do not call `fetch` directly
- empty inputs are rejected before submission
- partial garbage like `12abc` is rejected before submission
- non-finite values such as `Infinity` are rejected before submission
- backend error messages are surfaced in the UI

## Make Targets

| Target | Description |
| --- | --- |
| `make frontend.setup` | Install Node dependencies (`npm ci`) |
| `make frontend.run` | Run the Vite dev server |
| `make frontend.test` | Run Vitest |
| `make frontend.coverage` | Run Vitest with coverage |
| `make frontend.lint` | Run ESLint |
| `make frontend.format` | Auto-fix frontend lint issues |
| `make frontend.build` | Build `frontend/dist/` |
| `make frontend.clean` | Remove frontend build artifacts |
| `make frontend.docker.build` | Build the frontend Docker image |

## Testing

```sh
make frontend.test
make frontend.coverage
make frontend.lint
```

Tests cover three layers:

- API layer
- component layer
- full-app integration

## Build and Docker

```sh
make frontend.build
make frontend.docker.build
```

The production image serves static assets with nginx and includes its
own proxy configuration for `/api/`, so it no longer depends on a
Compose bind mount.
