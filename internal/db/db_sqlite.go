//go:build sqlite && !postgres

package db

import (
	"fmt"

	"github.com/eriner/burr/internal/ent"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db = func() database { return &client{} }
}

var _ database = &client{}

type client struct {
	ec *ent.Client
}

func (c *client) Open(_ map[string]any) (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:burr.db?mode=rwc&cache=shared&_fk=1")
	if err != nil {
		return nil, fmt.Errorf("failed connecting to the sqlite database: %w", err)
	}
	c.ec = client
	return client, nil
}
