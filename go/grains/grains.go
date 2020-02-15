// Package grains calculates the grains on a chessboard
// based on the amount on the first square with the rule that
// it doubles on each successive square.
package grains

import (
	"errors"
	"math"
)

// Square takes the index of the chessboard square and
// returns the number of grains of rine on it.
func Square(in int) (uint64, error) {
	switch {
	case in == 1:
		return 1, nil
	case in > 64 || in < 1:
		return 0, errors.New("invalid input")
	default:
		return uint64(math.Pow(2, float64(in-1))), nil
	}
}

// Total returns the total grains of rice on the board.
func Total() uint64 {
	var c uint64
	c = (uint64(math.Pow(2, 64) - float64(1<<10)))

	return uint64(c)
}
