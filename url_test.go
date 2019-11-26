package mantis

import (
	"bytes"
	"encoding/json"
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

var readerStruct = ReaderStruct{
	A: "123",
	B: "456",
}

var readerString, _ = json.Marshal(&readerStruct)

func getRequestHelper(useReaderString bool) *http.Request {
	w := httptest.NewRecorder()

	var r *http.Request
	if useReaderString {
		r = httptest.NewRequest(http.MethodPost, "/some/test", strings.NewReader(string(readerString)))
		r.Header.Set("Content-Type", "application/json")
	} else {
		r = httptest.NewRequest(http.MethodPost, "/?a=1&b=2", strings.NewReader("a=1&b=2"))
		r.Header.Set("Content-Type", "application/x-www-urlencoded-form")
	}

	func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
	}(w, r)
	return r
}

func TestParseBodyIntoStruct(t *testing.T) {
	val := getRequestHelper(true)
	rSVal := ReaderStruct{}

	if err := ParseBodyIntoStruct(val, &rSVal); err != nil {
		t.Errorf("ParseBodyIntoStruct Failed while reading: " + err.Error())
	}

	if rSVal.A != readerStruct.A {
		t.Fatalf("expected '%s', got '%s'", rSVal.A, readerStruct.A)
	}

	if rSVal.B != readerStruct.B {
		t.Fatalf("expected '%s', got '%s'", rSVal.B, readerStruct.B)
	}
}

func TestGetBody(t *testing.T) {
	val := getRequestHelper(true)
	body := GetBody(val)

	if bytes.Compare(body, readerString) == 1 {
		t.Fatalf("expected '%s', got '%s'", string(body), string(readerString))
	}
}

func TestGetQueryParameter(t *testing.T) {
	val := getRequestHelper(false)

	tests := []struct {
		actual   string
		expected string
	}{
		{GetQueryParameter(val, "a"), "1"},
		{GetQueryParameter(val, "b"), "2"},
		{GetQueryParameter(val, "z"), ""},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}

func TestGetQueryParameters(t *testing.T) {
	val := getRequestHelper(false)

	t1 := url.Values{}
	t1.Set("a", "1")
	t1.Set("b", "2")

	tests := []struct {
		actual   url.Values
		expected url.Values
	}{
		{GetQueryParameters(val), t1 },
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

	tests := []struct {
		actual   *url.URL
		expected *url.URL
	}{
		{ParseUrl(""), &t1},
		{ParseUrl("https://google.com/maps?v=1"), &t2},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if !reflect.DeepEqual(test.actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, test.actual)
			}
		})
	}
}