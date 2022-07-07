package payments

import (
	"errors"
	circlesdk "github.com/bryk-io/circle-sdk"
	ac "github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripper struct {
	response *http.Response
	err      error
}

func (r *roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.response, r.err
}

func TestAPI_CreateCard(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		roundTripper roundTripper
		expected     *circlesdk.Card
		expectedErr  error
	}{
		{
			name: "internal error",
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Post "v1/cards": some error`),
		}, {
			name: "api 500 error",
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "Internal error."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: Internal error."),
		}, {
			name: "api 401 error",
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "Malformed authorization."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: Malformed authorization."),
		}, {
			name: "api 400 error",
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "400",
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 400,
  "message": "Bad request."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("400: Bad request."),
		}, {
			name: "success",
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "201",
					StatusCode: http.StatusCreated,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "status": "pending",
    "billingDetails": {
      "name": "Satoshi Nakamoto",
      "city": "Boston",
      "country": "US",
      "line1": "100 Money Street",
      "line2": "Suite 1",
      "district": "MA",
      "postalCode": "01234"
    },
    "expMonth": 1,
    "expYear": 2020,
    "network": "VISA",
    "last4": "0123",
    "bin": "401230",
    "issuerCountry": "US",
    "fundingType": "credit",
    "fingerprint": "eb170539-9e1c-4e92-bf4f-1d09534fdca2",
    "errorCode": "verification_failed",
    "verification": {
      "avs": "not_requested",
      "cvv": "not_requested"
    },
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "metadata": {
      "email": "satoshi@circle.com",
      "phoneNumber": "+14155555555"
    },
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.Card{
				ID:     "b8627ae8-732b-4d25-b947-1df8f4007a29",
				Status: circlesdk.CardStatusPending,
				BillingDetails: &circlesdk.BillingDetails{
					Name:       "Satoshi Nakamoto",
					City:       "Boston",
					Country:    "US",
					Line1:      "100 Money Street",
					Line2:      "Suite 1",
					District:   "MA",
					PostalCode: "01234",
				},
				ExpMonth:      1,
				ExpYear:       2020,
				Network:       circlesdk.CardNetworkVISA,
				Last4:         "0123",
				Bin:           "401230",
				IssuerCountry: "US",
				FundingType:   circlesdk.CardFundingTypeCredit,
				Fingerprint:   "eb170539-9e1c-4e92-bf4f-1d09534fdca2",
				ErrorCode:     circlesdk.CardErrorCodeVerificationFailed,
				Verification: &circlesdk.CardVerification{
					Avs: circlesdk.CardVerificationAvsNotRequested,
					Cvv: circlesdk.CardVerificationCvvNotRequested,
				},
				RiskEvaluation: &circlesdk.RiskEvaluation{
					Decision: circlesdk.RiskEvaluationDecisionApproved,
					Reason:   "3000",
				},
				Metadata: &circlesdk.Metadata{
					Email:       "satoshi@circle.com",
					PhoneNumber: "+14155555555",
				},
				CreateDate: "2020-04-10T02:13:30.000Z",
				UpdateDate: "2020-04-10T02:13:30.000Z",
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.CreateCard(circlesdk.CreateCardRequest{})
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}
