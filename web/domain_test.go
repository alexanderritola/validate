package web

import (
	"github.com/daswasser/validate"
	"testing"
)

var domainTests = []struct {
	Domain string
	Valid  bool
}{
	{"code.google.com", true},
	{"code.google..com", false},
	{".com", false},
	{"com", true},
	{"one.2.three.4.com", true},
	{"invalid.bit", false},
}

func Test_Domain_Validate(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range domainTests {
		domain := NewDomain(d.Domain)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf("%d. IsValid(\"%s\") returned %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

func Benchmark_Domain_Validate(b *testing.B) {
	b.StopTimer()
	v := validate.NewValidator()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		for _, d := range domainTests {
			domain := NewDomain(d.Domain)
			v.Validate(domain)
		}
	}
}
