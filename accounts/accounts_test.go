package accounts

import (
	"testing"

	circlesdk "github.com/bryk-io/circle-sdk"
	ac "github.com/stretchr/testify/assert"
)

func TestAccounts(t *testing.T) {
	assert := ac.New(t)
	opts := []circlesdk.Option{
		// WithDebug(),
		circlesdk.WithUserAgent("circle-sdk-testing/0.1.0"),
		circlesdk.WithKeepAlive(10),
		circlesdk.WithMaxConnections(10),
		circlesdk.WithTimeout(10),
		circlesdk.WithAPIKeyFromEnv("CIRCLE_API_KEY"),
	}
	cl, err := circlesdk.NewClient(opts...)
	assert.Nil(err)
	accounts := API{cl: cl}
	t.Run("CreateWallet", func(t *testing.T) {
		t.Skip()
		ik := circlesdk.NewIdempotencyKey()
		wallet, err := accounts.CreateWallet("sample-test-wallet", circlesdk.WithIdempotencyKey(ik))
		assert.Nil(err)
		t.Logf("%+v", wallet)
	})

	t.Run("CreateWalletDepositAddress", func(t *testing.T) {
		t.Skip()
		ik := circlesdk.NewIdempotencyKey()
		addr, err := accounts.CreateWalletDepositAddress(
			"1000640329",
			circlesdk.USD,
			circlesdk.ChainALGO,
			circlesdk.WithIdempotencyKey(ik))
		assert.Nil(err)
		t.Logf("%+v", addr)
	})

	t.Run("GetWalletsList", func(t *testing.T) {
		list, err := accounts.GetWalletsList()
		assert.Nil(err)
		assert.NotEmpty(list)
		for _, wallet := range list {
			t.Logf("wallet: %s (%s) %s", wallet.ID, wallet.Kind, wallet.Entity)
		}
	})

	t.Run("GetWallet", func(t *testing.T) {
		wallet, err := accounts.GetWallet("1000581320")
		assert.Nil(err)
		assert.Equal(wallet.Kind, "merchant")
		assert.Equal(wallet.Entity, "26af371d-d4ac-4fef-a87c-da2e2f549f39")
	})

	t.Run("GetWalletAddressList", func(t *testing.T) {
		list, err := accounts.GetWalletAddressList("1000581320")
		assert.Nil(err)
		for _, addr := range list {
			t.Logf("%s (%s) on %s", addr.Address, addr.Currency, addr.Chain)
		}
	})

	t.Run("GetTransfersList", func(t *testing.T) {
		list, err := accounts.GetTransfersList("1000581320")
		assert.Nil(err)
		t.Logf("%+v", list)
	})
}
