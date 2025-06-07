package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"faspay-sendme-snap-go/snap" // Import the snap package using the module path
)

// This example demonstrates how to use the Faspay SendMe Snap SDK to perform an account inquiry.
// It shows how to initialize the client, make a request, and handle the response and errors.
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
	partnerId := "99999"              // Your 5-digit partner ID
	externalId := "20250607004236909" // Your 36-character external ID

	// Create a new client with a custom timeout
	client, err := snap.NewClient(
		partnerId,
		externalId,
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

	// Step 4: Create an account inquiry request
	request := &snap.ExternalAccountInquiryRequest{
		BeneficiaryBankCode:  "008",               // Bank code (e.g., "008" for Mandiri)
		BeneficiaryAccountNo: "60004400184",       // Account number
		PartnerReferenceNo:   "20250606234037372", // Your unique reference number
		AdditionalInfo: &snap.AdditionalInfoInquiryAccount{
			SourceAccount: "9920017573", // Source account number
		},
	}

	// Step 5: Perform the account inquiry
	println("Performing account inquiry...")
	response, err := client.AccountInquiry(ctx, request)

	// Step 6: Handle errors
	if err != nil {
		// Use the error helper functions to determine the type of error
		if snap.IsAuthenticationError(err) {
			log.Fatalf("Authentication error: %v", err)
		} else if snap.IsValidationError(err) {
			log.Fatalf("Validation error: %v", err)
		} else if snap.IsNotFoundError(err) {
			log.Fatalf("Not found error: %v", err)
		} else if snap.IsServerError(err) {
			log.Fatalf("Server error: %v", err)
		} else {
			log.Fatalf("Unknown error: %v", err)
		}
	}

	// Step 7: Process the response
	println("Account inquiry successful!")
	println("Response Code: %s\n", response.ResponseCode)
	println("Response Message: %s\n", response.ResponseMessage)
	println("Reference No: %s\n", response.ReferenceNo)
	println("Partner Reference No: %s\n", response.PartnerReferenceNo)
	println("Beneficiary Account Name: %s\n", response.BeneficiaryAccountName)
	println("Beneficiary Account No: %s\n", response.BeneficiaryAccountNo)
	println("Beneficiary Bank Code: %s\n", response.BeneficiaryBankCode)
	println("Beneficiary Bank Name: %s\n", response.BeneficiaryBankName)
	println("Currency: %s\n", response.Currency)

	if response.AdditionalInfo != nil {
		fmt.Printf("Additional Info Status: %s\n", response.AdditionalInfo.Status)
		fmt.Printf("Additional Info Message: %s\n", response.AdditionalInfo.Message)
	}
}
