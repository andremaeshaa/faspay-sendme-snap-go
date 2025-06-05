package snap

// Response is the base response structure for all API responses
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// DisbursementRequest represents a request to disburse funds
type DisbursementRequest struct {
	ReferenceID   string                 `json:"reference_id"`
	Amount        float64                `json:"amount"`
	Currency      string                 `json:"currency"`
	BankCode      string                 `json:"bank_code"`
	AccountName   string                 `json:"account_name"`
	AccountNumber string                 `json:"account_number"`
	Description   string                 `json:"description,omitempty"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// DisbursementResponse represents the response from a disbursement request
type DisbursementResponse struct {
	TransactionID string                 `json:"transaction_id"`
	ReferenceID   string                 `json:"reference_id"`
	Status        string                 `json:"status"`
	Amount        float64                `json:"amount"`
	Fee           float64                `json:"fee"`
	CreatedAt     string                 `json:"created_at"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
}

// TransactionStatusRequest represents a request to check transaction status
type TransactionStatusRequest struct {
	TransactionID string `json:"transaction_id,omitempty"`
	ReferenceID   string `json:"reference_id,omitempty"`
}

// TransactionStatusResponse represents the response from a transaction status request
type TransactionStatusResponse struct {
	TransactionID string                 `json:"transaction_id"`
	ReferenceID   string                 `json:"reference_id"`
	Status        string                 `json:"status"`
	Amount        float64                `json:"amount"`
	Fee           float64                `json:"fee"`
	CreatedAt     string                 `json:"created_at"`
	UpdatedAt     string                 `json:"updated_at"`
	Metadata      map[string]interface{} `json:"metadata,omitempty"`
	FailureReason string                 `json:"failure_reason,omitempty"`
}

// BalanceResponse represents the response from a balance inquiry
type BalanceResponse struct {
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

// BankListResponse represents the response from a bank list request
type BankListResponse struct {
	Banks []Bank `json:"banks"`
}

// Bank represents a bank in the system
type Bank struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	CountryCode string `json:"country_code"`
	IsActive    bool   `json:"is_active"`
}

// TransactionListRequest represents a request to list transactions
type TransactionListRequest struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	Status    string `json:"status,omitempty"`
	Page      int    `json:"page,omitempty"`
	Limit     int    `json:"limit,omitempty"`
}

// TransactionListResponse represents the response from a transaction list request
type TransactionListResponse struct {
	Transactions []TransactionStatusResponse `json:"transactions"`
	Pagination   Pagination                  `json:"pagination"`
}

// Pagination represents pagination information
type Pagination struct {
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalItems   int `json:"total_items"`
	ItemsPerPage int `json:"items_per_page"`
}
