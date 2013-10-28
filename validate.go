package validate

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

var (
	//	loAlphabet = []byte("abcdefghijklmnopqrstuvwxyz") // Slower than below

	// UTF-8 lowercase characters in common order
	loAlphabet = []byte("eitsanhurdmwgvlfbkopjxczyq")

	// UTF-8 uppercase characters in common order
	//	upAlphabet = []byte("EITSANHURDMWGVLFBKOPJXCZYQ")
)

func ValidateLowAlphabet(b []byte) bool {
	if utf8.Valid(b) {
		for _, r := range b {
			if bytes.IndexByte(loAlphabet, r) == -1 {
				return false
			}
		}
		return true
	}
	return false
}

func ValidatePrintableRunes(p []byte) bool {
	// Borrowed from utf.Valid() with added checks for printable runes
	for i := 0; i < len(p); {
		if p[i] < utf8.RuneSelf {
			// Check if this single byte run is printable
			if !unicode.IsPrint(rune(p[i])) {
				return false
			}
			i++
		} else {
			r, size := utf8.DecodeRune(p[i:])
			if size == 1 {
				// All valid runes of size 1 (those
				// below RuneSelf) were handled above.
				// This must be a RuneError.
				return false
			}
			// Check if this multi-byte rune is printable
			if !unicode.IsPrint(r) {
				return false
			}
			i += size
		}
	}
	return true
}
