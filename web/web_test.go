package web

import (
	"testing"
)

var domainTests = []struct {
	Domain []byte
	Valid  bool
}{
	{[]byte("code.google.com"), true},
	{[]byte("code.google..com"), false},
	{[]byte(".com"), false},
	{[]byte("one.2.three.4.com"), true},
	{[]byte("invalid.bit"), false},
}

func Test_Domain_1(t *testing.T) {
	for i, d := range domainTests {
		if v := IsDomain(d.Domain); v != d.Valid {
			t.Errorf("%d. IsDomain(\"%s\") returned %v, want %v",
				i, d.Domain, v, d.Valid)
		}
	}
}

func Benchmark_Domain_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range domainTests {
			IsDomain(d.Domain)
		}
	}
}
