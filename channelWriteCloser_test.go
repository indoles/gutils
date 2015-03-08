package gutils

import (
	"os"
	"testing"
)

func TestCW(t *testing.T) {
	cw := NewChannelWriteCloser(os.Stdout)
	cw.Write([]byte("Hello Channel!\n"))
	cw.Close()
}
