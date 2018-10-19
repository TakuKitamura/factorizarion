package factorization

import (
	"math/big"
	"testing"
)

func BenchmarkUintIsPrime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		UintIsPrime(uint64(i))
	}
}

func BenchmarkUintPollardsRho(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		UintPollardsRho(uint64(i))
	}
}

func BenchmarkUintTrialDivision(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		UintTrialDivision(uint64(i))
	}
}

func BenchmarkUintMain(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100000; i++ {
		UintMain(uint64(i))
	}
}

func BenchmarkBigintIsPrime(b *testing.B) {
	b.ResetTimer()
	start := new(big.Int)
	start, _ = start.SetString("100000000000000000000", 10)

	end := new(big.Int)
	end, _ = end.SetString("100000000000000000010", 10)

	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, big.NewInt(1)) {
		BigintIsPrime(*i)
	}
}

func BenchmarkBigintPollardsRho(b *testing.B) {
	b.ResetTimer()
	start := new(big.Int)
	start, _ = start.SetString("100000000000000000000", 10)

	end := new(big.Int)
	end, _ = end.SetString("100000000000000000010", 10)
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, big.NewInt(1)) {
		BigintPollardsRho(*i)
	}
}

func BenchmarkBigintTrialDivision(b *testing.B) {
	b.ResetTimer()
	start := new(big.Int)
	start, _ = start.SetString("100000000000000000000", 10)

	end := new(big.Int)
	end, _ = end.SetString("100000000000000000010", 10)
	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, big.NewInt(1)) {
		BigintTrialDivision(*i)
	}
}

func BenchmarkBigintMain(b *testing.B) {
	start := new(big.Int)
	start, _ = start.SetString("100000000000000000000", 10)

	end := new(big.Int)
	end, _ = end.SetString("100000000000000000010", 10)

	for i := new(big.Int).Set(start); i.Cmp(end) < 0; i.Add(i, big.NewInt(1)) {
		BigintMain(*i)
	}
}
