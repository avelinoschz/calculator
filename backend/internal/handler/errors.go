package handler

import (
	"errors"
	"net/http"

	"github.com/avelinoschz/calculator/backend/internal/calculator"
)

// Error code constants for HTTP transport-level errors.
// Domain error codes (INVALID_OPERATION, DIVISION_BY_ZERO, etc.) are owned
// by the calculator package and accessed via (*calculator.Error).Code().
const (
	ErrCodeInvalidRequest = "INVALID_REQUEST"
	ErrCodeMissingField   = "MISSING_FIELD"
	ErrCodeInternalError  = "INTERNAL_ERROR"
	// ErrCodeInvalidOperation is kept here because the handler uses it for
	// pre-validation (op.IsSupported) before calling the service.
	ErrCodeInvalidOperation = "INVALID_OPERATION"
)

// errStatusMap maps domain error codes to HTTP status codes.
// Errors not in this map default to 422 Unprocessable Entity.
var errStatusMap = map[string]int{
	"INVALID_OPERATION": http.StatusBadRequest,
}

func mapCalcError(err error) (status int, code string, message string) {
	var calcErr *calculator.Error
	if errors.As(err, &calcErr) {
		s, ok := errStatusMap[calcErr.Code()]
		if !ok {
			s = http.StatusUnprocessableEntity
		}
		return s, calcErr.Code(), calcErr.Error()
	}

	return http.StatusInternalServerError, ErrCodeInternalError, "an unexpected error occurred"
}
