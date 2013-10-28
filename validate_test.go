package validate

import (
	"testing"
)

func Test_ValidateLowAlphabet_1(t *testing.T) {
	if (ValidateLowAlphabet("goiscool") != true) {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_ValidateLowAlphabet_2(t *testing.T) {
	if (ValidateLowAlphabet("goiscoolandshit") != true) {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}
