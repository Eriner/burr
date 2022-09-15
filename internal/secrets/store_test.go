package secrets_test

import (
	"context"
	"testing"

	"github.com/eriner/burr/internal/secrets"
	"github.com/stretchr/testify/assert"
)

func testBasic(t *testing.T, s secrets.Store) {
	assert := assert.New(t)
	ctx := context.Background()
	_, err := s.WriteWithContext(ctx, "/test_kv/foo", map[string]any{
		"data": "bar",
	})
	assert.NoError(err)
	dat, err := s.ReadWithContext(ctx, "/test_kv/foo")
	assert.NoError(err)
	assert.Equal("bar", dat["data"])
	err = s.DeleteWithContext(ctx, "/test_kv/foo")
	assert.NoError(err)
	dat, err = s.ReadWithContext(ctx, "/test_kv/foo")
	assert.NoError(err)
	assert.Nil(dat)
}
