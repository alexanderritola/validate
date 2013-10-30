package simple

import (
	"testing"
)

func Benchmark_IsLower_1(b *testing.B) {
	b.StopTimer()
	by := []byte("abcdefghijklmnopqrstuvwxyz")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsLower(by)
	}
}

func Benchmark_IsLower_2(b *testing.B) {
	b.StopTimer()
	by := []byte("howimetyourmotherisagreatshow")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsLower(by)
	}
}

func Benchmark_IsLower_3(b *testing.B) {
	b.StopTimer()
	by := []byte("openupmyphoneanditsallwetinside")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsLower(by)
	}
}

func Benchmark_IsUpper_1(b *testing.B) {
	b.StopTimer()
	by := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsUpper(by)
	}
}

func Benchmark_IsUpper_2(b *testing.B) {
	b.StopTimer()
	by := []byte("HOWIMETYOURMOTHERISAGREATSHOW")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsUpper(by)
	}
}

func Benchmark_IsUpper_3(b *testing.B) {
	b.StopTimer()
	by := []byte("OPENUPMYPHONEANDITSALLWETINSIDE")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsUpper(by)
	}
}

func Benchmark_IsPrint_1(b *testing.B) {
	b.StopTimer()
	by := []byte("abcdefghijklmnopqrstuvwxyz")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsPrint(by)
	}
}

func Benchmark_IsPrint_2(b *testing.B) {
	b.StopTimer()
	by := []byte("howimetyourmotherisagreatshow")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsPrint(by)
	}
}

func Benchmark_IsPrint_3(b *testing.B) {
	b.StopTimer()
	by := []byte("openupmyphoneanditsallwetinside")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		IsPrint(by)
	}
}
