package main

import "testing"

func TestReverse(test *testing.T) {
	testCases := []struct {
		inputString, outputString string
	}{
		{"rajeev", "veejar"},
		{"manu", "unam"},
		{"monkey", "yeknom"},
		{" ", " "},
	}

	for _, c := range testCases {
		result := Reverse(c.inputString)
		if result != c.outputString {
			test.Errorf("Expected String (%q) but actually got (%q)", c.outputString, result)
		}
	}

}
