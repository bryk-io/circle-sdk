package payments

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	circlesdk "github.com/bryk-io/circle-sdk"
	ac "github.com/stretchr/testify/assert"
)

type roundTripper struct {
	response *http.Response
	err      error
}

func (r *roundTripper) RoundTrip(_ *http.Request) (*http.Response, error) {
	return r.response, r.err
}

func TestAPI_CreateCard(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Card
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Post "v1/cards": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 400 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "400",
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 400,
  "message": "bad request"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("400: bad request"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
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
		data, err := api.CreateCard(circlesdk.CreateCardRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_GetCard(t *testing.T) {
	if os.Getenv("CIRCLE_API_KEY") == "" {
		t.Skip("no API key available")
	}

	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Card
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/cards/": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
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
		data, err := api.GetCard("", tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error(), tt.name)
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_ListCards(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     []*circlesdk.Card
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/cards": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": [
    {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "status": "pending",
      "billingDetails": {
        "country": "US",
        "district": "MA"
      },
      "expMonth": 1,
      "expYear": 2020,
      "network": "VISA",
      "bin": "401230",
      "issuerCountry": "US",
      "fingerprint": "eb170539-9e1c-4e92-bf4f-1d09534fdca2",
      "verification": {
        "avs": "not_requested",
        "cvv": "not_requested"
      },
      "riskEvaluation": {
        "decision": "approved",
        "reason": "3000"
      },
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    }
  ]
}`)),
				},
			},
			expected: []*circlesdk.Card{
				{
					ID:     "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Status: circlesdk.CardStatusPending,
					BillingDetails: &circlesdk.BillingDetails{
						Country:  "US",
						District: "MA",
					},
					ExpMonth:      1,
					ExpYear:       2020,
					Network:       circlesdk.CardNetworkVISA,
					Bin:           "401230",
					IssuerCountry: "US",
					Fingerprint:   "eb170539-9e1c-4e92-bf4f-1d09534fdca2",
					Verification: &circlesdk.CardVerification{
						Avs: circlesdk.CardVerificationAvsNotRequested,
						Cvv: circlesdk.CardVerificationCvvNotRequested,
					},
					RiskEvaluation: &circlesdk.RiskEvaluation{
						Decision: circlesdk.RiskEvaluationDecisionApproved,
						Reason:   "3000",
					},
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.ListCards(tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_UpdateCard(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Card
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Put "v1/cards/": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 400 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "400",
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 400,
  "message": "bad request"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("400: bad request"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
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
		data, err := api.UpdateCard("", circlesdk.UpdateCardRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_CreateBankAccount(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.BankAccount
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Post "v1/banks/ach": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 400 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "400",
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 400,
  "message": "bad request"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("400: bad request"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "201",
					StatusCode: http.StatusCreated,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "status": "pending",
    "description": "WELLS FARGO BANK, NA ****0010",
    "trackingRef": "CIR13FB13A",
    "fingerprint": "eb170539-9e1c-4e92-bf4f-1d09534fdca2",
    "billingDetails": {
      "name": "Satoshi Nakamoto",
      "city": "Boston",
      "country": "US",
      "line1": "100 Money Street",
      "line2": "Suite 1",
      "district": "MA",
      "postalCode": "01234"
    },
    "bankAddress": {
      "bankName": "SAN FRANCISCO",
      "city": "SAN FRANCISCO",
      "country": "US",
      "line1": "100 Money Street",
      "line2": "Suite 1",
      "district": "CA"
    },
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.BankAccount{
				ID:          "b8627ae8-732b-4d25-b947-1df8f4007a29",
				Status:      circlesdk.BankAccountStatusPending,
				Description: "WELLS FARGO BANK, NA ****0010",
				TrackingRef: "CIR13FB13A",
				BillingDetails: &circlesdk.BillingDetails{
					Name:       "Satoshi Nakamoto",
					City:       "Boston",
					Country:    "US",
					Line1:      "100 Money Street",
					Line2:      "Suite 1",
					District:   "MA",
					PostalCode: "01234",
				},
				BankAddress: &circlesdk.BankAddress{
					BankName: "SAN FRANCISCO",
					City:     "SAN FRANCISCO",
					Country:  "US",
					Line1:    "100 Money Street",
					Line2:    "Suite 1",
					District: "CA",
				},
				Fingerprint: "eb170539-9e1c-4e92-bf4f-1d09534fdca2",
				CreateDate:  "2020-04-10T02:13:30.000Z",
				UpdateDate:  "2020-04-10T02:13:30.000Z",
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.CreateBankAccount(circlesdk.CreateBankAccountRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_GetBankAccount(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.BankAccount
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/banks/ach/": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "status": "pending",
    "description": "WELLS FARGO BANK, NA ****0010",
    "trackingRef": "CIR13FB13A",
    "fingerprint": "eb170539-9e1c-4e92-bf4f-1d09534fdca2",
    "billingDetails": {
      "name": "Satoshi Nakamoto",
      "city": "Boston",
      "country": "US",
      "line1": "100 Money Street",
      "line2": "Suite 1",
      "district": "MA",
      "postalCode": "01234"
    },
    "bankAddress": {
      "bankName": "SAN FRANCISCO",
      "city": "SAN FRANCISCO",
      "country": "US",
      "line1": "100 Money Street",
      "line2": "Suite 1",
      "district": "CA"
    },
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.BankAccount{
				ID:          "b8627ae8-732b-4d25-b947-1df8f4007a29",
				Status:      circlesdk.BankAccountStatusPending,
				Description: "WELLS FARGO BANK, NA ****0010",
				TrackingRef: "CIR13FB13A",
				BillingDetails: &circlesdk.BillingDetails{
					Name:       "Satoshi Nakamoto",
					City:       "Boston",
					Country:    "US",
					Line1:      "100 Money Street",
					Line2:      "Suite 1",
					District:   "MA",
					PostalCode: "01234",
				},
				BankAddress: &circlesdk.BankAddress{
					BankName: "SAN FRANCISCO",
					City:     "SAN FRANCISCO",
					Country:  "US",
					Line1:    "100 Money Street",
					Line2:    "Suite 1",
					District: "CA",
				},
				Fingerprint: "eb170539-9e1c-4e92-bf4f-1d09534fdca2",
				CreateDate:  "2020-04-10T02:13:30.000Z",
				UpdateDate:  "2020-04-10T02:13:30.000Z",
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.GetBankAccount("", tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_CreatePayment(t *testing.T) {
	captured := false
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Payment
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Post "v1/payments": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 400 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "400",
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 400,
  "message": "bad request"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("400: bad request"),
		},
		{
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "201",
					StatusCode: http.StatusCreated,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "type": "payment",
    "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "merchantWalletId": "212000",
    "amount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "source": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "card"
    },
    "description": "Payment",
    "status": "pending",
    "captured": false,
    "captureAmount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "captureDate": "2020-04-10T02:13:30.000Z",
    "requiredAction": {
      "type": "three_d_secure_required",
      "redirectUrl": "https://example.org"
    },
    "verification": {
      "avs": "not_requested",
      "cvv": "not_requested",
      "threeDSecure": "pass",
      "eci": "00"
    },
    "cancel": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "cancel",
      "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
      "merchantWalletId": "212000",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "source": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "card"
      },
      "description": "Payment",
      "status": "pending",
      "originalPayment": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "payment",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "description": "Payment",
        "status": "pending",
        "requiredAction": {
          "type": "three_d_secure_required",
          "redirectUrl": "https://example.org"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "createDate": "2020-04-10T02:13:30.000Z"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    },
    "refunds": [
      {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "refund",
        "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
        "merchantWalletId": "212000",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "source": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "card"
        },
        "description": "Payment",
        "status": "pending",
        "originalPayment": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "payment",
          "amount": {
            "amount": "3.14",
            "currency": "USD"
          },
          "description": "Payment",
          "status": "pending",
          "requiredAction": {
            "type": "three_d_secure_required",
            "redirectUrl": "https://example.org"
          },
          "fees": {
            "amount": "3.14",
            "currency": "USD"
          },
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "cancel": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "cancel",
          "description": "Payment",
          "status": "pending",
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
        "createDate": "2020-04-10T02:13:30.000Z",
        "updateDate": "2020-04-10T02:13:30.000Z"
      }
    ],
    "fees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "trackingRef": "24910599141085313498894",
    "errorCode": "payment_failed",
    "metadata": {
      "email": "satoshi@circle.com",
      "phoneNumber": "+14155555555"
    },
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.Payment{
				ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
				Type:             circlesdk.PaymentTypePayment,
				MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
				MerchantWalletID: "212000",
				Amount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				Source: &circlesdk.Source{
					ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type: circlesdk.SourceTypeCard,
				},
				Description: "Payment",
				Status:      circlesdk.PaymentStatusPending,
				Captured:    &captured,
				CaptureAmount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				CaptureDate: "2020-04-10T02:13:30.000Z",
				RequiredAction: &circlesdk.RequiredAction{
					Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
					RedirectURL: "https://example.org",
				},
				Verification: &circlesdk.PaymentVerification{
					Avs:          circlesdk.PaymentVerificationAvsNotRequested,
					Cvv:          circlesdk.PaymentVerificationCvvNotRequested,
					ThreeDSecure: circlesdk.PaymentVerificationThreeDSecurePass,
					Eci:          circlesdk.PaymentVerificationEci00,
				},
				Fees: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				TrackingRef: "24910599141085313498894",
				ErrorCode:   circlesdk.PaymentErrorCodePaymentFailed,
				Metadata: &circlesdk.Metadata{
					Email:       "satoshi@circle.com",
					PhoneNumber: "+14155555555",
				},
				RiskEvaluation: &circlesdk.RiskEvaluation{
					Decision: circlesdk.RiskEvaluationDecisionApproved,
					Reason:   "3000",
				},
				Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
				CreateDate: "2020-04-10T02:13:30.000Z",
				UpdateDate: "2020-04-10T02:13:30.000Z",
				Cancel: &circlesdk.Payment{
					ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type:             circlesdk.PaymentTypeCancel,
					MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
					MerchantWalletID: "212000",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Source: &circlesdk.Source{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.SourceTypeCard,
					},
					Description: "Payment",
					Status:      circlesdk.PaymentStatusPending,
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
					OriginalPayment: &circlesdk.Payment{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.PaymentTypePayment,
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						RequiredAction: &circlesdk.RequiredAction{
							Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
							RedirectURL: "https://example.org",
						},
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						CreateDate: "2020-04-10T02:13:30.000Z",
					},
				},
				Refunds: []circlesdk.Payment{
					{
						ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type:             circlesdk.PaymentTypeRefund,
						MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
						MerchantWalletID: "212000",
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Source: &circlesdk.Source{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.SourceTypeCard,
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
						CreateDate: "2020-04-10T02:13:30.000Z",
						UpdateDate: "2020-04-10T02:13:30.000Z",
						Cancel: &circlesdk.Payment{
							ID:          "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type:        circlesdk.PaymentTypeCancel,
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							CreateDate:  "2020-04-10T02:13:30.000Z",
						},
						OriginalPayment: &circlesdk.Payment{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.PaymentTypePayment,
							Amount: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							RequiredAction: &circlesdk.RequiredAction{
								Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
								RedirectURL: "https://example.org",
							},
							Fees: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							CreateDate: "2020-04-10T02:13:30.000Z",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.CreatePayment(circlesdk.CreatePaymentRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_GetPayment(t *testing.T) {
	captured := false
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Payment
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/payments/": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "type": "payment",
    "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "merchantWalletId": "212000",
    "amount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "source": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "card"
    },
    "description": "Payment",
    "status": "pending",
    "captured": false,
    "captureAmount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "captureDate": "2020-04-10T02:13:30.000Z",
    "requiredAction": {
      "type": "three_d_secure_required",
      "redirectUrl": "https://example.org"
    },
    "verification": {
      "avs": "not_requested",
      "cvv": "not_requested",
      "threeDSecure": "pass",
      "eci": "00"
    },
    "cancel": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "cancel",
      "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
      "merchantWalletId": "212000",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "source": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "card"
      },
      "description": "Payment",
      "status": "pending",
      "originalPayment": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "payment",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "description": "Payment",
        "status": "pending",
        "requiredAction": {
          "type": "three_d_secure_required",
          "redirectUrl": "https://example.org"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "createDate": "2020-04-10T02:13:30.000Z"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    },
    "refunds": [
      {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "refund",
        "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
        "merchantWalletId": "212000",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "source": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "card"
        },
        "description": "Payment",
        "status": "pending",
        "originalPayment": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "payment",
          "amount": {
            "amount": "3.14",
            "currency": "USD"
          },
          "description": "Payment",
          "status": "pending",
          "requiredAction": {
            "type": "three_d_secure_required",
            "redirectUrl": "https://example.org"
          },
          "fees": {
            "amount": "3.14",
            "currency": "USD"
          },
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "cancel": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "cancel",
          "description": "Payment",
          "status": "pending",
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
        "createDate": "2020-04-10T02:13:30.000Z",
        "updateDate": "2020-04-10T02:13:30.000Z"
      }
    ],
    "fees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "trackingRef": "24910599141085313498894",
    "errorCode": "payment_failed",
    "metadata": {
      "email": "satoshi@circle.com",
      "phoneNumber": "+14155555555"
    },
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.Payment{
				ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
				Type:             circlesdk.PaymentTypePayment,
				MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
				MerchantWalletID: "212000",
				Amount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				Source: &circlesdk.Source{
					ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type: circlesdk.SourceTypeCard,
				},
				Description: "Payment",
				Status:      circlesdk.PaymentStatusPending,
				Captured:    &captured,
				CaptureAmount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				CaptureDate: "2020-04-10T02:13:30.000Z",
				RequiredAction: &circlesdk.RequiredAction{
					Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
					RedirectURL: "https://example.org",
				},
				Verification: &circlesdk.PaymentVerification{
					Avs:          circlesdk.PaymentVerificationAvsNotRequested,
					Cvv:          circlesdk.PaymentVerificationCvvNotRequested,
					ThreeDSecure: circlesdk.PaymentVerificationThreeDSecurePass,
					Eci:          circlesdk.PaymentVerificationEci00,
				},
				Fees: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				TrackingRef: "24910599141085313498894",
				ErrorCode:   circlesdk.PaymentErrorCodePaymentFailed,
				Metadata: &circlesdk.Metadata{
					Email:       "satoshi@circle.com",
					PhoneNumber: "+14155555555",
				},
				RiskEvaluation: &circlesdk.RiskEvaluation{
					Decision: circlesdk.RiskEvaluationDecisionApproved,
					Reason:   "3000",
				},
				Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
				CreateDate: "2020-04-10T02:13:30.000Z",
				UpdateDate: "2020-04-10T02:13:30.000Z",
				Cancel: &circlesdk.Payment{
					ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type:             circlesdk.PaymentTypeCancel,
					MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
					MerchantWalletID: "212000",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Source: &circlesdk.Source{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.SourceTypeCard,
					},
					Description: "Payment",
					Status:      circlesdk.PaymentStatusPending,
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
					OriginalPayment: &circlesdk.Payment{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.PaymentTypePayment,
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						RequiredAction: &circlesdk.RequiredAction{
							Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
							RedirectURL: "https://example.org",
						},
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						CreateDate: "2020-04-10T02:13:30.000Z",
					},
				},
				Refunds: []circlesdk.Payment{
					{
						ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type:             circlesdk.PaymentTypeRefund,
						MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
						MerchantWalletID: "212000",
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Source: &circlesdk.Source{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.SourceTypeCard,
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
						CreateDate: "2020-04-10T02:13:30.000Z",
						UpdateDate: "2020-04-10T02:13:30.000Z",
						Cancel: &circlesdk.Payment{
							ID:          "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type:        circlesdk.PaymentTypeCancel,
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							CreateDate:  "2020-04-10T02:13:30.000Z",
						},
						OriginalPayment: &circlesdk.Payment{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.PaymentTypePayment,
							Amount: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							RequiredAction: &circlesdk.RequiredAction{
								Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
								RedirectURL: "https://example.org",
							},
							Fees: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							CreateDate: "2020-04-10T02:13:30.000Z",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.GetPayment("", tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_ListPayments(t *testing.T) {
	captured := false
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     []*circlesdk.Payment
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/payments": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": [{
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "type": "payment",
    "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "merchantWalletId": "212000",
    "amount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "source": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "card"
    },
    "description": "Payment",
    "status": "pending",
    "captured": false,
    "captureAmount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "captureDate": "2020-04-10T02:13:30.000Z",
    "requiredAction": {
      "type": "three_d_secure_required",
      "redirectUrl": "https://example.org"
    },
    "verification": {
      "avs": "not_requested",
      "cvv": "not_requested",
      "threeDSecure": "pass",
      "eci": "00"
    },
    "cancel": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "cancel",
      "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
      "merchantWalletId": "212000",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "source": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "card"
      },
      "description": "Payment",
      "status": "pending",
      "originalPayment": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "payment",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "description": "Payment",
        "status": "pending",
        "requiredAction": {
          "type": "three_d_secure_required",
          "redirectUrl": "https://example.org"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "createDate": "2020-04-10T02:13:30.000Z"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    },
    "refunds": [
      {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "refund",
        "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
        "merchantWalletId": "212000",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "source": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "card"
        },
        "description": "Payment",
        "status": "pending",
        "originalPayment": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "payment",
          "amount": {
            "amount": "3.14",
            "currency": "USD"
          },
          "description": "Payment",
          "status": "pending",
          "requiredAction": {
            "type": "three_d_secure_required",
            "redirectUrl": "https://example.org"
          },
          "fees": {
            "amount": "3.14",
            "currency": "USD"
          },
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "cancel": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "cancel",
          "description": "Payment",
          "status": "pending",
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
        "createDate": "2020-04-10T02:13:30.000Z",
        "updateDate": "2020-04-10T02:13:30.000Z"
      }
    ],
    "fees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "trackingRef": "24910599141085313498894",
    "errorCode": "payment_failed",
    "metadata": {
      "email": "satoshi@circle.com",
      "phoneNumber": "+14155555555"
    },
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }]
}`)),
				},
			},
			expected: []*circlesdk.Payment{
				{
					ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type:             circlesdk.PaymentTypePayment,
					MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
					MerchantWalletID: "212000",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Source: &circlesdk.Source{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.SourceTypeCard,
					},
					Description: "Payment",
					Status:      circlesdk.PaymentStatusPending,
					Captured:    &captured,
					CaptureAmount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					CaptureDate: "2020-04-10T02:13:30.000Z",
					RequiredAction: &circlesdk.RequiredAction{
						Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
						RedirectURL: "https://example.org",
					},
					Verification: &circlesdk.PaymentVerification{
						Avs:          circlesdk.PaymentVerificationAvsNotRequested,
						Cvv:          circlesdk.PaymentVerificationCvvNotRequested,
						ThreeDSecure: circlesdk.PaymentVerificationThreeDSecurePass,
						Eci:          circlesdk.PaymentVerificationEci00,
					},
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					TrackingRef: "24910599141085313498894",
					ErrorCode:   circlesdk.PaymentErrorCodePaymentFailed,
					Metadata: &circlesdk.Metadata{
						Email:       "satoshi@circle.com",
						PhoneNumber: "+14155555555",
					},
					RiskEvaluation: &circlesdk.RiskEvaluation{
						Decision: circlesdk.RiskEvaluationDecisionApproved,
						Reason:   "3000",
					},
					Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
					Cancel: &circlesdk.Payment{
						ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type:             circlesdk.PaymentTypeCancel,
						MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
						MerchantWalletID: "212000",
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Source: &circlesdk.Source{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.SourceTypeCard,
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
						CreateDate: "2020-04-10T02:13:30.000Z",
						UpdateDate: "2020-04-10T02:13:30.000Z",
						OriginalPayment: &circlesdk.Payment{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.PaymentTypePayment,
							Amount: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							RequiredAction: &circlesdk.RequiredAction{
								Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
								RedirectURL: "https://example.org",
							},
							Fees: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							CreateDate: "2020-04-10T02:13:30.000Z",
						},
					},
					Refunds: []circlesdk.Payment{
						{
							ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type:             circlesdk.PaymentTypeRefund,
							MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
							MerchantWalletID: "212000",
							Amount: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Source: &circlesdk.Source{
								ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
								Type: circlesdk.SourceTypeCard,
							},
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							Fees: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
							CreateDate: "2020-04-10T02:13:30.000Z",
							UpdateDate: "2020-04-10T02:13:30.000Z",
							Cancel: &circlesdk.Payment{
								ID:          "b8627ae8-732b-4d25-b947-1df8f4007a29",
								Type:        circlesdk.PaymentTypeCancel,
								Description: "Payment",
								Status:      circlesdk.PaymentStatusPending,
								CreateDate:  "2020-04-10T02:13:30.000Z",
							},
							OriginalPayment: &circlesdk.Payment{
								ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
								Type: circlesdk.PaymentTypePayment,
								Amount: &circlesdk.Amount{
									Amount:   "3.14",
									Currency: "USD",
								},
								Description: "Payment",
								Status:      circlesdk.PaymentStatusPending,
								RequiredAction: &circlesdk.RequiredAction{
									Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
									RedirectURL: "https://example.org",
								},
								Fees: &circlesdk.Amount{
									Amount:   "3.14",
									Currency: "USD",
								},
								CreateDate: "2020-04-10T02:13:30.000Z",
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.ListPayments(circlesdk.ListPaymentsRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_CapturePayment(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Post "v1/payments/test/capture": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 400 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "400",
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 400,
  "message": "bad request"
}`)),
				},
			},
			expectedErr: errors.New("400: bad request"),
		},
		{
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "202",
					StatusCode: http.StatusAccepted,
					Body:       io.NopCloser(strings.NewReader(`{}`)),
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		err := api.CapturePayment("test", circlesdk.CapturePaymentRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
	}
}

func TestAPI_CancelPayment(t *testing.T) {
	captured := false
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Payment
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Post "v1/payments/test/cancel": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 400 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "400",
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 400,
  "message": "bad request"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("400: bad request"),
		},
		{
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "201",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "type": "payment",
    "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "merchantWalletId": "212000",
    "amount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "source": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "card"
    },
    "description": "Payment",
    "status": "pending",
    "captured": false,
    "captureAmount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "captureDate": "2020-04-10T02:13:30.000Z",
    "requiredAction": {
      "type": "three_d_secure_required",
      "redirectUrl": "https://example.org"
    },
    "verification": {
      "avs": "not_requested",
      "cvv": "not_requested",
      "threeDSecure": "pass",
      "eci": "00"
    },
    "cancel": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "cancel",
      "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
      "merchantWalletId": "212000",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "source": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "card"
      },
      "description": "Payment",
      "status": "pending",
      "originalPayment": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "payment",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "description": "Payment",
        "status": "pending",
        "requiredAction": {
          "type": "three_d_secure_required",
          "redirectUrl": "https://example.org"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "createDate": "2020-04-10T02:13:30.000Z"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    },
    "refunds": [
      {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "refund",
        "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
        "merchantWalletId": "212000",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "source": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "card"
        },
        "description": "Payment",
        "status": "pending",
        "originalPayment": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "payment",
          "amount": {
            "amount": "3.14",
            "currency": "USD"
          },
          "description": "Payment",
          "status": "pending",
          "requiredAction": {
            "type": "three_d_secure_required",
            "redirectUrl": "https://example.org"
          },
          "fees": {
            "amount": "3.14",
            "currency": "USD"
          },
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "cancel": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "cancel",
          "description": "Payment",
          "status": "pending",
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
        "createDate": "2020-04-10T02:13:30.000Z",
        "updateDate": "2020-04-10T02:13:30.000Z"
      }
    ],
    "fees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "trackingRef": "24910599141085313498894",
    "errorCode": "payment_failed",
    "metadata": {
      "email": "satoshi@circle.com",
      "phoneNumber": "+14155555555"
    },
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.Payment{
				ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
				Type:             circlesdk.PaymentTypePayment,
				MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
				MerchantWalletID: "212000",
				Amount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				Source: &circlesdk.Source{
					ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type: circlesdk.SourceTypeCard,
				},
				Description: "Payment",
				Status:      circlesdk.PaymentStatusPending,
				Captured:    &captured,
				CaptureAmount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				CaptureDate: "2020-04-10T02:13:30.000Z",
				RequiredAction: &circlesdk.RequiredAction{
					Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
					RedirectURL: "https://example.org",
				},
				Verification: &circlesdk.PaymentVerification{
					Avs:          circlesdk.PaymentVerificationAvsNotRequested,
					Cvv:          circlesdk.PaymentVerificationCvvNotRequested,
					ThreeDSecure: circlesdk.PaymentVerificationThreeDSecurePass,
					Eci:          circlesdk.PaymentVerificationEci00,
				},
				Fees: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				TrackingRef: "24910599141085313498894",
				ErrorCode:   circlesdk.PaymentErrorCodePaymentFailed,
				Metadata: &circlesdk.Metadata{
					Email:       "satoshi@circle.com",
					PhoneNumber: "+14155555555",
				},
				RiskEvaluation: &circlesdk.RiskEvaluation{
					Decision: circlesdk.RiskEvaluationDecisionApproved,
					Reason:   "3000",
				},
				Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
				CreateDate: "2020-04-10T02:13:30.000Z",
				UpdateDate: "2020-04-10T02:13:30.000Z",
				Cancel: &circlesdk.Payment{
					ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type:             circlesdk.PaymentTypeCancel,
					MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
					MerchantWalletID: "212000",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Source: &circlesdk.Source{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.SourceTypeCard,
					},
					Description: "Payment",
					Status:      circlesdk.PaymentStatusPending,
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
					OriginalPayment: &circlesdk.Payment{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.PaymentTypePayment,
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						RequiredAction: &circlesdk.RequiredAction{
							Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
							RedirectURL: "https://example.org",
						},
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						CreateDate: "2020-04-10T02:13:30.000Z",
					},
				},
				Refunds: []circlesdk.Payment{
					{
						ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type:             circlesdk.PaymentTypeRefund,
						MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
						MerchantWalletID: "212000",
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Source: &circlesdk.Source{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.SourceTypeCard,
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
						CreateDate: "2020-04-10T02:13:30.000Z",
						UpdateDate: "2020-04-10T02:13:30.000Z",
						Cancel: &circlesdk.Payment{
							ID:          "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type:        circlesdk.PaymentTypeCancel,
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							CreateDate:  "2020-04-10T02:13:30.000Z",
						},
						OriginalPayment: &circlesdk.Payment{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.PaymentTypePayment,
							Amount: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							RequiredAction: &circlesdk.RequiredAction{
								Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
								RedirectURL: "https://example.org",
							},
							Fees: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							CreateDate: "2020-04-10T02:13:30.000Z",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.CancelPayment("test", circlesdk.CapturePaymentRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_RefundPayment(t *testing.T) {
	captured := false
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Payment
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Post "v1/payments/test/refund": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 400 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "400",
					StatusCode: http.StatusBadRequest,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 400,
  "message": "bad request"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("400: bad request"),
		},
		{
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "201",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "type": "payment",
    "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "merchantWalletId": "212000",
    "amount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "source": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "card"
    },
    "description": "Payment",
    "status": "pending",
    "captured": false,
    "captureAmount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "captureDate": "2020-04-10T02:13:30.000Z",
    "requiredAction": {
      "type": "three_d_secure_required",
      "redirectUrl": "https://example.org"
    },
    "verification": {
      "avs": "not_requested",
      "cvv": "not_requested",
      "threeDSecure": "pass",
      "eci": "00"
    },
    "cancel": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "type": "cancel",
      "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
      "merchantWalletId": "212000",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "source": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "card"
      },
      "description": "Payment",
      "status": "pending",
      "originalPayment": {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "payment",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "description": "Payment",
        "status": "pending",
        "requiredAction": {
          "type": "three_d_secure_required",
          "redirectUrl": "https://example.org"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "createDate": "2020-04-10T02:13:30.000Z"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    },
    "refunds": [
      {
        "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
        "type": "refund",
        "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
        "merchantWalletId": "212000",
        "amount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "source": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "card"
        },
        "description": "Payment",
        "status": "pending",
        "originalPayment": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "payment",
          "amount": {
            "amount": "3.14",
            "currency": "USD"
          },
          "description": "Payment",
          "status": "pending",
          "requiredAction": {
            "type": "three_d_secure_required",
            "redirectUrl": "https://example.org"
          },
          "fees": {
            "amount": "3.14",
            "currency": "USD"
          },
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "cancel": {
          "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
          "type": "cancel",
          "description": "Payment",
          "status": "pending",
          "createDate": "2020-04-10T02:13:30.000Z"
        },
        "fees": {
          "amount": "3.14",
          "currency": "USD"
        },
        "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
        "createDate": "2020-04-10T02:13:30.000Z",
        "updateDate": "2020-04-10T02:13:30.000Z"
      }
    ],
    "fees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "trackingRef": "24910599141085313498894",
    "errorCode": "payment_failed",
    "metadata": {
      "email": "satoshi@circle.com",
      "phoneNumber": "+14155555555"
    },
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "channel": "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.Payment{
				ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
				Type:             circlesdk.PaymentTypePayment,
				MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
				MerchantWalletID: "212000",
				Amount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				Source: &circlesdk.Source{
					ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type: circlesdk.SourceTypeCard,
				},
				Description: "Payment",
				Status:      circlesdk.PaymentStatusPending,
				Captured:    &captured,
				CaptureAmount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				CaptureDate: "2020-04-10T02:13:30.000Z",
				RequiredAction: &circlesdk.RequiredAction{
					Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
					RedirectURL: "https://example.org",
				},
				Verification: &circlesdk.PaymentVerification{
					Avs:          circlesdk.PaymentVerificationAvsNotRequested,
					Cvv:          circlesdk.PaymentVerificationCvvNotRequested,
					ThreeDSecure: circlesdk.PaymentVerificationThreeDSecurePass,
					Eci:          circlesdk.PaymentVerificationEci00,
				},
				Fees: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				TrackingRef: "24910599141085313498894",
				ErrorCode:   circlesdk.PaymentErrorCodePaymentFailed,
				Metadata: &circlesdk.Metadata{
					Email:       "satoshi@circle.com",
					PhoneNumber: "+14155555555",
				},
				RiskEvaluation: &circlesdk.RiskEvaluation{
					Decision: circlesdk.RiskEvaluationDecisionApproved,
					Reason:   "3000",
				},
				Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
				CreateDate: "2020-04-10T02:13:30.000Z",
				UpdateDate: "2020-04-10T02:13:30.000Z",
				Cancel: &circlesdk.Payment{
					ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type:             circlesdk.PaymentTypeCancel,
					MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
					MerchantWalletID: "212000",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Source: &circlesdk.Source{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.SourceTypeCard,
					},
					Description: "Payment",
					Status:      circlesdk.PaymentStatusPending,
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
					OriginalPayment: &circlesdk.Payment{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.PaymentTypePayment,
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						RequiredAction: &circlesdk.RequiredAction{
							Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
							RedirectURL: "https://example.org",
						},
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						CreateDate: "2020-04-10T02:13:30.000Z",
					},
				},
				Refunds: []circlesdk.Payment{
					{
						ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type:             circlesdk.PaymentTypeRefund,
						MerchantID:       "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
						MerchantWalletID: "212000",
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Source: &circlesdk.Source{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.SourceTypeCard,
						},
						Description: "Payment",
						Status:      circlesdk.PaymentStatusPending,
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Channel:    "ba943ff1-ca16-49b2-ba55-1057e70ca5c7",
						CreateDate: "2020-04-10T02:13:30.000Z",
						UpdateDate: "2020-04-10T02:13:30.000Z",
						Cancel: &circlesdk.Payment{
							ID:          "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type:        circlesdk.PaymentTypeCancel,
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							CreateDate:  "2020-04-10T02:13:30.000Z",
						},
						OriginalPayment: &circlesdk.Payment{
							ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
							Type: circlesdk.PaymentTypePayment,
							Amount: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Description: "Payment",
							Status:      circlesdk.PaymentStatusPending,
							RequiredAction: &circlesdk.RequiredAction{
								Type:        circlesdk.RequiredActionTypeThreeDSecureRequired,
								RedirectURL: "https://example.org",
							},
							Fees: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							CreateDate: "2020-04-10T02:13:30.000Z",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.RefundPayment("test", circlesdk.RefundPaymentRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_GetSettlement(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Settlement
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/settlements/": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "merchantWalletId": "212000",
    "walletId": "12345",
    "totalDebits": {
      "amount": "3.14",
      "currency": "USD"
    },
    "totalCredits": {
      "amount": "3.14",
      "currency": "USD"
    },
    "paymentFees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "chargebackFees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.Settlement{
				ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
				MerchantWalletID: "212000",
				WalletID:         "12345",
				TotalDebits: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				TotalCredits: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				PaymentFees: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				ChargebackFees: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				CreateDate: "2020-04-10T02:13:30.000Z",
				UpdateDate: "2020-04-10T02:13:30.000Z",
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.GetSettlement("", tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_ListSettlements(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     []*circlesdk.Settlement
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/settlements": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": [
    {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "merchantWalletId": "212000",
      "walletId": "12345",
      "totalDebits": {
        "amount": "3.14",
        "currency": "USD"
      },
      "totalCredits": {
        "amount": "3.14",
        "currency": "USD"
      },
      "paymentFees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "chargebackFees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    }
  ]
}`)),
				},
			},
			expected: []*circlesdk.Settlement{
				{
					ID:               "b8627ae8-732b-4d25-b947-1df8f4007a29",
					MerchantWalletID: "212000",
					WalletID:         "12345",
					TotalDebits: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					TotalCredits: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					PaymentFees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					ChargebackFees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.ListSettlements(tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_GetChargeback(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.ChargeBack
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/chargebacks/": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "api 404 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "404",
					StatusCode: http.StatusNotFound,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 404,
  "message": "not found"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: not found"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": {
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "paymentId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "reasonCode": "10.4",
    "category": "Canceled Recurring Payment",
    "history": [
      {
        "type": "1st Chargeback",
        "chargebackAmount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "fee": {
          "amount": "3.14",
          "currency": "USD"
        },
        "description": "The cardholder claims an unauthorized transaction occurred.",
        "settlementId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
        "createDate": "2020-04-10T02:13:30.000Z"
      }
    ]
  }
}`)),
				},
			},
			expected: &circlesdk.ChargeBack{
				ID:         "b8627ae8-732b-4d25-b947-1df8f4007a29",
				PaymentID:  "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
				MerchantID: "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
				ReasonCode: "10.4",
				Category:   circlesdk.ChargeBackCategoryCanceledRecurringPayment,
				History: []circlesdk.ChargeBackHistory{
					{
						Type: circlesdk.ChargeBackHistoryType1stChargeBack,
						ChargeBackAmount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Fee: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Description:  "The cardholder claims an unauthorized transaction occurred.",
						SettlementID: "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
						CreateDate:   "2020-04-10T02:13:30.000Z",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.GetChargeback("", tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_ListChargebacks(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     []*circlesdk.ChargeBack
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/chargebacks": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": [{
    "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
    "paymentId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "merchantId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
    "reasonCode": "10.4",
    "category": "Canceled Recurring Payment",
    "history": [
      {
        "type": "1st Chargeback",
        "chargebackAmount": {
          "amount": "3.14",
          "currency": "USD"
        },
        "fee": {
          "amount": "3.14",
          "currency": "USD"
        },
        "description": "The cardholder claims an unauthorized transaction occurred.",
        "settlementId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
        "createDate": "2020-04-10T02:13:30.000Z"
      }
    ]
  }]
}`)),
				},
			},
			expected: []*circlesdk.ChargeBack{
				{
					ID:         "b8627ae8-732b-4d25-b947-1df8f4007a29",
					PaymentID:  "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
					MerchantID: "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
					ReasonCode: "10.4",
					Category:   circlesdk.ChargeBackCategoryCanceledRecurringPayment,
					History: []circlesdk.ChargeBackHistory{
						{
							Type: circlesdk.ChargeBackHistoryType1stChargeBack,
							ChargeBackAmount: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Fee: &circlesdk.Amount{
								Amount:   "3.14",
								Currency: "USD",
							},
							Description:  "The cardholder claims an unauthorized transaction occurred.",
							SettlementID: "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
							CreateDate:   "2020-04-10T02:13:30.000Z",
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.ListChargebacks("", tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_ListReversals(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     []*circlesdk.Reversal
		expectedErr  error
	}{
		{
			name: "with option error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return errors.New("error")
			},
			expectedErr: errors.New("error"),
		},
		{
			name: "internal error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				err: errors.New("some error"),
			},
			expectedErr: errors.New(`Get "v1/reversals": some error`),
		}, {
			name: "api 500 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "500",
					StatusCode: http.StatusInternalServerError,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 500,
  "message": "internal error"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: internal error"),
		}, {
			name: "api 401 error",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "401",
					StatusCode: http.StatusUnauthorized,
					Body: io.NopCloser(strings.NewReader(`{
  "code": 401,
  "message": "malformed authorization"
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: malformed authorization"),
		}, {
			name: "success",
			withOptions: func(options *circlesdk.RequestOptions) error {
				return nil
			},
			roundTripper: roundTripper{
				response: &http.Response{
					Status:     "200",
					StatusCode: http.StatusOK,
					Body: io.NopCloser(strings.NewReader(`
{
  "data": [
    {
      "id": "key1",
      "paymentId": "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "description": "Merchant Payment Reversal",
      "status": "pending",
      "reason": "duplicate",
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    }
  ]
}`)),
				},
			},
			expected: []*circlesdk.Reversal{
				{
					ID:        "key1",
					PaymentID: "fc988ed5-c129-4f70-a064-e5beb7eb8e32",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Description: "Merchant Payment Reversal",
					Status:      circlesdk.ReversalStatusPending,
					Reason:      circlesdk.ReversalReasonDuplicate,
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.ListReversals("", tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}
