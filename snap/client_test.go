package snap

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	// Test with default options
	client := NewClient("https://api.example.com", "test-key", "test-secret")
	if client.baseURL != "https://api.example.com" {
		t.Errorf("Expected baseURL to be %s, got %s", "https://api.example.com", client.baseURL)
	}
	if client.apiKey != "test-key" {
		t.Errorf("Expected apiKey to be %s, got %s", "test-key", client.apiKey)
	}
	if client.apiSecret != "test-secret" {
		t.Errorf("Expected apiSecret to be %s, got %s", "test-secret", client.apiSecret)
	}
	if client.timeout != time.Duration(DefaultTimeout)*time.Second {
		t.Errorf("Expected timeout to be %s, got %s", time.Duration(DefaultTimeout)*time.Second, client.timeout)
	}

	// Test with custom timeout
	customTimeout := 60 * time.Second
	client = NewClient("https://api.example.com", "test-key", "test-secret", WithTimeout(customTimeout))
	if client.timeout != customTimeout {
		t.Errorf("Expected timeout to be %s, got %s", customTimeout, client.timeout)
	}

	// Test with custom HTTP client
	customHTTPClient := &http.Client{Timeout: 45 * time.Second}
	client = NewClient("https://api.example.com", "test-key", "test-secret", WithHTTPClient(customHTTPClient))
	if client.httpClient != customHTTPClient {
		t.Errorf("Expected httpClient to be %v, got %v", customHTTPClient, client.httpClient)
	}

	// Test with empty baseURL (should use default)
	client = NewClient("", "test-key", "test-secret")
	if client.baseURL != DefaultBaseURL {
		t.Errorf("Expected baseURL to be %s, got %s", DefaultBaseURL, client.baseURL)
	}
}

func TestDoRequest(t *testing.T) {
	// Create a test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check request headers
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type header to be application/json, got %s", r.Header.Get("Content-Type"))
		}
		if r.Header.Get("Accept") != "application/json" {
			t.Errorf("Expected Accept header to be application/json, got %s", r.Header.Get("Accept"))
		}
		if r.Header.Get("User-Agent") != UserAgent() {
			t.Errorf("Expected User-Agent header to be %s, got %s", UserAgent(), r.Header.Get("User-Agent"))
		}
		if r.Header.Get("X-API-KEY") != "test-key" {
			t.Errorf("Expected X-API-KEY header to be test-key, got %s", r.Header.Get("X-API-KEY"))
		}

		// Return a success response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"success","message":"OK","data":{"balance":1000,"currency":"IDR"}}`))
	}))
	defer server.Close()

	// Create a client using the test server URL
	client := NewClient(server.URL, "test-key", "test-secret")

	// Make a request
	resp, err := client.doRequest(context.Background(), http.MethodGet, "/test", nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Test parsing the response
	var response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Data    struct {
			Balance  float64 `json:"balance"`
			Currency string  `json:"currency"`
		} `json:"data"`
	}
	err = client.parseResponse(resp, &response)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if response.Status != "success" {
		t.Errorf("Expected status to be success, got %s", response.Status)
	}
	if response.Data.Balance != 1000 {
		t.Errorf("Expected balance to be 1000, got %f", response.Data.Balance)
	}
	if response.Data.Currency != "IDR" {
		t.Errorf("Expected currency to be IDR, got %s", response.Data.Currency)
	}
}
