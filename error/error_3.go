//Comparing error values

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var (
	FooError    = errors.New("I am foo")
	BarError    = errors.New("I am bar")
	FooBarError = errors.New("I am foobar")
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	err := ComparingErrorValues()
	if err != nil {
		switch err {
		case FooError:
			fmt.Println("Switched to foo")
			return
		case BarError:
			fmt.Println("Switched to bar error")
			return
		case FooBarError:
			fmt.Println("Switched to foobar error")
			return
		default:
			fmt.Println("I am crazy to write this")
			return
		}
	}
}

func ComparingErrorValues() error {
	errArray := [3]error{FooBarError, FooError, BarError}
	return errArray[rand.Intn(len(errArray))]
}
