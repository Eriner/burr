package kv

import "time"

// builtin is a poor-man's key-value cache backed by a Go map. Unlike a real
// LRU cache, there is zero optimization and keys are never evicted from the
// cache. This is used solely for testing and development when a redis cache
// is not available.
type builtin struct {
	m map[string]*data
}

func (b *builtin) Create(key string, value any, expiry time.Time) error {
	b.m[key] = &data{
		v: value,
		e: expiry,
	}
	return nil
}

func (b *builtin) Read(key string) (value any, expiry time.Time, err error) {
	if b.m[key] == nil {
		return nil, time.Now(), nil
	}
	return b.m[key].v, b.m[key].e, nil
}

func (b *builtin) Update(key string, value any, expiry time.Time) error {
	return b.Create(key, value, expiry)
}

func (b *builtin) Destroy(key string) error {
	b.m[key] = nil
	return nil
}

type data struct {
	v any
	e time.Time
}

func Builtin() (KV, error) {
	return &builtin{make(map[string]*data)}, nil
}
