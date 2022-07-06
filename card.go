package circlesdk

// CardVerificationAvs contains the avs value for the card verification.
type CardVerificationAvs string

const (
	// CardVerificationAvsNotRequested = "not_requested".
	CardVerificationAvsNotRequested CardVerificationAvs = "not_requested"

	// CardVerificationAvsPending = "pending".
	CardVerificationAvsPending CardVerificationAvs = "pending"
)

// CardVerificationCvv contains the cvv value for the card verification.
type CardVerificationCvv string

const (
	// CardVerificationCvvNotRequested = "not_requested".
	CardVerificationCvvNotRequested CardVerificationCvv = "not_requested"

	// CardVerificationCvvPass = "pass".
	CardVerificationCvvPass CardVerificationCvv = "pass"

	// CardVerificationCvvFail = "fail".
	CardVerificationCvvFail CardVerificationCvv = "fail"

	// CardVerificationCvvUnavailable = "unavailable".
	CardVerificationCvvUnavailable CardVerificationCvv = "unavailable"

	// CardVerificationCvvPending = "pending".
	CardVerificationCvvPending CardVerificationCvv = "pending"
)

// CardVerification indicates the status of the card for verification purposes.
type CardVerification struct {
	// Status of the AVS check. Raw AVS response, expressed as an upper-case letter.
	// not_requested indicates check was not made. pending is pending/processing.
	Avs CardVerificationAvs `json:"avs,omitempty"`

	// Enumerated status of the check.
	// not_requested indicates check was not made.
	// pass indicates value is correct.
	// fail indicates value is incorrect.
	// unavailable indicates card issuer did not do the provided check.
	// pending indicates check is pending/processing.
	Cvv CardVerificationCvv `json:"cvv,omitempty"`
}

// CardStatus contains the status value for the card.
type CardStatus string

const (
	// CardStatusPending = "pending".
	CardStatusPending CardStatus = "pending"

	// CardStatusComplete = "complete".
	CardStatusComplete CardStatus = "complete"

	// CardStatusFailed = "failed".
	CardStatusFailed CardStatus = "failed"
)

// CardNetwork contains the network value for the card.
type CardNetwork string

const (
	// CardNetworkVISA = "VISA".
	CardNetworkVISA CardNetwork = "VISA"

	// CardNetworkMASTERCARD = "MASTERCARD".
	CardNetworkMASTERCARD CardNetwork = "MASTERCARD"

	// CardNetworkAMEX = "AMEX".
	CardNetworkAMEX CardNetwork = "AMEX"

	// CardNetworkUNKNOWN = "UNKNOWN".
	CardNetworkUNKNOWN CardNetwork = "UNKNOWN"
)

// CardFundingType contains the funding type value for the card.
type CardFundingType string

//credit, debit, prepaid, and unknown.
const (
	// CardFundingTypeCredit = "credit".
	CardFundingTypeCredit CardFundingType = "credit"

	// CardFundingTypeDebit = "debit".
	CardFundingTypeDebit CardFundingType = "debit"

	// CardFundingTypePrepaid = "prepaid".
	CardFundingTypePrepaid CardFundingType = "prepaid"

	// CardFundingTypeUnknown = "unknown".
	CardFundingTypeUnknown CardFundingType = "unknown"
)

// CardErrorCode contains the error code value for the card.
type CardErrorCode string

const (
	// CardErrorCodeVerificationFailed = "verification_failed".
	CardErrorCodeVerificationFailed CardErrorCode = "verification_failed"

	// CardErrorCodeVerificationFraudDetected = "verification_fraud_detected".
	CardErrorCodeVerificationFraudDetected CardErrorCode = "verification_fraud_detected"

	// CardErrorCodeVerificationDenied = "verification_denied".
	CardErrorCodeVerificationDenied CardErrorCode = "verification_denied"

	// CardErrorCodeVerificationNotSupportedByIssuer = "verification_not_supported_by_issuer".
	CardErrorCodeVerificationNotSupportedByIssuer CardErrorCode = "verification_not_supported_by_issuer"

	// CardErrorCodeVerificationStoppedByIssuer = "verification_stopped_by_issuer".
	CardErrorCodeVerificationStoppedByIssuer CardErrorCode = "verification_stopped_by_issuer"

	// CardErrorCodeCardFailed = "card_failed".
	CardErrorCodeCardFailed CardErrorCode = "card_failed"

	// CardErrorCodeCardInvalid = "card_invalid".
	CardErrorCodeCardInvalid CardErrorCode = "card_invalid"

	// CardErrorCodeCardAddressMismatch = "card_address_mismatch".
	CardErrorCodeCardAddressMismatch CardErrorCode = "card_address_mismatch"

	// CardErrorCodeCardZipMismatch = "card_zip_mismatch".
	CardErrorCodeCardZipMismatch CardErrorCode = "card_zip_mismatch"

	// CardErrorCodeCardCvvInvalid = "card_cvv_invalid".
	CardErrorCodeCardCvvInvalid CardErrorCode = "card_cvv_invalid"

	// CardErrorCodeCardExpired = "card_expired".
	CardErrorCodeCardExpired CardErrorCode = "card_expired"

	// CardErrorCodeCardLimitViolated = "card_limit_violated".
	CardErrorCodeCardLimitViolated CardErrorCode = "card_limit_violated"

	// CardErrorCodeCardNotHonored = "card_not_honored".
	CardErrorCodeCardNotHonored CardErrorCode = "card_not_honored"

	// CardErrorCodeCardCvvRequired = "card_cvv_required".
	CardErrorCodeCardCvvRequired CardErrorCode = "card_cvv_required"

	// CardErrorCodeCreditCardNotAllowed = "credit_card_not_allowed".
	CardErrorCodeCreditCardNotAllowed CardErrorCode = "credit_card_not_allowed"

	// CardErrorCodeCardAccountIneligible = "card_account_ineligible".
	CardErrorCodeCardAccountIneligible CardErrorCode = "card_account_ineligible"

	// CardErrorCodeCardNetworkUnsupported = "card_network_unsupported".
	CardErrorCodeCardNetworkUnsupported CardErrorCode = "card_network_unsupported"
)

// Card is the object contain the card data returned from the API.
type Card struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Status of the account.
	// A pending status indicates that the linking is in-progress;
	// complete indicates the account was linked successfully;
	// failed indicates it failed.
	Status CardStatus `json:"status,omitempty"`

	// Object containing billing details for the card.
	BillingDetails *BillingDetails `json:"billingDetails,omitempty"`

	// Two digit number representing the card's expiration month.
	ExpMonth int `json:"expMonth,omitempty"`

	// Four digit number representing the card's expiration year.
	ExpYear int `json:"expYear,omitempty"`

	// The network of the card.
	// options: VISA, MASTERCARD, AMEX, UNKNOWN
	Network CardNetwork `json:"network,omitempty"`

	// The last 4 digits of the card.
	Last4 string `json:"last4,omitempty"`

	// The bank identification number (BIN), the first 6 digits of the card.
	Bin string `json:"bin,omitempty"`

	// The country code of the issuer bank. Follows the ISO 3166-1 alpha-2 standard.
	IssuerCountry string `json:"issuerCountry,omitempty"`

	// The funding type of the card. Possible values are credit, debit, prepaid, and unknown.
	FundingType CardFundingType `json:"fundingType,omitempty"`

	// A UUID that uniquely identifies the account number.
	// If the same account is used more than once, each card object will have a different id,
	// but the fingerprint will stay the same.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Indicates the failure reason of the card verification. Only present on cards with failed verification.
	// Possible values are [verification_failed, verification_fraud_detected, verification_denied,
	// verification_not_supported_by_issuer, verification_stopped_by_issuer, card_failed, card_invalid,
	// card_address_mismatch, card_zip_mismatch, card_cvv_invalid, card_expired, card_limit_violated,
	// card_not_honored, card_cvv_required, credit_card_not_allowed, card_account_ineligible, card_network_unsupported]'
	ErrorCode CardErrorCode `json:"errorCode,omitempty"`

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
