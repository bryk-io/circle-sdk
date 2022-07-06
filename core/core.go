package core

import (
	"errors"
	"net/http"

	"github.com/bryk-io/circle-sdk"
)

// API provides access to all core Circle APIs. This core set of APIs allow you to:
//   - Transfer digital currency (USDC) in and out of your Circle Account.
//   - Register your own business bank accounts - if you have them.
//   - Make transfers from / to your business bank account while seamlessly converting
//     those funds across digital currency and traditional FIAT.
// https://developers.circle.com/docs
type API struct {
	cl *circlesdk.Client
}

// Ping will perform a basic reachability test. Use it to make sure your
// client instance is properly setup.
func (mod *API) Ping() bool {
	type pingResponse struct {
		Message string `json:"message,omitempty"`
	}

	req := &circlesdk.RequestOptions{
		Method:   http.MethodGet,
		Endpoint: "ping",
		Input:    nil,
		Output:   &pingResponse{},
	}
	if err := mod.cl.Dispatch(req); err != nil {
		return false
	}
	res, ok := req.Output.(*pingResponse)
	if !ok {
		return false
	}
	return res.Message == "pong"
}

// GetMasterWalletID return your master wallet identifier.
func (mod *API) GetMasterWalletID() (string, error) {
	type configResponse struct {
		Payments struct {
			MasterWalletID string `json:"masterWalletId,omitempty"`
		} `json:"payments"`
	}

	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/configuration",
		Input:      nil,
		Output:     &configResponse{},
		UnwrapData: true,
	}
	if err := mod.cl.Dispatch(req); err != nil {
		return "", err
	}
	res, ok := req.Output.(*configResponse)
	if !ok {
		return "", errors.New("invalid response")
	}
	return res.Payments.MasterWalletID, nil
}

// GetBalance retrieves the balance of funds that are available for use.
// https://developers.circle.com/reference#balances-get
func (mod *API) GetBalance() (*circlesdk.Balance, error) {
	balance := new(circlesdk.Balance)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/businessAccount/balances",
		Input:      nil,
		Output:     balance,
		UnwrapData: true,
	}
	if err := mod.cl.Dispatch(req); err != nil {
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
func (mod *API) CreateDepositAddress(
	currency circlesdk.SupportedCurrency,
	chain circlesdk.SupportedChain,
	opts ...circlesdk.CallOption) (*circlesdk.DepositAddress, error) {
	address := new(circlesdk.DepositAddress)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   "v1/businessAccount/wallets/addresses/deposit",
		UnwrapData: true,
		Output:     address,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	input := map[string]interface{}{
		"currency": string(currency),
		"chain":    string(chain),
	}
	if req.IdempotencyKey != "" {
		input["idempotencyKey"] = req.IdempotencyKey
	}
	req.Input = input
	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return address, nil
}

// GetDepositAddressList returns a list of all available deposit address in the
// account.
// https://developers.circle.com/reference#addresses-deposit-get
func (mod *API) GetDepositAddressList(opts ...circlesdk.CallOption) ([]*circlesdk.DepositAddress, error) {
	var list []*circlesdk.DepositAddress
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/businessAccount/wallets/addresses/deposit",
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

// AddRecipientAddress stores an external blockchain address. Once added, the
// recipient address must be verified to ensure that you know and trust each
// new address.
// https://developers.circle.com/reference#addresses-recipient-create
func (mod *API) AddRecipientAddress(
	addr *circlesdk.DepositAddress,
	desc string,
	opts ...circlesdk.CallOption) (string, error) {
	res := map[string]string{}
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   "v1/businessAccount/wallets/addresses/recipient",
		UnwrapData: true,
		Output:     &res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return "", err
		}
	}
	input := map[string]interface{}{
		"address":     addr.Address,
		"addressTag":  addr.AddressTag,
		"currency":    addr.Currency,
		"chain":       addr.Chain,
		"description": desc,
	}
	if req.IdempotencyKey != "" {
		input["idempotencyKey"] = req.IdempotencyKey
	}
	req.Output = input
	if err := mod.cl.Dispatch(req); err != nil {
		return "", err
	}
	return res["id"], nil
}

// GetRecipientAddressList returns a list of recipient addresses that have each been verified
// and are eligible for transfers. Any recipient addresses pending verification are not
// included in the response.
// https://developers.circle.com/reference#addresses-recipient-get
func (mod *API) GetRecipientAddressList(opts ...circlesdk.CallOption) ([]*circlesdk.DepositAddress, error) {
	var list []*circlesdk.DepositAddress
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/businessAccount/wallets/addresses/recipient",
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
