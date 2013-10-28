package validate

import (
	"testing"
)

func Benchmark_ValidateLowAlphabet_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet("goissocool")
	}
}

func Benchmark_ValidateLowAlphabet_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet("howimetyourmotherisagreatshow")
	}
}

func Benchmark_ValidateLowAlphabet_3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet("openupmyphoneanditsallwetinside")
	}
}
