package validation

import "regexp"

// ContainsChinese determines if a string is Chinese
func ContainsChinese(str string) bool {
	return regexp.MustCompile("[\u4e00-\u9fa5]").MatchString(str)
}
