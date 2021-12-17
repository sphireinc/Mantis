package errors

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Errors struct {
	code      int32
	message   string
	timestamp time.Time
}

func (e *Errors) Code() int32 {
	return e.code
}

func (e *Errors) Message() string {
	return e.message
}

func (e *Errors) Time() string {
	return e.timestamp.String()
}

// New creates a new errors instance
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
	return errors.New(fmt.Sprintf("%v (error %v) %v", e.Time(), e.Code(), e.Message()))
}
