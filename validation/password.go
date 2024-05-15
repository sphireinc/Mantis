package validation

import "unicode"

// IsStrongPassword determines the complexity of a string
// Every character in the string is evaluated. For every character that:
// IsDigit add 3 complexity
// IsUpper add 2 complexity
// IsLower add 1 complexity
// IsSymbol OR IsPunctuation add 5 complexity
// A complexity of less than 15 (arbitrary) is deemed a weak password. Above is strong.
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
			complexity++
		case unicode.IsSymbol(r), unicode.IsPunct(r):
			complexity += 5
		}
	}

	if complexity <= 15 {
		return false, complexity
	}

	return true, complexity
}

// IsWeakPassword determines if a password is weak
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
