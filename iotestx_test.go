package iotestx

import (
	"testing"

	"github.com/aphistic/sweet"
	. "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	RegisterFailHandler(sweet.GomegaFail)

	sweet.T(func(s *sweet.S) {
		s.RunSuite(t, &PartialSuite{})
		s.RunSuite(t, &SequencedPartialSuite{})
		s.RunSuite(t, &FailedSuite{})
	})
}
