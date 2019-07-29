package mantis

import (
	"github.com/gorilla/mux" // TODO: Still relies on GMux for vars - need to find alternative
	"net/http"
)

// GetQueryParameter Fetch a URL query parameter based on a key and return a string array
func GetQueryParameter(r *http.Request, key string) []string {
	value, ok := r.URL.Query()[key]
	if !ok || len(value) < 1 {
		return nil
	}
	return value
}

// GetQueryParameterFirst Returns the first element of the requested URL query parameter (ie /some/url?key=value)
func GetQueryParameterFirst(r *http.Request, key string) string {
	return GetQueryParameter(r, key)[0]
}

// GetUrlParameter Returns a dynamic url based parameter based on the key (ie /some/url/:parameter - key = 'parameter')
func GetUrlParameter(r *http.Request, key string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		vars := mux.Vars(r)
		value = vars[key]
	}
	return value
}

// GetUrlParameters Returns all url based parameters
func GetUrlParameters(r *http.Request) map[string]string {
	return mux.Vars(r)
}
