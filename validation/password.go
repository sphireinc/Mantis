package validation

import "unicode"

func IsStrongPassword(str string, length int) (bool, int) {
	if len(str) < length {
		return false, 0
	}

	complexity := 0
	for _, r := range str {
		switch {
		case unicode.IsDigit(r):
			complexity += 3
		case unicode.IsUpper(r):
			complexity += 2
		case unicode.IsLower(r):
			complexity += 1
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			complexity += 5
		}
	}

	if complexity <= 15 {
		return false, complexity
	}

	return true, complexity
}

func IsWeakPassword(str string) bool {
	var num, letter, special bool
	for _, r := range str {
		switch {
		case unicode.IsDigit(r):
			num = true
		case unicode.IsLetter(r):
			letter = true
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			special = true
		}
	}

	return (num || letter) && !special
}
