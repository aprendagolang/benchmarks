package std

import "testing"

func BenchmarkListStd(b *testing.B) {
	c := setup()

	_ = InsertBenchStd(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ListStd(c)

		b.StopTimer()
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}
