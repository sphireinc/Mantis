package mantis

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// ParseBodyIntoStruct takes the body from an HTTP request and parses it into a JSON friendly struct
//     var s someStruct
//     mantis.ParseBodyIntoStruct(r, &s)
func ParseBodyIntoStruct(r *http.Request, obj interface{}) error {
	return json.NewDecoder(r.Body).Decode(obj)
}

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

// ParseUrl returns a *url.URL from a given URL string
func ParseUrl(rawurl string) *url.URL {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("unable to parse url %s: %s", rawurl, err)
	}
	return u
}