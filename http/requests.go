package http

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Request holds our http client and corresponding request data
type Request struct {
	Client       http.Client       `json:"-"`
	Request      *Request          `json:"-"`
	URL          string            `json:"url,omitempty"`
	Headers      http.Header       `json:"-"`
	PostBody     map[string]string `json:"post_body,omitempty"`
	PostBodyJSON []byte            `json:"post_body_json,omitempty"`
	ContentType  string            `json:"content_type,omitempty"`
}

// String converts our Request struct into a JSON string
func (r *Request) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// Response holds our response object, as well as a pointer to the original request
type Response struct {
	Request     *Request       `json:"-"`
	RawRequest  *http.Request  `json:"-"`
	Body        []byte         `json:"body,omitempty"`
	BodyString  string         `json:"body_string,omitempty"`
	RawResponse *http.Response `json:"-"`
	Error       error          `json:"error,omitempty"`
}

// String converts our Response struct into a JSON string
func (r *Response) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// Get performs a GET request hanging off of our Request pointer and returns a pointer to the Response object
func (r *Request) Get() *Response {
	req, err := http.NewRequest("GET", r.URL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header = r.Headers

	res, err := r.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	response := Response{
		Request:     r,
		RawRequest:  req,
		RawResponse: res,
	}
	response.Body, response.Error = ioutil.ReadAll(res.Body)
	if response.Error != nil {
		log.Fatalln(response.Error)
	}

	// convert the body to type string
	response.BodyString = strings.TrimRight(string(response.Body), "\r\n")
	return &response
}

// Post performs a POST request hanging off of our Request pointer and returns a pointer to the Response object
func (r *Request) Post() *Response {
	response := Response{}

	// Encode the data
	r.PostBodyJSON, _ = json.Marshal(r.PostBody)

	// Leverage Go's HTTP Post function to make request
	response.RawResponse, response.Error = http.Post(r.URL, r.ContentType, bytes.NewBuffer(r.PostBodyJSON))
	if response.Error != nil {
		log.Fatalln(response.Error)
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.RawResponse.Body)

	// Read the response body
	response.Body, response.Error = ioutil.ReadAll(response.RawResponse.Body)
	if response.Error != nil {
		log.Fatalln(response.Error)
	}

	// convert the body to type string
	response.BodyString = strings.TrimRight(string(response.Body), "\r\n")
	return &response
}
