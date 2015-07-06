package main

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
