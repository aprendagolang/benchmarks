package orm

import (
	"github.com/aprendagolang/benchmark-orm-vs-std/entities"
	"testing"
)

func BenchmarkInsertORM(b *testing.B) {
	db := setup()

	category := entities.Category{
		Name:        "Category 1",
		Description: "Description 1",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := InsertORM(db, category)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
