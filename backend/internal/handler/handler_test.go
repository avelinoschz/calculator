package handler_test

import (
	"bytes"
	"encoding/json"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avelinoschz/calculator/backend/internal/handler"
)

func noLimitsHandler() *handler.Handler {
	return &handler.Handler{Min: math.Inf(-1), Max: math.Inf(1)}
}

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.Health(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

	var resp map[string]string
	require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
	assert.Equal(t, "ok", resp["status"])
}

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name          string
		body          string
		wantStatus    int
		wantResult    *float64
		wantErrorCode string
	}{
		{
			name:       "add two numbers",
			body:       `{"op":"add","a":10,"b":5}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(15.0),
		},
		{
			name:       "subtract two numbers",
			body:       `{"op":"subtract","a":10,"b":3}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(7.0),
		},
		{
			name:       "multiply two numbers",
			body:       `{"op":"multiply","a":4,"b":5}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(20.0),
		},
		{
			name:       "divide two numbers",
			body:       `{"op":"divide","a":20,"b":4}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(5.0),
		},
		{
			name:       "power two numbers",
			body:       `{"op":"power","a":2,"b":3}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(8.0),
		},
		{
			name:       "percentage of value",
			body:       `{"op":"percentage","a":10,"b":200}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(20.0),
		},
		{
			name:       "square root with one operand",
			body:       `{"op":"sqrt","a":9}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(3.0),
		},
		{
			name:          "division by zero",
			body:          `{"op":"divide","a":10,"b":0}`,
			wantStatus:    http.StatusUnprocessableEntity,
			wantErrorCode: handler.ErrCodeDivisionByZero,
		},
		{
			name:          "invalid operation",
			body:          `{"op":"modulo","a":2,"b":3}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeInvalidOperation,
		},
		{
			name:          "negative square root",
			body:          `{"op":"sqrt","a":-1}`,
			wantStatus:    http.StatusUnprocessableEntity,
			wantErrorCode: handler.ErrCodeNegativeSquareRoot,
		},
		{
			name:          "non finite result",
			body:          `{"op":"power","a":1e308,"b":2}`,
			wantStatus:    http.StatusUnprocessableEntity,
			wantErrorCode: handler.ErrCodeNonFiniteResult,
		},
		{
			name:          "missing operation field",
			body:          `{"a":1,"b":2}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeMissingField,
		},
		{
			name:          "missing a field",
			body:          `{"op":"add","b":2}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeMissingField,
		},
		{
			name:          "missing b field",
			body:          `{"op":"add","a":2}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeMissingField,
		},
		{
			name:          "sqrt rejects second operand",
			body:          `{"op":"sqrt","a":9,"b":1}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeInvalidRequest,
		},
		{
			name:          "malformed JSON",
			body:          `{not valid json`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeInvalidRequest,
		},
		{
			name:          "empty body",
			body:          ``,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeInvalidRequest,
		},
		{
			name:          "unexpected extra field",
			body:          `{"op":"add","a":1,"b":2,"c":3}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeInvalidRequest,
		},
		{
			name:          "trailing JSON payload",
			body:          `{"op":"add","a":1,"b":2}{"op":"subtract","a":2,"b":1}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeInvalidRequest,
		},
	}

	h := noLimitsHandler()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculations", bytes.NewBufferString(tc.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			h.Calculate(rec, req)

			assert.Equal(t, tc.wantStatus, rec.Code)
			assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

			if tc.wantResult != nil {
				var resp handler.CalculateResponse
				require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
				assert.InDelta(t, *tc.wantResult, resp.Result, 1e-9)
			}

			if tc.wantErrorCode != "" {
				var resp handler.ErrorResponse
				require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
				assert.Equal(t, tc.wantErrorCode, resp.Error.Code)
				assert.NotEmpty(t, resp.Error.Message)
			}
		})
	}
}

func TestCalculateHandlerOperandLimits(t *testing.T) {
	h := &handler.Handler{Min: -100, Max: 100}

	tests := []struct {
		name          string
		body          string
		wantStatus    int
		wantResult    *float64
		wantErrorCode string
	}{
		{
			name:       "both operands within range",
			body:       `{"op":"add","a":50,"b":50}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(100.0),
		},
		{
			name:       "operands at exact boundaries",
			body:       `{"op":"add","a":-100,"b":100}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(0.0),
		},
		{
			name:       "sqrt operand within range",
			body:       `{"op":"sqrt","a":100}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(10.0),
		},
		{
			name:          "a below min",
			body:          `{"op":"add","a":-101,"b":0}`,
			wantStatus:    http.StatusUnprocessableEntity,
			wantErrorCode: handler.ErrCodeOperandOutOfRange,
		},
		{
			name:          "sqrt below min",
			body:          `{"op":"sqrt","a":-101}`,
			wantStatus:    http.StatusUnprocessableEntity,
			wantErrorCode: handler.ErrCodeOperandOutOfRange,
		},
		{
			name:          "b above max",
			body:          `{"op":"add","a":0,"b":101}`,
			wantStatus:    http.StatusUnprocessableEntity,
			wantErrorCode: handler.ErrCodeOperandOutOfRange,
		},
		{
			name:          "both operands out of range",
			body:          `{"op":"add","a":-999,"b":999}`,
			wantStatus:    http.StatusUnprocessableEntity,
			wantErrorCode: handler.ErrCodeOperandOutOfRange,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculations", bytes.NewBufferString(tc.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			h.Calculate(rec, req)

			assert.Equal(t, tc.wantStatus, rec.Code)
			assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))

			if tc.wantResult != nil {
				var resp handler.CalculateResponse
				require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
				assert.InDelta(t, *tc.wantResult, resp.Result, 1e-9)
			}

			if tc.wantErrorCode != "" {
				var resp handler.ErrorResponse
				require.NoError(t, json.NewDecoder(rec.Body).Decode(&resp))
				assert.Equal(t, tc.wantErrorCode, resp.Error.Code)
				assert.NotEmpty(t, resp.Error.Message)
			}
		})
	}
}

func ptr(f float64) *float64 { return &f }
