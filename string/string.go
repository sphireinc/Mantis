package string

func reverseString(s string) string {
	runes := []rune(s)
	reversed := make([]rune, len(s))

	for i, r := range runes {
		reversed[len(s)-1-i] = r
	}

	return string(reversed)
}

func isPalindrome(s string) bool {
	return s == reverseString(s)
}
