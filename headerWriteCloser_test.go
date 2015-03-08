package gutils

import (
	"os"
	"testing"
)

func TestHW(t *testing.T) {
	h := []byte("==========\nHELLO\n+++++++++++++")
	hw := NewHeaderWriteCloser(func() []byte { return h }, os.Stdout)
	hw.Write([]byte("Hello World!"))
	hw.Close()
}
