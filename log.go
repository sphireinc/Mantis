package mantis

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Log struct {
	Logger   *log.Logger
	Filename string
}

func (L *Log) NewLog(filename string) {
	if len(filename) == 0 {
		appName := os.Getenv("APP_NAME") + ".log"
		if len(appName) == 0 {
			appName = "go_lang_app.log"
		}
		ex, err := os.Executable()
		HandleError("NewLog get Executable", err)
		filename = filepath.Dir(ex) + appName
	}

	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
	HandleFatalError(err)
	L.Logger = log.New(logFile, "", log.LstdFlags)
	L.Write("Log successfully initiated")
	L.Filename = filename
}

func (L *Log) Write(msg string) {
	date := CurrentTime()
	L.Logger.Println(fmt.Sprintf("%s # Log := %s", date.DateToString(), msg))
}
