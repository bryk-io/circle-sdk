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

// BillingDetails is the object containing billing details for the entity.
type BillingDetails struct {
	// Full name of the card or bank account holder.
	Name string `json:"name,omitempty"`

	// City portion of the address.
	City string `json:"city,omitempty"`

	// Country portion of the address.
	// Formatted as a two-letter country code specified in ISO 3166-1 alpha-2.
	Country string `json:"country,omitempty"`

	// Line one of the street address.
	Line1 string `json:"line1,omitempty"`

	// Line two of the street address.
	Line2 string `json:"line2,omitempty"`

	// State / County / Province / Region portion of the address.
	// If the country is US or Canada, then district is required
	// and should use the two-letter code for the subdivision.
	District string `json:"district,omitempty"`

	// Postal / ZIP code of the address.
	PostalCode string `json:"postalCode,omitempty"`
}

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

// RiskEvaluation contains the Result of risk evaluation.
// Only present if the payment is denied by Circle's risk service.
type RiskEvaluation struct {
	// Enumerated decision of the account.
	// Options: approved, denied, review
	Decision string `json:"decision,omitempty"`

	// Risk reason for the definitive decision outcome.
	Reason string `json:"reason,omitempty"`
}

// Metadata is the object containing metadata for the entity.
type Metadata struct {
	// Email of the user.
	Email string `json:"email,omitempty"`

	// Phone number of the user in E.164 format.
	// We recommend using a library such as libphonenumber to parse and validate phone numbers.
	PhoneNumber string `json:"phoneNumber,omitempty"`
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

// CreateMetadataRequest contains the data to create metadata for entities.
type CreateMetadataRequest struct {
	// Email of the user.
	Email string `json:"email,omitempty"`

	// Phone number of the user in E.164 format.
	// We recommend using a library such as libphonenumber to parse and validate phone numbers.
	PhoneNumber string `json:"phoneNumber,omitempty"`

	// Hash of the session identifier; typically of the end user.
	// This helps us make risk decisions and prevent fraud.
	// IMPORTANT: Please hash the session identifier to prevent sending us actual session identifiers.
	SessionID string `json:"sessionId,omitempty"`

	// Single IPv4 or IPv6 address of user.
	IPAddress string `json:"ipAddress,omitempty"`
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

// BankAccount is the object contain the bank account data returned from the API.
type BankAccount struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Status of the account.
	// A pending status indicates that the linking is in-progress;
	// complete indicates the account was linked successfully;
	// failed indicates it failed.
	Status string `json:"status,omitempty"`

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
	ErrorCode string `json:"errorCode,omitempty"`

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

// Source object used for the payment.
type Source struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id"`

	// Type of the source.
	// options: card, ach, wire, sepa
	Type string `json:"type"`
}

// RequiredAction indicates when the payment status is action_required,
// this object summarizes the required additional steps.
type RequiredAction struct {
	// The type of action that is required to proceed with the payment. Currently only one type is supported.
	Type string `json:"type,omitempty"`

	// The URL to bring the user to in order to complete the payment.
	RedirectURL string `json:"redirectUrl,omitempty"`
}

// PaymentVerification indicates the status of the payment verification.
// This property will be present once the payment is confirmed.
type PaymentVerification struct {
	// Status of the AVS check. Raw AVS response, expressed as an upper-case letter.
	// not_requested indicates check was not made.
	// pending is pending/processing.
	Avs string `json:"avs,omitempty"`

	// Enumerated status of the check.
	// not_requested indicates check was not made.
	// pass indicates value is correct.
	// fail indicates value is incorrect.
	// unavailable indicates card issuer did not do the provided check.
	// pending indicates check is pending/processing.
	Cvv string `json:"cvv,omitempty"`

	// Enumerated status of the check.
	// pass indicates successful 3DS authentication.
	// fail indicates failed 3DS authentication.
	ThreeDSecure string `json:"threeDSecure"`

	// ECI (electronic commerce indicator) value returned by Directory Servers
	// (namely Visa, MasterCard, JCB, and American Express) indicating the outcome
	// of authentication attempted on transactions enforced by 3DS.
	Eci string `json:"eci,omitempty"`
}

// Payment is the object contain the payment data returned from the API.
type Payment struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Type of the payment object.
	// options: payment, refund, cancel
	Type string `json:"type"`

	// Unique system generated identifier for the merchant.
	MerchantID string `json:"merchantId"`

	// Unique system generated identifier for the wallet of the merchant.
	MerchantWalletID string `json:"merchantWalletId"`

	// Amount object for the payment
	Amount *BalanceEntry `json:"amount,omitempty"`

	// The payment source.
	Source *Source `json:"source,omitempty"`

	// Enumerated description of the payment.
	Description string `json:"description,omitempty"`

	// Enumerated status of the payment.
	// pending means the payment is waiting to be processed.
	// confirmed means the payment has been approved by the bank and the merchant can treat it as successful,
	// but settlement funds are not yet available to the merchant.
	// paid means settlement funds have been received and are available to the merchant.
	// failed means something went wrong (most commonly that the payment was denied).
	// action_required means that additional steps are required to process this payment;
	// refer to requiredAction for more details.
	// Terminal states are paid and failed.
	Status string `json:"status,omitempty"`

	// Determines if a payment has successfully been captured.
	// This property is only present for payments that did not use auto capture.
	Captured *bool `json:"captured,omitempty"`

	CaptureAmount *BalanceEntry `json:"captureAmount,omitempty"`

	// ISO-8601 UTC date/time format.
	CaptureDate string `json:"captureDate,omitempty"`

	// When the payment status is action_required, this object summarizes the required additional steps.
	RequiredAction *RequiredAction `json:"requiredAction,omitempty"`

	// Indicates the status of the payment verification. This property will be present once the payment is confirmed.
	Verification *PaymentVerification `json:"verification,omitempty"`

	// Fees object for the payment
	Fees *BalanceEntry `json:"fees,omitempty"`

	// Payment tracking reference. Will be present once known.
	TrackingRef string `json:"trackingRef,omitempty"`

	// External network identifier which will be present once provided from the applicable network.
	ExternalRef string `json:"externalRef,omitempty"`

	// Indicates the failure reason of a payment. Only present for payments in failed state.
	// Possible values are [payment_failed, payment_fraud_detected, payment_denied,
	// payment_not_supported_by_issuer, payment_not_funded, payment_unprocessable,
	// payment_stopped_by_issuer, payment_canceled, payment_returned, payment_failed_balance_check,
	// card_failed, card_invalid, card_address_mismatch, card_zip_mismatch, card_cvv_invalid,
	// card_expired, card_limit_violated, card_not_honored, card_cvv_required, credit_card_not_allowed,
	// card_account_ineligible, card_network_unsupported, channel_invalid, unauthorized_transaction,
	// bank_account_ineligible, bank_transaction_error, invalid_account_number, invalid_wire_rtn,
	// invalid_ach_rtn, vendor_inactive]'
	ErrorCode string `json:"errorCode,omitempty"`

	// Object containing metadata for the payment
	Metadata *Metadata `json:"metadata,omitempty"`

	// Results of risk evaluation. Only present if the payment is denied by Circle's risk service.
	RiskEvaluation *RiskEvaluation `json:"riskEvaluation,omitempty"`

	// The channel identifier that can be set for the payment. When not provided, the default channel is used.
	Channel string `json:"channel,omitempty"`

	// ISO-8601 UTC date/time format of the payment creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the payment update date.
	UpdateDate string `json:"updateDate,omitempty"`

	// Status information of the related cancel. This property is only present on canceled payment or refund items.
	Cancel *Payment `json:"cancel,omitempty"`

	// Status information of the related payment. This property is only present on refund or cancel items.
	OriginalPayment *Payment `json:"originalPayment,omitempty"`

	// Array of refunded payments.
	Refunds []Payment `json:"refunds,omitempty"`
}

// ListPaymentsRequest contains the data to list payments.
type ListPaymentsRequest struct {
	// Universally unique identifier (UUID v4) for the source.
	// Filter results to fetch only payments made from the provided source.
	Source string `json:"source,omitempty"`

	// Queries items with the specified settlement id. Matches any settlement id if unspecified.
	SettlementID string `json:"settlementId,omitempty"`

	// Source account type. Filters the results to fetch all payments made from a specified account type.
	// Matches any source type if unspecified.
	Type []string `json:"type,omitempty"`

	// Queries items with the specified status. Matches any status if unspecified.
	Status string `json:"status,omitempty"`
}

// CreatePaymentRequest contains the data to create a payment.
type CreatePaymentRequest struct {
	// Universally unique identifier (UUID v4) idempotency key.
	// This key is utilized to ensure exactly-once execution of mutating requests.
	IdempotencyKey string `json:"idempotencyKey,omitempty"`

	// Universally unique identifier (UUID v4) of the public key used in encryption.
	// NOTE the sandbox environment uses the default value of key1.
	// For this reason the example supplied is key1 rather than a UUID.
	KeyID string `json:"keyId,omitempty"`

	// Object containing metadata for the payment creation process
	Metadata *CreateMetadataRequest `json:"metadata,omitempty"`

	// Amount object for the payment
	Amount *BalanceEntry `json:"amount,omitempty"`

	// Triggers the automatic capture of the full payment amount.
	// If set to false the payment will only be authorized but not captured.
	AutoCapture bool `json:"autoCapture"`

	// Indicates the verification method for this payment.
	Verification string `json:"verification,omitempty"`

	// The URL to redirect users to after successful 3DS authentication.
	VerificationSuccessURL string `json:"verificationSuccessUrl,omitempty"`

	// The URL to redirect users to after failed 3DS authentication.
	VerificationFailureURL string `json:"verificationFailureUrl,omitempty"`

	// Source object used for the payment
	Source *Source `json:"source,omitempty"`

	// Description of the payment with length restriction of 240 characters.
	Description string `json:"description,omitempty"`

	// PGP encrypted base64 encoded string. Contains CVV.
	EncryptedData string `json:"encryptedData,omitempty"`

	// The channel identifier that can be set for the payment. When not provided, the default channel is used.
	Channel string `json:"channel,omitempty"`
}

// CapturePaymentRequest contains the data to capture a payment.
type CapturePaymentRequest struct {
	// Universally unique identifier (UUID v4) idempotency key.
	// This key is utilized to ensure exactly-once execution of mutating requests.
	IdempotencyKey string `json:"idempotencyKey,omitempty"`

	// Amount object for the payment capture
	Amount *BalanceEntry `json:"amount,omitempty"`
}

// CancelPaymentRequest contains the data to cancel a payment.
type CancelPaymentRequest struct {
	// Universally unique identifier (UUID v4) idempotency key.
	// This key is utilized to ensure exactly-once execution of mutating requests.
	IdempotencyKey string `json:"idempotencyKey,omitempty"`

	// Enumerated reason for a returned payment.
	// Providing this reason in the request is recommended (to improve risk evaluation) but not required.
	Reason string `json:"reason,omitempty"`
}

// RefundPaymentRequest contains the data to refund a payment.
type RefundPaymentRequest struct {
	// Universally unique identifier (UUID v4) idempotency key.
	// This key is utilized to ensure exactly-once execution of mutating requests.
	IdempotencyKey string `json:"idempotencyKey,omitempty"`

	// Amount object for the payment capture
	Amount *BalanceEntry `json:"amount,omitempty"`

	// Enumerated reason for a returned payment.
	// Providing this reason in the request is recommended (to improve risk evaluation) but not required.
	Reason string `json:"reason,omitempty"`
}

// PayoutDestination contains the bank account details.
type PayoutDestination struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// The destination bank account type.
	// options: wire, ach, sepa
	Type string `json:"type,omitempty"`

	// Bank name plus last four digits of the bank account number or IBAN.
	Name string `json:"name,omitempty"`
}

// PayoutAdjustment contains information about increases (credits) or decreases (debits)
// the total returned amount to the source wallet.
type PayoutAdjustment struct {
	// Credit object for the adjustment
	FxCredit *BalanceEntry `json:"fxCredit,omitempty"`

	// Debit object for the adjustment
	FxDebit *BalanceEntry `json:"fxDebit,omitempty"`
}

// PayoutReturn contains data if the payout is returned by the bank.
type PayoutReturn struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Universally unique identifier (UUID v4) of the payout that is associated with the return.
	PayoutID string `json:"payoutId,omitempty"`

	// Amount object for the return
	Amount *BalanceEntry `json:"amount,omitempty"`

	// Fees object for the return
	Fees *BalanceEntry `json:"fees,omitempty"`

	// Reason for the return.
	Reason string `json:"reason,omitempty"`

	// Status of the return. A pending status indicates that the return is in process;
	// complete indicates it finished successfully;
	// failed indicates it failed.
	Status string `json:"status,omitempty"`

	// ISO-8601 UTC date/time format of the return creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the return update date.
	UpdateDate string `json:"updateDate,omitempty"`
}

// Payout is the object contain the payout data returned from the API.
type Payout struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// The identifier of the source wallet used to fund a payout.
	SourceWalletID string `json:"sourceWalletId,omitempty"`

	// The destination bank account.
	Destination *PayoutDestination `json:"destination,omitempty"`

	// Amount object for the payout
	Amount *BalanceEntry `json:"amount,omitempty"`

	// Fees object for the payout
	Fees *BalanceEntry `json:"fees,omitempty"`

	// Status of the payout. Status pending indicates that the payout is in process;
	// complete indicates it finished successfully;
	// failed indicates it failed.
	Status string `json:"status,omitempty"`

	// A payout tracking reference. Will be present once known.
	TrackingRef string `json:"trackingRef,omitempty"`

	// External network identifier which will be present once provided from the applicable network.
	ExternalRef string `json:"externalRef,omitempty"`

	// Indicates the failure reason of a payout. Only present for payouts in failed state.
	// Possible values are [insufficient_funds, transaction_denied, transaction_failed,
	// transaction_returned, bank_transaction_error, fiat_account_limit_exceeded, invalid_bank_account_number,
	// invalid_ach_rtn, invalid_wire_rtn, vendor_inactive]'.
	ErrorCode string `json:"errorCode,omitempty"`

	// Results of risk evaluation. Only present if the payment is denied by Circle's risk service.
	RiskEvaluation *RiskEvaluation `json:"riskEvaluation,omitempty"`

	// Final adjustment which increases (credits) or decreases (debits) the total returned amount to the source wallet.
	Adjustments *PayoutAdjustment `json:"adjustments,omitempty"`

	// Return information if the payout is returned by bank.
	// Only present if errorCode of payout is transaction_returned.
	Return *PayoutReturn `json:"return,omitempty"`

	// ISO-8601 UTC date/time format of the payout creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the payout update date.
	UpdateDate string `json:"updateDate,omitempty"`
}

// ListPayoutsRequest contains the data to list payouts.
type ListPayoutsRequest struct {
	// Universally unique identifier (UUID v4) for the source wallet. Filters the results
	// to fetch all payouts made from a source wallet. If not provided,
	// payouts from all wallets will be returned.
	Source string `json:"source,omitempty"`

	// Destination bank account type. Filters the results to fetch all payouts made to a specified
	// destination bank account type. This query parameter can be passed multiple times to fetch results
	// matching multiple destination bank account types.
	Type []string `json:"type,omitempty"`

	// Queries items with the specified status. Matches any status if unspecified.
	Status []string `json:"status,omitempty"`

	// Universally unique identifier (UUID v4) for the destination bank account.
	// Filters the results to fetch all payouts made to a destination bank account.
	Destination string `json:"destination,omitempty"`
}

// CreatePayoutMetadataRequest contains data related to the payout beneficiary.
type CreatePayoutMetadataRequest struct {
	// Email of the user.
	BeneficiaryEmail string `json:"beneficiaryEmail,omitempty"`
}

// CreatePayoutRequest contains the data to create a payout.
type CreatePayoutRequest struct {
	// Universally unique identifier (UUID v4) idempotency key.
	// This key is utilized to ensure exactly-once execution of mutating requests.
	IdempotencyKey string `json:"idempotencyKey,omitempty"`

	// Source object for the payout
	Source *Source `json:"source,omitempty"`

	// The destination bank account.
	PayoutDestination *Source `json:"payoutDestination,omitempty"`

	// Amount object for the payout
	Amount *BalanceEntry `json:"amount,omitempty"`

	// Additional properties related to the payout beneficiary.
	Metadata *CreatePayoutMetadataRequest `json:"metadata,omitempty"`
}

// Settlement is the object contain the settlement data returned from the API.
type Settlement struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// If this settlement was used for a marketplace payment, the wallet involved in the settlement.
	// Not included for standard merchant settlements.
	MerchantWalletID string `json:"merchantWalletId,omitempty"`

	// Total debits for the settlement
	TotalDebits *BalanceEntry `json:"totalDebits,omitempty"`

	// Total credits for the settlement
	TotalCredits *BalanceEntry `json:"totalCredits,omitempty"`

	// Payment fees for the settlement
	PaymentFees *BalanceEntry `json:"paymentFees,omitempty"`

	// Chargeback fees for the settlement
	ChargebackFees *BalanceEntry `json:"chargebackFees,omitempty"`

	// ISO-8601 UTC date/time format of the settlement creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the settlement update date.
	UpdateDate string `json:"updateDate,omitempty"`
}

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
	ChargeBackAmount *BalanceEntry `json:"chargeBackAmount,omitempty"`

	// Fee object for the history
	Fee *BalanceEntry `json:"fee,omitempty"`

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

// Reversal is the object contain the reversal data returned from the API.
type Reversal struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Unique system generated identifier for the payment that is associated to the chargeback item.
	PaymentID string `json:"paymentId,omitempty"`

	// Amount object of the reversal
	Amount *BalanceEntry `json:"amount,omitempty"`

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
	Fees *BalanceEntry `json:"fees,omitempty"`

	// ISO-8601 UTC date/time format of the reversal creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the reversal update date.
	UpdateDate string `json:"updateDate,omitempty"`
}
