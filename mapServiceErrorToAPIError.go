package goErrorHandler

import (
	"net/http"
)

// APIError struct to represent API response errors
type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// MapServiceErrorToAPIError maps service errors to API errors
func MapServiceErrorToAPIError(err error) APIError {
	var apiError APIError

	// Check if the error is of type Error
	if svcError, ok := err.(Error); ok {
		apiError.Message = svcError.AppError().Error()

		// Use a type switch for the service error
		switch err := svcError.SvcError(); err {
		case ErrBadRequest:
			apiError.Status = http.StatusBadRequest
		case ErrNotFound:
			apiError.Status = http.StatusNotFound
		case ErrInternalFailure:
			apiError.Status = http.StatusInternalServerError
		case ErrUnauthorized:
			apiError.Status = http.StatusUnauthorized
		}
	}

	return apiError
}
