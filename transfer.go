package circlesdk

// TransferStatus contains the status value for the transfer.
type TransferStatus string

const (
	// TransferStatusPending = "pending".
	TransferStatusPending TransferStatus = "pending"

	// TransferStatusComplete = "complete".
	TransferStatusComplete TransferStatus = "complete"

	// TransferStatusFailed = "failed".
	TransferStatusFailed TransferStatus = "failed"
)

// TransferErrorCode contains the error code value for the transfer.
type TransferErrorCode string

const (
	// TransferErrorCodeInsufficientFunds = "insufficient_funds".
	TransferErrorCodeInsufficientFunds TransferErrorCode = "insufficient_funds"

	// TransferErrorCodeBlockChainError = "blockchain_error".
	TransferErrorCodeBlockChainError TransferErrorCode = "blockchain_error"

	// TransferErrorCodeTransferDenied = "transfer_denied".
	TransferErrorCodeTransferDenied TransferErrorCode = "transfer_denied"

	// TransferErrorCodeTransferFailed = "transfer_failed".
	TransferErrorCodeTransferFailed TransferErrorCode = "transfer_failed"
)

// Transfer of funds between a `source` and `destination` addresses.
type Transfer struct {
	// Unique identifier for this transfer.
	ID string `json:"id,omitempty"`

	// Source of the funds.
	Source *DepositAddress `json:"source,omitempty"`

	// Destination of the funds.
	Destination *DepositAddress `json:"destination,omitempty"`

	// Nominal value transferred.
	Amount *Amount `json:"amount,omitempty"`

	// A hash that uniquely identifies the on-chain transaction. This is only
	// available where either source or destination are of type blockchain.
	TxHash string `json:"transactionHash,omitempty"`

	// Status of the transfer. Status `pending` indicates that the transfer is in
	// the process of running; `complete` indicates it finished successfully;
	// `failed` indicates it failed.
	Status TransferStatus `json:"status,omitempty"`

	// Indicates the failure reason of a transfer. Only present for transfers in a
	// failed state. Possible values are `insufficient_funds`, `blockchain_error`,
	// `transfer_denied` and `transfer_failed`.
	ErrorCode TransferErrorCode `json:"errorCode,omitempty"`

	// The creation date of the transfer.
	CreateDate string `json:"createDate,omitempty"`
}
