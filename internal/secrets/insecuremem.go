//go:build !vault

package secrets

import (
	"context"
)

func Open(_ map[string]any) (Store, error) {
	return make(insecuremem), nil
}

// insecuremem is only used for local development and non-release builds where
// vault is not necessary.
type insecuremem map[string]map[string]any

func (i insecuremem) WriteWithContext(ctx context.Context, path string, opt map[string]any) (map[string]any, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		i[path] = opt
		return opt, nil
	}
}

func (i insecuremem) ReadWithContext(ctx context.Context, path string) (map[string]any, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		if i[path] == nil {
			return nil, nil
		}
		return i[path], nil
	}
}

func (i insecuremem) DeleteWithContext(ctx context.Context, path string) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		i[path] = nil
		return nil
	}
}
