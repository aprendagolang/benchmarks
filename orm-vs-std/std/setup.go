package std

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// Open a connection with a sqlite3 database
func setup() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=benchmark sslmode=disable")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS categories (id SERIAL, name TEXT, description TEXT, PRIMARY KEY (id))")
	if err != nil {
		panic(err)
	}

	return db
}
