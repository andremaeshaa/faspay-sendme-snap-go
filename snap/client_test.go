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

func TestClient_CustomerTopup(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc_stg.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	err = client.SetEnv("prod")
	if err != nil {
		panic(err)
	}

	res, err := client.CustomerTopup(context.Background(), &CustomerTopupRequest{
		PartnerReferenceNo: "20250609150352617",
		CustomerNumber:     "0812254830",
		Amount: &Amount{
			Value:    "76860.00",
			Currency: "IDR",
		},
		TransactionDate: "2025-06-09T15:03:52+07:00",
		AdditionalInfo: &AdditionalInfoCustomerTopupRequest{
			SourceAccount:          "9920017573",
			PlatformCode:           "gpy",
			InstructDate:           "",
			BeneficiaryEmail:       "aanfaspay2022@gmail.com,aan28setiawan@gmail.com",
			TransactionDescription: "Tunjangan Pulsa 20250609",
			CallbackUrl:            "https://245e-103-83-94-10.ngrok-free.app/v1/snap/callback",
		},
	})
	if err != nil {
		panic(err)
	}

	println(res.ResponseMessage)
}

func TestClient_CustomerTopupStatus(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc_stg.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	err = client.SetEnv("prod")
	if err != nil {
		panic(err)
	}

	response, err := client.CustomerTopupStatus(context.Background(), &CustomerTopupStatusRequest{
		OriginalPartnerReferenceNo: "20250609150352616",
		OriginalReferenceNo:        "59732",
		ServiceCode:                "38",
	})
	if err != nil {
		panic(err)
	}

	println(response.ResponseMessage)
}

func TestClient_BillInquiry(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc_stg.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	err = client.SetEnv("prod")
	if err != nil {
		panic(err)
	}

	res, err := client.BillInquiry(context.Background(), &BillInquiryRequest{
		PartnerReferenceNo: "20250609162756943",
		PartnerServiceId:   "    7008",
		CustomerNo:         "08000047816",
		VirtualAccountNo:   "700808000047816",
		AdditionalInfo: &AdditionalInfoBillInquiry{
			BillerCode:    "013",
			SourceAccount: "9920017573",
		},
	})
	if err != nil {
		panic(err)
	}

	println(res.ResponseMessage)
}

func TestClient_BillPayment(t *testing.T) {
	privateKey, err := os.ReadFile("../certs/enc_stg.key")
	if err != nil {
		t.Fatalf("Failed to read private key: %v", err)
	}

	client, err := NewClient("99999", privateKey)
	if err != nil {
		panic(err)
	}

	err = client.SetEnv("prod")
	if err != nil {
		panic(err)
	}

	response, err := client.BillPayment(context.Background(), &BillPaymentRequest{
		PartnerReferenceNo: "20250609162921210",
		PartnerServiceId:   "    7008",
		CustomerNo:         "08000047816",
		VirtualAccountNo:   "700808000047816",
		VirtualAccountName: "DUMMY VA",
		SourceAccount:      "9920017573",
		PaidAmount: &Amount{
			Value:    "41454.00",
			Currency: "IDR",
		},
		TrxDateTime: "2025-06-09T16:29:21",
		AdditionalInfo: &AdditionalInfoBillPayment{
			BillerCode:   "013",
			InstructDate: "2025-06-09T16:29:21+07:00",
			CallbackUrl:  "https://245e-103-83-94-10.ngrok-free.app/v1/snap/callback",
		},
	})
	if err != nil {
		panic(err)
	}

	println(response.ResponseMessage)
}
