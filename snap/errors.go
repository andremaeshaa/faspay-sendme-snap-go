package snap

import (
	"fmt"
	"net/http"
)

// Error represents an error returned by the Faspay SendMe Snap API
type Error struct {
	StatusCode int    // HTTP status code
	Code       string // Error code returned by the API
	Message    string // Error message
	Details    string // Additional error details
}

// Error returns the error message
func (e *Error) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("Faspay API error (HTTP %d, Code: %s): %s - %s", e.StatusCode, e.Code, e.Message, e.Details)
	}
	return fmt.Sprintf("Faspay API error (HTTP %d, Code: %s): %s", e.StatusCode, e.Code, e.Message)
}

// IsAPIError checks if an error is an API error
func IsAPIError(err error) bool {
	_, ok := err.(*Error)
	return ok
}

// IsNotFoundError checks if an error is a not found error
func IsNotFoundError(err error) bool {
	apiErr, ok := err.(*Error)
	return ok && apiErr.StatusCode == http.StatusNotFound
}

// IsAuthenticationError checks if an error is an authentication error
func IsAuthenticationError(err error) bool {
	apiErr, ok := err.(*Error)
	return ok && apiErr.StatusCode == http.StatusUnauthorized
}

// IsValidationError checks if an error is a validation error
func IsValidationError(err error) bool {
	apiErr, ok := err.(*Error)
	return ok && apiErr.StatusCode == http.StatusBadRequest
}

// IsServerError checks if an error is a server error
func IsServerError(err error) bool {
	apiErr, ok := err.(*Error)
	return ok && apiErr.StatusCode >= 500
}

// NewError creates a new API error
func NewError(statusCode int, code, message, details string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
		Details:    details,
	}
}
