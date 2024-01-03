package circlesdk

// RequiredActionType contains the type value for the required action.
type RequiredActionType string

// RequiredActionTypeThreeDSecureRequired = "three_d_secure_required".
const RequiredActionTypeThreeDSecureRequired RequiredActionType = "three_d_secure_required"

// RequiredAction indicates when the payment status is action_required,
// this object summarizes the required additional steps.
type RequiredAction struct {
	// The type of action that is required to proceed with the payment. Currently only one type is supported.
	Type RequiredActionType `json:"type,omitempty"`

	// The URL to bring the user to in order to complete the payment.
	RedirectURL string `json:"redirectUrl,omitempty"`
}

// PaymentVerificationAvs contains the avs value for the payment verification.
type PaymentVerificationAvs string

const (
	// PaymentVerificationAvsNotRequested = "not_requested".
	PaymentVerificationAvsNotRequested PaymentVerificationAvs = "not_requested"

	// PaymentVerificationAvsPending = "pending".
	PaymentVerificationAvsPending PaymentVerificationAvs = "pending"
)

// PaymentVerificationCvv contains the cvv value for the payment verification.
type PaymentVerificationCvv string

const (
	// PaymentVerificationCvvNotRequested = "not_requested".
	PaymentVerificationCvvNotRequested PaymentVerificationCvv = "not_requested"

	// PaymentVerificationCvvPass = "pass".
	PaymentVerificationCvvPass PaymentVerificationCvv = "pass"

	// PaymentVerificationCvvFail = "fail".
	PaymentVerificationCvvFail PaymentVerificationCvv = "fail"

	// PaymentVerificationCvvUnavailable = "unavailable".
	PaymentVerificationCvvUnavailable PaymentVerificationCvv = "unavailable"

	// PaymentVerificationCvvPending = "pending".
	PaymentVerificationCvvPending PaymentVerificationCvv = "pending"
)

// PaymentVerificationThreeDSecure contains the three d secure value for the payment verification.
type PaymentVerificationThreeDSecure string

const (
	//PaymentVerificationThreeDSecurePass = "pass".
	PaymentVerificationThreeDSecurePass PaymentVerificationThreeDSecure = "pass"

	// PaymentVerificationThreeDSecureFail = "fail".
	PaymentVerificationThreeDSecureFail PaymentVerificationThreeDSecure = "fail"
)

// PaymentVerificationEci contains the eci value for the payment verification.
type PaymentVerificationEci string

const (
	// PaymentVerificationEci00 = "00".
	PaymentVerificationEci00 PaymentVerificationEci = "00"

	// PaymentVerificationEci01 = "01".
	PaymentVerificationEci01 PaymentVerificationEci = "01"

	// PaymentVerificationEci02 = "02".
	PaymentVerificationEci02 PaymentVerificationEci = "02"

	// PaymentVerificationEci05 = "05".
	PaymentVerificationEci05 PaymentVerificationEci = "05"

	// PaymentVerificationEci06 = "06".
	PaymentVerificationEci06 PaymentVerificationEci = "06"

	// PaymentVerificationEci07 = "07".
	PaymentVerificationEci07 PaymentVerificationEci = "07"
)

// PaymentVerification indicates the status of the payment verification.
// This property will be present once the payment is confirmed.
type PaymentVerification struct {
	// Status of the AVS check. Raw AVS response, expressed as an upper-case letter.
	// not_requested indicates check was not made.
	// pending is pending/processing.
	Avs PaymentVerificationAvs `json:"avs,omitempty"`

	// Enumerated status of the check.
	// not_requested indicates check was not made.
	// pass indicates value is correct.
	// fail indicates value is incorrect.
	// unavailable indicates card issuer did not do the provided check.
	// pending indicates check is pending/processing.
	Cvv PaymentVerificationCvv `json:"cvv,omitempty"`

	// Enumerated status of the check.
	// pass indicates successful 3DS authentication.
	// fail indicates failed 3DS authentication.
	ThreeDSecure PaymentVerificationThreeDSecure `json:"threeDSecure"`

	// ECI (electronic commerce indicator) value returned by Directory Servers
	// (namely Visa, MasterCard, JCB, and American Express) indicating the outcome
	// of authentication attempted on transactions enforced by 3DS.
	Eci PaymentVerificationEci `json:"eci,omitempty"`
}

// PaymentType contains the type value for the payment.
type PaymentType string

const (
	// PaymentTypePayment = "payment".
	PaymentTypePayment PaymentType = "payment"

	// PaymentTypeRefund = "refund".
	PaymentTypeRefund PaymentType = "refund"

	// PaymentTypeCancel = "cancel".
	PaymentTypeCancel PaymentType = "cancel"
)

// PaymentStatus contains the status value for the payment.
type PaymentStatus string

const (
	// PaymentStatusPending = "pending".
	PaymentStatusPending PaymentStatus = "pending"

	// PaymentStatusConfirmed = "confirmed".
	PaymentStatusConfirmed PaymentStatus = "confirmed"

	// PaymentStatusPaid = "paid".
	PaymentStatusPaid PaymentStatus = "paid"

	// PaymentStatusFailed = "failed".
	PaymentStatusFailed PaymentStatus = "failed"

	// PaymentStatusActionRequired = "action_required".
	PaymentStatusActionRequired PaymentStatus = "action_required"
)

// PaymentErrorCode contains the error code value for the payment.
type PaymentErrorCode string

// nolint: gosec
const (
	// PaymentErrorCodePaymentFailed = "payment_failed".
	PaymentErrorCodePaymentFailed PaymentErrorCode = "payment_failed"

	// PaymentErrorCodePaymentFraudDetected = "payment_fraud_detected".
	PaymentErrorCodePaymentFraudDetected PaymentErrorCode = "payment_fraud_detected"

	// PaymentErrorCodePaymentDenied = "payment_denied".
	PaymentErrorCodePaymentDenied PaymentErrorCode = "payment_denied"

	// PaymentErrorCodePaymentNotSupportedByIssuer = "payment_not_supported_by_issuer".
	PaymentErrorCodePaymentNotSupportedByIssuer PaymentErrorCode = "payment_not_supported_by_issuer"

	// PaymentErrorCodePaymentNotFound = "payment_not_funded".
	PaymentErrorCodePaymentNotFound PaymentErrorCode = "payment_not_funded"

	// PaymentErrorCodePaymentUnprocessable = "payment_unprocessable".
	PaymentErrorCodePaymentUnprocessable PaymentErrorCode = "payment_unprocessable"

	// PaymentErrorCodePaymentStoppedByIssuer = "payment_stopped_by_issuer".
	PaymentErrorCodePaymentStoppedByIssuer PaymentErrorCode = "payment_stopped_by_issuer"

	// PaymentErrorCodePaymentCanceled = "payment_canceled".
	PaymentErrorCodePaymentCanceled PaymentErrorCode = "payment_canceled"

	// PaymentErrorCodePaymentReturned = "payment_returned".
	PaymentErrorCodePaymentReturned PaymentErrorCode = "payment_returned"

	// PaymentErrorCodePaymentFailedBalanceCheck = "payment_failed_balance_check".
	PaymentErrorCodePaymentFailedBalanceCheck PaymentErrorCode = "payment_failed_balance_check"

	// PaymentErrorCodeCardFailed = "card_failed".
	PaymentErrorCodeCardFailed PaymentErrorCode = "card_failed"

	// PaymentErrorCodeCardInvalid = "card_invalid".
	PaymentErrorCodeCardInvalid PaymentErrorCode = "card_invalid"

	// PaymentErrorCodeCardAddressMismatch = "card_address_mismatch".
	PaymentErrorCodeCardAddressMismatch PaymentErrorCode = "card_address_mismatch"

	// PaymentErrorCodeCardZipMismatch = "card_zip_mismatch".
	PaymentErrorCodeCardZipMismatch PaymentErrorCode = "card_zip_mismatch"

	// PaymentErrorCodeCardCvvInvalid = "card_cvv_invalid".
	PaymentErrorCodeCardCvvInvalid PaymentErrorCode = "card_cvv_invalid"

	// PaymentErrorCodeCardExpired = "card_expired".
	PaymentErrorCodeCardExpired PaymentErrorCode = "card_expired"

	// PaymentErrorCodeCardLimitViolated = "card_limit_violated".
	PaymentErrorCodeCardLimitViolated PaymentErrorCode = "card_limit_violated"

	// PaymentErrorCodeCardNotHonored = "card_not_honored".
	PaymentErrorCodeCardNotHonored PaymentErrorCode = "card_not_honored"

	// PaymentErrorCodeCardCvvRequired = "card_cvv_required".
	PaymentErrorCodeCardCvvRequired PaymentErrorCode = "card_cvv_required"

	// PaymentErrorCodeCreditCardNotAllowed = "credit_card_not_allowed".
	PaymentErrorCodeCreditCardNotAllowed PaymentErrorCode = "credit_card_not_allowed"

	// PaymentErrorCodeCardAccountIneligible = "card_account_ineligible".
	PaymentErrorCodeCardAccountIneligible PaymentErrorCode = "card_account_ineligible"

	// PaymentErrorCodeCardNetworkUnsupported = "card_network_unsupported".
	PaymentErrorCodeCardNetworkUnsupported PaymentErrorCode = "card_network_unsupported"

	// PaymentErrorCodeChannelInvalid = "channel_invalid".
	PaymentErrorCodeChannelInvalid PaymentErrorCode = "channel_invalid"

	// PaymentErrorCodeUnauthorized = "unauthorized_transaction".
	PaymentErrorCodeUnauthorized PaymentErrorCode = "unauthorized_transaction"

	// PaymentErrorCodeBankAccountIneligible = "bank_account_ineligible".
	PaymentErrorCodeBankAccountIneligible PaymentErrorCode = "bank_account_ineligible"

	// PaymentErrorCodeBankTransactionError = "bank_transaction_error".
	PaymentErrorCodeBankTransactionError PaymentErrorCode = "bank_transaction_error"

	// PaymentErrorCodeInvalidAccountNumber = "invalid_account_number".
	PaymentErrorCodeInvalidAccountNumber PaymentErrorCode = "invalid_account_number"

	// PaymentErrorCodeInvalidWireRtn = "invalid_wire_rtn".
	PaymentErrorCodeInvalidWireRtn PaymentErrorCode = "invalid_wire_rtn"

	// PaymentErrorCodeInvalidAchRtn = "invalid_ach_rtn".
	PaymentErrorCodeInvalidAchRtn PaymentErrorCode = "invalid_ach_rtn"

	// PaymentErrorCodeVendorInactive = "vendor_inactive".
	PaymentErrorCodeVendorInactive PaymentErrorCode = "vendor_inactive"
)

// Payment is the object contain the payment data returned from the API.
type Payment struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Type of the payment object.
	// options: payment, refund, cancel
	Type PaymentType `json:"type"`

	// Unique system generated identifier for the merchant.
	MerchantID string `json:"merchantId"`

	// Unique system generated identifier for the wallet of the merchant.
	MerchantWalletID string `json:"merchantWalletId"`

	// Amount object for the payment
	Amount *Amount `json:"amount,omitempty"`

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
	Status PaymentStatus `json:"status,omitempty"`

	// Determines if a payment has successfully been captured.
	// This property is only present for payments that did not use auto capture.
	Captured *bool `json:"captured,omitempty"`

	CaptureAmount *Amount `json:"captureAmount,omitempty"`

	// ISO-8601 UTC date/time format.
	CaptureDate string `json:"captureDate,omitempty"`

	// When the payment status is action_required, this object summarizes the required additional steps.
	RequiredAction *RequiredAction `json:"requiredAction,omitempty"`

	// Indicates the status of the payment verification. This property will be present once the payment is confirmed.
	Verification *PaymentVerification `json:"verification,omitempty"`

	// Fees object for the payment
	Fees *Amount `json:"fees,omitempty"`

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
	ErrorCode PaymentErrorCode `json:"errorCode,omitempty"`

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
	Status PaymentStatus `json:"status,omitempty"`
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
	Amount *Amount `json:"amount,omitempty"`

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
	Amount *Amount `json:"amount,omitempty"`
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
	Amount *Amount `json:"amount,omitempty"`

	// Enumerated reason for a returned payment.
	// Providing this reason in the request is recommended (to improve risk evaluation) but not required.
	Reason string `json:"reason,omitempty"`
}
