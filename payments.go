package circlesdk

import (
	"fmt"
	"net/http"
)

type paymentsAPI struct {
	cl *Client
}

// CreateCard creates a card for future usage, i.e. creating payments.
// https://developers.circle.com/reference/payments-cards-create
func (mod *paymentsAPI) CreateCard(c CreateCardRequest, opts ...CallOption) (*Card, error) {
	res := new(Card)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   "v1/cards",
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

// GetCard returns details of the specified card.
// https://developers.circle.com/reference/payments-cards-get-id
func (mod *paymentsAPI) GetCard(id string, opts ...CallOption) (*Card, error) {
	res := new(Card)
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/cards/" + id,
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

// ListCards returns a list of cards based on the passed options.
// https://developers.circle.com/reference/payments-cards-get
func (mod *paymentsAPI) ListCards(opts ...CallOption) ([]*Card, error) {
	var list []*Card
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/cards",
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

// UpdateCard updates the details of the specified card.
// https://developers.circle.com/reference/payments-cards-update-id
func (mod *paymentsAPI) UpdateCard(id string, c UpdateCardRequest, opts ...CallOption) (*Card, error) {
	res := new(Card)
	req := &requestOptions{
		method:     http.MethodPut,
		endpoint:   "v1/cards/" + id,
		unwrapData: true,
		output:     res,
		input: map[string]interface{}{
			"keyId":         c.KeyID,
			"encryptedData": c.EncryptedData,
			"expMonth":      c.ExpMonth,
			"expYear":       c.ExpYear,
		},
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

// CreateBankAccount creates a bank account for future usage, i.e. creating payments.
// https://developers.circle.com/reference/payments-bank-accounts-ach-create
func (mod *paymentsAPI) CreateBankAccount(b CreateBankAccountRequest, opts ...CallOption) (*BankAccount, error) {
	res := new(BankAccount)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   "v1/banks/ach",
		unwrapData: true,
		output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if req.idempotencyKey != "" {
		b.IdempotencyKey = req.idempotencyKey
	}
	req.input = &b
	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// GetBankAccount returns the details of the specified bank account.
// https://developers.circle.com/reference/payments-bank-accounts-ach-get-id
func (mod *paymentsAPI) GetBankAccount(id string, opts ...CallOption) (*BankAccount, error) {
	res := new(BankAccount)
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/banks/ach/" + id,
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

// CreatePayment creates a new payment.
// https://developers.circle.com/reference/payments-payments-create
func (mod *paymentsAPI) CreatePayment(c CreatePaymentRequest, opts ...CallOption) (*Payment, error) {
	res := new(Payment)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   "v1/payments",
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

// GetPayment returns the details of the specified payment.
// https://developers.circle.com/reference/payments-payments-get-id
func (mod *paymentsAPI) GetPayment(id string, opts ...CallOption) (*Payment, error) {
	res := new(Payment)
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/payments/" + id,
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

// ListPayments returns a list of payments based on the passed request and options.
// https://developers.circle.com/reference/payments-payments-get
func (mod *paymentsAPI) ListPayments(l ListPaymentsRequest, opts ...CallOption) ([]*Payment, error) {
	var list []*Payment
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/payments",
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
	if l.SettlementID != "" {
		req.addQueryParam("settlementId", l.SettlementID)
	}
	if l.Status != "" {
		req.addQueryParam("status", l.Status)
	}
	for _, t := range l.Type {
		req.addQueryParam("type", t)
	}

	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}

// CapturePayment the given amount will be captured for the authorized payment if possible.
// If no amount is specified, the full amount will be captured. You can only capture once per authorization.
// A successful response does not mean the payment has been captured. It only means the capture
// request was successfully submitted.
// https://developers.circle.com/reference/capturepayment
func (mod *paymentsAPI) CapturePayment(id string, c CapturePaymentRequest, opts ...CallOption) error {
	req := &requestOptions{
		method:   http.MethodPost,
		endpoint: fmt.Sprintf("v1/payments/%s/capture", id),
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return err
		}
	}
	if req.idempotencyKey != "" {
		c.IdempotencyKey = req.idempotencyKey
	}
	req.input = &c
	if err := mod.cl.dispatch(req); err != nil {
		return err
	}
	return nil
}

// CancelPayment the payment will be voided if possible meaning the payment source will not be charged
// & the payment will never settle. Otherwise, the payment will be refunded meaning the payment source will be
// charged & the payment will be refunded from deductions of future settlements. Not all payments are eligible
// to be canceled. A successful response does not mean the payment has been canceled;
// it only means the cancellation request is successfully submitted.
// https://developers.circle.com/reference/payments-payments-cancel-id
func (mod *paymentsAPI) CancelPayment(id string, c CapturePaymentRequest, opts ...CallOption) (*Payment, error) {
	res := new(Payment)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   fmt.Sprintf("v1/payments/%s/cancel", id),
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

// RefundPayment the payment source will be refunded if possible. Not all payments are eligible to be canceled.
// A successful response does not mean the payment has been refunded;
// it only means the refund request is successfully submitted.
// https://developers.circle.com/reference/payments-payments-refund-id
func (mod *paymentsAPI) RefundPayment(id string, r RefundPaymentRequest, opts ...CallOption) (*Payment, error) {
	res := new(Payment)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   fmt.Sprintf("v1/payments/%s/refund", id),
		unwrapData: true,
		output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if req.idempotencyKey != "" {
		r.IdempotencyKey = req.idempotencyKey
	}
	req.input = &r
	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}
