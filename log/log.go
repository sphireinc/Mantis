package log

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Log struct {
	logger      *log.Logger
	filename    string
	status      bool
	printToTerm bool
}

// NewLog Setup our log
func NewLog(filename string, printToTerm bool) (*Log, error) {
	// if filename is empty, we want to just create a local log wherever the executable is
	if filename == "" {
		absPath, err := os.Executable()
		if err != nil {
			return nil, err
		}
		filename = filepath.Dir(absPath) + string(os.PathSeparator) + "app.log"
	}

	L := Log{
		filename:    filename,
		printToTerm: printToTerm,
	}

	logFile, err := os.OpenFile(L.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	L.logger = log.New(logFile, "", log.LstdFlags)
	L.logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05"))
	L.Write("Log successfully initiated")
	L.status = true

	return &L, nil
}

// Write a message to log and prepend time
func (L *Log) Write(msg string) {
	if L.status {
		logMessage := fmt.Sprintf(" %s", msg)
		//fmt.Println(logMessage)
		L.logger.Println(logMessage)
	}
}

// LogHTTPRequest logs an HTTP request from writer with a given name
func (L *Log) LogHTTPRequest(name string, w http.ResponseWriter, r *http.Request) {
	if L.status {
		requestID := w.Header().Get("X-Request-Id")
		L.Write(fmt.Sprintf("%s %s %s %s %s %s", name, r.Method, r.URL, requestID, r.RemoteAddr, r.UserAgent()))
	}
}
