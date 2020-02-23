// Package prime implioemtns a utlity to caluclate the nth
// prime number of a given integer.
package prime

// Nth utilized a Sieve of Eratosthenes to calculate the
// the nth prime of a given integer.
func Nth(in int) (int, bool) {
	var sieve = make(map[int]int)
	var out int
	for i := 2; i <= in; i++ {
		sieve[i]=i
	}
			for i:=2;i<
	return out, true
}
