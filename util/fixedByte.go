package util

import (
	"errors"
)

type FixedWriter struct {
    buf    []byte
    offset int
}

func NewFixedWriter(size int) *FixedWriter {
    return &FixedWriter{
        buf: make([]byte, size),
    }
}

func (fw *FixedWriter) Write(p []byte) (n int, err error) {
    if fw.offset >= len(fw.buf) {
        return 0, errors.New("SizeExceed")
    }

    n = copy(fw.buf[fw.offset:], p)
    fw.offset += n

    if fw.offset >= len(fw.buf) {
        err = errors.New("SizeExceed")
    }

    return n, err
}

func (fw *FixedWriter) Buffer() []byte {
    return fw.buf
}