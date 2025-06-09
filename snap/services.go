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
	if c.environment == "sandbox" {
		println("Info: Transaction will be processed in sandbox mode")
	}

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
	if c.environment == "sandbox" {
		println("Info: Transaction will be processed in sandbox mode")
	}

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
