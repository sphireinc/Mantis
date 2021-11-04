package http

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetHTTPErrorCodeMessage(t *testing.T) {
	tests := []struct {
		responseCode int16
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
			if !reflect.DeepEqual(responseCode, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected.String(), responseCode.String())
			}
		})
	}
}
