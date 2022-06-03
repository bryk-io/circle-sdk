package circlesdk

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
	// Entity identifier. Numeric value but should be treated as a string as
	// format may change in the future
	ID string `json:"id,omitempty"`

	// Entity type.
	Kind string `json:"type,omitempty"`

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

// Wallet entry associated with a user's account.
type Wallet struct {
	// Wallet identifier. Numeric value but should be treated as a string as
	// format may change in the future
	ID string `json:"walletId,omitempty"`

	// Unique identifier of the entity that owns the wallet.
	Entity string `json:"entityId,omitempty"`

	// Wallet type.
	Kind string `json:"type,omitempty"`

	// A human-friendly, non-unique identifier for a wallet.
	Description string `json:"description,omitempty"`

	// A list of balances for currencies owned by the wallet.
	Balances []BalanceEntry `json:"balances,omitempty"`
}

// Transfer of funds between a `source` and `destination` addresses.
type Transfer struct {
	// Unique identifier for this transfer.
	ID string `json:"id,omitempty"`

	// Source of the funds.
	Source *DepositAddress `json:"source,omitempty"`

	// Destination of the funds.
	Destination *DepositAddress `json:"destination,omitempty"`

	// Nominal value transferred.
	Amount *BalanceEntry `json:"amount,omitempty"`

	// A hash that uniquely identifies the on-chain transaction. This is only
	// available where either source or destination are of type blockchain.
	TxHash string `json:"transactionHash,omitempty"`

	// Status of the transfer. Status `pending` indicates that the transfer is in
	// the process of running; `complete` indicates it finished successfully;
	// `failed` indicates it failed.
	Status string `json:"status,omitempty"`

	// Indicates the failure reason of a transfer. Only present for transfers in a
	// failed state. Possible values are `insufficient_funds`, `blockchain_error`,
	// `transfer_denied` and `transfer_failed`.
	ErrorCode string `json:"errorCode,omitempty"`

	// The creation date of the transfer.
	CreateDate string `json:"createDate,omitempty"`
}
