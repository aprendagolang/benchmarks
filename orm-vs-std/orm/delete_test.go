package orm

import "testing"

func BenchmarkDeleteORM(b *testing.B) {
	c := setup()

	categories := InsertBenchORM(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := DeleteORM(c, categories[i].ID)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
