package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eriner/burr/internal/ent"
	_ "github.com/eriner/burr/internal/ent/runtime"
	"github.com/eriner/burr/internal/policy"
)

// db is set in an init() function based on the build tags provided, postgres or sqlite.
var db func() database

// Open the database
func Open(cfg map[string]any) (*ent.Client, error) {
	if cfg == nil {
		cfg = make(map[string]any)
	}
	if db == nil {
		return nil, fmt.Errorf("burr was built without a database driver")
	}
	db := db()
	c, err := db.Open(cfg)
	if err != nil {
		return nil, err
	}
	err = initDB(c)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}
	return c, nil
}

// database is implemented by postgres and sqlite
type database interface {
	Open(cfg map[string]any) (*ent.Client, error)
}

// initDB sets runtime hooks, runs migrations, and other #dbthings
func initDB(c *ent.Client) error {
	if c == nil {
		return fmt.Errorf("nil client")
	}
	if err := c.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema: %w", err)
	}
	// Global runtime logging hook
	c.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				var email string
				if user := policy.UserFromContext(ctx); user != nil {
					email = user.Email
				}
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\tUser=%s\n",
					m.Op(), m.Type(), time.Since(start), m, email,
				)
			}()
			return next.Mutate(ctx, m)
		})
	})
	return nil
}
