package mantis

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Log struct {
	Logger   *log.Logger
	Filename string
}

// NewLog Setup our log
func (L *Log) NewLog(filename string) {
	L.Filename = filename
	if len(L.Filename) == 0 {
		ex, err := os.Executable()
		HandleError("NewLogGetExecutable", err)
		L.Filename = filepath.Dir(ex) + string(os.PathSeparator) + "app.log"
	}
	logFile, err := os.OpenFile(L.Filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	HandleFatalError(err)
	L.Logger = log.New(logFile, "", log.LstdFlags)
	L.Write("Log successfully initiated")
}

// Write Write a message to log and prepend time
func (L *Log) Write(msg string) {
	logMessage := fmt.Sprintf("%s", msg)
	fmt.Println(logMessage)
	L.Logger.Println(logMessage)
}

// LogHTTPRequest
func (L *Log) LogHTTPRequest(name string, w http.ResponseWriter, r *http.Request) {
	requestID := w.Header().Get("X-Request-Id")
	L.Write(fmt.Sprintf("%s %s %s %s %s %s", name, r.Method, r.URL, requestID, r.RemoteAddr, r.UserAgent()))
}
