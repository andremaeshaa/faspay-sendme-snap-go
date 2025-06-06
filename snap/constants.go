package snap

const (
	baseUrlSandbox = "https://account-dev.faspay.co.id"
	baseUrlProd    = "https://account-staging.faspay.co.id"
)

// API endpoint paths
const (
	EndpointTransferInterbank = "/account/v1.0/transfer-interbank"
	EndpointAccountInquiry    = "/account/v1.0/account-inquiry-external"
)

// Default configuration values
const (
	DefaultTimeout = 30                                 // Default timeout in seconds
	DefaultBaseURL = "https://account-dev.faspay.co.id" // Default API base URL
)
