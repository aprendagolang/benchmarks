package std

import "testing"

func BenchmarkDeleteStd(b *testing.B) {
	c := setup()

	categories := InsertBenchStd(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := DeleteStd(c, categories[i].ID)

		b.StopTimer()
		if err != nil {
			b.Error(err)
		}
		b.StartTimer()
	}
}
