package std

import "testing"

func BenchmarkListStd(b *testing.B) {
	db := setup()

	_ = InsertBenchStd(db, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ListStd(db)

		b.StopTimer()
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}
