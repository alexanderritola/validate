package validate

import (
	"errors"
	"fmt"
	"unicode"
	"unicode/utf8"
)

// Validator validates Methods
type Validator struct {
	NonLatin bool // Use only Latin characters by default
}

// Validate the given method
func (v *Validator) Validate(m *Method) (e *Error) {
	return
}

// Defines a specific validation method
type Method interface {
	// Run the validation check on the method
	Validate(*Validator) *Error

	// Set the message to be returned on validation failure
	Message(string) *Method
}

// Validation error
type Error struct {
	Level   ErrorLevel
	Message error
}

// I think I need to improve this type but it works for now
func (e *Error) Error() string {
	return fmt.Sprintf("%s", e.Message)
}

// Validation error level.
type ErrorLevel int

const (
	// Data did not pass validation formatting requirements
	ErrInvalid = iota + 1

	// Data contained invalid characters
	ErrInvalidChars
)

var (
	// Invalid UTF8 characters were encountered
	ErrInvalidUTF8 = &Error{
		Level:   ErrInvalidChars,
		Message: Critical,
	}
)

var (
	Format = errors.New(
		"validate: Data did not match the formatting requirements")
	Critical = errors.New(
		"validate: Data contained control or non-printable characters")
)

var (
	//	loAlphabet = []byte("abcdefghijklmnopqrstuvwxyz") // Slower than below

	// UTF-8 lowercase characters in common order
	loAlphabet = []byte("eitsanhurdmwgvlfbkopjxczyq")

	// UTF-8 uppercase characters in common order
	//	upAlphabet = []byte("EITSANHURDMWGVLFBKOPJXCZYQ")
)

// Check to ensure the byte slice only contains printable UTF-8 runes
func IsPrint(p []byte) bool {
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
