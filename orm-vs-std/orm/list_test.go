package orm

import "testing"

func BenchmarkListORM(b *testing.B) {
	db := setup()

	_ = InsertBenchORM(db, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ListORM(db)

		b.StopTimer()
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}
