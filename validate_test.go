package validate

import (
	"testing"
)

func Test_ValidateLowAlphabet_1(t *testing.T) {
	s := "goiscool"
	if (ValidateLowAlphabet([]byte(s)) != true) {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidateLowAlphabet_2(t *testing.T) {
	s := "goiscoolandshit"
	if (ValidateLowAlphabet([]byte(s)) != true) {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidateLowAlphabet_2_1(t *testing.T) {
	s := "goiscool"
	if (ValidateLowAlphabet_2([]byte(s)) != true) {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidateLowAlphabet_2_2(t *testing.T) {
	s := "goiscoolandshit"
	if (ValidateLowAlphabet_2([]byte(s)) != true) {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}
