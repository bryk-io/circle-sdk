package circlesdk

// PayoutDestinationType contains the type value for the payout destination.
type PayoutDestinationType string

const (
	// PayoutDestinationTypeWire = "wire".
	PayoutDestinationTypeWire PayoutDestinationType = "wire"

	// PayoutDestinationTypeAch = "ach".
	PayoutDestinationTypeAch PayoutDestinationType = "ach"

	// PayoutDestinationTypeSepa = "sepa".
	PayoutDestinationTypeSepa PayoutDestinationType = "sepa"
)

// PayoutDestination contains the bank account details.
type PayoutDestination struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// The destination bank account type.
	// options: wire, ach, sepa
	Type PayoutDestinationType `json:"type,omitempty"`

	// Bank name plus last four digits of the bank account number or IBAN.
	Name string `json:"name,omitempty"`
}

// PayoutAdjustment contains information about increases (credits) or decreases (debits)
// the total returned amount to the source wallet.
type PayoutAdjustment struct {
	// Credit object for the adjustment
	FxCredit *Amount `json:"fxCredit,omitempty"`

	// Debit object for the adjustment
	FxDebit *Amount `json:"fxDebit,omitempty"`
}

// PayoutReturnStatus contains the status value for the payout return.
type PayoutReturnStatus string

const (
	// PayoutReturnStatusPending = "pending".
	PayoutReturnStatusPending PayoutReturnStatus = "pending"

	// PayoutReturnStatusComplete = "complete".
	PayoutReturnStatusComplete PayoutReturnStatus = "complete"

	//PayoutReturnStatusFailed = "failed".
	PayoutReturnStatusFailed PayoutReturnStatus = "failed"
)

// PayoutReturn contains data if the payout is returned by the bank.
type PayoutReturn struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// Universally unique identifier (UUID v4) of the payout that is associated with the return.
	PayoutID string `json:"payoutId,omitempty"`

	// Amount object for the return
	Amount *Amount `json:"amount,omitempty"`

	// Fees object for the return
	Fees *Amount `json:"fees,omitempty"`

	// Reason for the return.
	Reason string `json:"reason,omitempty"`

	// Status of the return. A pending status indicates that the return is in process;
	// complete indicates it finished successfully;
	// failed indicates it failed.
	Status PayoutReturnStatus `json:"status,omitempty"`

	// ISO-8601 UTC date/time format of the return creation date.
	CreateDate string `json:"createDate,omitempty"`

	// ISO-8601 UTC date/time format of the return update date.
	UpdateDate string `json:"updateDate,omitempty"`
}

// PayoutStatus contains the status value for the payout.
type PayoutStatus string

const (
	// PayoutStatusPending = "pending".
	PayoutStatusPending PayoutStatus = "pending"

	// PayoutStatusComplete = "complete".
	PayoutStatusComplete PayoutStatus = "complete"

	// PayoutStatusFailed = "failed".
	PayoutStatusFailed PayoutStatus = "failed"
)

// PayoutErrorCode contains the error code value for the payout.
type PayoutErrorCode string

const (
	// PayoutErrorCodeInsufficientFunds = "insufficient_funds".
	PayoutErrorCodeInsufficientFunds PayoutErrorCode = "insufficient_funds"

	// PayoutErrorCodeTransactionDenied = "transaction_denied".
	PayoutErrorCodeTransactionDenied PayoutErrorCode = "transaction_denied"

	// PayoutErrorCodeTransactionFailed = "transaction_failed".
	PayoutErrorCodeTransactionFailed PayoutErrorCode = "transaction_failed"

	// PayoutErrorCodeTransactionReturned = "transaction_returned".
	PayoutErrorCodeTransactionReturned PayoutErrorCode = "transaction_returned"

	// PayoutErrorCodeBankTransactionError = "bank_transaction_error".
	PayoutErrorCodeBankTransactionError PayoutErrorCode = "bank_transaction_error"

	// PayoutErrorCodeFiatAccountLimitExceed = "fiat_account_limit_exceeded".
	PayoutErrorCodeFiatAccountLimitExceed PayoutErrorCode = "fiat_account_limit_exceeded"

	// PayoutErrorCodeInvalidBankAccountNumber = "invalid_bank_account_number".
	PayoutErrorCodeInvalidBankAccountNumber PayoutErrorCode = "invalid_bank_account_number"

	// PayoutErrorCodeInvalidAchRtn = "invalid_ach_rtn".
	PayoutErrorCodeInvalidAchRtn PayoutErrorCode = "invalid_ach_rtn"

	// PayoutErrorCodeInvalidWireRtn = "invalid_wire_rtn".
	PayoutErrorCodeInvalidWireRtn PayoutErrorCode = "invalid_wire_rtn"

	// PayoutErrorCodeVendorInactive = "vendor_inactive".
	PayoutErrorCodeVendorInactive PayoutErrorCode = "vendor_inactive"
)

// Payout is the object contain the payout data returned from the API.
type Payout struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id,omitempty"`

	// The identifier of the source wallet used to fund a payout.
	SourceWalletID string `json:"sourceWalletId,omitempty"`

	// The destination bank account.
	Destination *PayoutDestination `json:"destination,omitempty"`

	// Amount object for the payout
	Amount *Amount `json:"amount,omitempty"`

	// Fees object for the payout
	Fees *Amount `json:"fees,omitempty"`

	// Status of the payout. Status pending indicates that the payout is in process;
	// complete indicates it finished successfully;
	// failed indicates it failed.
	Status PayoutStatus `json:"status,omitempty"`

	// A payout tracking reference. Will be present once known.
	TrackingRef string `json:"trackingRef,omitempty"`

	// External network identifier which will be present once provided from the applicable network.
	ExternalRef string `json:"externalRef,omitempty"`

	// Indicates the failure reason of a payout. Only present for payouts in failed state.
	// Possible values are [insufficient_funds, transaction_denied, transaction_failed,
	// transaction_returned, bank_transaction_error, fiat_account_limit_exceeded, invalid_bank_account_number,
	// invalid_ach_rtn, invalid_wire_rtn, vendor_inactive]'.
	ErrorCode PayoutErrorCode `json:"errorCode,omitempty"`

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
	Status []PayoutStatus `json:"status,omitempty"`

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
	Amount *Amount `json:"amount,omitempty"`

	// Additional properties related to the payout beneficiary.
	Metadata *CreatePayoutMetadataRequest `json:"metadata,omitempty"`
}
