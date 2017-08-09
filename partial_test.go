package iotestx

import (
	"testing"

	. "github.com/onsi/gomega"
)

type PartialSuite struct{}

func (s *PartialSuite) TestShortWrite(t *testing.T) {
	w := NewPartialWriter(10)
	n, err := w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
}

func (s *PartialSuite) TestExactWrite(t *testing.T) {
	w := NewPartialWriter(5)
	n, err := w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
}

func (s *PartialSuite) TestLongWrite(t *testing.T) {
	w := NewPartialWriter(5)
	n, err := w.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
	n, err = w.Write([]byte{5, 6, 7, 8})
	Expect(n).To(Equal(4))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}

type SequencedPartialSuite struct{}

func (s *SequencedPartialSuite) TestSequenceTooShort(t *testing.T) {
	w := NewSequencePartialWriter(nil)
	n, err := w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(0))
	Expect(err).To(Equal(ErrSequenceTooShort))
	Expect(w.Buffer.Bytes()).To(HaveLen(0))

	w = NewSequencePartialWriter([]int{})
	n, err = w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(0))
	Expect(err).To(Equal(ErrSequenceTooShort))
	Expect(w.Buffer.Bytes()).To(HaveLen(0))
}

func (s *SequencedPartialSuite) TestShortWrite(t *testing.T) {
	w := NewSequencePartialWriter([]int{10})
	n, err := w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
}

func (s *SequencedPartialSuite) TestExactWrite(t *testing.T) {
	w := NewSequencePartialWriter([]int{5})
	n, err := w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
}

func (s *SequencedPartialSuite) TestLongWrite(t *testing.T) {
	w := NewSequencePartialWriter([]int{5, 2, 10})
	n, err := w.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
	n, err = w.Write([]byte{5, 6, 7, 8})
	Expect(n).To(Equal(2))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4, 5, 6}))
	n, err = w.Write([]byte{7, 8})
	Expect(n).To(Equal(2))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}
