package snap

type ExternalAccountInquiryRequest struct {
	BeneficiaryBankCode  string                 `json:"beneficiaryBankCode"`
	BeneficiaryAccountNo string                 `json:"beneficiaryAccountNo"`
	PartnerReferenceNo   string                 `json:"partnerReferenceNo"`
	AdditionalInfo       *AdditionalInfoRequest `json:"additionalInfo"`
}

type AdditionalInfoRequest struct {
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
