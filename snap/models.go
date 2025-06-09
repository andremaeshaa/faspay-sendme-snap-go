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
