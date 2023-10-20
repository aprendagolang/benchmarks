package connections

// User definition
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateUserTable creates table on sqlite3 database
func CreateUserTable() error {
	db, err := Open()
	if err != nil {
		return err
	}
	defer db.Close()

	// sql to create user table
	sql := `CREATE TABLE IF NOT EXISTS users(
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		username VARCHAR NOT NULL,
    		password VARCHAR NOT NULL);`

	// execute sql
	_, err = db.Exec(sql)

	return err
}
