// Package validate provides core functions for validating data using modular
// validation methods from outside packages with support for failed validation
// severity levels.
package validate

func Test() {
}

// The Validator interface is used to execute Methods to validate data.
type Validator interface {
	Validate(Method) *Error // Validate the method and return the error, if any.
}

// A simple Validator.
type BasicValidator struct{}

// Returns a new Validator value.
func NewValidator() *BasicValidator {
	return &BasicValidator{}
}

// Validate the given method.
func (v *BasicValidator) Validate(m Method) (e *Error) {
	// Run validation. If we get an error, append it to the error list.
	return m.Validate(v)
}

// Defines a specific validation method.
type Method interface {
	Validate(Validator) *Error // Run the validation check on the method.
	SetMessage(string) Method  // Sets the failure message.
	Message() string           // Gets the failure message.
}
