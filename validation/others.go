package validation

import (
	"encoding/json"
	"regexp"
)

// IsJSON determines if a string is valid JSON
func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsRegexMatch performs a regex match on a str
func IsRegexMatch(str string, regex string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}

// IsBase64 determines if a string is a valid base64
func IsBase64(str string) bool {
	b := regexp.MustCompile(`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`)
	return b.MatchString(str)
}
