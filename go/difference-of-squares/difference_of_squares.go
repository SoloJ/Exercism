// Package diffsquares impliments a simple function
// that calculates difference of the square of sums and
// the sum of squares.
package diffsquares

// SquareOfSum finds the square of sums.
func SquareOfSum(in int) int {
	return (((in * (in + 1)) / 2) * ((in * (in + 1)) / 2))
}

// SumOfSquares finds the sum of squares.
func SumOfSquares(in int) int {
	return (in * (in + 1) * (2*in + 1) / 6)
}

// Difference is difference between squares of sums and sum of squares.
func Difference(in int) int {
	return (SquareOfSum(in) - SumOfSquares(in))
}
