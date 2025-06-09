package snap

import (
	"context"
	"os"
	"testing"
)

func TestClient_AccountInquiry(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := client.AccountInquiry(context.Background(), &ExternalAccountInquiryRequest{
		BeneficiaryBankCode:  "008",
		BeneficiaryAccountNo: "60004400184",
		PartnerReferenceNo:   "20250606234037372",
		AdditionalInfo: &AdditionalInfoInquiryAccount{
			SourceAccount: "9920017573",
		},
	})
	if err != nil {
		panic(err)
	}

	println(res.ResponseMessage)
}

func TestClient_TransferInterBank(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := client.TransferInterBank(context.Background(), &TransferInterBankRequest{
		PartnerReferenceNo: "20250609103003235",
		Amount: &Amount{
			Value:    "59614.00",
			Currency: "IDR",
		},
		BeneficiaryAccountName: "GolangTestAjoji Ajojo",
		BeneficiaryAccountNo:   "60004400184",
		BeneficiaryBankCode:    "008",
		BeneficiaryEmail:       "aan28setiawan@gmail.com",
		SourceAccountNo:        "9920017573",
		TransactionDate:        "2025-06-09T10:30:03+07:00",
		AdditionalInfo: &AdditionalInfoTransferInterBank{
			InstructDate:           "",
			TransactionDescription: "snapmandiri20250609103003",
			CallbackUrl:            "http://account-service/account/api/mail/sendtotele",
		},
	})
	if err != nil {
		panic(err)
	}

	println(res.ResponseMessage)
}
