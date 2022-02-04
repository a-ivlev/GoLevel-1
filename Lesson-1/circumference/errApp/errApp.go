package errApp

import "errors"


var (
	ErrAreaNegativNumber = errors.New("area cannot be a negative number")
	ErrAreaZero = errors.New("area cannot be zero")
)