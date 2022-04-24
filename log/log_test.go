package log

import (
	"fmt"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	lg, err := New(os.TempDir()+string(os.PathSeparator)+"sph_test.log", true, true)
	if err != nil {
		fmt.Println(err)
	}

	lg.SetLogLevel(INFO)

	lg.Write("Write")
	lg.Info("INFO")
	lg.Warn("WARN")
	lg.Error("ERROR")
	lg.Fatal("FATAL")
	lg.Panic("PANIC")
}
