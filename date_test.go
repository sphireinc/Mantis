package mantis

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStringToDate(t *testing.T) {
	t1 := Date{2019, 1, 1, 0, 0, 0, 0, 0, 1, 1}
	t2 := Date{2019, 4, 3, 14, 45, 22, 0, 0, 1, 1}
	t3 := Date{1970, 1, 1, 00, 0, 0, 0, 0, 1, 1}
	t4 := Date{1870, 1, 1, 23, 59, 59, 0, 0, 1, 1}

	tests := []struct {
		date     string
		expected Date
	}{
		{"2019-01-01 00:00:00:00", t1},
		{"2019-04-03 14:45:22", t2},
		{"1970-01-01 00:00:00", t3},
		{"1870-01-01 23:59:59:17", t4},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual := StringToDate(test.date)
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%d', got '%d'", test.expected, actual)
			}
		})
	}
}

func TestDateToString(t *testing.T) {
	tests := []struct {
		expected string
		date     Date
	}{
		{"2019-01-01 00:00:00:00", Date{2019, 1, 1, 0, 0, 0, 0, 0, 1, 1}},
		{"2019-04-03 14:45:22", Date{2019, 4, 3, 14, 45, 22, 0, 0, 1, 1}},
		{"1970-01-01 00:00:00", Date{1970, 1, 1, 00, 0, 0, 0, 0, 1, 1}},
		{"1870-01-01 23:59:59:17", Date{1870, 1, 1, 23, 59, 59, 0, 0, 1, 1}},
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
