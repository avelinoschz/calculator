package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/avelinoschz/calculator/backend/internal/calculator"
)

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

// CalculatorService defines the calculation behavior the HTTP layer depends on.
// It is intentionally small so handler tests can mock the service boundary.
type CalculatorService interface {
	Calculate(op calculator.Operation, a float64, b *float64) (float64, error)
}

// Handler holds configuration for the calculate endpoint.
type Handler struct {
	service CalculatorService
}

// New creates a Handler with the given calculator service.
func New(service CalculatorService) *Handler {
	return &Handler{service: service}
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

	var secondOperand *float64

	if op.RequiresSecondOperand() {
		if req.B == nil {
			writeMissingField(w, "b")
			return
		}

		b := *req.B
		secondOperand = &b
	} else {
		if req.B != nil {
			writeError(w, http.StatusBadRequest, ErrCodeInvalidRequest,
				fmt.Sprintf("b is not allowed for %s", op))
			return
		}
	}

	result, calcErr := h.service.Calculate(op, a, secondOperand)
	if calcErr != nil {
		status, code, msg := mapCalcError(calcErr)
		if status == http.StatusInternalServerError {
			slog.Error("unexpected calculation error", "error", calcErr)
		}
		writeError(w, status, code, msg)
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
