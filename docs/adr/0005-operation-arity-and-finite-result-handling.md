# ADR 0005: Operation Arity and Finite Result Handling

## Status

Accepted

## Context

The calculator originally supported only binary operations and used a
single request shape: `{ "op": string, "a": number, "b": number }`.

Adding `power`, `sqrt`, and `percentage` creates two new design
pressures:

1. `sqrt` is naturally unary and should not require a fake second
   operand
2. advanced arithmetic increases the likelihood of non-finite results
   such as `NaN` or `Inf`, which cannot be encoded as valid JSON

The project still values a small API surface and explicit validation
over endpoint proliferation or hidden behavior.

## Decision

Keep the existing single calculation endpoint and evolve the request
contract to operation-specific arity:

- binary operations: `add`, `subtract`, `multiply`, `divide`, `power`,
  `percentage`
- unary operations: `sqrt`

Represent that distinction explicitly in the OpenAPI contract with
separate unary and binary schemas under a `oneOf`.

Reject mathematically invalid but well-formed requests with `422`
responses:

- `NEGATIVE_SQUARE_ROOT` for `sqrt(a)` when `a < 0`
- `NON_FINITE_RESULT` when a calculation would produce `NaN` or `Inf`

## Rationale

### 1. Preserve a small API surface

The existing `POST /api/v1/calculations` endpoint remains sufficient.
Adding operation-specific endpoints would duplicate handler behavior and
complicate the frontend for little benefit.

### 2. Avoid fake operands

Requiring `b` for `sqrt` would force clients and the UI to send a value
that is ignored by the backend. That is harder to understand, easier to
misuse, and less honest than an explicit unary request shape.

### 3. Keep the contract strict and machine-readable

Using explicit unary and binary schemas with `additionalProperties:
false` documents arity clearly and keeps generated or manually written
clients from guessing.

### 4. Prevent invalid JSON responses

Go's JSON encoder cannot serialize `NaN` or `Inf`. Catching non-finite
results in the domain layer preserves correctness and yields a stable,
intentional API error instead of an encoding failure.

### 5. Keep validation responsibility clear

Frontend validation remains a UX enhancement, while the backend remains
the source of truth for request shape, operand limits, and mathematical
validity.

## Consequences

### Positive

- No new endpoints
- Clear unary vs binary semantics
- Better UI ergonomics for `sqrt`
- Safer handling of advanced numeric edge cases
- Continued alignment between implementation and contract

### Trade-offs

- The request schema is slightly more complex than a single fixed object
- The backend and frontend both need operation-aware validation logic

## Notes

Percentage is defined as `a% of b`, computed as `(a / 100) * b`.
