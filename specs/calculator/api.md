# API Contract

Human-readable guide for the calculator API. The canonical contract is [`api/openapi.yaml`](../../api/openapi.yaml).

## Endpoints

### `GET /health`

Response:

```json
{
  "status": "ok"
}
```

### `POST /api/v1/calculations`

Request:

```json
{
  "op": "add",
  "a": 10,
  "b": 5
}
```

Unary request:

```json
{
  "op": "sqrt",
  "a": 9
}
```

Supported operations:

- `add`
- `subtract`
- `multiply`
- `divide`
- `power`
- `sqrt`
- `percentage`

Success:

```json
{
  "result": 15
}
```

Error:

```json
{
  "error": {
    "code": "INVALID_REQUEST",
    "message": "request body is invalid"
  }
}
```

## Status Codes

- `200 OK` — successful calculation
- `400 Bad Request` — malformed JSON, extra fields, trailing payload,
  missing field, invalid operation, or invalid unary/binary field shape
- `422 Unprocessable Entity` — division by zero, negative square root,
  non-finite result, or operand outside configured limits
- `500 Internal Server Error` — unexpected server-side failure

## Error Codes

- `INVALID_REQUEST` — malformed JSON, unknown field, or trailing payload
- `INVALID_REQUEST` — malformed JSON, unknown field, trailing payload,
  or disallowed field for the selected operation
- `MISSING_FIELD` — `op`, `a`, or `b` is absent
- `INVALID_OPERATION` — unsupported `op`
- `DIVISION_BY_ZERO` — `op=divide` and `b=0`
- `NEGATIVE_SQUARE_ROOT` — `op=sqrt` and `a<0`
- `NON_FINITE_RESULT` — the calculation would return `NaN` or `Inf`
- `OPERAND_OUT_OF_RANGE` — `a` or `b` is outside configured backend
  limits
- `INTERNAL_ERROR` — unexpected server-side failure

## Validation Rules

- request body must be valid JSON
- request body must contain only `op`, `a`, and `b`
- request body must contain exactly one JSON object
- `op` and `a` are always required
- `b` is required for `add`, `subtract`, `multiply`, `divide`, `power`,
  and `percentage`
- `b` is not allowed for `sqrt`
- `op` must be one of the seven supported operations
- backend operand limits are enforced after request validation

## Notes

- The backend is the source of truth for validation.
- The frontend mirrors some of this validation for UX, but backend
  responses define the final contract.
- Percentage is defined as `a% of b`, computed as `(a / 100) * b`.
