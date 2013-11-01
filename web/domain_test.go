package web

import (
	"fmt"
	"github.com/daswasser/validate"
	"testing"
)

type domainTest struct {
	Domain string
	Valid  bool
}

var invalidCharacters = []domainTest{
	{"a.com", true},
	{"a.b-c.com", true},
	{"a!b.com", false},
	{"a.com?", false},
	{"#hashtag.com", false},
}

// Test for setting maximum number of sub domains.
func Test_Domain_Validate_InvalidCharacters(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range invalidCharacters {
		domain := NewDomain(d.Domain)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf("%d. IsValid(\"%s\") error: %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

// Test the 127 sub-domain limit. Note that we have to add 'x' as a valid
// TLD to test for this otherwise its impossible to get 127 subdomains to
// fit in 255 characters.
var maxSubDomains = []domainTest{
	{"a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z.A.B.C.D.E.F.G.H.I.J.K.L.M.N.O.P.Q.R.S.T.U.V.W.X.Y.Z.a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z.A.B.C.D.E.F.G.H.I.J.K.L.M.N.O.P.Q.R.S.T.U.V.W.X.Y.Z.0.1.2.3.4.5.6.7.8.9.0.1.2.3.4.5.6.7.8.9.0.1.x.x", true},    // 127 sub-domains
	{"a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z.A.B.C.D.E.F.G.H.I.J.K.L.M.N.O.P.Q.R.S.T.U.V.W.X.Y.Z.a.b.c.d.e.f.g.h.i.j.k.l.m.n.o.p.q.r.s.t.u.v.w.x.y.z.A.B.C.D.E.F.G.H.I.J.K.L.M.N.O.P.Q.R.S.T.U.V.W.X.Y.Z.0.1.2.3.4.5.6.7.8.9.0.1.2.3.4.5.6.7.8.9.0.1.x.x.x", false}, // 128 sub-domains and length > 255
}

// Test for the maximum number of sub domains.
func Test_Domain_Validate_MaxSubDomains(t *testing.T) {
	// Append the .x TLD to the list so that we can test for all 127
	// sub domains and still fit within 255 characters.
	*TLDs = append(*TLDs, []byte("x"))
	v := validate.NewValidator()
	for i, d := range maxSubDomains {
		domain := NewDomain(d.Domain).MaxLength(257)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf("%d. IsValid(\"%s\") returned %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

var maxLength = []domainTest{
	{"g8j8lBR2ykFIdU7rfmKb0RDRODGT8g1xcxaAANU79QY23eG7apxmjCNXc4AkWqc.FKWtqTRiPqHVi3A9fuTeo8oudy8xGuv52MLj8J1OrGhNStXJ9GgrXBMtSNm1cW6.AiPbzMU8HhdAj3DklaF46e9ym0HsmaIwObZs5o2vwsLQuVfGfEzF3IxrTcjczgQ.Iu46lr4BF9Oo22cfUFNZBamgPBHif9iZ2WiuuOs5N2zNSe0t8FtJvUNocoY.com", true},
	{"g8j8lBR2ykFIdU7rfmKb0RDRODGT8g1xcxaAANU79QY23eG7apxmjCNXc4AkWqc.FKWtqTRiPqHVi3A9fuTeo8oudy8xGuv52MLj8J1OrGhNStXJ9GgrXBMtSNm1cW6.AiPbzMU8HhdAj3DklaF46e9ym0HsmaIwObZs5o2vwsLQuVfGfEzF3IxrTcjczgQ.Iu46lr4BF9Oo22cfUFNZBamgPBHif9iZ2WiuuOs5N2zNSe0t8FtJvUNocoYx.com", false},
}

// Test to make sure domains > 255 characters are invalid.
func Test_Domain_Validate_MaxLength(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range maxLength {
		domain := NewDomain(d.Domain)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			fmt.Println("ERROR:", d.Valid, err != nil, string(d.Domain))
			t.Errorf("%d. IsValid(\"%s\") error: %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

var maxSubs = []domainTest{
	{"a.com", true},
	{"a.b.com", true},
	{"a.b.c.com", false},
}

// Test for setting maximum number of sub domains.
func Test_Domain_Validate_MaxSubs(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range maxSubs {
		domain := NewDomain(d.Domain).MaxSubdomains(2)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf("%d. IsValid(\"%s\") error: %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

// Set max length to 5
var shorterMaxLength = []domainTest{
	{"a.com", true},
	{"ab.com", false},
}

// Test for setting custom domain max length.
func Test_Domain_Validate_ShorterMaxLength(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range shorterMaxLength {
		domain := NewDomain(d.Domain).MaxLength(5)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf("%d. IsValid(\"%s\") error: %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

var hyphens = []domainTest{
	{"a-b.com", true},
	{"a-.com", false},
	{"-a.com", false},
	{"a--b.com", true},
}

// Ensure labels do not start or end with hyphens.
func Test_Domain_ValidateHyphens(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range hyphens {
		domain := NewDomain(d.Domain)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf("%d. IsValid(\"%s\") returned %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

var tlds = []domainTest{
	{"golang.org", true},      // Test a known good TLD.
	{"oursite.test123", true}, // Test for our TLD that we will add to the list.
}

// Test for the maximum number of sub domains.
func Test_Domain_Validate_AddTLD(t *testing.T) {
	// Append the test123 TLD to the list.
	*TLDs = append(*TLDs, []byte("test123"))
	v := validate.NewValidator()
	for i, d := range tlds {
		domain := NewDomain(d.Domain)
		err := v.Validate(domain)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf("%d. IsValid(\"%s\") returned %v, want %v. Error: %v",
				i, d.Domain, err != nil, d.Valid, err)
		}
	}
}

// Package global variable to prevent compiler optimizations from removing
// pointless functions.
var benchErr validate.Error

func benchmarkDomain_Validate(domain string, b *testing.B) {
	var e validate.Error
	v := validate.NewValidator()
	for i := 0; i < b.N; i++ {
		d := NewDomain(domain)
		// Record th result to keep the compiler from eliminating
		// the function call.
		e = v.Validate(d)
	}
	// Save the result to package level variable to prevent
	// compiler from eliminating the benchmark entirely.
	benchErr = e
}

func Benchmark_Domain_ValidateTLD(b *testing.B) {
	benchmarkDomain_Validate("com", b)
}

func Benchmark_Domain_ValidateSimple(b *testing.B) {
	benchmarkDomain_Validate("golang.org", b)
}

func Benchmark_Domain_ValidateLessSimple(b *testing.B) {
	benchmarkDomain_Validate("www.golang.org", b)
}

func Benchmark_Domain_ValidateLongest(b *testing.B) {
	benchmarkDomain_Validate(maxLength[0].Domain, b)
}
func Benchmark_Domain_ValidateMostSubs(b *testing.B) {
	benchmarkDomain_Validate(maxSubDomains[0].Domain, b)
}
