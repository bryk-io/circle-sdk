package circlesdk

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
	Balances []Amount `json:"balances,omitempty"`
}
