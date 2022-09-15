package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/eriner/burr/internal/ent"
	"github.com/eriner/burr/internal/ent/enttest"
	_ "github.com/lib/pq"
)

func PostgresTest(t testing.TB) (*ent.Client, error) {
	host := "127.0.0.1"
	if os.Getenv("CI") != "" {
		host = "postgres"
	}
	return enttest.Open(t, "postgres", fmt.Sprintf("host=%s port=%d user=pgtest dbname=pgtest password=pgtest sslmode=disable", host, 5432)), nil
}
