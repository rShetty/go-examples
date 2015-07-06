package palindrome

import "testing"

func TestPalindrome(test *testing.T) {
	testCases := []struct {
		inputString string
		result      bool
	}{
		{"malayalam", true},
		{"madam", true},
		{"manu", false},
	}

	for _, c := range testCases {
		output := Palindrome(c.inputString)
		if output != c.result {
			test.Errorf("For (%q), Expected (%b) got (%b)", c.inputString, c.result, output)
		}
	}

}
