package util

import (
	"errors"
)

type FixedWriter struct {
	buf     []byte
	maxSize int
}

func NewFixedWriter(size int) *FixedWriter {
	return &FixedWriter{
		buf:     make([]byte, 0, size),
		maxSize: size,
	}
}

func (fw *FixedWriter) Write(p []byte) (n int, err error) {
	if len(fw.buf)+len(p) > fw.maxSize {
		return 0, errors.New("SizeExceed")
	}

	if len(fw.buf)+len(p) <= fw.maxSize {
		n = len(p)
	} else {
		n = fw.maxSize - len(fw.buf)
	}

	fw.buf = append(fw.buf, p[:n]...)

	return n, nil
}

func (fw *FixedWriter) Buffer() []byte {
	return fw.buf
}
