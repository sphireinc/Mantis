package mantis

import (
	"errors"
	"net/http"
	"github.com/gorilla/mux"
)

// GetQueryParameter Fetch a URL query parameter based on a key and return a string array
func GetQueryParameter(r *http.Request, key string) []string {
	value, ok := r.URL.Query()[key]
	if !ok {
		HandleError("GetQueryParameter ", errors.New("GetQueryParameter Failed to find key: "+key))
	}
	return value
}

// GetQueryParameterFirst Returns the first element of the requested URL query parameter (ie /some/url?key=value)
func GetQueryParameterFirst(r *http.Request, key string) string {
	return GetQueryParameter(r, key)[0]
}

// GetUrlParameter Returns a dynamic url based parameter based on the key (ie /some/url/:parameter - key = 'parameter')
func GetUrlParameter(r *http.Request, key string) string {
	value := r.URL.Query().Get(":" + key)
	if value == "" {
		vars := mux.Vars(r)
		value = vars[key]
	}
	return value
}
