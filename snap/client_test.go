package snap

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"
	"time"
)

// TestNewClient tests the NewClient function
func TestNewClient(t *testing.T) {
	// Test with valid parameters
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", "20250607004236909", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Verify client properties
	if client.PartnerId != "99999" {
		t.Errorf("Expected PartnerId to be '99999', got '%s'", client.PartnerId)
	}
	if client.ExternalId != "20250607004236909" {
		t.Errorf("Expected ExternalId to be '20250607004236909', got '%s'", client.ExternalId)
	}
	if client.environment != "sandbox" {
		t.Errorf("Expected environment to be 'sandbox', got '%s'", client.environment)
	}
	if client.baseURL != DefaultBaseURL {
		t.Errorf("Expected baseURL to be '%s', got '%s'", DefaultBaseURL, client.baseURL)
	}
	if client.timeout != time.Duration(DefaultTimeout)*time.Second {
		t.Errorf("Expected timeout to be %v, got %v", time.Duration(DefaultTimeout)*time.Second, client.timeout)
	}

	// Test with custom timeout
	customTimeout := 60 * time.Second
	client, err = NewClient("99999", "20250607004236909", privateKey, WithTimeout(customTimeout))
	if err != nil {
		t.Fatalf("Failed to create client with custom timeout: %v", err)
	}
	if client.timeout != customTimeout {
		t.Errorf("Expected timeout to be %v, got %v", customTimeout, client.timeout)
	}

	// Test with custom HTTP client
	customHTTPClient := &http.Client{Timeout: 45 * time.Second}
	client, err = NewClient("99999", "20250607004236909", privateKey, WithHTTPClient(customHTTPClient))
	if err != nil {
		t.Fatalf("Failed to create client with custom HTTP client: %v", err)
	}
	if client.httpClient != customHTTPClient {
		t.Errorf("Expected httpClient to be %v, got %v", customHTTPClient, client.httpClient)
	}
}

// TestSetEnv tests the SetEnv method
func TestSetEnv(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", "20250607004236909", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test setting environment to sandbox
	err = client.SetEnv("sandbox")
	if err != nil {
		t.Errorf("Failed to set environment to sandbox: %v", err)
	}
	if client.baseURL != baseUrlSandbox {
		t.Errorf("Expected baseURL to be '%s', got '%s'", baseUrlSandbox, client.baseURL)
	}
	if client.environment != "sandbox" {
		t.Errorf("Expected environment to be 'sandbox', got '%s'", client.environment)
	}

	// Test setting environment to prod
	err = client.SetEnv("prod")
	if err != nil {
		t.Errorf("Failed to set environment to prod: %v", err)
	}
	if client.baseURL != baseUrlProd {
		t.Errorf("Expected baseURL to be '%s', got '%s'", baseUrlProd, client.baseURL)
	}
	if client.environment != "prod" {
		t.Errorf("Expected environment to be 'prod', got '%s'", client.environment)
	}

	// Test setting environment to invalid value
	err = client.SetEnv("invalid")
	if err == nil {
		t.Error("Expected error when setting environment to invalid value, got nil")
	}
}

// TestGenerateSignatureSnap tests the GenerateSignatureSnap function
func TestGenerateSignatureSnap(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	// Test with valid parameters
	httpMethod := "POST"
	endpointUrl := "/account/v1.0/account-inquiry-external"
	requestBody := `{"beneficiaryBankCode":"008","beneficiaryAccountNo":"60004400184","partnerReferenceNo":"20250606234037372","additionalInfo":{"sourceAccount":"9920017573"}}`
	timeStamp := "2023-01-01T12:00:00+07:00"

	signature, err := GenerateSignatureSnap(httpMethod, endpointUrl, requestBody, timeStamp, privateKey)
	if err != nil {
		t.Fatalf("Failed to generate signature: %v", err)
	}
	if signature == "" {
		t.Error("Expected non-empty signature, got empty string")
	}

	// Test with invalid private key
	_, err = GenerateSignatureSnap(httpMethod, endpointUrl, requestBody, timeStamp, []byte("invalid-key"))
	if err == nil {
		t.Error("Expected error when generating signature with invalid private key, got nil")
	}
}

// TestDoRequest tests the doRequest method
func TestDoRequest(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	// Create a mock HTTP client that returns a success response
	mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
		// Verify request headers
		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be 'application/json', got '%s'", req.Header.Get("Content-Type"))
		}
		if req.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header to be 'application/json', got '%s'", req.Header.Get("Accept"))
		}
		if req.Header.Get("X-PARTNER-ID") != "99999" {
			t.Errorf("Expected X-PARTNER-ID header to be '99999', got '%s'", req.Header.Get("X-PARTNER-ID"))
		}
		if req.Header.Get("X-EXTERNAL-ID") != "20250607004236909" {
			t.Errorf("Expected X-EXTERNAL-ID header to be '20250607004236909', got '%s'", req.Header.Get("X-EXTERNAL-ID"))
		}
		if req.Header.Get("CHANNEL-ID") != "88001" {
			t.Errorf("Expected CHANNEL-ID header to be '88001', got '%s'", req.Header.Get("CHANNEL-ID"))
		}
		if req.Header.Get("X-TIMESTAMP") == "" {
			t.Error("Expected X-TIMESTAMP header to be non-empty")
		}
		if req.Header.Get("X-SIGNATURE") == "" {
			t.Error("Expected X-SIGNATURE header to be non-empty")
		}

		return MockSuccessResponse(), nil
	})

	// Create a client with the mock HTTP client
	client, err := NewClient("99999", "20250607004236909", privateKey, WithHTTPClient(mockHTTPClient))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test doRequest with valid parameters
	ctx := context.Background()
	request := &ExternalAccountInquiryRequest{
		BeneficiaryBankCode:  "008",
		BeneficiaryAccountNo: "60004400184",
		PartnerReferenceNo:   "20250606234037372",
		AdditionalInfo: &AdditionalInfoRequest{
			SourceAccount: "9920017573",
		},
	}

	resp, err := client.doRequest(ctx, http.MethodPost, EndpointAccountInquiry, request)
	if err != nil {
		t.Fatalf("Failed to do request: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code to be %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Test doRequest with HTTP client error
	mockHTTPClient = NewMockClient(func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("HTTP client error")
	})
	client, err = NewClient("99999", "20250607004236909", privateKey, WithHTTPClient(mockHTTPClient))
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	_, err = client.doRequest(ctx, http.MethodPost, EndpointAccountInquiry, request)
	if err == nil {
		t.Error("Expected error when doing request with HTTP client error, got nil")
	}
}

// TestParseResponse tests the parseResponse method
func TestParseResponse(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", "20250607004236909", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// Test parseResponse with valid response
	resp := MockSuccessResponse()
	var response ExternalAccountInquiryResponse
	err = client.parseResponse(resp, &response)
	if err != nil {
		t.Fatalf("Failed to parse response: %v", err)
	}

	// Verify response fields
	if response.ResponseCode != "00" {
		t.Errorf("Expected ResponseCode to be '00', got '%s'", response.ResponseCode)
	}
	if response.ResponseMessage != "Success" {
		t.Errorf("Expected ResponseMessage to be 'Success', got '%s'", response.ResponseMessage)
	}
	if response.ReferenceNo != "REF123456789" {
		t.Errorf("Expected ReferenceNo to be 'REF123456789', got '%s'", response.ReferenceNo)
	}
	if response.PartnerReferenceNo != "20250606234037372" {
		t.Errorf("Expected PartnerReferenceNo to be '20250606234037372', got '%s'", response.PartnerReferenceNo)
	}
	if response.BeneficiaryAccountName != "JOHN DOE" {
		t.Errorf("Expected BeneficiaryAccountName to be 'JOHN DOE', got '%s'", response.BeneficiaryAccountName)
	}
	if response.BeneficiaryAccountNo != "60004400184" {
		t.Errorf("Expected BeneficiaryAccountNo to be '60004400184', got '%s'", response.BeneficiaryAccountNo)
	}
	if response.BeneficiaryBankCode != "008" {
		t.Errorf("Expected BeneficiaryBankCode to be '008', got '%s'", response.BeneficiaryBankCode)
	}
	if response.BeneficiaryBankName != "MANDIRI" {
		t.Errorf("Expected BeneficiaryBankName to be 'MANDIRI', got '%s'", response.BeneficiaryBankName)
	}
	if response.Currency != "IDR" {
		t.Errorf("Expected Currency to be 'IDR', got '%s'", response.Currency)
	}
	if response.AdditionalInfo == nil {
		t.Error("Expected AdditionalInfo to be non-nil")
	} else {
		if response.AdditionalInfo.Status != "success" {
			t.Errorf("Expected AdditionalInfo.Status to be 'success', got '%s'", response.AdditionalInfo.Status)
		}
		if response.AdditionalInfo.Message != "Account inquiry successful" {
			t.Errorf("Expected AdditionalInfo.Message to be 'Account inquiry successful', got '%s'", response.AdditionalInfo.Message)
		}
	}

	// Test parseResponse with invalid JSON
	resp = MockResponse(http.StatusOK, "invalid-json")
	err = client.parseResponse(resp, &response)
	if err == nil {
		t.Error("Expected error when parsing invalid JSON, got nil")
	}
}
