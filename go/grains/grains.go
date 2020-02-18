// Package grains calculates the grains on a chessboard
// based on the amount on the first square with the rule that
// it doubles on each successive square.
package grains

import (
	"errors"
)

// Square takes the index of the chessboard square and
// returns the number of grains of rine on it.
func Square(in int) (uint64, error) {
	switch {
	case in > 64 || in < 1:
		return 0, errors.New("invalid input")
	default:
		return 1 << (in - 1), nil
	}
}

// Total returns the total grains of rice on the board.
func Total() uint64 {
	return 1<<64 - 1
}
