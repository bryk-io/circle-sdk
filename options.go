package circlesdk

import (
	"errors"
	"os"
)

// Option settings allow you to customize the behavior of a client
// instance.
type Option func(*Client) error

// WithUserAgent adjust the network agent value reported to the service.
func WithUserAgent(agent string) Option {
	return func(c *Client) error {
		c.userAgent = agent
		return nil
	}
}

// WithKeepAlive adjust the time to maintain open the connection with the
// service, in seconds.
func WithKeepAlive(val uint) Option {
	return func(c *Client) error {
		c.keepAlive = val
		return nil
	}
}

// WithMaxConnections adjusts the maximum number of network connections to
// keep open with the service.
func WithMaxConnections(val uint) Option {
	return func(c *Client) error {
		c.maxConnections = val
		return nil
	}
}

// WithTimeout specifies a time limit for requests made by this client. The
// timeout includes connection time, any redirects, and reading the response
// body. A timeout of zero means no timeout.
func WithTimeout(val uint) Option {
	return func(c *Client) error {
		c.timeout = val
		return nil
	}
}

// WithAPIKey specifies the Circle API key used to access the service.
func WithAPIKey(key string) Option {
	return func(c *Client) error {
		c.key = key
		return nil
	}
}

// WithAPIKeyFromEnv loads the Circle API key from the ENV variable specified
// in 'name'.
func WithAPIKeyFromEnv(name string) Option {
	return func(c *Client) error {
		c.key = os.Getenv(name)
		if c.key == "" {
			return errors.New("env variable not available")
		}
		return nil
	}
}

// WithDebug makes the client produce trace output of requests and responses.
// Recommended only for testing and development.
func WithDebug() Option {
	return func(c *Client) error {
		c.debug = true
		return nil
	}
}

// WithProductionBackend adjust the client to consume the production API.
// If not specified, the client will submit requests to the sandbox environment
// by default.
func WithProductionBackend() Option {
	return func(c *Client) error {
		c.backend = endpointProduction
		return nil
	}
}
