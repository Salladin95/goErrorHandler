package goErrorHandler

import (
	"fmt"
)

func BindRequestToBodyFailure(err error) error {
	return NewError(ErrBadRequest, fmt.Errorf("failed to bind request body: %s", err.Error()))
}

func ParseUUIDFailure() error {
	return NewError(ErrBadRequest, fmt.Errorf("invalid uuid"))
}

func Unauthorized() error {
	return NewError(ErrUnauthorized, fmt.Errorf("unauthorized"))
}

func IncorrectLoginOrPassword() error {
	return NewError(ErrBadRequest, fmt.Errorf("login or password is incorrect"))
}

func ValidationFailure(messages map[string]string) error {
	var errorMsg string

	for key, value := range messages {
		errorMsg += fmt.Sprintf("%s: %s\n", key, value)
	}
	return NewError(ErrBadRequest, fmt.Errorf(errorMsg))
}

func OperationFailure(operation string, err error) error {
	return NewError(ErrInternalFailure, fmt.Errorf("failed to %s: %s", operation, err.Error()))
}
