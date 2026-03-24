package handler

// CalculateRequest is the expected request body for POST /api/v1/calculations.
type CalculateRequest struct {
	Operation *string  `json:"op"`
	A         *float64 `json:"a"`
	B         *float64 `json:"b"`
}

// CalculateResponse is the success response body.
type CalculateResponse struct {
	Result float64 `json:"result"`
}

// ErrorDetail contains a machine-readable code and a human-readable message.
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ErrorResponse is the consistent error response body.
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// Error code constants matching the API contract.
const (
	ErrCodeInvalidRequest     = "INVALID_REQUEST"
	ErrCodeInvalidOperation   = "INVALID_OPERATION"
	ErrCodeMissingField       = "MISSING_FIELD"
	ErrCodeDivisionByZero     = "DIVISION_BY_ZERO"
	ErrCodeNegativeSquareRoot = "NEGATIVE_SQUARE_ROOT"
	ErrCodeNonFiniteResult    = "NON_FINITE_RESULT"
	ErrCodeOperandOutOfRange  = "OPERAND_OUT_OF_RANGE"
	ErrCodeInternalError      = "INTERNAL_ERROR"
)
