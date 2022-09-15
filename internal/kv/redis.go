package kv

import (
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

func Redis(addr, password string) (KV, error) {
	client := api.NewClient(&api.Options{
		Addr:     addr,
		Password: password,
	})
	return &redis{client}, nil
}
