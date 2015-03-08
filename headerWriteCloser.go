package gutils

import (
	"io"
)

type HeaderWriteCloser struct {
	header func() []byte
	w      io.WriteCloser
}

func NewHeaderWriteCloser(h func() []byte, w io.WriteCloser) *HeaderWriteCloser {
	return &HeaderWriteCloser{h, w}
}

func (h *HeaderWriteCloser) Write(p []byte) (int, error) {
	n, err := h.w.Write(h.header())
	if err != nil {
		return n, err
	}
	m, err := h.w.Write(p)
	return n + m, err
}

func (h *HeaderWriteCloser) Close() error {
	return h.w.Close()
}
