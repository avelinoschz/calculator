package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avelinoschz/calculator/backend/internal/handler"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		wantStatus     int
		wantResult     *float64
		wantErrorCode  string
	}{
		{
			name:       "add two numbers",
			body:       `{"operation":"add","a":10,"b":5}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(15.0),
		},
		{
			name:       "subtract two numbers",
			body:       `{"operation":"subtract","a":10,"b":3}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(7.0),
		},
		{
			name:       "multiply two numbers",
			body:       `{"operation":"multiply","a":4,"b":5}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(20.0),
		},
		{
			name:       "divide two numbers",
			body:       `{"operation":"divide","a":20,"b":4}`,
			wantStatus: http.StatusOK,
			wantResult: ptr(5.0),
		},
		{
			name:          "division by zero",
			body:          `{"operation":"divide","a":10,"b":0}`,
			wantStatus:    http.StatusUnprocessableEntity,
			wantErrorCode: handler.ErrCodeDivisionByZero,
		},
		{
			name:          "invalid operation",
			body:          `{"operation":"power","a":2,"b":3}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeInvalidOperation,
		},
		{
			name:          "missing operation field",
			body:          `{"a":1,"b":2}`,
			wantStatus:    http.StatusBadRequest,
			wantErrorCode: handler.ErrCodeMissingField,
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculations", bytes.NewBufferString(tc.body))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			handler.Calculate(rec, req)

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
