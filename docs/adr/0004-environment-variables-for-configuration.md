# ADR 0004: Environment Variables for Runtime Configuration

## Status

Accepted

## Context

The calculator needs a way to constrain the range of accepted operand
values (`a` and `b`). The limit values must be adjustable per deployment
— development, staging, and production environments may have different
acceptable ranges.

The key question is where and how these values should live: hardcoded
constants, a configuration file, a database, or environment variables.

## Decision

Use environment variables (`CALC_MIN`, `CALC_MAX` on the backend;
`VITE_CALC_MIN`, `VITE_CALC_MAX` on the frontend) to supply the operand
limits at runtime.

When a variable is absent, no limit is applied on that side. This keeps
the default behavior backward compatible — existing deployments without
these variables behave exactly as before.

## Rationale

### 1. Values are configurable at runtime without redeployment

Environment variables are resolved when the process starts. Changing a
limit requires only restarting the service with a different variable
value. No code change, rebuild, or redeployment of a new artifact is
needed.

This is one of the clearest demonstrations of runtime configuration: the
same binary or container image behaves differently depending solely on
the environment it runs in.

### 2. Follows the twelve-factor app principle

The [twelve-factor methodology](https://12factor.net/config) (factor III)
states that configuration that varies between deployments should be
stored in the environment, not in code. Operand limits are a textbook
example: they are deployment-specific, not application logic.

### 3. Environment-specific behavior with no code branches

Different environments (local dev, CI, staging, production) can apply
different limits without any conditional logic in the source code:

```sh
# Local dev — permissive
make backend.run

# Staging — moderate limits
CALC_MIN=-10000 CALC_MAX=10000 make backend.run

# Production — strict limits
CALC_MIN=-1000 CALC_MAX=1000 make backend.run
```

### 4. No sensitive values in source control

Environment variables keep configuration out of the codebase. Even
though operand limits are not secrets, establishing this pattern ensures
the mechanism is in place when values that _are_ sensitive (API keys,
credentials) need to follow the same path.

### 5. Native support in Docker and Docker Compose

Container runtimes treat environment variables as a first-class concept.
Limits can be injected without modifying the image or rebuilding:

```sh
docker run -e CALC_MIN=-1000 -e CALC_MAX=1000 calculator-backend
```

In Docker Compose they are declared under `environment:` or sourced
from a `.env` file — no changes to the Dockerfile required.

### 6. Minimal implementation complexity

Go's `os.Getenv` and `strconv.ParseFloat` are sufficient to read and
parse the variables. No configuration library or additional dependency
is needed, which aligns with the project's goal of minimal dependencies.

### 7. Backward compatibility by default

The absence of an environment variable is handled gracefully: the server
falls back to `±Inf`, meaning no limit is applied. Existing deployments
that do not set these variables continue to work without modification.

## Consequences

### Positive

- Limits can be changed without a code change or artifact rebuild
- Demonstrates runtime configuration via environment variables
- Aligns with twelve-factor app and container-native practices
- No new dependencies introduced
- Backward compatible — unset variables impose no limits

### Trade-offs

- Unparseable values (e.g. `CALC_MIN=abc`) fall back to the default and
  log a warning rather than failing fast at startup — this is intentional
  to preserve backward compatibility for deployments that omit these
  variables. Structurally invalid configurations (e.g. `CALC_MIN` greater
  than `CALC_MAX`, or NaN values) are caught by `calculator.NewService`
  at startup and cause an immediate fatal error.
- Frontend limits are baked in at build time (Vite inlines
  `VITE_CALC_*` during `npm run build`); changing them requires a
  frontend rebuild, unlike the backend which reads them on each start

## Notes

Frontend environment variables use the `VITE_` prefix, which is
required by Vite to expose them to browser code via `import.meta.env`.
They are static after the build, making them suitable for default UX
validation but not a substitute for the authoritative backend check.
