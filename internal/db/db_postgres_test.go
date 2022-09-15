//go:build prod

package db

import (
	"fmt"
	"os"
	"testing"

	"github.com/eriner/burr/internal/ent"
	"github.com/eriner/burr/internal/ent/enttest"
	_ "github.com/lib/pq"
)

func init() {
	testDB = func(t testing.TB) database { return &testClient{t: t} }
}

var _ database = (*testClient)(nil)

type testClient struct {
	t  testing.TB
	ec *ent.Client
}

func (t *testClient) Open() (*ent.Client, error) {
	host := "127.0.0.1"
	if os.Getenv("CI") != "" {
		host = "postgres"
	}
	if t.ec != nil {
		return t.ec, nil
	}
	return enttest.Open(t.t, "postgres", fmt.Sprintf("host=%s port=%s user=pgtest dbname=pgtest password=pgtest sslmode=disable", host, 5432)), nil
}

func (t *testClient) Close() error {
	if t.ec != nil {
		t.ec.Close()
	}
	return nil
}
