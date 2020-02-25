package sieve

import "sort"

func Sieve(in int) []int {
	temp := make([]bool, 0)
	out := make([]int, 0)
	for i := 2; i <= in; i++ {
		temp[i] = true
	}
	for key := range temp {
		for i := 2 * key; i <= in; i += key {
			if _, ok := temp[i]; ok {
				delete(temp, i)
			}
		}
	}
	for key := range temp {
		out = append(out, key)
	}
	sort.Ints(out)
	return out
}

/* Package sieve implements the Sieve of Eratosthenes.
package sieve

// Sieve gives a slice of primes up to and including n.
func Sieve(n int) []int {
	sieve := make([]bool, n-1)
	var primes []int
	for i := range sieve {
		if !sieve[i] {
			p := i + 2
			primes = append(primes, p)
			for j := p * 2; j <= n; j += p {
				sieve[j-2] = true
			}
		}
	}
	return primes
}