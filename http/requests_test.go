package http

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse(T *testing.T) {
	response := Response{
		Body:       []byte(`{"custom": "Content"}`),
		BodyString: "some string in the body",
		Error:      errors.New("some error"),
	}
	assert.Len(T, response.Byte(), 90)

	response = Response{
		Body:       []byte(`{"custom": "Content"}`),
		BodyString: "some string in the body",
	}
	assert.Len(T, response.Byte(), 80)

	response = Response{
		Body: []byte(`{"custom": "Content"}`),
	}
	assert.Len(T, response.Byte(), 57)

	response = Response{}
	assert.Len(T, response.Byte(), 41)
}
