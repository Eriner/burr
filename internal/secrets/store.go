package secrets

import "context"

// Store is a secrets storage backend.
type Store interface {
	// Basic read, write, and delete methods.
	WriteWithContext(ctx context.Context, path string, opt map[string]any) (map[string]any, error)
	ReadWithContext(ctx context.Context, path string) (map[string]any, error)
	DeleteWithContext(ctx context.Context, path string) error
}
