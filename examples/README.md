# Faspay SendMe Snap SDK Examples

This directory contains examples demonstrating how to use the Faspay SendMe Snap SDK.

## Account Inquiry Example

The `example.go` file demonstrates how to use the SDK to perform an account inquiry, which is a common operation when integrating with Faspay's payment system.

### Prerequisites

- Go 1.23 or higher
- A Faspay SendMe Snap account with:
  - Partner ID (5-digit code)
  - External ID (36-character code)
  - Private key for signing requests (stored in the `certs` directory)

### Running the Example

To run the example, navigate to the examples directory and execute:

```bash
go run example.go
```

### Example Explanation

The example demonstrates:

1. **Loading the private key**: Reading the RSA private key from a file
2. **Initializing the client**: Creating a new client with your credentials
3. **Setting the environment**: Choosing between sandbox and production
4. **Creating a request**: Building an account inquiry request with the necessary parameters
5. **Making the API call**: Sending the request to Faspay's API
6. **Error handling**: Properly handling different types of errors
7. **Processing the response**: Extracting and displaying the response data

### Customizing the Example

To use this example with your own Faspay account:

1. Replace the Partner ID and External ID with your own credentials
2. Ensure your private key is correctly stored in the certs directory
3. Update the request parameters (bank code, account number, etc.) as needed
4. Change the environment to "prod" when you're ready to use the production API

## Additional Resources

For more information about the Faspay SendMe Snap API, refer to:

- The SDK documentation in the main README.md file
- Faspay's official API documentation
- The source code in the `snap` directory