package std

import (
	"database/sql"
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
)

func GetStd(db *sql.DB, id int64) (*entities.Category, error) {
	var category entities.Category
	err := db.QueryRow("SELECT id, name, description FROM categories WHERE id = $1", id).
		Scan(&category.ID, &category.Name, &category.Description)

	if err != nil {
		return nil, err
	}

	return &category, nil
}
