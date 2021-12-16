package http

import (
	"encoding/json"
	"net/http"
)

type ResponseJsonError struct {
	Error string `json:"error,omitempty"`
}

func (r *ResponseJsonError) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

type ResponseJsonOk struct {
	Data string `json:"data,omitempty"`
}

func (r *ResponseJsonOk) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		jsonError := ResponseJsonError{
			Error: err.Error(),
		}
		return jsonError.String()
	}
	return string(marshaledStruct)
}

type ResponseCodes struct {
	code        int16
	description string
}

func (r *ResponseCodes) String() string {
	marshaledStruct, err := json.Marshal(r)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// GetHTTPResponseCode returns the description of a numeric HTTP code
func GetHTTPResponseCode(code int) ResponseCodes {
	text := http.StatusText(code)
	if text == "" {
		text = "Unknown"
	}

	return ResponseCodes{
		code:        int16(code),
		description: text,
	}
}
