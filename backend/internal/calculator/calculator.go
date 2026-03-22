package calculator

import "errors"

// Operation represents a supported calculator operation.
type Operation string

const (
	OperationAdd      Operation = "add"
	OperationSubtract Operation = "subtract"
	OperationMultiply Operation = "multiply"
	OperationDivide   Operation = "divide"
)

// ErrInvalidOperation is returned when the operation is not supported.
var ErrInvalidOperation = errors.New("invalid operation")

// ErrDivisionByZero is returned when dividing by zero.
var ErrDivisionByZero = errors.New("division by zero")

// Calculate performs the given operation on operands a and b.
func Calculate(op Operation, a, b float64) (float64, error) {
	switch op {
	case OperationAdd:
		return a + b, nil
	case OperationSubtract:
		return a - b, nil
	case OperationMultiply:
		return a * b, nil
	case OperationDivide:
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, ErrInvalidOperation
	}
}
