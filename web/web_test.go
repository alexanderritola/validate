package web

import (
	"fmt"
	"github.com/daswasser/validate"
	"testing"
)

func ExampleDomain_Validate() {
	// Setup a new validator
	v := validate.NewValidator()

	// Create a new Domain object and return the message on failure
	goodDomain :=
		NewDomain("www.golang.org").
			MaxSubdomains(2).
			SetMessage("Invalid domain specified!")

	badDomain :=
		NewDomain("gophersRock!com").SetMessage("Invalid domain specified!")

	// Validate the good domain
	err := v.Validate(goodDomain)
	if err != nil {
		fmt.Printf("%s error:\n", goodDomain)
		fmt.Println(err)
		fmt.Println(goodDomain.Message())
	} else {
		fmt.Printf("%s is a valid domain\n", goodDomain)
	}

	// Validate the bad domain
	err = v.Validate(badDomain)
	if err != nil {
		fmt.Printf("%s error:\n", badDomain)
		fmt.Println(err)
		fmt.Println(badDomain.Message())
	} else {
		fmt.Printf("%s is a valid domain\n", badDomain)
	}

	// Output:
	// www.golang.org is a valid domain
	// gophersRock!com error:
	// Error level 1: Invalid formatting
	// Invalid domain specified!
}

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
