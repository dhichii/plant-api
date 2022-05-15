package business

import "errors"

var (
	// Error when there is bad request
	ErrBadRequest = errors.New("bad request")
	// Error when data is not found
	ErrNotFound = errors.New("data was not found")
	// Error when there is duplicate data or column
	ErrConflict = errors.New("conflict")
)
