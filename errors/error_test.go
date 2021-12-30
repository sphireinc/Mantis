package errors

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	start := time.Now()
	err := New(100, "some message", nil)
	end := time.Now()
	if err.Code() != 100 {
		t.Fatalf("Code fail, test 1")
	}
	if err.Message() != "some message" {
		t.Fatalf("Message  fail, test 1")
	}
	if err.Time().After(end) || err.Time().Before(start) {
		t.Fatalf(fmt.Sprintf("Time was %v, after %v and before %v", err.Time(), start, end))
	}

	start = time.Now()
	err = New(100, "%v is %v years old", []any{"Kim", 22})
	end = time.Now()
	if err.Code() != 100 {
		t.Fatalf("Code fail, test 1")
	}
	if err.Message() != "Kim is 22 years old" {
		t.Fatalf("Message  fail, test 1")
	}
	if err.Time().After(end) || err.Time().Before(start) {
		t.Fatalf(fmt.Sprintf("Time was %v, after %v and before %v", err.Time(), start, end))
	}
}
