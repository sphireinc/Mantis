package http

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

type ReaderStruct struct {
	A string
	B string
}

func getRequestHelper(target string, readerString string, contentType string) *http.Request {
	var r *http.Request
	r = httptest.NewRequest(http.MethodPost, target, strings.NewReader(readerString))
	r.Header.Set("Content-Type", contentType)
	func(w http.ResponseWriter, r *http.Request) { _ = r.ParseForm() }(httptest.NewRecorder(), r)
	return r
}

func TestParseBodyIntoStruct(t *testing.T) {
	val := getRequestHelper("/some/test", `{"a": "123", "b": "456"}`, "application/json")
	rSVal := ReaderStruct{}

	if err := ParseBodyIntoStruct(val, &rSVal); err != nil {
		t.Errorf("ParseBodyIntoStruct Failed while reading: " + err.Error())
	}

	if rSVal.A != "123" {
		t.Fatalf("expected '%s', got '%s'", rSVal.A, "123")
	}

	if rSVal.B != "456" {
		t.Fatalf("expected '%s', got '%s'", rSVal.B, "456")
	}
}

func TestGetBody(t *testing.T) {
	val := getRequestHelper("/some/test", `{"a": "123", "b": "456"}`, "application/json")
	body, err := GetBody(val)

	if err != nil {
		t.Fatalf("Expected error to be nil, received error: %s", err.Error())
	}

	if bytes.Compare(body, []byte(`{"a": "123", "b": "456"}`)) == 1 {
		t.Fatalf("expected '%s', got '%s'", string(body), `{"a": "123", "b": "456"}`)
	}
}

func TestGetQueryParameter(t *testing.T) {
	val := getRequestHelper("/?a=1&b=2", "a=1&b=2", "application/x-www-urlencoded-form")

	tests := []struct {
		actual   string
		expected string
	}{
		{"a", "1"},
		{"b", "2"},
		{"z", ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			test.actual, _ = GetQueryParameter(val, test.actual)
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}

func TestGetQueryParameters(t *testing.T) {
	val := getRequestHelper("/?a=1&b=2", "a=1&b=2", "application/x-www-urlencoded-form")

	t1 := url.Values{}
	t1.Set("a", "1")
	t1.Set("b", "2")

	tests := []struct {
		actual   url.Values
		expected url.Values
	}{
		{GetQueryParameters(val), t1},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}

func TestParseUrl(t *testing.T) {
	t1 := url.URL{
		Scheme:     "",
		Opaque:     "",
		User:       nil,
		Host:       "",
		Path:       "",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   "",
		Fragment:   "",
	}

	t2 := url.URL{
		Scheme:     "https",
		Opaque:     "",
		User:       nil,
		Host:       "google.com",
		Path:       "/maps",
		RawPath:    "",
		ForceQuery: false,
		RawQuery:   "v=1",
		Fragment:   "",
	}

	firstTest, err1 := ParseUrl("")
	secondTest, err2 := ParseUrl("https://google.com/maps?v=1")

	if err1 != nil {
		t.Fatalf("Expected error to be nil, received error: %s", err1.Error())
	}

	if err2 != nil {
		t.Fatalf("Expected error to be nil, received error: %s", err2.Error())
	}

	tests := []struct {
		actual   *url.URL
		expected *url.URL
	}{
		{firstTest, &t1},
		{secondTest, &t2},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}
