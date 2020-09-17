package data

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestIsTrue(t *testing.T) {
	tests := []struct {
		actual   bool
		expected bool
	}{
		{IsTrue("true"), true},
		{IsTrue("TRUE"), true},
		{IsTrue("1"), true},
		{IsTrue("false"), false},
		{IsTrue("0"), false},
		{IsTrue("2"), false},
		{IsTrue("d7#$"), false},
		{IsTrue("trU3"), false},
		{IsTrue("FALSE"), false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%t', got '%t'", test.expected, test.actual)
			}
		})
	}
}

func TestJsonQuery(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
	}{
		{"t", "t"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}

func TestExists(t *testing.T) {
	dir, err := ioutil.TempDir("", "tmp")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer func() {
		_ = os.RemoveAll(dir)
	}()

	check, err := DirectoryExists(dir)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if check == false {
		t.Fatalf("expected: true, got %t", check)
	}
}

func TestContains(t *testing.T) {
	testData := make(map[string]string)
	testData["k1"] = "v1"
	testData["k2"] = "v2"
	testData["k3"] = "v3"

	tests := []struct {
		actual   bool
		expected bool
	}{
		{MapStringStringContains(testData, "k1"), true},
		{MapStringStringContains(testData, "k2"), true},
		{MapStringStringContains(testData, "k3"), true},
		{MapStringStringContains(testData, "k4"), false},
		{MapStringStringContains(testData, ""), false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%t', got '%t'", test.expected, test.actual)
			}
		})
	}
}
