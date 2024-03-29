package circlesdk

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// CallOption settings allow you to adjust the behavior of specific
// API calls.
type CallOption func(*RequestOptions) error

// WithContext allows you to provide a custom context to the API call.
func WithContext(ctx context.Context) CallOption {
	return func(req *RequestOptions) error {
		req.Ctx = ctx
		return nil
	}
}

// WithIdempotencyKey makes the request idempotent so that you can safely
// retry API calls when things go wrong before you receive a response. If
// "ik" is empty a new valid idempotency key will be generated.
// https://developers.circle.com/docs/a-note-on-idempotent-requests
func WithIdempotencyKey(ik string) CallOption {
	return func(req *RequestOptions) error {
		if ik == "" {
			ik = uuid.NewString()
		}
		req.IdempotencyKey = ik
		return nil
	}
}

// WithPageSize limits the number of items to be returned by API calls
// returning collections. Some collections have a strict upper bound that
// will disregard this value. In case the specified value is higher than
// the allowed limit, the collection limit will be used. If not provided,
// the collection will determine the page size itself.
func WithPageSize(size uint) CallOption {
	return func(req *RequestOptions) error {
		req.AddQueryParam("pageSize", fmt.Sprintf("%d", size))
		return nil
	}
}

// WithDateRange limits the collection items returned by API calls to the
// specified date range (inclusive).
func WithDateRange(from, to time.Time) CallOption {
	return func(req *RequestOptions) error {
		req.AddQueryParam("to", to.Format(time.RFC3339))
		req.AddQueryParam("from", from.Format(time.RFC3339))
		return nil
	}
}

// WithPageBefore marks the exclusive end of a page.
// When provided, the collection resource will return the next n items before
// the id, with n being specified by pageSize.
func WithPageBefore(id string) CallOption {
	return func(req *RequestOptions) error {
		req.AddQueryParam("pageBefore", id)
		return nil
	}
}

// WithPageAfter marks the exclusive begin of a page.
// When provided, the collection resource will return the next n items after
// the id, with n being specified by pageSize.
func WithPageAfter(id string) CallOption {
	return func(req *RequestOptions) error {
		req.AddQueryParam("pageAfter", id)
		return nil
	}
}
