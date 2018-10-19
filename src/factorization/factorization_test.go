package factorization

import (
	"testing"
)

func BenchmarkUintIsPrime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 1000000; i++ {
		UintIsPrime(uint64(i))
	}
}

func BenchmarkUintPollardsRho(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 1000000; i++ {
		UintPollardsRho(uint64(i))
	}
}

func BenchmarkUintTrialDivision(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100000000; i++ {
		UintTrialDivision(uint64(i))
	}
}

func BenchmarkUintMain(b *testing.B) {
	// fmt.Println(b.N, 123)
	b.ResetTimer()
	for i := 0; i < 1000000; i++ {
		UintMain(uint64(i))
	}
}
