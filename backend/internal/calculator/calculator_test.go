package calculator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avelinoschz/calculator/backend/internal/calculator"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		op      calculator.Operation
		a, b    float64
		want    float64
		wantErr error
	}{
		// Addition
		{name: "add positive numbers", op: calculator.OperationAdd, a: 10, b: 5, want: 15},
		{name: "add negative numbers", op: calculator.OperationAdd, a: -3, b: -7, want: -10},
		{name: "add with zero", op: calculator.OperationAdd, a: 0, b: 5, want: 5},
		{name: "add decimals", op: calculator.OperationAdd, a: 1.5, b: 2.5, want: 4},

		// Subtraction
		{name: "subtract positive numbers", op: calculator.OperationSubtract, a: 10, b: 3, want: 7},
		{name: "subtract resulting in negative", op: calculator.OperationSubtract, a: 3, b: 10, want: -7},
		{name: "subtract with zero", op: calculator.OperationSubtract, a: 5, b: 0, want: 5},

		// Multiplication
		{name: "multiply positive numbers", op: calculator.OperationMultiply, a: 4, b: 5, want: 20},
		{name: "multiply by zero", op: calculator.OperationMultiply, a: 100, b: 0, want: 0},
		{name: "multiply negative numbers", op: calculator.OperationMultiply, a: -3, b: -4, want: 12},
		{name: "multiply positive by negative", op: calculator.OperationMultiply, a: 3, b: -4, want: -12},
		{name: "multiply decimals", op: calculator.OperationMultiply, a: 0.5, b: 4, want: 2},

		// Division
		{name: "divide positive numbers", op: calculator.OperationDivide, a: 20, b: 4, want: 5},
		{name: "divide resulting in decimal", op: calculator.OperationDivide, a: 10, b: 3, want: 10.0 / 3.0},
		{name: "divide negative by positive", op: calculator.OperationDivide, a: -12, b: 4, want: -3},

		// Power
		{name: "power positive exponent", op: calculator.OperationPower, a: 2, b: 3, want: 8},
		{name: "power negative exponent", op: calculator.OperationPower, a: 2, b: -1, want: 0.5},

		// Percentage
		{name: "percentage of value", op: calculator.OperationPercentage, a: 10, b: 200, want: 20},
		{name: "percentage with decimal", op: calculator.OperationPercentage, a: 12.5, b: 80, want: 10},

		// Errors
		{name: "divide by zero", op: calculator.OperationDivide, a: 10, b: 0, wantErr: calculator.ErrDivisionByZero},
		{name: "power overflow", op: calculator.OperationPower, a: 1e308, b: 2, wantErr: calculator.ErrNonFiniteResult},
		{name: "empty operation", op: calculator.Operation(""), a: 1, b: 1, wantErr: calculator.ErrInvalidOperation},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.CalculateBinary(tc.op, tc.a, tc.b)
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			assert.InDelta(t, tc.want, got, 1e-9)
		})
	}
}

func TestCalculateUnary(t *testing.T) {
	tests := []struct {
		name    string
		op      calculator.Operation
		a       float64
		want    float64
		wantErr error
	}{
		{name: "square root of perfect square", op: calculator.OperationSqrt, a: 9, want: 3},
		{name: "square root of zero", op: calculator.OperationSqrt, a: 0, want: 0},
		{name: "square root of decimal", op: calculator.OperationSqrt, a: 2.25, want: 1.5},
		{name: "negative square root", op: calculator.OperationSqrt, a: -1, wantErr: calculator.ErrNegativeSquareRoot},
		{name: "invalid unary operation", op: calculator.OperationAdd, a: 1, wantErr: calculator.ErrInvalidOperation},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.CalculateUnary(tc.op, tc.a)
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			assert.InDelta(t, tc.want, got, 1e-9)
		})
	}
}
