package helper

import (
	"fmt"
	"reflect"
	"strconv"
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
			actual := StringToBool(test.given)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%t', got '%t'", test.expected, actual)
			}
		})
	}
}

func TestStrConvAtoiWithDefault(t *testing.T) {
	tests := []struct {
		given      string
		defaultVal int
		expected   int
	}{
		{"123", 15, 123},
		{"65432", 15, 65432},
		{"-65432", 15, -65432},
		{"definitelyNotANumber", 25, 25},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := AtoiWithDefault(test.given, test.defaultVal)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%d', got '%d'", test.expected, actual)
			}
		})
	}
}

func FuzzTestStrConvAtoiWithDefault(f *testing.F) {
	f.Fuzz(func(t *testing.T, given string, defaultVal int) {
		expected, _ := strconv.Atoi(given)
		value := AtoiWithDefault(given, defaultVal)
		if value != expected && value != defaultVal {
			t.Fatalf("expected '%v' or '%v, got '%v'", expected, defaultVal, value)
		}
	})
}

func FuzzDefault(f *testing.F) {
	f.Fuzz(func(t *testing.T, originalVal int, defaultVal int) {
		value := Default(originalVal, defaultVal)
		if value != originalVal && value != defaultVal {
			t.Fatalf("expected '%v' or '%v, got '%v'", originalVal, defaultVal, value)
		}
	})
}
