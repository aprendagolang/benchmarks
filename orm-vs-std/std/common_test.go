package std

import (
	"fmt"
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
)

func InsertBenchStd(size int) []entities.Category {
	c := setup()

	categories := make([]entities.Category, size)
	for i := 0; i < size; i++ {
		category := entities.Category{
			Name:        fmt.Sprintf("Category %d", i),
			Description: fmt.Sprintf("Description %d", i),
		}

		cat, err := InsertStd(c, category)
		if err != nil {
			panic(err)
		}

		categories[i] = cat
	}

	return categories
}
