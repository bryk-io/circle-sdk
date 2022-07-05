package circlesdk

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
	Status string `json:"status,omitempty"`

	// Indicates the failure reason of a transfer. Only present for transfers in a
	// failed state. Possible values are `insufficient_funds`, `blockchain_error`,
	// `transfer_denied` and `transfer_failed`.
	ErrorCode string `json:"errorCode,omitempty"`

	// The creation date of the transfer.
	CreateDate string `json:"createDate,omitempty"`
}
