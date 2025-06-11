package snap

import (
	"context"
	"errors"
	"net/http"
	"os"
	"testing"
)

// MockTransferInterBankSuccessResponse creates a mock successful HTTP response for transfer inter-bank
func MockTransferInterBankSuccessResponse() *http.Response {
	body := `{
		"responseCode": "00",
		"responseMessage": "Success",
		"referenceNo": "REF123456789",
		"partnerReferenceNo": "TRX123456789",
		"amount": {
			"value": "10000.00",
			"currency": "IDR"
		},
		"beneficiaryAccountNo": "60004400184",
		"beneficiaryBankCode": "008",
		"sourceAccountNo": "9920017573",
		"additionalInfo": {
			"beneficiaryAccountName": "JOHN DOE",
			"beneficiaryBankName": "MANDIRI",
			"instructDate": "",
			"transactionDescription": "Payment for services",
			"callbackUrl": "https://your-callback-url.com/callback",
			"latestTransactionStatus": "SUCCESS",
			"transactionStatusDesc": "Transaction successful"
		}
	}`
	return MockResponse(http.StatusOK, body)
}

// MockInquiryBalanceSuccessResponse creates a mock successful HTTP response for balance inquiry
func MockInquiryBalanceSuccessResponse() *http.Response {
	body := `{
		"responseCode": "00",
		"responseMessage": "Success",
		"accountNo": "9920017573",
		"accountInfos": [
			{
				"balanceType": "AVAILABLE",
				"amount": {
					"value": "100000.00",
					"currency": "IDR"
				},
				"availableBalance": {
					"value": "100000.00",
					"currency": "IDR"
				},
				"status": "ACTIVE"
			}
		]
	}`
	return MockResponse(http.StatusOK, body)
}

// MockStatusTransferSuccessResponse creates a mock successful HTTP response for status transfer
func MockStatusTransferSuccessResponse() *http.Response {
	body := `{
		"responseCode": "00",
		"responseMessage": "Success",
		"originalReferenceNo": "53883",
		"originalPartnerReferenceNo": "20250609103003234",
		"serviceCode": "18",
		"transactionDate": "2025-06-09T10:30:03+07:00",
		"amount": {
			"value": "10000.00",
			"currency": "IDR"
		},
		"beneficiaryAccountNo": "60004400184",
		"beneficiaryBankCode": "008",
		"referenceNumber": "REF123456789",
		"sourceAccountNo": "9920017573",
		"latestTransactionStatus": "SUCCESS",
		"transactionStatusDesc": "Transaction successful",
		"additionalInfo": {
			"beneficiaryAccountName": "JOHN DOE",
			"beneficiaryBankName": "MANDIRI",
			"transactionDescription": "Payment for services",
			"callbackUrl": "https://your-callback-url.com/callback",
			"transactionStatusDate": "2025-06-09T10:35:03+07:00"
		}
	}`
	return MockResponse(http.StatusOK, body)
}

// MockHistoryListSuccessResponse creates a mock successful HTTP response for history list
func MockHistoryListSuccessResponse() *http.Response {
	body := `{
		"responseCode": "00",
		"responseMessage": "Success",
		"detailData": [
			{
				"dateTime": "2024-12-15T10:30:03+07:00",
				"amount": {
					"value": "10000.00",
					"currency": "IDR"
				},
				"remark": "Payment for services",
				"sourceOfFunds": [
					{
						"source": "BANK_ACCOUNT"
					}
				],
				"status": "SUCCESS",
				"type": "TRANSFER",
				"additionalInfo": {
					"debitCredit": "DEBIT"
				}
			}
		],
		"additionalInfo": {
			"accountNo": "9920017573",
			"fromDateTime": "2024-12-01T00:00:00-07:00",
			"toDateTime": "2024-12-30T00:00:00-07:00",
			"message": "Transaction history retrieved successfully"
		}
	}`
	return MockResponse(http.StatusOK, body)
}

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
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
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
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
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
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
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
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
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
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
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
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
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

// TestTransferInterBank tests the TransferInterBank method
func TestTransferInterBank(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	// Test successful transfer
	t.Run("Success", func(t *testing.T) {
		// Create a mock HTTP client that returns a success response
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			// Verify request method and path
			if req.Method != http.MethodPost {
				t.Errorf("Expected request method to be POST, got %s", req.Method)
			}
			if req.URL.Path != EndpointTransferInterbank {
				t.Errorf("Expected request path to be %s, got %s", EndpointTransferInterbank, req.URL.Path)
			}

			return MockTransferInterBankSuccessResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &TransferInterBankRequest{
			PartnerReferenceNo: "TRX123456789",
			Amount: &Amount{
				Value:    "10000.00",
				Currency: "IDR",
			},
			BeneficiaryAccountName: "John Doe",
			BeneficiaryAccountNo:   "60004400184",
			BeneficiaryBankCode:    "008",
			BeneficiaryEmail:       "john@example.com",
			SourceAccountNo:        "9920017573",
			TransactionDate:        "2025-06-09T10:30:03+07:00",
			AdditionalInfo: &AdditionalInfoTransferInterBank{
				InstructDate:           "",
				TransactionDescription: "Payment for services",
				CallbackUrl:            "https://your-callback-url.com/callback",
			},
		}

		// Call TransferInterBank
		ctx := context.Background()
		response, err := client.TransferInterBank(ctx, request)
		if err != nil {
			t.Fatalf("Failed to call TransferInterBank: %v", err)
		}

		// Verify response
		if response.ResponseCode != "00" {
			t.Errorf("Expected ResponseCode to be '00', got '%s'", response.ResponseCode)
		}
		if response.ResponseMessage != "Success" {
			t.Errorf("Expected ResponseMessage to be 'Success', got '%s'", response.ResponseMessage)
		}
		if response.Amount.Value != "10000.00" {
			t.Errorf("Expected Amount.Value to be '10000.00', got '%s'", response.Amount.Value)
		}
		if response.AdditionalInfo.BeneficiaryAccountName != "JOHN DOE" {
			t.Errorf("Expected AdditionalInfo.BeneficiaryAccountName to be 'JOHN DOE', got '%s'", response.AdditionalInfo.BeneficiaryAccountName)
		}
	})

	// Test HTTP client error
	t.Run("HTTPClientError", func(t *testing.T) {
		// Create a mock HTTP client that returns an error
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			return nil, errors.New("HTTP client error")
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &TransferInterBankRequest{
			PartnerReferenceNo: "TRX123456789",
			Amount: &Amount{
				Value:    "10000.00",
				Currency: "IDR",
			},
			BeneficiaryAccountName: "John Doe",
			BeneficiaryAccountNo:   "60004400184",
			BeneficiaryBankCode:    "008",
			BeneficiaryEmail:       "john@example.com",
			SourceAccountNo:        "9920017573",
			TransactionDate:        "2025-06-09T10:30:03+07:00",
			AdditionalInfo: &AdditionalInfoTransferInterBank{
				InstructDate:           "",
				TransactionDescription: "Payment for services",
				CallbackUrl:            "https://your-callback-url.com/callback",
			},
		}

		// Call TransferInterBank
		ctx := context.Background()
		_, err = client.TransferInterBank(ctx, request)
		if err == nil {
			t.Error("Expected error when calling TransferInterBank with HTTP client error, got nil")
		}
	})
}

// TestInquiryBalance tests the InquiryBalance method
func TestInquiryBalance(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	// Test successful balance inquiry
	t.Run("Success", func(t *testing.T) {
		// Create a mock HTTP client that returns a success response
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			// Verify request method and path
			if req.Method != http.MethodPost {
				t.Errorf("Expected request method to be POST, got %s", req.Method)
			}
			if req.URL.Path != EndpointInquiryBalance {
				t.Errorf("Expected request path to be %s, got %s", EndpointInquiryBalance, req.URL.Path)
			}

			return MockInquiryBalanceSuccessResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &InquiryBalanceRequest{
			AccountNo: "9920017573",
		}

		// Call InquiryBalance
		ctx := context.Background()
		response, err := client.InquiryBalance(ctx, request)
		if err != nil {
			t.Fatalf("Failed to call InquiryBalance: %v", err)
		}

		// Verify response
		if response.ResponseCode != "00" {
			t.Errorf("Expected ResponseCode to be '00', got '%s'", response.ResponseCode)
		}
		if response.ResponseMessage != "Success" {
			t.Errorf("Expected ResponseMessage to be 'Success', got '%s'", response.ResponseMessage)
		}
		if response.AccountNo != "9920017573" {
			t.Errorf("Expected AccountNo to be '9920017573', got '%s'", response.AccountNo)
		}
		if len(response.AccountInfos) != 1 {
			t.Errorf("Expected 1 AccountInfo, got %d", len(response.AccountInfos))
		} else {
			if response.AccountInfos[0].Amount.Value != "100000.00" {
				t.Errorf("Expected Amount.Value to be '100000.00', got '%s'", response.AccountInfos[0].Amount.Value)
			}
		}
	})
}

// TestStatusTransfer tests the StatusTransfer method
func TestStatusTransfer(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	// Test successful status transfer
	t.Run("Success", func(t *testing.T) {
		// Create a mock HTTP client that returns a success response
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			// Verify request method and path
			if req.Method != http.MethodPost {
				t.Errorf("Expected request method to be POST, got %s", req.Method)
			}
			if req.URL.Path != EndpointInquiryStatus {
				t.Errorf("Expected request path to be %s, got %s", EndpointInquiryStatus, req.URL.Path)
			}

			return MockStatusTransferSuccessResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &StatusTransferRequest{
			OriginalPartnerReferenceNo: "20250609103003234",
			OriginalReferenceNo:        "53883",
			ServiceCode:                "18",
		}

		// Call StatusTransfer
		ctx := context.Background()
		response, err := client.StatusTransfer(ctx, request)
		if err != nil {
			t.Fatalf("Failed to call StatusTransfer: %v", err)
		}

		// Verify response
		if response.ResponseCode != "00" {
			t.Errorf("Expected ResponseCode to be '00', got '%s'", response.ResponseCode)
		}
		if response.ResponseMessage != "Success" {
			t.Errorf("Expected ResponseMessage to be 'Success', got '%s'", response.ResponseMessage)
		}
		if response.OriginalReferenceNo != "53883" {
			t.Errorf("Expected OriginalReferenceNo to be '53883', got '%s'", response.OriginalReferenceNo)
		}
		if response.LatestTransactionStatus != "SUCCESS" {
			t.Errorf("Expected LatestTransactionStatus to be 'SUCCESS', got '%s'", response.LatestTransactionStatus)
		}
	})
}

// TestHistoryList tests the HistoryList method
func TestHistoryList(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	// Test successful history list
	t.Run("Success", func(t *testing.T) {
		// Create a mock HTTP client that returns a success response
		mockHTTPClient := NewMockClient(func(req *http.Request) (*http.Response, error) {
			// Verify request method and path
			if req.Method != http.MethodPost {
				t.Errorf("Expected request method to be POST, got %s", req.Method)
			}
			if req.URL.Path != EndpointHistoryList {
				t.Errorf("Expected request path to be %s, got %s", EndpointHistoryList, req.URL.Path)
			}

			return MockHistoryListSuccessResponse(), nil
		})

		// Create a client with the mock HTTP client
		client, err := NewClient("99999", privateKey, getCertSSL(), WithHTTPClient(mockHTTPClient))
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}

		// Create a request
		request := &HistoryListRequest{
			FromDateTime: "2024-12-01T00:00:00-07:00",
			ToDateTime:   "2024-12-30T00:00:00-07:00",
			AdditionalInfo: &AdditionalHistoryListRequest{
				AccountNo: "9920017573",
			},
		}

		// Call HistoryList
		ctx := context.Background()
		response, err := client.HistoryList(ctx, request)
		if err != nil {
			t.Fatalf("Failed to call HistoryList: %v", err)
		}

		// Verify response
		if response.ResponseCode != "00" {
			t.Errorf("Expected ResponseCode to be '00', got '%s'", response.ResponseCode)
		}
		if response.ResponseMessage != "Success" {
			t.Errorf("Expected ResponseMessage to be 'Success', got '%s'", response.ResponseMessage)
		}
		if len(response.DetailData) != 1 {
			t.Errorf("Expected 1 DetailData, got %d", len(response.DetailData))
		} else {
			if response.DetailData[0].Amount.Value != "10000.00" {
				t.Errorf("Expected Amount.Value to be '10000.00', got '%s'", response.DetailData[0].Amount.Value)
			}
		}
	})
}
