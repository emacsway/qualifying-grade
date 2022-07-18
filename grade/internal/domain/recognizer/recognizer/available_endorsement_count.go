package recognizer

import (
	"errors"
	"fmt"
)

const yearlyEndorsementCount = uint8(20)

var (
	ErrInvalidAvailableEndorsementCount = errors.New(fmt.Sprintf(
		"endorsement count should be between 0 and %d", yearlyEndorsementCount,
	))
)

func NewAvailableEndorsementCount(value uint8) (AvailableEndorsementCount, error) {
	if value > yearlyEndorsementCount {
		return AvailableEndorsementCount(0), ErrInvalidAvailableEndorsementCount
	}
	return AvailableEndorsementCount(value), nil
}

type AvailableEndorsementCount uint8

func (c AvailableEndorsementCount) HasAvailable() bool {
	return uint8(c) > uint8(0)
}

func (c AvailableEndorsementCount) Decrease() (AvailableEndorsementCount, error) {
	n, err := NewAvailableEndorsementCount(uint8(c) - uint8(1))
	if err != nil {
		return AvailableEndorsementCount(0), err
	}
	return n, nil
}