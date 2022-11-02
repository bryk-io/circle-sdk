package payouts

import (
	"net/http"

	"github.com/bryk-io/circle-sdk"
)

// API The Circle Payouts API allows you to issue payouts to your customers, vendors, or
// suppliers in a variety of ways:
//   - Bank wires
//   - On-chain USDC transfers
//   - ACH (coming soon)
//
// Payouts are funded with your USDC denominated Circle Account, which can receive deposits
// from both traditional and blockchain payment rails.
type API struct {
	cl *circlesdk.Client
}

// CreatePayout creates a new payout.
// https://developers.circle.com/reference/payouts-payouts-create
func (mod *API) CreatePayout(c circlesdk.CreatePayoutRequest, opts ...circlesdk.CallOption) (*circlesdk.Payout, error) {
	res := new(circlesdk.Payout)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   "v1/payouts",
		UnwrapData: true,
		Output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if req.IdempotencyKey != "" {
		c.IdempotencyKey = req.IdempotencyKey
	}
	req.Input = &c
	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// GetPayout returns the details of the specified payout.
// https://developers.circle.com/reference/payouts-payouts-get-id
func (mod *API) GetPayout(id string, opts ...circlesdk.CallOption) (*circlesdk.Payout, error) {
	res := new(circlesdk.Payout)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/payouts/" + id,
		UnwrapData: true,
		Output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// ListPayouts returns a list of payouts based on the passed request and options.
// https://developers.circle.com/reference/payouts-payouts-get
func (mod *API) ListPayouts(l circlesdk.ListPayoutsRequest, opts ...circlesdk.CallOption) ([]*circlesdk.Payout, error) {
	var list []*circlesdk.Payout
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/payouts",
		UnwrapData: true,
		Output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if l.Source != "" {
		req.AddQueryParam("source", l.Source)
	}
	if l.Destination != "" {
		req.AddQueryParam("destination", l.Destination)
	}
	for _, s := range l.Status {
		req.AddQueryParam("status", string(s))
	}
	for _, t := range l.Type {
		req.AddQueryParam("type", t)
	}

	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}

// ListReturns returns a list of returns based on the passed options.
// https://developers.circle.com/reference/payouts-returns-get
func (mod *API) ListReturns(opts ...circlesdk.CallOption) ([]*circlesdk.PayoutReturn, error) {
	var list []*circlesdk.PayoutReturn
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/returns",
		UnwrapData: true,
		Output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}
