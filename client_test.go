package circlesdk

import (
	"context"
	"testing"
	"time"

	ac "github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	assert := ac.New(t)
	opts := []Option{
		// WithDebug(),
		WithUserAgent("circle-sdk-testing/0.1.0"),
		WithKeepAlive(10),
		WithMaxConnections(10),
		WithTimeout(10),
		WithAPIKeyFromEnv("CIRCLE_API_KEY"),
	}
	cl, err := NewClient(opts...)
	assert.Nil(err)

	t.Run("Core", func(t *testing.T) {
		t.Run("Ping", func(t *testing.T) {
			res := cl.Ping()
			assert.True(res)
		})

		t.Run("GetMasterWalletID", func(t *testing.T) {
			id, err := cl.GetMasterWalletID()
			assert.Nil(err)
			assert.NotEmpty(id)
			t.Logf("master wallet id: %s", id)
		})

		t.Run("GetBalance", func(t *testing.T) {
			balance, err := cl.GetBalance()
			assert.Nil(err)
			t.Logf("balance: %+v", balance)
		})

		t.Run("CreateDepositAddress", func(t *testing.T) {
			t.Skip()
			ik := NewIdempotencyKey()
			address, err := cl.CreateDepositAddress(USD, ChainALGO, WithIdempotencyKey(ik))
			assert.Nil(err)
			t.Logf("address: %+v", address)
		})

		t.Run("GetDepositAddressList", func(t *testing.T) {
			list, err := cl.GetDepositAddressList()
			assert.Nil(err)
			for _, addr := range list {
				t.Logf("chain: %s, currency: %s, address: %s\n", addr.Chain, addr.Currency, addr.Address)
			}
		})

		t.Run("AddRecipientAddress", func(t *testing.T) {
			t.Skip()
			ik := NewIdempotencyKey()
			addr := &DepositAddress{
				Address:  "25W7IOG63WJSPA63YN4XKCCXCE2ENZA3VBE4B6QRTKTZH3WVXU3Y474SOM",
				Currency: string(USD),
				Chain:    string(ChainALGO),
			}
			id, err := cl.AddRecipientAddress(addr, "test account", WithIdempotencyKey(ik))
			assert.Nil(err)
			t.Logf("recipient: %+v", id)
		})

		t.Run("GetRecipientAddressList", func(t *testing.T) {
			toDate := time.Now().UTC()
			fromDate := toDate.Add(time.Hour * -72) // 3 days ago
			opts := []CallOption{
				WithContext(context.Background()),
				WithPageSize(15),
				WithDateRange(fromDate, toDate),
			}
			list, err := cl.GetRecipientAddressList(opts...)
			assert.Nil(err)
			for _, addr := range list {
				t.Logf("chain: %s, currency: %s, address: %s, desc: %s\n", addr.Chain, addr.Currency, addr.Address, addr.Description)
			}
		})
	})

	t.Run("Accounts", func(t *testing.T) {
		t.Run("CreateWallet", func(t *testing.T) {
			t.Skip()
			ik := NewIdempotencyKey()
			wallet, err := cl.Accounts.CreateWallet("sample-test-wallet", WithIdempotencyKey(ik))
			assert.Nil(err)
			t.Logf("%+v", wallet)
		})

		t.Run("CreateWalletDepositAddress", func(t *testing.T) {
			t.Skip()
			ik := NewIdempotencyKey()
			addr, err := cl.Accounts.CreateWalletDepositAddress(
				"1000640329",
				USD,
				ChainALGO,
				WithIdempotencyKey(ik))
			assert.Nil(err)
			t.Logf("%+v", addr)
		})

		t.Run("GetWalletsList", func(t *testing.T) {
			list, err := cl.Accounts.GetWalletsList()
			assert.Nil(err)
			assert.NotEmpty(list)
			for _, wallet := range list {
				t.Logf("wallet: %s (%s) %s", wallet.ID, wallet.Kind, wallet.Entity)
			}
		})

		t.Run("GetWallet", func(t *testing.T) {
			wallet, err := cl.Accounts.GetWallet("1000581320")
			assert.Nil(err)
			assert.Equal(wallet.Kind, "merchant")
			assert.Equal(wallet.Entity, "26af371d-d4ac-4fef-a87c-da2e2f549f39")
		})

		t.Run("GetWalletAddressList", func(t *testing.T) {
			list, err := cl.Accounts.GetWalletAddressList("1000581320")
			assert.Nil(err)
			for _, addr := range list {
				t.Logf("%s (%s) on %s", addr.Address, addr.Currency, addr.Chain)
			}
		})

		t.Run("GetTransfersList", func(t *testing.T) {
			list, err := cl.Accounts.GetTransfersList("1000581320")
			assert.Nil(err)
			t.Logf("%+v", list)
		})
	})
}
