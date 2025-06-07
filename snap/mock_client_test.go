package snap

import (
	"bytes"
	"io"
	"net/http"
)

// MockTransport is a mock implementation of http.RoundTripper for testing
type MockTransport struct {
	// RoundTripFunc allows us to set a custom function for the RoundTrip method
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

// RoundTrip is the mock implementation of the http.RoundTripper's RoundTrip method
func (m *MockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.RoundTripFunc != nil {
		return m.RoundTripFunc(req)
	}
	// Default implementation returns a 200 OK response with empty body
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString("")),
	}, nil
}

// NewMockClient creates a new http.Client with a mock transport
func NewMockClient(roundTripFunc func(req *http.Request) (*http.Response, error)) *http.Client {
	return &http.Client{
		Transport: &MockTransport{
			RoundTripFunc: roundTripFunc,
		},
	}
}

// MockResponse creates a mock HTTP response with the given status code and body
func MockResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}
}

// MockSuccessResponse creates a mock successful HTTP response for account inquiry
func MockSuccessResponse() *http.Response {
	body := `{
		"responseCode": "00",
		"responseMessage": "Success",
		"referenceNo": "REF123456789",
		"partnerReferenceNo": "20250606234037372",
		"beneficiaryAccountName": "JOHN DOE",
		"beneficiaryAccountNo": "60004400184",
		"beneficiaryBankCode": "008",
		"beneficiaryBankName": "MANDIRI",
		"currency": "IDR",
		"additionalInfo": {
			"status": "success",
			"message": "Account inquiry successful"
		}
	}`
	return MockResponse(http.StatusOK, body)
}

// MockErrorResponse creates a mock HTTP error response with a given status code, error code, message, and optional details.
func MockErrorResponse(statusCode int, errorCode, errorMessage, errorDetails string) *http.Response {
	body := `{
		"responseCode": "` + errorCode + `",
		"responseMessage": "` + errorMessage + `"
	}`
	if errorDetails != "" {
		body = body[:len(body)-2] + `,
		"details": "` + errorDetails + `"
	}`
	}
	return MockResponse(statusCode, body)
}

// MockAuthenticationErrorResponse creates a mock authentication error HTTP response
func MockAuthenticationErrorResponse() *http.Response {
	return MockErrorResponse(http.StatusUnauthorized, "401", "Authentication failed", "Invalid credentials")
}

// MockValidationErrorResponse creates a mock validation error HTTP response
func MockValidationErrorResponse() *http.Response {
	return MockErrorResponse(http.StatusBadRequest, "400", "Validation failed", "Invalid request parameters")
}

// MockServerErrorResponse creates a mock server error HTTP response
func MockServerErrorResponse() *http.Response {
	return MockErrorResponse(http.StatusInternalServerError, "500", "Internal server error", "An unexpected error occurred")
}

// MockNotFoundErrorResponse creates a mock not found error HTTP response
func MockNotFoundErrorResponse() *http.Response {
	return MockErrorResponse(http.StatusNotFound, "404", "Not found", "Resource not found")
}
