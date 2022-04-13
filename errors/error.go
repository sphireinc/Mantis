package errors

import (
	"encoding/json"
	"fmt"
	"time"
)

// Errors holds our code, message, timestamp
type Errors struct {
	code      int32
	message   string
	timestamp time.Time
}

// Code returns the error code
func (e *Errors) Code() int32 {
	return e.code
}

// Message returns the error message
func (e *Errors) Message() string {
	return e.message
}

// Time returns the error timestamp
func (e *Errors) Time() time.Time {
	return e.timestamp
}

// New creates a new errors instance
// Usage:
//   err = New(100, "%v is %v years old", []any{"Kim", 22})
//  - or without args -
//   err := New(100, "some message", nil)
func New(code int32, message string, args []any) *Errors {
	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}
	return &Errors{
		code:      code,
		message:   message,
		timestamp: time.Now(),
	}
}

// MarshalJSON implements the JSON encoding interface
func (e *Errors) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"code":    e.code,
		"message": e.message,
		"time":    e.timestamp.String(),
	})
}

// Marshal creates a Go STD error from our Errors object
func (e *Errors) Marshal() error {
	return fmt.Errorf("%v (error %v) %v", e.Time(), e.Code(), e.Message())
}
