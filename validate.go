// Package validate provides core functions for validating data using modular
// validation methods from outside packages with support for failed validation
// severity levels.
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
func (v *Validator) Validate(m Method) (e *Error) {
	return m.Validate(v)
}

// Defines a specific validation method
type Method interface {
	Validate(*Validator) *Error // Run the validation check on the method
}

// Validation error
type Error struct {
	Level   ErrorLevel // Error severity level
	Message error      // Error message
}

// Returns the string representation of the error level and message
func (e *Error) Error() string {
	return fmt.Sprintf("Error level %v: %s", e.Level, e.Message)
}

// Validation error level. Higher levels are more severe.
type ErrorLevel int

const (
	// Data did not pass validation formatting requirements
	ErrInvalid = iota + 1

	// Data contained invalid characters
	ErrInvalidChars
)

var (
	// Invalid UTF8 characters were encountered. This could indicate a malicious
	// access attempt.
	ErrInvalidUTF8 = &Error{
		Level:   ErrInvalidChars,
		Message: Critical,
	}
)

var (
	Format = errors.New(
		"Data did not meet the formatting requirements")
	Critical = errors.New(
		"Data contained control or non-printable characters")
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
