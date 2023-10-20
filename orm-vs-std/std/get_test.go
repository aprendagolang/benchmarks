package std

import "testing"

func BenchmarkGetStd(b *testing.B) {
	db := setup()

	categories := InsertBenchStd(db, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetStd(db, categories[i].ID)

		b.StopTimer()
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}

}
