package iotestx

import (
	"bytes"
)

type FailedWriter struct {
	Buffer *bytes.Buffer

	length int

	failOn   int
	failWith error

	writes int
}

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
