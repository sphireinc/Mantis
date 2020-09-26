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

func TestStrConvParseBoolHideError(t *testing.T) {
	tests := []struct {
		given    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"anonboolean", false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := StrConvParseBoolHideError(test.given)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%t', got '%t'", test.expected, actual)
			}
		})
	}
}

func TestStrConvAtoiWithDefault(t *testing.T) {
	tests := []struct {
		given    string
		defaultGiven int
		expected int
	}{
		{"123",  15,123},
		{"65432", 15, 65432},
		{"-65432", 15, -65432},
		{"definitelyNotANumber", 25, 25},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := StrConvAtoiWithDefault(test.given, test.defaultGiven)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%d', got '%d'", test.expected, actual)
			}
		})
	}
}