package business

import "errors"

var (
	// Error when data is not found
	ErrNotFound = errors.New("data was not found")
	// Error when there is duplicate data or column
	ErrConflict = errors.New("conflict")
)
