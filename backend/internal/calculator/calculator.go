package calculator

import (
	"errors"
	"math"
)

// Operation represents a supported calculator operation.
type Operation string

const (
	OperationAdd        Operation = "add"
	OperationSubtract   Operation = "subtract"
	OperationMultiply   Operation = "multiply"
	OperationDivide     Operation = "divide"
	OperationPower      Operation = "power"
	OperationSqrt       Operation = "sqrt"
	OperationPercentage Operation = "percentage"
)

// SupportedOperations returns the stable, public list of supported operations.
func SupportedOperations() []Operation {
	return []Operation{
		OperationAdd,
		OperationSubtract,
		OperationMultiply,
		OperationDivide,
		OperationPower,
		OperationSqrt,
		OperationPercentage,
	}
}

// RequiresSecondOperand reports whether the operation is binary.
func (op Operation) RequiresSecondOperand() bool {
	switch op {
	case OperationAdd,
		OperationSubtract,
		OperationMultiply,
		OperationDivide,
		OperationPower,
		OperationPercentage:
		return true
	default:
		return false
	}
}

// IsSupported reports whether the operation is recognized.
func (op Operation) IsSupported() bool {
	for _, supported := range SupportedOperations() {
		if op == supported {
			return true
		}
	}
	return false
}

func calculateBinary(op Operation, a, b float64) (float64, error) {
	var result float64

	switch op {
	case OperationAdd:
		result = a + b
	case OperationSubtract:
		result = a - b
	case OperationMultiply:
		result = a * b
	case OperationDivide:
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		result = a / b
	case OperationPower:
		result = math.Pow(a, b)
	case OperationPercentage:
		result = (a / 100) * b
	default:
		return 0, ErrInvalidOperation
	}

	if !isFinite(result) {
		return 0, ErrNonFiniteResult
	}

	return result, nil
}

func calculateUnary(op Operation, a float64) (float64, error) {
	var result float64

	switch op {
	case OperationSqrt:
		if a < 0 {
			return 0, ErrNegativeSquareRoot
		}
		result = math.Sqrt(a)
	default:
		return 0, ErrInvalidOperation
	}

	if !isFinite(result) {
		return 0, ErrNonFiniteResult
	}

	return result, nil
}

func isFinite(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0)
}

// Service exposes the calculator domain through a small concrete type that
// can satisfy consumer-defined interfaces.
type Service struct {
	min float64
	max float64
}

// NewService creates a Service with the given operand limits.
// Use math.Inf(-1) and math.Inf(1) for no limits.
func NewService(min, max float64) (Service, error) {
	if math.IsNaN(min) || math.IsNaN(max) {
		return Service{}, errors.New("limits must not be NaN")
	}
	if min > max {
		return Service{}, errors.New("min must not exceed max")
	}
	return Service{min: min, max: max}, nil
}

// Calculate executes the requested operation using the domain functions.
func (s Service) Calculate(op Operation, a float64, b *float64) (float64, error) {
	if a < s.min || a > s.max {
		return 0, newErrOperandOutOfRange(s.min, s.max)
	}

	if op.RequiresSecondOperand() {
		if b == nil {
			return 0, ErrInvalidOperation
		}
		if *b < s.min || *b > s.max {
			return 0, newErrOperandOutOfRange(s.min, s.max)
		}
		return calculateBinary(op, a, *b)
	}

	return calculateUnary(op, a)
}
