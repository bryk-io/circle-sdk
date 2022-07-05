package circlesdk

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
	Status string `json:"status,omitempty"`

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
