package db

import (
	_ "embed"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
)

// Basically wraps db.New() with some additional configuration
func InitDB(name string) *Queries {
	db, err := sql.Open("sqlite3", name)
	if err != nil {
		Logger.Fatal("", err)
	}

	return New(db)
}
