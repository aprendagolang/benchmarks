package orm

import "testing"

func BenchmarkDeleteORM(b *testing.B) {
	db := setup()

	categories := InsertBenchORM(db, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := DeleteORM(db, categories[i].ID)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
