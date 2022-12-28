package validation

import (
	"encoding/json"
	"regexp"
)

func IsJSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

func IsRegexMatch(str string, regex string) bool {
	reg := regexp.MustCompile(regex)
	return reg.MatchString(str)
}

func IsBase64(str string) bool {
	b := regexp.MustCompile(`^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$`)
	return b.MatchString(str)
}
