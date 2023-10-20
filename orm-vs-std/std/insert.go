package std

import (
	"database/sql"
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
)

func InsertStd(db *sql.DB, data entities.Category) (entities.Category, error) {
	result, err := db.Exec("INSERT INTO categories (name, description) VALUES ($1, $2)", data.Name, data.Description)
	if err != nil {
		return data, err
	}

	data.ID, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return data, nil
}
