package iotestx

import (
	"bytes"
	"io"
)

type FailedWriter struct {
	Buffer *bytes.Buffer

	length int

	failOn   int
	failWith error

	writes int
}

var _ io.Writer = &FailedWriter{}

func NewFailedWriter(length int, failOn int, failWith error) *FailedWriter {
	return &FailedWriter{
		Buffer:   bytes.NewBuffer(nil),
		length:   length,
		failOn:   failOn,
		failWith: failWith,
		writes:   0,
	}
}

func (w *FailedWriter) Write(p []byte) (int, error) {
	if w.writes >= w.failOn {
		return 0, w.failWith
	}

	if len(p) > w.length {
		n, err := w.Buffer.Write(p[0:w.length])
		w.writes++
		return n, err
	}

	n, err := w.Buffer.Write(p)
	w.writes++
	return n, err
}

type FailedReader struct{}

var _ io.Reader = &FailedReader{}

func NewFailedReader() *FailedReader {
	return &FailedReader{}
}

func (r *FailedReader) Read(d []byte) (int, error) {
	return 0, io.ErrUnexpectedEOF
}
