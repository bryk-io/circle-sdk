package core

import (
	"context"
	"testing"
	"time"

	"github.com/bryk-io/circle-sdk"
	ac "github.com/stretchr/testify/assert"
)

func TestCore(t *testing.T) {
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
	core := API{cl: cl}
	assert.Nil(err)
	t.Run("Ping", func(t *testing.T) {
		res := core.Ping()
		assert.True(res)
	})

	t.Run("GetMasterWalletID", func(t *testing.T) {
		id, err := core.GetMasterWalletID()
		assert.Nil(err)
		assert.NotEmpty(id)
		t.Logf("master wallet id: %s", id)
	})

	t.Run("GetBalance", func(t *testing.T) {
		balance, err := core.GetBalance()
		assert.Nil(err)
		t.Logf("balance: %+v", balance)
	})

	t.Run("CreateDepositAddress", func(t *testing.T) {
		t.Skip()
		ik := circlesdk.NewIdempotencyKey()
		address, err := core.CreateDepositAddress(circlesdk.USD, circlesdk.ChainALGO, circlesdk.WithIdempotencyKey(ik))
		assert.Nil(err)
		t.Logf("address: %+v", address)
	})

	t.Run("GetDepositAddressList", func(t *testing.T) {
		list, err := core.GetDepositAddressList()
		assert.Nil(err)
		for _, addr := range list {
			t.Logf("chain: %s, currency: %s, address: %s\n", addr.Chain, addr.Currency, addr.Address)
		}
	})

	t.Run("AddRecipientAddress", func(t *testing.T) {
		t.Skip()
		ik := circlesdk.NewIdempotencyKey()
		addr := &circlesdk.DepositAddress{
			Address:  "25W7IOG63WJSPA63YN4XKCCXCE2ENZA3VBE4B6QRTKTZH3WVXU3Y474SOM",
			Currency: string(circlesdk.USD),
			Chain:    string(circlesdk.ChainALGO),
		}
		id, err := core.AddRecipientAddress(addr, "test account", circlesdk.WithIdempotencyKey(ik))
		assert.Nil(err)
		t.Logf("recipient: %+v", id)
	})

	t.Run("GetRecipientAddressList", func(t *testing.T) {
		toDate := time.Now().UTC()
		fromDate := toDate.Add(time.Hour * -72) // 3 days ago
		opts := []circlesdk.CallOption{
			circlesdk.WithContext(context.Background()),
			circlesdk.WithPageSize(15),
			circlesdk.WithDateRange(fromDate, toDate),
		}
		list, err := core.GetRecipientAddressList(opts...)
		assert.Nil(err)
		for _, addr := range list {
			t.Logf("chain: %s, currency: %s, address: %s, desc: %s\n", addr.Chain, addr.Currency, addr.Address, addr.Description)
		}
	})
}
