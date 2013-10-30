package simple

import (
	"github.com/daswasser/validate"
	"testing"
)

var isPrintTests = []struct {
	Data  string
	Valid bool
}{
	{"go is cool!", true},
	{string('\u0000'), false},
	{"http://مثال.إختبار", true},
}

func Test_IsPrint(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range isPrintTests {
		data := NewPrintable([]byte(d.Data))
		err := v.Validate(data)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf(
				"%d. IsLower(\"%s\") Error=%v, Valid=%v. Error value: %s\n",
				i, d.Data, err != nil, d.Valid, err)
		}
	}
}

func Benchmark_IsPrint(b *testing.B) {
	b.StopTimer()
	v := validate.NewValidator()
	for _, d := range isPrintTests {
		data := NewPrintable([]byte(d.Data))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			v.Validate(data)
		}
		b.StopTimer()
	}
}

var isLowerTests = []struct {
	Data  string
	Valid bool
}{
	{"abcdefghijklmnopqrstuv", true},
	{"CRUISECONTROLFORCOOL", false},
	{"spaces are bad", false},
	{"so is punctuation!", false},
	{"thisisokthough", true},
}

func Test_IsLower(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range isLowerTests {
		data := NewLower([]byte(d.Data))
		err := v.Validate(data)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf(
				"%d. IsLower(\"%s\") Error=%v, Valid=%v. Error value: %s\n",
				i, d.Data, err != nil, d.Valid, err)
		}
	}

}

func Benchmark_IsLower(b *testing.B) {
	b.StopTimer()
	v := validate.NewValidator()
	for _, d := range isLowerTests {
		data := NewLower([]byte(d.Data))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			v.Validate(data)
		}
		b.StopTimer()
	}
}

var isUpperTests = []struct {
	Data  string
	Valid bool
}{
	{"abcdefghijklmnopqrstuv", false},
	{"CRUISECONTROLFORCOOL", true},
	{"spaces are bad", false},
	{"so is punctuation!", false},
	{"THISISOKTHOUGH", true},
}

func Test_IsUpper(t *testing.T) {
	v := validate.NewValidator()
	for i, d := range isUpperTests {
		data := NewUpper([]byte(d.Data))
		err := v.Validate(data)
		if (err != nil && d.Valid) || (err == nil && !d.Valid) {
			t.Errorf(
				"%d. IsUpper(\"%s\") Error=%v, Valid=%v. Error value: %s\n",
				i, d.Data, err != nil, d.Valid, err)
		}
	}
}

func Benchmark_IsUpper(b *testing.B) {
	b.StopTimer()
	v := validate.NewValidator()
	for _, d := range isUpperTests {
		data := NewUpper([]byte(d.Data))
		b.StartTimer()
		for i := 0; i < b.N; i++ {
			v.Validate(data)
		}
	}
}
