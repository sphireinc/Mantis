package http

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// ParseBodyIntoStruct takes the body from an HTTP request and parses it into a JSON friendly struct
func ParseBodyIntoStruct(r *http.Request, obj interface{}) error {
	return json.NewDecoder(r.Body).Decode(obj)
}

// GetBody returns the body from the http request
func GetBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

// GetQueryParameter fetches a URL query parameter based on a key and return a string array
func GetQueryParameter(r *http.Request, key string) (string, error) {
	value, ok := r.URL.Query()[key]
	if !ok || len(value) < 1 {
		return "", errors.New("key not found")
	}
	return value[0], nil
}

// GetQueryParameters returns all query parameters
func GetQueryParameters(r *http.Request) url.Values {
	return r.URL.Query()
}

// ParseUrl returns a *http.URL from a given URL string
func ParseUrl(rawurl string) (*url.URL, error) {
	return url.Parse(rawurl)
}
