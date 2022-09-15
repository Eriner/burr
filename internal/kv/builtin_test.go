//go:build !redis

package kv_test

import (
	"testing"
	"time"

	"github.com/eriner/burr/internal/kv"
	"github.com/stretchr/testify/assert"
)

func TestBuiltinDriver(t *testing.T) {
	assert := assert.New(t)
	c, err := kv.Open(nil)
	assert.NoError(err)
	err = c.Create("foo", "value", time.Now().Add(4*time.Minute))
	assert.NoError(err)
	v, _, err := c.Read("foo")
	assert.NoError(err)
	assert.Equal("value", v)
	err = c.Update("foo", "newvalue", time.Now().Add(8*time.Minute))
	assert.NoError(err)
	v, _, err = c.Read("foo")
	assert.NoError(err)
	assert.Equal("newvalue", v)
	err = c.Destroy("foo")
	assert.NoError(err)
	v, _, err = c.Read("foo")
	assert.Nil(v)
}
