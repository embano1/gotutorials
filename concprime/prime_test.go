package concprime

import "testing"

func BenchmarkFindPrimes(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Find(1000, false)
	}
}
