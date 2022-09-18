//go:build ignore
package main

import (
	"context"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/eriner/burr/internal/ent/migrate"
)

func main() {
	ctx := context.Background()
	dir, err := atlas.NewGolangMigrateDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed to create the atlas migration directory: %v", err)
	}
	opts := []schema.MigrateOption{
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.Postgres),
	}
	if len(os.Args) != 2 {
		log.Fatalf("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}
	// Atlas runs migrations against a local database to validate them before they're committed.
	// NOTE: these credentials are for a local database. This has no security implication.
	if err := migrate.NamedDiff(ctx, "postgres://root:pass@localhost:6543/test", os.Args[1], opts...); err != nil {
		log.Fatalf("failed to generate migration file: %w", err)
	}
}
