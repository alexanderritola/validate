package validate

import (
	"testing"
)

func Test_ValidateLowAlphabet_1(t *testing.T) {
	s := "goiscool"
	if ValidateLowAlphabet([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidateLowAlphabet_2(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	if ValidateLowAlphabet([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidatePrintableRunes_1(t *testing.T) {
	s := "goiscool"
	if ValidatePrintableRunes([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidatePrintableRunes_2(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	if ValidatePrintableRunes([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidateUpAlphabet_1(t *testing.T) {
	s := "GOISCOOL"
	if ValidateUpAlphabet([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidateUpAlphabet_2(t *testing.T) {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if ValidateUpAlphabet([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}
