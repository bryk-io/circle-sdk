package circlesdk

import (
	"errors"
	"net/http"
)

// Balance of funds that are available for use.
type Balance struct {
	// List of currency balances (one for each currency) that are currently available
	// to spend.
	Available []BalanceEntry `json:"available,omitempty"`

	// List of currency balances (one for each currency) that have been captured but are
	// currently in the process of settling and will become available to spend at some
	// point in the future.
	Unsettled []BalanceEntry `json:"unsettled,omitempty"`
}

// BalanceEntry provides information for a specific amount/currency pair
// in the complete balance information.
type BalanceEntry struct {
	// Magnitude of the amount, in units of the currency.
	Amount string `json:"amount,omitempty"`

	// Currency code for the amount.
	Currency string `json:"currency,omitempty"`
}

// DepositAddress represents a blockchain account/destination where a
// user is available to receive funds.
type DepositAddress struct {
	// An alphanumeric string representing a blockchain address. Will be in
	// different formats for different chains. It is important to preserve the
	// exact formatting and capitalization of the address.
	Address string `json:"address,omitempty"`

	// The secondary identifier for a blockchain address. An example of this is
	// the memo field on the Stellar network, which can be text, id, or hash format.
	AddressTag string `json:"addressTag,omitempty"`

	// Currency associated with a balance or address.
	Currency string `json:"currency,omitempty"`

	// Blockchain that a given currency is available on.
	Chain string `json:"chain,omitempty"`

	// An identifier or sentence that describes the recipient.
	Description string `json:"description,omitempty"`
}

// GetMasterWalletID return your master wallet identifier.
func (cl *Client) GetMasterWalletID() (string, error) {
	type configResponse struct {
		Payments struct {
			MasterWalletID string `json:"masterWalletId,omitempty"`
		} `json:"payments"`
	}

	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/configuration",
		input:      nil,
		output:     &configResponse{},
		unwrapData: true,
	}
	if err := cl.dispatch(req); err != nil {
		return "", err
	}
	res, ok := req.output.(*configResponse)
	if !ok {
		return "", errors.New("invalid response")
	}
	return res.Payments.MasterWalletID, nil
}

// GetBalance retrieves the balance of funds that are available for use.
// https://developers.circle.com/reference#balances-get
func (cl *Client) GetBalance() (*Balance, error) {
	balance := new(Balance)
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/businessAccount/balances",
		input:      nil,
		output:     balance,
		unwrapData: true,
	}
	if err := cl.dispatch(req); err != nil {
		return nil, err
	}
	return balance, nil
}

// CreateDepositAddress generates a new blockchain address for a wallet for a
// given currency/chain pair. Circle may reuse addresses on blockchains that
// support reuse. For example, if you're requesting two addresses for depositing
// USD and ETH, both on Ethereum, you may see the same Ethereum address returned.
// Depositing cryptocurrency to a generated address will credit the associated
// wallet with the value of the deposit.
// https://developers.circle.com/reference#addresses-deposit-create
func (cl *Client) CreateDepositAddress(
	currency SupportedCurrency,
	chain SupportedChain,
	opts ...CallOption) (*DepositAddress, error) {
	address := new(DepositAddress)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   "v1/businessAccount/wallets/addresses/deposit",
		unwrapData: true,
		output:     address,
		input: map[string]interface{}{
			"currency": string(currency),
			"chain":    string(chain),
		},
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if req.idempotencyKey != "" {
		req.input["idempotencyKey"] = req.idempotencyKey
	}
	if err := cl.dispatch(req); err != nil {
		return nil, err
	}
	return address, nil
}

// GetDepositAddress returns a list of all available deposit address in the
// account.
// https://developers.circle.com/reference#addresses-deposit-get
func (cl *Client) GetDepositAddress(opts ...CallOption) ([]*DepositAddress, error) {
	var list []*DepositAddress
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/businessAccount/wallets/addresses/deposit",
		unwrapData: true,
		output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if err := cl.dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}

// AddRecipientAddress stores an external blockchain address. Once added, the
// recipient address must be verified to ensure that you know and trust each
// new address.
// https://developers.circle.com/reference#addresses-recipient-create
func (cl *Client) AddRecipientAddress(
	addr *DepositAddress,
	desc string,
	opts ...CallOption) (string, error) {
	res := map[string]string{}
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   "v1/businessAccount/wallets/addresses/recipient",
		unwrapData: true,
		output:     &res,
		input: map[string]interface{}{
			"address":     addr.Address,
			"addressTag":  addr.AddressTag,
			"currency":    addr.Currency,
			"chain":       addr.Chain,
			"description": desc,
		},
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return "", err
		}
	}
	if req.idempotencyKey != "" {
		req.input["idempotencyKey"] = req.idempotencyKey
	}
	if err := cl.dispatch(req); err != nil {
		return "", err
	}
	return res["id"], nil
}

// GetRecipientAddress returns a list of recipient addresses that have each been verified
// and are eligible for transfers. Any recipient addresses pending verification are not
// included in the response.
// https://developers.circle.com/reference#addresses-recipient-get
func (cl *Client) GetRecipientAddress(opts ...CallOption) ([]*DepositAddress, error) {
	var list []*DepositAddress
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/businessAccount/wallets/addresses/recipient",
		unwrapData: true,
		output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if err := cl.dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}
