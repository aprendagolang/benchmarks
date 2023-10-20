package std

import "testing"

func BenchmarkDeleteStd(b *testing.B) {
	db := setup()

	categories := InsertBenchStd(db, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := DeleteStd(db, categories[i].ID)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
