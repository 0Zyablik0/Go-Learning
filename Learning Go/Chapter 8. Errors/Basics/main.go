package ma

import (
	"errors"
)

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("zero division")
	}

	return numerator / denominator, numerator % denominator, nil
}
