package circlesdk

// Amount provides information for a specific amount/currency pair.
type Amount struct {
	// Magnitude of the amount, in units of the currency.
	Amount string `json:"amount,omitempty"`

	// Currency code for the amount.
	Currency string `json:"currency,omitempty"`
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

// Metadata is the object containing metadata for the entity.
type Metadata struct {
	// Email of the user.
	Email string `json:"email,omitempty"`

	// Phone number of the user in E.164 format.
	// We recommend using a library such as libphonenumber to parse and validate phone numbers.
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

// RiskEvaluationDecision contains decision value for the RiskEvaluation.
type RiskEvaluationDecision string

const (
	// RiskEvaluationDecisionApproved = "approved".
	RiskEvaluationDecisionApproved RiskEvaluationDecision = "approved"

	// RiskEvaluationDecisionDenied = "denied".
	RiskEvaluationDecisionDenied RiskEvaluationDecision = "denied"

	// RiskEvaluationDecisionReview = "review".
	RiskEvaluationDecisionReview RiskEvaluationDecision = "review"
)

// RiskEvaluation contains the Result of risk evaluation.
// Only present if the payment is denied by Circle's risk service.
type RiskEvaluation struct {
	// Enumerated decision of the account.
	// Options: approved, denied, review
	Decision RiskEvaluationDecision `json:"decision,omitempty"`

	// Risk reason for the definitive decision outcome.
	Reason string `json:"reason,omitempty"`
}

// SourceType contains the type value for the source.
type SourceType string

const (
	// SourceTypeCard = "card".
	SourceTypeCard SourceType = "card"

	// SourceTypeAch = "ach".
	SourceTypeAch SourceType = "ach"

	// SourceTypeWire = "wire".
	SourceTypeWire SourceType = "wire"

	// SourceTypeSepa = "sepa".
	SourceTypeSepa SourceType = "sepa"
)

// Source object used for the payment.
type Source struct {
	// Unique system generated identifier for the payment item.
	ID string `json:"id"`

	// Type of the source.
	// options: card, ach, wire, sepa
	Type SourceType `json:"type"`
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
