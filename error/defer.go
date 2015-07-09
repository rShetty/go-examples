package main

import (
	"fmt"
	"os"
)

func main() {
	OpenFile("/Users/rajeevnb/fun/dotfiles/filetype.vim")
	DeferCheck()
	LastInFirstOut()
	fmt.Println("I am the named return value %d", NamedReturnValues())
}

// Defer closing of opened files
func OpenFile(srcName string) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	// Close the opened file
	defer src.Close()
}

// Defered statement arguments are evaluated wheb defer is evaluated
func DeferCheck() {
	i := 0
	defer fmt.Println("Value of i: Argument Evaluation check:  %d", i)
	i++
	return

}

// Defered statements are executed in LIFO fashion
func LastInFirstOut() {
	for i := 0; i < 6; i++ {
		defer fmt.Println("Value of i is %d", i)
	}
}

// Assigning to named return values
// Good for modifying return values
func NamedReturnValues() (awesomeness int) {
	defer func() { awesomeness++ }()
	return 1
}
