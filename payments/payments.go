package payments

import (
	"fmt"
	"net/http"

	"github.com/bryk-io/circle-sdk"
)

// API The Circle Payments API allows you to take payments from your end users
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
type API struct {
	cl *circlesdk.Client
}

// CreateCard creates a card for future usage, i.e. creating payments.
// https://developers.circle.com/reference/payments-cards-create
func (mod *API) CreateCard(c circlesdk.CreateCardRequest, opts ...circlesdk.CallOption) (*circlesdk.Card, error) {
	res := new(circlesdk.Card)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   "v1/cards",
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

// GetCard returns details of the specified card.
// https://developers.circle.com/reference/payments-cards-get-id
func (mod *API) GetCard(id string, opts ...circlesdk.CallOption) (*circlesdk.Card, error) {
	res := new(circlesdk.Card)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/cards/" + id,
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

// ListCards returns a list of cards based on the passed options.
// https://developers.circle.com/reference/payments-cards-get
func (mod *API) ListCards(opts ...circlesdk.CallOption) ([]*circlesdk.Card, error) {
	var list []*circlesdk.Card
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/cards",
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

// UpdateCard updates the details of the specified card.
// https://developers.circle.com/reference/payments-cards-update-id
func (mod *API) UpdateCard(id string, c circlesdk.UpdateCardRequest, opts ...circlesdk.CallOption) (*circlesdk.Card, error) {
	res := new(circlesdk.Card)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPut,
		Endpoint:   "v1/cards/" + id,
		UnwrapData: true,
		Output:     res,
		Input: map[string]interface{}{
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
	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// CreateBankAccount creates a bank account for future usage, i.e. creating payments.
// https://developers.circle.com/reference/payments-bank-accounts-ach-create
func (mod *API) CreateBankAccount(b circlesdk.CreateBankAccountRequest, opts ...circlesdk.CallOption) (*circlesdk.BankAccount, error) {
	res := new(circlesdk.BankAccount)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   "v1/banks/ach",
		UnwrapData: true,
		Output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if req.IdempotencyKey != "" {
		b.IdempotencyKey = req.IdempotencyKey
	}
	req.Input = &b
	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// GetBankAccount returns the details of the specified bank account.
// https://developers.circle.com/reference/payments-bank-accounts-ach-get-id
func (mod *API) GetBankAccount(id string, opts ...circlesdk.CallOption) (*circlesdk.BankAccount, error) {
	res := new(circlesdk.BankAccount)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/banks/ach/" + id,
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

// CreatePayment creates a new payment.
// https://developers.circle.com/reference/payments-payments-create
func (mod *API) CreatePayment(c circlesdk.CreatePaymentRequest, opts ...circlesdk.CallOption) (*circlesdk.Payment, error) {
	res := new(circlesdk.Payment)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   "v1/payments",
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

// GetPayment returns the details of the specified payment.
// https://developers.circle.com/reference/payments-payments-get-id
func (mod *API) GetPayment(id string, opts ...circlesdk.CallOption) (*circlesdk.Payment, error) {
	res := new(circlesdk.Payment)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/payments/" + id,
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

// ListPayments returns a list of payments based on the passed request and options.
// https://developers.circle.com/reference/payments-payments-get
func (mod *API) ListPayments(l circlesdk.ListPaymentsRequest, opts ...circlesdk.CallOption) ([]*circlesdk.Payment, error) {
	var list []*circlesdk.Payment
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/payments",
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
	if l.SettlementID != "" {
		req.AddQueryParam("settlementId", l.SettlementID)
	}
	if l.Status != "" {
		req.AddQueryParam("status", l.Status)
	}
	for _, t := range l.Type {
		req.AddQueryParam("type", t)
	}

	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}

// CapturePayment the given amount will be captured for the authorized payment if possible.
// If no amount is specified, the full amount will be captured. You can only capture once per authorization.
// A successful response does not mean the payment has been captured. It only means the capture
// request was successfully submitted.
// https://developers.circle.com/reference/capturepayment
func (mod *API) CapturePayment(id string, c circlesdk.CapturePaymentRequest, opts ...circlesdk.CallOption) error {
	req := &circlesdk.RequestOptions{
		Method:   http.MethodPost,
		Endpoint: fmt.Sprintf("v1/payments/%s/capture", id),
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return err
		}
	}
	if req.IdempotencyKey != "" {
		c.IdempotencyKey = req.IdempotencyKey
	}
	req.Input = &c
	if err := mod.cl.Dispatch(req); err != nil {
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
func (mod *API) CancelPayment(id string, c circlesdk.CapturePaymentRequest, opts ...circlesdk.CallOption) (*circlesdk.Payment, error) {
	res := new(circlesdk.Payment)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   fmt.Sprintf("v1/payments/%s/cancel", id),
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

// RefundPayment the payment source will be refunded if possible. Not all payments are eligible to be canceled.
// A successful response does not mean the payment has been refunded;
// it only means the refund request is successfully submitted.
// https://developers.circle.com/reference/payments-payments-refund-id
func (mod *API) RefundPayment(id string, r circlesdk.RefundPaymentRequest, opts ...circlesdk.CallOption) (*circlesdk.Payment, error) {
	res := new(circlesdk.Payment)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   fmt.Sprintf("v1/payments/%s/refund", id),
		UnwrapData: true,
		Output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if req.IdempotencyKey != "" {
		r.IdempotencyKey = req.IdempotencyKey
	}
	req.Input = &r
	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// GetSettlement returns the details of the specified settlement.
// https://developers.circle.com/reference/payments-settlements-get-id
func (mod *API) GetSettlement(id string, opts ...circlesdk.CallOption) (*circlesdk.Settlement, error) {
	res := new(circlesdk.Settlement)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/settlements/" + id,
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

// ListSettlements returns a list of settlements based on the passed options.
// https://developers.circle.com/reference/payments-settlements-get
func (mod *API) ListSettlements(opts ...circlesdk.CallOption) ([]*circlesdk.Settlement, error) {
	var list []*circlesdk.Settlement
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/settlements",
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

// GetChargeback returns the details of the specified chargeback.
// https://developers.circle.com/reference/payments-chargebacks-get-id
func (mod *API) GetChargeback(id string, opts ...circlesdk.CallOption) (*circlesdk.ChargeBack, error) {
	res := new(circlesdk.ChargeBack)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/chargebacks/" + id,
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

// ListChargebacks returns a list of chargebacks based on the passed options.
// https://developers.circle.com/reference/payments-chargebacks-get
func (mod *API) ListChargebacks(paymentID string, opts ...circlesdk.CallOption) ([]*circlesdk.ChargeBack, error) {
	var list []*circlesdk.ChargeBack
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/chargebacks",
		UnwrapData: true,
		Output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if paymentID != "" {
		req.AddQueryParam("paymentID", paymentID)
	}

	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}

// ListReversals returns a list of reversals based on the passed options.
// https://developers.circle.com/reference/payments-reversals-get
func (mod *API) ListReversals(status string, opts ...circlesdk.CallOption) ([]*circlesdk.Reversal, error) {
	var list []*circlesdk.Reversal
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/reversals",
		UnwrapData: true,
		Output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	if status != "" {
		req.AddQueryParam("status", status)
	}

	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}
