package errors

import "errors"

var (
	ErrNotFound = errors.New("Requested item is not found!")
)
