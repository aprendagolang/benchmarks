package std

import (
	"database/sql"
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
)

func InsertStd(db *sql.DB, data entities.Category) (entities.Category, error) {
	err := db.QueryRow("INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id", data.Name, data.Description).Scan(&data.ID)
	if err != nil {
		return data, err
	}

	return data, nil
}
