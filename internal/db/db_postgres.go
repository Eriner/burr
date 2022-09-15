package db

import (
	"fmt"

	"github.com/eriner/burr/internal/ent"
)

func Postgres(host, port, user, password, database string) (*ent.Client, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		host, port, user, database, password))
	if err != nil {
		return nil, fmt.Errorf("failed connecting to the postgres database %s@%s:%s: %w", user, host, port, err)
	}
	return client, initDB(client)
}
