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
		switch svcError.SvcError() {
		case ErrBadRequest:
			apiError.Status = http.StatusBadRequest
		case ErrNotFound:
			apiError.Status = http.StatusNotFound
		case ErrInternalFailure:
			apiError.Status = http.StatusInternalServerError
		case ErrUnauthorized:
			apiError.Status = http.StatusUnauthorized
		}
	} else {
		// If not, treat it as a generic internal server error
		apiError.Message = err.Error()
		apiError.Status = http.StatusInternalServerError
	}

	return apiError
}
