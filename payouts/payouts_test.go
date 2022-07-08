package payouts

import (
	"errors"
	"io"
	"net/http"
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

func TestAPI_CreatePayout(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Payout
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
			expectedErr: errors.New(`Post "v1/payouts": some error`),
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
  "message": "Internal error."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: Internal error."),
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
  "message": "Malformed authorization."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: Malformed authorization."),
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
  "message": "Bad request."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("400: Bad request."),
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
    "sourceWalletId": "53535335",
    "destination": {
      "type": "wire",
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "name": "COMMERZBANK AG ****3000"
    },
    "amount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "fees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "status": "pending",
    "trackingRef": "CIR-6ESOQANEP3NAO",
    "errorCode": "insufficient_funds",
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "adjustments": {
      "fxCredit": {
        "amount": "3.14",
        "currency": "USD"
      },
      "fxDebit": {
        "amount": "3.14",
        "currency": "USD"
      }
    },
    "return": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "payoutId": "abdb500d-4a59-457c-801f-2d418c8703ac",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "reason": "payout_returned",
      "status": "pending",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    },
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.Payout{
				ID:             "b8627ae8-732b-4d25-b947-1df8f4007a29",
				SourceWalletID: "53535335",
				Destination: &circlesdk.PayoutDestination{
					ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type: circlesdk.PayoutDestinationTypeWire,
					Name: "COMMERZBANK AG ****3000",
				},
				Amount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				Fees: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				Status:      circlesdk.PayoutStatusPending,
				TrackingRef: "CIR-6ESOQANEP3NAO",
				ErrorCode:   circlesdk.PayoutErrorCodeInsufficientFunds,
				RiskEvaluation: &circlesdk.RiskEvaluation{
					Decision: circlesdk.RiskEvaluationDecisionApproved,
					Reason:   "3000",
				},
				Adjustments: &circlesdk.PayoutAdjustment{
					FxCredit: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					FxDebit: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
				},
				Return: &circlesdk.PayoutReturn{
					ID:       "b8627ae8-732b-4d25-b947-1df8f4007a29",
					PayoutID: "abdb500d-4a59-457c-801f-2d418c8703ac",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Reason:     "payout_returned",
					Status:     circlesdk.PayoutReturnStatusPending,
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
				},
				CreateDate: "2020-04-10T02:13:30.000Z",
				UpdateDate: "2020-04-10T02:13:30.000Z",
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.CreatePayout(circlesdk.CreatePayoutRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_GetPayout(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     *circlesdk.Payout
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
			expectedErr: errors.New(`Get "v1/payouts/": some error`),
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
 "message": "Internal error."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: Internal error."),
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
 "message": "Malformed authorization."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: Malformed authorization."),
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
 "message": "Not found."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("404: Not found."),
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
    "sourceWalletId": "53535335",
    "destination": {
      "type": "wire",
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "name": "COMMERZBANK AG ****3000"
    },
    "amount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "fees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "status": "pending",
    "trackingRef": "CIR-6ESOQANEP3NAO",
    "errorCode": "insufficient_funds",
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "adjustments": {
      "fxCredit": {
        "amount": "3.14",
        "currency": "USD"
      },
      "fxDebit": {
        "amount": "3.14",
        "currency": "USD"
      }
    },
    "return": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "payoutId": "abdb500d-4a59-457c-801f-2d418c8703ac",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "reason": "payout_returned",
      "status": "pending",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    },
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }
}`)),
				},
			},
			expected: &circlesdk.Payout{
				ID:             "b8627ae8-732b-4d25-b947-1df8f4007a29",
				SourceWalletID: "53535335",
				Destination: &circlesdk.PayoutDestination{
					ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
					Type: circlesdk.PayoutDestinationTypeWire,
					Name: "COMMERZBANK AG ****3000",
				},
				Amount: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				Fees: &circlesdk.Amount{
					Amount:   "3.14",
					Currency: "USD",
				},
				Status:      circlesdk.PayoutStatusPending,
				TrackingRef: "CIR-6ESOQANEP3NAO",
				ErrorCode:   circlesdk.PayoutErrorCodeInsufficientFunds,
				RiskEvaluation: &circlesdk.RiskEvaluation{
					Decision: circlesdk.RiskEvaluationDecisionApproved,
					Reason:   "3000",
				},
				Adjustments: &circlesdk.PayoutAdjustment{
					FxCredit: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					FxDebit: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
				},
				Return: &circlesdk.PayoutReturn{
					ID:       "b8627ae8-732b-4d25-b947-1df8f4007a29",
					PayoutID: "abdb500d-4a59-457c-801f-2d418c8703ac",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Reason:     "payout_returned",
					Status:     circlesdk.PayoutReturnStatusPending,
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
				},
				CreateDate: "2020-04-10T02:13:30.000Z",
				UpdateDate: "2020-04-10T02:13:30.000Z",
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.GetPayout("", tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_ListPayouts(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     []*circlesdk.Payout
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
			expectedErr: errors.New(`Get "v1/payouts": some error`),
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
 "message": "Internal error."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: Internal error."),
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
 "message": "Malformed authorization."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: Malformed authorization."),
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
    "sourceWalletId": "53535335",
    "destination": {
      "type": "wire",
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "name": "COMMERZBANK AG ****3000"
    },
    "amount": {
      "amount": "3.14",
      "currency": "USD"
    },
    "fees": {
      "amount": "3.14",
      "currency": "USD"
    },
    "status": "pending",
    "trackingRef": "CIR-6ESOQANEP3NAO",
    "errorCode": "insufficient_funds",
    "riskEvaluation": {
      "decision": "approved",
      "reason": "3000"
    },
    "adjustments": {
      "fxCredit": {
        "amount": "3.14",
        "currency": "USD"
      },
      "fxDebit": {
        "amount": "3.14",
        "currency": "USD"
      }
    },
    "return": {
      "id": "b8627ae8-732b-4d25-b947-1df8f4007a29",
      "payoutId": "abdb500d-4a59-457c-801f-2d418c8703ac",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "reason": "payout_returned",
      "status": "pending",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    },
    "createDate": "2020-04-10T02:13:30.000Z",
    "updateDate": "2020-04-10T02:13:30.000Z"
  }]
}`)),
				},
			},
			expected: []*circlesdk.Payout{
				{
					ID:             "b8627ae8-732b-4d25-b947-1df8f4007a29",
					SourceWalletID: "53535335",
					Destination: &circlesdk.PayoutDestination{
						ID:   "b8627ae8-732b-4d25-b947-1df8f4007a29",
						Type: circlesdk.PayoutDestinationTypeWire,
						Name: "COMMERZBANK AG ****3000",
					},
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Status:      circlesdk.PayoutStatusPending,
					TrackingRef: "CIR-6ESOQANEP3NAO",
					ErrorCode:   circlesdk.PayoutErrorCodeInsufficientFunds,
					RiskEvaluation: &circlesdk.RiskEvaluation{
						Decision: circlesdk.RiskEvaluationDecisionApproved,
						Reason:   "3000",
					},
					Adjustments: &circlesdk.PayoutAdjustment{
						FxCredit: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						FxDebit: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
					},
					Return: &circlesdk.PayoutReturn{
						ID:       "b8627ae8-732b-4d25-b947-1df8f4007a29",
						PayoutID: "abdb500d-4a59-457c-801f-2d418c8703ac",
						Amount: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Fees: &circlesdk.Amount{
							Amount:   "3.14",
							Currency: "USD",
						},
						Reason:     "payout_returned",
						Status:     circlesdk.PayoutReturnStatusPending,
						CreateDate: "2020-04-10T02:13:30.000Z",
						UpdateDate: "2020-04-10T02:13:30.000Z",
					},
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.ListPayouts(circlesdk.ListPayoutsRequest{}, tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}

func TestAPI_ListReturns(t *testing.T) {
	assert := ac.New(t)
	tests := []struct {
		name         string
		withOptions  circlesdk.CallOption
		roundTripper roundTripper
		expected     []*circlesdk.PayoutReturn
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
			expectedErr: errors.New(`Get "v1/returns": some error`),
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
 "message": "Internal error."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("500: Internal error."),
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
 "message": "Malformed authorization."
}`)),
				},
			},
			expected:    nil,
			expectedErr: errors.New("401: Malformed authorization."),
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
      "payoutId": "abdb500d-4a59-457c-801f-2d418c8703ac",
      "amount": {
        "amount": "3.14",
        "currency": "USD"
      },
      "fees": {
        "amount": "3.14",
        "currency": "USD"
      },
      "reason": "payout_returned",
      "status": "pending",
      "createDate": "2020-04-10T02:13:30.000Z",
      "updateDate": "2020-04-10T02:13:30.000Z"
    }
  ]
}`)),
				},
			},
			expected: []*circlesdk.PayoutReturn{
				{
					ID:       "b8627ae8-732b-4d25-b947-1df8f4007a29",
					PayoutID: "abdb500d-4a59-457c-801f-2d418c8703ac",
					Amount: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Fees: &circlesdk.Amount{
						Amount:   "3.14",
						Currency: "USD",
					},
					Reason:     "payout_returned",
					Status:     circlesdk.PayoutReturnStatusPending,
					CreateDate: "2020-04-10T02:13:30.000Z",
					UpdateDate: "2020-04-10T02:13:30.000Z",
				},
			},
		},
	}

	for _, tt := range tests {
		api := API{cl: &circlesdk.Client{Conn: &http.Client{Transport: &tt.roundTripper}}}
		data, err := api.ListReturns(tt.withOptions)
		if tt.expectedErr != nil {
			assert.Equal(tt.expectedErr.Error(), err.Error())
		}
		assert.Equal(tt.expected, data)
	}
}
