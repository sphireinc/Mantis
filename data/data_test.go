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
		actual   string
		expected bool
	}{
		{"true", true},
		{"TRUE", true},
		{"1", true},
		{"false", false},
		{"0", false},
		{"2", false},
		{"d7#$", false},
		{"trU3", false},
		{"FALSE", false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result := IsStringTrue(test.actual)
			if !reflect.DeepEqual(result, test.expected) {
				t.Fatalf("expected '%t', got '%t'", test.expected, result)
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

	check, err := Exists(dir, Directory)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if check == false {
		t.Fatalf("expected: true, got %t", check)
	}
}

func TestContains(t *testing.T) {
	testData := make(map[string]any)
	testData["k1"] = "v1"
	testData["k2"] = 5
	testData["k3"] = nil

	tests := []struct {
		actual   bool
		expected bool
	}{
		{MapHasKey(testData, "k1"), true},
		{MapHasKey(testData, "k2"), true},
		{MapHasKey(testData, "k3"), true},
		{MapHasKey(testData, "k4"), false},
		{MapHasKey(testData, ""), false},
	}

	testData2 := make(map[int]any)
	testData2[1] = "v1"
	testData2[2] = 3
	testData2[17] = "v3"

	tests2 := []struct {
		actual   bool
		expected bool
	}{
		{MapHasKey(testData2, 1), true},
		{MapHasKey(testData2, 2), true},
		{MapHasKey(testData2, 17), true},
		{MapHasKey(testData2, 3), false},
		{MapHasKey(testData2, -1), false},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%t', got '%t'", test.expected, test.actual)
			}
		})
	}

	for i, test := range tests2 {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%t', got '%t'", test.expected, test.actual)
			}
		})
	}
}
