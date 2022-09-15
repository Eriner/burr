//go:build redis
package kv

import (
	"fmt"
	"time"

	api "github.com/go-redis/redis/v8"
)

var _ KV = (*redis)(nil)

type redis struct {
	c *api.Client
}

func (r *redis) Create(key string, value any, expiry time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (r *redis) Read(key string) (value any, expiry time.Time, err error) {
	panic("not implemented") // TODO: Implement
}

func (r *redis) Update(key string, value any, expiry time.Time) error {
	panic("not implemented") // TODO: Implement
}

func (r *redis) Destroy(key string) error {
	panic("not implemented") // TODO: Implement
}

// Open a new redis connection.
func Open(cfg map[string]any) (KV, error) {
	if len(cfg) == 0 {
		return nil, fmt.Errorf("redis requires configuration")
	}
	addr, ok := cfg["addr"].(string)
	if !ok {
		return nil, fmt.Errorf("redis configuration must include \"addr\"")
	}
	password, ok := cfg["password"].(string)
	if !ok {
		password = ""
	}
	client := api.NewClient(&api.Options{
		Addr:     addr,
		Password: password,
	})
	return &redis{client}, nil
}
