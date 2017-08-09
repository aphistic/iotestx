package iotestx

import (
	"errors"
	"testing"

	. "github.com/onsi/gomega"
)

var (
	errFailedSuite = errors.New("Failed suite test error")
)

type FailedSuite struct{}

func (s *FailedSuite) TestInitialWriteFailure(t *testing.T) {
	w := NewFailedWriter(10, 0, errFailedSuite)
	n, err := w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(0))
	Expect(err).To(Equal(errFailedSuite))
	Expect(w.Buffer.Bytes()).To(HaveLen(0))
}

func (s *FailedSuite) TestShortWriteFailure(t *testing.T) {
	w := NewFailedWriter(5, 1, errFailedSuite)
	n, err := w.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
	n, err = w.Write([]byte{5, 6, 7, 8})
	Expect(n).To(Equal(0))
	Expect(err).To(Equal(errFailedSuite))
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
}

func (s *FailedSuite) TestShortWrite(t *testing.T) {
	w := NewFailedWriter(10, 10, nil)
	n, err := w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
}

func (s *FailedSuite) TestExactWrite(t *testing.T) {
	w := NewFailedWriter(5, 10, nil)
	n, err := w.Write([]byte{0, 1, 2, 3, 4})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
}

func (s *FailedSuite) TestLongWrite(t *testing.T) {
	w := NewFailedWriter(5, 10, nil)
	n, err := w.Write([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8})
	Expect(n).To(Equal(5))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4}))
	n, err = w.Write([]byte{5, 6, 7, 8})
	Expect(n).To(Equal(4))
	Expect(err).To(BeNil())
	Expect(w.Buffer.Bytes()).To(Equal([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8}))
}
