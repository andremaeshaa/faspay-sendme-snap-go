package snap

const (
	baseUrlSandbox = "https://account-dev.faspay.co.id"
	baseUrlProd    = "https://account-staging.faspay.co.id"
)

// API endpoint paths
const (
	EndpointTransferInterbank   = "/account/v1.0/transfer-interbank"
	EndpointAccountInquiry      = "/account/v1.0/account-inquiry-external"
	EndpointInquiryStatus       = "/account/v1.0/transfer/status"
	EndpointInquiryBalance      = "/account/v1.0/balance-inquiry"
	EndpointHistoryList         = "/account/v1.0/transaction-history-list"
	EndpointCustomerTopup       = "/account/v1.0/emoney/topup"
	EndpointCustomerTopupStatus = "/account/v1.0/emoney/topup-status"
	EndpointBillInquiry         = "/account/v1.0/transfer-va/inquiry-intrabank"
	EndpointBillPayment         = "/account/v1.0/transfer-va/payment-intrabank"
)

// Default configuration values
const (
	DefaultTimeout = 30                                 // Default timeout in seconds
	DefaultBaseURL = "https://account-dev.faspay.co.id" // Default API base URL
)
