package std

import (
	"fmt"
	"testing"
)

func BenchmarkUpdateStd(b *testing.B) {
	db := setup()

	categories := InsertBenchStd(db, b.N)

	for i := 0; i < b.N; i++ {
		category := categories[i]
		category.Name = fmt.Sprintf("Update Category %d", i)
		b.ResetTimer()

		_ = UpdateStd(db, category)
	}
}
