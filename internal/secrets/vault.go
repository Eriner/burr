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

func Vault(addr, token string, hc *http.Client) (*vault, error) {
	if hc == nil {
		hc = &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				// BUG: Disabling TLS validation for dev testing
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		}
	}
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
	log.Printf("Connected to Vault server %q at %s", resp.ClusterName, addr)
	return &vault{addr: addr, vault: client.Logical(), sys: client.Sys()}, nil
}

var _ Store = (*vault)(nil)

type vault struct {
	addr  string
	vault *api.Logical
	sys   *api.Sys
}

func (v *vault) Init() error {
	// TODO: create necessary secrets and storage backends
	return nil
}

func (vc *vault) WriteWithContext(ctx context.Context, path string, data map[string]any) (map[string]any, error) {
	s, err := vc.vault.WriteWithContext(ctx, path, data)
	if err != nil {
		return nil, err
	}
	return s.Data, nil
}

func (vc *vault) ReadWithContext(ctx context.Context, path string) (map[string]any, error) {
	s, err := vc.vault.ReadWithContext(ctx, path)
	if err != nil {
		return nil, err
	}
	return s.Data, nil
}

func (vc *vault) DeleteWithContext(ctx context.Context, path string) error {
	_, err := vc.vault.DeleteWithContext(ctx, path)
	return err
}
