package snap

import (
	"context"
	"fmt"
	"net/http"
)

// DisburseFunds initiates a fund disbursement to a bank account
func (c *Client) DisburseFunds(ctx context.Context, request DisbursementRequest) (*DisbursementResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodPost, EndpointDisbursements, request)
	if err != nil {
		return nil, err
	}

	var response struct {
		Response
		Data DisbursementResponse `json:"data"`
	}
	if err := c.parseResponse(resp, &response); err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, NewError(http.StatusBadRequest, response.Status, response.Message, response.Error)
	}

	return &response.Data, nil
}

// GetTransactionStatus retrieves the status of a transaction by ID or reference ID
func (c *Client) GetTransactionStatus(ctx context.Context, request TransactionStatusRequest) (*TransactionStatusResponse, error) {
	// Validate that at least one identifier is provided
	if request.TransactionID == "" && request.ReferenceID == "" {
		return nil, NewError(http.StatusBadRequest, ErrorCodeValidation, "either transaction_id or reference_id must be provided", "")
	}

	path := EndpointTransactionStatus
	if request.TransactionID != "" {
		path = fmt.Sprintf("%s?transaction_id=%s", path, request.TransactionID)
	} else {
		path = fmt.Sprintf("%s?reference_id=%s", path, request.ReferenceID)
	}

	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Response
		Data TransactionStatusResponse `json:"data"`
	}
	if err := c.parseResponse(resp, &response); err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, NewError(http.StatusBadRequest, response.Status, response.Message, response.Error)
	}

	return &response.Data, nil
}

// GetBalance retrieves the current account balance
func (c *Client) GetBalance(ctx context.Context) (*BalanceResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, EndpointBalance, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Response
		Data BalanceResponse `json:"data"`
	}
	if err := c.parseResponse(resp, &response); err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, NewError(http.StatusBadRequest, response.Status, response.Message, response.Error)
	}

	return &response.Data, nil
}

// GetBankList retrieves the list of supported banks
func (c *Client) GetBankList(ctx context.Context) (*BankListResponse, error) {
	resp, err := c.doRequest(ctx, http.MethodGet, EndpointBanks, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Response
		Data BankListResponse `json:"data"`
	}
	if err := c.parseResponse(resp, &response); err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, NewError(http.StatusBadRequest, response.Status, response.Message, response.Error)
	}

	return &response.Data, nil
}

// ListTransactions retrieves a list of transactions based on the provided filters
func (c *Client) ListTransactions(ctx context.Context, request TransactionListRequest) (*TransactionListResponse, error) {
	// Build query parameters
	path := EndpointTransactions
	queryParams := []string{}

	if request.StartDate != "" {
		queryParams = append(queryParams, fmt.Sprintf("start_date=%s", request.StartDate))
	}
	if request.EndDate != "" {
		queryParams = append(queryParams, fmt.Sprintf("end_date=%s", request.EndDate))
	}
	if request.Status != "" {
		queryParams = append(queryParams, fmt.Sprintf("status=%s", request.Status))
	}
	if request.Page > 0 {
		queryParams = append(queryParams, fmt.Sprintf("page=%d", request.Page))
	}
	if request.Limit > 0 {
		queryParams = append(queryParams, fmt.Sprintf("limit=%d", request.Limit))
	}

	// Add query parameters to path
	if len(queryParams) > 0 {
		path = fmt.Sprintf("%s?%s", path, queryParams[0])
		for i := 1; i < len(queryParams); i++ {
			path = fmt.Sprintf("%s&%s", path, queryParams[i])
		}
	}

	resp, err := c.doRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Response
		Data TransactionListResponse `json:"data"`
	}
	if err := c.parseResponse(resp, &response); err != nil {
		return nil, err
	}

	if response.Status != "success" {
		return nil, NewError(http.StatusBadRequest, response.Status, response.Message, response.Error)
	}

	return &response.Data, nil
}
