package orm

import (
	"fmt"
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
	"gorm.io/gorm"
)

func InsertBenchORM(db *gorm.DB, size int) []entities.Category {
	categories := make([]entities.Category, size)
	for i := 0; i < size; i++ {
		category := entities.Category{
			Name:        fmt.Sprintf("Category %d", i),
			Description: fmt.Sprintf("Description %d", i),
		}

		cat, err := InsertORM(db, category)
		if err != nil {
			panic(err)
		}

		categories[i] = cat
	}

	return categories
}
