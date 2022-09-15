package db

import (
	"testing"

	"github.com/eriner/burr/internal/ent"
	"github.com/eriner/burr/internal/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
)

func SqliteTest(t testing.TB) (*ent.Client, error) {
	ec := enttest.Open(t, "sqlite3", "file:burr-test.db?mode=memory&cache=shared&_fk=1")
	return ec, nil
}
