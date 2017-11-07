package prime

import "testing"

func BenchmarkFindPrimes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindPrimes(1000)
	}
}
