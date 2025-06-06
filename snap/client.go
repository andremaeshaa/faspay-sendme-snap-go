package snap

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// Client represents a Faspay SendMe Snap API client
type Client struct {
	baseURL          string
	httpClient       *http.Client
	PartnerId        string `validate:"required,len=5"`
	ExternalId       string `validate:"required,len=36"`
	privateKeyPath   string
	privateKeyString string
	timeout          time.Duration
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

// NewClient initializes and returns a new Client instance with the given API key, secret, and optional configurations.
func NewClient(partnerId, externalId string, privateKeyPath string, options ...ClientOption) (*Client, error) {
	client := &Client{
		httpClient: &http.Client{
			Timeout: time.Duration(DefaultTimeout) * time.Second,
		},
		PartnerId:      partnerId,
		ExternalId:     externalId,
		privateKeyPath: privateKeyPath,
		timeout:        time.Duration(DefaultTimeout) * time.Second,
	}

	if client.baseURL == "" {
		client.baseURL = DefaultBaseURL
	}

	err := client.setPrivateKey(privateKeyPath)
	if err != nil {
		return nil, err
	}

	// Apply options
	for _, option := range options {
		option(client)
	}

	return client, nil
}

// SetEnv sets the environment for the client, switching the base URL between "sandbox" and "prod" environments.
func (c *Client) SetEnv(envType string) error {
	if envType == "sandbox" {
		c.baseURL = baseUrlSandbox
	} else if envType == "prod" {
		c.baseURL = baseUrlProd
	} else {
		return fmt.Errorf("invalid env type")
	}
	return nil
}

func (c *Client) setPrivateKey(pathStr string) error {
	bytesPrivateKey, err := os.ReadFile(pathStr)
	if err != nil {
		return err
	}

	c.privateKeyString = string(bytesPrivateKey)

	return nil
}

// doRequest performs an HTTP request with the specified method, URL path, and request body, returning the HTTP response.
func (c *Client) doRequest(ctx context.Context, method, path string, body interface{}) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	var jsonBody []byte
	var reqBody io.Reader
	if body != nil {
		var err error
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error marshaling request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Generate timestamp for signature
	timestamp := time.Now().Format("2006-01-02T15:04:05-07:00")

	signature, err := GenerateSignatureSnap(method, path, string(jsonBody), timestamp, c.privateKeyString)
	if err != nil {
		return nil, fmt.Errorf("error generating signature: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", UserAgent())
	req.Header.Set("X-TIMESTAMP", timestamp)
	req.Header.Set("X-SIGNATURE", signature)
	req.Header.Set("X-PARTNER-ID", c.PartnerId)
	req.Header.Set("X-EXTERNAL-ID", c.ExternalId)
	req.Header.Set("CHANNEL-ID", "88001")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	return resp, nil
}

func GenerateSignatureSnap(httpMethod, endpointUrl, requestBody, timeStamp, privateKeyPEM string) (string, error) {
	// Remove escaped slashes (\/ → /)
	minifiedBody := strings.ReplaceAll(requestBody, `\/`, `/`)

	// SHA-256 hash of minified body
	hashed := sha256.Sum256([]byte(minifiedBody))
	lowercaseHash := fmt.Sprintf("%x", hashed[:])

	// Build string to sign
	stringToSign := fmt.Sprintf("%s:%s:%s:%s", httpMethod, endpointUrl, lowercaseHash, timeStamp)

	// Parse private key from PEM format
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil {
		return "", errors.New("failed to parse private key PEM")
	}

	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		// Try PKCS1 if PKCS8 fails
		parsedKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", errors.New("failed to parse private key")
		}
	}

	rsaKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("not an RSA private key")
	}

	// Sign using SHA256withRSA
	hash := sha256.New()
	hash.Write([]byte(stringToSign))
	hashedBytes := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, rsaKey, crypto.SHA256, hashedBytes)
	if err != nil {
		return "", fmt.Errorf("failed to sign: %v", err)
	}

	// Encode to base64
	encodedSignature := base64.StdEncoding.EncodeToString(signature)
	return encodedSignature, nil
}

// parseResponse parses the HTTP response into the provided response object
func (c *Client) parseResponse(resp *http.Response, v any) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %w", err)
	}
	println("resp: ", string(body))

	if v != nil {
		if err := json.Unmarshal(body, v); err != nil {
			return fmt.Errorf("error unmarshaling response: %w", err)
		}
	}

	return nil
}
