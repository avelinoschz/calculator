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

Supported operations:

- `add`
- `subtract`
- `multiply`
- `divide`

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
  missing field, or invalid operation
- `422 Unprocessable Entity` — division by zero or operand outside
  configured limits
- `500 Internal Server Error` — unexpected server-side failure

## Error Codes

- `INVALID_REQUEST` — malformed JSON, unknown field, or trailing payload
- `MISSING_FIELD` — `op`, `a`, or `b` is absent
- `INVALID_OPERATION` — unsupported `op`
- `DIVISION_BY_ZERO` — `op=divide` and `b=0`
- `OPERAND_OUT_OF_RANGE` — `a` or `b` is outside configured backend
  limits
- `INTERNAL_ERROR` — unexpected server-side failure

## Validation Rules

- request body must be valid JSON
- request body must contain only `op`, `a`, and `b`
- request body must contain exactly one JSON object
- `op`, `a`, and `b` are required
- `op` must be one of the four supported operations
- backend operand limits are enforced after request validation

## Notes

- The backend is the source of truth for validation.
- The frontend mirrors some of this validation for UX, but backend
  responses define the final contract.
