package mantis

import (
	"errors"
	"net/http"
)

// GetQueryParameter Fetch a URL query parameter based on a key and return a string array
func GetQueryParameter(r *http.Request, key string) []string {
	value, ok := r.URL.Query()[key]
	if !ok {
		HandleError("GetQueryParameter ", errors.New("GetQueryParameter Failed to find key: "+key))
	}
	return value
}

// GetQueryParameterFirst Returns the first element of the requested URL query parameter
func GetQueryParameterFirst(r *http.Request, key string) string {
	return GetQueryParameter(r, key)[0]
}
