package snap

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client represents a Faspay SendMe Snap API client
type Client struct {
	baseURL    string
	httpClient *http.Client
	apiKey     string
	apiSecret  string
	timeout    time.Duration
}

// ClientOption is a function that configures a Client
type ClientOption func(*Client)

// WithTimeout sets the timeout for the HTTP client
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.timeout = timeout
		c.httpClient.Timeout = timeout
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// NewClient creates a new Faspay SendMe Snap API client
func NewClient(baseURL, apiKey, apiSecret string, options ...ClientOption) *Client {
	// Use default base URL if none is provided
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	client := &Client{
		baseURL: baseURL,
		httpClient: &http.Client{
			Timeout: time.Duration(DefaultTimeout) * time.Second,
		},
		apiKey:    apiKey,
		apiSecret: apiSecret,
		timeout:   time.Duration(DefaultTimeout) * time.Second,
	}

	// Apply options
	for _, option := range options {
		option(client)
	}

	return client
}

// doRequest performs an HTTP request and returns the response
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	var reqBody io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", UserAgent())
	req.Header.Set("X-API-KEY", c.apiKey)
	// Add authentication headers or other required headers
	// This would depend on the specific authentication method required by the API

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	return resp, nil
}

// parseResponse parses the HTTP response into the provided response object
func (c *Client) parseResponse(resp *http.Response, v interface{}) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		// Try to parse error response
		var errorResp struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Error   string `json:"error"`
			Code    string `json:"code"`
			Details string `json:"details"`
		}

		if err := json.Unmarshal(body, &errorResp); err == nil {
			// Successfully parsed error response
			code := errorResp.Code
			if code == "" {
				code = errorResp.Status
			}

			message := errorResp.Message
			if message == "" {
				message = errorResp.Error
			}

			return NewError(resp.StatusCode, code, message, errorResp.Details)
		}

		// Couldn't parse error response, return generic error
		return NewError(resp.StatusCode, "unknown_error", fmt.Sprintf("API error: status code %d", resp.StatusCode), string(body))
	}

	if v != nil {
		if err := json.Unmarshal(body, v); err != nil {
			return fmt.Errorf("error unmarshaling response: %w", err)
		}
	}

	return nil
}
