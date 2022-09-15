//go:build !postgres

package db

import (
	"testing"

	"github.com/eriner/burr/internal/ent"
	"github.com/eriner/burr/internal/ent/enttest"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	testDB = func(t testing.TB) database { return &testClient{t: t} }
}

var _ database = &testClient{}

type testClient struct {
	t  testing.TB
	ec *ent.Client
}

func (tc *testClient) Open(_ map[string]any) (*ent.Client, error) {
	ec := enttest.Open(tc.t, "sqlite3", "file:burr-test.db?mode=memory&cache=shared&_fk=1")
	tc.ec = ec
	return ec, nil
}

func (tc *testClient) Close() error {
	if tc.ec == nil {
		return nil
	}
	return tc.ec.Close()
}
