// Implementation of inbuilt error type
package main

import (
	"errors"
	"fmt"
)

// Inbuilt type -> Implementing the error interface
type errorString struct {
	errorMessage string
}

func (error errorString) Error() string {
	return error.errorMessage
}

func (error errorString) New(message string) *errorString {
	return &errorString{errorMessage: message}
}

func main() {
	var ScrewedUpError = errors.New("This screwed up!")
	var BombError = errors.New("The Program bombed!")

	// Calling the Error method of the error types
	fmt.Println("ScrewedUpError", ScrewedUpError.Error())
	fmt.Println("BombError", BombError.Error())

	// fmt package formats the error by calling Error method
	fmt.Println("ScrewedUpError", ScrewedUpError)
	fmt.Println("BombError", BombError)

	fmt.Println(fmt.Errorf("This is a formated error (%q)", "Rajeev"))
}
