package data

import (
	"encoding/json"
	"fmt"
	"github.com/sphireinc/mantis/helper"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestExists(t *testing.T) {
	// Directory
	dir, err := ioutil.TempDir("", "tmp")
	if err != nil {
		t.Fatalf(err.Error())
	}
	defer func() { _ = os.RemoveAll(dir) }()

	check, err := Exists(dir, Directory)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if check == false {
		t.Fatalf("expected: true, got %t", check)
	}

	// File
	file, err := ioutil.TempFile("", "tmpfile")
	if err != nil {
		t.Fatalf(err.Error())
	}
	helper.DeferFileClose(file)

	_, err = Exists(file.Name(), File)
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Path
	_, err = Exists(dir, Path)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestIsStringTrue(t *testing.T) {
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

func TestQueryJson(t *testing.T) {
	// Test error paths first
	_, err := QueryJSON("", "a")
	assert.NotNil(t, err)

	_, err = QueryJSON("{}", "a")
	assert.NotNil(t, err)

	tests := []struct {
		json     string
		key      string
		expected float64
	}{
		{`{"a": 1, "b": 2}`, "b", 2},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			result, err := QueryJSON(test.json, test.key)
			assert.Nil(t, err)
			if !reflect.DeepEqual(result, test.expected) {
				t.Fatalf("expected '%f', got '%f'", test.expected, result)
			}
		})
	}
}

func TestMapHasKey(t *testing.T) {
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

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%t', got '%t'", test.expected, test.actual)
			}
		})
	}
}

func TestGetEnvVariables(t *testing.T) {
	assert.Greater(t, len(GetEnvVariables()), 0)
}

func TestUnmarshalFile(t *testing.T) {
	sampleData := make(map[string]string)
	sampleData["foo"] = "bar"
	sampleData["bar"] = "baz"

	yamlData, _ := yaml.Marshal(sampleData)
	jsonData, _ := json.Marshal(sampleData)

	// Create our temp files
	yamlFile, err := ioutil.TempFile("", "tmpfile.*.yaml")
	if err != nil {
		t.Fatalf(err.Error())
	}
	jsonFile, err := ioutil.TempFile("", "tmpfile.*.json")
	if err != nil {
		t.Fatalf(err.Error())
	}
	bsonFile, err := ioutil.TempFile("", "tmpfile.*.bson")
	if err != nil {
		t.Fatalf(err.Error())
	}

	// Defer close files
	defer helper.DeferFileClose(yamlFile)
	defer helper.DeferFileClose(jsonFile)
	defer helper.DeferFileClose(bsonFile)

	// Write our data to our files
	_, err = yamlFile.Write(yamlData)
	if err != nil {
		t.Fatalf(err.Error())
	}
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		t.Fatalf(err.Error())
	}
	_, err = bsonFile.Write(yamlData)
	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = UnmarshalFile(yamlFile.Name())
	assert.Nil(t, err)
	_, err = UnmarshalFile(jsonFile.Name())
	assert.Nil(t, err)
	_, err = UnmarshalFile(bsonFile.Name())
	assert.NotNil(t, err)
	_, err = UnmarshalFile("blahblahblah")
	assert.NotNil(t, err)
}

func TestMapToString(t *testing.T) {
	testOne := make(map[string]any)
	testOne["foo"] = "bar"
	testOne["bar"] = "baz"
	testOne["baz"] = 1
	testOne["bay"] = 4.0
	x, _ := MapToString(testOne)
	assert.Equal(t, `{"bar":"baz","bay":4,"baz":1,"foo":"bar"}`, x)
}
