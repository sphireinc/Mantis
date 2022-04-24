package log

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	INFO = iota
	WARN
	ERROR
	FATAL
	PANIC
)

// Log is our primary log struct
type Log struct {
	Logger      *log.Logger `json:"logger,omitempty"`
	Filename    string      `json:"filename,omitempty"`
	Status      bool        `json:"status,omitempty"`
	PrintToTerm bool        `json:"print_to_term,omitempty"`
	Overwrite   bool        `json:"overwrite,omitempty"`
	MinLogLevel int         `json:"minLogLevel,omitempty"`
}

// String converts our Log struct into a JSON string
func (l *Log) String() string {
	marshaledStruct, err := json.Marshal(l)
	if err != nil {
		return err.Error()
	}
	return string(marshaledStruct)
}

// New creates a new Log instance given filename
func New(filename string, printToTerm bool, overwrite bool) (*Log, error) {
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
		PrintToTerm: printToTerm,
		Overwrite:   overwrite,
		MinLogLevel: INFO,
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_APPEND
	if L.Overwrite {
		flags = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	}

	logFile, err := os.OpenFile(L.Filename, flags, 0755)

	if err != nil {
		return nil, err
	}

	L.Logger = log.New(logFile, "", log.LstdFlags|log.Lmicroseconds|log.LUTC)
	L.Status = true
	L.Write("Mantis.Log successfully initiated: " + filename)

	return &L, nil
}

func (l *Log) SetLogLevel(level int) {
	if level < INFO || level > PANIC {
		level = INFO
	}
	l.MinLogLevel = level
}

// Write a message to log and prepend time
func (l *Log) Write(msg string) {
	if l.Status {
		logMessage := fmt.Sprintf("%s", msg)
		if l.PrintToTerm {
			fmt.Println(logMessage)
		}
		l.Logger.Println(logMessage)
	}
}

// writer a message to log and prepend time
func (l *Log) writer(logLevel int, msg string) {
	logMessage := func(logLevelStr string, msg string) string {
		return fmt.Sprintf("%s %s", logLevelStr, msg)
	}

	fmt.Println(msg, logLevel, l.MinLogLevel)
	if logLevel < l.MinLogLevel {
		return
	}

	if l.Status {
		switch logLevel {
		case INFO:
			l.Logger.Println(logMessage("INFO", msg))
		case WARN:
			l.Logger.Println(logMessage("WARN", msg))
		case ERROR:
			l.Logger.Println(logMessage("ERROR", msg))
		case FATAL:
			l.Logger.Fatalln(logMessage("FATAL", msg))
		case PANIC:
			l.Logger.Panicln(logMessage("PANIC", msg))
		}
	}
}

// LogHTTPRequest logs an HTTP request from writer with a given name
func (l *Log) LogHTTPRequest(name string, w http.ResponseWriter, r *http.Request) {
	if l.Status {
		requestID := w.Header().Get("X-Request-Id")
		l.Write(fmt.Sprintf("%s %s %s %s %s %s", name, r.Method, r.URL.EscapedPath(), requestID, r.RemoteAddr, r.UserAgent()))
	}
}
