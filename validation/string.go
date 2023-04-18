package validation

import (
	"regexp"
	"strconv"
	"unicode"
)

// ContainLetter determines if a string contains letters
func ContainLetter(str string) bool {
	return regexp.MustCompile(`[a-zA-Z]`).MatchString(str)
}

// ContainLower determines if a string contains lowercase characters
func ContainLower(str string) bool {
	for _, r := range str {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// ContainUpper determines if a string contains uppercase characters
func ContainUpper(str string) bool {
	for _, r := range str {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// IsAlpha determines if a string is all alpha characters
func IsAlpha(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(str)
}

// IsAllUpper determines if a string is all uppercase
func IsAllUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return str != ""
}

// IsAllLower determines if a string is all lowercase
func IsAllLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return str != ""
}

// IsEmptyString determine if a string is an empty string (yes, seriously, lol)
func IsEmptyString(str string) bool {
	return len(str) == 0
}

// IsFloatStr determine if a string is a float
func IsFloatStr(str string) bool {
	_, e := strconv.ParseFloat(str, 64)
	return e == nil
}

// IsNumberStr determines if a string is either an int or a float
func IsNumberStr(str string) bool {
	return IsIntStr(str) || IsFloatStr(str)
}
