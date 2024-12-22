package calculation

import "errors"

var (
	ErrDivisionByZero        = errors.New("division by zero")//
	ErrEmptyInput            = errors.New("empty input") // добавить
	ErrMismatchedParentheses = errors.New("mismatched parentheses") //
	ErrInvalidOperator       = errors.New("invalid operator")//
	ErrOperatorAtEnd         = errors.New("operator at end of expression")
	ErrNotEnoughOperands     = errors.New("not enough operands")
	ErrExtraOperands         = errors.New("extra operands")
	ErrInvalidCharacter      = errors.New("invalid character")
)