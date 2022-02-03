package circlesdk

import "fmt"

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

// Error details for unsuccessful API requests.
// https://developers.circle.com/docs/common-error-responses
type Error struct {
	// General error type.
	// https://developers.circle.com/docs/api-response-errors
	Code int `json:"code,omitempty"`

	// Human-friendly message.
	Message string `json:"message,omitempty"`

	// Additional error details.
	// https://developers.circle.com/docs/entity-errors
	Details []ExtendedErrorDetails `json:"errors,omitempty"`
}

// ExtendedErrorDetails contains a list of one or multiple error descriptions
// associated with it.
type ExtendedErrorDetails struct {
	// Type of an error
	ErrorType string `json:"error,omitempty"`

	// Human-friendly message
	Message string `json:"message,omitempty"`

	// Period-separated path to the property that causes this error.
	// For example: address.billingCountry
	Location string `json:"location,omitempty"`

	// Actual value of the property specified in location key, as
	// received by the server.
	InvalidValue string `json:"invalidValue,omitempty"`

	// Special object that contains additional details about the error
	// and could be used for programmatic handling on the client side.
	Constraints ErrorConstraints `json:"constraints,omitempty"`
}

// ErrorConstraints contains additional details about the error and could be used
// for programmatic handling on the client side.
type ErrorConstraints struct {
	// Used to describe `max_value` or `min_value` errors.
	Min string `json:"min,omitempty"`

	// Used to describe `max_value` or `min_value` errors.
	Max string `json:"max,omitempty"`

	// Used to describe `max_value` or `min_value` errors.
	Inclusive bool `json:"inclusive,omitempty"`

	// Used to describe `pattern_mismatch` errors.
	Pattern string `json:"pattern,omitempty"`

	// Used to describe `number_format` errors.
	MaxIntegralDigits int `json:"max-integral-digits,omitempty"`

	// Used to described `number_format` errors.
	MaxFractionalDigits int `json:"max-fractional-digits,omitempty"`
}
