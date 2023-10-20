package std

import (
	"database/sql"
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
)

func UpdateStd(db *sql.DB, data entities.Category) error {
	_, err := db.Exec("UPDATE categories SET name = $1, description = $2 WHERE id = $3", data.Name, data.Description, data.ID)

	return err
}
