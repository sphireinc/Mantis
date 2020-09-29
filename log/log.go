package log

import (
	"fmt"
	mantisError "github.com/sphireco/mantis/error"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Log struct {
	Logger   *log.Logger
	Filename string
	Status   bool
}

// NewLog Setup our log
func (L *Log) NewLog(filename string) {
	L.Filename = filename
	if L.Filename == "" {
		ex, err := os.Executable()
		mantisError.HandleError("NewLogGetExecutable", err)
		L.Filename = filepath.Dir(ex) + string(os.PathSeparator) + "app.log"
	}
	logFile, err := os.OpenFile(L.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		mantisError.HandleFatalError(err)
	}

	L.Logger = log.New(logFile, "", log.LstdFlags)
	L.Write("Log successfully initiated")
}

// Write a message to log and prepend time
func (L *Log) Write(msg string) {
	if L.Status {
		logMessage := fmt.Sprintf("%s", msg)
		fmt.Println(logMessage)
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
