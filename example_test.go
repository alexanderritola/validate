package validate_test

import (
	"errors"
	"fmt"
	"github.com/alexanderritola/validate"
)

// This will be a method for verifying that we were given the string 'cat'.
type IsCat struct {
	data    string
	maxlen  int
	message string
}

// Check that the data we were given was something a cat would say (meow).
func (c *IsCat) Validate(v validate.Validator) validate.Error {
	// Check to see if the length is longer than allowed.
	if len(c.data) > c.maxlen {
		return &validate.Error{
			Message: errors.New("Not a cat, message too long!"),
		}
	}

	// Check to see if the message came from a cat.
	if c.data != "meow" {
		return &validate.Error{
			Message: errors.New("Not a cat, they don't say that!"),
		}
	} else {
		// Everything checks out, we have a cat!
		return nil
	}
}

// Set the failure message.
func (c *IsCat) SetMessage(msg string) validate.Method {
	c.message = msg
	return c
}

func (c *IsCat) MaxLength(length int) *IsCat {
	c.maxlen = length
	return c
}

// Get the failure message.
func (c *IsCat) Message() string {
	return c.message
}

// Create a new IsCat validation method.
func NewIsCat(data string) *IsCat {
	return &IsCat{data: data}
}

// Show how to define our own custom validation methods.
func ExampleMethod() {
	// Get a new BasicValidator
	v := validate.NewValidator()

	// We'll loop though these, validating each one.
	var cats = []string{"woof", "meow", "quack"}
	for i, c := range cats {
		// Create our validation method, including a custom MaxLength
		// requirement. Note that SetMessage() must be last in the chain.
		// We are going to assume that cats only say thing that are 4 characters
		// long
		IsCat := NewIsCat(c).MaxLength(4).SetMessage("Not a cat!")

		// Validate
		err := v.Validate(IsCat)
		if err != nil {
			fmt.Printf("%v. '%s' %s\n", i, c, IsCat.Message())
			fmt.Printf("    Returned error: %v\n", err)
		} else {
			fmt.Printf("%v. '%s' is a cat!\n", i, c)
		}
	}

	// Output:
	// 0. 'woof' Not a cat!
	//     Returned error: Error level 0: Not a cat, they don't say that!
	// 1. 'meow' is a cat!
	// 2. 'quack' Not a cat!
	//     Returned error: Error level 0: Not a cat, message too long!

}
