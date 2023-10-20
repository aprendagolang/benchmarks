package connections

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Open a connection with a sqlite3 database
func Open() (*sql.DB, error) {
	return sql.Open("sqlite3", "test.db")
}
