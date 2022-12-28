package validation

import (
	"regexp"
	"strconv"
	"unicode"
)

func ContainLetter(str string) bool {
	return regexp.MustCompile(`[a-zA-Z]`).MatchString(str)
}

func ContainLower(str string) bool {
	for _, r := range str {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func ContainUpper(str string) bool {
	for _, r := range str {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

func IsAlpha(str string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(str)
}

func IsAllUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return str != ""
}

func IsAllLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return str != ""
}

func IsEmptyString(str string) bool {
	return len(str) == 0
}

func IsFloatStr(str string) bool {
	_, e := strconv.ParseFloat(str, 64)
	return e == nil
}

func IsNumberStr(str string) bool {
	return IsIntStr(str) || IsFloatStr(str)
}
