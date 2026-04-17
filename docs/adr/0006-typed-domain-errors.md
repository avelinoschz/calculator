# ADR 0006: Typed Domain Errors with Generic HTTP Mapping

## Status

Accepted

## Context

The original error design used opaque sentinel errors created with `errors.New`:

```go
var ErrDivisionByZero = errors.New("division by zero")
```

These are sentinels in the classic sense: package-level variables compared by
identity via `errors.Is`. They carry no structured data — only a string — so
the handler must know each error individually to produce a meaningful response.

The HTTP handler mapped each error individually:

```go
if errors.Is(calcErr, calculator.ErrDivisionByZero) {
    writeError(w, http.StatusUnprocessableEntity, ErrCodeDivisionByZero, "division by zero is not allowed")
    return
}
```

This created three coupling points for every domain error:

1. The sentinel variable in `internal/calculator`
2. The per-error `if` branch in the handler
3. A duplicated error code constant in `internal/handler/models.go`

Adding a new domain error required touching all three locations and
keeping the human-readable message in the handler in sync with the
domain's intent.

## Decision

Replace opaque sentinel errors with typed sentinels — package-level `*Error`
variables that carry a machine-readable code and a canonical human-readable
message:

```go
type Error struct {
    code    string
    message string
}

func (e *Error) Error() string { return e.message }
func (e *Error) Code() string  { return e.code }

var ErrDivisionByZero = &Error{
    code:    "DIVISION_BY_ZERO",
    message: "division by zero is not allowed",
}
```

The handler maps errors generically via `errors.As` and a static
status-code table:

```go
var errStatusMap = map[string]int{
    "INVALID_OPERATION": http.StatusBadRequest,
    // all others default to 422 Unprocessable Entity
}

func mapCalcError(err error) (int, string, string) {
    var calcErr *calculator.Error
    if errors.As(err, &calcErr) {
        status, ok := errStatusMap[calcErr.Code()]
        if !ok {
            status = http.StatusUnprocessableEntity
        }
        return status, calcErr.Code(), calcErr.Error()
    }
    return http.StatusInternalServerError, ErrCodeInternalError, "an unexpected error occurred"
}
```

## Rationale

### 1. Single authorship per error

The domain owns both the code and the message. There is no duplicated
constant in the handler layer. The handler only decides the HTTP status
code, which is a transport concern.

### 2. O(1) extension cost

Adding a new domain error requires:

- One new `*Error` variable in `internal/calculator`
- One entry in `errStatusMap` if the default 422 is not appropriate

No handler branching logic needs to change.

### 3. Sentinel identity and code-based matching

The errors remain package-level variables — typed sentinels, not opaque
ones. `errors.Is` continues to work for all domain errors via the `Is`
method on `*Error`, which compares by `code` string:

```go
func (e *Error) Is(target error) bool {
    t, ok := target.(*Error)
    if !ok {
        return false
    }
    return e.code == t.code
}
```

For most sentinels (`ErrDivisionByZero`, `ErrNegativeSquareRoot`, etc.) the
domain function returns the sentinel directly, so pointer identity and code
comparison produce the same result.

`ErrOperandOutOfRange` is a special case. Its `message` field is a format
template rather than a final string:

```go
ErrOperandOutOfRange = &Error{
    code:    "OPERAND_OUT_OF_RANGE",
    message: "operands must be between %g and %g",
}
```

`Service.Calculate` never returns the sentinel itself. It always returns a
dynamically constructed instance via `newErrOperandOutOfRange(min, max)`,
which formats the message with the configured limits so the caller sees
`"operands must be between -100 and 100"` rather than a static fallback.
The sentinel is used only as the `errors.Is` comparison target.

`newErrOperandOutOfRange` keeps the sentinel as the single source of truth
for both `code` and message template:

```go
func newErrOperandOutOfRange(min, max float64) *Error {
    return &Error{
        code:    ErrOperandOutOfRange.code,
        message: fmt.Sprintf(ErrOperandOutOfRange.message, min, max),
    }
}
```

Existing test code and any caller relying on `errors.Is` requires no change.
What changed is that the errors also expose structure via `errors.As`,
enabling generic extraction of `Code()` and `Error()` without per-error
branching.

### 4. Message ownership is explicit

Previously, a message like "division by zero is not allowed" lived in
the handler, not in the domain. A developer reading `ErrDivisionByZero`
had no way to know the HTTP message without tracing to the handler. Now
the message is co-located with the error definition.

### 5. Handler becomes transport-only

With generic error mapping, the handler's responsibility shrinks to:
parse HTTP input, call the service, translate domain errors to HTTP
responses. It no longer needs to know the semantics of individual domain
errors.

## Consequences

### Positive

- Adding a domain error no longer requires touching the handler
- Error messages are defined once, in the domain
- Handler tests that assert on error codes use `calculator.ErrXxx.Code()`
  instead of `handler.ErrCodeXxx` constants, making the source of truth
  explicit
- The handler's error code constants are reduced to transport-only codes:
  `INVALID_REQUEST`, `MISSING_FIELD`, `INTERNAL_ERROR`

### Trade-offs

- Domain errors now carry a `Code()` string. If this service were to
  expose a non-HTTP transport (e.g. gRPC), that transport would receive
  the same code strings. This is generally desirable but means the domain
  is slightly more aware of how it is consumed.
- `errors.As` is less widely known than `errors.Is`. The pattern requires
  familiarity with Go's error wrapping model.
- The `Is` method means any two `*Error` values with the same `code` are
  equivalent under `errors.Is`. Pointer uniqueness of sentinels is no
  longer relied upon. This is intentional — it is what allows dynamically
  constructed errors (such as those returned by `newErrOperandOutOfRange`)
  to match their corresponding sentinel.
