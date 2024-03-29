package log

import (
	"encoding/json"
	"fmt"
)

// HandleError handles an error with an error message.
func (l *Log) HandleError(message string, err error) {
	if err != nil {
		l.Write(fmt.Sprintf("%s => %s", message, err.Error()))
	}
}

// HandleFatalError is (nicer) wrapper to panic.
func (l *Log) HandleFatalError(err error) {
	if err != nil {
		l.Write(fmt.Sprintf("Fatal panic => %s", err.Error()))
		panic(err)
	}
}

// JSONMarshalAndLogError logs and error then JSON Marshals it
func (l *Log) JSONMarshalAndLogError(message string, err error) string {
	l.HandleError(message, err)
	return JSONMarshalError(err)
}

// JSONMarshalError takes an error and JSON Marshals it
func JSONMarshalError(err error) string {
	type E struct {
		Error string `json:"error"`
	}
	output, _ := json.Marshal(&E{Error: err.Error()})
	return string(output)
}
