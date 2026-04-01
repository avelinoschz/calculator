package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"math"
	"net/http"
	"strings"

	"github.com/avelinoschz/calculator/backend/internal/calculator"
)

// CalculatorService defines the calculation behavior the HTTP layer depends on.
// It is intentionally small so handler tests can mock the service boundary.
type CalculatorService interface {
	Calculate(op calculator.Operation, a float64, b *float64) (float64, error)
}

// Handler holds configuration for the calculate endpoint.
type Handler struct {
	// Min and Max define the allowed range for operands.
	// Use math.Inf(-1) and math.Inf(1) for no limits (the defaults).
	Min float64
	Max float64

	Service CalculatorService
}

// New creates a Handler with the default calculator service when one is not supplied.
func New(min, max float64, service CalculatorService) *Handler {
	if service == nil {
		service = calculator.Service{}
	}

	return &Handler{
		Min:     min,
		Max:     max,
		Service: service,
	}
}

func (h *Handler) service() CalculatorService {
	if h.Service == nil {
		h.Service = calculator.Service{}
	}

	return h.Service
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

	op := calculator.Operation(*req.Operation)
	if !op.IsSupported() {
		writeError(w, http.StatusBadRequest, ErrCodeInvalidOperation, invalidOperationMessage())
		return
	}

	if req.A == nil {
		writeMissingField(w, "a")
		return
	}

	a := *req.A

	if a < h.Min || a > h.Max {
		writeError(w, http.StatusUnprocessableEntity, ErrCodeOperandOutOfRange,
			formatOperandRangeMessage(h.Min, h.Max))
		return
	}

	var result float64
	var calcErr error
	var secondOperand *float64

	if op.RequiresSecondOperand() {
		if req.B == nil {
			writeMissingField(w, "b")
			return
		}

		b := *req.B
		if b < h.Min || b > h.Max {
			writeError(w, http.StatusUnprocessableEntity, ErrCodeOperandOutOfRange,
				formatOperandRangeMessage(h.Min, h.Max))
			return
		}
		secondOperand = &b
	} else {
		if req.B != nil {
			writeError(w, http.StatusBadRequest, ErrCodeInvalidRequest, "b is not allowed for sqrt")
			return
		}
	}

	result, calcErr = h.service().Calculate(op, a, secondOperand)

	if calcErr != nil {
		if errors.Is(calcErr, calculator.ErrInvalidOperation) {
			writeError(w, http.StatusBadRequest, ErrCodeInvalidOperation, invalidOperationMessage())
			return
		}
		if errors.Is(calcErr, calculator.ErrDivisionByZero) {
			writeError(w, http.StatusUnprocessableEntity, ErrCodeDivisionByZero, "division by zero is not allowed")
			return
		}
		if errors.Is(calcErr, calculator.ErrNegativeSquareRoot) {
			writeError(w, http.StatusUnprocessableEntity, ErrCodeNegativeSquareRoot, "square root is only defined for non-negative numbers")
			return
		}
		if errors.Is(calcErr, calculator.ErrNonFiniteResult) {
			writeError(w, http.StatusUnprocessableEntity, ErrCodeNonFiniteResult, "calculation result is not a finite real number")
			return
		}
		slog.Error("unexpected calculation error", "error", calcErr)
		writeError(w, http.StatusInternalServerError, ErrCodeInternalError, "an unexpected error occurred")
		return
	}

	logAttrs := []any{"operation", op, "a", a, "result", result}
	if req.B != nil {
		logAttrs = append(logAttrs, "b", *req.B)
	}

	slog.Info("calculation completed", logAttrs...)
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

func invalidOperationMessage() string {
	operations := make([]string, 0, len(calculator.SupportedOperations()))
	for _, op := range calculator.SupportedOperations() {
		operations = append(operations, string(op))
	}

	return fmt.Sprintf("operation must be one of %s", strings.Join(operations, ", "))
}
