package orm

import (
	"fmt"
	"testing"
)

func BenchmarkUpdateORM(b *testing.B) {
	db := setup()

	categories := InsertBenchORM(db, b.N)

	for i := 0; i < b.N; i++ {
		category := categories[i]
		category.Name = fmt.Sprintf("Update Category %d", i)
		b.ResetTimer()

		_ = UpdateORM(db, category)
	}
}
