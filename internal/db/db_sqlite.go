package db

import (
	"fmt"

	"github.com/eriner/burr/internal/ent"
	_ "github.com/mattn/go-sqlite3"
)

func Sqlite() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:burr.db?mode=rwc&cache=shared&_fk=1")
	if err != nil {
		return nil, fmt.Errorf("failed connecting to the sqlite database: %w", err)
	}
	return client, initDB(client)
}
