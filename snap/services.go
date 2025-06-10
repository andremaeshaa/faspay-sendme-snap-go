package snap

import (
	"context"
	"fmt"
	"net/http"
)

type Services interface {
	SetEnv(envType string) error
	AccountInquiry(ctx context.Context, request *ExternalAccountInquiryRequest) (*ExternalAccountInquiryResponse, error)
	TransferInterBank(ctx context.Context, request *TransferInterBankRequest) (*TransferInterBankResponse, error)
	StatusTransfer(ctx context.Context, request *StatusTransferRequest) (*StatusTransferResponse, error)
	InquiryBalance(ctx context.Context, request *InquiryBalanceRequest) (*InquiryBalanceResponse, error)
	HistoryList(ctx context.Context, request *HistoryListRequest) (*HistoryListResponse, error)
	CustomerTopup(ctx context.Context, request *CustomerTopupRequest) (*CustomerTopupResponse, error)
	CustomerTopupStatus(ctx context.Context, request *CustomerTopupStatusRequest) (*CustomerTopupStatusResponse, error)
	BillInquiry(ctx context.Context, request *BillInquiryRequest) (*BillInquiryResponse, error)
	BillPayment(ctx context.Context, request *BillPaymentRequest) (*BillPaymentResponse, error)
}

// SetEnv sets the environment for the client, switching the base URL between "sandbox" and "prod" environments.
func (c *Client) SetEnv(envType string) error {
	if envType == "sandbox" {
		c.baseURL = baseUrlSandbox
		c.environment = "sandbox"
	} else if envType == "prod" {
		c.baseURL = baseUrlProd
		c.environment = "prod"
	} else {
		return fmt.Errorf("invalid env type")
	}
	return nil
}

// AccountInquiry performs an inquiry for external account details
func (c *Client) AccountInquiry(ctx context.Context, request *ExternalAccountInquiryRequest) (*ExternalAccountInquiryResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointAccountInquiry, request)
	if err != nil {
		return nil, err
	}

	// The API directly returns the account inquiry response without a wrapper
	var response ExternalAccountInquiryResponse

	if err := c.parseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) TransferInterBank(ctx context.Context, request *TransferInterBankRequest) (*TransferInterBankResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointTransferInterbank, request)
	if err != nil {
		return nil, err
	}

	var response TransferInterBankResponse

	if err := c.parseResponse(resp, &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) StatusTransfer(ctx context.Context, request *StatusTransferRequest) (*StatusTransferResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointInquiryStatus, request)
	if err != nil {
		return nil, err
	}

	var response StatusTransferResponse

	err = c.parseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) InquiryBalance(ctx context.Context, request *InquiryBalanceRequest) (*InquiryBalanceResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointInquiryBalance, request)
	if err != nil {
		return nil, err
	}

	var response InquiryBalanceResponse

	err = c.parseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) HistoryList(ctx context.Context, request *HistoryListRequest) (*HistoryListResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointHistoryList, request)
	if err != nil {
		return nil, err
	}

	var response HistoryListResponse

	err = c.parseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CustomerTopup(ctx context.Context, request *CustomerTopupRequest) (*CustomerTopupResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointCustomerTopup, request)
	if err != nil {
		return nil, err
	}

	var response CustomerTopupResponse

	err = c.parseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) CustomerTopupStatus(ctx context.Context, request *CustomerTopupStatusRequest) (*CustomerTopupStatusResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointCustomerTopupStatus, request)
	if err != nil {
		return nil, err
	}

	var response CustomerTopupStatusResponse

	err = c.parseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) BillInquiry(ctx context.Context, request *BillInquiryRequest) (*BillInquiryResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointBillInquiry, request)
	if err != nil {
		return nil, err
	}

	var response BillInquiryResponse

	err = c.parseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (c *Client) BillPayment(ctx context.Context, request *BillPaymentRequest) (*BillPaymentResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointBillPayment, request)
	if err != nil {
		return nil, err
	}

	var response BillPaymentResponse

	err = c.parseResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
