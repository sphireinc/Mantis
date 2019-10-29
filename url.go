package mantis

import (
	"github.com/gorilla/mux" // TODO: Still relies on GMux for vars - need to find alternative
	"io/ioutil"
	"net/http"
)

// GetBody returns the body from the http request
func GetBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	HandleError("Error reading body: %v", err)
	return body
}

// GetQueryParameter fetches a URL query parameter based on a key and return a string array
func GetQueryParameter(r *http.Request, key string) []string {
	value, ok := r.URL.Query()[key]
	if !ok || len(value) < 1 {
		return nil
	}
	return value
}

// GetQueryParameterFirst returns the first element of the requested URL query parameter (ie /some/url?key=value)
func GetQueryParameterFirst(r *http.Request, key string) string {
	return GetQueryParameter(r, key)[0]
}

// GetUrlParameter returns a dynamic url based parameter based on the key (ie /some/url/:parameter - key = 'parameter')
func GetUrlParameter(r *http.Request, key string) string {
	value := r.URL.Query().Get(key)
	if value == "" {
		vars := mux.Vars(r)
		value = vars[key]
	}
	return value
}

// GetUrlParameters returns all url based parameters
func GetUrlParameters(r *http.Request) map[string]string {
	return mux.Vars(r)
}
