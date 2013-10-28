package validate

import (
	"testing"
)

func Benchmark_ValidateLowAlphabet_1(b *testing.B) {
	b.StopTimer()
	by := []byte("abcdefghijklmnopqrstuvwxyz")
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

func Benchmark_ValidateUpAlphabet_1(b *testing.B) {
	b.StopTimer()
	by := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateUpAlphabet(by)
	}
}

func Benchmark_ValidateUpAlphabet_2(b *testing.B) {
	b.StopTimer()
	by := []byte("HOWIMETYOURMOTHERISAGREATSHOW")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateUpAlphabet(by)
	}
}

func Benchmark_ValidateUpAlphabet_3(b *testing.B) {
	b.StopTimer()
	by := []byte("OPENUPMYPHONEANDITSALLWETINSIDE")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		ValidateUpAlphabet(by)
	}
}

func Benchmark_ValidatePrintableRunes_1(b *testing.B) {
	b.StopTimer()
	by := []byte("abcdefghijklmnopqrstuvwxyz")
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
