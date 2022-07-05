package circlesdk

// CardVerification indicates the status of the card for verification purposes.
type CardVerification struct {
	// Status of the AVS check. Raw AVS response, expressed as an upper-case letter.
	// not_requested indicates check was not made. pending is pending/processing.
	Avs string `json:"avs,omitempty"`

	// Enumerated status of the check.
	// not_requested indicates check was not made.
	// pass indicates value is correct.
	// fail indicates value is incorrect.
	// unavailable indicates card issuer did not do the provided check.
	// pending indicates check is pending/processing.
	Cvv string `json:"cvv,omitempty"`
}

// Card is the object contain the card data returned from the API.
type Card struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Status of the account.
	// A pending status indicates that the linking is in-progress;
	// complete indicates the account was linked successfully;
	// failed indicates it failed.
	Status string `json:"status,omitempty"`

	// Object containing billing details for the card.
	BillingDetails *BillingDetails `json:"billingDetails,omitempty"`

	// Two digit number representing the card's expiration month.
	ExpMonth int `json:"expMonth,omitempty"`

	// Four digit number representing the card's expiration year.
	ExpYear int `json:"expYear,omitempty"`

	// The network of the card.
	// options: VISA, MASTERCARD, AMEX, UNKNOWN
	Network string `json:"network,omitempty"`

	// The last 4 digits of the card.
	Last4 string `json:"last4,omitempty"`

	// The bank identification number (BIN), the first 6 digits of the card.
	Bin string `json:"bin,omitempty"`

	// The country code of the issuer bank. Follows the ISO 3166-1 alpha-2 standard.
	IssuerCountry string `json:"issuerCountry,omitempty"`

	// The funding type of the card. Possible values are credit, debit, prepaid, and unknown.
	FundingType string `json:"fundingType,omitempty"`

	// A UUID that uniquely identifies the account number.
	// If the same account is used more than once, each card object will have a different id,
	// but the fingerprint will stay the same.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Indicates the failure reason of the card verification. Only present on cards with failed verification.
	// Possible values are [verification_failed, verification_fraud_detected, verification_denied,
	// verification_not_supported_by_issuer, verification_stopped_by_issuer, card_failed, card_invalid,
	// card_address_mismatch, card_zip_mismatch, card_cvv_invalid, card_expired, card_limit_violated,
	// card_not_honored, card_cvv_required, credit_card_not_allowed, card_account_ineligible, card_network_unsupported]'
	ErrorCode string `json:"errorCode,omitempty"`

	// Indicates the status of the card for verification purposes.
	Verification *CardVerification `json:"verification,omitempty"`

	// Results of risk evaluation. Only present if the payment is denied by Circle's risk service.
	RiskEvaluation *RiskEvaluation `json:"riskEvaluation,omitempty"`

	// Object containing metadata for the card
	Metadata *Metadata `json:"metadata,omitempty"`

	// ISO-8601 UTC date/time format of the card creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the card update date.
	UpdateDate string `json:"updateDate"`
}

// CreateCardRequest contains the data to create a card.
type CreateCardRequest struct {
	// Universally unique identifier (UUID v4) idempotency key.
	// This key is utilized to ensure exactly-once execution of mutating requests.
	IdempotencyKey string `json:"idempotencyKey,omitempty"`

	// Universally unique identifier (UUID v4) of the public key used in encryption.
	// NOTE the sandbox environment uses the default value of key1.
	// For this reason the example supplied is key1 rather than a UUID.
	KeyID string `json:"keyId,omitempty"`

	// PGP encrypted base64 encoded string. Contains Number and CVV.
	EncryptedData string `json:"encryptedData,omitempty"`

	// Object containing billing details for the card.
	BillingDetails *BillingDetails `json:"billingDetails,omitempty"`

	// Two digit number representing the card's expiration month.
	ExpMonth int `json:"expMonth,omitempty"`

	// Four digit number representing the card's expiration year.
	ExpYear int `json:"expYear,omitempty"`

	// Object containing metadata for the card creation process
	Metadata *CreateMetadataRequest `json:"metadata,omitempty"`
}

// UpdateCardRequest contains the data to update a card.
type UpdateCardRequest struct {
	// Universally unique identifier (UUID v4) of the public key used in encryption.
	// NOTE the sandbox environment uses the default value of key1.
	// For this reason the example supplied is key1 rather than a UUID.
	KeyID string `json:"keyId,omitempty"`

	// PGP encrypted base64 encoded string. Contains Number and CVV.
	EncryptedData string `json:"encryptedData,omitempty"`

	// Two digit number representing the card's expiration month.
	ExpMonth int `json:"expMonth,omitempty"`

	// Four digit number representing the card's expiration year.
	ExpYear int `json:"expYear,omitempty"`
}
