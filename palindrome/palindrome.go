package palindrome

func Palindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}
