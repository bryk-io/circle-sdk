package circlesdk

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/google/uuid"
)

const (
	endpointTesting    = "https://api-sandbox.circle.com/"
	endpointProduction = "https://api.circle.com/"
)

// Version identifier for the SDK.
const Version = "0.1.0"

// Client provices access to all core Circle APIs. This core set of APIs allow you to:
//   - Transfer digital currency (USDC) in and out of your Circle Account.
//   - Register your own business bank accounts - if you have them.
//   - Make transfers from / to your business bank account while seamlessly converting
//     those funds across digital currency and traditional FAIT.
type Client struct {
	// The Circle Payments API allows you to take payments from your end users
	// via traditional methods such as debit & credit cards, bank accounts, etc.,
	// and receive settlement in USDC. Businesses with users already holding USDC
	// are also able to take on-chain payments on supported blockchains.
	//
	// With the Circle Payments API you can:
	//   - Take card and bank transfer payments for goods or services.
	//   - Build a credit & debit card or bank transfer on-ramp for your crypto exchange.
	//   - Take card or bank transfer deposits for your savings, lending, investing or P2P
	//     payments product.
	//   - Take USDC payments directly through on-chain transfers.
	Payments *paymentsAPI

	// The Circle Payouts API allows you to issue payouts to your customers, vendors, or
	// suppliers in a variety of ways:
	//   - Bank wires
	//   - On-chain USDC transfers
	//   - ACH (coming soon)
	// Payouts are funded with your USDC denominated Circle Account, which can receive deposits
	// from both traditional and blockchain payment rails.
	Payouts *payoutsAPI

	// The Circle Accounts API allows you to easily create and manage accounts and balances
	// for your customers, and execute transfers of funds across accounts - whether they are
	// within the Circle platform, or in / out of the platform via on-chain USDC connectivity.
	//   - Embed US Dollar denominated accounts into your product or service without dealing
	//     with the complexity of legacy bank account structures.
	//   - Manage a multi-asset accounts infrastructure for your customers including seamless
	//     transfer of funds, across hosted accounts or via on-chain USDC connectivity.
	//   - Accept USDC deposits with minimum cost and no exposure to reversals.
	//   - Support BTC and ETH balances in addition to USDC.
	Accounts *accountsAPI

	// User agent value to report to the service.
	userAgent string

	// Time to maintain open the connection with the service, in seconds.
	keepAlive uint

	// Maximum network connections to keep open with the service.
	maxConnections uint

	// Network transport used to communicate with the service.
	conn *http.Client

	// Time to wait for requests, in seconds.
	timeout uint

	// Circle API key.
	key string

	// Produce trace output of requests and responses.
	debug bool

	// API backend to use.
	backend string
}

// Network request options.
type requestOptions struct {
	// HTTP method to use.
	method string

	// API operation endpoint.
	endpoint string

	// Operation parameters, if any
	input map[string]interface{}

	// Result holder, if any
	output interface{}

	// Produce idempotent results on POST requests; must be a valid UUID.
	// if none is provided a new one will be created by default.
	idempotencyKey string
}

// NewClient will construct a usable service handler using the provided API key and
// configuration options, if 'nil' options are provided default sane values will
// be used.
func NewClient(options ...Option) (*Client, error) {
	// New client instance
	cl := &Client{
		timeout:        30,
		keepAlive:      600,
		maxConnections: 50,
		userAgent:      "circlesdk-lib/" + Version,
		debug:          false,
		backend:        endpointTesting,
	}
	for _, opt := range options {
		if err := opt(cl); err != nil {
			return nil, err
		}
	}

	// Configure base HTTP transport
	t := &http.Transport{
		MaxIdleConns:        int(cl.maxConnections),
		MaxIdleConnsPerHost: int(cl.maxConnections),
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(cl.timeout) * time.Second,
			KeepAlive: time.Duration(cl.keepAlive) * time.Second,
		}).DialContext,
	}

	// Setup main client
	cl.conn = &http.Client{
		Transport: t,
		Timeout:   time.Duration(cl.timeout) * time.Second,
	}

	// Load API modules
	cl.Accounts = &accountsAPI{cl}
	cl.Payments = &paymentsAPI{cl}
	cl.Payouts = &payoutsAPI{cl}
	return cl, nil
}

// Ping will perform a basic reachability test. Use it to make sure your
// client instance is properly setup.
func (cl *Client) Ping() bool {
	type pingResponse struct {
		Message string `json:"message"`
	}

	req := &requestOptions{
		method:   "GET",
		endpoint: "ping",
		input:    nil,
		output:   &pingResponse{},
	}
	if err := cl.dispatch(req); err != nil {
		return false
	}
	return req.output.(*pingResponse).Message == "pong"
}

// Dispatch a network request to the service.
func (cl *Client) dispatch(r *requestOptions) error {
	// New idempotency key
	if r.idempotencyKey == "" {
		r.idempotencyKey = uuid.NewString()
	}

	req, _ := http.NewRequestWithContext(context.TODO(), r.method, cl.backend+r.endpoint, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+cl.key)
	req.Header.Add("Content-Type", "application/json")

	// Add user-agent header
	if cl.userAgent != "" {
		req.Header.Add("User-Agent", cl.userAgent)
	}

	// Dump request
	if cl.debug {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			fmt.Println("=== request ===")
			fmt.Printf("%s\n", dump)
		}
	}

	// Execute request
	res, err := cl.conn.Do(req)
	if res != nil {
		// Properly discard request content to be able to reuse the connection.
		defer func() {
			_, _ = io.Copy(ioutil.Discard, res.Body)
			_ = res.Body.Close()
		}()
	}

	// Dump response
	if cl.debug {
		dump, err := httputil.DumpResponse(res, true)
		if err == nil {
			fmt.Println("=== response ===")
			fmt.Printf("%s\n", dump)
		}
	}

	// Network level errors
	if err != nil {
		return err
	}

	// Get response contents
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Application level errors
	if res.StatusCode != 200 {
		return fmt.Errorf("unsuccessful request: %s", res.Status)
	}

	// Decode response
	if r.output != nil {
		if err = json.Unmarshal(body, r.output); err != nil {
			return fmt.Errorf("non JSON content returned: %s", body)
		}
	}

	// All good!
	return nil
}
