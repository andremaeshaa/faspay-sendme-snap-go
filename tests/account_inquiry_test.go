package tests

import (
	"context"
	"faspay-sendme-snap-go/snap"
	"fmt"
	"testing"
	"time"
)

func TestTimeStamp(t *testing.T) {
	now := time.Now()
	formattedTime := now.Format("2006-01-02T15:04:05-07:00")
	fmt.Println(formattedTime)
}

func TestAccountInquiry(t *testing.T) {
	ctx := context.Background()

	client, err := snap.NewClient("99999", "20250607004236908", "../certs/enc.key", snap.WithTimeout(60*time.Second))
	if err != nil {
		panic(err)
	}

	request := &snap.ExternalAccountInquiryRequest{
		BeneficiaryBankCode:  "008",
		BeneficiaryAccountNo: "60004400184",
		PartnerReferenceNo:   "20250606234037371",
		AdditionalInfo: &snap.AdditionalInfoRequest{
			SourceAccount: "9920017573",
		},
	}

	response, err := client.AccountInquiry(ctx, request)
	if err != nil {
		panic(err)
	}

	println(response.ResponseMessage)
}
