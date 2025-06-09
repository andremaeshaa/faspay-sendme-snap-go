package snap

import (
	"context"
	"encoding/json"
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

func TestClient_InquiryStatus(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := client.StatusTransfer(context.Background(), &StatusTransferRequest{
		OriginalPartnerReferenceNo: "20250609103003234",
		OriginalReferenceNo:        "53883",
		ServiceCode:                "18",
	})
	if err != nil {
		panic(err)
	}

	println(res.ResponseMessage)
}

func TestClient_InquiryBalance(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := client.InquiryBalance(context.Background(), &InquiryBalanceRequest{
		AccountNo: "9920017573",
	})
	if err != nil {
		panic(err)
	}

	println(res.ResponseMessage)
}

func TestClient_HistoryList(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	res, err := client.HistoryList(context.Background(), &HistoryListRequest{
		FromDateTime:   "2024-12-01T00:00:00-07:00",
		ToDateTime:     "2024-12-30T00:00:00-07:00",
		AdditionalInfo: &AdditionalHistoryListRequest{AccountNo: "9920017573"},
	})
	if err != nil {
		panic(err)
	}

	bytes, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	println("response: ", string(bytes))
}
