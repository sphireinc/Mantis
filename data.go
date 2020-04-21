package mantis

import (
	"os"
	"strconv"
	"strings"

	"github.com/jwilder/gojq"
)

// IsTrue determines if a string is boolean true/false
func IsTrue(str string) bool {
	converted, err := strconv.ParseBool(strings.ToLower(str))
	if err == nil {
		return converted
	}
	return false
}

// JsonQuery queries a json object for a given path
func JsonQuery(jsonObj string, query string) (interface{}, error) {
	parser, err := gojq.NewStringQuery(jsonObj)
	if err != nil {
		return "", err
	}
	res, err := parser.Query(query)
	if err != nil {
		return "", err
	}
	return res, nil
}

// Exists checks if a given directory path exists
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Contains checks if a map[string]string has a given key
func Contains(item map[string]string, key string) bool {
	if _, ok := item[key]; ok {
		return true
	}
	return false
}
