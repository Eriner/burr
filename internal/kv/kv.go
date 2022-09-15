package kv

import "time"

// KV implements CRUD interface with an expiry
type KV interface {
	Create(key string, value any, expiry time.Time) error
	Read(key string) (value any, expiry time.Time, err error)
	Update(key string, value any, expiry time.Time) error
	Destroy(key string) error
}
