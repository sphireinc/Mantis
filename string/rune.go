package string

func countRune(s string, targetRune rune) int {
	count := 0
	for _, r := range s {
		if r == targetRune {
			count++
		}
	}
	return count
}
