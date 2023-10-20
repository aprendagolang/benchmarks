package orm

import "testing"

func BenchmarkGetORM(b *testing.B) {
	db := setup()

	categories := InsertBenchORM(db, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetORM(db, categories[i].ID)

		b.StopTimer()
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}

}
