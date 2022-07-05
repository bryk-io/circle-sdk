package circlesdk

// Balance of funds that are available for use.
type Balance struct {
	// List of currency balances (one for each currency) that are currently available
	// to spend.
	Available []Amount `json:"available,omitempty"`

	// List of currency balances (one for each currency) that have been captured but are
	// currently in the process of settling and will become available to spend at some
	// point in the future.
	Unsettled []Amount `json:"unsettled,omitempty"`
}
