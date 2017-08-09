package iotestx

import "errors"

var (
	ErrSequenceTooShort = errors.New("Sequence is too short for the number of writes")
)
