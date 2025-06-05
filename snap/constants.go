package snap

// API endpoint paths
const (
	// Disbursement endpoints
	EndpointDisbursements = "/api/v1/disbursements"

	// Transaction endpoints
	EndpointTransactionStatus = "/api/v1/transactions/status"
	EndpointTransactions      = "/api/v1/transactions"

	// Balance endpoints
	EndpointBalance = "/api/v1/balance"

	// Bank endpoints
	EndpointBanks = "/api/v1/banks"
)

// Transaction status values
const (
	StatusPending    = "pending"
	StatusProcessing = "processing"
	StatusSuccess    = "success"
	StatusFailed     = "failed"
	StatusCancelled  = "cancelled"
	StatusRefunded   = "refunded"
)

// Currency codes
const (
	CurrencyIDR = "IDR" // Indonesian Rupiah
	CurrencyUSD = "USD" // US Dollar
	CurrencySGD = "SGD" // Singapore Dollar
)

// Error codes
const (
	ErrorCodeValidation         = "validation_error"
	ErrorCodeAuthentication     = "authentication_error"
	ErrorCodeAuthorization      = "authorization_error"
	ErrorCodeNotFound           = "not_found"
	ErrorCodeRateLimit          = "rate_limit_exceeded"
	ErrorCodeInternalServer     = "internal_server_error"
	ErrorCodeServiceUnavailable = "service_unavailable"
	ErrorCodeInsufficientFunds  = "insufficient_funds"
)

// Default configuration values
const (
	DefaultTimeout = 30                         // Default timeout in seconds
	DefaultBaseURL = "https://api.faspay.co.id" // Default API base URL
)
