package accounts

import (
	"fmt"
	"net/http"

	circlesdk "github.com/bryk-io/circle-sdk"
)

// API the Circle Accounts API allows you to easily create and manage accounts and balances
// for your customers, and execute transfers of funds across accounts - whether they are
// within the Circle platform, or in / out of the platform via on-chain USDC connectivity.
//   - Embed US Dollar denominated accounts into your product or service without dealing
//     with the complexity of legacy bank account structures.
//   - Manage a multi-asset accounts infrastructure for your customers including seamless
//     transfer of funds, across hosted accounts or via on-chain USDC connectivity.
//   - Accept USDC deposits with minimum cost and no exposure to reversals.
//   - Support BTC and ETH balances in addition to USDC.
type API struct {
	cl *circlesdk.Client
}

// GetTransfersList searches for transfers involving the provided wallet.
// `walletID` should be the unique identifier for the source or destination wallet
// of transfers; useful for fetching all transfers related to a wallet. If no
// wallet ids are provided, searches all wallets associated with your Circle API
// account. If the date parameters are omitted, returns the most recent transfers.
// This endpoint returns up to 50 transfers in descending chronological order or
// pageSize, if provided.
// https://developers.circle.com/reference#accounts-transfers-get
func (mod *API) GetTransfersList(walletID string, opts ...circlesdk.CallOption) ([]*circlesdk.Transfer, error) {
	var list []*circlesdk.Transfer
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/transfers",
		UnwrapData: true,
		Output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if walletID != "" {
		req.AddQueryParam("walletId", walletID)
	}
	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}

// GetWalletsList retrieves a list of a user's wallets.
// https://developers.circle.com/reference#accounts-wallets-get
func (mod *API) GetWalletsList(opts ...circlesdk.CallOption) ([]*circlesdk.Wallet, error) {
	var list []*circlesdk.Wallet
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/wallets",
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

// GetWallet retrieves the full details of a given wallet.
// https://developers.circle.com/reference#accounts-wallets-get-id
func (mod *API) GetWallet(id string, opts ...circlesdk.CallOption) (*circlesdk.Wallet, error) {
	res := new(circlesdk.Wallet)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   "v1/wallets/" + id,
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

// GetWalletAddressList retrieves a list of addresses associated with a wallet.
// https://developers.circle.com/reference#accounts-wallets-addresses-get
func (mod *API) GetWalletAddressList(id string, opts ...circlesdk.CallOption) ([]*circlesdk.DepositAddress, error) {
	var list []*circlesdk.DepositAddress
	req := &circlesdk.RequestOptions{
		Method:     http.MethodGet,
		Endpoint:   fmt.Sprintf("v1/wallets/%s/addresses", id),
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

// CreateWallet creates a new end user wallet.
// https://developers.circle.com/reference#accounts-wallets-create
func (mod *API) CreateWallet(description string, opts ...circlesdk.CallOption) (*circlesdk.Wallet, error) {
	res := new(circlesdk.Wallet)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   "v1/wallets",
		UnwrapData: true,
		Output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	input := map[string]interface{}{
		"description": description,
	}
	if req.IdempotencyKey != "" {
		input["idempotencyKey"] = req.IdempotencyKey
	}
	req.Input = input
	if err := mod.cl.Dispatch(req); err != nil {
		return nil, err
	}
	return res, nil
}

// CreateWalletDepositAddress generates a new blockchain address for a wallet for a
// given currency/chain pair. Circle may reuse addresses on blockchains that support
// reuse. For example, if you're requesting two addresses for depositing USD and ETH,
// both on Ethereum, you may see the same Ethereum address returned. Depositing
// cryptocurrency to a generated address will credit the associated wallet with the
// value of the deposit.
// https://developers.circle.com/reference#accounts-wallets-addresses-create
func (mod *API) CreateWalletDepositAddress(
	walletID string,
	currency circlesdk.SupportedCurrency,
	chain circlesdk.SupportedChain,
	opts ...circlesdk.CallOption) (*circlesdk.DepositAddress, error) {
	address := new(circlesdk.DepositAddress)
	req := &circlesdk.RequestOptions{
		Method:     http.MethodPost,
		Endpoint:   fmt.Sprintf("v1/wallets/%s/addresses", walletID),
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
