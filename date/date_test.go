package date

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestStringToDate(t *testing.T) {
	t1 := Date{2019, time.January, 1, 0, 0, 0, 0, 1546300800, time.Tuesday, 1}
	t2 := Date{2019, time.April, 3, 14, 45, 22, 0, 1554302722, time.Wednesday, 93}
	t3 := Date{1970, time.January, 1, 00, 0, 0, 0, 0, time.Thursday, 1}
	t4 := Date{1870, time.January, 1, 23, 59, 59, 0, -3155587201, time.Saturday, 1}

	tests := []struct {
		date     string
		expected Date
	}{
		{"2019-01-01 00:00:00", t1},
		{"2019-04-03 14:45:22", t2},
		{"1970-01-01 00:00:00", t3},
		{"1870-01-01 23:59:59", t4},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, _ := StringToDate(test.date)
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
		{"2019-01-01 00:00:00", Date{2019, 1, 1, 0, 0, 0, 0, 0, 1, 1}},
		{"2019-04-03 14:45:22", Date{2019, 4, 3, 14, 45, 22, 0, 0, 1, 1}},
		{"1970-01-01 00:00:00", Date{1970, 1, 1, 00, 0, 0, 0, 0, 1, 1}},
		{"1870-01-01 23:59:59", Date{1870, 1, 1, 23, 59, 59, 0, 0, 1, 1}},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, _ := test.date.DateToString("")
			if !reflect.DeepEqual(actual, test.expected) {
				t.Fatalf("expected '%s', got '%s'", test.expected, actual)
			}
		})
	}
}
