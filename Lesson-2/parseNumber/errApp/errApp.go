package errApp

import "errors"

var (
	ErrNumberLess100 = errors.New("the number cannot be less than 100")
	ErrNumberGreater999 = errors.New("the number cannot be greater than 999")
)