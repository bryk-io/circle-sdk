package circlesdk

import (
	"context"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	opts := []Option{
		// WithDebug(),
		WithUserAgent("circle-sdk-testing/0.1.0"),
		WithKeepAlive(10),
		WithMaxConnections(10),
		WithTimeout(10),
		WithAPIKeyFromEnv("CIRCLE_API_KEY"),
	}
	cl, err := NewClient(opts...)
	if err != nil {
		t.Error(err)
		return
	}

	t.Run("Ping", func(t *testing.T) {
		res := cl.Ping()
		if !res {
			t.Error("ping failed")
		}
	})

	t.Run("GetMasterWalletID", func(t *testing.T) {
		id, err := cl.GetMasterWalletID()
		if err != nil {
			t.Error(err)
		}
		t.Logf("master wallet id: %s", id)
	})

	t.Run("GetBalance", func(t *testing.T) {
		balance, err := cl.GetBalance()
		if err != nil {
			t.Error(err)
		}
		t.Logf("balance: %+v", balance)
	})

	t.Run("CreateDepositAddress", func(t *testing.T) {
		t.Skip()
		ik := NewIdempotencyKey()
		address, err := cl.CreateDepositAddress(USD, ChainALGO, WithIdempotencyKey(ik))
		if err != nil {
			t.Error(err)
		}
		t.Logf("address: %+v", address)
	})

	t.Run("GetDepositAddress", func(t *testing.T) {
		list, err := cl.GetDepositAddress()
		if err != nil {
			t.Error(err)
		}
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
		if err != nil {
			t.Error(err)
		}
		t.Logf("recipient: %+v", id)
	})

	t.Run("GetRecipientAddress", func(t *testing.T) {
		toDate := time.Now().UTC()
		fromDate := toDate.Add(time.Hour * -72) // 3 days ago
		opts := []CallOption{
			WithContext(context.Background()),
			WithPageSize(15),
			WithDateRange(fromDate, toDate),
		}
		list, err := cl.GetRecipientAddress(opts...)
		if err != nil {
			t.Error(err)
		}
		for _, addr := range list {
			t.Logf("chain: %s, currency: %s, address: %s, desc: %s\n", addr.Chain, addr.Currency, addr.Address, addr.Description)
		}
	})
}
