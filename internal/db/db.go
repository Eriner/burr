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
				var name string
				if actor := policy.ActorFromContext(ctx); actor != nil {
					name = actor.Name
				}
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\tActor=%s\n",
					m.Op(), m.Type(), time.Since(start), m, name,
				)
			}()
			return next.Mutate(ctx, m)
		})
	})
	return nil
}
