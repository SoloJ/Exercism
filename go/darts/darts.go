package darts

import (
	"math"
)

// Score is the function that the test file expects to call
func Score(x, y float64) int {
	attemptArea := math.Pi * math.Pow(math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)), 2) // i used the math package but this may not be nessesary?

	switch {
	case attemptArea <= math.Pi*math.Pow(1, 2):
		return 10

	case attemptArea <= math.Pi*math.Pow(5, 2):
		return 5

	case attemptArea <= math.Pi*math.Pow(10, 2):
		return 1

	default:
		return 0
	}
}
