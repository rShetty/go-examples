// Demonstration of custom error types

package main

import "fmt"

// Custom type implementing error interface
type MyScrewedUpError struct {
	errorMessage string
}

// MyScrewedUpError implements the builtin error interface
func (error MyScrewedUpError) Error() string {
	return "My Awesome custom error" + error.errorMessage
}

func main() {
	myScrewedUpError := new(MyScrewedUpError)
	fmt.Println("ScrewedUpError", myScrewedUpError.Error())
}
