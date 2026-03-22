package handler

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/avelinoschz/calculator/backend/internal/calculator"
)

// Health handles GET /health.
func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Calculate handles POST /api/v1/calculations.
func Calculate(w http.ResponseWriter, r *http.Request) {
	var req CalculateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Warn("failed to decode request body", "error", err)
		writeError(w, http.StatusBadRequest, ErrCodeInvalidRequest, "request body is invalid")
		return
	}

	if req.Operation == "" {
		writeError(w, http.StatusBadRequest, ErrCodeMissingField, "op is required")
		return
	}

	result, err := calculator.Calculate(calculator.Operation(req.Operation), req.A, req.B)
	if err != nil {
		if errors.Is(err, calculator.ErrInvalidOperation) {
			writeError(w, http.StatusBadRequest, ErrCodeInvalidOperation, "operation must be one of add, subtract, multiply, divide")
			return
		}
		if errors.Is(err, calculator.ErrDivisionByZero) {
			writeError(w, http.StatusUnprocessableEntity, ErrCodeDivisionByZero, "division by zero is not allowed")
			return
		}
		slog.Error("unexpected calculation error", "error", err)
		writeError(w, http.StatusInternalServerError, ErrCodeInternalError, "an unexpected error occurred")
		return
	}

	slog.Info("calculation completed", "operation", req.Operation, "a", req.A, "b", req.B, "result", result)
	writeJSON(w, http.StatusOK, CalculateResponse{Result: result})
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, ErrorResponse{
		Error: ErrorDetail{Code: code, Message: message},
	})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		slog.Error("failed to write response", "error", err)
	}
}
