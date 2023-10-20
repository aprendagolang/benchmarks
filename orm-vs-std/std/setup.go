package std

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Open a connection with a sqlite3 database
func setup() *sql.DB {
	db, err := sql.Open("sqlite3", "bench_std.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS categories (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, description TEXT)")
	if err != nil {
		panic(err)
	}

	return db
}
