package circlesdk

import (
	"fmt"
	"net/http"
)

type accountsAPI struct {
	cl *Client
}

// GetTransfersList searches for transfers involving the provided wallet.
// `walletID` should be the unique identifier for the source or destination wallet
// of transfers; useful for fetching all transfers related to a wallet. If no
// wallet ids are provided, searches all wallets associated with your Circle API
// account. If the date parameters are omitted, returns the most recent transfers.
// This endpoint returns up to 50 transfers in descending chronological order or
// pageSize, if provided.
// https://developers.circle.com/reference#accounts-transfers-get
func (mod *accountsAPI) GetTransfersList(walletID string, opts ...CallOption) ([]*Transfer, error) {
	var list []*Transfer
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/transfers",
		unwrapData: true,
		output:     &list,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	if walletID != "" {
		req.addQueryParam("walletId", walletID)
	}
	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return list, nil
}

// GetWalletsList retrieves a list of a user's wallets.
// https://developers.circle.com/reference#accounts-wallets-get
func (mod *accountsAPI) GetWalletsList(opts ...CallOption) ([]*Wallet, error) {
	var list []*Wallet
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/wallets",
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

// GetWallet retrieves the full details of a given wallet.
// https://developers.circle.com/reference#accounts-wallets-get-id
func (mod *accountsAPI) GetWallet(id string, opts ...CallOption) (*Wallet, error) {
	res := new(Wallet)
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   "v1/wallets/" + id,
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

// GetWalletAddressList retrieves a list of addresses associated with a wallet.
// https://developers.circle.com/reference#accounts-wallets-addresses-get
func (mod *accountsAPI) GetWalletAddressList(id string, opts ...CallOption) ([]*DepositAddress, error) {
	var list []*DepositAddress
	req := &requestOptions{
		method:     http.MethodGet,
		endpoint:   fmt.Sprintf("v1/wallets/%s/addresses", id),
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

// CreateWallet creates a new end user wallet.
// https://developers.circle.com/reference#accounts-wallets-create
func (mod *accountsAPI) CreateWallet(description string, opts ...CallOption) (*Wallet, error) {
	res := new(Wallet)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   "v1/wallets",
		unwrapData: true,
		output:     res,
	}
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}
	input := map[string]interface{}{
		"description": description,
	}
	if req.idempotencyKey != "" {
		input["idempotencyKey"] = req.idempotencyKey
	}
	req.input = input
	if err := mod.cl.dispatch(req); err != nil {
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
func (mod *accountsAPI) CreateWalletDepositAddress(
	walletID string,
	currency SupportedCurrency,
	chain SupportedChain,
	opts ...CallOption) (*DepositAddress, error) {
	address := new(DepositAddress)
	req := &requestOptions{
		method:     http.MethodPost,
		endpoint:   fmt.Sprintf("v1/wallets/%s/addresses", walletID),
		unwrapData: true,
		output:     address,
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
	if req.idempotencyKey != "" {
		input["idempotencyKey"] = req.idempotencyKey
	}
	req.input = input
	if err := mod.cl.dispatch(req); err != nil {
		return nil, err
	}
	return address, nil
}
