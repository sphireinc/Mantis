package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ResponseJSONError is our response JSON error struct
type ResponseJSONError struct {
	Error string `json:"error,omitempty"`
}

func (r *ResponseJSONError) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// ResponseJSONOk is our response JSON struct
type ResponseJSONOk struct {
	Data string `json:"data,omitempty"`
}

func (r *ResponseJSONOk) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		jsonError := ResponseJSONError{
			Error: err.Error(),
		}
		return jsonError.String()
	}
	return string(marshaledStruct)
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
	text := http.StatusText(code)
	if text == "" {
		code = 520
		text = "Unknown"
	}

	return ResponseCodes{
		code:        int16(code),
		description: text,
	}
}
