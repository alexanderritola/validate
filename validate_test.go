package validate

import (
	"testing"
)

func ExampleValidator_Validate() {

}
func Test_IsLower_1(t *testing.T) {
	s := "goiscool"
	if IsLower([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_IsLower_2(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	if IsLower([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_IsPrint_1(t *testing.T) {
	s := "goiscool"
	if IsPrint([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_IsPrint_2(t *testing.T) {
	s := "abcdefghijklmnopqrstuvwxyz"
	if IsPrint([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_IsUpper_1(t *testing.T) {
	s := "GOISCOOL"
	if IsUpper([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}

func Test_IsUpper_2(t *testing.T) {
	s := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if IsUpper([]byte(s)) != true {
		t.Error("Failed")
	} else {
		t.Log("Passed")
	}
}
