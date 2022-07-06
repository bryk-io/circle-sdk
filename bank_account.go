package circlesdk

// BankAddress contains address details for the bank, as provided during bank account creation.
type BankAddress struct {
	// Name of the bank.
	// This property is required for bank accounts outside of the US that do not support IBAN'
	BankName string `json:"bankName,omitempty"`

	// City portion of the address.
	// This property is required for bank accounts outside of the US.
	City string `json:"city,omitempty"`

	// Country portion of the address.
	// Formatted as a two-letter country code specified in ISO 3166-1 alpha-2.
	Country string `json:"country,omitempty"`

	// Line one of the street address.
	Line1 string `json:"line1,omitempty"`

	// Line two of the street address.
	Line2 string `json:"line2,omitempty"`

	// State / County / Province / Region portion of the address.
	// US and Canada use the two-letter code for the subdivision.
	District string `json:"district,omitempty"`
}

// BankAccountStatus contains the status value for the bank account.
type BankAccountStatus string

const (
	// BankAccountStatusPending = "pending".
	BankAccountStatusPending BankAccountStatus = "pending"

	// BankAccountStatusComplete = "complete".
	BankAccountStatusComplete BankAccountStatus = "complete"

	// BankAccountStatusFailed = "failed".
	BankAccountStatusFailed BankAccountStatus = "failed"
)

// BankAccountErrorCode contains the error code value for the bank account.
type BankAccountErrorCode string

const (
	// BankAccountErrorCodeBankAccountAuthorizationExpired = "bank_account_authorization_expired".
	BankAccountErrorCodeBankAccountAuthorizationExpired BankAccountErrorCode = "bank_account_authorization_expired"

	// BankAccountErrorCodeBankAccountError = "bank_account_error".
	BankAccountErrorCodeBankAccountError BankAccountErrorCode = "bank_account_error"

	// BankAccountErrorCodeBankAccountIneligible = "bank_account_ineligible".
	BankAccountErrorCodeBankAccountIneligible BankAccountErrorCode = "bank_account_ineligible"

	// BankAccountErrorCodeBankAccountNotFound = "bank_account_not_found".
	BankAccountErrorCodeBankAccountNotFound BankAccountErrorCode = "bank_account_not_found"

	// BankAccountErrorCodeBankAccountUnauthorized = "bank_account_unauthorized".
	BankAccountErrorCodeBankAccountUnauthorized BankAccountErrorCode = "bank_account_unauthorized"

	// BankAccountErrorCodeUnsupportedRoutingNumber = "unsupported_routing_number".
	BankAccountErrorCodeUnsupportedRoutingNumber BankAccountErrorCode = "unsupported_routing_number"

	// BankAccountErrorCodeVerificationFailed = "verification_failed".
	BankAccountErrorCodeVerificationFailed BankAccountErrorCode = "verification_failed"
)

// BankAccount is the object contain the bank account data returned from the API.
type BankAccount struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Status of the account.
	// A pending status indicates that the linking is in-progress;
	// complete indicates the account was linked successfully;
	// failed indicates it failed.
	Status BankAccountStatus `json:"status,omitempty"`

	// The redacted account number of the ACH account.
	AccountNumber string `json:"accountNumber,omitempty"`

	// The routing number of the ACH account.
	RoutingNumber string `json:"routingNumber,omitempty"`

	// Object containing billing details for the bank account.
	BillingDetails *BillingDetails `json:"billingDetails,omitempty"`

	// The address details for the bank, as provided during bank account creation.
	BankAddress *BankAddress `json:"bankAddress,omitempty"`

	// A UUID that uniquely identifies the account number.
	// If the same account is used more than once, each card object will have a different id,
	// but the fingerprint will stay the same.
	Fingerprint string `json:"fingerprint,omitempty"`

	// Indicates the failure reason of the ACH account. Only present on failed accounts.
	// Possible values are [bank_account_authorization_expired, bank_account_error,
	// bank_account_ineligible, bank_account_not_found, bank_account_unauthorized,
	// unsupported_routing_number, verification_failed].
	ErrorCode BankAccountErrorCode `json:"errorCode,omitempty"`

	// Results of risk evaluation. Only present if the payment is denied by Circle's risk service.
	RiskEvaluation *RiskEvaluation `json:"riskEvaluation,omitempty"`

	// Object containing metadata for the bank account
	Metadata *Metadata `json:"metadata,omitempty"`

	// ISO-8601 UTC date/time format of the bank account creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the bank account update date.
	UpdateDate string `json:"updateDate"`
}

// CreateBankAccountRequest contains the data to create a bank account (ACH).
type CreateBankAccountRequest struct {
	// Universally unique identifier (UUID v4) idempotency key.
	// This key is utilized to ensure exactly-once execution of mutating requests.
	IdempotencyKey string `json:"idempotencyKey,omitempty"`

	// The token for the ACH account provided by the processor (Plaid).
	PlaidProcessorToken string `json:"plaidProcessorToken,omitempty"`

	// Billing details of the account holder.
	BillingDetails *BillingDetails `json:"billingDetails,omitempty"`

	// Object containing metadata for the bank account creation process
	Metadata *CreateMetadataRequest `json:"metadata,omitempty"`
}
