package validate

import (
	"unicode/utf8"
)

var (
//	loAlphabet = []byte("abcdefghijklmnopqrstuvwxyz") // Slower than below
	loAlphabet = []byte("eitsanhurdmwgvlfbkopjxczyq") // UTF-8 lowercase characters in common order
//	upAlphabet = []byte("EITSANHURDMWGVLFBKOPJXCZYQ") // UTF-8 uppercase characters in common order
)

func ValidateLowAlphabet(b []byte) bool {
	match := 0
	if utf8.Valid(b) {
		for i, _  := range b {
			for ii, _ := range loAlphabet {
				if b[i] == loAlphabet[ii] {
					match += 1
					break
				}
			}
			if match != (i+1) {
				break
			} else {
				continue
			}
		}
	}
	if len(b) == match {
		return true
	}
	return false
}
/*
func ValidateLowAlphabet_2(b []byte) bool {
	if utf8.Valid(b) {
		for i, r  := range s {
			for ii, _ := range loAlphabet {
				if b[i] == loAlphabet[ii] {
					match += 1
					break
				}
			}
			if match != (i+1) {
				break
			} else {
				continue
			}
		}
	}
	if len(b) == match {
		return true
	}
	return false
}
*/
