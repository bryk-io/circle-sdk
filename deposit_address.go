package circlesdk

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
