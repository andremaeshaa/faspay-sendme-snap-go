package snap

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"
)

// TestAccountInquiry tests the AccountInquiry method
func TestAccountInquiry(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	// Test successful account inquiry
	t.Run("Success", func(t *testing.T) {
		// Create a mock HTTP client that returns a success response
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			// Verify request method and path
			if req.Method != http.MethodPost {
				t.Errorf("Expected request method to be POST, got %s", req.Method)
			}
			if req.URL.Path != EndpointAccountInquiry {
				t.Errorf("Expected request path to be %s, got %s", EndpointAccountInquiry, req.URL.Path)
			}

			return MockSuccessResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &ExternalAccountInquiryRequest{
			BeneficiaryBankCode:  "008",
			BeneficiaryAccountNo: "60004400184",
			PartnerReferenceNo:   "20250606234037372",
			AdditionalInfo: &AdditionalInfoInquiryAccount{
				SourceAccount: "9920017573",
			},
		}

		// Call AccountInquiry
		ctx := context.Background()
		response, err := client.AccountInquiry(ctx, request)
		if err != nil {
			t.Fatalf("Failed to call AccountInquiry: %v", err)
		}

		// Verify response
		if response.ResponseCode != "00" {
			t.Errorf("Expected ResponseCode to be '00', got '%s'", response.ResponseCode)
		}
		if response.ResponseMessage != "Success" {
			t.Errorf("Expected ResponseMessage to be 'Success', got '%s'", response.ResponseMessage)
		}
		if response.BeneficiaryAccountName != "JOHN DOE" {
			t.Errorf("Expected BeneficiaryAccountName to be 'JOHN DOE', got '%s'", response.BeneficiaryAccountName)
		}
	})

	// Test HTTP client error
	t.Run("HTTPClientError", func(t *testing.T) {
		// Create a mock HTTP client that returns an error
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("HTTP client error")
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &ExternalAccountInquiryRequest{
			BeneficiaryBankCode:  "008",
			BeneficiaryAccountNo: "60004400184",
			PartnerReferenceNo:   "20250606234037372",
			AdditionalInfo: &AdditionalInfoInquiryAccount{
				SourceAccount: "9920017573",
			},
		}

		// Call AccountInquiry
		ctx := context.Background()
		_, err = client.AccountInquiry(ctx, request)
		if err == nil {
			t.Error("Expected error when calling AccountInquiry with HTTP client error, got nil")
		}
	})

	// Test authentication error
	t.Run("AuthenticationError", func(t *testing.T) {
		// Create a mock HTTP client that returns an authentication error
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			return MockAuthenticationErrorResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &ExternalAccountInquiryRequest{
			BeneficiaryBankCode:  "008",
			BeneficiaryAccountNo: "60004400184",
			PartnerReferenceNo:   "20250606234037372",
			AdditionalInfo: &AdditionalInfoInquiryAccount{
				SourceAccount: "9920017573",
			},
		}

		// Call AccountInquiry
		ctx := context.Background()
		_, err = client.AccountInquiry(ctx, request)
		if err == nil {
			t.Error("Expected error when calling AccountInquiry with authentication error, got nil")
		}
		if !IsAuthenticationError(err) {
			t.Errorf("Expected IsAuthenticationError to return true, got false")
		}
	})

	// Test validation error
	t.Run("ValidationError", func(t *testing.T) {
		// Create a mock HTTP client that returns a validation error
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			return MockValidationErrorResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &ExternalAccountInquiryRequest{
			BeneficiaryBankCode:  "008",
			BeneficiaryAccountNo: "60004400184",
			PartnerReferenceNo:   "20250606234037372",
			AdditionalInfo: &AdditionalInfoInquiryAccount{
				SourceAccount: "9920017573",
			},
		}

		// Call AccountInquiry
		ctx := context.Background()
		_, err = client.AccountInquiry(ctx, request)
		if err == nil {
			t.Error("Expected error when calling AccountInquiry with validation error, got nil")
		}
		if !IsValidationError(err) {
			t.Errorf("Expected IsValidationError to return true, got false")
		}
	})

	// Test server error
	t.Run("ServerError", func(t *testing.T) {
		// Create a mock HTTP client that returns a server error
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			return MockServerErrorResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &ExternalAccountInquiryRequest{
			BeneficiaryBankCode:  "008",
			BeneficiaryAccountNo: "60004400184",
			PartnerReferenceNo:   "20250606234037372",
			AdditionalInfo: &AdditionalInfoInquiryAccount{
				SourceAccount: "9920017573",
			},
		}

		// Call AccountInquiry
		ctx := context.Background()
		_, err = client.AccountInquiry(ctx, request)
		println("error tot: ", err.Error())
		if err == nil {
			t.Error("Expected error when calling AccountInquiry with server error, got nil")
		}
		if !IsServerError(err) {
			t.Errorf("Expected IsServerError to return true, got false")
		}
	})

	// Test not found error
	t.Run("NotFoundError", func(t *testing.T) {
		// Create a mock HTTP client that returns a not found error
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			return MockNotFoundErrorResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &ExternalAccountInquiryRequest{
			BeneficiaryBankCode:  "008",
			BeneficiaryAccountNo: "60004400184",
			PartnerReferenceNo:   "20250606234037372",
			AdditionalInfo: &AdditionalInfoInquiryAccount{
				SourceAccount: "9920017573",
			},
		}

		// Call AccountInquiry
		ctx := context.Background()
		_, err = client.AccountInquiry(ctx, request)
		if err == nil {
			t.Error("Expected error when calling AccountInquiry with not found error, got nil")
		}
		if !IsNotFoundError(err) {
			t.Errorf("Expected IsNotFoundError to return true, got false")
		}
	})
}
