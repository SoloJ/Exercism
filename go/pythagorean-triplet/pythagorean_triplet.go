// Package pythagorean impliments a simple utlity
// to generate pythogarean triplets based of two methods.
package pythagorean

// Triplet is the return type that the tests expect
type Triplet [3]int

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	out := make([]Triplet, 0)
	for a := min; a < (max - 1); a++ {
		for b := a + 1; b < max; b++ {
			for c := b + 1; c <= max; c++ {
				if (a*a+b*b) == c*c && a < b && b < c {
					out = append(out, Triplet{a, b, c})
				}
			}
		}
	}
	return out
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
func Sum(p int) []Triplet {
	out := make([]Triplet, 0)
	for a := 1; a < (p - 1); a++ {
		for b := a + 1; b < p; b++ {
			for c := b + 1; c <= p; c++ {
				if (a*a+b*b) == c*c && a < b && b < c && a+b+c == p {
					out = append(out, Triplet{a, b, c})
				}
			}
		}
	}
	return out
}
