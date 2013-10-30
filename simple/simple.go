// Package validate/simple provides common methods for validating data.
package simple

import (
	"errors"
	"github.com/daswasser/validate"
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

var (
	// Invalid characters or punctuation.
	ErrFormatting = &validate.Error{
		Level:   validate.ErrInvalid,
		Message: errors.New("Invalid formatting"),
	}
)

type Printable struct {
	data    []byte
	message string
}

func NewPrintable(data []byte) *Printable {
	return &Printable{data: data}
}

// Check to ensure the byte slice only contains printable UTF-8 runes
func (m *Printable) Validate(v validate.Validator) *validate.Error {
	p := m.data
	// Borrowed from utf.Valid() with added checks for printable runes
	for i := 0; i < len(p); {
		if p[i] < utf8.RuneSelf {
			// Check if this single byte run is printable
			if !unicode.IsPrint(rune(p[i])) {
				return ErrFormatting
			}
			i++
		} else {
			r, size := utf8.DecodeRune(p[i:])
			if size == 1 {
				// All valid runes of size 1 (those
				// below RuneSelf) were handled above.
				// This must be a RuneError.
				return validate.ErrInvalidUTF8
			}
			// Check if this multi-byte rune is printable
			if !unicode.IsPrint(r) {
				return ErrFormatting
			}
			i += size
		}
	}
	return nil
}

// Defines a check for all lowercase unicode points. This will fail if there is
// any punctuation or spaces.
// TODO(inhies): Add option for allowing a space or similar characters to not
// fail validation.
type Lower struct {
	data    []byte
	message string
}

// Returns a new Lower value for the specified data.
func NewLower(data []byte) *Lower {
	return &Lower{data: data}
}

// Check to ensure the byte slice only contains lower case UTF-8 runes
func (m *Lower) Validate(v validate.Validator) *validate.Error {
	p := m.data
	// Borrowed from utf.Valid() with added checks for lower case runes
	for i := 0; i < len(p); {
		if p[i] < utf8.RuneSelf {
			// Check if this single byte run is a lower case rune
			if !unicode.IsLower(rune(p[i])) {
				return ErrFormatting
			}
			i++
		} else {
			r, size := utf8.DecodeRune(p[i:])
			if size == 1 {
				// All valid runes of size 1 (those
				// below RuneSelf) were handled above.
				// This must be a RuneError.
				return validate.ErrInvalidUTF8
			}

			// Check if this multi-byte rune is printable
			if !unicode.IsLower(r) {
				return ErrFormatting
			}
			i += size
		}
	}
	return nil
}

// Defines a check for all uppercase  unicode points. This will fail if there is
// any punctuation or spaces.
// TODO(inhies): Add option for allowing a space or similar characters to not
// fail validation.
type Upper struct {
	data    []byte
	message string
}

// Returns a new Upper  value for the specified data.
func NewUpper(data []byte) *Upper {
	return &Upper{data: data}
}

// Check to ensure the byte slice only contains lower case UTF-8 runes
func (m *Upper) Validate(v validate.Validator) *validate.Error {
	p := m.data
	// Borrowed from utf.Valid() with added checks for lower case runes
	for i := 0; i < len(p); {
		if p[i] < utf8.RuneSelf {
			// Check if this single byte run is a lower case rune
			if !unicode.IsUpper(rune(p[i])) {
				return ErrFormatting
			}
			i++
		} else {
			r, size := utf8.DecodeRune(p[i:])
			if size == 1 {
				// All valid runes of size 1 (those
				// below RuneSelf) were handled above.
				// This must be a RuneError.
				return validate.ErrInvalidUTF8
			}

			// Check if this multi-byte rune is printable
			if !unicode.IsUpper(r) {
				return ErrFormatting
			}
			i += size
		}
	}
	return nil
}
