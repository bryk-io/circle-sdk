package circlesdk

// ChargeBackHistory contains the data for one item of the chargeback object history property.
type ChargeBackHistory struct {
	// Enumerated type of the chargeback history event. 1st Chargeback represents the first stage of the dispute
	// procedure initiated by the cardholder’s issuing bank.  2nd Chargeback represents the second stage of the
	// dispute procedure initiated by the cardholder’s issuing bank (This stage is MasterCard only).
	// Chargeback Reversal represents when 1st Chargeback or 2nd Chargeback is withdrawn by the issuer.
	// Representment represents the stage when merchants decided to dispute 1st Chargeback or 2nd Chargeback.
	//Chargeback Settlement can imply one of the two: 1) If merchant or marketplace is taking the lost of the
	// chargeback, money will be debit from the wallet during this stage.
	//If merchant of marketplace successfully dispute the chargeback, money will be credit back to the wallet
	// during this stage.
	//1st Chargeback, 2nd Chargeback, Chargeback Reversal, Representment, Chargeback Settlement
	Type string `json:"type,omitempty"`

	// Chargeback amount object for the history
	ChargeBackAmount *Amount `json:"chargeBackAmount,omitempty"`

	// Fee object for the history
	Fee *Amount `json:"fee,omitempty"`

	// The reason the chargeback was created.
	Description string `json:"description,omitempty"`

	// Unique system generated identifier for the settlement related to the chargeback history.
	SettlementID string `json:"settlementId,omitempty"`

	// ISO-8601 UTC date/time format of the history creation date.
	CreateDate string `json:"createDate,omitempty"`
}

// ChargeBack is the object contain the chargeback data returned from the API.
type ChargeBack struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Unique system generated identifier for the payment that is associated to the chargeback item.
	PaymentID string `json:"paymentId,omitempty"`

	// Unique system generated identifier for the merchant.
	MerchantID string `json:"merchantId,omitempty"`

	// Reason code given by the card network for the chargeback item.
	ReasonCode string `json:"reasonCode,omitempty"`

	// Enumerated category of the chargeback status codes based on the chargeback status code.
	// options: Canceled Recurring Payment,  Customer Dispute, Fraudulent, General, Processing Error, Not Defined
	Category string `json:"category,omitempty"`

	// The chargeback item's history list will be sorted by create date descending:
	// more recent chargeback statuses will be at the beginning of the list.
	History []ChargeBackHistory `json:"history,omitempty"`
}
