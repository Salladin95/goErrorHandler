package goErrorHandler

import "errors"

// Service-specific errors
var (
	ErrNotFound        = errors.New("not found")
	ErrBadRequest      = errors.New("bad request")
	ErrInternalFailure = errors.New("internal failure")
	ErrUnauthorized    = errors.New("unauthorized")
)

// Error struct to represent both application and service errors
type Error struct {
	appErr error
	svcErr error
}

// AppError returns the application error
func (e Error) AppError() error {
	return e.appErr
}

// SvcError returns the service error
func (e Error) SvcError() error {
	return e.svcErr
}

// NewError creates a new instance of Error
func NewError(svcErr, appErr error) error {
	return Error{svcErr: svcErr, appErr: appErr}
}

// Error returns a string representation of the error
func (e Error) Error() string {
	return errors.Join(e.svcErr, e.appErr).Error()
}
