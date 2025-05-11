package customErrors

import "errors"

var (
	ErrUnableToFetchTask = errors.New("unable to fetch tasks")
)
