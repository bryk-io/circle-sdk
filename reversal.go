package circlesdk

// Reversal is the object contain the reversal data returned from the API.
type Reversal struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Unique system generated identifier for the payment that is associated to the chargeback item.
	PaymentID string `json:"paymentId,omitempty"`

	// Amount object of the reversal
	Amount *Amount `json:"amount,omitempty"`

	// Enumerated description of the payment.
	Description string `json:"description,omitempty"`

	// Enumerated status of the payment. pending means the payment is waiting to be processed.
	// confirmed means the payment has been approved by the bank and the merchant can treat it as successful,
	// but settlement funds are not yet available to the merchant.
	// paid means settlement funds have been received and are available to the merchant.
	// failed means something went wrong (most commonly that the payment was denied).
	// Terminal states are paid and failed.
	Status string `json:"status,omitempty"`

	// Enumerated reason for a returned payment.
	// Providing this reason in the request is recommended (to improve risk evaluation) but not required.
	// options: duplicate fraudulent, requested_by_customer, bank_transaction_error, invalid_account_number,
	// insufficient_funds, payment_stopped_by_issuer, payment_returned, bank_account_ineligible,
	// invalid_ach_rtn, unauthorized_transaction, payment_failed
	Reason string `json:"reason,omitempty"`

	// Fees object for the reversal
	Fees *Amount `json:"fees,omitempty"`

	// ISO-8601 UTC date/time format of the reversal creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the reversal update date.
	UpdateDate string `json:"updateDate,omitempty"`
}
