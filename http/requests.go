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

type Request struct {
	URL          string            `json:"url,omitempty"`
	Headers      map[string]string `json:"headers,omitempty"`
	PostBody     map[string]string `json:"post_body,omitempty"`
	PostBodyJSON []byte            `json:"post_body_json,omitempty"`
	ContentType  string            `json:"content_type,omitempty"`
}

func (r *Request) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

type Response struct {
	Request     *Request       `json:"request,omitempty"`
	Body        []byte         `json:"body,omitempty"`
	BodyString  string         `json:"body_string,omitempty"`
	RawResponse *http.Response `json:"raw_response,omitempty"`
	Error       error          `json:"error,omitempty"`
}

func (r *Response) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

func (r *Request) Get() *Response {
	response := Response{}

	response.RawResponse, response.Error = http.Get(r.URL)
	if response.Error != nil {
		log.Fatalln(response.Error)
	}

	// read the response body on the line below.
	response.Body, response.Error = ioutil.ReadAll(response.RawResponse.Body)
	if response.Error != nil {
		log.Fatalln(response.Error)
	}

	// convert the body to type string
	response.BodyString = strings.TrimRight(string(response.Body), "\r\n")
	return &response
}

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
		err := Body.Close()
		if err != nil {

		}
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
