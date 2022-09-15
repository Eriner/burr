//go:build vault

package secrets

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/vault/api"
)

func Open(cfg map[string]any) (Store, error) {
	if len(cfg) == 0 {
		return nil, fmt.Errorf("vault requires configuration options but none were given")
	}
	addr, ok := cfg["addr"].(string)
	if !ok {
		return nil, fmt.Errorf("vault requires \"addr\" configuration option")
	}
	// NOTE: token should be sourced from FS with a check that it is inaccessbile after the read,
	// or from stdin. This is bad practice and should be removed before release.
	token, ok := cfg["token"].(string)
	if !ok {
		return nil, fmt.Errorf("vault requires \"token\" but none was given.")
	}
	hc, ok := cfg["client"].(*http.Client)
	if !ok || hc == nil {
		hc = &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				// BUG: Disabling TLS validation for dev testing
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}
	kv, ok := cfg["kv"].(string)
	if !ok {
		kv = "burr_kv"
	}
	cfg = nil // GC
	client, err := api.NewClient(&api.Config{
		Address:    addr,
		HttpClient: hc,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create vault API client: %w", err)
	}
	client.SetToken(token)
	resp, err := client.Sys().Health()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to vault server: %w", err)
	}
	v := &vault{addr: addr, vault: client.Logical(), sys: client.Sys()}
	log.Printf("Connected to Vault server %q at %s", resp.ClusterName, addr)
	if err := v.InitKV(kv); err != nil {
		return nil, fmt.Errorf("failed to initilaize Vault KV: %w", err)
	}
	return v, nil
}

var _ Store = (*vault)(nil)

type vault struct {
	addr  string
	vault *api.Logical
	sys   *api.Sys
}

func (v *vault) InitKV(path string) error {
	mc, err := v.sys.MountConfig(path)
	if err != nil {
		err = v.sys.Mount(path, &api.MountInput{
			Type: "kv",
		})
		if err != nil {
			return fmt.Errorf("failed to create %q KV backend", path)
		}
		mc, err = v.sys.MountConfig(path)
		if err != nil {
			return fmt.Errorf("failed to get %q KV backend after create", path)
		}
	}
	_ = mc
	return nil
}

func (vc *vault) WriteWithContext(ctx context.Context, path string, data map[string]any) (map[string]any, error) {
	s, err := vc.vault.WriteWithContext(ctx, path, data)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}
	return s.Data, nil
}

func (vc *vault) ReadWithContext(ctx context.Context, path string) (map[string]any, error) {
	s, err := vc.vault.ReadWithContext(ctx, path)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, nil
	}
	return s.Data, nil
}

func (vc *vault) DeleteWithContext(ctx context.Context, path string) error {
	_, err := vc.vault.DeleteWithContext(ctx, path)
	return err
}
