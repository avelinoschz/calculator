# API Contract

## Overview

The backend exposes a minimal REST API for calculator operations.

The contract is intentionally small:

- one endpoint
- one request shape
- one success shape
- one error shape

This keeps the API easy to implement, test, and consume from the
frontend.

This document is a human-readable guide to the API. The canonical
machine-readable contract lives in `api/openapi.yaml`. If there is any
discrepancy, `api/openapi.yaml` prevails.

---

## Endpoint

### `POST /api/v1/calculations`

Executes a calculator operation using two numeric operands.

---

## Request

### Content-Type

`application/json`

### Request Body

```json
{
  "operation": "add",
  "a": 10,
  "b": 5
}
```

### Request Fields

| Field       | Type     | Required | Description          |
| ----------- | -------- | -------- | -------------------- |
| `operation` | `string` | Yes      | Operation to execute |
| `a`         | `number` | Yes      | First operand        |
| `b`         | `number` | Yes      | Second operand       |

### Allowed Operations

- `add`
- `subtract`
- `multiply`
- `divide`

---

## Success Response

### Success Status

`200 OK`

### Success Body

```json
{
  "result": 15
}
```

### Success Fields

| Field    | Type     | Description     |
| -------- | -------- | --------------- |
| `result` | `number` | Computed result |

---

## Error Response

All API errors return a consistent JSON structure.

### Error Body

```json
{
  "error": {
    "code": "INVALID_REQUEST",
    "message": "request body is invalid"
  }
}
```

### Error Fields

- `error.code` (`string`): machine-readable error code
- `error.message` (`string`): human-readable error message

---

## Status Codes

- `200 OK`: successful calculation
- `400 Bad Request`: invalid payload, malformed JSON, missing fields,
  or invalid operation
- `422 Unprocessable Entity`: mathematically invalid request, such as
  division by zero
- `500 Internal Server Error`: unexpected server-side error

---

## Error Codes

The set of error codes should remain small and predictable. Refer to
`api/openapi.yaml` for the canonical response contract.

- `INVALID_REQUEST`: request body could not be parsed or failed basic
  validation
- `INVALID_OPERATION`: operation is not supported
- `MISSING_FIELD`: a required field is missing
- `INVALID_NUMBER`: one or more operands are invalid
- `DIVISION_BY_ZERO`: division by zero is not allowed
- `INTERNAL_ERROR`: unexpected internal server error

---

## Validation Rules

### Request Validation

- Request body must be valid JSON
- `operation` must be one of:
  - `add`
  - `subtract`
  - `multiply`
  - `divide`
- `a` must be a valid number
- `b` must be a valid number

### Business Validation

- `divide` must reject `b = 0`

---

## Examples

### Addition

#### Addition Request

```json
{
  "operation": "add",
  "a": 10,
  "b": 5
}
```

#### Addition Response

```json
{
  "result": 15
}
```

---

### Division

#### Division Request

```json
{
  "operation": "divide",
  "a": 20,
  "b": 4
}
```

#### Division Response

```json
{
  "result": 5
}
```

---

### Invalid Operation

#### Invalid Operation Request

```json
{
  "operation": "power",
  "a": 2,
  "b": 3
}
```

#### Invalid Operation Response

```json
{
  "error": {
    "code": "INVALID_OPERATION",
    "message": "operation must be one of add, subtract, multiply, divide"
  }
}
```

---

### Division by Zero

#### Division by Zero Request

```json
{
  "operation": "divide",
  "a": 10,
  "b": 0
}
```

#### Division by Zero Response

```json
{
  "error": {
    "code": "DIVISION_BY_ZERO",
    "message": "division by zero is not allowed"
  }
}
```

---

## Contract Design Notes

### Why a single endpoint?

A single calculations endpoint keeps the API small and avoids
unnecessary duplication across multiple operation-specific routes.

### Why a consistent error shape?

A stable error model simplifies frontend integration and improves
maintainability.

### Why binary operands only?

The required scope only includes binary operations, which keeps the
contract simple and focused.
