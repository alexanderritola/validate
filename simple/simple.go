package simple

import (
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

type Printable struct {
	data    []byte
	message string
}

// Check to ensure the byte slice only contains printable UTF-8 runes
func (m *Printable) Validate(v *validate.Validator) *validate.Error {
	// Borrowed from utf.Valid() with added checks for printable runes
	for i := 0; i < len(p); {
		if p[i] < utf8.RuneSelf {
			// Check if this single byte run is printable
			if !unicode.IsPrint(rune(p[i])) {
				return nil
			}
			i++
		} else {
			r, size := utf8.DecodeRune(p[i:])
			if size == 1 {
				// All valid runes of size 1 (those
				// below RuneSelf) were handled above.
				// This must be a RuneError.
				return nil
			}
			// Check if this multi-byte rune is printable
			if !unicode.IsPrint(r) {
				return nil
			}
			i += size
		}
	}
	return nil
}

// Check to ensure the byte slice only contains lower case UTF-8 runes
func IsLower(p []byte) bool {
	// Borrowed from utf.Valid() with added checks for lower case runes
	for i := 0; i < len(p); {
		if p[i] < utf8.RuneSelf {
			// Check if this single byte run is a lower case rune
			if !unicode.IsLower(rune(p[i])) {
				return false
			}
			i++
		}
	}
	return true
}

// Check to ensure the byte slice only contains upper case UTF-8 runes
func IsUpper(p []byte) bool {
	// Borrowed from utf.Valid() with added checks for upper case runes
	for i := 0; i < len(p); {
		if p[i] < utf8.RuneSelf {
			// Check if this single byte run is a upper case rune
			if !unicode.IsUpper(rune(p[i])) {
				return false
			}
			i++
		}
	}
	return true
}
