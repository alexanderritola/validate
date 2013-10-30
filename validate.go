// Package validate provides core functions for validating data using modular
// validation methods from outside packages with support for failed validation
// severity levels.
package validate

import (
	"errors"
	"fmt"
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

	Format = errors.New(
		"Data did not meet the formatting requirements")

	Critical = errors.New(
		"Data contained control or non-printable characters")
)
