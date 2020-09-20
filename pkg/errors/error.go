package errors

import "net/http"

// Source:
//https://github.com/Optum/dce/blob/82521f9b906194df4b69ea8e852d9b3f763e4c89/pkg/errors/error.go

// StatusError is the custom error type we are using.
// Should satisfy errors interface
type StatusError struct {
	httpCode int
	cause    error
	message  string
}

// Error allows conversion to standard error object
func (se *StatusError) Error() string {
	return se.message
}

// HTTPCode returns the http code
func (se StatusError) HTTPCode() int { return se.httpCode }

// HTTPCode returns the API Code
type HTTPCode interface {
	HTTPCode() int
}

// NewBadRequest returns a new error representing a bad request
func NewBadRequest(m string) *StatusError {
	return &StatusError{
		httpCode: http.StatusBadRequest,
		cause:    nil,
		message:  m,
	}
}

// NewInternalServer returns an error for Internal Server Errors
func NewInternalServer(m string, err error) *StatusError {
	return &StatusError{
		httpCode: http.StatusInternalServerError,
		cause:    err,
		message:  m,
	}
}
