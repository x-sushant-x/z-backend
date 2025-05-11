package customErrors

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
)
