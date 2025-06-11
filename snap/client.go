package snap

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"math"
	mathRand "math/rand"
	"net/http"
	"strings"
	"time"
)

// Client represents a Faspay SendMe Snap API client
type Client struct {
	environment string
	baseURL     string
	httpClient  *http.Client
	PartnerId   string
	privateKey  []byte
	timeout     time.Duration
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
func NewClient(partnerId string, privateKey, sslCert []byte, options ...ClientOption) (Services, error) {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(sslCert)

	client := &Client{
		httpClient: &http.Client{
			Timeout: time.Duration(DefaultTimeout) * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
		},
		PartnerId:  partnerId,
		privateKey: privateKey,
		timeout:    time.Duration(DefaultTimeout) * time.Second,
	}

	if client.baseURL == "" {
		client.baseURL = DefaultBaseURL
		client.environment = "sandbox"
	}

	// Apply options
	for _, option := range options {
		option(client)
	}

	return client, nil
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

	signature, err := c.generateSignatureSnap(method, path, string(jsonBody), timestamp, c.privateKey)
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
	req.Header.Set("X-EXTERNAL-ID", c.generateRandomNumber())
	req.Header.Set("CHANNEL-ID", "88001")

	if c.environment == "sandbox" {
		println("Info: Transaction will be processed in sandbox mode")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}

	return resp, nil
}

func (c *Client) generateSignatureSnap(httpMethod, endpointUrl, requestBody, timeStamp string, privateKeyPEM []byte) (string, error) {
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

func (c *Client) generateRandomNumber() string {
	// Get current time in milliseconds (equivalent to microtime(true) * 1000)
	milliseconds := math.Round(float64(time.Now().UnixNano()) / 1e6)

	// Generate random number (assuming $rand is a random number)
	randomNum := mathRand.Intn(1000) // Adjust range as needed

	// Concatenate all parts (equivalent to '99999' . round(...) . $rand)
	return fmt.Sprintf("%s%d%d", c.PartnerId, int64(milliseconds), randomNum)
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
