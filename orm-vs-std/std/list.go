package std

import (
	"database/sql"
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
)

func ListStd(db *sql.DB) ([]entities.Category, error) {
	var categories []entities.Category
	rows, err := db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return categories, nil
}
