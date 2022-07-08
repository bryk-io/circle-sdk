package circlesdk

// Settlement is the object contain the settlement data returned from the API.
type Settlement struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// If this settlement was used for a marketplace payment, the wallet involved in the settlement.
	// Not included for standard merchant settlements.
	MerchantWalletID string `json:"merchantWalletId,omitempty"`

	// If this settlement was used for a marketplace payment, the wallet involved in the settlement.
	// Not included for standard merchant settlements.
	WalletID string `json:"walletId,omitempty"`

	// Total debits for the settlement
	TotalDebits *Amount `json:"totalDebits,omitempty"`

	// Total credits for the settlement
	TotalCredits *Amount `json:"totalCredits,omitempty"`

	// Payment fees for the settlement
	PaymentFees *Amount `json:"paymentFees,omitempty"`

	// Chargeback fees for the settlement
	ChargebackFees *Amount `json:"chargebackFees,omitempty"`

	// ISO-8601 UTC date/time format of the settlement creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the settlement update date.
	UpdateDate string `json:"updateDate,omitempty"`
}
