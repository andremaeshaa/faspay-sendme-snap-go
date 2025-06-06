package main

//func main() {
//	// Initialize the client with your API credentials
//	client := snap.NewClient(
//		snap.DefaultBaseURL,              // You can replace this with your custom base URL if needed
//		"your-api-key",                   // Replace with your actual API key
//		"your-api-secret",                // Replace with your actual API secret
//		snap.WithTimeout(60*time.Second), // Optional: Set custom timeout
//	)
//
//	// Create a context with timeout
//	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
//	defer cancel()
//
//	// Example 1: Get account balance
//	fmt.Println("Example 1: Get account balance")
//	balance, err := client.GetBalance(ctx)
//	if err != nil {
//		if snap.IsAuthenticationError(err) {
//			log.Fatalf("Authentication error: %v", err)
//		} else {
//			log.Fatalf("Error getting balance: %v", err)
//		}
//	}
//	fmt.Printf("Balance: %f %s\n\n", balance.Balance, balance.Currency)
//
//	// Example 2: Get list of supported banks
//	fmt.Println("Example 2: Get list of supported banks")
//	bankList, err := client.GetBankList(ctx)
//	if err != nil {
//		log.Fatalf("Error getting bank list: %v", err)
//	}
//	fmt.Printf("Found %d banks\n", len(bankList.Banks))
//	for i, bank := range bankList.Banks {
//		if i < 5 { // Print only first 5 banks
//			fmt.Printf("- %s (%s)\n", bank.Name, bank.Code)
//		}
//	}
//	fmt.Println()
//
//	// Example 3: Disburse funds
//	fmt.Println("Example 3: Disburse funds")
//	disbursementReq := snap.DisbursementRequest{
//		ReferenceID:   fmt.Sprintf("TRX-%d", time.Now().Unix()),
//		Amount:        1000000.0,
//		Currency:      snap.CurrencyIDR,
//		BankCode:      "BCA",
//		AccountName:   "John Doe",
//		AccountNumber: "1234567890",
//		Description:   "Payment for services",
//		Metadata: map[string]interface{}{
//			"customer_id": "CUST-123",
//			"order_id":    "ORDER-456",
//		},
//	}
//
//	disbursement, err := client.DisburseFunds(ctx, disbursementReq)
//	if err != nil {
//		if snap.IsValidationError(err) {
//			log.Fatalf("Validation error: %v", err)
//		} else {
//			log.Fatalf("Error disbursing funds: %v", err)
//		}
//	}
//	fmt.Printf("Disbursement successful! Transaction ID: %s, Status: %s\n\n",
//		disbursement.TransactionID, disbursement.Status)
//
//	// Example 4: Check transaction status
//	fmt.Println("Example 4: Check transaction status")
//	statusReq := snap.TransactionStatusRequest{
//		TransactionID: disbursement.TransactionID,
//	}
//
//	status, err := client.GetTransactionStatus(ctx, statusReq)
//	if err != nil {
//		if snap.IsNotFoundError(err) {
//			log.Fatalf("Transaction not found: %v", err)
//		} else {
//			log.Fatalf("Error checking transaction status: %v", err)
//		}
//	}
//	fmt.Printf("Transaction status: %s\n\n", status.Status)
//
//	// Example 5: List transactions
//	fmt.Println("Example 5: List transactions")
//	listReq := snap.TransactionListRequest{
//		StartDate: time.Now().AddDate(0, 0, -30).Format("2006-01-02"), // Last 30 days
//		EndDate:   time.Now().Format("2006-01-02"),
//		Status:    "success",
//		Page:      1,
//		Limit:     10,
//	}
//
//	transactions, err := client.ListTransactions(ctx, listReq)
//	if err != nil {
//		log.Fatalf("Error listing transactions: %v", err)
//	}
//	fmt.Printf("Found %d transactions\n", len(transactions.Transactions))
//	for i, tx := range transactions.Transactions {
//		if i < 5 { // Print only first 5 transactions
//			fmt.Printf("- %s: %f %s (Status: %s)\n",
//				tx.TransactionID, tx.Amount, snap.CurrencyIDR, tx.Status)
//		}
//	}
//}
