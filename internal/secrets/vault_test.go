//go:build vault

package secrets_test

import (
	"os"
	"testing"

	"github.com/eriner/burr/internal/secrets"
	"github.com/stretchr/testify/assert"
)

func TestVault(t *testing.T) {
	assert := assert.New(t)
	addr := "http://127.0.0.1:8200"
	if len(os.Getenv("CI")) > 0 {
		addr = "http://vault:8200"
	}
	s, err := secrets.Open(map[string]any{
		"addr":  addr,
		"token": "root",
		"kv":    "test_kv",
	})
	assert.NoError(err)
	assert.NotNil(s)
	testBasic(t, s)
}
