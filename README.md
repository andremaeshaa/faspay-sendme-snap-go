# Faspay SendMe Snap Go SDK

A Go client library for integrating Faspay's SendMe Snap API. This library provides an easy and secure way to interact with Faspay's payment services, supporting features like account inquiry, fund transfers, transaction tracking, and balance inquiry. Designed for simplicity and scalability in modern Go applications.

## Installation

```bash
go get github.com/faspay-team/faspay-sendme-snap-go
```

## Features

- Simple and intuitive API client
- Comprehensive error handling
- Support for all Faspay SendMe Snap API endpoints
- Configurable HTTP client with timeout options
- Detailed documentation and examples
- Secure request signing

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/faspay-team/faspay-sendme-snap-go/snap"
)

func main() {
	// Step 1: Load the private key from file
	privateKeyPath := "./certs/enc.key"
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Failed to read private key: %v", err)
	}

	// Step 2: Initialize the client
	partnerId := "99999" // Your 5-digit partner ID

	// Create a new client with a custom timeout
	client, err := snap.NewClient(
		partnerId,
		privateKey,
		snap.WithTimeout(60*time.Second), // Optional: Set a custom timeout
	)
	if err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}

	// Step 3: Set the environment (sandbox or prod)
	err = client.SetEnv("sandbox")
	if err != nil {
		log.Fatalf("Failed to set environment: %v", err)
	}

	// Step 4: Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Step 5: Perform an account inquiry
	request := &snap.ExternalAccountInquiryRequest{
		BeneficiaryBankCode:  "008",               // Bank code (e.g., "008" for Mandiri)
		BeneficiaryAccountNo: "60004400184",       // Account number
		PartnerReferenceNo:   "20250606234037372", // Your unique reference number
		AdditionalInfo: &snap.AdditionalInfoInquiryAccount{
			SourceAccount: "9920017573", // Source account number
		},
	}

	response, err := client.AccountInquiry(ctx, request)
	if err != nil {
		log.Fatalf("Error performing account inquiry: %v", err)
	}

	fmt.Printf("Account inquiry successful!\n")
	fmt.Printf("Account Name: %s\n", response.BeneficiaryAccountName)
	fmt.Printf("Account Number: %s\n", response.BeneficiaryAccountNo)
	fmt.Printf("Bank: %s (%s)\n", response.BeneficiaryBankName, response.BeneficiaryBankCode)
}
```

## API Reference

### Client Initialization

```go
// Load private key
privateKey, err := os.ReadFile("./certs/enc.key")
if err != nil {
    log.Fatalf("Failed to read private key: %v", err)
}

// Create a new client with default options
client, err := snap.NewClient("99999", privateKey)
if err != nil {
    log.Fatalf("Failed to initialize client: %v", err)
}

// Create a client with custom timeout
client, err := snap.NewClient(
    "99999", 
    privateKey, 
    snap.WithTimeout(60*time.Second)
)
if err != nil {
    log.Fatalf("Failed to initialize client: %v", err)
}

// Create a client with custom HTTP client
httpClient := &http.Client{
    Timeout: 45 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        10,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     30 * time.Second,
    },
}
client, err := snap.NewClient(
    "99999", 
    privateKey, 
    snap.WithHTTPClient(httpClient)
)
if err != nil {
    log.Fatalf("Failed to initialize client: %v", err)
}

// Set environment (sandbox or prod)
err = client.SetEnv("sandbox") // or "prod" for production
if err != nil {
    log.Fatalf("Failed to set environment: %v", err)
}
```

### Available Methods

#### Account Inquiry

Inquire about an external account's details.

```go
request := &snap.ExternalAccountInquiryRequest{
    BeneficiaryBankCode:  "008",               // Bank code
    BeneficiaryAccountNo: "60004400184",       // Account number
    PartnerReferenceNo:   "20250606234037372", // Your unique reference number
    AdditionalInfo: &snap.AdditionalInfoInquiryAccount{
        SourceAccount: "9920017573", // Source account number
    },
}

response, err := client.AccountInquiry(ctx, request)
```

#### Transfer Inter-Bank

Transfer funds between banks.

```go
request := &snap.TransferInterBankRequest{
    PartnerReferenceNo: "TRX123456789",
    Amount: &snap.Amount{
        Value:    "10000.00", // Amount to transfer
        Currency: "IDR",      // Currency
    },
    BeneficiaryAccountName: "John Doe",       // Recipient name
    BeneficiaryAccountNo:   "60004400184",    // Recipient account number
    BeneficiaryBankCode:    "008",            // Recipient bank code
    BeneficiaryEmail:       "john@example.com", // Recipient email
    SourceAccountNo:        "9920017573",     // Source account number
    TransactionDate:        time.Now().Format("2006-01-02T15:04:05-07:00"), // Current time
    AdditionalInfo: &snap.AdditionalInfoTransferInterBank{
        InstructDate:           "", // Optional instruction date
        TransactionDescription: "Payment for services", // Description
        CallbackUrl:            "https://your-callback-url.com/callback", // Callback URL
    },
}

response, err := client.TransferInterBank(ctx, request)
```

#### Check Transfer Status

Check the status of a transfer.

```go
request := &snap.StatusTransferRequest{
    OriginalPartnerReferenceNo: "20250609103003234", // Original reference number from transfer
    OriginalReferenceNo:        "53883",             // Original reference number from response
    ServiceCode:                "18",                // Service code (18 for transfer)
}

response, err := client.StatusTransfer(ctx, request)
```

#### Balance Inquiry

Check account balance.

```go
request := &snap.InquiryBalanceRequest{
    AccountNo: "9920017573", // Account number to check balance
}

response, err := client.InquiryBalance(ctx, request)
```

#### Transaction History

Get transaction history.

```go
request := &snap.HistoryListRequest{
    FromDateTime: "2024-12-01T00:00:00-07:00", // Start date
    ToDateTime:   "2024-12-30T00:00:00-07:00", // End date
    AdditionalInfo: &snap.AdditionalHistoryListRequest{
        AccountNo: "9920017573", // Account number
    },
}

response, err := client.HistoryList(ctx, request)
```

#### Customer Topup

Perform a customer top-up.

```go
request := &snap.CustomerTopupRequest{
    PartnerReferenceNo: "20250609150352617",
    CustomerNumber:     "0812254830",
    Amount: &snap.Amount{
        Value:    "76860.00",
        Currency: "IDR",
    },
    TransactionDate: "2025-06-09T15:03:52+07:00",
    AdditionalInfo: &snap.AdditionalInfoCustomerTopupRequest{
        SourceAccount:          "9920017573",
        PlatformCode:           "gpy",
        InstructDate:           "",
        BeneficiaryEmail:       "customer@example.com",
        TransactionDescription: "Tunjangan Pulsa 20250609",
        CallbackUrl:            "https://your-callback-url.com/callback",
    },
}

response, err := client.CustomerTopup(ctx, request)
```

#### Customer Topup Status

Check the status of a customer top-up.

```go
request := &snap.CustomerTopupStatusRequest{
    OriginalPartnerReferenceNo: "20250609150352616",
    OriginalReferenceNo:        "59732",
    ServiceCode:                "38",
}

response, err := client.CustomerTopupStatus(ctx, request)
```

#### Bill Inquiry

Inquire about a bill.

```go
request := &snap.BillInquiryRequest{
    PartnerReferenceNo: "20250609162756943",
    PartnerServiceId:   "7008",
    CustomerNo:         "08000047816",
    VirtualAccountNo:   "700808000047816",
    AdditionalInfo: &snap.AdditionalInfoBillInquiry{
        BillerCode:    "013",
        SourceAccount: "9920017573",
    },
}

response, err := client.BillInquiry(ctx, request)
```

#### Bill Payment

Pay a bill.

```go
request := &snap.BillPaymentRequest{
    PartnerReferenceNo: "20250609162921210",
    PartnerServiceId:   "7008",
    CustomerNo:         "08000047816",
    VirtualAccountNo:   "700808000047816",
    VirtualAccountName: "DUMMY VA",
    SourceAccount:      "9920017573",
    PaidAmount: &snap.Amount{
        Value:    "41454.00",
        Currency: "IDR",
    },
    TrxDateTime: "2025-06-09T16:29:21",
    AdditionalInfo: &snap.AdditionalInfoBillPayment{
        BillerCode:   "013",
        InstructDate: "2025-06-09T16:29:21+07:00",
        CallbackUrl:  "https://your-callback-url.com/callback",
    },
}

response, err := client.BillPayment(ctx, request)
```

### Error Handling

The SDK provides custom error types and helper functions for better error handling:

```go
response, err := client.AccountInquiry(ctx, request)
if err != nil {
    if snap.IsAuthenticationError(err) {
        // Handle authentication error
    } else if snap.IsValidationError(err) {
        // Handle validation error
    } else if snap.IsNotFoundError(err) {
        // Handle not found error
    } else if snap.IsServerError(err) {
        // Handle server error
    } else {
        // Handle other errors
    }
}
```

## Examples

For more detailed examples, see the [examples](./examples) directory. The examples demonstrate:

1. Account Inquiry - Inquire about an external account's details
2. Balance Inquiry - Check account balance
3. Transfer Inter-Bank - Transfer funds between banks
4. Check Transfer Status - Check the status of a transfer
5. Transaction History - Get transaction history
6. Customer Topup - Perform a customer top-up
7. Bill Inquiry and Payment - Inquire about and pay bills

## Certificate Files

The SDK requires a private key file for signing API requests. The private key should be stored in the `certs` directory:

- `path/to-private-key/enc.key` - Private key for you're environment

Make sure to keep these files secure and not commit them to version control.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
