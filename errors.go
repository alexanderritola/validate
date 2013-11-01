package validate

import (
	"errors"
	"fmt"
)

// Error should be able to return a valid error message and a validation
// failure severity level.
type Error interface {
	Error() string
	Level() int
}

// Validation error
type ValidatorError struct {
	ErrLevel int   // Error severity level
	Message  error // Error message
}

// Returns the string representation of the error level and message
func (e *ValidatorError) Error() string {
	return fmt.Sprintf("Error level %v: %s", e.ErrLevel, e.Message)
}

// Returns the error level.
func (e *ValidatorError) Level() int {
	return e.ErrLevel
}

// Validation error level. Higher levels are more severe.
type ErrorLevel int

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
	ErrInvalidUTF8 = &ValidatorError{ErrInvalidChars, Critical}

	Format = errors.New(
		"Data did not meet the formatting requirements")

	Critical = errors.New(
		"Data contained control or non-printable characters")
)
