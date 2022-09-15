package db

import "testing"

var testDB func(t testing.TB) database

func TestDriver(t *testing.T) {
	if db == nil {
		t.Fatal("db package requires a driver to be selected via go tags")
	}
}
