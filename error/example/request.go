package main

import (
	"./requesterror"
	"fmt"
)

type Request struct {
	url string
}

func (r Request) malformed() bool {
	if r.url == "i/am/malformed" {
		return true
	}
	return false
}

func (r Request) Submit() (string, error) {
	customError := new(requesterror.CustomUserError)
	if r.url == "" {
		return "", customError.InvalidRequest()
	} else if r.malformed() {
		return "", customError.MalformedRequest()
	} else {
		return "Request Successfull", nil
	}
}

func main() {
	request := Request{url: ""}
	_, err := request.Submit()
	if err != nil {
		fmt.Println("1st request Successful")
		return
	}
	fmt.Println("1st Request", err)

	request2 := Request{url: "i/am/malformed"}
	_, err2 := request2.Submit()
	if err != nil {
		fmt.Println("2nd request Successful")
		return
	}
	fmt.Println("2nd Request", err2)
}
