package circlesdk

import "net/http"

type payoutsAPI struct {
	cl *Client
}

// CreatePayout creates a new payout.
// https://developers.circle.com/reference/payouts-payouts-create
func (mod *payoutsAPI) CreatePayout(c CreatePayoutRequest, opts ...CallOption) (*Payout, error) {
	res := new(Payout)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   "v1/payouts",
		unwrapData: true,
		output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if req.idempotencyKey != "" {
		c.IdempotencyKey = req.idempotencyKey
	}
	req.input = &c
	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// GetPayout returns the details of the specified payout.
// https://developers.circle.com/reference/payouts-payouts-get-id
func (mod *paymentsAPI) GetPayout(id string, opts ...CallOption) (*Payout, error) {
	res := new(Payout)
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/payouts/" + id,
		unwrapData: true,
		output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// ListPayouts returns a list of payouts based on the passed request and options.
// https://developers.circle.com/reference/payouts-payouts-get
func (mod *payoutsAPI) ListPayouts(l ListPayoutsRequest, opts ...CallOption) ([]*Payout, error) {
	var list []*Payout
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/payouts",
		unwrapData: true,
		output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if l.Source != "" {
		req.addQueryParam("source", l.Source)
	}
	if l.Destination != "" {
		req.addQueryParam("destination", l.Destination)
	}
	for _, s := range l.Status {
		req.addQueryParam("status", s)
	}
	for _, t := range l.Type {
		req.addQueryParam("type", t)
	}

	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}

// ListReturns returns a list of returns based on the passed options.
// https://developers.circle.com/reference/payouts-returns-get
func (mod *payoutsAPI) ListReturns(opts ...CallOption) ([]*PayoutReturn, error) {
	var list []*PayoutReturn
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/returns",
		unwrapData: true,
		output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}
