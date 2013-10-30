package validate

import (
	"errors"
	"fmt"
)

// Validation error
type Error struct {
	Level   ErrorLevel // Error severity level
	Message error      // Error message
}

// Validation error level. Higher levels are more severe.
type ErrorLevel int

// Returns the string representation of the error level and message
func (e *Error) Error() string {
	return fmt.Sprintf("Error level %v: %s", e.Level, e.Message)
}

// Error constants
const (
	// Data did not pass validation formatting requirements
	ErrInvalid = iota + 1

	// Data contained invalid characters
	ErrInvalidChars
)

// Pre-defined errors
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
