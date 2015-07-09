package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("somefile.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error")
	}
}
