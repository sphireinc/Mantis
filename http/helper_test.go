package http

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHTTPErrorCodeMessage(t *testing.T) {
	tests := []struct {
		responseCode int
		expected     ResponseCodes
	}{
		{200, ResponseCodes{
			code:        200,
			description: "OK",
		}},
		{700, ResponseCodes{
			code:        001,
			description: "Unknown",
		}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			responseCode := GetHTTPResponseCode(test.responseCode)
			assert.Equal(t, responseCode.String(), test.expected.String())
		})
	}
}
