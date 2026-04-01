package calculator

import "fmt"

// Error represents a domain calculation error with a machine-readable code.
type Error struct {
	code    string
	message string
}

func (e *Error) Error() string { return e.message }

// Code returns the machine-readable error code.
func (e *Error) Code() string { return e.code }

// Is reports whether this error has the same code as the target.
// This allows errors.Is to match dynamically constructed errors (e.g. those
// with formatted messages) against sentinel errors of the same code.
func (e *Error) Is(target error) bool {
	t, ok := target.(*Error)
	if !ok {
		return false
	}
	return e.code == t.code
}

var (
	// ErrInvalidOperation is returned when the operation is not supported.
	ErrInvalidOperation = &Error{code: "INVALID_OPERATION", message: "invalid operation"}

	// ErrDivisionByZero is returned when dividing by zero.
	ErrDivisionByZero = &Error{code: "DIVISION_BY_ZERO", message: "division by zero is not allowed"}

	// ErrNegativeSquareRoot is returned when taking the square root of a negative number.
	ErrNegativeSquareRoot = &Error{code: "NEGATIVE_SQUARE_ROOT", message: "square root is only defined for non-negative numbers"}

	// ErrNonFiniteResult is returned when an operation produces NaN or Inf.
	ErrNonFiniteResult = &Error{code: "NON_FINITE_RESULT", message: "calculation result is not a finite real number"}

	// ErrOperandOutOfRange is the sentinel used as the comparison target in errors.Is.
	// Its message is a format template; use newErrOperandOutOfRange to produce
	// an instance with the configured limits substituted in.
	ErrOperandOutOfRange = &Error{code: "OPERAND_OUT_OF_RANGE", message: "operands must be between %g and %g"}
)

// newErrOperandOutOfRange returns an OPERAND_OUT_OF_RANGE error whose message
// includes the configured min and max limits.
func newErrOperandOutOfRange(min, max float64) *Error {
	return &Error{
		code:    ErrOperandOutOfRange.code,
		message: fmt.Sprintf(ErrOperandOutOfRange.message, min, max),
	}
}
