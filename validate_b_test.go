package validate

import (
	"testing"
)

func Benchmark_ValidateLowAlphabet_1(b *testing.B) {
	b.StopTimer()
	by := []byte("goissocool")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet(by)
	}
}

func Benchmark_ValidateLowAlphabet_2(b *testing.B) {
	b.StopTimer()
	by := []byte("howimetyourmotherisagreatshow")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet(by)
	}
}

func Benchmark_ValidateLowAlphabet_3(b *testing.B) {
	b.StopTimer()
	by := []byte("openupmyphoneanditsallwetinside")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet(by)
	}
}

func Benchmark_ValidateLowAlphabet_2_1(b *testing.B) {
	b.StopTimer()
	by := []byte("goissocool")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet_2(by)
	}
}

func Benchmark_ValidateLowAlphabet_2_2(b *testing.B) {
	b.StopTimer()
	by := []byte("howimetyourmotherisagreatshow")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet_2(by)
	}
}

func Benchmark_ValidateLowAlphabet_2_3(b *testing.B) {
	b.StopTimer()
	by := []byte("openupmyphoneanditsallwetinside")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateLowAlphabet_2(by)
	}
}
