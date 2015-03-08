package gutils

import (
	"io"
)

type reply struct {
	n   int
	err error
}

type ChannelWriteCloser struct {
	c chan []byte
	r chan reply
	w io.WriteCloser
}

func writer(w io.WriteCloser, c chan []byte, r chan reply) {
	for {
		b, more := <-c
		if more {
			n, err := w.Write(b)
			r <- reply{n, err}
		} else {
			err := w.Close()
			r <- reply{0, err}
			close(r)
			break
		}
	}
}

func NewChannelWriteCloser(w io.WriteCloser) *ChannelWriteCloser {
	cw := ChannelWriteCloser{make(chan []byte), make(chan reply), w}
	go writer(cw.w, cw.c, cw.r)
	return &cw
}

func (c *ChannelWriteCloser) Write(b []byte) (int, error) {
	c.c <- b
	r := <-c.r
	return r.n, r.err
}

func (c *ChannelWriteCloser) Close() error {
	close(c.c)
	r := <-c.r
	return r.err
}
