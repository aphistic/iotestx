package iotestx

import (
	"bytes"
)

type PartialWriter struct {
	Buffer *bytes.Buffer
	length int
}

func NewPartialWriter(length int) *PartialWriter {
	return &PartialWriter{
		Buffer: bytes.NewBuffer(nil),
		length: length,
	}
}

func (w *PartialWriter) Write(p []byte) (int, error) {
	if len(p) > w.length {
		n, err := w.Buffer.Write(p[0:w.length])
		return n, err
	}

	n, err := w.Buffer.Write(p)
	return n, err
}

type SequencedPartialWriter struct {
	Buffer   *bytes.Buffer
	sequence []int

	writes int
}

func NewSequencePartialWriter(sequence []int) *SequencedPartialWriter {
	if sequence == nil {
		sequence = make([]int, 0)
	}

	return &SequencedPartialWriter{
		Buffer:   bytes.NewBuffer(nil),
		sequence: sequence,
		writes:   0,
	}
}

func (w *SequencedPartialWriter) Write(p []byte) (int, error) {
	if w.writes >= len(w.sequence) {
		return 0, ErrSequenceTooShort
	}

	if len(p) > w.sequence[w.writes] {
		n, err := w.Buffer.Write(p[0:w.sequence[w.writes]])
		w.writes++
		return n, err
	}

	n, err := w.Buffer.Write(p)
	w.writes++
	return n, err
}
