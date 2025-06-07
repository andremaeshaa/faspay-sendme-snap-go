package snap

import (
	"context"
	"net/http"
)

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
