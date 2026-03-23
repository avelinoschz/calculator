package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"net/http"

	"github.com/avelinoschz/calculator/backend/internal/calculator"
)

// Handler holds configuration for the calculate endpoint.
type Handler struct {
	// Min and Max define the allowed range for operands.
	// Use math.Inf(-1) and math.Inf(1) for no limits (the defaults).
	Min float64
	Max float64
}

// Health handles GET /health.
func Health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Calculate handles POST /api/v1/calculations.
func (h *Handler) Calculate(w http.ResponseWriter, r *http.Request) {
	req, err := decodeCalculateRequest(r)
	if err != nil {
		slog.Warn("failed to decode request body", "error", err)
		writeError(w, http.StatusBadRequest, ErrCodeInvalidRequest, "request body is invalid")
		return
	}

	if req.Operation == nil {
		writeMissingField(w, "op")
		return
	}

	if req.A == nil {
		writeMissingField(w, "a")
		return
	}

	if req.B == nil {
		writeMissingField(w, "b")
		return
	}

	op := *req.Operation
	a := *req.A
	b := *req.B

	if a < h.Min || a > h.Max || b < h.Min || b > h.Max {
		writeError(w, http.StatusUnprocessableEntity, ErrCodeOperandOutOfRange,
			formatOperandRangeMessage(h.Min, h.Max))
		return
	}

	result, calcErr := calculator.Calculate(calculator.Operation(op), a, b)
	if calcErr != nil {
		if errors.Is(calcErr, calculator.ErrInvalidOperation) {
			writeError(w, http.StatusBadRequest, ErrCodeInvalidOperation, "operation must be one of add, subtract, multiply, divide")
			return
		}
		if errors.Is(calcErr, calculator.ErrDivisionByZero) {
			writeError(w, http.StatusUnprocessableEntity, ErrCodeDivisionByZero, "division by zero is not allowed")
			return
		}
		slog.Error("unexpected calculation error", "error", calcErr)
		writeError(w, http.StatusInternalServerError, ErrCodeInternalError, "an unexpected error occurred")
		return
	}

	slog.Info("calculation completed", "operation", op, "a", a, "b", b, "result", result)
	writeJSON(w, http.StatusOK, CalculateResponse{Result: result})
}

func decodeCalculateRequest(r *http.Request) (CalculateRequest, error) {
	var req CalculateRequest

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&req); err != nil {
		return CalculateRequest{}, err
	}

	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return CalculateRequest{}, errors.New("unexpected trailing payload")
		}
		return CalculateRequest{}, err
	}

	return req, nil
}

func formatOperandRangeMessage(min, max float64) string {
	if !math.IsInf(min, -1) && !math.IsInf(max, 1) {
		return fmt.Sprintf("operands must be between %g and %g", min, max)
	}

	if !math.IsInf(min, -1) {
		return fmt.Sprintf("operands must be at least %g", min)
	}

	if !math.IsInf(max, 1) {
		return fmt.Sprintf("operands must be at most %g", max)
	}

	return "operands are outside the allowed range"
}

func writeMissingField(w http.ResponseWriter, field string) {
	writeError(w, http.StatusBadRequest, ErrCodeMissingField, fmt.Sprintf("%s is required", field))
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
