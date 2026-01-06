package main

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

/* Fixtures: Setup & Teardown
Scope:
- Function
	- setup: call a function
	- teardown: defer, t.Cleanup
- Group of tests / suite
	- Like the above with t.Run
- Package
	- TestMain
- All tests / session
	- No solution in in Go SDK, write your own code around the tests
*/

func setupDB(t *testing.T) *DB {
	dbFile := path.Join(t.TempDir(), "unter.db")
	db, err := NewDB(dbFile)
	require.NoError(t, err)
	t.Cleanup(func() { db.Close() })
	return db
}

func TestFunctionScope(t *testing.T) {
	db := setupDB(t)
	_ = db // ...
}

func testA(t *testing.T, db *DB) {}
func testB(t *testing.T, db *DB) {}

func TestGroupScope(t *testing.T) {
	db := setupDB(t)

	t.Run("TestA", func(t *testing.T) { testA(t, db) })
	t.Run("TestB", func(t *testing.T) { testB(t, db) })
}

func TestMain(m *testing.M) {
	// setup
	code := m.Run()
	// teardown

	os.Exit(code)
}
