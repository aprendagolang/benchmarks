package std

import "testing"

func BenchmarkGetStd(b *testing.B) {
	c := setup()

	categories := InsertBenchStd(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetStd(c, categories[i].ID)

		b.StopTimer()
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}

}
