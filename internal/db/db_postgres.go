//go:build postgres && !sqlite

package db

import (
	"fmt"

	"github.com/eriner/burr/internal/ent"
)

func init() {
	db = func() database { return &client{} }
}

var _ database = (*client)(nil)

type client struct {
	ec *ent.Client
}

func (c *client) Open(cfg map[string]any) (*ent.Client, error) {
	if c.ec != nil {
		return c.ec, nil
	}
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		cfg["host"], cfg["port"], cfg["user"], cfg["database"], cfg["password"]))
	if err != nil {
		return nil, fmt.Errorf("failed connecting to the postgres database %s@%s:%s: %w", cfg["user"], cfg["host"], cfg["port"], err)
	}
	cfg = nil // contained sensitive info. trash it
	c.ec = client
	return client, nil
}
