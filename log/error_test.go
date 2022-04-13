package log

import (
	"errors"
	"fmt"
	"testing"

	validation "github.com/go-ozzo/ozzo-validation"
)

func TestHandleError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			L := Log{}
			L.HandleError("Error test", errors.New("something went wrong"))
		}
	}()
}

func TestHandleFatalError(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			L := Log{}
			L.HandleFatalError(errors.New("something went wrong"))
		}
	}()
}

func TestJSONMarshalError(t *testing.T) {
	type Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
		Zip    string `json:"zip"`
	}
	a := Address{
		Street: "",
		City:   "",
		State:  "",
		Zip:    "",
	}
	err := validation.ValidateStruct(&a,
		validation.Field(&a.City, validation.Required),
	)

	tests := []struct {
		err      error
		expected string
	}{
		{errors.New("not an empty test"), `{"error":"not an empty test"}`},
		{errors.New(""), `{"error":""}`},
		{errors.New("{invalid:json}"), `{"error":"{invalid:json}"}`},
		{errors.New("{'valid':'json?'}"), `{"error":"{'valid':'json?'}"}`},
		{err, `{"error":"city: cannot be blank."}`},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := JSONMarshalError(test.err)
			if actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}

func TestJSONMarshalAndLogError(t *testing.T) {
	type Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
		Zip    string `json:"zip"`
	}
	a := Address{
		Street: "",
		City:   "",
		State:  "",
		Zip:    "",
	}
	err := validation.ValidateStruct(&a,
		validation.Field(&a.City, validation.Required),
	)

	tests := []struct {
		err      error
		expected string
	}{
		{errors.New("not an empty test"), `{"error":"not an empty test"}`},
		{errors.New(""), `{"error":""}`},
		{errors.New("{invalid:json}"), `{"error":"{invalid:json}"}`},
		{errors.New("{'valid':'json?'}"), `{"error":"{'valid':'json?'}"}`},
		{err, `{"error":"city: cannot be blank."}`},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			L := Log{}
			actual := L.JSONMarshalAndLogError(test.expected, test.err)
			if actual != test.expected {
				t.Fatalf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}
