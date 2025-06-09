package snap

type ExternalAccountInquiryRequest struct {
	BeneficiaryBankCode  string                        `json:"beneficiaryBankCode"`
	BeneficiaryAccountNo string                        `json:"beneficiaryAccountNo"`
	PartnerReferenceNo   string                        `json:"partnerReferenceNo"`
	AdditionalInfo       *AdditionalInfoInquiryAccount `json:"additionalInfo"`
}

type TransferInterBankRequest struct {
	PartnerReferenceNo     string                           `json:"partnerReferenceNo"`
	Amount                 *Amount                          `json:"amount"`
	BeneficiaryAccountName string                           `json:"beneficiaryAccountName"`
	BeneficiaryAccountNo   string                           `json:"beneficiaryAccountNo"`
	BeneficiaryBankCode    string                           `json:"beneficiaryBankCode"`
	BeneficiaryEmail       string                           `json:"beneficiaryEmail"`
	SourceAccountNo        string                           `json:"sourceAccountNo"`
	TransactionDate        string                           `json:"transactionDate"`
	AdditionalInfo         *AdditionalInfoTransferInterBank `json:"additionalInfo"`
}

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type AvailableBalance struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type AdditionalInfoTransferInterBank struct {
	InstructDate           string `json:"instructDate"`
	TransactionDescription string `json:"transactionDescription"`
	CallbackUrl            string `json:"callbackUrl"`
}

type AdditionalInfoTransferInterBankResponse struct {
	BeneficiaryAccountName  string `json:"beneficiaryAccountName"`
	BeneficiaryBankName     string `json:"beneficiaryBankName"`
	InstructDate            string `json:"instructDate"`
	TransactionDescription  string `json:"transactionDescription"`
	CallbackUrl             string `json:"callbackUrl"`
	LatestTransactionStatus string `json:"latestTransactionStatus"`
	TransactionStatusDesc   string `json:"transactionStatusDesc"`
}

type TransferInterBankResponse struct {
	ResponseCode         string                                   `json:"responseCode"`
	ResponseMessage      string                                   `json:"responseMessage"`
	ReferenceNo          string                                   `json:"referenceNo"`
	PartnerReferenceNo   string                                   `json:"partnerReferenceNo"`
	Amount               *Amount                                  `json:"amount"`
	BeneficiaryAccountNo string                                   `json:"beneficiaryAccountNo"`
	BeneficiaryBankCode  string                                   `json:"beneficiaryBankCode"`
	SourceAccountNo      string                                   `json:"sourceAccountNo"`
	AdditionalInfo       *AdditionalInfoTransferInterBankResponse `json:"additionalInfo"`
}

type AdditionalInfoInquiryAccount struct {
	SourceAccount string `json:"sourceAccount"`
}

type AdditionalInfoResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ExternalAccountInquiryResponse struct {
	ResponseCode           string                  `json:"responseCode,omitempty"`
	ResponseMessage        string                  `json:"responseMessage,omitempty"`
	ReferenceNo            string                  `json:"referenceNo,omitempty"`
	PartnerReferenceNo     string                  `json:"partnerReferenceNo,omitempty"`
	BeneficiaryAccountName string                  `json:"beneficiaryAccountName,omitempty"`
	BeneficiaryAccountNo   string                  `json:"beneficiaryAccountNo,omitempty"`
	BeneficiaryBankCode    string                  `json:"beneficiaryBankCode,omitempty"`
	BeneficiaryBankName    string                  `json:"beneficiaryBankName,omitempty"`
	Currency               string                  `json:"currency,omitempty"`
	AdditionalInfo         *AdditionalInfoResponse `json:"additionalInfo,omitempty"`
}

type StatusTransferRequest struct {
	OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
	OriginalReferenceNo        string `json:"originalReferenceNo"`
	ServiceCode                string `json:"serviceCode"`
}

type StatusTransferResponse struct {
	ResponseCode               string                                `json:"responseCode"`
	ResponseMessage            string                                `json:"responseMessage"`
	OriginalReferenceNo        string                                `json:"originalReferenceNo"`
	OriginalPartnerReferenceNo string                                `json:"originalPartnerReferenceNo"`
	ServiceCode                string                                `json:"serviceCode"`
	TransactionDate            string                                `json:"transactionDate"`
	Amount                     *Amount                               `json:"amount"`
	BeneficiaryAccountNo       string                                `json:"beneficiaryAccountNo"`
	BeneficiaryBankCode        string                                `json:"beneficiaryBankCode"`
	ReferenceNumber            string                                `json:"referenceNumber"`
	SourceAccountNo            string                                `json:"sourceAccountNo"`
	LatestTransactionStatus    string                                `json:"latestTransactionStatus"`
	TransactionStatusDesc      string                                `json:"transactionStatusDesc"`
	AdditionalInfo             *AdditionalInfoStatusTransferResponse `json:"additionalInfo"`
}

type AdditionalInfoStatusTransferResponse struct {
	BeneficiaryAccountName string `json:"beneficiaryAccountName"`
	BeneficiaryBankName    string `json:"beneficiaryBankName"`
	TransactionDescription string `json:"transactionDescription"`
	CallbackUrl            string `json:"callbackUrl"`
	TransactionStatusDate  string `json:"transactionStatusDate"`
}

type InquiryBalanceRequest struct {
	AccountNo string `json:"accountNo"`
}

type InquiryBalanceResponse struct {
	ResponseCode    string          `json:"responseCode"`
	ResponseMessage string          `json:"responseMessage"`
	AccountNo       string          `json:"accountNo"`
	AccountInfos    []*AccountInfos `json:"accountInfos"`
}

type AccountInfos struct {
	BalanceType      string            `json:"balanceType"`
	Amount           *Amount           `json:"amount"`
	AvailableBalance *AvailableBalance `json:"availableBalance"`
	Status           string            `json:"status"`
}

type HistoryListRequest struct {
	FromDateTime   string                        `json:"fromDateTime"`
	ToDateTime     string                        `json:"toDateTime"`
	AdditionalInfo *AdditionalHistoryListRequest `json:"additionalInfo"`
}

type AdditionalHistoryListRequest struct {
	AccountNo string `json:"accountNo"`
}

type HistoryListResponse struct {
	ResponseCode    string                             `json:"responseCode"`
	ResponseMessage string                             `json:"responseMessage"`
	DetailData      []*DetailData                      `json:"detailData"`
	AdditionalInfo  *AdditionalInfoHistoryListResponse `json:"additionalInfo"`
}

type SourceOfFunds struct {
	Source string `json:"source"`
}

type AdditionalInfoDetailData struct {
	DebitCredit string `json:"debitCredit"`
}

type AdditionalInfoHistoryListResponse struct {
	AccountNo    string `json:"accountNo"`
	FromDateTime string `json:"fromDateTime"`
	ToDateTime   string `json:"toDateTime"`
	Message      string `json:"message"`
}

type DetailData struct {
	DateTime       string                    `json:"dateTime"`
	Amount         *Amount                   `json:"amount"`
	Remark         string                    `json:"remark"`
	SourceOfFunds  []*SourceOfFunds          `json:"sourceOfFunds"`
	Status         string                    `json:"status"`
	Type           string                    `json:"type"`
	AdditionalInfo *AdditionalInfoDetailData `json:"additionalInfo"`
}

type CustomerTopupRequest struct {
	PartnerReferenceNo string                              `json:"partnerReferenceNo"`
	CustomerNumber     string                              `json:"customerNumber"`
	Amount             *Amount                             `json:"amount"`
	TransactionDate    string                              `json:"transactionDate"`
	AdditionalInfo     *AdditionalInfoCustomerTopupRequest `json:"additionalInfo"`
}

type AdditionalInfoCustomerTopupRequest struct {
	SourceAccount          string `json:"sourceAccount"`
	PlatformCode           string `json:"platformCode"`
	InstructDate           string `json:"instructDate"`
	BeneficiaryEmail       string `json:"beneficiaryEmail"`
	TransactionDescription string `json:"transactionDescription"`
	CallbackUrl            string `json:"callbackUrl"`
}

type AdditionalInfoCustomerTopup struct {
	SourceAccount           string `json:"sourceAccount"`
	PlatformCode            string `json:"platformCode"`
	BeneficiaryEmail        string `json:"beneficiaryEmail"`
	TransactionDate         string `json:"transactionDate"`
	InstructDate            string `json:"instructDate"`
	TransactionDescription  string `json:"transactionDescription"`
	CallbackUrl             string `json:"callbackUrl"`
	TransactionReference    string `json:"transactionReference"`
	LatestTransactionStatus string `json:"latestTransactionStatus"`
	TransactionStatusDesc   string `json:"transactionStatusDesc"`
}

type CustomerTopupResponse struct {
	ResponseCode       string                       `json:"responseCode"`
	ResponseMessage    string                       `json:"responseMessage"`
	ReferenceNo        string                       `json:"referenceNo"`
	PartnerReferenceNo string                       `json:"partnerReferenceNo"`
	CustomerNumber     string                       `json:"customerNumber"`
	Amount             *Amount                      `json:"amount"`
	AdditionalInfo     *AdditionalInfoCustomerTopup `json:"additionalInfo"`
}
