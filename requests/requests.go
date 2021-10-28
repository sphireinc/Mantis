package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	URL          string            `json:"url"`
	Headers      map[string]string `json:"headers"`
	PostBody     map[string]string `json:"post_body"`
	PostBodyJSON []byte            `json:"post_body_json"`
	ContentType  string            `json:"content_type"`
}

type Response struct {
	Request     *Request       `json:"request"`
	Body        []byte         `json:"body"`
	BodyString  string         `json:"body_string"`
	RawResponse *http.Response `json:"raw_response"`
	Error       error          `json:"error"`
}

func (R *Request) Get() *Response {
	response := Response{}

	response.RawResponse, response.Error = http.Get(R.URL)
	if response.Error != nil {
		log.Fatalln(response.Error)
	}

	// read the response body on the line below.
	response.Body, response.Error = ioutil.ReadAll(response.RawResponse.Body)
	if response.Error != nil {
		log.Fatalln(response.Error)
	}

	// convert the body to type string
	response.BodyString = string(response.Body)
	return &response
}

func (R *Request) Post() *Response {
	response := Response{}

	// Encode the data
	R.PostBodyJSON, _ = json.Marshal(R.PostBody)

	// Leverage Go's HTTP Post function to make request
	response.RawResponse, response.Error = http.Post(R.URL, R.ContentType, bytes.NewBuffer(R.PostBodyJSON))
	if response.Error != nil {
		log.Fatalln(response.Error)
	}
	defer response.RawResponse.Body.Close()

	// Read the response body
	response.Body, response.Error = ioutil.ReadAll(response.RawResponse.Body)
	if response.Error != nil {
		log.Fatalln(response.Error)
	}

	// convert the body to type string
	response.BodyString = string(response.Body)
	return &response
}
