package mantis

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToDate(t *testing.T) {
	tests := []struct {
		date     string
		expected Date
	}{
		{"2019-01-01 00:00:00:00", Date{"2019", "01", "01", "00", "00", "00"}},
		{"2019-04-03 14:45:22", Date{"2019", "04", "03", "14", "45", "22"}},
		{"1970-01-01 00:00:00", Date{"1970", "01", "01", "00", "00", "00"}},
		{"1870-01-01 23:59:59:17", Date{"1870", "01", "01", "23", "59", "59"}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			act := StringToDate(test.date)
			if !reflect.DeepEqual(act, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, act)
			}
		})
	}
}

func TestDateToString(t *testing.T) {
	tests := []struct {
		expected string
		date     Date
	}{
		{"2019-01-01 00:00:00", Date{"2019", "01", "01", "00", "00", "00"}},
		{"2019-04-03 14:45:22", Date{"2019", "04", "03", "14", "45", "22"}},
		{"1970-01-01 00:00:00", Date{"1970", "01", "01", "00", "00", "00"}},
		{"1870-01-01 23:59:59", Date{"1870", "01", "01", "23", "59", "59"}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			act := test.date.DateToString()
			if !reflect.DeepEqual(act, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, act)
			}
		})
	}
}
