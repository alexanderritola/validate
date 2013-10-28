package validate

import (
	"unicode/utf8"
	"bytes"
)

var (
//	loAlphabet = []byte("abcdefghijklmnopqrstuvwxyz") // Slower than below
	loAlphabet = []byte("eitsanhurdmwgvlfbkopjxczyq") // UTF-8 lowercase characters in common order
//	upAlphabet = []byte("EITSANHURDMWGVLFBKOPJXCZYQ") // UTF-8 uppercase characters in common order
)

func ValidateLowAlphabet(b []byte) bool {
	if utf8.Valid(b) {
		for _, r  := range b {
			if bytes.IndexByte(loAlphabet, r) == -1 {
				return false
			} 
		}
		return true
	}
	return false
}

func ValidateLowAlphabet_2(b []byte) bool {
	if utf8.Valid(b) {
		for _, r  := range b {
			if bytes.IndexByte(loAlphabet, r) == -1 {
				return false
			} 
		}
		return true
	}
	return false
}
