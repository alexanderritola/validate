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

func Benchmark_ValidatePrintableRunes_1(b *testing.B) {
	b.StopTimer()
	by := []byte("goissocool")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidatePrintableRunes(by)
	}
}

func Benchmark_ValidatePrintableRunes_2(b *testing.B) {
	b.StopTimer()
	by := []byte("howimetyourmotherisagreatshow")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidatePrintableRunes(by)
	}
}

func Benchmark_ValidatePrintableRunes_3(b *testing.B) {
	b.StopTimer()
	by := []byte("openupmyphoneanditsallwetinside")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidatePrintableRunes(by)
	}
}
