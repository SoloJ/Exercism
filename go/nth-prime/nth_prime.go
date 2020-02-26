// Package prime implioemtns a utlity to caluclate the nth
// prime number of a given integer.
package prime

// Nth utilized a Sieve of Eratosthenes to calculate the
// the nth prime of a given integer.
func Nth(in int) (int, bool)){
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