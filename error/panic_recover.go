package main

import "fmt"

func main() {
	// Recover a panicking method
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OMG! I got recovered")
		}
	}()
	PanicDemonstration()
}

// Panic executed deferred functions normally
func PanicDemonstration() {
	defer fmt.Println("I am defered but executed on panic")
	panic("This thing just blowed up")
}
