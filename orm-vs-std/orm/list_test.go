package orm

import "testing"

func BenchmarkListORM(b *testing.B) {
	c := setup()

	_ = InsertBenchORM(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := ListORM(c)

		b.StopTimer()
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}
}
