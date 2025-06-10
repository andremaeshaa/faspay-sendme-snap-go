package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"faspay-sendme-snap-go/snap" // Import the snap package using the module path
)

// This example demonstrates how to use the Faspay SendMe Snap SDK to perform various operations.
// It shows how to initialize the client, make requests, and handle responses and errors.
func main() {
	// Step 1: Load the private key from file
	// The private key is used for signing API requests
	privateKeyPath := "../certs/enc.key" // Path relative to the examples directory
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalf("Failed to read private key: %v", err)
	}

	// Step 2: Initialize the client
	// Replace these values with your actual credentials
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

	// Optional: Set the environment (sandbox or prod)
	// By default, the client uses the sandbox environment
	err = client.SetEnv("sandbox")
	if err != nil {
		log.Fatalf("Failed to set environment: %v", err)
	}

	// Step 3: Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Example 1: Account Inquiry
	exampleAccountInquiry(ctx, client)

	// Example 2: Balance Inquiry
	exampleBalanceInquiry(ctx, client)

	// Example 3: Transfer Inter-Bank
	exampleTransferInterBank(ctx, client)

	// Example 4: Check Transfer Status
	exampleCheckTransferStatus(ctx, client)

	// Example 5: Transaction History
	exampleTransactionHistory(ctx, client)
}

// Example 1: Account Inquiry
func exampleAccountInquiry(ctx context.Context, client snap.Services) {
	fmt.Println("\n=== Example 1: Account Inquiry ===")

	// Create an account inquiry request
	request := &snap.ExternalAccountInquiryRequest{
		BeneficiaryBankCode:  "008",               // Bank code (e.g., "008" for Mandiri)
		BeneficiaryAccountNo: "60004400184",       // Account number
		PartnerReferenceNo:   "20250606234037372", // Your unique reference number
		AdditionalInfo: &snap.AdditionalInfoInquiryAccount{
			SourceAccount: "9920017573", // Source account number
		},
	}

	// Perform the account inquiry
	fmt.Println("Performing account inquiry...")
	response, err := client.AccountInquiry(ctx, request)

	// Handle errors
	if err != nil {
		handleError(err)
		return
	}

	// Process the response
	fmt.Println("Account inquiry successful!")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Message: %s\n", response.ResponseMessage)
	fmt.Printf("Reference No: %s\n", response.ReferenceNo)
	fmt.Printf("Partner Reference No: %s\n", response.PartnerReferenceNo)
	fmt.Printf("Beneficiary Account Name: %s\n", response.BeneficiaryAccountName)
	fmt.Printf("Beneficiary Account No: %s\n", response.BeneficiaryAccountNo)
	fmt.Printf("Beneficiary Bank Code: %s\n", response.BeneficiaryBankCode)
	fmt.Printf("Beneficiary Bank Name: %s\n", response.BeneficiaryBankName)
	fmt.Printf("Currency: %s\n", response.Currency)

	if response.AdditionalInfo != nil {
		fmt.Printf("Additional Info Status: %s\n", response.AdditionalInfo.Status)
		fmt.Printf("Additional Info Message: %s\n", response.AdditionalInfo.Message)
	}
}

// Example 2: Balance Inquiry
func exampleBalanceInquiry(ctx context.Context, client snap.Services) {
	fmt.Println("\n=== Example 2: Balance Inquiry ===")

	// Create a balance inquiry request
	request := &snap.InquiryBalanceRequest{
		AccountNo: "9920017573", // Account number to check balance
	}

	// Perform the balance inquiry
	fmt.Println("Performing balance inquiry...")
	response, err := client.InquiryBalance(ctx, request)

	// Handle errors
	if err != nil {
		handleError(err)
		return
	}

	// Process the response
	fmt.Println("Balance inquiry successful!")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Message: %s\n", response.ResponseMessage)
	fmt.Printf("Account No: %s\n", response.AccountNo)

	// Display account information
	for i, accountInfo := range response.AccountInfos {
		fmt.Printf("\nAccount Info #%d:\n", i+1)
		fmt.Printf("  Balance Type: %s\n", accountInfo.BalanceType)
		fmt.Printf("  Amount: %s %s\n", accountInfo.Amount.Value, accountInfo.Amount.Currency)
		fmt.Printf("  Available Balance: %s %s\n", accountInfo.AvailableBalance.Value, accountInfo.AvailableBalance.Currency)
		fmt.Printf("  Status: %s\n", accountInfo.Status)
	}
}

// Example 3: Transfer Inter-Bank
func exampleTransferInterBank(ctx context.Context, client snap.Services) {
	fmt.Println("\n=== Example 3: Transfer Inter-Bank ===")

	// Generate a unique reference number based on timestamp
	referenceNo := fmt.Sprintf("TRX%d", time.Now().Unix())

	// Create a transfer request
	request := &snap.TransferInterBankRequest{
		PartnerReferenceNo: referenceNo,
		Amount: &snap.Amount{
			Value:    "10000.00", // Amount to transfer
			Currency: "IDR",      // Currency
		},
		BeneficiaryAccountName: "John Doe",                                     // Recipient name
		BeneficiaryAccountNo:   "60004400184",                                  // Recipient account number
		BeneficiaryBankCode:    "008",                                          // Recipient bank code (e.g., "008" for Mandiri)
		BeneficiaryEmail:       "john@example.com",                             // Recipient email
		SourceAccountNo:        "9920017573",                                   // Source account number
		TransactionDate:        time.Now().Format("2006-01-02T15:04:05-07:00"), // Current time
		AdditionalInfo: &snap.AdditionalInfoTransferInterBank{
			InstructDate:           "",                                       // Optional instruction date
			TransactionDescription: "Payment for services",                   // Description
			CallbackUrl:            "https://your-callback-url.com/callback", // Callback URL
		},
	}

	// Perform the transfer
	fmt.Println("Performing inter-bank transfer...")
	response, err := client.TransferInterBank(ctx, request)

	// Handle errors
	if err != nil {
		handleError(err)
		return
	}

	// Process the response
	fmt.Println("Transfer successful!")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Message: %s\n", response.ResponseMessage)
	fmt.Printf("Reference No: %s\n", response.ReferenceNo)
	fmt.Printf("Partner Reference No: %s\n", response.PartnerReferenceNo)
	fmt.Printf("Amount: %s %s\n", response.Amount.Value, response.Amount.Currency)
	fmt.Printf("Beneficiary Account No: %s\n", response.BeneficiaryAccountNo)
	fmt.Printf("Beneficiary Bank Code: %s\n", response.BeneficiaryBankCode)
	fmt.Printf("Source Account No: %s\n", response.SourceAccountNo)

	if response.AdditionalInfo != nil {
		fmt.Printf("Beneficiary Account Name: %s\n", response.AdditionalInfo.BeneficiaryAccountName)
		fmt.Printf("Beneficiary Bank Name: %s\n", response.AdditionalInfo.BeneficiaryBankName)
		fmt.Printf("Transaction Description: %s\n", response.AdditionalInfo.TransactionDescription)
		fmt.Printf("Latest Transaction Status: %s\n", response.AdditionalInfo.LatestTransactionStatus)
		fmt.Printf("Transaction Status Description: %s\n", response.AdditionalInfo.TransactionStatusDesc)
	}
}

// Example 4: Check Transfer Status
func exampleCheckTransferStatus(ctx context.Context, client snap.Services) {
	fmt.Println("\n=== Example 4: Check Transfer Status ===")

	// Create a status request
	request := &snap.StatusTransferRequest{
		OriginalPartnerReferenceNo: "20250609103003234", // Original reference number from transfer
		OriginalReferenceNo:        "53883",             // Original reference number from response
		ServiceCode:                "18",                // Service code (18 for transfer)
	}

	// Check the transfer status
	fmt.Println("Checking transfer status...")
	response, err := client.StatusTransfer(ctx, request)

	// Handle errors
	if err != nil {
		handleError(err)
		return
	}

	// Process the response
	fmt.Println("Status check successful!")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Message: %s\n", response.ResponseMessage)
	fmt.Printf("Original Reference No: %s\n", response.OriginalReferenceNo)
	fmt.Printf("Original Partner Reference No: %s\n", response.OriginalPartnerReferenceNo)
	fmt.Printf("Service Code: %s\n", response.ServiceCode)
	fmt.Printf("Transaction Date: %s\n", response.TransactionDate)
	fmt.Printf("Amount: %s %s\n", response.Amount.Value, response.Amount.Currency)
	fmt.Printf("Beneficiary Account No: %s\n", response.BeneficiaryAccountNo)
	fmt.Printf("Beneficiary Bank Code: %s\n", response.BeneficiaryBankCode)
	fmt.Printf("Reference Number: %s\n", response.ReferenceNumber)
	fmt.Printf("Source Account No: %s\n", response.SourceAccountNo)
	fmt.Printf("Latest Transaction Status: %s\n", response.LatestTransactionStatus)
	fmt.Printf("Transaction Status Description: %s\n", response.TransactionStatusDesc)

	if response.AdditionalInfo != nil {
		fmt.Printf("Beneficiary Account Name: %s\n", response.AdditionalInfo.BeneficiaryAccountName)
		fmt.Printf("Beneficiary Bank Name: %s\n", response.AdditionalInfo.BeneficiaryBankName)
		fmt.Printf("Transaction Description: %s\n", response.AdditionalInfo.TransactionDescription)
		fmt.Printf("Transaction Status Date: %s\n", response.AdditionalInfo.TransactionStatusDate)
	}
}

// Example 5: Transaction History
func exampleTransactionHistory(ctx context.Context, client snap.Services) {
	fmt.Println("\n=== Example 5: Transaction History ===")

	// Create a history request
	request := &snap.HistoryListRequest{
		FromDateTime: "2024-12-01T00:00:00-07:00", // Start date
		ToDateTime:   "2024-12-30T00:00:00-07:00", // End date
		AdditionalInfo: &snap.AdditionalHistoryListRequest{
			AccountNo: "9920017573", // Account number
		},
	}

	// Get transaction history
	fmt.Println("Getting transaction history...")
	response, err := client.HistoryList(ctx, request)

	// Handle errors
	if err != nil {
		handleError(err)
		return
	}

	// Process the response
	fmt.Println("Transaction history retrieved successfully!")
	fmt.Printf("Response Code: %s\n", response.ResponseCode)
	fmt.Printf("Response Message: %s\n", response.ResponseMessage)
	fmt.Printf("Number of transactions: %d\n", len(response.DetailData))

	// Display transaction details
	for i, transaction := range response.DetailData {
		fmt.Printf("\nTransaction #%d:\n", i+1)
		fmt.Printf("  Date/Time: %s\n", transaction.DateTime)
		fmt.Printf("  Amount: %s %s\n", transaction.Amount.Value, transaction.Amount.Currency)
		fmt.Printf("  Remark: %s\n", transaction.Remark)
		fmt.Printf("  Status: %s\n", transaction.Status)
		fmt.Printf("  Type: %s\n", transaction.Type)

		if transaction.AdditionalInfo != nil {
			fmt.Printf("  Debit/Credit: %s\n", transaction.AdditionalInfo.DebitCredit)
		}

		if len(transaction.SourceOfFunds) > 0 {
			fmt.Println("  Source of Funds:")
			for j, source := range transaction.SourceOfFunds {
				fmt.Printf("    Source #%d: %s\n", j+1, source.Source)
			}
		}
	}

	if response.AdditionalInfo != nil {
		fmt.Printf("\nAdditional Info:\n")
		fmt.Printf("  Account No: %s\n", response.AdditionalInfo.AccountNo)
		fmt.Printf("  From Date/Time: %s\n", response.AdditionalInfo.FromDateTime)
		fmt.Printf("  To Date/Time: %s\n", response.AdditionalInfo.ToDateTime)
		fmt.Printf("  Message: %s\n", response.AdditionalInfo.Message)
	}
}

// Helper function to handle errors
func handleError(err error) {
	if snap.IsAuthenticationError(err) {
		log.Printf("Authentication error: %v", err)
	} else if snap.IsValidationError(err) {
		log.Printf("Validation error: %v", err)
	} else if snap.IsNotFoundError(err) {
		log.Printf("Not found error: %v", err)
	} else if snap.IsServerError(err) {
		log.Printf("Server error: %v", err)
	} else {
		log.Printf("Unknown error: %v", err)
	}
}
