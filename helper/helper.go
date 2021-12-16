package helper

import (
	"os"
	"strconv"
)

// Reverse a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// DeferFileClose prevents non-closure file closing error
func DeferFileClose(file *os.File) {
	err := file.Close()
	if err != nil {

	}
}

// StringToBool same as strconv.ParseBool except hides the error (returns false)
func StringToBool(boolean string) bool {
	ret, err := strconv.ParseBool(boolean)
	if err != nil {
		return false
	}
	return ret
}

// AtoiWithDefault same as strconv.Atoi except only returns the value or a default value if nil
func AtoiWithDefault(value string, defaultValue int) int {
	intFromStr, intFromStrErr := strconv.Atoi(value)
	if intFromStrErr != nil {
		return defaultValue
	}
	return intFromStr
}

// Default returns the defaultVal given if originalVal is empty/nil
func Default[T comparable](originalVal T, defaultVal T) T {
	var zero T
	if originalVal == zero {
		return defaultVal
	}
	return originalVal
}
