package data

import (
	"github.com/jwilder/gojq"
	"os"
	"strconv"
	"strings"
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

// DirectoryExists checks if a given directory path exists
func DirectoryExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// MapStringStringContains contains checks if a map[string]string has a given key
func MapStringStringContains(item map[string]string, key string) bool {
	if _, ok := item[key]; ok {
		return true
	}
	return false
}
