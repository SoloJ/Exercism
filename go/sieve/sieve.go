// Package sieve impliments a small utlity that
// performs the Sieve of Eratosthenes.
package sieve

// Sieve takes in an integer and uses the Sieve of
// Eratosthenes to return all prime numbers up to it.
func Sieve(in int) []int {
	temp := make([]bool, in+1)
	out := make([]int, 0)
	for i := 2; i < in+1; i++ {
		if !temp[i] {
			out = append(out, i)
		}
		for p := i; p < in+1; p += i {
			temp[p] = true
		}
	}
	return out
}
