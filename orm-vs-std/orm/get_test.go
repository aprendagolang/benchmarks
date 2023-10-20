package orm

import "testing"

func BenchmarkGetORM(b *testing.B) {
	c := setup()

	categories := InsertBenchORM(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := GetORM(c, categories[i].ID)

		b.StopTimer()
		if err != nil {
			b.Fatal(err)
		}
		b.StartTimer()
	}

}
