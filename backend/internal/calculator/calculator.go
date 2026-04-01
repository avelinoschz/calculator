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

// ErrInvalidOperation is returned when the operation is not supported.
var ErrInvalidOperation = errors.New("invalid operation")

// ErrDivisionByZero is returned when dividing by zero.
var ErrDivisionByZero = errors.New("division by zero")

// ErrNegativeSquareRoot is returned when taking the square root of a negative number.
var ErrNegativeSquareRoot = errors.New("negative square root")

// ErrNonFiniteResult is returned when an operation produces NaN or Inf.
var ErrNonFiniteResult = errors.New("non-finite result")

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

// CalculateBinary performs the given binary operation on operands a and b.
func CalculateBinary(op Operation, a, b float64) (float64, error) {
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

// CalculateUnary performs the given unary operation on operand a.
func CalculateUnary(op Operation, a float64) (float64, error) {
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
type Service struct{}

// Calculate executes the requested operation using the domain functions.
func (Service) Calculate(op Operation, a float64, b *float64) (float64, error) {
	if op.RequiresSecondOperand() {
		if b == nil {
			return 0, ErrInvalidOperation
		}
		return CalculateBinary(op, a, *b)
	}

	return CalculateUnary(op, a)
}
