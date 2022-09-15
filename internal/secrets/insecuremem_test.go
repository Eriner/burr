//go:build !vault

package secrets_test

import (
	"testing"

	"github.com/eriner/burr/internal/secrets"
	"github.com/stretchr/testify/assert"
)

func TestInsecuremem(t *testing.T) {
	assert := assert.New(t)
	s, err := secrets.Open(nil)
	assert.NoError(err)
	assert.NotNil(s)

	testBasic(t, s)
}
