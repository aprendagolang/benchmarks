package std

import (
	"database/sql"
	"fmt"
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
)

func InsertBenchStd(db *sql.DB, size int) []entities.Category {
	categories := make([]entities.Category, size)
	for i := 0; i < size; i++ {
		category := entities.Category{
			Name:        fmt.Sprintf("Category %d", i),
			Description: fmt.Sprintf("Description %d", i),
		}

		cat, err := InsertStd(db, category)
		if err != nil {
			panic(err)
		}

		categories[i] = cat
	}

	return categories
}
