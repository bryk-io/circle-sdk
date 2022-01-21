package circlesdk

import (
	"testing"
)

func TestClient(t *testing.T) {
	opts := []Option{
		WithDebug(),
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
}
