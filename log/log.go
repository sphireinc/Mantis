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
	Logger      *log.Logger `json:"logger,omitempty"`
	Filename    string      `json:"filename,omitempty"`
	Status      bool        `json:"status,omitempty"`
	PrintToTerm bool        `json:"print_to_term,omitempty"`
}

// New creates a new Log instance given filename
func New(filename string) (*Log, error) {
	// if filename is empty, create a local log where executable is
	if filename == "" {
		absPath, err := os.Executable()
		if err != nil {
			return nil, err
		}
		filename = filepath.Dir(absPath) + string(os.PathSeparator) + "app.log"
	}

	L := Log{
		Filename:    filename,
		PrintToTerm: false,
	}

	logFile, err := os.OpenFile(L.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return nil, err
	}

	L.Logger = log.New(logFile, "", log.LstdFlags)
	L.Logger.SetPrefix(time.Now().Format("2006-01-02 15:04:05"))
	L.Write("Log successfully initiated")
	L.Status = true

	return &L, nil
}

// Write a message to log and prepend time
func (L *Log) Write(msg string) {
	if L.Status {
		logMessage := fmt.Sprintf(" %s", msg)
		if L.PrintToTerm {
			fmt.Println(logMessage)
		}
		L.Logger.Println(logMessage)
	}
}

// LogHTTPRequest logs an HTTP request from writer with a given name
func (L *Log) LogHTTPRequest(name string, w http.ResponseWriter, r *http.Request) {
	if L.Status {
		requestID := w.Header().Get("X-Request-Id")
		L.Write(fmt.Sprintf("%s %s %s %s %s %s", name, r.Method, r.URL, requestID, r.RemoteAddr, r.UserAgent()))
	}
}
