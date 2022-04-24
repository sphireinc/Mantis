package log

import (
	"fmt"
	"os"
	"testing"
)

func TestLog(t *testing.T) {
	lg, err := New(os.TempDir()+string(os.PathSeparator)+"sph_test.log", true, false)
	if err != nil {
		fmt.Println(err)
	}

	lg.Write("HELLO")
	lg.Write("HELLO")
	lg.Write("HELLO")
	lg.Write("HELLO")
}
