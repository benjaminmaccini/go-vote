package utils

import (
	"database/sql"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var (
	TestDBPath = filepath.Join(ROOT_DIR, "test.sqlite")
)

func SetupTeardown(tb testing.TB) func(tb testing.TB) {
	InitLogger("DEBUG")

	// Clear tables before each test
	// This is useful for debugging by investigating the state of the database
	clearTables()

	// Cleanup function
	return func(tb testing.TB) {}
}

func clearTables() {
	db, err := sql.Open("sqlite3", TestDBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Get all table names
	rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table' AND name != 'schema_migrations'")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			panic(err)
		}
		tables = append(tables, name)
	}

	// Clear all tables
	for _, table := range tables {
		_, err := db.Exec("DELETE FROM " + table)
		if err != nil {
			panic(err)
		}
	}
}
