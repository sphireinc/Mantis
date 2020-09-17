package helper

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		given    string
		expected string
	}{
		{"hello", "olleh"},
		{"654321", "123456"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := Reverse(test.given)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}