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
	{"one.2.three.4.com", true},
	{"invalid.bit", false},
}

func Test_Domain_1(t *testing.T) {
	v := validate.Validator{}
	for i, d := range domainTests {
		domain := NewDomain(d.Domain)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf("%d. IsValid(\"%s\") returned %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

/*
func Benchmark_Domain_1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, d := range domainTests {
			IsDomain(d.Domain)
		}
	}
}
*/
