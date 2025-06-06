# faspay-sendme-snap-go

A Go client library for integrating Faspay's SendMe Snap API. This library provides an easy and secure way to interact with Faspay's payout services, supporting features like fund disbursement, transaction tracking, and status inquiry. Designed for simplicity and scalability in modern Go applications.
test
## Installation

```bash
go get github.com/yourusername/faspay-sendme-snap-go
```

## Features

- Simple and intuitive API client
- Comprehensive error handling
- Support for all Faspay SendMe Snap API endpoints
- Configurable HTTP client with timeout options
- Detailed documentation and examples

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/yourusername/faspay-sendme-snap-go/snap"
)

func main() {
	// Initialize the client with your API credentials
	client := snap.NewClient(
		"https://api.faspay.co.id", // Replace with the actual API base URL
		"your-api-key",             // Replace with your actual API key
		"your-api-secret",          // Replace with your actual API secret
	)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Get account balance
	balance, err := client.GetBalance(ctx)
	if err != nil {
		log.Fatalf("Error getting balance: %v", err)
	}
	fmt.Printf("Balance: %f %s\n", balance.Balance, balance.Currency)

	// Disburse funds
	disbursementReq := snap.DisbursementRequest{
		ReferenceID:    fmt.Sprintf("TRX-%d", time.Now().Unix()),
		Amount:         1000000.0,
		Currency:       "IDR",
		BankCode:       "BCA",
		AccountName:    "John Doe",
		AccountNumber:  "1234567890",
		Description:    "Payment for services",
	}

	disbursement, err := client.DisburseFunds(ctx, disbursementReq)
	if err != nil {
		log.Fatalf("Error disbursing funds: %v", err)
	}
	fmt.Printf("Disbursement successful! Transaction ID: %s\n", disbursement.TransactionID)
}
```

## API Reference

### Client Initialization

```go
// Create a new client with default options
client := snap.NewClient(baseURL, apiKey, apiSecret)

// Create a client with custom timeout
client := snap.NewClient(
    baseURL, 
    apiKey, 
    apiSecret, 
    snap.WithTimeout(60*time.Second)
)

// Create a client with custom HTTP client
httpClient := &http.Client{
    Timeout: 45 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        10,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     30 * time.Second,
    },
}
client := snap.NewClient(
    baseURL, 
    apiKey, 
    apiSecret, 
    snap.WithHTTPClient(httpClient)
)
```

### Available Methods

#### Disbursement

```go
// Disburse funds to a bank account
disbursement, err := client.DisburseFunds(ctx, disbursementRequest)
```

#### Transaction Status

```go
// Check status by transaction ID
status, err := client.GetTransactionStatus(ctx, snap.TransactionStatusRequest{
    TransactionID: "TRX-123456",
})

// Check status by reference ID
status, err := client.GetTransactionStatus(ctx, snap.TransactionStatusRequest{
    ReferenceID: "REF-123456",
})
```

#### Balance Inquiry

```go
// Get account balance
balance, err := client.GetBalance(ctx)
```

#### Bank List

```go
// Get list of supported banks
banks, err := client.GetBankList(ctx)
```

#### Transaction List

```go
// List transactions with filters
transactions, err := client.ListTransactions(ctx, snap.TransactionListRequest{
    StartDate: "2023-01-01",
    EndDate:   "2023-01-31",
    Status:    "success",
    Page:      1,
    Limit:     10,
})
```

### Error Handling

The SDK provides custom error types and helper functions for better error handling:

```go
balance, err := client.GetBalance(ctx)
if err != nil {
    if snap.IsAuthenticationError(err) {
        // Handle authentication error
    } else if snap.IsValidationError(err) {
        // Handle validation error
    } else if snap.IsRateLimitError(err) {
        // Handle rate limit error
    } else if snap.IsServerError(err) {
        // Handle server error
    } else {
        // Handle other errors
    }
}
```

## Examples

For more detailed examples, see the [examples](./examples) directory.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
