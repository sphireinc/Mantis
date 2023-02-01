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
	assert.Len(T, response.Byte(), 81)

	response = Response{
		Body:       []byte(`{"custom": "Content"}`),
		BodyString: "some string in the body",
	}
	assert.Len(T, response.Byte(), 60)

	response = Response{
		Body: []byte(`{"custom": { "Content": {"x": 3}}}`),
	}
	assert.Len(T, response.Byte(), 30)

	response = Response{}
	assert.Len(T, response.Byte(), 4)
}
