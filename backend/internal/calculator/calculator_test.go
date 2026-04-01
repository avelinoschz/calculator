package calculator_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/avelinoschz/calculator/backend/internal/calculator"
)

func noLimitsService(t *testing.T) calculator.Service {
	t.Helper()
	svc, err := calculator.NewService(math.Inf(-1), math.Inf(1))
	require.NoError(t, err)
	return svc
}

func TestNewService(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		min     float64
		max     float64
		wantErr bool
	}{
		{name: "valid no limits", min: math.Inf(-1), max: math.Inf(1)},
		{name: "valid finite limits", min: -100, max: 100},
		{name: "valid equal limits", min: 0, max: 0},
		{name: "min greater than max", min: 100, max: -100, wantErr: true},
		{name: "NaN min", min: math.NaN(), max: 100, wantErr: true},
		{name: "NaN max", min: -100, max: math.NaN(), wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := calculator.NewService(tc.min, tc.max)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestServiceCalculate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		op      calculator.Operation
		a       float64
		b       *float64
		want    float64
		wantErr error
	}{
		// Addition
		{name: "add positive numbers", op: calculator.OperationAdd, a: 10, b: ptr(5), want: 15},
		{name: "add negative numbers", op: calculator.OperationAdd, a: -3, b: ptr(-7), want: -10},
		{name: "add with zero", op: calculator.OperationAdd, a: 0, b: ptr(5), want: 5},
		{name: "add decimals", op: calculator.OperationAdd, a: 1.5, b: ptr(2.5), want: 4},

		// Subtraction
		{name: "subtract positive numbers", op: calculator.OperationSubtract, a: 10, b: ptr(3), want: 7},
		{name: "subtract resulting in negative", op: calculator.OperationSubtract, a: 3, b: ptr(10), want: -7},
		{name: "subtract with zero", op: calculator.OperationSubtract, a: 5, b: ptr(0), want: 5},

		// Multiplication
		{name: "multiply positive numbers", op: calculator.OperationMultiply, a: 4, b: ptr(5), want: 20},
		{name: "multiply by zero", op: calculator.OperationMultiply, a: 100, b: ptr(0), want: 0},
		{name: "multiply negative numbers", op: calculator.OperationMultiply, a: -3, b: ptr(-4), want: 12},
		{name: "multiply positive by negative", op: calculator.OperationMultiply, a: 3, b: ptr(-4), want: -12},
		{name: "multiply decimals", op: calculator.OperationMultiply, a: 0.5, b: ptr(4), want: 2},

		// Division
		{name: "divide positive numbers", op: calculator.OperationDivide, a: 20, b: ptr(4), want: 5},
		{name: "divide resulting in decimal", op: calculator.OperationDivide, a: 10, b: ptr(3), want: 10.0 / 3.0},
		{name: "divide negative by positive", op: calculator.OperationDivide, a: -12, b: ptr(4), want: -3},

		// Power
		{name: "power positive exponent", op: calculator.OperationPower, a: 2, b: ptr(3), want: 8},
		{name: "power negative exponent", op: calculator.OperationPower, a: 2, b: ptr(-1), want: 0.5},

		// Percentage
		{name: "percentage of value", op: calculator.OperationPercentage, a: 10, b: ptr(200), want: 20},
		{name: "percentage with decimal", op: calculator.OperationPercentage, a: 12.5, b: ptr(80), want: 10},

		// Sqrt (unary)
		{name: "sqrt of perfect square", op: calculator.OperationSqrt, a: 9, want: 3},
		{name: "sqrt of zero", op: calculator.OperationSqrt, a: 0, want: 0},
		{name: "sqrt of decimal", op: calculator.OperationSqrt, a: 2.25, want: 1.5},

		// Errors
		{name: "divide by zero", op: calculator.OperationDivide, a: 10, b: ptr(0), wantErr: calculator.ErrDivisionByZero},
		{name: "power overflow", op: calculator.OperationPower, a: 1e308, b: ptr(2), wantErr: calculator.ErrNonFiniteResult},
		{name: "negative square root", op: calculator.OperationSqrt, a: -1, wantErr: calculator.ErrNegativeSquareRoot},
		{name: "unsupported operation", op: calculator.Operation(""), a: 1, b: ptr(1), wantErr: calculator.ErrInvalidOperation},

		// Service dispatch: binary op with nil b
		{name: "binary op with nil b", op: calculator.OperationAdd, a: 1, wantErr: calculator.ErrInvalidOperation},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			svc := noLimitsService(t)
			got, err := svc.Calculate(tc.op, tc.a, tc.b)
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
			assert.InDelta(t, tc.want, got, 1e-9)
		})
	}
}

func TestServiceCalculateOperandLimits(t *testing.T) {
	t.Parallel()

	svc, err := calculator.NewService(-100, 100)
	require.NoError(t, err)

	tests := []struct {
		name    string
		op      calculator.Operation
		a       float64
		b       *float64
		wantErr error
	}{
		{name: "a within range", op: calculator.OperationAdd, a: 50, b: ptr(50)},
		{name: "a at min boundary", op: calculator.OperationAdd, a: -100, b: ptr(0)},
		{name: "a at max boundary", op: calculator.OperationAdd, a: 100, b: ptr(0)},
		{name: "a below min", op: calculator.OperationAdd, a: -101, b: ptr(0), wantErr: calculator.ErrOperandOutOfRange},
		{name: "a above max", op: calculator.OperationAdd, a: 101, b: ptr(0), wantErr: calculator.ErrOperandOutOfRange},
		{name: "b below min", op: calculator.OperationAdd, a: 0, b: ptr(-101), wantErr: calculator.ErrOperandOutOfRange},
		{name: "b above max", op: calculator.OperationAdd, a: 0, b: ptr(101), wantErr: calculator.ErrOperandOutOfRange},
		{name: "sqrt a within range", op: calculator.OperationSqrt, a: 9},
		{name: "sqrt a out of range", op: calculator.OperationSqrt, a: -101, wantErr: calculator.ErrOperandOutOfRange},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, err := svc.Calculate(tc.op, tc.a, tc.b)
			if tc.wantErr != nil {
				require.ErrorIs(t, err, tc.wantErr)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestServiceCalculateOperandOutOfRangeMessage(t *testing.T) {
	t.Parallel()

	svc, err := calculator.NewService(-100, 100)
	require.NoError(t, err)

	_, err = svc.Calculate(calculator.OperationAdd, -101, ptr(0))
	require.ErrorIs(t, err, calculator.ErrOperandOutOfRange)
	assert.Equal(t, "operands must be between -100 and 100", err.Error())
}

func ptr(f float64) *float64 { return &f }
