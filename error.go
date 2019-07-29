package mantis

import (
	"fmt"
)

var logger Log

// SetErrorLog Sets the logger interface to utilize to log
func SetErrorLog(logr Log) {
	logger = logr
}

// HandleError Handles an error with an error message
func HandleError(message string, err error) {
	if err != nil {
		logger.Write(fmt.Sprintf("%s => %s", message, err.Error()))
	}
}

// HandleFatalError Wrapper to panic
func HandleFatalError(err error) {
	if err != nil {
		logger.Write(fmt.Sprintf("Fatal panic => %s", err.Error()))
		panic(err)
	}
}