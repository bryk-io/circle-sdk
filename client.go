package circlesdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	endpointTesting    = "https://api-sandbox.circle.com/"
	endpointProduction = "https://api.circle.com/"
)

// Version identifier for the SDK.
const Version = "0.1.0"

// RequestOptions represents network request options.
type RequestOptions struct {
	// HTTP method to use.
	Method string

	// API operation endpoint.
	Endpoint string

	// Operation parameters, if any.
	Input interface{}

	// Result holder, if any.
	Output interface{}

	// Automatically unwrap the "data" key in the response to the output
	// holder provided.
	UnwrapData bool

	// Produce idempotent results on POST requests; must be a valid UUID.
	// if none is provided a new one will be created by default.
	IdempotencyKey string

	// Custom request context.
	Ctx context.Context

	// Custom query parameters.
	QueryParams url.Values
}

// AddQueryParam register a new query parameter.
func (req *RequestOptions) AddQueryParam(key, value string) {
	if req.QueryParams == nil {
		req.QueryParams = url.Values{}
	}
	req.QueryParams.Add(key, value)
}

// Client contains the properties of the connection.
type Client struct {
	// User agent value to report to the service.
	UserAgent string

	// Time to maintain open the connection with the service, in seconds.
	KeepAlive uint

	// Maximum network connections to keep open with the service.
	MaxConnections uint

	// Network transport used to communicate with the service.
	Conn *http.Client

	// Time to wait for requests, in seconds.
	Timeout uint

	// Circle API key.
	Key string

	// Produce trace output of requests and responses.
	Debug bool

	// API backend to use.
	Backend string
}

// NewClient will construct a usable service handler using the provided API key and
// configuration options, if 'nil' options are provided default sane values will
// be used.
func NewClient(options ...Option) (*Client, error) {
	// New client instance
	cl := &Client{
		Timeout:        30,
		KeepAlive:      600,
		MaxConnections: 50,
		UserAgent:      "circlesdk-lib/" + Version,
		Debug:          false,
		Backend:        endpointTesting,
	}
	for _, opt := range options {
		if err := opt(cl); err != nil {
			return nil, err
		}
	}

	// Configure base HTTP transport
	t := &http.Transport{
		MaxIdleConns:        int(cl.MaxConnections),
		MaxIdleConnsPerHost: int(cl.MaxConnections),
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(cl.Timeout) * time.Second,
			KeepAlive: time.Duration(cl.KeepAlive) * time.Second,
		}).DialContext,
	}

	// Setup main client
	cl.Conn = &http.Client{
		Transport: t,
		Timeout:   time.Duration(cl.Timeout) * time.Second,
	}
	return cl, nil
}

// Dispatch a network request to the service.
func (cl *Client) Dispatch(r *RequestOptions) error {
	// Build request
	req, err := cl.newRequest(r)
	if err != nil {
		return err
	}

	// Dump request
	if cl.Debug {
		dump, err := httputil.DumpRequest(req, true)
		if err == nil {
			fmt.Println("=== request ===")
			fmt.Printf("%s\n", dump)
		}
	}

	// Execute request
	res, err := cl.Conn.Do(req)
	if res != nil {
		// Properly discard request content to be able to reuse the connection.
		defer func() {
			_, _ = io.Copy(io.Discard, res.Body)
			_ = res.Body.Close()
		}()
	}

	// Network level errors
	if err != nil {
		return err
	}

	// Dump response
	if cl.Debug {
		dump, err := httputil.DumpResponse(res, true)
		if err == nil {
			fmt.Println("=== response ===")
			fmt.Printf("%s\n", dump)
		}
	}

	// Get response contents
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// API level errors
	if res.StatusCode > 299 {
		e := new(Error)
		if err := json.Unmarshal(body, e); err == nil {
			return e
		}
		return fmt.Errorf("unsuccessful request: %s", res.Status)
	}

	// Decode response
	if r.Output != nil {
		// Unwrap "data" key in the response
		if r.UnwrapData {
			temp := map[string]interface{}{}
			if err := json.Unmarshal(body, &temp); err != nil {
				return fmt.Errorf("non JSON content returned: %s", body)
			}
			if data, ok := temp["data"]; ok {
				body, err = json.Marshal(data)
				if err != nil {
					return fmt.Errorf("non JSON content returned: %s", data)
				}
			}
		}

		// Load response data in the provided output interface
		if err = json.Unmarshal(body, r.Output); err != nil {
			return fmt.Errorf("non JSON content returned: %s", body)
		}
	}

	// All good!
	return nil
}

// Prepare a new API request.
func (cl *Client) newRequest(r *RequestOptions) (*http.Request, error) {
	// Default context
	if r.Ctx == nil {
		r.Ctx = context.TODO()
	}

	// Prepare input data
	var input io.Reader
	if r.Input != nil {
		data, err := json.Marshal(r.Input)
		if err != nil {
			return nil, fmt.Errorf("invalid input data: %w", err)
		}
		input = bytes.NewReader(data)
	}

	// Build basic request
	req, err := http.NewRequestWithContext(r.Ctx, r.Method, cl.Backend+r.Endpoint, input)
	if err != nil {
		return nil, err
	}

	// Add additional headers
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+cl.Key)
	req.Header.Add("Content-Type", "application/json")
	if cl.UserAgent != "" {
		req.Header.Add("User-Agent", cl.UserAgent)
	}

	// Add additional query parameters
	if r.QueryParams != nil {
		req.URL.RawQuery = r.QueryParams.Encode()
	}

	return req, nil
}
