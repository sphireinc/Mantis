package log

import (
	"encoding/json"
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

func (l *Log) String() string {
	marshaledStruct, err := json.Marshal(l)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
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
func (l *Log) Write(msg string) {
	if l.Status {
		logMessage := fmt.Sprintf(" %s", msg)
		if l.PrintToTerm {
			fmt.Println(logMessage)
		}
		l.Logger.Println(logMessage)
	}
}

// LogHTTPRequest logs an HTTP request from writer with a given name
func (l *Log) LogHTTPRequest(name string, w http.ResponseWriter, r *http.Request) {
	if l.Status {
		requestID := w.Header().Get("X-Request-Id")
		l.Write(fmt.Sprintf("%s %s %s %s %s %s", name, r.Method, r.URL.EscapedPath(), requestID, r.RemoteAddr, r.UserAgent()))
	}
}
