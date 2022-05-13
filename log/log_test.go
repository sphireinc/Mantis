package log

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	lg, err := New(os.TempDir()+string(os.PathSeparator)+"sph_test.log", true, true)
	assert.Nil(t, err)

	lg.SetLogLevel(INFO)

	lg.Write("Write")
	lg.Info("INFO")
	lg.Warn("WARN")
	lg.Error("ERROR")
	//lg.Fatal("FATAL")
	//lg.Panic("PANIC")
}
