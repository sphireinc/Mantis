package url

import (
	"encoding/json"
	mantisError "github.com/sphireinc/mantis/error"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// ParseBodyIntoStruct takes the body from an HTTP request and parses it into a JSON friendly struct
// Usage:
//      var s someStruct
//      mantis.ParseBodyIntoStruct(r, &s)
func ParseBodyIntoStruct(r *http.Request, obj interface{}) error {
	return json.NewDecoder(r.Body).Decode(obj)
}

// GetBody returns the body from the http request
func GetBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		mantisError.HandleError("Error reading body: %v", err)
		return nil, err
	}
	return body, nil
}

// GetQueryParameter fetches a URL query parameter based on a key and return a string array
func GetQueryParameter(r *http.Request, key string) string {
	value, ok := r.URL.Query()[key]
	if !ok || len(value) < 1 {
		return ""
	}
	return value[0]
}

// GetQueryParameters returns all query parameters
func GetQueryParameters(r *http.Request) url.Values {
	return r.URL.Query()
}

// ParseUrl returns a *url.URL from a given URL string
func ParseUrl(rawurl string) (*url.URL, error) {
	u, err := url.Parse(rawurl)
	if err != nil {
		log.Fatalf("unable to parse url %s: %s", rawurl, err)
		return nil, err
	}
	return u, nil
}
