package http

import (
	"encoding/json"
	"fmt"
)

// ResponseJSONError is our response JSON error struct
type ResponseJSONError struct {
	Error string `json:"error,omitempty"`
}

// Byte converts our ResponseJSONError struct into a JSON []byte
func (r *ResponseJSONError) Byte() []byte {
	marshaledStruct, _ := json.Marshal(r)
	return marshaledStruct
}

// ResponseJSONOk is our response JSON struct
type ResponseJSONOk struct {
	Data string `json:"data,omitempty"`
}

// Byte converts our ResponseJSONOk struct into a JSON []byte
func (r *ResponseJSONOk) Byte() []byte {
	marshaledStruct, _ := json.Marshal(r)
	return marshaledStruct
}

// ResponseCodes holds our response codes and their description
type ResponseCodes struct {
	code        int16
	description string
}

func (r *ResponseCodes) String() string {
	return fmt.Sprintf("%d %s", r.code, r.description)
}

// GetHTTPResponseCode returns the description of a numeric HTTP code
func GetHTTPResponseCode(code int) ResponseCodes {
	text := StatusText(code)
	if text == "" {
		code = 520
		text = "Unknown"
	}

	return ResponseCodes{
		code:        int16(code),
		description: text,
	}
}
