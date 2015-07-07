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
	if r.url == "" {
		return "", requesterror.InvalidRequest()
	} else if r.malformed() {
		return "", requesterror.MalformedRequest()
	} else {
		return "Request Successful", nil
	}
}

func main() {
	// Request 1
	request := Request{url: ""}
	_, err := request.Submit()
	if err != nil {
		fmt.Println("1st Request Error:", err)
	}

	// Request 2
	request2 := Request{url: "i/am/malformed"}
	_, err2 := request2.Submit()
	if err != nil {
		fmt.Println("2nd Request Error:", err2)
	}
}
